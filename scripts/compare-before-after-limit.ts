import { ethers } from "hardhat";

async function main() {
    console.log("\n📊 对比：加限制前后的行为\n");

    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const staking = await ethers.getContractAt("Staking", stakingAddress);
    
    const tokenId = 4812;
    const stakeInfo = await staking.stakes(tokenId);
    
    const currentTimestamp = await staking.getCurrentTimestamp();
    const totalDays = currentTimestamp.sub(stakeInfo.lastClaimTime).div(86400);
    
    console.log("========================================");
    console.log("测试数据");
    console.log("========================================");
    console.log("NFT ID:", tokenId);
    console.log("实际待领取天数:", totalDays.toString());
    console.log();
    
    console.log("========================================");
    console.log("当前行为（加了90天限制后）");
    console.log("========================================");
    
    try {
        const pendingRewards = await staking.calculatePendingRewards(tokenId);
        console.log("calculatePendingRewards() 返回:");
        console.log("  ", ethers.utils.formatEther(pendingRewards), "CPOP");
        console.log("  ", "（最多90天的奖励）");
        console.log();
        
        // 模拟领取
        console.log("如果调用 claimRewards():");
        console.log("  将领取:", ethers.utils.formatEther(pendingRewards), "CPOP");
        console.log("  剩余:", totalDays.gt(90) ? `${totalDays.sub(90)} 天" : "0 天");
        console.log();
        
    } catch (e: any) {
        console.log("❌ 计算失败:", e.message);
    }
    
    console.log("========================================");
    console.log("影响分析");
    console.log("========================================");
    
    const scenarios = [
        { days: 30, label: "短期质押（30天）" },
        { days: 90, label: "中期质押（90天）" },
        { days: 180, label: "长期质押（180天）" },
        { days: 365, label: "超长期质押（365天）" },
    ];
    
    for (const scenario of scenarios) {
        console.log(`\n${scenario.label}:`);
        
        if (scenario.days <= 90) {
            console.log("  ✅ 显示完整奖励");
            console.log("  ✅ 一次领取完");
            console.log("  ✅ 无影响");
        } else {
            const batches = Math.ceil(scenario.days / 90);
            console.log(`  ⚠️  限制影响:`);
            console.log(`     - 只显示 90 天奖励（${Math.floor(scenario.days / 90) * 100}% 完整度）`);
            console.log(`     - 需要分 ${batches} 次领取`);
            console.log(`     - 用户可能感到困惑`);
        }
    }
    
    console.log("\n========================================");
    console.log("改进建议");
    console.log("========================================");
    
    console.log(`
当前方案的问题：
❌ calculatePendingRewards() 显示不完整
❌ 用户看到的金额不对
❌ 需要猜测实际应领取多少

建议改进方案：保持 calculatePendingRewards 完整，仅在领取时限制

方案 1: 修改 claimRewards 逻辑（推荐）
优点：
  ✅ calculatePendingRewards() 显示完整
  ✅ 用户体验好
  ✅ 自动分批领取
  
实现：
  - calculatePendingRewards: 不加限制，准确显示总奖励
  - claimRewards: 如果总天数 > 90，只领取90天，更新 lastClaimTime
  
方案 2: 添加新函数 claimRewardsPartial
优点：
  ✅ 保留原有逻辑不变
  ✅ 用户可以主动选择
  
实现：
  - claimRewards: 保持原样（可能会Gas超限）
  - claimRewardsPartial(days): 新增，指定领取天数
  - calculatePendingRewards: 不加限制

方案 3: 前端分批显示和领取
优点：
  ✅ 不需要修改合约
  
实现：
  - 前端检测 totalDays > 90
  - 显示: "总计 ${totalDays} 天，将分 ${batches} 次领取"
  - 自动分批调用 claimRewards
    `);
    
    console.log("\n========================================");
    console.log("推荐实施：方案 1");
    console.log("========================================");
    console.log(`
步骤：
1. 移除 calculatePendingRewards 中的限制
2. 在 claimRewards 中添加智能分批逻辑
3. 检测 totalDays > 90 时，只领取 90 天并更新 lastClaimTime
4. 提示用户需要多次领取
    `);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });

