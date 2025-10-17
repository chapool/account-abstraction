# BatchTransfer 合约部署记录

## 📋 部署信息

### Sepolia 测试网部署

- **合约名称**: BatchTransfer
- **合约地址**: `0x29aAD71f97Da6AA0F4096cfd50DA395eAa18D964`
- **网络**: Sepolia Testnet (Chain ID: 11155111)
- **部署者地址**: `0xa3B605fB633AD0A0DC4B74b10bBfc40fDB050d35`
- **部署交易**: `0x4ca33acf2b57787c5d2a8e940b95978e725bbf9ffe8c75f8d93e0c7b7834e26b`
- **Gas 消耗**: 1,578,999
- **部署时间**: 2025-10-17
- **区块浏览器**: [查看合约](https://sepolia.etherscan.io/address/0x29aAD71f97Da6AA0F4096cfd50DA395eAa18D964)

## ✨ 合约功能

BatchTransfer 是一个用于批量转移 NFT（ERC721）和 FT 同质化代币（ERC20）的智能合约。

### 核心功能

1. **NFT 批量转移**
   - ✅ 批量转移 NFT 给多个接收者
   - ✅ 批量转移多个 NFT 给单个接收者
   - ✅ 转移来自不同 NFT 集合的代币

2. **FT（代币）批量转移**
   - ✅ 批量转移不同金额的 ERC20 代币
   - ✅ 批量转移相同金额的 ERC20 代币
   - ✅ 批量转移多种 ERC20 代币

3. **组合转移**
   - ✅ 在单个交易中同时转移 NFT 和 FT

4. **查询功能**
   - ✅ 检查 NFT 授权状态
   - ✅ 检查代币授权额度

## 🔧 技术规格

- **Solidity 版本**: 0.8.28
- **EVM 版本**: Cancun
- **优化**: 启用（Runs: 1,000,000）
- **安全特性**: ReentrancyGuardTransient（防重入攻击，节省 gas）
- **依赖库**: OpenZeppelin Contracts

## 🚀 快速使用

### 在 JavaScript/TypeScript 中使用

```javascript
const { ethers } = require("ethers");

// 合约地址和 ABI
const BATCH_TRANSFER_ADDRESS = "0x29aAD71f97Da6AA0F4096cfd50DA395eAa18D964";
const BATCH_TRANSFER_ABI = [
  "function batchTransferNFT(address nftContract, address[] recipients, uint256[] tokenIds)",
  "function batchTransferToken(address tokenContract, address[] recipients, uint256[] amounts)",
  "function combinedBatchTransfer(address nftContract, address[] nftRecipients, uint256[] tokenIds, address tokenContract, address[] tokenRecipients, uint256[] tokenAmounts)"
];

// 连接到合约
const provider = new ethers.providers.JsonRpcProvider("YOUR_SEPOLIA_RPC_URL");
const wallet = new ethers.Wallet("YOUR_PRIVATE_KEY", provider);
const batchTransfer = new ethers.Contract(BATCH_TRANSFER_ADDRESS, BATCH_TRANSFER_ABI, wallet);

// 批量转移 NFT
await batchTransfer.batchTransferNFT(
  nftContractAddress,
  [recipient1, recipient2],
  [tokenId1, tokenId2]
);
```

## 📖 完整文档

详细使用指南请查看：[BatchTransfer 使用指南](./docs/BatchTransfer-Usage-Guide.md)

## 🧪 测试

运行测试：
```bash
npx hardhat test test/BatchTransfer.test.ts
```

## 🔍 验证合约

合约源码可在 Etherscan 上查看和验证：
- 访问: https://sepolia.etherscan.io/address/0x29aAD71f97Da6AA0F4096cfd50DA395eAa18D964
- 合约已开源，可直接查看代码

## 📝 使用案例

1. **NFT 空投**: 向多个用户批量分发 NFT
2. **代币分发**: 向团队成员或社区成员批量发放代币
3. **奖励发放**: 向活动获胜者批量发放 NFT 和代币奖励
4. **资产迁移**: 批量转移资产到新钱包
5. **游戏道具分发**: 批量发放游戏内 NFT 道具和代币

## ⚠️ 安全提示

1. 使用前务必授权合约（`setApprovalForAll` for NFTs 或 `approve` for tokens）
2. 确认所有接收地址正确无误
3. 测试网先测试，再在主网使用
4. 批量操作前检查 gas 费用
5. 合约已实现防重入保护

## 💡 Gas 优化

- 使用 `batchTransferTokenEqual` 发送相同金额可节省 gas
- 使用 `combinedBatchTransfer` 同时转移 NFT 和代币可节省交易费用
- 批量操作比单个操作节省约 20-40% 的 gas

## 📞 支持

如有问题或建议，请：
1. 查看[完整文档](./docs/BatchTransfer-Usage-Guide.md)
2. 查看[示例代码](./examples/batch-transfer-usage.ts)
3. 查看[测试用例](./test/BatchTransfer.test.ts)

## 📜 License

MIT License

