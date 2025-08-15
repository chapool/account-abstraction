/**
 * @fileoverview Off-chain gas calculation utility for CPOP token payments
 * @description Provides utilities to calculate gas costs in CPOP tokens using Oracle data
 */

const { ethers } = require('ethers');

/**
 * Gas Calculator for CPOP payments
 */
class CPOPGasCalculator {
    constructor(provider, oracleAddress, oracleAbi) {
        this.provider = provider;
        this.oracle = new ethers.Contract(oracleAddress, oracleAbi, provider);
        this.cache = new Map();
        this.cacheTimeout = 30000; // 30 seconds cache
    }

    /**
     * Get current gas price from network
     * @returns {Promise<bigint>} Current gas price in wei
     */
    async getCurrentGasPrice() {
        const feeData = await this.provider.getFeeData();
        return feeData.gasPrice || feeData.maxFeePerGas;
    }

    /**
     * Estimate gas limit for a transaction
     * @param {Object} transaction - Transaction object
     * @returns {Promise<bigint>} Estimated gas limit
     */
    async estimateGasLimit(transaction) {
        try {
            return await this.provider.estimateGas(transaction);
        } catch (error) {
            console.warn('Gas estimation failed, using default:', error.message);
            return BigInt(21000); // Default gas limit
        }
    }

    /**
     * Get ETH price in USD from oracle with caching
     * @returns {Promise<{price: bigint, timestamp: bigint}>}
     */
    async getETHPriceUSD() {
        const cacheKey = 'eth_price_usd';
        const cached = this.cache.get(cacheKey);
        
        if (cached && Date.now() - cached.timestamp < this.cacheTimeout) {
            return cached.data;
        }

        try {
            const [price, timestamp] = await this.oracle.getETHPriceUSD();
            const result = { price, timestamp };
            
            this.cache.set(cacheKey, {
                data: result,
                timestamp: Date.now()
            });
            
            return result;
        } catch (error) {
            console.error('Failed to get ETH price from oracle:', error);
            throw new Error('Oracle price fetch failed');
        }
    }

    /**
     * Get CPOP price in USD from oracle with caching
     * @returns {Promise<{price: bigint, timestamp: bigint}>}
     */
    async getCPOPPriceUSD() {
        const cacheKey = 'cpop_price_usd';
        const cached = this.cache.get(cacheKey);
        
        if (cached && Date.now() - cached.timestamp < this.cacheTimeout) {
            return cached.data;
        }

        try {
            const [price, timestamp] = await this.oracle.getCPOPPriceUSD();
            const result = { price, timestamp };
            
            this.cache.set(cacheKey, {
                data: result,
                timestamp: Date.now()
            });
            
            return result;
        } catch (error) {
            console.error('Failed to get CPOP price from oracle:', error);
            throw new Error('Oracle price fetch failed');
        }
    }

    /**
     * Calculate gas cost in CPOP tokens
     * @param {bigint} gasLimit - Gas limit for the transaction
     * @param {bigint} gasPrice - Gas price in wei
     * @returns {Promise<bigint>} Gas cost in CPOP tokens
     */
    async calculateGasCostInCPOP(gasLimit, gasPrice) {
        try {
            return await this.oracle.estimateGasCostInCPOP(gasLimit, gasPrice);
        } catch (error) {
            console.warn('Oracle gas calculation failed, using fallback calculation:', error.message);
            return await this.fallbackGasCostCalculation(gasLimit, gasPrice);
        }
    }

    /**
     * Fallback gas cost calculation when oracle is unavailable
     * @param {bigint} gasLimit - Gas limit
     * @param {bigint} gasPrice - Gas price in wei
     * @returns {Promise<bigint>} Gas cost in CPOP tokens
     */
    async fallbackGasCostCalculation(gasLimit, gasPrice) {
        const totalGasCostWei = gasLimit * gasPrice;
        
        try {
            const ethPrice = await this.getETHPriceUSD();
            const cpopPrice = await this.getCPOPPriceUSD();
            
            // Calculate ETH value in USD (considering 18 decimals for ETH, 8 for price)
            const ethValueUSD = (totalGasCostWei * ethPrice.price) / BigInt(10 ** 18);
            
            // Calculate CPOP amount needed (considering 18 decimals for CPOP, 8 for price)
            const cpopAmount = (ethValueUSD * BigInt(10 ** 18)) / cpopPrice.price;
            
            return cpopAmount;
        } catch (error) {
            console.error('Fallback calculation failed:', error);
            throw new Error('Unable to calculate gas cost in CPOP');
        }
    }

    /**
     * Estimate total transaction cost for a user operation
     * @param {Object} userOp - User operation object
     * @returns {Promise<Object>} Cost breakdown
     */
    async estimateUserOpCost(userOp) {
        try {
            const gasPrice = await this.getCurrentGasPrice();
            
            // Extract gas limits from user operation
            const callGasLimit = BigInt(userOp.callGasLimit || 0);
            const verificationGasLimit = BigInt(userOp.verificationGasLimit || 0);
            const preVerificationGas = BigInt(userOp.preVerificationGas || 0);
            
            // Calculate total gas limit
            const totalGasLimit = callGasLimit + verificationGasLimit + preVerificationGas;
            
            // Calculate costs
            const totalGasCostWei = totalGasLimit * gasPrice;
            const cpopCost = await this.calculateGasCostInCPOP(totalGasLimit, gasPrice);
            
            return {
                gasLimit: {
                    call: callGasLimit,
                    verification: verificationGasLimit,
                    preVerification: preVerificationGas,
                    total: totalGasLimit
                },
                gasPrice: gasPrice,
                costs: {
                    wei: totalGasCostWei,
                    eth: ethers.formatEther(totalGasCostWei),
                    cpop: ethers.formatEther(cpopCost),
                    cpopWei: cpopCost
                },
                timestamp: Date.now()
            };
        } catch (error) {
            console.error('Failed to estimate user operation cost:', error);
            throw error;
        }
    }

    /**
     * Check if user has sufficient CPOP balance for transaction
     * @param {string} userAddress - User's address
     * @param {Object} userOp - User operation
     * @param {Object} cpopToken - CPOP token contract instance
     * @returns {Promise<Object>} Balance check result
     */
    async checkUserBalance(userAddress, userOp, cpopToken) {
        try {
            const costEstimate = await this.estimateUserOpCost(userOp);
            const userBalance = await cpopToken.balanceOf(userAddress);
            
            const hasSufficientBalance = userBalance >= costEstimate.costs.cpopWei;
            
            return {
                userBalance: ethers.formatEther(userBalance),
                userBalanceWei: userBalance,
                requiredCPOP: costEstimate.costs.cpop,
                requiredCPOPWei: costEstimate.costs.cpopWei,
                hasSufficientBalance,
                shortfall: hasSufficientBalance ? 
                    BigInt(0) : 
                    costEstimate.costs.cpopWei - userBalance,
                costEstimate
            };
        } catch (error) {
            console.error('Failed to check user balance:', error);
            throw error;
        }
    }

    /**
     * Calculate optimal gas parameters for cost efficiency
     * @param {Object} transaction - Transaction details
     * @returns {Promise<Object>} Optimized gas parameters
     */
    async optimizeGasParameters(transaction) {
        try {
            const currentGasPrice = await this.getCurrentGasPrice();
            const estimatedGasLimit = await this.estimateGasLimit(transaction);
            
            // Calculate costs for different gas price scenarios
            const scenarios = [
                { name: 'slow', multiplier: 0.8 },
                { name: 'standard', multiplier: 1.0 },
                { name: 'fast', multiplier: 1.2 },
                { name: 'fastest', multiplier: 1.5 }
            ];
            
            const optimizations = await Promise.all(
                scenarios.map(async (scenario) => {
                    const gasPrice = BigInt(Math.floor(Number(currentGasPrice) * scenario.multiplier));
                    const cpopCost = await this.calculateGasCostInCPOP(estimatedGasLimit, gasPrice);
                    
                    return {
                        speed: scenario.name,
                        gasPrice: gasPrice,
                        gasPriceGwei: ethers.formatUnits(gasPrice, 'gwei'),
                        gasLimit: estimatedGasLimit,
                        totalCostWei: estimatedGasLimit * gasPrice,
                        totalCostETH: ethers.formatEther(estimatedGasLimit * gasPrice),
                        cpopCost: ethers.formatEther(cpopCost),
                        cpopCostWei: cpopCost
                    };
                })
            );
            
            return {
                current: {
                    gasPrice: currentGasPrice,
                    gasPriceGwei: ethers.formatUnits(currentGasPrice, 'gwei')
                },
                optimizations,
                recommended: optimizations[1] // Standard speed
            };
        } catch (error) {
            console.error('Failed to optimize gas parameters:', error);
            throw error;
        }
    }

    /**
     * Get price feed health status
     * @returns {Promise<Object>} Oracle health status
     */
    async getOracleHealthStatus() {
        try {
            const [ethValid, cpopValid] = await Promise.all([
                this.oracle.isPriceValid('ETH/USD'),
                this.oracle.isPriceValid('CPOP/USD')
            ]);
            
            const [maxAge, deviationThreshold] = await Promise.all([
                this.oracle.getMaxPriceAge(),
                this.oracle.getPriceDeviationThreshold()
            ]);
            
            return {
                isHealthy: ethValid && cpopValid,
                feeds: {
                    'ETH/USD': ethValid,
                    'CPOP/USD': cpopValid
                },
                config: {
                    maxPriceAge: Number(maxAge),
                    deviationThreshold: Number(deviationThreshold)
                },
                timestamp: Date.now()
            };
        } catch (error) {
            console.error('Failed to get oracle health status:', error);
            return {
                isHealthy: false,
                error: error.message,
                timestamp: Date.now()
            };
        }
    }

    /**
     * Clear price cache
     */
    clearCache() {
        this.cache.clear();
    }

    /**
     * Set cache timeout
     * @param {number} timeout - Cache timeout in milliseconds
     */
    setCacheTimeout(timeout) {
        this.cacheTimeout = timeout;
    }
}

module.exports = {
    CPOPGasCalculator
};