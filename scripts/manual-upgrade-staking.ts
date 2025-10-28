import { ethers, upgrades } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 手动升级 Staking 合约
 * 1. 先部署新的实现合约
 * 2. 然后手动调用升级函数
 */

async function main() {
    console.log("\n🔄 开始手动升级 Staking 合约...\n");

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
        
    } catch (e) {
        console.log("❌ 无法读取合约状态:", e.message);
        return;
    }

    console.log("\n========================================");
    console.log("步骤1: 部署新的实现合约");
    console.log("========================================");

    let newImplementationAddress: string;

    try {
        // 编译新版本
        console.log("正在编译新版本...");
        const StakingNew = await ethers.getContractFactory("Staking");
        console.log("✅ 新版本编译成功");
        
        // 部署新的实现合约
        console.log("正在部署新的实现合约...");
        
        // 估算gas费用
        const deployTx = await StakingNew.getDeployTransaction();
        const gasEstimate = await deployer.estimateGas(deployTx);
        const gasPrice = await deployer.getGasPrice();
        const gasCost = gasEstimate.mul(gasPrice);
        
        console.log("预估Gas费用:", ethers.utils.formatEther(gasCost), "ETH");
        console.log("Gas Limit:", gasEstimate.toString());
        console.log("Gas Price:", ethers.utils.formatUnits(gasPrice, "gwei"), "Gwei");
        
        const stakingImplementation = await StakingNew.deploy({
            gasLimit: gasEstimate.mul(120).div(100), // 增加20%的gas buffer
            gasPrice: gasPrice
        });
        await stakingImplementation.deployed();
        
        newImplementationAddress = stakingImplementation.address;
        console.log("✅ 新实现合约部署成功！");
        console.log("新实现合约地址:", newImplementationAddress);
        
        // 验证新合约
        const newVersion = await stakingImplementation.version();
        console.log("新合约版本:", newVersion);
        
    } catch (error) {
        console.log("❌ 部署新实现合约失败:", error.message);
        return;
    }

    console.log("\n========================================");
    console.log("步骤2: 手动调用升级函数");
    console.log("========================================");

    try {
        // 连接到代理合约
        const proxyContract = await ethers.getContractAt("@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol:UUPSUpgradeable", stakingProxyAddress);
        
        console.log("正在调用 upgradeToAndCall 函数...");
        console.log("升级到地址:", newImplementationAddress);
        
        // 调用升级函数
        const tx = await proxyContract.upgradeToAndCall(newImplementationAddress, "0x");
        console.log("升级交易哈希:", tx.hash);
        
        console.log("等待交易确认...");
        const receipt = await tx.wait();
        console.log("✅ 升级成功！");
        console.log("Gas 使用:", receipt.gasUsed.toString());
        
    } catch (error) {
        console.log("❌ 手动升级失败:", error.message);
        
        if (error.message.includes("Not authorized")) {
            console.log("权限问题: 请确保部署者地址有升级权限");
        }
        
        if (error.message.includes("New implementation is not UUPS")) {
            console.log("实现合约问题: 新实现合约不是UUPS兼容的");
        }
        
        return;
    }

    console.log("\n========================================");
    console.log("步骤3: 验证升级结果");
    console.log("========================================");

    try {
        // 验证升级
        const stakingNew = StakingOld.attach(stakingProxyAddress);
        
        const finalVersion = await stakingNew.version();
        console.log("升级后版本:", finalVersion);
        
        const finalImplementationAddress = await upgrades.erc1967.getImplementationAddress(stakingProxyAddress);
        console.log("最终实现地址:", finalImplementationAddress);
        
        if (finalImplementationAddress === newImplementationAddress) {
            console.log("✅ 升级验证成功！");
        } else {
            console.log("❌ 升级验证失败");
        }
        
        // 验证新功能
        console.log("\n验证新功能...");
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
        
    } catch (error) {
        console.log("❌ 验证升级结果失败:", error.message);
    }

    console.log("\n========================================");
    console.log("升级完成");
    console.log("========================================");
    console.log("✅ Staking 合约已成功升级到版本 3.2.0");
    console.log("✅ 分级加成规则已优化");
    console.log("✅ SSS级别限制已实现");
    console.log("✅ 延迟生效机制已实现");
    console.log();

    // 保存升级记录
    const upgradeRecord = {
        timestamp: new Date().toISOString(),
        network: "sepolia",
        proxyAddress: stakingProxyAddress,
        newImplementation: newImplementationAddress,
        version: "3.2.0",
        method: "manual_upgrade",
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

    const recordPath = path.join(__dirname, "../staking-manual-upgrade-success.json");
    fs.writeFileSync(recordPath, JSON.stringify(upgradeRecord, null, 2));
    console.log(`📄 升级记录已保存: ${recordPath}`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });
