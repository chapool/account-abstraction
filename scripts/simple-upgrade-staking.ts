import { ethers, upgrades } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 简单升级 Staking 合约
 * 先尝试直接升级，如果失败则提供诊断信息
 */

async function main() {
    console.log("\n🔄 开始简单升级 Staking 合约...\n");

    const [deployer] = await ethers.getSigners();
    console.log("部署者地址:", deployer.address);
    const balance = await ethers.provider.getBalance(deployer.address);
    console.log("部署者余额:", ethers.utils.formatEther(balance), "ETH");
    console.log();

    // Staking 合约代理地址
    const stakingProxyAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    console.log("Staking 代理地址:", stakingProxyAddress);

    // 连接到现有合约
    const StakingOld = await ethers.getContractFactory("Staking");
    const stakingOld = StakingOld.attach(stakingProxyAddress);
    
    try {
        const oldVersion = await stakingOld.version();
        console.log("当前合约版本:", oldVersion);
    } catch (e) {
        console.log("无法读取当前版本");
    }

    // 检查合约状态
    try {
        const configAddress = await stakingOld.configContract();
        console.log("配置合约地址:", configAddress);
        
        const totalStaked = await stakingOld.totalStakedCount();
        console.log("当前质押总数:", totalStaked.toString());
        
        const paused = await stakingOld.paused();
        console.log("合约是否暂停:", paused);
        
    } catch (e) {
        console.log("无法读取合约状态:", e.message);
    }

    console.log("\n========================================");
    console.log("尝试升级");
    console.log("========================================");

    try {
        // 编译新版本
        console.log("正在编译新版本...");
        const StakingNew = await ethers.getContractFactory("Staking");
        
        // 尝试升级
        console.log("正在升级合约...");
        const stakingNew = await upgrades.upgradeProxy(
            stakingProxyAddress,
            StakingNew,
            {
                kind: "uups",
                unsafeSkipStorageCheck: true,
                gasLimit: 10000000, // 更高的 gas 限制
            }
        );

        console.log("✅ 升级成功！");
        
        // 验证升级
        const newVersion = await stakingNew.version();
        console.log("新合约版本:", newVersion);
        
        const implementationAddress = await upgrades.erc1967.getImplementationAddress(stakingProxyAddress);
        console.log("新实现合约地址:", implementationAddress);

    } catch (error) {
        console.log("❌ 升级失败:", error.message);
        
        // 提供诊断信息
        console.log("\n========================================");
        console.log("诊断信息");
        console.log("========================================");
        
        if (error.message.includes("execution reverted")) {
            console.log("可能的原因:");
            console.log("1. 存储布局不兼容");
            console.log("2. 构造函数参数不匹配");
            console.log("3. 依赖合约地址变更");
            console.log("4. 权限问题");
        }
        
        if (error.message.includes("UNPREDICTABLE_GAS_LIMIT")) {
            console.log("Gas 估算失败，可能的原因:");
            console.log("1. 合约逻辑错误");
            console.log("2. 存储冲突");
            console.log("3. 依赖合约问题");
        }
        
        console.log("\n建议解决方案:");
        console.log("1. 检查存储布局是否兼容");
        console.log("2. 验证依赖合约地址");
        console.log("3. 检查权限设置");
        console.log("4. 考虑重新部署而不是升级");
        
        // 保存错误信息
        const errorRecord = {
            timestamp: new Date().toISOString(),
            network: "sepolia",
            proxyAddress: stakingProxyAddress,
            error: error.message,
            stack: error.stack,
            deployer: deployer.address,
            balance: ethers.utils.formatEther(balance)
        };

        const errorPath = path.join(__dirname, "../staking-upgrade-error.json");
        fs.writeFileSync(errorPath, JSON.stringify(errorRecord, null, 2));
        console.log(`\n📄 错误记录已保存: ${errorPath}`);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });
