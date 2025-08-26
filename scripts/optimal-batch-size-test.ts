import { ethers } from "hardhat";

async function main() {
    console.log("🎯 CPOPToken最优批处理规模测试");
    console.log("=".repeat(50));
    
    // 连接到BSC测试网的CPOPToken
    const cpopTokenAddress = "0xc2DEe82B70C60eDbF8a8E180Cb72A2e727574260";
    const cpopToken = await ethers.getContractAt("contracts/cpop/CPOPToken.sol:CPOPToken", cpopTokenAddress);
    
    const [deployer] = await ethers.getSigners();
    const deployerAddress = await deployer.getAddress();
    
    console.log("🔑 测试账户:", deployerAddress);
    console.log("💰 余额:", ethers.utils.formatEther(await deployer.getBalance()), "BNB");
    console.log("📝 CPOPToken:", cpopTokenAddress);
    
    // 测试不同的批量规模
    const batchSizes = [1, 2, 3, 5, 8, 10, 15, 20, 25, 30, 50, 100];
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
            console.log(`  效率: ${costSavings}`);
            
            // 防止gas limit问题，如果单次操作gas过高就停止
            if (batchGas > 8000000) {
                console.log("⚠️  接近gas limit，停止更大规模测试");
                break;
            }
            
            // 小延迟避免网络拥塞
            await new Promise(resolve => setTimeout(resolve, 2000));
            
        } catch (error) {
            console.error(`❌ 批量规模 ${size} 测试失败:`, (error as Error).message);
            if (size > 20) {
                console.log("⚠️  大批量可能受到网络限制，停止测试");
                break;
            }
        }
    }
    
    // 分析结果
    console.log("\n📈 步骤3: 结果分析");
    console.log("=".repeat(80));
    console.log("批量大小 | 总Gas     | 平均Gas/操作 | 效率提升  | 成本节省");
    console.log("-".repeat(80));
    
    let maxEfficiency = -Infinity;
    let optimalSize = 1;
    let bestCostPerOperation = Infinity;
    let mostEfficientSize = 1;
    
    for (const result of testResults) {
        console.log(
            `${result.size.toString().padStart(8)} | ` +
            `${result.totalGas.toLocaleString().padStart(9)} | ` +
            `${result.gasPerOperation.toLocaleString().padStart(12)} | ` +
            `${(result.totalGas / (singleGas * result.size)).toFixed(2)}x`.padStart(9) + ` | ` +
            `${result.costSavings.padStart(8)}`
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
    
    console.log("=".repeat(80));
    
    // 找到效率拐点
    let diminishingReturnsPoint = 1;
    for (let i = 1; i < testResults.length; i++) {
        const current = testResults[i];
        const previous = testResults[i - 1];
        
        const efficiencyIncrease = current.efficiency - previous.efficiency;
        
        // 如果效率提升小于1%，认为是收益递减点
        if (efficiencyIncrease < 1 && current.size >= 10) {
            diminishingReturnsPoint = current.size;
            break;
        }
    }
    
    // 综合建议
    console.log("\n🎯 最优批处理规模建议:");
    console.log(`📊 最高效率规模: ${mostEfficientSize} (节省 ${maxEfficiency.toFixed(2)}%)`);
    console.log(`💰 最低平均成本: ${optimalSize} (${bestCostPerOperation.toLocaleString()} gas/操作)`);
    console.log(`⚡ 收益递减拐点: ${diminishingReturnsPoint} (效率增长开始放缓)`);
    
    // 应用场景建议
    console.log("\n💡 应用场景建议:");
    console.log("🚀 小规模操作 (1-5个地址): 适合日常转账、小额奖励");
    console.log("🎯 中规模操作 (10-20个地址): 适合活动奖励、积分发放");
    console.log("🌟 大规模操作 (50+个地址): 适合空投、大型活动");
    
    // 风险提示
    console.log("\n⚠️  注意事项:");
    console.log("• 批量过大可能遇到gas limit限制");
    console.log("• 网络拥堵时大批量可能失败");
    console.log("• 建议根据实际gas limit调整批量大小");
    
    // 最终推荐
    const recommendedSize = Math.min(optimalSize, 25); // 不超过25，平衡效率和风险
    console.log(`\n🌟 综合推荐批量规模: ${recommendedSize}`);
    console.log(`理由: 在效率、安全性和实用性之间取得最佳平衡`);
    
    return {
        optimalSize,
        mostEfficientSize,
        diminishingReturnsPoint,
        recommendedSize,
        maxEfficiency,
        bestCostPerOperation,
        testResults
    };
}

main()
    .then((result) => {
        console.log("\n🎉 最优批处理规模测试完成!");
        console.log(`📈 推荐使用批量规模: ${result.recommendedSize}`);
        console.log(`🏆 最大效率提升: ${result.maxEfficiency.toFixed(2)}%`);
        console.log(`💎 最优Gas成本: ${result.bestCostPerOperation.toLocaleString()} gas/操作`);
        process.exit(0);
    })
    .catch((error) => {
        console.error("❌ 测试失败:", error.message);
        process.exit(1);
    });