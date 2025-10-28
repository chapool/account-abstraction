import { ethers } from "hardhat";

async function main() {
    console.log("\n🔍 分析 90 天限制对奖励计算的影响\n");

    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const staking = await ethers.getContractAt("Staking", stakingAddress);
    
    console.log("========================================");
    console.log("测试场景：质押 180 天未领取");
    console.log("========================================");
    
    // 获取质押信息
    const tokenId = 4812;
    const stakeInfo = await staking.stakes(tokenId);
    
    console.log("\n当前状态:");
    console.log("质押时间:", new Date(Number(stakeInfo.stakeTime) * 1000).toISOString());
    console.log("最后领取时间:", new Date(Number(stakeInfo.lastClaimTime) * 1000).toISOString());
    
    const currentTimestamp = await staking.getCurrentTimestamp();
    const totalDays = currentTimestamp.sub(stakeInfo.lastClaimTime).div(86400);
    
    console.log("当前时间:", new Date(Number(currentTimestamp) * 1000).toISOString());
    console.log("实际待领取天数:", totalDays.toString());
    console.log();
    
    // 测试计算奖励
    console.log("========================================");
    console.log("测试：calculatePendingRewards");
    console.log("========================================");
    
    try {
        const pendingRewards = await staking.calculatePendingRewards(tokenId);
        console.log("✅ 计算成功");
        console.log("待领取奖励:", ethers.utils.formatEther(pendingRewards), "CPOP");
        
        // 计算理论应该领取多少
        const expectedRewards = await calculateExpectedRewards(
            staking,
            tokenId,
            totalDays.toNumber()
        );
        
        const limitedRewards = await calculateExpectedRewards(
            staking,
            tokenId,
            90 // 限制为90天
        );
        
        console.log("\n对比分析:");
        console.log("理论总奖励（不受限）:", ethers.utils.formatEther(expectedRewards), "CPOP");
        console.log("实际返回奖励（受90天限制）:", ethers.utils.formatEther(pendingRewards), "CPOP");
        console.log("限制后应返回:", ethers.utils.formatEther(limitedRewards), "CPOP");
        
        const diff = expectedRewards.sub(pendingRewards);
        console.log("未领取的奖励:", ethers.utils.formatEther(diff), "CPOP");
        
        if (diff.gt(0)) {
            console.log("\n⚠️ 影响分析:");
            console.log("- 用户会看到:", ethers.utils.formatEther(pendingRewards), "CPOP");
            console.log("- 但实际应领取:", ethers.utils.formatEther(expectedRewards), "CPOP");
            console.log("- 差额:", ethers.utils.formatEther(diff), "CPOP 被限制");
            console.log("\n💡 这意味着:");
            console.log("1. 展示的奖励不完整（只有90天）");
            console.log("2. 第一次领取只能领到90天的奖励");
            console.log("3. 需要多次领取才能领完所有奖励");
        } else {
            console.log("\n✅ 没有影响（总天数 < 90 天）");
        }
        
    } catch (e: any) {
        console.log("❌ 计算失败:", e.message);
    }
    
    console.log("\n========================================");
    console.log("详细影响分析");
    console.log("========================================");
    
    // 模拟不同质押天数
    const testCases = [30, 90, 180, 360, 1434];
    
    for (const days of testCases) {
        console.log(`\n质押 ${days} 天:`);
        
        const theoreticalRewards = await calculateExpectedRewards(staking, tokenId, days);
        const limitedRewards = await calculateExpectedRewards(staking, tokenId, Math.min(days, 90));
        
        const percentage = (limitedRewards.mul(10000).div(theoreticalRewards)).toNumber() / 100;
        
        console.log(`  理论奖励: ${ethers.utils.formatEther(theoreticalRewards)} CPOP`);
        console.log(`  实际返回: ${ethers.utils.formatEther(limitedRewards)} CPOP`);
        console.log(`  显示比例: ${percentage}%`);
        
        if (days > 90) {
            const missed = theoreticalRewards.sub(limitedRewards);
            console.log(`  未显示: ${ethers.utils.formatEther(missed)} CPOP`);
            console.log(`  ⚠️ 需要分 ${Math.ceil(days / 90)} 次领取`);
        }
    }
    
    console.log("\n========================================");
    console.log("解决方案建议");
    console.log("========================================");
    console.log(`
问题：
1. ✅ 解决了 Gas 超限问题
2. ❌ 但会限制显示/领取的奖励数量

当前行为：
- calculatePendingRewards() 只返回最多 90 天的奖励
- 用户会看到不完整的奖励金额
- 需要多次领取才能领取完所有奖励

推荐解决方案：

方案 1: 分离显示和领取逻辑（推荐）
- calculatePendingRewards() 不加限制（用于准确显示）
- claimRewards() 内部使用限制（防止Gas超限）
- 在 claimRewards 中如果计算天数 > 90，只领取90天，更新 lastClaimTime

方案 2: 前端分批显示
- 前端检测到待领取天数 > 90 时
- 提示："您的奖励已累积 ${totalDays} 天，将分 ${Math.ceil(totalDays/90)} 次领取"
- 自动分批领取

方案 3: 优化计算逻辑
- 使用数学公式代替循环
- 无需限制天数
- 开发成本较高
    `);
}

// 辅助函数：模拟计算奖励（不使用合约限制）
async function calculateExpectedRewards(
    staking: any,
    tokenId: number,
    days: number
): Promise<ethers.BigNumber> {
    // 这里简化，实际需要从合约配置中获取
    const stakeInfo = await staking.stakes(tokenId);
    const baseReward = 100; // SSS级示例
    const level = stakeInfo.level;
    
    // 简单的线性计算（不考虑衰减和调整）
    // 实际应该和合约逻辑一致
    const rewardPerDay = baseReward * (1 - (level - 1) * 0.1);
    return ethers.utils.parseEther((rewardPerDay * days).toString());
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });

