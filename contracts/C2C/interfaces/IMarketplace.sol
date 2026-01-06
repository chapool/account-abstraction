// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/**
 * @title IMarketplace
 * @dev Interface for C2C NFT Marketplace contract
 */
interface IMarketplace {
    // ============================================
    // ENUMS
    // ============================================
    
    enum ListingType {
        FixedPrice,  // 一口价
        Auction      // 拍卖
    }
    
    enum ListingStatus {
        Active,      // 上架中
        Cancelled,   // 已取消
        Sold         // 已售出
    }
    
    // ============================================
    // STRUCTS
    // ============================================
    
    struct Listing {
        uint256 listingId;
        address seller;
        address nftContract;
        uint256 tokenId;
        ListingType listingType;
        uint256 price;              // 一口价或起拍价
        ListingStatus status;
        uint256 startTime;
        uint256 endTime;            // 拍卖结束时间（仅拍卖）
        uint256 minBidIncrement;     // 最小加价幅度（仅拍卖）
    }
    
    struct Bid {
        address bidder;
        address payer;
        uint256 amount;
        uint256 timestamp;
        bool refunded;  // Packed with timestamp to save gas
    }
    
    // ============================================
    // EVENTS
    // ============================================
    
    event ListingCreated(
        uint256 indexed listingId,
        address indexed seller,
        address indexed nftContract,
        uint256 tokenId,
        ListingType listingType,
        uint256 price,
        uint256 endTime
    );
    
    event ListingCancelled(
        uint256 indexed listingId,
        address indexed seller
    );
    
    event ItemSold(
        uint256 indexed listingId,
        address indexed seller,
        address indexed buyer,
        address payer,
        uint256 price,
        uint256 platformFee
    );
    
    event BidPlaced(
        uint256 indexed listingId,
        address indexed bidder,
        address payer,
        uint256 amount,
        address indexed previousBidder,
        uint256 refundAmount
    );
    
    event AuctionSettled(
        uint256 indexed listingId,
        address indexed seller,
        address indexed winner,
        address payer,
        uint256 finalPrice,
        uint256 platformFee
    );
    
    event BidRefunded(
        uint256 indexed listingId,
        address indexed bidder,
        address payer,
        uint256 amount
    );
    
    // ============================================
    // FUNCTIONS
    // ============================================
    
    function createListing(
        address seller,
        uint256 tokenId,
        ListingType listingType,
        uint256 price,
        uint256 duration,           // 拍卖时长（秒）
        uint256 minBidIncrement     // 最小加价幅度（仅拍卖）
    ) external returns (uint256 listingId);
    
    function cancelListing(uint256 listingId) external;
    
    function buyItem(uint256 listingId, address buyer, address payer) external;
    
    function placeBid(uint256 listingId, address bidder, address payer, uint256 bidAmount) external;
    
    function settleAuction(uint256 listingId, address winner) external;
    
    function getListing(uint256 listingId) external view returns (Listing memory);
    
    function getBids(uint256 listingId) external view returns (Bid[] memory);
    
    function getHighestBid(uint256 listingId) external view returns (Bid memory);
    
    // Admin configuration setters
    function updatePlatformFeeRecipient(address _platformFeeRecipient) external;
    function updatePlatformFeeRate(uint256 _platformFeeRate) external;
    function updateDelistingLimit(uint256 _delistingLimit) external;
    function updateDelistingWindow(uint256 _delistingWindow) external;
}
