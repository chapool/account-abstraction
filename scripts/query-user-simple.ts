import { ethers } from "hardhat";

/**
 * 查询用户质押状态 - 简化版本
 * 用户地址: 0xDf3715f4693CC308c961AaF0AacD56400E229F43
 */

async function main() {
    console.log("\n🔍 查询用户质押状态（简化版）...\n");

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
        
        // 获取用户各等级的 NFT 数量
        console.log("各等级 NFT 数量统计:");
        const levelCounts = {};
        let totalStaked = 0;
        
        for (let level = 1; level <= 6; level++) {
            try {
                const count = await stakingContract.userLevelCounts(userAddress, level);
                if (count.gt(0)) {
                    levelCounts[level] = count.toString();
                    totalStaked += Number(count);
                    const levelName = getLevelName(level);
                    console.log(`  ${levelName} (Level ${level}): ${count.toString()} 个`);
                }
            } catch (e) {
                console.log(`  Level ${level}: 查询失败`);
            }
        }
        
        console.log(`\n总质押 NFT 数量: ${totalStaked}`);
        
        if (totalStaked === 0) {
            console.log("❌ 该用户没有质押任何 NFT");
            return;
        }
        
        console.log("\n========================================");
        console.log("组合加成状态");
        console.log("========================================");
        
        // 检查各等级的组合加成状态
        for (let level = 1; level <= 6; level++) {
            if (levelCounts[level]) {
                const count = Number(levelCounts[level]);
                const levelName = getLevelName(level);
                console.log(`\n${levelName} (Level ${level}) - ${count} 个:`);
                
                try {
                    // 获取有效组合加成
                    const effectiveBonus = await stakingContract.getEffectiveComboBonus(userAddress, level);
                    const bonusPercent = (Number(effectiveBonus) / 100).toFixed(1);
                    console.log(`  当前有效加成: ${bonusPercent}%`);
                    
                    // 获取组合状态
                    const comboStatus = await stakingContract.getComboStatus(userAddress, level);
                    console.log(`  组合状态:`);
                    console.log(`    数量: ${comboStatus.count.toString()}`);
                    console.log(`    待生效加成: ${(Number(comboStatus.bonus) / 100).toFixed(1)}%`);
                    console.log(`    是否待生效: ${comboStatus.isPending}`);
                    
                    if (comboStatus.effectiveFrom.gt(0)) {
                        const effectiveTime = new Date(Number(comboStatus.effectiveFrom) * 1000);
                        console.log(`    生效时间: ${effectiveTime.toLocaleString()}`);
                        
                        const now = new Date();
                        if (effectiveTime > now) {
                            const timeLeft = Math.ceil((effectiveTime.getTime() - now.getTime()) / (1000 * 60 * 60 * 24));
                            console.log(`    剩余等待时间: ${timeLeft} 天`);
                        }
                    }
                    
                    // 分析组合等级
                    if (count >= 10) {
                        console.log(`  🎯 组合等级: 10个NFT组合 (20%加成)`);
                        if (level === 6) {
                            console.log(`  ⚠️  SSS级别限制: 10个NFT组合不适用`);
                        }
                    } else if (count >= 5) {
                        console.log(`  🎯 组合等级: 5个NFT组合 (10%加成)`);
                    } else if (count >= 3) {
                        console.log(`  🎯 组合等级: 3个NFT组合 (5%加成)`);
                    } else {
                        console.log(`  ❌ 无组合加成 (需要至少3个同等级NFT)`);
                    }
                    
                } catch (e) {
                    console.log(`  ❌ 无法获取组合状态: ${e.message}`);
                }
            }
        }
        
        console.log("\n========================================");
        console.log("总结");
        console.log("========================================");
        console.log(`用户 ${userAddress} 质押状态:`);
        console.log(`- 总质押 NFT 数量: ${totalStaked}`);
        console.log(`- 各等级分布:`, Object.entries(levelCounts).map(([level, count]) => 
            `${getLevelName(Number(level))}: ${count}`).join(", "));
        
        // 检查是否有组合加成
        const hasComboBonus = Object.keys(levelCounts).some(level => {
            const count = Number(levelCounts[level]);
            return count >= 3;
        });
        
        if (hasComboBonus) {
            console.log("✅ 该用户有组合加成资格");
            
            // 统计组合加成等级
            const comboLevels = Object.entries(levelCounts).filter(([level, count]) => {
                const numCount = Number(count);
                return numCount >= 3;
            });
            
            console.log(`- 有组合加成的等级: ${comboLevels.length} 个`);
            comboLevels.forEach(([level, count]) => {
                const levelName = getLevelName(Number(level));
                const numCount = Number(count);
                let comboType = "";
                if (numCount >= 10 && Number(level) !== 6) {
                    comboType = "10个NFT组合 (20%加成)";
                } else if (numCount >= 5) {
                    comboType = "5个NFT组合 (10%加成)";
                } else if (numCount >= 3) {
                    comboType = "3个NFT组合 (5%加成)";
                }
                console.log(`  ${levelName}: ${comboType}`);
            });
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
