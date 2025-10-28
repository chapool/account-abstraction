import { ethers } from "hardhat";

async function main() {
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    const tokenIds = ['4812', '3416', '3393', '3378', '3372', '3366'];
    
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const staking = await ethers.getContractAt("Staking", stakingAddress);
    
    console.log("=== 检查质押时间 ===");
    console.log("当前时间:", new Date().toLocaleString());
    
    // 获取当前区块时间
    const latestBlock = await ethers.provider.getBlock("latest");
    const currentTimestamp = latestBlock?.timestamp;
    console.log("链上当前时间:", currentTimestamp ? new Date(Number(currentTimestamp) * 1000).toLocaleString() : "无法获取");
    
    console.log("\n=== 检查每个NFT的质押时间和计算 ===");
    for (const tokenIdStr of tokenIds) {
        const tokenId = BigInt(tokenIdStr);
        const stake = await staking.stakes(tokenId);
        
        console.log(`\nToken ${tokenId}:`);
        console.log(`  质押时间: ${new Date(Number(stake.stakeTime) * 1000).toLocaleString()} (${stake.stakeTime})`);
        console.log(`  最后领取时间: ${new Date(Number(stake.lastClaimTime) * 1000).toLocaleString()} (${stake.lastClaimTime})`);
        
        if (currentTimestamp) {
            const stakeDuration = Number(currentTimestamp) - Number(stake.stakeTime);
            const stakeDays = stakeDuration / 86400;
            console.log(`  质押时长: ${stakeDays.toFixed(2)} 天 (${stakeDuration} 秒)`);
            
            const lastClaimDuration = Number(currentTimestamp) - Number(stake.lastClaimTime);
            const lastClaimDays = lastClaimDuration / 86400;
            console.log(`  距离最后领取: ${lastClaimDays.toFixed(2)} 天 (${lastClaimDays * 86400} 秒)`);
            
            // 检查时间是否异常
            if (stake.lastClaimTime > stake.stakeTime) {
                const claimBeforeStake = Number(stake.lastClaimTime) - Number(stake.stakeTime);
                if (claimBeforeStake > 86400) {
                    console.log(`  ⚠️ 最后领取时间晚于质押时间`);
                }
            }
            
            if (stake.stakeTime > currentTimestamp) {
                console.log(`  ❌ 质押时间在未来! (差异: ${Number(stake.stakeTime) - Number(currentTimestamp)} 秒)`);
            }
            if (stake.lastClaimTime > currentTimestamp) {
                console.log(`  ❌ 最后领取时间在未来! (差异: ${Number(stake.lastClaimTime) - Number(currentTimestamp)} 秒)`);
            }
        }
        
        // 尝试计算奖励
        try {
            const rewards = await staking.calculateTotalRewards(tokenId);
            console.log(`  总奖励: ${ethers.formatUnits(rewards, 18)} CPOP`);
        } catch (error: any) {
            console.log(`  ❌ 计算奖励失败:`, error.message);
        }
    }
    
    // 检查测试模式
    console.log("\n=== 检查测试模式 ===");
    try {
        const testMode = await staking.testMode();
        const testTimestamp = await staking.testTimestamp();
        console.log("测试模式:", testMode);
        if (testMode) {
            console.log("测试时间戳:", testTimestamp.toString());
            console.log("测试时间:", new Date(Number(testTimestamp) * 1000).toLocaleString());
        }
    } catch (error: any) {
        console.log("无法获取测试模式:", error.message);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });

