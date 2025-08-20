import { ethers } from "hardhat";

async function main() {
    console.log("Deploying CPOPToken to BSC Testnet...");
    
    // Get the deployer account
    const [deployer] = await ethers.getSigners();
    const deployerAddress = await deployer.getAddress();
    
    console.log("Deploying contracts with the account:", deployerAddress);
    console.log("Account balance:", (await deployer.getBalance()).toString());
    
    // Deploy CPOPToken
    const initialSupply = ethers.utils.parseEther("1000000"); // 1M tokens
    
    console.log("Deploying CPOPToken...");
    const CPOPTokenFactory = await ethers.getContractFactory("contracts/cpop/CPOPToken.sol:CPOPToken");
    
    const cpopToken = await CPOPTokenFactory.deploy(deployerAddress, initialSupply);
    
    console.log("Waiting for deployment...");
    await cpopToken.deployed();
    
    console.log("✅ CPOPToken deployed to:", cpopToken.address);
    console.log("✅ Admin address:", deployerAddress);
    console.log("✅ Initial supply:", ethers.utils.formatEther(initialSupply), "CPOP");
    
    // Wait for a few block confirmations
    console.log("Waiting for block confirmations...");
    await cpopToken.deployTransaction.wait(5);
    
    console.log("\n🎉 Deployment completed successfully!");
    
    // Verify deployment by checking token details
    console.log("\n📋 Token Details:");
    console.log("Name:", await cpopToken.name());
    console.log("Symbol:", await cpopToken.symbol());
    console.log("Decimals:", await cpopToken.decimals());
    console.log("Total Supply:", ethers.utils.formatEther(await cpopToken.totalSupply()));
    console.log("Admin Balance:", ethers.utils.formatEther(await cpopToken.balanceOf(deployerAddress)));
    
    // Check roles
    console.log("\n🔐 Admin Roles:");
    console.log("Has ADMIN_ROLE:", await cpopToken.hasRole(deployerAddress, 1));
    console.log("Has MINTER_ROLE:", await cpopToken.hasRole(deployerAddress, 2));
    console.log("Has BURNER_ROLE:", await cpopToken.hasRole(deployerAddress, 4));
    
    console.log("\n📝 Contract Address for verification:");
    console.log("CPOPToken:", cpopToken.address);
    console.log("Network: BSC Testnet (ChainID: 97)");
    console.log("Constructor args:", deployerAddress, initialSupply.toString());
    
    return {
        cpopToken: cpopToken.address,
        admin: deployerAddress,
        network: "BSC Testnet"
    };
}

main()
    .then((result) => {
        console.log("\n🚀 Deployment Summary:");
        console.log(result);
        process.exit(0);
    })
    .catch((error) => {
        console.error("❌ Deployment failed:", error);
        process.exit(1);
    });