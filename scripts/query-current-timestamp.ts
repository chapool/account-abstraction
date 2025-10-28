import { ethers } from "hardhat";

async function main() {
    console.log("\n⏰ 查询 Staking 合约的 currentTimestamp\n");

    const [deployer] = await ethers.getSigners();
    console.log("查询者地址:", deployer.address);

    // Staking 合约地址
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    
    // 连接到 Staking 合约
    const staking = await ethers.getContractAt("Staking", stakingAddress);

    console.log("========================================");
    console.log("currentTimestamp 说明");
    console.log("========================================");
    
    // 获取 testMode
    const testMode = await staking.testMode();
    console.log("\n测试模式状态:", testMode ? "✅ 已启用" : "❌ 未启用");
    
    // 获取 testTimestamp
    const testTimestamp = await staking.testTimestamp();
    const testDate = new Date(Number(testTimestamp) * 1000);
    
    console.log("\ntestTimestamp (测试时间戳):");
    console.log("  数值:", testTimestamp.toString());
    console.log("  日期:", testDate.toISOString());
    console.log("  本地时间:", testDate.toLocaleString());
    
    // 获取 currentTimestamp
    const currentTimestamp = await staking.getCurrentTimestamp();
    const currentDate = new Date(Number(currentTimestamp) * 1000);
    
    console.log("\ncurrentTimestamp (当前使用的时间戳):");
    console.log("  数值:", currentTimestamp.toString());
    console.log("  日期:", currentDate.toISOString());
    console.log("  本地时间:", currentDate.toLocaleString());
    
    // 获取区块时间戳
    const blockNumber = await ethers.provider.getBlockNumber();
    const block = await ethers.provider.getBlock(blockNumber);
    const blockTimestamp = block?.timestamp || 0;
    const blockDate = new Date(blockTimestamp * 1000);
    
    console.log("\n区块时间戳 (block.timestamp):");
    console.log("  数值:", blockTimestamp.toString());
    console.log("  日期:", blockDate.toISOString());
    console.log("  本地时间:", blockDate.toLocaleString());
    
    console.log("\n========================================");
    console.log("时间差分析");
    console.log("========================================");
    
    if (testMode) {
        const diff = Number(currentTimestamp) - blockTimestamp;
        console.log("\n测试时间 vs 真实区块时间:");
        console.log("  时间差:", diff, "秒");
        console.log("  相当于:", (diff / 60).toFixed(2), "分钟");
        console.log("  相当于:", (diff / 3600).toFixed(2), "小时");
        console.log("  相当于:", (diff / 86400).toFixed(2), "天");
        
        if (diff > 0) {
            console.log("\n✅ 测试时间比真实区块时间快了", (diff / 86400).toFixed(2), "天");
        } else if (diff < 0) {
            console.log("\n⚠️ 测试时间比真实区块时间慢了", Math.abs(diff / 86400).toFixed(2), "天");
        } else {
            console.log("\n⏸️ 测试时间与真实区块时间相同");
        }
    }
    
    console.log("\n========================================");
    console.log("关键说明");
    console.log("========================================");
    console.log(`
currentTimestamp 的逻辑:

if (testMode) {
    return testTimestamp;  // 测试模式：返回测试时间戳
} else {
    return block.timestamp;  // 生产模式：返回区块时间戳
}

当前状态:
- testMode: ${testMode}
- testTimestamp: ${testTimestamp} (${testDate.toISOString()})
- currentTimestamp: ${currentTimestamp} (${currentDate.toISOString()})
- block.timestamp: ${blockTimestamp} (${blockDate.toISOString()})

结论:
- 合约内部所有的 _getCurrentTimestamp() 调用都会返回 ${currentTimestamp}
- 质押、领取、奖励计算等所有时间相关的操作都使用这个时间戳
- ${testMode ? '当前处于测试模式，时间可人为控制' : '当前处于生产模式，使用真实区块时间'}
    `);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 查询失败:", error);
        process.exit(1);
    });

