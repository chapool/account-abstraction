import { ethers } from "hardhat";

async function main() {
    console.log("💰 Hashkey测试网合约部署Gas费用估算");
    console.log("=".repeat(45));
    
    const [deployer] = await ethers.getSigners();
    const deployerAddress = await deployer.getAddress();
    
    console.log("🔑 部署账户:", deployerAddress);
    const balance = await deployer.getBalance();
    console.log("💰 当前余额:", ethers.utils.formatEther(balance), "HSK");
    
    // Gas价格 (从配置中获取)
    const gasPrice = ethers.BigNumber.from("30000000000"); // 30 gwei
    console.log("⛽ Gas价格:", ethers.utils.formatUnits(gasPrice, "gwei"), "gwei");
    
    // 估算各合约部署Gas消耗
    const contractEstimates = {
        "CPOPToken": 1500000, // 已部署成功，实际可能更少
        "GasPriceOracle": 800000,
        "SessionKeyManager": 1200000,
        "AAWallet": 2000000,
        "WalletManager": 2500000,
        "MasterAggregator": 2200000
    };
    
    console.log("\n📊 各合约估算Gas消耗:");
    let totalGas = 0;
    let totalCost = ethers.BigNumber.from(0);
    
    for (const [name, gas] of Object.entries(contractEstimates)) {
        const cost = gasPrice.mul(gas);
        totalGas += gas;
        totalCost = totalCost.add(cost);
        
        const status = name === "CPOPToken" ? "✅ 已部署" : "⏳ 待部署";
        console.log(`${name}: ${gas.toLocaleString()} gas = ${ethers.utils.formatEther(cost)} HSK ${status}`);
    }
    
    console.log("\n📈 总计:");
    console.log("总Gas消耗:", totalGas.toLocaleString());
    console.log("总费用:", ethers.utils.formatEther(totalCost), "HSK");
    
    // 减去已部署的CPOPToken
    const remainingGas = totalGas - contractEstimates.CPOPToken;
    const remainingCost = totalCost.sub(gasPrice.mul(contractEstimates.CPOPToken));
    
    console.log("\n🎯 剩余合约部署需求:");
    console.log("剩余Gas:", remainingGas.toLocaleString());
    console.log("剩余费用:", ethers.utils.formatEther(remainingCost), "HSK");
    
    // 安全余量 (20%)
    const safetyMargin = remainingCost.mul(20).div(100);
    const recommendedAmount = remainingCost.add(safetyMargin);
    
    console.log("\n💡 建议准备:");
    console.log("基础需求:", ethers.utils.formatEther(remainingCost), "HSK");
    console.log("安全余量 (20%):", ethers.utils.formatEther(safetyMargin), "HSK");
    console.log("推荐总额:", ethers.utils.formatEther(recommendedAmount), "HSK");
    
    // 当前余额状态
    console.log("\n📋 当前状态:");
    console.log("当前余额:", ethers.utils.formatEther(balance), "HSK");
    
    if (balance.gte(recommendedAmount)) {
        console.log("✅ 余额充足，可以继续部署");
    } else {
        const needed = recommendedAmount.sub(balance);
        console.log("❌ 余额不足，还需要:", ethers.utils.formatEther(needed), "HSK");
    }
    
    // 分阶段部署建议
    console.log("\n🔄 分阶段部署建议:");
    const phases = [
        {
            name: "Phase 1 - 基础合约",
            contracts: ["GasPriceOracle", "SessionKeyManager"],
            gas: contractEstimates.GasPriceOracle + contractEstimates.SessionKeyManager
        },
        {
            name: "Phase 2 - 核心AA合约", 
            contracts: ["AAWallet", "WalletManager"],
            gas: contractEstimates.AAWallet + contractEstimates.WalletManager
        },
        {
            name: "Phase 3 - 聚合器",
            contracts: ["MasterAggregator"],
            gas: contractEstimates.MasterAggregator
        }
    ];
    
    for (const phase of phases) {
        const phaseCost = gasPrice.mul(phase.gas);
        console.log(`${phase.name}: ${ethers.utils.formatEther(phaseCost)} HSK`);
        console.log(`  合约: ${phase.contracts.join(", ")}`);
        console.log(`  Gas: ${phase.gas.toLocaleString()}`);
        console.log("");
    }
    
    return {
        totalCost: ethers.utils.formatEther(totalCost),
        remainingCost: ethers.utils.formatEther(remainingCost),
        recommendedAmount: ethers.utils.formatEther(recommendedAmount),
        currentBalance: ethers.utils.formatEther(balance)
    };
}

main()
    .then((result) => {
        console.log("🎉 Gas费用估算完成!");
        process.exit(0);
    })
    .catch((error) => {
        console.error("❌ 估算失败:", error.message);
        process.exit(1);
    });