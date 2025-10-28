import { ethers } from "hardhat";

async function main() {
    console.log("\n🎁 测试领取奖励功能\n");

    const [deployer] = await ethers.getSigners();
    console.log("操作者地址:", deployer.address);

    // Staking 合约地址
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    
    // 要测试的 NFT IDs
    const tokenIds = ['4812', '3416', '3393', '3378', '3372', '3366'];
    
    console.log("目标 NFT IDs:", tokenIds.join(", "));
    console.log();
    
    // 连接到 Staking 合约
    const staking = await ethers.getContractAt("Staking", stakingAddress);
    
    // 查询当前时间状态
    const testMode = await staking.testMode();
    const testTimestamp = await staking.testTimestamp();
    const currentTimestamp = await staking.getCurrentTimestamp();
    
    console.log("========================================");
    console.log("时间状态");
    console.log("========================================");
    console.log("测试模式:", testMode);
    console.log("testTimestamp:", testTimestamp.toString());
    console.log("currentTimestamp:", currentTimestamp.toString());
    console.log();
    
    // 检查每个 NFT 的状态
    console.log("========================================");
    console.log("检查 NFT 状态");
    console.log("========================================");
    
    for (const tokenId of tokenIds) {
        try {
            console.log(`\n检查 NFT #${tokenId}...`);
            
            const stakeInfo = await staking.stakes(tokenId);
            
            if (!stakeInfo.isActive) {
                console.log("  ❌ 未质押");
                continue;
            }
            
            console.log("  所有者:", stakeInfo.owner);
            console.log("  质押时间:", stakeInfo.stakeTime.toString());
            console.log("  最后领取时间:", stakeInfo.lastClaimTime.toString());
            
            // 计算质押天数（使用 currentTimestamp）
            const stakingDuration = currentTimestamp.sub(stakeInfo.stakeTime);
            const stakingDays = stakingDuration.div(86400);
            console.log("  已质押天数:", stakingDays.toString());
            
            // 计算待领取奖励
            try {
                console.log("  正在计算待领取奖励...");
                const pendingRewards = await staking.calculatePendingRewards(tokenId);
                console.log("  ✅ 待领取奖励:", ethers.utils.formatEther(pendingRewards), "CPOP");
                
                if (pendingRewards.gt(0)) {
                    console.log("  准备领取奖励...");
                    // 这里先不实际领取，只检查状态
                }
                
            } catch (e: any) {
                console.log("  ❌ 计算待领取奖励失败:", e.message);
                console.log("  错误详情:", e);
            }
            
        } catch (e: any) {
            console.log("  ❌ 查询失败:", e.message);
        }
    }
    
    console.log("\n========================================");
    console.log("测试实际领取")
    console.log("========================================");
    
    // 只测试第一个 NFT
    const testTokenId = '4812';
    console.log(`\n测试领取 NFT #${testTokenId} 的奖励...`);
    
    try {
        const stakeInfo = await staking.stakes(testTokenId);
        
        if (!stakeInfo.isActive) {
            console.log("❌ NFT 未质押");
            return;
        }
        
        // 检查是否是所有者
        if (stakeInfo.owner.toLowerCase() !== deployer.address.toLowerCase()) {
            console.log("❌ 不是该 NFT 的所有者");
            console.log("所有者:", stakeInfo.owner);
            console.log("当前地址:", deployer.address);
            return;
        }
        
        // 计算待领取奖励
        const pendingRewards = await staking.calculatePendingRewards(testTokenId);
        console.log("待领取奖励:", ethers.utils.formatEther(pendingRewards), "CPOP");
        
        if (pendingRewards.gt(0)) {
            console.log("\n尝试领取奖励...");
            const tx = await staking.claimRewards(testTokenId);
            console.log("交易哈希:", tx.hash);
            console.log("等待确认...");
            const receipt = await tx.wait();
            console.log("✅ 领取成功!");
            console.log("区块号:", receipt.blockNumber);
            console.log("Gas 使用:", receipt.gasUsed.toString());
        } else {
            console.log("❌ 没有可领取的奖励");
        }
        
    } catch (e: any) {
        console.log("❌ 领取失败:", e.message);
        console.log("错误详情:", e);
        if (e.data) {
            console.log("错误数据:", e.data);
        }
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });

