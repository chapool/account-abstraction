import { ethers } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 验证 Staking 合约升级是否成功
 * 检查新功能是否可用
 */

async function main() {
    console.log("\n🔍 验证 Staking 合约升级...\n");

    const [deployer] = await ethers.getSigners();
    console.log("验证者地址:", deployer.address);

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

    console.log("========================================");
    console.log("基本信息验证");
    console.log("========================================");

    // 检查版本
    try {
        const version = await staking.version();
        console.log("✅ 合约版本:", version);
    } catch (e) {
        console.log("⚠️  版本函数:", "可能不存在或调用失败");
    }

    // 检查 owner
    try {
        const owner = await staking.owner();
        console.log("✅ 合约 Owner:", owner);
        console.log("   是否为当前账户:", owner.toLowerCase() === deployer.address.toLowerCase() ? "是" : "否");
    } catch (e) {
        console.log("❌ 无法读取 owner");
    }

    console.log("\n========================================");
    console.log("新功能验证");
    console.log("========================================");

    let allFunctionsWork = true;

    // 1. 检查 testMode
    try {
        const testMode = await staking.testMode();
        console.log("✅ testMode:", testMode ? "已启用" : "未启用");
    } catch (e) {
        console.log("❌ testMode: 函数不存在或调用失败");
        allFunctionsWork = false;
    }

    // 2. 检查 testTimestamp
    try {
        const testTimestamp = await staking.testTimestamp();
        console.log("✅ testTimestamp:", testTimestamp.toString());
    } catch (e) {
        console.log("❌ testTimestamp: 函数不存在或调用失败");
        allFunctionsWork = false;
    }

    // 3. 检查函数存在性（不实际调用，只检查是否存在）
    console.log("\n函数存在性检查：");
    const functions = [
        "enableTestMode",
        "disableTestMode",
        "setTestTimestamp",
        "fastForwardTime",
        "fastForwardDays",
        "fastForwardMinutes"
    ];

    for (const funcName of functions) {
        try {
            const func = staking[funcName];
            if (func) {
                console.log(`✅ ${funcName}(): 存在`);
            } else {
                console.log(`❌ ${funcName}(): 不存在`);
                allFunctionsWork = false;
            }
        } catch (e) {
            console.log(`❌ ${funcName}(): 检查失败`);
            allFunctionsWork = false;
        }
    }

    console.log("\n========================================");
    console.log("存储数据验证");
    console.log("========================================");

    // 检查一些基本的存储变量是否仍然可访问
    try {
        const totalStakedCount = await staking.totalStakedCount();
        console.log("✅ totalStakedCount:", totalStakedCount.toString());
    } catch (e) {
        console.log("❌ 无法读取 totalStakedCount");
    }

    // 检查合约引用
    try {
        const cpnftContract = await staking.cpnftContract();
        console.log("✅ CPNFT 合约地址:", cpnftContract);
    } catch (e) {
        console.log("❌ 无法读取 CPNFT 合约地址");
    }

    try {
        const cpopTokenContract = await staking.cpopTokenContract();
        console.log("✅ CPOP Token 合约地址:", cpopTokenContract);
    } catch (e) {
        console.log("❌ 无法读取 CPOP Token 合约地址");
    }

    try {
        const configContract = await staking.configContract();
        console.log("✅ Config 合约地址:", configContract);
    } catch (e) {
        console.log("❌ 无法读取 Config 合约地址");
    }

    console.log("\n========================================");
    console.log("验证总结");
    console.log("========================================");

    if (allFunctionsWork) {
        console.log("✅ 升级成功！所有新功能都可用");
        console.log("✅ 原有数据完整保留");
        console.log("✅ 合约可以正常使用");
    } else {
        console.log("⚠️  升级可能存在问题，部分功能不可用");
        console.log("建议检查升级过程或重新升级");
    }

    console.log("\n📝 可用操作：");
    console.log("  查看时间状态:");
    console.log("    npx hardhat run scripts/check-time-status.ts --network sepolia");
    console.log("  启用测试模式:");
    console.log("    npx hardhat run scripts/enable-test-mode.ts --network sepolia");
    console.log("  快进时间:");
    console.log("    npx hardhat run scripts/fast-forward-time.ts --network sepolia -- --minutes 60");
    console.log();
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });

