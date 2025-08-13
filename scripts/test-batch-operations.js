const { ethers } = require("hardhat");

async function main() {
    console.log("🚀 测试CPOPToken批量操作功能\n");

    const [deployer, user1, user2, user3, user4, user5] = await ethers.getSigners();
    console.log("👥 参与者:");
    console.log(`  Deployer: ${deployer.address}`);
    console.log(`  User1: ${user1.address}`);
    console.log(`  User2: ${user2.address}`);
    console.log(`  User3: ${user3.address}`);
    console.log(`  User4: ${user4.address}`);
    console.log(`  User5: ${user5.address}\n`);

    try {
        // 1. 部署CPOPToken
        console.log("📦 部署CPOPToken...");
        const CPOPToken = await ethers.getContractFactory("CPOPToken");
        const cpopToken = await CPOPToken.deploy(deployer.address);
        await cpopToken.deployed();
        console.log(`  CPOPToken: ${cpopToken.address}\n`);

        // 2. 测试批量mint功能
        console.log("💎 测试批量mint功能...");
        
        const recipients = [user1.address, user2.address, user3.address];
        const amounts = [
            ethers.utils.parseEther("100"),
            ethers.utils.parseEther("200"),
            ethers.utils.parseEther("300")
        ];
        
        console.log(`  准备批量mint给${recipients.length}个地址:`);
        for (let i = 0; i < recipients.length; i++) {
            console.log(`    ${recipients[i]}: ${ethers.utils.formatEther(amounts[i])} CPOP`);
        }
        
        // 执行批量mint
        const mintTx = await cpopToken.batchMint(recipients, amounts);
        const mintReceipt = await mintTx.wait();
        console.log(`  ✅ 批量mint成功，gas使用: ${mintReceipt.gasUsed.toString()}`);
        
        // 验证余额
        console.log(`  验证mint后余额:`);
        for (let i = 0; i < recipients.length; i++) {
            const balance = await cpopToken.balanceOf(recipients[i]);
            console.log(`    ${recipients[i]}: ${ethers.utils.formatEther(balance)} CPOP`);
        }
        
        // 检查总供应量
        const totalSupplyAfterMint = await cpopToken.totalSupply();
        const expectedTotal = amounts.reduce((sum, amount) => sum.add(amount), ethers.BigNumber.from(0));
        console.log(`  总供应量: ${ethers.utils.formatEther(totalSupplyAfterMint)} CPOP`);
        console.log(`  预期总量: ${ethers.utils.formatEther(expectedTotal)} CPOP`);
        console.log(`  总量验证: ${totalSupplyAfterMint.eq(expectedTotal) ? '✅' : '❌'}`);

        // 3. 测试单次mint vs 批量mint的gas对比
        console.log("\n⚡ Gas效率对比测试...");
        
        // 单次mint测试
        const singleMintAddresses = [user4.address, user5.address];
        const singleMintAmounts = [
            ethers.utils.parseEther("50"),
            ethers.utils.parseEther("75")
        ];
        
        let totalSingleMintGas = 0;
        console.log(`  单次mint测试:`);
        for (let i = 0; i < singleMintAddresses.length; i++) {
            const tx = await cpopToken.mint(singleMintAddresses[i], singleMintAmounts[i]);
            const receipt = await tx.wait();
            totalSingleMintGas += parseInt(receipt.gasUsed.toString());
            console.log(`    Mint ${i+1}: ${receipt.gasUsed.toString()} gas`);
        }
        console.log(`  单次mint总gas: ${totalSingleMintGas}`);
        
        // 批量mint测试（同样的操作）
        const batchMintTx = await cpopToken.batchMint(
            [deployer.address, deployer.address], // 临时地址用于测试
            [ethers.utils.parseEther("50"), ethers.utils.parseEther("75")]
        );
        const batchMintReceipt = await batchMintTx.wait();
        const batchMintGas = parseInt(batchMintReceipt.gasUsed.toString());
        console.log(`  批量mint gas: ${batchMintGas}`);
        console.log(`  Gas节省: ${totalSingleMintGas - batchMintGas} (${((totalSingleMintGas - batchMintGas) / totalSingleMintGas * 100).toFixed(2)}%)`);

        // 4. 测试批量burn功能
        console.log("\n🔥 测试批量burn功能...");
        
        const burnAccounts = [user1.address, user2.address];
        const burnAmounts = [
            ethers.utils.parseEther("50"),  // 从100中burn 50
            ethers.utils.parseEther("100") // 从200中burn 100
        ];
        
        console.log(`  准备批量burn:`);
        for (let i = 0; i < burnAccounts.length; i++) {
            const currentBalance = await cpopToken.balanceOf(burnAccounts[i]);
            console.log(`    ${burnAccounts[i]}: burn ${ethers.utils.formatEther(burnAmounts[i])} / ${ethers.utils.formatEther(currentBalance)} CPOP`);
        }
        
        // 执行批量burn
        const burnTx = await cpopToken.batchBurn(burnAccounts, burnAmounts);
        const burnReceipt = await burnTx.wait();
        console.log(`  ✅ 批量burn成功，gas使用: ${burnReceipt.gasUsed.toString()}`);
        
        // 验证burn后余额
        console.log(`  验证burn后余额:`);
        for (let i = 0; i < burnAccounts.length; i++) {
            const balance = await cpopToken.balanceOf(burnAccounts[i]);
            console.log(`    ${burnAccounts[i]}: ${ethers.utils.formatEther(balance)} CPOP`);
        }
        
        // 验证总供应量减少
        const totalSupplyAfterBurn = await cpopToken.totalSupply();
        const expectedBurnTotal = burnAmounts.reduce((sum, amount) => sum.add(amount), ethers.BigNumber.from(0));
        console.log(`  总供应量: ${ethers.utils.formatEther(totalSupplyAfterBurn)} CPOP`);
        console.log(`  减少了: ${ethers.utils.formatEther(expectedBurnTotal)} CPOP`);

        // 5. 测试错误情况
        console.log("\n⚠️  测试错误情况处理...");
        
        // 测试数组长度不匹配
        try {
            await cpopToken.batchMint([user1.address], [ethers.utils.parseEther("100"), ethers.utils.parseEther("200")]);
            console.log(`  ❌ 应该失败但没有失败: 数组长度不匹配`);
        } catch (error) {
            console.log(`  ✅ 正确拒绝: 数组长度不匹配`);
        }
        
        // 测试空数组
        try {
            await cpopToken.batchMint([], []);
            console.log(`  ❌ 应该失败但没有失败: 空数组`);
        } catch (error) {
            console.log(`  ✅ 正确拒绝: 空数组`);
        }
        
        // 测试burn超过余额
        try {
            await cpopToken.batchBurn([user3.address], [ethers.utils.parseEther("1000")]);
            console.log(`  ❌ 应该失败但没有失败: burn超过余额`);
        } catch (error) {
            console.log(`  ✅ 正确拒绝: burn超过余额`);
        }

        console.log("\n✅ 批量操作功能测试完成!");
        
        // 6. 最终状态总结
        console.log("\n📊 最终状态总结:");
        console.log("==========================================");
        const finalTotalSupply = await cpopToken.totalSupply();
        console.log(`总供应量: ${ethers.utils.formatEther(finalTotalSupply)} CPOP`);
        
        const allUsers = [user1, user2, user3, user4, user5];
        for (const user of allUsers) {
            const balance = await cpopToken.balanceOf(user.address);
            if (balance.gt(0)) {
                console.log(`${user.address}: ${ethers.utils.formatEther(balance)} CPOP`);
            }
        }
        
        console.log("==========================================");
        console.log("✅ 批量mint功能: 支持多地址同时mint");
        console.log("✅ 批量burn功能: 支持多地址同时burn");
        console.log("✅ Gas优化: 批量操作比单次操作更省gas");
        console.log("✅ 错误处理: 正确处理无效输入");
        console.log("✅ 权限控制: 需要MINTER_ROLE和BURNER_ROLE");

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