# CPNFT Staking 可升级合约使用指南

## 概述

CPNFT Staking 是一个可升级的NFT质押合约，使用UUPS（Universal Upgradeable Proxy Standard）代理模式。该合约支持多等级收益计算、组合加成，并通过AccountManager自动计算收益地址。

## 主要特性

- ✅ **可升级合约**：使用OpenZeppelin UUPS代理模式
- ✅ **自动收益地址计算**：通过AccountManager计算AAccount地址
- ✅ **多等级NFT支持**：SSS、SS、S、A、B、C六个等级
- ✅ **复杂收益计算**：支持衰减、组合加成、动态调整
- ✅ **批量操作**：支持批量质押、批量领取
- ✅ **管理员功能**：支持参数调整、紧急提取

## 合约架构

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   UUPS Proxy    │────│  Implementation  │────│  AccountManager │
│  (User calls)   │    │   (Logic)        │    │  (Address calc) │
└─────────────────┘    └──────────────────┘    └─────────────────┘
         │                       │
         │                       │
    ┌────▼────┐              ┌───▼───┐
    │ CPNFT   │              │ CPOP  │
    │ (NFTs)  │              │ Token │
    └─────────┘              └───────┘
```

## 部署流程

### 1. 部署实现合约

```typescript
const CPNFTStakingFactory = await ethers.getContractFactory("CPNFTStaking");
const implementation = await CPNFTStakingFactory.deploy();
```

### 2. 部署UUPS代理

```typescript
const initData = CPNFTStakingFactory.interface.encodeFunctionData("initialize", [
  cpnftAddress,           // CPNFT合约地址
  cpopTokenAddress,       // CPOPToken合约地址
  accountManagerAddress,  // AccountManager合约地址
  defaultMasterSignerAddress, // 默认主签名者地址
  ownerAddress           // 合约拥有者地址
]);

const proxy = await UUPSProxyFactory.deploy(implementation.address, initData);
```

### 3. 验证部署

```typescript
const staking = CPNFTStakingFactory.attach(proxy.address);
const version = await staking.getVersion(); // "1.0.0"
```

## 核心功能

### 1. 质押NFT

```typescript
// 质押单个NFT
await staking.stake(tokenId);

// 批量质押
await staking.batchStake([tokenId1, tokenId2, tokenId3]);
```

### 2. 收益地址计算

```typescript
// 计算收益地址
const beneficiary = await staking.calculateBeneficiaryAddress(tokenId, owner);

// 计算masterSigner
const masterSigner = await staking.getCalculatedMasterSigner(tokenId);

// 验证收益地址
const isValid = await staking.verifyBeneficiaryAddress(tokenId);
```

### 3. 奖励管理

```typescript
// 领取单个奖励
await staking.claimReward(tokenId);

// 批量领取奖励
await staking.batchClaimRewards([tokenId1, tokenId2]);

// 查询待领取奖励
const pendingRewards = await staking.getUserPendingRewards(userAddress);
```

### 4. 取消质押

```typescript
// 取消质押（可能有惩罚）
await staking.unstake(tokenId);
```

## 管理员功能

### 1. 系统参数调整

```typescript
// 更新动态乘数（80%-120%）
await staking.updateDynamicMultiplier(level, multiplier);

// 设置等级总供应量
await staking.setLevelTotalSupply(level, supply);

// 更新默认主签名者
await staking.updateDefaultMasterSigner(newMasterSigner);
```

### 2. 系统控制

```typescript
// 暂停/恢复质押
await staking.setStakingPaused(true);  // 暂停
await staking.setStakingPaused(false); // 恢复

// 紧急提取
await staking.emergencyWithdraw(tokenId);
```

## 升级流程

### 1. 部署新实现

```typescript
const NewCPNFTStakingFactory = await ethers.getContractFactory("CPNFTStakingV2");
const newImplementation = await NewCPNFTStakingFactory.deploy();
```

### 2. 执行升级

```typescript
const staking = await ethers.getContractAt("CPNFTStaking", proxyAddress);
const upgradeTx = await staking.upgradeTo(newImplementation.address);
await upgradeTx.wait();
```

### 3. 验证升级

```typescript
const newVersion = await staking.getVersion(); // "1.1.0"
```

## 收益计算逻辑

### 1. 基础收益

| 等级 | 初始日收益 | 衰减间隔 | 衰减率 | 最大衰减率 |
|------|------------|----------|--------|------------|
| SSS  | 10 CPOP    | 180天    | 10%    | 20%        |
| SS   | 5 CPOP     | 90天     | 15%    | 40%        |
| S    | 3 CPOP     | 60天     | 20%    | 50%        |
| A    | 1.5 CPOP   | 45天     | 25%    | 60%        |
| B    | 0.8 CPOP   | 30天     | 30%    | 70%        |
| C    | 0.3 CPOP   | 20天     | 35%    | 80%        |

### 2. 组合加成

- 3个同等级NFT：+5%
- 5个同等级NFT：+10%
- 10个同等级NFT：+20%（SSS级除外）

### 3. 计算公式

```
单NFT每日收益 = 初始日收益 × 衰减系数 × 组合加成 × 动态乘数
```

## 地址计算逻辑

### 1. MasterSigner计算

```solidity
bytes32 hash = keccak256(abi.encodePacked(
    "CPNFT_MASTER_SIGNER",
    tokenId,
    uint8(level),
    block.chainid
));
masterSigner = address(uint160(uint256(hash)));
```

### 2. 收益地址计算

```typescript
// 通过AccountManager计算AAccount地址
const beneficiary = await accountManager.getAccountAddress(owner, masterSigner);
```

## 安全考虑

### 1. 升级安全

- 只有合约拥有者可以执行升级
- 升级前会验证新实现合约的有效性
- 升级后自动验证核心功能

### 2. 质押安全

- 质押期间NFT被锁定，无法转移
- 提前赎回有惩罚机制
- 支持紧急提取功能

### 3. 权限控制

- 使用OpenZeppelin的Ownable模式
- 关键功能只有拥有者可以调用
- 支持所有权转移

## 事件监控

### 1. 质押事件

```solidity
event Staked(address indexed owner, address indexed beneficiary, uint256 indexed tokenId, uint256 stakeTime);
event Unstaked(address indexed owner, address indexed beneficiary, uint256 indexed tokenId, uint256 reward);
```

### 2. 奖励事件

```solidity
event RewardsClaimed(address indexed beneficiary, uint256 amount);
event ComboBonusUpdated(address indexed user, uint256 bonusMultiplier);
```

### 3. 系统事件

```solidity
event DynamicMultiplierUpdated(NFTLevel level, uint256 multiplier);
event DefaultMasterSignerUpdated(address indexed oldSigner, address indexed newSigner);
```

## 最佳实践

### 1. 部署前准备

- 确保所有依赖合约已部署
- 验证合约地址的正确性
- 设置合适的默认参数

### 2. 升级前检查

- 充分测试新实现合约
- 备份当前状态
- 准备回滚方案

### 3. 监控建议

- 监控质押和提取事件
- 跟踪奖励发放情况
- 关注系统参数变化

## 故障排除

### 1. 常见错误

- "Account calculation failed": AccountManager地址不正确
- "Token already staked": NFT已被质押
- "Not stake owner": 非质押者调用

### 2. 解决方案

- 检查合约地址配置
- 验证NFT所有权状态
- 确认调用者权限

## 版本历史

- **v1.0.0**: 初始版本，基础质押功能
- **v1.1.0**: 添加组合加成功能
- **v1.2.0**: 优化gas消耗

## 联系支持

如有问题，请联系开发团队或查看相关文档。
