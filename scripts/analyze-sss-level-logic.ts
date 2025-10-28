import { ethers } from "hardhat";

async function main() {
    console.log("🔍 分析SSS级别的组合加成逻辑...\n");

    // 获取合约实例
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const stakingContract = await ethers.getContractAt("Staking", stakingAddress);

    // 获取配置合约实例
    const configAddress = await stakingContract.configContract();
    const configContract = await ethers.getContractAt("StakingConfig", configAddress);

    console.log("==========================================");
    console.log("SSS级别组合加成逻辑分析");
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
        console.log("当前SSS级别逻辑分析");
        console.log("==========================================");

        console.log("📋 当前 _calculateComboBonusByCount 函数:");
        console.log("```solidity");
        console.log("function _calculateComboBonusByCount(uint256 count, uint8 level) internal view returns (uint256) {");
        console.log("    // SSS级别(level 6)不适用10个NFT组合");
        console.log("    if (level == 6 && count >= 10) {");
        console.log("        return 0;  // ❌ 只限制了10个NFT组合");
        console.log("    }");
        console.log("    ");
        console.log("    // 其他逻辑...");
        console.log("    // ❌ 问题：SSS级别仍然可以参与3个和5个NFT组合");
        console.log("}");
        console.log("```");

        console.log("\n==========================================");
        console.log("问题分析");
        console.log("==========================================");

        console.log("❌ 问题1：SSS级别仍然可以参与3个NFT组合");
        console.log("   - 用户有3个SSS级NFT");
        console.log("   - 当前逻辑：返回5%加成");
        console.log("   - 期望逻辑：应该返回0%加成");

        console.log("\n❌ 问题2：SSS级别仍然可以参与5个NFT组合");
        console.log("   - 用户有5个SSS级NFT");
        console.log("   - 当前逻辑：返回10%加成");
        console.log("   - 期望逻辑：应该返回0%加成");

        console.log("\n❌ 问题3：逻辑不一致");
        console.log("   - 10个SSS级NFT：0%加成（正确）");
        console.log("   - 3个SSS级NFT：5%加成（错误）");
        console.log("   - 5个SSS级NFT：10%加成（错误）");

        console.log("\n==========================================");
        console.log("正确的SSS级别逻辑");
        console.log("==========================================");

        console.log("✅ 正确的逻辑应该是：");
        console.log("```solidity");
        console.log("function _calculateComboBonusByCount(uint256 count, uint8 level) internal view returns (uint256) {");
        console.log("    // SSS级别(level 6)完全不参与组合加成");
        console.log("    if (level == 6) {");
        console.log("        return 0;  // ✅ 所有SSS级别都返回0%加成");
        console.log("    }");
        console.log("    ");
        console.log("    // 其他级别的组合加成逻辑...");
        console.log("}");
        console.log("```");

        console.log("\n==========================================");
        console.log("修复后的效果");
        console.log("==========================================");

        console.log("🎯 修复后的SSS级别行为:");
        console.log("   - 1个SSS级NFT：0%加成");
        console.log("   - 3个SSS级NFT：0%加成");
        console.log("   - 5个SSS级NFT：0%加成");
        console.log("   - 10个SSS级NFT：0%加成");
        console.log("   - 任何数量的SSS级NFT：0%加成");

        console.log("\n✅ 优势:");
        console.log("   1. 逻辑一致：SSS级别完全不参与组合加成");
        console.log("   2. 符合设计：SSS级别因稀缺性不适用组合加成");
        console.log("   3. 简化逻辑：不需要复杂的条件判断");

        console.log("\n==========================================");
        console.log("建议修复");
        console.log("==========================================");

        console.log("🔧 建议修改 _calculateComboBonusByCount 函数:");
        console.log("```solidity");
        console.log("function _calculateComboBonusByCount(uint256 count, uint8 level) internal view returns (uint256) {");
        console.log("    // SSS级别(level 6)完全不参与组合加成");
        console.log("    if (level == 6) {");
        console.log("        return 0;");
        console.log("    }");
        console.log("    ");
        console.log("    // Get combo config");
        console.log("    uint256[3] memory thresholds = configContract.getComboThresholds();");
        console.log("    uint256[3] memory bonuses = configContract.getComboBonuses();");
        console.log("    ");
        console.log("    // Check combo thresholds in reverse order (highest first)");
        console.log("    for (uint256 i = thresholds.length; i > 0; i--) {");
        console.log("        uint256 threshold = thresholds[i - 1];");
        console.log("        ");
        console.log("        if (count >= threshold) {");
        console.log("            return bonuses[i - 1];");
        console.log("        }");
        console.log("    }");
        console.log("    ");
        console.log("    return 0;");
        console.log("}");
        console.log("```");

    } catch (error) {
        console.error("❌ 分析过程中出现错误:", error);
    }

    console.log("\n🎉 分析完成！");
    console.log("\n📋 总结：");
    console.log("   - SSS级别应该完全不参与组合加成");
    console.log("   - 当前逻辑只限制了10个NFT组合");
    console.log("   - 需要修复为完全排除SSS级别");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });


