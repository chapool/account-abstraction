import { ethers } from "hardhat";

/**
 * 测试用户每日收益统计功能
 */
async function main() {
  const stakingReaderAddress = "0x3a42F2F6962Fdc9bb8392F73C94e79071De5eC8d"; // 新部署的地址
  const userAddress = "0xc5cCc3c5e4bbb9519Deaf7a8afA29522DA49E33D"; // token 2645 的所有者
  
  console.log("\n" + "=".repeat(80));
  console.log("📊 测试用户每日收益统计功能");
  console.log("=".repeat(80));
  console.log(`StakingReader: ${stakingReaderAddress}`);
  console.log(`用户地址: ${userAddress}`);
  console.log("=".repeat(80) + "\n");

  const StakingReader = await ethers.getContractFactory("StakingReader");
  const reader = StakingReader.attach(stakingReaderAddress);

  try {
    // ===========================
    // 1. 获取用户所有NFT的每日收益
    // ===========================
    console.log("🎯 1. 用户所有NFT每日收益汇总:\n");
    
    const dailyRewards = await reader.getUserDailyRewards(userAddress);
    
    console.log("📈 总体统计:");
    console.log(`  总基础收益: ${ethers.utils.formatEther(dailyRewards.totalBaseReward)} CPOP/天`);
    console.log(`  总衰减后收益: ${ethers.utils.formatEther(dailyRewards.totalDecayedReward)} CPOP/天`);
    console.log(`  总最终收益: ${ethers.utils.formatEther(dailyRewards.totalFinalReward)} CPOP/天`);
    
    // totalBonus 格式：10000=100%, 11000=110%, 9000=90%
    const bonusValue = Number(dailyRewards.totalBonus);
    const bonusPercent = (bonusValue / 100).toFixed(2);
    const bonusChange = ((bonusValue - 10000) / 100).toFixed(2);
    const bonusSign = bonusValue > 10000 ? "+" : "";
    
    console.log(`  ⭐ 总乘数: ${bonusPercent}% (${bonusSign}${bonusChange}%)`);
    console.log(`  └─ Combo加成: +${(Number(dailyRewards.totalComboBonus) / 100).toFixed(2)}%`);
    console.log(`  └─ 动态乘数: ${(Number(dailyRewards.totalDynamicMultiplier) / 100).toFixed(2)}%\n`);

    console.log("📊 各等级详情:");
    const levelNames = ["C", "B", "A", "S", "SS", "SSS"];
    
    for (let i = 0; i < 6; i++) {
      const nftCount = Number(dailyRewards.nftCountPerLevel[i]);
      if (nftCount > 0) {
        const baseReward = ethers.utils.formatEther(dailyRewards.baseRewardPerLevel[i]);
        const finalReward = ethers.utils.formatEther(dailyRewards.finalRewardPerLevel[i]);
        const bonusPercent = ((Number(dailyRewards.finalRewardPerLevel[i]) / Number(dailyRewards.baseRewardPerLevel[i]) - 1) * 100).toFixed(2);
        
        console.log(`  Level ${levelNames[i]}:`);
        console.log(`    NFT数量: ${nftCount}`);
        console.log(`    基础收益: ${baseReward} CPOP/天`);
        console.log(`    最终收益: ${finalReward} CPOP/天`);
        console.log(`    加成幅度: +${bonusPercent}%\n`);
      }
    }

    // ===========================
    // 2. 按等级查询每日收益
    // ===========================
    console.log("\n🔍 2. 按等级查询每日收益:\n");
    
    for (let level = 1; level <= 6; level++) {
      try {
        const { baseReward, finalReward, nftCount } = await reader.getUserDailyRewardsByLevel(userAddress, level);
        
        if (Number(nftCount) > 0) {
          console.log(`Level ${levelNames[level - 1]}:`);
          console.log(`  质押数量: ${nftCount}`);
          console.log(`  基础收益: ${ethers.utils.formatEther(baseReward)} CPOP/天`);
          console.log(`  最终收益: ${ethers.utils.formatEther(finalReward)} CPOP/天`);
          
          if (Number(baseReward) > 0) {
            const totalBonus = ((Number(finalReward) / Number(baseReward) - 1) * 100).toFixed(2);
            console.log(`  总加成: +${totalBonus}%\n`);
          }
        }
      } catch (error: any) {
        // Level 可能没有质押的NFT
      }
    }

    // ===========================
    // 3. 查询单个NFT的每日收益详情
    // ===========================
    console.log("\n🔬 3. 单个NFT每日收益详情:\n");
    
    const tokenId = 2645;
    
    try {
      const breakdown = await reader.getNFTDailyRewardBreakdown(tokenId);
      
      console.log(`Token #${tokenId}:`);
      console.log(`  基础收益: ${ethers.utils.formatEther(breakdown.baseReward)} CPOP/天`);
      console.log(`  衰减后收益: ${ethers.utils.formatEther(breakdown.decayedReward)} CPOP/天`);
      console.log(`  Combo加成: +${(Number(breakdown.comboBonus) / 100).toFixed(2)}%`);
      console.log(`  动态乘数: ${(Number(breakdown.dynamicMultiplier) / 100).toFixed(2)}%`);
      console.log(`  最终收益: ${ethers.utils.formatEther(breakdown.finalReward)} CPOP/天\n`);
      
      // 计算各个加成的影响
      const decayEffect = Number(breakdown.decayedReward) > 0 ? 
        ((Number(breakdown.baseReward) / Number(breakdown.decayedReward) - 1) * 100).toFixed(2) : "0";
      const comboEffect = (Number(breakdown.comboBonus) / 100).toFixed(2);
      const dynamicEffect = ((Number(breakdown.dynamicMultiplier) / 10000 - 1) * 100).toFixed(2);
      
      console.log(`  📉 衰减影响: ${decayEffect}%`);
      console.log(`  🔥 Combo影响: +${comboEffect}%`);
      console.log(`  📊 动态乘数影响: ${dynamicEffect}%`);
      
    } catch (error: any) {
      console.log(`  ❌ Token ${tokenId} 未质押或查询失败: ${error.message}`);
    }

    // ===========================
    // 4. 计算预期月收益
    // ===========================
    console.log("\n\n💰 4. 预期收益估算:\n");
    
    const dailyTotal = Number(ethers.utils.formatEther(dailyRewards.totalFinalReward));
    const weeklyEstimate = dailyTotal * 7;
    const monthlyEstimate = dailyTotal * 30;
    const yearlyEstimate = dailyTotal * 365;
    
    console.log(`  每日收益: ${dailyTotal.toFixed(6)} CPOP`);
    console.log(`  每周收益: ${weeklyEstimate.toFixed(6)} CPOP`);
    console.log(`  每月收益: ${monthlyEstimate.toFixed(6)} CPOP`);
    console.log(`  每年收益: ${yearlyEstimate.toFixed(6)} CPOP`);
    
    console.log("\n⚠️  注意: 此估算基于当前收益率，实际收益会受以下因素影响:");
    console.log("   • 衰减机制（随时间递减）");
    console.log("   • Combo状态变化");
    console.log("   • 平台质押率变化");
    console.log("   • 季度调整");

    // ===========================
    // 5. 与用户质押汇总对比
    // ===========================
    console.log("\n\n📋 5. 与质押汇总对比:\n");
    
    const summary = await reader.getUserStakingSummary(userAddress);
    const pendingRewards = Number(ethers.utils.formatEther(summary.totalPendingRewards));
    
    console.log(`  质押NFT总数: ${summary.totalStakedCount}`);
    console.log(`  待领取奖励: ${pendingRewards.toFixed(6)} CPOP`);
    console.log(`  每日新增收益: ${dailyTotal.toFixed(6)} CPOP`);
    
    if (dailyTotal > 0) {
      const daysToDouble = pendingRewards / dailyTotal;
      console.log(`  \n  如果保持当前收益率:`);
      console.log(`    • ${daysToDouble.toFixed(1)} 天后，pending 翻倍`);
      console.log(`    • ${(30 * dailyTotal + pendingRewards).toFixed(6)} CPOP 可在30天后领取`);
    }

  } catch (error: any) {
    console.error("\n❌ 错误:", error.message);
    if (error.data) {
      console.error("Error data:", error.data);
    }
  }

  console.log("\n" + "=".repeat(80));
  console.log("✅ 测试完成!");
  console.log("=".repeat(80) + "\n");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

