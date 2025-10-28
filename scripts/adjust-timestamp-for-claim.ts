import { ethers } from "hardhat";

async function main() {
    console.log("\n🕐 调整时间戳以便领取奖励\n");

    // 合约地址
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    
    // 连接到合约
    const staking = await ethers.getContractAt("Staking", stakingAddress);
    
    const [deployer] = await ethers.getSigners();
    console.log("操作者:", deployer.address);
    console.log();
    
    // 检查当前状态
    console.log("========================================");
    console.log("当前状态");
    console.log("========================================");
    
    const testMode = await staking.testMode();
    const currentTimestamp = await staking.testTimestamp();
    
    if (!testMode) {
        console.log("❌ 测试模式未启用");
        console.log("💡 请先运行: npx hardhat run scripts/enable-test-mode.ts --network sepoliaCustom");
        return;
    }
    
    console.log("✅ 测试模式已启用");
    console.log("当前时间戳:", currentTimestamp.toString());
    console.log("当前日期:", new Date(Number(currentTimestamp) * 1000).toISOString());
    console.log();
    
    // 获取质押信息
    const tokenId = 4812;
    const stakeInfo = await staking.stakes(tokenId);
    const stakeTime = Number(stakeInfo.stakeTime);
    
    console.log("========================================");
    console.log("质押信息");
    console.log("========================================");
    console.log("质押时间:", new Date(stakeTime * 1000).toISOString());
    console.log("最后领取时间:", new Date(Number(stakeInfo.lastClaimTime) * 1000).toISOString());
    
    const currentDays = currentTimestamp.sub(stakeInfo.lastClaimTime).div(86400);
    console.log("距离最后领取天数:", currentDays.toString());
    console.log();
    
    // 计算建议的新时间戳
    // 建议设置为：质押时间 + 90 天（便于测试）
    const suggestedDays = 90;
    const suggestedTimestamp = ethers.BigNumber.from(stakeTime + suggestedDays * 86400);
    
    console.log("========================================");
    console.log("建议方案");
    console.log("========================================");
    console.log("当前需要计算:", currentDays.toString(), "天（导致 Gas 超限）");
    console.log("建议调整为:", suggestedDays, "天");
    console.log("新时间戳:", suggestedTimestamp.toString());
    console.log("新日期:", new Date(Number(suggestedTimestamp) * 1000).toISOString());
    console.log();
    
    console.log("========================================");
    console.log("执行调整");
    console.log("========================================");
    
    try {
        console.log("设置 testTimestamp 为:", suggestedTimestamp.toString());
        const tx = await staking.setTestTimestamp(suggestedTimestamp);
        console.log("交易哈希:", tx.hash);
        console.log("等待确认...");
        const receipt = await tx.wait();
        console.log("✅ 成功！");
        console.log("区块号:", receipt.blockNumber);
        console.log();
        
        // 验证新时间
        const newTimestamp = await staking.testTimestamp();
        const newDate = new Date(Number(newTimestamp) * 1000);
        console.log("新的时间戳:", newTimestamp.toString());
        console.log("新的日期:", newDate.toISOString());
        console.log();
        
        // 计算现在的天数
        const newDays = newTimestamp.sub(stakeInfo.lastClaimTime).div(86400);
        console.log("现在需要计算的天数:", newDays.toString());
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
            
            if (estimateGas.lt(25000000)) {
                console.log("✅ Gas 消耗在安全范围内");
            }
            
        } catch (e: any) {
            console.log("⚠️ 仍无法计算:", e.message);
        }
        
        console.log("\n========================================");
        console.log("总结");
        console.log("========================================");
        console.log(`
✅ 时间戳已调整为 ${suggestedDays} 天
💡 现在可以尝试领取奖励了

注意:
- 这会将计算天数从 ${currentDays.toString()} 天减少到 ${suggestedDays} 天
- 可以成功领取约 ${suggestedDays} 天的奖励
- 剩余的 ${currentDays.sub(suggestedDays).toString()} 天奖励可以分批领取

下一步:
1. 尝试领取奖励
2. 如果成功，可以再次调整时间戳继续领取剩余奖励
3. 重复此过程直到领取完所有奖励
    `);
        
    } catch (e: any) {
        console.log("❌ 调整失败:", e.message);
        if (e.data) {
            console.log("错误数据:", e.data);
        }
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });

