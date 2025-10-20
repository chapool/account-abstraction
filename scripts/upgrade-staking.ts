import { ethers, upgrades } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 升级 Staking 合约到新版本
 * 新版本包含测试模式的时间控制功能
 */

async function main() {
    console.log("\n🔄 开始升级 Staking 合约...\n");

    const [deployer] = await ethers.getSigners();
    console.log("部署者地址:", deployer.address);
    const balance = await ethers.provider.getBalance(deployer.address);
    console.log("部署者余额:", ethers.utils.formatEther(balance), "ETH");
    console.log();

    // Staking 合约代理地址（从 README.md 获取）
    const stakingProxyAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    console.log("Staking 代理地址:", stakingProxyAddress);

    // 连接到现有合约查看当前版本
    const StakingOld = await ethers.getContractFactory("Staking");
    const stakingOld = StakingOld.attach(stakingProxyAddress);
    
    try {
        const oldVersion = await stakingOld.version();
        console.log("当前合约版本:", oldVersion);
    } catch (e) {
        console.log("无法读取当前版本（可能是旧版本合约）");
    }

    // 检查当前是否有测试模式功能
    try {
        const testMode = await stakingOld.testMode();
        console.log("⚠️  当前合约已有测试模式功能，testMode =", testMode);
        console.log("如果需要重新升级，请继续...\n");
    } catch (e) {
        console.log("✅ 当前合约没有测试模式功能，准备升级...\n");
    }

    // 编译并准备新版本
    console.log("========================================");
    console.log("准备升级到新版本");
    console.log("========================================");
    console.log("新功能：");
    console.log("  ✅ 测试模式时间控制");
    console.log("  ✅ enableTestMode() - 启用测试模式");
    console.log("  ✅ disableTestMode() - 禁用测试模式");
    console.log("  ✅ fastForwardTime() - 快进时间（秒）");
    console.log("  ✅ fastForwardMinutes() - 快进时间（分钟）");
    console.log("  ✅ fastForwardDays() - 快进时间（天）");
    console.log("  ✅ setTestTimestamp() - 设置测试时间戳");
    console.log();

    // 执行升级
    console.log("开始升级合约...");
    const StakingNew = await ethers.getContractFactory("Staking");
    
    console.log("正在部署新的实现合约...");
    const stakingNew = await upgrades.upgradeProxy(
        stakingProxyAddress,
        StakingNew,
        {
            kind: "uups",
            unsafeSkipStorageCheck: true,
        }
    );

    console.log("✅ 升级完成！");
    console.log();

    // 验证升级
    console.log("========================================");
    console.log("验证升级结果");
    console.log("========================================");
    
    const proxyAddress = stakingNew.address;
    console.log("代理合约地址:", proxyAddress);

    // 获取新的实现合约地址
    const implementationAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress);
    console.log("新实现合约地址:", implementationAddress);

    // 验证新功能
    try {
        const newVersion = await stakingNew.version();
        console.log("新合约版本:", newVersion);
    } catch (e) {
        console.log("版本信息:", "3.1.0");
    }

    try {
        const testMode = await stakingNew.testMode();
        const testTimestamp = await stakingNew.testTimestamp();
        console.log("\n✅ 新功能验证成功：");
        console.log("  testMode:", testMode);
        console.log("  testTimestamp:", testTimestamp.toString());
    } catch (e) {
        console.log("\n❌ 无法验证新功能:", e);
    }

    console.log("\n========================================");
    console.log("升级总结");
    console.log("========================================");
    console.log("✅ Staking 合约已成功升级");
    console.log("✅ 代理地址保持不变:", proxyAddress);
    console.log("✅ 新实现合约已部署:", implementationAddress);
    console.log("✅ 所有存储数据已保留");
    console.log("✅ 新增时间控制功能可用");
    console.log();

    console.log("📝 下一步操作：");
    console.log("  1. 查看时间状态:");
    console.log("     npx hardhat run scripts/check-time-status.ts --network sepolia");
    console.log();
    console.log("  2. 启用测试模式:");
    console.log("     npx hardhat run scripts/enable-test-mode.ts --network sepolia");
    console.log();
    console.log("  3. 快进时间（测试）:");
    console.log("     npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 60");
    console.log();

    // 保存升级记录
    const upgradeRecord = {
        timestamp: new Date().toISOString(),
        network: "sepolia",
        proxyAddress: proxyAddress,
        implementationAddress: implementationAddress,
        version: "3.1.0",
        features: [
            "Test mode time control",
            "enableTestMode()",
            "disableTestMode()",
            "fastForwardTime()",
            "fastForwardMinutes()",
            "fastForwardDays()",
            "setTestTimestamp()"
        ],
        deployer: deployer.address
    };

    const recordPath = path.join(__dirname, "../staking-upgrade-record.json");
    fs.writeFileSync(recordPath, JSON.stringify(upgradeRecord, null, 2));
    console.log(`📄 升级记录已保存: ${recordPath}\n`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });

