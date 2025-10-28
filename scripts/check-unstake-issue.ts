import { ethers } from "hardhat";

async function main() {
    const [signer] = await ethers.getSigners();
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    const tokenIds = ['4812', '3416', '3393', '3378', '3372', '3366'];
    
    console.log("用户地址:", userAddress);
    console.log("Token IDs:", tokenIds);
    
    const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    const staking = await ethers.getContractAt("Staking", stakingAddress);
    
    // 检查每个NFT状态
    console.log("\n=== 检查NFT状态 ===");
    for (const tokenIdStr of tokenIds) {
        const tokenId = BigInt(tokenIdStr);
        const stake = await staking.stakes(tokenId);
        console.log(`\nToken ${tokenId}:`);
        console.log(`  所有者: ${stake.owner}`);
        console.log(`  质押中: ${stake.isActive}`);
        console.log(`  等级: ${stake.level}`);
        console.log(`  是否属于用户: ${stake.owner.toLowerCase() === userAddress.toLowerCase()}`);
        
        if (stake.owner.toLowerCase() !== userAddress.toLowerCase()) {
            console.log(`  ❌ 不是用户的NFT!`);
        }
        if (!stake.isActive) {
            console.log(`  ❌ NFT未质押!`);
        }
    }
    
    // 检查AccountManager
    console.log("\n=== 检查AccountManager ===");
    const accountManagerAddr = await staking.accountManagerContract();
    console.log("AccountManager地址:", accountManagerAddr);
    
    try {
        const accountManager = await ethers.getContractAt("IAccountManager", accountManagerAddr);
        const masterSigner = await accountManager.getDefaultMasterSigner();
        console.log("Master Signer:", masterSigner);
        
        if (masterSigner !== ethers.ZeroAddress) {
            const aaAccount = await accountManager.getAccountAddress(userAddress, masterSigner);
            console.log("用户的AA账户:", aaAccount);
            
            if (aaAccount === ethers.ZeroAddress) {
                console.log("❌ AA账户地址为0 - 这会导致batchUnstake失败!");
            } else {
                const code = await ethers.provider.getCode(aaAccount);
                console.log("AA账户代码大小:", code.length);
                if (code === "0x") {
                    console.log("⚠️ AA账户未部署");
                }
            }
        } else {
            console.log("❌ Master Signer未设置!");
        }
    } catch (error: any) {
        console.log("AccountManager检查失败:", error.message);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });

