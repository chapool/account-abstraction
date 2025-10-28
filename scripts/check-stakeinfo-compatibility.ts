import { ethers } from "hardhat";

/**
 * 检查当前部署合约的StakeInfo结构
 * 验证升级兼容性
 */

async function main() {
    console.log("\n🔍 检查StakeInfo结构升级兼容性...\n");

    const [deployer] = await ethers.getSigners();
    console.log("检查者地址:", deployer.address);

    // 合约地址
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";

    try {
        // 连接到当前部署的合约
        const stakingContract = await ethers.getContractAt("Staking", stakingAddress);
        
        console.log("========================================");
        console.log("检查当前合约状态");
        console.log("========================================");
        
        const version = await stakingContract.version();
        console.log(`当前合约版本: ${version}`);
        
        // 获取用户的Token ID
        const tokenId = await stakingContract.userStakes(userAddress, 0);
        console.log(`用户Token ID: ${tokenId.toString()}`);
        
        console.log("\n========================================");
        console.log("检查StakeInfo结构");
        console.log("========================================");
        
        try {
            // 尝试获取StakeInfo（使用8个字段的解构）
            const stakeInfo8 = await stakingContract.stakes(tokenId);
            console.log("✅ 当前合约支持8个字段的StakeInfo结构:");
            console.log(`  所有者: ${stakeInfo8.owner}`);
            console.log(`  Token ID: ${stakeInfo8.tokenId.toString()}`);
            console.log(`  等级: ${stakeInfo8.level}`);
            console.log(`  质押时间: ${new Date(Number(stakeInfo8.stakeTime) * 1000).toLocaleString()}`);
            console.log(`  最后领取时间: ${new Date(Number(stakeInfo8.lastClaimTime) * 1000).toLocaleString()}`);
            console.log(`  是否活跃: ${stakeInfo8.isActive}`);
            console.log(`  总奖励: ${ethers.utils.formatEther(stakeInfo8.totalRewards)} CPOP`);
            console.log(`  待领取奖励: ${ethers.utils.formatEther(stakeInfo8.pendingRewards)} CPOP`);
            
            // 检查是否有第9个字段
            try {
                // 尝试访问第9个字段（continuousBonusClaimed）
                const stakeInfo9 = await stakingContract.stakes(tokenId);
                console.log(`\n🔍 检查第9个字段 (continuousBonusClaimed):`);
                
                // 如果合约支持9个字段，stakeInfo9应该包含continuousBonusClaimed
                if (stakeInfo9.length >= 9) {
                    console.log(`✅ 当前合约已支持continuousBonusClaimed字段: ${stakeInfo9[8]}`);
                } else {
                    console.log(`❌ 当前合约不支持continuousBonusClaimed字段`);
                    console.log(`   字段数量: ${stakeInfo9.length}`);
                }
                
            } catch (e) {
                console.log(`❌ 无法访问continuousBonusClaimed字段: ${e.message}`);
            }
            
        } catch (e) {
            console.log(`❌ 无法获取StakeInfo: ${e.message}`);
        }
        
        console.log("\n========================================");
        console.log("升级兼容性分析");
        console.log("========================================");
        
        console.log("📋 升级兼容性检查:");
        console.log("1. 当前合约版本: 3.2.0");
        console.log("2. 目标合约版本: 3.3.0");
        console.log("3. StakeInfo结构变化: 添加了continuousBonusClaimed字段");
        
        console.log("\n🔍 可能的升级问题:");
        console.log("❌ 如果当前合约的StakeInfo只有8个字段");
        console.log("   - 升级后新字段会初始化为默认值(false)");
        console.log("   - 现有质押记录的continuousBonusClaimed将为false");
        console.log("   - 可能导致连续质押奖励重复发放");
        
        console.log("\n✅ 如果当前合约的StakeInfo已经有9个字段");
        console.log("   - 升级完全兼容");
        console.log("   - 现有数据保持不变");
        console.log("   - 可以安全升级");
        
        console.log("\n========================================");
        console.log("建议的升级策略");
        console.log("========================================");
        
        console.log("🛡️ 安全升级方案:");
        console.log("1. 检查当前合约是否已有continuousBonusClaimed字段");
        console.log("2. 如果没有，需要数据迁移策略");
        console.log("3. 如果有，可以直接升级");
        
        console.log("\n📝 数据迁移策略（如果需要）:");
        console.log("1. 在升级前记录所有质押状态");
        console.log("2. 升级后设置continuousBonusClaimed为true（对于已领取过奖励的NFT）");
        console.log("3. 确保连续质押奖励不会重复发放");
        
    } catch (error) {
        console.log("❌ 检查失败:", error.message);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });
