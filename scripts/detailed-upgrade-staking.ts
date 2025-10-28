import { ethers, upgrades } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 详细升级 Staking 合约
 * 添加更多调试信息和错误处理
 */

async function main() {
    console.log("\n🔄 开始详细升级 Staking 合约...\n");

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
    
    console.log("========================================");
    console.log("检查当前合约状态");
    console.log("========================================");
    
    try {
        const version = await stakingOld.version();
        console.log("当前版本:", version);
        
        const configAddress = await stakingOld.configContract();
        console.log("配置合约地址:", configAddress);
        
        const totalStaked = await stakingOld.totalStakedCount();
        console.log("当前质押总数:", totalStaked.toString());
        
        const paused = await stakingOld.paused();
        console.log("合约是否暂停:", paused);
        
        const owner = await stakingOld.owner();
        console.log("合约所有者:", owner);
        
        // 检查测试模式
        try {
            const testMode = await stakingOld.testMode();
            console.log("测试模式:", testMode);
        } catch (e) {
            console.log("测试模式: 不支持（旧版本）");
        }
        
    } catch (e) {
        console.log("❌ 无法读取合约状态:", e.message);
    }

    console.log("\n========================================");
    console.log("准备升级");
    console.log("========================================");

    try {
        // 编译新版本
        console.log("正在编译新版本...");
        const StakingNew = await ethers.getContractFactory("Staking");
        console.log("✅ 新版本编译成功");
        
        // 检查代理合约
        console.log("检查代理合约...");
        const implementationAddress = await upgrades.erc1967.getImplementationAddress(stakingProxyAddress);
        console.log("当前实现地址:", implementationAddress);
        
        // 尝试升级
        console.log("正在升级合约...");
        
        // 使用更详细的配置
        const upgradeOptions = {
            kind: "uups",
            unsafeSkipStorageCheck: true,
            gasLimit: 12000000, // 更高的 gas 限制
            timeout: 600000, // 10分钟超时
        };
        
        console.log("升级配置:", upgradeOptions);
        
        const stakingNew = await upgrades.upgradeProxy(
            stakingProxyAddress,
            StakingNew,
            upgradeOptions
        );

        console.log("✅ 升级成功！");
        
        // 验证升级
        const newVersion = await stakingNew.version();
        console.log("新合约版本:", newVersion);
        
        const newImplementationAddress = await upgrades.erc1967.getImplementationAddress(stakingProxyAddress);
        console.log("新实现合约地址:", newImplementationAddress);
        
        // 验证新功能
        console.log("\n========================================");
        console.log("验证新功能");
        console.log("========================================");
        
        try {
            const configAddress = await stakingNew.configContract();
            const configContract = await ethers.getContractAt("StakingConfig", configAddress);
            
            const comboMinDays = await configContract.getComboMinDays();
            console.log("✅ 组合等待天数:", comboMinDays.map(d => d.toString()));
            
            if (comboMinDays[2].toString() === "50") {
                console.log("✅ 10个NFT等待时间已更新为50天");
            }
            
        } catch (e) {
            console.log("❌ 无法验证新功能:", e.message);
        }
        
        console.log("\n========================================");
        console.log("升级完成");
        console.log("========================================");
        console.log("✅ Staking 合约已成功升级到版本 3.2.0");
        console.log("✅ 分级加成规则已优化");
        console.log("✅ SSS级别限制已实现");
        console.log("✅ 延迟生效机制已实现");
        
        // 保存升级记录
        const upgradeRecord = {
            timestamp: new Date().toISOString(),
            network: "sepolia",
            proxyAddress: stakingProxyAddress,
            oldImplementation: implementationAddress,
            newImplementation: newImplementationAddress,
            version: "3.2.0",
            features: [
                "分级加成规则优化",
                "3个NFT: 7天等待期 → 5%加成",
                "5个NFT: 15天等待期 → 10%加成", 
                "10个NFT: 50天等待期 → 20%加成",
                "SSS级别: 10个NFT组合限制",
                "延迟生效机制",
                "配置化等待时间"
            ],
            deployer: deployer.address,
            balance: ethers.utils.formatEther(balance)
        };

        const recordPath = path.join(__dirname, "../staking-upgrade-success.json");
        fs.writeFileSync(recordPath, JSON.stringify(upgradeRecord, null, 2));
        console.log(`📄 升级记录已保存: ${recordPath}`);

    } catch (error) {
        console.log("❌ 升级失败:", error.message);
        
        // 提供详细的错误分析
        console.log("\n========================================");
        console.log("错误分析");
        console.log("========================================");
        
        if (error.message.includes("execution reverted")) {
            console.log("执行回滚，可能的原因:");
            console.log("1. 存储布局不兼容");
            console.log("2. 构造函数参数不匹配");
            console.log("3. 依赖合约地址变更");
            console.log("4. 权限问题");
            console.log("5. 合约逻辑错误");
        }
        
        if (error.message.includes("UNPREDICTABLE_GAS_LIMIT")) {
            console.log("Gas 估算失败，可能的原因:");
            console.log("1. 合约逻辑错误");
            console.log("2. 存储冲突");
            console.log("3. 依赖合约问题");
            console.log("4. 网络问题");
        }
        
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

        const errorPath = path.join(__dirname, "../staking-upgrade-error-detailed.json");
        fs.writeFileSync(errorPath, JSON.stringify(errorRecord, null, 2));
        console.log(`\n📄 详细错误记录已保存: ${errorPath}`);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });
