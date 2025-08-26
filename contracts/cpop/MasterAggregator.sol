// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import "../interfaces/IAggregator.sol";
import "../interfaces/IEntryPoint.sol";
import "./interfaces/IAAccount.sol";
import "./interfaces/IMasterAggregator.sol";

/**
 * @title MasterAggregator
 * @notice Signature aggregator for Accounts controlled by master signers
 * @dev Optimized for scenarios where one master signer controls multiple Accounts
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
    mapping(address => mapping(address => bool)) public masterToAccounts; // master => Account => authorized
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
     * @dev Validates that all operations come from Accounts controlled by the same master signer
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
        
        // Verify all Accounts are controlled by this master
        for (uint256 i = 0; i < userOps.length; i++) {
            address Account = userOps[i].sender;
            require(
                masterToAccounts[masterSigner][Account] || _isValidAccount(Account, masterSigner),
                "MasterAggregator: Account not controlled by master"
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
        
        // For master aggregation, we need to identify which master controls these Accounts
        // Check if all operations belong to Accounts controlled by the same master
        address masterSigner = address(0);
        
        for (uint256 i = 0; i < userOps.length; i++) {
            address Account = userOps[i].sender;
            
            // Try to find the master for this Account
            address AccountMaster = _findMasterForAccount(Account);
            require(AccountMaster != address(0), "MasterAggregator: Account has no master");
            
            if (masterSigner == address(0)) {
                masterSigner = AccountMaster;
            } else {
                require(masterSigner == AccountMaster, "MasterAggregator: operations from different masters");
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
     * @notice Create master aggregated signature for multiple Accounts
     * @dev Master signer can control multiple Accounts with one signature
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
        
        // Validate all operations belong to Accounts controlled by master
        for (uint256 i = 0; i < userOps.length; i++) {
            require(
                this.isAccountControlledByMaster(userOps[i].sender, masterSigner),
                "MasterAggregator: Account not controlled by master"
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
     * @dev Find the master signer that controls a given Account
     * @param Account The Account address to search for
     * @return master The master address that controls the Account, or address(0) if not found
     */
    function _findMasterForAccount(address Account) internal view returns (address master) {
        // Iterate through all authorized masters to find which one controls this Account
        for (uint256 i = 0; i < authorizedMastersList.length; i++) {
            address potentialMaster = authorizedMastersList[i];
            if (authorizedMasters[potentialMaster] && masterToAccounts[potentialMaster][Account]) {
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
     * @notice Authorize/deauthorize Account for a master signer
     * @param master Master signer address
     * @param Account Account address
     * @param authorized Whether to authorize or deauthorize
     */
    function setAccountAuthorization(
        address master, 
        address Account, 
        bool authorized
    ) external override onlyOwner {
        require(authorizedMasters[master], "MasterAggregator: master not authorized");
        require(Account != address(0), "MasterAggregator: invalid Account");
        
        masterToAccounts[master][Account] = authorized;
        emit AccountAuthorized(master, Account, authorized);
    }

    /**
     * @notice Batch authorize Accounts for a master signer
     * @param master Master signer address
     * @param Accounts Array of Account addresses
     * @param authorized Whether to authorize or deauthorize
     */
    function batchSetAccountAuthorization(
        address master,
        address[] calldata Accounts,
        bool authorized
    ) external override onlyOwner {
        require(authorizedMasters[master], "MasterAggregator: master not authorized");
        
        for (uint256 i = 0; i < Accounts.length; i++) {
            if (Accounts[i] != address(0)) {
                masterToAccounts[master][Accounts[i]] = authorized;
                emit AccountAuthorized(master, Accounts[i], authorized);
            }
        }
    }

    /**
     * @notice Auto-authorize Account by verifying master signer relationship
     * @param Account Account address
     * @param master Master signer address
     */
    function autoAuthorizeAccount(address Account, address master) external override {
        require(authorizedMasters[master], "MasterAggregator: master not authorized");
        require(_isValidAccount(Account, master), "MasterAggregator: invalid Account-master relationship");
        
        masterToAccounts[master][Account] = true;
        emit AccountAuthorized(master, Account, true);
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
     * @notice Check if Account is controlled by master signer
     * @param Account Account address
     * @param master Master signer address
     * @return isValid True if Account is controlled by master
     */
    function isAccountControlledByMaster(
        address Account, 
        address master
    ) external view override returns (bool isValid) {
        return masterToAccounts[master][Account] || _isValidAccount(Account, master);
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
     * @notice Check if Account is a valid  Account controlled by master
     * @param Account Account address
     * @param master Master signer address
     * @return isValid True if valid relationship
     */
    function _isValidAccount(address Account, address master) internal view returns (bool isValid) {
        // Try to call the Account to check if master signer matches
        try IAAccount(Account).getMasterSigner() returns (address AccountMaster) {
            return AccountMaster == master;
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