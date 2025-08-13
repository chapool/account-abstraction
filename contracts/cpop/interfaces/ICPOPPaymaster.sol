// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "../../interfaces/IPaymaster.sol";

/**
 * @title ICPOPPaymaster
 * @notice Interface for CPOP Paymaster - enables paying gas fees with CPOP tokens
 * @dev Extends IPaymaster with CPOP-specific functionality
 */
interface ICPOPPaymaster is IPaymaster {
    /**
     * @notice Emitted when CPOP tokens are used to pay for gas
     * @param user The user whose gas was paid
     * @param cpopAmount The amount of CPOP tokens consumed
     * @param gasAmount The amount of gas covered
     */
    event CPOPGasPayment(address indexed user, uint256 cpopAmount, uint256 gasAmount);

    /**
     * @notice Emitted when the exchange rate is updated
     * @param oldRate The previous exchange rate
     * @param newRate The new exchange rate
     */
    event ExchangeRateUpdated(uint256 oldRate, uint256 newRate);

    /**
     * @notice Emitted when the daily limit is updated
     * @param user The user whose limit was updated
     * @param newLimit The new daily limit
     */
    event DailyLimitUpdated(address indexed user, uint256 newLimit);

    /**
     * @notice Get the current exchange rate (CPOP per wei)
     * @return The exchange rate
     */
    function getExchangeRate() external view returns (uint256);

    /**
     * @notice Set the exchange rate for CPOP to ETH conversion
     * @dev Only callable by owner
     * @param newRate The new exchange rate (CPOP per wei)
     */
    function setExchangeRate(uint256 newRate) external;

    /**
     * @notice Get the daily gas limit for a user
     * @param user The user address
     * @return The daily limit in wei
     */
    function getDailyLimit(address user) external view returns (uint256);

    /**
     * @notice Set the daily gas limit for a user
     * @dev Only callable by owner
     * @param user The user address
     * @param limit The daily limit in wei
     */
    function setDailyLimit(address user, uint256 limit) external;

    /**
     * @notice Get the amount of gas used by a user today
     * @param user The user address
     * @return The gas used today in wei
     */
    function getDailyUsage(address user) external view returns (uint256);

    /**
     * @notice Estimate the CPOP cost for a given gas amount
     * @param gasAmount The gas amount in wei
     * @return The estimated CPOP cost
     */
    function estimateCPOPCost(uint256 gasAmount) external view returns (uint256);

    /**
     * @notice Check if a user can pay for a given gas amount
     * @param user The user address
     * @param gasAmount The gas amount in wei
     * @return True if the user can pay
     */
    function canPayForGas(address user, uint256 gasAmount) external view returns (bool);

    /**
     * @notice Get the CPOP token address
     * @return The CPOP token contract address
     */
    function getCPOPToken() external view returns (address);
}