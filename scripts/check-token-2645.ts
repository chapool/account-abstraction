import { ethers } from "hardhat";

async function main() {
  const tokenId = 2645;

  console.log(`\n🔍 检查 Token ID: ${tokenId}\n`);
  console.log("=".repeat(60));

  // 获取合约地址（从部署记录中获取）
  const cpnftAddress = "0xcC63bf57EF4b4fE5635cF0745Ae7E2C75A63c7Ed";
  const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5"; // 正确的 Staking 合约地址

  const CPNFT = await ethers.getContractFactory("CPNFT");
  const Staking = await ethers.getContractFactory("Staking");

  const cpnft = CPNFT.attach(cpnftAddress);
  const staking = Staking.attach(stakingAddress);

  try {
    // 1. 检查 NFT 基本信息
    console.log("\n📋 NFT 基本信息:");
    console.log("-".repeat(60));

    let owner;
    try {
      owner = await cpnft.ownerOf(tokenId);
      console.log(`✅ Owner: ${owner}`);
    } catch (error) {
      console.log(`❌ NFT 不存在或已销毁`);
      return;
    }

    // 2. 检查 NFT level
    let nftLevel;
    try {
      nftLevel = await cpnft.getTokenLevel(tokenId);
      const levelNames = ["NORMAL", "C", "B", "A", "S", "SS", "SSS"];
      console.log(`✅ NFT Level: ${nftLevel} (${levelNames[nftLevel] || "Unknown"})`);
    } catch (error) {
      console.log(`❌ 无法获取 Level: ${error.message}`);
    }

    // 3. 检查质押状态
    console.log("\n🔒 质押状态:");
    console.log("-".repeat(60));

    const stakeInfo = await staking.stakes(tokenId);
    
    console.log(`Owner in Staking: ${stakeInfo.owner}`);
    console.log(`Token ID: ${stakeInfo.tokenId}`);
    console.log(`Level in Staking: ${stakeInfo.level}`);
    console.log(`Is Active: ${stakeInfo.isActive}`);
    
    if (stakeInfo.isActive) {
      console.log(`✅ NFT 已质押`);
      console.log(`Stake Time: ${new Date(Number(stakeInfo.stakeTime) * 1000).toLocaleString()}`);
      console.log(`Last Claim Time: ${new Date(Number(stakeInfo.lastClaimTime) * 1000).toLocaleString()}`);
      console.log(`Total Rewards: ${ethers.utils.formatEther(stakeInfo.totalRewards)} CPOP`);
      console.log(`Pending Rewards: ${ethers.utils.formatEther(stakeInfo.pendingRewards)} CPOP`);
      
      // 计算质押天数
      const stakeDays = Math.floor((Date.now() / 1000 - Number(stakeInfo.stakeTime)) / 86400);
      console.log(`Staked Days: ${stakeDays} 天`);
    } else {
      console.log(`❌ NFT 未质押`);
    }

    // 4. 验证问题
    console.log("\n🐛 问题分析:");
    console.log("-".repeat(60));

    if (!stakeInfo.isActive) {
      console.log(`❌ 问题原因: NFT 未质押`);
      console.log(`   stakes[${tokenId}].level = ${stakeInfo.level} (默认值 0)`);
      console.log(`   调用 getPendingRewards 会失败，因为 level=0 不在有效范围 1-6`);
    } else if (stakeInfo.level < 1 || stakeInfo.level > 6) {
      console.log(`❌ 问题原因: Level 值无效`);
      console.log(`   stakes[${tokenId}].level = ${stakeInfo.level}`);
      console.log(`   有效范围: 1-6 (C, B, A, S, SS, SSS)`);
    } else if (nftLevel === 0) {
      console.log(`⚠️  问题原因: NFT 是 NORMAL level (不应该被质押)`);
      console.log(`   但质押记录中 level = ${stakeInfo.level}`);
      console.log(`   可能是质押后 level 被修改了`);
    } else {
      console.log(`✅ Level 值有效: ${stakeInfo.level}`);
      console.log(`   可以尝试查询 pending rewards`);
      
      try {
        const pendingRewards = await staking.calculatePendingRewards(tokenId);
        console.log(`   Pending Rewards: ${ethers.utils.formatEther(pendingRewards)} CPOP`);
      } catch (error: any) {
        console.log(`   ❌ 查询失败: ${error.message}`);
      }
    }

    // 5. 检查用户所有质押的 NFT
    console.log("\n👤 用户质押列表:");
    console.log("-".repeat(60));

    const userStakes = await staking.getUserStakes(owner);
    console.log(`用户 ${owner} 质押了 ${userStakes.length} 个 NFT:`);
    
    for (const id of userStakes) {
      const info = await staking.stakes(id);
      const levelNames = ["NORMAL", "C", "B", "A", "S", "SS", "SSS"];
      console.log(`  - Token ${id}: Level ${info.level} (${levelNames[info.level]}), Active: ${info.isActive}`);
    }

  } catch (error: any) {
    console.error("\n❌ 错误:", error.message);
    if (error.data) {
      console.error("Error data:", error.data);
    }
  }

  console.log("\n" + "=".repeat(60) + "\n");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

