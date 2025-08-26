// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "../../interfaces/IAccount.sol";
import "@openzeppelin/contracts/utils/introspection/IERC165.sol";

/**
 * @title IAAccount  
 * @notice Interface for AA Wallet (Account Abstraction wallet)
 * @dev Extends IAccount with Web2-friendly functionality and master signer system
 */
interface IAAccount is IAccount, IERC165 {
    /**
     * @notice Emitted when the account is initialized
     * @param owner The owner of the account
     * @param masterSigner The master signer address for Web2 users
     */
    event AAccountInitialized(address indexed owner, address indexed masterSigner);

    /**
     * @notice Emitted when master signer is updated
     * @param oldMasterSigner The old master signer address
     * @param newMasterSigner The new master signer address
     */
    event MasterSignerUpdated(address indexed oldMasterSigner, address indexed newMasterSigner);

    /**
     * @notice Emitted when master signer validation is used
     * @param masterSigner The master signer that validated the operation
     * @param userOpHash The hash of the user operation
     */
    event MasterSignerValidation(address indexed masterSigner, bytes32 userOpHash);

    /**
     * @notice Get the owner of this account
     * @return The owner address
     */
    function getOwner() external view returns (address);

    /**
     * @notice Get the master signer of this account
     * @return The master signer address
     */
    function getMasterSigner() external view returns (address);

    /**
     * @notice Set a new master signer (only current master signer can call)
     * @param newMasterSigner The new master signer address
     */
    function setMasterSigner(address newMasterSigner) external;

    /**
     * @notice Check if master signer validation is enabled
     * @return True if master signer validation is enabled
     */
    function isMasterSignerEnabled() external view returns (bool);

    // Session Keys Events
    /**
     * @notice Emitted when a session key is added
     * @param sessionKey The session key address
     * @param validAfter Start time for the session key
     * @param validUntil End time for the session key
     * @param permissions Encoded permissions for the session key
     */
    event SessionKeyAdded(
        address indexed sessionKey, 
        uint48 validAfter, 
        uint48 validUntil, 
        bytes32 permissions
    );

    /**
     * @notice Emitted when a session key is revoked
     * @param sessionKey The session key address that was revoked
     */
    event SessionKeyRevoked(address indexed sessionKey);

    /**
     * @notice Emitted when a session key is used for validation
     * @param sessionKey The session key that validated the operation
     * @param userOpHash The hash of the user operation
     */
    event SessionKeyValidation(address indexed sessionKey, bytes32 userOpHash);


    // Session Keys Functions
    /**
     * @notice Add a session key with time and permission constraints
     * @param sessionKey The session key address
     * @param validAfter Start time for the session key (0 for immediate)
     * @param validUntil End time for the session key
     * @param permissions Encoded permissions (targets, methods, etc.)
     */
    function addSessionKey(
        address sessionKey,
        uint48 validAfter,
        uint48 validUntil,
        bytes32 permissions
    ) external;

    /**
     * @notice Revoke a session key
     * @param sessionKey The session key address to revoke
     */
    function revokeSessionKey(address sessionKey) external;

    /**
     * @notice Check if a session key is valid
     * @param sessionKey The session key address to check
     * @return isValid True if the session key is currently valid
     * @return validAfter Start time of validity
     * @return validUntil End time of validity
     * @return permissions Encoded permissions
     */
    function getSessionKeyInfo(address sessionKey) 
        external 
        view 
        returns (bool isValid, uint48 validAfter, uint48 validUntil, bytes32 permissions);

    /**
     * @notice Check if a session key can execute a specific operation
     * @param sessionKey The session key address
     * @param target The target contract address
     * @param selector The function selector
     * @return canExecute True if the session key can execute this operation
     */
    function canSessionKeyExecute(
        address sessionKey,
        address target,
        bytes4 selector
    ) external view returns (bool canExecute);

    // Aggregator Functions
    /**
     * @notice Set aggregator address for signature aggregation
     * @param _aggregator The aggregator contract address
     */
    function setAggregator(address _aggregator) external;

    /**
     * @notice Get the current aggregator address
     * @return The aggregator contract address
     */
    function getAggregator() external view returns (address);

}