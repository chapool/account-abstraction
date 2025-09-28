import { ethers } from "hardhat";
import { Contract } from "ethers";

async function main() {
  console.log("🚀 Deploying CPNFTStaking Contract...");

  // Get the contract factory
  const CPNFTStakingFactory = await ethers.getContractFactory("CPNFTStaking");

  // Get deployed contracts addresses (adjust these based on your deployment)
  const CPNFT_ADDRESS = process.env.CPNFT_ADDRESS || "0x..."; // Replace with actual CPNFT address
  const CPP_TOKEN_ADDRESS = process.env.CPP_TOKEN_ADDRESS || "0x..."; // Replace with actual CPP token address
  const ACCOUNT_MANAGER_ADDRESS = process.env.ACCOUNT_MANAGER_ADDRESS || "0x..."; // Replace with actual AccountManager address

  // Get the deployer
  const [deployer] = await ethers.getSigners();
  console.log("Deploying contracts with the account:", deployer.address);
  console.log("Account balance:", (await deployer.getBalance()).toString());

  // Deploy the contract
  console.log("Deploying CPNFTStaking...");
  const staking = await CPNFTStakingFactory.deploy();
  await staking.deployed();

  console.log("✅ CPNFTStaking deployed to:", staking.address);

  // Initialize the contract
  console.log("Initializing CPNFTStaking...");
  const initTx = await staking.initialize(
    CPNFT_ADDRESS,
    CPP_TOKEN_ADDRESS,
    ACCOUNT_MANAGER_ADDRESS,
    deployer.address
  );
  await initTx.wait();

  console.log("✅ CPNFTStaking initialized successfully");

  // Set up CPNFT contract to recognize this staking contract
  console.log("Setting up CPNFT contract...");
  const cpnft = await ethers.getContractAt("CPNFT", CPNFT_ADDRESS);
  const setStakingTx = await cpnft.setStakingContract(staking.address);
  await setStakingTx.wait();

  console.log("✅ CPNFT staking contract set successfully");

  // Grant MINTER_ROLE to staking contract for CPP token
  console.log("Granting MINTER_ROLE to staking contract...");
  const cppToken = await ethers.getContractAt("CPOPToken", CPP_TOKEN_ADDRESS);
  const MINTER_ROLE = 2; // MINTER_ROLE constant from CPOPToken
  const grantRoleTx = await cppToken.grantRole(staking.address, MINTER_ROLE);
  await grantRoleTx.wait();

  console.log("✅ MINTER_ROLE granted to staking contract");

  // Set initial total supply per level (example values - adjust as needed)
  console.log("Setting initial total supply per level...");
  const setSupplyTx = await staking.setTotalSupplyPerLevel(1, 1000); // C level: 1000 supply
  await setSupplyTx.wait();
  
  const setSupplyTx2 = await staking.setTotalSupplyPerLevel(2, 800); // B level: 800 supply
  await setSupplyTx2.wait();
  
  const setSupplyTx3 = await staking.setTotalSupplyPerLevel(3, 500); // A level: 500 supply
  await setSupplyTx3.wait();
  
  const setSupplyTx4 = await staking.setTotalSupplyPerLevel(4, 200); // S level: 200 supply
  await setSupplyTx4.wait();
  
  const setSupplyTx5 = await staking.setTotalSupplyPerLevel(5, 100); // SS level: 100 supply
  await setSupplyTx5.wait();
  
  const setSupplyTx6 = await staking.setTotalSupplyPerLevel(6, 50); // SSS level: 50 supply
  await setSupplyTx6.wait();

  console.log("✅ Initial total supply per level set");

  // Verify deployment
  console.log("\n🔍 Verifying deployment...");
  
  const config = await staking.getConfiguration();
  console.log("Staking config loaded:", {
    minStakeDays: config.staking.minStakeDays.toString(),
    earlyWithdrawPenalty: config.staking.earlyWithdrawPenalty.toString(),
    dailyRewards: config.staking.dailyRewards.map(r => r.toString())
  });

  const platformStats = await staking.getPlatformStats();
  console.log("Platform stats:", {
    staked: platformStats.staked.map(s => s.toString()),
    supply: platformStats.supply.map(s => s.toString())
  });

  console.log("\n📋 Deployment Summary:");
  console.log("====================");
  console.log("CPNFTStaking Address:", staking.address);
  console.log("CPNFT Address:", CPNFT_ADDRESS);
  console.log("CPP Token Address:", CPP_TOKEN_ADDRESS);
  console.log("Account Manager Address:", ACCOUNT_MANAGER_ADDRESS);
  console.log("Owner:", deployer.address);
  console.log("Version:", await staking.version());

  // Save deployment info
  const deploymentInfo = {
    network: (await ethers.provider.getNetwork()).name,
    chainId: (await ethers.provider.getNetwork()).chainId,
    timestamp: new Date().toISOString(),
    contracts: {
      CPNFTStaking: {
        address: staking.address,
        transactionHash: staking.deployTransaction.hash,
        blockNumber: staking.deployTransaction.blockNumber,
        gasUsed: staking.deployTransaction.gasLimit?.toString()
      },
      CPNFT: CPNFT_ADDRESS,
      CPPToken: CPP_TOKEN_ADDRESS,
      AccountManager: ACCOUNT_MANAGER_ADDRESS
    },
    configuration: {
      minStakeDays: config.staking.minStakeDays.toString(),
      earlyWithdrawPenalty: config.staking.earlyWithdrawPenalty.toString(),
      dailyRewards: config.staking.dailyRewards.map(r => r.toString()),
      decayIntervals: config.staking.decayIntervals.map(d => d.toString()),
      decayRates: config.staking.decayRates.map(r => r.toString()),
      maxDecayRates: config.staking.maxDecayRates.map(r => r.toString()),
      comboThresholds: config.combo.comboThresholds.map(t => t.toString()),
      comboBonuses: config.combo.comboBonuses.map(b => b.toString()),
      comboMinDays: config.combo.comboMinDays.map(d => d.toString()),
      continuousThresholds: config.continuous.thresholds.map(t => t.toString()),
      continuousBonuses: config.continuous.bonuses.map(b => b.toString())
    }
  };

  // Write to file
  const fs = require('fs');
  const path = require('path');
  const outputPath = path.join(__dirname, '..', 'staking-deployment-results.json');
  fs.writeFileSync(outputPath, JSON.stringify(deploymentInfo, null, 2));
  console.log(`\n📄 Deployment info saved to: ${outputPath}`);

  console.log("\n🎉 CPNFTStaking deployment completed successfully!");
  console.log("\n📝 Next steps:");
  console.log("1. Update your frontend to use the new staking contract address");
  console.log("2. Set up monitoring for the staking events");
  console.log("3. Configure the reward parameters as needed");
  console.log("4. Test the staking functionality with test NFTs");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error("❌ Deployment failed:", error);
  process.exitCode = 1;
});
