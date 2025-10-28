import { ethers } from "hardhat";
import * as dotenv from "dotenv";

dotenv.config();

async function main() {
  console.log("🔍 检查组合加成...\n");

  const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
  const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
  
  // 连接合约
  const [deployer] = await ethers.getSigners();
  const staking = await ethers.getContractAt("Staking", stakingAddress);

  console.log("👤 用户地址:", userAddress);
  console.log("📋 Staking 合约:", stakingAddress);
  console.log();

  // 1. 查询用户质押的 A 级 NFT（Level 3）
  console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
  console.log("📊 查询用户质押的 A 级 NFT");
  console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");

  try {
    // 获取用户质押的所有 NFT
    const userStakes: any = {};
    const tokenIds: number[] = [];
    
    // 遍历查找用户的质押 NFT
    let found = 0;
    for (let i = 1; i < 10000; i++) {
      try {
        const stakeInfo = await staking.stakes(i);
        if (stakeInfo.owner.toLowerCase() === userAddress.toLowerCase() && stakeInfo.isActive) {
          tokenIds.push(i);
          userStakes[i] = stakeInfo;
          found++;
          if (found >= 20) break; // 最多查 20 个
        }
      } catch (e) {
        continue;
      }
    }

    // 筛选 A 级 NFT（Level 3）
    const level3NFTs = tokenIds.filter(id => userStakes[id].level === 3);
    
    console.log(`✅ 找到 ${level3NFTs.length} 个 A 级 NFT\n`);
    
    if (level3NFTs.length === 0) {
      console.log("❌ 没有找到 A 级 NFT");
      return;
    }

    // 显示详细信息
    console.log("📦 质押的 A 级 NFT:");
    for (const tokenId of level3NFTs) {
      const stakeInfo = userStakes[tokenId];
      const stakeTime = Number(stakeInfo.stakeTime);
      const currentTime = Number(await staking.getCurrentTimestamp());
      const stakingDays = Math.floor((currentTime - stakeTime) / 86400);
      
      console.log(`  NFT #${tokenId}:`);
      console.log(`    质押天数: ${stakingDays} 天`);
      console.log(`    质押时间: ${new Date(stakeTime * 1000).toLocaleString()}`);
      
      // 查询待领取奖励
      const pendingRewards = await staking.calculatePendingRewards(tokenId);
      console.log(`    待领取: ${ethers.utils.formatEther(pendingRewards)} CPOP`);
    }
    console.log();

    // 2. 检查组合加成
    console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
    console.log("🎯 检查组合加成");
    console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");

    const comboStatus = await staking.getComboStatus(userAddress, 3);
    console.log("组合状态:");
    console.log("  当前数量:", comboStatus.currentCount.toString());
    console.log("  目标数量:", comboStatus.targetCount.toString());
    console.log("  等待期到期时间:", new Date(Number(comboStatus.deadline) * 1000).toLocaleString());
    console.log("  是否有效:", comboStatus.isEffective);
    console.log();

    const effectiveComboBonus = await staking.getEffectiveComboBonus(userAddress, 3);
    console.log("当前有效的组合加成:", effectiveComboBonus.toString(), "基点 (1 基点 = 0.01%)");
    console.log("预期加成:", "10 个 A 级 NFT 应该获得 2000 基点 (20%)");
    console.log();

    // 3. 计算理论加成
    console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
    console.log("📐 计算理论加成");
    console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");

    // 查询配置中的组合加成规则
    const configAddress = await staking.configContract();
    const config = await ethers.getContractAt("StakingConfig", configAddress);

    // 查询 10 个 NFT 应该获得的加成
    console.log("A 级 (Level 3) 组合加成规则:");
    
    // 直接调用合约的计算函数来查看
    console.log("\n尝试使用 StakingReader 查询详情...");
    
    const readerAddress = "0xbbEe6e5FF90f0B6EFF185F73c71b0deE6Fe9D0A6"; // 从文档中找到
    try {
      const reader = await ethers.getContractAt("StakingReader", readerAddress);
      const rewardBreakdown = await reader.getUserLevelRewardBreakdown(userAddress, 3);
      
      console.log("\n用户 A 级 NFT 奖励分解:");
      console.log("  总 NFT 数量:", rewardBreakdown.totalNFTs.toString());
      console.log("  基础奖励总和:", ethers.utils.formatEther(rewardBreakdown.totalBaseReward), "CPOP");
      console.log("  最终奖励总和:", ethers.utils.formatEther(rewardBreakdown.totalFinalReward), "CPOP");
      console.log("  综合乘数:", (Number(rewardBreakdown.combinedMultiplier) / 10000).toFixed(2), "x");
      
      const totalComboBonusBp = rewardBreakdown.totalComboBonus;
      const comboBonusPercent = Number(totalComboBonusBp) / 100;
      console.log("  组合加成:", comboBonusPercent.toFixed(2) + "%");
      
      if (comboBonusPercent < 20) {
        console.log("\n❌ 问题发现: 组合加成不足 20%");
        console.log("  原因分析:");
        console.log("  1. 检查目标数量:", comboStatus.targetCount.toString());
        console.log("  2. 检查等待期:", new Date(Number(comboStatus.deadline) * 1000).toLocaleString());
        console.log("  3. 检查当前数量:", comboStatus.currentCount.toString());
        console.log("  4. 检查是否有效:", comboStatus.isEffective);
        
        // 检查是否有等待期
        const now = await staking.getCurrentTimestamp();
        if (Number(comboStatus.deadline) > Number(now)) {
          const remainingDays = Math.floor((Number(comboStatus.deadline) - Number(now)) / 86400);
          console.log(`\n  ⚠️  等待期未结束，还需 ${remainingDays} 天`);
        }
      } else {
        console.log("\n✅ 组合加成正常 (≥20%)");
      }
      
    } catch (error: any) {
      console.log("⚠️  无法查询 Reader:", error.message);
    }

  } catch (error: any) {
    console.error("❌ 查询失败:", error.message);
  }

  console.log("\n✅ 查询完成");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

