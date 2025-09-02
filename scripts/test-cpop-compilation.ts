import { ethers } from "hardhat";

async function testCompilation() {
  console.log("🧪 Testing CPOP contract compilation...\n");

  try {
    // Test all contract factories
    const contracts = [
      "EntryPoint",
      "CPOPToken", 
      "GasPriceOracle",
      "GasPaymaster",
      "MasterAggregator",
      "SessionKeyManager",
      "AAccount",
      "AccountManager"
    ];

    for (const contractName of contracts) {
      console.log(`📋 Testing ${contractName}...`);
      const factory = await ethers.getContractFactory(contractName);
      console.log(`✅ ${contractName} compiled successfully`);
    }

    console.log("\n🎉 All CPOP contracts compile successfully!");
    console.log("✅ Ready for deployment");

  } catch (error) {
    console.error("❌ Compilation test failed:", error);
    throw error;
  }
}

testCompilation()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });