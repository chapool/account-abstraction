import { ethers } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 启用测试模式并设置初始时间戳
 * 这样就可以通过 fastForwardMinutes 等函数来模拟时间流逝
 */

async function main() {
    console.log("\n⏰ 启用 Staking 合约测试模式...\n");

    const [deployer] = await ethers.getSigners();
    console.log("操作者地址:", deployer.address);

    // 读取部署记录
    const deploymentPath = path.join(__dirname, "../.openzeppelin/sepolia.json");
    if (!fs.existsSync(deploymentPath)) {
        throw new Error("找不到部署记录文件");
    }

    const deployment = JSON.parse(fs.readFileSync(deploymentPath, "utf-8"));
    
    // 查找 Staking 合约地址
    let stakingAddress: string | undefined;
    for (const [address, info] of Object.entries(deployment.proxies)) {
        const proxyInfo = info as any;
        if (proxyInfo.kind === "transparent" && proxyInfo.address) {
            const implementationAddress = proxyInfo.address;
            const implementationInfo = deployment.impls?.[implementationAddress];
            if (implementationInfo?.layout?.types) {
                const typesStr = JSON.stringify(implementationInfo.layout.types);
                if (typesStr.includes("StakeInfo") && typesStr.includes("ComboStatus")) {
                    stakingAddress = address;
                    break;
                }
            }
        }
    }

    if (!stakingAddress) {
        throw new Error("找不到 Staking 合约地址");
    }

    console.log("Staking 地址:", stakingAddress);
    console.log();

    // 连接到合约
    const Staking = await ethers.getContractFactory("Staking");
    const staking = Staking.attach(stakingAddress);

    // 检查当前状态
    const currentTestMode = await staking.testMode();
    console.log("当前测试模式状态:", currentTestMode ? "已启用" : "未启用");

    if (currentTestMode) {
        const currentTestTimestamp = await staking.testTimestamp();
        console.log("当前测试时间戳:", currentTestTimestamp.toString());
        const date = new Date(Number(currentTestTimestamp) * 1000);
        console.log("对应日期:", date.toISOString());
        console.log("\n⚠️  测试模式已经启用，如果需要重置，请先禁用再重新启用。\n");
        return;
    }

    // 启用测试模式，使用当前区块时间作为初始时间
    console.log("正在启用测试模式...");
    const tx = await staking.enableTestMode(0); // 0 表示使用当前 block.timestamp
    await tx.wait();

    const testTimestamp = await staking.testTimestamp();
    const testMode = await staking.testMode();

    console.log("\n✅ 测试模式已启用！");
    console.log("========================================");
    console.log("测试模式:", testMode ? "✓ 已启用" : "✗ 未启用");
    console.log("测试时间戳:", testTimestamp.toString());
    const date = new Date(Number(testTimestamp) * 1000);
    console.log("对应日期:", date.toISOString());
    console.log("========================================");

    console.log("\n📝 现在可以使用以下函数来控制时间：");
    console.log("  - fastForwardMinutes(分钟数)：快进指定分钟");
    console.log("  - fastForwardDays(天数)：快进指定天数");
    console.log("  - fastForwardTime(秒数)：快进指定秒数");
    console.log("  - setTestTimestamp(时间戳)：设置到指定时间戳");
    console.log("  - disableTestMode()：禁用测试模式，恢复使用真实时间\n");

    console.log("💡 示例用法：");
    console.log("  npx hardhat run scripts/fast-forward-time.ts --network sepolia\n");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });

