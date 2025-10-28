import { ethers } from "hardhat";

async function main() {
    console.log("\n🔍 检查 AA 账户配置\n");

    // 合约地址
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const accountManagerAddress = "0x2E4f862Ba3ee6D84dd19ae9f002F5D8c0C5675ef";
    
    console.log("Staking 地址:", stakingAddress);
    console.log("AccountManager 地址:", accountManagerAddress);
    console.log();
    
    // 连接到合约
    const accountManager = await ethers.getContractAt("AccountManager", accountManagerAddress);
    
    // 检查 master signer
    console.log("========================================");
    console.log("1. 检查 Master Signer");
    console.log("========================================");
    
    try {
        const masterSigner = await accountManager.getDefaultMasterSigner();
        console.log("Master Signer:", masterSigner);
        
        if (masterSigner === ethers.constants.AddressZero) {
            console.log("❌ Master Signer 未设置");
        } else {
            console.log("✅ Master Signer 已设置");
        }
    } catch (e: any) {
        console.log("❌ 获取 Master Signer 失败:", e.message);
    }
    
    // 检查用户 AA 账户
    console.log("\n========================================");
    console.log("2. 检查用户的 AA 账户");
    console.log("========================================");
    
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    console.log("用户地址:", userAddress);
    
    try {
        const masterSigner = await accountManager.getDefaultMasterSigner();
        console.log("Master Signer:", masterSigner);
        
        if (masterSigner === ethers.constants.AddressZero) {
            console.log("❌ Master Signer 未设置，无法获取 AA 账户");
        } else {
            try {
                const aaAccount = await accountManager.getAccountAddress(userAddress, masterSigner);
                console.log("AA 账户:", aaAccount);
                
                if (aaAccount === ethers.constants.AddressZero) {
                    console.log("❌ AA 账户未找到");
                    console.log("\n💡 解决方案:");
                    console.log("   需要先为该用户创建 AA 账户");
                } else {
                    console.log("✅ AA 账户已存在");
                    
                    // 检查账户余额
                    const cpopTokenAddress = "0xA2d58d11c2752b010C2444fa2f795c6cf4cb76Bc";
                    const cpopToken = await ethers.getContractAt("CPOPToken", cpopTokenAddress);
                    
                    try {
                        const balance = await cpopToken.balanceOf(aaAccount);
                        console.log("AA 账户 CPOP 余额:", ethers.utils.formatEther(balance), "CPOP");
                    } catch (e) {
                        console.log("无法查询余额");
                    }
                }
            } catch (e: any) {
                console.log("❌ 获取 AA 账户失败:", e.message);
            }
        }
    } catch (e: any) {
        console.log("❌ 检查失败:", e.message);
    }
    
    // 测试 claimRewards 的完整流程
    console.log("\n========================================");
    console.log("3. 诊断 claimRewards 错误原因");
    console.log("========================================");
    
    const staking = await ethers.getContractAt("Staking", stakingAddress);
    const tokenId = 4812;
    
    try {
        // 获取质押信息
        const stakeInfo = await staking.stakes(tokenId);
        console.log("✅ 质押信息正常");
        console.log("  所有者:", stakeInfo.owner);
        console.log("  是否活跃:", stakeInfo.isActive);
        
        // 检查是否有待领取奖励
        const pendingRewards = await staking.calculatePendingRewards(tokenId);
        console.log("  待领取奖励:", ethers.utils.formatEther(pendingRewards), "CPOP");
        
        if (pendingRewards.isZero()) {
            console.log("⚠️ 待领取奖励为 0");
        }
        
        // 尝试获取 AA 账户（模拟 claimRewards 的内部逻辑）
        console.log("\n模拟 claimRewards 内部逻辑:");
        
        if (stakeInfo.owner.toLowerCase() !== userAddress.toLowerCase()) {
            console.log("❌ 用户不是 NFT 所有者");
            console.log("  NFT 所有者:", stakeInfo.owner);
            console.log("  用户地址:", userAddress);
        } else {
            console.log("✅ 用户是 NFT 所有者");
        }
        
        if (!stakeInfo.isActive) {
            console.log("❌ NFT 未质押");
        } else {
            console.log("✅ NFT 已质押");
        }
        
        // 检查 AA 账户
        const masterSigner = await accountManager.getDefaultMasterSigner();
        if (masterSigner === ethers.constants.AddressZero) {
            console.log("❌ Master Signer 未设置");
        } else {
            try {
                const aaAccount = await accountManager.getAccountAddress(userAddress, masterSigner);
                if (aaAccount === ethers.constants.AddressZero) {
                    console.log("❌ AA 账户不存在");
                } else {
                    console.log("✅ AA 账户存在:", aaAccount);
                }
            } catch (e: any) {
                console.log("❌ 获取 AA 账户失败:", e.message);
            }
        }
        
    } catch (e: any) {
        console.log("❌ 诊断失败:", e.message);
    }
    
    console.log("\n========================================");
    console.log("总结");
    console.log("========================================");
    console.log(`
领取奖励可能的错误原因:

1. ❌ 不是 NFT 所有者
   - 检查用户地址是否与 stakeInfo.owner 匹配
   
2. ❌ NFT 未质押
   - 检查 stakeInfo.isActive 是否为 true
   
3. ❌ 没有待领取奖励
   - 检查 _calculatePendingRewards 是否返回 > 0
   
4. ❌ Master Signer 未设置
   - 检查 accountManager.getDefaultMasterSigner() 是否为 0
   
5. ❌ AA 账户不存在
   - 检查 accountManager.getAccountAddress() 是否返回 0
   - 需要先调用 accountManager.createAccount() 创建 AA 账户
   
6. ❌ 合约已暂停
   - 检查 staking.paused() 是否为 true
    `);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });

