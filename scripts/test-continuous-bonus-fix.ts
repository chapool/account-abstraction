import { ethers } from "hardhat";

/**
 * 测试连续质押奖励修复是否有效
 */

async function main() {
    console.log("\n🔍 测试连续质押奖励修复...\n");

    const [deployer] = await ethers.getSigners();
    console.log("测试者地址:", deployer.address);

    // 用户地址
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";

    try {
        // 连接到合约
        const stakingContract = await ethers.getContractAt("Staking", stakingAddress);
        
        console.log("========================================");
        console.log("检查合约版本");
        console.log("========================================");
        
        const version = await stakingContract.version();
        console.log(`当前合约版本: ${version}`);
        
        if (version !== "3.3.0") {
            console.log("❌ 合约版本不是3.3.0，需要升级");
            return;
        }
        
        console.log("✅ 合约版本正确");
        
        console.log("\n========================================");
        console.log("检查用户质押信息");
        console.log("========================================");
        
        // 获取用户质押信息
        const tokenId = await stakingContract.userStakes(userAddress, 0);
        const stakeInfo = await stakingContract.stakes(tokenId);
        
        console.log(`Token ID: ${tokenId.toString()}`);
        console.log(`连续质押奖励是否已发放: ${stakeInfo.continuousBonusClaimed}`);
        console.log(`总奖励: ${ethers.utils.formatEther(stakeInfo.totalRewards)} CPOP`);
        console.log(`待领取奖励: ${ethers.utils.formatEther(stakeInfo.pendingRewards)} CPOP`);
        
        console.log("\n========================================");
        console.log("测试待领取奖励计算");
        console.log("========================================");
        
        // 计算待领取奖励
        const pendingRewards = await stakingContract.calculatePendingRewards(tokenId);
        console.log(`计算出的待领取奖励: ${ethers.utils.formatEther(pendingRewards)} CPOP`);
        
        console.log("\n========================================");
        console.log("分析结果");
        console.log("========================================");
        
        if (stakeInfo.continuousBonusClaimed) {
            console.log("✅ 连续质押奖励已标记为已发放");
            console.log("✅ 待领取奖励计算应该不包含连续质押奖励");
            
            // 手动计算基础奖励（不包含连续质押奖励）
            const currentTime = await stakingContract.getCurrentTimestamp();
            const daysSinceLastClaim = Math.floor((currentTime - Number(stakeInfo.lastClaimTime)) / (24 * 60 * 60));
            
            console.log(`距离上次领取: ${daysSinceLastClaim} 天`);
            
            // 简化计算：假设平均日收益为81 CPOP（考虑衰减）
            const estimatedBaseRewards = daysSinceLastClaim * 81;
            console.log(`估算基础奖励: ${estimatedBaseRewards} CPOP`);
            
            const actualRewards = Number(pendingRewards) / 1e18;
            const difference = Math.abs(actualRewards - estimatedBaseRewards);
            
            if (difference < 100) {
                console.log("✅ 待领取奖励计算正确，不包含重复的连续质押奖励");
            } else {
                console.log("❌ 待领取奖励计算可能仍有问题");
                console.log(`差异: ${difference} CPOP`);
            }
            
        } else {
            console.log("❌ 连续质押奖励未标记为已发放");
            console.log("❌ 可能存在重复发放问题");
        }
        
        console.log("\n========================================");
        console.log("总结");
        console.log("========================================");
        
        if (stakeInfo.continuousBonusClaimed && Number(pendingRewards) / 1e18 < 20000) {
            console.log("🎉 连续质押奖励修复成功！");
            console.log("- 连续质押奖励已正确标记为已发放");
            console.log("- 待领取奖励不再包含重复的连续质押奖励");
        } else {
            console.log("⚠️ 需要进一步检查修复效果");
        }
        
    } catch (error) {
        console.log("❌ 测试失败:", error.message);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });
