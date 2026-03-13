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
  ├──[存 USDT]──▶ ChapoolEarnVault ◀──[注入收益]── ChapoolRewardDistributor
  │                     │
  ├──[锁 CPOT]──▶ VeCPOTLocker ────[boostBps(user)]──▶ ChapoolEarnVault
  │
  └──[持 CPNFT]──▶ NFTBoostController ─[boostBps(user)]──▶ ChapoolEarnVault
  
  ChapoolVaultReader ◀──[只读聚合]── 以上所有合约
  
  前端 / DefiLlama ──▶ ChapoolVaultReader
```

### 1.2 新增合约清单

| 合约 | 类型 | Phase | 核心职责 |
|------|------|-------|----------|
| `ChapoolEarnVault` | 可升级（UUPS） | P1 | 存取款、份额发行、位置记录 |
| `ChapoolRewardDistributor` | 可升级（UUPS） | P1 | 收益注入、指数更新 |
| `VeCPOTLocker` | 可升级（UUPS） | P2 | CPOT 锁仓、veCPOT 计算、CPP 奖励资格 |
| `NFTBoostController` | 可升级（UUPS） | P2 | 读取 CPNFT 持仓、计算 NFT boost |
| `ChapoolVaultReader` | 非升级 | P1 | 聚合只读，供前端和第三方查询 |

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

- 资产：USDT（单资产，Phase 1-3 不支持多资产）
- 份额凭证：`ceUSDT`（ERC20，1:1 初始兑换比，随收益升值）
- 模式：活期，随存随取，Phase 3 增加定期模式

### 2.2 存款

#### 接口

```solidity
function deposit(uint256 assets, address receiver) external returns (uint256 shares);
```

#### 逻辑

1. 从 `msg.sender` 转入 `assets` USDT
2. 计算应发份额：`shares = assets * totalShares / totalAssets`（首次存款时 1:1）
3. 向 `receiver` mint `shares` ceUSDT
4. 更新 `totalAssets` 和 `totalShares`
5. 记录 `UserVaultPosition`
6. emit `Deposited(user, assets, shares)`

#### 前置条件

- 用户已 `approve` Vault 合约 ≥ `assets` USDT
- `assets > 0`
- 合约未暂停

#### 前端调用序列

```
1. 读取用户 USDT balance → 展示可用余额
2. 读取 vault.previewDeposit(assets) → 展示预计获得 ceUSDT
3. 调用 USDT.approve(vaultAddress, assets)
4. 调用 vault.deposit(assets, userAddress)
5. 监听 Deposited 事件
6. 更新页面持仓数据
```

### 2.3 取款（赎回）

#### 接口

```solidity
function withdraw(uint256 assets, address receiver, address owner) external returns (uint256 shares);
function redeem(uint256 shares, address receiver, address owner) external returns (uint256 assets);
```

> 前端统一使用 `redeem`（输入份额）或 `withdraw`（输入资产量），两种方式都支持。  
> 活期产品无罚金，T+0 到账。

#### 逻辑

1. 验证 `owner` 拥有足够 ceUSDT
2. 计算 `assets = shares * totalAssets / totalShares`
3. burn `shares` ceUSDT
4. 转出 `assets` USDT 给 `receiver`
5. 更新 `totalAssets` 和 `totalShares`
6. emit `Withdrawn(user, assets, shares)`

### 2.4 只读查询

```solidity
function totalAssets() external view returns (uint256);
function totalShares() external view returns (uint256);
function balanceOf(address user) external view returns (uint256 shares);
function convertToAssets(uint256 shares) external view returns (uint256);
function convertToShares(uint256 assets) external view returns (uint256);
function previewDeposit(uint256 assets) external view returns (uint256 shares);
function previewRedeem(uint256 shares) external view returns (uint256 assets);
```

### 2.5 数据结构

```solidity
struct UserVaultPosition {
    uint256 shares;         // 持有份额
    uint256 principal;      // 存入本金（用于前端展示，不影响份额计算）
    uint256 depositedAt;    // 首次存入时间戳
    uint256 lastActionAt;   // 最后操作时间戳
}

mapping(address => UserVaultPosition) public positions;
uint256 public totalAssets;
uint256 public totalShares;
address public asset;       // USDT 合约地址
```

### 2.6 管理员接口

```solidity
function pause() external onlyOwner;
function unpause() external onlyOwner;
function setRewardDistributor(address dist) external onlyOwner;
function setBoostController(address boost) external onlyOwner;
```

> **关键约束**：owner 无法直接提取 TVL 主池资产，只能管理参数与状态。

---

## 3 模块二：VeCPOTLocker

### 3.1 合约概述

- 输入：CPOT（ERC20）
- 输出：veCPOT（不可转让积分，只读计算，不发行 ERC20）
- 锁仓到期后：CPOT 自动可提取，veCPOT 归零
- 奖励：CPP（ERC20，通过 `CPPToken.mint` 发放）

### 3.2 锁仓

#### 接口

```solidity
function lock(uint256 amount, uint256 durationDays) external;
```

> `durationDays` 取值：30 / 90 / 180 / 360

#### 逻辑

1. 从 `msg.sender` 转入 `amount` CPOT
2. 计算 `unlockTime = block.timestamp + durationDays * 1 days`
3. 计算 veCPOT 积分（见 3.5 节）
4. 记录 `LockPosition`
5. emit `LockedCPOT(user, amount, unlockTime, veAmount)`

#### 前端调用序列

```
1. 用户选择档位（30/90/180/360天）
2. 读取 locker.previewVeCPOT(amount, duration) → 展示 veCPOT 积分
3. 读取 vault.getUserBoostBps(user) after → 展示锁仓后年化
4. 调用 CPOT.approve(lockerAddress, amount)
5. 调用 locker.lock(amount, durationDays)
6. 监听 LockedCPOT 事件
7. 刷新页面加速数据
```

### 3.3 解锁

#### 接口

```solidity
function unlock() external;
```

#### 逻辑

1. 验证 `block.timestamp >= position.unlockTime`
2. 将 CPOT 原路退回 `msg.sender`
3. 清除 `LockPosition`
4. veCPOT 积分归零（影响 Boost 系数）
5. emit `UnlockedCPOT(user, amount)`

### 3.4 增量锁仓

用户可在锁仓期内追加锁仓：

```solidity
function lockMore(uint256 additionalAmount) external;
```

> 追加时以当前剩余时间重新计算 veCPOT，不延长或缩短解锁时间。

### 3.5 veCPOT 计算规则

```
veCPOT = amount × (durationDays / 360)
```

示例：

| 锁仓量 | 时长 | veCPOT |
|--------|------|--------|
| 5,000 CPOT | 30天 | 416.7 |
| 5,000 CPOT | 90天 | 1,250 |
| 5,000 CPOT | 180天 | 2,500 |
| 5,000 CPOT | 360天 | 5,000 |

### 3.6 APY Boost 换算

veCPOT → Earn APY boost（基点，`bps`，1% = 100 bps）：

| 锁仓时长 | APY 加成 |
|---------|---------|
| 30天 | +1.0% |
| 90天 | +2.0% |
| 180天 | +3.5% |
| 360天 | +5.0% |

> **说明**：Boost 档位由 `durationDays` 决定，与 `amount` 无关（原型实现逻辑）。  
> 如需更精细的线性模型，可按 `veCPOT` 积分绝对值分档，Phase 1-2 使用简单档位即可。

#### 只读接口

```solidity
function getVeCPOT(address user) external view returns (uint256);
function getBoostBps(address user) external view returns (uint256);
function getLockPosition(address user) external view returns (LockPosition memory);
function previewVeCPOT(uint256 amount, uint256 durationDays) external view returns (uint256);
function previewBoostBps(uint256 durationDays) external view returns (uint256);
```

### 3.7 CPP 奖励

CPP 奖励**不自动发放**，必须由用户主动调用 `claimCPP()` 领取：

- `getPendingCPP(user)` 返回当前可领取数量，前端轮询或事件触发刷新
- 领取后链上 mint CPP 至用户地址，emit `CPPRewardClaimed` 事件
- 奖励来源由 `ChapoolRewardDistributor` 定期注入 CPP 额度，按用户 veCPOT 比例分配

```solidity
function claimCPP() external returns (uint256 amount);
function getPendingCPP(address user) external view returns (uint256);
```

> **前端对应**：`VeCPOTLocker.getPendingCPP(userAddress)` 的返回值直接显示为"待领取 CPP"，  
> 不换算为 USDT，保持 CPP 原始数量展示。

### 3.8 数据结构

```solidity
struct LockPosition {
    uint256 amount;       // 锁仓 CPOT 数量
    uint256 startTime;    // 开始时间戳
    uint256 unlockTime;   // 解锁时间戳
    uint256 durationDays; // 锁仓天数（30/90/180/360）
}

mapping(address => LockPosition) public lockPositions;
```

---

## 4 模块三：NFTBoostController

### 4.1 产品规则

Phase 2 采用**持有即加成**模式：

- 用户持有 CPNFT 即自动计入 boost，无需额外操作
- 不要求重新质押 NFT（避免与现有 `Staking` 冲突）
- 同一地址最多计算 **1 个 NFT** 的 boost（Phase 2，后续可扩展）
- boost 取最高等级 NFT

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
3. 遍历用户持有的 tokenId（最多查询 N 个）
4. 取最高 level 的 tokenId
5. 按 boost 对照表返回 boostBps
```

### 4.4 NFT 等级来源

从现有 `CPNFT` 合约读取 `level` 字段（已有实现）。

---

## 5 模块四：ChapoolRewardDistributor

### 5.1 职责

- 接收平台收益（USDT 手续费、补贴注入）
- 更新 Vault 收益指数
- 接收 CPP 分配额度，供 `VeCPOTLocker` 用于奖励发放

### 5.2 接口

```solidity
function depositRewards(uint256 usdtAmount) external;
function depositCPPRewards(uint256 cppAmount) external;
function getVaultAPR() external view returns (uint256);
```

### 5.3 收益注入规则

- 操作频率：每日或每周一次（由运营后台触发）
- 注入来源：Marketplace 手续费收入 + 平台补贴
- 注入后 `accRewardPerShare` 更新，`Vault` 的 `convertToAssets` 自动升值

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

## 7 APY 计算规则

### 7.1 公式

```
用户实际 APY = 基础 APY + veCPOT Boost + NFT Boost
```

### 7.2 示例数据（原型参数）

```
基础 APY:       8.0%
veCPOT Boost:  +0.0% / +1.0% / +2.0% / +3.5% / +5.0%（按锁仓档位）
NFT Boost:     +0.5% / +1.0% / +2.0% / +3.5%（按 NFT 等级）

最大叠加上限:   20%（防止极端激励失控）
```

### 7.3 Boost 叠加示例

| 场景 | 基础 APY | veCPOT | NFT | 最终 APY |
|------|---------|--------|-----|---------|
| 无加速 | 8.0% | 0 | 0 | 8.0% |
| 锁仓 90天 | 8.0% | +2.0% | 0 | 10.0% |
| 持 NFT L2 | 8.0% | 0 | +1.0% | 9.0% |
| 锁仓 180天 + NFT L3 | 8.0% | +3.5% | +2.0% | 13.5% |
| 锁仓 360天 + NFT L4 | 8.0% | +5.0% | +3.5% | 16.5% |

---

## 8 链上事件规范

所有关键操作都必须 emit 以下标准事件，方便 Dune 统计和 DefiLlama 接入：

```solidity
// ChapoolEarnVault
event Deposited(address indexed user, uint256 assets, uint256 shares, uint256 timestamp);
event Withdrawn(address indexed user, uint256 assets, uint256 shares, uint256 timestamp);

// ChapoolRewardDistributor
event RewardsDeposited(address indexed sender, uint256 usdtAmount, uint256 timestamp);
event CPPRewardsDeposited(address indexed sender, uint256 cppAmount, uint256 timestamp);

// VeCPOTLocker
event LockedCPOT(address indexed user, uint256 amount, uint256 unlockTime, uint256 veAmount, uint256 durationDays);
event UnlockedCPOT(address indexed user, uint256 amount, uint256 timestamp);
event CPPRewardClaimed(address indexed user, uint256 amount, uint256 timestamp);
event MoreCPOTLocked(address indexed user, uint256 additionalAmount, uint256 newVeAmount);
```

---

## 9 前端页面与交互规范

> 参考实现：`frontend-prototype/src/App.tsx`  
> 设计原则：用户视角文案，不使用 PM 语言。

### 9.1 页面结构（单页滚动）

```
[Header: 理财 / Earn]
│
├── [Section 1] 总资产卡片
│     ├── 总资产金额（本金 + 累计收益折算 USDT）
│     ├── 当前年化（含 boost）
│     ├── [存入 USDT] 按钮 → 触发存款 Sheet
│     └── [取出] 按钮 → 触发取款 Sheet
│
├── [Section 2] 我的存款
│     ├── 本金
│     ├── 流动性（活期，随时可取）
│     └── 待领取 CPP（N,NNN CPP，绿色）+ [领取] 按钮
│           ├── 有待领取：按钮高亮，可点击
│           ├── 点击后按钮进入 loading 态（"···"）
│           ├── 领取成功：数量归零变"—"，按钮变灰 disabled
│           └── 最近记录顶部实时追加"CPP 已领取 +N,NNN CPP"
│
├── [Section 3] 收益加速
│     ├── 右上角：当前实际年化 badge
│     ├── [Row] 锁仓 CPOT
│     │     ├── 当前状态：已锁 X CPOT · 解锁于 X月X日 / 未锁仓
│     │     ├── 右侧 boost badge：+N.N%（绿色）/ +0%（灰色）
│     │     └── 按钮：继续锁仓 / 立即锁仓 → 触发锁仓 Sheet
│     ├── [Divider]
│     ├── [Row] NFT 加速
│     │     ├── 当前状态：持有 N 个 CPNFT / 未持有
│     │     ├── 右侧 boost badge：+N.N%（绿色）/ +0%（灰色）
│     │     └── 按钮：获取 NFT → 触发 NFT 信息 Sheet
│     └── [APY 明细]
│           ├── 基础年化 8.0%
│           ├── veCPOT 加速 +N.N%（若有）
│           ├── NFT 加速 +N.N%（若有）
│           └── 当前年化 = 合计
│
└── [Section 4] 最近记录
      ├── 动态列表，每次操作成功后实时追加至顶部
      └── 初始加载：从后端 API 拉取历史（含 deposit / withdraw / lock / cpp_claim）
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
| 预览卡片（输入 > 0 时展示） | 你将获得 X ceUSDT、当前年化 N%、活期可取 ✓ |
| 确认按钮 | 金额 = 0 时 disabled |

#### success 态内容

| 元素 | 内容 |
|------|------|
| 大勾图标 | ✓ 绿色圆形 |
| 标题 | 存入成功 / 已发起取出 |
| 说明（存入） | "已存入 X USDT，收益从明天开始计算。" |
| 说明（取出） | "取出申请已提交，通常 T+0 到账。" |
| 完成按钮 | 关闭 sheet；本金数字、总资产数字联动更新；最近记录顶部追加新条目 |

#### 数据联动

- 存款成功后：页面本金 +X，总资产 +X，最近记录顶部追加 `↓ 存入 USDT +$X`
- 取出成功后：页面本金 -X，总资产 -X，最近记录顶部追加 `↑ 取出 USDT -$X`（金额显示红色）

### 9.3 锁仓 Sheet（底部弹出层）

#### 状态机

```
[form 态] → 下一步 → [confirm 态] → 确认锁仓 → [success 态] → 完成 → [关闭]
```

#### form 态内容

| 元素 | 说明 |
|------|------|
| 标题 | "锁 CPOT 提升年化" |
| 副文案 | "锁仓时间越长，年化加成越高。到期自动解锁，CPOT 原路退回。" |
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

- 成功后：收益加速卡片中 veCPOT 的 +N% badge 更新
- APY 明细中 veCPOT 一行数值更新
- 总资产卡片年化数字更新
- 最近记录顶部追加 `🔒 锁仓 CPOT N,NNN CPOT`

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
| 暂停/恢复 | `DEFAULT_ADMIN_ROLE`（多签） |
| 更新参数（APY上限、最小存款等） | `DEFAULT_ADMIN_ROLE` |
| **提取 TVL 主池** | **❌ 任何角色均不可直接提取** |

### 12.2 关键约束

- TVL 主池资产只能通过用户赎回流程流出
- 收益注入只能通过 `depositRewards()` 路径进入，不得与主池混淆
- veCPOT 不可转让、不可出售
- Boost 叠加有上限（`MAX_BOOST_BPS = 2000`，即 20%）
- Phase 2 NFT Boost 为只读（不需要用户再次授权或 approve）

### 12.3 升级约束

- 所有核心合约使用 UUPS 模式
- 升级操作需要多签或 timelock（≥ 48小时）
- 升级时不得修改现有存储布局（`_gap` 占位）

---

## 附录：原型与合约的参数对应

| 原型参数 | 合约参数 | 说明 |
|---------|---------|------|
| `BASE_APY = 8.0%` | `baseAPRBps = 800` | 初始基础年化 |
| `DURATION_BOOST[30] = 1.0%` | `durationBoostBps[30] = 100` | 30天锁仓加速 |
| `DURATION_BOOST[90] = 2.0%` | `durationBoostBps[90] = 200` | 90天锁仓加速 |
| `DURATION_BOOST[180] = 3.5%` | `durationBoostBps[180] = 350` | 180天锁仓加速 |
| `DURATION_BOOST[360] = 5.0%` | `durationBoostBps[360] = 500` | 360天锁仓加速 |
| NFT boost `+1.5%`（原型简化） | `nftBoostBps[L2] = 100, [L3] = 200` | 按等级精确计算 |
| `ceUSDT` | 份额凭证 ERC20 token | 名称待定 |
