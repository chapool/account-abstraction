// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title StakingConfig - Configuration Management Contract
 * @dev Optimized storage using packed structs and efficient data structures
 */
contract StakingConfig is Ownable {
    
    // Reference to Staking contract for automatic historical recording
    address public stakingContract;
    
    // ============================================
    // STORAGE STRUCTS (Packed for Gas Efficiency)
    // ============================================
    
    // Basic configuration (3 slots)
    struct BasicConfig {
        uint64 minStakeDays;        // 0-365 days
        uint64 earlyWithdrawPenalty; // 0-10000 (basis points)
        uint64 quarterlyMultiplier;  // 8000-12000 (basis points)
        uint64 lastQuarterlyUpdate;  // timestamp (fits in 64 bits until 2106)
    }
    
    // Level-specific configuration (packed)
    struct LevelConfig {
        uint32 dailyReward;    // Max 4.2B tokens
        uint32 decayInterval;  // Max 4.2B days
        uint16 decayRate;      // Max 65535 (basis points)
        uint16 maxDecayRate;   // Max 65535 (basis points)
    }
    
    // Combo configuration (packed)
    struct ComboConfig {
        uint8 threshold;       // Max 255 NFTs
        uint16 bonus;          // Max 65535 (basis points)
        uint8 minDays;         // Max 255 days
        uint8 _padding;        // Unused padding
    }
    
    // Continuous staking configuration (packed)
    struct ContinuousConfig {
        uint16 threshold;      // Max 65535 days
        uint16 bonus;          // Max 65535 (basis points)
        uint16 _padding1;      // Unused padding
        uint16 _padding2;      // Unused padding
    }
    
    // Dynamic balance configuration (packed)
    struct DynamicConfig {
        uint16 highStakeThreshold;    // Max 65535 (basis points)
        uint16 lowStakeThreshold;     // Max 65535 (basis points)
        uint16 highStakeMultiplier;   // Max 65535 (basis points)
        uint16 lowStakeMultiplier;    // Max 65535 (basis points)
        uint16 quarterlyAdjustmentMax; // Max 65535 (basis points)
        uint16 _padding1;             // Unused padding
        uint16 _padding2;             // Unused padding
        uint16 _padding3;             // Unused padding
    }
    
    // ============================================
    // STORAGE VARIABLES
    // ============================================
    
    BasicConfig public basicConfig;
    LevelConfig[7] public levelConfigs;        // Index 0 = NORMAL (unused), 1-6 = C to SSS
    ComboConfig[3] public comboConfigs;        // 3, 5, 10 NFT combos
    ContinuousConfig[2] public continuousConfigs; // 30, 90 day rewards
    DynamicConfig public dynamicConfig;
    
    
    // Quarterly adjustment history
    struct QuarterlyAdjustmentRecord {
        uint256 multiplier;
        uint256 timestamp;
        uint256 announcementTime;
        bool isActive;
    }
    
    QuarterlyAdjustmentRecord[] public quarterlyAdjustments;
    uint256 public nextQuarterlyUpdate;
    
    // ============================================
    // EVENTS
    // ============================================
    
    event ConfigUpdated(string configType, address indexed updater);
    event QuarterlyAdjustment(uint256 newMultiplier, uint256 timestamp);
    event QuarterlyAdjustmentAnnounced(uint256 newMultiplier, uint256 effectiveTime);
    event QuarterlyAdjustmentExecuted(uint256 index, uint256 multiplier, uint256 timestamp);
    
    // ============================================
    // CONSTRUCTOR
    // ============================================
    
    constructor() Ownable(msg.sender) {
        _initializeDefaultConfig();
    }
    
    // ============================================
    // INITIALIZATION
    // ============================================
    
    function _initializeDefaultConfig() internal {
        // Basic configuration
        basicConfig = BasicConfig({
            minStakeDays: 7,
            earlyWithdrawPenalty: 5000, // 50%
            quarterlyMultiplier: 10000,  // 1.0x
            lastQuarterlyUpdate: uint64(block.timestamp)
        });
        
        // Initialize quarterly update schedule
        nextQuarterlyUpdate = block.timestamp + 90 days;
        
        // Level configurations (C=1, B=2, A=3, S=4, SS=5, SSS=6)
        levelConfigs[1] = LevelConfig(3, 20, 3500, 8000);   // C level
        levelConfigs[2] = LevelConfig(8, 30, 3000, 7000);   // B level
        levelConfigs[3] = LevelConfig(15, 45, 2500, 6000);  // A level
        levelConfigs[4] = LevelConfig(30, 60, 2000, 5000);  // S level
        levelConfigs[5] = LevelConfig(50, 90, 1500, 4000);  // SS level
        levelConfigs[6] = LevelConfig(100, 180, 1000, 2000); // SSS level
        
        // Combo configurations
        comboConfigs[0] = ComboConfig(3, 500, 7, 0);    // 3 NFTs: 5% bonus, 7 days
        comboConfigs[1] = ComboConfig(5, 1000, 15, 0);  // 5 NFTs: 10% bonus, 15 days
        comboConfigs[2] = ComboConfig(10, 2000, 30, 0); // 10 NFTs: 20% bonus, 30 days
        
        // Continuous staking configurations
        continuousConfigs[0] = ContinuousConfig(30, 1000, 0, 0);  // 30 days: 10% bonus
        continuousConfigs[1] = ContinuousConfig(90, 3000, 0, 0);  // 90 days: 30% bonus
        
        // Dynamic balance configuration
        dynamicConfig = DynamicConfig({
            highStakeThreshold: 6000,    // 60%
            lowStakeThreshold: 1000,     // 10%
            highStakeMultiplier: 8000,   // 0.8x
            lowStakeMultiplier: 10500,   // 1.05x
            quarterlyAdjustmentMax: 2000, // 20%
            _padding1: 0,
            _padding2: 0,
            _padding3: 0
        });
    }
    
    // ============================================
    // ADMIN FUNCTIONS
    // ============================================
    
    /**
     * @dev Set the Staking contract address
     * @param _stakingContract Address of the Staking contract
     */
    function setStakingContract(address _stakingContract) external onlyOwner {
        require(_stakingContract != address(0), "Invalid address");
        stakingContract = _stakingContract;
        emit ConfigUpdated("staking_contract_set", msg.sender);
    }
    
    function updateBasicConfig(
        uint64 _minStakeDays,
        uint64 _earlyWithdrawPenalty
    ) external onlyOwner {
        require(_minStakeDays > 0 && _minStakeDays <= 365, "Invalid min stake days");
        require(_earlyWithdrawPenalty <= 10000, "Invalid penalty rate");
        
        basicConfig.minStakeDays = _minStakeDays;
        basicConfig.earlyWithdrawPenalty = _earlyWithdrawPenalty;
        
        emit ConfigUpdated("basic", msg.sender);
    }
    
    function updateRewards(
        uint32[6] calldata _rewards
    ) external onlyOwner {
        require(_rewards.length == 6, "Invalid rewards array length");
        
        for (uint256 i = 0; i < 6; i++) {
            levelConfigs[i + 1].dailyReward = _rewards[i];
        }
        
        emit ConfigUpdated("rewards", msg.sender);
    }
    
    function updateDecayConfig(
        uint32[6] calldata _intervals,
        uint16[6] calldata _rates,
        uint16[6] calldata _maxRates
    ) external onlyOwner {
        require(_intervals.length == 6 && _rates.length == 6 && _maxRates.length == 6, "Invalid array lengths");
        
        for (uint256 i = 0; i < 6; i++) {
            levelConfigs[i + 1].decayInterval = _intervals[i];
            levelConfigs[i + 1].decayRate = _rates[i];
            levelConfigs[i + 1].maxDecayRate = _maxRates[i];
        }
        
        emit ConfigUpdated("decay", msg.sender);
    }
    
    function updateComboConfig(
        uint8[3] calldata _thresholds,
        uint16[3] calldata _bonuses,
        uint8[3] calldata _minDays
    ) external onlyOwner {
        require(_thresholds.length == 3 && _bonuses.length == 3 && _minDays.length == 3, "Invalid array lengths");
        
        for (uint256 i = 0; i < 3; i++) {
            comboConfigs[i] = ComboConfig({
                threshold: _thresholds[i],
                bonus: _bonuses[i],
                minDays: _minDays[i],
                _padding: 0
            });
        }
        
        emit ConfigUpdated("combo", msg.sender);
    }
    
    function updateContinuousRewards(
        uint16[2] calldata _thresholds,
        uint16[2] calldata _bonuses
    ) external onlyOwner {
        require(_thresholds.length == 2 && _bonuses.length == 2, "Invalid array lengths");
        
        for (uint256 i = 0; i < 2; i++) {
            continuousConfigs[i] = ContinuousConfig({
                threshold: _thresholds[i],
                bonus: _bonuses[i],
                _padding1: 0,
                _padding2: 0
            });
        }
        
        emit ConfigUpdated("continuous", msg.sender);
    }
    
    function updateDynamicConfig(
        uint16 _highStakeThreshold,
        uint16 _lowStakeThreshold,
        uint16 _highStakeMultiplier,
        uint16 _lowStakeMultiplier,
        uint16 _quarterlyAdjustmentMax
    ) external onlyOwner {
        dynamicConfig = DynamicConfig({
            highStakeThreshold: _highStakeThreshold,
            lowStakeThreshold: _lowStakeThreshold,
            highStakeMultiplier: _highStakeMultiplier,
            lowStakeMultiplier: _lowStakeMultiplier,
            quarterlyAdjustmentMax: _quarterlyAdjustmentMax,
            _padding1: 0,
            _padding2: 0,
            _padding3: 0
        });
        
        emit ConfigUpdated("dynamic", msg.sender);
    }
    
    function updateQuarterlyMultiplier(uint64 newMultiplier) external onlyOwner {
        require(newMultiplier >= 8000 && newMultiplier <= 12000, "Multiplier out of range");
        basicConfig.quarterlyMultiplier = newMultiplier;
        basicConfig.lastQuarterlyUpdate = uint64(block.timestamp);
        emit QuarterlyAdjustment(newMultiplier, block.timestamp);
    }
    
    // ============================================
    // VIEW FUNCTIONS
    // ============================================
    
    function getDailyReward(uint256 level) external view returns (uint256) {
        require(level > 0 && level < 7, "Invalid level");
        return levelConfigs[level].dailyReward;
    }
    
    function getDecayInterval(uint256 level) external view returns (uint256) {
        require(level > 0 && level < 7, "Invalid level");
        return levelConfigs[level].decayInterval;
    }
    
    function getDecayRate(uint256 level) external view returns (uint256) {
        require(level > 0 && level < 7, "Invalid level");
        return levelConfigs[level].decayRate;
    }
    
    function getMaxDecayRate(uint256 level) external view returns (uint256) {
        require(level > 0 && level < 7, "Invalid level");
        return levelConfigs[level].maxDecayRate;
    }
    
    function getMinStakeDays() external view returns (uint256) {
        return basicConfig.minStakeDays;
    }
    
    function getEarlyWithdrawPenalty() external view returns (uint256) {
        return basicConfig.earlyWithdrawPenalty;
    }
    
    function getQuarterlyMultiplier() external view returns (uint256) {
        return basicConfig.quarterlyMultiplier;
    }
    
    function getLastQuarterlyUpdate() external view returns (uint256) {
        return basicConfig.lastQuarterlyUpdate;
    }
    
    function getComboThresholds() external view returns (uint256[3] memory) {
        return [
            uint256(comboConfigs[0].threshold),
            uint256(comboConfigs[1].threshold),
            uint256(comboConfigs[2].threshold)
        ];
    }
    
    function getComboBonuses() external view returns (uint256[3] memory) {
        return [
            uint256(comboConfigs[0].bonus),
            uint256(comboConfigs[1].bonus),
            uint256(comboConfigs[2].bonus)
        ];
    }
    
    function getComboMinDays() external view returns (uint256[3] memory) {
        return [
            uint256(comboConfigs[0].minDays),
            uint256(comboConfigs[1].minDays),
            uint256(comboConfigs[2].minDays)
        ];
    }
    
    function getContinuousThresholds() external view returns (uint256[2] memory) {
        return [
            uint256(continuousConfigs[0].threshold),
            uint256(continuousConfigs[1].threshold)
        ];
    }
    
    function getContinuousBonuses() external view returns (uint256[2] memory) {
        return [
            uint256(continuousConfigs[0].bonus),
            uint256(continuousConfigs[1].bonus)
        ];
    }
    
    function getDynamicConfig() external view returns (
        uint256 highStakeThreshold,
        uint256 lowStakeThreshold,
        uint256 highStakeMultiplier,
        uint256 lowStakeMultiplier,
        uint256 quarterlyAdjustmentMax
    ) {
        return (
            dynamicConfig.highStakeThreshold,
            dynamicConfig.lowStakeThreshold,
            dynamicConfig.highStakeMultiplier,
            dynamicConfig.lowStakeMultiplier,
            dynamicConfig.quarterlyAdjustmentMax
        );
    }
    
    // ============================================
    // BATCH GETTERS FOR EFFICIENCY
    // ============================================
    
    function getLevelConfig(uint256 level) external view returns (
        uint256 dailyReward,
        uint256 decayInterval,
        uint256 decayRate,
        uint256 maxDecayRate
    ) {
        require(level > 0 && level < 7, "Invalid level");
        LevelConfig memory config = levelConfigs[level];
        return (config.dailyReward, config.decayInterval, config.decayRate, config.maxDecayRate);
    }
    
    function getAllLevelConfigs() external view returns (
        uint256[6] memory dailyRewards,
        uint256[6] memory decayIntervals,
        uint256[6] memory decayRates,
        uint256[6] memory maxDecayRates
    ) {
        for (uint256 i = 0; i < 6; i++) {
            LevelConfig memory config = levelConfigs[i + 1];
            dailyRewards[i] = config.dailyReward;
            decayIntervals[i] = config.decayInterval;
            decayRates[i] = config.decayRate;
            maxDecayRates[i] = config.maxDecayRate;
        }
    }
    
    function getBasicConfig() external view returns (
        uint256 minStakeDays,
        uint256 earlyWithdrawPenalty,
        uint256 quarterlyMultiplier,
        uint256 lastQuarterlyUpdate
    ) {
        return (
            basicConfig.minStakeDays,
            basicConfig.earlyWithdrawPenalty,
            basicConfig.quarterlyMultiplier,
            basicConfig.lastQuarterlyUpdate
        );
    }
    
    // ============================================
    // SUPPLY CONFIGURATION FUNCTIONS
    // ============================================
    
    
    // ============================================
    // QUARTERLY ADJUSTMENT FUNCTIONS
    // ============================================
    
    /**
     * @dev Announce a quarterly adjustment (7 days before execution)
     * @param newMultiplier New quarterly multiplier (8000-12000, representing 0.8x-1.2x)
     */
    function announceQuarterlyAdjustment(uint256 newMultiplier) external onlyOwner {
        require(newMultiplier >= 8000 && newMultiplier <= 12000, "Multiplier out of range");
        require(block.timestamp >= nextQuarterlyUpdate - 90 days, "Too early for next announcement");
        
        uint256 effectiveTime = block.timestamp + 7 days;
        
        quarterlyAdjustments.push(QuarterlyAdjustmentRecord({
            multiplier: newMultiplier,
            timestamp: 0, // Will be set when executed
            announcementTime: block.timestamp,
            isActive: false
        }));
        
        emit QuarterlyAdjustmentAnnounced(newMultiplier, effectiveTime);
        emit ConfigUpdated("quarterly_announcement", msg.sender);
    }
    
    /**
     * @dev Execute the latest announced quarterly adjustment
     */
    function executeQuarterlyAdjustment() external onlyOwner {
        require(quarterlyAdjustments.length > 0, "No pending adjustments");
        
        QuarterlyAdjustmentRecord storage latest = quarterlyAdjustments[quarterlyAdjustments.length - 1];
        require(!latest.isActive, "Already executed");
        require(block.timestamp >= latest.announcementTime + 7 days, "Announcement period not complete");
        
        // Update basic config
        basicConfig.quarterlyMultiplier = uint64(latest.multiplier);
        basicConfig.lastQuarterlyUpdate = uint64(block.timestamp);
        
        // Mark as active and set execution time
        latest.isActive = true;
        latest.timestamp = block.timestamp;
        
        // Schedule next quarterly update
        nextQuarterlyUpdate = block.timestamp + 90 days;
        
        // Automatically record historical adjustment in Staking contract
        if (stakingContract != address(0)) {
            stakingContract.call(
                abi.encodeWithSignature("recordHistoricalAdjustment()")
            );
            // Continue with quarterly adjustment regardless of historical recording success
            // This ensures quarterly adjustment is not blocked by historical recording issues
        }
        
        emit QuarterlyAdjustmentExecuted(quarterlyAdjustments.length - 1, latest.multiplier, block.timestamp);
        emit ConfigUpdated("quarterly_execution", msg.sender);
    }
    
    /**
     * @dev Get quarterly adjustment history
     * @param index Index of the adjustment record
     * @return multiplier The quarterly multiplier
     * @return timestamp When the adjustment was executed
     * @return announcementTime When the adjustment was announced
     * @return isActive Whether the adjustment is currently active
     */
    function getQuarterlyAdjustment(uint256 index) external view returns (
        uint256 multiplier,
        uint256 timestamp,
        uint256 announcementTime,
        bool isActive
    ) {
        require(index < quarterlyAdjustments.length, "Index out of range");
        QuarterlyAdjustmentRecord memory record = quarterlyAdjustments[index];
        return (record.multiplier, record.timestamp, record.announcementTime, record.isActive);
    }
    
    /**
     * @dev Get the latest quarterly adjustment
     * @return multiplier The quarterly multiplier
     * @return timestamp When the adjustment was executed
     * @return announcementTime When the adjustment was announced
     * @return isActive Whether the adjustment is currently active
     */
    function getLatestQuarterlyAdjustment() external view returns (
        uint256 multiplier,
        uint256 timestamp,
        uint256 announcementTime,
        bool isActive
    ) {
        require(quarterlyAdjustments.length > 0, "No adjustments yet");
        QuarterlyAdjustmentRecord memory record = quarterlyAdjustments[quarterlyAdjustments.length - 1];
        return (record.multiplier, record.timestamp, record.announcementTime, record.isActive);
    }
    
    /**
     * @dev Get the number of quarterly adjustments
     * @return Number of adjustment records
     */
    function getQuarterlyAdjustmentCount() external view returns (uint256) {
        return quarterlyAdjustments.length;
    }
    
    /**
     * @dev Get next quarterly update timestamp
     * @return Timestamp when next quarterly update can be announced
     */
    function getNextQuarterlyUpdate() external view returns (uint256) {
        return nextQuarterlyUpdate;
    }
    
    function version() public pure returns (string memory) {
        return "3.2.0";
    }
}