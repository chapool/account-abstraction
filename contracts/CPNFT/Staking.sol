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
    mapping(uint8 => uint256) public totalSupplyPerLevel;
    
    // ============================================
    // EVENTS
    // ============================================
    
    event NFTStaked(address indexed user, uint256 indexed tokenId, uint8 level, uint256 timestamp);
    event NFTUnstaked(address indexed user, uint256 indexed tokenId, uint256 rewards, uint256 timestamp);
    event RewardsClaimed(address indexed user, uint256 indexed tokenId, uint256 amount, uint256 timestamp);
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
        emit PlatformStatsUpdated(level, totalStakedPerLevel[level], totalSupplyPerLevel[level]);
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
        emit PlatformStatsUpdated(stakeInfo.level, totalStakedPerLevel[stakeInfo.level], totalSupplyPerLevel[stakeInfo.level]);
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
            tokenId,
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
     * @dev Calculate pending rewards since last claim
     */
    function _calculatePendingRewards(uint256 tokenId) internal view returns (uint256) {
        StakeInfo memory stakeInfo = stakes[tokenId];
        
        // Calculate base rewards
        uint256 baseRewards = _calculateRewards(
            tokenId,
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
     * @dev Calculate rewards for a specific NFT
     */
    function _calculateRewards(
        uint256 tokenId,
        uint8 level,
        uint256 stakeTime,
        uint256 lastClaimTime
    ) internal view returns (uint256) {
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
        uint256 supply = totalSupplyPerLevel[level];
        
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
    
    function setTotalSupplyPerLevel(uint8 level, uint256 supply) external onlyOwner {
        totalSupplyPerLevel[level] = supply;
        emit PlatformStatsUpdated(level, totalStakedPerLevel[level], supply);
    }
    
    // ============================================
    // MINIMAL VIEW FUNCTIONS
    // ============================================
    
    function version() public pure returns (string memory) {
        return "3.0.0";
    }
}