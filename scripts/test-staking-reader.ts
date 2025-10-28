import { ethers } from "hardhat";

/**
 * 测试 StakingReader 合约的两个关键函数
 * 1. getUserStakingSummary - 获取用户质押概览
 * 2. getUserDailyRewards - 获取用户每日收益详情
 */

async function main() {
    console.log("\n🔍 测试 StakingReader 合约函数...\n");

    const [deployer] = await ethers.getSigners();
    console.log("测试者地址:", deployer.address);

    // 用户地址
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    console.log("目标用户地址:", userAddress);
    console.log();

    // StakingReader 合约地址
    const stakingReaderAddress = "0x3a42F2F6962Fdc9bb8392F73C94e79071De5eC8d";
    
    try {
        // 连接到 StakingReader 合约
        const stakingReaderContract = await ethers.getContractAt("StakingReader", stakingReaderAddress);
        
        console.log("========================================");
        console.log("测试 getUserStakingSummary 函数");
        console.log("========================================");
        
        try {
            const summary = await stakingReaderContract.getUserStakingSummary(userAddress);
            
            console.log("📊 用户质押概览:");
            console.log(`  总质押数量: ${summary.totalStakedCount}`);
            console.log(`  总已领取奖励: ${ethers.utils.formatEther(summary.totalClaimedRewards)} CPOP`);
            console.log(`  总待领取奖励: ${ethers.utils.formatEther(summary.totalPendingRewards)} CPOP`);
            console.log(`  最长质押时长: ${Math.floor(summary.longestStakingDuration / (24 * 60 * 60))} 天`);
            console.log(`  总有效乘数: ${(Number(summary.totalEffectiveMultiplier) / 100).toFixed(2)}%`);
            
            console.log("\n📈 各等级质押数量:");
            const levelNames = ["C级", "B级", "A级", "S级", "SS级", "SSS级"];
            for (let i = 0; i < 6; i++) {
                if (summary.levelStakingCounts[i] > 0) {
                    console.log(`  ${levelNames[i]}: ${summary.levelStakingCounts[i]} 个`);
                }
            }
            
        } catch (e) {
            console.log("❌ getUserStakingSummary 调用失败:", e.message);
        }
        
        console.log("\n========================================");
        console.log("测试 getUserDailyRewards 函数");
        console.log("========================================");
        
        try {
            const dailyRewards = await stakingReaderContract.getUserDailyRewards(userAddress);
            
            console.log("💰 每日收益详情:");
            console.log(`  总基础收益: ${ethers.utils.formatEther(dailyRewards.totalBaseReward)} CPOP`);
            console.log(`  总衰减后收益: ${ethers.utils.formatEther(dailyRewards.totalDecayedReward)} CPOP`);
            console.log(`  总Combo加成: ${(Number(dailyRewards.totalComboBonus) / 100).toFixed(2)}%`);
            console.log(`  总动态乘数: ${(Number(dailyRewards.totalDynamicMultiplier) / 100).toFixed(2)}%`);
            console.log(`  总最终收益: ${ethers.utils.formatEther(dailyRewards.totalFinalReward)} CPOP`);
            console.log(`  总乘数: ${(Number(dailyRewards.totalBonus) / 100).toFixed(2)}%`);
            
            console.log("\n📊 各等级收益分布:");
            const levelNames = ["C级", "B级", "A级", "S级", "SS级", "SSS级"];
            for (let i = 0; i < 6; i++) {
                if (dailyRewards.nftCountPerLevel[i] > 0) {
                    console.log(`  ${levelNames[i]}:`);
                    console.log(`    NFT数量: ${dailyRewards.nftCountPerLevel[i]}`);
                    console.log(`    基础收益: ${ethers.utils.formatEther(dailyRewards.baseRewardPerLevel[i])} CPOP`);
                    console.log(`    最终收益: ${ethers.utils.formatEther(dailyRewards.finalRewardPerLevel[i])} CPOP`);
                }
            }
            
        } catch (e) {
            console.log("❌ getUserDailyRewards 调用失败:", e.message);
        }
        
    } catch (error) {
        console.log("❌ 合约连接失败:", error.message);
        console.log("请确保 StakingReader 合约地址正确");
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });
