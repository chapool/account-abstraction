import { ethers } from "hardhat";

async function simpleLocalTest() {
  console.log("🧪 Simple Local Contract Test\n");

  try {
    const accounts = await ethers.getSigners();
    const deployer = accounts[0];
    
    console.log(`📋 Deployer: ${deployer.address}`);
    const balance = await deployer.getBalance();
    console.log(`💰 Balance: ${ethers.utils.formatEther(balance)} ETH\n`);

    // Test deploying a simple contract first
    console.log("1️⃣ Testing CPOPToken deployment...");
    const CPOPToken = await ethers.getContractFactory("CPOPToken");
    const token = await CPOPToken.deploy("Test Token", "TEST", 18);
    await token.deployed();
    console.log(`✅ Token deployed: ${token.address}\n`);

    // Check token details
    const name = await token.name();
    const symbol = await token.symbol();
    const decimals = await token.decimals();
    console.log(`Token: ${name} (${symbol}) - Decimals: ${decimals}`);

    console.log("\n🎉 Simple local test successful!");
    console.log("🔄 Ready to test full deployment script!");

  } catch (error) {
    console.error("❌ Simple test failed:", error);
    throw error;
  }
}

simpleLocalTest()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("Simple test failed:", error);
    process.exit(1);
  });