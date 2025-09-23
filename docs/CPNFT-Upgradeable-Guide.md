# CPNFT 可升级合约指南

## 概述

CPNFT 合约已经升级为可升级合约，使用 OpenZeppelin 的 UUPS (Universal Upgradeable Proxy Standard) 模式。这允许我们在不改变合约地址的情况下升级合约逻辑。

## 架构

### 合约结构

1. **CPNFTUpgradeable.sol** - 可升级的实现合约
2. **CPNFTProxy.sol** - ERC1967 代理合约
3. **CPNFTInitializer.sol** - 初始化辅助合约

### 升级模式

使用 UUPS (Universal Upgradeable Proxy Standard) 模式：
- 升级逻辑存储在实现合约中
- 只有合约拥有者可以执行升级
- 使用 `_authorizeUpgrade` 函数控制升级权限

## 部署

### 1. 部署可升级版本

```bash
# 部署可升级的 CPNFT 合约
npx hardhat deploy --tags CPNFTUpgradeable --network sepolia
```

### 2. 部署信息

部署完成后，会在 `deployments/sepoliaCustom/` 目录下生成部署信息文件，包含：
- 代理合约地址
- 实现合约地址
- 合约拥有者
- 部署配置

## 升级流程

### 1. 准备新版本

在升级之前，确保新版本的实现合约：
- 保持存储布局兼容性
- 继承相同的父合约
- 实现 `_authorizeUpgrade` 函数

### 2. 执行升级

```bash
# 运行升级脚本
npx hardhat run scripts/upgradeCPNFT.ts --network sepolia
```

### 3. 验证升级

升级完成后，验证：
- 合约地址保持不变
- 新功能正常工作
- 现有数据完整

## 重要注意事项

### 存储布局兼容性

⚠️ **警告**: 升级时必须保持存储布局兼容性：

```solidity
// ✅ 正确：在末尾添加新变量
mapping(uint256 => bool) private _newFeature;

// ❌ 错误：修改现有变量类型
mapping(uint256 => uint8) private _tokenLevels; // 原来是 NFTLevel

// ❌ 错误：删除现有变量
// mapping(uint256 => NFTLevel) private _tokenLevels; // 已删除
```

### 初始化函数

- 实现合约使用 `initialize` 函数而不是构造函数
- 只能调用一次初始化函数
- 代理合约在部署时自动调用初始化

### 权限控制

```solidity
function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
```

只有合约拥有者可以执行升级。

## 测试

### 1. 本地测试

```bash
# 编译合约
npx hardhat compile

# 运行测试
npx hardhat test test/CPNFTUpgradeable.test.ts
```

### 2. 升级测试

```bash
# 测试升级流程
npx hardhat test test/CPNFTUpgrade.test.ts
```

## 最佳实践

### 1. 版本控制

- 使用语义化版本号
- 在合约中添加版本信息
- 记录每次升级的变更

### 2. 测试策略

- 全面测试新功能
- 验证现有功能不受影响
- 测试存储数据完整性

### 3. 安全考虑

- 仔细审查升级代码
- 使用多签钱包控制升级权限
- 在测试网充分测试后再部署到主网

## 故障排除

### 常见问题

1. **存储布局冲突**
   ```
   Error: New implementation is not compatible
   ```
   解决：检查存储变量布局是否兼容

2. **初始化失败**
   ```
   Error: Initializable: contract is already initialized
   ```
   解决：确保只调用一次初始化函数

3. **权限不足**
   ```
   Error: Ownable: caller is not the owner
   ```
   解决：确保使用正确的拥有者账户

### 回滚策略

如果升级出现问题：
1. 部署修复版本
2. 执行新的升级
3. 验证功能正常

## 示例代码

### 添加新功能

```solidity
// CPNFTUpgradeableV2.sol
contract CPNFTUpgradeableV2 is CPNFTUpgradeable {
    // 添加新的存储变量（在末尾）
    mapping(uint256 => string) private _tokenMetadata;
    
    // 添加新功能
    function setTokenMetadata(uint256 tokenId, string memory metadata) external onlyOwner {
        require(_mintedTokens[tokenId], "Token does not exist");
        _tokenMetadata[tokenId] = metadata;
    }
    
    function getTokenMetadata(uint256 tokenId) external view returns (string memory) {
        require(_mintedTokens[tokenId], "Token does not exist");
        return _tokenMetadata[tokenId];
    }
    
    // 更新版本
    function version() public pure override returns (string memory) {
        return "2.0.0";
    }
}
```

### 升级脚本

```typescript
// scripts/upgradeToV2.ts
async function upgradeToV2() {
  const CPNFTUpgradeableV2 = await ethers.getContractFactory('CPNFTUpgradeableV2')
  
  // 部署新实现
  const newImplementation = await CPNFTUpgradeableV2.deploy()
  await newImplementation.deployed()
  
  // 升级代理
  const proxy = await ethers.getContractAt('CPNFTUpgradeable', PROXY_ADDRESS)
  await proxy.upgradeTo(newImplementation.address)
  
  console.log('Upgrade completed!')
}
```

## 相关资源

- [OpenZeppelin Upgradeable Contracts](https://docs.openzeppelin.com/contracts/4.x/upgradeable)
- [UUPS Proxy Pattern](https://eips.ethereum.org/EIPS/eip-1822)
- [ERC1967 Proxy](https://eips.ethereum.org/EIPS/eip-1967)
