// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "../cpop/interfaces/IGasPriceOracle.sol";

/**
 * @title MockGasPriceOracle
 * @notice Mock Oracle for testing CPOP Paymaster functionality
 * @dev Provides configurable mock prices for testing environments
 */
contract MockGasPriceOracle is IGasPriceOracle {
    struct MockPriceFeed {
        uint256 price;
        uint256 timestamp;
        bool isValid;
    }

    mapping(string => MockPriceFeed) private mockFeeds;
    mapping(address => bool) public authorized;
    
    uint256 public maxPriceAge = 3600; // 1 hour
    uint256 public priceDeviationThreshold = 1000; // 10%
    
    address public owner;
    bool public simulateFailure;
    bool public returnStaleData;

    modifier onlyOwner() {
        require(msg.sender == owner, "MockOracle: not owner");
        _;
    }

    modifier onlyAuthorized() {
        require(authorized[msg.sender] || msg.sender == owner, "MockOracle: not authorized");
        _;
    }

    constructor() {
        owner = msg.sender;
        authorized[msg.sender] = true;
        
        // Set default test prices
        _setMockPrice("ETH/USD", 2000e8, true); // $2000 ETH
        _setMockPrice("CPOP/USD", 5e6, true);   // $0.05 CPOP
    }

    function setOwner(address newOwner) external onlyOwner {
        owner = newOwner;
    }

    function setAuthorized(address user, bool isAuthorized) external onlyOwner {
        authorized[user] = isAuthorized;
    }

    /**
     * @notice Set mock price for testing
     */
    function setMockPrice(string calldata feedType, uint256 price, bool isValid) external onlyAuthorized {
        _setMockPrice(feedType, price, isValid);
    }

    function _setMockPrice(string memory feedType, uint256 price, bool isValid) internal {
        mockFeeds[feedType] = MockPriceFeed({
            price: price,
            timestamp: block.timestamp,
            isValid: isValid
        });
    }

    /**
     * @notice Configure Oracle behavior for testing
     */
    function configureTestBehavior(bool _simulateFailure, bool _returnStaleData) external onlyOwner {
        simulateFailure = _simulateFailure;
        returnStaleData = _returnStaleData;
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getETHPriceUSD() external view override returns (uint256 price, uint256 timestamp) {
        if (simulateFailure) {
            revert("MockOracle: simulated failure");
        }
        
        MockPriceFeed memory feed = mockFeeds["ETH/USD"];
        price = feed.price;
        timestamp = returnStaleData ? block.timestamp - maxPriceAge - 1 : feed.timestamp;
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getTokenPriceUSD() external view override returns (uint256 price, uint256 timestamp) {
        if (simulateFailure) {
            revert("MockOracle: simulated failure");
        }
        
        MockPriceFeed memory feed = mockFeeds["TOKEN/USD"];
        price = feed.price;
        timestamp = returnStaleData ? block.timestamp - maxPriceAge - 1 : feed.timestamp;
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getTokenForETH(uint256 weiAmount) external view override returns (uint256 tokenAmount) {
        if (simulateFailure) {
            revert("MockOracle: simulated failure");
        }
        
        MockPriceFeed memory ethFeed = mockFeeds["ETH/USD"];
        MockPriceFeed memory tokenFeed = mockFeeds["TOKEN/USD"];
        
        require(ethFeed.isValid && tokenFeed.isValid, "MockOracle: invalid price data");
        require(tokenFeed.price > 0, "MockOracle: invalid Token price");
        
        // Calculate: weiAmount * ethPriceUSD / tokenPriceUSD
        uint256 ethValueUSD = (weiAmount * ethFeed.price) / 1e18;
        tokenAmount = (ethValueUSD * 1e18) / tokenFeed.price;
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getETHForToken(uint256 tokenAmount) external view override returns (uint256 weiAmount) {
        if (simulateFailure) {
            revert("MockOracle: simulated failure");
        }
        
        MockPriceFeed memory ethFeed = mockFeeds["ETH/USD"];
        MockPriceFeed memory tokenFeed = mockFeeds["TOKEN/USD"];
        
        require(ethFeed.isValid && tokenFeed.isValid, "MockOracle: invalid price data");
        require(ethFeed.price > 0, "MockOracle: invalid ETH price");
        
        // Calculate: tokenAmount * tokenPriceUSD / ethPriceUSD
        uint256 tokenValueUSD = (tokenAmount * tokenFeed.price) / 1e18;
        weiAmount = (tokenValueUSD * 1e18) / ethFeed.price;
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
        if (simulateFailure) {
            return false;
        }
        
        MockPriceFeed memory feed = mockFeeds[feedType];
        return feed.isValid && 
               feed.price > 0 && 
               (!returnStaleData ? (block.timestamp - feed.timestamp) <= maxPriceAge : false);
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
        MockPriceFeed memory feed = mockFeeds[feedType];
        price = feed.price;
        timestamp = returnStaleData ? block.timestamp - maxPriceAge - 1 : feed.timestamp;
        
        if (simulateFailure) {
            isValid = false;
        } else {
            isValid = feed.isValid && 
                     feed.price > 0 && 
                     (!returnStaleData ? (block.timestamp - timestamp) <= maxPriceAge : false);
        }
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function updatePriceFeed(string calldata feedType, uint256 newPrice) external override onlyAuthorized {
        require(newPrice > 0, "MockOracle: invalid price");
        _setMockPrice(feedType, newPrice, true);
        emit PriceFeedUpdated(feedType, mockFeeds[feedType].price, newPrice, block.timestamp);
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function setFallbackPrice(string calldata feedType, uint256 fallbackPrice) external override onlyOwner {
        _setMockPrice(feedType, fallbackPrice, true);
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getMaxPriceAge() external view override returns (uint256) {
        return maxPriceAge;
    }

    /**
     * @inheritdoc IGasPriceOracle
     */
    function getPriceDeviationThreshold() external view override returns (uint256) {
        return priceDeviationThreshold;
    }

    /**
     * @notice Set realistic test scenario
     */
    function setRealisticTestScenario() external onlyOwner {
        _setMockPrice("ETH/USD", 2000e8, true);
        _setMockPrice("CPOP/USD", 5e6, true);
        simulateFailure = false;
        returnStaleData = false;
    }

    /**
     * @notice Set bear market scenario
     */
    function setBearMarketScenario() external onlyOwner {
        _setMockPrice("ETH/USD", 1200e8, true);
        _setMockPrice("CPOP/USD", 2e6, true);
        simulateFailure = false;
        returnStaleData = false;
    }

    /**
     * @notice Simulate Oracle downtime
     */
    function simulateDowntime() external onlyOwner {
        simulateFailure = true;
    }

    /**
     * @notice Restore Oracle functionality
     */
    function restoreService() external onlyOwner {
        simulateFailure = false;
        returnStaleData = false;
    }

    /**
     * @notice Get current CPOP per ETH rate
     */
    function getCurrentCPOPPerETHRate() external view returns (uint256 cpopPerETH) {
        MockPriceFeed memory ethFeed = mockFeeds["ETH/USD"];
        MockPriceFeed memory cpopFeed = mockFeeds["CPOP/USD"];
        
        if (ethFeed.price > 0 && cpopFeed.price > 0) {
            cpopPerETH = (ethFeed.price * 1e18) / cpopFeed.price;
        }
    }
}