import { ethers, upgrades } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 更新 StakingConfig 合约
 * 将10个NFT的等待时间从30天改为50天
 */

async function main() {
    console.log("\n🔄 开始更新 StakingConfig 合约...\n");

    const [deployer] = await ethers.getSigners();
    console.log("部署者地址:", deployer.address);
    const balance = await ethers.provider.getBalance(deployer.address);
    console.log("部署者余额:", ethers.utils.formatEther(balance), "ETH");
    console.log();

    // StakingConfig 合约地址
    const configAddress = "0x37196054B23Be5CB977AA3284A3A844a8fe77861";
    console.log("StakingConfig 合约地址:", configAddress);

    // 连接到现有合约
    const configContract = await ethers.getContractAt("StakingConfig", configAddress);
    
    console.log("========================================");
    console.log("检查当前配置");
    console.log("========================================");
    
    try {
        const comboThresholds = await configContract.getComboThresholds();
        const comboBonuses = await configContract.getComboBonuses();
        const comboMinDays = await configContract.getComboMinDays();
        
        console.log("当前组合阈值:", comboThresholds.map(t => t.toString()));
        console.log("当前组合加成:", comboBonuses.map(b => (Number(b)/100).toFixed(1) + "%"));
        console.log("当前等待天数:", comboMinDays.map(d => d.toString()));
        
        if (comboMinDays[2].toString() === "30") {
            console.log("⚠️  需要更新: 10个NFT等待时间仍为30天");
        } else {
            console.log("✅ 已更新: 10个NFT等待时间为", comboMinDays[2].toString(), "天");
            console.log("无需更新配置合约");
            return;
        }
        
    } catch (e) {
        console.log("❌ 无法读取当前配置:", e.message);
        return;
    }

    console.log("\n========================================");
    console.log("更新配置");
    console.log("========================================");

    try {
        // 准备新的配置
        const newThresholds = [3, 5, 10];
        const newBonuses = [500, 1000, 2000]; // 5%, 10%, 20%
        const newMinDays = [7, 15, 50]; // 7天, 15天, 50天
        
        console.log("新配置:");
        console.log("阈值:", newThresholds);
        console.log("加成:", newBonuses.map(b => (b/100).toFixed(1) + "%"));
        console.log("等待天数:", newMinDays);
        
        // 更新组合配置
        console.log("正在更新组合配置...");
        const tx = await configContract.updateComboConfig(
            newThresholds,
            newBonuses,
            newMinDays
        );
        
        console.log("交易哈希:", tx.hash);
        console.log("等待交易确认...");
        
        const receipt = await tx.wait();
        console.log("✅ 配置更新成功！");
        console.log("Gas 使用:", receipt.gasUsed.toString());
        
        // 验证更新
        console.log("\n========================================");
        console.log("验证更新结果");
        console.log("========================================");
        
        const updatedThresholds = await configContract.getComboThresholds();
        const updatedBonuses = await configContract.getComboBonuses();
        const updatedMinDays = await configContract.getComboMinDays();
        
        console.log("更新后的阈值:", updatedThresholds.map(t => t.toString()));
        console.log("更新后的加成:", updatedBonuses.map(b => (Number(b)/100).toFixed(1) + "%"));
        console.log("更新后的等待天数:", updatedMinDays.map(d => d.toString()));
        
        if (updatedMinDays[2].toString() === "50") {
            console.log("✅ 配置更新验证成功！");
        } else {
            console.log("❌ 配置更新验证失败");
        }
        
    } catch (error) {
        console.log("❌ 配置更新失败:", error.message);
        
        if (error.message.includes("Not authorized")) {
            console.log("权限问题: 请确保部署者地址有更新配置的权限");
        }
        
        if (error.message.includes("Invalid array lengths")) {
            console.log("参数错误: 数组长度不匹配");
        }
        
        return;
    }

    console.log("\n========================================");
    console.log("更新完成");
    console.log("========================================");
    console.log("✅ StakingConfig 合约已成功更新");
    console.log("✅ 10个NFT等待时间已从30天改为50天");
    console.log("✅ 现在可以升级 Staking 合约了");
    console.log();

    // 保存更新记录
    const updateRecord = {
        timestamp: new Date().toISOString(),
        network: "sepolia",
        configAddress: configAddress,
        changes: {
            "comboThresholds": [3, 5, 10],
            "comboBonuses": [500, 1000, 2000],
            "comboMinDays": [7, 15, 50]
        },
        deployer: deployer.address,
        balance: ethers.utils.formatEther(balance)
    };

    const recordPath = path.join(__dirname, "../staking-config-update-record.json");
    fs.writeFileSync(recordPath, JSON.stringify(updateRecord, null, 2));
    console.log(`📄 更新记录已保存: ${recordPath}`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });
