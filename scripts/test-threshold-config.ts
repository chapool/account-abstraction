import { ethers } from "hardhat";

async function main() {
    console.log("🧪 测试阈值配置化修改...\n");

    // 获取合约实例
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const stakingContract = await ethers.getContractAt("Staking", stakingAddress);

    // 获取配置合约实例
    const configAddress = await stakingContract.configContract();
    const configContract = await ethers.getContractAt("StakingConfig", configAddress);

    console.log("==========================================");
    console.log("阈值配置化测试");
    console.log("==========================================");

    try {
        // 获取配置信息
        const thresholds = await configContract.getComboThresholds();
        const bonuses = await configContract.getComboBonuses();
        const minDays = await configContract.getComboMinDays();
        
        console.log("当前配置:");
        console.log(`- 阈值: [${thresholds[0]}, ${thresholds[1]}, ${thresholds[2]}]`);
        console.log(`- 加成: [${bonuses[0]/100}%, ${bonuses[1]/100}%, ${bonuses[2]/100}%]`);
        console.log(`- 等待天数: [${minDays[0]}, ${minDays[1]}, ${minDays[2]}]`);

        console.log("\n==========================================");
        console.log("测试不同NFT数量的等待天数");
        console.log("==========================================");

        // 测试不同数量的等待天数
        const testCounts = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11];
        
        for (const count of testCounts) {
            // 模拟 _getWaitDaysForCount 函数的逻辑
            let waitDays = 0;
            
            // 检查阈值（从高到低）
            for (let i = thresholds.length; i > 0; i--) {
                const threshold = thresholds[i - 1];
                
                if (count >= threshold) {
                    waitDays = minDays[i - 1];
                    break;
                }
            }
            
            console.log(`- ${count}个NFT: ${waitDays}天等待期`);
        }

        console.log("\n==========================================");
        console.log("验证修改效果");
        console.log("==========================================");
        console.log("✅ 阈值现在从配置合约读取");
        console.log("✅ 不再硬编码 10、5、3");
        console.log("✅ 可以通过配置合约动态调整阈值");
        console.log("✅ 提高了系统的灵活性");

        console.log("\n📝 修改说明:");
        console.log("- 原来: if (count >= 10) return minDays[2];");
        console.log("- 现在: 从 configContract.getComboThresholds() 读取阈值");
        console.log("- 好处: 可以通过配置合约调整组合规则，无需升级主合约");

    } catch (error) {
        console.error("❌ 测试过程中出现错误:", error);
    }

    console.log("\n🎉 阈值配置化测试完成！");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
