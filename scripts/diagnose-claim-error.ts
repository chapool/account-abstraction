import { ethers } from "hardhat";

async function main() {
    console.log("\n🔬 诊断领取奖励错误\n");

    const [deployer] = await ethers.getSigners();
    console.log("部署者地址:", deployer.address);

    // Staking 合约地址
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    
    // 要测试的 NFT IDs
    const tokenIds = ['4812', '3416', '3393', '3378', '3372', '3366'];
    
    // 连接到 Staking 合约
    const staking = await ethers.getContractAt("Staking", stakingAddress);
    
    // 查询时间状态
    const testMode = await staking.testMode();
    const testTimestamp = await staking.testTimestamp();
    const currentTimestamp = await staking.getCurrentTimestamp();
    
    console.log("========================================");
    console.log("合约状态");
    console.log("========================================");
    console.log("测试模式:", testMode);
    console.log("testTimestamp:", testTimestamp.toString());
    console.log("currentTimestamp:", currentTimestamp.toString());
    console.log();
    
    // 详细检查每个 NFT
    console.log("========================================");
    console.log("NFT 详细状态");
    console.log("========================================");
    
    for (const tokenIdStr of tokenIds) {
        const tokenId = ethers.BigNumber.from(tokenIdStr);
        console.log(`\n--- NFT #${tokenIdStr} ---`);
        
        try {
            // 获取质押信息
            const stakeInfo = await staking.stakes(tokenId);
            
            console.log("✅ 质押信息:");
            console.log("  所有者:", stakeInfo.owner);
            console.log("  等级:", stakeInfo.level.toString());
            console.log("  是否活跃:", stakeInfo.isActive);
            
            const stakeTime = stakeInfo.stakeTime;
            const lastClaimTime = stakeInfo.lastClaimTime;
            
            console.log("  质押时间戳:", stakeTime.toString());
            console.log("  质押日期:", new Date(Number(stakeTime) * 1000).toISOString());
            console.log("  最后领取时间戳:", lastClaimTime.toString());
            console.log("  最后领取日期:", new Date(Number(lastClaimTime) * 1000).toISOString());
            
            // 计算时间差
            const stakingDuration = currentTimestamp.sub(stakeTime);
            const stakingDays = stakingDuration.div(86400);
            const lastClaimDays = currentTimestamp.sub(lastClaimTime).div(86400);
            
            console.log("  已质押天数:", stakingDays.toString());
            console.log("  距离最后领取天数:", lastClaimDays.toString());
            
            // 计算奖励
            console.log("\n📊 奖励计算:");
            
            try {
                const pendingRewards = await staking.calculatePendingRewards(tokenId);
                console.log("  待领取奖励:", ethers.utils.formatEther(pendingRewards), "CPOP");
                
                // 如果奖励为 0，分析原因
                if (pendingRewards.isZero()) {
                    console.log("  ⚠️ 奖励为 0，分析原因...");
                    
                    // 检查距离最后领取是否超过 1 天
                    const hoursSinceLastClaim = currentTimestamp.sub(lastClaimTime).div(3600);
                    console.log("  距离最后领取小时数:", hoursSinceLastClaim.toString());
                    
                    if (hoursSinceLastClaim.lt(24)) {
                        console.log("  ⚠️ 距离最后领取不到 24 小时，可能无法领取");
                    }
                }
                
            } catch (e: any) {
                console.log("  ❌ 计算奖励失败:", e.message);
                console.log("  错误详情:", JSON.stringify(e, null, 2));
            }
            
            // 检查合约状态
            console.log("\n🔍 合约状态检查:");
            
            try {
                const isPaused = await staking.paused();
                console.log("  合约是否暂停:", isPaused ? "是" : "否");
                
                if (isPaused) {
                    console.log("  ❌ 合约已暂停，无法领取奖励");
                }
            } catch (e: any) {
                console.log("  ⚠️ 无法检查暂停状态:", e.message);
            }
            
            // 检查所有者权限
            console.log("\n🔑 权限检查:");
            console.log("  NFT 所有者:", stakeInfo.owner);
            console.log("  部署者地址:", deployer.address);
            console.log("  是否匹配:", stakeInfo.owner.toLowerCase() === deployer.address.toLowerCase() ? "✅ 是" : "❌ 否");
            
            if (stakeInfo.owner.toLowerCase() !== deployer.address.toLowerCase()) {
                console.log("  ⚠️ 部署者不是 NFT 所有者，无法领取");
                console.log("  💡 使用正确的所有者账户才能领取");
            }
            
            // 尝试模拟交易（不实际发送）
            console.log("\n🧪 模拟交易:");
            try {
                const estimatedGas = await staking.estimateGas.claimRewards(tokenId);
                console.log("  估算 Gas:", estimatedGas.toString());
                console.log("  状态: 可以调用");
            } catch (e: any) {
                console.log("  ❌ 模拟失败:", e.message);
                console.log("  错误原因:", e.reason);
                if (e.data) {
                    console.log("  错误数据:", e.data);
                }
            }
            
        } catch (e: any) {
            console.log("❌ 查询失败:", e.message);
            console.log("错误详情:", JSON.stringify(e, null, 2));
        }
    }
    
    console.log("\n========================================");
    console.log("总结");
    console.log("========================================");
    console.log(`
诊断要点:

1. 时间状态:
   - 测试模式已启用
   - currentTimestamp: ${currentTimestamp.toString()} (${new Date(Number(currentTimestamp) * 1000).toISOString()})
   
2. NFT 质押时间都是: 1769079240 (2026-01-22)
   
3. 已质押约 1434 天 (约 3.9 年)
   
4. 所有 NFT 都有可观的可领取奖励

可能的错误原因:
- ❌ 不是 NFT 所有者
- ❌ 合约已暂停
- ❌ Gas 不足
- ❌ 奖励为 0（距离最后领取不到 24 小时）
- ❌ 网络问题
- ❌ 合约逻辑错误
    `);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });

