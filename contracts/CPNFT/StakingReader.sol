// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./Staking.sol";
import "./StakingConfig.sol";
import "./CPNFT.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @title StakingReader - 前端查询合约
 * @dev 轻量级只读合约，专门用于前端查询功能
 * @notice 本合约提供优化的查询接口，减少Gas消耗，提高查询效率
 * @dev 所有函数都是view类型，不会修改状态，适合频繁调用
 * @dev 支持UUPS升级模式，可以通过升级添加新的查询功能
 */
contract StakingReader is 
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable
{
    // ============================================
    // 数据结构定义
    // ============================================
    
    /**
     * @dev 用户质押汇总信息
     * @notice 包含用户所有质押NFT的统计数据
     */

    struct UserStakingSummary {
        uint256 totalStakedCount;           // 总质押数量
        uint256 totalClaimedRewards;        // 已领取总收益
        uint256 totalPendingRewards;        // 待领取总收益
        uint256[6] levelStakingCounts;      // 每个等级的质押数量
        uint256 longestStakingDuration;     // 最长质押时间
        uint256 totalEffectiveMultiplier;   // 综合收益倍率
    }

    /**
     * @dev 单个质押NFT的详细信息
     * @notice 包含NFT的质押状态、收益和倍率信息
     */
    struct StakedNFTInfo {
        uint256 tokenId;                    // NFT ID
        uint8 level;                        // NFT 等级
        uint256 stakingDuration;            // 质押时长
        uint256 pendingRewards;             // 待领取收益
        uint256 totalRewards;               // 总收益
        uint256 effectiveMultiplier;        // 有效倍率
    }

    /**
     * @dev 用户收益统计信息
     * @notice 包含历史收益、待领取收益和每个等级的收益统计
     */
    struct UserRewardStats {
        uint256 totalHistoricalRewards;     // 历史总收益
        uint256 totalPendingRewards;        // 当前待领取总收益
        uint256[6] rewardsPerLevel;         // 每个等级的收益统计
        uint256 last24HoursRewards;         // 最近24小时收益
        uint256 averageDailyRewards;        // 平均每日收益
    }

    /**
     * @dev 用户Combo加成汇总信息
     * @notice 包含每个等级的Combo状态、加成和下一个阈值信息
     */
    struct UserComboSummary {
        uint256[6] currentComboCounts;      // 每个等级当前 Combo 数量
        uint256[6] comboBonus;              // 每个等级的 Combo 加成
        uint256[6] nextComboThreshold;      // 每个等级距离下一个 Combo 还需数量
        bool[6] hasPendingCombo;            // 每个等级是否有待生效的 Combo
    }
    
    /// @dev 主质押合约引用
    Staking public stakingContract;
    /// @dev 配置合约引用
    StakingConfig public configContract;
    /// @dev NFT合约引用
    CPNFT public cpnftContract;
    
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }
    
    /**
     * @notice 初始化合约
     * @dev 只能调用一次，设置合约依赖和所有者
     * @param _stakingContract 主质押合约地址
     * @param _configContract 配置合约地址
     * @param _cpnftContract NFT合约地址
     * @param _owner 合约所有者地址
     */
    function initialize(
        address _stakingContract,
        address _configContract,
        address _cpnftContract,
        address _owner
    ) public initializer {
        __UUPSUpgradeable_init();
        __Ownable_init(_owner);
        
        require(_stakingContract != address(0), "Staking contract address cannot be zero");
        require(_configContract != address(0), "Config contract address cannot be zero");
        require(_cpnftContract != address(0), "CPNFT contract address cannot be zero");
        
        stakingContract = Staking(_stakingContract);
        configContract = StakingConfig(_configContract);
        cpnftContract = CPNFT(_cpnftContract);
    }
    
    /**
     * @dev 实现UUPS升级授权检查
     * @param newImplementation 新的实现合约地址
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /**
     * @notice 更新合约依赖
     * @dev 只有合约所有者可以调用
     * @param _stakingContract 新的主质押合约地址（如果不更新则传0地址）
     * @param _configContract 新的配置合约地址（如果不更新则传0地址）
     * @param _cpnftContract 新的NFT合约地址（如果不更新则传0地址）
     */
    function updateContracts(
        address _stakingContract,
        address _configContract,
        address _cpnftContract
    ) external onlyOwner {
        if (_stakingContract != address(0)) {
            stakingContract = Staking(_stakingContract);
        }
        if (_configContract != address(0)) {
            configContract = StakingConfig(_configContract);
        }
        if (_cpnftContract != address(0)) {
            cpnftContract = CPNFT(_cpnftContract);
        }
    }
    
    // ============================================
    // 基础查询函数
    // ============================================
    
    /**
     * @notice 这些函数提供基本的质押信息查询功能
     * @dev 所有函数都是view类型，不消耗gas
     */
    
    /**
     * @notice 获取指定NFT的质押详情
     * @dev 返回包含实时待领取收益的质押信息
     * @param tokenId NFT的ID
     * @return 质押详情，包括所有者、质押时间、收益等信息
     */
    function getStakeDetails(uint256 tokenId) external view returns (Staking.StakeInfo memory) {
        (
            address owner,
            uint256 tokenId_,
            uint8 level,
            uint256 stakeTime,
            uint256 lastClaimTime,
            bool isActive,
            uint256 totalRewards,
            uint256 pendingRewards
        ) = stakingContract.stakes(tokenId);
        
        // Calculate real-time pending rewards if stake is active
        uint256 realTimePendingRewards = pendingRewards;
        if (isActive && block.timestamp > lastClaimTime) {
            realTimePendingRewards = stakingContract.calculatePendingRewards(tokenId);
        }
        
        return Staking.StakeInfo({
            owner: owner,
            tokenId: tokenId_,
            level: level,
            stakeTime: stakeTime,
            lastClaimTime: lastClaimTime,
            isActive: isActive,
            totalRewards: totalRewards,
            pendingRewards: realTimePendingRewards
        });
    }
    
    /**
     * @notice 获取指定NFT的待领取收益
     * @dev 计算从上次领取到现在的累积收益，包含所有加成和衰减
     * @param tokenId NFT的ID
     * @return 待领取的收益数量（以wei为单位）
     */
    function getPendingRewards(uint256 tokenId) external view returns (uint256) {
        // Only get the values we need
        (, , uint8 level, uint256 stakeTime, uint256 lastClaimTime, bool isActive, , ) = stakingContract.stakes(tokenId);
        
        if (!isActive) return 0;
        
        // Calculate pending rewards using internal logic
        uint256 totalDays = (block.timestamp - lastClaimTime) / 1 days;
        if (totalDays == 0) return 0;
        
        uint256 baseReward = configContract.getDailyReward(level);
        
        // Apply decay
        uint256 decayedReward = _calculateDecayedReward(baseReward, level, stakeTime);
        
        // Apply quarterly adjustment
        uint256 quarterlyMultiplier = configContract.getQuarterlyMultiplier();
        uint256 adjustedReward = decayedReward * quarterlyMultiplier / 10000;
        
        return adjustedReward * totalDays;
    }
    
    /**
     * @notice 获取详细的收益计算明细
     * @dev 返回收益计算的各个组成部分，包括基础收益、衰减后收益、各种加成等
     * @param tokenId NFT的ID
     * @return baseReward 基础收益（未应用任何加成）
     * @return decayedReward 衰减后的收益
     * @return comboBonus Combo加成比例（以基点表示，10000=100%）
     * @return dynamicMultiplier 动态倍率（以基点表示，10000=100%）
     * @return finalReward 最终收益（应用所有加成后）
     */
    function getDetailedRewardCalculation(uint256 tokenId) external view returns (
        uint256 baseReward,
        uint256 decayedReward,
        uint256 comboBonus,
        uint256 dynamicMultiplier,
        uint256 finalReward
    ) {
        // Only get the values we need
        (address owner, , uint8 level, uint256 stakeTime, uint256 lastClaimTime, , , ) = stakingContract.stakes(tokenId);
        
        baseReward = configContract.getDailyReward(level);
        decayedReward = _calculateDecayedReward(baseReward, level, stakeTime);
        
        comboBonus = _calculateComboBonus(owner, level);
        dynamicMultiplier = _calculateDynamicMultiplier(level);
        
        uint256 totalDays = (block.timestamp - lastClaimTime) / 1 days;
        uint256 quarterlyMultiplier = configContract.getQuarterlyMultiplier();
        
        finalReward = decayedReward * (10000 + comboBonus) / 10000;
        finalReward = finalReward * dynamicMultiplier * quarterlyMultiplier / (10000 * 10000);
        finalReward = finalReward * totalDays;
    }
    
    /**
     * @notice 获取用户特定等级的Combo信息
     * @dev 返回当前Combo状态和可能的下一个Combo阈值
     * @param user 用户地址
     * @param level NFT等级（1-6: C, B, A, S, SS, SSS）
     * @return count 当前该等级的NFT数量
     * @return currentBonus 当前生效的Combo加成（以基点表示）
     * @return thresholds Combo阈值数组（达到这些数量可获得加成）
     * @return bonuses 对应的加成比例数组（以基点表示）
     */
    function getComboInfo(address user, uint8 level) external view returns (
        uint256 count,
        uint256 currentBonus,
        uint256[] memory thresholds,
        uint256[] memory bonuses
    ) {
        count = stakingContract.userLevelCounts(user, level);
        currentBonus = stakingContract.getEffectiveComboBonus(user, level);
        
        uint256[3] memory configThresholds = configContract.getComboThresholds();
        uint256[3] memory configBonuses = configContract.getComboBonuses();
        
        thresholds = new uint256[](3);
        bonuses = new uint256[](3);
        
        for (uint256 i = 0; i < 3; i++) {
            thresholds[i] = configThresholds[i];
            bonuses[i] = configBonuses[i];
        }
    }
    
    /**
     * @notice 获取用户特定等级的详细Combo状态
     * @dev 包含待生效的Combo变更信息和生效倒计时
     * @param user 用户地址
     * @param level NFT等级（1-6: C, B, A, S, SS, SSS）
     * @return level_ NFT等级
     * @return count 当前该等级的NFT数量
     * @return effectiveFrom Combo变更生效时间（时间戳）
     * @return bonus Combo加成比例（以基点表示）
     * @return isPending 是否有待生效的变更
     * @return timeUntilEffective 距离生效还需要的时间（秒）
     */
    function getComboStatus(address user, uint8 level) external view returns (
        uint256 level_,
        uint256 count,
        uint256 effectiveFrom,
        uint256 bonus,
        bool isPending,
        uint256 timeUntilEffective
    ) {
        Staking.ComboStatus memory status = stakingContract.getComboStatus(user, level);
        
        level_ = status.level;
        count = status.count;
        effectiveFrom = status.effectiveFrom;
        bonus = status.bonus;
        isPending = status.isPending;
        
        if (status.isPending && status.effectiveFrom > block.timestamp) {
            timeUntilEffective = status.effectiveFrom - block.timestamp;
        } else {
            timeUntilEffective = 0;
        }
    }
    
    /**
     * @notice 获取各等级的统计数据
     * @dev 返回每个等级的质押数量、总供应量、质押比例和动态倍率
     * @return stakedCounts 每个等级的质押数量数组[NORMAL, C, B, A, S, SS, SSS]
     * @return supplies 每个等级的总供应量数组
     * @return stakingRatios 每个等级的质押比例数组（以基点表示，10000=100%）
     * @return dynamicMultipliers 每个等级的动态倍率数组（以基点表示，10000=100%）
     */
    function getLevelStats() external view returns (
        uint256[7] memory stakedCounts,
        uint256[7] memory supplies,
        uint256[7] memory stakingRatios,
        uint256[7] memory dynamicMultipliers
    ) {
        for (uint8 i = 0; i < 7; i++) {
            stakedCounts[i] = stakingContract.totalStakedPerLevel(i);
            
            // Get supply from CPNFT contract (convert index to enum)
            if (i == 0) supplies[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.NORMAL);
            else if (i == 1) supplies[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.C);
            else if (i == 2) supplies[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.B);
            else if (i == 3) supplies[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.A);
            else if (i == 4) supplies[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.S);
            else if (i == 5) supplies[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.SS);
            else if (i == 6) supplies[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.SSS);
            
            if (supplies[i] > 0) {
                stakingRatios[i] = (stakedCounts[i] * 10000) / supplies[i];
            }
            
            dynamicMultipliers[i] = _calculateDynamicMultiplier(i);
        }
    }
    
    /**
     * @notice 获取平台基础统计数据
     * @dev 返回简化版的平台统计信息，主要包含质押数量和供应量
     * @return staked 每个等级的质押数量数组[NORMAL, C, B, A, S, SS, SSS]
     * @return supply 每个等级的总供应量数组
     */
    function getPlatformStats() external view returns (uint256[7] memory staked, uint256[7] memory supply) {
        for (uint256 i = 0; i < 7; i++) {
            staked[i] = stakingContract.totalStakedPerLevel(uint8(i));
            
            // Get supply from CPNFT contract (convert index to enum)
            if (i == 0) supply[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.NORMAL);
            else if (i == 1) supply[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.C);
            else if (i == 2) supply[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.B);
            else if (i == 3) supply[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.A);
            else if (i == 4) supply[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.S);
            else if (i == 5) supply[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.SS);
            else if (i == 6) supply[i] = cpnftContract.getLevelSupply(CPNFT.NFTLevel.SSS);
        }
    }
    
    // ============================================
    // INTERNAL CALCULATION FUNCTIONS
    // ============================================

    /**
     * @dev 获取用户质押的所有NFT ID列表
     * @param user 用户地址
     * @return 用户质押的NFT ID数组
     * @notice 使用try-catch遍历数组，因为无法直接获取数组长度
     */
    function _getUserTokenIds(address user) internal view returns (uint256[] memory) {
        uint256 count = 0;
        // First count how many tokens the user has staked
        for (uint256 i = 0; ; i++) {
            try stakingContract.userStakes(user, i) returns (uint256) {
                count++;
            } catch {
                break;
            }
        }
        
        // Now create the array and fill it
        uint256[] memory tokenIds = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            tokenIds[i] = stakingContract.userStakes(user, i);
        }
        
        return tokenIds;
    }
    
    /**
     * @dev 获取NFT的质押信息
     * @param tokenId NFT的ID
     * @return 质押信息结构体，包含所有者、等级、时间等信息
     * @notice 从主质押合约获取信息并转换为结构体格式
     */
    function _getStakeInfo(uint256 tokenId) internal view returns (Staking.StakeInfo memory) {
        (
            address owner,
            uint256 tokenId_,
            uint8 level,
            uint256 stakeTime,
            uint256 lastClaimTime,
            bool isActive,
            uint256 totalRewards,
            uint256 pendingRewards
        ) = stakingContract.stakes(tokenId);
        
        return Staking.StakeInfo({
            owner: owner,
            tokenId: tokenId_,
            level: level,
            stakeTime: stakeTime,
            lastClaimTime: lastClaimTime,
            isActive: isActive,
            totalRewards: totalRewards,
            pendingRewards: pendingRewards
        });
    }
    
    /**
     * @dev 计算衰减后的收益
     * @param baseReward 基础收益（未衰减）
     * @param level NFT等级
     * @param stakeTime 质押开始时间
     * @return 应用衰减后的收益
     * @notice 根据质押时长和配置计算衰减后的收益，考虑最大衰减限制
     */
    function _calculateDecayedReward(
        uint256 baseReward,
        uint8 level,
        uint256 stakeTime
    ) internal view returns (uint256) {
        uint256 stakingDays = (block.timestamp - stakeTime) / 1 days;
        uint256 decayInterval = configContract.getDecayInterval(level);
        uint256 decayRate = configContract.getDecayRate(level);
        uint256 maxDecay = configContract.getMaxDecayRate(level);
        
        if (decayInterval == 0 || stakingDays < decayInterval) {
            return baseReward;
        }
        
        uint256 decayCycles = stakingDays / decayInterval;
        uint256 totalDecay = decayCycles * decayRate;
        
        if (totalDecay > maxDecay) {
            totalDecay = maxDecay;
        }
        
        return baseReward * (10000 - totalDecay) / 10000;
    }
    
    /**
     * @dev 计算用户特定等级的Combo加成
     * @param user 用户地址
     * @param level NFT等级
     * @return Combo加成比例（以基点表示，10000=100%）
     * @notice 根据用户持有的NFT数量和配置的阈值计算Combo加成
     */
    function _calculateComboBonus(address user, uint8 level) internal view returns (uint256) {
        uint256 count = stakingContract.userLevelCounts(user, level);
        
        // Get combo config
        uint256[3] memory thresholds = configContract.getComboThresholds();
        uint256[3] memory bonuses = configContract.getComboBonuses();
        
        // Check combo thresholds in reverse order (highest first)
        for (uint256 i = thresholds.length; i > 0; i--) {
            uint256 threshold = thresholds[i - 1];
            
            if (count >= threshold) {
                return bonuses[i - 1];
            }
        }
        
        return 0;
    }
    
    /**
     * @dev 计算动态平衡倍率
     * @param level NFT等级
     * @return 动态倍率（以基点表示，10000=100%）
     * @notice 根据质押比例计算动态倍率，用于平衡质押量
     * - 高质押率（>60%）时返回较低倍率（0.8x）
     * - 低质押率（<10%）时返回较高倍率（1.05x）
     * - 正常范围内返回标准倍率（1.0x）
     */
    /**
     * @dev 从NFT合约获取特定等级的总供应量
     * @param level NFT等级（1-6: C, B, A, S, SS, SSS）
     * @return 指定等级的NFT总供应量
     * @notice 将数字等级转换为对应的枚举值后查询
     */
    function _getLevelSupply(uint8 level) internal view returns (uint256) {
        if (level == 1) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.C);
        if (level == 2) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.B);
        if (level == 3) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.A);
        if (level == 4) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.S);
        if (level == 5) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.SS);
        if (level == 6) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.SSS);
        return 0; // Invalid level
    }
    
    function _calculateDynamicMultiplier(uint8 level) internal view returns (uint256) {
        uint256 staked = stakingContract.totalStakedPerLevel(level);
        uint256 supply = _getLevelSupply(level);
        
        if (supply == 0) return 10000; // 1.0x
        
        uint256 stakingRatio = (staked * 10000) / supply;
        
        (
            uint256 highStakeThreshold,
            uint256 lowStakeThreshold,
            uint256 highStakeMultiplier,
            uint256 lowStakeMultiplier,
            // Unused: uint256 quarterlyAdjustmentMax
        ) = configContract.getDynamicConfig();
        
        if (stakingRatio > highStakeThreshold) {
            return highStakeMultiplier;
        } else if (stakingRatio < lowStakeThreshold) {
            return lowStakeMultiplier;
        }
        
        return 10000; // 1.0x
    }
    
    /**
     * @dev 计算连续质押奖励
     * @param stakingDays 质押天数
     * @return 连续质押奖励（以基点表示，10000=100%）
     * @notice 根据质押天数计算连续质押奖励
     * - 达到第一个阈值（30天）获得10%奖励
     * - 达到第二个阈值（90天）获得30%奖励
     */
    function _calculateContinuousBonus(uint256 stakingDays) internal view returns (uint256) {
        uint256[2] memory thresholds = configContract.getContinuousThresholds();
        uint256[2] memory bonuses = configContract.getContinuousBonuses();
        
        // Check thresholds in reverse order (highest first)
        for (uint256 i = thresholds.length; i > 0; i--) {
            if (stakingDays >= thresholds[i - 1]) {
                return bonuses[i - 1];
            }
        }
        
        return 0; // No bonus
    }
    
    // ============================================
    // CONFIGURATION QUERIES
    // ============================================
    
    /**
     * @dev Get all configuration data
     */
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
    ) {
        minStakeDays = configContract.getMinStakeDays();
        earlyWithdrawPenalty = configContract.getEarlyWithdrawPenalty();
        quarterlyMultiplier = configContract.getQuarterlyMultiplier();
        
        // Daily rewards (excluding NORMAL level)
        for (uint8 i = 1; i < 7; i++) {
            dailyRewards[i-1] = configContract.getDailyReward(i);
            decayIntervals[i-1] = configContract.getDecayInterval(i);
            decayRates[i-1] = configContract.getDecayRate(i);
            maxDecayRates[i-1] = configContract.getMaxDecayRate(i);
        }
        
        comboThresholds = configContract.getComboThresholds();
        comboBonuses = configContract.getComboBonuses();
        comboMinDays = configContract.getComboMinDays();
        continuousThresholds = configContract.getContinuousThresholds();
        continuousBonuses = configContract.getContinuousBonuses();
    }
    
    /**
     * @dev Get platform statistics
     */
    function getPlatformStatistics() external view returns (
        uint256 totalStakedNFTs,
        uint256[7] memory stakedPerLevel,
        uint256[7] memory supplyPerLevel,
        uint256[7] memory stakingRatios,
        uint256[7] memory dynamicMultipliers,
        uint256 totalRewardsToday
    ) {
        totalStakedNFTs = stakingContract.totalStakedCount();
        (stakedPerLevel, supplyPerLevel, stakingRatios, dynamicMultipliers) = this.getLevelStats();
        
        // Calculate estimated daily rewards
        for (uint8 level = 1; level < 7; level++) {
            uint256 dailyReward = configContract.getDailyReward(level);
            uint256 stakedCount = stakedPerLevel[level];
            uint256 multiplier = dynamicMultipliers[level];
            uint256 quarterlyMultiplier = configContract.getQuarterlyMultiplier();
            
            totalRewardsToday += (dailyReward * stakedCount * multiplier * quarterlyMultiplier) / (10000 * 10000);
        }
    }
    
    /**
     * @dev Get contract versions
     */
    function getVersions() external view returns (
        string memory stakingVersion,
        string memory configVersion
    ) {
        stakingVersion = stakingContract.version();
        configVersion = configContract.version();
    }
    
    // ============================================
    // HISTORICAL ADJUSTMENT QUERIES
    // ============================================
    
    /**
     * @dev Get historical adjustment record
     * @param index Index of the adjustment record
     * @return timestamp When the adjustment was recorded
     * @return quarterlyMultiplier The quarterly multiplier at that time
     */
    function getHistoricalAdjustment(uint256 index) external view returns (
        uint256 timestamp,
        uint256 quarterlyMultiplier
    ) {
        return stakingContract.getHistoricalAdjustment(index);
    }
    
    /**
     * @dev Get historical dynamic multiplier for a specific level and time
     * @param index Index of the adjustment record
     * @param level NFT level (1-6)
     * @return The dynamic multiplier for that level at that time
     */
    function getHistoricalDynamicMultiplier(uint256 index, uint8 level) external view returns (uint256) {
        return stakingContract.getHistoricalDynamicMultiplier(index, level);
    }
    
    /**
     * @dev Get number of historical adjustments
     * @return Number of adjustment records
     */
    function getHistoricalAdjustmentCount() external view returns (uint256) {
        return stakingContract.getHistoricalAdjustmentCount();
    }
    
    /**
     * @dev Get all historical adjustments for frontend display
     * @return timestamps Array of timestamps
     * @return quarterlyMultipliers Array of quarterly multipliers
     */
    function getAllHistoricalAdjustments() external view returns (
        uint256[] memory timestamps,
        uint256[] memory quarterlyMultipliers
    ) {
        uint256 count = stakingContract.getHistoricalAdjustmentCount();
        
        timestamps = new uint256[](count);
        quarterlyMultipliers = new uint256[](count);
        
        for (uint256 i = 0; i < count; i++) {
            (timestamps[i], quarterlyMultipliers[i]) = stakingContract.getHistoricalAdjustment(i);
        }
    }
    
    /**
     * @dev Get historical dynamic multipliers for all levels at a specific time
     * @param index Index of the adjustment record
     * @return multipliers Array of dynamic multipliers for levels 1-6
     */
    function getHistoricalDynamicMultipliersForAllLevels(uint256 index) external view returns (uint256[6] memory multipliers) {
        for (uint8 level = 1; level <= 6; level++) {
            multipliers[level - 1] = stakingContract.getHistoricalDynamicMultiplier(index, level);
        }
    }
    
    /**
     * @dev Get complete historical adjustment data for frontend
     * @param index Index of the adjustment record
     * @return timestamp When the adjustment was recorded
     * @return quarterlyMultiplier The quarterly multiplier at that time
     * @return dynamicMultipliers Array of dynamic multipliers for levels 1-6
     */
    function getCompleteHistoricalAdjustment(uint256 index) external view returns (
        uint256 timestamp,
        uint256 quarterlyMultiplier,
        uint256[6] memory dynamicMultipliers
    ) {
        (timestamp, quarterlyMultiplier) = stakingContract.getHistoricalAdjustment(index);
        
        for (uint8 level = 1; level <= 6; level++) {
            dynamicMultipliers[level - 1] = stakingContract.getHistoricalDynamicMultiplier(index, level);
        }
    }
    
    /**
     * @dev Get historical quarterly multiplier for a specific timestamp
     * @param timestamp The timestamp to get the quarterly multiplier for
     * @return The quarterly multiplier that was active at that timestamp
     */
    function getHistoricalQuarterlyMultiplier(uint256 timestamp) external view returns (uint256) {
        return stakingContract.getHistoricalQuarterlyMultiplier(timestamp);
    }
    
    /**
     * @dev Get detailed reward breakdown for a specific day
     * @param tokenId The NFT token ID
     * @param dayOffset Days from stake start (0 = first day)
     * @return baseReward The base reward before adjustments
     * @return decayedReward The reward after decay
     * @return quarterlyMultiplier The quarterly multiplier applied
     * @return dynamicMultiplier The dynamic multiplier applied
     * @return finalReward The final reward for this day
     */
    function getDailyRewardBreakdown(uint256 tokenId, uint256 dayOffset) external view returns (
        uint256 baseReward,
        uint256 decayedReward,
        uint256 quarterlyMultiplier,
        uint256 dynamicMultiplier,
        uint256 finalReward
    ) {
        return stakingContract.getDailyRewardBreakdown(tokenId, dayOffset);
    }
    
    /**
     * @dev Get comprehensive reward calculation for a staked NFT
     * Shows how dynamic adjustments affect rewards over time
     * @param tokenId The NFT token ID
     * @return totalDays Total staking days
     * @return totalRewards Total rewards earned
     * @return baseRewards Array of base rewards for each day
     * @return decayedRewards Array of decayed rewards for each day
     * @return quarterlyMultipliers Array of quarterly multipliers for each day
     * @return dynamicMultipliers Array of dynamic multipliers for each day
     * @return dailyRewards Array of final daily rewards
     */
    // ============================================
    // 用户统计查询函数
    // ============================================
    
    /**
     * @notice 这些函数提供用户级别的统计数据查询功能
     * @dev 包含质押汇总、NFT列表、收益统计和Combo汇总等功能
     */

    /**
     * @notice 获取用户的质押汇总信息
     * @dev 返回用户所有质押NFT的统计数据，包括数量、收益和倍率等
     * @param user 用户地址
     * @return summary 用户质押汇总信息，包括：
     *         - totalStakedCount: 总质押数量
     *         - totalClaimedRewards: 已领取总收益
     *         - totalPendingRewards: 待领取总收益
     *         - levelStakingCounts: 每个等级的质押数量
     *         - longestStakingDuration: 最长质押时间
     *         - totalEffectiveMultiplier: 综合收益倍率
     */
    /**
     * @notice 获取用户质押的NFT列表（支持分页）
     * @dev 返回带有详细信息的NFT列表，包括质押时长、收益等
     * @param user 用户地址
     * @param offset 分页起始位置
     * @param limit 每页数量限制
     * @return nfts NFT详细信息数组，每个元素包含：
     *         - tokenId: NFT的ID
     *         - level: NFT等级
     *         - stakingDuration: 质押时长
     *         - pendingRewards: 待领取收益
     *         - totalRewards: 总收益
     *         - effectiveMultiplier: 有效倍率
     * @return total 用户质押的NFT总数
     */
    /**
     * @notice 获取用户的详细收益统计
     * @dev 返回用户的收益统计信息，包括历史收益、待领取收益和每个等级的收益等
     * @param user 用户地址
     * @return stats 用户收益统计信息，包括：
     *         - totalHistoricalRewards: 历史总收益
     *         - totalPendingRewards: 当前待领取总收益
     *         - rewardsPerLevel: 每个等级的收益统计
     *         - last24HoursRewards: 最近24小时收益
     *         - averageDailyRewards: 平均每日收益
     */
    /**
     * @notice 获取用户的Combo加成汇总信息
     * @dev 返回用户所有等级的Combo状态，包括当前加成和下一个阈值信息
     * @param user 用户地址
     * @return summary 用户Combo汇总信息，包括：
     *         - currentComboCounts: 每个等级当前的Combo数量
     *         - comboBonus: 每个等级的Combo加成
     *         - nextComboThreshold: 每个等级距离下一个Combo还需数量
     *         - hasPendingCombo: 每个等级是否有待生效的Combo
     */
    function getUserComboSummary(address user) external view returns (UserComboSummary memory summary) {
        // Get combo configuration
        uint256[3] memory thresholds = configContract.getComboThresholds();
        
        // Process each level
        for (uint8 level = 1; level <= 6; level++) {
            // Get current count and status
            summary.currentComboCounts[level - 1] = stakingContract.userLevelCounts(user, level);
            summary.comboBonus[level - 1] = stakingContract.getEffectiveComboBonus(user, level);
            
            // Get combo status to check for pending changes
            Staking.ComboStatus memory status = stakingContract.getComboStatus(user, level);
            summary.hasPendingCombo[level - 1] = status.isPending;
            
            // Calculate next threshold
            uint256 currentCount = summary.currentComboCounts[level - 1];
            summary.nextComboThreshold[level - 1] = type(uint256).max; // Default to max if no next threshold
            
            // Find next threshold
            for (uint256 i = 0; i < thresholds.length; i++) {
                if (currentCount < thresholds[i]) {
                    summary.nextComboThreshold[level - 1] = thresholds[i] - currentCount;
                    break;
                }
            }
        }
    }

    function getUserRewardStats(address user) external view returns (UserRewardStats memory stats) {
        uint256[] memory tokenIds = _getUserTokenIds(user);
        if (tokenIds.length == 0) {
            return stats;
        }
        
        // Get current timestamp for time-based calculations
        uint256 currentTime = block.timestamp;
        uint256 oneDayAgo = currentTime - 1 days;
        
        // Initialize variables for average calculation
        uint256 totalStakingDays = 0;
        
        // Process each staked NFT
        for (uint256 i = 0; i < tokenIds.length; i++) {
            Staking.StakeInfo memory stakeInfo = _getStakeInfo(tokenIds[i]);
            
            // Skip if not active
            if (!stakeInfo.isActive) continue;
            
            // Update total rewards
            stats.totalHistoricalRewards += stakeInfo.totalRewards;
            uint256 pendingRewards = stakingContract.calculatePendingRewards(tokenIds[i]);
            stats.totalPendingRewards += pendingRewards;
            
            // Update level-specific rewards
            if (stakeInfo.level >= 1 && stakeInfo.level <= 6) {
                stats.rewardsPerLevel[stakeInfo.level - 1] += stakeInfo.totalRewards + pendingRewards;
            }
            
            // Calculate last 24 hours rewards
            if (stakeInfo.lastClaimTime < oneDayAgo) {
                // If last claim was more than 24 hours ago, calculate rewards for last 24 hours
                uint256 dailyReward = configContract.getDailyReward(stakeInfo.level);
                uint256 comboBonus = stakingContract.getEffectiveComboBonus(user, stakeInfo.level);
                uint256 dynamicMultiplier = _calculateDynamicMultiplier(stakeInfo.level);
                uint256 quarterlyMultiplier = configContract.getQuarterlyMultiplier();
                
                uint256 effectiveDaily = dailyReward * (10000 + comboBonus) * dynamicMultiplier * quarterlyMultiplier / (10000 * 10000 * 10000);
                stats.last24HoursRewards += effectiveDaily;
            }
            
            // Update total staking days for average calculation
            totalStakingDays += (currentTime - stakeInfo.stakeTime) / 1 days;
        }
        
        // Calculate average daily rewards
        if (totalStakingDays > 0) {
            stats.averageDailyRewards = (stats.totalHistoricalRewards + stats.totalPendingRewards) / totalStakingDays;
        }
    }

    function getUserStakedNFTs(
        address user,
        uint256 offset,
        uint256 limit
    ) external view returns (
        StakedNFTInfo[] memory nfts,
        uint256 total
    ) {
        uint256[] memory tokenIds = _getUserTokenIds(user);
        total = tokenIds.length;
        
        if (total == 0 || offset >= total) {
            return (new StakedNFTInfo[](0), total);
        }
        
        // Calculate actual number of items to return
        uint256 end = offset + limit;
        if (end > total) {
            end = total;
        }
        uint256 resultCount = end - offset;
        
        // Initialize return array
        nfts = new StakedNFTInfo[](resultCount);
        uint256 resultIndex = 0;
        
        // Process each NFT in the requested range
        for (uint256 i = offset; i < end; i++) {
            uint256 tokenId = tokenIds[i];
            Staking.StakeInfo memory stakeInfo = _getStakeInfo(tokenId);
            
            // Skip if not active
            if (!stakeInfo.isActive) continue;
            
            // Calculate effective multiplier
            uint256 comboBonus = stakingContract.getEffectiveComboBonus(user, stakeInfo.level);
            uint256 dynamicMultiplier = _calculateDynamicMultiplier(stakeInfo.level);
            uint256 quarterlyMultiplier = configContract.getQuarterlyMultiplier();
            uint256 effectiveMultiplier = (10000 + comboBonus) * dynamicMultiplier * quarterlyMultiplier / (10000 * 10000);
            
            // Populate NFT info
            nfts[resultIndex] = StakedNFTInfo({
                tokenId: tokenId,
                level: stakeInfo.level,
                stakingDuration: block.timestamp - stakeInfo.stakeTime,
                pendingRewards: stakingContract.calculatePendingRewards(tokenId),
                totalRewards: stakeInfo.totalRewards,
                effectiveMultiplier: effectiveMultiplier
            });
            
            resultIndex++;
        }
        
        // If some NFTs were skipped because they were inactive, resize the array
        if (resultIndex < resultCount) {
            assembly {
                mstore(nfts, resultIndex)
            }
        }
    }

    function getUserStakingSummary(address user) external view returns (UserStakingSummary memory summary) {
        // Get user's staked token IDs
        uint256[] memory tokenIds = _getUserTokenIds(user);
        summary.totalStakedCount = tokenIds.length;
        
        if (summary.totalStakedCount == 0) {
            return summary;
        }

        // Initialize variables
        uint256 longestDuration = 0;
        uint256 totalMultiplier = 0;
        
        // Process each staked NFT
        for (uint256 i = 0; i < tokenIds.length; i++) {
            Staking.StakeInfo memory stakeInfo = _getStakeInfo(tokenIds[i]);
            
            // Skip if not active
            if (!stakeInfo.isActive) continue;
            
            // Update level counts
            if (stakeInfo.level >= 1 && stakeInfo.level <= 6) {
                summary.levelStakingCounts[stakeInfo.level - 1]++;
            }
            
            // Update rewards
            summary.totalClaimedRewards += stakeInfo.totalRewards;
            summary.totalPendingRewards += stakingContract.calculatePendingRewards(tokenIds[i]);
            
            // Calculate and update duration
            uint256 duration = block.timestamp - stakeInfo.stakeTime;
            if (duration > longestDuration) {
                longestDuration = duration;
            }
            
            // Calculate effective multiplier
            uint256 comboBonus = stakingContract.getEffectiveComboBonus(user, stakeInfo.level);
            uint256 dynamicMultiplier = _calculateDynamicMultiplier(stakeInfo.level);
            uint256 quarterlyMultiplier = configContract.getQuarterlyMultiplier();
            
            uint256 effectiveMultiplier = (10000 + comboBonus) * dynamicMultiplier * quarterlyMultiplier / (10000 * 10000);
            totalMultiplier += effectiveMultiplier;
        }
        
        // Set final values
        summary.longestStakingDuration = longestDuration;
        summary.totalEffectiveMultiplier = summary.totalStakedCount > 0 ? 
            totalMultiplier / summary.totalStakedCount : 0;
    }

    /**
     * @dev Get comprehensive reward calculation for a staked NFT
     * Shows how dynamic adjustments affect rewards over time
     * @param tokenId The NFT token ID
     * @return totalDays Total staking days
     * @return totalRewards Total rewards earned
     * @return baseRewards Array of base rewards for each day
     * @return decayedRewards Array of decayed rewards for each day
     * @return quarterlyMultipliers Array of quarterly multipliers for each day
     * @return dynamicMultipliers Array of dynamic multipliers for each day
     * @return dailyRewards Array of final daily rewards
     */
    function getComprehensiveRewardCalculation(uint256 tokenId) external view returns (
        uint256 totalDays,
        uint256 totalRewards,
        uint256[] memory baseRewards,
        uint256[] memory decayedRewards,
        uint256[] memory quarterlyMultipliers,
        uint256[] memory dynamicMultipliers,
        uint256[] memory dailyRewards
    ) {
        // Only get the values we need
        (, , uint8 level, uint256 stakeTime, uint256 lastClaimTime, bool isActive, , uint256 pendingRewards) = stakingContract.stakes(tokenId);
        
        require(isActive, "NFT is not staked");
        
        totalDays = (block.timestamp - lastClaimTime) / 1 days;
        totalRewards = pendingRewards;
        
        if (totalDays == 0) {
            return (0, 0, new uint256[](0), new uint256[](0), new uint256[](0), new uint256[](0), new uint256[](0));
        }
        
        baseRewards = new uint256[](totalDays);
        decayedRewards = new uint256[](totalDays);
        quarterlyMultipliers = new uint256[](totalDays);
        dynamicMultipliers = new uint256[](totalDays);
        dailyRewards = new uint256[](totalDays);
        
        // Calculate breakdown for each day
        for (uint256 day = 0; day < totalDays; day++) {
            (
                baseRewards[day],
                decayedRewards[day],
                quarterlyMultipliers[day],
                dynamicMultipliers[day],
                dailyRewards[day]
            ) = stakingContract.getDailyRewardBreakdown(tokenId, (block.timestamp - stakeTime) / 1 days - (totalDays - 1 - day));
        }
    }
}