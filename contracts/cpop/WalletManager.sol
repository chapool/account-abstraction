// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "../interfaces/ISenderCreator.sol";
import "../interfaces/IEntryPoint.sol";
import "./interfaces/IWalletManager.sol";
import "./interfaces/IMasterAggregator.sol";
import "./AAWallet.sol";

/**
 * @title WalletManager
 * @notice Upgradeable factory contract for creating AAWallets
 * @dev Creates deterministic wallet addresses using CREATE2 for Web2 user experience
 */
contract WalletManager is Initializable, IWalletManager, OwnableUpgradeable, UUPSUpgradeable {
    address public accountImplementation;
    ISenderCreator public senderCreator;
    address public entryPointAddress; // Store entryPoint address for later use
    address public masterAggregatorAddress; // MasterAggregator for signature aggregation
    
    mapping(address => bool) private authorizedCreators;

    modifier onlyAuthorizedCreator() {
        require(
            authorizedCreators[msg.sender] || msg.sender == owner(),
            "WalletManager: unauthorized creator"
        );
        _;
    }

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize the wallet manager
     * @param entryPoint The EntryPoint contract address
     * @param Token The  token contract address
     * @param owner The owner of this contract
     */
    function initialize(address entryPoint, address Token, address owner) external initializer {
        require(entryPoint != address(0), "WalletManager: invalid entryPoint");
        require(Token != address(0), "WalletManager: invalid Token");
        require(owner != address(0), "WalletManager: invalid owner");
        
        // Initialize parent contracts
        __Ownable_init(owner);
        __UUPSUpgradeable_init();
        
        // Initialize contract state
        entryPointAddress = entryPoint;
        accountImplementation = address(new AAWallet());
        senderCreator = IEntryPoint(entryPoint).senderCreator();
        
        // Authorize the owner as a creator by default
        authorizedCreators[owner] = true;
        emit CreatorAuthorized(owner);
    }

    /**
     * @notice Create a new  account with deterministic address
     */
    function createAccount(address owner, bytes32 salt) 
        external 
        override 
        returns (address account) 
    {
        require(msg.sender == address(senderCreator), "WalletManager: only SenderCreator");
        account = _createAccount(owner, salt);
    }

    /**
     * @notice Create a new  account with master signer via EntryPoint
     * @dev This function can be called by EntryPoint for Web2 users via initCode
     */
    function createAccountWithMasterSigner(
        address generatedOwner, 
        bytes32 salt, 
        address masterSigner
    ) external returns (address account) {
        require(msg.sender == address(senderCreator), "WalletManager: only SenderCreator");
        require(masterSigner != address(0), "WalletManager: invalid master signer");
        
        address addr = _getAccountAddressWithMasterSigner(generatedOwner, salt, masterSigner);
        uint256 codeSize = addr.code.length;
        
        if (codeSize > 0) {
            account = addr;
        } else {
            // Use initializeWithAggregator if aggregator is available
            bytes memory initData;
            if (masterAggregatorAddress != address(0) && masterAggregatorAddress.code.length > 0) {
                initData = abi.encodeCall(AAWallet.initializeWithAggregator, 
                    (entryPointAddress, generatedOwner, masterSigner, masterAggregatorAddress));
            } else {
                initData = abi.encodeCall(AAWallet.initialize, 
                    (entryPointAddress, generatedOwner, masterSigner));
            }
            
            account = address(
                new ERC1967Proxy{salt: salt}(
                    accountImplementation,
                    initData
                )
            );
            
            // Register wallet-master relationship in aggregator if configured
            if (masterAggregatorAddress != address(0) && masterAggregatorAddress.code.length > 0) {
                try IMasterAggregator(masterAggregatorAddress).autoAuthorizeWallet(masterSigner, account) {
                    // Successfully registered wallet-master relationship
                } catch {
                    // Aggregator authorization failed, wallet can still function without aggregation
                }
            }
        }
        
        emit AccountCreated(account, generatedOwner, salt);
    }

    /**
     * @notice Internal function to create account
     */
    function _createAccount(address owner, bytes32 salt) 
        internal 
        returns (address account) 
    {
        require(owner != address(0), "WalletManager: invalid owner");

        address addr = getAccountAddress(owner, salt);
        uint256 codeSize = addr.code.length;
        
        if (codeSize > 0) {
            return addr;
        }

        account = address(
            new ERC1967Proxy{salt: salt}(
                accountImplementation,
                abi.encodeCall(AAWallet.initialize, (entryPointAddress, owner, address(0)))
            )
        );

        emit AccountCreated(account, owner, salt);
    }

    /**
     * @notice Create account using string identifier (Web2 friendly)
     */
    function createAccountWithIdentifier(address owner, string calldata identifier) 
        external 
        override 
        returns (address account) 
    {
        bytes32 salt = identifierToSalt(identifier);
        return _createAccount(owner, salt);
    }


    /**
     * @notice Create account for Web2 users with master signer support
     */
    function createWeb2AccountWithMasterSigner(
        string calldata identifier, 
        address masterSigner
    ) 
        external 
        override 
        onlyAuthorizedCreator
        returns (address account, address generatedOwner) 
    {
        require(masterSigner != address(0), "WalletManager: invalid master signer");
        
        // Generate owner based on master signer + identifier for better uniqueness
        generatedOwner = generateOwnerFromMasterSigner(masterSigner, identifier);
        bytes32 salt = identifierToSalt(identifier);
        
        // Create account with master signer
        address addr = _getAccountAddressWithMasterSigner(generatedOwner, salt, masterSigner);
        uint256 codeSize = addr.code.length;
        
        if (codeSize > 0) {
            account = addr;
        } else {
            account = address(
                new ERC1967Proxy{salt: salt}(
                    accountImplementation,
                    abi.encodeCall(AAWallet.initialize, (entryPointAddress, generatedOwner, masterSigner))
                )
            );
        }
        
        emit AccountCreated(account, generatedOwner, salt);
    }


    /**
     * @notice Get deterministic account address
     */
    function getAccountAddress(address owner, bytes32 salt) 
        public 
        view 
        override 
        returns (address account) 
    {
        return _getAccountAddressWithMasterSigner(owner, salt, address(0));
    }
    
    /**
     * @notice Get deterministic account address with master signer support
     */
    function _getAccountAddressWithMasterSigner(address owner, bytes32 salt, address masterSigner) 
        internal 
        view 
        returns (address account) 
    {
        return Create2.computeAddress(
            salt,
            keccak256(
                abi.encodePacked(
                    type(ERC1967Proxy).creationCode,
                    abi.encode(
                        accountImplementation,
                        abi.encodeCall(AAWallet.initialize, (entryPointAddress, owner, masterSigner))
                    )
                )
            )
        );
    }

    /**
     * @notice Get deterministic address using string identifier
     */
    function getAccountAddressWithIdentifier(address owner, string calldata identifier) 
        external 
        view 
        override 
        returns (address account) 
    {
        bytes32 salt = identifierToSalt(identifier);
        return getAccountAddress(owner, salt);
    }

    /**
     * @notice Get Web2 account address using identifier and master signer
     */
    function getWeb2AccountAddress(string calldata identifier, address masterSigner) 
        external 
        view 
        returns (address account) 
    {
        address generatedOwner = generateOwnerFromMasterSigner(masterSigner, identifier);
        bytes32 salt = identifierToSalt(identifier);
        return _getAccountAddressWithMasterSigner(generatedOwner, salt, masterSigner);
    }

    /**
     * @notice Generate initCode for Web2 users to enable EntryPoint deployment
     * @param identifier User identifier string
     * @param masterSigner Master signer address for Web2 user
     * @return initCode Bytes for EntryPoint to deploy the account
     */
    function getWeb2InitCode(string calldata identifier, address masterSigner) 
        external 
        view 
        returns (bytes memory initCode) 
    {
        address generatedOwner = generateOwnerFromMasterSigner(masterSigner, identifier);
        bytes32 salt = identifierToSalt(identifier);
        
        initCode = abi.encodePacked(
            address(this),
            abi.encodeCall(this.createAccountWithMasterSigner, (generatedOwner, salt, masterSigner))
        );
    }

    /**
     * @notice Check if Web2 account is already deployed
     * @param identifier User identifier string
     * @param masterSigner Master signer address
     * @return isDeployed True if account is already deployed
     */
    function isWeb2AccountDeployed(string calldata identifier, address masterSigner) 
        external 
        view 
        returns (bool isDeployed) 
    {
        address accountAddress = this.getWeb2AccountAddress(identifier, masterSigner);
        return accountAddress.code.length > 0;
    }


    /**
     * @notice Get the account implementation address
     */
    function getImplementation() external view override returns (address) {
        return accountImplementation;
    }

    /**
     * @notice Convert string identifier to deterministic salt
     */
    function identifierToSalt(string calldata identifier) 
        public 
        pure 
        override 
        returns (bytes32 salt) 
    {
        return keccak256(abi.encodePacked(identifier));
    }


    /**
     * @notice Generate deterministic owner address from master signer and identifier
     * @dev This creates unique owner addresses even with the same identifier across different master signers
     */
    function generateOwnerFromMasterSigner(address masterSigner, string calldata identifier) 
        public 
        pure 
        override 
        returns (address owner) 
    {
        // Combine master signer address with identifier for uniqueness
        bytes32 hash = keccak256(abi.encodePacked("_MASTER_OWNER:", masterSigner, ":", identifier));
        return address(uint160(uint256(hash)));
    }

    /**
     * @notice Authorize contract upgrades
     * @dev Only owner can authorize upgrades
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {
        // Additional upgrade logic can be added here if needed
    }

    /**
     * @notice Update the account implementation
     * @dev Only owner can update implementation
     * @param newImplementation New AAWallet implementation address
     */
    function updateAccountImplementation(address newImplementation) external onlyOwner {
        require(newImplementation != address(0), "WalletManager: invalid implementation");
        accountImplementation = newImplementation;
    }

    /**
     * @notice Set MasterAggregator address
     * @dev Only owner can set the aggregator
     * @param aggregator MasterAggregator contract address
     */
    function setMasterAggregator(address aggregator) external onlyOwner {
        require(aggregator != address(0), "WalletManager: invalid aggregator");
        masterAggregatorAddress = aggregator;
    }

    /**
     * @notice Authorize an address to create accounts
     */
    function authorizeCreator(address creator) external override onlyOwner {
        require(creator != address(0), "WalletManager: invalid creator");
        require(!authorizedCreators[creator], "WalletManager: already authorized");
        
        authorizedCreators[creator] = true;
        emit CreatorAuthorized(creator);
    }

    /**
     * @notice Revoke authorization from an address
     */
    function revokeCreator(address creator) external override onlyOwner {
        require(authorizedCreators[creator], "WalletManager: not authorized");
        require(creator != owner(), "WalletManager: cannot revoke owner");
        
        authorizedCreators[creator] = false;
        emit CreatorRevoked(creator);
    }

    /**
     * @notice Check if an address is authorized to create accounts
     */
    function isAuthorizedCreator(address creator) external view override returns (bool) {
        return authorizedCreators[creator] || creator == owner();
    }

}