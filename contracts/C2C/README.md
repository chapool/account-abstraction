# C2C NFT 交易市场合约

## 📋 概述

C2C NFT 交易市场是一个基于 Web3 的去中心化 NFT 交易平台，支持一口价和拍卖两种交易模式。合约采用"第三方合约托管"机制保障交易安全，全面采用 USDT（或任意 ERC20）作为计价单位。

## 🏗️ 系统架构

### 核心合约

```
contracts/C2C/
├── Marketplace.sol          # 主交易合约（可升级）
├── interfaces/
│   └── IMarketplace.sol     # 接口定义
└── README.md                # 本文档
```

### 架构特点

- **半中心化设计**：参考 `Staking.sol` 设计模式，使用 `onlyOwner` 修饰符，允许后端平台代用户执行操作
- **UUPS 可升级**：支持合约升级，便于后续功能迭代
- **安全机制**：集成 ReentrancyGuard 和 Pausable，防止重入攻击和紧急暂停
- **托管模式**：NFT 和资金由合约托管，保障交易安全
- **Gas 优化**：使用结构体打包存储，支持批量操作

## 🚀 核心功能

### 1. 上架功能 (createListing)

支持两种上架模式：

- **一口价 (FixedPrice)**：设置固定价格，买家可直接购买
- **拍卖 (Auction)**：设置起拍价、拍卖时长和最小加价幅度

**流程**：
1. 用户在前端选择 NFT 并设置参数
2. 用户授权 NFT 给合约
3. 后端调用 `createListing()` 代用户执行
4. NFT 从用户钱包转移到合约托管

### 2. 下架功能 (cancelListing)

- 检查下架频次限制（24小时内最多3次）
- NFT 从合约转回卖家钱包
- 如果是拍卖，自动退还所有出价者的 USDT

### 3. 购买功能 (buyItem)

一口价商品的购买流程：

1. 买家在前端点击"立即购买"
2. 买家授权 USDT 给合约
3. 后端调用 `buyItem()` 代用户执行
4. 合约自动计算平台抽成（5%）
5. 转账给卖家（95%）和平台（5%）
6. NFT 转移给买家

### 4. 竞拍功能 (placeBid)

拍卖商品的出价流程：

1. 买家输入出价金额（需高于当前最高价 + 最小步长）
2. 买家授权 USDT 给合约
3. 后端调用 `placeBid()` 代用户执行
4. 新出价者的 USDT 被锁定
5. 自动退还前一个最高出价者的 USDT
6. 记录出价历史

### 5. 拍卖结算 (settleAuction)

拍卖结束后的结算流程：

1. 卖家点击"立即成交"或买家点击"领取"
2. 后端调用 `settleAuction()` 代用户执行
3. 合约计算平台抽成（5%）
4. 转账给卖家（95%）和平台（5%）
5. NFT 转移给买家

## 💰 商业规则

### 平台抽成

- **费率**：默认 5%（500 basis points）
- **扣费时机**：交易结算时自动执行
- **计算公式**：卖家实收 = 成交价 × (1 - 5%)
- **示例**：100 USDT 成交 → 卖家收到 95 USDT，平台收到 5 USDT

### 下架频次限制

- **限制规则**：每个用户地址在 24 小时内最多执行 3 次下架操作
- **重置机制**：超过 24 小时后自动重置计数
- **提示机制**：前端需提示用户剩余下架次数

## 📝 合约部署

### 初始化参数

```solidity
function initialize(
    address _nftContract,              // CPNFT 合约地址
    address _paymentToken,             // USDT 或任意 ERC20 合约地址
    address _platformFeeRecipient,     // 平台抽成接收地址
    uint256 _platformFeeRate,         // 平台费率（500 = 5%）
    uint256 _delistingLimit,          // 下架限制次数（默认：3）
    uint256 _delistingWindow,          // 下架时间窗口（默认：86400 = 24小时）
    address _owner                     // 合约所有者地址
)
```

### 部署步骤

1. 部署 Marketplace 实现合约
2. 部署 UUPS 代理合约，指向实现合约
3. 调用 `initialize()` 初始化合约
4. 配置 NFT 合约和支付代币地址
5. 设置平台抽成接收地址

## 🔧 管理员功能

### 配置更新

- `updatePlatformFeeRecipient()` - 更新平台抽成接收地址
- `updatePlatformFeeRate()` - 更新平台费率
- `updateDelistingLimit()` - 更新下架限制次数
- `updateDelistingWindow()` - 更新下架时间窗口
- `updateNFTContract()` - 更新 NFT 合约地址
- `updatePaymentToken()` - 更新支付代币地址

### 紧急功能

- `pause()` / `unpause()` - 暂停/恢复合约
- `emergencyWithdraw()` - 紧急提取 ERC20 代币
- `emergencyWithdrawNFT()` - 紧急提取 ERC721 NFT

### 测试模式

- `enableTestMode()` - 启用测试模式（时间控制）
- `disableTestMode()` - 禁用测试模式
- `setTestTimestamp()` - 设置测试时间戳

## 📊 查询功能

### 基础查询

- `getListing(uint256 listingId)` - 获取订单详情
- `getBids(uint256 listingId)` - 获取所有出价记录
- `getHighestBid(uint256 listingId)` - 获取最高出价
- `getUserListings(address user)` - 获取用户的订单列表

### 下架记录查询

- `getDelistingRecord(address user)` - 获取用户下架记录
  - 返回：当前计数、上次重置时间、剩余次数

## 🔐 安全特性

1. **重入保护**：使用 ReentrancyGuard 防止重入攻击
2. **暂停机制**：支持紧急暂停所有交易
3. **权限控制**：关键操作仅限 owner（后端）执行
4. **输入验证**：所有输入参数都经过严格验证
5. **状态检查**：操作前检查订单状态和有效性
6. **资金安全**：使用 SafeERC20 和 safeTransferFrom 确保资金安全

## 📈 Gas 优化

1. **结构体打包**：优化存储布局，减少存储槽使用
2. **批量操作**：支持批量下架等操作（未来扩展）
3. **事件索引**：关键字段使用 indexed，便于链下查询

## 🧪 测试建议

### 测试场景

1. **上架测试**
   - 一口价上架
   - 拍卖上架
   - 无效参数测试

2. **购买测试**
   - 正常购买流程
   - 余额不足测试
   - 重复购买测试

3. **竞拍测试**
   - 首次出价
   - 加价出价
   - 出价不足测试
   - 拍卖结束后出价测试

4. **下架测试**
   - 正常下架
   - 下架频次限制测试
   - 拍卖下架退款测试

5. **平台抽成测试**
   - 抽成计算准确性
   - 不同费率测试

## 📚 参考实现

本合约参考了以下设计模式：

- `contracts/CPNFT/Staking.sol` - 半中心化批量操作模式
- OpenZeppelin Contracts - 安全标准和最佳实践

## 🔄 版本历史

- **v1.0.0** - 初始版本
  - 支持一口价和拍卖两种模式
  - 实现平台抽成机制
  - 实现下架频次限制
  - 支持 UUPS 可升级

## 📞 支持

如有问题或建议，请联系开发团队。
