import { ethers } from "hardhat";

async function main() {
    console.log("🔍 分析新NFT质押对用户加成的影响...\n");

    // 获取合约实例
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const stakingContract = await ethers.getContractAt("Staking", stakingAddress);

    // 获取配置合约实例
    const configAddress = await stakingContract.configContract();
    const configContract = await ethers.getContractAt("StakingConfig", configAddress);

    console.log("==========================================");
    console.log("新NFT质押对加成的影响分析");
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
        console.log("关键问题分析");
        console.log("==========================================");

        console.log("❓ 问题：新NFT质押进来会不会导致用户加成延后？");
        console.log("\n📋 _updateComboStatus 函数逻辑:");
        console.log("1. 检查NFT数量是否发生变化");
        console.log("2. 如果变化，重新计算等待期");
        console.log("3. 设置新的生效时间 = 当前时间 + 等待天数");

        console.log("\n==========================================");
        console.log("具体场景分析");
        console.log("==========================================");

        console.log("🎯 场景1：用户已有5个NFT，等待15天后获得10%加成");
        console.log("   - 当前状态：5个NFT，等待15天");
        console.log("   - 用户再质押1个NFT");
        console.log("   - 新状态：6个NFT，仍然等待15天（5个NFT组合）");
        console.log("   - 结果：✅ 不会延后，因为6个NFT仍然属于5个NFT组合");

        console.log("\n🎯 场景2：用户已有4个NFT，等待7天后获得5%加成");
        console.log("   - 当前状态：4个NFT，等待7天");
        console.log("   - 用户再质押1个NFT");
        console.log("   - 新状态：5个NFT，等待15天（5个NFT组合）");
        console.log("   - 结果：❌ 会延后！从7天变成15天");

        console.log("\n🎯 场景3：用户已有9个NFT，等待15天后获得10%加成");
        console.log("   - 当前状态：9个NFT，等待15天");
        console.log("   - 用户再质押1个NFT");
        console.log("   - 新状态：10个NFT，等待50天（10个NFT组合）");
        console.log("   - 结果：❌ 会延后！从15天变成50天");

        console.log("\n==========================================");
        console.log("问题根源");
        console.log("==========================================");

        console.log("❌ 问题：");
        console.log("   - 每次NFT数量变化都会重新计算等待期");
        console.log("   - 新的生效时间 = 当前时间 + 等待天数");
        console.log("   - 这会导致之前等待的时间被重置");

        console.log("\n✅ 解决方案建议：");
        console.log("   1. 检查新组合是否比当前组合更好");
        console.log("   2. 如果更好，才更新等待期");
        console.log("   3. 如果更差，保持当前状态");

        console.log("\n==========================================");
        console.log("代码分析");
        console.log("==========================================");

        console.log("📋 当前代码问题：");
        console.log("```solidity");
        console.log("if (status.count != currentCount) {");
        console.log("    // 任何数量变化都会重新计算等待期");
        console.log("    status.effectiveFrom = _getCurrentTimestamp() + (waitDays * 1 days);");
        console.log("}");
        console.log("```");

        console.log("\n💡 建议的修复：");
        console.log("```solidity");
        console.log("if (status.count != currentCount) {");
        console.log("    uint256 newBonus = _calculateComboBonusByCount(currentCount, level);");
        console.log("    // 只有当新加成更好时才更新");
        console.log("    if (newBonus > status.bonus) {");
        console.log("        status.effectiveFrom = _getCurrentTimestamp() + (waitDays * 1 days);");
        console.log("        status.bonus = newBonus;");
        console.log("    }");
        console.log("}");
        console.log("```");

    } catch (error) {
        console.error("❌ 分析过程中出现错误:", error);
    }

    console.log("\n🎉 分析完成！");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
