import { ethers } from "hardhat";

async function main() {
    console.log("\n🔄 重置并调整时间戳\n");

    // 合约地址
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    
    // 连接到合约
    const staking = await ethers.getContractAt("Staking", stakingAddress);
    
    const [deployer] = await ethers.getSigners();
    console.log("操作者:", deployer.address);
    console.log();
    
    // 获取质押信息
    const tokenId = 4812;
    const stakeInfo = await staking.stakes(tokenId);
    const stakeTime = Number(stakeInfo.stakeTime);
    
    console.log("========================================");
    console.log("当前状态");
    console.log("========================================");
    
    const testMode = await staking.testMode();
    const currentTimestamp = await staking.testTimestamp();
    
    console.log("测试模式:", testMode ? "启用" : "禁用");
    console.log("当前时间戳:", currentTimestamp.toString());
    console.log("当前日期:", new Date(Number(currentTimestamp) * 1000).toISOString());
    console.log();
    
    console.log("质押时间:", new Date(stakeTime * 1000).toISOString());
    console.log();
    
    // 步骤 1: 禁用测试模式
    console.log("========================================");
    console.log("步骤 1: 禁用测试模式");
    console.log("========================================");
    
    if (testMode) {
        try {
            console.log("禁用测试模式...");
            const tx1 = await staking.disableTestMode();
            console.log("交易哈希:", tx1.hash);
            console.log("等待确认...");
            await tx1.wait();
            console.log("✅ 测试模式已禁用");
            console.log();
        } catch (e: any) {
            console.log("⚠️ 禁用失败:", e.message);
            console.log("继续执行...");
            console.log();
        }
    } else {
        console.log("✅ 测试模式已禁用");
        console.log();
    }
    
    // 步骤 2: 重新启用测试模式，设置为质押时间
    console.log("========================================");
    console.log("步骤 2: 启用测试模式");
    console.log("========================================");
    
    // 设置为质押时间
    const initialTimestamp = ethers.BigNumber.from(stakeTime);
    
    console.log("启用测试模式，初始时间戳:", initialTimestamp.toString());
    console.log("初始日期:", new Date(Number(initialTimestamp) * 1000).toISOString());
    
    try {
        const tx2 = await staking.enableTestMode(initialTimestamp);
        console.log("交易哈希:", tx2.hash);
        console.log("等待确认...");
        await tx2.wait();
        console.log("✅ 测试模式已启用");
        console.log();
    } catch (e: any) {
        console.log("❌ 启用失败:", e.message);
        return;
    }
    
    // 验证
    const newTestMode = await staking.testMode();
    const newTimestamp = await staking.testTimestamp();
    console.log("验证:");
    console.log("  测试模式:", newTestMode ? "启用" : "禁用");
    console.log("  时间戳:", newTimestamp.toString());
    console.log("  日期:", new Date(Number(newTimestamp) * 1000).toISOString());
    console.log();
    
    // 步骤 3: 快进 90 天
    console.log("========================================");
    console.log("步骤 3: 快进 90 天");
    console.log("========================================");
    
    const targetDays = 90;
    console.log(`快进 ${targetDays} 天...`);
    
    try {
        const tx3 = await staking.fastForwardDays(ethers.BigNumber.from(targetDays));
        console.log("交易哈希:", tx3.hash);
        console.log("等待确认...");
        await tx3.wait();
        console.log("✅ 快进成功");
        console.log();
        
        // 验证新时间
        const finalTimestamp = await staking.testTimestamp();
        const finalDate = new Date(Number(finalTimestamp) * 1000);
        console.log("最终时间戳:", finalTimestamp.toString());
        console.log("最终日期:", finalDate.toISOString());
        
        // 计算天数
        const days = finalTimestamp.sub(stakeInfo.lastClaimTime).div(86400);
        console.log("需要计算的天数:", days.toString());
        console.log();
        
        // 测试计算奖励
        console.log("========================================");
        console.log("测试计算奖励");
        console.log("========================================");
        
        try {
            const estimateGas = await staking.estimateGas.calculatePendingRewards(tokenId);
            console.log("✅ 可以计算奖励");
            console.log("估算 Gas:", estimateGas.toString());
            console.log("   Gas 限制: 30,000,000");
            console.log("   使用率:", ((Number(estimateGas) / 30000000) * 100).toFixed(2) + "%");
            console.log();
            
            const pendingRewards = await staking.calculatePendingRewards(tokenId);
            console.log("待领取奖励:", ethers.utils.formatEther(pendingRewards), "CPOP");
            console.log();
            
            if (estimateGas.lt(20000000)) {
                console.log("✅ Gas 消耗在安全范围内");
                console.log("✅ 现在可以领取奖励了！");
            }
            
        } catch (e: any) {
            console.log("⚠️ 仍无法计算:", e.message);
        }
        
        console.log("\n========================================");
        console.log("总结");
        console.log("========================================");
        console.log(`
✅ 时间戳已重置并快进到 ${targetDays} 天
💡 现在可以尝试领取约 ${targetDays} 天的奖励

注意:
- 质押时间: ${new Date(stakeTime * 1000).toISOString()}
- 当前时间: ${finalDate.toISOString()}
- 计算天数: ${days.toString()}
- 可以领取 ${targetDays} 天的奖励

下一步:
1. 使用正确的账户（0xDf3715f4693CC308c961AaF0AacD56400E229F43）尝试领取
2. 如果成功，可以再快进更多天数继续领取
3. 重复此过程直到领取完所有奖励
    `);
        
    } catch (e: any) {
        console.log("❌ 快进失败:", e.message);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });

