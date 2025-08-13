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
import "./interfaces/ICPOPWalletManager.sol";
import "./CPOPAccount.sol";

/**
 * @title CPOPWalletManager
 * @notice Upgradeable factory contract for creating CPOP Account Abstraction wallets
 * @dev Creates deterministic wallet addresses using CREATE2 for Web2 user experience
 */
contract CPOPWalletManager is Initializable, ICPOPWalletManager, OwnableUpgradeable, UUPSUpgradeable {
    address public accountImplementation;
    ISenderCreator public senderCreator;
    
    mapping(address => bool) private authorizedCreators;

    modifier onlyAuthorizedCreator() {
        require(
            authorizedCreators[msg.sender] || msg.sender == owner(),
            "CPOPWalletManager: unauthorized creator"
        );
        _;
    }

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize the wallet manager
     * @param entryPoint The EntryPoint contract address
     * @param cpopToken The CPOP token contract address
     * @param owner The owner of this contract
     */
    function initialize(address entryPoint, address cpopToken, address owner) external initializer {
        require(entryPoint != address(0), "CPOPWalletManager: invalid entryPoint");
        require(cpopToken != address(0), "CPOPWalletManager: invalid cpopToken");
        require(owner != address(0), "CPOPWalletManager: invalid owner");
        
        // Initialize parent contracts
        __Ownable_init(owner);
        __UUPSUpgradeable_init();
        
        // Initialize contract state
        accountImplementation = address(new CPOPAccount(entryPoint));
        senderCreator = IEntryPoint(entryPoint).senderCreator();
        
        // Authorize the owner as a creator by default
        authorizedCreators[owner] = true;
        emit CreatorAuthorized(owner);
    }

    /**
     * @notice Create a new CPOP account with deterministic address
     */
    function createAccount(address owner, bytes32 salt) 
        external 
        override 
        returns (address account) 
    {
        require(msg.sender == address(senderCreator), "CPOPWalletManager: only SenderCreator");
        account = _createAccount(owner, salt);
    }

    /**
     * @notice Internal function to create account
     */
    function _createAccount(address owner, bytes32 salt) 
        internal 
        returns (address account) 
    {
        require(owner != address(0), "CPOPWalletManager: invalid owner");

        address addr = getAccountAddress(owner, salt);
        uint256 codeSize = addr.code.length;
        
        if (codeSize > 0) {
            return addr;
        }

        account = address(
            new ERC1967Proxy{salt: salt}(
                accountImplementation,
                abi.encodeCall(CPOPAccount.initialize, (owner))
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
        require(masterSigner != address(0), "CPOPWalletManager: invalid master signer");
        
        // Generate owner based on master signer + identifier for better uniqueness
        generatedOwner = generateOwnerFromMasterSigner(masterSigner, identifier);
        bytes32 salt = identifierToSalt(identifier);
        
        // Create account with master signer
        address addr = getAccountAddress(generatedOwner, salt);
        uint256 codeSize = addr.code.length;
        
        if (codeSize > 0) {
            account = addr;
        } else {
            account = address(
                new ERC1967Proxy{salt: salt}(
                    accountImplementation,
                    abi.encodeCall(CPOPAccount.initializeWithMasterSigner, (generatedOwner, masterSigner))
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
        return Create2.computeAddress(
            salt,
            keccak256(
                abi.encodePacked(
                    type(ERC1967Proxy).creationCode,
                    abi.encode(
                        accountImplementation,
                        abi.encodeCall(CPOPAccount.initialize, (owner))
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
        bytes32 hash = keccak256(abi.encodePacked("CPOP_MASTER_OWNER:", masterSigner, ":", identifier));
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
     * @param newImplementation New CPOPAccount implementation address
     */
    function updateAccountImplementation(address newImplementation) external onlyOwner {
        require(newImplementation != address(0), "CPOPWalletManager: invalid implementation");
        accountImplementation = newImplementation;
    }

    /**
     * @notice Authorize an address to create accounts
     */
    function authorizeCreator(address creator) external override onlyOwner {
        require(creator != address(0), "CPOPWalletManager: invalid creator");
        require(!authorizedCreators[creator], "CPOPWalletManager: already authorized");
        
        authorizedCreators[creator] = true;
        emit CreatorAuthorized(creator);
    }

    /**
     * @notice Revoke authorization from an address
     */
    function revokeCreator(address creator) external override onlyOwner {
        require(authorizedCreators[creator], "CPOPWalletManager: not authorized");
        require(creator != owner(), "CPOPWalletManager: cannot revoke owner");
        
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