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
     * @notice Emitted when a wallet authorization status changes for a master signer
     * @param master Master signer address
     * @param wallet Wallet address
     * @param authorized Whether the wallet is authorized
     */
    event WalletAuthorized(address indexed master, address indexed wallet, bool authorized);

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
     * @notice Check if a wallet is controlled by a master signer
     * @param master Master signer address
     * @param wallet Wallet address
     * @return authorized True if the wallet is authorized for the master
     */
    function masterToWallets(address master, address wallet) external view returns (bool authorized);

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
     * @notice Set wallet authorization for a master signer
     * @param master Master signer address
     * @param wallet Wallet address
     * @param authorized Whether to authorize or deauthorize
     */
    function setWalletAuthorization(address master, address wallet, bool authorized) external;

    /**
     * @notice Batch set wallet authorization for a master signer
     * @param master Master signer address
     * @param wallets Array of wallet addresses
     * @param authorized Whether to authorize or deauthorize
     */
    function batchSetWalletAuthorization(
        address master,
        address[] calldata wallets,
        bool authorized
    ) external;

    /**
     * @notice Auto-authorize wallet by verifying master signer relationship
     * @param wallet Wallet address
     * @param master Master signer address
     */
    function autoAuthorizeWallet(address wallet, address master) external;

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

    // Wallet Validation

    /**
     * @notice Check if wallet is controlled by master signer
     * @param wallet Wallet address
     * @param master Master signer address
     * @return isValid True if wallet is controlled by master
     */
    function isWalletControlledByMaster(address wallet, address master) external view returns (bool isValid);

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

    /**
     * @notice Create aggregated signature for session key operations
     * @param userOps Array of user operations using session keys
     * @param sessionKey Session key that will sign all operations
     * @param sessionKeyPrivateKey Session key private key (used off-chain)
     * @return aggregatedSignature Encoded session key signature
     */
    function createSessionKeyAggregatedSignature(
        PackedUserOperation[] calldata userOps,
        address sessionKey,
        bytes32 sessionKeyPrivateKey
    ) external view returns (bytes memory aggregatedSignature);

    /**
     * @notice Validate session key aggregated signature
     * @param userOps Array of user operations
     * @param signature Session key aggregated signature
     * @return isValid True if signature is valid
     */
    function validateSessionKeyAggregatedSignature(
        PackedUserOperation[] calldata userOps,
        bytes calldata signature
    ) external view returns (bool isValid);

    // Master Signature Creation (Off-chain Helper)

    /**
     * @notice Create aggregated signature for master signer (off-chain helper)
     * @param userOps Array of user operations to aggregate
     * @param masterSigner Master signer address
     * @param masterPrivateKey Master signer private key (used off-chain)
     * @return aggregatedSignature Encoded aggregated signature
     */
    function createMasterAggregatedSignature(
        PackedUserOperation[] calldata userOps,
        address masterSigner,
        bytes32 masterPrivateKey
    ) external view returns (bytes memory aggregatedSignature);
}