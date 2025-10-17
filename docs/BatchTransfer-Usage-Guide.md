# BatchTransfer 使用指南

## 概述

BatchTransfer 是一个用于批量转移 NFT（ERC721）和 FT 同质化代币（ERC20）的智能合约。它提供了多种批量操作功能，可以大幅降低 gas 费用并提高操作效率。

## 合约地址

**Sepolia 测试网**: `0x29aAD71f97Da6AA0F4096cfd50DA395eAa18D964`

- **网络**: Sepolia Testnet (Chain ID: 11155111)
- **部署者**: `0xa3B605fB633AD0A0DC4B74b10bBfc40fDB050d35`
- **交易哈希**: `0x4ca33acf2b57787c5d2a8e940b95978e725bbf9ffe8c75f8d93e0c7b7834e26b`
- **区块浏览器**: [查看合约](https://sepolia.etherscan.io/address/0x29aAD71f97Da6AA0F4096cfd50DA395eAa18D964)

## 快速开始

### 使用已部署的合约

在 Sepolia 测试网上，你可以直接使用已部署的 BatchTransfer 合约：

```javascript
const { ethers } = require("ethers");

// 连接到 Sepolia 测试网
const provider = new ethers.providers.JsonRpcProvider("YOUR_SEPOLIA_RPC_URL");
const wallet = new ethers.Wallet("YOUR_PRIVATE_KEY", provider);

// BatchTransfer 合约地址
const BATCH_TRANSFER_ADDRESS = "0x29aAD71f97Da6AA0F4096cfd50DA395eAa18D964";

// BatchTransfer ABI (仅列出常用方法)
const BATCH_TRANSFER_ABI = [
  "function batchTransferNFT(address nftContract, address[] recipients, uint256[] tokenIds)",
  "function batchTransferNFTToSingle(address nftContract, address to, uint256[] tokenIds)",
  "function batchTransferToken(address tokenContract, address[] recipients, uint256[] amounts)",
  "function batchTransferTokenEqual(address tokenContract, address[] recipients, uint256 amountPerRecipient)",
  "function combinedBatchTransfer(address nftContract, address[] nftRecipients, uint256[] tokenIds, address tokenContract, address[] tokenRecipients, uint256[] tokenAmounts)",
  "function isNFTApproved(address nftContract, address owner) view returns (bool)",
  "function hasTokenAllowance(address tokenContract, address owner, uint256 amount) view returns (bool)"
];

// 创建合约实例
const batchTransfer = new ethers.Contract(
  BATCH_TRANSFER_ADDRESS,
  BATCH_TRANSFER_ABI,
  wallet
);

// 使用示例
async function example() {
  // 假设你的 NFT 合约地址
  const nftAddress = "YOUR_NFT_CONTRACT_ADDRESS";
  
  // 1. 授权 BatchTransfer 管理你的 NFT
  const nft = new ethers.Contract(
    nftAddress,
    ["function setApprovalForAll(address operator, bool approved)"],
    wallet
  );
  await nft.setApprovalForAll(BATCH_TRANSFER_ADDRESS, true);
  
  // 2. 批量转移 NFT
  await batchTransfer.batchTransferNFT(
    nftAddress,
    ["0xRecipient1...", "0xRecipient2..."],
    [1, 2] // tokenIds
  );
}
```

## 主要功能

### 1. NFT 批量转移功能

#### 1.1 批量转移 NFT 给多个接收者

将多个 NFT 分别转移给不同的接收者。

```solidity
function batchTransferNFT(
    address nftContract,
    address[] calldata recipients,
    uint256[] calldata tokenIds
) external
```

**参数说明：**
- `nftContract`: ERC721 合约地址
- `recipients`: 接收者地址数组
- `tokenIds`: 要转移的 NFT token ID 数组（长度必须与 recipients 相同）

**使用示例：**

```javascript
// 1. 授权 BatchTransfer 管理你的 NFT
await cpnft.setApprovalForAll(batchTransferAddress, true);

// 2. 批量转移 NFT
await batchTransfer.batchTransferNFT(
  cpnftAddress,
  [user1.address, user2.address, user3.address],
  [1, 2, 3]  // tokenIds
);
```

#### 1.2 批量转移多个 NFT 给单个接收者

将多个 NFT 转移给同一个接收者。

```solidity
function batchTransferNFTToSingle(
    address nftContract,
    address to,
    uint256[] calldata tokenIds
) external
```

**使用示例：**

```javascript
await batchTransfer.batchTransferNFTToSingle(
  cpnftAddress,
  recipientAddress,
  [1, 2, 3, 4, 5]  // 转移 5 个 NFT 给同一个地址
);
```

#### 1.3 批量转移来自不同集合的 NFT

将来自不同 NFT 集合的代币转移给单个接收者。

```solidity
function batchTransferMultipleNFTCollections(
    address[] calldata nftContracts,
    address to,
    uint256[] calldata tokenIds
) external
```

**使用示例：**

```javascript
await batchTransfer.batchTransferMultipleNFTCollections(
  [collection1Address, collection2Address, collection3Address],
  recipientAddress,
  [tokenId1, tokenId2, tokenId3]
);
```

### 2. ERC20 代币批量转移功能

#### 2.1 批量转移代币（不同金额）

向多个接收者转移不同数量的代币。

```solidity
function batchTransferToken(
    address tokenContract,
    address[] calldata recipients,
    uint256[] calldata amounts
) external
```

**使用示例：**

```javascript
// 1. 授权代币
const totalAmount = ethers.utils.parseEther("1000");
await token.approve(batchTransferAddress, totalAmount);

// 2. 批量转移
await batchTransfer.batchTransferToken(
  tokenAddress,
  [user1.address, user2.address, user3.address],
  [
    ethers.utils.parseEther("100"),
    ethers.utils.parseEther("200"),
    ethers.utils.parseEther("700")
  ]
);
```

#### 2.2 批量转移代币（相同金额）

向多个接收者转移相同数量的代币。

```solidity
function batchTransferTokenEqual(
    address tokenContract,
    address[] calldata recipients,
    uint256 amountPerRecipient
) external
```

**使用示例：**

```javascript
// 每人发送 100 代币
await batchTransfer.batchTransferTokenEqual(
  tokenAddress,
  [user1.address, user2.address, user3.address],
  ethers.utils.parseEther("100")
);
```

#### 2.3 批量转移多种代币

将多种 ERC20 代币转移给单个接收者。

```solidity
function batchTransferMultipleTokens(
    address[] calldata tokenContracts,
    address to,
    uint256[] calldata amounts
) external
```

**使用示例：**

```javascript
// 转移 USDT 和 USDC
await batchTransfer.batchTransferMultipleTokens(
  [usdtAddress, usdcAddress],
  recipientAddress,
  [
    ethers.utils.parseUnits("100", 6),  // USDT (6 decimals)
    ethers.utils.parseUnits("200", 6)   // USDC (6 decimals)
  ]
);
```

### 3. 组合批量转移功能

#### 3.1 同时转移 NFT 和代币

在单个交易中同时转移 NFT 和 ERC20 代币。

```solidity
function combinedBatchTransfer(
    address nftContract,
    address[] calldata nftRecipients,
    uint256[] calldata tokenIds,
    address tokenContract,
    address[] calldata tokenRecipients,
    uint256[] calldata tokenAmounts
) external
```

**使用示例：**

```javascript
// 1. 授权 NFT 和代币
await cpnft.setApprovalForAll(batchTransferAddress, true);
await token.approve(batchTransferAddress, totalAmount);

// 2. 组合转移
await batchTransfer.combinedBatchTransfer(
  cpnftAddress,
  [user1.address, user2.address],  // NFT 接收者
  [1, 2],                           // NFT tokenIds
  tokenAddress,
  [user1.address, user2.address],  // 代币接收者
  [
    ethers.utils.parseEther("100"),
    ethers.utils.parseEther("200")
  ]
);
```

**只转移 NFT（跳过代币）：**

```javascript
await batchTransfer.combinedBatchTransfer(
  cpnftAddress,
  [user1.address, user2.address],
  [1, 2],
  ethers.constants.AddressZero,  // 设为 0 地址跳过代币转移
  [],
  []
);
```

**只转移代币（跳过 NFT）：**

```javascript
await batchTransfer.combinedBatchTransfer(
  ethers.constants.AddressZero,  // 设为 0 地址跳过 NFT 转移
  [],
  [],
  tokenAddress,
  [user1.address, user2.address],
  [amount1, amount2]
);
```

### 4. 查询功能

#### 4.1 检查 NFT 授权状态

```solidity
function isNFTApproved(address nftContract, address owner) external view returns (bool)
```

**使用示例：**

```javascript
const isApproved = await batchTransfer.isNFTApproved(cpnftAddress, userAddress);
console.log("NFT已授权:", isApproved);
```

#### 4.2 检查代币授权额度

```solidity
function hasTokenAllowance(
    address tokenContract,
    address owner,
    uint256 amount
) external view returns (bool)
```

**使用示例：**

```javascript
const amount = ethers.utils.parseEther("1000");
const hasAllowance = await batchTransfer.hasTokenAllowance(
  tokenAddress,
  userAddress,
  amount
);
console.log("代币授权充足:", hasAllowance);
```

## 完整使用流程

### 场景 1: 空投 NFT 给多个用户

```javascript
const { ethers } = require("hardhat");

async function airdropNFTs() {
  const [owner] = await ethers.getSigners();
  
  const batchTransfer = await ethers.getContractAt("BatchTransfer", BATCH_TRANSFER_ADDRESS);
  const cpnft = await ethers.getContractAt("CPNFT", CPNFT_ADDRESS);

  // 准备接收者列表和 tokenIds
  const recipients = [
    "0xRecipient1...",
    "0xRecipient2...",
    "0xRecipient3...",
    // ... 更多接收者
  ];
  
  const tokenIds = [1, 2, 3, /* ... */];

  // 1. 检查授权状态
  const isApproved = await batchTransfer.isNFTApproved(CPNFT_ADDRESS, owner.address);
  
  if (!isApproved) {
    console.log("授权 BatchTransfer...");
    const tx = await cpnft.setApprovalForAll(BATCH_TRANSFER_ADDRESS, true);
    await tx.wait();
    console.log("✅ 授权完成");
  }

  // 2. 批量转移 NFT
  console.log(`开始空投 ${tokenIds.length} 个 NFT...`);
  const tx = await batchTransfer.batchTransferNFT(
    CPNFT_ADDRESS,
    recipients,
    tokenIds
  );
  
  const receipt = await tx.wait();
  console.log(`✅ 空投完成! Gas used: ${receipt.gasUsed.toString()}`);
}

airdropNFTs();
```

### 场景 2: 发放代币奖励给多个用户

```javascript
async function distributeRewards() {
  const [owner] = await ethers.getSigners();
  
  const batchTransfer = await ethers.getContractAt("BatchTransfer", BATCH_TRANSFER_ADDRESS);
  const token = await ethers.getContractAt("IERC20", TOKEN_ADDRESS);

  const recipients = [
    "0xWinner1...",
    "0xWinner2...",
    "0xWinner3..."
  ];
  
  const rewards = [
    ethers.utils.parseEther("1000"),  // 第一名
    ethers.utils.parseEther("500"),   // 第二名
    ethers.utils.parseEther("250")    // 第三名
  ];

  // 计算总金额
  const totalAmount = rewards.reduce((a, b) => a.add(b), ethers.BigNumber.from(0));

  // 1. 检查授权
  const hasAllowance = await batchTransfer.hasTokenAllowance(
    TOKEN_ADDRESS,
    owner.address,
    totalAmount
  );

  if (!hasAllowance) {
    console.log("授权代币...");
    const tx = await token.approve(BATCH_TRANSFER_ADDRESS, totalAmount);
    await tx.wait();
    console.log("✅ 授权完成");
  }

  // 2. 批量转移代币
  console.log("开始发放奖励...");
  const tx = await batchTransfer.batchTransferToken(
    TOKEN_ADDRESS,
    recipients,
    rewards
  );
  
  await tx.wait();
  console.log("✅ 奖励发放完成!");
}

distributeRewards();
```

### 场景 3: 同时发放 NFT 和代币

```javascript
async function distributeMixedRewards() {
  const [owner] = await ethers.getSigners();
  
  const batchTransfer = await ethers.getContractAt("BatchTransfer", BATCH_TRANSFER_ADDRESS);
  const cpnft = await ethers.getContractAt("CPNFT", CPNFT_ADDRESS);
  const token = await ethers.getContractAt("IERC20", TOKEN_ADDRESS);

  const recipients = [
    "0xWinner1...",
    "0xWinner2...",
    "0xWinner3..."
  ];

  // 1. 授权 NFT
  const isNFTApproved = await batchTransfer.isNFTApproved(CPNFT_ADDRESS, owner.address);
  if (!isNFTApproved) {
    await (await cpnft.setApprovalForAll(BATCH_TRANSFER_ADDRESS, true)).wait();
  }

  // 2. 授权代币
  const totalTokens = ethers.utils.parseEther("1750");
  const hasTokenAllowance = await batchTransfer.hasTokenAllowance(
    TOKEN_ADDRESS,
    owner.address,
    totalTokens
  );
  if (!hasTokenAllowance) {
    await (await token.approve(BATCH_TRANSFER_ADDRESS, totalTokens)).wait();
  }

  // 3. 组合转移
  console.log("开始发放奖励包（NFT + 代币）...");
  const tx = await batchTransfer.combinedBatchTransfer(
    CPNFT_ADDRESS,
    recipients,
    [1, 2, 3],  // NFT tokenIds
    TOKEN_ADDRESS,
    recipients,
    [
      ethers.utils.parseEther("1000"),
      ethers.utils.parseEther("500"),
      ethers.utils.parseEther("250")
    ]
  );
  
  await tx.wait();
  console.log("✅ 混合奖励发放完成!");
}

distributeMixedRewards();
```

## Gas 优化建议

1. **使用相同金额转账函数**: 当向多个地址发送相同数量的代币时，使用 `batchTransferTokenEqual()` 可以节省 gas。

2. **组合操作**: 使用 `combinedBatchTransfer()` 可以在单个交易中完成 NFT 和代币转移，节省交易费用。

3. **一次授权，多次使用**: 授权操作只需执行一次，后续可以多次调用转移函数。

4. **批量操作**: 批量操作比逐个转移可以节省约 20-40% 的 gas 费用。

## 常见应用场景

- ✅ **空投活动**: 向社区成员批量发放 NFT 或代币
- ✅ **奖励发放**: 向活动获胜者分发奖品
- ✅ **团队分配**: 向团队成员分发项目代币
- ✅ **资产迁移**: 将资产从一个钱包批量转移到另一个钱包
- ✅ **游戏奖励**: 向游戏玩家批量发放游戏道具（NFT）和代币
- ✅ **投资组合管理**: 批量管理和转移投资资产

## 安全注意事项

⚠️ **重要提醒:**

1. **检查授权状态**: 在批量操作前，始终检查授权状态
2. **验证接收地址**: 确保接收地址正确，避免转错地址
3. **确认所有权**: 确保你拥有所有要转移的 NFT
4. **检查余额和授权额度**: 确保有足够的代币余额和授权额度
5. **先在测试网测试**: 在主网操作前，先在测试网测试
6. **小额测试**: 首次使用时先用小额进行测试
7. **防重入保护**: 合约已实现 ReentrancyGuard，防止重入攻击
8. **检查数组长度**: 确保所有数组参数长度匹配

## 错误处理

常见错误及解决方案：

| 错误信息 | 原因 | 解决方案 |
|---------|------|---------|
| `Invalid NFT contract address` | NFT 合约地址为零地址 | 检查并提供正确的合约地址 |
| `Arrays length mismatch` | 数组长度不匹配 | 确保 recipients 和 tokenIds/amounts 数组长度相同 |
| `Not token owner` | 不是 NFT 所有者 | 确保你拥有要转移的 NFT |
| `Invalid recipient address` | 接收地址无效 | 检查接收地址，不能是零地址 |
| `Amount must be greater than 0` | 转账金额为 0 | 提供大于 0 的转账金额 |
| `Empty arrays` | 数组为空 | 提供至少一个接收者/tokenId |
| `ERC20: insufficient allowance` | 代币授权额度不足 | 增加授权额度 |
| `ERC721: caller is not token owner or approved` | 未授权 NFT 操作 | 调用 setApprovalForAll 授权 |

## 部署指南

### 1. 部署合约

```bash
# 使用 Hardhat Deploy
npx hardhat deploy --tags BatchTransfer --network <network-name>

# 或者使用脚本部署
npx hardhat run scripts/deploy-batch-transfer.ts --network <network-name>
```

### 2. 验证合约

**自动验证**（如果安装了 hardhat-etherscan）：
```bash
npx hardhat verify --network sepoliaCustom 0x29aAD71f97Da6AA0F4096cfd50DA395eAa18D964
```

**手动验证**（在 Etherscan 上）：

1. 访问 [Sepolia Etherscan](https://sepolia.etherscan.io/address/0x29aAD71f97Da6AA0F4096cfd50DA395eAa18D964#code)
2. 点击 "Contract" 标签
3. 点击 "Verify and Publish"
4. 选择以下选项：
   - Compiler Type: Solidity (Single file)
   - Compiler Version: v0.8.28
   - License Type: MIT
5. 复制 `contracts/utils/BatchTransfer.sol` 的完整代码（包含所有 imports）
6. 优化设置：
   - Optimization: Yes
   - Runs: 1000000
   - EVM Version: cancun
7. 提交验证

**已验证**: ✅ 合约已在 Sepolia 测试网上部署并可供使用

### 3. 测试合约

```bash
# 运行测试
npx hardhat test test/BatchTransfer.test.ts

# 查看测试覆盖率
npx hardhat coverage
```

## 技术规格

- **Solidity 版本**: ^0.8.19
- **依赖库**: 
  - OpenZeppelin Contracts (ERC721, ERC20, SafeERC20, ReentrancyGuard)
- **安全特性**:
  - ReentrancyGuard: 防止重入攻击
  - SafeERC20: 安全的 ERC20 操作
  - 完善的输入验证
- **Gas 优化**:
  - 使用 calldata 而非 memory
  - 优化的循环结构
  - 批量操作减少交易次数

## 许可证

MIT License

## 支持

如有问题或建议，请联系开发团队或提交 Issue。

