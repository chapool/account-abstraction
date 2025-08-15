/**
 * @fileoverview Web2 Wallet Hybrid Deployment Service
 * @description Supports both immediate deployment and EntryPoint-based lazy deployment for Web2 users
 */

const { ethers } = require('ethers');

/**
 * Web2 Wallet Deployment Service with hybrid deployment modes
 */
class Web2WalletDeployer {
    constructor(provider, walletManagerAddress, walletManagerAbi, config = {}) {
        this.provider = provider;
        this.walletManager = new ethers.Contract(walletManagerAddress, walletManagerAbi, provider);
        
        // Configuration
        this.config = {
            preferLazyDeployment: config.preferLazyDeployment !== false, // Default: true
            gasThreshold: config.gasThreshold || ethers.parseUnits('50', 'gwei'), // 50 Gwei threshold
            highValueThreshold: config.highValueThreshold || ethers.parseEther('1'), // 1 ETH
            ...config
        };
    }

    /**
     * Get deployment strategy for a Web2 user
     * @param {string} identifier - User identifier
     * @param {string} masterSigner - Master signer address
     * @param {Object} options - Additional options
     * @returns {Promise<Object>} Deployment strategy recommendation
     */
    async getDeploymentStrategy(identifier, masterSigner, options = {}) {
        try {
            // Check if account is already deployed
            const isDeployed = await this.walletManager.isWeb2AccountDeployed(identifier, masterSigner);
            if (isDeployed) {
                const accountAddress = await this.walletManager.getWeb2AccountAddress(identifier, masterSigner);
                return {
                    strategy: 'already_deployed',
                    accountAddress,
                    isDeployed: true,
                    initCode: '0x',
                    estimatedGas: 0
                };
            }

            // Get current gas price
            const currentGasPrice = await this.provider.getFeeData().then(data => 
                data.gasPrice || data.maxFeePerGas
            );

            // Calculate deployment cost
            const deploymentGas = await this.estimateDeploymentGas(identifier, masterSigner);
            const deploymentCost = deploymentGas * currentGasPrice;

            // Determine strategy based on conditions
            let recommendedStrategy;
            let reasoning = [];

            if (options.forceImmediate) {
                recommendedStrategy = 'immediate';
                reasoning.push('Force immediate deployment requested');
            } else if (options.userValue && options.userValue > this.config.highValueThreshold) {
                recommendedStrategy = 'immediate';
                reasoning.push('High value user - immediate deployment for better UX');
            } else if (currentGasPrice > this.config.gasThreshold) {
                recommendedStrategy = 'lazy';
                reasoning.push('High gas price - use lazy deployment');
            } else if (options.expectsMultipleOps) {
                recommendedStrategy = 'immediate';
                reasoning.push('Multiple operations expected - deploy immediately');
            } else if (this.config.preferLazyDeployment) {
                recommendedStrategy = 'lazy';
                reasoning.push('Default lazy deployment for cost optimization');
            } else {
                recommendedStrategy = 'immediate';
                reasoning.push('Immediate deployment for guaranteed availability');
            }

            const accountAddress = await this.walletManager.getWeb2AccountAddress(identifier, masterSigner);
            const initCode = recommendedStrategy === 'lazy' ? 
                await this.walletManager.getWeb2InitCode(identifier, masterSigner) : '0x';

            return {
                strategy: recommendedStrategy,
                accountAddress,
                isDeployed: false,
                initCode,
                estimatedGas: deploymentGas,
                estimatedCost: deploymentCost,
                currentGasPrice,
                reasoning,
                alternatives: {
                    immediate: {
                        cost: deploymentCost,
                        gasEstimate: deploymentGas,
                        pros: ['Immediate availability', 'No first-tx delay'],
                        cons: ['Higher upfront cost', 'Potential gas waste']
                    },
                    lazy: {
                        cost: deploymentCost, // Same cost, different timing
                        gasEstimate: deploymentGas,
                        pros: ['Cost optimization', 'Only deploy when needed'],
                        cons: ['First transaction delay', 'Requires initCode']
                    }
                }
            };
        } catch (error) {
            console.error('Failed to get deployment strategy:', error);
            throw error;
        }
    }

    /**
     * Deploy Web2 wallet immediately
     * @param {string} identifier - User identifier
     * @param {string} masterSigner - Master signer address
     * @param {Object} signer - Ethereum signer (authorized creator)
     * @returns {Promise<Object>} Deployment result
     */
    async deployImmediately(identifier, masterSigner, signer) {
        try {
            console.log(`🚀 Deploying Web2 wallet immediately for ${identifier}...`);
            
            const walletManagerWithSigner = this.walletManager.connect(signer);
            
            // Use the legacy method for immediate deployment
            const tx = await walletManagerWithSigner.createWeb2AccountWithMasterSigner(
                identifier, 
                masterSigner
            );
            
            console.log(`📝 Transaction submitted: ${tx.hash}`);
            const receipt = await tx.wait();
            
            // Extract account address from events
            const accountCreatedEvent = receipt.logs.find(log => {
                try {
                    const parsed = walletManagerWithSigner.interface.parseLog(log);
                    return parsed.name === 'AccountCreated';
                } catch {
                    return false;
                }
            });
            
            let accountAddress;
            if (accountCreatedEvent) {
                const parsed = walletManagerWithSigner.interface.parseLog(accountCreatedEvent);
                accountAddress = parsed.args.account;
            } else {
                // Fallback: calculate address
                accountAddress = await this.walletManager.getWeb2AccountAddress(identifier, masterSigner);
            }
            
            console.log(`✅ Wallet deployed at: ${accountAddress}`);
            
            return {
                success: true,
                deploymentStrategy: 'immediate',
                accountAddress,
                transactionHash: tx.hash,
                gasUsed: receipt.gasUsed,
                effectiveGasPrice: receipt.effectiveGasPrice,
                deploymentCost: receipt.gasUsed * receipt.effectiveGasPrice,
                blockNumber: receipt.blockNumber
            };
        } catch (error) {
            console.error('❌ Immediate deployment failed:', error);
            throw error;
        }
    }

    /**
     * Prepare lazy deployment for Web2 wallet
     * @param {string} identifier - User identifier  
     * @param {string} masterSigner - Master signer address
     * @returns {Promise<Object>} Lazy deployment setup
     */
    async prepareLazyDeployment(identifier, masterSigner) {
        try {
            console.log(`⏳ Preparing lazy deployment for ${identifier}...`);
            
            const accountAddress = await this.walletManager.getWeb2AccountAddress(identifier, masterSigner);
            const initCode = await this.walletManager.getWeb2InitCode(identifier, masterSigner);
            
            console.log(`📍 Predicted address: ${accountAddress}`);
            console.log(`🔧 InitCode generated (${initCode.length / 2 - 1} bytes)`);
            
            return {
                success: true,
                deploymentStrategy: 'lazy',
                accountAddress,
                initCode,
                isDeployed: false,
                instructions: {
                    usage: 'Include initCode in first UserOperation',
                    note: 'Account will be deployed automatically on first transaction'
                }
            };
        } catch (error) {
            console.error('❌ Lazy deployment preparation failed:', error);
            throw error;
        }
    }

    /**
     * Smart deployment - automatically choose best strategy
     * @param {string} identifier - User identifier
     * @param {string} masterSigner - Master signer address
     * @param {Object} signer - Ethereum signer (for immediate deployment)
     * @param {Object} options - Deployment options
     * @returns {Promise<Object>} Deployment result
     */
    async smartDeploy(identifier, masterSigner, signer, options = {}) {
        try {
            const strategy = await this.getDeploymentStrategy(identifier, masterSigner, options);
            
            console.log(`🎯 Recommended strategy: ${strategy.strategy}`);
            console.log(`💭 Reasoning: ${strategy.reasoning.join(', ')}`);
            
            if (strategy.strategy === 'already_deployed') {
                console.log(`✅ Account already deployed at: ${strategy.accountAddress}`);
                return strategy;
            }
            
            if (strategy.strategy === 'immediate') {
                return await this.deployImmediately(identifier, masterSigner, signer);
            } else {
                return await this.prepareLazyDeployment(identifier, masterSigner);
            }
        } catch (error) {
            console.error('❌ Smart deployment failed:', error);
            throw error;
        }
    }

    /**
     * Estimate gas cost for deployment
     * @param {string} identifier - User identifier
     * @param {string} masterSigner - Master signer address
     * @returns {Promise<bigint>} Estimated gas
     */
    async estimateDeploymentGas(identifier, masterSigner) {
        try {
            // This is an estimation - actual gas may vary
            // Based on proxy deployment + initialization
            return BigInt(250000); // Approximate gas for account deployment
        } catch (error) {
            console.error('Failed to estimate deployment gas:', error);
            return BigInt(300000); // Conservative fallback
        }
    }

    /**
     * Monitor gas prices and suggest optimal deployment timing
     * @param {number} durationMinutes - Monitoring duration in minutes
     * @returns {Promise<Object>} Gas monitoring result
     */
    async monitorGasPrices(durationMinutes = 60) {
        console.log(`📊 Monitoring gas prices for ${durationMinutes} minutes...`);
        
        const prices = [];
        const interval = Math.min(durationMinutes * 1000, 60000); // Max 1 minute intervals
        const iterations = Math.ceil((durationMinutes * 60 * 1000) / interval);
        
        for (let i = 0; i < iterations; i++) {
            try {
                const feeData = await this.provider.getFeeData();
                const gasPrice = feeData.gasPrice || feeData.maxFeePerGas;
                
                prices.push({
                    timestamp: Date.now(),
                    gasPrice: gasPrice,
                    gasPriceGwei: ethers.formatUnits(gasPrice, 'gwei')
                });
                
                console.log(`⛽ Gas price: ${ethers.formatUnits(gasPrice, 'gwei')} Gwei`);
                
                if (i < iterations - 1) {
                    await new Promise(resolve => setTimeout(resolve, interval));
                }
            } catch (error) {
                console.error('Error monitoring gas price:', error);
            }
        }
        
        // Analyze results
        const gasPrices = prices.map(p => p.gasPrice);
        const minPrice = gasPrices.reduce((min, price) => price < min ? price : min);
        const maxPrice = gasPrices.reduce((max, price) => price > max ? price : max);
        const avgPrice = gasPrices.reduce((sum, price) => sum + price, BigInt(0)) / BigInt(gasPrices.length);
        
        const recommendation = minPrice < this.config.gasThreshold ? 
            'Deploy now - gas prices are favorable' :
            'Wait for lower gas prices or use lazy deployment';
        
        return {
            monitoringDuration: durationMinutes,
            samplesCount: prices.length,
            prices,
            analysis: {
                min: {
                    price: minPrice,
                    gwei: ethers.formatUnits(minPrice, 'gwei')
                },
                max: {
                    price: maxPrice,
                    gwei: ethers.formatUnits(maxPrice, 'gwei')
                },
                average: {
                    price: avgPrice,
                    gwei: ethers.formatUnits(avgPrice, 'gwei')
                }
            },
            recommendation,
            shouldDeployNow: minPrice < this.config.gasThreshold
        };
    }

    /**
     * Get deployment statistics
     * @returns {Object} Configuration and thresholds
     */
    getConfig() {
        return {
            preferLazyDeployment: this.config.preferLazyDeployment,
            gasThreshold: {
                wei: this.config.gasThreshold,
                gwei: ethers.formatUnits(this.config.gasThreshold, 'gwei')
            },
            highValueThreshold: {
                wei: this.config.highValueThreshold,
                eth: ethers.formatEther(this.config.highValueThreshold)
            }
        };
    }
}

module.exports = {
    Web2WalletDeployer
};