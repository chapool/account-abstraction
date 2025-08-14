// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "./interfaces/ICPOPToken.sol";
import "./interfaces/IAAWallet.sol";

/**
 * @title CPOPToken
 * @notice A gas-optimized token for CPOP ecosystem internal circulation
 * @dev Non-standard ERC20 - only allows transfers between whitelisted addresses
 * Optimized for gas efficiency by removing unnecessary features like allowances
 */
contract CPOPToken is ICPOPToken, AccessControl, Pausable {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant BURNER_ROLE = keccak256("BURNER_ROLE");
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    string public constant name = "CPOP Token";
    string public constant symbol = "CPOP";
    uint8 public constant decimals = 18;

    uint256 private _totalSupply;
    mapping(address => uint256) private _balances;
    mapping(address => bool) private _whitelist;

    // Removed onlyWhitelisted modifier - now using smart interface detection only

    modifier validTransfer(address from, address to) {
        require(isAuthorizedTransfer(from, to), "CPOPToken: unauthorized transfer");
        _;
    }

    constructor(address admin) {
        require(admin != address(0), "CPOPToken: admin cannot be zero address");
        
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
        _grantRole(MINTER_ROLE, admin);
        _grantRole(BURNER_ROLE, admin);
    }

    /**
     * @notice Get the total supply of tokens
     */
    function totalSupply() external view override returns (uint256) {
        return _totalSupply;
    }

    /**
     * @notice Get the balance of an account
     */
    function balanceOf(address account) external view override returns (uint256) {
        return _balances[account];
    }

    /**
     * @notice Transfer tokens between whitelisted addresses
     */
    function transfer(address to, uint256 amount) 
        external 
        override 
        whenNotPaused 
        validTransfer(msg.sender, to)
        returns (bool) 
    {
        _transfer(msg.sender, to, amount);
        return true;
    }

    /**
     * @notice Mint new tokens to any address
     */
    function mint(address to, uint256 amount) 
        external 
        override 
        onlyRole(MINTER_ROLE) 
        whenNotPaused 
    {
        require(to != address(0), "CPOPToken: mint to zero address");
        require(amount > 0, "CPOPToken: mint amount must be positive");

        _totalSupply += amount;
        _balances[to] += amount;

        emit Mint(to, amount);
        emit Transfer(address(0), to, amount);
    }

    /**
     * @notice Burn tokens from any address
     */
    function burn(address from, uint256 amount) 
        external 
        override 
        onlyRole(BURNER_ROLE) 
        whenNotPaused 
    {
        require(from != address(0), "CPOPToken: burn from zero address");
        require(amount > 0, "CPOPToken: burn amount must be positive");
        require(_balances[from] >= amount, "CPOPToken: burn amount exceeds balance");

        _balances[from] -= amount;
        _totalSupply -= amount;

        emit Burn(from, amount);
        emit Transfer(from, address(0), amount);
    }

    /**
     * @notice Mint tokens to multiple addresses in batch
     * @dev Gas-optimized batch operation for minting
     * @param recipients Array of addresses to mint tokens to
     * @param amounts Array of amounts to mint to each address
     */
    function batchMint(address[] calldata recipients, uint256[] calldata amounts) 
        external 
        override 
        onlyRole(MINTER_ROLE) 
        whenNotPaused 
    {
        require(recipients.length == amounts.length, "CPOPToken: arrays length mismatch");
        require(recipients.length > 0, "CPOPToken: empty arrays");
        
        for (uint256 i = 0; i < recipients.length; i++) {
            address to = recipients[i];
            uint256 amount = amounts[i];
            
            require(to != address(0), "CPOPToken: mint to zero address");
            require(amount > 0, "CPOPToken: mint amount must be positive");
            
            _totalSupply += amount;
            _balances[to] += amount;
            
            emit Mint(to, amount);
            emit Transfer(address(0), to, amount);
        }
    }

    /**
     * @notice Burn tokens from multiple addresses in batch
     * @dev Gas-optimized batch operation for burning
     * @param accounts Array of addresses to burn tokens from
     * @param amounts Array of amounts to burn from each address
     */
    function batchBurn(address[] calldata accounts, uint256[] calldata amounts) 
        external 
        override 
        onlyRole(BURNER_ROLE) 
        whenNotPaused 
    {
        require(accounts.length == amounts.length, "CPOPToken: arrays length mismatch");
        require(accounts.length > 0, "CPOPToken: empty arrays");
        
        for (uint256 i = 0; i < accounts.length; i++) {
            address from = accounts[i];
            uint256 amount = amounts[i];
            
            require(from != address(0), "CPOPToken: burn from zero address");
            require(amount > 0, "CPOPToken: burn amount must be positive");
            require(_balances[from] >= amount, "CPOPToken: burn amount exceeds balance");
            
            _balances[from] -= amount;
            _totalSupply -= amount;
            
            emit Burn(from, amount);
            emit Transfer(from, address(0), amount);
        }
    }

    /**
     * @notice Check if an address is whitelisted
     */
    function isWhitelisted(address account) external view override returns (bool) {
        return _whitelist[account];
    }

    /**
     * @notice Check if a transfer is authorized between two addresses
     * @dev Only allows transfers between system contracts and AAWallet contracts
     */
    function isAuthorizedTransfer(address from, address to) public view override returns (bool) {
        // Allow transfers if both addresses are whitelisted (for system contracts)
        if (_whitelist[from] && _whitelist[to]) {
            return true;
        }
        
        // Allow transfers involving AAWallet contracts
        bool fromIsAAWallet = isAAWallet(from);
        bool toIsAAWallet = isAAWallet(to);
        
        if (fromIsAAWallet || toIsAAWallet) {
            // Allow transfers between AAWallets
            if (fromIsAAWallet && toIsAAWallet) {
                return true;
            }
            // Allow transfers between AAWallet and whitelisted system contracts
            return (_whitelist[from] || fromIsAAWallet) && 
                   (_whitelist[to] || toIsAAWallet);
        }
        
        return false;
    }

    /**
     * @notice Add an address to the whitelist
     */
    function addToWhitelist(address account) external override onlyRole(ADMIN_ROLE) {
        require(account != address(0), "CPOPToken: cannot whitelist zero address");
        require(!_whitelist[account], "CPOPToken: address already whitelisted");

        _whitelist[account] = true;
        emit WhitelistAdded(account);
    }

    /**
     * @notice Remove an address from the whitelist
     */
    function removeFromWhitelist(address account) external override onlyRole(ADMIN_ROLE) {
        require(_whitelist[account], "CPOPToken: address not whitelisted");

        _whitelist[account] = false;
        emit WhitelistRemoved(account);
    }

    /**
     * @notice Check if an address is an AAWallet contract
     * @dev Uses interface detection to identify AAWallet contracts
     * @param account The address to check
     * @return True if the address is an AAWallet contract
     */
    function isAAWallet(address account) public view override returns (bool) {
        if (account.code.length == 0) {
            return false; // EOA cannot be AAWallet
        }
        
        try IAAWallet(account).supportsInterface(type(IAAWallet).interfaceId) returns (bool supported) {
            return supported;
        } catch {
            // Fallback: check if the contract has AAWallet-specific functions
            try IAAWallet(account).getOwner() returns (address) {
                try IAAWallet(account).getMasterSigner() returns (address) {
                    return true; // Has both functions, likely AAWallet
                } catch {
                    return false;
                }
            } catch {
                return false;
            }
        }
    }

    /**
     * @notice Internal transfer function
     */
    function _transfer(address from, address to, uint256 amount) internal {
        require(from != address(0), "CPOPToken: transfer from zero address");
        require(to != address(0), "CPOPToken: transfer to zero address");
        require(amount > 0, "CPOPToken: transfer amount must be positive");
        require(_balances[from] >= amount, "CPOPToken: transfer amount exceeds balance");

        _balances[from] -= amount;
        _balances[to] += amount;

        emit Transfer(from, to, amount);
    }

    /**
     * @notice Pause the contract (emergency function)
     */
    function pause() external onlyRole(ADMIN_ROLE) {
        _pause();
    }

    /**
     * @notice Unpause the contract
     */
    function unpause() external onlyRole(ADMIN_ROLE) {
        _unpause();
    }
}