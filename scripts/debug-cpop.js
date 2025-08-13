const { ethers } = require("hardhat");

async function main() {
    console.log("🔍 调试CPOP系统部署...\n");

    const [deployer, masterSigner] = await ethers.getSigners();
    console.log(`Deployer: ${deployer.address}`);
    console.log(`Master Signer: ${masterSigner.address}\n`);

    try {
        // 1. 部署EntryPoint
        console.log("1️⃣ 部署EntryPoint...");
        const EntryPoint = await ethers.getContractFactory("EntryPoint");
        const entryPoint = await EntryPoint.deploy();
        await entryPoint.deployed();
        console.log(`✅ EntryPoint: ${entryPoint.address}\n`);

        // 2. 部署CPOPToken
        console.log("2️⃣ 部署CPOPToken...");
        const CPOPToken = await ethers.getContractFactory("CPOPToken");
        const cpopToken = await CPOPToken.deploy(deployer.address);
        await cpopToken.deployed();
        console.log(`✅ CPOPToken: ${cpopToken.address}\n`);

        // 3. 部署CPOPWalletManager
        console.log("3️⃣ 部署CPOPWalletManager...");
        const CPOPWalletManager = await ethers.getContractFactory("CPOPWalletManager");
        const walletManager = await CPOPWalletManager.deploy();
        await walletManager.deployed();
        console.log(`✅ CPOPWalletManager: ${walletManager.address}`);

        // 4. 初始化WalletManager
        console.log("4️⃣ 初始化WalletManager...");
        try {
            const initTx = await walletManager.initialize(
                entryPoint.address,
                cpopToken.address,
                deployer.address
            );
            await initTx.wait();
            console.log(`✅ WalletManager初始化成功\n`);
        } catch (error) {
            console.log(`❌ WalletManager初始化失败: ${error.message}`);
            console.log(`错误代码: ${error.code}`);
            if (error.data) {
                console.log(`错误数据: ${error.data}`);
            }
            return;
        }

        // 5. 测试创建钱包
        console.log("5️⃣ 测试创建钱包...");
        try {
            const identifier = "test_user@example.com";
            const createTx = await walletManager.createWeb2AccountWithMasterSigner(
                identifier,
                masterSigner.address
            );
            const receipt = await createTx.wait();
            console.log(`✅ 钱包创建成功`);
            
            // 从事件中获取钱包地址
            const walletAddress = receipt.events[0].args.account;
            console.log(`钱包地址: ${walletAddress}\n`);
            
            // 6. 测试CPOP功能
            console.log("6️⃣ 测试CPOP代币功能...");
            
            // 将钱包加入白名单
            await cpopToken.addToWhitelist(walletAddress);
            console.log(`✅ 钱包已加入白名单`);
            
            // Mint代币到钱包
            const mintAmount = ethers.parseEther("100");
            await cpopToken.mint(walletAddress, mintAmount);
            
            // 检查余额
            const balance = await cpopToken.balanceOf(walletAddress);
            console.log(`钱包CPOP余额: ${ethers.formatEther(balance)} CPOP`);
            
            // 测试接口检测
            const isCPOPAccount = await cpopToken.isCPOPAccount(walletAddress);
            console.log(`是CPOPAccount: ${isCPOPAccount}`);
            
            console.log("\n✅ 所有测试通过!");
            
        } catch (error) {
            console.log(`❌ 钱包创建失败: ${error.message}`);
            console.log(`错误代码: ${error.code}`);
            if (error.data) {
                console.log(`错误数据: ${error.data}`);
            }
        }

    } catch (error) {
        console.error(`❌ 部署失败:`, error.message);
        if (error.code) {
            console.log(`错误代码: ${error.code}`);
        }
        if (error.data) {
            console.log(`错误数据: ${error.data}`);
        }
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("❌ 脚本执行失败:", error);
        process.exit(1);
    });