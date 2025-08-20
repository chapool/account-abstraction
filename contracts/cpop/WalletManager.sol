// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "../interfaces/ISenderCreator.sol";
import "../interfaces/IEntryPoint.sol";
import "./interfaces/IMasterAggregator.sol";
import "./interfaces/IWalletManager.sol";
import "./AAWallet.sol";

/**
 * @title WalletManager v2 - Simplified EOA + Master Architecture
 * @notice Simplified factory contract for creating AA wallets using EOA + Master signer pattern
 * @dev Creates deterministic wallet addresses using owner EOA + master signer combination
 */
contract WalletManager is Initializable, IWalletManager, OwnableUpgradeable, UUPSUpgradeable {
    address public accountImplementation;
    ISenderCreator public senderCreator;
    address public entryPointAddress;
    address public masterAggregatorAddress;
    address public defaultMasterSigner;
    
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
     * @param owner The owner of this contract
     */
    function initialize(address entryPoint, address owner) external initializer {
        require(entryPoint != address(0), "WalletManager: invalid entryPoint");
        require(owner != address(0), "WalletManager: invalid owner");
        
        // Initialize parent contracts
        __Ownable_init(owner);
        __UUPSUpgradeable_init();
        
        // Initialize contract state
        entryPointAddress = entryPoint;
        accountImplementation = address(new AAWallet());
        senderCreator = IEntryPoint(entryPoint).senderCreator();
        
        // Set default master signer to owner initially
        defaultMasterSigner = owner;
        
        // Authorize the owner as a creator by default
        authorizedCreators[owner] = true;
        emit CreatorAuthorized(owner);
    }

    // ============================================
    // CORE WALLET CREATION FUNCTIONS
    // ============================================

    /**
     * @notice Core wallet creation function - ALL wallet creation goes through this
     * @param owner The wallet owner EOA address
     * @param masterSigner Master signer for aggregation (use address(0) for default)
     * @return account The created wallet address
     */
    function createWallet(address owner, address masterSigner) public returns (address account) {
        require(owner != address(0), "WalletManager: invalid owner");
        
        // Use default master signer if none provided
        address actualMasterSigner = masterSigner != address(0) ? masterSigner : defaultMasterSigner;
        
        // Generate deterministic salt from owner + master combination
        bytes32 salt = _generateSalt(owner, actualMasterSigner);
        
        // Check if wallet already exists
        address addr = _getAccountAddress(owner, actualMasterSigner);
        uint256 codeSize = addr.code.length;
        
        if (codeSize > 0) {
            account = addr;
        } else {
            // All wallets use MasterAggregator initialization if available
            bytes memory initData;
            if (masterAggregatorAddress != address(0) && masterAggregatorAddress.code.length > 0) {
                initData = abi.encodeCall(AAWallet.initializeWithAggregator, 
                    (entryPointAddress, owner, actualMasterSigner, masterAggregatorAddress));
            } else {
                initData = abi.encodeCall(AAWallet.initialize, 
                    (entryPointAddress, owner, actualMasterSigner));
            }
            
            account = address(
                new ERC1967Proxy{salt: salt}(
                    accountImplementation,
                    initData
                )
            );
            
            // Register wallet-master relationship in aggregator if configured
            if (masterAggregatorAddress != address(0) && masterAggregatorAddress.code.length > 0) {
                try IMasterAggregator(masterAggregatorAddress).autoAuthorizeWallet(actualMasterSigner, account) {
                    // Successfully registered wallet-master relationship
                } catch {
                    // Aggregator authorization failed, wallet can still function without aggregation
                }
            }
            
            // Only emit event when wallet is actually created
            emit AccountCreated(account, owner, actualMasterSigner);
        }
    }

    /**
     * @notice Create account for Web2 users (simplified)
     * @param owner User's EOA address
     * @param masterSigner Master signer for transaction control
     * @return account The created AA wallet address
     */
    function createUserAccount(
        address owner,
        address masterSigner
    ) external onlyAuthorizedCreator returns (address account) {
        require(owner != address(0), "WalletManager: invalid owner");
        require(masterSigner != address(0), "WalletManager: invalid master signer");
        
        return createWallet(owner, masterSigner);
    }

    /**
     * @notice Legacy function for EntryPoint compatibility
     * @dev Routes through core createWallet function with provided parameters
     */
    function createAccountWithMasterSigner(
        address owner, 
        address masterSigner
    ) external returns (address account) {
        require(msg.sender == address(senderCreator), "WalletManager: only SenderCreator");
        require(masterSigner != address(0), "WalletManager: invalid master signer");
        return createWallet(owner, masterSigner);
    }

    // ============================================
    // ADDRESS PREDICTION FUNCTIONS
    // ============================================

    /**
     * @notice Get deterministic AA wallet address
     * @param owner User's EOA address
     * @param masterSigner Master signer address
     * @return account Predicted AA wallet address
     */
    function getAccountAddress(address owner, address masterSigner) 
        external 
        view 
        returns (address account) 
    {
        return _getAccountAddress(owner, masterSigner != address(0) ? masterSigner : defaultMasterSigner);
    }

    /**
     * @notice Internal function to compute deterministic address
     */
    function _getAccountAddress(address owner, address masterSigner) 
        internal 
        view 
        returns (address account) 
    {
        bytes32 salt = _generateSalt(owner, masterSigner);
        
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
     * @notice Generate deterministic salt from owner and master signer
     */
    function _generateSalt(address owner, address masterSigner) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("WALLET_V2:", owner, ":", masterSigner));
    }

    // ============================================
    // ENTRYPOINT INTEGRATION
    // ============================================

    /**
     * @notice Generate initCode for EntryPoint deployment
     * @param owner User's EOA address
     * @param masterSigner Master signer address
     * @return initCode Bytes for EntryPoint to deploy the account
     */
    function getInitCode(address owner, address masterSigner) 
        external 
        view 
        returns (bytes memory initCode) 
    {
        initCode = abi.encodePacked(
            address(this),
            abi.encodeCall(this.createAccountWithMasterSigner, (owner, masterSigner))
        );
    }

    /**
     * @notice Check if account is already deployed
     * @param owner User's EOA address
     * @param masterSigner Master signer address
     * @return isDeployed True if account is already deployed
     */
    function isAccountDeployed(address owner, address masterSigner) 
        external 
        view 
        returns (bool isDeployed) 
    {
        address accountAddress = _getAccountAddress(
            owner, 
            masterSigner != address(0) ? masterSigner : defaultMasterSigner
        );
        return accountAddress.code.length > 0;
    }

    // ============================================
    // MANAGEMENT FUNCTIONS
    // ============================================

    /**
     * @notice Get the account implementation address
     */
    function getImplementation() external view returns (address) {
        return accountImplementation;
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
     * @notice Set default master signer
     * @dev Only owner can set the default master signer
     * @param masterSigner Default master signer address
     */
    function setDefaultMasterSigner(address masterSigner) external onlyOwner {
        require(masterSigner != address(0), "WalletManager: invalid master signer");
        defaultMasterSigner = masterSigner;
    }

    /**
     * @notice Get default master signer
     * @return Default master signer address
     */
    function getDefaultMasterSigner() external view returns (address) {
        return defaultMasterSigner;
    }

    /**
     * @notice Authorize an address to create accounts
     */
    function authorizeCreator(address creator) external onlyOwner {
        require(creator != address(0), "WalletManager: invalid creator");
        require(!authorizedCreators[creator], "WalletManager: already authorized");
        
        authorizedCreators[creator] = true;
        emit CreatorAuthorized(creator);
    }

    /**
     * @notice Revoke authorization from an address
     */
    function revokeCreator(address creator) external onlyOwner {
        require(authorizedCreators[creator], "WalletManager: not authorized");
        require(creator != owner(), "WalletManager: cannot revoke owner");
        
        authorizedCreators[creator] = false;
        emit CreatorRevoked(creator);
    }

    /**
     * @notice Check if an address is authorized to create accounts
     */
    function isAuthorizedCreator(address creator) external view returns (bool) {
        return authorizedCreators[creator] || creator == owner();
    }

    /**
     * @notice Authorize contract upgrades
     * @dev Only owner can authorize upgrades
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {
        // Additional upgrade logic can be added here if needed
    }
}