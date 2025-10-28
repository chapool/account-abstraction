import { ethers } from "hardhat";

async function main() {
    console.log("🔍 分析组合加成计算逻辑...\n");

    // 获取合约实例
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const stakingContract = await ethers.getContractAt("Staking", stakingAddress);

    // 测试用户地址
    const testUser = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";

    console.log("==========================================");
    console.log("组合加成逻辑分析");
    console.log("==========================================");

    try {
        // 获取用户C级NFT数量
        const level1Count = await stakingContract.userLevelCounts(testUser, 1);
        console.log(`用户C级NFT数量: ${level1Count}`);

        // 获取组合状态
        const comboStatus = await stakingContract.userComboStatus(testUser, 1);
        console.log(`\n组合状态详情:`);
        console.log(`- 等级: ${comboStatus.level}`);
        console.log(`- 数量: ${comboStatus.count}`);
        console.log(`- 加成: ${comboStatus.bonus / 100}%`);
        console.log(`- 是否待生效: ${comboStatus.isPending}`);
        console.log(`- 生效时间: ${new Date(Number(comboStatus.effectiveFrom) * 1000).toLocaleString()}`);

        // 获取当前时间
        const currentTime = Math.floor(Date.now() / 1000);
        console.log(`- 当前时间: ${new Date(currentTime * 1000).toLocaleString()}`);

        // 计算剩余等待时间
        const remainingTime = Number(comboStatus.effectiveFrom) - currentTime;
        const remainingDays = Math.ceil(remainingTime / (24 * 60 * 60));
        console.log(`- 剩余等待时间: ${remainingDays} 天`);

        console.log("\n==========================================");
        console.log("关键问题分析");
        console.log("==========================================");

        // 分析 _calculateComboBonus 函数的逻辑
        console.log("📋 _calculateComboBonus 函数逻辑:");
        console.log("1. 检查是否有待生效的状态 (isPending = true)");
        console.log("2. 如果当前时间 >= 生效时间，返回待生效的加成");
        console.log("3. 否则，基于当前NFT数量计算加成");

        console.log("\n🔍 当前情况分析:");
        if (comboStatus.isPending) {
            console.log("✅ 有待生效的组合状态");
            if (currentTime >= Number(comboStatus.effectiveFrom)) {
                console.log("✅ 当前时间已超过生效时间");
                console.log("✅ 应该使用待生效的加成:", comboStatus.bonus / 100 + "%");
            } else {
                console.log("⚠️  当前时间未到生效时间");
                console.log("⚠️  应该使用当前NFT数量的加成");
                
                // 计算当前NFT数量对应的加成
                const currentBonus = await stakingContract.getEffectiveComboBonus(testUser, 1);
                console.log("📊 当前NFT数量对应的加成:", currentBonus / 100 + "%");
            }
        } else {
            console.log("ℹ️  没有待生效的组合状态");
            console.log("ℹ️  使用当前NFT数量的加成");
        }

        console.log("\n==========================================");
        console.log("问题根源");
        console.log("==========================================");
        console.log("❌ 问题: 查询脚本可能直接读取了 userComboStatus 中的 bonus 值");
        console.log("❌ 但没有检查 isPending 和 effectiveFrom 的状态");
        console.log("✅ 正确做法: 应该调用 calculateComboBonus 函数获取实际生效的加成");

        // 获取实际生效的加成
        const actualBonus = await stakingContract.getEffectiveComboBonus(testUser, 1);
        console.log(`\n🎯 实际生效的加成: ${actualBonus / 100}%`);

        if (actualBonus === 0) {
            console.log("✅ 确认: 当前没有生效的组合加成");
            console.log("✅ 需要等待到生效时间才能享受10%加成");
        } else {
            console.log("⚠️  意外: 当前已有生效的组合加成");
        }

    } catch (error) {
        console.error("❌ 分析过程中出现错误:", error);
    }

    console.log("\n🎉 组合加成逻辑分析完成！");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
