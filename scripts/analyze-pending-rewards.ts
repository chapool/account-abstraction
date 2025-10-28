import { ethers } from "hardhat";

/**
 * 分析用户待领取奖励的计算过程
 * 用户地址: 0xDf3715f4693CC308c961AaF0AacD56400E229F43
 */

async function main() {
    console.log("\n🔍 分析用户待领取奖励计算过程...\n");

    const [deployer] = await ethers.getSigners();
    console.log("分析者地址:", deployer.address);

    // 用户地址
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    console.log("目标用户地址:", userAddress);
    console.log();

    // 合约地址
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const configAddress = "0x37196054B23Be5CB977AA3284A3A844a8fe77861";

    try {
        // 连接到合约
        const stakingContract = await ethers.getContractAt("Staking", stakingAddress);
        const configContract = await ethers.getContractAt("StakingConfig", configAddress);
        
        console.log("========================================");
        console.log("获取用户质押信息");
        console.log("========================================");
        
        // 获取用户各等级的 NFT 数量
        let userTokenId = null;
        for (let level = 1; level <= 6; level++) {
            const count = await stakingContract.userLevelCounts(userAddress, level);
            if (count.gt(0)) {
                console.log(`Level ${level}: ${count.toString()} 个NFT`);
                
                // 获取第一个NFT的Token ID（简化处理）
                try {
                    const tokenId = await stakingContract.userStakes(userAddress, 0);
                    userTokenId = tokenId;
                    break;
                } catch (e) {
                    console.log("无法获取Token ID:", e.message);
                }
            }
        }
        
        if (!userTokenId) {
            console.log("❌ 无法获取用户的NFT Token ID");
            return;
        }
        
        console.log(`\n分析的NFT Token ID: ${userTokenId.toString()}`);
        
        // 获取质押信息
        const stakeInfo = await stakingContract.stakes(userTokenId);
        console.log("\n📊 质押信息:");
        console.log(`  所有者: ${stakeInfo.owner}`);
        console.log(`  Token ID: ${stakeInfo.tokenId.toString()}`);
        console.log(`  等级: ${stakeInfo.level}`);
        console.log(`  质押时间: ${new Date(Number(stakeInfo.stakeTime) * 1000).toLocaleString()}`);
        console.log(`  最后领取时间: ${new Date(Number(stakeInfo.lastClaimTime) * 1000).toLocaleString()}`);
        console.log(`  是否活跃: ${stakeInfo.isActive}`);
        console.log(`  总奖励: ${ethers.utils.formatEther(stakeInfo.totalRewards)} CPOP`);
        console.log(`  待领取奖励: ${ethers.utils.formatEther(stakeInfo.pendingRewards)} CPOP`);
        
        // 计算时间差
        const currentTime = await stakingContract.getCurrentTimestamp();
        const stakeTime = Number(stakeInfo.stakeTime);
        const lastClaimTime = Number(stakeInfo.lastClaimTime);
        
        const totalStakingDays = Math.floor((currentTime - stakeTime) / (24 * 60 * 60));
        const daysSinceLastClaim = Math.floor((currentTime - lastClaimTime) / (24 * 60 * 60));
        
        console.log(`\n⏰ 时间分析:`);
        console.log(`  当前时间戳: ${currentTime}`);
        console.log(`  总质押天数: ${totalStakingDays} 天`);
        console.log(`  距离上次领取: ${daysSinceLastClaim} 天`);
        
        console.log("\n========================================");
        console.log("配置参数分析");
        console.log("========================================");
        
        // 获取配置参数
        const dailyReward = await configContract.getDailyReward(stakeInfo.level);
        const decayInterval = await configContract.getDecayInterval(stakeInfo.level);
        const decayRate = await configContract.getDecayRate(stakeInfo.level);
        const maxDecayRate = await configContract.getMaxDecayRate(stakeInfo.level);
        
        console.log(`📋 Level ${stakeInfo.level} 配置:`);
        console.log(`  每日基础奖励: ${ethers.utils.formatEther(dailyReward)} CPOP`);
        console.log(`  衰减间隔: ${decayInterval.toString()} 天`);
        console.log(`  衰减率: ${(Number(decayRate) / 100).toFixed(2)}%`);
        console.log(`  最大衰减率: ${(Number(maxDecayRate) / 100).toFixed(2)}%`);
        
        console.log("\n========================================");
        console.log("奖励计算分析");
        console.log("========================================");
        
        // 计算待领取奖励
        const pendingRewards = await stakingContract.calculatePendingRewards(userTokenId);
        console.log(`💰 待领取奖励: ${ethers.utils.formatEther(pendingRewards)} CPOP`);
        
        // 获取每日奖励分解
        try {
            const dayOffset = Math.floor((currentTime - stakeTime) / (24 * 60 * 60));
            const breakdown = await stakingContract.getDailyRewardBreakdown(userTokenId, dayOffset);
            
            console.log(`\n📈 当前日奖励分解 (第${dayOffset}天):`);
            console.log(`  基础奖励: ${ethers.utils.formatEther(breakdown.baseReward)} CPOP`);
            console.log(`  衰减后奖励: ${ethers.utils.formatEther(breakdown.decayedReward)} CPOP`);
            console.log(`  季度乘数: ${(Number(breakdown.quarterlyMultiplier) / 100).toFixed(2)}%`);
            console.log(`  动态乘数: ${(Number(breakdown.dynamicMultiplier) / 100).toFixed(2)}%`);
            console.log(`  最终奖励: ${ethers.utils.formatEther(breakdown.finalReward)} CPOP`);
            
        } catch (e) {
            console.log("无法获取每日奖励分解:", e.message);
        }
        
        // 获取Combo加成
        const comboBonus = await stakingContract.getEffectiveComboBonus(userAddress, stakeInfo.level);
        console.log(`\n🎯 Combo加成: ${(Number(comboBonus) / 100).toFixed(2)}%`);
        
        // 计算连续质押奖励
        try {
            const continuousBonus = await stakingContract.calculateContinuousBonus(userTokenId, totalStakingDays);
            console.log(`🏆 连续质押奖励: ${ethers.utils.formatEther(continuousBonus)} CPOP`);
        } catch (e) {
            console.log("无法计算连续质押奖励:", e.message);
        }
        
        console.log("\n========================================");
        console.log("计算逻辑分析");
        console.log("========================================");
        
        console.log("📝 待领取奖励计算公式:");
        console.log("1. 基础计算: 从最后领取时间到现在的天数");
        console.log(`   - 天数: ${daysSinceLastClaim} 天`);
        console.log(`   - 基础日奖励: ${ethers.utils.formatEther(dailyReward)} CPOP`);
        
        console.log("\n2. 衰减计算:");
        if (decayInterval > 0) {
            console.log(`   - 衰减间隔: ${decayInterval} 天`);
            console.log(`   - 衰减率: ${(Number(decayRate) / 100).toFixed(2)}%`);
            console.log(`   - 最大衰减: ${(Number(maxDecayRate) / 100).toFixed(2)}%`);
            
            const completedCycles = Math.floor((totalStakingDays - 1) / Number(decayInterval));
            console.log(`   - 已完成周期: ${completedCycles} 个`);
            
            if (completedCycles > 0) {
                let decayedReward = Number(dailyReward);
                for (let i = 0; i < completedCycles; i++) {
                    const totalDecaySoFar = (i + 1) * Number(decayRate);
                    if (totalDecaySoFar > Number(maxDecayRate)) {
                        const remainingDecay = Number(maxDecayRate) - (i * Number(decayRate));
                        decayedReward = decayedReward * (10000 - remainingDecay) / 10000;
                        break;
                    }
                    decayedReward = decayedReward * (10000 - Number(decayRate)) / 10000;
                }
                console.log(`   - 衰减后日奖励: ${ethers.utils.formatEther(decayedReward.toString())} CPOP`);
            }
        } else {
            console.log("   - 无衰减");
        }
        
        console.log("\n3. 乘数应用:");
        console.log("   - 季度乘数: 根据历史季度调整");
        console.log("   - 动态乘数: 根据等级和历史时间调整");
        console.log(`   - Combo加成: ${(Number(comboBonus) / 100).toFixed(2)}%`);
        
        console.log("\n4. 连续质押奖励:");
        console.log("   - 基于质押天数的额外奖励");
        
        console.log("\n========================================");
        console.log("总结");
        console.log("========================================");
        
        const estimatedDailyReward = Number(pendingRewards) / daysSinceLastClaim;
        console.log(`📊 分析结果:`);
        console.log(`  待领取奖励: ${ethers.utils.formatEther(pendingRewards)} CPOP`);
        console.log(`  平均日收益: ${ethers.utils.formatEther(estimatedDailyReward.toString())} CPOP/天`);
        console.log(`  计算天数: ${daysSinceLastClaim} 天`);
        console.log(`  等级: Level ${stakeInfo.level} (SSS级)`);
        console.log(`  Combo加成: ${(Number(comboBonus) / 100).toFixed(2)}%`);
        
    } catch (error) {
        console.log("❌ 分析失败:", error.message);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });
