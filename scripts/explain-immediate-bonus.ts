import { ethers } from "hardhat";

async function main() {
    console.log("🔍 解释'基于当前NFT数量的加成是立即生效的'概念...\n");

    // 获取合约实例
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const stakingContract = await ethers.getContractAt("Staking", stakingAddress);

    // 获取配置合约实例
    const configAddress = await stakingContract.configContract();
    const configContract = await ethers.getContractAt("StakingConfig", configAddress);

    console.log("==========================================");
    console.log("概念解释：立即生效 vs 延迟生效");
    console.log("==========================================");

    try {
        // 获取配置信息
        const thresholds = await configContract.getComboThresholds();
        const bonuses = await configContract.getComboBonuses();
        const minDays = await configContract.getComboMinDays();
        
        console.log("组合配置:");
        console.log(`- 3个NFT阈值: ${thresholds[0]}, 加成: ${bonuses[0]/100}%, 等待: ${minDays[0]}天`);
        console.log(`- 5个NFT阈值: ${thresholds[1]}, 加成: ${bonuses[1]/100}%, 等待: ${minDays[1]}天`);
        console.log(`- 10个NFT阈值: ${thresholds[2]}, 加成: ${bonuses[2]/100}%, 等待: ${minDays[2]}天`);

        console.log("\n==========================================");
        console.log("关键概念解释");
        console.log("==========================================");

        console.log("📋 两种不同的加成机制:");
        console.log("\n1️⃣ 【立即生效】- 基于当前NFT数量");
        console.log("   - 用户质押5个C级NFT");
        console.log("   - 立即获得10%加成");
        console.log("   - 不需要等待任何时间");
        console.log("   - 这是用户实际享受的加成");

        console.log("\n2️⃣ 【延迟生效】- 基于组合状态更新");
        console.log("   - 当用户NFT数量发生变化时");
        console.log("   - 新的组合状态需要等待期");
        console.log("   - 等待期结束后，状态更新生效");
        console.log("   - 这主要用于状态管理");

        console.log("\n==========================================");
        console.log("实际例子分析");
        console.log("==========================================");

        console.log("🎯 场景：用户质押5个C级NFT");
        console.log("\n步骤1: 用户质押第1个NFT");
        console.log("   - 当前NFT数量: 1个");
        console.log("   - 立即生效加成: 0% (1 < 3)");
        console.log("   - 组合状态: 无");

        console.log("\n步骤2: 用户质押第3个NFT");
        console.log("   - 当前NFT数量: 3个");
        console.log("   - 立即生效加成: 5% (3 >= 3)");
        console.log("   - 组合状态: 3个NFT组合，等待7天");

        console.log("\n步骤3: 用户质押第5个NFT");
        console.log("   - 当前NFT数量: 5个");
        console.log("   - 立即生效加成: 10% (5 >= 5)");
        console.log("   - 组合状态: 5个NFT组合，等待15天");

        console.log("\n==========================================");
        console.log("关键理解");
        console.log("==========================================");

        console.log("✅ 用户立即享受10%加成");
        console.log("   - 不需要等待15天");
        console.log("   - 基于当前5个NFT数量计算");

        console.log("\n⚠️  组合状态显示'等待15天'");
        console.log("   - 这只是状态更新的等待期");
        console.log("   - 不影响当前的奖励计算");
        console.log("   - 用于下次状态变化时的管理");

        console.log("\n==========================================");
        console.log("代码逻辑分析");
        console.log("==========================================");

        console.log("📋 _calculateComboBonus 函数逻辑:");
        console.log("1. 检查是否有待生效状态且时间已到");
        console.log("2. 如果是，返回待生效的加成");
        console.log("3. 否则，基于当前NFT数量计算加成");

        console.log("\n🔍 当前情况:");
        console.log("- 用户有5个C级NFT");
        console.log("- 5个NFT >= 5个NFT阈值 ✅");
        console.log("- 因此立即获得10%加成");

        console.log("\n💡 为什么这样设计？");
        console.log("- 用户质押NFT后应该立即享受加成");
        console.log("- 等待期只影响状态更新，不影响奖励");
        console.log("- 这样用户体验更好");

    } catch (error) {
        console.error("❌ 解释过程中出现错误:", error);
    }

    console.log("\n🎉 概念解释完成！");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
