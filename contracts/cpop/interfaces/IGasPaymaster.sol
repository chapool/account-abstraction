// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "../../interfaces/IPaymaster.sol" as BaseIPaymaster;

/**
 * @title IGasPaymaster
 * @notice Interface for GasPaymaster - enables paying gas fees with ERC20 tokens
 * @dev Extends IPaymaster with token-specific functionality and Oracle integration
 */
interface IGasPaymaster is BaseIPaymaster.IPaymaster {
    /**
     * @notice Emitted when tokens are used to pay for gas
     * @param user The user whose gas was paid
     * @param amount The amount of tokens consumed
     * @param gasAmount The amount of gas covered
     */
    event GasPayment(address indexed user, uint256 amount, uint256 gasAmount);

    /**
     * @notice Emitted when the Oracle address is updated
     * @param oldOracle The previous Oracle address
     * @param newOracle The new Oracle address
     */
    event OracleUpdated(address indexed oldOracle, address indexed newOracle);

    /**
     * @notice Emitted when the daily limit is updated
     * @param user The user whose limit was updated
     * @param newLimit The new daily limit
     */
    event DailyLimitUpdated(address indexed user, uint256 newLimit);

    /**
     * @notice Emitted when fallback exchange rate is used
     * @param user The user for whom fallback was used
     * @param fallbackRate The fallback rate used
     * @param reason The reason for using fallback
     */
    event FallbackRateUsed(address indexed user, uint256 fallbackRate, string reason);

    /**
     * @notice Get the current Oracle address
     * @return The Oracle contract address
     */
    function getOracle() external view returns (address);

    /**
     * @notice Set the Oracle address for price feeds
     * @dev Only callable by owner
     * @param newOracle The new Oracle contract address
     */
    function setOracle(address newOracle) external;

    /**
     * @notice Get the fallback exchange rate (tokens per wei)
     * @return The fallback exchange rate
     */
    function getFallbackExchangeRate() external view returns (uint256);

    /**
     * @notice Set the fallback exchange rate for token to ETH conversion
     * @dev Only callable by owner, used when Oracle is unavailable
     * @param newRate The new fallback exchange rate (tokens per wei)
     */
    function setFallbackExchangeRate(uint256 newRate) external;

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
     * @notice Estimate the token cost for a given gas amount
     * @param gasAmount The gas amount in wei
     * @return The estimated token cost
     */
    function estimateCost(uint256 gasAmount) external view returns (uint256);

    /**
     * @notice Check if a user can pay for a given gas amount
     * @param user The user address
     * @param gasAmount The gas amount in wei
     * @return True if the user can pay
     */
    function canPayForGas(address user, uint256 gasAmount) external view returns (bool);

    /**
     * @notice Get the token address
     * @return The token contract address
     */
    function getToken() external view returns (address);

    /**
     * @notice Emitted when token handling mode is updated
     * @param burnTokens Whether to burn tokens or transfer to beneficiary
     * @param beneficiary Address to receive tokens when not burning
     */
    event TokenHandlingModeUpdated(bool burnTokens, address beneficiary);

    /**
     * @notice Emitted when beneficiary address is updated
     * @param oldBeneficiary Previous beneficiary address
     * @param newBeneficiary New beneficiary address
     */
    event BeneficiaryUpdated(address indexed oldBeneficiary, address indexed newBeneficiary);

    /**
     * @notice Emitted when tokens are burned for gas payment
     * @param user The user whose tokens were burned
     * @param amount The amount of tokens burned
     */
    event TokensBurned(address indexed user, uint256 amount);

    /**
     * @notice Emitted when tokens are transferred to beneficiary for gas payment
     * @param user The user whose tokens were transferred
     * @param beneficiary The beneficiary who received the tokens
     * @param amount The amount of tokens transferred
     */
    event TokensTransferred(address indexed user, address indexed beneficiary, uint256 amount);

    /**
     * @notice Get the current token handling mode
     * @return burnTokens True if burning tokens, false if transferring to beneficiary
     * @return beneficiary Address receiving tokens when not burning
     */
    function getTokenHandlingMode() external view returns (bool burnTokens, address beneficiary);

    /**
     * @notice Update token handling mode
     * @param _burnTokens If true, burn tokens; if false, transfer to beneficiary
     * @param _beneficiary Address to receive tokens when not burning
     */
    function setTokenHandlingMode(bool _burnTokens, address _beneficiary) external;

    /**
     * @notice Update beneficiary address
     * @param _newBeneficiary New beneficiary address
     */
    function setBeneficiary(address _newBeneficiary) external;
}