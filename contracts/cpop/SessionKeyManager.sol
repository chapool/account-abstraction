// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "./interfaces/IAAWallet.sol";
import "./interfaces/ISessionKeyManager.sol";

/**
 * @title SessionKeyManager
 * @notice Centralized management for session keys across multiple CPOP accounts
 * @dev Allows batch operations and template-based session key management
 */
contract SessionKeyManager is 
    Initializable, 
    ISessionKeyManager,
    OwnableUpgradeable, 
    UUPSUpgradeable 
{
    using ECDSA for bytes32;

    // Structs are defined in ISessionKeyManager interface

    
    // Master signer to accounts mapping (optional for batch operations)
    mapping(address => address[]) public masterSignerAccounts;
    mapping(address => bool) public authorizedMasters;
    
    // Session key templates
    mapping(string => SessionKeyTemplate) public sessionKeyTemplates;
    string[] public templateNames;
    
    // Events are defined in ISessionKeyManager interface

    modifier onlyAuthorizedMaster() {
        require(authorizedMasters[msg.sender], "SessionKeyManager: not authorized master");
        _;
    }

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize the session key manager
     * @param _owner Owner of the contract
     */
    function initialize(address _owner) external initializer {
        require(_owner != address(0), "SessionKeyManager: invalid owner");
        
        __Ownable_init(_owner);
        __UUPSUpgradeable_init();
        
        // Create default templates
        _createDefaultTemplates();
    }

    /**
     * @notice Authorize a master signer
     * @param masterSigner Master signer address to authorize
     */
    function authorizeMasterSigner(address masterSigner) external override onlyOwner {
        require(masterSigner != address(0), "SessionKeyManager: invalid master signer");
        require(!authorizedMasters[masterSigner], "SessionKeyManager: already authorized");
        
        authorizedMasters[masterSigner] = true;
        emit MasterSignerAuthorized(masterSigner, true);
    }

    /**
     * @notice Revoke master signer authorization
     * @param masterSigner Master signer address to revoke
     */
    function revokeMasterSigner(address masterSigner) external override onlyOwner {
        require(authorizedMasters[masterSigner], "SessionKeyManager: not authorized");
        
        authorizedMasters[masterSigner] = false;
        emit MasterSignerAuthorized(masterSigner, false);
    }

    /**
     * @notice Register an account under a master signer
     * @param masterSigner Master signer address
     * @param account Account address to register
     */
    function registerAccount(address masterSigner, address account) external override {
        require(authorizedMasters[masterSigner], "SessionKeyManager: master not authorized");
        require(account != address(0), "SessionKeyManager: invalid account");
        
        // Verify the account is actually controlled by the master signer
        try IAAWallet(account).getMasterSigner() returns (address accountMaster) {
            require(accountMaster == masterSigner, "SessionKeyManager: master signer mismatch");
        } catch {
            revert("SessionKeyManager: account verification failed");
        }
        
        // Check if already registered
        address[] storage accounts = masterSignerAccounts[masterSigner];
        for (uint256 i = 0; i < accounts.length; i++) {
            require(accounts[i] != account, "SessionKeyManager: account already registered");
        }
        
        accounts.push(account);
        emit AccountRegistered(masterSigner, account);
    }

    /**
     * @notice Unregister an account from a master signer
     * @param masterSigner Master signer address
     * @param account Account address to unregister
     */
    function unregisterAccount(address masterSigner, address account) external override onlyAuthorizedMaster {
        require(msg.sender == masterSigner || msg.sender == owner(), "SessionKeyManager: unauthorized");
        
        address[] storage accounts = masterSignerAccounts[masterSigner];
        for (uint256 i = 0; i < accounts.length; i++) {
            if (accounts[i] == account) {
                // Move last element to current position and pop
                accounts[i] = accounts[accounts.length - 1];
                accounts.pop();
                emit AccountUnregistered(masterSigner, account);
                return;
            }
        }
        revert("SessionKeyManager: account not found");
    }

    /**
     * @notice Create a session key template
     * @param name Template name
     * @param defaultDuration Default duration in seconds
     * @param permissions Default permissions
     */
    function createSessionKeyTemplate(
        string calldata name,
        uint48 defaultDuration,
        bytes32 permissions
    ) external override onlyOwner {
        require(bytes(name).length > 0, "SessionKeyManager: invalid template name");
        require(defaultDuration > 0, "SessionKeyManager: invalid duration");
        require(!sessionKeyTemplates[name].isActive, "SessionKeyManager: template already exists");
        
        sessionKeyTemplates[name] = SessionKeyTemplate({
            name: name,
            defaultDuration: defaultDuration,
            permissions: permissions,
            isActive: true
        });
        
        templateNames.push(name);
        emit SessionKeyTemplateCreated(name, defaultDuration, permissions);
    }

    /**
     * @notice Add session keys to multiple accounts using a template
     * @param masterSigner Master signer (must be authorized)
     * @param sessionKey Session key address
     * @param templateName Template to use
     * @param customDuration Custom duration (0 to use template default)
     */
    function addSessionKeyWithTemplate(
        address masterSigner,
        address sessionKey,
        string calldata templateName,
        uint48 customDuration
    ) external override onlyAuthorizedMaster {
        require(msg.sender == masterSigner, "SessionKeyManager: can only manage own accounts");
        require(sessionKeyTemplates[templateName].isActive, "SessionKeyManager: template not found");
        
        SessionKeyTemplate memory template = sessionKeyTemplates[templateName];
        uint48 duration = customDuration > 0 ? customDuration : template.defaultDuration;
        uint48 validAfter = uint48(block.timestamp);
        uint48 validUntil = validAfter + duration;
        
        // Add session key to all registered accounts
        address[] memory accounts = masterSignerAccounts[masterSigner];
        for (uint256 i = 0; i < accounts.length; i++) {
            try IAAWallet(accounts[i]).addSessionKey(
                sessionKey,
                validAfter,
                validUntil,
                template.permissions
            ) {
                // Success - continue to next account
            } catch {
                // Log error but continue with other accounts
            }
        }
        
        emit BatchSessionKeysAdded(masterSigner, accounts.length);
    }

    /**
     * @notice Batch add session keys to multiple accounts
     * @param operations Array of batch operations
     */
    function batchAddSessionKeys(
        BatchSessionKeyOp[] calldata operations
    ) external override onlyAuthorizedMaster {
        require(operations.length > 0, "SessionKeyManager: no operations");
        
        address masterSigner = msg.sender;
        uint256 successCount = 0;
        
        for (uint256 i = 0; i < operations.length; i++) {
            BatchSessionKeyOp memory op = operations[i];
            
            // Verify account belongs to master signer
            if (!_isAccountRegistered(masterSigner, op.account)) {
                continue;
            }
            
            try IAAWallet(op.account).addSessionKey(
                op.sessionKey,
                op.validAfter,
                op.validUntil,
                op.permissions
            ) {
                successCount++;
            } catch {
                // Continue with next operation
            }
        }
        
        emit BatchSessionKeysAdded(masterSigner, successCount);
    }

    /**
     * @notice Batch revoke session keys from multiple accounts
     * @param sessionKey Session key to revoke
     */
    function batchRevokeSessionKey(address sessionKey) external override onlyAuthorizedMaster {
        address masterSigner = msg.sender;
        address[] memory accounts = masterSignerAccounts[masterSigner];
        uint256 successCount = 0;
        
        for (uint256 i = 0; i < accounts.length; i++) {
            try IAAWallet(accounts[i]).revokeSessionKey(sessionKey) {
                successCount++;
            } catch {
                // Continue with next account
            }
        }
        
        emit BatchSessionKeysRevoked(masterSigner, successCount);
    }

    /**
     * @notice Get accounts registered under a master signer
     * @param masterSigner Master signer address
     * @return accounts Array of registered account addresses
     */
    function getRegisteredAccounts(address masterSigner) external view override returns (address[] memory accounts) {
        return masterSignerAccounts[masterSigner];
    }

    /**
     * @notice Get all template names
     * @return names Array of template names
     */
    function getTemplateNames() external view override returns (string[] memory names) {
        return templateNames;
    }

    /**
     * @notice Get template information
     * @param templateName Template name
     * @return template Template data
     */
    function getTemplate(string calldata templateName) external view override returns (SessionKeyTemplate memory template) {
        return sessionKeyTemplates[templateName];
    }



    /**
     * @notice Check if account is registered under master signer
     */
    function _isAccountRegistered(address masterSigner, address account) internal view returns (bool registered) {
        address[] memory accounts = masterSignerAccounts[masterSigner];
        for (uint256 i = 0; i < accounts.length; i++) {
            if (accounts[i] == account) {
                return true;
            }
        }
        return false;
    }

    /**
     * @notice Create default session key templates
     */
    function _createDefaultTemplates() internal {
        // DApp interaction template (24 hours)
        sessionKeyTemplates["DAPP_INTERACTION"] = SessionKeyTemplate({
            name: "DAPP_INTERACTION",
            defaultDuration: 24 * 60 * 60, // 24 hours
            permissions: bytes32(0), // Allow all operations
            isActive: true
        });
        templateNames.push("DAPP_INTERACTION");
        
        // Trading template (1 hour)
        sessionKeyTemplates["TRADING"] = SessionKeyTemplate({
            name: "TRADING",
            defaultDuration: 60 * 60, // 1 hour
            permissions: keccak256("TRADING_PERMISSIONS"), // Specific trading permissions
            isActive: true
        });
        templateNames.push("TRADING");
        
        // Game interaction template (7 days)
        sessionKeyTemplates["GAME_INTERACTION"] = SessionKeyTemplate({
            name: "GAME_INTERACTION", 
            defaultDuration: 7 * 24 * 60 * 60, // 7 days
            permissions: keccak256("GAME_PERMISSIONS"), // Game-specific permissions
            isActive: true
        });
        templateNames.push("GAME_INTERACTION");
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
     * @notice Authorize contract upgrades
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {
        // Additional upgrade logic can be added here if needed
    }
}