// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title ICPOPToken
 * @notice Interface for CPOP token - a gas-optimized token for internal circulation
 * @dev Only allows transfers between authorized addresses (AAWallet and system contracts)
 */
interface ICPOPToken {
    /**
     * @notice Emitted when tokens are minted
     * @param to The address that received the tokens
     * @param amount The amount of tokens minted
     */
    event Mint(address indexed to, uint256 amount);

    /**
     * @notice Emitted when tokens are burned
     * @param from The address that tokens were burned from
     * @param amount The amount of tokens burned
     */
    event Burn(address indexed from, uint256 amount);

    /**
     * @notice Emitted when tokens are transferred
     * @param from The address that sent the tokens
     * @param to The address that received the tokens
     * @param amount The amount of tokens transferred
     */
    event Transfer(address indexed from, address indexed to, uint256 amount);

    /**
     * @notice Emitted when an address is added to whitelist
     * @param account The address added to whitelist
     */
    event WhitelistAdded(address indexed account);

    /**
     * @notice Emitted when an address is removed from whitelist
     * @param account The address removed from whitelist
     */
    event WhitelistRemoved(address indexed account);

    /**
     * @notice Get the total supply of tokens
     * @return The total supply
     */
    function totalSupply() external view returns (uint256);

    /**
     * @notice Get the balance of an account
     * @param account The address to query
     * @return The token balance
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @notice Transfer tokens to another address
     * @dev Only works between whitelisted addresses
     * @param to The recipient address
     * @param amount The amount to transfer
     * @return success True if transfer succeeded
     */
    function transfer(address to, uint256 amount) external returns (bool);

    /**
     * @notice Mint new tokens
     * @dev Only callable by accounts with MINTER_ROLE
     * @param to The address to mint tokens to
     * @param amount The amount to mint
     */
    function mint(address to, uint256 amount) external;

    /**
     * @notice Burn tokens from an address
     * @dev Only callable by accounts with BURNER_ROLE
     * @param from The address to burn tokens from
     * @param amount The amount to burn
     */
    function burn(address from, uint256 amount) external;

    /**
     * @notice Check if an address is whitelisted for transfers
     * @param account The address to check
     * @return True if the address is whitelisted
     */
    function isWhitelisted(address account) external view returns (bool);

    /**
     * @notice Check if a transfer between two addresses is authorized
     * @param from The sender address
     * @param to The recipient address
     * @return True if the transfer is authorized
     */
    function isAuthorizedTransfer(address from, address to) external view returns (bool);

    /**
     * @notice Add an address to the transfer whitelist
     * @dev Only callable by accounts with ADMIN_ROLE
     * @param account The address to add
     */
    function addToWhitelist(address account) external;

    /**
     * @notice Remove an address from the transfer whitelist
     * @dev Only callable by accounts with ADMIN_ROLE
     * @param account The address to remove
     */
    function removeFromWhitelist(address account) external;

    /**
     * @notice Check if an address is a CPOPAccount contract
     * @param account The address to check
     * @return True if the address is a CPOPAccount contract
     */
    function isCPOPAccount(address account) external view returns (bool);

    /**
     * @notice Mint tokens to multiple addresses in batch
     * @dev Only callable by accounts with MINTER_ROLE
     * @param recipients Array of addresses to mint tokens to
     * @param amounts Array of amounts to mint to each address
     */
    function batchMint(address[] calldata recipients, uint256[] calldata amounts) external;

    /**
     * @notice Burn tokens from multiple addresses in batch
     * @dev Only callable by accounts with BURNER_ROLE
     * @param accounts Array of addresses to burn tokens from
     * @param amounts Array of amounts to burn from each address
     */
    function batchBurn(address[] calldata accounts, uint256[] calldata amounts) external;
}