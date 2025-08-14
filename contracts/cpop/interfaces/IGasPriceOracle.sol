// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title IGasPriceOracle
 * @notice Interface for gas price oracle that provides ETH/Token exchange rates
 * @dev Supports multiple price feeds and fallback mechanisms
 */
interface IGasPriceOracle {
    /**
     * @notice Price feed update event
     * @param feedType The type of price feed (e.g., "ETH/USD", "TOKEN/USD")
     * @param oldPrice Previous price
     * @param newPrice Updated price
     * @param timestamp Update timestamp
     */
    event PriceFeedUpdated(
        string indexed feedType,
        uint256 oldPrice,
        uint256 newPrice,
        uint256 timestamp
    );

    /**
     * @notice Oracle configuration update event
     * @param maxPriceAge Maximum allowed price age in seconds
     * @param priceDeviationThreshold Maximum allowed price deviation percentage
     */
    event OracleConfigUpdated(
        uint256 maxPriceAge,
        uint256 priceDeviationThreshold
    );

    /**
     * @notice Fallback price used event
     * @param feedType The feed type that used fallback
     * @param fallbackPrice The fallback price used
     * @param reason Reason for using fallback
     */
    event FallbackPriceUsed(
        string indexed feedType,
        uint256 fallbackPrice,
        string reason
    );

    /**
     * @notice Get current ETH price in USD (with 8 decimals)
     * @return price ETH price in USD
     * @return timestamp Last update timestamp
     */
    function getETHPriceUSD() external view returns (uint256 price, uint256 timestamp);

    /**
     * @notice Get current Token price in USD (with 8 decimals)
     * @return price Token price in USD
     * @return timestamp Last update timestamp
     */
    function getTokenPriceUSD() external view returns (uint256 price, uint256 timestamp);

    /**
     * @notice Get Token amount needed to pay for wei amount of ETH
     * @param weiAmount Amount in wei
     * @return tokenAmount Amount of tokens needed
     */
    function getTokenForETH(uint256 weiAmount) external view returns (uint256 tokenAmount);

    /**
     * @notice Get ETH amount equivalent to Token amount
     * @param tokenAmount Amount of tokens
     * @return weiAmount Equivalent amount in wei
     */
    function getETHForToken(uint256 tokenAmount) external view returns (uint256 weiAmount);

    /**
     * @notice Estimate gas cost in tokens
     * @param gasLimit Gas limit for the transaction
     * @param gasPrice Gas price in wei
     * @return tokenCost Cost in tokens
     */
    function estimateGasCostInToken(
        uint256 gasLimit,
        uint256 gasPrice
    ) external view returns (uint256 tokenCost);

    /**
     * @notice Check if price data is fresh and valid
     * @param feedType Type of price feed to check
     * @return isValid True if price is valid and fresh
     */
    function isPriceValid(string calldata feedType) external view returns (bool isValid);

    /**
     * @notice Get price with metadata
     * @param feedType Type of price feed
     * @return price Current price
     * @return timestamp Last update timestamp
     * @return isValid Whether the price is valid
     */
    function getPriceWithMetadata(string calldata feedType) 
        external 
        view 
        returns (uint256 price, uint256 timestamp, bool isValid);

    /**
     * @notice Update price feed (restricted to authorized oracles)
     * @param feedType Type of price feed to update
     * @param newPrice New price value
     */
    function updatePriceFeed(string calldata feedType, uint256 newPrice) external;

    /**
     * @notice Set fallback price for a feed type
     * @param feedType Type of price feed
     * @param fallbackPrice Fallback price to use when main feed fails
     */
    function setFallbackPrice(string calldata feedType, uint256 fallbackPrice) external;

    /**
     * @notice Get maximum allowed price age
     * @return maxAge Maximum age in seconds
     */
    function getMaxPriceAge() external view returns (uint256 maxAge);

    /**
     * @notice Get price deviation threshold
     * @return threshold Deviation threshold as percentage (e.g., 500 = 5%)
     */
    function getPriceDeviationThreshold() external view returns (uint256 threshold);
}