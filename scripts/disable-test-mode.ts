import { ethers } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 禁用测试模式，恢复使用真实的 block.timestamp
 */

async function main() {
    console.log("\n⏰ 禁用 Staking 合约测试模式...\n");

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

    if (!currentTestMode) {
        console.log("\n⚠️  测试模式已经是禁用状态，无需再次禁用。\n");
        return;
    }

    const currentTestTimestamp = await staking.testTimestamp();
    const date = new Date(Number(currentTestTimestamp) * 1000);
    console.log("当前测试时间戳:", currentTestTimestamp.toString());
    console.log("对应日期:", date.toISOString());
    console.log();

    // 禁用测试模式
    console.log("正在禁用测试模式...");
    const tx = await staking.disableTestMode();
    await tx.wait();

    const testMode = await staking.testMode();
    const testTimestamp = await staking.testTimestamp();

    console.log("\n✅ 测试模式已禁用！");
    console.log("========================================");
    console.log("测试模式:", testMode ? "✓ 已启用" : "✗ 已禁用");
    console.log("测试时间戳已重置:", testTimestamp.toString());
    console.log("========================================");
    console.log("\n现在合约将使用真实的 block.timestamp\n");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });

