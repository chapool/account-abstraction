/**
 * @fileoverview Example usage of CPOP Gas Calculator
 * @description Demonstrates how to use the gas calculator for off-chain cost estimation
 */

const { ethers } = require('ethers');
const { CPOPGasCalculator } = require('./gas-calculator');

// Example Oracle ABI (simplified)
const ORACLE_ABI = [
    "function getETHPriceUSD() external view returns (uint256 price, uint256 timestamp)",
    "function getCPOPPriceUSD() external view returns (uint256 price, uint256 timestamp)",
    "function estimateGasCostInCPOP(uint256 gasLimit, uint256 gasPrice) external view returns (uint256 cpopCost)",
    "function isPriceValid(string calldata feedType) external view returns (bool isValid)",
    "function getMaxPriceAge() external view returns (uint256 maxAge)",
    "function getPriceDeviationThreshold() external view returns (uint256 threshold)"
];

// Example CPOP Token ABI
const CPOP_TOKEN_ABI = [
    "function balanceOf(address account) external view returns (uint256)",
    "function transfer(address to, uint256 amount) external returns (bool)",
    "function burn(address from, uint256 amount) external"
];

async function main() {
    // Setup provider (replace with your RPC URL)
    const provider = new ethers.JsonRpcProvider(process.env.RPC_URL || 'http://localhost:8545');
    
    // Oracle and token contract addresses (replace with actual deployed addresses)
    const ORACLE_ADDRESS = process.env.ORACLE_ADDRESS || '0x...';
    const CPOP_TOKEN_ADDRESS = process.env.CPOP_TOKEN_ADDRESS || '0x...';
    
    // Initialize calculator
    const gasCalculator = new CPOPGasCalculator(provider, ORACLE_ADDRESS, ORACLE_ABI);
    const cpopToken = new ethers.Contract(CPOP_TOKEN_ADDRESS, CPOP_TOKEN_ABI, provider);
    
    try {
        console.log('🔮 CPOP Gas Calculator Example\n');
        
        // 1. Check oracle health
        console.log('1. Checking Oracle Health...');
        const healthStatus = await gasCalculator.getOracleHealthStatus();
        console.log('Oracle Health:', healthStatus.isHealthy ? '✅ Healthy' : '❌ Unhealthy');
        console.log('Feed Status:', healthStatus.feeds);
        console.log('Config:', healthStatus.config);
        console.log();
        
        // 2. Get current prices
        console.log('2. Current Prices:');
        const ethPrice = await gasCalculator.getETHPriceUSD();
        const cpopPrice = await gasCalculator.getCPOPPriceUSD();
        console.log(`ETH Price: $${ethers.formatUnits(ethPrice.price, 8)}`);
        console.log(`CPOP Price: $${ethers.formatUnits(cpopPrice.price, 8)}`);
        console.log();
        
        // 3. Example user operation
        console.log('3. Estimating User Operation Cost...');
        const exampleUserOp = {
            callGasLimit: '100000',
            verificationGasLimit: '150000',
            preVerificationGas: '21000',
            sender: '0x742d35Cc6436Cc296Cb4c5F170DCa1c8B3ce1234' // Example address
        };
        
        const costEstimate = await gasCalculator.estimateUserOpCost(exampleUserOp);
        console.log('Gas Limits:');
        console.log(`  Call: ${costEstimate.gasLimit.call.toLocaleString()}`);
        console.log(`  Verification: ${costEstimate.gasLimit.verification.toLocaleString()}`);
        console.log(`  Pre-verification: ${costEstimate.gasLimit.preVerification.toLocaleString()}`);
        console.log(`  Total: ${costEstimate.gasLimit.total.toLocaleString()}`);
        console.log();
        console.log('Costs:');
        console.log(`  ETH: ${costEstimate.costs.eth} ETH`);
        console.log(`  CPOP: ${costEstimate.costs.cpop} CPOP`);
        console.log(`  Gas Price: ${ethers.formatUnits(costEstimate.gasPrice, 'gwei')} Gwei`);
        console.log();
        
        // 4. Check user balance
        console.log('4. User Balance Check...');
        try {
            const balanceCheck = await gasCalculator.checkUserBalance(
                exampleUserOp.sender,
                exampleUserOp,
                cpopToken
            );
            
            console.log(`User CPOP Balance: ${balanceCheck.userBalance} CPOP`);
            console.log(`Required CPOP: ${balanceCheck.requiredCPOP} CPOP`);
            console.log(`Sufficient Balance: ${balanceCheck.hasSufficientBalance ? '✅ Yes' : '❌ No'}`);
            
            if (!balanceCheck.hasSufficientBalance) {
                console.log(`Shortfall: ${ethers.formatEther(balanceCheck.shortfall)} CPOP`);
            }
        } catch (error) {
            console.log('⚠️  Could not check user balance (likely example address)');
        }
        console.log();
        
        // 5. Gas optimization scenarios
        console.log('5. Gas Optimization Scenarios...');
        const exampleTx = {
            to: '0x742d35Cc6436Cc296Cb4c5F170DCa1c8B3ce1234',
            value: ethers.parseEther('0.1'),
            data: '0x'
        };
        
        const optimizations = await gasCalculator.optimizeGasParameters(exampleTx);
        console.log('Current Network Gas Price:', optimizations.current.gasPriceGwei, 'Gwei');
        console.log();
        console.log('Optimization Scenarios:');
        
        optimizations.optimizations.forEach(scenario => {
            console.log(`  ${scenario.speed.toUpperCase()}:`);
            console.log(`    Gas Price: ${scenario.gasPriceGwei} Gwei`);
            console.log(`    Total Cost: ${scenario.totalCostETH} ETH`);
            console.log(`    CPOP Cost: ${scenario.cpopCost} CPOP`);
            console.log();
        });
        
        console.log('✅ Recommended:', optimizations.recommended.speed);
        console.log();
        
        // 6. Real-time monitoring example
        console.log('6. Real-time Price Monitoring (5 updates)...');
        for (let i = 0; i < 5; i++) {
            try {
                const currentETHPrice = await gasCalculator.getETHPriceUSD();
                const currentCPOPPrice = await gasCalculator.getCPOPPriceUSD();
                const currentGasPrice = await gasCalculator.getCurrentGasPrice();
                
                console.log(`Update ${i + 1}:`);
                console.log(`  ETH: $${ethers.formatUnits(currentETHPrice.price, 8)}`);
                console.log(`  CPOP: $${ethers.formatUnits(currentCPOPPrice.price, 8)}`);
                console.log(`  Gas: ${ethers.formatUnits(currentGasPrice, 'gwei')} Gwei`);
                console.log(`  Time: ${new Date().toLocaleTimeString()}`);
                console.log();
                
                // Wait 2 seconds between updates
                if (i < 4) {
                    await new Promise(resolve => setTimeout(resolve, 2000));
                }
            } catch (error) {
                console.log(`Update ${i + 1}: ❌ Failed -`, error.message);
            }
        }
        
        console.log('🎉 Example completed successfully!');
        
    } catch (error) {
        console.error('❌ Error running example:', error.message);
        console.error('Stack:', error.stack);
    }
}

// Helper function to run with environment setup
async function runExample() {
    console.log('🚀 Starting CPOP Gas Calculator Example...\n');
    
    // Check environment variables
    const requiredEnvVars = ['RPC_URL', 'ORACLE_ADDRESS', 'CPOP_TOKEN_ADDRESS'];
    const missingVars = requiredEnvVars.filter(varName => !process.env[varName]);
    
    if (missingVars.length > 0) {
        console.log('⚠️  Missing environment variables:');
        missingVars.forEach(varName => {
            console.log(`   - ${varName}`);
        });
        console.log('\nUsing default/example values for demonstration...\n');
    }
    
    await main();
}

// Export for use in other scripts
module.exports = {
    main,
    runExample,
    ORACLE_ABI,
    CPOP_TOKEN_ABI
};

// Run if this file is executed directly
if (require.main === module) {
    runExample().catch(console.error);
}