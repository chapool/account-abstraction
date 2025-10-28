import { ethers } from "hardhat";

async function main() {
    console.log("🔍 分析加成等级变化时的等待期更新逻辑...\n");

    // 获取合约实例
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const stakingContract = await ethers.getContractAt("Staking", stakingAddress);

    // 获取配置合约实例
    const configAddress = await stakingContract.configContract();
    const configContract = await ethers.getContractAt("StakingConfig", configAddress);

    console.log("==========================================");
    console.log("加成等级变化时的等待期更新逻辑分析");
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
        console.log("新的逻辑分析");
        console.log("==========================================");

        console.log("✅ 新逻辑：加成等级变化时更新等待期");
        console.log("```solidity");
        console.log("if (status.count != currentCount) {");
        console.log("    status.level = level;");
        console.log("    status.count = currentCount;");
        console.log("    ");
        console.log("    // Update if the bonus level changes (upgrade or downgrade)");
        console.log("    if (newBonus != status.bonus) {");
        console.log("        // 重新计算等待期");
        console.log("        status.effectiveFrom = _getCurrentTimestamp() + (waitDays * 1 days);");
        console.log("        status.bonus = newBonus;");
        console.log("        status.isPending = true;");
        console.log("    }");
        console.log("}");
        console.log("```");

        console.log("\n==========================================");
        console.log("具体场景分析");
        console.log("==========================================");

        console.log("🎯 场景1：用户已有4个NFT，等待7天后获得5%加成");
        console.log("   - 当前状态：4个NFT，等待7天，5%加成");
        console.log("   - 用户再质押1个NFT");
        console.log("   - 新状态：5个NFT，等待15天，10%加成");
        console.log("   - 结果：✅ 更新！从5%升级到10%，重新等待15天");

        console.log("\n🎯 场景2：用户已有5个NFT，等待15天后获得10%加成");
        console.log("   - 当前状态：5个NFT，等待15天，10%加成");
        console.log("   - 用户再质押1个NFT");
        console.log("   - 新状态：6个NFT，保持15天，10%加成");
        console.log("   - 结果：✅ 不更新！加成等级相同，保持当前状态");

        console.log("\n🎯 场景3：用户已有10个NFT，等待50天后获得20%加成");
        console.log("   - 当前状态：10个NFT，等待50天，20%加成");
        console.log("   - 用户取消质押1个NFT");
        console.log("   - 新状态：9个NFT，等待15天，10%加成");
        console.log("   - 结果：✅ 更新！从20%降级到10%，重新等待15天");

        console.log("\n🎯 场景4：用户已有3个NFT，等待7天后获得5%加成");
        console.log("   - 当前状态：3个NFT，等待7天，5%加成");
        console.log("   - 用户再质押2个NFT");
        console.log("   - 新状态：5个NFT，等待15天，10%加成");
        console.log("   - 结果：✅ 更新！从5%升级到10%，重新等待15天");

        console.log("\n==========================================");
        console.log("逻辑优势");
        console.log("==========================================");

        console.log("✅ 优势1：公平性");
        console.log("   - 升级：重新等待，获得更高加成");
        console.log("   - 降级：重新等待，避免立即损失");

        console.log("\n✅ 优势2：一致性");
        console.log("   - 所有加成等级变化都遵循相同规则");
        console.log("   - 用户行为可预测");

        console.log("\n✅ 优势3：平衡性");
        console.log("   - 防止频繁切换NFT数量来规避等待期");
        console.log("   - 鼓励用户稳定持有");

        console.log("\n==========================================");
        console.log("特殊情况处理");
        console.log("==========================================");

        console.log("🔍 特殊情况1：NFT数量变化但加成等级不变");
        console.log("   - 5个NFT → 6个NFT：都是10%加成");
        console.log("   - 结果：不更新等待期，保持当前状态");

        console.log("\n🔍 特殊情况2：NFT数量变化且加成等级变化");
        console.log("   - 4个NFT → 5个NFT：5% → 10%");
        console.log("   - 结果：更新等待期，重新计算生效时间");

        console.log("\n🔍 特殊情况3：NFT数量减少导致加成等级降低");
        console.log("   - 10个NFT → 9个NFT：20% → 10%");
        console.log("   - 结果：更新等待期，重新计算生效时间");

    } catch (error) {
        console.error("❌ 分析过程中出现错误:", error);
    }

    console.log("\n🎉 分析完成！");
    console.log("\n📋 总结：");
    console.log("   - 加成等级变化时（升级或降级）都会更新等待期");
    console.log("   - 加成等级不变时保持当前状态");
    console.log("   - 逻辑更加公平和一致");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
