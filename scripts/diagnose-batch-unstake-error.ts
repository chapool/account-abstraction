/**
 * @fileoverview 诊断 batchUnstake 错误
 */
import { ethers } from "ethers";
import * as dotenv from "dotenv";

// Load environment variables
dotenv.config();

async function main() {
    const provider = new ethers.JsonRpcProvider(process.env.ETH_RPC_URL);
    const signer = new ethers.Wallet(process.env.PRIVATE_KEY!, provider);

    console.log("Using account:", signer.address);

    const STAKING_ADDRESS = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
    
    const tokenIds = ['4812', '3416', '3393', '3378', '3372', '3366'].map(id => BigInt(id));
    const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
    
    console.log("\n=== 诊断信息 ===");
    console.log("Staking地址:", STAKING_ADDRESS);
    console.log("用户地址:", userAddress);
    console.log("要解除质押的Token IDs:", tokenIds);
    console.log("脚本运行者余额:", (await provider.getBalance(signer.address)).toString(), "wei");
    console.log("用户地址余额:", (await provider.getBalance(userAddress)).toString(), "wei");

    // 1. 检查Staking合约是否正常运行
    try {
        const staking = new ethers.Contract(STAKING_ADDRESS, [
            "function stakes(uint256) view returns (address owner, uint256 tokenId, uint8 level, uint256 stakeTime, uint256 lastClaimTime, bool isActive, uint256 totalRewards, uint256 pendingRewards, bool continuousBonusClaimed)",
            "function paused() view returns (bool)",
            "function accountManagerContract() view returns (address)",
            "function cpopTokenContract() view returns (address)"
        ], provider);
        console.log("\n1. 检查Staking合约状态...");
        
        // 检查是否暂停
        const paused = await staking.paused();
        console.log("   合约是否暂停:", paused);
        if (paused) {
            console.log("   ❌ 合约已暂停!");
            return;
        }

        // 检查AccountManager
        const accountManagerAddr = await staking.accountManagerContract();
        console.log("   AccountManager地址:", accountManagerAddr);

        // 检查CPOPToken
        const cpopTokenAddr = await staking.cpopTokenContract();
        console.log("   CPOPToken地址:", cpopTokenAddr);

        // 2. 检查AccountManager的masterSigner
        if (accountManagerAddr) {
            const accountManager = new ethers.Contract(accountManagerAddr, [
                "function getDefaultMasterSigner() view returns (address)",
                "function getAccountAddress(address owner, address masterSigner) view returns (address)"
            ], provider);
            
            console.log("\n2. 检查AccountManager...");
            const masterSigner = await accountManager.getDefaultMasterSigner();
            console.log("   Master Signer:", masterSigner);
            
            if (masterSigner === ethers.ZeroAddress) {
                console.log("   ❌ Master Signer未设置!");
            } else {
                const aaAccount = await accountManager.getAccountAddress(userAddress, masterSigner);
                console.log("   用户AA账户:", aaAccount);
                
                if (aaAccount === ethers.ZeroAddress) {
                    console.log("   ❌ AA账户地址为0!");
                } else {
                    const codeSize = await provider.getCode(aaAccount);
                    if (codeSize === "0x") {
                        console.log("   ⚠️ AA账户未部署 (这是正常的，账户会在首次使用时自动部署)");
                    } else {
                        console.log("   ✓ AA账户已部署");
                    }
                }
            }
        }

        // 3. 检查每个NFT的状态
        console.log("\n3. 检查NFT状态...");
        for (const tokenId of tokenIds) {
            const stake = await staking.stakes(tokenId);
            console.log(`   Token ${tokenId}:`);
            console.log(`      所有者: ${stake.owner}`);
            console.log(`      是否质押: ${stake.isActive}`);
            console.log(`      质押时间: ${new Date(Number(stake.stakeTime) * 1000).toLocaleString()}`);
            console.log(`      等级: ${stake.level}`);
            console.log(`      是否属于用户: ${stake.owner.toLowerCase() === userAddress.toLowerCase()}`);
            
            if (stake.owner.toLowerCase() !== userAddress.toLowerCase()) {
                console.log(`      ❌ 所有权不匹配!`);
            }
            if (!stake.isActive) {
                console.log(`      ❌ NFT未质押!`);
            }
        }

        // 4. 尝试estimateGas (使用用户地址)
        console.log("\n4. 尝试估计gas (以用户身份)...");
        try {
            const iface = new ethers.Interface([
                "function batchUnstake(uint256[] calldata tokenIds) nonReentrant whenNotPaused"
            ]);
            const data = iface.encodeFunctionData("batchUnstake", [tokenIds]);
            const gasEstimate = await provider.estimateGas({
                to: STAKING_ADDRESS,
                data: data,
                from: userAddress
            });
            console.log("   ✓ Gas估算成功:", gasEstimate.toString());
        } catch (error: any) {
            console.log("   ❌ Gas估算失败:");
            console.log("   错误信息:", error.message);
            if (error.reason) {
                console.log("   原因:", error.reason);
            }
            if (error.data && error.data !== "0x") {
                console.log("   错误数据:", error.data);
            }
        }

    } catch (error: any) {
        console.error("错误:", error.message);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
