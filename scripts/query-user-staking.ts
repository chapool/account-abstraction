import { ethers } from "hardhat";
import * as fs from "fs";

/**
 * 查询用户质押状态
 * 用户地址: 0xDf3715f4693CC308c961AaF0AacD56400E229F43
 */

async function main() {
    console.log("\n🔍 查询用户质押状态...\n");

    const [deployer] = await ethers.getSigners();
    console.log("查询者地址:", deployer.address);

    // 用户地址
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    console.log("目标用户地址:", userAddress);
    console.log();

    // Staking 合约地址
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    console.log("Staking 合约地址:", stakingAddress);

    try {
        // 连接到 Staking 合约
        const stakingContract = await ethers.getContractAt("Staking", stakingAddress);
        
        console.log("========================================");
        console.log("用户质押概览");
        console.log("========================================");
        
        // 获取用户质押的 NFT 列表
        let userStakes = [];
        try {
            // 由于 userStakes 是公共映射，我们需要通过循环来获取所有质押的 NFT
            // 先尝试获取第一个 NFT，如果失败说明用户没有质押
            let i = 0;
            while (true) {
                try {
                    const tokenId = await stakingContract.userStakes(userAddress, i);
                    userStakes.push(tokenId);
                    i++;
                } catch (e) {
                    // 如果获取失败，说明已经获取完所有质押的 NFT
                    break;
                }
            }
            console.log("用户质押的 NFT 数量:", userStakes.length);
        } catch (e) {
            console.log("无法获取用户质押列表:", e.message);
            return;
        }
        
        if (userStakes.length === 0) {
            console.log("❌ 该用户没有质押任何 NFT");
            return;
        }
        
        console.log("质押的 NFT Token IDs:", userStakes.map(id => id.toString()));
        
        // 获取用户各等级的 NFT 数量
        console.log("\n========================================");
        console.log("各等级 NFT 数量统计");
        console.log("========================================");
        
        const levelCounts = {};
        for (let level = 1; level <= 6; level++) {
            const count = await stakingContract.userLevelCounts(userAddress, level);
            if (count.gt(0)) {
                levelCounts[level] = count.toString();
                const levelName = getLevelName(level);
                console.log(`${levelName} (Level ${level}): ${count.toString()} 个`);
            }
        }
        
        console.log("\n========================================");
        console.log("详细质押信息");
        console.log("========================================");
        
        // 获取每个 NFT 的详细信息
        for (let i = 0; i < userStakes.length; i++) {
            const tokenId = userStakes[i];
            console.log(`\n--- NFT #${tokenId} ---`);
            
            try {
                const stakeInfo = await stakingContract.stakes(tokenId);
                
                console.log("所有者:", stakeInfo.owner);
                console.log("Token ID:", stakeInfo.tokenId.toString());
                console.log("等级:", getLevelName(stakeInfo.level));
                console.log("质押时间:", new Date(Number(stakeInfo.stakeTime) * 1000).toLocaleString());
                console.log("最后领取时间:", new Date(Number(stakeInfo.lastClaimTime) * 1000).toLocaleString());
                console.log("是否活跃:", stakeInfo.isActive);
                console.log("总奖励:", ethers.utils.formatEther(stakeInfo.totalRewards), "CPOP");
                console.log("待领取奖励:", ethers.utils.formatEther(stakeInfo.pendingRewards), "CPOP");
                
                // 计算质押天数
                const currentTime = Math.floor(Date.now() / 1000);
                const stakingDays = Math.floor((currentTime - Number(stakeInfo.stakeTime)) / (24 * 60 * 60));
                console.log("质押天数:", stakingDays, "天");
                
                // 计算待领取奖励
                try {
                    const pendingRewards = await stakingContract.calculatePendingRewards(tokenId);
                    console.log("当前待领取奖励:", ethers.utils.formatEther(pendingRewards), "CPOP");
                } catch (e) {
                    console.log("无法计算待领取奖励:", e.message);
                }
                
            } catch (e) {
                console.log("❌ 无法获取 NFT 详细信息:", e.message);
            }
        }
        
        console.log("\n========================================");
        console.log("组合加成状态");
        console.log("========================================");
        
        // 检查各等级的组合加成状态
        for (let level = 1; level <= 6; level++) {
            const count = await stakingContract.userLevelCounts(userAddress, level);
            if (count.gt(0)) {
                const levelName = getLevelName(level);
                console.log(`\n${levelName} (Level ${level}) - ${count.toString()} 个:`);
                
                try {
                    // 获取组合状态
                    const comboStatus = await stakingContract.getComboStatus(userAddress, level);
                    console.log("  组合状态:");
                    console.log("    等级:", comboStatus.level);
                    console.log("    数量:", comboStatus.count.toString());
                    console.log("    生效时间:", comboStatus.effectiveFrom.gt(0) ? 
                        new Date(Number(comboStatus.effectiveFrom) * 1000).toLocaleString() : "立即生效");
                    console.log("    加成:", (Number(comboStatus.bonus) / 100).toFixed(1) + "%");
                    console.log("    是否待生效:", comboStatus.isPending);
                    
                    // 获取有效组合加成
                    const effectiveBonus = await stakingContract.getEffectiveComboBonus(userAddress, level);
                    console.log("    当前有效加成:", (Number(effectiveBonus) / 100).toFixed(1) + "%");
                    
                } catch (e) {
                    console.log("  ❌ 无法获取组合状态:", e.message);
                }
            }
        }
        
        console.log("\n========================================");
        console.log("总结");
        console.log("========================================");
        console.log(`用户 ${userAddress} 质押状态:`);
        console.log(`- 总质押 NFT 数量: ${userStakes.length}`);
        console.log(`- 各等级分布:`, Object.entries(levelCounts).map(([level, count]) => 
            `${getLevelName(Number(level))}: ${count}`).join(", "));
        
        // 检查是否有组合加成
        const hasComboBonus = Object.keys(levelCounts).some(level => {
            const count = Number(levelCounts[level]);
            return count >= 3;
        });
        
        if (hasComboBonus) {
            console.log("✅ 该用户有组合加成资格");
        } else {
            console.log("❌ 该用户暂无组合加成资格（需要至少3个同等级NFT）");
        }
        
    } catch (error) {
        console.log("❌ 查询失败:", error.message);
    }
}

function getLevelName(level: number): string {
    const levelNames = {
        1: "C级",
        2: "B级", 
        3: "A级",
        4: "S级",
        5: "SS级",
        6: "SSS级"
    };
    return levelNames[level] || `Level ${level}`;
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });
