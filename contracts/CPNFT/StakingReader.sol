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
            
            uint256 duration = block.timestamp - stakeTime;
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
                stakingDuration: block.timestamp - st,
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
        
        uint256 currentTime = block.timestamp;
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
