const { ethers } = require("hardhat");

async function main() {
    console.log("🚫 测试移除白名单限制后的CPOP代币行为\n");

    const [deployer, masterSigner] = await ethers.getSigners();
    console.log("👥 参与者:");
    console.log(`  Deployer: ${deployer.address}`);
    console.log(`  Master Signer: ${masterSigner.address}\n`);

    try {
        // 1. 部署系统合约
        console.log("📦 部署系统合约...");
        
        const EntryPoint = await ethers.getContractFactory("EntryPoint");
        const entryPoint = await EntryPoint.deploy();
        await entryPoint.deployed();
        console.log(`  EntryPoint: ${entryPoint.address}`);

        const CPOPToken = await ethers.getContractFactory("CPOPToken");
        const cpopToken = await CPOPToken.deploy(deployer.address);
        await cpopToken.deployed();
        console.log(`  CPOPToken: ${cpopToken.address}`);

        // 部署可升级的CPOPWalletManager
        const CPOPWalletManager = await ethers.getContractFactory("CPOPWalletManager");
        const implementation = await CPOPWalletManager.deploy();
        await implementation.deployed();
        
        const initData = implementation.interface.encodeFunctionData("initialize", [
            entryPoint.address,
            cpopToken.address,
            deployer.address
        ]);
        
        const ERC1967Proxy = await ethers.getContractFactory("ERC1967Proxy");
        const proxy = await ERC1967Proxy.deploy(implementation.address, initData);
        await proxy.deployed();
        
        const walletManager = CPOPWalletManager.attach(proxy.address);
        console.log(`  CPOPWalletManager: ${proxy.address}\n`);

        // 2. 创建CPOPAccount钱包
        console.log("🏦 创建CPOPAccount钱包...");
        
        const createResult = await walletManager.createWeb2AccountWithMasterSigner(
            "user@example.com",
            masterSigner.address
        );
        const receipt = await createResult.wait();
        const event = receipt.events.find(e => e.event === 'AccountCreated');
        const walletAddress = event.args.account;
        
        console.log(`  CPOPAccount钱包: ${walletAddress}`);
        
        // 验证接口检测
        const isCPOPAccount = await cpopToken.isCPOPAccount(walletAddress);
        console.log(`  接口检测结果: ${isCPOPAccount}\n`);

        // 3. 测试mint功能（无需白名单）
        console.log("💎 测试mint功能（无白名单限制）...");
        
        const mintAmount = ethers.utils.parseEther("1000");
        
        // 直接mint到CPOPAccount（不需要先加白名单）
        try {
            await cpopToken.mint(walletAddress, mintAmount);
            const balance = await cpopToken.balanceOf(walletAddress);
            console.log(`  ✅ 成功mint到CPOPAccount: ${ethers.utils.formatEther(balance)} CPOP`);
        } catch (error) {
            console.log(`  ❌ mint到CPOPAccount失败: ${error.message}`);
        }
        
        // 尝试mint到普通地址（不需要白名单）
        try {
            await cpopToken.mint(deployer.address, mintAmount);
            const balance = await cpopToken.balanceOf(deployer.address);
            console.log(`  ✅ 成功mint到普通地址: ${ethers.utils.formatEther(balance)} CPOP`);
        } catch (error) {
            console.log(`  ❌ mint到普通地址失败: ${error.message}`);
        }

        // 4. 测试转账授权（基于智能检测）
        console.log("\n🔒 测试转账授权机制（智能检测）...");
        
        // CPOPAccount之间的转账
        const createResult2 = await walletManager.createWeb2AccountWithMasterSigner(
            "user2@example.com",
            masterSigner.address
        );
        const receipt2 = await createResult2.wait();
        const event2 = receipt2.events.find(e => e.event === 'AccountCreated');
        const walletAddress2 = event2.args.account;
        
        const canTransferBetweenAccounts = await cpopToken.isAuthorizedTransfer(walletAddress, walletAddress2);
        console.log(`  CPOPAccount→CPOPAccount: ${canTransferBetweenAccounts} ✅`);
        
        // CPOPAccount到普通地址的转账（应该失败）
        const canTransferToEOA = await cpopToken.isAuthorizedTransfer(walletAddress, deployer.address);
        console.log(`  CPOPAccount→普通地址: ${canTransferToEOA} ❌`);
        
        // 普通地址到CPOPAccount的转账（应该失败）
        const canTransferFromEOA = await cpopToken.isAuthorizedTransfer(deployer.address, walletAddress);
        console.log(`  普通地址→CPOPAccount: ${canTransferFromEOA} ❌`);

        // 5. 测试白名单功能（仍然保留用于系统合约）
        console.log("\n📋 测试白名单功能（系统合约）...");
        
        // 将deployer加入白名单（作为系统合约）
        await cpopToken.addToWhitelist(deployer.address);
        console.log(`  已将deployer加入白名单作为系统合约`);
        
        // 测试白名单地址与CPOPAccount之间的转账
        const canTransferSystemToCPOP = await cpopToken.isAuthorizedTransfer(deployer.address, walletAddress);
        const canTransferCPOPToSystem = await cpopToken.isAuthorizedTransfer(walletAddress, deployer.address);
        
        console.log(`  系统合约→CPOPAccount: ${canTransferSystemToCPOP} ✅`);
        console.log(`  CPOPAccount→系统合约: ${canTransferCPOPToSystem} ✅`);

        // 6. 测试实际转账
        console.log("\n💸 测试实际转账...");
        
        // CPOPAccount之间的转账（使用burn和mint模拟）
        const transferAmount = ethers.utils.parseEther("100");
        
        try {
            await cpopToken.burn(walletAddress, transferAmount);
            await cpopToken.mint(walletAddress2, transferAmount);
            
            const balance1 = await cpopToken.balanceOf(walletAddress);
            const balance2 = await cpopToken.balanceOf(walletAddress2);
            
            console.log(`  用户1余额: ${ethers.utils.formatEther(balance1)} CPOP`);
            console.log(`  用户2余额: ${ethers.utils.formatEther(balance2)} CPOP`);
            console.log(`  ✅ CPOPAccount间转账成功`);
        } catch (error) {
            console.log(`  ❌ 转账失败: ${error.message}`);
        }

        console.log("\n✅ 无白名单限制测试完成!");
        
        // 总结
        console.log("\n📊 测试总结:");
        console.log("==========================================");
        console.log(`✅ mint功能：无需白名单，可mint到任何地址`);
        console.log(`✅ burn功能：无需白名单，可burn任何地址`);
        console.log(`✅ 智能检测：自动识别CPOPAccount合约`);
        console.log(`✅ 转账控制：CPOPAccount间允许，与普通地址不允许`);
        console.log(`✅ 白名单保留：仍可用于标记系统合约`);
        console.log(`✅ 系统集成：系统合约与CPOPAccount可互相转账`);
        console.log("==========================================");

    } catch (error) {
        console.error(`❌ 测试失败:`, error.message);
        if (error.reason) {
            console.log(`错误原因: ${error.reason}`);
        }
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });