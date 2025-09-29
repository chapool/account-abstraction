import { ethers, upgrades } from "hardhat";
import { Contract } from "ethers";

async function main() {
  console.log("🚀 Deploying Staking Contracts...");

  // Contract addresses provided by user
  const ACCOUNT_MANAGER_ADDRESS = "0x2E4f862Ba3ee6D84dd19ae9f002F5D8c0C5675ef";
  const CPOP_TOKEN_ADDRESS = "0xA2d58d11c2752b010C2444fa2f795c6cf4cb76Bc";
  const CPNFT_ADDRESS = "0xcC63bf57EF4b4fE5635cF0745Ae7E2C75A63c7Ed";

  // Get the deployer
  const [deployer] = await ethers.getSigners();
  console.log("Deploying contracts with the account:", deployer.address);
  console.log("Account balance:", ethers.utils.formatEther(await deployer.getBalance()), "ETH");

  // ============================================
  // STEP 1: Deploy StakingConfig Contract
  // ============================================
  console.log("\n📋 Step 1: Deploying StakingConfig...");
  
  const StakingConfigFactory = await ethers.getContractFactory("StakingConfig");
  const stakingConfig = await StakingConfigFactory.deploy();
  await stakingConfig.deployed();

  console.log("✅ StakingConfig deployed to:", stakingConfig.address);

  // ============================================
  // STEP 2: Deploy Staking Contract (Proxy)
  // ============================================
  console.log("\n📋 Step 2: Deploying Staking Contract (Proxy)...");
  
  const StakingFactory = await ethers.getContractFactory("Staking");
  
  // Deploy as upgradeable proxy
  const staking = await upgrades.deployProxy(
    StakingFactory,
    [
      CPNFT_ADDRESS,
      CPOP_TOKEN_ADDRESS,
      ACCOUNT_MANAGER_ADDRESS,
      stakingConfig.address,
      deployer.address
    ],
    {
      initializer: "initialize",
      kind: "uups"
    }
  );
  
  await staking.deployed();
  console.log("✅ Staking deployed to:", staking.address);

  // ============================================
  // STEP 3: Configure StakingConfig
  // ============================================
  console.log("\n📋 Step 3: Configuring StakingConfig...");
  
  // Set the staking contract address in config
  const setStakingTx = await stakingConfig.setStakingContract(staking.address);
  await setStakingTx.wait();
  console.log("✅ Staking contract address set in config");

  // ============================================
  // STEP 4: Set up CPNFT Contract
  // ============================================
  console.log("\n📋 Step 4: Setting up CPNFT Contract...");
  
  const cpnft = await ethers.getContractAt("CPNFT", CPNFT_ADDRESS);
  
  // Set staking contract in CPNFT
  const setStakingContractTx = await cpnft.setStakingContract(staking.address);
  await setStakingContractTx.wait();
  console.log("✅ Staking contract set in CPNFT");

  // ============================================
  // STEP 5: Set up CPOP Token Permissions
  // ============================================
  console.log("\n📋 Step 5: Setting up CPOP Token Permissions...");
  
  const cpopToken = await ethers.getContractAt("CPOPToken", CPOP_TOKEN_ADDRESS);
  
  // Check if deployer has ADMIN_ROLE
  const ADMIN_ROLE = 1;
  const MINTER_ROLE = 2;
  const hasAdminRole = await cpopToken.hasRole(deployer.address, ADMIN_ROLE);
  console.log("Deployer has ADMIN_ROLE:", hasAdminRole);
  
  if (hasAdminRole) {
    // Grant MINTER_ROLE to staking contract
    const grantRoleTx = await cpopToken.grantRole(staking.address, MINTER_ROLE);
    await grantRoleTx.wait();
    console.log("✅ MINTER_ROLE granted to staking contract");
  } else {
    console.log("⚠️  Deployer does not have ADMIN_ROLE. Please grant MINTER_ROLE manually.");
    console.log("   Call: cpopToken.grantRole(staking.address, 2)");
  }

  // ============================================
  // STEP 6: Verify Deployment
  // ============================================
  console.log("\n🔍 Step 6: Verifying deployment...");
  
  // Check contract versions
  const stakingVersion = await staking.version();
  const configVersion = await stakingConfig.version();
  
  console.log("Staking contract version:", stakingVersion);
  console.log("StakingConfig contract version:", configVersion);

  // Verify contract addresses
  const stakingConfigAddress = await staking.configContract();
  const stakingCpnftAddress = await staking.cpnftContract();
  const stakingCpopAddress = await staking.cpopTokenContract();
  const stakingAccountManagerAddress = await staking.accountManagerContract();

  console.log("Contract addresses verification:");
  console.log("- StakingConfig:", stakingConfigAddress);
  console.log("- CPNFT:", stakingCpnftAddress);
  console.log("- CPOP Token:", stakingCpopAddress);
  console.log("- Account Manager:", stakingAccountManagerAddress);

  // Check basic configuration
  const minStakeDays = await stakingConfig.getMinStakeDays();
  const earlyWithdrawPenalty = await stakingConfig.getEarlyWithdrawPenalty();
  const quarterlyMultiplier = await stakingConfig.getQuarterlyMultiplier();

  console.log("Basic configuration:");
  console.log("- Min stake days:", minStakeDays.toString());
  console.log("- Early withdraw penalty:", earlyWithdrawPenalty.toString(), "basis points");
  console.log("- Quarterly multiplier:", quarterlyMultiplier.toString(), "basis points");

  // Check level configurations
  console.log("Level configurations:");
  for (let level = 1; level <= 6; level++) {
    const dailyReward = await stakingConfig.getDailyReward(level);
    const decayInterval = await stakingConfig.getDecayInterval(level);
    const decayRate = await stakingConfig.getDecayRate(level);
    const maxDecayRate = await stakingConfig.getMaxDecayRate(level);
    
    console.log(`Level ${level}:`, {
      dailyReward: ethers.utils.formatEther(dailyReward) + " CPP",
      decayInterval: decayInterval.toString() + " days",
      decayRate: decayRate.toString() + " bp",
      maxDecayRate: maxDecayRate.toString() + " bp"
    });
  }

  // ============================================
  // STEP 7: Save Deployment Info
  // ============================================
  console.log("\n📄 Step 7: Saving deployment information...");
  
  const deploymentInfo = {
    network: (await ethers.provider.getNetwork()).name,
    chainId: (await ethers.provider.getNetwork()).chainId,
    timestamp: new Date().toISOString(),
    deployer: deployer.address,
    contracts: {
      Staking: {
        address: staking.address,
        implementation: await upgrades.erc1967.getImplementationAddress(staking.address),
        admin: await upgrades.erc1967.getAdminAddress(staking.address),
        version: stakingVersion
      },
      StakingConfig: {
        address: stakingConfig.address,
        version: configVersion
      },
      dependencies: {
        CPNFT: CPNFT_ADDRESS,
        CPOPToken: CPOP_TOKEN_ADDRESS,
        AccountManager: ACCOUNT_MANAGER_ADDRESS
      }
    },
    configuration: {
      basic: {
        minStakeDays: minStakeDays.toString(),
        earlyWithdrawPenalty: earlyWithdrawPenalty.toString(),
        quarterlyMultiplier: quarterlyMultiplier.toString()
      },
      levels: {}
    }
  };

  // Add level configurations
  for (let level = 1; level <= 6; level++) {
    const levelConfig = await stakingConfig.getLevelConfig(level);
    deploymentInfo.configuration.levels[level] = {
      dailyReward: levelConfig.dailyReward.toString(),
      decayInterval: levelConfig.decayInterval.toString(),
      decayRate: levelConfig.decayRate.toString(),
      maxDecayRate: levelConfig.maxDecayRate.toString()
    };
  }

  // Write to file
  const fs = require('fs');
  const path = require('path');
  const outputPath = path.join(__dirname, '..', 'staking-deployment-results.json');
  fs.writeFileSync(outputPath, JSON.stringify(deploymentInfo, null, 2));
  console.log(`✅ Deployment info saved to: ${outputPath}`);

  // ============================================
  // DEPLOYMENT SUMMARY
  // ============================================
  console.log("\n🎉 Staking Contracts Deployment Completed Successfully!");
  console.log("==================================================");
  console.log("📋 Contract Addresses:");
  console.log("  Staking (Proxy):", staking.address);
  console.log("  StakingConfig:", stakingConfig.address);
  console.log("  CPNFT:", CPNFT_ADDRESS);
  console.log("  CPOP Token:", CPOP_TOKEN_ADDRESS);
  console.log("  Account Manager:", ACCOUNT_MANAGER_ADDRESS);
  console.log("  Owner:", deployer.address);
  
  console.log("\n📝 Next Steps:");
  console.log("1. Update your frontend to use the new staking contract addresses");
  console.log("2. Test staking functionality with test NFTs");
  console.log("3. Monitor staking events and rewards");
  console.log("4. Configure additional parameters as needed");
  console.log("5. Set up quarterly adjustment schedule");
  
  console.log("\n🔧 Available Functions:");
  console.log("- stake(uint256 tokenId): Stake a single NFT");
  console.log("- batchStake(uint256[] tokenIds): Stake multiple NFTs");
  console.log("- unstake(uint256 tokenId): Unstake and claim rewards");
  console.log("- batchUnstake(uint256[] tokenIds): Unstake multiple NFTs");
  console.log("- claimRewards(uint256 tokenId): Claim rewards without unstaking");
  console.log("- batchClaimRewards(uint256[] tokenIds): Claim rewards for multiple NFTs");
  console.log("- calculatePendingRewards(uint256 tokenId): View pending rewards");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error("❌ Deployment failed:", error);
  process.exitCode = 1;
});
