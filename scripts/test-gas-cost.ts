import { ethers } from "hardhat";

async function main() {
    console.log("\n⛽ 测试 Gas 消耗\n");

    // 合约地址
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const tokenId = 4812;
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    
    // 连接到合约
    const staking = await ethers.getContractAt("Staking", stakingAddress);
    
    // 查询质押信息
    console.log("========================================");
    console.log("质押信息");
    console.log("========================================");
    const stakeInfo = await staking.stakes(tokenId);
    const currentTimestamp = await staking.getCurrentTimestamp();
    
    const totalDays = currentTimestamp.sub(stakeInfo.lastClaimTime).div(86400);
    
    console.log("质押时间:", new Date(Number(stakeInfo.stakeTime) * 1000).toISOString());
    console.log("最后领取时间:", new Date(Number(stakeInfo.lastClaimTime) * 1000).toISOString());
    console.log("当前时间:", new Date(Number(currentTimestamp) * 1000).toISOString());
    console.log("需要计算的天数:", totalDays.toString());
    console.log();
    
    if (totalDays.gt(100)) {
        console.log("⚠️ 警告：需要计算超过 100 天的奖励！");
        console.log("   这可能会导致 Gas 超限问题");
        console.log();
    }
    
    console.log("========================================");
    console.log("测试计算奖励的 Gas 消耗");
    console.log("========================================");
    
    try {
        // 使用静态调用模拟计算
        console.log("尝试调用 calculatePendingRewards...");
        
        // 使用 estimateGas
        const estimateGas = await staking.estimateGas.calculatePendingRewards(tokenId);
        console.log("✅ 估算 Gas:", estimateGas.toString());
        console.log("   Gas 限制: 30,000,000");
        console.log("   使用率:", ((Number(estimateGas) / 30000000) * 100).toFixed(2) + "%");
        
        if (estimateGas.gt(25000000)) {
            console.log("\n❌ Gas 消耗接近或超过限制！");
            console.log("   这会导致交易失败");
        }
        
    } catch (e: any) {
        console.log("❌ 无法估算 Gas:", e.message);
        
        if (e.message.includes("gas")) {
            console.log("\n💡 这就是问题所在！");
            console.log("   calculatePendingRewards 函数计算的天数太多");
            console.log("   超过了 Gas 限制");
        }
    }
    
    console.log("\n========================================");
    console.log("解决方案");
    console.log("========================================");
    console.log(`
问题原因:
1. NFT 质押了 ${totalDays.toString()} 天
2. _calculatePendingRewards 函数会循环 ${totalDays.toString()} 次
3. 每次循环都包含复杂的计算（衰减、历史调整等）
4. 总计算量超过了 Gas 限制

可能的解决方案:

方案 1: 分批领取（推荐）
- 不要一次领取 ${totalDays.toString()} 天的奖励
- 可以分多次领取，每次领取一定天数的奖励
- 例如：每 90 天领取一次

方案 2: 修改合约逻辑
- 优化计算逻辑，使用数学公式代替循环
- 或者限制单次计算的天数

方案 3: 快进时间（仅测试环境）
- 将 testTimestamp 调整到接近质押时间
- 这样可以先测试短时间的领取逻辑
    `);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });

