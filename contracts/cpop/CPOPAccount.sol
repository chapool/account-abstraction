// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import "../core/BaseAccount.sol";
import "../interfaces/IAccount.sol";
import "./interfaces/ICPOPAccount.sol";

/**
 * @title CPOPAccount
 * @notice Account Abstraction wallet optimized for Web2 users
 * @dev Inherits from BaseAccount and implements ICPOPAccount interface
 * Simplified design without token-specific functionality
 */
contract CPOPAccount is Initializable, BaseAccount, ICPOPAccount, UUPSUpgradeable, ERC165 {
    using ECDSA for bytes32;

    uint256 internal constant SIG_VALIDATION_FAILED = 1;
    uint256 internal constant SIG_VALIDATION_SUCCESS = 0;

    address public owner;
    address public masterSigner; // Master signer for Web2 users
    address public immutable entryPointAddress;

    modifier onlyOwner() {
        require(msg.sender == owner || msg.sender == address(this), "CPOPAccount: not owner");
        _;
    }

    modifier onlyMasterSigner() {
        require(msg.sender == masterSigner || msg.sender == address(this), "CPOPAccount: not master signer");
        _;
    }

    constructor(address _entryPoint) {
        require(_entryPoint != address(0), "CPOPAccount: invalid entryPoint");
        
        entryPointAddress = _entryPoint;
        _disableInitializers();
    }

    /**
     * @notice Initialize the account with an owner
     */
    function initialize(address _owner) external initializer {
        require(_owner != address(0), "CPOPAccount: invalid owner");
        owner = _owner;
        emit CPOPAccountInitialized(_owner, address(0));
    }

    /**
     * @notice Initialize the account with an owner and master signer for Web2 users
     */
    function initializeWithMasterSigner(address _owner, address _masterSigner) external initializer {
        require(_owner != address(0), "CPOPAccount: invalid owner");
        require(_masterSigner != address(0), "CPOPAccount: invalid master signer");
        owner = _owner;
        masterSigner = _masterSigner;
        emit CPOPAccountInitialized(_owner, _masterSigner);
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
     * @dev Supports both owner signature and master signer signature for Web2 users
     */
    function _validateSignature(
        PackedUserOperation calldata userOp,
        bytes32 userOpHash
    ) internal override returns (uint256 validationData) {
        bytes32 hash = MessageHashUtils.toEthSignedMessageHash(userOpHash);
        address recovered = ECDSA.recover(hash, userOp.signature);
        
        // First try owner signature (for Web3 users)
        if (recovered == owner) {
            return SIG_VALIDATION_SUCCESS;
        }
        
        // Then try master signer signature (for Web2 users)
        if (masterSigner != address(0) && recovered == masterSigner) {
            emit MasterSignerValidation(masterSigner, userOpHash);
            return SIG_VALIDATION_SUCCESS;
        }
        
        return SIG_VALIDATION_FAILED;
    }

    /**
     * @notice Check execution authorization
     * @dev Only EntryPoint can call execute functions for security
     */
    function _requireForExecute() internal view override {
        require(
            msg.sender == address(entryPoint()),
            "CPOPAccount: only EntryPoint can execute"
        );
    }

    /**
     * @notice Authorize contract upgrades
     */
    function _authorizeUpgrade(address newImplementation) internal view override {
        (newImplementation); // suppress unused parameter warning
        require(msg.sender == owner, "CPOPAccount: only owner can upgrade");
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

    /**
     * @notice Check interface support for ERC165
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC165, IERC165) returns (bool) {
        return interfaceId == type(ICPOPAccount).interfaceId || 
               interfaceId == type(IAccount).interfaceId ||
               super.supportsInterface(interfaceId);
    }

    /**
     * @notice Receive ETH
     */
    receive() external payable {}
}