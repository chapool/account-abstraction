import { ethers } from "hardhat";

async function main() {
    console.log("🚀 Deploying core CPOP contracts to BSC Testnet...");
    
    // Get the deployer account
    const [deployer] = await ethers.getSigners();
    const deployerAddress = await deployer.getAddress();
    
    console.log("Deploying with account:", deployerAddress);
    console.log("Account balance:", ethers.utils.formatEther(await deployer.getBalance()), "BNB");
    
    // BSC Testnet EntryPoint address
    const entryPointAddress = "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789";
    console.log("Using EntryPoint:", entryPointAddress);
    
    const deployedContracts: Record<string, string> = {};
    
    try {
        // 1. CPOPToken (already deployed)
        console.log("\n📝 CPOPToken already deployed");
        deployedContracts.CPOPToken = "0xc2DEe82B70C60eDbF8a8E180Cb72A2e727574260";
        
        // 2. Deploy WalletManager
        console.log("\n👤 Deploying WalletManager...");
        const WalletManagerFactory = await ethers.getContractFactory("contracts/cpop/WalletManager.sol:WalletManager");
        const walletManager = await WalletManagerFactory.deploy();
        await walletManager.deployed();
        deployedContracts.WalletManager = walletManager.address;
        console.log("✅ WalletManager deployed to:", walletManager.address);
        
        // Initialize WalletManager
        console.log("Initializing WalletManager...");
        const initTx = await walletManager.initialize(entryPointAddress, deployerAddress);
        await initTx.wait();
        console.log("✅ WalletManager initialized");
        
        // 3. Deploy MasterAggregator
        console.log("\n🔗 Deploying MasterAggregator...");
        const MasterAggregatorFactory = await ethers.getContractFactory("contracts/cpop/MasterAggregator.sol:MasterAggregator");
        const masterAggregator = await MasterAggregatorFactory.deploy(entryPointAddress);
        await masterAggregator.deployed();
        deployedContracts.MasterAggregator = masterAggregator.address;
        console.log("✅ MasterAggregator deployed to:", masterAggregator.address);
        
        // 4. Deploy SessionKeyManager
        console.log("\n🔑 Deploying SessionKeyManager...");
        const SessionKeyManagerFactory = await ethers.getContractFactory("contracts/cpop/SessionKeyManager.sol:SessionKeyManager");
        const sessionKeyManager = await SessionKeyManagerFactory.deploy();
        await sessionKeyManager.deployed();
        deployedContracts.SessionKeyManager = sessionKeyManager.address;
        console.log("✅ SessionKeyManager deployed to:", sessionKeyManager.address);
        
        // Setup integrations
        console.log("\n🔧 Setting up integrations...");
        
        // Set MasterAggregator in WalletManager
        const setAggregatorTx = await walletManager.setMasterAggregator(masterAggregator.address);
        await setAggregatorTx.wait();
        console.log("✅ MasterAggregator set in WalletManager");
        
        // Test wallet creation
        console.log("\n🧪 Testing wallet creation...");
        const testOwner = deployerAddress;
        const testMasterSigner = deployerAddress;
        
        const walletAddress = await walletManager.getAccountAddress(testOwner, testMasterSigner);
        console.log("Predicted wallet address:", walletAddress);
        
        const createTx = await walletManager.createWallet(testOwner, testMasterSigner);
        await createTx.wait();
        console.log("✅ Test wallet created successfully");
        
        console.log("\n🎉 Core CPOP contracts deployed successfully!");
        
    } catch (error) {
        console.error("❌ Deployment failed:", error);
        throw error;
    }
    
    // Print summary
    console.log("\n📋 Deployment Summary:");
    console.log("=".repeat(50));
    console.log("Network: BSC Testnet (ChainID: 97)");
    console.log("Deployer:", deployerAddress);
    console.log("EntryPoint:", entryPointAddress);
    console.log("\n📍 Contract Addresses:");
    
    for (const [name, address] of Object.entries(deployedContracts)) {
        console.log(`${name}: ${address}`);
        console.log(`  BSCScan: https://testnet.bscscan.com/address/${address}`);
    }
    
    console.log("\n🔗 Integration Status:");
    console.log("✅ WalletManager initialized with EntryPoint");
    console.log("✅ MasterAggregator set in WalletManager");
    console.log("✅ Test wallet created successfully");
    console.log("✅ Core AA infrastructure ready");
    
    return deployedContracts;
}

main()
    .then((contracts) => {
        console.log("\n🚀 Core CPOP deployment completed!");
        console.log("Account abstraction infrastructure is ready on BSC Testnet.");
        process.exit(0);
    })
    .catch((error) => {
        console.error("❌ Deployment failed:", error);
        process.exit(1);
    });