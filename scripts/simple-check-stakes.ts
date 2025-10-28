import { ethers } from "hardhat";

/**
 * 简单检查当前合约的stakes函数
 */

async function main() {
    console.log("\n🔍 简单检查stakes函数...\n");

    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";

    try {
        const stakingContract = await ethers.getContractAt("Staking", stakingAddress);
        
        // 获取用户的Token ID
        const tokenId = await stakingContract.userStakes(userAddress, 0);
        console.log(`Token ID: ${tokenId.toString()}`);
        
        // 尝试不同的调用方式
        console.log("\n尝试获取stakes信息...");
        
        try {
            // 方式1：直接调用
            const result = await stakingContract.stakes(tokenId);
            console.log("✅ 成功获取stakes信息");
            console.log(`返回结果长度: ${result.length}`);
            
            if (result.length === 8) {
                console.log("📋 当前合约StakeInfo有8个字段:");
                console.log(`  0. owner: ${result[0]}`);
                console.log(`  1. tokenId: ${result[1].toString()}`);
                console.log(`  2. level: ${result[2]}`);
                console.log(`  3. stakeTime: ${result[3].toString()}`);
                console.log(`  4. lastClaimTime: ${result[4].toString()}`);
                console.log(`  5. isActive: ${result[5]}`);
                console.log(`  6. totalRewards: ${ethers.utils.formatEther(result[6])} CPOP`);
                console.log(`  7. pendingRewards: ${ethers.utils.formatEther(result[7])} CPOP`);
                console.log("\n❌ 缺少continuousBonusClaimed字段");
            } else if (result.length === 9) {
                console.log("📋 当前合约StakeInfo有9个字段:");
                console.log(`  0. owner: ${result[0]}`);
                console.log(`  1. tokenId: ${result[1].toString()}`);
                console.log(`  2. level: ${result[2]}`);
                console.log(`  3. stakeTime: ${result[3].toString()}`);
                console.log(`  4. lastClaimTime: ${result[4].toString()}`);
                console.log(`  5. isActive: ${result[5]}`);
                console.log(`  6. totalRewards: ${ethers.utils.formatEther(result[6])} CPOP`);
                console.log(`  7. pendingRewards: ${ethers.utils.formatEther(result[7])} CPOP`);
                console.log(`  8. continuousBonusClaimed: ${result[8]}`);
                console.log("\n✅ 已有continuousBonusClaimed字段");
            } else {
                console.log(`❓ 意外的字段数量: ${result.length}`);
            }
            
        } catch (e) {
            console.log(`❌ 调用失败: ${e.message}`);
            
            // 尝试使用call方法
            try {
                console.log("\n尝试使用call方法...");
                const data = stakingContract.interface.encodeFunctionData("stakes", [tokenId]);
                const result = await stakingContract.provider.call({
                    to: stakingAddress,
                    data: data
                });
                
                const decoded = stakingContract.interface.decodeFunctionResult("stakes", result);
                console.log("✅ call方法成功");
                console.log(`解码结果长度: ${decoded.length}`);
                
            } catch (e2) {
                console.log(`❌ call方法也失败: ${e2.message}`);
            }
        }
        
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
