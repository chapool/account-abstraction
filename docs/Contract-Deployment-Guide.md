# 合约部署指南

本文档提供完整的合约部署流程，适用于 BNB Chain 生态系统（BSC 和 opBNB）的所有网络。

## 📋 目录

- [前置要求](#前置要求)
- [环境配置](#环境配置)
- [网络配置](#网络配置)
- [部署流程](#部署流程)
- [合约部署顺序](#合约部署顺序)
- [部署后配置](#部署后配置)
- [合约验证](#合约验证)
- [升级合约](#升级合约)
- [故障排除](#故障排除)

## 前置要求

### 1. 安装依赖

```bash
# 克隆仓库
git clone <repository-url>
cd account-abstraction

# 安装依赖
yarn install
# 或
npm install
```

### 2. 编译合约

```bash
yarn compile
# 或
npx hardhat compile
```

### 3. 准备部署账户

确保部署账户有足够的 BNB 用于支付 Gas 费用：

- **BSC Testnet**: 从 [BSC Testnet Faucet](https://testnet.bnbchain.org/faucet-smart) 获取测试币
- **opBNB Testnet**: 从 [opBNB Testnet Faucet](https://opbnb-testnet-bridge.bnbchain.org/) 获取测试币
- **Mainnet**: 确保账户有足够的 BNB

## 环境配置

### 1. 创建环境变量文件

根据目标网络创建对应的环境变量文件：

#### opBNB Testnet (`.env.opBNBTestnet`)

```env
# ============================================
# opBNB Testnet Environment Configuration
# ============================================
# Network: opBNB Testnet
# Chain ID: 5611
# Block Explorer: https://opbnb-testnet.bscscan.com
# RPC URL: https://opbnb-testnet-rpc.bnbchain.org

# ============================================
# Network RPC URLs
# ============================================
OPBNB_TESTNET_RPC_URL=https://opbnb-testnet-rpc.bnbchain.org
# 或使用备用 RPC
# OPBNB_TESTNET_RPC_URL=https://opbnb-testnet.nodereal.io/v1/YOUR_API_KEY

# ============================================
# Account Configuration
# ============================================
# 使用私钥（二选一）
PRIVATE_KEY=0x你的私钥

# 或使用助记词（二选一）
# MNEMONIC=your twelve word mnemonic phrase here
# ACCOUNT_INDEX=0

# ============================================
# Gas Price Configuration
# ============================================
OPBNB_TESTNET_GAS_PRICE=1000000000  # 1 gwei

# ============================================
# Network Timeout
# ============================================
NETWORK_TIMEOUT=60000

# ============================================
# Deployment Configuration
# ============================================
SALT=0x0a59dbff790c23c976a548690c27297883cc66b4c67024f9117b0238995e35e9

# EntryPoint 地址（必需）
# 如果网络已有 EntryPoint，使用现有地址
# 否则需要先部署 EntryPoint
ENTRYPOINT=0x44484D3380A8D0C7555D9a4d2064db14f6A24FBb

# ============================================
# Environment Mode
# ============================================
NODE_ENV=development
HARDHAT_NETWORK=opbnbTestnet
NETWORK=opbnbTestnet

# ============================================
# Optional: GasPaymaster Configuration
# ============================================
# 设置为 'true' 启用代币燃烧模式
# GP_BURN=false
```

#### BSC Testnet (`.env.bnbTestnet`)

```env
# ============================================
# BSC Testnet Environment Configuration
# ============================================
BSC_TESTNET_RPC_URL=https://bsc-testnet.nodereal.io/v1/YOUR_API_KEY
# 或
BNB_TESTNET_RPC_URL=https://bsc-testnet.nodereal.io/v1/YOUR_API_KEY

PRIVATE_KEY=0x你的私钥
# 或
# MNEMONIC=your twelve word mnemonic phrase here
# ACCOUNT_INDEX=0

BSC_TESTNET_GAS_PRICE=20000000000  # 20 gwei
NETWORK_TIMEOUT=60000
ENTRYPOINT=0xYourEntryPointAddressHere
```

#### opBNB Mainnet (`.env.opBNB`)

```env
# ============================================
# opBNB Mainnet Environment Configuration
# ============================================
OPBNB_RPC_URL=https://opbnb-mainnet.nodereal.io/v1/YOUR_API_KEY
# 或
OPBNB_MAINNET_RPC_URL=https://opbnb-mainnet.nodereal.io/v1/YOUR_API_KEY

PRIVATE_KEY=0x你的私钥
# 或
# MNEMONIC=your twelve word mnemonic phrase here
# ACCOUNT_INDEX=0

OPBNB_GAS_PRICE=1000000000  # 1 gwei
NETWORK_TIMEOUT=60000
ENTRYPOINT=0xYourEntryPointAddressHere
```

### 2. 安全注意事项

⚠️ **重要**：
- 永远不要将包含真实私钥的 `.env` 文件提交到 Git
- 使用 `.gitignore` 确保环境文件被忽略
- 生产环境使用硬件钱包或安全的密钥管理服务
- 测试网络可以使用测试账户

## 网络配置

### 支持的网络

| 网络 | Chain ID | RPC URL | Gas Price | 区块浏览器 |
|------|----------|---------|-----------|------------|
| BSC Testnet | 97 | `https://bsc-testnet.nodereal.io/v1/...` | 20 Gwei | [BSCScan Testnet](https://testnet.bscscan.com) |
| opBNB Testnet | 5611 | `https://opbnb-testnet-rpc.bnbchain.org` | 1 Gwei | [opBNB Testnet](https://testnet.opbnbscan.com) |
| opBNB Mainnet | 204 | `https://opbnb-mainnet.nodereal.io/v1/...` | 1 Gwei | [opBNB Explorer](https://opbnb.bscscan.com) |

### RPC 端点选择

#### 免费公共端点
- **opBNB Testnet**: `https://opbnb-testnet-rpc.bnbchain.org`
- **BSC Testnet**: `https://bsc-testnet-rpc.bnbchain.org`

#### 需要 API Key 的端点
- **NodeReal**: `https://nodereal.io/` (推荐，稳定)
- **Alchemy**: `https://www.alchemy.com/`
- **Infura**: `https://infura.io/`

## 部署流程

### 方法一：一键部署所有核心合约（推荐）

#### 1. 部署到 opBNB Testnet

```bash
# 设置环境变量
export HARDHAT_NETWORK=opbnbTestnet
export NETWORK=opbnbTestnet

# 运行部署脚本
npx hardhat run scripts/deploy-core-opbnb-testnet.ts --network opbnbTestnet
```

#### 2. 部署到 BSC Testnet

```bash
# 设置环境变量
export HARDHAT_NETWORK=bnbTestnet
export NETWORK=bnbTestnet

# 运行部署脚本
npx hardhat run scripts/deploy-core-bnb.ts --network bnbTestnet
```

#### 3. 部署到 opBNB Mainnet

```bash
# 设置环境变量
export HARDHAT_NETWORK=opbnb
export NETWORK=opbnb

# 运行部署脚本
npx hardhat run scripts/deploy-core-opbnb.ts --network opbnb
```

### 方法二：使用 Hardhat Deploy 插件

```bash
# 部署所有合约（按 deploy/ 目录中的顺序）
npx hardhat deploy --network <network-name>

# 例如：部署到 opBNB Testnet
npx hardhat deploy --network opbnbTestnet
```

### 方法三：单独部署特定合约

#### 部署 EntryPoint

```bash
# 部署 EntryPoint 到 opBNB Testnet
npx hardhat run scripts/deploy-entrypoint-opbnb-testnet.ts --network opbnbTestnet
```

#### 部署 MockUSDT

```bash
# 部署 MockUSDT 到 BSC Testnet
npx hardhat run scripts/deploy-mockusdt-bnb-testnet.ts --network bnbTestnet
```

#### 部署 StakingConfig

```bash
# 部署 StakingConfig 到 BSC Testnet
npx hardhat run scripts/deploy-staking-config-bnb.ts --network bnbTestnet
```

## 合约部署顺序

部署脚本会自动按以下顺序部署合约，并处理依赖关系：

### 1. 基础合约（无依赖）

1. **CPPToken** - ERC20 代币合约
   - 构造函数参数：`admin`, `initialSupply`
   - 类型：普通部署

2. **StakingConfig** - 质押配置合约
   - 构造函数参数：无
   - 类型：普通部署

3. **Payment** - 支付合约
   - 构造函数参数：无
   - 类型：普通部署

4. **MockUSDT** - 测试 USDT 代币
   - 构造函数参数：无
   - 类型：普通部署

5. **BatchTransfer** - 批量转账合约
   - 构造函数参数：无
   - 类型：普通部署

### 2. UUPS 可升级代理合约

6. **GasPriceOracle** - Gas 价格预言机
   - 初始化参数：`owner`, `updateInterval`, `maxDeviation`
   - 类型：UUPS 代理

7. **MasterAggregator** - 签名聚合器
   - 初始化参数：`owner`, `initialMasters`
   - 类型：UUPS 代理

8. **SessionKeyManager** - 会话密钥管理器
   - 初始化参数：`owner`
   - 类型：UUPS 代理

9. **AccountManager** - 账户管理器
   - 初始化参数：`entryPoint`, `owner`
   - 类型：UUPS 代理
   - 依赖：EntryPoint

10. **CPNFT** - NFT 合约
    - 初始化参数：`name`, `symbol`, `baseURI`
    - 类型：UUPS 代理

11. **Staking** - 质押合约
    - 初始化参数：`cpnft`, `cppToken`, `accountManager`, `stakingConfig`, `owner`
    - 类型：UUPS 代理
    - 依赖：CPNFT, CPPToken, AccountManager, StakingConfig

12. **StakingReader** - 质押读取器
    - 初始化参数：`staking`, `stakingConfig`, `cpnft`, `owner`
    - 类型：UUPS 代理
    - 依赖：Staking, StakingConfig, CPNFT

### 3. 依赖其他合约的合约

13. **GasPaymaster** - Gas 支付主合约
    - 构造函数参数：`entryPoint`, `token`, `gasPriceOracle`, `fallbackRate`, `burnTokens`, `beneficiary`
    - 类型：普通部署
    - 依赖：EntryPoint, CPPToken, GasPriceOracle

### 部署脚本特性

- ✅ **幂等性**：可以重复运行，不会重复部署已存在的合约
- ✅ **智能检测**：自动检测合约是否已部署
- ✅ **依赖管理**：自动处理合约间的依赖关系
- ✅ **配置保存**：部署信息自动保存到 `deployments/<network>/core.json`

## 部署后配置

### 1. 权限配置

部署脚本会自动配置以下权限：

#### CPPToken 权限
- 授予 `Staking` 合约 `MINTER_ROLE`，用于铸造奖励代币
- 可选：如果启用 `GP_BURN=true`，授予 `GasPaymaster` `BURNER_ROLE`

#### 合约关联配置
- `AccountManager.setMasterAggregator()` - 设置聚合器
- `StakingConfig.setStakingContract()` - 设置质押合约
- `CPNFT.setStakingContract()` - 设置质押合约

### 2. Payment 合约白名单

部署脚本会自动将以下代币添加到 Payment 白名单：
- `MockUSDT`
- `CPPToken`

手动添加其他代币：

```bash
# 使用 Hardhat console
npx hardhat console --network opbnbTestnet

# 在 console 中
const Payment = await ethers.getContractFactory('Payment')
const payment = Payment.attach('0xPayment合约地址')
await payment.addAllowedToken('0x代币地址')
```

### 3. GasPaymaster 配置（可选）

启用代币燃烧模式：

```env
# 在 .env 文件中设置
GP_BURN=true
```

然后重新运行部署脚本，或手动配置：

```bash
npx hardhat console --network opbnbTestnet

const GasPaymaster = await ethers.getContractFactory('GasPaymaster')
const gp = GasPaymaster.attach('0xGasPaymaster地址')
const CPPToken = await ethers.getContractFactory('CPPToken')
const token = CPPToken.attach('0xCPPToken地址')

# 授予 BURNER_ROLE
await token.grantRole(4, '0xGasPaymaster地址')

# 启用燃烧模式
await gp.setTokenHandlingMode(true, ethers.constants.AddressZero)
```

## 合约验证

### 1. 使用 Hardhat Verify 插件

```bash
# 验证单个合约
npx hardhat verify --network <network> <合约地址> <构造函数参数>

# 示例：验证 Payment 合约（无构造函数参数）
npx hardhat verify --network opbnbTestnet 0x8ab3CD39295CF002103c183963e527f2536949Df

# 示例：验证 CPPToken（有构造函数参数）
npx hardhat verify --network opbnbTestnet \
  0xCPPToken地址 \
  "0x管理员地址" \
  "1000000000000000000000000000"
```

### 2. 验证 UUPS 代理合约

对于 UUPS 代理合约，需要验证实现合约：

```bash
# 1. 获取实现合约地址
npx hardhat console --network opbnbTestnet
const proxy = await ethers.getContractAt('ERC1967Proxy', '0x代理地址')
const impl = await proxy.implementation()
console.log('Implementation:', impl)

# 2. 验证实现合约
npx hardhat verify --network opbnbTestnet <实现合约地址> <初始化参数>
```

### 3. 批量验证脚本

可以创建自定义验证脚本：

```typescript
// scripts/verify-contracts.ts
import { run } from 'hardhat'

async function main() {
  const network = process.env.HARDHAT_NETWORK || 'opbnbTestnet'
  const deployments = require(`../deployments/${network}/core.json`)
  
  // 验证 Payment
  await run('verify:verify', {
    address: deployments.contracts.Payment,
    network: network
  })
  
  // 验证其他合约...
}

main().catch(console.error)
```

## 升级合约

### UUPS 可升级合约

以下合约支持 UUPS 升级：
- GasPriceOracle
- MasterAggregator
- SessionKeyManager
- AccountManager
- CPNFT
- Staking
- StakingReader

### 升级步骤

#### 1. 编译新版本

```bash
yarn compile
```

#### 2. 运行升级脚本

```bash
# 升级 Staking 合约（BSC Testnet）
npx hardhat run scripts/upgrade-staking-bnb.ts --network bnbTestnet

# 升级 CPNFT 合约（BSC Testnet）
npx hardhat run scripts/upgrade-cpnft-bnb.ts --network bnbTestnet

# 升级 Staking 合约（Sepolia）
npx hardhat run scripts/upgrade-staking-sepolia.ts --network sepoliaCustom
```

#### 3. 验证升级

```bash
# 检查实现合约地址
npx hardhat run scripts/check-implementation.ts --network bnbTestnet

# 验证升级
npx hardhat run scripts/validate-staking-upgrade.ts --network bnbTestnet
```

### 升级注意事项

⚠️ **重要**：
- 升级前备份当前实现合约地址
- 确保新版本合约兼容现有存储布局
- 在生产环境升级前，先在测试网络充分测试
- 升级后验证所有功能正常

## 故障排除

### 常见问题

#### 1. RPC 端点错误

**错误信息**：
```
Error: RPC endpoint returned HTTP client error
```

**解决方案**：
- 检查 RPC URL 是否正确
- 尝试使用备用 RPC 端点
- 检查网络连接和防火墙设置
- 确认 API Key 有效（如果使用）

#### 2. Gas 不足

**错误信息**：
```
Error: insufficient funds for gas
```

**解决方案**：
- 确保部署账户有足够的 BNB
- 检查 Gas Price 设置是否合理
- 对于测试网络，从水龙头获取测试币

#### 3. EntryPoint 未设置

**错误信息**：
```
请设置环境变量 ENTRYPOINT 为已部署的 EntryPoint 地址
```

**解决方案**：
- 如果网络已有 EntryPoint，在 `.env` 文件中设置 `ENTRYPOINT` 变量
- 如果没有，先部署 EntryPoint：
  ```bash
  npx hardhat run scripts/deploy-entrypoint-opbnb-testnet.ts --network opbnbTestnet
  ```

#### 4. 合约已部署

**现象**：脚本跳过部署，显示 "已部署"

**解决方案**：
- 这是正常行为，脚本会检测已部署的合约
- 如果需要重新部署，删除 `deployments/<network>/core.json` 中的对应地址
- 或手动删除合约（仅测试网络）

#### 5. 权限错误

**错误信息**：
```
Error: execution reverted: AccessControl: account is missing role
```

**解决方案**：
- 检查部署账户是否有正确的权限
- 运行权限配置脚本：
  ```bash
  npx hardhat run scripts/grant-minter-role.ts --network <network>
  ```

### 调试技巧

#### 1. 启用详细日志

```bash
# 设置环境变量
export DEBUG=hardhat:*

# 运行部署
npx hardhat run scripts/deploy-core-opbnb-testnet.ts --network opbnbTestnet
```

#### 2. 检查部署状态

```bash
# 查看部署记录
cat deployments/opbnbTestnet/core.json

# 检查合约代码
npx hardhat console --network opbnbTestnet
const code = await ethers.provider.getCode('0x合约地址')
console.log(code !== '0x' ? '合约已部署' : '合约未部署')
```

#### 3. 估算 Gas

```bash
# 估算部署 Gas
npx hardhat run scripts/estimate-deployment-gas.ts --network opbnbTestnet
```

#### 4. EARN 合约部署（opBNB Testnet）

在已部署 core 合约的前提下，部署 Chapool Earn 体系（金库、veCPOT 锁仓、NFT 加速、收益注入、只读聚合）：

```bash
# 确保 deployments/opbnbTestnet/core.json 存在且包含 CPPToken、AccountManager、CPNFT、MockUSDT
export HARDHAT_NETWORK=opbnbTestnet
npx hardhat run scripts/deploy-earn-opbnb-testnet.ts --network opbnbTestnet
```

部署结果写入 `deployments/opbnbTestnet/earn.json`，包含：`MockCPOT`、`ChapoolEarnVault`、`VeCPOTLocker`、`NFTBoostController`、`ChapoolRewardDistributor`、`ChapoolVaultReader`。脚本会为 `ChapoolEarnVault` 授予 CPP 的 `MINTER_ROLE`，并完成各合约间的相互引用配置。

## 部署信息存储

部署信息自动保存到以下位置：

```
deployments/
├── opbnbTestnet/
│   ├── core.json          # 核心合约部署信息
│   ├── earn.json          # EARN 合约部署信息（见 deploy-earn-opbnb-testnet.ts）
│   └── entrypoint.json    # EntryPoint 部署信息
├── bnbTestnet/
│   └── core.json
└── opbnb/
    └── core.json
```

### core.json 格式

```json
{
  "network": "opbnbTestnet",
  "chainId": 5611,
  "deployer": "0x部署者地址",
  "entryPoint": "0xEntryPoint地址",
  "timestamp": "2025-12-05T10:00:00.000Z",
  "contracts": {
    "CPPToken": "0x...",
    "GasPriceOracle": "0x...",
    "MasterAggregator": "0x...",
    "SessionKeyManager": "0x...",
    "AccountManager": "0x...",
    "GasPaymaster": "0x...",
    "CPNFT": "0x...",
    "StakingConfig": "0x...",
    "Staking": "0x...",
    "StakingReader": "0x...",
    "BatchTransfer": "0x...",
    "MockUSDT": "0x...",
    "Payment": "0x..."
  }
}
```

## 最佳实践

### 1. 部署前检查清单

- [ ] 环境变量文件已正确配置
- [ ] 部署账户有足够的 BNB
- [ ] 合约已成功编译
- [ ] 测试网络已测试所有功能
- [ ] 备份了重要配置信息

### 2. 部署流程

1. **测试网络部署**：先在测试网络部署并测试
2. **代码审查**：确保合约代码经过审查
3. **主网部署**：确认无误后部署到主网
4. **验证合约**：部署后立即验证合约代码
5. **配置权限**：设置必要的权限和关联关系
6. **测试功能**：验证所有功能正常工作

### 3. 安全建议

- 使用多签钱包管理主网部署
- 定期备份部署信息
- 记录所有部署和升级操作
- 使用版本控制管理部署脚本
- 生产环境使用硬件钱包

## 相关文档

- [BNB Chain 部署指南](./BNB-Chain-Deployment.md) - BNB Chain 特定部署信息
- [合约升级指南](./CPNFT-Upgradeable-Guide.md) - UUPS 升级详细说明
- [Payment 合约使用](./Payment-Usage-Examples.md) - Payment 合约使用示例
- [Staking 系统文档](../contracts/CPNFT/README.md) - Staking 系统完整文档

## 支持

如遇到问题：

1. 查看本文档的故障排除部分
2. 检查 [GitHub Issues](https://github.com/your-repo/issues)
3. 参考 [BNB Chain 文档](https://docs.bnbchain.org/)
4. 联系开发团队

---

**最后更新**: 2025-12-05  
**版本**: 1.0.0
