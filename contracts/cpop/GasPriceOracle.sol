// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "./interfaces/IGasPriceOracle.sol";

/**
 * @title GasPriceOracle
 * @notice Oracle contract for ETH/CPOP exchange rates and gas price calculations
 * @dev Supports multiple price feeds with fallback mechanisms and validation
 */
contract GasPriceOracle is 
    Initializable, 
    IGasPriceOracle, 
    OwnableUpgradeable, 
    UUPSUpgradeable,
    PausableUpgradeable 
{
    struct PriceFeed {
        uint256 price;
        uint256 timestamp;
        uint256 fallbackPrice;
        bool isActive;
    }

    // Price feeds mapping: feedType => PriceFeed
    mapping(string => PriceFeed) private priceFeeds;
    
    // Authorized oracle addresses
    mapping(address => bool) public authorizedOracles;
    
    // Oracle configuration
    uint256 public maxPriceAge; // Maximum allowed price age in seconds
    uint256 public priceDeviationThreshold; // Maximum allowed price deviation (basis points, e.g., 500 = 5%)
    
    // Constants
    uint256 private constant BASIS_POINTS = 10000;
    uint256 private constant PRECISION = 1e18;
    string private constant ETH_USD_FEED = "ETH/USD";
    string private constant CPOP_USD_FEED = "CPOP/USD";

    modifier onlyAuthorizedOracle() {
        require(
            authorizedOracles[msg.sender] || msg.sender == owner(),
            "GasPriceOracle: unauthorized oracle"
        );
        _;
    }

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize the oracle
     * @param _owner Owner of the contract
     * @param _maxPriceAge Maximum price age in seconds
     * @param _priceDeviationThreshold Price deviation threshold in basis points
     */
    function initialize(
        address _owner,
        uint256 _maxPriceAge,
        uint256 _priceDeviationThreshold
    ) external initializer {
        require(_owner != address(0), "GasPriceOracle: invalid owner");
        require(_maxPriceAge > 0, "GasPriceOracle: invalid max price age");
        require(_priceDeviationThreshold <= BASIS_POINTS, "GasPriceOracle: invalid threshold");

        __Ownable_init(_owner);
        __UUPSUpgradeable_init();
        __Pausable_init();

        maxPriceAge = _maxPriceAge;
        priceDeviationThreshold = _priceDeviationThreshold;

        // Initialize default feeds
        priceFeeds[ETH_USD_FEED].isActive = true;
        priceFeeds[CPOP_USD_FEED].isActive = true;

        emit OracleConfigUpdated(_maxPriceAge, _priceDeviationThreshold);
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getETHPriceUSD() external view override returns (uint256 price, uint256 timestamp) {
        PriceFeed memory feed = priceFeeds[ETH_USD_FEED];
        
        if (_isPriceFresh(feed.timestamp) && feed.price > 0) {
            return (feed.price, feed.timestamp);
        }
        
        // Use fallback price if main price is stale
        if (feed.fallbackPrice > 0) {
            return (feed.fallbackPrice, block.timestamp);
        }
        
        revert("GasPriceOracle: ETH price unavailable");
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getTokenPriceUSD() external view override returns (uint256 price, uint256 timestamp) {
        PriceFeed memory feed = priceFeeds[CPOP_USD_FEED];
        
        if (_isPriceFresh(feed.timestamp) && feed.price > 0) {
            return (feed.price, feed.timestamp);
        }
        
        // Use fallback price if main price is stale
        if (feed.fallbackPrice > 0) {
            return (feed.fallbackPrice, block.timestamp);
        }
        
        revert("GasPriceOracle: Token price unavailable");
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getTokenForETH(uint256 weiAmount) external view override returns (uint256 tokenAmount) {
        (uint256 ethPriceUSD,) = this.getETHPriceUSD();
        (uint256 tokenPriceUSD,) = this.getTokenPriceUSD();
        
        require(tokenPriceUSD > 0, "GasPriceOracle: invalid Token price");
        
        // Calculate: weiAmount * ethPriceUSD / cpopPriceUSD
        // Note: ETH has 18 decimals, prices have 8 decimals
        uint256 ethValueUSD = (weiAmount * ethPriceUSD) / 1e18;
        tokenAmount = (ethValueUSD * 1e18) / tokenPriceUSD;
        
        return tokenAmount;
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getETHForToken(uint256 tokenAmount) external view override returns (uint256 weiAmount) {
        (uint256 ethPriceUSD,) = this.getETHPriceUSD();
        (uint256 tokenPriceUSD,) = this.getTokenPriceUSD();
        
        require(ethPriceUSD > 0, "GasPriceOracle: invalid ETH price");
        
        // Calculate: tokenAmount * tokenPriceUSD / ethPriceUSD
        uint256 tokenValueUSD = (tokenAmount * tokenPriceUSD) / 1e18;
        weiAmount = (tokenValueUSD * 1e18) / ethPriceUSD;
        
        return weiAmount;
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function estimateGasCostInToken(
        uint256 gasLimit,
        uint256 gasPrice
    ) external view override returns (uint256 tokenCost) {
        uint256 totalGasCostWei = gasLimit * gasPrice;
        return this.getTokenForETH(totalGasCostWei);
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function isPriceValid(string calldata feedType) external view override returns (bool isValid) {
        PriceFeed memory feed = priceFeeds[feedType];
        return feed.isActive && _isPriceFresh(feed.timestamp) && feed.price > 0;
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getPriceWithMetadata(string calldata feedType) 
        external 
        view 
        override 
        returns (uint256 price, uint256 timestamp, bool isValid) 
    {
        PriceFeed memory feed = priceFeeds[feedType];
        price = feed.price;
        timestamp = feed.timestamp;
        isValid = feed.isActive && _isPriceFresh(feed.timestamp) && feed.price > 0;
        
        // If main price is invalid, try fallback
        if (!isValid && feed.fallbackPrice > 0) {
            price = feed.fallbackPrice;
            timestamp = block.timestamp;
            isValid = true;
        }
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function updatePriceFeed(
        string calldata feedType, 
        uint256 newPrice
    ) external override onlyAuthorizedOracle whenNotPaused {
        require(newPrice > 0, "GasPriceOracle: invalid price");
        require(priceFeeds[feedType].isActive, "GasPriceOracle: feed not active");
        
        PriceFeed storage feed = priceFeeds[feedType];
        uint256 oldPrice = feed.price;
        
        // Validate price deviation if there's a previous price
        if (oldPrice > 0 && _isPriceFresh(feed.timestamp)) {
            uint256 deviation = _calculatePriceDeviation(oldPrice, newPrice);
            require(
                deviation <= priceDeviationThreshold,
                "GasPriceOracle: price deviation too high"
            );
        }
        
        feed.price = newPrice;
        feed.timestamp = block.timestamp;
        
        emit PriceFeedUpdated(feedType, oldPrice, newPrice, block.timestamp);
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function setFallbackPrice(
        string calldata feedType, 
        uint256 fallbackPrice
    ) external override onlyOwner {
        require(fallbackPrice > 0, "GasPriceOracle: invalid fallback price");
        priceFeeds[feedType].fallbackPrice = fallbackPrice;
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getMaxPriceAge() external view override returns (uint256 maxAge) {
        return maxPriceAge;
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getPriceDeviationThreshold() external view override returns (uint256 threshold) {
        return priceDeviationThreshold;
    }

    /**
     * @notice Add new price feed
     * @param feedType Type of price feed
     * @param initialPrice Initial price
     * @param fallbackPrice Fallback price
     */
    function addPriceFeed(
        string calldata feedType,
        uint256 initialPrice,
        uint256 fallbackPrice
    ) external onlyOwner {
        require(initialPrice > 0, "GasPriceOracle: invalid initial price");
        require(!priceFeeds[feedType].isActive, "GasPriceOracle: feed already exists");
        
        priceFeeds[feedType] = PriceFeed({
            price: initialPrice,
            timestamp: block.timestamp,
            fallbackPrice: fallbackPrice,
            isActive: true
        });
    }

    /**
     * @notice Remove price feed
     * @param feedType Type of price feed to remove
     */
    function removePriceFeed(string calldata feedType) external onlyOwner {
        require(priceFeeds[feedType].isActive, "GasPriceOracle: feed not active");
        delete priceFeeds[feedType];
    }

    /**
     * @notice Authorize oracle address
     * @param oracle Oracle address to authorize
     */
    function authorizeOracle(address oracle) external onlyOwner {
        require(oracle != address(0), "GasPriceOracle: invalid oracle");
        require(!authorizedOracles[oracle], "GasPriceOracle: already authorized");
        
        authorizedOracles[oracle] = true;
    }

    /**
     * @notice Revoke oracle authorization
     * @param oracle Oracle address to revoke
     */
    function revokeOracle(address oracle) external onlyOwner {
        require(authorizedOracles[oracle], "GasPriceOracle: not authorized");
        authorizedOracles[oracle] = false;
    }

    /**
     * @notice Update oracle configuration
     * @param _maxPriceAge New maximum price age
     * @param _priceDeviationThreshold New price deviation threshold
     */
    function updateOracleConfig(
        uint256 _maxPriceAge,
        uint256 _priceDeviationThreshold
    ) external onlyOwner {
        require(_maxPriceAge > 0, "GasPriceOracle: invalid max price age");
        require(_priceDeviationThreshold <= BASIS_POINTS, "GasPriceOracle: invalid threshold");
        
        maxPriceAge = _maxPriceAge;
        priceDeviationThreshold = _priceDeviationThreshold;
        
        emit OracleConfigUpdated(_maxPriceAge, _priceDeviationThreshold);
    }

    /**
     * @notice Pause oracle operations
     */
    function pause() external onlyOwner {
        _pause();
    }

    /**
     * @notice Unpause oracle operations
     */
    function unpause() external onlyOwner {
        _unpause();
    }

    /**
     * @notice Check if price timestamp is fresh
     * @param timestamp Price timestamp to check
     * @return fresh True if price is fresh
     */
    function _isPriceFresh(uint256 timestamp) internal view returns (bool fresh) {
        return timestamp > 0 && (block.timestamp - timestamp) <= maxPriceAge;
    }

    /**
     * @notice Calculate price deviation percentage
     * @param oldPrice Previous price
     * @param newPrice New price
     * @return deviation Deviation in basis points
     */
    function _calculatePriceDeviation(
        uint256 oldPrice, 
        uint256 newPrice
    ) internal pure returns (uint256 deviation) {
        uint256 diff = oldPrice > newPrice ? oldPrice - newPrice : newPrice - oldPrice;
        return (diff * BASIS_POINTS) / oldPrice;
    }

    /**
     * @notice Authorize contract upgrades
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {
        // Additional upgrade logic can be added here if needed
    }
}