// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./Staking.sol";
import "./StakingConfig.sol";
import "./CPNFT.sol";

/**
 * @title StakingReader - Frontend Query Contract
 * @dev Lightweight read-only contract for frontend queries
 */
contract StakingReader {
    
    Staking public immutable stakingContract;
    StakingConfig public immutable configContract;
    CPNFT public immutable cpnftContract;
    
    constructor(address _stakingContract, address _configContract, address _cpnftContract) {
        stakingContract = Staking(_stakingContract);
        configContract = StakingConfig(_configContract);
        cpnftContract = CPNFT(_cpnftContract);
    }
    
    // ============================================
    // BASIC VIEW FUNCTIONS
    // ============================================
    
    /**
     * @dev Get stake details for a specific token
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
     * @dev Get pending rewards for a specific token
     */
    function getPendingRewards(uint256 tokenId) external view returns (uint256) {
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
     * @dev Get detailed reward calculation breakdown
     */
    function getDetailedRewardCalculation(uint256 tokenId) external view returns (
        uint256 baseReward,
        uint256 decayedReward,
        uint256 comboBonus,
        uint256 dynamicMultiplier,
        uint256 finalReward
    ) {
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
     * @dev Get combo information for a user and level
     */
    function getComboInfo(address user, uint8 level) external view returns (
        uint256 count,
        uint256 currentBonus,
        uint256[] memory thresholds,
        uint256[] memory bonuses
    ) {
        count = stakingContract.userLevelCounts(user, level);
        currentBonus = _calculateComboBonus(user, level);
        
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
     * @dev Get level statistics
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
     * @dev Get platform statistics
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
     * @dev Calculate decayed reward based on staking duration
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
     * @dev Calculate combo bonus for user's level
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
     * @dev Calculate dynamic balance multiplier
     */
    function _calculateDynamicMultiplier(uint8 level) internal view returns (uint256) {
        uint256 staked = stakingContract.totalStakedPerLevel(level);
        
        // Get supply from CPNFT contract (convert level 1-6 to enum 1-6)
        uint256 supply;
        if (level == 1) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.C);
        else if (level == 2) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.B);
        else if (level == 3) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.A);
        else if (level == 4) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.S);
        else if (level == 5) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.SS);
        else if (level == 6) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.SSS);
        
        if (supply == 0) return 10000; // 1.0x
        
        uint256 stakingRatio = (staked * 10000) / supply;
        
        (
            uint256 highStakeThreshold,
            uint256 lowStakeThreshold,
            uint256 highStakeMultiplier,
            uint256 lowStakeMultiplier,
            uint256 quarterlyAdjustmentMax
        ) = configContract.getDynamicConfig();
        
        if (stakingRatio > highStakeThreshold) {
            return highStakeMultiplier;
        } else if (stakingRatio < lowStakeThreshold) {
            return lowStakeMultiplier;
        }
        
        return 10000; // 1.0x
    }
    
    /**
     * @dev Calculate continuous staking bonus based on staking duration
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
}