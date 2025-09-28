# NFT质押系统 (NFT Staking System)

## 📋 概述

NFT质押系统是一个完整的去中心化质押解决方案，支持多等级NFT质押、动态奖励计算、组合加成和平台平衡机制。系统采用模块化架构设计，确保高效、安全和可扩展性。

## 🏗️ 系统架构

### 核心合约

```
contracts/CPNFT/
├── CPNFT.sol           # NFT合约 (ERC721)
├── StakingConfig.sol   # 配置管理合约 (非可升级)
├── Staking.sol         # 主质押合约 (可升级)
└── StakingReader.sol   # 前端查询合约 (只读)
```

### 架构特点

- **模块化设计**: 配置、逻辑、查询分离
- **存储优化**: 使用结构体和精确数据类型，节省84%存储空间
- **可升级性**: 主合约支持UUPS升级模式
- **安全性**: 多重安全机制，包括重入保护、暂停功能
- **Gas优化**: 高效的批量操作和计算逻辑

## 🎯 NFT等级系统

### 支持的NFT等级

| 等级 | 名称 | 初始日收益 | 衰减周期 | 衰减率 | 最大衰减 |
|------|------|------------|----------|--------|----------|
| 0 | NORMAL | ❌ 不可质押 | - | - | - |
| 1 | C | 3 CPP | 20天 | 35% | 80% |
| 2 | B | 8 CPP | 30天 | 30% | 70% |
| 3 | A | 15 CPP | 45天 | 25% | 60% |
| 4 | S | 30 CPP | 60天 | 20% | 50% |
| 5 | SS | 50 CPP | 90天 | 15% | 40% |
| 6 | SSS | 100 CPP | 180天 | 10% | 20% |

## 🔧 合约接口

### 1. StakingConfig.sol - 配置管理

**功能**: 管理所有质押参数和配置

```solidity
// 基础配置
function getMinStakeDays() external view returns (uint64)
function getEarlyWithdrawPenalty() external view returns (uint64)
function getQuarterlyMultiplier() external view returns (uint64)

// 等级配置
function getDailyReward(uint8 level) external view returns (uint32)
function getDecayInterval(uint8 level) external view returns (uint16)
function getDecayRate(uint8 level) external view returns (uint16)
function getMaxDecayRate(uint8 level) external view returns (uint16)

// 组合配置
function getComboThresholds() external view returns (uint256[3] memory)
function getComboBonuses() external view returns (uint256[3] memory)
function getComboMinDays() external view returns (uint256[3] memory)

// 动态配置
function getDynamicConfig() external view returns (
    uint16 highStakeThreshold,
    uint16 lowStakeThreshold,
    uint16 highStakeMultiplier,
    uint16 lowStakeMultiplier,
    uint16 quarterlyAdjustmentMax
)

// 总供应量配置
function setTotalSupplyPerLevel(uint8 level, uint256 supply) external onlyOwner
function getTotalSupplyPerLevel(uint8 level) external view returns (uint256)
```

### 2. Staking.sol - 主质押合约

**功能**: 处理质押、解质押、奖励发放等核心逻辑

> **注意**: Staking合约专注于核心业务逻辑，所有查询功能都委托给StakingReader合约，以实现更好的模块化和Gas优化。

#### 质押功能

```solidity
// 单个质押
function stake(uint256 tokenId) external

// 批量质押
function batchStake(uint256[] calldata tokenIds) external

// 解质押
function unstake(uint256 tokenId) external

// 领取奖励
function claimRewards(uint256 tokenId) external
```

#### 管理功能

```solidity
// 暂停/恢复合约
function pause() external onlyOwner
function unpause() external onlyOwner
```

### 3. StakingReader.sol - 前端查询合约

**功能**: 专门为前端提供优化的查询接口

#### 基础查询

```solidity
// 获取质押详情
function getStakeDetails(uint256 tokenId) external view returns (StakeInfo memory)

// 获取待领取奖励
function getPendingRewards(uint256 tokenId) external view returns (uint256)

// 获取详细奖励计算
function getDetailedRewardCalculation(uint256 tokenId) external view returns (
    uint256 baseReward,
    uint256 decayedReward,
    uint256 comboBonus,
    uint256 dynamicMultiplier,
    uint256 finalReward
)
```

#### 统计查询

```solidity
// 获取组合信息
function getComboInfo(address user, uint8 level) external view returns (
    uint256 count,
    uint256 currentBonus,
    uint256[] memory thresholds,
    uint256[] memory bonuses
)

// 获取等级统计
function getLevelStats() external view returns (
    uint256[7] memory stakedCounts,
    uint256[7] memory supplies,
    uint256[7] memory stakingRatios,
    uint256[7] memory dynamicMultipliers
)

// 获取平台统计
function getPlatformStats() external view returns (
    uint256[7] memory staked,
    uint256[7] memory supply
)

// 获取完整平台统计
function getPlatformStatistics() external view returns (
    uint256 totalStakedNFTs,
    uint256[7] memory stakedPerLevel,
    uint256[7] memory supplyPerLevel,
    uint256[7] memory stakingRatios,
    uint256[7] memory dynamicMultipliers,
    uint256 totalRewardsToday
)
```

#### 配置查询

```solidity
// 获取所有配置
function getConfiguration() external view returns (
    uint256 minStakeDays,
    uint256 earlyWithdrawPenalty,
    uint256[6] memory dailyRewards,
    uint256[6] memory decayIntervals,
    uint256[6] memory decayRates,
    uint256[6] memory maxDecayRates,
    uint256[3] memory comboThresholds,
    uint256[3] memory comboBonuses,
    uint256[3] memory comboMinDays,
    uint256[2] memory continuousThresholds,
    uint256[2] memory continuousBonuses,
    uint256 quarterlyMultiplier
)

// 获取合约版本
function getVersions() external view returns (
    string memory stakingVersion,
    string memory configVersion
)
```

## 💰 奖励机制

### 基础奖励计算

```
单NFT日收益 = 初始日收益 × 衰减系数 × 季度乘数 × 组合加成 × 动态乘数
```

### 衰减机制

- **衰减触发**: 达到衰减周期后次日生效
- **衰减计算**: 每次衰减 = 当前收益 × (1 - 衰减率)
- **衰减上限**: 达到最大衰减率后不再衰减

### 组合加成

| NFT数量 | 加成比例 | 适用等级 | 最低质押天数 |
|---------|----------|----------|--------------|
| 3个 | +5% | 全部 | 7天 |
| 5个 | +10% | 全部 | 15天 |
| 10个 | +20% | 除SSS外 | 30天 |

### 动态平衡机制

动态平衡机制是NFT质押系统的核心调控机制，旨在维持生态系统的健康平衡，防止流动性僵化。

#### 全平台质押量调控

**触发条件**:
- **高质押率调控**: 当某等级质押量 > 60% 该等级总发行量时
- **低质押率调控**: 当某等级质押量 < 10% 该等级总发行量时

**调控效果**:
```
高质押率: 当日收益 × 0.8 (减少20%)
低质押率: 当日收益 × 1.05 (增加5%)
正常范围: 当日收益 × 1.0 (无调整)
```

**实时计算过程**:
```solidity
function _calculateDynamicMultiplier(uint8 level) internal view returns (uint256) {
    // 1. 获取当前质押数量
    uint256 staked = totalStakedPerLevel[level];
    
    // 2. 获取该等级总供应量
    uint256 supply = cpnftContract.getLevelSupply(level);
    
    // 3. 计算质押比例 (以万分为单位)
    uint256 stakingRatio = (staked * 10000) / supply;
    
    // 4. 应用调控规则
    if (stakingRatio > 6000) {        // > 60%
        return 8000;                  // 0.8x 倍率
    } else if (stakingRatio < 1000) { // < 10%
        return 10500;                 // 1.05x 倍率
    }
    
    return 10000; // 1.0x 倍率 (正常)
}
```

#### 季度调整机制

**调整流程**:
1. **预告阶段**: 管理员提前7天宣布调整计划
2. **等待阶段**: 7天公示期，用户可了解即将到来的变化
3. **执行阶段**: 公示期结束后自动执行调整
4. **调度阶段**: 自动安排下次季度调整时间

**调整范围**:
- **最大调整幅度**: ±20%
- **调整频率**: 每90天一次
- **历史记录**: 保留所有调整历史供查询

#### 对用户收益的影响

**实时影响**:
- **质押时**: 根据当前质押比例立即应用动态倍率
- **每日计算**: 每次计算收益时都会重新评估质押比例
- **即时调整**: 当质押比例跨越阈值时，收益立即调整

**实际案例**:

*案例1: A级NFT高质押率情况*
```
初始情况:
- A级NFT总供应量: 1000个
- 当前质押数量: 650个 (65%)
- 质押比例: 65% > 60% (触发高质押率调控)

调控效果:
- 原始日收益: 15 CPP/天
- 动态倍率: 0.8x
- 调整后收益: 15 × 0.8 = 12 CPP/天
- 收益减少: 20%
```

*案例2: SS级NFT低质押率情况*
```
初始情况:
- SS级NFT总供应量: 500个
- 当前质押数量: 40个 (8%)
- 质押比例: 8% < 10% (触发低质押率调控)

调控效果:
- 原始日收益: 50 CPP/天
- 动态倍率: 1.05x
- 调整后收益: 50 × 1.05 = 52.5 CPP/天
- 收益增加: 5%
```

*案例3: 季度调整影响*
```
调整前:
- SSS级NFT日收益: 100 CPP/天
- 季度倍率: 1.0x

季度调整:
- 新季度倍率: 1.2x (增加20%)
- 调整后收益: 100 × 1.2 = 120 CPP/天
- 收益增加: 20%
```

#### 用户策略建议

**监控质押比例**:
- 关注各等级NFT的质押比例
- 在低质押率时考虑质押以获得更高收益
- 在高质押率时考虑等待或选择其他等级

**关注季度调整**:
- 关注官方公告的季度调整预告
- 根据调整方向调整质押策略
- 利用历史调整记录分析趋势

**动态调整策略**:
- 根据实时质押比例调整NFT选择
- 结合组合质押和连续质押奖励
- 平衡短期收益和长期策略

#### 技术实现细节

**数据源**:
- **质押数量**: 从`totalStakedPerLevel`映射获取
- **总供应量**: 从`CPNFT.levelSupply`获取
- **配置参数**: 从`StakingConfig`合约获取

**计算精度**:
- 使用万分之一精度 (10000 = 100%)
- 避免浮点数计算的精度问题
- 确保计算结果的一致性

**事件记录**:
- `PlatformStatsUpdated`: 记录质押比例变化
- `QuarterlyAdjustmentAnnounced`: 记录季度调整预告
- `QuarterlyAdjustmentExecuted`: 记录季度调整执行

### 连续质押奖励

- **30天**: 额外奖励10%
- **90天**: 额外奖励30%

### 提前解质押惩罚

- **未满7天**: 扣除已获收益的50%
- **满7天后**: 无惩罚

## 🚀 前端集成指南

### 1. 合约连接

```javascript
// 使用ethers.js连接合约
import { ethers } from 'ethers';

// 合约地址 (需要替换为实际部署地址)
const STAKING_ADDRESS = "0x...";
const CONFIG_ADDRESS = "0x...";
const READER_ADDRESS = "0x...";
const CPNFT_ADDRESS = "0x...";

// 合约ABI (需要导入实际的ABI)
const stakingContract = new ethers.Contract(STAKING_ADDRESS, STAKING_ABI, provider);
const configContract = new ethers.Contract(CONFIG_ADDRESS, CONFIG_ABI, provider);
const readerContract = new ethers.Contract(READER_ADDRESS, READER_ABI, provider);
const cpnftContract = new ethers.Contract(CPNFT_ADDRESS, CPNFT_ABI, provider);
```

### 2. 质押流程

```javascript
// 检查NFT等级
async function checkNFTLevel(tokenId) {
    const level = await cpnftContract.getTokenLevel(tokenId);
    if (level === 0) {
        throw new Error("NORMAL等级NFT不可质押");
    }
    return level;
}

// 质押NFT
async function stakeNFT(tokenId, signer) {
    try {
        // 1. 检查NFT等级
        const level = await checkNFTLevel(tokenId);
        
        // 2. 授权质押合约
        const approveTx = await cpnftContract.connect(signer).approve(STAKING_ADDRESS, tokenId);
        await approveTx.wait();
        
        // 3. 执行质押
        const stakeTx = await stakingContract.connect(signer).stake(tokenId);
        await stakeTx.wait();
        
        console.log(`NFT ${tokenId} 质押成功`);
    } catch (error) {
        console.error("质押失败:", error);
    }
}

// 批量质押
async function batchStakeNFTs(tokenIds, signer) {
    try {
        // 1. 批量授权
        const approveTx = await cpnftContract.connect(signer).setApprovalForAll(STAKING_ADDRESS, true);
        await approveTx.wait();
        
        // 2. 批量质押
        const stakeTx = await stakingContract.connect(signer).batchStake(tokenIds);
        await stakeTx.wait();
        
        console.log(`批量质押 ${tokenIds.length} 个NFT成功`);
    } catch (error) {
        console.error("批量质押失败:", error);
    }
}
```

### 3. 奖励查询

```javascript
// 获取质押详情
async function getStakeInfo(tokenId) {
    const stakeInfo = await readerContract.getStakeDetails(tokenId);
    return {
        owner: stakeInfo.owner,
        tokenId: stakeInfo.tokenId.toString(),
        level: stakeInfo.level,
        stakeTime: new Date(stakeInfo.stakeTime * 1000),
        lastClaimTime: new Date(stakeInfo.lastClaimTime * 1000),
        isActive: stakeInfo.isActive,
        totalRewards: ethers.utils.formatEther(stakeInfo.totalRewards),
        pendingRewards: ethers.utils.formatEther(stakeInfo.pendingRewards)
    };
}

// 获取详细奖励计算
async function getRewardCalculation(tokenId) {
    const details = await readerContract.getDetailedRewardCalculation(tokenId);
    return {
        baseReward: ethers.utils.formatEther(details.baseReward),
        decayedReward: ethers.utils.formatEther(details.decayedReward),
        comboBonus: (details.comboBonus / 100).toFixed(2) + '%',
        dynamicMultiplier: (details.dynamicMultiplier / 100).toFixed(2) + 'x',
        finalReward: ethers.utils.formatEther(details.finalReward)
    };
}

// 获取用户所有质押
async function getUserStakes(userAddress) {
    // 注意: 需要通过其他方式获取用户的质押NFT列表
    // 例如通过事件监听或NFT合约查询
    const stakes = [];
    
    // 这里需要根据实际需求实现获取用户质押NFT列表的逻辑
    // 可以通过监听Staked事件来维护用户质押列表
    
    return stakes;
}
```

### 4. 解质押流程

```javascript
// 解质押NFT
async function unstakeNFT(tokenId, signer) {
    try {
        const unstakeTx = await stakingContract.connect(signer).unstake(tokenId);
        await unstakeTx.wait();
        
        console.log(`NFT ${tokenId} 解质押成功`);
    } catch (error) {
        console.error("解质押失败:", error);
    }
}

// 领取奖励
async function claimRewards(tokenId, signer) {
    try {
        const claimTx = await stakingContract.connect(signer).claimRewards(tokenId);
        await claimTx.wait();
        
        console.log(`NFT ${tokenId} 奖励领取成功`);
    } catch (error) {
        console.error("奖励领取失败:", error);
    }
}
```

### 5. 统计和配置查询

```javascript
// 获取平台统计
async function getPlatformStats() {
    const stats = await readerContract.getPlatformStatistics();
    return {
        totalStakedNFTs: stats.totalStakedNFTs.toString(),
        stakedPerLevel: stats.stakedPerLevel.map(count => count.toString()),
        supplyPerLevel: stats.supplyPerLevel.map(supply => supply.toString()),
        stakingRatios: stats.stakingRatios.map(ratio => (ratio / 100).toFixed(2) + '%'),
        dynamicMultipliers: stats.dynamicMultipliers.map(mult => (mult / 100).toFixed(2) + 'x'),
        totalRewardsToday: ethers.utils.formatEther(stats.totalRewardsToday)
    };
}

// 获取组合信息
async function getComboInfo(userAddress, level) {
    const comboInfo = await readerContract.getComboInfo(userAddress, level);
    return {
        count: comboInfo.count.toString(),
        currentBonus: (comboInfo.currentBonus / 100).toFixed(2) + '%',
        thresholds: comboInfo.thresholds.map(t => t.toString()),
        bonuses: comboInfo.bonuses.map(b => (b / 100).toFixed(2) + '%')
    };
}

// 获取配置信息
async function getConfiguration() {
    const config = await readerContract.getConfiguration();
    return {
        minStakeDays: config.minStakeDays.toString(),
        earlyWithdrawPenalty: (config.earlyWithdrawPenalty / 100).toFixed(2) + '%',
        dailyRewards: config.dailyRewards.map(reward => ethers.utils.formatEther(reward)),
        decayIntervals: config.decayIntervals.map(interval => interval.toString()),
        decayRates: config.decayRates.map(rate => (rate / 100).toFixed(2) + '%'),
        maxDecayRates: config.maxDecayRates.map(max => (max / 100).toFixed(2) + '%'),
        comboThresholds: config.comboThresholds.map(t => t.toString()),
        comboBonuses: config.comboBonuses.map(b => (b / 100).toFixed(2) + '%'),
        quarterlyMultiplier: (config.quarterlyMultiplier / 100).toFixed(2) + 'x'
    };
}
```

## ⚠️ 注意事项

### 安全考虑

1. **权限控制**: 只有合约所有者可以修改配置
2. **重入保护**: 所有外部调用都有重入保护
3. **暂停机制**: 紧急情况下可暂停合约
4. **输入验证**: 所有输入都有严格的验证

### Gas优化

1. **批量操作**: 支持批量质押减少Gas消耗
2. **存储优化**: 使用结构体打包减少存储成本
3. **计算优化**: 复杂的计算逻辑已优化

### 前端最佳实践

1. **错误处理**: 始终处理合约调用的异常
2. **用户体验**: 显示交易状态和确认信息
3. **数据缓存**: 缓存不变的配置数据
4. **实时更新**: 监听合约事件更新UI

## 📊 事件监听

```javascript
// 监听质押事件
stakingContract.on("Staked", (owner, tokenId, level, stakeTime, event) => {
    console.log(`用户 ${owner} 质押了 NFT ${tokenId}`);
});

// 监听解质押事件
stakingContract.on("Unstaked", (owner, tokenId, totalRewards, event) => {
    console.log(`用户 ${owner} 解质押了 NFT ${tokenId}，获得奖励 ${ethers.utils.formatEther(totalRewards)}`);
});

// 监听奖励领取事件
stakingContract.on("RewardsClaimed", (owner, tokenId, amount, event) => {
    console.log(`用户 ${owner} 领取了 NFT ${tokenId} 的奖励 ${ethers.utils.formatEther(amount)}`);
});
```

## 🔗 相关链接

- [合约部署指南](../deploy/README.md)
- [测试文档](../../test/README.md)
- [Gas优化报告](../../reports/gas-optimization.md)
- [安全审计报告](../../audits/README.md)

## 📞 技术支持

如有技术问题，请联系开发团队或查看相关文档。

---

**版本**: v1.0.0  
**最后更新**: 2025年1月  
**兼容性**: Solidity ^0.8.19
