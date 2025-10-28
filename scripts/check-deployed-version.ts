import { ethers } from "hardhat";

async function main() {
    console.log("🔍 检查部署的合约版本...\n");

    // 获取合约实例
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const stakingContract = await ethers.getContractAt("Staking", stakingAddress);

    try {
        console.log("==========================================");
        console.log("合约版本检查");
        console.log("==========================================");
        
        console.log("Staking合约地址:", stakingAddress);
        
        const version = await stakingContract.version();
        console.log("当前部署版本:", version);
        
        console.log("\n==========================================");
        console.log("版本对比");
        console.log("==========================================");
        
        if (version === "3.9.0") {
            console.log("✅ 部署版本与代码版本一致");
            console.log("✅ 合约已是最新版本");
        } else {
            console.log("❌ 部署版本与代码版本不一致");
            console.log("📋 需要升级合约");
        }
        
        console.log("\n==========================================");
        console.log("合约状态");
        console.log("==========================================");
        
        const configAddress = await stakingContract.configContract();
        console.log("配置合约地址:", configAddress);
        
        const totalStaked = await stakingContract.totalStakedCount();
        console.log("总质押数量:", totalStaked.toString());
        
        const isPaused = await stakingContract.paused();
        console.log("合约是否暂停:", isPaused);
        
        const owner = await stakingContract.owner();
        console.log("合约所有者:", owner);

    } catch (error) {
        console.error("❌ 检查过程中出现错误:", error);
    }

    console.log("\n🎉 检查完成！");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });


