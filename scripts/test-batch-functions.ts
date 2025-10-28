import { ethers } from "hardhat";
import * as dotenv from "dotenv";

dotenv.config();

async function main() {
  console.log("🧪 测试批量函数...\n");

  // 获取部署者账户（owner）
  const [deployer] = await ethers.getSigners();
  console.log("📝 Deployer/Owner:", deployer.address);

  // 测试用的用户地址
  const testUserAddress = process.env.TEST_USER_ADDRESS || "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
  console.log("👤 Test User:", testUserAddress);
  console.log();

  // 连接合约
  const stakingAddress = process.env.STAKING_ADDRESS || "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
  const staking = await ethers.getContractAt("Staking", stakingAddress);

  console.log("📋 合约信息:");
  console.log("  - Address:", stakingAddress);
  console.log("  - Owner:", deployer.address);
  console.log();

  // 测试 1: 查询用户质押的 NFT
  console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
  console.log("📊 测试 1: 查询用户质押的 NFT");
  console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");

  try {
    // 获取用户质押的所有 NFT
    const userStakes = await staking.getUserStakes(testUserAddress);
    console.log(`✅ 用户质押了 ${userStakes.length} 个 NFT`);
    
    if (userStakes.length === 0) {
      console.log("⚠️  用户没有质押的 NFT，无法测试");
      return;
    }

    // 显示前 5 个 NFT
    console.log("\n📦 质押的 NFT:");
    for (let i = 0; i < Math.min(5, userStakes.length); i++) {
      const tokenId = userStakes[i];
      const stakeInfo = await staking.stakes(tokenId);
      const pendingRewards = await staking.calculatePendingRewards(tokenId);
      
      console.log(`  - NFT #${tokenId}:`);
      console.log(`    质押天数: ${Math.floor((Date.now() / 1000 - Number(stakeInfo.stakeTime)) / 86400)} 天`);
      console.log(`    待领取: ${ethers.utils.formatEther(pendingRewards)} CPOP`);
    }
    console.log();

  } catch (error: any) {
    console.log("❌ 查询失败:", error.message);
    return;
  }

  // 测试 2: 测试 batchClaimRewards
  console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
  console.log("🧪 测试 2: batchClaimRewards");
  console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");

  const testTokenId = userStakes[0]; // 使用第一个 NFT
  
  try {
    // 查询领取前的奖励
    const rewardsBefore = await staking.calculatePendingRewards(testTokenId);
    console.log(`📊 NFT #${testTokenId} 待领取奖励:`, ethers.utils.formatEther(rewardsBefore), "CPOP");
    
    if (rewardsBefore.isZero()) {
      console.log("⚠️  没有待领取的奖励，跳过测试");
    } else {
      console.log("\n📤 调用 batchClaimRewards...");
      
      // 估算 Gas
      const gasEstimate = await staking.estimateGas.batchClaimRewards(
        testUserAddress,
        [testTokenId]
      );
      console.log(`⛽ 预计 Gas: ${gasEstimate.toString()}`);

      // 执行领取（需要 owner 权限）
      console.log("⏳ 执行交易...");
      const tx = await staking.batchClaimRewards(
        testUserAddress,
        [testTokenId]
      );
      
      console.log("✅ 交易已发送:", tx.hash);
      
      const receipt = await tx.wait();
      console.log("✅ 交易已确认，区块:", receipt.blockNumber);
      console.log("💸 Gas 使用:", receipt.gasUsed.toString());

      // 查询领取后的奖励
      const rewardsAfter = await staking.calculatePendingRewards(testTokenId);
      console.log(`📊 NFT #${testTokenId} 剩余待领取:`, ethers.utils.formatEther(rewardsAfter), "CPOP");
    }
    
  } catch (error: any) {
    console.log("❌ 领取失败:", error.message);
  }

  console.log();

  // 测试 3: 测试 batchUnstake（仅显示信息，不实际执行）
  console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
  console.log("🧪 测试 3: batchUnstake（预览）");
  console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");

  try {
    const testTokenIds = userStakes.slice(0, 3); // 取前 3 个
    console.log("📝 准备解质押的 NFT:", testTokenIds);
    
    // 计算每个 NFT 的奖励
    let totalRewards = ethers.BigNumber.from(0);
    for (const tokenId of testTokenIds) {
      const rewards = await staking.calculatePendingRewards(tokenId);
      const total = await staking._calculateTotalRewards(tokenId);
      totalRewards = totalRewards.add(total);
      
      console.log(`  - NFT #${tokenId}: 总奖励 ${ethers.utils.formatEther(total)} CPOP`);
    }
    
    console.log(`\n💰 总奖励: ${ethers.utils.formatEther(totalRewards)} CPOP`);
    
    // 估算 Gas
    const gasEstimate = await staking.estimateGas.batchUnstake(
      testUserAddress,
      testTokenIds
    );
    console.log(`⛽ 预计 Gas: ${gasEstimate.toString()}`);
    
    console.log("\n⚠️  注意: 这只是预览，没有实际执行解质押");
    
  } catch (error: any) {
    console.log("❌ 预览失败:", error.message);
  }

  console.log();
  console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
  console.log("✅ 测试完成");
  console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

