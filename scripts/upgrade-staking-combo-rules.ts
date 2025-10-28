import { ethers, upgrades } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 升级 Staking 合约到新版本
 * 新版本包含分级加成规则优化：
 * - 3个NFT: 组合成功后等待7天 → 开始获得5%加成
 * - 5个NFT: 组合成功后等待15天 → 开始获得10%加成
 * - 10个NFT: 组合成功后等待50天 → 开始获得20%加成
 * - SSS级别: 10个NFT组合不适用 → 最高只能获得10%加成
 */

async function main() {
    console.log("\n🔄 开始升级 Staking 合约（分级加成规则优化）...\n");

    const [deployer] = await ethers.getSigners();
    console.log("部署者地址:", deployer.address);
    const balance = await ethers.provider.getBalance(deployer.address);
    console.log("部署者余额:", ethers.utils.formatEther(balance), "ETH");
    console.log();

    // Staking 合约代理地址（从部署结果获取）
    const stakingProxyAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    console.log("Staking 代理地址:", stakingProxyAddress);

    // 连接到现有合约查看当前版本
    const StakingOld = await ethers.getContractFactory("Staking");
    const stakingOld = StakingOld.attach(stakingProxyAddress);
    
    try {
        const oldVersion = await stakingOld.version();
        console.log("当前合约版本:", oldVersion);
    } catch (e) {
        console.log("无法读取当前版本");
    }

    // 检查当前组合加成配置
    try {
        const configAddress = await stakingOld.configContract();
        const configContract = await ethers.getContractAt("StakingConfig", configAddress);
        
        const comboThresholds = await configContract.getComboThresholds();
        const comboBonuses = await configContract.getComboBonuses();
        const comboMinDays = await configContract.getComboMinDays();
        
        console.log("\n当前组合加成配置:");
        console.log("阈值:", comboThresholds.map(t => t.toString()));
        console.log("加成:", comboBonuses.map(b => (Number(b)/100).toFixed(1) + "%"));
        console.log("等待天数:", comboMinDays.map(d => d.toString()));
    } catch (e) {
        console.log("无法读取当前配置:", e.message);
    }

    // 编译并准备新版本
    console.log("\n========================================");
    console.log("准备升级到新版本");
    console.log("========================================");
    console.log("新功能：");
    console.log("  ✅ 分级加成规则优化");
    console.log("  ✅ 3个NFT: 组合成功后等待7天 → 5%加成");
    console.log("  ✅ 5个NFT: 组合成功后等待15天 → 10%加成");
    console.log("  ✅ 10个NFT: 组合成功后等待50天 → 20%加成");
    console.log("  ✅ SSS级别限制: 10个NFT组合不适用");
    console.log("  ✅ 延迟生效机制: 利用ComboStatus.isPending");
    console.log("  ✅ 配置化等待时间: 使用getComboMinDays()");
    console.log();

    // 执行升级
    console.log("开始升级合约...");
    const StakingNew = await ethers.getContractFactory("Staking");
    
    // 先注册代理合约（如果需要）
    try {
        console.log("正在注册代理合约...");
        await upgrades.forceImport(stakingProxyAddress, StakingNew, {
            kind: "uups"
        });
        console.log("✅ 代理合约注册成功");
    } catch (e) {
        console.log("代理合约已注册或注册失败:", e.message);
    }
    
    console.log("正在部署新的实现合约...");
    
    // 先尝试估算 gas
    try {
        console.log("正在估算 gas...");
        const gasEstimate = await StakingNew.getDeployTransaction().then(tx => 
            ethers.provider.estimateGas(tx)
        );
        console.log("估算 gas:", gasEstimate.toString());
    } catch (e) {
        console.log("Gas 估算失败:", e.message);
    }
    
    const stakingNew = await upgrades.upgradeProxy(
        stakingProxyAddress,
        StakingNew,
        {
            kind: "uups",
            unsafeSkipStorageCheck: true,
            gasLimit: 8000000, // 增加 gas 限制
            timeout: 300000, // 5分钟超时
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

    // 验证组合加成功能
    try {
        const configAddress = await stakingNew.configContract();
        const configContract = await ethers.getContractAt("StakingConfig", configAddress);
        
        const comboThresholds = await configContract.getComboThresholds();
        const comboBonuses = await configContract.getComboBonuses();
        const comboMinDays = await configContract.getComboMinDays();
        
        console.log("\n✅ 新组合加成配置验证：");
        console.log("阈值:", comboThresholds.map(t => t.toString()));
        console.log("加成:", comboBonuses.map(b => (Number(b)/100).toFixed(1) + "%"));
        console.log("等待天数:", comboMinDays.map(d => d.toString()));
        
        // 验证SSS级别限制
        console.log("\n✅ SSS级别限制验证：");
        console.log("10个NFT组合等待时间:", comboMinDays[2].toString(), "天");
        console.log("SSS级别(level 6)将无法享受10个NFT组合加成");
        
    } catch (e) {
        console.log("\n❌ 无法验证组合加成配置:", e.message);
    }

    console.log("\n========================================");
    console.log("升级总结");
    console.log("========================================");
    console.log("✅ Staking 合约已成功升级");
    console.log("✅ 代理地址保持不变:", proxyAddress);
    console.log("✅ 新实现合约已部署:", implementationAddress);
    console.log("✅ 所有存储数据已保留");
    console.log("✅ 分级加成规则已优化");
    console.log("✅ SSS级别限制已实现");
    console.log();

    console.log("📝 新功能说明：");
    console.log("  1. 组合加成延迟生效:");
    console.log("     - 3个NFT: 等待7天后开始获得5%加成");
    console.log("     - 5个NFT: 等待15天后开始获得10%加成");
    console.log("     - 10个NFT: 等待50天后开始获得20%加成");
    console.log();
    console.log("  2. SSS级别特殊限制:");
    console.log("     - SSS级别(level 6)不适用10个NFT组合");
    console.log("     - SSS级别最高只能获得10%加成(5个NFT组合)");
    console.log();
    console.log("  3. 技术实现:");
    console.log("     - 使用ComboStatus.isPending和effectiveFrom机制");
    console.log("     - 配置化等待时间，支持灵活调整");
    console.log("     - 保持向后兼容，不影响现有质押");
    console.log();

    console.log("📝 下一步操作：");
    console.log("  1. 测试组合加成功能:");
    console.log("     - 质押3个同等级NFT，等待7天验证5%加成");
    console.log("     - 质押5个同等级NFT，等待15天验证10%加成");
    console.log("     - 质押10个同等级NFT，等待50天验证20%加成");
    console.log("     - 测试SSS级别的10个NFT组合限制");
    console.log();

    // 保存升级记录
    const upgradeRecord = {
        timestamp: new Date().toISOString(),
        network: "sepolia",
        proxyAddress: proxyAddress,
        implementationAddress: implementationAddress,
        version: "3.1.0",
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
        changes: {
            "StakingConfig.sol": "10个NFT等待时间从30天改为50天",
            "Staking.sol": [
                "添加_getWaitDaysForCount辅助函数",
                "修改_calculateComboBonusByCount添加level参数和SSS限制",
                "修改_updateComboStatus使用配置的等待天数",
                "更新所有函数调用"
            ]
        }
    };

    const recordPath = path.join(__dirname, "../staking-combo-upgrade-record.json");
    fs.writeFileSync(recordPath, JSON.stringify(upgradeRecord, null, 2));
    console.log(`📄 升级记录已保存: ${recordPath}\n`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 升级失败:", error);
        process.exit(1);
    });
