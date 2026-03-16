# Chapool Earn 前端开发文档（opBNB Testnet）

> 面向 **frontend-prototype** 及 Earn 产品前端，基于 opBNB Testnet 已部署合约。  
> 参考原型：`frontend-prototype/src/App.tsx`；交互与文案规范：`docs/Earn-Feature-Spec.md` §9。

---

## 1. 网络与链配置

| 项目 | 值 |
|------|-----|
| 网络名称 | opBNB Testnet |
| Chain ID | **5611** |
| RPC URL | `https://opbnb-testnet.nodereal.io/v1/e59b1bd6360a4cdf9a39530a77e815ca` 或环境变量 `OPBNB_TESTNET_RPC_URL` |
| 区块浏览器 | https://opbnb-testnet.bscscan.com |

前端连接钱包或 Provider 时请使用 **chainId: 5611**。

---

## 2. 合约地址（opBNB Testnet）

### 2.1 核心合约（core）

来自 `deployments/opbnbTestnet/core.json`：

| 合约 | 地址 | 说明 |
|------|------|------|
| MockUSDT | `0x1950cDE2DECb98Ee93a1D636fA923Fe8a3f09094` | 测试网 USDT，Earn 存款资产 |
| CPPToken | `0x2AB38D1fc6Dd71d79AEa3b279F25965267383393` | CPP 代币，领取奖励用 |
| AccountManager | `0x043FBF4A20540a67200A9f084a8E27Cff0167519` | 解析用户 AA 账户 |
| CPNFT | `0x27593d4Fed6c4B812Ae8e2B64e38d8B039c8Fb91` | NFT 合约，用于 NFT 加成 |

### 2.2 EARN 合约（earn）

来自 `deployments/opbnbTestnet/earn.json`：

| 合约 | 地址 | 说明 |
|------|------|------|
| **ChapoolVaultReader** | `0x5895A836679AaFeD205cD6689627aC8160860b82` | **前端主读入口**，聚合用户/协议数据 |
| ChapoolEarnVault | `0xe57E96e423306847990877b8334BDB711efdfD10` | 存款/取款/领取 CPP |
| VeCPOTLocker | `0xfF226A6D5A8F3Ff5E621cD9C1564310beC65509f` | CPOT 锁仓、veCPOT 加成 |
| NFTBoostController | `0xC6366FbE6A7C1a0b7053EDc3E91232Da6d884946` | NFT 加成注册与查询 |
| ChapoolRewardDistributor | `0x385D801f556e676bd22e40B8bd1388C46833EC3a` | 运营设置 rewardRate（前端只读） |
| MockCPOT | `0xDECC4966A5fe63fF0c8f8545AEAB025390195b5d` | 测试网治理代币，锁仓用 |

---

## 3. 前端主读：ChapoolVaultReader

建议**仅通过 ChapoolVaultReader** 拉取用户与协议数据，减少 RPC 次数。

### 3.1 用户仪表盘（一屏数据）

**调用**：`ChapoolVaultReader.getUserDashboard(userAddress)`

**返回结构**（与合约 `UserDashboard` 一致）：

```ts
interface UserDashboard {
  usdtBalance: bigint;        // 存款本金（USDT，18 位小数）
  weightedUSDT: bigint;       // 加权 USDT（含 veCPOT + NFT 加成）
  pendingCPP: bigint;         // 可领取 CPP
  estimatedDailyCPP: bigint; // 按当前速率估算每日 CPP
  boostBps: bigint;           // 总加成（bps，10000 = 100%）
  vecpotBoostBps: bigint;     // veCPOT 加成 bps
  nftBoostBps: bigint;        // NFT 加成 bps
  totalVeCPOT: bigint;        // 当前 veCPOT 总量（用于展示）
  lockedCPOT: bigint;         // 锁仓中的 CPOT 总量
  depositedAt: bigint;        // 首次存款时间戳
  lastActionAt: bigint;       // 最近一次操作时间戳
}
```

用于：我的存款、待领取 CPP、预计日收益、加速倍率、锁仓状态等。

### 3.2 协议概览（TVL / 年化参考）

**调用**：`ChapoolVaultReader.getProtocolOverview()`

**返回结构**：

```ts
interface ProtocolOverview {
  tvl: bigint;                // 金库 USDT 余额（TVL）
  totalWeightedUSDT: bigint;  // 全网加权 USDT 合计
  totalUsers: bigint;         // 存款用户数
  rewardRate: bigint;         // 全网每秒 CPP 排放（18 位小数）
  dailyCPPEmission: bigint;   // rewardRate × 86400
  emergencyMode: boolean;     // 是否紧急模式
}
```

用于：首页 TVL、基础年化估算、是否展示紧急提示等。

### 3.3 其他只读（按需）

- `ChapoolVaultReader.getTVL()`：仅返回 TVL（uint256），适配 DefiLlama 等。
- 锁仓列表：`ChapoolVaultReader.getUserLockPositions(user)`，返回 `LockPosition[]`。
- NFT 激活信息：`NFTBoostController.getActiveNFT(user)`，返回 `(tokenId, level, boostBps)`。

---

## 4. 用户写操作（合约直调）

以下由用户钱包对**对应合约**发起交易（需先连接 chainId 5611）。

### 4.1 存款 / 取款 / 领 CPP

| 操作 | 合约 | 方法 | 说明 |
|------|------|------|------|
| 存入 USDT | ChapoolEarnVault | `deposit(amount, receiver)` | 先对 Vault 授权 MockUSDT |
| 取出 USDT | ChapoolEarnVault | `withdraw(amount, receiver)` | receiver 通常为 user |
| 领取 CPP | ChapoolEarnVault | `claimCPP()` | CPP 会 mint 到用户 AA 账户 |

合约地址：`0xe57E96e423306847990877b8334BDB711efdfD10`（ChapoolEarnVault）。

### 4.2 锁仓 CPOT（veCPOT）

| 操作 | 合约 | 方法 |
|------|------|------|
| 锁仓 | VeCPOTLocker | `lock(amount, durationDays)`，durationDays 为 30 / 90 / 180 / 360 |
| 解锁 | VeCPOTLocker | `unlock(lockId)` |
| 锁仓列表 | 通过 ChapoolVaultReader.getUserLockPositions(user) 或 VeCPOTLocker.getLockPositions(user) |

合约地址：`0xfF226A6D5A8F3Ff5E621cD9C1564310beC65509f`（VeCPOTLocker）。  
资产：MockCPOT `0xDECC4966A5fe63fF0c8f8545AEAB025390195b5d`，需先 approve。

### 4.3 NFT 加成

| 操作 | 合约 | 方法 |
|------|------|------|
| 激活 | NFTBoostController | `activateBoost(tokenId)`，需持有该 CPNFT 且未质押 |
| 取消激活 | NFTBoostController | `deactivateBoost()` |

合约地址：`0xC6366FbE6A7C1a0b7053EDc3E91232Da6d884946`（NFTBoostController）。  
激活/取消后建议调用 **ChapoolEarnVault.syncBoost(user)** 以更新金库内权重（或由合约内部/后端触发）。

---

## 5. 数据流与展示建议

1. **首屏**：`getProtocolOverview()` + `getUserDashboard(userAddress)`，渲染 TVL、用户存款、待领 CPP、预计日收益、加成状态。
2. **存款/取款 Sheet**：输入金额后可按需调用 Vault 的 view（如 `estimatedDailyCPP`）做预览；提交时调 `deposit` / `withdraw`。
3. **锁仓 Sheet**：档位 30/90/180/360 天对应不同加成；可调 `VeCPOTLocker.getBoostBps(user)` 或 Reader 的 dashboard 展示当前/预览加成。
4. **领取 CPP**：展示 `pendingCPP`，按钮触发 `claimCPP()`，成功后刷新 dashboard。
5. **收益规则说明**：文案见 `frontend-prototype/src/copy/earnRewardRules.ts`，中英文已就绪。

---

## 6. 历史记录（待开发）

> **说明**：**历史记录**（存入/取出/锁仓/解锁/领取 CPP 等操作记录）**尚未开发完成**。  
> 链上已有部分事件（如 `Deposited`、`Withdrawn`、`CPPRewardClaimed`、`LockedCPOT`、`UnlockedCPOT`），但**统一的历史列表接口或后端聚合接口尚未就绪**。  
> 前端可先做**占位展示**或**本地 mock 列表**，待历史接口定义后再对接。

---

## 7. 参考

- 合约与交互规范：`docs/Earn-Feature-Spec.md`
- 产品设计：`docs/Chapool-TVL-Product-Design.md`
- 部署信息：`deployments/opbnbTestnet/core.json`、`deployments/opbnbTestnet/earn.json`
- 收益规则文案：`frontend-prototype/src/copy/earnRewardRules.ts`
