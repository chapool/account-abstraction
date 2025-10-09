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
- **自动化管理**: 季度调整自动记录历史状态，无需手动干预
- **完整追踪**: 提供历史调整记录查询，支持收益重新计算

## 🚀 合约部署信息

### 已部署合约地址 (Sepolia测试网)

**质押系统合约 (最新部署):**
- **Staking Contract (Proxy)**: `0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5`
  - Implementation: 自动管理
- **StakingConfig Contract**: `0x37196054B23Be5CB977AA3284A3A844a8fe77861`
- **StakingReader Contract (Proxy)**: `0x6C9f7Fb0376C961FE79cED8cf09EbBBaDBfF0051`
  - Implementation: `0xb700544fB85d95A09Db71E2BE29Bb76d06386E7c`

**依赖合约:**
- **CPNFT Contract**: `0xcC63bf57EF4b4fE5635cF0745Ae7E2C75A63c7Ed`
- **CPOP Token Contract**: `0xA2d58d11c2752b010C2444fa2f795c6cf4cb76Bc`
- **Account Manager Contract**: `0x2E4f862Ba3ee6D84dd19ae9f002F5D8c0C5675ef`

### 部署配置

**基础配置:**
- 最小质押天数: 7天
- 提前解质押惩罚: 50% (5000 basis points)
- 季度乘数: 100% (10000 basis points)

**等级奖励配置 (每日):**
- Level 1 (C): 3 CPP
- Level 2 (B): 8 CPP  
- Level 3 (A): 15 CPP
- Level 4 (S): 30 CPP
- Level 5 (SS): 50 CPP
- Level 6 (SSS): 100 CPP

**衰减配置:**
- Level 1: 20天周期, 35%衰减率, 最大80%衰减
- Level 2: 30天周期, 30%衰减率, 最大70%衰减
- Level 3: 45天周期, 25%衰减率, 最大60%衰减
- Level 4: 60天周期, 20%衰减率, 最大50%衰减
- Level 5: 90天周期, 15%衰减率, 最大40%衰减
- Level 6: 180天周期, 10%衰减率, 最大20%衰减

### 权限设置

✅ **已完成的权限配置:**
- Staking合约已获得CPOP Token的MINTER_ROLE权限
- CPNFT合约已识别Staking合约地址
- StakingConfig已正确链接到Staking合约

### 部署时间

**最新部署**: 2025-01-09 (优化版本)  
**网络**: Sepolia测试网 (Chain ID: 11155111)  
**部署者**: `0xa3B605fB633AD0A0DC4B74b10bBfc40fDB050d35`

**更新说明:**
- Staking 合约: v3.1.0 (批量操作事件优化)
- StakingReader 合约: v2.0.0 (可升级 + 大小优化)
- 新增用户统计查询功能
- 合约大小优化，符合24KB限制

### 合约版本

- **Staking Contract**: v3.1.0
- **StakingConfig Contract**: v3.2.0
- **StakingReader Contract**: v1.0.0

### 前端集成地址

```javascript
// 合约地址配置 (最新版本)
const CONTRACT_ADDRESSES = {
  // 核心质押合约
  STAKING: "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5",           // 主质押合约 (可升级)
  STAKING_CONFIG: "0x37196054B23Be5CB977AA3284A3A844a8fe77861",   // 配置合约
  STAKING_READER: "0x6C9f7Fb0376C961FE79cED8cf09EbBBaDBfF0051",   // 查询合约 (可升级)
  
  // 依赖合约
  CPNFT: "0xcC63bf57EF4b4fE5635cF0745Ae7E2C75A63c7Ed",           // NFT合约
  CPOP_TOKEN: "0xA2d58d11c2752b010C2444fa2f795c6cf4cb76Bc",      // 代币合约
  ACCOUNT_MANAGER: "0x2E4f862Ba3ee6D84dd19ae9f002F5D8c0C5675ef"  // 账户管理
};

// 网络配置
const NETWORK_CONFIG = {
  chainId: 11155111,        // Sepolia
  name: "Sepolia Testnet",
  rpcUrl: "https://sepolia.infura.io/v3/YOUR_INFURA_KEY"
};
```

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

// 季度调整管理 (自动化功能)
function announceQuarterlyAdjustment(uint256 multiplier) external onlyOwner
function executeQuarterlyAdjustment() external onlyOwner
function getQuarterlyAdjustment(uint256 index) external view returns (...)
function getLatestQuarterlyAdjustment() external view returns (...)
function setStakingContract(address _stakingContract) external onlyOwner

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

// 批量解质押
function batchUnstake(uint256[] calldata tokenIds) external

// 领取奖励
function claimRewards(uint256 tokenId) external

// 批量领取奖励
function batchClaimRewards(uint256[] calldata tokenIds) external
```

#### 管理功能

```solidity
// 暂停/恢复合约
function pause() external onlyOwner
function unpause() external onlyOwner

// 历史调整记录 (自动化调用)
function recordHistoricalAdjustment() external
function getHistoricalAdjustment(uint256 index) external view returns (...)
function getHistoricalDynamicMultiplier(uint256 index, uint8 level) external view returns (uint256)
function getHistoricalAdjustmentCount() external view returns (uint256)
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

#### 历史调整查询

```solidity
// 获取历史调整记录数量
function getHistoricalAdjustmentCount() external view returns (uint256)

// 获取单条历史调整记录
function getHistoricalAdjustment(uint256 index) external view returns (
    uint256 timestamp,
    uint256 quarterlyMultiplier
)

// 获取特定等级的历史动态乘数
function getHistoricalDynamicMultiplier(uint256 index, uint8 level) external view returns (uint256)

// 获取所有历史调整记录
function getAllHistoricalAdjustments() external view returns (
    uint256[] memory timestamps,
    uint256[] memory quarterlyMultipliers
)

// 获取所有等级的历史动态乘数
function getHistoricalDynamicMultipliersForAllLevels(uint256 index) external view returns (uint256[6] memory multipliers)

// 获取完整的历史调整记录
function getCompleteHistoricalAdjustment(uint256 index) external view returns (
    uint256 timestamp,
    uint256 quarterlyMultiplier,
    uint256[6] memory dynamicMultipliers
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

**自动化历史记录**:
- **自动触发**: 季度调整执行时自动记录历史状态
- **无需手动**: 系统自动调用历史记录功能，无需管理员干预
- **完整快照**: 记录调整时刻的季度乘数和所有等级的动态乘数
- **容错设计**: 历史记录失败不会影响季度调整的正常执行

#### 对用户收益的影响

**实时影响**:
- **质押时**: 根据当前质押比例立即应用动态倍率
- **每日计算**: 每次计算收益时都会重新评估质押比例
- **即时调整**: 当质押比例跨越阈值时，收益立即调整

**历史影响与重新计算**:
- **历史记录**: 系统自动保存每次调整的完整状态快照
- **收益追溯**: 可根据历史参数重新计算用户在不同时期的收益
- **透明度**: 用户可查询历史调整记录，了解系统参数变化
- **合规性**: 提供完整的调整历史，满足审计和合规要求

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

// 合约地址 (Sepolia测试网已部署地址)
const STAKING_ADDRESS = "0x23983f63C7Eb0e920Fa73146293A51B215310Ac2";
const CONFIG_ADDRESS = "0x50fd41550bB5f6116a8b1330Cb50FAc41E658A69";
const READER_ADDRESS = "0x3243Fac23cfa3196525de9d1C28d3AD34E9783E8";
const CPNFT_ADDRESS = "0xcC63bf57EF4b4fE5635cF0745Ae7E2C75A63c7Ed";
const CPOP_TOKEN_ADDRESS = "0xA2d58d11c2752b010C2444fa2f795c6cf4cb76Bc";
const ACCOUNT_MANAGER_ADDRESS = "0x2E4f862Ba3ee6D84dd19ae9f002F5D8c0C5675ef";

// 合约ABI (需要导入实际的ABI)
const stakingContract = new ethers.Contract(STAKING_ADDRESS, StakingABI, provider);
const configContract = new ethers.Contract(CONFIG_ADDRESS, StakingConfigABI, provider);
const readerContract = new ethers.Contract(READER_ADDRESS, StakingReaderABI, provider);
const cpnftContract = new ethers.Contract(CPNFT_ADDRESS, CPNFTABI, provider);
const cpopTokenContract = new ethers.Contract(CPOP_TOKEN_ADDRESS, CPOPTokenABI, provider);
const accountManagerContract = new ethers.Contract(ACCOUNT_MANAGER_ADDRESS, AccountManagerABI, provider);
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
        // 检查批量大小限制 (最多50个)
        if (tokenIds.length > 50) {
            throw new Error("批量质押最多支持50个NFT");
        }
        
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

// 获取用户质押汇总信息 (新增功能)
async function getUserStakingSummary(userAddress) {
    const summary = await readerContract.getUserStakingSummary(userAddress);
    return {
        totalStakedCount: summary.totalStakedCount.toString(),
        totalClaimedRewards: ethers.utils.formatEther(summary.totalClaimedRewards),
        totalPendingRewards: ethers.utils.formatEther(summary.totalPendingRewards),
        levelStakingCounts: summary.levelStakingCounts.map(c => c.toString()),
        longestStakingDuration: Math.floor(summary.longestStakingDuration / 86400) + ' 天',
        totalEffectiveMultiplier: (summary.totalEffectiveMultiplier / 100).toFixed(2) + 'x'
    };
}

// 获取用户质押的NFT列表 (分页，新增功能)
async function getUserStakedNFTs(userAddress, offset = 0, limit = 10) {
    const result = await readerContract.getUserStakedNFTs(userAddress, offset, limit);
    return {
        nfts: result.nfts.map(nft => ({
            tokenId: nft.tokenId.toString(),
            level: nft.level,
            stakingDuration: Math.floor(nft.stakingDuration / 86400) + ' 天',
            pendingRewards: ethers.utils.formatEther(nft.pendingRewards),
            totalRewards: ethers.utils.formatEther(nft.totalRewards),
            effectiveMultiplier: (nft.effectiveMultiplier / 100).toFixed(2) + 'x'
        })),
        total: result.total.toString()
    };
}

// 获取用户收益统计 (新增功能)
async function getUserRewardStats(userAddress) {
    const stats = await readerContract.getUserRewardStats(userAddress);
    return {
        totalHistoricalRewards: ethers.utils.formatEther(stats.totalHistoricalRewards),
        totalPendingRewards: ethers.utils.formatEther(stats.totalPendingRewards),
        rewardsPerLevel: stats.rewardsPerLevel.map(r => ethers.utils.formatEther(r)),
        last24HoursRewards: ethers.utils.formatEther(stats.last24HoursRewards),
        averageDailyRewards: ethers.utils.formatEther(stats.averageDailyRewards)
    };
}

// 获取用户Combo汇总 (新增功能)
async function getUserComboSummary(userAddress) {
    const summary = await readerContract.getUserComboSummary(userAddress);
    return {
        currentComboCounts: summary.currentComboCounts.map(c => c.toString()),
        comboBonus: summary.comboBonus.map(b => (b / 100).toFixed(2) + '%'),
        nextComboThreshold: summary.nextComboThreshold.map(t => 
            t.toString() === ethers.constants.MaxUint256.toString() ? '已达最高' : t.toString()
        ),
        hasPendingCombo: summary.hasPendingCombo
    };
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

// 批量解质押NFT
async function batchUnstakeNFTs(tokenIds, signer) {
    try {
        // 检查批量大小限制 (最多50个)
        if (tokenIds.length > 50) {
            throw new Error("批量解质押最多支持50个NFT");
        }
        
        const unstakeTx = await stakingContract.connect(signer).batchUnstake(tokenIds);
        await unstakeTx.wait();
        
        console.log(`批量解质押 ${tokenIds.length} 个NFT成功`);
    } catch (error) {
        console.error("批量解质押失败:", error);
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

// 批量领取奖励
async function batchClaimRewards(tokenIds, signer) {
    try {
        // 检查批量大小限制 (最多50个)
        if (tokenIds.length > 50) {
            throw new Error("批量领取奖励最多支持50个NFT");
        }
        
        const claimTx = await stakingContract.connect(signer).batchClaimRewards(tokenIds);
        await claimTx.wait();
        
        console.log(`批量领取 ${tokenIds.length} 个NFT的奖励成功`);
    } catch (error) {
        console.error("批量领取奖励失败:", error);
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

// 查询历史调整记录
const historicalCount = await stakingReader.getHistoricalAdjustmentCount();
console.log(`历史调整记录数: ${historicalCount}`);

if (historicalCount > 0) {
    // 获取所有历史调整
    const allAdjustments = await stakingReader.getAllHistoricalAdjustments();
    console.log("历史调整记录:");
    
    for (let i = 0; i < allAdjustments.timestamps.length; i++) {
        const timestamp = allAdjustments.timestamps[i];
        const multiplier = allAdjustments.quarterlyMultipliers[i];
        const date = new Date(timestamp * 1000);
        
        console.log(`记录${i}: ${date.toISOString()} - 季度乘数 ${multiplier/10000}x`);
        
        // 获取该时刻的动态乘数
        const dynamicMultipliers = await stakingReader.getHistoricalDynamicMultipliersForAllLevels(i);
        console.log(`  动态乘数: [${dynamicMultipliers.map(m => m/10000 + 'x').join(', ')}]`);
    }
    
    // 获取完整的历史调整记录
    const completeRecord = await stakingReader.getCompleteHistoricalAdjustment(0);
    console.log("最新完整记录:", {
        timestamp: completeRecord.timestamp,
        quarterlyMultiplier: completeRecord.quarterlyMultiplier/10000 + 'x',
        dynamicMultipliers: completeRecord.dynamicMultipliers.map(m => m/10000 + 'x')
    });
}
```

### 📊 StakingReader 新功能说明 (v2.0)

#### 🆕 用户统计查询功能

StakingReader v2.0 新增了一系列用户级别的统计查询功能，专为前端展示优化：

**1. 用户质押汇总 (`getUserStakingSummary`)**
- 总质押数量
- 已领取总收益
- 待领取总收益
- 各等级质押数量统计
- 最长质押时间
- 综合收益倍率

```javascript
const summary = await readerContract.getUserStakingSummary(userAddress);
// 返回用户的完整质押概览数据
```

**2. 用户质押NFT列表 (`getUserStakedNFTs`)**
- 支持分页查询
- 包含每个NFT的详细信息：tokenId、等级、质押时长、收益、倍率
- 适合展示用户的质押资产列表

```javascript
const { nfts, total } = await readerContract.getUserStakedNFTs(userAddress, 0, 10);
// offset: 起始位置, limit: 每页数量
```

**3. 用户收益统计 (`getUserRewardStats`)**
- 历史总收益
- 当前待领取收益
- 各等级收益分布
- 最近24小时收益
- 平均每日收益

```javascript
const stats = await readerContract.getUserRewardStats(userAddress);
// 返回用户的收益统计数据
```

**4. 用户Combo汇总 (`getUserComboSummary`)**
- 各等级当前Combo数量
- 各等级Combo加成
- 距离下一个Combo的数量
- 待生效Combo状态

```javascript
const combo = await readerContract.getUserComboSummary(userAddress);
// 返回用户的Combo状态信息
```

#### 🔧 合约优化特性

**可升级性**
- 采用UUPS代理模式，支持未来功能扩展
- 合约依赖可动态更新

**大小优化**
- 移除冗余功能，合约大小从25KB优化到符合24KB限制
- 保留核心查询功能和新增的用户统计功能

**Gas优化**
- 所有函数都是view类型，查询不消耗gas
- 优化了内部计算逻辑，减少栈深度

### 📊 历史调整记录查询

StakingReader提供了完整的历史调整记录查询功能，用于追踪系统参数的历史变化：

#### 基础查询函数

```javascript
// 获取历史调整记录数量
const count = await stakingReader.getHistoricalAdjustmentCount();

// 获取单条历史调整记录
const [timestamp, quarterlyMultiplier] = await stakingReader.getHistoricalAdjustment(index);

// 获取特定等级的历史动态乘数
const dynamicMultiplier = await stakingReader.getHistoricalDynamicMultiplier(index, level);
```

#### 批量查询函数

```javascript
// 获取所有历史调整记录
const { timestamps, quarterlyMultipliers } = await stakingReader.getAllHistoricalAdjustments();

// 获取所有等级的历史动态乘数
const allLevelMultipliers = await stakingReader.getHistoricalDynamicMultipliersForAllLevels(index);

// 获取完整的历史调整记录（包含季度乘数和所有动态乘数）
const { timestamp, quarterlyMultiplier, dynamicMultipliers } = 
    await stakingReader.getCompleteHistoricalAdjustment(index);
```

#### 历史调整记录用途

1. **收益重新计算**: 根据历史参数准确计算用户在特定时期的收益
2. **透明度展示**: 向用户展示系统参数的历史变化
3. **数据分析**: 分析系统参数调整对用户收益的影响
4. **合规审计**: 提供完整的参数调整历史记录

## 🔗 相关链接

- [合约部署指南](../deploy/README.md)
- [测试文档](../../test/README.md)
- [Gas优化报告](../../reports/gas-optimization.md)
- [安全审计报告](../../audits/README.md)

## 📞 技术支持

如有技术问题，请联系开发团队或查看相关文档。

---

**版本**: v1.1.0  
**最后更新**: 2025年1月  
**兼容性**: Solidity ^0.8.19

## 🔄 版本更新记录

### v1.1.0 (2025年1月)
- ✅ **自动化历史记录**: 季度调整执行时自动记录历史状态，无需手动干预
- ✅ **完整历史查询**: StakingReader新增6个历史调整查询函数
- ✅ **容错设计**: 历史记录失败不影响季度调整正常执行
- ✅ **批处理功能**: 支持批量质押、批量解质押、批量领取奖励，提升用户体验
- ✅ **前端优化**: 提供批量查询和历史数据展示功能
- ✅ **文档完善**: 更新API文档和使用示例，包含完整的批处理函数示例

### v1.0.0 (2025年1月)
- ✅ **核心功能**: 多等级NFT质押系统
- ✅ **动态平衡**: 平台质押比例调控机制
- ✅ **季度调整**: 参数调整和历史记录功能
- ✅ **模块化架构**: 配置、逻辑、查询分离设计
- ✅ **批量操作**: 支持批量质押和奖励领取

## 🎉 部署状态

**✅ 已成功部署到Sepolia测试网**

- **Staking部署时间**: 2025-09-29T08:33:30.943Z
- **StakingReader部署时间**: 2025-09-29T08:36:26.468Z
- **网络**: Sepolia测试网 (Chain ID: 11155111)
- **状态**: 所有合约已部署并完成初始化
- **权限**: 所有必要权限已正确设置
- **可用性**: 系统已准备就绪，可开始质押操作

**合约部署清单**:
- ✅ Staking Contract (Proxy) - 主质押合约
- ✅ StakingConfig Contract - 配置管理合约  
- ✅ StakingReader Contract - 前端查询合约
- ✅ 所有依赖合约权限设置完成

**下一步**: 前端集成和用户测试
