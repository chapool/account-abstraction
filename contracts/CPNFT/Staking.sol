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
        
        emit BatchStaked(msg.sender, tokenIds, block.timestamp);
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
            stakeTime: block.timestamp,
            lastClaimTime: block.timestamp,
            isActive: true,
            totalRewards: 0,
            pendingRewards: 0
        });
        
        // Update user data
        userStakes[msg.sender].push(tokenId);
        userLevelCounts[msg.sender][level]++;
        
        // Update platform stats
        totalStakedPerLevel[level]++;
        totalStakedCount++;
        
        // Set NFT as staked in CPNFT contract
        cpnftContract.setStakeStatus(tokenId, true);
        
        emit NFTStaked(msg.sender, tokenId, level, block.timestamp);
        // Get supply from CPNFT contract
        uint256 supply;
        if (level == 1) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.C);
        else if (level == 2) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.B);
        else if (level == 3) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.A);
        else if (level == 4) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.S);
        else if (level == 5) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.SS);
        else if (level == 6) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.SSS);
        
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
        uint256 stakingDays = (block.timestamp - stakeInfo.stakeTime) / 1 days;
        
        if (stakingDays < minStakeDays) {
            uint256 penalty = configContract.getEarlyWithdrawPenalty();
            rewards = rewards * (10000 - penalty) / 10000;
        }
        
        // Add continuous staking bonus
        uint256 continuousBonus = _calculateContinuousBonus(rewards, stakingDays);
        rewards += continuousBonus;
        
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
        
        // Update platform stats
        totalStakedPerLevel[stakeInfo.level]--;
        totalStakedCount--;
        
        // Set NFT as not staked in CPNFT contract
        cpnftContract.setStakeStatus(tokenId, false);
        
        emit NFTUnstaked(msg.sender, tokenId, rewards, block.timestamp);
        // Get supply from CPNFT contract
        uint256 supply;
        if (stakeInfo.level == 1) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.C);
        else if (stakeInfo.level == 2) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.B);
        else if (stakeInfo.level == 3) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.A);
        else if (stakeInfo.level == 4) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.S);
        else if (stakeInfo.level == 5) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.SS);
        else if (stakeInfo.level == 6) supply = cpnftContract.getLevelSupply(CPNFT.NFTLevel.SSS);
        
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
            uint256 stakingDays = (block.timestamp - stakeInfo.stakeTime) / 1 days;
            
            if (stakingDays < minStakeDays) {
                uint256 penalty = configContract.getEarlyWithdrawPenalty();
                rewards = rewards * (10000 - penalty) / 10000;
            }
            
            // Add continuous staking bonus
            uint256 continuousBonus = _calculateContinuousBonus(rewards, stakingDays);
            rewards += continuousBonus;
            
            totalRewards += rewards;
            
            // Update stake info
            stakeInfo.isActive = false;
            stakeInfo.totalRewards = rewards;
            stakeInfo.pendingRewards = 0;
            
            // Update user data
            _removeUserStake(msg.sender, tokenId);
            userLevelCounts[msg.sender][stakeInfo.level]--;
            
            // Update platform stats
            totalStakedPerLevel[stakeInfo.level]--;
            totalStakedCount--;
            
            // Set NFT as not staked in CPNFT contract
            cpnftContract.setStakeStatus(tokenId, false);
        }
        
        // Send total rewards to AA account
        if (totalRewards > 0) {
            address aaAccount = _getAAAccount(msg.sender);
            cpopTokenContract.mint(aaAccount, totalRewards);
        }
        
        emit BatchUnstaked(msg.sender, tokenIds, totalRewards, block.timestamp);
    }
    
    // ============================================
    // REWARD CALCULATION FUNCTIONS
    // ============================================
    
    /**
     * @dev Calculate pending rewards since last claim (internal)
     */
    function _calculatePendingRewards(uint256 tokenId) internal view returns (uint256) {
        StakeInfo memory stakeInfo = stakes[tokenId];
        
        // Calculate base rewards with phase-based decay and dynamic adjustment
        uint256 totalDays = (block.timestamp - stakeInfo.lastClaimTime) / 1 days;
        if (totalDays == 0) return 0;
        
        uint256 baseReward = configContract.getDailyReward(stakeInfo.level);
        uint256 decayInterval = configContract.getDecayInterval(stakeInfo.level);
        uint256 decayRate = configContract.getDecayRate(stakeInfo.level);
        uint256 quarterlyMultiplier = configContract.getQuarterlyMultiplier();
        
        uint256 totalRewards = 0;
        
        // Calculate rewards day by day with phase-based decay and dynamic adjustment
        for (uint256 day = 0; day < totalDays; day++) {
            uint256 currentDayFromStake = (block.timestamp - stakeInfo.stakeTime) / 1 days - (totalDays - 1 - day);
            
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
            
            // Apply quarterly adjustment
            dailyReward = dailyReward * quarterlyMultiplier / 10000;
            
            // Apply historical dynamic multiplier
            uint256 dynamicMultiplier = _getHistoricalDynamicMultiplier(
                stakeInfo.level, 
                stakeInfo.stakeTime + (currentDayFromStake * 1 days)
            );
            dailyReward = dailyReward * dynamicMultiplier / 10000;
            
            totalRewards += dailyReward;
        }
        
        // Calculate combo bonus
        uint256 comboBonus = _calculateComboBonus(msg.sender, stakeInfo.level);
        totalRewards = totalRewards * (10000 + comboBonus) / 10000;
        
        // Add continuous staking bonus
        uint256 stakingDays = (block.timestamp - stakeInfo.stakeTime) / 1 days;
        uint256 continuousBonus = _calculateContinuousBonus(totalRewards, stakingDays);
        totalRewards += continuousBonus;
        
        return totalRewards;
    }

    /**
     * @dev Calculate pending rewards since last claim (external for Reader)
     */
    function calculatePendingRewards(uint256 tokenId) external view returns (uint256) {
        return _calculatePendingRewards(tokenId);
    }

    /**
     * @dev Claim pending rewards without unstaking
     */
    function claimRewards(uint256 tokenId) external nonReentrant whenNotPaused {
        require(stakes[tokenId].owner == msg.sender, "Not the owner of this stake");
        require(stakes[tokenId].isActive, "NFT is not staked");
        
        StakeInfo storage stakeInfo = stakes[tokenId];
        
        // Calculate rewards since last claim
        uint256 rewards = _calculatePendingRewards(tokenId);
        require(rewards > 0, "No pending rewards");
        
        // Update last claim time
        stakeInfo.lastClaimTime = block.timestamp;
        stakeInfo.pendingRewards = 0;
        stakeInfo.totalRewards += rewards;
        
        // Send rewards to AA account
        address aaAccount = _getAAAccount(msg.sender);
        cpopTokenContract.mint(aaAccount, rewards);
        
        emit RewardsClaimed(msg.sender, tokenId, rewards, block.timestamp);
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
            
            // Calculate rewards since last claim
            uint256 rewards = _calculatePendingRewards(tokenId);
            
            if (rewards > 0) {
                // Update last claim time
                stakeInfo.lastClaimTime = block.timestamp;
                stakeInfo.pendingRewards = 0;
                stakeInfo.totalRewards += rewards;
                totalRewards += rewards;
                
                emit RewardsClaimed(msg.sender, tokenId, rewards, block.timestamp);
            }
        }
        
        require(totalRewards > 0, "No pending rewards");
        
        // Send total rewards to AA account
        cpopTokenContract.mint(aaAccount, totalRewards);
        
        emit BatchRewardsClaimed(msg.sender, tokenIds.length, totalRewards, block.timestamp);
    }
    
    // ============================================
    // REWARD CALCULATION FUNCTIONS
    // ============================================
    
    /**
     * @dev Calculate total rewards for a staked NFT
     */
    function _calculateTotalRewards(uint256 tokenId) internal view returns (uint256) {
        StakeInfo memory stakeInfo = stakes[tokenId];
        
        // Calculate base rewards
        uint256 baseRewards = _calculateRewards(
            stakeInfo.level,
            stakeInfo.stakeTime,
            stakeInfo.lastClaimTime
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
        uint256 lastClaimTime
    ) internal view returns (uint256) {
        uint256 totalDays = (block.timestamp - lastClaimTime) / 1 days;
        if (totalDays == 0) return 0;
        
        uint256 baseReward = configContract.getDailyReward(level);
        uint256 decayInterval = configContract.getDecayInterval(level);
        uint256 decayRate = configContract.getDecayRate(level);
        uint256 quarterlyMultiplier = configContract.getQuarterlyMultiplier();
        
        uint256 totalRewards = 0;
        
        // Calculate rewards day by day with proper phase-based decay
        for (uint256 day = 0; day < totalDays; day++) {
            uint256 currentDayFromStake = (block.timestamp - stakeTime) / 1 days - (totalDays - 1 - day);
            
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
            
            // Apply quarterly adjustment
            dailyReward = dailyReward * quarterlyMultiplier / 10000;
            
            // Apply historical dynamic multiplier
            uint256 dynamicMultiplier = _getHistoricalDynamicMultiplier(
                level, 
                stakeTime + (currentDayFromStake * 1 days)
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
        uint256 stakingDays = (block.timestamp - stakeTime) / 1 days;
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
     * @dev Calculate combo bonus for user's level
     */
    function _calculateComboBonus(address user, uint8 level) internal view returns (uint256) {
        uint256 count = userLevelCounts[user][level];
        
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
        uint256 staked = totalStakedPerLevel[level];
        
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
     * @dev Calculate continuous staking bonus
     */
    function _calculateContinuousBonus(uint256 baseRewards, uint256 totalDays) internal view returns (uint256) {
        uint256[2] memory thresholds = configContract.getContinuousThresholds();
        uint256[2] memory bonuses = configContract.getContinuousBonuses();
        
        for (uint256 i = thresholds.length; i > 0; i--) {
            if (totalDays >= thresholds[i - 1]) {
                return baseRewards * bonuses[i - 1] / 10000;
            }
        }
        return 0;
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
        try accountManagerContract.getDefaultMasterSigner() returns (address masterSigner) {
            if (masterSigner != address(0)) {
                try accountManagerContract.getAccountAddress(user, masterSigner) returns (address aaAccount) {
                    if (aaAccount != address(0)) {
                        return aaAccount;
                    }
                } catch {
                    // Fallback to user address if AA account not found
                }
            }
        } catch {
            // Fallback to user address if master signer not found
        }
        return user;
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
        
        newRecord.timestamp = block.timestamp;
        
        // Get current quarterly multiplier
        (uint256 minStakeDays, uint256 earlyWithdrawPenalty, uint256 quarterlyMultiplier, uint256 lastQuarterlyUpdate) = configContract.getBasicConfig();
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
    
    // ============================================
    // MINIMAL VIEW FUNCTIONS
    // ============================================
    
    function version() public pure returns (string memory) {
        return "3.1.0";
    }
}