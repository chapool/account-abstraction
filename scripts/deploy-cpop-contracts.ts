import { ethers } from "hardhat";

async function main() {
    console.log("🚀 Deploying CPOP contracts to BSC Testnet...");
    
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
        // 1. Deploy CPOPToken (already deployed)
        console.log("\n📝 CPOPToken already deployed at: 0xc2DEe82B70C60eDbF8a8E180Cb72A2e727574260");
        deployedContracts.CPOPToken = "0xc2DEe82B70C60eDbF8a8E180Cb72A2e727574260";
        
        // 2. Deploy GasPriceOracle
        console.log("\n📊 Deploying GasPriceOracle...");
        const GasPriceOracleFactory = await ethers.getContractFactory("contracts/cpop/GasPriceOracle.sol:GasPriceOracle");
        const gasPriceOracle = await GasPriceOracleFactory.deploy();
        await gasPriceOracle.deployed();
        deployedContracts.GasPriceOracle = gasPriceOracle.address;
        console.log("✅ GasPriceOracle deployed to:", gasPriceOracle.address);
        
        // 3. Deploy GasPaymaster
        console.log("\n💰 Deploying GasPaymaster...");
        const GasPaymasterFactory = await ethers.getContractFactory("contracts/cpop/GasPaymaster.sol:GasPaymaster");
        const fallbackExchangeRate = ethers.utils.parseEther("1000"); // 1000 tokens per ETH
        const burnTokens = true; // Burn tokens instead of transferring
        const beneficiary = ethers.constants.AddressZero; // Not needed when burning
        
        const gasPaymaster = await GasPaymasterFactory.deploy(
            entryPointAddress,
            deployedContracts.CPOPToken,
            gasPriceOracle.address,
            fallbackExchangeRate,
            burnTokens,
            beneficiary
        );
        await gasPaymaster.deployed();
        deployedContracts.GasPaymaster = gasPaymaster.address;
        console.log("✅ GasPaymaster deployed to:", gasPaymaster.address);
        
        // 4. Deploy WalletManager
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
        
        // 5. Deploy MasterAggregator
        console.log("\n🔗 Deploying MasterAggregator...");
        const MasterAggregatorFactory = await ethers.getContractFactory("contracts/cpop/MasterAggregator.sol:MasterAggregator");
        const masterAggregator = await MasterAggregatorFactory.deploy(entryPointAddress);
        await masterAggregator.deployed();
        deployedContracts.MasterAggregator = masterAggregator.address;
        console.log("✅ MasterAggregator deployed to:", masterAggregator.address);
        
        // 6. Deploy SessionKeyManager
        console.log("\n🔑 Deploying SessionKeyManager...");
        const SessionKeyManagerFactory = await ethers.getContractFactory("contracts/cpop/SessionKeyManager.sol:SessionKeyManager");
        const sessionKeyManager = await SessionKeyManagerFactory.deploy();
        await sessionKeyManager.deployed();
        deployedContracts.SessionKeyManager = sessionKeyManager.address;
        console.log("✅ SessionKeyManager deployed to:", sessionKeyManager.address);
        
        // Wait for confirmations
        console.log("\n⏳ Waiting for block confirmations...");
        await new Promise(resolve => setTimeout(resolve, 30000)); // Wait 30 seconds
        
        // Setup integrations
        console.log("\n🔧 Setting up contract integrations...");
        
        // Set MasterAggregator in WalletManager
        console.log("Setting MasterAggregator in WalletManager...");
        const setAggregatorTx = await walletManager.setMasterAggregator(masterAggregator.address);
        await setAggregatorTx.wait();
        console.log("✅ MasterAggregator set in WalletManager");
        
        // Grant roles to contracts
        console.log("Setting up CPOPToken roles...");
        const cpopToken = await ethers.getContractAt("contracts/cpop/CPOPToken.sol:CPOPToken", deployedContracts.CPOPToken);
        
        // Grant MINTER_ROLE to WalletManager and MasterAggregator
        await cpopToken.grantRole(walletManager.address, 2); // MINTER_ROLE
        await cpopToken.grantRole(masterAggregator.address, 2); // MINTER_ROLE
        console.log("✅ MINTER_ROLE granted to WalletManager and MasterAggregator");
        
        // Grant BURNER_ROLE to GasPaymaster and MasterAggregator
        await cpopToken.grantRole(gasPaymaster.address, 4); // BURNER_ROLE
        await cpopToken.grantRole(masterAggregator.address, 4); // BURNER_ROLE
        console.log("✅ BURNER_ROLE granted to GasPaymaster and MasterAggregator");
        
        console.log("\n🎉 All CPOP contracts deployed and configured successfully!");
        
    } catch (error) {
        console.error("❌ Deployment failed:", error);
        throw error;
    }
    
    // Print summary
    console.log("\n📋 Deployment Summary:");
    console.log("=" * 50);
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
    console.log("✅ CPOPToken roles configured for all contracts");
    console.log("✅ All contracts ready for account abstraction operations");
    
    return deployedContracts;
}

main()
    .then((contracts) => {
        console.log("\n🚀 CPOP deployment completed successfully!");
        console.log("All contracts are deployed and ready to use on BSC Testnet.");
        process.exit(0);
    })
    .catch((error) => {
        console.error("❌ Deployment failed:", error);
        process.exit(1);
    });