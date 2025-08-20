// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title ICPOPToken
 * @notice Interface for gas-optimized role-based CPOP token
 * @dev Lightweight role-based access control for maximum gas efficiency
 */
interface ICPOPToken {
    // Custom errors
    error AccessDenied();
    error InvalidRole();
    error ArrayLengthMismatch();
    error EmptyArray();

    // Role constants (bit flags) - as public constants
    function ADMIN_ROLE() external view returns (uint8);
    function MINTER_ROLE() external view returns (uint8);
    function BURNER_ROLE() external view returns (uint8);

    // Standard ERC20 events
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    
    // Role events
    event RoleGranted(address indexed account, uint8 role);
    event RoleRevoked(address indexed account, uint8 role);

    // Standard ERC20 functions
    function name() external pure returns (string memory);
    function symbol() external pure returns (string memory);
    function decimals() external pure returns (uint8);
    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function transfer(address to, uint256 amount) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address from, address to, uint256 amount) external returns (bool);

    // Role management functions
    function hasRole(address account, uint8 role) external view returns (bool);
    function grantRole(address account, uint8 role) external;
    function revokeRole(address account, uint8 role) external;
    function roles(address account) external view returns (uint8);

    // Token management functions
    function mint(address to, uint256 amount) external;
    function burn(uint256 amount) external;
    function burnFrom(address from, uint256 amount) external;
    function adminBurn(address from, uint256 amount) external;
    
    // Batch operations for gas efficiency
    function batchMint(address[] calldata recipients, uint256[] calldata amounts) external;
    function batchBurn(address[] calldata accounts, uint256[] calldata amounts) external;
}