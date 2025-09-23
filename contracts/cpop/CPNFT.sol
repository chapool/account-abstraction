// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/**
 * @title ERC721 token receiver interface
 * @dev Interface for any contract that wants to support safeTransfers
 * from ERC721 asset contracts.
 */
interface IERC721Receiver {
    /**
     * @dev Whenever an {IERC721} `tokenId` token is transferred to this contract via {IERC721-safeTransferFrom}
     * by `operator` from `from`, this function is called.
     *
     * It must return its Solidity selector to confirm the token transfer.
     * If any other value is returned or the interface is not implemented by the recipient, the transfer will be reverted.
     *
     * The selector can be obtained in Solidity with `IERC721Receiver.onERC721Received.selector`.
     */
    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external returns (bytes4);
}

/**
 * @title Gas-Optimized CPNFT Contract
 * @dev Custom implementation of NFT with batch operations and owner privileges
 * @dev Completely independent of OpenZeppelin libraries to minimize gas costs
 */
contract CPNFT {
    // ============================================
    // STORAGE VARIABLES
    // ============================================
    
    // Token metadata
    string private _name;
    string private _symbol;
    string private _baseTokenURI;
    
    // Token management
    uint256 private _tokenIdCounter;
    address private _owner;
    
    // ERC721 mappings
    mapping(uint256 => address) private _owners;
    mapping(address => uint256) private _balances;
    mapping(uint256 => address) private _tokenApprovals;
    mapping(address => mapping(address => bool)) private _operatorApprovals;
    mapping(uint256 => bool) private _mintedTokens;
    
    // NFT level and staking
    mapping(uint256 => NFTLevel) private _tokenLevels;
    mapping(uint256 => bool) private _isStaked;

    // ============================================
    // ENUMS AND EVENTS
    // ============================================
    
    enum NFTLevel { C, B, A, S, SS, SSS }
    
    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
    event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);
    event ApprovalForAll(address indexed owner, address indexed operator, bool approved);
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    event TokenLevelSet(uint256 indexed tokenId, NFTLevel level);
    event TokenStakeStatusChanged(uint256 indexed tokenId, bool isStaked);

    // ============================================
    // CONSTRUCTOR
    // ============================================
    
    /**
     * @dev Constructor function
     * @param name_ Name of the NFT collection
     * @param symbol_ Symbol of the NFT collection
     * @param baseTokenURI_ Base URI for token metadata
     */
    constructor(
        string memory name_,
        string memory symbol_,
        string memory baseTokenURI_
    ) {
        _name = name_;
        _symbol = symbol_;
        _baseTokenURI = baseTokenURI_;
        _owner = msg.sender;
        _tokenIdCounter = 1; // Start token IDs from 1
    }

    // ============================================
    // MODIFIERS
    // ============================================
    
    /**
     * @dev Throws if called by any account other than the owner
     */
    modifier onlyOwner() {
        require(_owner == msg.sender, "Caller is not the owner");
        _;
    }

    // ============================================
    // OWNERSHIP FUNCTIONS
    // ============================================
    
    /**
     * @dev Returns the address of the current owner
     */
    function owner() public view returns (address) {
        return _owner;
    }

    /**
     * @dev Transfers ownership of the contract to a new account
     * @param newOwner The address to transfer ownership to
     */
    function transferOwnership(address newOwner) public onlyOwner {
        require(newOwner != address(0), "New owner is the zero address");
        emit OwnershipTransferred(_owner, newOwner);
        _owner = newOwner;
    }

    // ============================================
    // METADATA FUNCTIONS
    // ============================================
    
    /**
     * @dev Returns the name of the token
     */
    function name() public view returns (string memory) {
        return _name;
    }

    /**
     * @dev Returns the symbol of the token
     */
    function symbol() public view returns (string memory) {
        return _symbol;
    }

    /**
     * @dev Returns the Uniform Resource Identifier (URI) for a token
     * @param tokenId The token ID to query
     */
    function tokenURI(uint256 tokenId) public view returns (string memory) {
        require(_mintedTokens[tokenId], "Token does not exist");
        return string(abi.encodePacked(_baseTokenURI, _toString(tokenId)));
    }

    /**
     * @dev Set the base URI for token metadata
     * @param baseURI New base URI
     */
    function setBaseURI(string memory baseURI) external onlyOwner {
        _baseTokenURI = baseURI;
    }

    // ============================================
    // ERC721 VIEW FUNCTIONS
    // ============================================
    
    /**
     * @dev Returns the number of tokens in an owner's account
     * @param owner_ The address to query the balance of
     */
    function balanceOf(address owner_) public view returns (uint256) {
        require(owner_ != address(0), "Balance query for the zero address");
        return _balances[owner_];
    }

    /**
     * @dev Returns the owner of a token
     * @param tokenId The token ID to query
     */
    function ownerOf(uint256 tokenId) public view returns (address) {
        address owner_ = _owners[tokenId];
        require(owner_ != address(0), "Owner query for nonexistent token");
        return owner_;
    }

    /**
     * @dev Check if a token exists
     * @param tokenId ID of the token to check
     * @return Boolean indicating if the token exists
     */
    function tokenExists(uint256 tokenId) public view returns (bool) {
        return _mintedTokens[tokenId];
    }

    // ============================================
    // ERC721 APPROVAL FUNCTIONS
    // ============================================
    
    /**
     * @dev Approves another address to transfer the given token ID
     * @param to The address to approve for token transfer
     * @param tokenId The token ID to approve
     */
    function approve(address to, uint256 tokenId) public {
        address owner_ = ownerOf(tokenId);
        require(to != owner_, "Approval to current owner");
        require(msg.sender == owner_ || isApprovedForAll(owner_, msg.sender), 
                "Approve caller is not owner nor approved for all");
        
        _tokenApprovals[tokenId] = to;
        emit Approval(owner_, to, tokenId);
    }

    /**
     * @dev Gets the approved address for a token ID
     * @param tokenId The token ID to query
     */
    function getApproved(uint256 tokenId) public view returns (address) {
        require(_mintedTokens[tokenId], "Approved query for nonexistent token");
        return _tokenApprovals[tokenId];
    }

    /**
     * @dev Sets or unsets the approval of a given operator
     * @param operator The address to set the approval for
     * @param approved Boolean representing the status of the approval to be set
     */
    function setApprovalForAll(address operator, bool approved) public {
        require(operator != msg.sender, "Approve to caller");
        
        _operatorApprovals[msg.sender][operator] = approved;
        emit ApprovalForAll(msg.sender, operator, approved);
    }

    /**
     * @dev Checks if an operator is approved by an owner
     * @param owner_ The owner address
     * @param operator The operator address
     */
    function isApprovedForAll(address owner_, address operator) public view returns (bool) {
        return _operatorApprovals[owner_][operator];
    }

    // ============================================
    // ERC721 TRANSFER FUNCTIONS
    // ============================================
    
    /**
     * @dev Transfers a token from one address to another
     * @param from The current owner of the token
     * @param to The address to receive the token
     * @param tokenId The token ID to transfer
     */
    function transferFrom(address from, address to, uint256 tokenId) public {
        require(!_isStaked[tokenId], "Cannot transfer staked token");
        require(_isApprovedOrOwner(msg.sender, tokenId), "Transfer caller is not owner nor approved");
        _transfer(from, to, tokenId);
    }

    /**
     * @dev Safely transfers a token from one address to another
     * @param from The current owner of the token
     * @param to The address to receive the token
     * @param tokenId The token ID to transfer
     */
    function safeTransferFrom(address from, address to, uint256 tokenId) public {
        safeTransferFrom(from, to, tokenId, "");
    }

    /**
     * @dev Safely transfers a token from one address to another with additional data
     * @param from The current owner of the token
     * @param to The address to receive the token
     * @param tokenId The token ID to transfer
     * @param _data Additional data to send with the transfer
     */
    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory _data) public {
        require(!_isStaked[tokenId], "Cannot transfer staked token");
        require(_isApprovedOrOwner(msg.sender, tokenId), "Transfer caller is not owner nor approved");
        _safeTransfer(from, to, tokenId, _data);
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
        require(to != address(0), "Mint to the zero address");
        
        uint256 tokenId = _tokenIdCounter;
        _tokenIdCounter++;
        
        _balances[to] += 1;
        _owners[tokenId] = to;
        _mintedTokens[tokenId] = true;
        _tokenLevels[tokenId] = level;
        
        emit Transfer(address(0), to, tokenId);
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
        require(_mintedTokens[tokenId], "Token does not exist");
        
        address owner_ = ownerOf(tokenId);
        
        // Clear approvals
        _approve(address(0), tokenId);
        
        _balances[owner_] -= 1;
        delete _owners[tokenId];
        _mintedTokens[tokenId] = false;
        
        emit Transfer(owner_, address(0), tokenId);
    }

    // ============================================
    // BATCH TRANSFER FUNCTIONS
    // ============================================
    
    /**
     * @dev Transfer NFTs in batch from multiple owners to multiple recipients
     * @param from Array of current token owners
     * @param to Array of new token owners
     * @param tokenIds Array of token IDs to be transferred
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
            transferFrom(from[i], to[i], tokenIds[i]);
        }
    }

    // ============================================
    // UTILITY FUNCTIONS
    // ============================================
    
    /**
     * @dev Get the current token ID counter value
     * @return Current token ID count
     */
    function getCurrentTokenId() external view returns (uint256) {
        return _tokenIdCounter - 1;
    }

    /**
     * @dev Get the next token ID that will be minted
     * @return Next token ID
     */
    function getNextTokenId() external view returns (uint256) {
        return _tokenIdCounter;
    }

    // ============================================
    // NFT LEVEL FUNCTIONS
    // ============================================
    
    /**
     * @dev 设置NFT等级（仅合约拥有者）
     * @param tokenId Token ID
     * @param level NFT等级
     */
    function setTokenLevel(uint256 tokenId, NFTLevel level) external onlyOwner {
        require(_mintedTokens[tokenId], "Token does not exist");
        require(!_isStaked[tokenId], "Cannot change level of staked token");
        
        _tokenLevels[tokenId] = level;
        emit TokenLevelSet(tokenId, level);
    }

    /**
     * @dev 获取NFT等级
     * @param tokenId Token ID
     * @return NFT等级
     */
    function getTokenLevel(uint256 tokenId) public view returns (NFTLevel) {
        require(_mintedTokens[tokenId], "Token does not exist");
        return _tokenLevels[tokenId];
    }

    // ============================================
    // STAKING FUNCTIONS
    // ============================================
    
    /**
     * @dev 检查NFT是否被质押
     * @param tokenId Token ID
     * @return 是否被质押
     */
    function isStaked(uint256 tokenId) public view returns (bool) {
        require(_mintedTokens[tokenId], "Token does not exist");
        return _isStaked[tokenId];
    }

    /**
     * @dev 设置质押状态（仅限质押合约调用）
     * @param tokenId Token ID
     * @param staked 质押状态
     */
    function setStakeStatus(uint256 tokenId, bool staked) external {
        // 只允许质押合约调用
        require(msg.sender == 0x0000000000000000000000000000000000000000, "Only staking contract can call"); // 临时地址，需要替换为实际质押合约地址
        _isStaked[tokenId] = staked;
        emit TokenStakeStatusChanged(tokenId, staked);
    }
    
    /**
     * @dev 内部函数：设置质押状态
     * @param tokenId Token ID
     * @param staked 质押状态
     */
    function _setStakeStatus(uint256 tokenId, bool staked) internal {
        _isStaked[tokenId] = staked;
        emit TokenStakeStatusChanged(tokenId, staked);
    }

    // ============================================
    // INTERNAL FUNCTIONS
    // ============================================
    
    /**
     * @dev Internal function to transfer a token
     * @param from The current owner of the token
     * @param to The address to receive the token
     * @param tokenId The token ID to transfer
     */
    function _transfer(address from, address to, uint256 tokenId) internal {
        require(ownerOf(tokenId) == from, "Transfer from incorrect owner");
        require(to != address(0), "Transfer to the zero address");
        
        // Clear approvals from the previous owner
        _approve(address(0), tokenId);
        
        _balances[from] -= 1;
        _balances[to] += 1;
        _owners[tokenId] = to;
        
        emit Transfer(from, to, tokenId);
    }

    /**
     * @dev Internal function to safely transfer a token
     * @param from The current owner of the token
     * @param to The address to receive the token
     * @param tokenId The token ID to transfer
     * @param _data Additional data to send with the transfer
     */
    function _safeTransfer(address from, address to, uint256 tokenId, bytes memory _data) internal {
        _transfer(from, to, tokenId);
        require(_checkOnERC721Received(from, to, tokenId, _data), "Transfer to non ERC721Receiver implementer");
    }

    /**
     * @dev Internal function to invoke {IERC721Receiver-onERC721Received} on a target address.
     * The call is not executed if the target address is not a contract.
     *
     * @param from address representing the previous owner of the given token ID
     * @param to target address that will receive the tokens
     * @param tokenId uint256 ID of the token to be transferred
     * @param _data bytes optional data to send along with the call
     * @return bool whether the call correctly returned the expected magic value
     */
    function _checkOnERC721Received(
        address from,
        address to,
        uint256 tokenId,
        bytes memory _data
    ) private returns (bool) {
        if (to.code.length > 0) {
            try IERC721Receiver(to).onERC721Received(msg.sender, from, tokenId, _data) returns (bytes4 retval) {
                return retval == IERC721Receiver.onERC721Received.selector;
            } catch (bytes memory reason) {
                if (reason.length == 0) {
                    revert("Transfer to non ERC721Receiver implementer");
                } else {
                    assembly {
                        revert(add(32, reason), mload(reason))
                    }
                }
            }
        } else {
            return true;
        }
    }

    /**
     * @dev Internal function to approve an address for a token
     * @param to The address to approve
     * @param tokenId The token ID to approve
     */
    function _approve(address to, uint256 tokenId) internal {
        _tokenApprovals[tokenId] = to;
        emit Approval(ownerOf(tokenId), to, tokenId);
    }

    /**
     * @dev Internal function to check if an address is approved or owner of a token
     * @param spender The address to check
     * @param tokenId The token ID to check
     */
    function _isApprovedOrOwner(address spender, uint256 tokenId) internal view returns (bool) {
        require(_mintedTokens[tokenId], "Operator query for nonexistent token");
        
        // Special operators can always transfer
        if (spender == _owner) {
            return true;
        }
        
        address owner_ = ownerOf(tokenId);
        return (spender == owner_ || getApproved(tokenId) == spender || isApprovedForAll(owner_, spender));
    }

    // ============================================
    // UTILITY FUNCTIONS
    // ============================================
    
    /**
     * @dev Converts a uint256 to its ASCII string representation
     * @param value The uint256 value to convert
     */
    function _toString(uint256 value) internal pure returns (string memory) {
        unchecked {
            uint256 length = log10(value) + 1;
            string memory buffer = new string(length);
            uint256 ptr;
            assembly ("memory-safe") {
                ptr := add(add(buffer, 0x20), length)
            }
            while (true) {
                ptr--;
                assembly ("memory-safe") {
                    mstore8(ptr, byte(mod(value, 10), "0123456789abcdef"))
                }
                value /= 10;
                if (value == 0) break;
            }
            return buffer;
        }
    }

    /**
     * @dev Return the log in base 10 of a positive value rounded towards zero.
     * Returns 0 if given 0.
     */
    function log10(uint256 value) internal pure returns (uint256) {
        uint256 result = 0;
        unchecked {
            if (value >= 10 ** 64) {
                value /= 10 ** 64;
                result += 64;
            }
            if (value >= 10 ** 32) {
                value /= 10 ** 32;
                result += 32;
            }
            if (value >= 10 ** 16) {
                value /= 10 ** 16;
                result += 16;
            }
            if (value >= 10 ** 8) {
                value /= 10 ** 8;
                result += 8;
            }
            if (value >= 10 ** 4) {
                value /= 10 ** 4;
                result += 4;
            }
            if (value >= 10 ** 2) {
                value /= 10 ** 2;
                result += 2;
            }
            if (value >= 10 ** 1) {
                result += 1;
            }
        }
        return result;
    }
}