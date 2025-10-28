import { ethers } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 查询特定 NFT 的质押信息
 */

async function main() {
    console.log("\n🔍 查询特定 NFT 质押信息...\n");

    const [deployer] = await ethers.getSigners();
    console.log("查询者地址:", deployer.address);

    // 要查询的 NFT Token IDs
    const tokenIds = ['4812', '3416', '3393', '3378', '3372', '3366'];
    console.log("目标 NFT IDs:", tokenIds.join(", "));
    console.log();

    // 读取部署记录以获取 Staking 地址
    const deploymentPath = path.join(__dirname, "../.openzeppelin/sepolia.json");
    if (!fs.existsSync(deploymentPath)) {
        throw new Error("找不到部署记录文件");
    }

    // 直接使用已知的 Staking 合约地址
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    
    console.log("Staking 合约地址:", stakingAddress);
    console.log();

    // 连接到 Staking 合约
    const stakingContract = await ethers.getContractAt("Staking", stakingAddress);

    // 查询合约的时间状态
    console.log("========================================");
    console.log("合约时间状态");
    console.log("========================================");
    try {
        const testMode = await stakingContract.testMode();
        const testTimestamp = await stakingContract.testTimestamp();
        const currentTimestamp = await stakingContract.getCurrentTimestamp();
        
        console.log("测试模式:", testMode ? "✅ 已启用" : "❌ 未启用");
        console.log("测试时间戳:", testTimestamp.toString());
        
        if (testMode) {
            const testDate = new Date(Number(testTimestamp) * 1000);
            console.log("测试日期:", testDate.toISOString());
            console.log("本地时间:", testDate.toLocaleString());
        }
        
        const currentDate = new Date(Number(currentTimestamp) * 1000);
        console.log("当前使用时间戳:", currentTimestamp.toString());
        console.log("当前使用日期:", currentDate.toISOString());
        console.log("本地时间:", currentDate.toLocaleString());
        console.log();
    } catch (e) {
        console.log("无法获取时间状态:", e.message);
        console.log();
    }

    // 获取当前区块时间
    const currentBlock = await ethers.provider.getBlock("latest");
    const blockTimestamp = currentBlock?.timestamp || 0;
    const blockDate = new Date(blockTimestamp * 1000);
    console.log("区块时间戳:", blockTimestamp.toString());
    console.log("区块时间:", blockDate.toISOString());
    console.log("本地时间:", blockDate.toLocaleString());
    console.log();

    // 查询每个 NFT 的质押信息
    console.log("========================================");
    console.log("NFT 质押详情");
    console.log("========================================");

    for (const tokenId of tokenIds) {
        console.log(`\n--- NFT #${tokenId} ---`);
        
        try {
            const stakeInfo = await stakingContract.stakes(tokenId);
            
            if (!stakeInfo.isActive) {
                console.log("❌ 该 NFT 未质押或已取消质押");
                continue;
            }
            
            console.log("所有者:", stakeInfo.owner);
            console.log("Token ID:", stakeInfo.tokenId.toString());
            console.log("等级:", getLevelName(stakeInfo.level));
            
            const stakeDate = new Date(Number(stakeInfo.stakeTime) * 1000);
            console.log("质押时间戳:", stakeInfo.stakeTime.toString());
            console.log("质押日期:", stakeDate.toISOString());
            console.log("本地时间:", stakeDate.toLocaleString());
            
            const lastClaimDate = new Date(Number(stakeInfo.lastClaimTime) * 1000);
            console.log("最后领取时间戳:", stakeInfo.lastClaimTime.toString());
            console.log("最后领取日期:", lastClaimDate.toISOString());
            console.log("本地时间:", lastClaimDate.toLocaleString());
            
            console.log("是否活跃:", stakeInfo.isActive);
            console.log("总奖励:", ethers.utils.formatEther(stakeInfo.totalRewards), "CPOP");
            console.log("待领取奖励:", ethers.utils.formatEther(stakeInfo.pendingRewards), "CPOP");
            
            // 计算质押天数
            const currentTime = Number(currentTimestamp);
            const stakingDuration = currentTime - Number(stakeInfo.stakeTime);
            const stakingDays = Math.floor(stakingDuration / 86400);
            const stakingHours = Math.floor(stakingDuration / 3600);
            const stakingMinutes = Math.floor(stakingDuration / 60);
            
            console.log("质押时长:", stakingDays, "天", 
                stakingHours % 24, "小时", 
                stakingMinutes % 60, "分钟");
            
            // 计算已质押天数（使用当前时间）
            console.log(`实际已质押: ${stakingDays} 天 ${Math.floor((stakingDuration % 86400) / 3600)} 小时`);
            
            // 计算待领取奖励
            try {
                const pendingRewards = await stakingContract.calculatePendingRewards(tokenId);
                console.log("当前待领取奖励:", ethers.utils.formatEther(pendingRewards), "CPOP");
            } catch (e) {
                console.log("无法计算待领取奖励:", e.message);
            }
            
        } catch (e) {
            console.log("❌ 无法获取 NFT 信息:", e.message);
        }
    }
    
    console.log("\n========================================");
    console.log("查询完成");
    console.log("========================================");
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

