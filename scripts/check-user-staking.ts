import { ethers } from "hardhat";

/**
 * 检查用户的NFT质押和pending奖励情况
 */
async function main() {
  // 配置
  const userAddress = "0xa3B605fB633AD0A0DC4B74b10bBfc40fDB050d35";
  const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
  const stakingReaderAddress = "0x3eCA9E28583C469CAA654fcb49Bd94Ce25C0262C"; // ✅ 更新为新地址
  
  console.log("=".repeat(80));
  console.log("检查用户NFT质押和奖励状态");
  console.log("=".repeat(80));
  console.log(`用户地址: ${userAddress}`);
  console.log(`Staking合约: ${stakingAddress}`);
  console.log(`StakingReader合约: ${stakingReaderAddress}`);
  console.log(`检查时间: ${new Date().toLocaleString('zh-CN', { timeZone: 'Asia/Shanghai' })}`);
  console.log("=".repeat(80));
  console.log("");

  try {
    // 获取合约实例
    const Staking = await ethers.getContractAt("Staking", stakingAddress);
    const StakingReader = await ethers.getContractAt("StakingReader", stakingReaderAddress);

    // 1. 获取用户质押汇总
    console.log("📊 用户质押汇总:");
    console.log("-".repeat(80));
    
    const summary = await StakingReader.getUserStakingSummary(userAddress);
    console.log(`总质押数量: ${summary.totalStakedCount.toString()}`);
    console.log(`总已领取奖励: ${ethers.utils.formatEther(summary.totalClaimedRewards)} CPOP`);
    console.log(`总待领取奖励: ${ethers.utils.formatEther(summary.totalPendingRewards)} CPOP`);
    console.log(`最长质押时长: ${(Number(summary.longestStakingDuration) / 86400).toFixed(2)} 天`);
    console.log(`平均有效乘数: ${(Number(summary.totalEffectiveMultiplier) / 100).toFixed(2)}%`);
    
    console.log("\n各等级质押数量:");
    const levelNames = ["C", "B", "A", "S", "SS", "SSS"];
    for (let i = 0; i < 6; i++) {
      if (Number(summary.levelStakingCounts[i]) > 0) {
        console.log(`  Level ${levelNames[i]}: ${summary.levelStakingCounts[i].toString()} 个`);
      }
    }
    console.log("");

    // 2. 获取用户质押的NFT详情
    console.log("🎯 质押的NFT详情:");
    console.log("-".repeat(80));
    
    if (Number(summary.totalStakedCount) > 0) {
      const { nfts, total } = await StakingReader.getUserStakedNFTs(userAddress, 0, 100);
      
      console.log(`找到 ${total.toString()} 个质押的NFT:\n`);
      
      for (let i = 0; i < nfts.length; i++) {
        const nft = nfts[i];
        if (nft.tokenId == 0n) continue; // 跳过空条目
        
        console.log(`NFT #${nft.tokenId.toString()}:`);
        console.log(`  等级: Level ${levelNames[Number(nft.level) - 1]} (${nft.level})`);
        console.log(`  质押时长: ${(Number(nft.stakingDuration) / 86400).toFixed(2)} 天`);
        console.log(`  待领取奖励: ${ethers.utils.formatEther(nft.pendingRewards)} CPOP`);
        console.log(`  已领取总奖励: ${ethers.utils.formatEther(nft.totalRewards)} CPOP`);
        console.log(`  有效乘数: ${(Number(nft.effectiveMultiplier) / 100).toFixed(2)}%`);
        
        // 获取详细的质押信息
        const stakeInfo = await StakingReader.getStakeDetails(nft.tokenId);
        const stakeTime = new Date(Number(stakeInfo.stakeTime) * 1000);
        const lastClaimTime = new Date(Number(stakeInfo.lastClaimTime) * 1000);
        
        console.log(`  质押时间: ${stakeTime.toLocaleString('zh-CN', { timeZone: 'Asia/Shanghai' })}`);
        console.log(`  上次领取: ${lastClaimTime.toLocaleString('zh-CN', { timeZone: 'Asia/Shanghai' })}`);
        console.log(`  状态: ${stakeInfo.isActive ? '✅ 活跃' : '❌ 已取消'}`);
        
        // 计算每日收益
        const dailyReward = await StakingReader.getDailyReward(nft.level);
        console.log(`  基础日收益: ${ethers.utils.formatEther(dailyReward)} CPOP/天`);
        
        // 计算预期收益率
        if (Number(nft.stakingDuration) > 0) {
          const daysStaked = Number(nft.stakingDuration) / 86400;
          const totalEarned = Number(ethers.utils.formatEther(nft.totalRewards)) + Number(ethers.utils.formatEther(nft.pendingRewards));
          const dailyAverage = totalEarned / daysStaked;
          console.log(`  实际日均收益: ${dailyAverage.toFixed(6)} CPOP/天`);
        }
        
        console.log("");
      }
    } else {
      console.log("❌ 该用户当前没有质押任何NFT\n");
    }

    // 3. 获取用户收益统计
    console.log("💰 收益统计:");
    console.log("-".repeat(80));
    
    const rewardStats = await StakingReader.getUserRewardStats(userAddress);
    console.log(`历史总奖励: ${ethers.utils.formatEther(rewardStats.totalHistoricalRewards)} CPOP`);
    console.log(`当前待领取: ${ethers.utils.formatEther(rewardStats.totalPendingRewards)} CPOP`);
    console.log(`过去24小时收益: ${ethers.utils.formatEther(rewardStats.last24HoursRewards)} CPOP`);
    console.log(`平均日收益: ${ethers.utils.formatEther(rewardStats.averageDailyRewards)} CPOP/天`);
    
    console.log("\n各等级收益分布:");
    for (let i = 0; i < 6; i++) {
      if (Number(rewardStats.rewardsPerLevel[i]) > 0) {
        console.log(`  Level ${levelNames[i]}: ${ethers.utils.formatEther(rewardStats.rewardsPerLevel[i])} CPOP`);
      }
    }
    console.log("");

    // 4. 获取Combo状态
    console.log("🔥 Combo状态:");
    console.log("-".repeat(80));
    
    const comboSummary = await StakingReader.getUserComboSummary(userAddress);
    const comboConfig = await StakingReader.getComboConfig();
    
    console.log("当前各等级Combo状态:");
    for (let i = 0; i < 6; i++) {
      const count = Number(comboSummary.currentComboCounts[i]);
      if (count > 0) {
        const bonus = Number(comboSummary.comboBonus[i]);
        const isPending = comboSummary.hasPendingCombo[i];
        const nextThreshold = Number(comboSummary.nextComboThreshold[i]);
        
        console.log(`  Level ${levelNames[i]}:`);
        console.log(`    质押数量: ${count}`);
        console.log(`    当前加成: +${(bonus / 100).toFixed(2)}%`);
        console.log(`    待生效: ${isPending ? '是' : '否'}`);
        if (nextThreshold !== Number.MAX_SAFE_INTEGER) {
          console.log(`    距离下一档: 还需 ${nextThreshold} 个`);
        } else {
          console.log(`    状态: 已达最高档`);
        }
      }
    }
    
    console.log("\nCombo配置:");
    console.log(`  档位: ${comboConfig.thresholds.map(t => t.toString()).join(', ')}`);
    console.log(`  加成: ${comboConfig.bonuses.map(b => `+${Number(b) / 100}%`).join(', ')}`);
    console.log("");

    // 5. 获取平台统计
    console.log("📈 平台统计:");
    console.log("-".repeat(80));
    
    const levelStats = await StakingReader.getLevelStats();
    console.log("各等级质押情况:");
    for (let i = 1; i <= 6; i++) {
      const staked = Number(levelStats.stakedCounts[i]);
      const supply = Number(levelStats.supplies[i]);
      const ratio = Number(levelStats.stakingRatios[i]) / 100;
      const multiplier = Number(levelStats.dynamicMultipliers[i]) / 100;
      
      if (supply > 0) {
        console.log(`  Level ${levelNames[i - 1]}:`);
        console.log(`    已质押: ${staked} / ${supply} (${ratio.toFixed(2)}%)`);
        console.log(`    动态乘数: ${multiplier.toFixed(2)}%`);
      }
    }
    console.log("");

    // 6. 验证奖励计算是否正常
    console.log("🔍 奖励验证:");
    console.log("-".repeat(80));
    
    if (Number(summary.totalStakedCount) > 0) {
      const { nfts } = await StakingReader.getUserStakedNFTs(userAddress, 0, 1);
      
      if (nfts.length > 0 && nfts[0].tokenId != 0n) {
        const tokenId = nfts[0].tokenId;
        const stakeInfo = await StakingReader.getStakeDetails(tokenId);
        const stakeTimestamp = Number(stakeInfo.stakeTime);
        const nowTimestamp = Math.floor(Date.now() / 1000);
        const daysStaked = Math.floor((nowTimestamp - stakeTimestamp) / 86400);
        
        console.log(`检查NFT #${tokenId.toString()}的奖励计算:`);
        console.log(`  质押天数: ${daysStaked} 天`);
        
        // 检查每天的奖励分解
        if (daysStaked > 0) {
          const dayToCheck = Math.min(daysStaked - 1, 0); // 检查第一天
          try {
            const breakdown = await Staking.getDailyRewardBreakdown(tokenId, dayToCheck);
            console.log(`\n  第 ${dayToCheck + 1} 天奖励分解:`);
            console.log(`    基础奖励: ${ethers.utils.formatEther(breakdown.baseReward)} CPOP`);
            console.log(`    衰减后奖励: ${ethers.utils.formatEther(breakdown.decayedReward)} CPOP`);
            console.log(`    季度乘数: ${Number(breakdown.quarterlyMultiplier) / 100}%`);
            console.log(`    动态乘数: ${Number(breakdown.dynamicMultiplier) / 100}%`);
            console.log(`    最终奖励: ${ethers.utils.formatEther(breakdown.finalReward)} CPOP`);
          } catch (error: any) {
            console.log(`    ⚠️ 无法获取详细分解: ${error.message}`);
          }
        }
        
        // 检查pending奖励是否合理
        const pendingRewards = Number(ethers.utils.formatEther(stakeInfo.pendingRewards));
        const dailyReward = Number(ethers.utils.formatEther(await StakingReader.getDailyReward(stakeInfo.level)));
        
        // 获取上次领取时间到现在的天数
        const lastClaimTimestamp = Number(stakeInfo.lastClaimTime);
        const daysSinceLastClaim = (nowTimestamp - lastClaimTimestamp) / 86400;
        
        console.log(`\n  奖励合理性检查:`);
        console.log(`    距上次领取: ${daysSinceLastClaim.toFixed(2)} 天`);
        console.log(`    基础日收益: ${dailyReward.toFixed(6)} CPOP`);
        console.log(`    预期收益范围: ${(dailyReward * daysSinceLastClaim * 0.2).toFixed(6)} - ${(dailyReward * daysSinceLastClaim * 1.5).toFixed(6)} CPOP`);
        console.log(`    实际pending: ${pendingRewards.toFixed(6)} CPOP`);
        
        // 判断是否在合理范围内
        const minExpected = dailyReward * daysSinceLastClaim * 0.2; // 考虑衰减，最低20%
        const maxExpected = dailyReward * daysSinceLastClaim * 1.5; // 考虑combo加成，最高150%
        
        if (pendingRewards >= minExpected && pendingRewards <= maxExpected) {
          console.log(`    ✅ 奖励计算正常`);
        } else if (pendingRewards < minExpected) {
          console.log(`    ⚠️ 奖励可能偏低，请检查衰减和乘数配置`);
        } else {
          console.log(`    ⚠️ 奖励可能偏高，请检查combo和乘数配置`);
        }
      }
    }
    console.log("");

    // 7. 检查昨天的质押记录
    console.log("📅 昨天的质押检查:");
    console.log("-".repeat(80));
    
    const now = Math.floor(Date.now() / 1000);
    const oneDayAgo = now - 86400;
    const twoDaysAgo = now - 172800;
    
    console.log(`检查时间范围: ${new Date(twoDaysAgo * 1000).toLocaleString('zh-CN')} - ${new Date(oneDayAgo * 1000).toLocaleString('zh-CN')}`);
    
    if (Number(summary.totalStakedCount) > 0) {
      const { nfts } = await StakingReader.getUserStakedNFTs(userAddress, 0, 100);
      let yesterdayStakes = 0;
      
      for (let i = 0; i < nfts.length; i++) {
        const nft = nfts[i];
        if (nft.tokenId == 0n) continue;
        
        const stakeInfo = await StakingReader.getStakeDetails(nft.tokenId);
        const stakeTime = Number(stakeInfo.stakeTime);
        
        if (stakeTime >= twoDaysAgo && stakeTime <= oneDayAgo) {
          yesterdayStakes++;
          console.log(`\n  ✅ NFT #${nft.tokenId.toString()} 在昨天质押:`);
          console.log(`     质押时间: ${new Date(stakeTime * 1000).toLocaleString('zh-CN')}`);
          console.log(`     等级: Level ${levelNames[Number(nft.level) - 1]}`);
          console.log(`     当前pending: ${ethers.utils.formatEther(nft.pendingRewards)} CPOP`);
          console.log(`     状态: ${stakeInfo.isActive ? '✅ 正常' : '❌ 已取消'}`);
        }
      }
      
      if (yesterdayStakes === 0) {
        console.log(`  ℹ️ 昨天没有新的质押记录`);
      } else {
        console.log(`\n  总计: 昨天质押了 ${yesterdayStakes} 个NFT`);
      }
    }
    console.log("");

    // 8. 总结
    console.log("=".repeat(80));
    console.log("📋 检查总结:");
    console.log("=".repeat(80));
    console.log(`✅ 用户地址: ${userAddress}`);
    console.log(`✅ 总质押数量: ${summary.totalStakedCount.toString()} 个NFT`);
    console.log(`✅ 总待领取奖励: ${ethers.utils.formatEther(summary.totalPendingRewards)} CPOP`);
    console.log(`✅ 总已领取奖励: ${ethers.utils.formatEther(summary.totalClaimedRewards)} CPOP`);
    
    if (Number(summary.totalPendingRewards) > 0) {
      console.log(`\n💡 建议: 用户有 ${ethers.utils.formatEther(summary.totalPendingRewards)} CPOP 待领取，可以调用 claimRewards 或 batchClaimRewards 领取`);
    }
    
    console.log("\n✅ 检查完成!");
    console.log("=".repeat(80));

  } catch (error: any) {
    console.error("\n❌ 错误:");
    console.error(error.message);
    if (error.stack) {
      console.error("\n堆栈跟踪:");
      console.error(error.stack);
    }
    process.exit(1);
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

