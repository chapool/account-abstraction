import { ethers } from "hardhat";

/**
 * 精确验证待领取奖励计算
 */

async function main() {
    console.log("\n🔍 精确验证待领取奖励计算...\n");

    const [deployer] = await ethers.getSigners();
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const configAddress = "0x37196054B23Be5CB977AA3284A3A844a8fe77861";

    try {
        const stakingContract = await ethers.getContractAt("Staking", stakingAddress);
        const configContract = await ethers.getContractAt("StakingConfig", configAddress);
        
        // 获取Token ID
        const tokenId = await stakingContract.userStakes(userAddress, 0);
        console.log(`Token ID: ${tokenId.toString()}`);
        
        // 获取质押信息
        const stakeInfo = await stakingContract.stakes(tokenId);
        const currentTime = await stakingContract.getCurrentTimestamp();
        
        console.log("\n📊 质押信息:");
        console.log(`质押时间: ${new Date(Number(stakeInfo.stakeTime) * 1000).toLocaleString()}`);
        console.log(`最后领取: ${new Date(Number(stakeInfo.lastClaimTime) * 1000).toLocaleString()}`);
        console.log(`当前时间: ${new Date(currentTime * 1000).toLocaleString()}`);
        
        // 计算天数
        const daysSinceLastClaim = Math.floor((currentTime - Number(stakeInfo.lastClaimTime)) / (24 * 60 * 60));
        const totalStakingDays = Math.floor((currentTime - Number(stakeInfo.stakeTime)) / (24 * 60 * 60));
        
        console.log(`\n⏰ 时间计算:`);
        console.log(`距离上次领取: ${daysSinceLastClaim} 天`);
        console.log(`总质押天数: ${totalStakingDays} 天`);
        
        // 获取配置
        const dailyReward = await configContract.getDailyReward(stakeInfo.level);
        const decayInterval = await configContract.getDecayInterval(stakeInfo.level);
        const decayRate = await configContract.getDecayRate(stakeInfo.level);
        const maxDecayRate = await configContract.getMaxDecayRate(stakeInfo.level);
        
        console.log(`\n📋 配置参数:`);
        console.log(`每日基础奖励: ${ethers.utils.formatEther(dailyReward)} CPOP`);
        console.log(`衰减间隔: ${decayInterval} 天`);
        console.log(`衰减率: ${(Number(decayRate) / 100).toFixed(2)}%`);
        console.log(`最大衰减率: ${(Number(maxDecayRate) / 100).toFixed(2)}%`);
        
        // 手动计算每日奖励
        console.log(`\n🧮 逐日计算验证:`);
        let totalCalculatedRewards = 0;
        
        for (let day = 0; day < daysSinceLastClaim; day++) {
            const currentDayFromStake = totalStakingDays - (daysSinceLastClaim - 1 - day);
            const currentDayTimestamp = Number(stakeInfo.stakeTime) + (currentDayFromStake * 24 * 60 * 60);
            
            let dailyRewardAmount = Number(dailyReward);
            
            // 应用衰减
            if (decayInterval > 0 && currentDayFromStake > decayInterval) {
                const completedCycles = Math.floor((currentDayFromStake - 1) / Number(decayInterval));
                
                for (let i = 0; i < completedCycles; i++) {
                    const totalDecaySoFar = (i + 1) * Number(decayRate);
                    if (totalDecaySoFar > Number(maxDecayRate)) {
                        const remainingDecay = Number(maxDecayRate) - (i * Number(decayRate));
                        dailyRewardAmount = dailyRewardAmount * (10000 - remainingDecay) / 10000;
                        break;
                    }
                    dailyRewardAmount = dailyRewardAmount * (10000 - Number(decayRate)) / 10000;
                }
            }
            
            // 应用历史季度乘数
            try {
                const historicalQuarterlyMultiplier = await stakingContract._getHistoricalQuarterlyMultiplier(currentDayTimestamp);
                dailyRewardAmount = dailyRewardAmount * Number(historicalQuarterlyMultiplier) / 10000;
            } catch (e) {
                // 如果函数不存在，使用默认值
                dailyRewardAmount = dailyRewardAmount * 10000 / 10000;
            }
            
            // 应用历史动态乘数
            try {
                const dynamicMultiplier = await stakingContract._getHistoricalDynamicMultiplier(stakeInfo.level, currentDayTimestamp);
                dailyRewardAmount = dailyRewardAmount * Number(dynamicMultiplier) / 10000;
            } catch (e) {
                // 如果函数不存在，使用默认值
                dailyRewardAmount = dailyRewardAmount * 10000 / 10000;
            }
            
            totalCalculatedRewards += dailyRewardAmount;
            
            if (day < 5 || day >= daysSinceLastClaim - 5) {
                console.log(`第${currentDayFromStake}天: ${ethers.utils.formatEther(dailyRewardAmount.toString())} CPOP`);
            } else if (day === 5) {
                console.log(`... (省略中间${daysSinceLastClaim - 10}天) ...`);
            }
        }
        
        console.log(`\n📊 计算结果:`);
        console.log(`手动计算总奖励: ${(totalCalculatedRewards / 1e18).toFixed(2)} CPOP`);
        
        // 获取实际待领取奖励
        const actualPendingRewards = await stakingContract.calculatePendingRewards(tokenId);
        console.log(`合约计算总奖励: ${ethers.utils.formatEther(actualPendingRewards)} CPOP`);
        
        const difference = Number(actualPendingRewards) - totalCalculatedRewards;
        console.log(`差异: ${(difference / 1e18).toFixed(2)} CPOP`);
        
        if (Math.abs(difference) < 1000000000000000000) { // 1 CPOP的精度
            console.log("✅ 计算结果一致！");
        } else {
            console.log("❌ 计算结果有差异，可能原因:");
            console.log("1. 连续质押奖励");
            console.log("2. Combo加成");
            console.log("3. 其他乘数调整");
        }
        
    } catch (error) {
        console.log("❌ 验证失败:", error.message);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });
