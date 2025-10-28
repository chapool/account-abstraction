import { ethers } from "hardhat";

/**
 * 分析连续质押奖励计算失败的原因
 * 并手动计算连续质押奖励
 */

async function main() {
    console.log("\n🔍 分析连续质押奖励计算...\n");

    const [deployer] = await ethers.getSigners();
    console.log("分析者地址:", deployer.address);

    // 用户地址
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const configAddress = "0x37196054B23Be5CB977AA3284A3A844a8fe77861";

    try {
        // 连接到合约
        const stakingContract = await ethers.getContractAt("Staking", stakingAddress);
        const configContract = await ethers.getContractAt("StakingConfig", configAddress);
        
        console.log("========================================");
        console.log("检查函数可见性");
        console.log("========================================");
        
        // 尝试调用 internal 函数（应该失败）
        try {
            await stakingContract._calculateContinuousBonus(3335, 361);
            console.log("❌ 意外成功：internal 函数被外部调用");
        } catch (e) {
            console.log("✅ 预期失败：_calculateContinuousBonus 是 internal 函数");
            console.log(`   错误: ${e.message}`);
        }
        
        // 尝试调用不存在的 external 函数（应该失败）
        try {
            await stakingContract.calculateContinuousBonus(3335, 361);
            console.log("❌ 意外成功：calculateContinuousBonus 函数存在");
        } catch (e) {
            console.log("✅ 预期失败：calculateContinuousBonus 函数不存在");
            console.log(`   错误: ${e.message}`);
        }
        
        console.log("\n========================================");
        console.log("连续质押奖励配置");
        console.log("========================================");
        
        // 获取连续质押奖励配置
        const thresholds = await configContract.getContinuousThresholds();
        const bonuses = await configContract.getContinuousBonuses();
        
        console.log("📋 连续质押奖励配置:");
        console.log(`  阈值: [${thresholds[0]}, ${thresholds[1]}] 天`);
        console.log(`  奖励: [${(Number(bonuses[0]) / 100).toFixed(1)}%, ${(Number(bonuses[1]) / 100).toFixed(1)}%]`);
        
        console.log("\n========================================");
        console.log("手动计算连续质押奖励");
        console.log("========================================");
        
        // 获取用户质押信息
        const tokenId = await stakingContract.userStakes(userAddress, 0);
        const stakeInfo = await stakingContract.stakes(tokenId);
        const currentTime = await stakingContract.getCurrentTimestamp();
        const totalStakingDays = Math.floor((currentTime - Number(stakeInfo.stakeTime)) / (24 * 60 * 60));
        
        console.log(`Token ID: ${tokenId.toString()}`);
        console.log(`总质押天数: ${totalStakingDays} 天`);
        
        // 手动计算连续质押奖励
        let applicableThreshold = 0;
        let applicableBonus = 0;
        
        // 找到适用的阈值和奖励
        for (let i = thresholds.length; i > 0; i--) {
            if (totalStakingDays >= Number(thresholds[i - 1])) {
                applicableThreshold = Number(thresholds[i - 1]);
                applicableBonus = Number(bonuses[i - 1]);
                break;
            }
        }
        
        console.log(`\n🎯 适用的连续质押奖励:`);
        console.log(`  阈值: ${applicableThreshold} 天`);
        console.log(`  奖励率: ${(applicableBonus / 100).toFixed(1)}%`);
        
        if (applicableBonus === 0) {
            console.log("❌ 无连续质押奖励");
            return;
        }
        
        // 计算阈值期间的奖励
        console.log(`\n🧮 计算阈值期间奖励 (前${applicableThreshold}天):`);
        
        const baseReward = await configContract.getDailyReward(stakeInfo.level);
        const decayInterval = await configContract.getDecayInterval(stakeInfo.level);
        const decayRate = await configContract.getDecayRate(stakeInfo.level);
        const maxDecayRate = await configContract.getMaxDecayRate(stakeInfo.level);
        
        console.log(`基础日奖励: ${ethers.utils.formatEther(baseReward)} CPOP`);
        console.log(`衰减间隔: ${decayInterval} 天`);
        console.log(`衰减率: ${(Number(decayRate) / 100).toFixed(2)}%`);
        
        let rewardsAtThreshold = 0;
        
        // 计算前 applicableThreshold 天的奖励
        for (let day = 0; day < applicableThreshold; day++) {
            const currentDayTimestamp = Number(stakeInfo.stakeTime) + (day * 24 * 60 * 60);
            let dailyReward = Number(baseReward);
            
            // 应用衰减
            if (decayInterval > 0 && day >= Number(decayInterval)) {
                const completedCycles = Math.floor(day / Number(decayInterval));
                for (let i = 0; i < completedCycles; i++) {
                    const totalDecaySoFar = (i + 1) * Number(decayRate);
                    if (totalDecaySoFar > Number(maxDecayRate)) {
                        const remainingDecay = Number(maxDecayRate) - (i * Number(decayRate));
                        dailyReward = dailyReward * (10000 - remainingDecay) / 10000;
                        break;
                    }
                    dailyReward = dailyReward * (10000 - Number(decayRate)) / 10000;
                }
            }
            
            // 应用历史季度乘数
            try {
                const historicalQuarterlyMultiplier = await stakingContract._getHistoricalQuarterlyMultiplier(currentDayTimestamp);
                dailyReward = dailyReward * Number(historicalQuarterlyMultiplier) / 10000;
            } catch (e) {
                // 使用默认值
                dailyReward = dailyReward * 10000 / 10000;
            }
            
            // 应用历史动态乘数
            try {
                const dynamicMultiplier = await stakingContract._getHistoricalDynamicMultiplier(stakeInfo.level, currentDayTimestamp);
                dailyReward = dailyReward * Number(dynamicMultiplier) / 10000;
            } catch (e) {
                // 使用默认值
                dailyReward = dailyReward * 10000 / 10000;
            }
            
            rewardsAtThreshold += dailyReward;
            
            if (day < 5 || day >= applicableThreshold - 5) {
                console.log(`第${day + 1}天: ${ethers.utils.formatEther(dailyReward.toString())} CPOP`);
            } else if (day === 5) {
                console.log(`... (省略中间${applicableThreshold - 10}天) ...`);
            }
        }
        
        console.log(`\n📊 阈值期间总奖励: ${(rewardsAtThreshold / 1e18).toFixed(2)} CPOP`);
        
        // 应用Combo加成
        const comboBonus = await stakingContract.getEffectiveComboBonus(userAddress, stakeInfo.level);
        rewardsAtThreshold = rewardsAtThreshold * (10000 + Number(comboBonus)) / 10000;
        
        console.log(`Combo加成后: ${(rewardsAtThreshold / 1e18).toFixed(2)} CPOP`);
        console.log(`Combo加成: ${(Number(comboBonus) / 100).toFixed(2)}%`);
        
        // 计算最终连续质押奖励
        const continuousBonus = rewardsAtThreshold * applicableBonus / 10000;
        
        console.log(`\n🎉 连续质押奖励计算结果:`);
        console.log(`  阈值期间奖励: ${(rewardsAtThreshold / 1e18).toFixed(2)} CPOP`);
        console.log(`  奖励率: ${(applicableBonus / 100).toFixed(1)}%`);
        console.log(`  连续质押奖励: ${(continuousBonus / 1e18).toFixed(2)} CPOP`);
        
        console.log("\n========================================");
        console.log("总结");
        console.log("========================================");
        
        console.log("🔍 函数调用失败的原因:");
        console.log("1. _calculateContinuousBonus 是 internal 函数，无法外部调用");
        console.log("2. calculateContinuousBonus 函数不存在");
        console.log("3. 合约设计：连续质押奖励只在内部计算中使用");
        
        console.log("\n💰 连续质押奖励分析:");
        console.log(`- 用户质押 ${totalStakingDays} 天`);
        console.log(`- 触发 ${applicableThreshold} 天阈值`);
        console.log(`- 获得 ${(applicableBonus / 100).toFixed(1)}% 奖励`);
        console.log(`- 奖励金额: ${(continuousBonus / 1e18).toFixed(2)} CPOP`);
        
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
