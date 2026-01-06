// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../CPNFT/CPNFT.sol";
import "./interfaces/IMarketplace.sol";

/**
 * @title Marketplace - C2C NFT Trading Marketplace
 * @dev Main marketplace contract supporting fixed price and auction listings
 * @notice Uses escrow model for secure C2C transactions
 */
contract Marketplace is 
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    PausableUpgradeable,
    IMarketplace,
    IERC721Receiver
{
    using SafeERC20 for IERC20;
    
    // ============================================
    // STORAGE VARIABLES
    // ============================================
    
    // Contract dependencies
    address public nftContract;              // CPNFT contract address
    address public paymentToken;             // USDT or other ERC20 token address
    
    // Platform configuration
    address public platformFeeRecipient;     // Platform fee recipient address
    uint256 public platformFeeRate;         // Platform fee rate in basis points (500 = 5%)
    uint256 public delistingLimit;          // Maximum delisting count per window (default: 3)
    uint256 public delistingWindow;         // Delisting time window in seconds (default: 24 hours)
    
    // Listing management
    uint256 public nextListingId;
    mapping(uint256 => Listing) public listings;
    mapping(address => uint256[]) public userListings;  // User's listing IDs
    
    // Auction bidding
    mapping(uint256 => Bid[]) public bids;               // Listing ID => Bids array
    mapping(uint256 => uint256) public highestBidIndex; // Listing ID => Index of highest bid
    
    // Delisting frequency limit
    struct DelistingRecord {
        uint256 count;
        uint256 lastResetTime;
    }
    mapping(address => DelistingRecord) public delistingRecords;
    
    // Testing mode for time manipulation
    bool public testMode;
    uint256 public testTimestamp;
    
    // ============================================
    // MODIFIERS
    // ============================================
    
    modifier onlyValidListing(uint256 listingId) {
        require(listings[listingId].listingId != 0, "Listing does not exist");
        _;
    }
    
    modifier onlyActiveListing(uint256 listingId) {
        require(listings[listingId].status == ListingStatus.Active, "Listing is not active");
        _;
    }
    
    // ============================================
    // INITIALIZER
    // ============================================
    
    /**
     * @dev Initialize the marketplace contract
     * @param _nftContract CPNFT contract address
     * @param _paymentToken USDT or other ERC20 token address
     * @param _platformFeeRecipient Platform fee recipient address
     * @param _platformFeeRate Platform fee rate in basis points (500 = 5%)
     * @param _delistingLimit Maximum delisting count per window (default: 3)
     * @param _delistingWindow Delisting time window in seconds (default: 86400 = 24 hours)
     * @param _owner Contract owner address
     */
    function initialize(
        address _nftContract,
        address _paymentToken,
        address _platformFeeRecipient,
        uint256 _platformFeeRate,
        uint256 _delistingLimit,
        uint256 _delistingWindow,
        address _owner
    ) public initializer {
        __UUPSUpgradeable_init();
        __Ownable_init(_owner);
        __ReentrancyGuard_init();
        __Pausable_init();
        
        require(_nftContract != address(0), "NFT contract address cannot be zero");
        require(_paymentToken != address(0), "Payment token address cannot be zero");
        require(_platformFeeRecipient != address(0), "Platform fee recipient cannot be zero");
        require(_platformFeeRate <= 10000, "Platform fee rate cannot exceed 100%");
        require(_delistingLimit > 0, "Delisting limit must be greater than 0");
        require(_delistingWindow > 0, "Delisting window must be greater than 0");
        
        nftContract = _nftContract;
        paymentToken = _paymentToken;
        platformFeeRecipient = _platformFeeRecipient;
        platformFeeRate = _platformFeeRate;
        delistingLimit = _delistingLimit;
        delistingWindow = _delistingWindow;
        
        nextListingId = 1;
    }
    
    // ============================================
    // UUPS UPGRADE AUTHORIZATION
    // ============================================
    
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
    
    // ============================================
    // TIME MANAGEMENT FOR TESTING
    // ============================================
    
    /**
     * @dev Get current timestamp (real or test)
     */
    function _getCurrentTimestamp() internal view returns (uint256) {
        return testMode ? testTimestamp : block.timestamp;
    }
    
    /**
     * @dev Get current timestamp (public version for external contracts)
     */
    function getCurrentTimestamp() external view returns (uint256) {
        return _getCurrentTimestamp();
    }
    
    /**
     * @dev Enable test mode and set initial test timestamp
     */
    function enableTestMode(uint256 initialTimestamp) external onlyOwner {
        testMode = true;
        testTimestamp = initialTimestamp > 0 ? initialTimestamp : block.timestamp;
    }
    
    /**
     * @dev Disable test mode (back to real time)
     */
    function disableTestMode() external onlyOwner {
        testMode = false;
        testTimestamp = 0;
    }
    
    /**
     * @dev Set test timestamp (only in test mode)
     */
    function setTestTimestamp(uint256 timestamp) external onlyOwner {
        require(testMode, "Not in test mode");
        require(timestamp >= testTimestamp, "Cannot go back in time");
        testTimestamp = timestamp;
    }
    
    // ============================================
    // LISTING FUNCTIONS
    // ============================================
    
    /**
     * @dev Create a new listing (Backend/Admin only)
     * @param seller The seller address
     * @param tokenId NFT token ID
     * @param listingType Listing type (FixedPrice or Auction)
     * @param price Fixed price or starting bid price
     * @param duration Auction duration in seconds (only for Auction type)
     * @param minBidIncrement Minimum bid increment in basis points (only for Auction type)
     * @return listingId The created listing ID
     * @notice Only contract owner or authorized backend can call this function
     * @notice Uses the NFT contract address configured during initialization
     */
    function createListing(
        address seller,
        uint256 tokenId,
        ListingType listingType,
        uint256 price,
        uint256 duration,
        uint256 minBidIncrement
    ) external override nonReentrant whenNotPaused onlyOwner returns (uint256 listingId) {
        require(seller != address(0), "Invalid seller address");
        require(nftContract != address(0), "NFT contract not configured");
        require(price > 0, "Price must be greater than 0");
        
        CPNFT nft = CPNFT(nftContract);
        require(nft.ownerOf(tokenId) == seller, "Seller is not the owner");
        require(!nft.isStaked(tokenId), "Cannot list staked token");
        
        uint256 currentTime = _getCurrentTimestamp();
        uint256 endTime = 0;
        
        if (listingType == ListingType.Auction) {
            require(duration > 0, "Auction duration must be greater than 0");
            require(minBidIncrement > 0, "Min bid increment must be greater than 0");
            endTime = currentTime + duration;
        }
        
        listingId = nextListingId++;
        
        listings[listingId] = Listing({
            listingId: listingId,
            seller: seller,
            nftContract: nftContract,
            tokenId: tokenId,
            listingType: listingType,
            price: price,
            status: ListingStatus.Active,
            startTime: currentTime,
            endTime: endTime,
            minBidIncrement: minBidIncrement
        });
        
        userListings[seller].push(listingId);
        
        // Transfer NFT from seller to contract (escrow) - no approval needed
        nft.marketplaceSafeTransferFrom(seller, address(this), tokenId);
        
        emit ListingCreated(
            listingId,
            seller,
            nftContract,
            tokenId,
            listingType,
            price,
            endTime
        );
        
        return listingId;
    }
    
    /**
     * @dev Cancel a listing (Backend/Admin only)
     * @param listingId The listing ID to cancel
     * @notice Only contract owner or authorized backend can call this function
     * @notice Enforces delisting frequency limit (24 hours, 3 times max)
     */
    function cancelListing(uint256 listingId) 
        external 
        override 
        nonReentrant 
        whenNotPaused 
        onlyOwner 
        onlyValidListing(listingId)
        onlyActiveListing(listingId)
    {
        Listing storage listing = listings[listingId];
        address seller = listing.seller;
        
        // Check delisting frequency limit
        _checkDelistingLimit(seller);
        
        // Update listing status
        listing.status = ListingStatus.Cancelled;
        
        // If auction, refund all bids
        if (listing.listingType == ListingType.Auction) {
            _refundAllBids(listingId);
        }
        
        // Transfer NFT back to seller - no approval needed
        CPNFT(listing.nftContract).marketplaceSafeTransferFrom(address(this), seller, listing.tokenId);
        
        emit ListingCancelled(listingId, seller);
    }
    
    /**
     * @dev Buy an item at fixed price (Backend/Admin only)
     * @param listingId The listing ID
     * @param buyer The buyer address
     * @notice Only contract owner or authorized backend can call this function
     * @notice Buyer must have approved sufficient payment token amount
     */
    function buyItem(uint256 listingId, address buyer) 
        external 
        override 
        nonReentrant 
        whenNotPaused 
        onlyOwner 
        onlyValidListing(listingId)
        onlyActiveListing(listingId)
    {
        Listing storage listing = listings[listingId];
        require(listing.listingType == ListingType.FixedPrice, "Not a fixed price listing");
        require(buyer != address(0), "Invalid buyer address");
        require(buyer != listing.seller, "Seller cannot buy own item");
        
        uint256 price = listing.price;
        
        // Transfer payment token from buyer to contract
        IERC20(paymentToken).safeTransferFrom(buyer, address(this), price);
        
        // Calculate platform fee
        uint256 platformFee = (price * platformFeeRate) / 10000;
        uint256 sellerAmount = price - platformFee;
        
        // Transfer payment to seller and platform
        IERC20(paymentToken).safeTransfer(listing.seller, sellerAmount);
        IERC20(paymentToken).safeTransfer(platformFeeRecipient, platformFee);
        
        // Transfer NFT to buyer - no approval needed
        CPNFT(listing.nftContract).marketplaceSafeTransferFrom(address(this), buyer, listing.tokenId);
        
        // Update listing status
        listing.status = ListingStatus.Sold;
        
        emit ItemSold(listingId, listing.seller, buyer, price, platformFee);
    }
    
    // ============================================
    // AUCTION FUNCTIONS
    // ============================================
    
    /**
     * @dev Place a bid on an auction (Backend/Admin only)
     * @param listingId The listing ID
     * @param bidder The bidder address
     * @param bidAmount The bid amount
     * @notice Only contract owner or authorized backend can call this function
     * @notice Bidder must have approved sufficient payment token amount
     */
    function placeBid(uint256 listingId, address bidder, uint256 bidAmount) 
        external 
        override 
        nonReentrant 
        whenNotPaused 
        onlyOwner 
        onlyValidListing(listingId)
        onlyActiveListing(listingId)
    {
        Listing storage listing = listings[listingId];
        require(listing.listingType == ListingType.Auction, "Not an auction listing");
        require(bidder != address(0), "Invalid bidder address");
        require(bidder != listing.seller, "Seller cannot bid on own auction");
        
        uint256 currentTime = _getCurrentTimestamp();
        require(currentTime < listing.endTime, "Auction has ended");
        
        uint256 highestBidIdx = highestBidIndex[listingId];
        uint256 currentHighestBid = 0;
        address previousBidder = address(0);
        
        if (bids[listingId].length > 0 && highestBidIdx < bids[listingId].length) {
            Bid storage currentBid = bids[listingId][highestBidIdx];
            require(!currentBid.refunded, "Highest bid is refunded");
            currentHighestBid = currentBid.amount;
            previousBidder = currentBid.bidder;
        } else {
            // First bid must be at least starting price
            require(bidAmount >= listing.price, "Bid must be at least starting price");
        }
        
        // Check bid increment
        uint256 minBid = currentHighestBid > 0 
            ? currentHighestBid + (currentHighestBid * listing.minBidIncrement) / 10000
            : listing.price;
        require(bidAmount >= minBid, "Bid amount too low");
        
        // Transfer payment token from bidder to contract
        IERC20(paymentToken).safeTransferFrom(bidder, address(this), bidAmount);
        
        // Refund previous highest bidder if exists
        uint256 refundAmount = 0;
        if (previousBidder != address(0)) {
            refundAmount = currentHighestBid;
            bids[listingId][highestBidIdx].refunded = true;
            IERC20(paymentToken).safeTransfer(previousBidder, refundAmount);
            emit BidRefunded(listingId, previousBidder, refundAmount);
        }
        
        // Create new bid
        bids[listingId].push(Bid({
            bidder: bidder,
            amount: bidAmount,
            timestamp: currentTime,
            refunded: false
        }));
        
        highestBidIndex[listingId] = bids[listingId].length - 1;
        
        emit BidPlaced(listingId, bidder, bidAmount, previousBidder, refundAmount);
    }
    
    /**
     * @dev Settle an auction (Backend/Admin only)
     * @param listingId The listing ID
     * @param winner The winner address (can be called by seller to accept or winner to claim)
     * @notice Only contract owner or authorized backend can call this function
     */
    function settleAuction(uint256 listingId, address winner) 
        external 
        override 
        nonReentrant 
        whenNotPaused 
        onlyOwner 
        onlyValidListing(listingId)
        onlyActiveListing(listingId)
    {
        Listing storage listing = listings[listingId];
        require(listing.listingType == ListingType.Auction, "Not an auction listing");
        
        uint256 currentTime = _getCurrentTimestamp();
        require(currentTime >= listing.endTime || winner == listing.seller, "Auction not ended");
        
        uint256 highestBidIdx = highestBidIndex[listingId];
        require(bids[listingId].length > 0, "No bids placed");
        require(highestBidIdx < bids[listingId].length, "Invalid highest bid index");
        
        Bid storage winningBid = bids[listingId][highestBidIdx];
        require(!winningBid.refunded, "Winning bid already refunded");
        
        // If winner is specified, verify it matches the highest bidder
        if (winner != address(0)) {
            require(winner == winningBid.bidder, "Winner does not match highest bidder");
        } else {
            winner = winningBid.bidder;
        }
        
        uint256 finalPrice = winningBid.amount;
        
        // Calculate platform fee
        uint256 platformFee = (finalPrice * platformFeeRate) / 10000;
        uint256 sellerAmount = finalPrice - platformFee;
        
        // Transfer payment to seller and platform
        IERC20(paymentToken).safeTransfer(listing.seller, sellerAmount);
        IERC20(paymentToken).safeTransfer(platformFeeRecipient, platformFee);
        
        // Transfer NFT to winner - no approval needed
        CPNFT(listing.nftContract).marketplaceSafeTransferFrom(address(this), winner, listing.tokenId);
        
        // Mark bid as used (refunded flag repurposed)
        winningBid.refunded = true;
        
        // Update listing status
        listing.status = ListingStatus.Sold;
        
        emit AuctionSettled(listingId, listing.seller, winner, finalPrice, platformFee);
    }
    
    // ============================================
    // VIEW FUNCTIONS
    // ============================================
    
    /**
     * @dev Get listing details
     * @param listingId The listing ID
     * @return Listing struct
     */
    function getListing(uint256 listingId) 
        external 
        view 
        override 
        onlyValidListing(listingId) 
        returns (Listing memory) 
    {
        return listings[listingId];
    }
    
    /**
     * @dev Get all bids for a listing
     * @param listingId The listing ID
     * @return Array of Bid structs
     */
    function getBids(uint256 listingId) 
        external 
        view 
        override 
        onlyValidListing(listingId) 
        returns (Bid[] memory) 
    {
        return bids[listingId];
    }
    
    /**
     * @dev Get highest bid for a listing
     * @param listingId The listing ID
     * @return Highest Bid struct
     */
    function getHighestBid(uint256 listingId) 
        external 
        view 
        override 
        onlyValidListing(listingId) 
        returns (Bid memory) 
    {
        uint256 highestBidIdx = highestBidIndex[listingId];
        require(bids[listingId].length > 0, "No bids placed");
        require(highestBidIdx < bids[listingId].length, "Invalid highest bid index");
        return bids[listingId][highestBidIdx];
    }
    
    /**
     * @dev Get user's listings
     * @param user User address
     * @return Array of listing IDs
     */
    function getUserListings(address user) external view returns (uint256[] memory) {
        return userListings[user];
    }
    
    /**
     * @dev Get delisting record for a user
     * @param user User address
     * @return count Current delisting count
     * @return lastResetTime Last reset timestamp
     * @return remainingCount Remaining delisting count in current window
     */
    function getDelistingRecord(address user) 
        external 
        view 
        returns (uint256 count, uint256 lastResetTime, uint256 remainingCount) 
    {
        DelistingRecord memory record = delistingRecords[user];
        uint256 currentTime = _getCurrentTimestamp();
        
        // Check if window has expired
        if (currentTime >= record.lastResetTime + delistingWindow) {
            return (0, currentTime, delistingLimit);
        }
        
        uint256 remaining = record.count < delistingLimit 
            ? delistingLimit - record.count 
            : 0;
        
        return (record.count, record.lastResetTime, remaining);
    }
    
    // ============================================
    // INTERNAL FUNCTIONS
    // ============================================
    
    /**
     * @dev Check and update delisting frequency limit
     * @param user User address
     */
    function _checkDelistingLimit(address user) internal {
        uint256 currentTime = _getCurrentTimestamp();
        DelistingRecord storage record = delistingRecords[user];
        
        // Reset if window has expired
        if (currentTime >= record.lastResetTime + delistingWindow) {
            record.count = 0;
            record.lastResetTime = currentTime;
        }
        
        // Check limit
        require(record.count < delistingLimit, "Delisting limit exceeded");
        
        // Increment count
        record.count++;
    }
    
    /**
     * @dev Refund all bids for a cancelled auction
     * @param listingId The listing ID
     */
    function _refundAllBids(uint256 listingId) internal {
        Bid[] storage listingBids = bids[listingId];
        IERC20 token = IERC20(paymentToken);
        
        for (uint256 i = 0; i < listingBids.length; i++) {
            if (!listingBids[i].refunded) {
                token.safeTransfer(listingBids[i].bidder, listingBids[i].amount);
                listingBids[i].refunded = true;
                emit BidRefunded(listingId, listingBids[i].bidder, listingBids[i].amount);
            }
        }
    }
    
    // ============================================
    // ADMIN FUNCTIONS
    // ============================================
    
    /**
     * @dev Pause the contract
     */
    function pause() external onlyOwner {
        _pause();
    }
    
    /**
     * @dev Unpause the contract
     */
    function unpause() external onlyOwner {
        _unpause();
    }
    
    /**
     * @dev Update platform fee recipient
     * @param _platformFeeRecipient New platform fee recipient address
     */
    function updatePlatformFeeRecipient(address _platformFeeRecipient) external onlyOwner {
        require(_platformFeeRecipient != address(0), "Invalid address");
        platformFeeRecipient = _platformFeeRecipient;
    }
    
    /**
     * @dev Update platform fee rate
     * @param _platformFeeRate New platform fee rate in basis points
     */
    function updatePlatformFeeRate(uint256 _platformFeeRate) external onlyOwner {
        require(_platformFeeRate <= 10000, "Fee rate cannot exceed 100%");
        platformFeeRate = _platformFeeRate;
    }
    
    /**
     * @dev Update delisting limit
     * @param _delistingLimit New delisting limit
     */
    function updateDelistingLimit(uint256 _delistingLimit) external onlyOwner {
        require(_delistingLimit > 0, "Limit must be greater than 0");
        delistingLimit = _delistingLimit;
    }
    
    /**
     * @dev Update delisting window
     * @param _delistingWindow New delisting window in seconds
     */
    function updateDelistingWindow(uint256 _delistingWindow) external onlyOwner {
        require(_delistingWindow > 0, "Window must be greater than 0");
        delistingWindow = _delistingWindow;
    }
    
    /**
     * @dev Update NFT contract address
     * @param _nftContract New NFT contract address
     */
    function updateNFTContract(address _nftContract) external onlyOwner {
        require(_nftContract != address(0), "Invalid address");
        nftContract = _nftContract;
    }
    
    /**
     * @dev Update payment token address
     * @param _paymentToken New payment token address
     */
    function updatePaymentToken(address _paymentToken) external onlyOwner {
        require(_paymentToken != address(0), "Invalid address");
        paymentToken = _paymentToken;
    }
    
    /**
     * @dev Emergency withdraw ERC20 tokens
     * @param token Token address
     * @param to Recipient address
     * @param amount Amount to withdraw
     */
    function emergencyWithdraw(address token, address to, uint256 amount) external onlyOwner {
        require(to != address(0), "Invalid recipient");
        IERC20(token).safeTransfer(to, amount);
    }
    
    /**
     * @dev Emergency withdraw ERC721 tokens from the configured NFT contract
     * @param tokenId Token ID
     * @param to Recipient address
     */
    function emergencyWithdrawNFT(
        uint256 tokenId,
        address to
    ) external onlyOwner {
        require(to != address(0), "Invalid recipient");
        CPNFT(nftContract).marketplaceSafeTransferFrom(address(this), to, tokenId);
    }
    
    /**
     * @dev Get contract version
     */
    function version() public pure returns (string memory) {
        return "1.0.0";
    }
    
    function onERC721Received(address, address, uint256, bytes calldata) external pure returns (bytes4) {
        return IERC721Receiver.onERC721Received.selector;
    }
}
