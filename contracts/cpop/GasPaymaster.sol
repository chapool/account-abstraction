// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/utils/Pausable.sol";
import "../core/BasePaymaster.sol";
import "./interfaces/IGasPaymaster.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces/IGasPriceOracle.sol";

/**
 * @notice Interface for burnable ERC20 tokens
 */
interface IERC20Burnable is IERC20 {
    function burnFrom(address account, uint256 amount) external;
}

/**
 * @title GasPaymaster
 * @notice Enhanced Gas Paymaster with Oracle integration for dynamic Token/ETH pricing
 * @dev Inherits from BasePaymaster and implements IGasPaymaster interface
 * Uses Oracle for real-time price feeds with fallback mechanisms
 */
contract GasPaymaster is BasePaymaster, IGasPaymaster, Pausable {
    IERC20Burnable public immutable token;
    IGasPriceOracle public oracle;
    
    uint256 public fallbackExchangeRate; // Fallback Token per wei
    uint256 public constant DEFAULT_DAILY_LIMIT = 0.01 ether; // 0.01 ETH worth of gas per day
    uint256 public constant ORACLE_PRICE_TOLERANCE = 3600; // 1 hour tolerance for Oracle prices
    
    // Token handling configuration
    bool public burnTokens; // If true, burn tokens; if false, transfer to beneficiary
    address public beneficiary; // Address to receive tokens when not burning
    
    mapping(address => uint256) private dailyLimits;
    mapping(address => uint256) private dailyUsage;
    mapping(address => uint256) private lastUsageDay;

    // Events are defined in IGasPaymaster interface

    modifier validOracle() {
        require(address(oracle) != address(0), "GasPaymaster: Oracle not set");
        _;
    }

    constructor(
        IEntryPoint _entryPoint,
        address _token,
        address _oracle,
        uint256 _fallbackExchangeRate,
        bool _burnTokens,
        address _beneficiary
    ) BasePaymaster(_entryPoint) {
        require(_token != address(0), "GasPaymaster: invalid token");
        require(_fallbackExchangeRate > 0, "GasPaymaster: invalid fallback rate");
        if (!_burnTokens) {
            require(_beneficiary != address(0), "GasPaymaster: invalid beneficiary");
        }
        
        token = IERC20Burnable(_token);
        oracle = IGasPriceOracle(_oracle);
        fallbackExchangeRate = _fallbackExchangeRate;
        burnTokens = _burnTokens;
        beneficiary = _beneficiary;
    }

    /**
     * @inheritdoc IGasPaymaster
     */
    function getOracle() external view override returns (address) {
        return address(oracle);
    }

    /**
     * @inheritdoc IGasPaymaster
     */
    function setOracle(address newOracle) external override onlyOwner {
        require(newOracle != address(0), "GasPaymaster: invalid Oracle");
        address oldOracle = address(oracle);
        oracle = IGasPriceOracle(newOracle);
        emit OracleUpdated(oldOracle, newOracle);
    }

    /**
     * @inheritdoc IGasPaymaster
     */
    function getFallbackExchangeRate() external view override returns (uint256) {
        return fallbackExchangeRate;
    }

    /**
     * @inheritdoc IGasPaymaster
     */
    function setFallbackExchangeRate(uint256 newRate) external override onlyOwner {
        require(newRate > 0, "GasPaymaster: invalid fallback rate");
        fallbackExchangeRate = newRate;
    }

    /**
     * @inheritdoc IGasPaymaster
     */
    function getDailyLimit(address user) public view override returns (uint256) {
        uint256 limit = dailyLimits[user];
        return limit > 0 ? limit : DEFAULT_DAILY_LIMIT;
    }

    /**
     * @inheritdoc IGasPaymaster
     */
    function setDailyLimit(address user, uint256 limit) external override onlyOwner {
        dailyLimits[user] = limit;
        emit DailyLimitUpdated(user, limit);
    }

    /**
     * @inheritdoc IGasPaymaster
     */
    function getDailyUsage(address user) external view override returns (uint256) {
        if (_getCurrentDay() != lastUsageDay[user]) {
            return 0;
        }
        return dailyUsage[user];
    }

    /**
     * @inheritdoc IGasPaymaster
     */
    function estimateCost(uint256 gasAmount) public view override returns (uint256) {
        return _getExchangeRateWithFallback() * gasAmount;
    }

    /**
     * @notice Get exchange rate from Oracle with fallback mechanism
     * @return rate Exchange rate (Token per wei)
     */
    function _getExchangeRateWithFallback() internal view returns (uint256 rate) {
        if (address(oracle) == address(0)) {
            return fallbackExchangeRate;
        }

        try oracle.getTokenForETH(1 ether) returns (uint256 tokenForOneETH) {
            // Convert to Token per wei
            rate = tokenForOneETH / 1 ether;
            
            // Sanity check: rate should be reasonable
            if (rate == 0 || rate > fallbackExchangeRate * 10) {
                return fallbackExchangeRate;
            }
            
            return rate;
        } catch {
            return fallbackExchangeRate;
        }
    }

    /**
     * @notice Get exchange rate from Oracle with metadata
     * @return rate Exchange rate (Token per wei)
     * @return isFromOracle True if rate is from Oracle, false if fallback
     */
    function _getExchangeRateWithMetadata() internal view returns (uint256 rate, bool isFromOracle) {
        if (address(oracle) == address(0)) {
            return (fallbackExchangeRate, false);
        }

        try oracle.getTokenForETH(1 ether) returns (uint256 tokenForOneETH) {
            rate = tokenForOneETH / 1 ether;
            
            // Sanity check
            if (rate == 0 || rate > fallbackExchangeRate * 10) {
                return (fallbackExchangeRate, false);
            }
            
            return (rate, true);
        } catch {
            return (fallbackExchangeRate, false);
        }
    }

    /**
     * @inheritdoc IGasPaymaster
     */
    function canPayForGas(address user, uint256 gasAmount) public view override returns (bool) {
        // Check daily limit
        uint256 currentDay = _getCurrentDay();
        uint256 todayUsage = lastUsageDay[user] == currentDay ? dailyUsage[user] : 0;
        if (todayUsage + gasAmount > getDailyLimit(user)) {
            return false;
        }

        // Check Token balance
        uint256 tokenCost = estimateCost(gasAmount);
        return token.balanceOf(user) >= tokenCost;
    }

    /**
     * @inheritdoc IGasPaymaster
     */
    function getToken() external view override returns (address) {
        return address(token);
    }

    /**
     * @notice Validate paymaster user operation with Oracle integration
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

        // Get exchange rate with Oracle
        (uint256 exchangeRate, bool isFromOracle) = _getExchangeRateWithMetadata();
        uint256 tokenCost = exchangeRate * maxCost;
        
        // Emit fallback event if Oracle was not used
        if (!isFromOracle) {
            emit FallbackRateUsed(sender, exchangeRate, "Oracle unavailable or invalid");
        }
        
        // Handle tokens - burn or transfer to beneficiary
        _handleTokenPayment(sender, tokenCost);
        
        // Update daily usage
        _updateDailyUsage(sender, maxCost);
        
        // Return context with user address, gas cost, token cost, and rate source
        context = abi.encode(sender, maxCost, tokenCost, isFromOracle);
        return (context, 0); // validation success
    }

    /**
     * @notice Post operation handler with Oracle rate tracking
     */
    function _postOp(
        PostOpMode mode,
        bytes calldata context,
        uint256 actualGasCost,
        uint256 actualUserOpFeePerGas
    ) internal override {
        (actualUserOpFeePerGas); // suppress unused parameter warning
        
        (address sender, /* uint256 maxCost */, uint256 tokenCost, bool isFromOracle) = abi.decode(
            context, 
            (address, uint256, uint256, bool)
        );

        emit GasPayment(sender, tokenCost, actualGasCost);

        // Handle operation failure - could implement partial refund logic
        if (mode == PostOpMode.postOpReverted) {
            // For v2, we keep it simple and don't refund
            // Future versions could implement more sophisticated refund logic
        }
        
        // Log Oracle health for monitoring
        if (!isFromOracle && address(oracle) != address(0)) {
            // Oracle was available but returned invalid data
            // This could trigger alerts in monitoring systems
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
     * @notice Handle token payment - burn or transfer to beneficiary
     * @param user The user who is paying
     * @param amount The amount of tokens to handle
     */
    function _handleTokenPayment(address user, uint256 amount) internal {
        if (burnTokens) {
            // Burn tokens from user
            token.burnFrom(user, amount);
            emit TokensBurned(user, amount);
        } else {
            // Transfer tokens from user to beneficiary
            token.transferFrom(user, beneficiary, amount);
            emit TokensTransferred(user, beneficiary, amount);
        }
    }

    /**
     * @notice Get Oracle health status
     * @return isHealthy True if Oracle is responding and providing valid data
     * @return lastUpdate Timestamp of last successful Oracle update
     */
    function getOracleHealthStatus() external view returns (bool isHealthy, uint256 lastUpdate) {
        if (address(oracle) == address(0)) {
            return (false, 0);
        }
        
        try oracle.getTokenPriceUSD() returns (uint256, uint256 timestamp) {
            isHealthy = (block.timestamp - timestamp) <= ORACLE_PRICE_TOLERANCE;
            lastUpdate = timestamp;
        } catch {
            isHealthy = false;
            lastUpdate = 0;
        }
    }

    /**
     * @notice Emergency function to estimate gas cost with detailed breakdown
     * @param gasLimit Gas limit for the operation
     * @param gasPrice Current gas price
     * @return oracleCost Cost calculated using Oracle
     * @return fallbackCost Cost calculated using fallback rate
     * @return recommendedCost Recommended cost to use
     * @return useOracle Whether Oracle should be used
     */
    function getDetailedGasCostEstimate(uint256 gasLimit, uint256 gasPrice) 
        external 
        view 
        returns (
            uint256 oracleCost,
            uint256 fallbackCost,
            uint256 recommendedCost,
            bool useOracle
        ) 
    {
        uint256 totalGasCostWei = gasLimit * gasPrice;
        
        // Always calculate fallback cost
        fallbackCost = fallbackExchangeRate * totalGasCostWei;
        
        // Try to get Oracle cost
        if (address(oracle) != address(0)) {
            try oracle.getTokenForETH(totalGasCostWei) returns (uint256 cost) {
                oracleCost = cost;
                
                // Use Oracle if cost is reasonable (within 50% of fallback)
                uint256 deviation = oracleCost > fallbackCost ? 
                    oracleCost - fallbackCost : 
                    fallbackCost - oracleCost;
                    
                if (deviation <= fallbackCost / 2) {
                    useOracle = true;
                    recommendedCost = oracleCost;
                } else {
                    useOracle = false;
                    recommendedCost = fallbackCost;
                }
            } catch {
                oracleCost = 0;
                useOracle = false;
                recommendedCost = fallbackCost;
            }
        } else {
            oracleCost = 0;
            useOracle = false;
            recommendedCost = fallbackCost;
        }
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

    /**
     * @notice Update token handling mode
     * @param _burnTokens If true, burn tokens; if false, transfer to beneficiary
     * @param _beneficiary Address to receive tokens when not burning (ignored if burning)
     */
    function setTokenHandlingMode(bool _burnTokens, address _beneficiary) external onlyOwner {
        if (!_burnTokens) {
            require(_beneficiary != address(0), "GasPaymaster: invalid beneficiary");
        }
        
        burnTokens = _burnTokens;
        beneficiary = _beneficiary;
        
        emit TokenHandlingModeUpdated(_burnTokens, _beneficiary);
    }

    /**
     * @notice Update beneficiary address
     * @param _newBeneficiary New beneficiary address
     */
    function setBeneficiary(address _newBeneficiary) external onlyOwner {
        require(_newBeneficiary != address(0), "GasPaymaster: invalid beneficiary");
        require(!burnTokens, "GasPaymaster: burning mode enabled");
        
        address oldBeneficiary = beneficiary;
        beneficiary = _newBeneficiary;
        
        emit BeneficiaryUpdated(oldBeneficiary, _newBeneficiary);
    }

    /**
     * @inheritdoc IGasPaymaster
     */
    function getTokenHandlingMode() external view override returns (bool, address) {
        return (burnTokens, beneficiary);
    }
}