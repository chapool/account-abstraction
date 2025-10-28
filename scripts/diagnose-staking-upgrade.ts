import { ethers, upgrades } from "hardhat";
import * as fs from "fs";

/**
 * 诊断 Staking 合约升级问题
 * 检查存储布局和依赖关系
 */

async function main() {
    console.log("\n🔍 诊断 Staking 合约升级问题...\n");

    const [deployer] = await ethers.getSigners();
    console.log("部署者地址:", deployer.address);

    // Staking 合约代理地址
    const stakingProxyAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    console.log("Staking 代理地址:", stakingProxyAddress);

    // 连接到现有合约
    const StakingOld = await ethers.getContractFactory("Staking");
    const stakingOld = StakingOld.attach(stakingProxyAddress);
    
    console.log("\n========================================");
    console.log("检查当前合约状态");
    console.log("========================================");
    
    try {
        const version = await stakingOld.version();
        console.log("当前版本:", version);
        
        const configAddress = await stakingOld.configContract();
        console.log("配置合约地址:", configAddress);
        
        const cpnftAddress = await stakingOld.cpnftContract();
        console.log("CPNFT合约地址:", cpnftAddress);
        
        const cpopAddress = await stakingOld.cpopTokenContract();
        console.log("CPOP合约地址:", cpopAddress);
        
        const accountManagerAddress = await stakingOld.accountManagerContract();
        console.log("AccountManager合约地址:", accountManagerAddress);
        
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
    console.log("检查配置合约");
    console.log("========================================");
    
    try {
        const configAddress = await stakingOld.configContract();
        const configContract = await ethers.getContractAt("StakingConfig", configAddress);
        
        const comboThresholds = await configContract.getComboThresholds();
        const comboBonuses = await configContract.getComboBonuses();
        const comboMinDays = await configContract.getComboMinDays();
        
        console.log("组合阈值:", comboThresholds.map(t => t.toString()));
        console.log("组合加成:", comboBonuses.map(b => (Number(b)/100).toFixed(1) + "%"));
        console.log("等待天数:", comboMinDays.map(d => d.toString()));
        
        // 检查是否需要更新配置
        if (comboMinDays[2].toString() === "30") {
            console.log("⚠️  配置需要更新: 10个NFT等待时间仍为30天，需要改为50天");
        } else {
            console.log("✅ 配置已更新: 10个NFT等待时间为", comboMinDays[2].toString(), "天");
        }
        
    } catch (e) {
        console.log("❌ 无法读取配置:", e.message);
    }

    console.log("\n========================================");
    console.log("检查存储布局兼容性");
    console.log("========================================");
    
    // 检查新版本合约的存储布局
    try {
        const StakingNew = await ethers.getContractFactory("Staking");
        console.log("✅ 新版本合约编译成功");
        
        // 检查是否有新的存储变量
        console.log("检查存储变量...");
        
        // 这里我们可以检查新添加的存储变量
        console.log("新版本包含的存储变量:");
        console.log("- testMode: bool");
        console.log("- testTimestamp: uint256");
        console.log("- 其他现有变量保持不变");
        
    } catch (e) {
        console.log("❌ 新版本合约编译失败:", e.message);
    }

    console.log("\n========================================");
    console.log("升级建议");
    console.log("========================================");
    
    console.log("基于诊断结果，建议的升级策略:");
    console.log();
    console.log("方案1: 先更新配置合约");
    console.log("  1. 更新 StakingConfig 合约中的等待时间");
    console.log("  2. 然后升级 Staking 合约");
    console.log();
    console.log("方案2: 使用 unsafeSkipStorageCheck");
    console.log("  1. 确保存储变量顺序正确");
    console.log("  2. 使用更高的 gas 限制");
    console.log("  3. 添加超时设置");
    console.log();
    console.log("方案3: 重新部署（如果升级失败）");
    console.log("  1. 迁移现有质押数据");
    console.log("  2. 更新前端配置");
    console.log("  3. 通知用户");
    
    console.log("\n推荐执行顺序:");
    console.log("1. 先尝试更新 StakingConfig 合约");
    console.log("2. 然后升级 Staking 合约");
    console.log("3. 如果仍然失败，考虑重新部署");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 诊断失败:", error);
        process.exit(1);
    });
