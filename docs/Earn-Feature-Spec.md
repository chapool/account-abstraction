# Chapool Earn — Phase 1 + Phase 2 功能开发说明

> **文档类型**：开发交付规范  
> **覆盖范围**：Phase 1（Chapool Earn MVP）+ Phase 2（veCPOT + NFT Boost）  
> **参考原型**：`frontend-prototype/src/App.tsx`  
> **基础产品设计**：`docs/Chapool-TVL-Product-Design.md`  
> **日期**：2026-03-11  

---

## 目录

1. [整体架构](#1-整体架构)
2. [模块一：ChapoolEarnVault（活期金库）](#2-模块一chapoolEarnVault)
3. [模块二：VeCPOTLocker（CPOT 锁仓）](#3-模块二vecpotlocker)
4. [模块三：NFTBoostController（NFT 加速）](#4-模块三nftboostcontroller)
5. [模块四：ChapoolRewardDistributor（收益注入）](#5-模块四chapoolrewarddistributor)
6. [模块五：ChapoolVaultReader（聚合只读）](#6-模块五chapoolvaultreader)
7. [APY 计算规则](#7-apy-计算规则)
8. [链上事件规范](#8-链上事件规范)
9. [前端页面与交互规范](#9-前端页面与交互规范)
10. [前端数据接入接口](#10-前端数据接入接口)
11. [错误处理与边界条件](#11-错误处理与边界条件)
12. [安全与权限边界](#12-安全与权限边界)

---

## 1 整体架构

### 1.1 模块关系

```
用户钱包
  │
  ├──[存 USDT]──────────▶ ChapoolEarnVault
  │                            │  │
  ├──[锁 CPOT]──▶ VeCPOTLocker│  │ ◀──[getBoostBps]
  │                   │syncBoost  │
  └──[激活 NFT]──▶ NFTBoostController ◀──[getBoostBps]
  
  ChapoolEarnVault ──[claimCPP → mint CPP]──▶ 用户 AA 账户
  
  ChapoolRewardDistributor ──[setRewardRate]──▶ ChapoolEarnVault
  
  前端 / DefiLlama ──▶ ChapoolVaultReader ──[只读聚合]──▶ 以上所有合约
```

**核心公式：**
```
userWeightedUSDT = usdtBalance × (10000 + veCPOTBoost + nftBoost) / 10000
userCPP/秒       = rewardRate × userWeightedUSDT / totalWeightedUSDT
```

- USDT 存款：1:1 托管，随时取回，不产生 USDT 收益
- CPP 奖励：按加权 USDT 比例持续累积，用户手动 `claimCPP()` 领取，mint 到 AA 账户
- 1 CPP = 0.002 USDT（参考换算，CPP 可在站内消费或充值 U 卡）

### 1.2 新增合约清单

| 合约 | 类型 | Phase | 核心职责 |
|------|------|-------|----------|
| `ChapoolEarnVault` | 可升级（UUPS） | P1 | USDT 托管、CPP 累加器、claimCPP |
| `ChapoolRewardDistributor` | 可升级（UUPS） | P1 | 控制 CPP 排放速率（纯控制平面） |
| `VeCPOTLocker` | 可升级（UUPS） | P2 | CPOT 锁仓、veCPOT 计算、提供 Boost |
| `NFTBoostController` | 可升级（UUPS） | P2 | 读取 CPNFT 持仓等级、提供 NFT Boost |
| `ChapoolVaultReader` | 非升级 | P1 | 聚合只读，供前端和 DefiLlama 查询 |

### 1.3 链下服务（不上链）

| 服务 | 作用 |
|------|------|
| CPP Consumption Service | CPP 站内消费余额与订单管理 |
| UCard Redemption Service | CPP → U卡提现申请、审批、到账 |
| Settlement Service | 线下消费清算、对账与财务核销 |

CPP 相关的前端展示由后端 API 提供数据，前端只负责渲染，无需新增链上合约。

---

## 2 模块一：ChapoolEarnVault

### 2.1 合约概述

- 资产：USDT（单资产托管，随存随取，1:1 返还，无 USDT 收益）
- 收益形式：CPP Token（按加权 USDT 比例连续累积，手动 claimCPP）
- 模式：活期，T+0 到账
- CPP mint 目标：用户 AA 账户（与 Staking 合约一致）

### 2.2 存款

#### 接口

```solidity
function deposit(uint256 assets, address receiver) external returns (uint256);
```

#### 逻辑

1. 从 `msg.sender` 转入 `assets` USDT
2. 调用 `_updateAccumulator()` 更新全局累加器
3. 调用 `_settleUser(receiver)` 结算待领取 CPP
4. 更新 `usdtBalance[receiver] += assets`
5. 调用 `_syncUserWeight(receiver)` 重新计算加权 USDT 并更新 `totalWeightedUSDT`
6. emit `Deposited(user, assets, timestamp)`

#### 前端调用序列

```
1. 读取用户 USDT balance → 展示可用余额
2. 调用 USDT.approve(vaultAddress, assets)
3. 调用 vault.deposit(assets, userAddress)
4. 监听 Deposited 事件
5. 刷新 usdtBalance、pendingCPP、estimatedDailyCPP
```

### 2.3 取款

#### 接口

```solidity
function withdraw(uint256 assets, address receiver) external returns (uint256);
```

> 活期产品，取出金额 = 存入金额，1:1 返还，T+0 到账，无手续费。

#### 逻辑

1. 验证 `usdtBalance[msg.sender] >= assets`
2. 调用 `_updateAccumulator()` + `_settleUser(msg.sender)`
3. 更新 `usdtBalance[msg.sender] -= assets`
4. 调用 `_syncUserWeight(msg.sender)`
5. 转出 USDT 给 `receiver`
6. emit `Withdrawn(user, assets, timestamp)`

### 2.4 CPP 领取

#### 接口

```solidity
function claimCPP() external returns (uint256 amount);
function getPendingCPP(address user) external view returns (uint256);
```

#### 逻辑

1. `_updateAccumulator()` + `_settleUser(msg.sender)`
2. 读取 `pendingCPP[msg.sender]`，清零
3. 调用 `_getAAAccount(msg.sender)` 解析 AA 账户
4. `cppToken.mint(aaAccount, amount)`
5. emit `CPPRewardClaimed(user, aaAccount, amount, timestamp)`

### 2.5 Boost 同步

#### 接口

```solidity
function syncBoost(address user) external;
```

调用时机：
- VeCPOTLocker 在 `lock()` / `unlock()` / `lockMore()` 后自动调用
- NFTBoostController `activateBoost()` / `deactivateBoost()` 后由用户手动调用
- 任何人可调用，无权限限制

#### 逻辑

1. `_updateAccumulator()` + `_settleUser(user)`
2. `_syncUserWeight(user)` — 重新读取最新 boost，更新 `weightedUSDT[user]` 和 `totalWeightedUSDT`

### 2.6 累加器机制（Synthetix 模式）

```solidity
// 每次状态变化前更新
function _updateAccumulator() internal {
    elapsed = block.timestamp - lastUpdateTime;
    accCPPPerWeightedUSDT += elapsed × rewardRate × PRECISION / totalWeightedUSDT;
    lastUpdateTime = block.timestamp;
}

// 结算用户待领取 CPP
function _settleUser(address user) internal {
    pendingCPP[user] += weightedUSDT[user] × (accCPPPerWeightedUSDT - rewardDebt[user]) / PRECISION;
    rewardDebt[user] = accCPPPerWeightedUSDT;
}

// 同步用户加权 USDT
function _syncUserWeight(address user) internal {
    boostBps = veCPOTBoost + nftBoost;  // 读取外部合约
    newWeighted = usdtBalance[user] × (10000 + boostBps) / 10000;
    totalWeightedUSDT = totalWeightedUSDT - weightedUSDT[user] + newWeighted;
    weightedUSDT[user] = newWeighted;
    rewardDebt[user] = accCPPPerWeightedUSDT;
}
```

### 2.7 只读查询

```solidity
function totalVaultAssets() external view returns (uint256);  // TVL for DefiLlama
function usdtBalance(address user) external view returns (uint256);
function weightedUSDT(address user) external view returns (uint256);
function getPendingCPP(address user) external view returns (uint256);
function getUserBoostBps(address user) external view returns (uint256);
function estimatedDailyCPP(address user) external view returns (uint256);
```

### 2.8 数据结构

```solidity
uint256 public rewardRate;              // CPP/秒（全局）
uint256 public lastUpdateTime;
uint256 public accCPPPerWeightedUSDT;   // 累加器 × PRECISION
uint256 public totalWeightedUSDT;

mapping(address => uint256) public usdtBalance;
mapping(address => uint256) public weightedUSDT;
mapping(address => uint256) public rewardDebt;
mapping(address => uint256) public pendingCPP;

struct UserVaultPosition {
    uint256 depositedAt;
    uint256 lastActionAt;
}
mapping(address => UserVaultPosition) public positions;
```

### 2.9 管理员接口

```solidity
function setRewardRate(uint256 cppPerSecond) external onlyOwner;
function setVecpotLocker(address locker) external onlyOwner;
function setNftBoostController(address controller) external onlyOwner;
function setAccountManager(address am) external onlyOwner;
function pause() / unpause() external onlyOwner;
```

> **关键约束**：owner 无法直接提取 TVL 资产，仅能通过 24 小时 timelock 的紧急提款流程操作。

---

## 3 模块二：VeCPOTLocker

### 3.1 合约概述

- 输入：CPOT（ERC20）
- 输出：veCPOT（不可转让积分，只读计算，不发行 ERC20）
- 锁仓到期后：CPOT 自动可提取，veCPOT 归零
- 奖励：CPP（ERC20，通过 `CPPToken.mint` 发放）

### 3.2 锁仓

每次锁仓创建一个独立仓位，返回 `lockId`，用户可同时持有多个不同周期的仓位。

#### 接口

```solidity
function lock(uint256 amount, uint256 durationDays) external returns (uint256 lockId);
```

> `durationDays` 取值：30 / 90 / 180 / 360

#### 逻辑

1. 从 `msg.sender` 转入 `amount` CPOT
2. 生成递增 `lockId`（每个用户独立计数）
3. 计算 `unlockTime = block.timestamp + durationDays * 1 days`
4. 计算该仓位的 veCPOT 积分（见 3.5 节）
5. 追加到 `lockPositions[user]` 数组
6. emit `LockedCPOT(user, lockId, amount, unlockTime, veAmount, durationDays)`

#### 前端调用序列

```
1. 用户选择档位（30/90/180/360天）+ 输入金额
2. 读取 locker.previewVeCPOT(amount, duration) → 展示新仓位 veCPOT
3. 读取 locker.previewBoostBps(user, duration) → 展示锁仓后年化
   （若新档位 > 当前最高档位，boost 会提升；否则不变）
4. 调用 CPOT.approve(lockerAddress, amount)
5. 调用 locker.lock(amount, durationDays) → 获得 lockId
6. 监听 LockedCPOT 事件
7. 刷新锁仓列表和加速数据
```

### 3.3 解锁

每个仓位独立解锁，通过 `lockId` 指定。

#### 接口

```solidity
function unlock(uint256 lockId) external;
```

#### 逻辑

1. 验证 `lockPositions[msg.sender][lockId]` 存在且归属正确
2. 验证 `block.timestamp >= position.unlockTime`
3. 将该仓位的 CPOT 原路退回 `msg.sender`
4. 从 `lockPositions[user]` 数组中移除该仓位
5. 重新计算用户的 veCPOT 总量，boost 随 veCPOT 减少而自动降低
6. emit `UnlockedCPOT(user, lockId, amount)`

> **注意**：解锁仓位后该仓位的 veCPOT 从总量中扣除，boost 随之降低。若还有其他活跃仓位，boost 不会归零。

### 3.4 向已有仓位追加金额

用户可向某个特定 `lockId` 追加 CPOT，解锁时间不变：

```solidity
function lockMore(uint256 lockId, uint256 additionalAmount) external;
```

> 追加后以当前仓位剩余时间重新计算该仓位的 veCPOT，不延长或缩短解锁时间。

### 3.5 veCPOT 计算规则

**单个仓位**：

```
position.veAmount = amount × (durationDays / 360)
```

**用户总 veCPOT = 所有活跃仓位的 veAmount 之和**：

```
totalVeCPOT(user) = Σ position[i].veAmount  （仅计算 unlockTime > now 的仓位）
```

示例（多仓位叠加）：

| 仓位 | 锁仓量 | 时长 | veAmount |
|------|--------|------|---------|
| Lock #1 | 5,000 CPOT | 360天 | 5,000 |
| Lock #2 | 2,000 CPOT | 90天 | 500 |
| Lock #3 | 1,000 CPOT | 30天 | 83 |
| **合计** | **8,000 CPOT** | — | **5,583** |

### 3.6 veCPOT 加速倍率换算

**Boost 由用户的 总 veCPOT 数量决定，数量越多倍率越高**，上限 +5%（500 bps）：

```
veUnits(user)  = totalVeCPOT(user) / 1e18
boostBps(user) = min( (veUnits × boostPerVeUnit) / BOOST_VE_PRECISION,  maxVecpotBoostBps )
```

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `boostPerVeUnit` | 1 | boost 分子（与精度配合使用） |
| `BOOST_VE_PRECISION` | 10 | 精度除数（常量，不可修改） |
| `maxVecpotBoostBps` | 500 | CPOT 加速上限（+5.0%） |

`boostPerVeUnit` 和 `maxVecpotBoostBps` 可由管理员通过 `setBoostParams()` 调整，无需升级合约。

**示例计算**（默认参数，实际效果：每 10 veCPOT 单位 = 1 bps）：

| 锁仓方案 | veCPOT 单位 | boostBps | 加速 |
|---------|------------|---------|------|
| 500 CPOT × 30天 | 41 | 4 | +0.04% |
| 1,000 CPOT × 90天 | 250 | 25 | +0.25% |
| 5,000 CPOT × 90天 | 1,250 | 125 | +1.25% |
| 10,000 CPOT × 90天 | 2,500 | 250 | +2.50% |
| 5,000 CPOT × 360天 | 5,000 | **500（达上限）** | +5.0% |
| 10,000 CPOT × 180天 | 5,000 | **500（达上限）** | +5.0% |

**关键特性**：
- 锁仓**数量**和**时长**同时影响 veCPOT，二者缺一不可
- 多个仓位的 veCPOT 可以叠加，持续增加直至上限
- 解锁仓位后 veCPOT 减少，boost 自动下降

#### 只读接口

```solidity
function getTotalVeCPOT(address user) external view returns (uint256);
function getBoostBps(address user) external view returns (uint256);
function getLockPositions(address user) external view returns (LockPosition[] memory);
function getLockPosition(address user, uint256 lockId) external view returns (LockPosition memory);
function previewVeCPOT(uint256 amount, uint256 durationDays) external view returns (uint256);
function previewBoostBps(address user, uint256 additionalAmount, uint256 durationDays) external view returns (uint256);
// 返回：若新增 additionalAmount CPOT 锁仓 durationDays 天后，用户 boost 会变为多少
```

### 3.7 Vault 通知（syncBoost 回调）

每次锁仓状态变化后，VeCPOTLocker 自动通知 ChapoolEarnVault 重新同步用户权重：

```solidity
// VeCPOTLocker 内部，lock/unlock/lockMore 后调用
function _notifyVault(address user) internal {
    if (vault != address(0)) {
        try IEarnVault(vault).syncBoost(user) {} catch {}
    }
}
```

> **前端无需手动调用 syncBoost**，VeCPOTLocker 的锁仓操作会自动触发。

### 3.8 数据结构

```solidity
struct LockPosition {
    uint256 lockId;       // 仓位 ID（用户维度递增）
    uint256 amount;       // 锁仓 CPOT 数量
    uint256 veAmount;     // 对应 veCPOT 积分
    uint256 startTime;    // 开始时间戳
    uint256 unlockTime;   // 解锁时间戳
    uint256 durationDays; // 锁仓天数（30/90/180/360）
}

// 每个用户有独立的仓位数组
mapping(address => LockPosition[]) public lockPositions;

// 用户下一个 lockId（递增计数器）
mapping(address => uint256) public nextLockId;
```

#### 辅助计算函数（内部）

```solidity
// Boost 计算核心：基于 veCPOT 总量，精度除数 BOOST_VE_PRECISION = 10
function _calcBoostBps(uint256 totalVe) internal view returns (uint256) {
    uint256 veUnits = totalVe / 1e18;
    uint256 bps     = (veUnits * boostPerVeUnit) / BOOST_VE_PRECISION;
    return bps > maxVecpotBoostBps ? maxVecpotBoostBps : bps;
}
```

---

## 4 模块三：NFTBoostController

### 4.1 产品规则

Phase 2 采用**注册激活、取最高等级**模式：

- 用户调用 `activateBoost(tokenId)` 注册一张 NFT 作为加速凭证
- **质押中的 NFT 不能激活**（避免与 Staking 合约双重获益）
- 卖出后 `ownerOf` 验证失败，boost 自动归零
- 多张 NFT **不叠加**，只有注册的那一张生效
- 随时可切换：`deactivateBoost()` 后重新 `activateBoost(newTokenId)`

> **示例**：用户持有 L1 + L3 + L2 三张，boost = L3 对应的 +2.0%，而非三张相加。

### 4.2 Boost 对照表

| NFT Level | APY Boost |
|-----------|-----------|
| L1 | +0.5% |
| L2 | +1.0% |
| L3 | +2.0% |
| L4 | +3.5% |

> 原型中统一展示 +1.5%（L2/L3 均值），合约实现时按 NFT 等级精确计算。

### 4.3 接口

```solidity
function getNFTBoostBps(address user) external view returns (uint256);
function getNFTLevel(uint256 tokenId) external view returns (uint256);
function getActiveNFT(address user) external view returns (uint256 tokenId, uint256 level);
```

#### 内部逻辑

```
1. 读取 CPNFT.balanceOf(user)
2. 如果余额 = 0，返回 boostBps = 0
3. 遍历用户持有的所有 tokenId（建议设置最大遍历数 MAX_SCAN = 20，防止 gas 超限）
4. 找出 level 最高的 tokenId（多张取最高，不累加）
5. 按 boost 对照表返回该 level 对应的 boostBps
```

> **注意**：`getNFTBoostBps` 为 view 函数，链下调用无 gas 限制；  
> 若需在链上被其他合约调用（如 Vault 计算实时 APY），需控制遍历上限避免 out-of-gas。

### 4.4 NFT 等级来源

从现有 `CPNFT` 合约读取 `level` 字段（已有实现）。

---

## 5 模块四：ChapoolRewardDistributor

### 5.1 职责

纯控制平面。负责根据平台 USDT 收入，换算为 CPP/秒排放速率并写入 ChapoolEarnVault。
无代币转移经过本合约，CPP 在用户 claimCPP 时由 Vault 直接 mint。

### 5.2 接口

```solidity
function setVaultRewardRate(uint256 cppPerSecond) external;  // 运营钱包调用
function pauseEmission() external;                            // 停止排放（rate = 0）
function getVaultRewardRate() external view returns (uint256);
```

### 5.3 排放速率计算

运营团队线下计算，基于公式：

```
cppPerSecond = (周期内平台USDT收入 / 0.002) / 周期秒数

例：本周平台收入 1,000 USDT，发放周期 7天（604800秒）
  → cppPerSecond = (1000 / 0.002) / 604800 = 500000 / 604800 ≈ 0.826 CPP/秒
  → 实际值：826_388_888_888_888（wei，18位精度）
```

- 操作频率：每周一次（运营后台触发）
- 调整前自动触发累加器更新，不影响存量用户奖励
- 需提前确保 CPPToken 已对 ChapoolEarnVault 授予 `MINTER_ROLE`

---

## 6 模块五：ChapoolVaultReader

### 6.1 职责

聚合多合约只读数据，供前端和 DefiLlama 查询，减少前端多次 RPC 调用。

### 6.2 接口

```solidity
// 全局概览
function getProtocolOverview() external view returns (
    uint256 tvl,              // Vault USDT 余额
    uint256 baseAPRBps,       // 基础年化（bps）
    uint256 maxAPRBps,        // 最高年化（含最高 boost）
    uint256 totalUsers,       // 存款用户数
    uint256 weeklyNetInflow   // 7日净流入
);

// 用户持仓详情
function getUserDashboard(address user) external view returns (
    uint256 principal,        // 存款本金（USDT）
    uint256 currentValue,     // 当前价值（含收益）
    uint256 earnedToDate,     // 累计收益
    uint256 effectiveAPRBps,  // 用户实际年化
    uint256 vecpotBoostBps,   // veCPOT 加速
    uint256 nftBoostBps,      // NFT 加速
    uint256 lockedCPOT,       // 锁仓 CPOT 量
    uint256 cpotUnlockTime,   // 解锁时间
    uint256 pendingCPP        // 可领取 CPP
);

// DefiLlama 最简口径
function getTVL() external view returns (uint256);
```

---

## 7 CPP 收益计算规则

### 7.1 核心公式

```
用户加权 USDT = USDT 存款 × (10000 + veCPOT_boost + NFT_boost) / 10000

用户每秒 CPP  = rewardRate × 用户加权USDT / 全网totalWeightedUSDT

每日 CPP 估算 = 用户每秒CPP × 86400
```

### 7.2 Boost 对照表

**veCPOT Boost（取最长活跃锁仓时长）：**

| 锁仓时长 | Boost | 加权系数示例（存1000 USDT） |
|---------|-------|--------------------------|
| 无锁仓  | +0%   | 1000 |
| 30天    | +1%   | 1010 |
| 90天    | +2%   | 1020 |
| 180天   | +3.5% | 1035 |
| 360天   | +5%   | 1050 |

**NFT Boost（注册的未质押 NFT 等级）：**

| NFT 等级 | Boost (bps) | Boost |
|---------|------------|-------|
| 无 / NORMAL / C | 0 | +0% |
| B (L1) | 10 | +0.1% |
| A (L2) | 20 | +0.2% |
| S (L3) | 50 | +0.5% |
| SS (L4) | 350 | +3.5% |

### 7.3 综合示例

假设：全网 totalWeightedUSDT = 1,000,000，rewardRate = 1 CPP/秒

| 用户 | USDT | veCPOT boost | NFT boost | 加权USDT | 每日CPP |
|------|------|-------------|-----------|---------|--------|
| A | 10,000 | 0% | 0% | 10,000 | 864 |
| B | 10,000 | 5% | 0% | 10,500 | 907 |
| C | 10,000 | 5% | 3.5% | 10,850 | 937 |
| D | 10,000 | 0% | 0% | 10,000 | 864×30天=**25,920** CPP/月 |

CPP 参考价值：25,920 CPP × 0.002 = **51.84 USDT/月**（D 用户，10,000 USDT 存款）

### 7.4 存款时长影响

收益按秒累积，存得越久领取的 CPP 越多。中途取款后剩余 USDT 的累积不中断：

```
存款1天 → 领 864 CPP
存款30天 → 领 25,920 CPP（在全网权重不变情况下）
```

---

### 7.5 实际样本数据演算

#### 全局参数假设

```
平台本周收入：10,000 USDT
发放周期：7 天（604,800 秒）
CPP 换算：1 CPP = 0.002 USDT → 总 CPP 池 = 10,000 / 0.002 = 5,000,000 CPP/周

rewardRate = 5,000,000 / 604,800 ≈ 8.267 CPP/秒

全网 totalWeightedUSDT = 1,000,000（假设平台共有 100 万加权 USDT 存款）
```

#### Boost 速查

**CPOT 加速**（默认参数：每 10 veCPOT 单位 = 1 bps，上限 500 bps）

```
veCPOT 单位 = Σ (锁仓 CPOT × 锁仓天数 / 360)
boostBps    = min(veCPOT 单位 / 10,  500)
```

| 锁仓方案示例 | veCPOT 单位 | CPOT Boost (bps) |
|------------|------------|-----------------|
| 无 | 0 | 0 |
| 1,000 CPOT × 90天 | 250 | 25 (+0.25%) |
| 5,000 CPOT × 90天 | 1,250 | 125 (+1.25%) |
| 10,000 CPOT × 90天 | 2,500 | 250 (+2.50%) |
| 5,000 CPOT × 360天 | 5,000 | **500（达上限，+5%）** |
| 10,000 CPOT × 180天 | 5,000 | **500（达上限，+5%）** |
| 50,000 CPOT × 360天 | 50,000 | **500（达上限，+5%）** |

**NFT 加速**（注册未质押 NFT，取最高等级）

| NFT 等级 | NFT Boost (bps) |
|---------|----------------|
| 无 / NORMAL / C | 0 |
| B (L1) | 10 (+0.1%) |
| A (L2) | 20 (+0.2%) |
| S (L3) | 50 (+0.5%) |
| SS (L4) | 350 (+3.5%) |

---

#### 用户 A — 小额存款，无加速

```
USDT = 500，veCPOT boost = 0 bps，NFT boost = 0 bps

加权USDT = 500 × (10000 + 0) / 10000 = 500

每日CPP = 8.267 × (500 / 1,000,000) × 86,400 = 357 CPP/天
每月CPP = 357 × 30 = 10,710 CPP
月参考价值 = 10,710 × 0.002 ≈ 21.4 USDT
等效月化率 ≈ 4.28%（年化约 51.4%）
```

---

#### 用户 B — 中等存款，锁仓 5,000 CPOT × 90 天

```
USDT = 5,000
锁仓：5,000 CPOT × 90天 → veCPOT 单位 = 5,000×90/360 = 1,250
veCPOT boost = min(1,250 / 10, 500) = 125 bps (+1.25%)，NFT boost = 0 bps

加权USDT = 5,000 × (10000 + 125) / 10000 = 5,062.5

每日CPP = 8.267 × (5,062.5 / 1,000,000) × 86,400 ≈ 3,615 CPP/天
每月CPP = 3,615 × 30 = 108,450 CPP
月参考价值 = 108,450 × 0.002 ≈ 216.9 USDT
等效月化率 ≈ 4.34%（年化约 52.0%）
```

---

#### 用户 C — 大额存款，锁仓 50,000 CPOT × 360 天 + NFT SS 级

```
USDT = 50,000
锁仓：50,000 CPOT × 360天 → veCPOT 单位 = 50,000×360/360 = 50,000
veCPOT boost = min(50,000 / 10, 500) = 500 bps（达上限），NFT SS boost = 350 bps，合计 850 bps

加权USDT = 50,000 × (10000 + 850) / 10000 = 54,250

每日CPP = 8.267 × (54,250 / 1,000,000) × 86,400 = 38,720 CPP/天
每月CPP = 38,720 × 30 = 1,161,600 CPP
月参考价值 = 1,161,600 × 0.002 ≈ 2,323 USDT
等效月化率 ≈ 4.65%（年化约 55.7%）
```

---

#### 用户 D — 中等存款，锁仓 20,000 CPOT × 180 天 + NFT A 级

```
USDT = 10,000
锁仓：20,000 CPOT × 180天 → veCPOT 单位 = 20,000×180/360 = 10,000
veCPOT boost = min(10,000 / 10, 500) = 500 bps（达上限），NFT A boost = 20 bps，合计 520 bps

加权USDT = 10,000 × (10000 + 520) / 10000 = 10,520

每日CPP = 8.267 × (10,520 / 1,000,000) × 86,400 ≈ 7,511 CPP/天
每月CPP = 7,511 × 30 = 225,330 CPP
月参考价值 = 225,330 × 0.002 ≈ 450.7 USDT
等效月化率 ≈ 4.51%（年化约 54.1%）
```

---

#### 汇总对比

| 用户 | USDT | CPOT 锁仓方案 | veCPOT单位 | CPOT boost | NFT boost | 加权USDT | 每日CPP | 每月CPP | 月参考价值 | 等效月化率 |
|-----|------|-------------|----------|-----------|-----------|---------|--------|--------|----------|----------|
| A | 500 | 无 | 0 | 0 bps | 0 bps | 500 | 357 | 10,710 | 21.4 USDT | 4.28% |
| B | 5,000 | 5,000 CPOT × 90天 | 1,250 | 125 bps | 0 bps | 5,062.5 | 3,615 | 108,450 | 216.9 USDT | 4.34% |
| C | 50,000 | 50,000 CPOT × 360天 | 50,000 | **500 bps（上限）** | 350 bps（SS） | 54,250 | 38,720 | 1,161,600 | 2,323 USDT | 4.65% |
| D | 10,000 | 20,000 CPOT × 180天 | 10,000 | **500 bps（上限）** | 20 bps（A） | 10,520 | 7,511 | 225,330 | 450.7 USDT | 4.51% |

> **注**：等效月化率 = 月参考价值 / USDT 存款 × 100%。CPP 实际价值随市场波动，此处按 1 CPP = 0.002 USDT 换算仅供参考。

#### 关键结论

1. **存款金额是首要因素** — 直接决定用户在全网 totalWeightedUSDT 中的份额
2. **Boost 提升边际收益约 0.3%~0.4%** — 从无 boost 到最高 boost（+8.5%），月化率从 4.28% 提升至 4.65%，因为全网其他用户也可能同时 boost
3. **时长影响两个维度**：锁仓越久 veCPOT boost 越高；在池中存得越久累积的 CPP 越多（按秒线性增长）
4. **平台收入增加 → rewardRate 增大 → 所有用户收益等比例提升**（收益池为零和竞争）
5. **CPP 流向**：用户 claimCPP() 后，CPP 直接 mint 到其 AA 账户，可用于站内消费或 U 卡充值线下消费

---

## 8 链上事件规范

所有关键操作都必须 emit 以下标准事件，方便 Dune 统计和 DefiLlama 接入：

```solidity
// ChapoolEarnVault
event Deposited(address indexed user, uint256 assets, uint256 timestamp);
event Withdrawn(address indexed user, uint256 assets, uint256 timestamp);
event CPPRewardClaimed(address indexed user, address indexed aaAccount, uint256 amount, uint256 timestamp);
event RewardRateSet(uint256 cppPerSecond, uint256 timestamp);
event BoostSynced(address indexed user, uint256 oldWeighted, uint256 newWeighted);

// ChapoolRewardDistributor
event VaultRewardRateSet(address indexed sender, uint256 cppPerSecond, uint256 timestamp);

// VeCPOTLocker
event LockedCPOT(address indexed user, uint256 lockId, uint256 amount, uint256 unlockTime, uint256 veAmount, uint256 durationDays);
event UnlockedCPOT(address indexed user, uint256 lockId, uint256 amount, uint256 timestamp);
event MoreCPOTLocked(address indexed user, uint256 lockId, uint256 additionalAmount, uint256 newVeAmount);

// NFTBoostController
event BoostActivated(address indexed user, uint256 indexed tokenId, uint8 level, uint256 boostBps);
event BoostDeactivated(address indexed user, uint256 indexed tokenId);
```

---

## 9 前端页面与交互规范

> 参考实现：`frontend-prototype/src/App.tsx`  
> 设计原则：用户视角文案，不使用 PM 语言。

### 9.1 页面结构（单页滚动）

```
[Header: 理财 / Earn]
│
├── [Section 1] 我的存款卡片（主 hero card）
│     ├── USDT 存款金额（大字，e.g. "$10,000"）
│     ├── 预计日收益（小字，e.g. "≈ 864 CPP / 天"）
│     ├── [存入 USDT] 按钮 → 触发存款 Sheet
│     └── [取出] 按钮 → 触发取款 Sheet
│
├── [Section 2] 我的收益
│     ├── 待领取 CPP（N,NNN CPP，绿色大字）+ [领取] 按钮
│     │     ├── 有待领取：按钮高亮，可点击
│     │     ├── 点击后按钮进入 loading 态（"···"）
│     │     ├── 领取成功：数量归零变"—"，按钮变灰 disabled
│     │     └── 最近记录顶部追加"CPP 已领取 +N,NNN CPP"
│     ├── 流动性：活期，随时可取
│     └── CPP 价值参考：1 CPP = 0.002 USDT（灰色小字）
│
├── [Section 3] 收益加速
│     ├── 右上角：当前加速倍率 badge（如 ×1.085）
│     ├── [Row] 锁仓 CPOT
│     │     ├── 状态：已锁 X CPOT · 解锁于 X月X日 / 未锁仓
│     │     ├── 右侧 boost badge：+N.N%（绿色）/ +0%（灰色）
│     │     └── 按钮：新增锁仓 / 立即锁仓 → 锁仓 Sheet
│     ├── [Divider]
│     ├── [Row] NFT 加速
│     │     ├── 状态：已激活 NFT #XXXX（等级 S）/ 未激活
│     │     ├── 右侧 boost badge：+N.N%（绿色）/ +0%（灰色）
│     │     └── 按钮：激活加速 / 获取 NFT → NFT Sheet
│     └── [收益明细]
│           ├── 基础权重：×1.000（无加速）
│           ├── CPOT 加速：+N.N%
│           ├── NFT 加速：+N.N%
│           └── 综合倍率：×N.NNN
│
└── [Section 4] 最近记录
      ├── 动态列表，每次操作成功后实时追加至顶部
      └── 记录类型：deposit（存入）/ withdraw（取出）/ lock（锁仓）/ cpp_claim（领取 CPP）
```

### 9.2 存款 Sheet（底部弹出层）

#### 状态机

```
[form 态] → 点击确认 → [success 态] → 点击完成 → [关闭 sheet]
```

#### form 态内容

| 元素 | 说明 |
|------|------|
| 顶部 Tab | 存入 / 取出（切换 mode） |
| 金额输入框 | 数字键盘，右对齐，单位 USDT |
| 全部按钮 | 存入模式填 USDT 余额；取出模式填当前本金 |
| 预览卡片（输入 > 0 时展示） | 预计日收益 ≈ N CPP、活期可取 ✓ |
| 确认按钮 | 金额 = 0 时 disabled |

#### success 态内容

| 元素 | 内容 |
|------|------|
| 大勾图标 | ✓ 绿色圆形 |
| 标题 | 存入成功 / 已发起取出 |
| 说明（存入） | "已存入 X USDT，CPP 奖励实时累积，随时可领取。" |
| 说明（取出） | "取出申请已提交，通常 T+0 到账。" |
| 完成按钮 | 关闭 sheet；本金数字、总资产数字联动更新；最近记录顶部追加新条目 |

#### 数据联动

- 存款成功后：USDT 存款 +X，预计日收益重新计算，最近记录顶部追加 `↓ 存入 USDT +$X`
- 取出成功后：USDT 存款 -X，预计日收益重新计算，最近记录顶部追加 `↑ 取出 USDT -$X`（金额显示红色）

### 9.3 锁仓 Sheet（底部弹出层）

#### 状态机

```
[form 态] → 下一步 → [confirm 态] → 确认锁仓 → [success 态] → 完成 → [关闭]
```

#### form 态内容

| 元素 | 说明 |
|------|------|
| 标题 | "锁 CPOT 提升年化" |
| 副文案 | "锁仓越多越久，veCPOT 越高，CPP 加速倍率越大。到期自动解锁，CPOT 原路退回。" |
| 档位选择（2×2 格子） | 30天 +1.0%、90天 +2.0%、180天 +3.5%、360天 +5.0% |
| CPOT 金额输入 | 数字键盘 |
| 全部按钮 | 填入用户 CPOT 余额 |
| APY 计算区 | 基础年化 + 已有加速 + 本次加速 = 锁仓后年化（实时计算） |
| 下一步按钮 | 金额 = 0 时 disabled |

#### confirm 态内容

| 项目 | 值 |
|------|-----|
| 锁仓数量 | X CPOT |
| 锁仓时长 | X 天 |
| 新增加速 | +N.N%（绿色） |
| 锁仓后年化 | N.N%（绿色） |
| 风险提示 | "锁仓期间无法取出 CPOT，到期自动解锁并退回钱包" |
| 确认锁仓 | 主按钮 |
| 返回修改 | 次按钮 |

#### success 态内容

| 元素 | 内容 |
|------|------|
| 大勾 | ✓ |
| 标题 | "锁仓成功" |
| 说明 | "已锁 X CPOT，锁期 N 天" |
| APY 提示 | "你的 Earn 年化已提升至 N.N%"（绿色） |

#### 数据联动

- 成功后：锁仓列表追加新条目
- 若新仓位时长 > 当前最高时长：boost badge 升级，APY 明细更新，总资产年化更新
- 若新仓位时长 ≤ 当前最高时长：boost 不变（无需刷新年化）
- 最近记录顶部追加 `🔒 锁仓 CPOT N,NNN CPOT`

### 9.3.1 锁仓列表展示

位于"收益加速"卡片的 veCPOT 区块内，格式如下：

```
锁仓 CPOT                           +5.0%  [新增锁仓]
─────────────────────────────────────────────────────
Lock #1  5,000 CPOT  解锁于 6月1日          [6月1日解锁]（灰色 disabled）
Lock #2  2,000 CPOT  解锁于 3月20日  已到期  [解锁]（高亮可点击）
Lock #3  1,000 CPOT  解锁于 12月1日         [12月1日解锁]（灰色 disabled）
```

#### 解锁按钮规则

| 状态 | 按钮文案 | 样式 |
|------|---------|------|
| 未到期 | "X月X日解锁" / "Unlocks X" | 灰色 disabled |
| 已到期 | "解锁" / "Unlock" | 高亮，可点击 |
| 点击中 | "···" | loading disabled |
| 解锁成功 | 该行从列表中移除，追加历史记录 `🔓 解锁 CPOT N,NNN CPOT` |

#### 解锁后 boost 变化

- 解锁后该仓位的 veCPOT 从总量中扣除，boost 按新的 totalVeCPOT 重新计算
- boost badge / 加速倍率明细 / 预计日收益自动联动更新

### 9.4 NFT 信息 Sheet

| 元素 | 内容 |
|------|------|
| 图标 | 🖼 |
| 标题 | "如何获得 NFT 加速" |
| 说明 | "持有 Chapool CPNFT 即可自动激活 +1.5% 年化加速。NFT 可在 Marketplace 购买，或参与平台活动获取。" |
| 主按钮 | "去 Marketplace"（跳转） |
| 次按钮 | "关闭" |

### 9.5 最近记录

#### 记录类型与格式

| 图标 | 类型 key | 文案（中） | 文案（英） | 金额格式 | 金额颜色 |
|------|---------|-----------|-----------|---------|---------|
| 💰 | `cpp_claim` | CPP 已领取 | CPP claimed | +N,NNN CPP | 默认 |
| ↓ | `deposit` | 存入 USDT | Deposited USDT | +$N,NNN | 默认 |
| ↑ | `withdraw` | 取出 USDT | Withdrew USDT | -$N,NNN | 红色 |
| 🔒 | `lock` | 锁仓 CPOT | Locked CPOT | N,NNN CPOT | 默认 |
| 🔓 | `unlock` | 解锁 CPOT | Unlocked CPOT | N,NNN CPOT | 默认 |

> **重要**：CPP 收益不自动到账，必须由用户主动点击"领取"按钮触发链上 `claimCPP()` 调用后才会计入历史。  
> 历史记录为动态列表，每次操作成功后立即在顶部追加新条目，无需重新加载页面。

#### CPP 金额换算说明

后端 API 和链上返回的 CPP 数量以原始 token 单位展示，前端**不做 USDT 换算**，直接显示 CPP 数量。  
参考汇率（仅供用户理解，不在 UI 中主动换算）：1 CPP = 0.002 USDT。

### 9.6 CPP 领取交互规范

#### 入口位置

"我的存款"卡片最后一行，标签"待领取 CPP"，右侧显示金额和领取按钮。

#### 状态机

```
[有待领取 CPP] → 点击"领取" → [loading，按钮显示"···"] → 链上确认 → [成功]
                                                                  │
                                                         pendingCPP 归零 → 按钮 disabled
                                                         最近记录顶部追加 "CPP 已领取 +N CPP"
```

#### 边界条件

| 场景 | 处理 |
|------|------|
| `pendingCPP = 0` | 按钮 disabled，灰色样式 |
| 点击中（loading） | 禁止重复点击 |
| 链上 revert | Toast 提示失败原因，按钮恢复可点击 |

---

## 10 前端数据接入接口

### 10.1 页面初始化数据（一次性调用）

调用 `ChapoolVaultReader.getUserDashboard(userAddress)` 获取：

```typescript
interface UserDashboard {
  principal: bigint          // 存款本金（USDT，18位精度）
  currentValue: bigint       // 当前价值（含收益）
  earnedToDate: bigint       // 累计收益
  effectiveAPRBps: number    // 用户实际年化（bps，1% = 100）
  vecpotBoostBps: number     // veCPOT 加速 bps
  nftBoostBps: number        // NFT 加速 bps
  lockedCPOT: bigint         // 锁仓 CPOT 数量
  cpotUnlockTime: number     // 解锁时间戳（0 表示未锁仓）
  pendingCPP: bigint         // 可领取 CPP
}
```

### 10.2 全局协议数据

调用 `ChapoolVaultReader.getProtocolOverview()` 获取：

```typescript
interface ProtocolOverview {
  tvl: bigint           // TVL（USDT）
  baseAPRBps: number    // 基础年化（bps）
  maxAPRBps: number     // 最高年化（含最高 boost）
  totalUsers: number    // 存款用户数
  weeklyNetInflow: bigint
}
```

### 10.3 锁仓预览（选档位时实时调用）

```typescript
// 调用 VeCPOTLocker.previewBoostBps(durationDays)
// 返回 boostBps，前端计算：新年化 = 基础APY + 当前boost + 新boost
```

### 10.4 存款预览（输入金额时实时调用）

```typescript
// 调用 ChapoolEarnVault.previewDeposit(assets)
// 返回将获得的 shares（即 ceUSDT 数量）
```

### 10.5 CPP 数据（调用后端 API，非链上）

```typescript
// GET /api/v1/user/:address/cpp
interface CPPProfile {
  claimable: string      // 可领取 CPP
  monthlyEarned: string  // 本月已赚 CPP
  spent: string          // 已消费 CPP
  redeemedToCard: string // 已提现到 U卡 CPP
}
```

### 10.6 交易历史（链上事件 + 后端聚合）

```typescript
// GET /api/v1/user/:address/history
interface HistoryItem {
  type: 'deposit' | 'withdraw' | 'lock' | 'unlock' | 'cpp_claim'
  amount: string       // 原始数量字符串，前端按 type 决定单位
  // deposit/withdraw:  "$N,NNN" USDT
  // lock/unlock:       "N,NNN CPOT"
  // cpp_claim:         "N,NNN CPP"（不换算 USDT）
  timestamp: number
  txHash: string
}
```

> **说明**：`reward`（自动到账）类型已移除，CPP 收益统一改为用户主动领取，对应 `cpp_claim`。  
> 前端列表为**动态追加**模式：初始化时拉取历史列表，每次操作成功后在客户端顶部插入新条目，无需重新请求。

---

## 11 错误处理与边界条件

### 11.1 合约层

| 场景 | 错误处理 |
|------|---------|
| 存款金额为 0 | `revert ZeroAmount()` |
| USDT 余额不足 | `revert InsufficientBalance()` |
| 未授权 | ERC20 标准 revert |
| 合约暂停 | `revert ContractPaused()` |
| 锁仓档位非法 | `revert InvalidDuration()` |
| 未到期提前解锁 | `revert NotUnlockedYet(unlockTime)` |
| 取款超出余额 | `revert InsufficientShares()` |

### 11.2 前端层

| 场景 | UI 处理 |
|------|---------|
| 金额为 0 | 确认按钮 disabled |
| 超出余额 | 按钮 disabled + 提示"余额不足" |
| 钱包未连接 | 显示"连接钱包"占位按钮 |
| 交易 pending | 按钮 loading 态，禁止重复提交 |
| 交易失败 | Toast 提示错误原因 |
| 网络错误 | Toast 提示"网络异常，请稍后重试" |

---

## 12 安全与权限边界

### 12.1 合约权限设计

| 操作 | 权限 |
|------|------|
| 存款/取款 | 任意用户 |
| 锁仓/解锁 CPOT | 任意用户 |
| 注入收益 | `REWARD_DEPOSITOR_ROLE`（后台服务账户） |
| 暂停/恢复 | `DEFAULT_ADMIN_ROLE` |
| 更新参数（APY上限、最小存款等） | `DEFAULT_ADMIN_ROLE` |
| 发起紧急提款 | `DEFAULT_ADMIN_ROLE` + emergencyMode 开启 |
| 执行紧急提款 | `DEFAULT_ADMIN_ROLE` + timelock 到期 |

### 12.2 关键约束

- 收益注入只能通过 `depositRewards()` 路径进入，不得与主池混淆
- veCPOT 不可转让、不可出售
- Boost 叠加有上限（`MAX_BOOST_BPS = 2000`，即 20%）
- Phase 2 NFT Boost 为只读（不需要用户再次授权或 approve）
- 紧急提款有 timelock 保护，不可绕过（详见 12.4）

### 12.4 紧急提款机制

#### 设计目标

在合约出现严重风险时，允许 admin 将资金转移到安全地址，同时给用户足够时间知情并自行赎回。

#### 合约接口

```solidity
// Step 1: 开启紧急模式（同时暂停新存款）
function setEmergencyMode(bool enabled) external onlyRole(DEFAULT_ADMIN_ROLE);

// Step 2: 发起提款申请（必须已开启 emergencyMode）
function initiateEmergencyWithdraw(
    uint256 amount,
    address to
) external onlyRole(DEFAULT_ADMIN_ROLE);

// Step 3: timelock 到期后执行
function executeEmergencyWithdraw() external onlyRole(DEFAULT_ADMIN_ROLE);

// 取消未执行的申请
function cancelEmergencyWithdraw() external onlyRole(DEFAULT_ADMIN_ROLE);
```

#### 流程说明

```
admin 调用 setEmergencyMode(true)
  → 合约暂停新存款
  → emit EmergencyModeSet(true, block.timestamp)
  → 链上公开可见，用户可自行赎回

admin 调用 initiateEmergencyWithdraw(amount, to)
  → 记录 PendingWithdraw { amount, to, executeAfter: now + EMERGENCY_TIMELOCK }
  → emit EmergencyWithdrawInitiated(amount, to, executeAfter)

[等待 EMERGENCY_TIMELOCK 时间]

admin 调用 executeEmergencyWithdraw()
  → 验证 block.timestamp >= executeAfter
  → 转出 amount USDT 至 to 地址
  → emit EmergencyWithdrawExecuted(amount, to)
```

#### 参数

| 参数 | 值 | 说明 |
|------|-----|------|
| `EMERGENCY_TIMELOCK` | 24 小时 | admin 发起申请到可执行的最短等待时间 |
| `to` | 任意地址 | 由 admin 在发起时指定，无地址白名单限制 |
| `amount` | ≤ vault 当前余额 | 不可超出合约实际持有量 |

#### 数据结构

```solidity
bool public emergencyMode;

struct PendingWithdraw {
    uint256 amount;
    address to;
    uint256 executeAfter; // block.timestamp + EMERGENCY_TIMELOCK
    bool exists;
}

PendingWithdraw public pendingEmergencyWithdraw;

uint256 public constant EMERGENCY_TIMELOCK = 24 hours;
```

#### 关键事件

```solidity
event EmergencyModeSet(bool enabled, uint256 timestamp);
event EmergencyWithdrawInitiated(uint256 amount, address to, uint256 executeAfter);
event EmergencyWithdrawExecuted(uint256 amount, address to, uint256 timestamp);
event EmergencyWithdrawCancelled(uint256 timestamp);
```

#### 用户保护机制

1. `EmergencyModeSet` 事件发出后，用户可通过区块链浏览器或前端提醒知悉
2. 24 小时窗口期内，用户仍可正常赎回（取款功能不暂停，只暂停存款）
3. 申请发出后 admin 可随时取消（`cancelEmergencyWithdraw`）
4. 所有操作全程链上可查，无隐式后门

### 12.3 升级约束

- 所有核心合约使用 UUPS 模式
- 升级操作需要多签或 timelock（≥ 48小时）
- 升级时不得修改现有存储布局（`_gap` 占位）

---

## 附录：原型与合约的参数对应

| 原型参数 | 合约参数 | 说明 |
|---------|---------|------|
| `BOOST_PER_VE_UNIT = 1` | `boostPerVeUnit = 1` | veCPOT 加速系数（分子） |
| `BOOST_VE_PRECISION = 10` | `BOOST_VE_PRECISION = 10` | veCPOT 加速精度除数（常量） |
| `MAX_VECPOT_BOOST_BPS = 500` | `maxVecpotBoostBps = 500` | veCPOT 加速上限（+5%） |
| `NFT_BOOST_BPS = 50`（S级演示） | `BOOST_B/A/S/SS_BPS = 10/20/50/350` | NFT 分等级精确计算 |
