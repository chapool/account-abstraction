// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title MockUSDT
 * @notice Mock USDT token for testing purposes
 * @dev Standard ERC20 with mint and burn functions
 */
contract MockUSDT is ERC20, Ownable {
    uint8 private _decimals;
    
    constructor() ERC20("Mock USDT", "mUSDT") Ownable(msg.sender) {
        _decimals = 6; // USDT has 6 decimals
        // Mint initial supply to deployer (1M USDT)
        _mint(msg.sender, 1_000_000 * 10**_decimals);
    }
    
    /**
     * @notice Override decimals to return 6 (like real USDT)
     */
    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }
    
    /**
     * @notice Mint tokens to any address (for testing)
     * @param to Recipient address
     * @param amount Amount to mint
     */
    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }
    
    /**
     * @notice Burn tokens from any address (for testing)
     * @param from Address to burn from
     * @param amount Amount to burn
     */
    function burn(address from, uint256 amount) external onlyOwner {
        _burn(from, amount);
    }
    
    /**
     * @notice Mint tokens to caller (for easy testing)
     * @param amount Amount to mint
     */
    function mintToSelf(uint256 amount) external {
        _mint(msg.sender, amount);
    }
    
    /**
     * @notice Faucet function - anyone can get 1000 USDT for testing
     */
    function faucet() external {
        uint256 faucetAmount = 1000 * 10**_decimals; // 1000 USDT
        _mint(msg.sender, faucetAmount);
    }
}