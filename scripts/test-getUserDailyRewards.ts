import { ethers } from "hardhat";

async function main() {
    console.log("🔍 测试 getUserDailyRewards 函数...\n");

    const [deployer] = await ethers.getSigners();
    console.log("查询者地址:", deployer.address);

    // 用户地址
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    console.log("目标用户地址:", userAddress);
    console.log();

    // StakingReader 合约地址
    const stakingReaderAddress = "0x3a42F2F6962Fdc9bb8392F73C94e79071De5eC8d";
    console.log("StakingReader 合约地址:", stakingReaderAddress);

    try {
        // 连接到 StakingReader 合约
        const stakingReaderContract = await ethers.getContractAt("StakingReader", stakingReaderAddress);
        
        console.log("========================================");
        console.log("getUserDailyRewards 函数测试");
        console.log("========================================");
        
        // 调用 getUserDailyRewards 函数
        const dailyRewards = await stakingReaderContract.getUserDailyRewards(userAddress);
        
        console.log("📊 用户每日收益详情:");
        console.log("----------------------------------------");
        console.log("总基础收益:", ethers.utils.formatEther(dailyRewards.totalBaseReward), "CPOP");
        console.log("总衰减后收益:", ethers.utils.formatEther(dailyRewards.totalDecayedReward), "CPOP");
        console.log("总Combo加成:", (Number(dailyRewards.totalComboBonus) / 100).toFixed(2) + "%");
        console.log("总动态乘数:", (Number(dailyRewards.totalDynamicMultiplier) / 100).toFixed(2) + "%");
        console.log("总最终收益:", ethers.utils.formatEther(dailyRewards.totalFinalReward), "CPOP");
        console.log("总乘数:", (Number(dailyRewards.totalBonus) / 100).toFixed(2) + "%");
        
        console.log("\n📈 各等级收益分布:");
        console.log("----------------------------------------");
        const levelNames = ["C级", "B级", "A级", "S级", "SS级", "SSS级"];
        
        for (let i = 0; i < 6; i++) {
            const level = i + 1;
            const baseReward = ethers.utils.formatEther(dailyRewards.baseRewardPerLevel[i]);
            const finalReward = ethers.utils.formatEther(dailyRewards.finalRewardPerLevel[i]);
            const nftCount = dailyRewards.nftCountPerLevel[i].toString();
            
            if (Number(nftCount) > 0) {
                console.log(`${levelNames[i]} (Level ${level}):`);
                console.log(`  基础收益: ${baseReward} CPOP`);
                console.log(`  最终收益: ${finalReward} CPOP`);
                console.log(`  NFT数量: ${nftCount} 个`);
                
                if (Number(baseReward) > 0) {
                    const bonus = (Number(dailyRewards.finalRewardPerLevel[i]) * 10000) / Number(dailyRewards.baseRewardPerLevel[i]);
                    console.log(`  等级加成: ${(bonus / 100).toFixed(2)}%`);
                }
                console.log();
            }
        }
        
        console.log("========================================");
        console.log("函数调用成功！");
        console.log("========================================");
        
    } catch (error) {
        console.log("❌ 函数调用失败:", error.message);
        
        // 如果是合约地址问题，尝试查找正确的地址
        if (error.message.includes("call exception") || error.message.includes("revert")) {
            console.log("\n🔍 尝试查找 StakingReader 合约地址...");
            
            // 检查部署结果文件
            try {
                const fs = require('fs');
                const deploymentFiles = [
                    'staking-reader-deployment-results.json',
                    'sepolia-deployment-results.md',
                    'deployments/sepoliaCustom/StakingReader.json'
                ];
                
                for (const file of deploymentFiles) {
                    try {
                        if (fs.existsSync(file)) {
                            const content = fs.readFileSync(file, 'utf8');
                            console.log(`📄 找到部署文件: ${file}`);
                            
                            if (file.endsWith('.json')) {
                                const data = JSON.parse(content);
                                if (data.address) {
                                    console.log(`📍 StakingReader 地址: ${data.address}`);
                                }
                            } else if (file.endsWith('.md')) {
                                const addressMatch = content.match(/StakingReader.*?0x[a-fA-F0-9]{40}/);
                                if (addressMatch) {
                                    console.log(`📍 可能的地址: ${addressMatch[0]}`);
                                }
                            }
                        }
                    } catch (e) {
                        // 忽略文件读取错误
                    }
                }
            } catch (e) {
                console.log("无法查找部署文件");
            }
        }
    }

    console.log("\n🎉 测试完成！");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });
