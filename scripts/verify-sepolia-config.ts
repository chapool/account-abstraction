import { ethers } from "hardhat";
import * as fs from "fs";

// Load Sepolia environment variables
if (fs.existsSync('.env.sepolia')) {
    require('dotenv').config({ path: '.env.sepolia' });
}

async function verifySepoliaConfig() {
    console.log("🔍 Verifying Sepolia configuration...\n");
    
    try {
        // 1. Check environment variables
        const rpcUrl = process.env.ETH_RPC_URL;
        const privateKey = process.env.PRIVATE_KEY;
        const chainId = process.env.CHAIN_ID;
        
        console.log("📋 Configuration:");
        console.log(`RPC URL: ${rpcUrl}`);
        console.log(`Chain ID: ${chainId}`);
        console.log(`Private Key: ${privateKey ? `${privateKey.substring(0, 8)}...` : 'NOT SET'}\n`);
        
        // 2. Connect to provider
        const provider = new ethers.providers.JsonRpcProvider(rpcUrl);
        
        // 3. Verify network
        const network = await provider.getNetwork();
        console.log("🌐 Network Info:");
        console.log(`Name: ${network.name}`);
        console.log(`Chain ID: ${network.chainId}`);
        console.log(`Is Sepolia: ${network.chainId === 11155111 ? '✅ Yes' : '❌ No'}\n`);
        
        // 4. Check wallet
        if (!privateKey) {
            throw new Error("Private key not found in environment");
        }
        
        const wallet = new ethers.Wallet(privateKey, provider);
        console.log("👛 Wallet Info:");
        console.log(`Address: ${wallet.address}`);
        
        // 5. Check balance
        const balance = await provider.getBalance(wallet.address);
        const balanceEth = ethers.utils.formatEther(balance);
        console.log(`Balance: ${balanceEth} ETH`);
        
        if (parseFloat(balanceEth) === 0) {
            console.log("⚠️  Warning: Wallet has no ETH. Get test ETH from:");
            console.log("   - https://sepoliafaucet.com/");
            console.log("   - https://faucet.sepolia.dev/");
        } else if (parseFloat(balanceEth) < 0.01) {
            console.log("⚠️  Warning: Low balance. Consider getting more test ETH");
        } else {
            console.log("✅ Sufficient balance for testing");
        }
        
        // 6. Check gas price
        const gasPrice = await provider.getGasPrice();
        console.log(`\n⛽ Gas Info:`);
        console.log(`Gas Price: ${ethers.utils.formatUnits(gasPrice, "gwei")} Gwei`);
        
        // 7. Test transaction simulation (dry run)
        console.log("\n🧪 Testing transaction simulation...");
        try {
            const tx = {
                to: wallet.address,
                value: ethers.utils.parseEther("0.001"),
                gasLimit: 21000
            };
            
            const estimatedGas = await provider.estimateGas(tx);
            console.log(`✅ Transaction simulation successful`);
            console.log(`   Estimated Gas: ${estimatedGas}`);
        } catch (error: any) {
            console.log(`❌ Transaction simulation failed: ${error.message}`);
        }
        
        console.log("\n🎉 Sepolia configuration verification complete!");
        
    } catch (error: any) {
        console.error("❌ Configuration verification failed:", error.message);
        process.exit(1);
    }
}

// Run verification
verifySepoliaConfig()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("Script failed:", error);
        process.exit(1);
    });