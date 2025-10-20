import { ethers } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 快进时间的脚本
 * 支持按分钟、天或秒来快进
 */

async function main() {
    console.log("\n⏰ Staking 合约时间快进工具\n");

    // 从命令行参数获取快进时间
    const args = process.argv.slice(2);
    let minutes = 0;
    let days = 0;
    let seconds = 0;

    // 解析参数
    for (let i = 0; i < args.length; i++) {
        if (args[i] === '--minutes' || args[i] === '-m') {
            minutes = parseInt(args[i + 1]);
            i++;
        } else if (args[i] === '--days' || args[i] === '-d') {
            days = parseInt(args[i + 1]);
            i++;
        } else if (args[i] === '--seconds' || args[i] === '-s') {
            seconds = parseInt(args[i + 1]);
            i++;
        }
    }

    if (minutes === 0 && days === 0 && seconds === 0) {
        console.log("用法:");
        console.log("  npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 180");
        console.log("  npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --days 1");
        console.log("  npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --seconds 3600");
        console.log("\n选项:");
        console.log("  -m, --minutes <数量>  快进指定分钟数");
        console.log("  -d, --days <数量>     快进指定天数");
        console.log("  -s, --seconds <数量>  快进指定秒数");
        console.log("\n示例（1天 = 1分钟的测试模式）:");
        console.log("  快进 180 分钟（相当于 180 天）:");
        console.log("    npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 180\n");
        return;
    }

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

    // 检查测试模式
    const testMode = await staking.testMode();
    if (!testMode) {
        console.log("❌ 错误：测试模式未启用");
        console.log("请先运行: npx hardhat run scripts/enable-test-mode.ts --network sepolia\n");
        return;
    }

    // 获取当前测试时间
    const beforeTimestamp = await staking.testTimestamp();
    const beforeDate = new Date(Number(beforeTimestamp) * 1000);
    
    console.log("========================================");
    console.log("当前时间");
    console.log("========================================");
    console.log("时间戳:", beforeTimestamp.toString());
    console.log("日期:", beforeDate.toISOString());
    console.log();

    // 快进时间
    let tx;
    let description = "";
    
    if (minutes > 0) {
        console.log(`⏩ 快进 ${minutes} 分钟...`);
        tx = await staking.fastForwardMinutes(minutes);
        description = `${minutes} 分钟`;
    } else if (days > 0) {
        console.log(`⏩ 快进 ${days} 天...`);
        tx = await staking.fastForwardDays(days);
        description = `${days} 天`;
    } else if (seconds > 0) {
        console.log(`⏩ 快进 ${seconds} 秒...`);
        tx = await staking.fastForwardTime(seconds);
        description = `${seconds} 秒`;
    }

    await tx.wait();

    // 获取新的测试时间
    const afterTimestamp = await staking.testTimestamp();
    const afterDate = new Date(Number(afterTimestamp) * 1000);
    const elapsedSeconds = Number(afterTimestamp) - Number(beforeTimestamp);
    
    console.log("\n✅ 时间快进成功！");
    console.log("========================================");
    console.log("新时间");
    console.log("========================================");
    console.log("时间戳:", afterTimestamp.toString());
    console.log("日期:", afterDate.toISOString());
    console.log();
    console.log("实际快进:", description);
    console.log("快进秒数:", elapsedSeconds, "秒");
    console.log("快进分钟数:", (elapsedSeconds / 60).toFixed(2), "分钟");
    console.log("快进天数:", (elapsedSeconds / 86400).toFixed(4), "天");
    console.log("========================================\n");

    console.log("💡 提示：如果使用配置中 1天=1分钟 的模式，");
    console.log(`   本次快进相当于经过了 ${(elapsedSeconds / 60).toFixed(2)} 天\n`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });

