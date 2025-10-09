import { ethers, upgrades } from "hardhat";
import { Contract } from "ethers";
import * as fs from "fs";
import * as path from "path";

async function main() {
  console.log("🚀 Deploying StakingReader Contract...");

  // Contract addresses (defaults, will try to override from staking-deployment-results.json)
  let STAKING_ADDRESS = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
  let CONFIG_ADDRESS = "0x37196054B23Be5CB977AA3284A3A844a8fe77861";
  let CPNFT_ADDRESS = "0xcC63bf57EF4b4fE5635cF0745Ae7E2C75A63c7Ed";

  try {
    const resultsPath = path.join(__dirname, "..", "staking-deployment-results.json");
    if (fs.existsSync(resultsPath)) {
      const deploymentInfo = JSON.parse(fs.readFileSync(resultsPath, "utf-8"));
      STAKING_ADDRESS = deploymentInfo?.contracts?.Staking?.address ?? STAKING_ADDRESS;
      CONFIG_ADDRESS = deploymentInfo?.contracts?.StakingConfig?.address ?? CONFIG_ADDRESS;
      CPNFT_ADDRESS = deploymentInfo?.contracts?.dependencies?.CPNFT ?? CPNFT_ADDRESS;
      console.log("Using addresses from staking-deployment-results.json");
    }
  } catch (err) {
    console.log("⚠️  Could not read staking-deployment-results.json, using default addresses");
  }

  // Get the deployer
  const [deployer] = await ethers.getSigners();
  console.log("Deploying contracts with the account:", deployer.address);
  console.log("Account balance:", ethers.utils.formatEther(await deployer.getBalance()), "ETH");

  // ============================================
  // STEP 1: Deploy StakingReader Contract (UUPS proxy)
  // ============================================
  console.log("\n📋 Step 1: Deploying StakingReader...");
  
  const StakingReaderFactory = await ethers.getContractFactory("StakingReader");
  const stakingReader = await upgrades.deployProxy(
    StakingReaderFactory,
    [STAKING_ADDRESS, CONFIG_ADDRESS, CPNFT_ADDRESS, (await ethers.getSigners())[0].address],
    { initializer: "initialize", kind: "uups" }
  );
  await stakingReader.deployed();
  const implementation = await upgrades.erc1967.getImplementationAddress(stakingReader.address);

  console.log("✅ StakingReader deployed to:", stakingReader.address);
  console.log("✅ StakingReader implementation:", implementation);

  // ============================================
  // STEP 2: Verify Deployment
  // ============================================
  console.log("\n🔍 Step 2: Verifying deployment...");
  
  // Test basic functionality
  try {
    // Test configuration query
    const config = await stakingReader.getConfiguration();
    console.log("✅ Configuration query successful");
    console.log("  - Min stake days:", config.minStakeDays.toString());
    console.log("  - Early withdraw penalty:", config.earlyWithdrawPenalty.toString());
    console.log("  - Quarterly multiplier:", config.quarterlyMultiplier.toString());

    // Test platform stats
    const platformStats = await stakingReader.getPlatformStats();
    console.log("✅ Platform stats query successful");
    console.log("  - Staked counts:", platformStats.staked.map(s => s.toString()));
    console.log("  - Supply counts:", platformStats.supply.map(s => s.toString()));

    // Test level stats
    const levelStats = await stakingReader.getLevelStats();
    console.log("✅ Level stats query successful");
    console.log("  - Staked per level:", levelStats.stakedCounts.map(s => s.toString()));
    console.log("  - Dynamic multipliers:", levelStats.dynamicMultipliers.map(m => (m/100).toFixed(2) + 'x'));

    // Test versions
    const versions = await stakingReader.getVersions();
    console.log("✅ Version query successful");
    console.log("  - Staking version:", versions.stakingVersion);
    console.log("  - Config version:", versions.configVersion);

    // Test historical adjustments
    const historicalCount = await stakingReader.getHistoricalAdjustmentCount();
    console.log("✅ Historical adjustments query successful");
    console.log("  - Historical adjustment count:", historicalCount.toString());

  } catch (error) {
    console.log("⚠️  Some queries failed:", error.message);
  }

  // ============================================
  // STEP 3: Save Deployment Info
  // ============================================
  console.log("\n📄 Step 3: Saving deployment information...");
  
  const deploymentInfo = {
    network: (await ethers.provider.getNetwork()).name,
    chainId: (await ethers.provider.getNetwork()).chainId,
    timestamp: new Date().toISOString(),
    deployer: deployer.address,
    contracts: {
      StakingReader: {
        address: stakingReader.address,
        version: "1.0.0"
      },
      dependencies: {
        Staking: STAKING_ADDRESS,
        StakingConfig: CONFIG_ADDRESS,
        CPNFT: CPNFT_ADDRESS
      }
    }
  };

  // Write to file
  const fs = require('fs');
  const path = require('path');
  const outputPath = path.join(__dirname, '..', 'staking-reader-deployment-results.json');
  fs.writeFileSync(outputPath, JSON.stringify(deploymentInfo, null, 2));
  console.log(`✅ Deployment info saved to: ${outputPath}`);

  // ============================================
  // DEPLOYMENT SUMMARY
  // ============================================
  console.log("\n🎉 StakingReader Contract Deployment Completed Successfully!");
  console.log("==================================================");
  console.log("📋 Contract Addresses:");
  console.log("  StakingReader:", stakingReader.address);
  console.log("  Staking:", STAKING_ADDRESS);
  console.log("  StakingConfig:", CONFIG_ADDRESS);
  console.log("  CPNFT:", CPNFT_ADDRESS);
  console.log("  Deployer:", deployer.address);
  
  console.log("\n📝 Next Steps:");
  console.log("1. Update frontend to use StakingReader contract for all queries");
  console.log("2. Test all query functions with the deployed contract");
  console.log("3. Update documentation with StakingReader address");
  
  console.log("\n🔧 Available Query Functions:");
  console.log("- getStakeDetails(uint256 tokenId): Get detailed stake information");
  console.log("- getPendingRewards(uint256 tokenId): Get pending rewards for a token");
  console.log("- getDetailedRewardCalculation(uint256 tokenId): Get reward breakdown");
  console.log("- getComboInfo(address user, uint8 level): Get combo bonus information");
  console.log("- getLevelStats(): Get statistics for all levels");
  console.log("- getPlatformStats(): Get platform statistics");
  console.log("- getPlatformStatistics(): Get comprehensive platform data");
  console.log("- getConfiguration(): Get all configuration parameters");
  console.log("- getVersions(): Get contract versions");
  console.log("- getHistoricalAdjustmentCount(): Get number of historical adjustments");
  console.log("- getAllHistoricalAdjustments(): Get all historical adjustment records");
  console.log("- getCompleteHistoricalAdjustment(uint256 index): Get complete historical data");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error("❌ Deployment failed:", error);
  process.exitCode = 1;
});
