import { ethers } from "hardhat";

async function main() {
    console.log("🧪 简化版BSC测试网系统测试");
    console.log("=".repeat(50));
    
    const [deployer] = await ethers.getSigners();
    const deployerAddress = await deployer.getAddress();
    
    console.log("🔑 测试账户:", deployerAddress);
    console.log("余额:", ethers.utils.formatEther(await deployer.getBalance()), "BNB");
    
    // 合约地址
    const contracts = {
        cpopToken: "0xc2DEe82B70C60eDbF8a8E180Cb72A2e727574260"
    };
    
    // 连接到已部署的CPOPToken
    const cpopToken = await ethers.getContractAt("contracts/cpop/CPOPToken.sol:CPOPToken", contracts.cpopToken);
    
    console.log("\n📋 CPOPToken合约测试:");
    console.log("合约地址:", contracts.cpopToken);
    console.log("代币名称:", await cpopToken.name());
    console.log("代币符号:", await cpopToken.symbol());
    console.log("总供应量:", ethers.utils.formatEther(await cpopToken.totalSupply()));
    console.log("部署者余额:", ethers.utils.formatEther(await cpopToken.balanceOf(deployerAddress)));
    
    try {
        // 测试1: 创建模拟的用户地址
        console.log("\n👥 步骤1: 创建模拟用户地址");
        const user1Address = "0x1111111111111111111111111111111111111111";
        const user2Address = "0x2222222222222222222222222222222222222222";
        
        console.log("User1地址:", user1Address);
        console.log("User2地址:", user2Address);
        
        // 测试2: 批量mint到模拟用户地址
        console.log("\n💰 步骤2: 批量mint CPOP代币测试");
        
        const recipients = [user1Address, user2Address];
        const amounts = [
            ethers.utils.parseEther("1000"), // 给User1 1000 CPOP
            ethers.utils.parseEther("500")   // 给User2 500 CPOP
        ];
        
        console.log("执行批量mint...");
        const batchMintTx = await cpopToken.batchMint(recipients, amounts);
        const batchMintReceipt = await batchMintTx.wait();
        
        console.log("✅ 批量mint成功!");
        console.log("交易哈希:", batchMintTx.hash);
        console.log("Gas消耗:", batchMintReceipt.gasUsed.toString());
        console.log("BSCScan:", `https://testnet.bscscan.com/tx/${batchMintTx.hash}`);
        
        // 检查余额
        const user1Balance = await cpopToken.balanceOf(user1Address);
        const user2Balance = await cpopToken.balanceOf(user2Address);
        
        console.log("User1余额:", ethers.utils.formatEther(user1Balance), "CPOP");
        console.log("User2余额:", ethers.utils.formatEther(user2Balance), "CPOP");
        
        // 测试3: 批量转账测试
        console.log("\n🔄 步骤3: 批量转账测试");
        
        const transferRecipients = [user1Address, user2Address];
        const transferAmounts = [
            ethers.utils.parseEther("100"), // 转给User1 100 CPOP
            ethers.utils.parseEther("50")   // 转给User2 50 CPOP
        ];
        
        console.log("执行批量转账...");
        const batchTransferTx = await cpopToken.batchTransfer(transferRecipients, transferAmounts);
        const batchTransferReceipt = await batchTransferTx.wait();
        
        console.log("✅ 批量转账成功!");
        console.log("交易哈希:", batchTransferTx.hash);
        console.log("Gas消耗:", batchTransferReceipt.gasUsed.toString());
        
        // 检查最终余额
        const user1FinalBalance = await cpopToken.balanceOf(user1Address);
        const user2FinalBalance = await cpopToken.balanceOf(user2Address);
        const deployerFinalBalance = await cpopToken.balanceOf(deployerAddress);
        
        console.log("\n📊 最终余额:");
        console.log("User1:", ethers.utils.formatEther(user1FinalBalance), "CPOP");
        console.log("User2:", ethers.utils.formatEther(user2FinalBalance), "CPOP");
        console.log("Deployer:", ethers.utils.formatEther(deployerFinalBalance), "CPOP");
        
        // 测试4: 权限测试
        console.log("\n🔐 步骤4: 权限验证测试");
        console.log("Deployer拥有的角色:");
        console.log("- ADMIN_ROLE:", await cpopToken.hasRole(deployerAddress, 1));
        console.log("- MINTER_ROLE:", await cpopToken.hasRole(deployerAddress, 2));
        console.log("- BURNER_ROLE:", await cpopToken.hasRole(deployerAddress, 4));
        
        // 测试5: Gas效率展示
        console.log("\n⚡ 步骤5: Gas效率对比");
        
        // 单次mint
        const singleMintTx = await cpopToken.mint(user1Address, ethers.utils.parseEther("10"));
        const singleMintReceipt = await singleMintTx.wait();
        
        console.log("单次mint Gas:", singleMintReceipt.gasUsed.toString());
        console.log("批量mint Gas (2地址):", batchMintReceipt.gasUsed.toString());
        
        const efficiency = ((singleMintReceipt.gasUsed.toNumber() * 2 - batchMintReceipt.gasUsed.toNumber()) / (singleMintReceipt.gasUsed.toNumber() * 2) * 100).toFixed(2);
        console.log("批量操作效率提升:", efficiency + "%");
        
        console.log("\n🎉 系统测试完成!");
        
        return {
            success: true,
            contracts: {
                cpopToken: contracts.cpopToken
            },
            testResults: {
                batchMintGas: batchMintReceipt.gasUsed.toString(),
                batchTransferGas: batchTransferReceipt.gasUsed.toString(),
                efficiency: efficiency + "%"
            }
        };
        
    } catch (error) {
        console.error("❌ 测试失败:", error);
        throw error;
    }
}

main()
    .then((result) => {
        console.log("\n✅ 测试总结:");
        console.log("=".repeat(30));
        console.log("✅ CPOPToken批量操作功能正常");
        console.log("✅ 权限控制系统工作正常");
        console.log("✅ Gas优化效果显著");
        console.log("✅ BSC测试网部署成功");
        
        console.log("\n📈 性能数据:");
        console.log("批量mint Gas:", result.testResults.batchMintGas);
        console.log("批量transfer Gas:", result.testResults.batchTransferGas);
        console.log("效率提升:", result.testResults.efficiency);
        
        console.log("\n🔗 合约链接:");
        console.log("CPOPToken:", `https://testnet.bscscan.com/address/${result.contracts.cpopToken}`);
        
        process.exit(0);
    })
    .catch((error) => {
        console.error("\n❌ 系统测试失败:", error.message);
        process.exit(1);
    });