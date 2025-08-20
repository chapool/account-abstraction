// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title IWalletManager v2 - Simplified EOA + Master Interface
 * @notice Simplified interface for Wallet Manager using EOA + Master signer pattern
 * @dev Creates deterministic wallet addresses using owner EOA + master signer combination
 */
interface IWalletManager {
    /**
     * @notice Emitted when a new AA wallet is created
     * @param account The address of the created AA wallet
     * @param owner The EOA owner address  
     * @param masterSigner The master signer address
     */
    event AccountCreated(address indexed account, address indexed owner, address indexed masterSigner);

    /**
     * @notice Emitted when an address is authorized as account creator
     * @param creator The authorized creator address
     */
    event CreatorAuthorized(address indexed creator);

    /**
     * @notice Emitted when an address is removed from authorized creators
     * @param creator The removed creator address
     */
    event CreatorRevoked(address indexed creator);

    // ============================================
    // CORE WALLET CREATION FUNCTIONS
    // ============================================

    /**
     * @notice Core wallet creation function
     * @param owner The wallet owner EOA address
     * @param masterSigner Master signer for aggregation (use address(0) for default)
     * @return account The created wallet address
     */
    function createWallet(address owner, address masterSigner) external returns (address account);

    /**
     * @notice Create account for Web2 users (simplified)
     * @param owner User's EOA address
     * @param masterSigner Master signer for transaction control
     * @return account The created AA wallet address
     */
    function createUserAccount(
        address owner,
        address masterSigner
    ) external returns (address account);

    /**
     * @notice Legacy function for EntryPoint compatibility
     * @param owner Wallet owner address
     * @param masterSigner Master signer address
     * @return account The created account address
     */
    function createAccountWithMasterSigner(
        address owner, 
        address masterSigner
    ) external returns (address account);

    // ============================================
    // ADDRESS PREDICTION FUNCTIONS
    // ============================================

    /**
     * @notice Get deterministic AA wallet address
     * @param owner User's EOA address
     * @param masterSigner Master signer address  
     * @return account Predicted AA wallet address
     */
    function getAccountAddress(address owner, address masterSigner) external view returns (address account);

    // ============================================
    // ENTRYPOINT INTEGRATION
    // ============================================

    /**
     * @notice Generate initCode for EntryPoint deployment
     * @param owner User's EOA address
     * @param masterSigner Master signer address
     * @return initCode Bytes for EntryPoint to deploy the account
     */
    function getInitCode(address owner, address masterSigner) external view returns (bytes memory initCode);

    /**
     * @notice Check if account is already deployed
     * @param owner User's EOA address
     * @param masterSigner Master signer address
     * @return isDeployed True if account is already deployed
     */
    function isAccountDeployed(address owner, address masterSigner) external view returns (bool isDeployed);

    // ============================================
    // MANAGEMENT FUNCTIONS
    // ============================================

    /**
     * @notice Initialize the wallet manager (for upgradeable contracts)
     * @param entryPoint The EntryPoint contract address
     * @param owner The owner of this contract
     */
    function initialize(address entryPoint, address owner) external;

    /**
     * @notice Get the account implementation address
     * @return The address of the AAWallet implementation
     */
    function getImplementation() external view returns (address);

    /**
     * @notice Update the account implementation
     * @param newImplementation New AAWallet implementation address
     */
    function updateAccountImplementation(address newImplementation) external;

    /**
     * @notice Set MasterAggregator address
     * @param aggregator MasterAggregator contract address
     */
    function setMasterAggregator(address aggregator) external;

    /**
     * @notice Set default master signer
     * @param masterSigner Default master signer address
     */
    function setDefaultMasterSigner(address masterSigner) external;

    /**
     * @notice Get default master signer
     * @return Default master signer address
     */
    function getDefaultMasterSigner() external view returns (address);

    /**
     * @notice Authorize an address to create accounts
     * @param creator The address to authorize
     */
    function authorizeCreator(address creator) external;

    /**
     * @notice Revoke authorization from an address
     * @param creator The address to revoke authorization from
     */
    function revokeCreator(address creator) external;

    /**
     * @notice Check if an address is authorized to create accounts
     * @param creator The address to check
     * @return True if authorized
     */
    function isAuthorizedCreator(address creator) external view returns (bool);
}