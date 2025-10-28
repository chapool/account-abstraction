import { ethers } from "hardhat";

/**
 * 重新分析连续质押奖励的发放机制
 * 连续质押奖励只发放一次，不是持续发放
 */

async function main() {
    console.log("\n🔍 重新分析连续质押奖励发放机制...\n");

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
        console.log("用户质押时间线分析");
        console.log("========================================");
        
        // 获取用户质押信息
        const tokenId = await stakingContract.userStakes(userAddress, 0);
        const stakeInfo = await stakingContract.stakes(tokenId);
        const currentTime = await stakingContract.getCurrentTimestamp();
        
        const stakeTime = Number(stakeInfo.stakeTime);
        const lastClaimTime = Number(stakeInfo.lastClaimTime);
        const totalStakingDays = Math.floor((currentTime - stakeTime) / (24 * 60 * 60));
        const daysSinceLastClaim = Math.floor((currentTime - lastClaimTime) / (24 * 60 * 60));
        
        console.log(`Token ID: ${tokenId.toString()}`);
        console.log(`质押开始时间: ${new Date(stakeTime * 1000).toLocaleString()}`);
        console.log(`最后领取时间: ${new Date(lastClaimTime * 1000).toLocaleString()}`);
        console.log(`当前时间: ${new Date(currentTime * 1000).toLocaleString()}`);
        console.log(`总质押天数: ${totalStakingDays} 天`);
        console.log(`距离上次领取: ${daysSinceLastClaim} 天`);
        
        console.log("\n========================================");
        console.log("连续质押奖励发放分析");
        console.log("========================================");
        
        // 获取连续质押奖励配置
        const thresholds = await configContract.getContinuousThresholds();
        const bonuses = await configContract.getContinuousBonuses();
        
        console.log("📋 连续质押奖励配置:");
        console.log(`  阈值: [${thresholds[0]}, ${thresholds[1]}] 天`);
        console.log(`  奖励: [${(Number(bonuses[0]) / 100).toFixed(1)}%, ${(Number(bonuses[1]) / 100).toFixed(1)}%]`);
        
        // 分析连续质押奖励的发放时机
        console.log("\n🎯 连续质押奖励发放时机分析:");
        
        // 检查在质押开始到上次领取期间是否触发了连续质押奖励
        const daysFromStakeToLastClaim = Math.floor((lastClaimTime - stakeTime) / (24 * 60 * 60));
        console.log(`质押开始到上次领取: ${daysFromStakeToLastClaim} 天`);
        
        // 检查是否在第一次领取时已经获得了连续质押奖励
        let continuousBonusTriggered = false;
        let continuousBonusAmount = 0;
        
        for (let i = thresholds.length; i > 0; i--) {
            if (daysFromStakeToLastClaim >= Number(thresholds[i - 1])) {
                console.log(`✅ 在第${daysFromStakeToLastClaim}天时已触发${thresholds[i - 1]}天阈值`);
                console.log(`   奖励率: ${(Number(bonuses[i - 1]) / 100).toFixed(1)}%`);
                continuousBonusTriggered = true;
                
                // 计算连续质押奖励金额
                const baseReward = await configContract.getDailyReward(stakeInfo.level);
                const threshold = Number(thresholds[i - 1]);
                const bonus = Number(bonuses[i - 1]);
                
                // 简化计算：前threshold天的奖励
                let thresholdRewards = Number(baseReward) * threshold;
                
                // 应用Combo加成
                const comboBonus = await stakingContract.getEffectiveComboBonus(userAddress, stakeInfo.level);
                thresholdRewards = thresholdRewards * (10000 + Number(comboBonus)) / 10000;
                
                continuousBonusAmount = thresholdRewards * bonus / 10000;
                
                console.log(`   连续质押奖励金额: ${(continuousBonusAmount / 1e18).toFixed(2)} CPOP`);
                break;
            }
        }
        
        if (!continuousBonusTriggered) {
            console.log("❌ 在第一次领取时未触发连续质押奖励");
        }
        
        console.log("\n========================================");
        console.log("当前待领取奖励分析");
        console.log("========================================");
        
        // 计算当前待领取奖励（不包含连续质押奖励）
        console.log("🧮 当前待领取奖励计算 (不包含连续质押奖励):");
        
        const baseReward = await configContract.getDailyReward(stakeInfo.level);
        const decayInterval = await configContract.getDecayInterval(stakeInfo.level);
        const decayRate = await configContract.getDecayRate(stakeInfo.level);
        const maxDecayRate = await configContract.getMaxDecayRate(stakeInfo.level);
        
        console.log(`基础日奖励: ${ethers.utils.formatEther(baseReward)} CPOP`);
        console.log(`衰减间隔: ${decayInterval} 天`);
        console.log(`衰减率: ${(Number(decayRate) / 100).toFixed(2)}%`);
        
        let currentPendingRewards = 0;
        
        // 计算从上次领取到现在的奖励
        for (let day = 0; day < daysSinceLastClaim; day++) {
            const currentDayFromStake = totalStakingDays - (daysSinceLastClaim - 1 - day);
            const currentDayTimestamp = stakeTime + (currentDayFromStake * 24 * 60 * 60);
            
            let dailyReward = Number(baseReward);
            
            // 应用衰减
            if (decayInterval > 0 && currentDayFromStake > Number(decayInterval)) {
                const completedCycles = Math.floor((currentDayFromStake - 1) / Number(decayInterval));
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
                dailyReward = dailyReward * 10000 / 10000;
            }
            
            // 应用历史动态乘数
            try {
                const dynamicMultiplier = await stakingContract._getHistoricalDynamicMultiplier(stakeInfo.level, currentDayTimestamp);
                dailyReward = dailyReward * Number(dynamicMultiplier) / 10000;
            } catch (e) {
                dailyReward = dailyReward * 10000 / 10000;
            }
            
            currentPendingRewards += dailyReward;
            
            if (day < 3 || day >= daysSinceLastClaim - 3) {
                console.log(`第${currentDayFromStake}天: ${ethers.utils.formatEther(dailyReward.toString())} CPOP`);
            } else if (day === 3) {
                console.log(`... (省略中间${daysSinceLastClaim - 6}天) ...`);
            }
        }
        
        // 应用Combo加成
        const comboBonus = await stakingContract.getEffectiveComboBonus(userAddress, stakeInfo.level);
        currentPendingRewards = currentPendingRewards * (10000 + Number(comboBonus)) / 10000;
        
        console.log(`\n📊 当前待领取奖励计算:`);
        console.log(`  基础计算: ${(currentPendingRewards / 1e18).toFixed(2)} CPOP`);
        console.log(`  Combo加成: ${(Number(comboBonus) / 100).toFixed(2)}%`);
        console.log(`  最终奖励: ${(currentPendingRewards / 1e18).toFixed(2)} CPOP`);
        
        // 获取实际待领取奖励
        const actualPendingRewards = await stakingContract.calculatePendingRewards(tokenId);
        console.log(`  实际待领取: ${ethers.utils.formatEther(actualPendingRewards)} CPOP`);
        
        const difference = Number(actualPendingRewards) - currentPendingRewards;
        console.log(`  差异: ${(difference / 1e18).toFixed(2)} CPOP`);
        
        console.log("\n========================================");
        console.log("结论分析");
        console.log("========================================");
        
        if (continuousBonusTriggered) {
            console.log("✅ 连续质押奖励已在第一次领取时发放");
            console.log(`   发放金额: ${(continuousBonusAmount / 1e18).toFixed(2)} CPOP`);
            console.log("❌ 当前待领取奖励不应包含连续质押奖励");
        } else {
            console.log("❌ 连续质押奖励尚未发放");
            console.log("✅ 当前待领取奖励可能包含连续质押奖励");
        }
        
        console.log("\n🔍 17,991 CPOP 构成分析:");
        console.log(`- 基础衰减计算: ${(currentPendingRewards / 1e18).toFixed(2)} CPOP`);
        console.log(`- 额外奖励: ${(difference / 1e18).toFixed(2)} CPOP`);
        
        if (Math.abs(difference) > 100) {
            console.log("\n❓ 额外奖励的可能来源:");
            console.log("1. 连续质押奖励（如果尚未发放）");
            console.log("2. 其他奖励机制");
            console.log("3. 计算精度差异");
        } else {
            console.log("\n✅ 计算结果基本一致，差异在合理范围内");
        }
        
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
