// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "./CPNFT.sol";
import "../cpop/interfaces/ICPOPToken.sol";
import "../cpop/interfaces/IAccountManager.sol";
import "./StakingConfig.sol";

/**
 * @title Staking - Complete NFT Staking Contract with Integrated Reward Logic
 * @dev Main staking contract with integrated reward calculation and platform statistics
 */
contract Staking is 
    Initializable, 
    UUPSUpgradeable, 
    OwnableUpgradeable, 
    ReentrancyGuardUpgradeable, 
    PausableUpgradeable 
{
    
    // ============================================
    // STORAGE VARIABLES
    // ============================================
    
    // Contract dependencies
    CPNFT public cpnftContract;
    ICPOPToken public cpopTokenContract;
    IAccountManager public accountManagerContract;
    StakingConfig public configContract;
    
    // Staking data
    struct StakeInfo {
        address owner;
        uint256 tokenId;
        uint8 level;
        uint256 stakeTime;
        uint256 lastClaimTime;
        bool isActive;
        uint256 totalRewards;
        uint256 pendingRewards;
        bool continuousBonusClaimed;  // 标记连续质押奖励是否已发放
    }
    
    mapping(uint256 => StakeInfo) public stakes;
    mapping(address => uint256[]) public userStakes;
    mapping(address => mapping(uint8 => uint256)) public userLevelCounts;
    
    uint256 public totalStakedCount;
    
    // Platform statistics for dynamic balancing
    mapping(uint8 => uint256) public totalStakedPerLevel;
    
    // Historical adjustment records for reward recalculation
    struct HistoricalAdjustment {
        uint256 timestamp;
        uint256 quarterlyMultiplier;
        mapping(uint8 => uint256) dynamicMultipliers; // Level-specific multipliers at this time
    }
    
    HistoricalAdjustment[] public historicalAdjustments;
    
    // Combo status for next-day effect mechanism
    struct ComboStatus {
        uint8 level;
        uint256 count;
        uint256 effectiveFrom; // Timestamp when this status becomes effective
        uint256 bonus;         // The bonus that will be applied
        bool isPending;        // Whether there's a pending status change
    }
    
    mapping(address => mapping(uint8 => ComboStatus)) public userComboStatus;
    
    // Testing mode for time manipulation (NEW - added in upgrade)
    bool public testMode;
    uint256 public testTimestamp;
    
    // ============================================
    // EVENTS
    // ============================================
    
    event NFTStaked(address indexed user, uint256 indexed tokenId, uint8 level, uint256 timestamp);
    event NFTUnstaked(address indexed user, uint256 indexed tokenId, uint256 rewards, uint256 timestamp);
    event RewardsClaimed(address indexed user, uint256 indexed tokenId, uint256 amount, uint256 timestamp);
    event BatchRewardsClaimed(address indexed user, uint256 tokenCount, uint256 totalAmount, uint256 timestamp);
    event BatchStaked(address indexed user, uint256[] tokenIds, uint256 timestamp);
    event BatchUnstaked(address indexed user, uint256[] tokenIds, uint256 totalRewards, uint256 timestamp);
    event PlatformStatsUpdated(uint8 level, uint256 staked, uint256 supply);
    
    // ============================================
    // INITIALIZER
    // ============================================
    
    function initialize(
        address _cpnftContract,
        address _cpopTokenContract,
        address _accountManagerContract,
        address _configContract,
        address _owner
    ) public initializer {
        __UUPSUpgradeable_init();
        __Ownable_init(_owner);
        __ReentrancyGuard_init();
        __Pausable_init();
        
        require(_cpnftContract != address(0), "CPNFT contract address cannot be zero");
        require(_cpopTokenContract != address(0), "CPOP token contract address cannot be zero");
        require(_accountManagerContract != address(0), "Account manager contract address cannot be zero");
        require(_configContract != address(0), "Config contract address cannot be zero");
        
        cpnftContract = CPNFT(_cpnftContract);
        cpopTokenContract = ICPOPToken(_cpopTokenContract);
        accountManagerContract = IAccountManager(_accountManagerContract);
        configContract = StakingConfig(_configContract);
        
        totalStakedCount = 0;
    }
    
    // ============================================
    // UUPS UPGRADE AUTHORIZATION
    // ============================================
    
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
    
    // ============================================
    // TIME MANAGEMENT FOR TESTING
    // ============================================
    
    /**
     * @dev Get current timestamp (real or test)
     */
    function _getCurrentTimestamp() internal view returns (uint256) {
        return testMode ? testTimestamp : block.timestamp;
    }
    
    /**
     * @dev Get current timestamp (public version for external contracts)
     */
    function getCurrentTimestamp() external view returns (uint256) {
        return _getCurrentTimestamp();
    }
    
    /**
     * @dev Enable test mode and set initial test timestamp
     */
    function enableTestMode(uint256 initialTimestamp) external onlyOwner {
        testMode = true;
        testTimestamp = initialTimestamp > 0 ? initialTimestamp : block.timestamp;
    }
    
    /**
     * @dev Disable test mode (back to real time)
     */
    function disableTestMode() external onlyOwner {
        testMode = false;
        testTimestamp = 0;
    }
    
    /**
     * @dev Set test timestamp (only in test mode)
     */
    function setTestTimestamp(uint256 timestamp) external onlyOwner {
        require(testMode, "Not in test mode");
        require(timestamp >= testTimestamp, "Cannot go back in time");
        testTimestamp = timestamp;
    }
    
    /**
     * @dev Fast forward time (only in test mode)
     */
    function fastForwardTime(uint256 seconds_) external onlyOwner {
        require(testMode, "Not in test mode");
        testTimestamp += seconds_;
    }
    
    /**
     * @dev Fast forward by days (only in test mode)
     */
    function fastForwardDays(uint256 days_) external onlyOwner {
        require(testMode, "Not in test mode");
        testTimestamp += days_ * 1 days;
    }
    
    /**
     * @dev Fast forward by minutes (only in test mode)
     */
    function fastForwardMinutes(uint256 minutes_) external onlyOwner {
        require(testMode, "Not in test mode");
        testTimestamp += minutes_ * 1 minutes;
    }
    
    // ============================================
    // STAKING FUNCTIONS
    // ============================================
    
    /**
     * @dev Stake a single NFT
     */
    function stake(uint256 tokenId) external nonReentrant whenNotPaused {
        _stake(tokenId);
    }
    
    /**
     * @dev Stake multiple NFTs in batch
     */
    function batchStake(uint256[] calldata tokenIds) external nonReentrant whenNotPaused {
        require(tokenIds.length > 0, "Token IDs array cannot be empty");
        require(tokenIds.length <= 50, "Too many tokens in batch");
        
        for (uint256 i = 0; i < tokenIds.length; i++) {
            _stake(tokenIds[i]);
        }
        
        emit BatchStaked(msg.sender, tokenIds, _getCurrentTimestamp());
    }
    
    /**
     * @dev Internal stake function
     */
    function _stake(uint256 tokenId) internal {
        require(cpnftContract.ownerOf(tokenId) == msg.sender, "Not the owner of this NFT");
        require(!stakes[tokenId].isActive, "NFT is already staked");
        
        uint8 level = uint8(cpnftContract.getTokenLevel(tokenId));
        require(level != 0, "NORMAL level NFTs cannot be staked");
        
        // Create stake info
        stakes[tokenId] = StakeInfo({
            owner: msg.sender,
            tokenId: tokenId,
            level: level,
            stakeTime: _getCurrentTimestamp(),
            lastClaimTime: _getCurrentTimestamp(),
            isActive: true,
            totalRewards: 0,
            pendingRewards: 0,
            continuousBonusClaimed: false
        });
        
        // Update user data
        userStakes[msg.sender].push(tokenId);
        userLevelCounts[msg.sender][level]++;
        
        // Update combo status for next-day effect
        _updateComboStatus(msg.sender, level);
        
        // Update platform stats
        totalStakedPerLevel[level]++;
        totalStakedCount++;
        
        // Set NFT as staked in CPNFT contract
        cpnftContract.setStakeStatus(tokenId, true);
        
        emit NFTStaked(msg.sender, tokenId, level, _getCurrentTimestamp());
        
        // Get supply from CPNFT contract
        uint256 supply = _getLevelSupply(level);
        emit PlatformStatsUpdated(level, totalStakedPerLevel[level], supply);
    }
    
    /**
     * @dev Unstake a single NFT
     */
    function unstake(uint256 tokenId) external nonReentrant whenNotPaused {
        require(stakes[tokenId].owner == msg.sender, "Not the owner of this stake");
        require(stakes[tokenId].isActive, "NFT is not staked");
        
        StakeInfo storage stakeInfo = stakes[tokenId];
        
        // Calculate rewards
        uint256 rewards = _calculateTotalRewards(tokenId);
        
        // Check for early withdrawal penalty
        uint256 minStakeDays = configContract.getMinStakeDays();
        uint256 stakingDays = (_getCurrentTimestamp() - stakeInfo.stakeTime) / 1 days;
        
        if (stakingDays < minStakeDays) {
            uint256 penalty = configContract.getEarlyWithdrawPenalty();
            rewards = rewards * (10000 - penalty) / 10000;
        }
        
        // Add continuous staking bonus (based on rewards at threshold) - only if not already claimed
        if (!stakeInfo.continuousBonusClaimed) {
            uint256 continuousBonus = _calculateContinuousBonus(tokenId, stakingDays);
            rewards += continuousBonus;
            stakeInfo.continuousBonusClaimed = true;
        }
        
        // Send rewards to AA account
        if (rewards > 0) {
            address aaAccount = _getAAAccount(msg.sender);
            cpopTokenContract.mint(aaAccount, rewards);
        }
        
        // Update stake info
        stakeInfo.isActive = false;
        stakeInfo.totalRewards = rewards;
        stakeInfo.pendingRewards = 0;
        
        // Update user data
        _removeUserStake(msg.sender, tokenId);
        userLevelCounts[msg.sender][stakeInfo.level]--;
        
        // Update combo status for next-day effect
        _updateComboStatus(msg.sender, stakeInfo.level);
        
        // Update platform stats
        totalStakedPerLevel[stakeInfo.level]--;
        totalStakedCount--;
        
        // Set NFT as not staked in CPNFT contract
        cpnftContract.setStakeStatus(tokenId, false);
        
        emit NFTUnstaked(msg.sender, tokenId, rewards, _getCurrentTimestamp());
        
        // Get supply from CPNFT contract
        uint256 supply = _getLevelSupply(stakeInfo.level);
        emit PlatformStatsUpdated(stakeInfo.level, totalStakedPerLevel[stakeInfo.level], supply);
    }
    
    /**
     * @dev Unstake multiple NFTs in batch
     */
    function batchUnstake(uint256[] calldata tokenIds) external nonReentrant whenNotPaused {
        require(tokenIds.length > 0, "Token IDs array cannot be empty");
        require(tokenIds.length <= 50, "Too many tokens in batch");
        
        uint256 totalRewards = 0;
        
        for (uint256 i = 0; i < tokenIds.length; i++) {
            uint256 tokenId = tokenIds[i];
            require(stakes[tokenId].owner == msg.sender, "Not the owner of this stake");
            require(stakes[tokenId].isActive, "NFT is not staked");
            
            StakeInfo storage stakeInfo = stakes[tokenId];
            
            // Calculate rewards
            uint256 rewards = _calculateTotalRewards(tokenId);
            
            // Check for early withdrawal penalty
            uint256 minStakeDays = configContract.getMinStakeDays();
            uint256 stakingDays = (_getCurrentTimestamp() - stakeInfo.stakeTime) / 1 days;
            
            if (stakingDays < minStakeDays) {
                uint256 penalty = configContract.getEarlyWithdrawPenalty();
                rewards = rewards * (10000 - penalty) / 10000;
            }
            
            // Add continuous staking bonus (based on rewards at threshold) - only if not already claimed
            if (!stakeInfo.continuousBonusClaimed) {
                uint256 continuousBonus = _calculateContinuousBonus(tokenId, stakingDays);
                rewards += continuousBonus;
                stakeInfo.continuousBonusClaimed = true;
            }
            
            totalRewards += rewards;
            
            // Update stake info
            stakeInfo.isActive = false;
            stakeInfo.totalRewards = rewards;
            stakeInfo.pendingRewards = 0;
            
            // Update user data
            _removeUserStake(msg.sender, tokenId);
            userLevelCounts[msg.sender][stakeInfo.level]--;
            
            // Update combo status for next-day effect
            _updateComboStatus(msg.sender, stakeInfo.level);
            
            // Update platform stats
            totalStakedPerLevel[stakeInfo.level]--;
            totalStakedCount--;
            
            // Set NFT as not staked in CPNFT contract
            cpnftContract.setStakeStatus(tokenId, false);
            
            // Emit individual event for each token
            emit NFTUnstaked(msg.sender, tokenId, rewards, _getCurrentTimestamp());
        }
        
        // Send total rewards to AA account
        if (totalRewards > 0) {
            address aaAccount = _getAAAccount(msg.sender);
            cpopTokenContract.mint(aaAccount, totalRewards);
        }
        
        emit BatchUnstaked(msg.sender, tokenIds, totalRewards, _getCurrentTimestamp());
    }
    
    // ============================================
    // REWARD CALCULATION FUNCTIONS
    // ============================================
    
    /**
     * @dev Calculate pending rewards since last claim (internal)
     * @param tokenId The NFT token ID
     * @param applyLimit Whether to apply 90-day limit (for Gas safety)
     */
    function _calculatePendingRewards(uint256 tokenId, bool applyLimit) internal view returns (uint256) {
        StakeInfo memory stakeInfo = stakes[tokenId];
        
        // Calculate base rewards with phase-based decay and dynamic adjustment
        uint256 totalDays = (_getCurrentTimestamp() - stakeInfo.lastClaimTime) / 1 days;
        if (totalDays == 0) return 0;
        
        // Gas optimization: Limit maximum calculation days to prevent gas overflow
        // Only apply when claiming (to ensure Gas safety), not when displaying
        if (applyLimit) {
            uint256 MAX_CALCULATION_DAYS = 90;
            if (totalDays > MAX_CALCULATION_DAYS) {
                totalDays = MAX_CALCULATION_DAYS;
            }
        }
        
        uint256 baseReward = configContract.getDailyReward(stakeInfo.level);
        uint256 decayInterval = configContract.getDecayInterval(stakeInfo.level);
        uint256 decayRate = configContract.getDecayRate(stakeInfo.level);
        
        uint256 totalRewards = 0;
        
        // Calculate rewards day by day with phase-based decay and dynamic adjustment
        for (uint256 day = 0; day < totalDays; day++) {
            uint256 currentDayFromStake = (_getCurrentTimestamp() - stakeInfo.stakeTime) / 1 days - (totalDays - 1 - day);
            uint256 currentDayTimestamp = stakeInfo.stakeTime + (currentDayFromStake * 1 days);
            
            uint256 dailyReward = baseReward;
            
            // Apply decay based on current day from stake
            if (decayInterval > 0 && currentDayFromStake > decayInterval) {
                uint256 completedCycles = (currentDayFromStake - 1) / decayInterval;
                
                // Apply compound decay for each completed cycle
                for (uint256 i = 0; i < completedCycles; i++) {
                    uint256 totalDecaySoFar = (i + 1) * decayRate;
                    if (totalDecaySoFar > configContract.getMaxDecayRate(stakeInfo.level)) {
                        uint256 remainingDecay = configContract.getMaxDecayRate(stakeInfo.level) - (i * decayRate);
                        dailyReward = dailyReward * (10000 - remainingDecay) / 10000;
                        break;
                    }
                    
                    dailyReward = dailyReward * (10000 - decayRate) / 10000;
                }
            }
            
            // Apply historical quarterly adjustment for this specific day
            uint256 historicalQuarterlyMultiplier = _getHistoricalQuarterlyMultiplier(currentDayTimestamp);
            dailyReward = dailyReward * historicalQuarterlyMultiplier / 10000;
            
            // Apply historical dynamic multiplier for this specific day
            uint256 dynamicMultiplier = _getHistoricalDynamicMultiplier(
                stakeInfo.level, 
                currentDayTimestamp
            );
            dailyReward = dailyReward * dynamicMultiplier / 10000;
            
            totalRewards += dailyReward;
        }
        
        // Calculate combo bonus based on current NFT's individual decay state
        uint256 comboBonus = _calculateComboBonus(stakeInfo.owner, stakeInfo.level);
        totalRewards = totalRewards * (10000 + comboBonus) / 10000;
        
        // Add continuous staking bonus only if not already claimed
        if (!stakeInfo.continuousBonusClaimed) {
            uint256 stakingDays = (_getCurrentTimestamp() - stakeInfo.stakeTime) / 1 days;
            uint256 continuousBonus = _calculateContinuousBonus(tokenId, stakingDays);
            totalRewards += continuousBonus;
        }
        
        return totalRewards;
    }

    /**
     * @dev Calculate pending rewards since last claim (external for Reader)
     * Note: This function has NO day limit for accurate display, but may timeout for very long periods
     */
    function calculatePendingRewards(uint256 tokenId) external view returns (uint256) {
        return _calculatePendingRewards(tokenId, false); // false = 不加限制，显示完整奖励
    }

    /**
     * @dev Claim pending rewards without unstaking
     */
    function claimRewards(uint256 tokenId) external nonReentrant whenNotPaused {
        require(stakes[tokenId].owner == msg.sender, "Not the owner of this stake");
        require(stakes[tokenId].isActive, "NFT is not staked");
        
        StakeInfo storage stakeInfo = stakes[tokenId];
        
        // Calculate rewards since last claim (with 90-day limit to prevent Gas overflow)
        uint256 rewards = _calculatePendingRewards(tokenId, true); // true = 限制90天，防止Gas超限
        require(rewards > 0, "No pending rewards");
        
        // Update last claim time
        stakeInfo.lastClaimTime = _getCurrentTimestamp();
        stakeInfo.pendingRewards = 0;
        stakeInfo.totalRewards += rewards;
        
        // Mark continuous bonus as claimed if it was included in this claim
        if (!stakeInfo.continuousBonusClaimed) {
            uint256 stakingDays = (_getCurrentTimestamp() - stakeInfo.stakeTime) / 1 days;
            uint256 continuousBonus = _calculateContinuousBonus(tokenId, stakingDays);
            if (continuousBonus > 0) {
                stakeInfo.continuousBonusClaimed = true;
            }
        }
        
        // Send rewards to AA account
        address aaAccount = _getAAAccount(msg.sender);
        cpopTokenContract.mint(aaAccount, rewards);
        
        emit RewardsClaimed(msg.sender, tokenId, rewards, _getCurrentTimestamp());
    }

    /**
     * @dev Batch claim rewards for multiple NFTs
     * @param tokenIds Array of token IDs to claim rewards for
     */
    function batchClaimRewards(uint256[] calldata tokenIds) external nonReentrant whenNotPaused {
        require(tokenIds.length > 0, "Empty token IDs array");
        require(tokenIds.length <= 50, "Too many tokens in batch");
        
        address aaAccount = _getAAAccount(msg.sender);
        uint256 totalRewards = 0;
        
        for (uint256 i = 0; i < tokenIds.length; i++) {
            uint256 tokenId = tokenIds[i];
            
            require(stakes[tokenId].owner == msg.sender, "Not the owner of this stake");
            require(stakes[tokenId].isActive, "NFT is not staked");
            
            StakeInfo storage stakeInfo = stakes[tokenId];
            
            // Calculate rewards since last claim (with 90-day limit to prevent Gas overflow)
            uint256 rewards = _calculatePendingRewards(tokenId, true); // true = 限制90天
            
            if (rewards > 0) {
                // Update last claim time
                stakeInfo.lastClaimTime = _getCurrentTimestamp();
                stakeInfo.pendingRewards = 0;
                stakeInfo.totalRewards += rewards;
                totalRewards += rewards;
                
                // Mark continuous bonus as claimed if it was included in this claim
                if (!stakeInfo.continuousBonusClaimed) {
                    uint256 stakingDays = (_getCurrentTimestamp() - stakeInfo.stakeTime) / 1 days;
                    uint256 continuousBonus = _calculateContinuousBonus(tokenId, stakingDays);
                    if (continuousBonus > 0) {
                        stakeInfo.continuousBonusClaimed = true;
                    }
                }
                
                emit RewardsClaimed(msg.sender, tokenId, rewards, _getCurrentTimestamp());
            }
        }
        
        require(totalRewards > 0, "No pending rewards");
        
        // Send total rewards to AA account
        cpopTokenContract.mint(aaAccount, totalRewards);
        
        emit BatchRewardsClaimed(msg.sender, tokenIds.length, totalRewards, _getCurrentTimestamp());
    }
    
    // ============================================
    // REWARD CALCULATION FUNCTIONS
    // ============================================
    
    /**
     * @dev Calculate total rewards for a staked NFT
     */
    function _calculateTotalRewards(uint256 tokenId) internal view returns (uint256) {
        StakeInfo memory stakeInfo = stakes[tokenId];
        
        // Calculate base rewards (with limit for Gas safety)
        uint256 baseRewards = _calculateRewards(
            stakeInfo.level,
            stakeInfo.stakeTime,
            stakeInfo.lastClaimTime,
            true // true = 限制90天，防止Gas超限
        );
        
        // Calculate combo bonus
        uint256 comboBonus = _calculateComboBonus(msg.sender, stakeInfo.level);
        
        // Calculate dynamic multiplier
        uint256 dynamicMultiplier = _calculateDynamicMultiplier(stakeInfo.level);
        
        // Apply bonuses
        uint256 finalRewards = baseRewards * (10000 + comboBonus) / 10000;
        finalRewards = finalRewards * dynamicMultiplier / 10000;
        
        return finalRewards;
    }
    
    /**
     * @dev Calculate rewards for a specific NFT (simplified version)
     */
    function _calculateRewards(
        uint8 level,
        uint256 stakeTime,
        uint256 lastClaimTime,
        bool applyLimit
    ) internal view returns (uint256) {
        uint256 totalDays = (_getCurrentTimestamp() - lastClaimTime) / 1 days;
        if (totalDays == 0) return 0;
        
        // Gas optimization: Limit maximum calculation days to prevent gas overflow
        if (applyLimit) {
            uint256 MAX_CALCULATION_DAYS = 90;
            if (totalDays > MAX_CALCULATION_DAYS) {
                totalDays = MAX_CALCULATION_DAYS;
            }
        }
        
        uint256 baseReward = configContract.getDailyReward(level);
        uint256 decayInterval = configContract.getDecayInterval(level);
        uint256 decayRate = configContract.getDecayRate(level);
        
        uint256 totalRewards = 0;
        
        // Calculate rewards day by day with proper phase-based decay
        for (uint256 day = 0; day < totalDays; day++) {
            uint256 currentDayFromStake = (_getCurrentTimestamp() - stakeTime) / 1 days - (totalDays - 1 - day);
            uint256 currentDayTimestamp = stakeTime + (currentDayFromStake * 1 days);
            
            uint256 dailyReward = baseReward;
            
            // Apply decay based on current day from stake
            if (decayInterval > 0 && currentDayFromStake > decayInterval) {
                uint256 completedCycles = (currentDayFromStake - 1) / decayInterval;
                
                // Apply compound decay for each completed cycle
                for (uint256 i = 0; i < completedCycles; i++) {
                    uint256 totalDecaySoFar = (i + 1) * decayRate;
                    if (totalDecaySoFar > configContract.getMaxDecayRate(level)) {
                        uint256 remainingDecay = configContract.getMaxDecayRate(level) - (i * decayRate);
                        dailyReward = dailyReward * (10000 - remainingDecay) / 10000;
                        break;
                    }
                    
                    dailyReward = dailyReward * (10000 - decayRate) / 10000;
                }
            }
            
            // Apply historical quarterly adjustment for this specific day
            uint256 historicalQuarterlyMultiplier = _getHistoricalQuarterlyMultiplier(currentDayTimestamp);
            dailyReward = dailyReward * historicalQuarterlyMultiplier / 10000;
            
            // Apply historical dynamic multiplier for this specific day
            uint256 dynamicMultiplier = _getHistoricalDynamicMultiplier(
                level, 
                currentDayTimestamp
            );
            dailyReward = dailyReward * dynamicMultiplier / 10000;
            
            totalRewards += dailyReward;
        }
        
        return totalRewards;
    }
    
    /**
     * @dev Calculate decayed reward based on staking duration
     * Uses phase-based decay as described in the document
     * Each phase has a fixed daily reward rate
     */
    function _calculateDecayedReward(
        uint256 baseReward,
        uint8 level,
        uint256 stakeTime
    ) internal view returns (uint256) {
        uint256 stakingDays = (_getCurrentTimestamp() - stakeTime) / 1 days;
        uint256 decayInterval = configContract.getDecayInterval(level);
        uint256 decayRate = configContract.getDecayRate(level);
        uint256 maxDecay = configContract.getMaxDecayRate(level);
        
        if (decayInterval == 0 || stakingDays < decayInterval) {
            return baseReward;
        }
        
        // Calculate which phase we're in
        uint256 completedCycles = stakingDays / decayInterval;
        
        // Apply compound decay for each completed cycle
        uint256 decayedReward = baseReward;
        for (uint256 i = 0; i < completedCycles; i++) {
            // Check if we've reached max decay
            uint256 currentDecay = (i + 1) * decayRate;
            if (currentDecay > maxDecay) {
                currentDecay = maxDecay;
                break;
            }
            
            // Apply compound decay: each cycle reduces by decayRate
            decayedReward = decayedReward * (10000 - decayRate) / 10000;
        }
        
        return decayedReward;
    }
    
    /**
     * @dev Calculate combo bonus for user's level with next-day effect
     */
    function _calculateComboBonus(address user, uint8 level) internal view returns (uint256) {
        ComboStatus memory status = userComboStatus[user][level];
        
        // If there's a pending status change and it's time to apply it
        if (status.isPending && _getCurrentTimestamp() >= status.effectiveFrom) {
            return status.bonus;
        }
        
        // No combo bonus until the waiting period is over
        return 0;
    }
    
    /**
     * @dev Get wait days for combo bonus based on count
     */
    function _getWaitDaysForCount(uint256 count, uint256[3] memory minDays) internal view returns (uint256) {
        // Get thresholds from config contract
        uint256[3] memory thresholds = configContract.getComboThresholds();
        
        // Check thresholds in reverse order (highest first)
        for (uint256 i = thresholds.length; i > 0; i--) {
            uint256 threshold = thresholds[i - 1];
            
            if (count >= threshold) {
                return minDays[i - 1];
            }
        }
        
        return 0;
    }
    
    /**
     * @dev Calculate combo bonus based on count and level
     */
    function _calculateComboBonusByCount(uint256 count, uint8 level) internal view returns (uint256) {
        // SSS级别(level 6)不适用10个NFT组合
        if (level == 6 && count >= 10) {
            return 0;
        }
        
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
    
    /**
     * @dev Calculate dynamic balance multiplier
     */
    function _calculateDynamicMultiplier(uint8 level) internal view returns (uint256) {
        uint256 staked = totalStakedPerLevel[level];
        uint256 supply = _getLevelSupply(level);
        
        if (supply == 0) return 10000; // 1.0x
        
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
        
        return 10000; // 1.0x
    }
    
    /**
     * @dev Get historical dynamic multiplier for a specific level and timestamp
     * @param level NFT level (1-6)
     * @param timestamp The timestamp to get the dynamic multiplier for
     * @return The dynamic multiplier that was active at that timestamp
     */
    function _getHistoricalDynamicMultiplier(uint8 level, uint256 timestamp) internal view returns (uint256) {
        // Find the most recent historical adjustment before or at the given timestamp
        for (uint256 i = historicalAdjustments.length; i > 0; i--) {
            HistoricalAdjustment storage adjustment = historicalAdjustments[i - 1];
            if (adjustment.timestamp <= timestamp) {
                return adjustment.dynamicMultipliers[level];
            }
        }
        
        // If no historical adjustment found, use current dynamic multiplier as fallback
        return _calculateDynamicMultiplier(level);
    }
    
    /**
     * @dev Get historical quarterly multiplier for a specific timestamp
     * @param timestamp The timestamp to get the quarterly multiplier for
     * @return The quarterly multiplier that was active at that timestamp
     */
    function _getHistoricalQuarterlyMultiplier(uint256 timestamp) internal view returns (uint256) {
        // Find the most recent historical adjustment before or at the given timestamp
        for (uint256 i = historicalAdjustments.length; i > 0; i--) {
            HistoricalAdjustment storage adjustment = historicalAdjustments[i - 1];
            if (adjustment.timestamp <= timestamp) {
                return adjustment.quarterlyMultiplier;
            }
        }
        
        // If no historical adjustment found, use current quarterly multiplier as fallback
        return configContract.getQuarterlyMultiplier();
    }
    
    /**
     * @dev Calculate continuous staking bonus (based on rewards at threshold)
     */
    function _calculateContinuousBonus(uint256 tokenId, uint256 totalStakingDays) internal view returns (uint256) {
        uint256[2] memory thresholds = configContract.getContinuousThresholds();
        uint256[2] memory bonuses = configContract.getContinuousBonuses();
        
        uint256 applicableThreshold = 0;
        uint256 applicableBonus = 0;
        
        for (uint256 i = thresholds.length; i > 0; i--) {
            if (totalStakingDays >= thresholds[i - 1]) {
                applicableThreshold = thresholds[i - 1];
                applicableBonus = bonuses[i - 1];
                break;
            }
        }
        
        if (applicableBonus == 0) return 0;
        
        StakeInfo memory stakeInfo = stakes[tokenId];
        uint256 rewardsAtThreshold = 0;
        uint256 baseReward = configContract.getDailyReward(stakeInfo.level);
        uint256 decayInterval = configContract.getDecayInterval(stakeInfo.level);
        uint256 decayRate = configContract.getDecayRate(stakeInfo.level);
        
        for (uint256 day = 0; day < applicableThreshold; day++) {
            uint256 currentDayTimestamp = stakeInfo.stakeTime + (day * 1 days);
            uint256 dailyReward = baseReward;
            
            if (decayInterval > 0 && day >= decayInterval) {
                uint256 completedCycles = day / decayInterval;
                for (uint256 i = 0; i < completedCycles; i++) {
                    uint256 totalDecaySoFar = (i + 1) * decayRate;
                    if (totalDecaySoFar > configContract.getMaxDecayRate(stakeInfo.level)) {
                        uint256 remainingDecay = configContract.getMaxDecayRate(stakeInfo.level) - (i * decayRate);
                        dailyReward = dailyReward * (10000 - remainingDecay) / 10000;
                        break;
                    }
                    dailyReward = dailyReward * (10000 - decayRate) / 10000;
                }
            }
            
            dailyReward = dailyReward * _getHistoricalQuarterlyMultiplier(currentDayTimestamp) / 10000;
            dailyReward = dailyReward * _getHistoricalDynamicMultiplier(stakeInfo.level, currentDayTimestamp) / 10000;
            rewardsAtThreshold += dailyReward;
        }
        
        uint256 comboBonus = _calculateComboBonus(stakeInfo.owner, stakeInfo.level);
        rewardsAtThreshold = rewardsAtThreshold * (10000 + comboBonus) / 10000;
        
        return rewardsAtThreshold * applicableBonus / 10000;
    }
    
    /**
     * @dev Update combo status for next-day effect
     */
    function _updateComboStatus(address user, uint8 level) internal {
        uint256 currentCount = userLevelCounts[user][level];
        ComboStatus storage status = userComboStatus[user][level];
        
        // Calculate what the new bonus should be
        uint256 newBonus = _calculateComboBonusByCount(currentCount, level);
        
        // If there's a change in count, schedule it for the configured wait period
        if (status.count != currentCount) {
            status.level = level;
            status.count = currentCount;
            
            // Update if the bonus level changes (upgrade or downgrade)
            if (newBonus != status.bonus) {
                // Get wait days from config based on count
                uint256[3] memory minDays = configContract.getComboMinDays();
                uint256 waitDays = _getWaitDaysForCount(currentCount, minDays);
                
                status.effectiveFrom = _getCurrentTimestamp() + (waitDays * 1 days);
                status.bonus = newBonus;
                status.isPending = true;
            }
            // If bonus level is the same, keep the current status unchanged
        }
    }
    
    // ============================================
    // INTERNAL FUNCTIONS
    // ============================================
    
    /**
     * @dev Remove user stake from array
     */
    function _removeUserStake(address user, uint256 tokenId) internal {
        uint256[] storage userTokenIds = userStakes[user];
        for (uint256 i = 0; i < userTokenIds.length; i++) {
            if (userTokenIds[i] == tokenId) {
                userTokenIds[i] = userTokenIds[userTokenIds.length - 1];
                userTokenIds.pop();
                break;
            }
        }
    }
    
    /**
     * @dev Get AA account address for user
     */
    function _getAAAccount(address user) internal view returns (address) {
        address masterSigner = accountManagerContract.getDefaultMasterSigner();
        require(masterSigner != address(0), "Master signer not found");
        
        address aaAccount = accountManagerContract.getAccountAddress(user, masterSigner);
        require(aaAccount != address(0), "AA account not found");
        
        return aaAccount;
    }
    
    // ============================================
    // ADMIN FUNCTIONS
    // ============================================
    
    function pause() external onlyOwner {
        _pause();
    }
    
    function unpause() external onlyOwner {
        _unpause();
    }
    
    function updateContracts(
        address _cpnftContract,
        address _cpopTokenContract,
        address _accountManagerContract,
        address _configContract
    ) external onlyOwner {
        if (_cpnftContract != address(0)) {
            cpnftContract = CPNFT(_cpnftContract);
        }
        if (_cpopTokenContract != address(0)) {
            cpopTokenContract = ICPOPToken(_cpopTokenContract);
        }
        if (_accountManagerContract != address(0)) {
            accountManagerContract = IAccountManager(_accountManagerContract);
        }
        if (_configContract != address(0)) {
            configContract = StakingConfig(_configContract);
        }
    }
    
    
    // ============================================
    // HISTORICAL ADJUSTMENT FUNCTIONS
    // ============================================
    
    /**
     * @dev Record current state as historical adjustment
     * Called when quarterly adjustments are executed
     */
    function recordHistoricalAdjustment() external {
        require(msg.sender == owner() || msg.sender == address(configContract), "Not authorized");
        historicalAdjustments.push();
        HistoricalAdjustment storage newRecord = historicalAdjustments[historicalAdjustments.length - 1];
        
        newRecord.timestamp = _getCurrentTimestamp();
        
        // Get current quarterly multiplier
        (, , uint256 quarterlyMultiplier, ) = configContract.getBasicConfig();
        newRecord.quarterlyMultiplier = quarterlyMultiplier;
        
        // Record current dynamic multipliers for all levels
        for (uint8 level = 1; level <= 6; level++) {
            newRecord.dynamicMultipliers[level] = _calculateDynamicMultiplier(level);
        }
    }
    
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
        require(index < historicalAdjustments.length, "Index out of range");
        HistoricalAdjustment storage record = historicalAdjustments[index];
        return (record.timestamp, record.quarterlyMultiplier);
    }
    
    /**
     * @dev Get historical dynamic multiplier for a specific level and time
     * @param index Index of the adjustment record
     * @param level NFT level (1-6)
     * @return Dynamic multiplier at that time
     */
    function getHistoricalDynamicMultiplier(uint256 index, uint8 level) external view returns (uint256) {
        require(index < historicalAdjustments.length, "Index out of range");
        require(level >= 1 && level <= 6, "Invalid level");
        return historicalAdjustments[index].dynamicMultipliers[level];
    }
    
    /**
     * @dev Get number of historical adjustments
     * @return Number of adjustment records
     */
    function getHistoricalAdjustmentCount() external view returns (uint256) {
        return historicalAdjustments.length;
    }
    
    /**
     * @dev Get combo status for a user and level
     * @param user User address
     * @param level NFT level
     * @return ComboStatus struct with current combo information
     */
    function getComboStatus(address user, uint8 level) external view returns (ComboStatus memory) {
        return userComboStatus[user][level];
    }
    
    /**
     * @dev Get effective combo bonus for a user and level
     * @param user User address
     * @param level NFT level
     * @return The currently effective combo bonus
     */
    function getEffectiveComboBonus(address user, uint8 level) external view returns (uint256) {
        return _calculateComboBonus(user, level);
    }
    
    /**
     * @dev Get historical quarterly multiplier for a specific timestamp
     * @param timestamp The timestamp to get the quarterly multiplier for
     * @return The quarterly multiplier that was active at that timestamp
     */
    function getHistoricalQuarterlyMultiplier(uint256 timestamp) external view returns (uint256) {
        return _getHistoricalQuarterlyMultiplier(timestamp);
    }
    
    /**
     * @dev Calculate detailed reward breakdown for a specific day
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
        StakeInfo memory stakeInfo = stakes[tokenId];
        require(stakeInfo.isActive, "NFT is not staked");
        
        uint256 currentDayTimestamp = stakeInfo.stakeTime + (dayOffset * 1 days);
        
        baseReward = configContract.getDailyReward(stakeInfo.level);
        decayedReward = baseReward;
        
        // Apply decay if applicable
        uint256 decayInterval = configContract.getDecayInterval(stakeInfo.level);
        uint256 decayRate = configContract.getDecayRate(stakeInfo.level);
        
        if (decayInterval > 0 && dayOffset >= decayInterval) {
            uint256 completedCycles = dayOffset / decayInterval;
            
            for (uint256 i = 0; i < completedCycles; i++) {
                uint256 totalDecaySoFar = (i + 1) * decayRate;
                if (totalDecaySoFar > configContract.getMaxDecayRate(stakeInfo.level)) {
                    uint256 remainingDecay = configContract.getMaxDecayRate(stakeInfo.level) - (i * decayRate);
                    decayedReward = decayedReward * (10000 - remainingDecay) / 10000;
                    break;
                }
                
                decayedReward = decayedReward * (10000 - decayRate) / 10000;
            }
        }
        
        // Get historical multipliers for this day
        quarterlyMultiplier = _getHistoricalQuarterlyMultiplier(currentDayTimestamp);
        dynamicMultiplier = _getHistoricalDynamicMultiplier(stakeInfo.level, currentDayTimestamp);
        
        // Calculate final reward
        finalReward = decayedReward * quarterlyMultiplier / 10000;
        finalReward = finalReward * dynamicMultiplier / 10000;
    }
    
    // ============================================
    // MINIMAL VIEW FUNCTIONS
    // ============================================
    
    function version() public pure returns (string memory) {
        return "4.0.0";
    }
}