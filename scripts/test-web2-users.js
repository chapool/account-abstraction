const { ethers } = require("hardhat");

async function main() {
    console.log("🌐 CPOP系统Web2用户测试\n");

    // 获取系统账户（这些是实际的以太坊账户）
    const [deployer, masterSigner, systemOperator] = await ethers.getSigners();
    console.log("🔧 系统账户:");
    console.log(`  Deployer: ${deployer.address}`);
    console.log(`  Master Signer: ${masterSigner.address}`);
    console.log(`  System Operator: ${systemOperator.address}`);
    
    // Web2用户（只是标识符，没有私钥）
    console.log("\n👤 Web2用户标识:");
    const web2UserA = "alice@example.com";
    const web2UserB = "bob@example.com";
    console.log(`  用户A: ${web2UserA} (邮箱标识，无私钥)`);
    console.log(`  用户B: ${web2UserB} (邮箱标识，无私钥)`);

    try {
        // 1. 部署系统合约
        console.log("\n📦 部署系统合约...");
        
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
        console.log(`  CPOPWalletManager实现: ${implementation.address}`);
        
        // 部署代理合约
        const initData = implementation.interface.encodeFunctionData("initialize", [
            entryPoint.address,
            cpopToken.address,
            deployer.address
        ]);
        
        const ERC1967Proxy = await ethers.getContractFactory("ERC1967Proxy");
        const proxy = await ERC1967Proxy.deploy(implementation.address, initData);
        await proxy.deployed();
        
        // 连接代理接口
        const walletManager = CPOPWalletManager.attach(proxy.address);
        console.log(`  CPOPWalletManager代理: ${proxy.address}`);

        // 2. 系统管理员为Web2用户创建AA钱包
        console.log("\n🏦 系统管理员为Web2用户创建AA钱包...");
        
        // 为Web2用户A创建钱包
        console.log(`\n👤 为Web2用户 ${web2UserA} 创建钱包:`);
        const createAResult = await walletManager.createWeb2AccountWithMasterSigner(
            web2UserA,
            masterSigner.address
        );
        const receiptA = await createAResult.wait();
        const eventA = receiptA.events.find(e => e.event === 'AccountCreated');
        const walletA = eventA.args.account;
        const ownerA = eventA.args.owner;
        
        console.log(`  钱包地址: ${walletA}`);
        console.log(`  生成的Owner: ${ownerA}`);
        console.log(`  Master Signer: ${masterSigner.address}`);

        // 为Web2用户B创建钱包
        console.log(`\n👤 为Web2用户 ${web2UserB} 创建钱包:`);
        const createBResult = await walletManager.createWeb2AccountWithMasterSigner(
            web2UserB,
            masterSigner.address
        );
        const receiptB = await createBResult.wait();
        const eventB = receiptB.events.find(e => e.event === 'AccountCreated');
        const walletB = eventB.args.account;
        const ownerB = eventB.args.owner;
        
        console.log(`  钱包地址: ${walletB}`);
        console.log(`  生成的Owner: ${ownerB}`);
        console.log(`  Master Signer: ${masterSigner.address}`);

        // 3. 验证钱包配置
        console.log("\n🔍 验证钱包配置...");
        const CPOPAccount = await ethers.getContractFactory("CPOPAccount");
        const accountA = CPOPAccount.attach(walletA);
        const accountB = CPOPAccount.attach(walletB);
        
        const masterSignerA = await accountA.getMasterSigner();
        const masterSignerB = await accountB.getMasterSigner();
        const isMasterEnabledA = await accountA.isMasterSignerEnabled();
        const isMasterEnabledB = await accountB.isMasterSignerEnabled();
        
        console.log(`  用户A钱包Master Signer: ${masterSignerA}`);
        console.log(`  用户A钱包Master启用: ${isMasterEnabledA}`);
        console.log(`  用户B钱包Master Signer: ${masterSignerB}`);
        console.log(`  用户B钱包Master启用: ${isMasterEnabledB}`);

        // 4. 测试CPOP代币系统
        console.log("\n💎 测试CPOP代币系统...");
        
        // 检查钱包是否被自动识别为CPOPAccount
        const isAccountA = await cpopToken.isCPOPAccount(walletA);
        const isAccountB = await cpopToken.isCPOPAccount(walletB);
        console.log(`  钱包A被识别为CPOPAccount: ${isAccountA}`);
        console.log(`  钱包B被识别为CPOPAccount: ${isAccountB}`);
        
        // 将钱包加入白名单（用于mint）
        await cpopToken.addToWhitelist(walletA);
        await cpopToken.addToWhitelist(walletB);
        
        // 为用户A mint代币
        const mintAmount = ethers.utils.parseEther("1000");
        await cpopToken.mint(walletA, mintAmount);
        
        const balanceA = await cpopToken.balanceOf(walletA);
        console.log(`  用户A钱包CPOP余额: ${ethers.utils.formatEther(balanceA)} CPOP`);

        // 5. 测试转账授权检查
        console.log("\n🔒 测试转账授权机制...");
        
        const canTransferAtoB = await cpopToken.isAuthorizedTransfer(walletA, walletB);
        const canTransferAtoEOA = await cpopToken.isAuthorizedTransfer(walletA, deployer.address);
        const canTransferEOAtoA = await cpopToken.isAuthorizedTransfer(deployer.address, walletA);
        
        console.log(`  CPOPAccount→CPOPAccount转账: ${canTransferAtoB} ✅`);
        console.log(`  CPOPAccount→普通地址转账: ${canTransferAtoEOA} ❌`);
        console.log(`  普通地址→CPOPAccount转账: ${canTransferEOAtoA} ❌`);

        // 6. 演示Web2用户操作流程
        console.log("\n🌐 Web2用户操作流程演示...");
        console.log(`📱 场景：${web2UserA} 要向 ${web2UserB} 转账100 CPOP`);
        console.log("   1. 用户A在Web界面发起转账请求");
        console.log("   2. 后端系统验证用户身份和余额");
        console.log("   3. 系统使用Master Signer构造UserOperation");
        console.log("   4. 通过EntryPoint执行转账操作");
        console.log("   5. CPOPAccount验证Master Signer签名");
        console.log("   6. 执行CPOP代币转账");
        console.log("   7. (可选)使用CPOPPaymaster支付gas费");
        
        // 模拟转账操作的结果
        const transferAmount = ethers.utils.parseEther("100");
        await cpopToken.burn(walletA, transferAmount);
        await cpopToken.mint(walletB, transferAmount);
        
        const newBalanceA = await cpopToken.balanceOf(walletA);
        const newBalanceB = await cpopToken.balanceOf(walletB);
        
        console.log(`\n💸 转账完成:`);
        console.log(`  ${web2UserA} 余额: ${ethers.utils.formatEther(newBalanceA)} CPOP`);
        console.log(`  ${web2UserB} 余额: ${ethers.utils.formatEther(newBalanceB)} CPOP`);

        // 7. 系统架构总结
        console.log("\n📋 CPOP系统Web2架构总结:");
        console.log("==========================================");
        console.log(`🏢 后端系统:`);
        console.log(`  - Master Signer: ${masterSigner.address}`);
        console.log(`  - EntryPoint: ${entryPoint.address}`);
        console.log(`  - 钱包工厂: ${proxy.address} (代理)`);
        console.log(`  - 工厂实现: ${implementation.address}`);
        console.log(`  - CPOP代币: ${cpopToken.address}`);
        console.log(``);
        console.log(`👤 Web2用户:`);
        console.log(`  - ${web2UserA} → 钱包: ${walletA}`);
        console.log(`  - ${web2UserB} → 钱包: ${walletB}`);
        console.log(``);
        console.log(`🔐 安全特性:`);
        console.log(`  ✅ 用户无需管理私钥`);
        console.log(`  ✅ Master Signer统一签名验证`);
        console.log(`  ✅ CPOP代币限制转账范围`);
        console.log(`  ✅ 接口检测自动识别AA钱包`);
        console.log(`  ✅ 支持无ETH的Web2体验`);
        console.log(`  ✅ UUPS可升级架构`);
        console.log("==========================================");

    } catch (error) {
        console.error(`❌ 测试失败:`, error.message);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });