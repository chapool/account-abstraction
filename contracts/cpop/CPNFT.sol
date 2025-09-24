// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

/**
 * @title CPNFT Contract - Standard Implementation
 * @dev Standard ERC721 with batch operations and owner privileges
 * @dev Owner can mint/burn/transfer without authorization in batch operations
 */
contract CPNFT is 
    Initializable,
    ERC721Upgradeable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    // ============================================
    // STORAGE VARIABLES
    // ============================================
    
    // Token metadata
    string private _baseTokenURI;
    
    // Token management
    uint256 private _tokenIdCounter;
    
    // NFT level and staking
    mapping(uint256 => NFTLevel) private _tokenLevels;
    mapping(uint256 => bool) private _isStaked;
    
    // Staking contract address
    address public stakingContract;

    // ============================================
    // ENUMS AND EVENTS
    // ============================================
    
    enum NFTLevel { NORMAL, C, B, A, S, SS, SSS }
    
    event TokenLevelSet(uint256 indexed tokenId, NFTLevel level);
    event TokenStakeStatusChanged(uint256 indexed tokenId, bool isStaked);
    event StakingContractSet(address indexed stakingContract);

    // ============================================
    // INITIALIZER
    // ============================================
    
    /**
     * @dev Initialize function (replaces constructor for upgradeable contracts)
     * @param name_ Name of the NFT collection
     * @param symbol_ Symbol of the NFT collection
     * @param baseTokenURI_ Base URI for token metadata
     */
    function initialize(
        string memory name_,
        string memory symbol_,
        string memory baseTokenURI_
    ) public initializer {
        __ERC721_init(name_, symbol_);
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
        
        _baseTokenURI = baseTokenURI_;
        _tokenIdCounter = 1; // Start token IDs from 1
    }

    // ============================================
    // UUPS UPGRADE AUTHORIZATION
    // ============================================
    
    /**
     * @dev Authorize upgrade (only owner can upgrade)
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {
        // Upgrade authorization logic can be added here if needed
        // For now, only owner can upgrade
    }


    // ============================================
    // METADATA FUNCTIONS (Standard Implementation)
    // ============================================
    
    /**
     * @dev Returns the Uniform Resource Identifier (URI) for a token
     * @param tokenId The token ID to query
     */
    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        _requireOwned(tokenId);
        return string(abi.encodePacked(_baseTokenURI, Strings.toString(tokenId)));
    }

    /**
     * @dev Set the base URI for token metadata
     * @param baseURI New base URI
     */
    function setBaseURI(string memory baseURI) external onlyOwner {
        _baseTokenURI = baseURI;
    }



    // ============================================
    // ERC721 TRANSFER FUNCTIONS (Override with staking check)
    // ============================================
    
    /**
     * @dev Transfers a token from one address to another
     * @param from The current owner of the token
     * @param to The address to receive the token
     * @param tokenId The token ID to transfer
     */
    function transferFrom(address from, address to, uint256 tokenId) public override {
        require(!_isStaked[tokenId], "Cannot transfer staked token");
        super.transferFrom(from, to, tokenId);
    }

    /**
     * @dev Safely transfers a token from one address to another with additional data
     * @param from The current owner of the token
     * @param to The address to receive the token
     * @param tokenId The token ID to transfer
     * @param data Additional data to send with the transfer
     */
    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory data) public override {
        require(!_isStaked[tokenId], "Cannot transfer staked token");
        super.safeTransferFrom(from, to, tokenId, data);
    }

    // ============================================
    // MINTING FUNCTIONS
    // ============================================
    
    /**
     * @dev Mint NFTs in batch to multiple addresses with specified levels
     * @param to Array of addresses to receive the minted NFTs
     * @param levels Array of NFT levels corresponding to each address
     */
    function batchMint(address[] calldata to, NFTLevel[] calldata levels) external onlyOwner() {
        require(to.length == levels.length, "Arrays length mismatch");
        
        for (uint256 i = 0; i < to.length; i++) {
            mint(to[i], levels[i]);
        }
    }
    

    /**
     * @dev Mint a single NFT to an address with specified level
     * @param to Address to receive the minted NFT
     * @param level NFT level
     * @return tokenId The ID of the minted token
     */
    function mint(address to, NFTLevel level) public onlyOwner returns (uint256) {
        uint256 tokenId = _tokenIdCounter;
        _tokenIdCounter++;
        
        _safeMint(to, tokenId);
        _tokenLevels[tokenId] = level;
        
        emit TokenLevelSet(tokenId, level);
        return tokenId;
    }
    

    // ============================================
    // BURNING FUNCTIONS
    // ============================================
    
    /**
     * @dev Burn NFTs in batch
     * @param tokenIds Array of token IDs to be burned
     */
    function batchBurn(uint256[] calldata tokenIds) external onlyOwner {
        for (uint256 i = 0; i < tokenIds.length; i++) {
            burn(tokenIds[i]);
        }
    }

    /**
     * @dev Burn a single NFT
     * @param tokenId ID of the token to be burned
     */
    function burn(uint256 tokenId) public onlyOwner {
        _requireOwned(tokenId);
        _burn(tokenId);
    }

    // ============================================
    // BATCH TRANSFER FUNCTIONS
    // ============================================
    
    /**
     * @dev Transfer NFTs in batch from multiple owners to multiple recipients
     * @param from Array of current token owners
     * @param to Array of new token owners
     * @param tokenIds Array of token IDs to be transferred
     * @dev Owner can transfer without authorization
     */
    function batchTransferFrom(
        address[] calldata from,
        address[] calldata to,
        uint256[] calldata tokenIds
    ) external onlyOwner {
        require(
            from.length == to.length && to.length == tokenIds.length,
            "Arrays length mismatch"
        );
        
        for (uint256 i = 0; i < tokenIds.length; i++) {
            require(!_isStaked[tokenIds[i]], "Cannot transfer staked token");
            _transfer(from[i], to[i], tokenIds[i]);
        }
    }


    // ============================================
    // NFT LEVEL FUNCTIONS
    // ============================================
    
    /**
     * @dev Set NFT level (only contract owner)
     * @param tokenId Token ID
     * @param level NFT level
     */
    function setTokenLevel(uint256 tokenId, NFTLevel level) external onlyOwner {
        _requireOwned(tokenId);
        require(!_isStaked[tokenId], "Cannot change level of staked token");
        
        _tokenLevels[tokenId] = level;
        emit TokenLevelSet(tokenId, level);
    }

    /**
     * @dev Get NFT level
     * @param tokenId Token ID
     * @return NFT level
     */
    function getTokenLevel(uint256 tokenId) public view returns (NFTLevel) {
        _requireOwned(tokenId);
        return _tokenLevels[tokenId];
    }

    // ============================================
    // STAKING FUNCTIONS
    // ============================================
    
    /**
     * @dev Check if NFT is staked
     * @param tokenId Token ID
     * @return Whether the NFT is staked
     */
    function isStaked(uint256 tokenId) public view returns (bool) {
        _requireOwned(tokenId);
        return _isStaked[tokenId];
    }

    /**
     * @dev Set staking contract address (only contract owner)
     * @param _stakingContract Staking contract address
     */
    function setStakingContract(address _stakingContract) external onlyOwner {
        stakingContract = _stakingContract;
        emit StakingContractSet(_stakingContract);
    }

    /**
     * @dev Set staking status (only callable by staking contract)
     * @param tokenId Token ID
     * @param staked Staking status
     */
    function setStakeStatus(uint256 tokenId, bool staked) external {
        require(msg.sender == stakingContract, "Only staking contract can call");
        _isStaked[tokenId] = staked;
        emit TokenStakeStatusChanged(tokenId, staked);
    }
    



    // ============================================
    // VERSIONING
    // ============================================
    
    /**
     * @dev Returns the version of the contract
     */
    function version() public pure returns (string memory) {
        return "1.0.0";
    }
}