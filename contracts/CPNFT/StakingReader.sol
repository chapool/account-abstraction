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
 * @dev 轻量级只读合约，专门用于前端查询功能（优化版本）
 * @notice 本合约提供优化的查询接口，减少Gas消耗，提高查询效率
 */
contract StakingReader is 
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable
{
    // ============================================
    // 数据结构定义
    // ============================================
    
    struct UserStakingSummary {
        uint256 totalStakedCount;
        uint256 totalClaimedRewards;
        uint256 totalPendingRewards;
        uint256[6] levelStakingCounts;
        uint256 longestStakingDuration;
        uint256 totalEffectiveMultiplier;
    }

    struct StakedNFTInfo {
        uint256 tokenId;
        uint8 level;
        uint256 stakingDuration;
        uint256 pendingRewards;
        uint256 totalRewards;
        uint256 effectiveMultiplier;
    }

    struct UserRewardStats {
        uint256 totalHistoricalRewards;
        uint256 totalPendingRewards;
        uint256[6] rewardsPerLevel;
        uint256 last24HoursRewards;
        uint256 averageDailyRewards;
    }

    struct UserComboSummary {
        uint256[6] currentComboCounts;
        uint256[6] comboBonus;
        uint256[6] nextComboThreshold;
        bool[6] hasPendingCombo;
    }

    /// @dev 用户每日收益详情
    struct UserDailyRewards {
        uint256 totalBaseReward;           // 总基础收益
        uint256 totalDecayedReward;        // 总衰减后收益
        uint256 totalComboBonus;           // 总Combo加成（百分比，基于衰减后）
        uint256 totalDynamicMultiplier;    // 总动态乘数（百分比，基于基础收益）
        uint256 totalFinalReward;          // 总最终收益
        uint256 totalBonus;                // 总乘数（10000=100%, 11000=110%, 9000=90%）
        uint256[6] baseRewardPerLevel;     // 各等级基础收益
        uint256[6] finalRewardPerLevel;    // 各等级最终收益
        uint256[6] nftCountPerLevel;       // 各等级NFT数量
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
    
    function initialize(
        address _stakingContract,
        address _configContract,
        address _cpnftContract,
        address _owner
    ) public initializer {
        __UUPSUpgradeable_init();
        __Ownable_init(_owner);
        
        require(_stakingContract != address(0), "Invalid staking address");
        require(_configContract != address(0), "Invalid config address");
        require(_cpnftContract != address(0), "Invalid CPNFT address");
        
        stakingContract = Staking(_stakingContract);
        configContract = StakingConfig(_configContract);
        cpnftContract = CPNFT(_cpnftContract);
    }
    
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    function updateContracts(
        address _stakingContract,
        address _configContract,
        address _cpnftContract
    ) external onlyOwner {
        if (_stakingContract != address(0)) stakingContract = Staking(_stakingContract);
        if (_configContract != address(0)) configContract = StakingConfig(_configContract);
        if (_cpnftContract != address(0)) cpnftContract = CPNFT(_cpnftContract);
    }
    
    // ============================================
    // 核心查询函数（保留最常用的）
    // ============================================
    
    /// @notice 获取NFT待领取收益（直接调用Staking合约）
    function getPendingRewards(uint256 tokenId) external view returns (uint256) {
        return stakingContract.calculatePendingRewards(tokenId);
    }
    
    /// @notice 获取NFT质押详情
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
        
        uint256 realTimePendingRewards = pendingRewards;
        if (isActive && stakingContract.getCurrentTimestamp() > lastClaimTime) {
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
    
    /// @notice 获取用户Combo信息
    function getComboInfo(address user, uint8 level) external view returns (
        uint256 count,
        uint256 currentBonus,
        uint256[3] memory thresholds,
        uint256[3] memory bonuses
    ) {
        count = stakingContract.userLevelCounts(user, level);
        currentBonus = stakingContract.getEffectiveComboBonus(user, level);
        thresholds = configContract.getComboThresholds();
        bonuses = configContract.getComboBonuses();
    }
    
    /// @notice 获取等级统计
    function getLevelStats() external view returns (
        uint256[7] memory stakedCounts,
        uint256[7] memory supplies,
        uint256[7] memory stakingRatios,
        uint256[7] memory dynamicMultipliers
    ) {
        for (uint256 i = 0; i < 7; i++) {
            stakedCounts[i] = stakingContract.totalStakedPerLevel(uint8(i));
            supplies[i] = _getLevelSupply(uint8(i));
            if (supplies[i] > 0) {
                stakingRatios[i] = (stakedCounts[i] * 10000) / supplies[i];
            }
            if (i > 0) {
                dynamicMultipliers[i] = _calculateDynamicMultiplier(uint8(i));
            }
        }
    }
    
    /// @notice 获取平台统计（简化版）
    function getPlatformStats() external view returns (
        uint256[7] memory staked, 
        uint256[7] memory supply
    ) {
        for (uint256 i = 0; i < 7; i++) {
            staked[i] = stakingContract.totalStakedPerLevel(uint8(i));
            supply[i] = _getLevelSupply(uint8(i));
        }
    }
    
    // ============================================
    // 用户统计查询（新增的核心功能）
    // ============================================
    
    /// @notice 获取用户质押汇总
    function getUserStakingSummary(address user) external view returns (UserStakingSummary memory summary) {
        uint256[] memory tokenIds = _getUserTokenIds(user);
        summary.totalStakedCount = tokenIds.length;
        
        if (summary.totalStakedCount == 0) return summary;

        uint256 longestDuration = 0;
        uint256 totalMultiplier = 0;
        
        for (uint256 i = 0; i < tokenIds.length; i++) {
            (address owner_, uint256 tokenId_, uint8 level, uint256 stakeTime, uint256 lastClaimTime, bool isActive, uint256 totalRewards,) = stakingContract.stakes(tokenIds[i]);
            
            if (!isActive) continue;
            
            if (level >= 1 && level <= 6) {
                summary.levelStakingCounts[level - 1]++;
            }
            
            summary.totalClaimedRewards += totalRewards;
            summary.totalPendingRewards += stakingContract.calculatePendingRewards(tokenIds[i]);
            
            uint256 duration = stakingContract.getCurrentTimestamp() - stakeTime;
            if (duration > longestDuration) longestDuration = duration;
            
            uint256 comboBonus = stakingContract.getEffectiveComboBonus(user, level);
            uint256 dynamicMultiplier = _calculateDynamicMultiplier(level);
            uint256 quarterlyMultiplier = configContract.getQuarterlyMultiplier();
            
            uint256 effectiveMultiplier = (10000 + comboBonus) * dynamicMultiplier * quarterlyMultiplier / (10000 * 10000);
            totalMultiplier += effectiveMultiplier;
        }
        
        summary.longestStakingDuration = longestDuration;
        summary.totalEffectiveMultiplier = summary.totalStakedCount > 0 ? 
            totalMultiplier / summary.totalStakedCount : 0;
    }

    /// @notice 获取用户质押的NFT列表（分页）
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
        
        uint256 end = offset + limit > total ? total : offset + limit;
        nfts = new StakedNFTInfo[](end - offset);
        uint256 idx = 0;
        
        for (uint256 i = offset; i < end; i++) {
            (,, uint8 lv, uint256 st,, bool active, uint256 tr,) = stakingContract.stakes(tokenIds[i]);
            
            if (!active) continue;
            
            uint256 mult = _getEffectiveMultiplier(user, lv);
            
            nfts[idx++] = StakedNFTInfo({
                tokenId: tokenIds[i],
                level: lv,
                stakingDuration: stakingContract.getCurrentTimestamp() - st,
                pendingRewards: stakingContract.calculatePendingRewards(tokenIds[i]),
                totalRewards: tr,
                effectiveMultiplier: mult
            });
        }
        
        if (idx < nfts.length) {
            assembly {
                mstore(nfts, idx)
            }
        }
    }
    
    function _getEffectiveMultiplier(address user, uint8 level) internal view returns (uint256) {
        uint256 cb = stakingContract.getEffectiveComboBonus(user, level);
        uint256 dm = _calculateDynamicMultiplier(level);
        uint256 qm = configContract.getQuarterlyMultiplier();
        return (10000 + cb) * dm * qm / (10000 * 10000);
    }

    /// @notice 获取用户收益统计
    function getUserRewardStats(address user) external view returns (UserRewardStats memory stats) {
        uint256[] memory tokenIds = _getUserTokenIds(user);
        if (tokenIds.length == 0) return stats;
        
        uint256 currentTime = stakingContract.getCurrentTimestamp();
        uint256 oneDayAgo = currentTime - 1 days;
        uint256 totalStakingDays = 0;
        
        for (uint256 i = 0; i < tokenIds.length; i++) {
            (address owner_, uint256 tokenId_, uint8 level, uint256 stakeTime, uint256 lastClaimTime, bool isActive, uint256 totalRewards,) = stakingContract.stakes(tokenIds[i]);
            
            if (!isActive) continue;
            
            stats.totalHistoricalRewards += totalRewards;
            uint256 pendingRewards = stakingContract.calculatePendingRewards(tokenIds[i]);
            stats.totalPendingRewards += pendingRewards;
            
            if (level >= 1 && level <= 6) {
                stats.rewardsPerLevel[level - 1] += totalRewards + pendingRewards;
            }
            
            if (lastClaimTime < oneDayAgo) {
                uint256 dailyReward = configContract.getDailyReward(level);
                uint256 comboBonus = stakingContract.getEffectiveComboBonus(user, level);
                uint256 dynamicMultiplier = _calculateDynamicMultiplier(level);
                uint256 quarterlyMultiplier = configContract.getQuarterlyMultiplier();
                
                uint256 effectiveDaily = dailyReward * (10000 + comboBonus) * dynamicMultiplier * quarterlyMultiplier / (10000 * 10000 * 10000);
                stats.last24HoursRewards += effectiveDaily;
            }
            
            totalStakingDays += (currentTime - stakeTime) / 1 days;
        }
        
        if (totalStakingDays > 0) {
            stats.averageDailyRewards = (stats.totalHistoricalRewards + stats.totalPendingRewards) / totalStakingDays;
        }
    }

    /// @notice 获取用户Combo汇总
    function getUserComboSummary(address user) external view returns (UserComboSummary memory summary) {
        uint256[3] memory thresholds = configContract.getComboThresholds();
        
        for (uint8 level = 1; level <= 6; level++) {
            summary.currentComboCounts[level - 1] = stakingContract.userLevelCounts(user, level);
            summary.comboBonus[level - 1] = stakingContract.getEffectiveComboBonus(user, level);
            
            Staking.ComboStatus memory status = stakingContract.getComboStatus(user, level);
            summary.hasPendingCombo[level - 1] = status.isPending;
            
            uint256 currentCount = summary.currentComboCounts[level - 1];
            summary.nextComboThreshold[level - 1] = type(uint256).max;
            
            for (uint256 i = 0; i < thresholds.length; i++) {
                if (currentCount < thresholds[i]) {
                    summary.nextComboThreshold[level - 1] = thresholds[i] - currentCount;
                    break;
                }
            }
        }
    }
    
    // ============================================
    // 配置查询（简化版）
    // ============================================
    
    /// @notice 获取基础配置
    function getBasicConfig() external view returns (
        uint256 minStakeDays,
        uint256 earlyWithdrawPenalty,
        uint256 quarterlyMultiplier
    ) {
        (minStakeDays, earlyWithdrawPenalty, quarterlyMultiplier, ) = configContract.getBasicConfig();
    }
    
    /// @notice 获取每日奖励配置
    function getDailyReward(uint8 level) external view returns (uint256) {
        return configContract.getDailyReward(level);
    }
    
    /// @notice 获取Combo配置
    function getComboConfig() external view returns (
        uint256[3] memory thresholds,
        uint256[3] memory bonuses
    ) {
        thresholds = configContract.getComboThresholds();
        bonuses = configContract.getComboBonuses();
    }
    
    /// @notice 获取合约版本
    function getVersions() external view returns (
        string memory stakingVersion,
        string memory configVersion
    ) {
        stakingVersion = stakingContract.version();
        configVersion = configContract.version();
    }
    
    // ============================================
    // 每日收益统计（新增）
    // ============================================
    
    /// @notice 获取用户所有NFT的每日收益详情
    /// @param user 用户地址
    /// @return 用户每日收益详细信息
    function getUserDailyRewards(address user) external view returns (UserDailyRewards memory) {
        UserDailyRewards memory rewards;
        uint256[] memory tokenIds = _getUserTokenIds(user);
        
        if (tokenIds.length == 0) return rewards;
        
        // 统计各等级的Combo加成
        uint256[6] memory comboBonuses;
        uint256 totalComboBonus = 0;
        uint256 totalRewardBeforeCombo = 0;
        
        for (uint8 level = 1; level <= 6; level++) {
            comboBonuses[level - 1] = stakingContract.getEffectiveComboBonus(user, level);
        }
        
        // 遍历所有质押的NFT
        for (uint256 i = 0; i < tokenIds.length; i++) {
            (,, uint8 level, uint256 stakeTime,, bool isActive,,) = stakingContract.stakes(tokenIds[i]);
            
            if (!isActive) continue;
            
            // 计算当前天数偏移
            uint256 dayOffset = (stakingContract.getCurrentTimestamp() - stakeTime) / 1 days;
            
            // 获取每日奖励分解
            try stakingContract.getDailyRewardBreakdown(tokenIds[i], dayOffset) returns (
                uint256 baseReward,
                uint256 decayedReward,
                uint256 quarterlyMultiplier,
                uint256 dynamicMultiplier,
                uint256 finalReward
            ) {
                // 累加基础收益
                rewards.totalBaseReward += baseReward;
                rewards.baseRewardPerLevel[level - 1] += baseReward;
                
                // 累加衰减后收益
                rewards.totalDecayedReward += decayedReward;
                
                // 应用Combo加成（finalReward已包含quarterly和dynamic，需要再应用combo）
                uint256 comboBonus = comboBonuses[level - 1];
                uint256 rewardWithCombo = finalReward * (10000 + comboBonus) / 10000;
                
                // 累加最终收益
                rewards.totalFinalReward += rewardWithCombo;
                rewards.finalRewardPerLevel[level - 1] += rewardWithCombo;
                
                // 计算 combo 贡献（用于加权平均）
                if (comboBonus > 0) {
                    totalRewardBeforeCombo += finalReward;
                    totalComboBonus += finalReward * comboBonus;
                }
                
                // 统计NFT数量
                rewards.nftCountPerLevel[level - 1]++;
                
            } catch {
                // 如果某个NFT计算失败，跳过
                continue;
            }
        }
        
        // 计算加权平均Combo加成
        if (totalRewardBeforeCombo > 0) {
            rewards.totalComboBonus = totalComboBonus / totalRewardBeforeCombo;
        }
        
        // 计算平均动态乘数（基于 decayed 到 final before combo 的增益）
        // finalReward = decayedReward * quarterlyMultiplier * dynamicMultiplier / (10000 * 10000)
        // 我们需要反推综合乘数
        if (rewards.totalDecayedReward > 0) {
            // totalFinalReward 包含 combo，需要移除 combo 影响
            uint256 finalBeforeCombo = (rewards.totalFinalReward * 10000) / (10000 + rewards.totalComboBonus);
            rewards.totalDynamicMultiplier = (finalBeforeCombo * 10000) / rewards.totalDecayedReward;
        }
        
        // 计算总乘数：最终收益 / 基础收益 * 10000
        // 10000 = 100% (无变化), 11000 = 110% (+10%), 9000 = 90% (-10%)
        if (rewards.totalBaseReward > 0) {
            rewards.totalBonus = (rewards.totalFinalReward * 10000) / rewards.totalBaseReward;
        } else {
            rewards.totalBonus = 10000; // 默认 100%
        }
        
        return rewards;
    }
    
    /// @notice 获取用户指定等级的每日收益
    /// @param user 用户地址
    /// @param level NFT等级 (1-6)
    /// @return baseReward 基础收益
    /// @return finalReward 最终收益（含加成）
    /// @return nftCount NFT数量
    function getUserDailyRewardsByLevel(address user, uint8 level) external view returns (
        uint256 baseReward,
        uint256 finalReward,
        uint256 nftCount
    ) {
        require(level > 0 && level <= 6, "Invalid level");
        
        uint256[] memory tokenIds = _getUserTokenIds(user);
        uint256 comboBonus = stakingContract.getEffectiveComboBonus(user, level);
        
        for (uint256 i = 0; i < tokenIds.length; i++) {
            (,, uint8 nftLevel, uint256 stakeTime,, bool isActive,,) = stakingContract.stakes(tokenIds[i]);
            
            if (!isActive || nftLevel != level) continue;
            
            uint256 dayOffset = (stakingContract.getCurrentTimestamp() - stakeTime) / 1 days;
            
            try stakingContract.getDailyRewardBreakdown(tokenIds[i], dayOffset) returns (
                uint256 base,
                uint256,
                uint256,
                uint256,
                uint256 finalRwd
            ) {
                baseReward += base;
                finalReward += finalRwd * (10000 + comboBonus) / 10000;
                nftCount++;
            } catch {
                continue;
            }
        }
    }
    
    /// @notice 获取用户单个NFT的每日收益详情
    /// @param tokenId NFT token ID
    /// @return baseReward 基础收益
    /// @return decayedReward 衰减后收益
    /// @return comboBonus Combo加成（百分比）
    /// @return dynamicMultiplier 动态乘数（百分比）
    /// @return finalReward 最终收益
    function getNFTDailyRewardBreakdown(uint256 tokenId) external view returns (
        uint256 baseReward,
        uint256 decayedReward,
        uint256 comboBonus,
        uint256 dynamicMultiplier,
        uint256 finalReward
    ) {
        (address owner,, uint8 level, uint256 stakeTime,, bool isActive,,) = stakingContract.stakes(tokenId);
        require(isActive, "NFT is not staked");
        
        uint256 dayOffset = (stakingContract.getCurrentTimestamp() - stakeTime) / 1 days;
        
        // 获取基础分解
        uint256 quarterlyMult;
        (baseReward, decayedReward, quarterlyMult, dynamicMultiplier, finalReward) = 
            stakingContract.getDailyRewardBreakdown(tokenId, dayOffset);
        
        // 获取Combo加成
        comboBonus = stakingContract.getEffectiveComboBonus(owner, level);
        
        // 应用Combo到最终收益
        finalReward = finalReward * (10000 + comboBonus) / 10000;
    }

    // ============================================
    // 内部辅助函数
    // ============================================
    
    function _getUserTokenIds(address user) internal view returns (uint256[] memory) {
        uint256 count = 0;
        for (uint256 i = 0; ; i++) {
            try stakingContract.userStakes(user, i) returns (uint256) {
                count++;
            } catch {
                break;
            }
        }
        
        uint256[] memory tokenIds = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            tokenIds[i] = stakingContract.userStakes(user, i);
        }
        
        return tokenIds;
    }
    
    function _getLevelSupply(uint8 level) internal view returns (uint256) {
        if (level == 0) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.NORMAL);
        if (level == 1) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.C);
        if (level == 2) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.B);
        if (level == 3) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.A);
        if (level == 4) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.S);
        if (level == 5) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.SS);
        if (level == 6) return cpnftContract.getLevelSupply(CPNFT.NFTLevel.SSS);
        return 0;
    }
    
    function _calculateDynamicMultiplier(uint8 level) internal view returns (uint256) {
        uint256 staked = stakingContract.totalStakedPerLevel(level);
        uint256 supply = _getLevelSupply(level);
        
        if (supply == 0) return 10000;
        
        uint256 stakingRatio = (staked * 10000) / supply;
        
        (
            uint256 highStakeThreshold,
            uint256 lowStakeThreshold,
            uint256 highStakeMultiplier,
            uint256 lowStakeMultiplier,
        ) = configContract.getDynamicConfig();
        
        if (stakingRatio > highStakeThreshold) {
            return highStakeMultiplier;
        } else if (stakingRatio < lowStakeThreshold) {
            return lowStakeMultiplier;
        }
        
        return 10000;
    }
}
