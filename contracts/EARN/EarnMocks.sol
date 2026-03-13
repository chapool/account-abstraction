// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title MockCPOT
 * @dev Mock CPOT token for Earn/VeCPOTLocker testing
 */
contract MockCPOT is ERC20 {
    constructor() ERC20("Mock CPOT", "MCPOT") {
        _mint(msg.sender, 10000000 * 10**18);
    }
    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

/**
 * @title EarnMockAccountManager
 * @dev Returns user as AA account so CPP mints to user (simplifies test assertions)
 */
contract EarnMockAccountManager {
    function getDefaultMasterSigner() external pure returns (address) {
        return address(0x1);
    }
    function getAccountAddress(address owner, address) external pure returns (address) {
        return owner;
    }
}

/**
 * @title MockCPNFTForEarn
 * @dev Minimal CPNFT for NFTBoostController tests: ownerOf, getTokenLevel, isStaked
 */
contract MockCPNFTForEarn {
    mapping(uint256 => address) private _owners;
    mapping(uint256 => uint8) private _levels;  // 0=NORMAL, 1=C, 2=B, 3=A, 4=S, 5=SS, 6=SSS
    mapping(uint256 => bool) private _staked;

    function mint(address to, uint256 tokenId, uint8 level) external {
        _owners[tokenId] = to;
        _levels[tokenId] = level;
    }
    function ownerOf(uint256 tokenId) external view returns (address) {
        return _owners[tokenId];
    }
    function getTokenLevel(uint256 tokenId) external view returns (uint8) {
        return _levels[tokenId];
    }
    function isStaked(uint256 tokenId) external view returns (bool) {
        return _staked[tokenId];
    }
    function setStakeStatus(uint256 tokenId, bool staked) external {
        _staked[tokenId] = staked;
    }
}
