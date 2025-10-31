// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title ISessionKeyManager
 * @notice Interface for centralized session key management across multiple CPOP accounts
 * @dev Allows batch operations and template-based session key management
 */
interface ISessionKeyManager {
    
    // Structs
    
    /**
     * @notice Session key template for common use cases
     * @param name Template name
     * @param defaultDuration Default session duration in seconds
     * @param permissions Default permissions
     * @param isActive Whether the template is active
     */
    struct SessionKeyTemplate {
        string name;
        uint48 defaultDuration;
        bytes32 permissions;
        bool isActive;
    }

    /**
     * @notice Batch session key operation
     * @param account Target account address
     * @param sessionKey Session key address
     * @param validAfter Start time for the session key
     * @param validUntil End time for the session key
     * @param permissions Encoded permissions
     */
    struct BatchSessionKeyOp {
        address account;
        address sessionKey;
        uint48 validAfter;
        uint48 validUntil;
        bytes32 permissions;
    }

    // Events
    
    /**
     * @notice Emitted when a master signer authorization status changes
     * @param masterSigner Master signer address
     * @param authorized Whether the master signer is authorized
     */
    event MasterSignerAuthorized(address indexed masterSigner, bool authorized);

    /**
     * @notice Emitted when an account is registered under a master signer
     * @param masterSigner Master signer address
     * @param account Account address
     */
    event AccountRegistered(address indexed masterSigner, address indexed account);

    /**
     * @notice Emitted when an account is unregistered from a master signer
     * @param masterSigner Master signer address
     * @param account Account address
     */
    event AccountUnregistered(address indexed masterSigner, address indexed account);

    /**
     * @notice Emitted when a session key template is created
     * @param name Template name
     * @param duration Default duration
     * @param permissions Default permissions
     */
    event SessionKeyTemplateCreated(string indexed name, uint48 duration, bytes32 permissions);

    /**
     * @notice Emitted when session keys are added in batch
     * @param masterSigner Master signer address
     * @param operationCount Number of operations performed
     */
    event BatchSessionKeysAdded(address indexed masterSigner, uint256 operationCount);

    /**
     * @notice Emitted when session keys are revoked in batch
     * @param masterSigner Master signer address
     * @param operationCount Number of operations performed
     */
    event BatchSessionKeysRevoked(address indexed masterSigner, uint256 operationCount);

    // Master Signer Management

    /**
     * @notice Authorize a master signer
     * @param masterSigner Master signer address to authorize
     */
    function authorizeMasterSigner(address masterSigner) external;

    /**
     * @notice Revoke master signer authorization
     * @param masterSigner Master signer address to revoke
     */
    function revokeMasterSigner(address masterSigner) external;

    /**
     * @notice Check if a master signer is authorized
     * @param masterSigner Master signer address to check
     * @return authorized True if the master signer is authorized
     */
    function authorizedMasters(address masterSigner) external view returns (bool authorized);

    // Account Management

    /**
     * @notice Register an account under a master signer
     * @param masterSigner Master signer address
     * @param account Account address to register
     */
    function registerAccount(address masterSigner, address account) external;

    /**
     * @notice Unregister an account from a master signer
     * @param masterSigner Master signer address
     * @param account Account address to unregister
     */
    function unregisterAccount(address masterSigner, address account) external;

    /**
     * @notice Get accounts registered under a master signer
     * @param masterSigner Master signer address
     * @return accounts Array of registered account addresses
     */
    function getRegisteredAccounts(address masterSigner) external view returns (address[] memory accounts);

    // Template Management

    /**
     * @notice Create a session key template
     * @param name Template name
     * @param defaultDuration Default duration in seconds
     * @param permissions Default permissions
     */
    function createSessionKeyTemplate(
        string calldata name,
        uint48 defaultDuration,
        bytes32 permissions
    ) external;

    /**
     * @notice Get template information
     * @param templateName Template name
     * @return template Template data
     */
    function getTemplate(string calldata templateName) external view returns (SessionKeyTemplate memory template);

    /**
     * @notice Get all template names
     * @return names Array of template names
     */
    function getTemplateNames() external view returns (string[] memory names);


    // Session Key Operations

    /**
     * @notice Add session keys to multiple accounts using a template
     * @param masterSigner Master signer (must be authorized)
     * @param sessionKey Session key address
     * @param templateName Template to use
     * @param customDuration Custom duration (0 to use template default)
     */
    function addSessionKeyWithTemplate(
        address masterSigner,
        address sessionKey,
        string calldata templateName,
        uint48 customDuration
    ) external;

    /**
     * @notice Batch add session keys to multiple accounts
     * @param operations Array of batch operations
     */
    function batchAddSessionKeys(BatchSessionKeyOp[] calldata operations) external;

    /**
     * @notice Batch revoke session keys from multiple accounts
     * @param sessionKey Session key to revoke
     */
    function batchRevokeSessionKey(address sessionKey) external;


    // View Functions

    /**
     * @notice Get master signer accounts mapping
     * @param masterSigner Master signer address
     * @param index Index in the accounts array
     * @return account Account address at the given index
     */
    function masterSignerAccounts(address masterSigner, uint256 index) external view returns (address account);
}