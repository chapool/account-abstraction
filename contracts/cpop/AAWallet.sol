// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import "../core/BaseAccount.sol";
import "../core/Helpers.sol";
import "./interfaces/IAAWallet.sol";

/**
 * @title AAWallet
 * @notice Account Abstraction wallet optimized for Web2 users
 * @dev Inherits from BaseAccount and implements IAAWallet interface
 * Simplified design without token-specific functionality
 */
contract AAWallet is Initializable, BaseAccount, IAAWallet, UUPSUpgradeable, ERC165 {
    using ECDSA for bytes32;
    using UserOperationLib for PackedUserOperation;

    // Using constants from Helpers.sol instead of redefining

    // Additional events for aggregation support
    event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator);

    // Session key data structure
    struct SessionKeyData {
        uint48 validAfter;  // Start time (0 for immediate)
        uint48 validUntil;  // End time
        bytes32 permissions; // Encoded permissions
        bool isActive;      // Whether the session key is active
    }

    address public owner;
    address public masterSigner; // Master signer for Web2 users
    address public aggregatorAddress; // Aggregator for signature aggregation
    address public immutable entryPointAddress;
    
    // Session keys mapping
    mapping(address => SessionKeyData) public sessionKeys;

    modifier onlyOwner() {
        require(msg.sender == owner || msg.sender == address(this), "AAWallet: not owner");
        _;
    }

    modifier onlyMasterSigner() {
        require(msg.sender == masterSigner || msg.sender == address(this), "AAWallet: not master signer");
        _;
    }

    constructor(address _entryPoint) {
        require(_entryPoint != address(0), "AAWallet: invalid entryPoint");
        
        entryPointAddress = _entryPoint;
        _disableInitializers();
    }

    /**
     * @notice Initialize the account with an owner
     */
    function initialize(address _owner) external initializer {
        require(_owner != address(0), "AAWallet: invalid owner");
        owner = _owner;
        emit AccountInitialized(_owner, address(0));
    }

    /**
     * @notice Initialize the account with an owner and master signer for Web2 users
     */
    function initializeWithMasterSigner(address _owner, address _masterSigner) external initializer {
        require(_owner != address(0), "AAWallet: invalid owner");
        require(_masterSigner != address(0), "AAWallet: invalid master signer");
        owner = _owner;
        masterSigner = _masterSigner;
        emit AccountInitialized(_owner, _masterSigner);
    }

    /**
     * @notice Get the EntryPoint used by this account
     */
    function entryPoint() public view override returns (IEntryPoint) {
        return IEntryPoint(entryPointAddress);
    }

    /**
     * @notice Get the owner of this account
     */
    function getOwner() external view override returns (address) {
        return owner;
    }

    /**
     * @notice Get the master signer of this account
     */
    function getMasterSigner() external view override returns (address) {
        return masterSigner;
    }

    /**
     * @notice Set a new master signer (only current master signer can call)
     */
    function setMasterSigner(address newMasterSigner) external override onlyMasterSigner {
        address oldMasterSigner = masterSigner;
        masterSigner = newMasterSigner;
        emit MasterSignerUpdated(oldMasterSigner, newMasterSigner);
    }

    /**
     * @notice Check if master signer validation is enabled
     */
    function isMasterSignerEnabled() external view override returns (bool) {
        return masterSigner != address(0);
    }




    /**
     * @notice Validate user operation signature
     * @dev Supports owner signature, master signer signature, session keys, and aggregated signatures
     */
    function _validateSignature(
        PackedUserOperation calldata userOp,
        bytes32 userOpHash
    ) internal override returns (uint256 validationData) {
        // Check if this is an aggregated operation (empty signature means aggregator will handle)
        if (userOp.signature.length == 0) {
            // Return aggregator address if master signer is set
            if (masterSigner != address(0) && aggregatorAddress != address(0)) {
                return _packValidationData(ValidationData(aggregatorAddress, 0, 0));
            }
            return SIG_VALIDATION_FAILED;
        }
        
        bytes32 hash = MessageHashUtils.toEthSignedMessageHash(userOpHash);
        address recovered = ECDSA.recover(hash, userOp.signature);
        
        // First try owner signature (highest priority)
        if (recovered == owner) {
            return SIG_VALIDATION_SUCCESS;
        }
        
        // Then try master signer signature
        if (masterSigner != address(0) && recovered == masterSigner) {
            emit MasterSignerValidation(masterSigner, userOpHash);
            return SIG_VALIDATION_SUCCESS;
        }
        
        // Finally try session key validation
        SessionKeyData memory sessionKey = sessionKeys[recovered];
        if (sessionKey.isActive) {
            // Check time validity
            uint48 currentTime = uint48(block.timestamp);
            if (currentTime >= sessionKey.validAfter && currentTime <= sessionKey.validUntil) {
                // Check permission for the specific operation
                if (_validateSessionKeyPermission(recovered, userOp)) {
                    emit SessionKeyValidation(recovered, userOpHash);
                    return _packValidationData(ValidationData(address(0), sessionKey.validUntil, sessionKey.validAfter));
                }
            }
        }
        
        return SIG_VALIDATION_FAILED;
    }

    /**
     * @notice Get the aggregator address for this account
     * @return aggregator Address of the aggregator contract
     */
    function getAggregator() public view returns (address aggregator) {
        // This could be set during initialization or managed by the wallet manager
        // For now, return a default aggregator address stored in the implementation
        return aggregatorAddress;
    }

    /**
     * @notice Set aggregator address (only owner can set)
     * @param _aggregator New aggregator address
     */
    function setAggregator(address _aggregator) external onlyOwner {
        aggregatorAddress = _aggregator;
        emit AggregatorUpdated(aggregatorAddress, _aggregator);
    }


    /**
     * @notice Check execution authorization
     * @dev Only EntryPoint can call execute functions for security
     */
    function _requireForExecute() internal view override {
        require(
            msg.sender == address(entryPoint()),
            "AAWallet: only EntryPoint can execute"
        );
    }

    /**
     * @notice Authorize contract upgrades
     */
    function _authorizeUpgrade(address newImplementation) internal view override {
        (newImplementation); // suppress unused parameter warning
        require(msg.sender == owner, "AAWallet: only owner can upgrade");
    }

    /**
     * @notice Get account nonce
     */
    function getNonce() public view override returns (uint256) {
        return entryPoint().getNonce(address(this), 0);
    }

    /**
     * @notice Get account deposit in EntryPoint
     */
    function getDeposit() external view returns (uint256) {
        return entryPoint().balanceOf(address(this));
    }

    /**
     * @notice Add deposit to EntryPoint
     */
    function addDeposit() external payable {
        entryPoint().depositTo{value: msg.value}(address(this));
    }

    /**
     * @notice Withdraw deposit from EntryPoint
     */
    function withdrawDepositTo(address payable withdrawAddress, uint256 amount) external onlyOwner {
        entryPoint().withdrawTo(withdrawAddress, amount);
    }

    // Session Keys Management Functions
    
    /**
     * @notice Add a session key with time and permission constraints
     */
    function addSessionKey(
        address sessionKey,
        uint48 validAfter,
        uint48 validUntil,
        bytes32 permissions
    ) external override onlyOwner {
        require(sessionKey != address(0), "AAWallet: invalid session key");
        require(validUntil > validAfter, "AAWallet: invalid time range");
        require(validUntil > block.timestamp, "AAWallet: session key already expired");
        require(!sessionKeys[sessionKey].isActive, "AAWallet: session key already exists");
        
        sessionKeys[sessionKey] = SessionKeyData({
            validAfter: validAfter,
            validUntil: validUntil,
            permissions: permissions,
            isActive: true
        });
        
        emit SessionKeyAdded(sessionKey, validAfter, validUntil, permissions);
    }
    
    /**
     * @notice Revoke a session key
     */
    function revokeSessionKey(address sessionKey) external override onlyOwner {
        require(sessionKeys[sessionKey].isActive, "AAWallet: session key not active");
        
        delete sessionKeys[sessionKey];
        emit SessionKeyRevoked(sessionKey);
    }
    
    /**
     * @notice Get session key information
     */
    function getSessionKeyInfo(address sessionKey) 
        external 
        view 
        override 
        returns (bool isValid, uint48 validAfter, uint48 validUntil, bytes32 permissions) 
    {
        SessionKeyData memory data = sessionKeys[sessionKey];
        isValid = data.isActive && 
                 block.timestamp >= data.validAfter && 
                 block.timestamp <= data.validUntil;
        validAfter = data.validAfter;
        validUntil = data.validUntil;
        permissions = data.permissions;
    }
    
    /**
     * @notice Check if a session key can execute a specific operation
     */
    function canSessionKeyExecute(
        address sessionKey,
        address target,
        bytes4 selector
    ) external view override returns (bool canExecute) {
        SessionKeyData memory data = sessionKeys[sessionKey];
        if (!data.isActive || 
            block.timestamp < data.validAfter || 
            block.timestamp > data.validUntil) {
            return false;
        }
        
        // Decode permissions to check target and selector allowances
        return _checkPermissions(data.permissions, target, selector);
    }
    
    /**
     * @notice Validate session key permission for a user operation
     */
    function _validateSessionKeyPermission(
        address sessionKey, 
        PackedUserOperation calldata userOp
    ) internal view returns (bool isValid) {
        SessionKeyData memory data = sessionKeys[sessionKey];
        
        // Extract target and selector from callData
        if (userOp.callData.length < 4) {
            return false;
        }
        
        address target;
        bytes4 selector;
        
        // Parse execute call to get target and selector
        bytes4 executeSelector = bytes4(userOp.callData[:4]);
        if (executeSelector == this.execute.selector) {
            // execute(address,uint256,bytes)
            if (userOp.callData.length < 68) return false;
            target = address(bytes20(userOp.callData[16:36]));
            if (userOp.callData.length >= 72) {
                selector = bytes4(userOp.callData[68:72]);
            }
        } else if (executeSelector == this.executeBatch.selector) {
            // For batch operations, check if all targets/selectors are allowed
            // Simplified check - in production, would parse all batch items
            return _checkBatchPermissions(data.permissions, userOp.callData);
        } else {
            // Direct account function calls
            target = address(this);
            selector = executeSelector;
        }
        
        return _checkPermissions(data.permissions, target, selector);
    }
    
    /**
     * @notice Check permissions for target and selector
     */
    function _checkPermissions(
        bytes32 permissions, 
        address target, 
        bytes4 selector
    ) internal pure returns (bool allowed) {
        // Simple permission encoding: 
        // - 0x0: allow all
        // - specific encoding for target/selector combinations
        if (permissions == bytes32(0)) {
            return true; // Allow all operations
        }
        
        // Extract permission data
        // Format: [targetHash(16 bytes)][selectorHash(4 bytes)][flags(12 bytes)]
        bytes16 targetHash = bytes16(permissions);
        bytes4 selectorHash = bytes4(permissions << 128);
        
        // Check if target matches (simplified hash comparison)
        bool targetMatches = (targetHash == bytes16(0)) || 
                            (targetHash == bytes16(keccak256(abi.encode(target))));
        
        // Check if selector matches
        bool selectorMatches = (selectorHash == bytes4(0)) || 
                              (selectorHash == selector);
        
        return targetMatches && selectorMatches;
    }
    
    /**
     * @notice Check permissions for batch operations
     */
    function _checkBatchPermissions(
        bytes32 permissions,
        bytes calldata /* callData */
    ) internal pure returns (bool allowed) {
        // Simplified batch permission check
        // In production, would parse batch data and check each operation
        return permissions == bytes32(0); // For now, allow all if unrestricted
    }

    /**
     * @notice Check interface support for ERC165
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC165, IERC165) returns (bool) {
        return interfaceId == type(IAAWallet).interfaceId || 
               interfaceId == type(IAAWallet).interfaceId ||
               super.supportsInterface(interfaceId);
    }

    /**
     * @notice Receive ETH
     */
    receive() external payable {}
}