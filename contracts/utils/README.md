# Utils Contracts

## BatchTransfer

`BatchTransfer.sol` 是一个用于批量转移 NFT 和 FT（同质化代币）的智能合约。

### 主要功能

- **批量转移 NFT**: 支持将多个 NFT 转移给多个接收者
- **批量转移 FT**: 支持将 ERC20 代币批量转移给多个接收者  
- **组合转移**: 在单个交易中同时转移 NFT 和 FT
- **多集合支持**: 支持从不同的 NFT 集合批量转移
- **多代币支持**: 支持批量转移多种 ERC20 代币

### 文档

详细使用指南请查看：[BatchTransfer 使用指南](../../docs/BatchTransfer-Usage-Guide.md)

### 示例代码

查看使用示例：[batch-transfer-usage.ts](../../examples/batch-transfer-usage.ts)

### 测试

运行测试：
```bash
npx hardhat test test/BatchTransfer.test.ts
```

### 部署

```bash
npx hardhat deploy --tags BatchTransfer --network <network-name>
```

