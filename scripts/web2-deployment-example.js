/**
 * @fileoverview Web2 Hybrid Deployment Example
 * @description Demonstrates how to use the hybrid deployment mode for Web2 users
 */

const { ethers } = require('ethers');
const { Web2WalletDeployer } = require('./web2-wallet-deployer');

// Example ABI for CPOPWalletManager
const WALLET_MANAGER_ABI = [
    "function getWeb2AccountAddress(string calldata identifier, address masterSigner) external view returns (address)",
    "function getWeb2InitCode(string calldata identifier, address masterSigner) external view returns (bytes memory)",
    "function isWeb2AccountDeployed(string calldata identifier, address masterSigner) external view returns (bool)",
    "function createWeb2AccountWithMasterSigner(string calldata identifier, address masterSigner) external returns (address, address)",
    "function createAccountWithMasterSigner(address generatedOwner, bytes32 salt, address masterSigner) external returns (address)",
    "event AccountCreated(address indexed account, address indexed owner, bytes32 indexed salt)"
];

async function demonstrateHybridDeployment() {
    console.log('🌟 Web2 Hybrid Deployment Demonstration\n');
    
    // Setup (replace with your actual values)
    const provider = new ethers.JsonRpcProvider(process.env.RPC_URL || 'http://localhost:8545');
    const WALLET_MANAGER_ADDRESS = process.env.WALLET_MANAGER_ADDRESS || '0x...';
    const MASTER_SIGNER_ADDRESS = process.env.MASTER_SIGNER_ADDRESS || '0x...';
    
    // Initialize deployer
    const deployer = new Web2WalletDeployer(
        provider,
        WALLET_MANAGER_ADDRESS,
        WALLET_MANAGER_ABI,
        {
            preferLazyDeployment: true,
            gasThreshold: ethers.parseUnits('30', 'gwei'), // 30 Gwei threshold
            highValueThreshold: ethers.parseEther('0.5')    // 0.5 ETH
        }
    );

    console.log('📋 Deployer Configuration:');
    console.log(deployer.getConfig());
    console.log();

    // Example user scenarios
    const users = [
        {
            identifier: 'user123',
            masterSigner: MASTER_SIGNER_ADDRESS,
            userValue: ethers.parseEther('0.1'),
            expectsMultipleOps: false,
            scenario: 'Regular user'
        },
        {
            identifier: 'vip_user456',
            masterSigner: MASTER_SIGNER_ADDRESS,
            userValue: ethers.parseEther('2.0'),
            expectsMultipleOps: true,
            scenario: 'High-value user'
        },
        {
            identifier: 'casual_user789',
            masterSigner: MASTER_SIGNER_ADDRESS,
            userValue: ethers.parseEther('0.01'),
            expectsMultipleOps: false,
            scenario: 'Casual user'
        }
    ];

    // Demonstrate strategy selection for each user
    console.log('🎯 Strategy Analysis for Different User Types:\n');
    
    for (const user of users) {
        console.log(`📋 ${user.scenario} (${user.identifier}):`);
        
        try {
            const strategy = await deployer.getDeploymentStrategy(
                user.identifier,
                user.masterSigner,
                {
                    userValue: user.userValue,
                    expectsMultipleOps: user.expectsMultipleOps
                }
            );
            
            console.log(`   Strategy: ${strategy.strategy}`);
            console.log(`   Account: ${strategy.accountAddress}`);
            console.log(`   Deployed: ${strategy.isDeployed}`);
            console.log(`   Reasoning: ${strategy.reasoning.join(', ')}`);
            
            if (strategy.estimatedCost) {
                console.log(`   Est. Cost: ${ethers.formatEther(strategy.estimatedCost)} ETH`);
            }
            
            console.log();
        } catch (error) {
            console.log(`   ❌ Error: ${error.message}\n`);
        }
    }

    // Demonstrate lazy deployment (initCode generation)
    console.log('⏳ Lazy Deployment Example:\n');
    
    try {
        const lazyResult = await deployer.prepareLazyDeployment(
            'lazy_user_001',
            MASTER_SIGNER_ADDRESS
        );
        
        console.log('✅ Lazy deployment prepared:');
        console.log(`   Account Address: ${lazyResult.accountAddress}`);
        console.log(`   InitCode Length: ${lazyResult.initCode.length / 2 - 1} bytes`);
        console.log(`   InitCode: ${lazyResult.initCode.slice(0, 42)}...`);
        console.log();
        
        // Example UserOperation with initCode
        const exampleUserOp = {
            sender: lazyResult.accountAddress,
            nonce: '0x0',
            initCode: lazyResult.initCode,
            callData: '0x', // Empty for demonstration
            callGasLimit: '100000',
            verificationGasLimit: '150000',
            preVerificationGas: '21000',
            maxFeePerGas: '20000000000', // 20 Gwei
            maxPriorityFeePerGas: '1000000000', // 1 Gwei
            paymasterAndData: '0x',
            signature: '0x' // Will be filled by wallet
        };
        
        console.log('📝 Example UserOperation for lazy deployment:');
        console.log(JSON.stringify(exampleUserOp, null, 2));
        console.log();
        
    } catch (error) {
        console.log(`❌ Lazy deployment preparation failed: ${error.message}\n`);
    }

    // Demonstrate immediate deployment (requires authorized signer)
    console.log('🚀 Immediate Deployment Example:\n');
    
    if (process.env.PRIVATE_KEY) {
        try {
            const signer = new ethers.Wallet(process.env.PRIVATE_KEY, provider);
            console.log(`Using signer: ${signer.address}`);
            
            const immediateResult = await deployer.deployImmediately(
                'immediate_user_002',
                MASTER_SIGNER_ADDRESS,
                signer
            );
            
            console.log('✅ Immediate deployment completed:');
            console.log(`   Account: ${immediateResult.accountAddress}`);
            console.log(`   Tx Hash: ${immediateResult.transactionHash}`);
            console.log(`   Gas Used: ${immediateResult.gasUsed.toLocaleString()}`);
            console.log(`   Cost: ${ethers.formatEther(immediateResult.deploymentCost)} ETH`);
            console.log();
            
        } catch (error) {
            console.log(`❌ Immediate deployment failed: ${error.message}`);
            console.log('Note: This requires an authorized creator private key\n');
        }
    } else {
        console.log('ℹ️  Skipping immediate deployment (PRIVATE_KEY not provided)\n');
    }

    // Gas price monitoring example
    console.log('⛽ Gas Price Monitoring Example:\n');
    
    try {
        const gasMonitoring = await deployer.monitorGasPrices(2); // 2 minutes
        
        console.log('📊 Gas monitoring results:');
        console.log(`   Samples: ${gasMonitoring.samplesCount}`);
        console.log(`   Min: ${gasMonitoring.analysis.min.gwei} Gwei`);
        console.log(`   Max: ${gasMonitoring.analysis.max.gwei} Gwei`);
        console.log(`   Avg: ${gasMonitoring.analysis.average.gwei} Gwei`);
        console.log(`   Recommendation: ${gasMonitoring.recommendation}`);
        console.log();
        
    } catch (error) {
        console.log(`❌ Gas monitoring failed: ${error.message}\n`);
    }

    // Smart deployment example
    console.log('🧠 Smart Deployment Example:\n');
    
    try {
        // This will automatically choose the best strategy
        const smartResult = await deployer.smartDeploy(
            'smart_user_003',
            MASTER_SIGNER_ADDRESS,
            process.env.PRIVATE_KEY ? new ethers.Wallet(process.env.PRIVATE_KEY, provider) : null,
            {
                userValue: ethers.parseEther('0.3'),
                expectsMultipleOps: false
            }
        );
        
        console.log('✅ Smart deployment result:');
        console.log(`   Strategy: ${smartResult.deploymentStrategy || smartResult.strategy}`);
        console.log(`   Account: ${smartResult.accountAddress}`);
        
        if (smartResult.transactionHash) {
            console.log(`   Tx Hash: ${smartResult.transactionHash}`);
        }
        
        if (smartResult.initCode) {
            console.log(`   InitCode: ${smartResult.initCode.slice(0, 42)}...`);
        }
        
        console.log();
        
    } catch (error) {
        console.log(`❌ Smart deployment failed: ${error.message}\n`);
    }

    console.log('🎉 Demonstration completed!');
}

// Helper function for production use
async function deployWeb2Wallet(identifier, masterSigner, options = {}) {
    const provider = new ethers.JsonRpcProvider(options.rpcUrl || process.env.RPC_URL);
    const deployer = new Web2WalletDeployer(
        provider,
        options.walletManagerAddress || process.env.WALLET_MANAGER_ADDRESS,
        WALLET_MANAGER_ABI,
        options.config
    );

    const signer = options.privateKey ? 
        new ethers.Wallet(options.privateKey, provider) : null;

    return await deployer.smartDeploy(identifier, masterSigner, signer, options);
}

// Export for use in other modules
module.exports = {
    demonstrateHybridDeployment,
    deployWeb2Wallet,
    WALLET_MANAGER_ABI
};

// Run demonstration if this file is executed directly
if (require.main === module) {
    demonstrateHybridDeployment().catch(console.error);
}