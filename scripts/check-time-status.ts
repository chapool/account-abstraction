import { ethers } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 查看 Staking 合约的时间状态
 */

async function main() {
    console.log("\n⏰ Staking 合约时间状态\n");

    const [deployer] = await ethers.getSigners();
    console.log("查询者地址:", deployer.address);

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

    // 获取状态
    const testMode = await staking.testMode();
    const testTimestamp = await staking.testTimestamp();

    // 获取当前区块时间
    const currentBlock = await ethers.provider.getBlock("latest");
    const blockTimestamp = currentBlock?.timestamp || 0;
    const blockDate = new Date(blockTimestamp * 1000);

    console.log("========================================");
    console.log("时间状态");
    console.log("========================================");
    console.log("测试模式:", testMode ? "✅ 已启用" : "❌ 未启用");
    console.log();

    if (testMode) {
        const testDate = new Date(Number(testTimestamp) * 1000);
        console.log("📍 当前使用时间（测试时间）:");
        console.log("  时间戳:", testTimestamp.toString());
        console.log("  日期:", testDate.toISOString());
        console.log("  本地时间:", testDate.toLocaleString());
        console.log();
        console.log("🕐 真实区块时间:");
        console.log("  时间戳:", blockTimestamp.toString());
        console.log("  日期:", blockDate.toISOString());
        console.log("  本地时间:", blockDate.toLocaleString());
        console.log();
        
        const diff = Number(testTimestamp) - blockTimestamp;
        if (diff > 0) {
            console.log("⏩ 测试时间领先真实时间:", Math.abs(diff), "秒");
            console.log("   相当于:", (Math.abs(diff) / 60).toFixed(2), "分钟");
            console.log("   相当于:", (Math.abs(diff) / 3600).toFixed(2), "小时");
            console.log("   相当于:", (Math.abs(diff) / 86400).toFixed(4), "天");
        } else if (diff < 0) {
            console.log("⏪ 测试时间落后真实时间:", Math.abs(diff), "秒");
            console.log("   相当于:", (Math.abs(diff) / 60).toFixed(2), "分钟");
            console.log("   相当于:", (Math.abs(diff) / 3600).toFixed(2), "小时");
            console.log("   相当于:", (Math.abs(diff) / 86400).toFixed(4), "天");
        } else {
            console.log("⏸️  测试时间与真实时间相同");
        }
    } else {
        console.log("📍 当前使用时间（真实区块时间）:");
        console.log("  时间戳:", blockTimestamp.toString());
        console.log("  日期:", blockDate.toISOString());
        console.log("  本地时间:", blockDate.toLocaleString());
    }
    
    console.log("========================================");

    if (testMode) {
        console.log("\n💡 可用操作：");
        console.log("  - 快进时间: npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 60");
        console.log("  - 禁用测试模式: npx hardhat run scripts/disable-test-mode.ts --network sepolia");
    } else {
        console.log("\n💡 可用操作：");
        console.log("  - 启用测试模式: npx hardhat run scripts/enable-test-mode.ts --network sepolia");
    }
    console.log();
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });

