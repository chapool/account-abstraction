/**
 * @fileoverview CPOP Master Signature Aggregation Tool
 * @description Provides utilities to create and manage aggregated signatures for multiple CPOP wallets
 */

const { ethers } = require('ethers');

/**
 * CPOP Signature Aggregator for Master Signer operations
 */
class CPOPSignatureAggregator {
    constructor(provider, aggregatorAddress, aggregatorAbi, masterPrivateKey) {
        this.provider = provider;
        this.aggregator = new ethers.Contract(aggregatorAddress, aggregatorAbi, provider);
        this.masterWallet = new ethers.Wallet(masterPrivateKey, provider);
        this.cache = new Map();
    }

    /**
     * Create aggregated signature for multiple user operations
     * @param {Array} userOps - Array of user operations
     * @param {Object} options - Additional options
     * @returns {Promise<Object>} Aggregated signature data
     */
    async createAggregatedSignature(userOps, options = {}) {
        try {
            console.log(`🔗 Creating aggregated signature for ${userOps.length} operations...`);
            
            // Validate user operations
            this._validateUserOperations(userOps);
            
            // Prepare operations for aggregation (remove individual signatures)
            const aggregatedOps = userOps.map(op => ({
                ...op,
                signature: '0x' // Empty signature for aggregated operations
            }));
            
            // Get current nonce for master signer
            const masterNonce = await this.aggregator.getMasterNonce(this.masterWallet.address);
            
            // Create aggregated hash
            const aggregatedHash = await this._createAggregatedHash(
                aggregatedOps, 
                this.masterWallet.address, 
                masterNonce
            );
            
            // Sign with master wallet
            const masterSignature = await this.masterWallet.signMessage(
                ethers.getBytes(aggregatedHash)
            );
            
            // Encode aggregated signature
            const encodedSignature = ethers.AbiCoder.defaultAbiCoder().encode(
                ['address', 'uint256', 'bytes'],
                [this.masterWallet.address, masterNonce, masterSignature]
            );
            
            const result = {
                userOps: aggregatedOps,
                aggregator: await this.aggregator.getAddress(),
                signature: encodedSignature,
                masterSigner: this.masterWallet.address,
                nonce: masterNonce,
                hash: aggregatedHash,
                gasEstimate: await this._estimateAggregatedGas(aggregatedOps.length)
            };
            
            console.log(`✅ Aggregated signature created successfully`);
            console.log(`   Master: ${result.masterSigner}`);
            console.log(`   Nonce: ${result.nonce}`);
            console.log(`   Operations: ${aggregatedOps.length}`);
            
            return result;
            
        } catch (error) {
            console.error('❌ Failed to create aggregated signature:', error);
            throw error;
        }
    }

    /**
     * Execute aggregated operations via EntryPoint
     * @param {Object} aggregatedData - Result from createAggregatedSignature
     * @param {Object} entryPoint - EntryPoint contract instance
     * @param {string} beneficiary - Beneficiary address for gas refunds
     * @returns {Promise<Object>} Transaction result
     */
    async executeAggregatedOperations(aggregatedData, entryPoint, beneficiary) {
        try {
            console.log(`🚀 Executing ${aggregatedData.userOps.length} aggregated operations...`);
            
            const opsPerAggregator = [{
                userOps: aggregatedData.userOps,
                aggregator: aggregatedData.aggregator,
                signature: aggregatedData.signature
            }];
            
            // Estimate gas for the aggregated transaction
            const gasEstimate = await entryPoint.estimateGas.handleAggregatedOps(
                opsPerAggregator,
                beneficiary
            );
            
            // Execute aggregated operations
            const tx = await entryPoint.handleAggregatedOps(
                opsPerAggregator,
                beneficiary,
                {
                    gasLimit: gasEstimate + BigInt(50000) // Add buffer
                }
            );
            
            console.log(`📝 Transaction submitted: ${tx.hash}`);
            const receipt = await tx.wait();
            
            console.log(`✅ Aggregated operations executed successfully`);
            console.log(`   Gas used: ${receipt.gasUsed.toLocaleString()}`);
            console.log(`   Block: ${receipt.blockNumber}`);
            
            return {
                success: true,
                transactionHash: tx.hash,
                receipt,
                gasUsed: receipt.gasUsed,
                operationsCount: aggregatedData.userOps.length,
                gasSavings: await this._calculateGasSavings(aggregatedData.userOps.length, receipt.gasUsed)
            };
            
        } catch (error) {
            console.error('❌ Failed to execute aggregated operations:', error);
            throw error;
        }
    }

    /**
     * Batch multiple operations from different wallets
     * @param {Array} operations - Array of {wallet, calls} objects
     * @param {Object} gasConfig - Gas configuration
     * @returns {Promise<Array>} Array of UserOperations ready for aggregation
     */
    async batchOperationsFromWallets(operations, gasConfig = {}) {
        try {
            console.log(`📦 Batching operations from ${operations.length} wallets...`);
            
            const userOps = [];
            const defaultGasConfig = {
                callGasLimit: '200000',
                verificationGasLimit: '150000',
                preVerificationGas: '21000',
                maxFeePerGas: ethers.parseUnits('20', 'gwei').toString(),
                maxPriorityFeePerGas: ethers.parseUnits('2', 'gwei').toString(),
                ...gasConfig
            };
            
            for (const operation of operations) {
                const { wallet, calls, options = {} } = operation;
                
                // Get wallet nonce
                const nonce = await this._getWalletNonce(wallet);
                
                // Encode batch call data
                const callData = await this._encodeBatchCalls(calls);
                
                // Create user operation
                const userOp = {
                    sender: wallet,
                    nonce: nonce.toString(),
                    initCode: '0x',
                    callData,
                    callGasLimit: options.callGasLimit || defaultGasConfig.callGasLimit,
                    verificationGasLimit: options.verificationGasLimit || defaultGasConfig.verificationGasLimit,
                    preVerificationGas: options.preVerificationGas || defaultGasConfig.preVerificationGas,
                    maxFeePerGas: options.maxFeePerGas || defaultGasConfig.maxFeePerGas,
                    maxPriorityFeePerGas: options.maxPriorityFeePerGas || defaultGasConfig.maxPriorityFeePerGas,
                    paymasterAndData: options.paymasterAndData || '0x',
                    signature: '0x' // Will be handled by aggregator
                };
                
                userOps.push(userOp);
            }
            
            console.log(`✅ Created ${userOps.length} user operations for aggregation`);
            return userOps;
            
        } catch (error) {
            console.error('❌ Failed to batch operations:', error);
            throw error;
        }
    }

    /**
     * Auto-authorize wallets with the aggregator
     * @param {Array} walletAddresses - Array of wallet addresses to authorize
     * @returns {Promise<Object>} Authorization result
     */
    async autoAuthorizeWallets(walletAddresses) {
        try {
            console.log(`🔐 Auto-authorizing ${walletAddresses.length} wallets...`);
            
            const aggregatorWithSigner = this.aggregator.connect(this.masterWallet);
            const results = [];
            
            for (const wallet of walletAddresses) {
                try {
                    // Check if wallet is controlled by master signer
                    const isControlled = await this.aggregator.isWalletControlledByMaster(
                        wallet, 
                        this.masterWallet.address
                    );
                    
                    if (isControlled) {
                        // Auto-authorize if not already authorized
                        const tx = await aggregatorWithSigner.autoAuthorizeWallet(
                            wallet, 
                            this.masterWallet.address
                        );
                        await tx.wait();
                        
                        results.push({ wallet, authorized: true, tx: tx.hash });
                        console.log(`  ✅ Authorized: ${wallet}`);
                    } else {
                        results.push({ wallet, authorized: false, reason: 'Not controlled by master' });
                        console.log(`  ❌ Skipped: ${wallet} (not controlled by master)`);
                    }
                } catch (error) {
                    results.push({ wallet, authorized: false, error: error.message });
                    console.log(`  ❌ Failed: ${wallet} - ${error.message}`);
                }
            }
            
            const successCount = results.filter(r => r.authorized).length;
            console.log(`✅ Authorization completed: ${successCount}/${walletAddresses.length} successful`);
            
            return {
                success: successCount > 0,
                totalWallets: walletAddresses.length,
                authorizedCount: successCount,
                results
            };
            
        } catch (error) {
            console.error('❌ Failed to auto-authorize wallets:', error);
            throw error;
        }
    }

    /**
     * Calculate gas savings from aggregation
     * @param {number} operationCount - Number of operations
     * @returns {Promise<Object>} Gas savings analysis
     */
    async calculateGasSavings(operationCount) {
        const savings = await this.aggregator.calculateGasSavings(operationCount);
        const currentGasPrice = await this.provider.getFeeData().then(data => 
            data.gasPrice || data.maxFeePerGas
        );
        
        return {
            gasSaved: savings,
            gasSavedFormatted: savings.toLocaleString(),
            estimatedETHSaved: ethers.formatEther(savings * currentGasPrice),
            currentGasPrice: ethers.formatUnits(currentGasPrice, 'gwei') + ' Gwei',
            operationCount,
            savingsPercentage: operationCount > 1 ? 
                Math.round((Number(savings) / (operationCount * 3000)) * 100) : 0
        };
    }

    /**
     * Monitor aggregator events
     * @param {Function} callback - Callback function for events
     */
    startEventMonitoring(callback) {
        console.log('📡 Starting aggregator event monitoring...');
        
        this.aggregator.on('AggregatedValidation', (master, operationCount, aggregatedHash, event) => {
            callback({
                type: 'AggregatedValidation',
                master,
                operationCount: Number(operationCount),
                aggregatedHash,
                blockNumber: event.blockNumber,
                transactionHash: event.transactionHash
            });
        });
        
        this.aggregator.on('MasterAuthorized', (master, authorized, event) => {
            callback({
                type: 'MasterAuthorized',
                master,
                authorized,
                blockNumber: event.blockNumber,
                transactionHash: event.transactionHash
            });
        });
        
        this.aggregator.on('WalletAuthorized', (master, wallet, authorized, event) => {
            callback({
                type: 'WalletAuthorized',
                master,
                wallet,
                authorized,
                blockNumber: event.blockNumber,
                transactionHash: event.transactionHash
            });
        });
    }

    /**
     * Validate user operations before aggregation
     * @param {Array} userOps - User operations to validate
     */
    _validateUserOperations(userOps) {
        if (!Array.isArray(userOps) || userOps.length === 0) {
            throw new Error('User operations must be a non-empty array');
        }
        
        if (userOps.length > 50) {
            throw new Error('Too many operations for aggregation (max 50)');
        }
        
        for (const op of userOps) {
            if (!op.sender || !ethers.isAddress(op.sender)) {
                throw new Error('Invalid sender address in user operation');
            }
            
            if (!op.callData) {
                throw new Error('Missing callData in user operation');
            }
        }
    }

    /**
     * Create aggregated hash for signing
     * @param {Array} userOps - User operations
     * @param {string} masterSigner - Master signer address
     * @param {number} nonce - Current nonce
     * @returns {Promise<string>} Aggregated hash
     */
    async _createAggregatedHash(userOps, masterSigner, nonce) {
        const opHashes = userOps.map(op => 
            ethers.keccak256(ethers.AbiCoder.defaultAbiCoder().encode(
                ['address', 'uint256', 'bytes', 'uint256', 'uint256', 'uint256', 'uint256', 'uint256', 'bytes'],
                [
                    op.sender,
                    op.nonce,
                    op.callData,
                    op.callGasLimit || '200000',
                    op.verificationGasLimit || '150000',
                    op.preVerificationGas || '21000',
                    op.maxFeePerGas || '20000000000',
                    op.maxPriorityFeePerGas || '2000000000',
                    op.paymasterAndData || '0x'
                ]
            ))
        );
        
        return ethers.keccak256(ethers.AbiCoder.defaultAbiCoder().encode(
            ['string', 'address', 'uint256', 'uint256', 'address', 'bytes32[]'],
            [
                'CPOP_MASTER_AGGREGATION',
                masterSigner,
                nonce,
                (await this.provider.getNetwork()).chainId,
                await this.aggregator.getAddress(),
                opHashes
            ]
        ));
    }

    /**
     * Get wallet nonce
     * @param {string} walletAddress - Wallet address
     * @returns {Promise<number>} Current nonce
     */
    async _getWalletNonce(walletAddress) {
        // This would typically call the wallet contract or EntryPoint
        // For demonstration, return a mock nonce
        return 0;
    }

    /**
     * Encode batch calls
     * @param {Array} calls - Array of call objects
     * @returns {Promise<string>} Encoded call data
     */
    async _encodeBatchCalls(calls) {
        // Encode calls for executeBatch function
        const iface = new ethers.Interface([
            'function executeBatch((address target, uint256 value, bytes data)[] calls)'
        ]);
        
        return iface.encodeFunctionData('executeBatch', [calls]);
    }

    /**
     * Estimate gas for aggregated operations
     * @param {number} operationCount - Number of operations
     * @returns {Promise<Object>} Gas estimates
     */
    async _estimateAggregatedGas(operationCount) {
        const baseGas = 100000; // Base EntryPoint overhead
        const perOpGas = 50000;  // Gas per operation
        const aggregationGas = 20000; // Aggregation overhead
        
        return {
            estimated: baseGas + (operationCount * perOpGas) + aggregationGas,
            breakdown: {
                base: baseGas,
                operations: operationCount * perOpGas,
                aggregation: aggregationGas
            }
        };
    }

    /**
     * Calculate actual gas savings
     * @param {number} operationCount - Number of operations
     * @param {bigint} actualGasUsed - Actual gas used
     * @returns {Promise<Object>} Savings calculation
     */
    async _calculateGasSavings(operationCount, actualGasUsed) {
        const individualGasCost = operationCount * 150000; // Estimated individual cost
        const savings = Math.max(0, individualGasCost - Number(actualGasUsed));
        
        return {
            individualEstimate: individualGasCost,
            aggregatedActual: Number(actualGasUsed),
            saved: savings,
            percentage: Math.round((savings / individualGasCost) * 100)
        };
    }
}

module.exports = {
    CPOPSignatureAggregator
};