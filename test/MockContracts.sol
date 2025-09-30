// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title MockCPOPToken
 * @dev Mock CPOP token for testing
 */
contract MockCPOPToken is ERC20 {
    constructor() ERC20("Mock CPOP", "MCPOP") {
        _mint(msg.sender, 1000000 * 10**18); // Mint 1M tokens to deployer
    }
    
    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

/**
 * @title MockAccountManager
 * @dev Mock account manager for testing
 */
contract MockAccountManager {
    function getDefaultMasterSigner() external pure returns (address) {
        return address(0x1234567890123456789012345678901234567890); // Mock address
    }
    
    function getAccountAddress(address user, address masterSigner) external pure returns (address) {
        // Return a deterministic address based on user and master signer
        return address(uint160(uint256(keccak256(abi.encodePacked(user, masterSigner)))));
    }
}
