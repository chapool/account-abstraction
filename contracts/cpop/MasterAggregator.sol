// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import "../interfaces/IAggregator.sol";
import "../interfaces/IEntryPoint.sol";
import "./interfaces/IAAWallet.sol";
import "./interfaces/IMasterAggregator.sol";

/**
 * @title MasterAggregator
 * @notice Signature aggregator for wallets controlled by master signers
 * @dev Optimized for scenarios where one master signer controls multiple wallets
 */
contract MasterAggregator is 
    Initializable, 
    IMasterAggregator, 
    OwnableUpgradeable, 
    UUPSUpgradeable 
{
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    // Master signer management
    mapping(address => bool) public authorizedMasters;
    mapping(address => mapping(address => bool)) public masterToWallets; // master => wallet => authorized
    
    // Aggregation configuration
    uint256 public maxAggregatedOps = 50; // Maximum operations per aggregation
    uint256 public validationWindow = 300; // 5 minutes validation window
    
    // Nonce management for aggregated operations
    mapping(address => uint256) public masterNonces;
    
    // Events are defined in IMasterAggregator interface

    modifier onlyAuthorizedMaster() {
        require(authorizedMasters[msg.sender], "MasterAggregator: not authorized master");
        _;
    }

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize the aggregator
     * @param _owner Owner of the contract
     * @param _initialMasters Initial authorized master signers
     */
    function initialize(
        address _owner,
        address[] memory _initialMasters
    ) external initializer {
        require(_owner != address(0), "MasterAggregator: invalid owner");
        
        __Ownable_init(_owner);
        __UUPSUpgradeable_init();
        
        // Authorize initial masters
        for (uint256 i = 0; i < _initialMasters.length; i++) {
            if (_initialMasters[i] != address(0)) {
                authorizedMasters[_initialMasters[i]] = true;
                emit MasterAuthorized(_initialMasters[i], true);
            }
        }
    }

    /**
     * @inheritdoc IAggregator
     * @dev Validates that all operations come from wallets controlled by the same master signer
     */
    function validateSignatures(
        PackedUserOperation[] calldata userOps,
        bytes calldata signature
    ) external override {
        require(userOps.length > 0, "MasterAggregator: empty operations");
        require(userOps.length <= maxAggregatedOps, "MasterAggregator: too many operations");
        
        // Decode aggregated signature
        (address masterSigner, uint256 nonce, bytes memory masterSignature) = abi.decode(
            signature, 
            (address, uint256, bytes)
        );
        
        // Verify master signer is authorized
        require(authorizedMasters[masterSigner], "MasterAggregator: unauthorized master");
        
        // Verify nonce
        require(nonce == masterNonces[masterSigner], "MasterAggregator: invalid nonce");
        masterNonces[masterSigner]++;
        
        // Create hash of all user operations
        bytes32 aggregatedHash = _createAggregatedHash(userOps, masterSigner, nonce);
        
        // Verify master signature
        bytes32 ethHash = aggregatedHash.toEthSignedMessageHash();
        address recovered = ethHash.recover(masterSignature);
        require(recovered == masterSigner, "MasterAggregator: invalid master signature");
        
        // Verify all wallets are controlled by this master
        for (uint256 i = 0; i < userOps.length; i++) {
            address wallet = userOps[i].sender;
            require(
                masterToWallets[masterSigner][wallet] || _isValidWallet(wallet, masterSigner),
                "MasterAggregator: wallet not controlled by master"
            );
        }
        
        emit AggregatedValidation(masterSigner, userOps.length, aggregatedHash);
    }

    /**
     * @inheritdoc IAggregator
     * @dev Returns empty signature as validation is done by master aggregation
     */
    function validateUserOpSignature(
        PackedUserOperation calldata /* userOp */
    ) external pure override returns (bytes memory) {
        // For aggregated operations, individual signatures are not needed
        // The master signature covers all operations
        return "";
    }

    /**
     * @inheritdoc IAggregator
     * @dev Creates aggregated signature from individual user operations
     */
    function aggregateSignatures(
        PackedUserOperation[] calldata userOps
    ) external pure override returns (bytes memory aggregatedSignature) {
        require(userOps.length > 0, "MasterAggregator: empty operations");
        
        // For this implementation, we expect the master signer to be provided externally
        // This function is mainly for interface compliance
        // Real aggregation happens off-chain with master signer
        
        return abi.encode(address(0), uint256(0), bytes(""));
    }

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
        bytes32 masterPrivateKey // This would be used off-chain only
    ) external view override returns (bytes memory aggregatedSignature) {
        require(authorizedMasters[masterSigner], "MasterAggregator: unauthorized master");
        
        uint256 nonce = masterNonces[masterSigner];
        bytes32 aggregatedHash = _createAggregatedHash(userOps, masterSigner, nonce);
        
        // Note: In real implementation, this would be done off-chain
        // This is just for demonstration
        bytes32 ethHash = aggregatedHash.toEthSignedMessageHash();
        
        // Create signature (this would be done off-chain with actual private key)
        bytes memory masterSignature = _signHash(ethHash, masterPrivateKey);
        
        return abi.encode(masterSigner, nonce, masterSignature);
    }

    /**
     * @notice Authorize/deauthorize master signer
     * @param master Master signer address
     * @param authorized Whether to authorize or deauthorize
     */
    function setMasterAuthorization(address master, bool authorized) external override onlyOwner {
        require(master != address(0), "MasterAggregator: invalid master");
        authorizedMasters[master] = authorized;
        emit MasterAuthorized(master, authorized);
    }

    /**
     * @notice Authorize/deauthorize wallet for a master signer
     * @param master Master signer address
     * @param wallet Wallet address
     * @param authorized Whether to authorize or deauthorize
     */
    function setWalletAuthorization(
        address master, 
        address wallet, 
        bool authorized
    ) external override onlyOwner {
        require(authorizedMasters[master], "MasterAggregator: master not authorized");
        require(wallet != address(0), "MasterAggregator: invalid wallet");
        
        masterToWallets[master][wallet] = authorized;
        emit WalletAuthorized(master, wallet, authorized);
    }

    /**
     * @notice Batch authorize wallets for a master signer
     * @param master Master signer address
     * @param wallets Array of wallet addresses
     * @param authorized Whether to authorize or deauthorize
     */
    function batchSetWalletAuthorization(
        address master,
        address[] calldata wallets,
        bool authorized
    ) external override onlyOwner {
        require(authorizedMasters[master], "MasterAggregator: master not authorized");
        
        for (uint256 i = 0; i < wallets.length; i++) {
            if (wallets[i] != address(0)) {
                masterToWallets[master][wallets[i]] = authorized;
                emit WalletAuthorized(master, wallets[i], authorized);
            }
        }
    }

    /**
     * @notice Auto-authorize wallet by verifying master signer relationship
     * @param wallet Wallet address
     * @param master Master signer address
     */
    function autoAuthorizeWallet(address wallet, address master) external override {
        require(authorizedMasters[master], "MasterAggregator: master not authorized");
        require(_isValidWallet(wallet, master), "MasterAggregator: invalid wallet-master relationship");
        
        masterToWallets[master][wallet] = true;
        emit WalletAuthorized(master, wallet, true);
    }

    /**
     * @notice Update aggregation configuration
     * @param _maxAggregatedOps Maximum operations per aggregation
     * @param _validationWindow Validation window in seconds
     */
    function updateConfig(
        uint256 _maxAggregatedOps,
        uint256 _validationWindow
    ) external override onlyOwner {
        require(_maxAggregatedOps > 0 && _maxAggregatedOps <= 100, "MasterAggregator: invalid max ops");
        require(_validationWindow > 0 && _validationWindow <= 3600, "MasterAggregator: invalid window");
        
        maxAggregatedOps = _maxAggregatedOps;
        validationWindow = _validationWindow;
    }

    /**
     * @notice Check if wallet is controlled by master signer
     * @param wallet Wallet address
     * @param master Master signer address
     * @return isValid True if wallet is controlled by master
     */
    function isWalletControlledByMaster(
        address wallet, 
        address master
    ) external view override returns (bool isValid) {
        return masterToWallets[master][wallet] || _isValidWallet(wallet, master);
    }

    /**
     * @notice Get current nonce for master signer
     * @param master Master signer address
     * @return nonce Current nonce
     */
    function getMasterNonce(address master) external view override returns (uint256 nonce) {
        return masterNonces[master];
    }

    /**
     * @notice Calculate gas savings from aggregation
     * @param operationCount Number of operations
     * @return savings Estimated gas savings
     */
    function calculateGasSavings(uint256 operationCount) external pure override returns (uint256 savings) {
        if (operationCount <= 1) return 0;
        
        // Each individual operation has signature validation overhead (~3000 gas)
        // Aggregated validation has fixed cost (~5000 gas) + marginal cost per op (~500 gas)
        uint256 individualCost = operationCount * 3000;
        uint256 aggregatedCost = 5000 + (operationCount * 500);
        
        return individualCost > aggregatedCost ? individualCost - aggregatedCost : 0;
    }

    /**
     * @notice Create aggregated hash for signature
     * @param userOps Array of user operations
     * @param masterSigner Master signer address
     * @param nonce Current nonce
     * @return hash Aggregated hash
     */
    function _createAggregatedHash(
        PackedUserOperation[] calldata userOps,
        address masterSigner,
        uint256 nonce
    ) internal view returns (bytes32 hash) {
        bytes32[] memory opHashes = new bytes32[](userOps.length);
        
        for (uint256 i = 0; i < userOps.length; i++) {
            opHashes[i] = keccak256(abi.encode(
                userOps[i].sender,
                userOps[i].nonce,
                userOps[i].callData,
                userOps[i].accountGasLimits,
                userOps[i].preVerificationGas,
                userOps[i].gasFees,
                userOps[i].paymasterAndData
            ));
        }
        
        return keccak256(abi.encode(
            "MASTER_AGGREGATION",
            masterSigner,
            nonce,
            block.chainid,
            address(this),
            opHashes
        ));
    }

    /**
     * @notice Check if wallet is a valid  wallet controlled by master
     * @param wallet Wallet address
     * @param master Master signer address
     * @return isValid True if valid relationship
     */
    function _isValidWallet(address wallet, address master) internal view returns (bool isValid) {
        // Try to call the wallet to check if master signer matches
        try IAAWallet(wallet).getMasterSigner() returns (address walletMaster) {
            return walletMaster == master;
        } catch {
            return false;
        }
    }

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
        bytes32 sessionKeyPrivateKey // This would be used off-chain only
    ) external view override returns (bytes memory aggregatedSignature) {
        require(userOps.length > 0, "MasterAggregator: empty operations");
        
        // Verify all wallets support session keys and this session key is valid
        for (uint256 i = 0; i < userOps.length; i++) {
            address wallet = userOps[i].sender;
            try IAAWallet(wallet).getSessionKeyInfo(sessionKey) returns (
                bool isValid,
                uint48 /* validAfter */,
                uint48 /* validUntil */,
                bytes32 /* permissions */
            ) {
                require(isValid, "MasterAggregator: invalid session key for wallet");
            } catch {
                revert("MasterAggregator: wallet does not support session keys");
            }
        }
        
        // Create hash for session key signing
        bytes32 sessionKeyHash = _createSessionKeyAggregatedHash(userOps, sessionKey);
        bytes32 ethHash = sessionKeyHash.toEthSignedMessageHash();
        
        // Create signature (this would be done off-chain with actual private key)
        bytes memory sessionKeySignature = _signHash(ethHash, sessionKeyPrivateKey);
        
        return abi.encode("SESSION_KEY", sessionKey, sessionKeySignature);
    }

    /**
     * @notice Validate session key aggregated signature
     * @param userOps Array of user operations
     * @param signature Session key aggregated signature
     */
    function validateSessionKeyAggregatedSignature(
        PackedUserOperation[] calldata userOps,
        bytes calldata signature
    ) external view override returns (bool isValid) {
        require(userOps.length > 0, "MasterAggregator: empty operations");
        
        // Decode session key signature
        (string memory sigType, address sessionKey, bytes memory sessionKeySignature) = abi.decode(
            signature,
            (string, address, bytes)
        );
        
        require(
            keccak256(abi.encode(sigType)) == keccak256(abi.encode("SESSION_KEY")),
            "MasterAggregator: invalid signature type"
        );
        
        // Verify session key is valid for all wallets
        for (uint256 i = 0; i < userOps.length; i++) {
            address wallet = userOps[i].sender;
            try IAAWallet(wallet).getSessionKeyInfo(sessionKey) returns (
                bool keyIsValid,
                uint48 /* validAfter */,
                uint48 /* validUntil */,
                bytes32 /* permissions */
            ) {
                if (!keyIsValid) {
                    return false;
                }
            } catch {
                return false;
            }
        }
        
        // Verify session key signature
        bytes32 sessionKeyHash = _createSessionKeyAggregatedHash(userOps, sessionKey);
        bytes32 ethHash = sessionKeyHash.toEthSignedMessageHash();
        address recovered = ethHash.recover(sessionKeySignature);
        
        return recovered == sessionKey;
    }

    /**
     * @notice Create aggregated hash for session key signature
     * @param userOps Array of user operations
     * @param sessionKey Session key address
     * @return hash Session key aggregated hash
     */
    function _createSessionKeyAggregatedHash(
        PackedUserOperation[] calldata userOps,
        address sessionKey
    ) internal view returns (bytes32 hash) {
        bytes32[] memory opHashes = new bytes32[](userOps.length);
        
        for (uint256 i = 0; i < userOps.length; i++) {
            opHashes[i] = keccak256(abi.encode(
                userOps[i].sender,
                userOps[i].nonce,
                userOps[i].callData,
                userOps[i].accountGasLimits,
                userOps[i].preVerificationGas,
                userOps[i].gasFees,
                userOps[i].paymasterAndData
            ));
        }
        
        return keccak256(abi.encode(
            "SESSION_KEY_AGGREGATION",
            sessionKey,
            block.chainid,
            address(this),
            block.timestamp,
            opHashes
        ));
    }

    /**
     * @notice Sign hash with private key (off-chain helper)
     * @dev This is for demonstration only, real signing should be done off-chain
     */
    function _signHash(bytes32 /* hash */, bytes32 /* privateKey */) internal pure returns (bytes memory signature) {
        // This is a placeholder - in real implementation, signing is done off-chain
        // Using a mock signature for demonstration
        return abi.encodePacked(bytes32(0), bytes32(0), uint8(27));
    }

    /**
     * @notice Add stake to EntryPoint for this aggregator
     * @param entryPoint EntryPoint address
     * @param delay Unstake delay
     */
    function addStake(IEntryPoint entryPoint, uint32 delay) external payable onlyOwner {
        entryPoint.addStake{value: msg.value}(delay);
    }

    /**
     * @notice Unlock stake from EntryPoint
     * @param entryPoint EntryPoint address
     */
    function unlockStake(IEntryPoint entryPoint) external onlyOwner {
        entryPoint.unlockStake();
    }

    /**
     * @notice Withdraw stake from EntryPoint
     * @param entryPoint EntryPoint address
     * @param withdrawAddress Address to receive withdrawn stake
     */
    function withdrawStake(IEntryPoint entryPoint, address payable withdrawAddress) external onlyOwner {
        entryPoint.withdrawStake(withdrawAddress);
    }

    /**
     * @notice Authorize contract upgrades
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /**
     * @notice Emergency pause mechanism
     */
    function pause() external onlyOwner {
        // Could implement pausable functionality if needed
    }
}