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
    address[] public authorizedMastersList; // Array to track all authorized masters
    
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
                authorizedMastersList.push(_initialMasters[i]); // Add to list
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
        // Note: masterSignature was created by signing the ethHash from getMasterSigningData
        // which already applied toEthSignedMessageHash(), so we need to recover from that
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
     * @dev Creates aggregated signature placeholder for EntryPoint
     * @notice This returns a placeholder that indicates master aggregation is needed
     */
    function aggregateSignatures(
        PackedUserOperation[] calldata userOps
    ) external view override returns (bytes memory aggregatedSignature) {
        require(userOps.length > 0, "MasterAggregator: empty operations");
        require(userOps.length <= maxAggregatedOps, "MasterAggregator: too many operations");
        
        // For master aggregation, we need to identify which master controls these wallets
        // Check if all operations belong to wallets controlled by the same master
        address masterSigner = address(0);
        
        for (uint256 i = 0; i < userOps.length; i++) {
            address wallet = userOps[i].sender;
            
            // Try to find the master for this wallet
            address walletMaster = _findMasterForWallet(wallet);
            require(walletMaster != address(0), "MasterAggregator: wallet has no master");
            
            if (masterSigner == address(0)) {
                masterSigner = walletMaster;
            } else {
                require(masterSigner == walletMaster, "MasterAggregator: operations from different masters");
            }
        }
        
        uint256 nonce = masterNonces[masterSigner];
        
        // Return aggregation info for EntryPoint
        // The actual master signature will be provided via validateSignatures
        return abi.encode(
            masterSigner,
            nonce, 
            "MASTER_AGGREGATION_PLACEHOLDER"
        );
    }

    /**
     * @notice Create master aggregated signature for multiple wallets
     * @dev Master signer can control multiple wallets with one signature
     * @param userOps Array of user operations to aggregate  
     * @param masterSigner Master signer address
     * @param masterSignature Pre-computed signature from master signer
     * @return aggregatedSignature Encoded aggregated signature for EntryPoint
     */
    function createMasterAggregatedSignature(
        PackedUserOperation[] calldata userOps,
        address masterSigner,
        bytes calldata masterSignature
    ) external override returns (bytes memory aggregatedSignature) {
        require(authorizedMasters[masterSigner], "MasterAggregator: unauthorized master");
        require(userOps.length > 0 && userOps.length <= maxAggregatedOps, "MasterAggregator: invalid operation count");
        
        // Validate all operations belong to wallets controlled by master
        for (uint256 i = 0; i < userOps.length; i++) {
            require(
                this.isWalletControlledByMaster(userOps[i].sender, masterSigner),
                "MasterAggregator: wallet not controlled by master"
            );
        }
        
        uint256 nonce = masterNonces[masterSigner];
        
        // Emit aggregation event for tracking
        emit AggregatedValidation(masterSigner, userOps.length, _createAggregatedHash(userOps, masterSigner, nonce));
        
        return abi.encode(masterSigner, nonce, masterSignature);
    }

    /**
     * @notice Get data that master needs to sign (call this off-chain)
     * @param userOps Array of user operations to aggregate
     * @param masterSigner Master signer address  
     * @return hashToSign The hash that master should sign off-chain
     * @return nonce The nonce to use
     */
    function getMasterSigningData(
        PackedUserOperation[] calldata userOps,
        address masterSigner
    ) external view override returns (bytes32 hashToSign, uint256 nonce) {
        require(authorizedMasters[masterSigner], "MasterAggregator: unauthorized master");
        require(userOps.length > 0, "MasterAggregator: empty operations");
        
        nonce = masterNonces[masterSigner];
        bytes32 aggregatedHash = _createAggregatedHash(userOps, masterSigner, nonce);
        hashToSign = aggregatedHash.toEthSignedMessageHash();
        
        return (hashToSign, nonce);
    }

    /**
     * @dev Find the master signer that controls a given wallet
     * @param wallet The wallet address to search for
     * @return master The master address that controls the wallet, or address(0) if not found
     */
    function _findMasterForWallet(address wallet) internal view returns (address master) {
        // Iterate through all authorized masters to find which one controls this wallet
        for (uint256 i = 0; i < authorizedMastersList.length; i++) {
            address potentialMaster = authorizedMastersList[i];
            if (authorizedMasters[potentialMaster] && masterToWallets[potentialMaster][wallet]) {
                return potentialMaster;
            }
        }
        return address(0);
    }

    /**
     * @dev Remove master from the authorized list
     */
    function _removeMasterFromList(address master) internal {
        for (uint256 i = 0; i < authorizedMastersList.length; i++) {
            if (authorizedMastersList[i] == master) {
                // Move last element to current position and remove last
                authorizedMastersList[i] = authorizedMastersList[authorizedMastersList.length - 1];
                authorizedMastersList.pop();
                break;
            }
        }
    }

    /**
     * @notice Authorize/deauthorize master signer
     * @param master Master signer address
     * @param authorized Whether to authorize or deauthorize
     */
    function setMasterAuthorization(address master, bool authorized) external override onlyOwner {
        require(master != address(0), "MasterAggregator: invalid master");
        
        bool wasAuthorized = authorizedMasters[master];
        authorizedMasters[master] = authorized;
        
        if (authorized && !wasAuthorized) {
            // Add to list when first authorized
            authorizedMastersList.push(master);
        } else if (!authorized && wasAuthorized) {
            // Remove from list when deauthorized
            _removeMasterFromList(master);
        }
        
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