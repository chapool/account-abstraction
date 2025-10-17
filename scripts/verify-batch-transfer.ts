import { ethers } from "hardhat";

/**
 * 验证 BatchTransfer 合约部署
 * 
 * 此脚本验证合约是否正确部署并可以调用
 */

async function main() {
  console.log("\n" + "=".repeat(60));
  console.log("验证 BatchTransfer 合约部署");
  console.log("=".repeat(60));

  // Sepolia 测试网上的 BatchTransfer 合约地址
  const BATCH_TRANSFER_ADDRESS = "0x29aAD71f97Da6AA0F4096cfd50DA395eAa18D964";

  console.log(`\n📍 合约地址: ${BATCH_TRANSFER_ADDRESS}`);
  console.log(`🌐 网络: ${(await ethers.provider.getNetwork()).name}`);
  console.log(`⛓️  Chain ID: ${(await ethers.provider.getNetwork()).chainId}`);

  try {
    // 获取合约实例
    const batchTransfer = await ethers.getContractAt("BatchTransfer", BATCH_TRANSFER_ADDRESS);

    console.log("\n✅ 合约实例创建成功");

    // 测试一个简单的查询函数
    const testAddress = "0x0000000000000000000000000000000000000000";
    const testNFTAddress = "0x0000000000000000000000000000000000000001";

    try {
      const isApproved = await batchTransfer.isNFTApproved(testNFTAddress, testAddress);
      console.log(`✅ 查询函数工作正常 (isNFTApproved 返回: ${isApproved})`);
    } catch (error: any) {
      console.log(`⚠️  查询函数调用失败: ${error.message}`);
    }

    try {
      const testTokenAddress = "0x0000000000000000000000000000000000000002";
      const testAmount = ethers.utils.parseEther("100");
      const hasAllowance = await batchTransfer.hasTokenAllowance(
        testTokenAddress,
        testAddress,
        testAmount
      );
      console.log(`✅ 查询函数工作正常 (hasTokenAllowance 返回: ${hasAllowance})`);
    } catch (error: any) {
      console.log(`⚠️  查询函数调用失败: ${error.message}`);
    }

    // 获取合约代码
    const code = await ethers.provider.getCode(BATCH_TRANSFER_ADDRESS);
    if (code === "0x") {
      console.log("❌ 错误: 合约地址没有部署代码");
      return;
    }
    console.log(`✅ 合约代码存在 (大小: ${code.length / 2 - 1} 字节)`);

    console.log("\n" + "=".repeat(60));
    console.log("✅ 验证完成！BatchTransfer 合约部署成功！");
    console.log("=".repeat(60));

    console.log("\n📖 查看完整文档:");
    console.log("   docs/BatchTransfer-Usage-Guide.md");
    console.log("\n🔗 在 Etherscan 上查看:");
    console.log(`   https://sepolia.etherscan.io/address/${BATCH_TRANSFER_ADDRESS}`);
    console.log("\n💡 使用示例:");
    console.log("   examples/batch-transfer-usage.ts");
    console.log("\n🧪 运行测试:");
    console.log("   npx hardhat test test/BatchTransfer.test.ts");

  } catch (error: any) {
    console.log("\n❌ 验证失败:");
    console.log(error.message);
    process.exit(1);
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

