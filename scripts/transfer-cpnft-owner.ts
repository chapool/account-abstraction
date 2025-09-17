import { ethers } from "ethers";

// CPNFT ABI - 包含我们需要的函数
const ABI = [
  "function owner() view returns (address)",
  "function transferOwnership(address newOwner) public",
  "event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)"
];

async function main() {
  // 配置
  const CPNFT_ADDRESS = "0xEa81A317a4Bc82084359028A207e282F8F503d16";
  const NEW_OWNER = "0xa3B605fB633AD0A0DC4B74b10bBfc40fDB050d35";
  
  // 连接到 Sepolia
  const provider = new ethers.providers.JsonRpcProvider("https://eth-sepolia.g.alchemy.com/v2/_x4NAgu50ejHAhTH2-gJoRNS4PQv7Tjp");
  
  // 需要私钥来发送交易
  if (!process.env.PRIVATE_KEY) {
    throw new Error("请设置 PRIVATE_KEY 环境变量");
  }
  const signer = new ethers.Wallet(process.env.PRIVATE_KEY, provider);
  
  console.log("准备转移 CPNFT 所有权...");
  console.log("合约地址:", CPNFT_ADDRESS);
  console.log("新所有者:", NEW_OWNER);

  // 获取合约实例
  const cpnft = new ethers.Contract(CPNFT_ADDRESS, ABI, signer);

  try {
    // 检查当前所有者
    const currentOwner = await cpnft.owner();
    console.log("\n当前所有者:", currentOwner);
    
    if (currentOwner.toLowerCase() === NEW_OWNER.toLowerCase()) {
      console.log("该地址已经是合约所有者了！");
      return;
    }

    console.log("\n发送转移所有权交易...");
    const tx = await cpnft.transferOwnership(NEW_OWNER);
    console.log("交易哈希:", tx.hash);
    
    console.log("等待交易确认...");
    const receipt = await tx.wait();
    console.log("交易已确认！区块号:", receipt.blockNumber);

    // 验证新的所有者
    const newOwner = await cpnft.owner();
    console.log("\n新的所有者:", newOwner);
    
    if (newOwner.toLowerCase() === NEW_OWNER.toLowerCase()) {
      console.log("✅ 所有权转移成功！");
    } else {
      console.log("❌ 所有权转移似乎没有成功，请检查！");
    }

  } catch (error: any) {
    console.error("\n错误:", error);
    if (error?.message?.includes("caller is not the owner")) {
      console.error("交易失败：发送者不是合约所有者");
    }
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });