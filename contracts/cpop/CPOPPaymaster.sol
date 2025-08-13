// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/utils/Pausable.sol";
import "../core/BasePaymaster.sol";
import "./interfaces/ICPOPPaymaster.sol";
import "./interfaces/ICPOPToken.sol";

/**
 * @title CPOPPaymaster
 * @notice Paymaster that allows users to pay gas fees with CPOP tokens
 * @dev Inherits from BasePaymaster and implements ICPOPPaymaster interface
 */
contract CPOPPaymaster is BasePaymaster, ICPOPPaymaster, Pausable {
    ICPOPToken public immutable cpopToken;
    
    uint256 public exchangeRate; // CPOP per wei (e.g., 1000 means 1000 CPOP = 1 wei)
    uint256 public constant DEFAULT_DAILY_LIMIT = 0.01 ether; // 0.01 ETH worth of gas per day
    
    mapping(address => uint256) private dailyLimits;
    mapping(address => uint256) private dailyUsage;
    mapping(address => uint256) private lastUsageDay;

    constructor(
        IEntryPoint _entryPoint,
        address _cpopToken,
        uint256 _initialExchangeRate
    ) BasePaymaster(_entryPoint) {
        require(_cpopToken != address(0), "CPOPPaymaster: invalid token");
        require(_initialExchangeRate > 0, "CPOPPaymaster: invalid rate");
        
        cpopToken = ICPOPToken(_cpopToken);
        exchangeRate = _initialExchangeRate;
    }

    /**
     * @notice Get the current exchange rate
     */
    function getExchangeRate() external view override returns (uint256) {
        return exchangeRate;
    }

    /**
     * @notice Set new exchange rate
     */
    function setExchangeRate(uint256 newRate) external override onlyOwner {
        require(newRate > 0, "CPOPPaymaster: invalid rate");
        uint256 oldRate = exchangeRate;
        exchangeRate = newRate;
        emit ExchangeRateUpdated(oldRate, newRate);
    }

    /**
     * @notice Get daily gas limit for a user
     */
    function getDailyLimit(address user) public view override returns (uint256) {
        uint256 limit = dailyLimits[user];
        return limit > 0 ? limit : DEFAULT_DAILY_LIMIT;
    }

    /**
     * @notice Set daily gas limit for a user
     */
    function setDailyLimit(address user, uint256 limit) external override onlyOwner {
        dailyLimits[user] = limit;
        emit DailyLimitUpdated(user, limit);
    }

    /**
     * @notice Get daily gas usage for a user
     */
    function getDailyUsage(address user) external view override returns (uint256) {
        if (_getCurrentDay() != lastUsageDay[user]) {
            return 0;
        }
        return dailyUsage[user];
    }

    /**
     * @notice Estimate CPOP cost for gas amount
     */
    function estimateCPOPCost(uint256 gasAmount) public view override returns (uint256) {
        return gasAmount * exchangeRate;
    }

    /**
     * @notice Check if user can pay for gas
     */
    function canPayForGas(address user, uint256 gasAmount) public view override returns (bool) {
        // Check daily limit
        uint256 currentDay = _getCurrentDay();
        uint256 todayUsage = lastUsageDay[user] == currentDay ? dailyUsage[user] : 0;
        if (todayUsage + gasAmount > getDailyLimit(user)) {
            return false;
        }

        // Check CPOP balance
        uint256 cpopCost = estimateCPOPCost(gasAmount);
        return cpopToken.balanceOf(user) >= cpopCost;
    }

    /**
     * @notice Get CPOP token address
     */
    function getCPOPToken() external view override returns (address) {
        return address(cpopToken);
    }

    /**
     * @notice Validate paymaster user operation
     */
    function _validatePaymasterUserOp(
        PackedUserOperation calldata userOp,
        bytes32 userOpHash,
        uint256 maxCost
    ) internal override whenNotPaused returns (bytes memory context, uint256 validationData) {
        (userOpHash); // suppress unused parameter warning
        
        address sender = userOp.sender;
        
        // Check if user can pay for gas
        if (!canPayForGas(sender, maxCost)) {
            return ("", 1); // validation failed
        }

        // Calculate CPOP cost
        uint256 cpopCost = estimateCPOPCost(maxCost);
        
        // Burn CPOP tokens immediately
        cpopToken.burn(sender, cpopCost);
        
        // Update daily usage
        _updateDailyUsage(sender, maxCost);
        
        // Return context with user address and gas cost
        context = abi.encode(sender, maxCost, cpopCost);
        return (context, 0); // validation success
    }

    /**
     * @notice Post operation handler
     */
    function _postOp(
        PostOpMode mode,
        bytes calldata context,
        uint256 actualGasCost,
        uint256 actualUserOpFeePerGas
    ) internal override {
        (actualUserOpFeePerGas); // suppress unused parameter warning
        
        (address sender, uint256 maxCost, uint256 cpopCost) = abi.decode(
            context, 
            (address, uint256, uint256)
        );

        emit CPOPGasPayment(sender, cpopCost, actualGasCost);

        // If operation failed, we could implement refund logic here
        // For v1, we keep it simple and don't refund
        if (mode == PostOpMode.postOpReverted) {
            // Handle post-op revert if needed
        }
    }

    /**
     * @notice Update daily usage for a user
     */
    function _updateDailyUsage(address user, uint256 gasAmount) internal {
        uint256 currentDay = _getCurrentDay();
        
        if (lastUsageDay[user] != currentDay) {
            dailyUsage[user] = gasAmount;
            lastUsageDay[user] = currentDay;
        } else {
            dailyUsage[user] += gasAmount;
        }
    }

    /**
     * @notice Get current day number
     */
    function _getCurrentDay() internal view returns (uint256) {
        return block.timestamp / 1 days;
    }

    /**
     * @notice Pause the paymaster
     */
    function pause() external onlyOwner {
        _pause();
    }

    /**
     * @notice Unpause the paymaster
     */
    function unpause() external onlyOwner {
        _unpause();
    }
}