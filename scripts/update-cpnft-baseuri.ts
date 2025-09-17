import { ethers } from 'hardhat'

async function main() {
  const CPNFT_ADDRESS = "0x0C69Fc117ebe241d7Aefa748b30F43c774eD70a5";
  const NEW_BASE_URI = "http://chapool.net:8080/api/v1/meta/";

  console.log("准备更新 CPNFT base URI...");
  console.log("合约地址:", CPNFT_ADDRESS);
  console.log("新 Base URI:", NEW_BASE_URI);

  // 获取合约实例
  const cpnft = await ethers.getContractAt("CPNFT", CPNFT_ADDRESS);

  try {
    // 检查当前所有者
    const owner = await cpnft.owner();
    console.log("\n当前所有者:", owner);

    // 发送交易
    console.log("\n发送更新 base URI 交易...");
    const tx = await cpnft.setBaseURI(NEW_BASE_URI);
    console.log("交易哈希:", tx.hash);
    
    console.log("等待交易确认...");
    const receipt = await tx.wait();
    console.log("交易已确认！区块号:", receipt.blockNumber);

    // 尝试获取一个示例 tokenURI 来验证（如果有token的话）
    try {
      const exampleTokenURI = await cpnft.tokenURI(1);
      console.log("\n示例 Token #1 URI:", exampleTokenURI);
    } catch (error) {
      console.log("\n注意：目前还没有铸造任何代币，无法验证完整的 tokenURI");
    }

    console.log("\n✅ Base URI 更新成功！");

  } catch (error: any) {
    console.error("\n错误:", error);
    if (error?.message?.includes("caller is not the owner")) {
      console.error("交易失败：发送者不是合约所有者");
    } else {
      console.error("详细错误:", error);
    }
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
