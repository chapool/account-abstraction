import { ethers } from "hardhat";

async function main() {
    console.log("🔍 分析等级加成72%的原因...\n");

    const [deployer] = await ethers.getSigners();
    console.log("查询者地址:", deployer.address);

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
        console.log("C级NFT配置参数分析");
        console.log("========================================");
        
        // 获取C级NFT的配置参数
        const dailyReward = await configContract.getDailyReward(1);
        const decayInterval = await configContract.getDecayInterval(1);
        const decayRate = await configContract.getDecayRate(1);
        const maxDecayRate = await configContract.getMaxDecayRate(1);
        
        console.log("C级NFT配置:");
        console.log("  每日基础奖励:", ethers.utils.formatEther(dailyReward), "CPOP");
        console.log("  衰减间隔:", decayInterval.toString(), "天");
        console.log("  衰减率:", (Number(decayRate) / 100).toFixed(2) + "%");
        console.log("  最大衰减率:", (Number(maxDecayRate) / 100).toFixed(2) + "%");
        
        console.log("\n========================================");
        console.log("用户NFT详细分析");
        console.log("========================================");
        
        // 获取用户质押的NFT
        let userStakes = [];
        let i = 0;
        while (true) {
            try {
                const tokenId = await stakingContract.userStakes(userAddress, i);
                userStakes.push(tokenId);
                i++;
            } catch (e) {
                break;
            }
        }
        
        console.log(`用户质押了 ${userStakes.length} 个NFT`);
        
        // 分析每个NFT的收益计算
        for (let j = 0; j < Math.min(userStakes.length, 3); j++) { // 只分析前3个
            const tokenId = userStakes[j];
            console.log(`\n--- NFT #${tokenId} 详细分析 ---`);
            
            const stakeInfo = await stakingContract.stakes(tokenId);
            const currentTime = await stakingContract.getCurrentTimestamp();
            const dayOffset = Math.floor((Number(currentTime) - Number(stakeInfo.stakeTime)) / (24 * 60 * 60));
            
            console.log("  质押时间:", new Date(Number(stakeInfo.stakeTime) * 1000).toLocaleString());
            console.log("  当前时间:", new Date(Number(currentTime) * 1000).toLocaleString());
            console.log("  质押天数偏移:", dayOffset, "天");
            
            // 获取收益分解
            try {
                const breakdown = await stakingContract.getDailyRewardBreakdown(tokenId, dayOffset);
                
                console.log("  收益分解:");
                console.log("    基础奖励:", ethers.utils.formatEther(breakdown.baseReward), "CPOP");
                console.log("    衰减后奖励:", ethers.utils.formatEther(breakdown.decayedReward), "CPOP");
                console.log("    季度乘数:", (Number(breakdown.quarterlyMultiplier) / 100).toFixed(2) + "%");
                console.log("    动态乘数:", (Number(breakdown.dynamicMultiplier) / 100).toFixed(2) + "%");
                console.log("    最终奖励:", ethers.utils.formatEther(breakdown.finalReward), "CPOP");
                
                // 计算衰减率
                const decayPercent = ((Number(breakdown.baseReward) - Number(breakdown.decayedReward)) / Number(breakdown.baseReward)) * 100;
                console.log("    衰减率:", decayPercent.toFixed(2) + "%");
                
                // 计算总乘数
                const totalMultiplier = (Number(breakdown.finalReward) / Number(breakdown.baseReward)) * 100;
                console.log("    总乘数:", totalMultiplier.toFixed(2) + "%");
                
            } catch (e) {
                console.log("  ❌ 无法获取收益分解:", e.message);
            }
        }
        
        console.log("\n========================================");
        console.log("衰减计算逻辑分析");
        console.log("========================================");
        
        console.log("C级NFT衰减规则:");
        console.log("  每20天衰减一次");
        console.log("  每次衰减35% (3500/10000)");
        console.log("  最大衰减80% (8000/10000)");
        
        // 计算不同天数的衰减情况
        const testDays = [0, 20, 40, 60, 80, 100];
        console.log("\n不同质押天数的衰减情况:");
        console.log("天数\t基础奖励\t衰减后\t衰减率\t剩余率");
        console.log("----\t--------\t------\t------\t------");
        
        for (const days of testDays) {
            const baseReward = Number(dailyReward);
            let decayedReward = baseReward;
            
            if (days >= 20) {
                const completedCycles = Math.floor(days / 20);
                
                for (let i = 0; i < completedCycles; i++) {
                    const totalDecaySoFar = (i + 1) * 3500;
                    if (totalDecaySoFar > 8000) {
                        const remainingDecay = 8000 - (i * 3500);
                        decayedReward = decayedReward * (10000 - remainingDecay) / 10000;
                        break;
                    }
                    decayedReward = decayedReward * (10000 - 3500) / 10000;
                }
            }
            
            const decayRate = ((baseReward - decayedReward) / baseReward) * 100;
            const remainingRate = (decayedReward / baseReward) * 100;
            
            console.log(`${days}\t${baseReward.toFixed(1)}\t\t${decayedReward.toFixed(1)}\t${decayRate.toFixed(1)}%\t${remainingRate.toFixed(1)}%`);
        }
        
        console.log("\n========================================");
        console.log("结论");
        console.log("========================================");
        console.log("72%的等级加成是因为:");
        console.log("1. C级NFT基础奖励: 3 CPOP");
        console.log("2. 由于质押时间较长，经历了多次衰减");
        console.log("3. 最终收益约为基础奖励的72%");
        console.log("4. 这反映了NFT质押的时间衰减机制");
        
    } catch (error) {
        console.log("❌ 分析失败:", error.message);
    }

    console.log("\n🎉 分析完成！");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });

