// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";

/**
 * @title BatchTransfer
 * @dev Batch transfer contract for NFTs (ERC721) and Fungible Tokens (ERC20)
 * @notice This contract allows users to batch transfer NFT and FT tokens efficiently
 */
contract BatchTransfer is ReentrancyGuardTransient {
    using SafeERC20 for IERC20;

    // ============================================
    // EVENTS
    // ============================================
    
    event BatchNFTTransfer(
        address indexed nftContract,
        address indexed from,
        uint256[] tokenIds,
        address[] recipients
    );
    
    event BatchTokenTransfer(
        address indexed tokenContract,
        address indexed from,
        uint256 totalAmount,
        address[] recipients
    );
    
    event SingleNFTBatchTransfer(
        address indexed nftContract,
        address indexed from,
        address indexed to,
        uint256[] tokenIds
    );

    // ============================================
    // NFT BATCH TRANSFER FUNCTIONS
    // ============================================
    
    /**
     * @dev Batch transfer NFTs from msg.sender to multiple recipients
     * @param nftContract Address of the ERC721 contract
     * @param recipients Array of recipient addresses (must match tokenIds length)
     * @param tokenIds Array of token IDs to transfer
     * @notice msg.sender must own all tokenIds and have approved this contract
     */
    function batchTransferNFT(
        address nftContract,
        address[] calldata recipients,
        uint256[] calldata tokenIds
    ) external nonReentrant {
        require(nftContract != address(0), "Invalid NFT contract address");
        require(recipients.length == tokenIds.length, "Arrays length mismatch");
        require(recipients.length > 0, "Empty arrays");
        
        IERC721 nft = IERC721(nftContract);
        
        for (uint256 i = 0; i < tokenIds.length; i++) {
            require(recipients[i] != address(0), "Invalid recipient address");
            require(nft.ownerOf(tokenIds[i]) == msg.sender, "Not token owner");
            
            nft.safeTransferFrom(msg.sender, recipients[i], tokenIds[i]);
        }
        
        emit BatchNFTTransfer(nftContract, msg.sender, tokenIds, recipients);
    }
    
    /**
     * @dev Transfer multiple NFTs from msg.sender to a single recipient
     * @param nftContract Address of the ERC721 contract
     * @param to Recipient address
     * @param tokenIds Array of token IDs to transfer
     * @notice msg.sender must own all tokenIds and have approved this contract
     */
    function batchTransferNFTToSingle(
        address nftContract,
        address to,
        uint256[] calldata tokenIds
    ) external nonReentrant {
        require(nftContract != address(0), "Invalid NFT contract address");
        require(to != address(0), "Invalid recipient address");
        require(tokenIds.length > 0, "Empty token array");
        
        IERC721 nft = IERC721(nftContract);
        
        for (uint256 i = 0; i < tokenIds.length; i++) {
            require(nft.ownerOf(tokenIds[i]) == msg.sender, "Not token owner");
            nft.safeTransferFrom(msg.sender, to, tokenIds[i]);
        }
        
        emit SingleNFTBatchTransfer(nftContract, msg.sender, to, tokenIds);
    }
    
    /**
     * @dev Transfer multiple NFTs from different collections to a single recipient
     * @param nftContracts Array of ERC721 contract addresses
     * @param to Recipient address
     * @param tokenIds Array of token IDs (must match nftContracts length)
     * @notice msg.sender must own all tokenIds and have approved this contract for each collection
     */
    function batchTransferMultipleNFTCollections(
        address[] calldata nftContracts,
        address to,
        uint256[] calldata tokenIds
    ) external nonReentrant {
        require(nftContracts.length == tokenIds.length, "Arrays length mismatch");
        require(to != address(0), "Invalid recipient address");
        require(nftContracts.length > 0, "Empty arrays");
        
        for (uint256 i = 0; i < nftContracts.length; i++) {
            require(nftContracts[i] != address(0), "Invalid NFT contract address");
            
            IERC721 nft = IERC721(nftContracts[i]);
            require(nft.ownerOf(tokenIds[i]) == msg.sender, "Not token owner");
            
            nft.safeTransferFrom(msg.sender, to, tokenIds[i]);
        }
    }

    // ============================================
    // ERC20 BATCH TRANSFER FUNCTIONS
    // ============================================
    
    /**
     * @dev Batch transfer ERC20 tokens from msg.sender to multiple recipients
     * @param tokenContract Address of the ERC20 contract
     * @param recipients Array of recipient addresses
     * @param amounts Array of token amounts to transfer (must match recipients length)
     * @notice msg.sender must have approved this contract for the total amount
     */
    function batchTransferToken(
        address tokenContract,
        address[] calldata recipients,
        uint256[] calldata amounts
    ) external nonReentrant {
        require(tokenContract != address(0), "Invalid token contract address");
        require(recipients.length == amounts.length, "Arrays length mismatch");
        require(recipients.length > 0, "Empty arrays");
        
        IERC20 token = IERC20(tokenContract);
        uint256 totalAmount = 0;
        
        for (uint256 i = 0; i < recipients.length; i++) {
            require(recipients[i] != address(0), "Invalid recipient address");
            require(amounts[i] > 0, "Amount must be greater than 0");
            
            totalAmount += amounts[i];
            token.safeTransferFrom(msg.sender, recipients[i], amounts[i]);
        }
        
        emit BatchTokenTransfer(tokenContract, msg.sender, totalAmount, recipients);
    }
    
    /**
     * @dev Batch transfer equal amounts of ERC20 tokens to multiple recipients
     * @param tokenContract Address of the ERC20 contract
     * @param recipients Array of recipient addresses
     * @param amountPerRecipient Amount of tokens to send to each recipient
     * @notice msg.sender must have approved this contract for (amountPerRecipient * recipients.length)
     */
    function batchTransferTokenEqual(
        address tokenContract,
        address[] calldata recipients,
        uint256 amountPerRecipient
    ) external nonReentrant {
        require(tokenContract != address(0), "Invalid token contract address");
        require(recipients.length > 0, "Empty recipients array");
        require(amountPerRecipient > 0, "Amount must be greater than 0");
        
        IERC20 token = IERC20(tokenContract);
        uint256 totalAmount = amountPerRecipient * recipients.length;
        
        for (uint256 i = 0; i < recipients.length; i++) {
            require(recipients[i] != address(0), "Invalid recipient address");
            token.safeTransferFrom(msg.sender, recipients[i], amountPerRecipient);
        }
        
        emit BatchTokenTransfer(tokenContract, msg.sender, totalAmount, recipients);
    }
    
    /**
     * @dev Batch transfer multiple ERC20 tokens to a single recipient
     * @param tokenContracts Array of ERC20 contract addresses
     * @param to Recipient address
     * @param amounts Array of token amounts (must match tokenContracts length)
     * @notice msg.sender must have approved this contract for each token
     */
    function batchTransferMultipleTokens(
        address[] calldata tokenContracts,
        address to,
        uint256[] calldata amounts
    ) external nonReentrant {
        require(tokenContracts.length == amounts.length, "Arrays length mismatch");
        require(to != address(0), "Invalid recipient address");
        require(tokenContracts.length > 0, "Empty arrays");
        
        for (uint256 i = 0; i < tokenContracts.length; i++) {
            require(tokenContracts[i] != address(0), "Invalid token contract address");
            require(amounts[i] > 0, "Amount must be greater than 0");
            
            IERC20 token = IERC20(tokenContracts[i]);
            token.safeTransferFrom(msg.sender, to, amounts[i]);
        }
    }

    // ============================================
    // COMBINED TRANSFER FUNCTIONS
    // ============================================
    
    /**
     * @dev Combined batch transfer of both NFTs and ERC20 tokens
     * @param nftContract Address of the ERC721 contract (set to address(0) to skip NFT transfer)
     * @param nftRecipients Array of NFT recipient addresses
     * @param tokenIds Array of NFT token IDs
     * @param tokenContract Address of the ERC20 contract (set to address(0) to skip token transfer)
     * @param tokenRecipients Array of token recipient addresses
     * @param tokenAmounts Array of token amounts
     * @notice This function allows transferring both NFTs and tokens in a single transaction
     */
    function combinedBatchTransfer(
        address nftContract,
        address[] calldata nftRecipients,
        uint256[] calldata tokenIds,
        address tokenContract,
        address[] calldata tokenRecipients,
        uint256[] calldata tokenAmounts
    ) external nonReentrant {
        // Transfer NFTs if specified
        if (nftContract != address(0)) {
            require(nftRecipients.length == tokenIds.length, "NFT arrays length mismatch");
            IERC721 nft = IERC721(nftContract);
            
            for (uint256 i = 0; i < tokenIds.length; i++) {
                require(nftRecipients[i] != address(0), "Invalid NFT recipient");
                require(nft.ownerOf(tokenIds[i]) == msg.sender, "Not NFT owner");
                nft.safeTransferFrom(msg.sender, nftRecipients[i], tokenIds[i]);
            }
            
            if (tokenIds.length > 0) {
                emit BatchNFTTransfer(nftContract, msg.sender, tokenIds, nftRecipients);
            }
        }
        
        // Transfer tokens if specified
        if (tokenContract != address(0)) {
            require(tokenRecipients.length == tokenAmounts.length, "Token arrays length mismatch");
            IERC20 token = IERC20(tokenContract);
            uint256 totalAmount = 0;
            
            for (uint256 i = 0; i < tokenRecipients.length; i++) {
                require(tokenRecipients[i] != address(0), "Invalid token recipient");
                require(tokenAmounts[i] > 0, "Invalid token amount");
                token.safeTransferFrom(msg.sender, tokenRecipients[i], tokenAmounts[i]);
                totalAmount += tokenAmounts[i];
            }
            
            if (tokenRecipients.length > 0) {
                emit BatchTokenTransfer(tokenContract, msg.sender, totalAmount, tokenRecipients);
            }
        }
    }

    // ============================================
    // VIEW FUNCTIONS
    // ============================================
    
    /**
     * @dev Check if user has approved this contract for a specific NFT collection
     * @param nftContract Address of the ERC721 contract
     * @param owner Address of the token owner
     * @return bool Whether the contract is approved
     */
    function isNFTApproved(address nftContract, address owner) external view returns (bool) {
        IERC721 nft = IERC721(nftContract);
        return nft.isApprovedForAll(owner, address(this));
    }
    
    /**
     * @dev Check if user has sufficient token allowance for this contract
     * @param tokenContract Address of the ERC20 contract
     * @param owner Address of the token owner
     * @param amount Required amount
     * @return bool Whether the allowance is sufficient
     */
    function hasTokenAllowance(
        address tokenContract,
        address owner,
        uint256 amount
    ) external view returns (bool) {
        IERC20 token = IERC20(tokenContract);
        return token.allowance(owner, address(this)) >= amount;
    }

    // ============================================
    // ERC721 RECEIVER
    // ============================================
    
    /**
     * @dev Whenever an {IERC721} tokenId token is transferred to this contract via {IERC721-safeTransferFrom}
     * by operator from from, this function is called.
     */
    function onERC721Received(
        address,
        address,
        uint256,
        bytes calldata
    ) external pure returns (bytes4) {
        return this.onERC721Received.selector;
    }
}

