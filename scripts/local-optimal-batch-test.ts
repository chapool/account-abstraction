import { ethers } from "hardhat";

async function main() {
    console.log("🎯 CPOPToken本地最优批处理规模测试");
    console.log("=".repeat(50));
    
    const [deployer] = await ethers.getSigners();
    const deployerAddress = await deployer.getAddress();
    
    console.log("🔑 测试账户:", deployerAddress);
    console.log("💰 余额:", ethers.utils.formatEther(await deployer.getBalance()), "ETH");
    
    // 部署CPOPToken到本地网络
    console.log("\n📝 部署CPOPToken到本地网络...");
    const initialSupply = ethers.utils.parseEther("1000000");
    
    const CPOPTokenFactory = await ethers.getContractFactory("contracts/cpop/CPOPToken.sol:CPOPToken");
    const cpopToken = await CPOPTokenFactory.deploy(deployerAddress, initialSupply);
    await cpopToken.deployed();
    
    console.log("✅ CPOPToken部署成功:", cpopToken.address);
    
    // 测试不同的批量规模
    const batchSizes = [1, 2, 3, 5, 8, 10, 15, 20, 25, 30, 40, 50, 75, 100];
    const testResults: Array<{
        size: number;
        totalGas: number;
        gasPerOperation: number;
        efficiency: number;
        costSavings: string;
    }> = [];
    
    // 单次操作基准测试
    console.log("\n📊 步骤1: 单次操作基准测试");
    const singleMintTx = await cpopToken.mint("0x1111111111111111111111111111111111111111", ethers.utils.parseEther("1"));
    const singleMintReceipt = await singleMintTx.wait();
    const singleGas = singleMintReceipt.gasUsed.toNumber();
    console.log("✅ 单次mint基准:", singleGas, "gas");
    
    console.log("\n🧪 步骤2: 批量规模效率测试");
    
    for (const size of batchSizes) {
        try {
            console.log(`\n测试批量规模: ${size}`);
            
            // 生成测试地址和金额
            const recipients: string[] = [];
            const amounts: ethers.BigNumber[] = [];
            
            for (let i = 0; i < size; i++) {
                // 生成唯一的测试地址
                const testAddress = ethers.utils.getAddress(
                    `0x${(i + 1).toString(16).padStart(40, '0')}`
                );
                recipients.push(testAddress);
                amounts.push(ethers.utils.parseEther("1"));
            }
            
            // 执行批量操作
            const batchTx = await cpopToken.batchMint(recipients, amounts);
            const batchReceipt = await batchTx.wait();
            const batchGas = batchReceipt.gasUsed.toNumber();
            
            // 计算效率指标
            const gasPerOperation = Math.round(batchGas / size);
            const theoreticalGas = singleGas * size;
            const efficiency = ((theoreticalGas - batchGas) / theoreticalGas * 100);
            const costSavings = efficiency > 0 ? `${efficiency.toFixed(2)}%` : `-${Math.abs(efficiency).toFixed(2)}%`;
            
            testResults.push({
                size,
                totalGas: batchGas,
                gasPerOperation,
                efficiency,
                costSavings
            });
            
            console.log(`  总Gas: ${batchGas.toLocaleString()}`);
            console.log(`  平均Gas/操作: ${gasPerOperation.toLocaleString()}`);
            console.log(`  效率提升: ${costSavings}`);
            console.log(`  vs单次操作: ${(batchGas / theoreticalGas).toFixed(3)}x`);
            
        } catch (error) {
            console.error(`❌ 批量规模 ${size} 测试失败:`, (error as Error).message);
            break;
        }
    }
    
    // 分析结果
    console.log("\n📈 步骤3: 结果分析");
    console.log("=".repeat(90));
    console.log("批量大小 | 总Gas     | 平均Gas/操作 | 效率提升  | 成本节省  | 相对成本");
    console.log("-".repeat(90));
    
    let maxEfficiency = -Infinity;
    let optimalSize = 1;
    let bestCostPerOperation = Infinity;
    let mostEfficientSize = 1;
    
    for (const result of testResults) {
        const relativeCost = result.totalGas / (singleGas * result.size);
        console.log(
            `${result.size.toString().padStart(8)} | ` +
            `${result.totalGas.toLocaleString().padStart(9)} | ` +
            `${result.gasPerOperation.toLocaleString().padStart(12)} | ` +
            `${result.efficiency.toFixed(2)}%`.padStart(9) + ` | ` +
            `${result.costSavings.padStart(8)} | ` +
            `${relativeCost.toFixed(3)}x`.padStart(8)
        );
        
        // 找到最高效率
        if (result.efficiency > maxEfficiency) {
            maxEfficiency = result.efficiency;
            mostEfficientSize = result.size;
        }
        
        // 找到最低平均成本
        if (result.gasPerOperation < bestCostPerOperation) {
            bestCostPerOperation = result.gasPerOperation;
            optimalSize = result.size;
        }
    }
    
    console.log("=".repeat(90));
    
    // 找到效率拐点
    let diminishingReturnsPoint = 1;
    for (let i = 1; i < testResults.length; i++) {
        const current = testResults[i];
        const previous = testResults[i - 1];
        
        const efficiencyIncrease = current.efficiency - previous.efficiency;
        
        // 如果效率提升小于1%，认为是收益递减点
        if (efficiencyIncrease < 1 && current.size >= 10) {
            diminishingReturnsPoint = previous.size;
            break;
        }
    }
    
    // 寻找最佳平衡点 (效率>40%且规模合理)
    let balancedSize = 10;
    for (const result of testResults) {
        if (result.efficiency > 40 && result.size <= 50) {
            balancedSize = result.size;
        }
    }
    
    // 综合建议
    console.log("\n🎯 最优批处理规模分析:");
    console.log(`📊 最高效率规模: ${mostEfficientSize} (节省 ${maxEfficiency.toFixed(2)}%)`);
    console.log(`💰 最低平均成本: ${optimalSize} (${bestCostPerOperation.toLocaleString()} gas/操作)`);
    console.log(`⚡ 收益递减拐点: ${diminishingReturnsPoint} (效率增长开始放缓)`);
    console.log(`🎯 平衡最优规模: ${balancedSize} (效率与实用性平衡)`);
    
    // 详细效率分析
    console.log("\n📊 效率区间分析:");
    const ranges = [
        { name: "小批量 (1-5)", min: 1, max: 5 },
        { name: "中批量 (10-25)", min: 10, max: 25 },
        { name: "大批量 (30-50)", min: 30, max: 50 },
        { name: "超大批量 (50+)", min: 50, max: 100 }
    ];
    
    for (const range of ranges) {
        const rangeResults = testResults.filter(r => r.size >= range.min && r.size <= range.max);
        if (rangeResults.length > 0) {
            const avgEfficiency = rangeResults.reduce((sum, r) => sum + r.efficiency, 0) / rangeResults.length;
            const minGas = Math.min(...rangeResults.map(r => r.gasPerOperation));
            const maxGas = Math.max(...rangeResults.map(r => r.gasPerOperation));
            console.log(`${range.name}: 平均效率 ${avgEfficiency.toFixed(2)}%, Gas范围 ${minGas}-${maxGas}`);
        }
    }
    
    // 应用场景建议
    console.log("\n💡 应用场景建议:");
    console.log("🚀 日常转账 (1-3个地址): 适合小额转账、个人交易");
    console.log("🎯 活动奖励 (10-20个地址): 适合游戏奖励、积分发放");
    console.log("🌟 营销活动 (25-40个地址): 适合推广活动、用户激励");
    console.log("🏆 大型空投 (50+个地址): 适合代币分发、大规模营销");
    
    // 风险评估
    console.log("\n⚠️  风险评估:");
    console.log("• 批量1-10: 低风险，适合所有网络环境");
    console.log("• 批量10-25: 中等风险，主流网络建议规模");
    console.log("• 批量25-50: 较高风险，需要充足的gas limit");
    console.log("• 批量50+: 高风险，可能遇到区块gas limit");
    
    // 最终推荐
    console.log(`\n🌟 综合推荐批量规模: ${balancedSize}`);
    console.log(`理由: 在效率 (${testResults.find(r => r.size === balancedSize)?.efficiency.toFixed(2)}%)、安全性和实用性之间取得最佳平衡`);
    
    return {
        optimalSize,
        mostEfficientSize,
        diminishingReturnsPoint,
        balancedSize,
        maxEfficiency,
        bestCostPerOperation,
        testResults,
        singleGas,
        contractAddress: cpopToken.address
    };
}

main()
    .then((result) => {
        console.log("\n🎉 本地最优批处理规模测试完成!");
        console.log(`📈 推荐使用批量规模: ${result.balancedSize}`);
        console.log(`🏆 最大效率提升: ${result.maxEfficiency.toFixed(2)}%`);
        console.log(`💎 最优Gas成本: ${result.bestCostPerOperation.toLocaleString()} gas/操作`);
        console.log(`🔗 测试合约地址: ${result.contractAddress}`);
        
        // 生成最终建议
        console.log("\n📝 生产环境建议:");
        console.log("• 日常操作: 5-10个地址");
        console.log("• 批量奖励: 15-25个地址");
        console.log("• 大型活动: 30-50个地址");
        console.log("• 根据网络状况动态调整批量大小");
        
        process.exit(0);
    })
    .catch((error) => {
        console.error("❌ 测试失败:", error.message);
        process.exit(1);
    });