const { ethers } = require("hardhat");

async function main() {
    console.log("🔍 单独测试CPOPAccount部署...\n");

    const [deployer] = await ethers.getSigners();
    console.log(`Deployer: ${deployer.address}\n`);

    try {
        // 1. 部署EntryPoint
        console.log("1️⃣ 部署EntryPoint...");
        const EntryPoint = await ethers.getContractFactory("EntryPoint");
        const entryPoint = await EntryPoint.deploy();
        await entryPoint.deployed();
        console.log(`✅ EntryPoint: ${entryPoint.address}\n`);

        // 2. 测试单独部署CPOPAccount
        console.log("2️⃣ 测试部署CPOPAccount...");
        const CPOPAccount = await ethers.getContractFactory("CPOPAccount");
        const cpopAccount = await CPOPAccount.deploy(entryPoint.address);
        await cpopAccount.deployed();
        console.log(`✅ CPOPAccount: ${cpopAccount.address}\n`);

        console.log("✅ CPOPAccount单独部署成功!");

    } catch (error) {
        console.error(`❌ 部署失败:`, error.message);
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