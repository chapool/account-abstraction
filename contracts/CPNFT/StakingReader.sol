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
     * @dev Get detailed combo status for a user and level
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
    /**
     * @dev Get level supply from CPNFT contract
     * @param level NFT level (1-6: C, B, A, S, SS, SSS)
     * @return The total supply for the specified level
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
    function getComprehensiveRewardCalculation(uint256 tokenId) external view returns (
        uint256 totalDays,
        uint256 totalRewards,
        uint256[] memory baseRewards,
        uint256[] memory decayedRewards,
        uint256[] memory quarterlyMultipliers,
        uint256[] memory dynamicMultipliers,
        uint256[] memory dailyRewards
    ) {
        (
            address owner,
            uint256 tokenId_,
            uint8 level,
            uint256 stakeTime,
            uint256 lastClaimTime,
            bool isActive,
            uint256 totalRewards_,
            uint256 pendingRewards
        ) = stakingContract.stakes(tokenId);
        
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