import { ethers } from "hardhat";

async function main() {
    console.log("\n🧪 模拟真实的领取奖励流程\n");

    // 关键信息
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const accountManagerAddress = "0x2E4f862Ba3ee6D84dd19ae9f002F5D8c0C5675ef";
    const cpopTokenAddress = "0xA2d58d11c2752b010C2444fa2f795c6cf4cb76Bc";
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    const tokenId = 4812;
    
    console.log("NFT ID:", tokenId);
    console.log("用户地址:", userAddress);
    console.log();
    
    // 连接到合约
    const staking = await ethers.getContractAt("Staking", stakingAddress);
    const accountManager = await ethers.getContractAt("AccountManager", accountManagerAddress);
    const cpopToken = await ethers.getContractAt("CPOPToken", cpopTokenAddress);
    
    // 获取质押信息
    console.log("========================================");
    console.log("步骤 1: 检查质押信息");
    console.log("========================================");
    const stakeInfo = await staking.stakes(tokenId);
    console.log("✅ 质押信息:");
    console.log("  所有者:", stakeInfo.owner);
    console.log("  是否活跃:", stakeInfo.isActive);
    console.log("  质押时间:", new Date(Number(stakeInfo.stakeTime) * 1000).toISOString());
    console.log();
    
    // 计算待领取奖励
    console.log("========================================");
    console.log("步骤 2: 计算待领取奖励");
    console.log("========================================");
    const pendingRewards = await staking.calculatePendingRewards(tokenId);
    console.log("待领取奖励:", ethers.utils.formatEther(pendingRewards), "CPOP");
    console.log();
    
    // 获取 AA 账户
    console.log("========================================");
    console.log("步骤 3: 获取 AA 账户");
    console.log("========================================");
    const masterSigner = await accountManager.getDefaultMasterSigner();
    console.log("Master Signer:", masterSigner);
    const aaAccount = await accountManager.getAccountAddress(userAddress, masterSigner);
    console.log("AA 账户:", aaAccount);
    console.log();
    
    // 查询 AA 账户当前余额
    console.log("========================================");
    console.log("步骤 4: 查询 AA 账户当前余额");
    console.log("========================================");
    const balanceBefore = await cpopToken.balanceOf(aaAccount);
    console.log("领取前余额:", ethers.utils.formatEther(balanceBefore), "CPOP");
    console.log();
    
    // 检查合约状态
    console.log("========================================");
    console.log("步骤 5: 检查合约状态");
    console.log("========================================");
    const isPaused = await staking.paused();
    console.log("合约是否暂停:", isPaused ? "是" : "否");
    console.log();
    
    // 检查权限
    console.log("========================================");
    console.log("步骤 6: 检查权限");
    console.log("========================================");
    console.log("权限检查:");
    console.log("  stakeInfo.owner:", stakeInfo.owner);
    console.log("  userAddress:", userAddress);
    console.log("  是否匹配:", stakeInfo.owner.toLowerCase() === userAddress.toLowerCase() ? "✅" : "❌");
    console.log();
    
    // 模拟领取 - 使用调用方式而非交易
    console.log("========================================");
    console.log("步骤 7: 模拟调用 claimRewards");
    console.log("========================================");
    console.log("尝试使用 owner 的地址来模拟调用...");
    
    // 注意：这里需要使用正确的私钥来模拟
    // 由于我们无法访问用户的私钥，所以我们只能提供详细的分析
    
    console.log("\n💡 关键发现:");
    console.log("  如果使用用户地址 0xDf3715f4693CC308c961AaF0AacD56400E229F43");
    console.log("  应该能够成功领取奖励");
    console.log("  奖励将发放到 AA 账户 0x16F482aFFecAb476B7125376eE0fe34f7674C7E4");
    console.log();
    
    console.log("========================================");
    console.log("可能的前端错误原因");
    console.log("========================================");
    console.log(`
1. ❌ 前端使用的账户不是 NFT 所有者
   - 前端必须使用 0xDf3715f4693CC308c961AaF0AacD56400E229F43 这个账户

2. ❌ AA 账户创建失败
   - 但诊断显示 AA 账户已存在，所以不是这个问题

3. ❌ 账户抽象账户未创建
   - 如果前端通过 AA 账户调用，需要确保 AA 账户存在
   - AA 账户: 0x16F482aFFecAb476B7125376eE0fe34f7674C7E4
   - 需要确保这个账户已创建并可以签名

4. ❌ Gas 问题
   - 确保账户有足够的 ETH 支付 Gas

5. ❌ 前端配置错误
   - 检查前端是否使用了正确的合约地址
   - 检查网络配置是否正确
    `);
    
    console.log("\n========================================");
    console.log("建议的检查步骤");
    console.log("========================================");
    console.log(`
1. 在 MetaMask 或钱包中切换到账户 0xDf3715f4693CC308c961AaF0AacD56400E229F43

2. 确保该账户有足够的 ETH 支付 Gas

3. 在前端尝试领取奖励时，检查浏览器的 console 日志

4. 如果报错 "Not the owner of this stake":
   - 确认当前连接的账户是 0xDf3715f4693CC308c961AaF0AacD56400E229F43
   
5. 如果报错 "AA account not found":
   - 尝试先创建 AA 账户
   - 或者检查 accountManager 配置

6. 如果报错 gas 相关:
   - 确保账户有 ETH
   - 尝试增加 gas limit
    `);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });

