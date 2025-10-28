import { ethers } from "hardhat";

async function main() {
    console.log("🧪 测试组合加成等待期修复...\n");

    // 获取合约实例
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const stakingContract = await ethers.getContractAt("Staking", stakingAddress);

    // 测试用户地址
    const testUser = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";

    console.log("==========================================");
    console.log("组合加成等待期测试");
    console.log("==========================================");

    try {
        // 获取用户C级NFT数量
        const level1Count = await stakingContract.userLevelCounts(testUser, 1);
        console.log(`用户C级NFT数量: ${level1Count}`);

        // 获取组合状态
        const comboStatus = await stakingContract.getComboStatus(testUser, 1);
        console.log(`\n组合状态:`);
        console.log(`- 数量: ${comboStatus.count}`);
        console.log(`- 待生效加成: ${comboStatus.bonus/100}%`);
        console.log(`- 是否待生效: ${comboStatus.isPending}`);
        console.log(`- 生效时间: ${new Date(Number(comboStatus.effectiveFrom) * 1000).toLocaleString()}`);

        // 获取当前时间
        const currentTime = Math.floor(Date.now() / 1000);
        console.log(`- 当前时间: ${new Date(currentTime * 1000).toLocaleString()}`);

        // 计算剩余等待时间
        const remainingTime = Number(comboStatus.effectiveFrom) - currentTime;
        const remainingDays = Math.ceil(remainingTime / (24 * 60 * 60));
        console.log(`- 剩余等待时间: ${remainingDays} 天`);

        // 获取实际生效的加成
        const effectiveBonus = await stakingContract.getEffectiveComboBonus(testUser, 1);
        console.log(`\n实际生效的加成: ${effectiveBonus/100}%`);

        console.log("\n==========================================");
        console.log("修复验证");
        console.log("==========================================");

        if (effectiveBonus === 0) {
            console.log("✅ 修复成功！");
            console.log("✅ 用户当前没有组合加成");
            console.log("✅ 需要等待到生效时间才能获得10%加成");
            console.log(`✅ 还需要等待 ${remainingDays} 天`);
        } else {
            console.log("❌ 修复失败！");
            console.log("❌ 用户仍然有组合加成");
            console.log("❌ 应该等待到生效时间");
        }

        console.log("\n==========================================");
        console.log("修复说明");
        console.log("==========================================");
        console.log("📋 修复前的问题:");
        console.log("   - 用户质押5个NFT后立即获得10%加成");
        console.log("   - 忽略了等待期的设计");
        console.log("   - 不符合'次日生效'的规则");

        console.log("\n📋 修复后的逻辑:");
        console.log("   - 用户质押5个NFT后，组合状态设为待生效");
        console.log("   - 需要等待15天后才能获得10%加成");
        console.log("   - 符合'次日生效'的设计意图");

    } catch (error) {
        console.error("❌ 测试过程中出现错误:", error);
    }

    console.log("\n🎉 组合加成等待期修复测试完成！");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
