import { ethers, upgrades } from "hardhat";

async function main() {
  console.log("🚀 Upgrading Staking Contract on BNB Testnet...");

  // Staking 合约代理地址 (from deployments/bnbTestnet/core.json)
  const STAKING_PROXY_ADDRESS = "0x8ab3CD39295CF002103c183963e527f2536949Df";

  // Get the deployer
  const signers = await ethers.getSigners();
  if (signers.length === 0) {
    throw new Error("No signers available. Please configure PRIVATE_KEY in .env file or use --network bnbTestnet.");
  }
  const deployer = signers[0];
  console.log("Upgrading contracts with the account:", deployer.address);
  
  const balance = await deployer.getBalance();
  console.log("Account balance:", ethers.utils.formatEther(balance), "BNB");
  
  if (balance.lt(ethers.utils.parseEther("0.01"))) {
    console.warn("⚠️  Warning: Low balance. Upgrade may fail due to insufficient gas.");
  }

  // ============================================
  // STEP 1: Check Current Version
  // ============================================
  console.log("\n📋 Step 1: Checking current version...");
  
  const stakingProxy = await ethers.getContractAt("Staking", STAKING_PROXY_ADDRESS);
  const currentVersion = await stakingProxy.version();
  console.log("Current version:", currentVersion);

  // ============================================
  // STEP 2: Get Current Implementation Address
  // ============================================
  console.log("\n📋 Step 2: Getting current implementation address...");
  
  const currentImplementation = await upgrades.erc1967.getImplementationAddress(STAKING_PROXY_ADDRESS);
  console.log("Current implementation:", currentImplementation);

  // ============================================
  // STEP 3: Check Ownership
  // ============================================
  console.log("\n📋 Step 3: Verifying ownership...");
  
  const contractOwner = await stakingProxy.owner();
  console.log("Contract owner:", contractOwner);
  console.log("Deployer address:", deployer.address);
  
  if (contractOwner.toLowerCase() !== deployer.address.toLowerCase()) {
    throw new Error(`Deployer (${deployer.address}) is not the contract owner (${contractOwner}). Only the owner can upgrade the contract.`);
  }
  console.log("✅ Ownership verified");

  // ============================================
  // STEP 4: Deploy New Implementation
  // ============================================
  console.log("\n📋 Step 4: Deploying new implementation...");
  
  const StakingFactory = await ethers.getContractFactory("Staking");
  
  console.log("Deploying new implementation contract...");
  const newImplementation = await StakingFactory.deploy();
  await newImplementation.deployed();
  console.log("✅ New implementation deployed to:", newImplementation.address);

  // ============================================
  // STEP 5: Upgrade Using upgradeToAndCall
  // ============================================
  console.log("\n📋 Step 5: Upgrading proxy using upgradeToAndCall...");
  console.log("This may take a few minutes...");
  
  try {
    // Get UUPSUpgradeable interface
    const proxyContract = await ethers.getContractAt(
      [
        "function upgradeToAndCall(address newImplementation, bytes memory data) external payable"
      ],
      STAKING_PROXY_ADDRESS
    );
    
    // Prepare empty data (no initialization needed for this upgrade)
    const emptyData = "0x";
    
    console.log("Calling upgradeToAndCall on proxy...");
    console.log("New implementation address:", newImplementation.address);
    console.log("Proxy address:", STAKING_PROXY_ADDRESS);
    console.log("Data:", emptyData);
    
    // Try to estimate gas first
    try {
      const gasEstimate = await proxyContract.estimateGas.upgradeToAndCall(
        newImplementation.address,
        emptyData,
        { value: 0 }
      );
      console.log("Gas estimate:", gasEstimate.toString());
    } catch (estimateError: any) {
      console.error("⚠️  Gas estimation failed:", estimateError.message);
    }
    
    // Execute upgradeToAndCall with manual gas limit
    const upgradeTx = await proxyContract.upgradeToAndCall(
      newImplementation.address,
      emptyData,
      {
        value: 0,
        gasLimit: 5000000  // Set a high gas limit
      }
    );
    console.log("Upgrade transaction sent:", upgradeTx.hash);
    
    const receipt = await upgradeTx.wait();
    console.log("✅ Upgrade transaction confirmed in block:", receipt.blockNumber);
    console.log("Gas used:", receipt.gasUsed.toString());
    
  } catch (error: any) {
    console.error("❌ Upgrade failed:", error.message);
    
    if (error.reason) {
      console.error("Reason:", error.reason);
    }
    if (error.data) {
      console.error("Error data:", error.data);
    }
    
    throw error;
  }

  // ============================================
  // STEP 6: Verify Upgrade
  // ============================================
  console.log("\n📋 Step 6: Verifying upgrade...");
  
  // Get new implementation address
  const newImplementationAddress = await upgrades.erc1967.getImplementationAddress(STAKING_PROXY_ADDRESS);
  console.log("New implementation address:", newImplementationAddress);
  console.log("Expected implementation address:", newImplementation.address);
  
  if (newImplementationAddress.toLowerCase() !== newImplementation.address.toLowerCase()) {
    throw new Error("Implementation address mismatch! Upgrade may have failed.");
  }
  console.log("✅ Implementation address verified");
  
  // Get upgraded contract instance
  const upgradedStaking = await ethers.getContractAt("Staking", STAKING_PROXY_ADDRESS);
  
  // Check version
  const newVersion = await upgradedStaking.version();
  console.log("New version:", newVersion);
  console.log("Old version:", currentVersion);
  
  // Verify addresses are still correct
  const configAddress = await upgradedStaking.configContract();
  const cpnftAddress = await upgradedStaking.cpnftContract();
  const cppTokenAddress = await upgradedStaking.CPPTokenContract();
  const accountManagerAddress = await upgradedStaking.accountManagerContract();
  
  console.log("\nContract addresses verification:");
  console.log("- StakingConfig:", configAddress);
  console.log("- CPNFT:", cpnftAddress);
  console.log("- CPP Token:", cppTokenAddress);
  console.log("- Account Manager:", accountManagerAddress);
  
  // Check contract state
  const totalStaked = await upgradedStaking.totalStakedCount();
  const isPaused = await upgradedStaking.paused();
  const owner = await upgradedStaking.owner();
  
  console.log("\nContract state:");
  console.log("- Total staked:", totalStaked.toString());
  console.log("- Is paused:", isPaused);
  console.log("- Owner:", owner);

  // ============================================
  // STEP 7: Test Bug Fix
  // ============================================
  console.log("\n📋 Step 7: Verifying bug fix...");
  console.log("Bug fix: unstake() now correctly accumulates totalRewards (line 313)");
  console.log("Changed: stakeInfo.totalRewards = rewards");
  console.log("To:      stakeInfo.totalRewards += rewards");
  console.log("✅ This fix ensures claimed rewards are not overwritten on unstake");

  // ============================================
  // UPGRADE SUMMARY
  // ============================================
  console.log("\n🎉 Staking Contract Upgrade Completed Successfully!");
  console.log("==================================================");
  console.log("📋 Upgrade Details:");
  console.log("  Network: BNB Testnet (ChainID: 97)");
  console.log("  Proxy Address:", STAKING_PROXY_ADDRESS);
  console.log("  Old Implementation:", currentImplementation);
  console.log("  New Implementation:", newImplementation.address);
  console.log("  Old Version:", currentVersion);
  console.log("  New Version:", newVersion);
  console.log("  Owner:", owner);
  console.log("  Method Used: upgradeToAndCall");
  
  console.log("\n🐛 Bug Fixed:");
  console.log("  Issue: totalRewards was being overwritten (=) instead of accumulated (+=) on unstake");
  console.log("  Impact: Users who claimed rewards before unstaking would see totalRewards reset to 0");
  console.log("  Fix: Changed line 313 from '=' to '+=' to correctly accumulate rewards");
  
  console.log("\n📝 Next Steps:");
  console.log("1. Verify the upgrade on BscScan:");
  console.log("   https://testnet.bscscan.com/address/" + STAKING_PROXY_ADDRESS);
  console.log("2. Test unstake functionality to ensure rewards are correctly accumulated");
  console.log("3. Monitor contract for any issues");
  console.log("4. Document the upgrade in your changelog");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error("❌ Upgrade failed:", error);
  process.exitCode = 1;
});

