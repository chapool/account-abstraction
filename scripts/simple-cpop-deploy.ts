import { ethers } from "hardhat";

async function main() {
    console.log("🚀 Simple CPOP deployment to BSC Testnet...");
    
    const [deployer] = await ethers.getSigners();
    const deployerAddress = await deployer.getAddress();
    
    console.log("Deployer:", deployerAddress);
    console.log("Balance:", ethers.utils.formatEther(await deployer.getBalance()), "BNB");
    
    const entryPointAddress = "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789";
    const deployedContracts: Record<string, string> = {};
    
    try {
        // 1. CPOPToken (已部署)
        deployedContracts.CPOPToken = "0xc2DEe82B70C60eDbF8a8E180Cb72A2e727574260";
        console.log("✅ CPOPToken:", deployedContracts.CPOPToken);
        
        // 2. Deploy GasPriceOracle (simple)
        console.log("\n📊 Deploying GasPriceOracle...");
        const GasPriceOracleFactory = await ethers.getContractFactory("contracts/cpop/GasPriceOracle.sol:GasPriceOracle");
        const gasPriceOracle = await GasPriceOracleFactory.deploy();
        await gasPriceOracle.deployed();
        deployedContracts.GasPriceOracle = gasPriceOracle.address;
        console.log("✅ GasPriceOracle:", gasPriceOracle.address);
        
        // 3. Deploy SessionKeyManager
        console.log("\n🔑 Deploying SessionKeyManager...");
        const SessionKeyManagerFactory = await ethers.getContractFactory("contracts/cpop/SessionKeyManager.sol:SessionKeyManager");
        const sessionKeyManager = await SessionKeyManagerFactory.deploy();
        await sessionKeyManager.deployed();
        deployedContracts.SessionKeyManager = sessionKeyManager.address;
        console.log("✅ SessionKeyManager:", sessionKeyManager.address);
        
        console.log("\n🎉 Simple deployment completed!");
        
    } catch (error) {
        console.error("❌ Deployment failed:", error);
        throw error;
    }
    
    // Summary
    console.log("\n📋 Deployed Contracts:");
    console.log("Network: BSC Testnet");
    console.log("EntryPoint:", entryPointAddress);
    console.log("");
    
    for (const [name, address] of Object.entries(deployedContracts)) {
        console.log(`${name}: ${address}`);
        console.log(`BSCScan: https://testnet.bscscan.com/address/${address}`);
        console.log("");
    }
    
    return deployedContracts;
}

main()
    .then((contracts) => {
        console.log("🎉 Deployment successful!");
        process.exit(0);
    })
    .catch((error) => {
        console.error("❌ Error:", error.message);
        process.exit(1);
    });