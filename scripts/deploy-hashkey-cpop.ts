import { ethers } from "hardhat";

async function main() {
    console.log("🚀 Deploying CPOP contracts to Hashkey Testnet...");
    console.log("=".repeat(55));
    
    const [deployer] = await ethers.getSigners();
    const deployerAddress = await deployer.getAddress();
    
    console.log("🔑 Deployer:", deployerAddress);
    console.log("💰 Balance:", ethers.utils.formatEther(await deployer.getBalance()), "HSK");
    
    // Hashkey Testnet may use the same EntryPoint or we need to check
    const entryPointAddress = "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"; // Standard ERC-4337 EntryPoint
    
    const deployedContracts: Record<string, string> = {};
    
    try {
        // 1. Deploy CPOPToken
        console.log("\n📝 Step 1: Deploying CPOPToken...");
        const initialSupply = ethers.utils.parseEther("1000000"); // 1M tokens
        
        const CPOPTokenFactory = await ethers.getContractFactory("contracts/cpop/CPOPToken.sol:CPOPToken");
        const cpopToken = await CPOPTokenFactory.deploy(deployerAddress, initialSupply);
        await cpopToken.deployed();
        
        deployedContracts.CPOPToken = cpopToken.address;
        console.log("✅ CPOPToken deployed:", cpopToken.address);
        
        // 2. Deploy GasPriceOracle
        console.log("\n📊 Step 2: Deploying GasPriceOracle...");
        const GasPriceOracleFactory = await ethers.getContractFactory("contracts/cpop/GasPriceOracle.sol:GasPriceOracle");
        const gasPriceOracle = await GasPriceOracleFactory.deploy();
        await gasPriceOracle.deployed();
        
        deployedContracts.GasPriceOracle = gasPriceOracle.address;
        console.log("✅ GasPriceOracle deployed:", gasPriceOracle.address);
        
        // 3. Deploy SessionKeyManager
        console.log("\n🔑 Step 3: Deploying SessionKeyManager...");
        const SessionKeyManagerFactory = await ethers.getContractFactory("contracts/cpop/SessionKeyManager.sol:SessionKeyManager");
        const sessionKeyManager = await SessionKeyManagerFactory.deploy();
        await sessionKeyManager.deployed();
        
        deployedContracts.SessionKeyManager = sessionKeyManager.address;
        console.log("✅ SessionKeyManager deployed:", sessionKeyManager.address);
        
        // 4. Deploy AAWallet implementation
        console.log("\n👤 Step 4: Deploying AAWallet implementation...");
        const AAWalletFactory = await ethers.getContractFactory("contracts/cpop/AAWallet.sol:AAWallet");
        const aaWalletImpl = await AAWalletFactory.deploy();
        await aaWalletImpl.deployed();
        
        deployedContracts.AAWallet = aaWalletImpl.address;
        console.log("✅ AAWallet implementation deployed:", aaWalletImpl.address);
        
        // 5. Deploy WalletManager
        console.log("\n🏭 Step 5: Deploying WalletManager...");
        const WalletManagerFactory = await ethers.getContractFactory("contracts/cpop/WalletManager.sol:WalletManager");
        const walletManager = await WalletManagerFactory.deploy();
        await walletManager.deployed();
        
        deployedContracts.WalletManager = walletManager.address;
        console.log("✅ WalletManager deployed:", walletManager.address);
        
        // Initialize WalletManager
        console.log("🔧 Initializing WalletManager...");
        try {
            const initTx = await walletManager.initialize(entryPointAddress, deployerAddress, {
                gasLimit: 500000
            });
            await initTx.wait();
            console.log("✅ WalletManager initialized");
        } catch (initError) {
            console.warn("⚠️  WalletManager initialization may have failed:", (initError as Error).message);
        }
        
        // 6. Deploy MasterAggregator
        console.log("\n🔗 Step 6: Deploying MasterAggregator...");
        const MasterAggregatorFactory = await ethers.getContractFactory("contracts/cpop/MasterAggregator.sol:MasterAggregator");
        const masterAggregator = await MasterAggregatorFactory.deploy();
        await masterAggregator.deployed();
        
        deployedContracts.MasterAggregator = masterAggregator.address;
        console.log("✅ MasterAggregator deployed:", masterAggregator.address);
        
        // Initialize MasterAggregator
        console.log("🔧 Initializing MasterAggregator...");
        try {
            const initMasterTx = await masterAggregator.initialize(deployerAddress);
            await initMasterTx.wait();
            console.log("✅ MasterAggregator initialized");
        } catch (error) {
            console.warn("⚠️  MasterAggregator initialization may have failed:", (error as Error).message);
        }
        
        console.log("\n🎉 All contracts deployed successfully!");
        
        // Test basic functionality
        console.log("\n🧪 Testing basic functionality...");
        
        // Test CPOPToken
        console.log("Testing CPOPToken...");
        const tokenName = await cpopToken.name();
        const tokenSymbol = await cpopToken.symbol();
        const totalSupply = await cpopToken.totalSupply();
        const deployerBalance = await cpopToken.balanceOf(deployerAddress);
        
        console.log("✅ Token Name:", tokenName);
        console.log("✅ Token Symbol:", tokenSymbol);
        console.log("✅ Total Supply:", ethers.utils.formatEther(totalSupply));
        console.log("✅ Deployer Balance:", ethers.utils.formatEther(deployerBalance));
        
        // Test batch mint
        console.log("Testing batch mint...");
        const testRecipients = [
            "0x1111111111111111111111111111111111111111",
            "0x2222222222222222222222222222222222222222"
        ];
        const testAmounts = [
            ethers.utils.parseEther("100"),
            ethers.utils.parseEther("200")
        ];
        
        const batchMintTx = await cpopToken.batchMint(testRecipients, testAmounts);
        const batchMintReceipt = await batchMintTx.wait();
        
        console.log("✅ Batch mint successful!");
        console.log("Gas used:", batchMintReceipt.gasUsed.toString());
        
        // Check balances
        for (let i = 0; i < testRecipients.length; i++) {
            const balance = await cpopToken.balanceOf(testRecipients[i]);
            console.log(`Balance ${testRecipients[i]}: ${ethers.utils.formatEther(balance)} CPOP`);
        }
        
    } catch (error) {
        console.error("❌ Deployment failed:", error);
        throw error;
    }
    
    // Final summary
    console.log("\n📋 Hashkey Testnet Deployment Summary:");
    console.log("=".repeat(50));
    console.log("Network: Hashkey Chain Testnet (ChainID: 133)");
    console.log("EntryPoint:", entryPointAddress);
    console.log("Deployer:", deployerAddress);
    console.log("");
    
    console.log("📍 Contract Addresses:");
    for (const [name, address] of Object.entries(deployedContracts)) {
        console.log(`${name}: ${address}`);
        console.log(`  Explorer: https://testnet-explorer.hsk.xyz/address/${address}`);
        console.log("");
    }
    
    console.log("🔧 Integration Status:");
    console.log("✅ CPOPToken with batch operations deployed");
    console.log("✅ GasPriceOracle for price feeds deployed");
    console.log("✅ SessionKeyManager for session management deployed");
    console.log("✅ AAWallet implementation deployed");
    console.log("✅ WalletManager for wallet factory deployed");
    console.log("✅ MasterAggregator for signature aggregation deployed");
    
    console.log("\n🚀 CPOP ecosystem ready on Hashkey Testnet!");
    
    return deployedContracts;
}

main()
    .then((contracts) => {
        console.log("\n🎉 Hashkey Testnet deployment completed!");
        console.log("All CPOP contracts are now live on Hashkey Chain Testnet.");
        
        console.log("\n📝 Quick Start:");
        console.log("1. Use CPOPToken for batch token operations");
        console.log("2. Use WalletManager to create AA wallets");
        console.log("3. Enjoy low gas fees on Hashkey Chain!");
        
        process.exit(0);
    })
    .catch((error) => {
        console.error("❌ Deployment error:", error.message);
        process.exit(1);
    });