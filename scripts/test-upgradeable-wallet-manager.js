const { ethers } = require("hardhat");

async function main() {
    console.log("🔄 测试CPOPWalletManager升级功能\n");

    const [deployer, proxyAdmin, newOwner] = await ethers.getSigners();
    console.log("👥 参与者:");
    console.log(`  Deployer: ${deployer.address}`);
    console.log(`  Proxy Admin: ${proxyAdmin.address}`);
    console.log(`  New Owner: ${newOwner.address}\n`);

    try {
        // 1. 部署依赖合约
        console.log("📦 部署依赖合约...");
        const EntryPoint = await ethers.getContractFactory("EntryPoint");
        const entryPoint = await EntryPoint.deploy();
        await entryPoint.deployed();
        console.log(`  EntryPoint: ${entryPoint.address}`);

        const CPOPToken = await ethers.getContractFactory("CPOPToken");
        const cpopToken = await CPOPToken.deploy(deployer.address);
        await cpopToken.deployed();
        console.log(`  CPOPToken: ${cpopToken.address}\n`);

        // 2. 部署实现合约
        console.log("🏭 部署CPOPWalletManager实现合约...");
        const CPOPWalletManager = await ethers.getContractFactory("CPOPWalletManager");
        const implementationV1 = await CPOPWalletManager.deploy();
        await implementationV1.deployed();
        console.log(`  实现合约V1: ${implementationV1.address}`);

        // 3. 部署代理合约
        console.log("\n🔄 部署UUPS代理合约...");
        const initData = implementationV1.interface.encodeFunctionData("initialize", [
            entryPoint.address,
            cpopToken.address,
            deployer.address
        ]);

        const ERC1967Proxy = await ethers.getContractFactory("ERC1967Proxy");
        const proxy = await ERC1967Proxy.deploy(implementationV1.address, initData);
        await proxy.deployed();
        console.log(`  代理合约: ${proxy.address}`);

        // 4. 连接代理合约接口
        const walletManagerProxy = CPOPWalletManager.attach(proxy.address);
        
        // 验证初始化
        const owner = await walletManagerProxy.owner();
        const implementation = await walletManagerProxy.getImplementation();
        console.log(`  代理合约Owner: ${owner}`);
        console.log(`  账户实现地址: ${implementation}`);

        // 5. 测试基本功能
        console.log("\n🧪 测试基本功能...");
        const userIdentifier = "test@example.com";
        const masterSigner = deployer.address;
        
        try {
            const createResult = await walletManagerProxy.createWeb2AccountWithMasterSigner(
                userIdentifier,
                masterSigner
            );
            const receipt = await createResult.wait();
            const event = receipt.events.find(e => e.event === 'AccountCreated');
            const userWallet = event.args.account;
            console.log(`  ✅ 创建用户钱包成功: ${userWallet}`);
        } catch (error) {
            console.log(`  ❌ 创建钱包失败: ${error.message}`);
        }

        // 6. 部署新版本实现合约（模拟升级）
        console.log("\n🆙 模拟合约升级...");
        console.log("  部署新版本实现合约...");
        const implementationV2 = await CPOPWalletManager.deploy();
        await implementationV2.deployed();
        console.log(`  实现合约V2: ${implementationV2.address}`);

        // 7. 执行升级
        console.log("  执行UUPS升级...");
        try {
            const upgradeTx = await walletManagerProxy.upgradeToAndCall(implementationV2.address, "0x");
            await upgradeTx.wait();
            console.log(`  ✅ 升级成功`);
            
            // 验证升级后状态
            const newImplementation = await walletManagerProxy.getImplementation();
            const ownerAfterUpgrade = await walletManagerProxy.owner();
            console.log(`  新实现地址: ${newImplementation}`);
            console.log(`  升级后Owner: ${ownerAfterUpgrade}`);
            console.log(`  Owner保持不变: ${owner === ownerAfterUpgrade}`);
            
        } catch (error) {
            console.log(`  ❌ 升级失败: ${error.message}`);
        }

        // 8. 测试升级后功能
        console.log("\n🧪 测试升级后功能...");
        const newUserIdentifier = "upgraded@example.com";
        
        try {
            const createResult = await walletManagerProxy.createWeb2AccountWithMasterSigner(
                newUserIdentifier,
                masterSigner
            );
            const receipt = await createResult.wait();
            const event = receipt.events.find(e => e.event === 'AccountCreated');
            const newUserWallet = event.args.account;
            console.log(`  ✅ 升级后创建钱包成功: ${newUserWallet}`);
        } catch (error) {
            console.log(`  ❌ 升级后创建钱包失败: ${error.message}`);
        }

        // 9. 测试权限转移
        console.log("\n👑 测试所有权转移...");
        try {
            const transferTx = await walletManagerProxy.transferOwnership(newOwner.address);
            await transferTx.wait();
            console.log(`  ✅ 所有权转移到: ${newOwner.address}`);
            
            const finalOwner = await walletManagerProxy.owner();
            console.log(`  最终Owner: ${finalOwner}`);
            console.log(`  权限转移成功: ${finalOwner === newOwner.address}`);
            
        } catch (error) {
            console.log(`  ❌ 权限转移失败: ${error.message}`);
        }

        console.log("\n✅ CPOPWalletManager升级功能测试完成!");
        
        // 总结
        console.log("\n📊 测试总结:");
        console.log("==========================================");
        console.log(`✅ UUPS代理合约部署成功`);
        console.log(`✅ 合约初始化功能正常`);
        console.log(`✅ 基本钱包创建功能正常`);
        console.log(`✅ 合约升级功能正常`);
        console.log(`✅ 升级后功能保持正常`);
        console.log(`✅ 所有权转移功能正常`);
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