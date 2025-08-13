// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "../../interfaces/IAccount.sol";
import "@openzeppelin/contracts/utils/introspection/IERC165.sol";

/**
 * @title ICPOPAccount
 * @notice Interface for CPOP Account Abstraction wallet
 * @dev Extends IAccount with Web2-friendly functionality and master signer system
 */
interface ICPOPAccount is IAccount, IERC165 {
    /**
     * @notice Emitted when the account is initialized
     * @param owner The owner of the account
     * @param masterSigner The master signer address for Web2 users
     */
    event CPOPAccountInitialized(address indexed owner, address indexed masterSigner);

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
}