/**
 * @fileoverview CPOP Signature Aggregation Usage Example
 * @description Demonstrates how to use signature aggregation for multiple CPOP wallets
 */

const { ethers } = require('ethers');
const { CPOPSignatureAggregator } = require('./cpop-signature-aggregator');

// Example ABIs
const AGGREGATOR_ABI = [
    "function validateSignatures(tuple(address sender, uint256 nonce, bytes initCode, bytes callData, uint256 callGasLimit, uint256 verificationGasLimit, uint256 preVerificationGas, uint256 maxFeePerGas, uint256 maxPriorityFeePerGas, bytes paymasterAndData, bytes signature)[] userOps, bytes signature) external",
    "function getMasterNonce(address master) external view returns (uint256)",
    "function isWalletControlledByMaster(address wallet, address master) external view returns (bool)",
    "function calculateGasSavings(uint256 operationCount) external view returns (uint256)",
    "function autoAuthorizeWallet(address wallet, address master) external",
    "event AggregatedValidation(address indexed master, uint256 operationCount, bytes32 aggregatedHash)",
    "event MasterAuthorized(address indexed master, bool authorized)",
    "event WalletAuthorized(address indexed master, address indexed wallet, bool authorized)"
];

const ENTRYPOINT_ABI = [
    "function handleAggregatedOps(tuple(tuple(address sender, uint256 nonce, bytes initCode, bytes callData, uint256 callGasLimit, uint256 verificationGasLimit, uint256 preVerificationGas, uint256 maxFeePerGas, uint256 maxPriorityFeePerGas, bytes paymasterAndData, bytes signature)[] userOps, address aggregator, bytes signature)[] opsPerAggregator, address payable beneficiary) external"
];

async function demonstrateAggregation() {
    console.log('🎯 CPOP Signature Aggregation Demonstration\n');
    
    // Setup (replace with actual values)
    const provider = new ethers.JsonRpcProvider(process.env.RPC_URL || 'http://localhost:8545');
    const AGGREGATOR_ADDRESS = process.env.AGGREGATOR_ADDRESS || '0x...';
    const ENTRYPOINT_ADDRESS = process.env.ENTRYPOINT_ADDRESS || '0x...';
    const MASTER_PRIVATE_KEY = process.env.MASTER_PRIVATE_KEY || '0x...';
    
    // Initialize aggregator
    const aggregator = new CPOPSignatureAggregator(
        provider,
        AGGREGATOR_ADDRESS,
        AGGREGATOR_ABI,
        MASTER_PRIVATE_KEY
    );
    
    const entryPoint = new ethers.Contract(ENTRYPOINT_ADDRESS, ENTRYPOINT_ABI, provider);
    
    console.log('📋 Configuration:');
    console.log(`   Aggregator: ${AGGREGATOR_ADDRESS}`);
    console.log(`   EntryPoint: ${ENTRYPOINT_ADDRESS}`);
    console.log(`   Master: ${aggregator.masterWallet.address}`);
    console.log();

    // Example wallet addresses (controlled by master signer)
    const walletAddresses = [
        '0x742d35Cc6436Cc296Cb4c5F170DCa1c8B3ce1001',
        '0x742d35Cc6436Cc296Cb4c5F170DCa1c8B3ce1002',
        '0x742d35Cc6436Cc296Cb4c5F170DCa1c8B3ce1003',
        '0x742d35Cc6436Cc296Cb4c5F170DCa1c8B3ce1004'
    ];

    try {
        // 1. Auto-authorize wallets
        console.log('🔐 Step 1: Auto-authorizing wallets...\n');
        
        const authResult = await aggregator.autoAuthorizeWallets(walletAddresses);
        console.log(`✅ Authorization completed: ${authResult.authorizedCount}/${authResult.totalWallets}`);
        console.log();

        // 2. Prepare batch operations for multiple wallets
        console.log('📦 Step 2: Preparing batch operations...\n');
        
        const operations = [
            {
                wallet: walletAddresses[0],
                calls: [
                    {
                        target: '0xA0b86a33E6441B8C60612240b4A54f2C1B11608c', // Example token
                        value: '0',
                        data: '0xa9059cbb000000000000000000000000742d35cc6436cc296cb4c5f170dca1c8b3ce1234000000000000000000000000000000000000000000000000016345785d8a0000' // transfer(address,uint256)
                    }
                ]
            },
            {
                wallet: walletAddresses[1],
                calls: [
                    {
                        target: '0xA0b86a33E6441B8C60612240b4A54f2C1B11608c',
                        value: '0',
                        data: '0xa9059cbb000000000000000000000000742d35cc6436cc296cb4c5f170dca1c8b3ce1234000000000000000000000000000000000000000000000000016345785d8a0000'
                    },
                    {
                        target: '0xB0b86a33E6441B8C60612240b4A54f2C1B11608c', // Another contract
                        value: '0',
                        data: '0x095ea7b3000000000000000000000000742d35cc6436cc296cb4c5f170dca1c8b3ce1234ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff' // approve(address,uint256)
                    }
                ]
            },
            {
                wallet: walletAddresses[2],
                calls: [
                    {
                        target: '0xC0b86a33E6441B8C60612240b4A54f2C1B11608c',
                        value: ethers.parseEther('0.1').toString(),
                        data: '0x' // ETH transfer
                    }
                ]
            },
            {
                wallet: walletAddresses[3],
                calls: [
                    {
                        target: '0xA0b86a33E6441B8C60612240b4A54f2C1B11608c',
                        value: '0',
                        data: '0xa9059cbb000000000000000000000000742d35cc6436cc296cb4c5f170dca1c8b3ce1234000000000000000000000000000000000000000000000000016345785d8a0000'
                    }
                ]
            }
        ];
        
        const userOps = await aggregator.batchOperationsFromWallets(operations, {
            maxFeePerGas: ethers.parseUnits('25', 'gwei').toString(),
            maxPriorityFeePerGas: ethers.parseUnits('2', 'gwei').toString()
        });
        
        console.log(`✅ Created ${userOps.length} user operations`);
        console.log();

        // 3. Calculate potential gas savings
        console.log('💰 Step 3: Calculating gas savings...\n');
        
        const gasSavings = await aggregator.calculateGasSavings(userOps.length);
        console.log('Gas Savings Analysis:');
        console.log(`   Operations: ${gasSavings.operationCount}`);
        console.log(`   Gas Saved: ${gasSavings.gasSavedFormatted}`);
        console.log(`   ETH Saved: ${gasSavings.estimatedETHSaved} ETH`);
        console.log(`   Savings: ${gasSavings.savingsPercentage}%`);
        console.log(`   Gas Price: ${gasSavings.currentGasPrice}`);
        console.log();

        // 4. Create aggregated signature
        console.log('🔗 Step 4: Creating aggregated signature...\n');
        
        const aggregatedData = await aggregator.createAggregatedSignature(userOps, {
            validateWallets: true
        });
        
        console.log('Aggregated Signature Created:');
        console.log(`   Master: ${aggregatedData.masterSigner}`);
        console.log(`   Nonce: ${aggregatedData.nonce}`);
        console.log(`   Hash: ${aggregatedData.hash}`);
        console.log(`   Signature: ${aggregatedData.signature.slice(0, 42)}...`);
        console.log();

        // 5. Execute aggregated operations
        console.log('🚀 Step 5: Executing aggregated operations...\n');
        
        if (process.env.EXECUTE_TX === 'true') {
            const beneficiary = aggregator.masterWallet.address;
            const executeResult = await aggregator.executeAggregatedOperations(
                aggregatedData,
                entryPoint,
                beneficiary
            );
            
            console.log('Execution Results:');
            console.log(`   Success: ${executeResult.success}`);
            console.log(`   Tx Hash: ${executeResult.transactionHash}`);
            console.log(`   Gas Used: ${executeResult.gasUsed.toLocaleString()}`);
            console.log(`   Operations: ${executeResult.operationsCount}`);
            console.log(`   Actual Savings: ${executeResult.gasSavings.percentage}%`);
            console.log();
        } else {
            console.log('ℹ️  Execution skipped (set EXECUTE_TX=true to execute)');
            console.log(`   Would execute ${aggregatedData.userOps.length} operations`);
            console.log(`   Estimated gas: ${aggregatedData.gasEstimate.estimated.toLocaleString()}`);
            console.log();
        }

        // 6. Monitor events
        console.log('📡 Step 6: Event monitoring example...\n');
        
        aggregator.startEventMonitoring((event) => {
            console.log(`🔔 Event: ${event.type}`);
            console.log(`   Block: ${event.blockNumber}`);
            console.log(`   Tx: ${event.transactionHash}`);
            
            if (event.type === 'AggregatedValidation') {
                console.log(`   Master: ${event.master}`);
                console.log(`   Operations: ${event.operationCount}`);
            } else if (event.type === 'WalletAuthorized') {
                console.log(`   Master: ${event.master}`);
                console.log(`   Wallet: ${event.wallet}`);
                console.log(`   Authorized: ${event.authorized}`);
            }
            console.log();
        });
        
        console.log('Event monitoring started (listening for aggregation events)...');
        
        // 7. Advanced usage scenarios
        console.log('🎛️  Step 7: Advanced scenarios...\n');
        
        console.log('Scenario A: High-frequency trading operations');
        const tradingOps = await simulateTrading(aggregator, walletAddresses.slice(0, 2));
        console.log(`   Created ${tradingOps.length} trading operations`);
        
        console.log('Scenario B: Token distribution');
        const distributionOps = await simulateTokenDistribution(aggregator, walletAddresses);
        console.log(`   Created ${distributionOps.length} distribution operations`);
        
        console.log('Scenario C: Cross-wallet DeFi interactions');
        const defiOps = await simulateDeFiOperations(aggregator, walletAddresses.slice(0, 3));
        console.log(`   Created ${defiOps.length} DeFi operations`);
        
        console.log();
        console.log('🎉 Aggregation demonstration completed successfully!');
        
    } catch (error) {
        console.error('❌ Demonstration failed:', error.message);
        if (error.stack) {
            console.error('Stack:', error.stack);
        }
    }
}

// Simulation helpers
async function simulateTrading(aggregator, walletAddresses) {
    const operations = walletAddresses.map((wallet, index) => ({
        wallet,
        calls: [
            {
                target: '0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57', // Example DEX
                value: '0',
                data: `0x38ed1739${index.toString().padStart(64, '0')}` // Mock swap data
            }
        ]
    }));
    
    return await aggregator.batchOperationsFromWallets(operations);
}

async function simulateTokenDistribution(aggregator, walletAddresses) {
    const operations = walletAddresses.map((wallet, index) => ({
        wallet,
        calls: [
            {
                target: '0xA0b86a33E6441B8C60612240b4A54f2C1B11608c',
                value: '0',
                data: '0xa9059cbb000000000000000000000000' + wallet.slice(2) + '000000000000000000000000000000000000000000000000016345785d8a0000'
            }
        ]
    }));
    
    return await aggregator.batchOperationsFromWallets(operations);
}

async function simulateDeFiOperations(aggregator, walletAddresses) {
    const operations = walletAddresses.map((wallet, index) => {
        const calls = [];
        
        // Approve token
        calls.push({
            target: '0xA0b86a33E6441B8C60612240b4A54f2C1B11608c',
            value: '0',
            data: '0x095ea7b3000000000000000000000000DEF171Fe48CF0115B1d80b88dc8eAB59176FEe57ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
        });
        
        // Provide liquidity or stake
        calls.push({
            target: '0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57',
            value: '0',
            data: `0xe8e33700${index.toString().padStart(64, '0')}`
        });
        
        return { wallet, calls };
    });
    
    return await aggregator.batchOperationsFromWallets(operations);
}

// Production helper function
async function createAndExecuteAggregation(walletOperations, config = {}) {
    const provider = new ethers.JsonRpcProvider(config.rpcUrl || process.env.RPC_URL);
    const aggregator = new CPOPSignatureAggregator(
        provider,
        config.aggregatorAddress || process.env.AGGREGATOR_ADDRESS,
        AGGREGATOR_ABI,
        config.masterPrivateKey || process.env.MASTER_PRIVATE_KEY
    );
    
    // Create user operations
    const userOps = await aggregator.batchOperationsFromWallets(walletOperations);
    
    // Create aggregated signature
    const aggregatedData = await aggregator.createAggregatedSignature(userOps);
    
    // Execute if requested
    if (config.execute) {
        const entryPoint = new ethers.Contract(
            config.entryPointAddress || process.env.ENTRYPOINT_ADDRESS,
            ENTRYPOINT_ABI,
            provider
        );
        
        return await aggregator.executeAggregatedOperations(
            aggregatedData,
            entryPoint,
            config.beneficiary || aggregator.masterWallet.address
        );
    }
    
    return aggregatedData;
}

// Export for use in other modules
module.exports = {
    demonstrateAggregation,
    createAndExecuteAggregation,
    simulateTrading,
    simulateTokenDistribution,
    simulateDeFiOperations,
    AGGREGATOR_ABI,
    ENTRYPOINT_ABI
};

// Run demonstration if this file is executed directly
if (require.main === module) {
    demonstrateAggregation().catch(console.error);
}