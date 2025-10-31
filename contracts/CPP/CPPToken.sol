// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "./interfaces/ICPPToken.sol";

/**
 * @title CPPToken - Gas Optimized Role-Based ERC20
 * @notice Ultra gas-efficient ERC20 implementation with lightweight role system
 * @dev Uses bit-packed roles for minimal gas overhead while supporting multiple contracts
 */
contract CPPToken is ICPPToken {
    // Role constants (bit flags for gas efficiency)
    uint8 public constant ADMIN_ROLE = 1;     // 0001 - Can manage roles
    uint8 public constant MINTER_ROLE = 2;    // 0010 - Can mint tokens
    uint8 public constant BURNER_ROLE = 4;    // 0100 - Can burn tokens from any address
    
    // ERC20 basic storage - packed into single slot where possible
    string public constant NAME = "CPP Token";
    string public constant SYMBOL = "CPP";
    uint8 public constant DECIMALS = 18;
    
    uint256 public totalSupply;
    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;
    
    // Role management - gas efficient bit-packed roles
    mapping(address => uint8) public roles;
    
    constructor(address _admin, uint256 _initialSupply) {
        // Grant all roles to initial admin
        roles[_admin] = ADMIN_ROLE | MINTER_ROLE | BURNER_ROLE;
        
        totalSupply = _initialSupply;
        balanceOf[_admin] = _initialSupply;
        
        emit Transfer(address(0), _admin, _initialSupply);
        emit RoleGranted(_admin, ADMIN_ROLE);
        emit RoleGranted(_admin, MINTER_ROLE);
        emit RoleGranted(_admin, BURNER_ROLE);
    }
    
    /**
     * @notice Check if an address has a specific role
     * @dev Uses bitwise operations for gas efficiency
     */
    function hasRole(address account, uint8 role) public view returns (bool) {
        return (roles[account] & role) != 0;
    }
    
    /**
     * @notice Grant a role to an address
     * @dev Only ADMIN_ROLE can grant roles
     */
    function grantRole(address account, uint8 role) external {
        if (!hasRole(msg.sender, ADMIN_ROLE)) revert AccessDenied();
        if (role == 0 || role > 7) revert InvalidRole(); // Valid roles: 1, 2, 4 and combinations
        
        roles[account] |= role;
        emit RoleGranted(account, role);
    }
    
    /**
     * @notice Revoke a role from an address
     * @dev Only ADMIN_ROLE can revoke roles
     */
    function revokeRole(address account, uint8 role) external {
        if (!hasRole(msg.sender, ADMIN_ROLE)) revert AccessDenied();
        if (role == 0 || role > 7) revert InvalidRole();
        
        roles[account] &= ~role;
        emit RoleRevoked(account, role);
    }
    
    /**
     * @notice Transfer tokens
     * @dev Gas optimized - minimal checks
     */
    function transfer(address to, uint256 amount) external returns (bool) {
        balanceOf[msg.sender] -= amount; // Will revert on underflow (Solidity 0.8+)
        
        // Unchecked block for gas optimization - safe since we checked underflow above
        unchecked {
            balanceOf[to] += amount;
        }
        
        emit Transfer(msg.sender, to, amount);
        return true;
    }
    
    /**
     * @notice Approve spender
     * @dev Gas optimized approval
     */
    function approve(address spender, uint256 amount) external returns (bool) {
        allowance[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }
    
    /**
     * @notice Transfer from approved amount
     * @dev Gas optimized transferFrom
     */
    function transferFrom(address from, address to, uint256 amount) external returns (bool) {
        uint256 allowed = allowance[from][msg.sender];
        
        // Check allowance (will revert if insufficient)
        if (allowed != type(uint256).max) {
            allowance[from][msg.sender] = allowed - amount;
        }
        
        // Transfer tokens
        balanceOf[from] -= amount; // Will revert on underflow
        
        unchecked {
            balanceOf[to] += amount;
        }
        
        emit Transfer(from, to, amount);
        return true;
    }
    
    /**
     * @notice Mint new tokens (MINTER_ROLE required)
     * @dev Can be called by any address with MINTER_ROLE
     */
    function mint(address to, uint256 amount) external {
        if (!hasRole(msg.sender, MINTER_ROLE)) revert AccessDenied();
        
        unchecked {
            totalSupply += amount;
            balanceOf[to] += amount;
        }
        
        emit Transfer(address(0), to, amount);
    }
    
    /**
     * @notice Burn tokens from caller's balance
     * @dev Gas optimized burn
     */
    function burn(uint256 amount) external {
        balanceOf[msg.sender] -= amount; // Will revert on underflow
        
        unchecked {
            totalSupply -= amount;
        }
        
        emit Transfer(msg.sender, address(0), amount);
    }
    
    /**
     * @notice Burn tokens from specific address (with approval or BURNER_ROLE)
     * @dev BURNER_ROLE can burn from any address, others need approval
     */
    function burnFrom(address from, uint256 amount) external {
        // BURNER_ROLE can burn from any address without approval
        if (!hasRole(msg.sender, BURNER_ROLE)) {
            uint256 allowed = allowance[from][msg.sender];
            if (allowed != type(uint256).max) {
                allowance[from][msg.sender] = allowed - amount;
            }
        }
        
        balanceOf[from] -= amount; // Will revert on underflow
        
        unchecked {
            totalSupply -= amount;
        }
        
        emit Transfer(from, address(0), amount);
    }
    
    /**
     * @notice Admin burn - burn tokens from any address without approval
     * @dev Only BURNER_ROLE can call this function
     */
    function adminBurn(address from, uint256 amount) external {
        if (!hasRole(msg.sender, BURNER_ROLE)) revert AccessDenied();
        
        balanceOf[from] -= amount; // Will revert on underflow
        
        unchecked {
            totalSupply -= amount;
        }
        
        emit Transfer(from, address(0), amount);
    }
    
    // ERC20 metadata functions for compatibility
    function name() external pure returns (string memory) {
        return NAME;
    }
    
    function symbol() external pure returns (string memory) {
        return SYMBOL;
    }
    
    function decimals() external pure returns (uint8) {
        return DECIMALS;
    }
    
    /**
     * @notice Mint tokens to multiple addresses in batch
     * @dev Gas-optimized batch operation for minting - only MINTER_ROLE
     */
    function batchMint(address[] calldata recipients, uint256[] calldata amounts) external {
        if (!hasRole(msg.sender, MINTER_ROLE)) revert AccessDenied();
        if (recipients.length != amounts.length) revert ArrayLengthMismatch();
        if (recipients.length == 0) revert EmptyArray();
        
        unchecked {
            for (uint256 i = 0; i < recipients.length; i++) {
                address to = recipients[i];
                uint256 amount = amounts[i];
                
                totalSupply += amount;
                balanceOf[to] += amount;
                
                emit Transfer(address(0), to, amount);
            }
        }
    }
    
    /**
     * @notice Burn tokens from multiple addresses in batch  
     * @dev Gas-optimized batch operation for burning - only BURNER_ROLE
     */
    function batchBurn(address[] calldata accounts, uint256[] calldata amounts) external {
        if (!hasRole(msg.sender, BURNER_ROLE)) revert AccessDenied();
        if (accounts.length != amounts.length) revert ArrayLengthMismatch();
        if (accounts.length == 0) revert EmptyArray();
        
        for (uint256 i = 0; i < accounts.length; i++) {
            address from = accounts[i];
            uint256 amount = amounts[i];
            
            balanceOf[from] -= amount; // Will revert on underflow (checked math)
            
            unchecked {
                totalSupply -= amount; // Safe since we checked balance above
            }
            
            emit Transfer(from, address(0), amount);
        }
    }
    
    /**
     * @notice Transfer tokens to multiple addresses in batch
     * @dev Gas-optimized batch operation for transfers from caller
     */
    function batchTransfer(address[] calldata recipients, uint256[] calldata amounts) external {
        if (recipients.length != amounts.length) revert ArrayLengthMismatch();
        if (recipients.length == 0) revert EmptyArray();
        
        uint256 totalAmount = 0;
        
        // Calculate total amount to transfer
        unchecked {
            for (uint256 i = 0; i < amounts.length; i++) {
                totalAmount += amounts[i];
            }
        }
        
        // Check sender has sufficient balance (will revert on underflow)
        balanceOf[msg.sender] -= totalAmount;
        
        // Execute transfers
        unchecked {
            for (uint256 i = 0; i < recipients.length; i++) {
                address to = recipients[i];
                uint256 amount = amounts[i];
                
                balanceOf[to] += amount;
                emit Transfer(msg.sender, to, amount);
            }
        }
    }
    
    /**
     * @notice Transfer tokens between multiple address pairs in batch
     * @dev Gas-optimized batch operation - ADMIN_ROLE bypasses allowance checks
     */
    function batchTransferFrom(
        address[] calldata from,
        address[] calldata to, 
        uint256[] calldata amounts
    ) external {
        if (from.length != to.length || from.length != amounts.length) revert ArrayLengthMismatch();
        if (from.length == 0) revert EmptyArray();
        
        bool isAdmin = hasRole(msg.sender, ADMIN_ROLE);
        
        for (uint256 i = 0; i < from.length; i++) {
            address fromAddr = from[i];
            address toAddr = to[i];
            uint256 amount = amounts[i];
            
            // Admin role bypasses allowance checks
            if (!isAdmin) {
                uint256 allowed = allowance[fromAddr][msg.sender];
                if (allowed != type(uint256).max) {
                    allowance[fromAddr][msg.sender] = allowed - amount; // Will revert on underflow
                }
            }
            
            // Transfer tokens
            balanceOf[fromAddr] -= amount; // Will revert on underflow
            
            unchecked {
                balanceOf[toAddr] += amount;
            }
            
            emit Transfer(fromAddr, toAddr, amount);
        }
    }
}