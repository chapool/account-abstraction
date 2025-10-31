// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "../../interfaces/IAggregator.sol";

/**
 * @title IMasterAggregator
 * @notice Interface for Master Aggregator with session key support
 * @dev Extends IAggregator with master signer and session key functionality
 */
interface IMasterAggregator is IAggregator {
    
    // Events
    
    /**
     * @notice Emitted when a master signer authorization status changes
     * @param master Master signer address
     * @param authorized Whether the master signer is authorized
     */
    event MasterAuthorized(address indexed master, bool authorized);

    /**
     * @notice Emitted when a Account authorization status changes for a master signer
     * @param master Master signer address
     * @param Account Account address
     * @param authorized Whether the Account is authorized
     */
    event AccountAuthorized(address indexed master, address indexed Account, bool authorized);

    /**
     * @notice Emitted when aggregated validation is performed
     * @param master Master signer address
     * @param operationCount Number of operations in the aggregation
     * @param aggregatedHash Hash of the aggregated operations
     */
    event AggregatedValidation(address indexed master, uint256 operationCount, bytes32 aggregatedHash);

    // Master Signer Management

    /**
     * @notice Check if a master signer is authorized
     * @param master Master signer address
     * @return authorized True if the master signer is authorized
     */
    function authorizedMasters(address master) external view returns (bool authorized);

    /**
     * @notice Check if a Account is controlled by a master signer
     * @param master Master signer address
     * @param Account Account address
     * @return authorized True if the Account is authorized for the master
     */
    function masterToAccounts(address master, address Account) external view returns (bool authorized);

    /**
     * @notice Get the current nonce for a master signer
     * @param master Master signer address
     * @return nonce Current nonce value
     */
    function masterNonces(address master) external view returns (uint256 nonce);

    /**
     * @notice Set master signer authorization status
     * @param master Master signer address
     * @param authorized Whether to authorize or deauthorize
     */
    function setMasterAuthorization(address master, bool authorized) external;

    /**
     * @notice Set Account authorization for a master signer
     * @param master Master signer address
     * @param Account Account address
     * @param authorized Whether to authorize or deauthorize
     */
    function setAccountAuthorization(address master, address Account, bool authorized) external;

    /**
     * @notice Batch set Account authorization for a master signer
     * @param master Master signer address
     * @param Accounts Array of Account addresses
     * @param authorized Whether to authorize or deauthorize
     */
    function batchSetAccountAuthorization(
        address master,
        address[] calldata Accounts,
        bool authorized
    ) external;

    /**
     * @notice Auto-authorize Account by verifying master signer relationship
     * @param Account Account address
     * @param master Master signer address
     */
    function autoAuthorizeAccount(address Account, address master) external;

    // Configuration

    /**
     * @notice Get maximum number of operations per aggregation
     * @return maxOps Maximum operations allowed
     */
    function maxAggregatedOps() external view returns (uint256 maxOps);

    /**
     * @notice Get validation window duration
     * @return window Validation window in seconds
     */
    function validationWindow() external view returns (uint256 window);

    /**
     * @notice Update aggregation configuration
     * @param _maxAggregatedOps Maximum operations per aggregation
     * @param _validationWindow Validation window in seconds
     */
    function updateConfig(uint256 _maxAggregatedOps, uint256 _validationWindow) external;

    // Account Validation

    /**
     * @notice Check if Account is controlled by master signer
     * @param Account Account address
     * @param master Master signer address
     * @return isValid True if Account is controlled by master
     */
    function isAccountControlledByMaster(address Account, address master) external view returns (bool isValid);

    /**
     * @notice Get current nonce for master signer
     * @param master Master signer address
     * @return nonce Current nonce
     */
    function getMasterNonce(address master) external view returns (uint256 nonce);

    // Gas Optimization

    /**
     * @notice Calculate gas savings from aggregation
     * @param operationCount Number of operations
     * @return savings Estimated gas savings
     */
    function calculateGasSavings(uint256 operationCount) external pure returns (uint256 savings);

    // Session Key Support


    // Master Control Signature Creation (Secure)

    /**
     * @notice Create master aggregated signature for multiple Accounts
     * @dev Master signer can control multiple Accounts with one signature
     * @param userOps Array of user operations to aggregate  
     * @param masterSigner Master signer address
     * @param masterSignature Pre-computed signature from master signer
     * @return aggregatedSignature Encoded aggregated signature for EntryPoint
     */
    function createMasterAggregatedSignature(
        PackedUserOperation[] calldata userOps,
        address masterSigner,
        bytes calldata masterSignature
    ) external returns (bytes memory aggregatedSignature);

    /**
     * @notice Get data that master needs to sign (call this off-chain)
     * @param userOps Array of user operations to aggregate
     * @param masterSigner Master signer address  
     * @return hashToSign The hash that master should sign off-chain
     * @return nonce The nonce to use
     */
    function getMasterSigningData(
        PackedUserOperation[] calldata userOps,
        address masterSigner
    ) external view returns (bytes32 hashToSign, uint256 nonce);
}