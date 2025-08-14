// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title IWalletManager
 * @notice Interface for  Wallet Manager - creates AA wallets for Web2 users
 * @dev Factory pattern for creating AAWallet instances with deterministic addresses
 */
interface IWalletManager {
    /**
     * @notice Emitted when a new  account is created
     * @param account The address of the created account
     * @param owner The owner of the account
     * @param salt The salt used for creation
     */
    event AccountCreated(address indexed account, address indexed owner, bytes32 salt);

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

    /**
     * @notice Create a new  account
     * @dev Uses CREATE2 for deterministic address generation
     * @param owner The owner of the new account
     * @param salt A unique salt for address generation
     * @return account The address of the created account
     */
    function createAccount(address owner, bytes32 salt) external returns (address account);

    /**
     * @notice Create a  account using a string identifier
     * @dev Converts string to bytes32 salt for Web2 user convenience
     * @param owner The owner of the new account
     * @param identifier A string identifier (email, phone, etc.)
     * @return account The address of the created account
     */
    function createAccountWithIdentifier(address owner, string calldata identifier) external returns (address account);

    /**
     * @notice Create a  account for Web2 users with master signer
     * @dev Uses master signer for signing on behalf of Web2 users  
     * @param identifier A string identifier (email, phone, etc.)
     * @param masterSigner The master signer address for Web2 users
     * @return account The address of the created account
     * @return generatedOwner The generated owner address  
     */
    function createWeb2AccountWithMasterSigner(
        string calldata identifier, 
        address masterSigner
    ) external returns (address account, address generatedOwner);

    /**
     * @notice Get the deterministic address for an account
     * @dev Returns the address without creating the account
     * @param owner The owner of the account
     * @param salt The salt for address generation
     * @return account The deterministic address
     */
    function getAccountAddress(address owner, bytes32 salt) external view returns (address account);

    /**
     * @notice Get the deterministic address using a string identifier
     * @dev Converts string to bytes32 salt and returns deterministic address
     * @param owner The owner of the account
     * @param identifier A string identifier
     * @return account The deterministic address
     */
    function getAccountAddressWithIdentifier(address owner, string calldata identifier) external view returns (address account);

    /**
     * @notice Get the account implementation address
     * @return The address of the AAWallet implementation
     */
    function getImplementation() external view returns (address);

    /**
     * @notice Create account with master signer via EntryPoint (for hybrid deployment)
     * @param generatedOwner The generated owner address
     * @param salt The salt for deterministic deployment
     * @param masterSigner The master signer address
     * @return account The deployed account address
     */
    function createAccountWithMasterSigner(
        address generatedOwner, 
        bytes32 salt, 
        address masterSigner
    ) external returns (address account);

    /**
     * @notice Get Web2 account address using identifier and master signer
     * @param identifier User identifier string
     * @param masterSigner Master signer address
     * @return account The predicted account address
     */
    function getWeb2AccountAddress(string calldata identifier, address masterSigner) 
        external view returns (address account);

    /**
     * @notice Generate initCode for Web2 users to enable EntryPoint deployment
     * @param identifier User identifier string
     * @param masterSigner Master signer address for Web2 user
     * @return initCode Bytes for EntryPoint to deploy the account
     */
    function getWeb2InitCode(string calldata identifier, address masterSigner) 
        external view returns (bytes memory initCode);

    /**
     * @notice Check if Web2 account is already deployed
     * @param identifier User identifier string
     * @param masterSigner Master signer address
     * @return isDeployed True if account is already deployed
     */
    function isWeb2AccountDeployed(string calldata identifier, address masterSigner) 
        external view returns (bool isDeployed);

    /**
     * @notice Convert string identifier to bytes32 salt
     * @param identifier The string identifier
     * @return salt The generated salt
     */
    function identifierToSalt(string calldata identifier) external pure returns (bytes32 salt);

    /**
     * @notice Generate deterministic owner address from master signer and identifier
     * @param masterSigner The master signer address
     * @param identifier The string identifier
     * @return owner The generated owner address
     */
    function generateOwnerFromMasterSigner(address masterSigner, string calldata identifier) external pure returns (address owner);

    /**
     * @notice Initialize the wallet manager (for upgradeable contracts)
     * @param entryPoint The EntryPoint contract address
     * @param Token The  token contract address  
     * @param owner The owner of this contract
     */
    function initialize(address entryPoint, address Token, address owner) external;

    /**
     * @notice Update the account implementation
     * @param newImplementation New AAWallet implementation address
     */
    function updateAccountImplementation(address newImplementation) external;

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