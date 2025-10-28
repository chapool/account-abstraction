import { ethers } from "hardhat";

async function main() {
    console.log("🔍 检查账户余额和Gas费用...\n");

    const [deployer] = await ethers.getSigners();
    console.log("部署者地址:", deployer.address);
    
    const balance = await deployer.getBalance();
    console.log("当前余额:", ethers.utils.formatEther(balance), "ETH");
    
    // 估算部署新实现合约的gas费用
    const Staking = await ethers.getContractFactory("Staking");
    const deployTx = await Staking.getDeployTransaction();
    const gasEstimate = await deployer.estimateGas(deployTx);
    const gasPrice = await deployer.getGasPrice();
    const gasCost = gasEstimate.mul(gasPrice);
    
    console.log("预估Gas费用:", ethers.utils.formatEther(gasCost), "ETH");
    console.log("Gas Limit:", gasEstimate.toString());
    console.log("Gas Price:", ethers.utils.formatUnits(gasPrice, "gwei"), "Gwei");
    
    if (balance.lt(gasCost)) {
        console.log("❌ 余额不足！需要充值");
        console.log("需要充值:", ethers.utils.formatEther(gasCost.sub(balance)), "ETH");
    } else {
        console.log("✅ 余额充足，可以部署");
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
