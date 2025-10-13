import { ethers } from "hardhat";

async function main() {
  const cpnftAddress = "0xcC63bf57EF4b4fE5635cF0745Ae7E2C75A63c7Ed";
  const tokenId = 2645;

  console.log(`\n🔍 检查 NFT 合约: ${cpnftAddress}`);
  console.log(`📋 Token ID: ${tokenId}\n`);
  console.log("=".repeat(70));

  const CPNFT = await ethers.getContractFactory("CPNFT");
  const cpnft = CPNFT.attach(cpnftAddress);

  // 1. 检查合约基本信息
  console.log("\n📦 合约基本信息:");
  console.log("-".repeat(70));
  
  try {
    const name = await cpnft.name();
    const symbol = await cpnft.symbol();
    
    console.log(`名称: ${name}`);
    console.log(`符号: ${symbol}`);
    
    // 获取各个等级的供应量
    const supplies = await cpnft.getAllLevelSupplies();
    const levelNames = ["NORMAL", "C", "B", "A", "S", "SS", "SSS"];
    let totalSupply = 0n;
    
    console.log(`\n各等级供应量:`);
    supplies.forEach((supply: bigint, index: number) => {
      console.log(`  ${levelNames[index]}: ${supply}`);
      totalSupply += supply;
    });
    console.log(`总供应量: ${totalSupply}`);
  } catch (error: any) {
    console.log(`❌ 无法获取合约信息: ${error.message}`);
  }

  // 2. 检查特定 token ID
  console.log(`\n🔎 检查 Token ID ${tokenId}:`);
  console.log("-".repeat(70));

  try {
    const owner = await cpnft.ownerOf(tokenId);
    console.log(`✅ Token ${tokenId} 存在`);
    console.log(`   Owner: ${owner}`);
    
    try {
      const level = await cpnft.getTokenLevel(tokenId);
      const levelNames = ["NORMAL", "C", "B", "A", "S", "SS", "SSS"];
      console.log(`   Level: ${level} (${levelNames[level] || "Unknown"})`);
    } catch (e: any) {
      console.log(`   ⚠️  无法获取 Level: ${e.message}`);
    }
  } catch (error: any) {
    console.log(`❌ Token ${tokenId} 不存在`);
    console.log(`   原因: ${error.message}`);
  }

  // 3. 查找合约中存在的 token IDs（采样方式）
  console.log(`\n🔍 查找存在的 Token IDs (采样检查 0-100):`);
  console.log("-".repeat(70));

  const existingTokens: number[] = [];
  const maxCheck = 100; // 检查前 100 个 token ID

  console.log(`检查范围: Token ID 0 到 ${maxCheck}...`);

  // 检查 token ID 0 到 maxCheck
  for (let i = 0; i <= maxCheck; i++) {
    try {
      const owner = await cpnft.ownerOf(i);
      const level = await cpnft.getTokenLevel(i);
      existingTokens.push(i);
      
      if (existingTokens.length <= 10) {
        const levelNames = ["NORMAL", "C", "B", "A", "S", "SS", "SSS"];
        console.log(`  ✅ Token ${i}: Owner=${owner.slice(0, 10)}..., Level=${levelNames[level]}`);
      }
    } catch (e) {
      // Token 不存在，跳过
    }
  }

  console.log(`\n找到 ${existingTokens.length} 个存在的 tokens`);
  if (existingTokens.length > 10) {
    console.log(`前 10 个已显示，其他 token IDs: ${existingTokens.slice(10, 20).join(", ")}...`);
  }

  // 4. 检查 token ID 2645 附近的范围
  console.log(`\n🔍 检查 Token ${tokenId} 附近的范围 (±10):`);
  console.log("-".repeat(70));

  const nearbyTokens: number[] = [];
  for (let i = tokenId - 10; i <= tokenId + 10; i++) {
    if (i < 0) continue;
    
    try {
      const owner = await cpnft.ownerOf(i);
      const level = await cpnft.getTokenLevel(i);
      const levelNames = ["NORMAL", "C", "B", "A", "S", "SS", "SSS"];
      nearbyTokens.push(i);
      console.log(`  ✅ Token ${i}: Level=${levelNames[level]}, Owner=${owner.slice(0, 10)}...`);
    } catch (e) {
      console.log(`  ❌ Token ${i}: 不存在`);
    }
  }

  // 5. 尝试通过事件查找 minted tokens
  console.log(`\n📜 查询最近的 Mint 事件:`);
  console.log("-".repeat(70));

  try {
    const currentBlock = await ethers.provider.getBlockNumber();
    const fromBlock = currentBlock - 10000; // 最近 10000 个区块
    
    const transferFilter = cpnft.filters.Transfer(ethers.ZeroAddress, null);
    const mintEvents = await cpnft.queryFilter(transferFilter, fromBlock, currentBlock);
    
    console.log(`找到 ${mintEvents.length} 个 Mint 事件 (最近 10000 区块)`);
    
    const mintedTokenIds = mintEvents.slice(-10).map((event: any) => {
      return {
        tokenId: event.args?.tokenId?.toString() || "Unknown",
        to: event.args?.to || "Unknown",
        blockNumber: event.blockNumber
      };
    });
    
    console.log(`\n最近 Mint 的 10 个 NFT:`);
    mintedTokenIds.forEach((mint) => {
      console.log(`  Token ${mint.tokenId} -> ${mint.to.slice(0, 10)}... (Block: ${mint.blockNumber})`);
    });

    // 检查是否有 token 2645
    const token2645Event = mintEvents.find((event: any) => 
      event.args?.tokenId?.toString() === tokenId.toString()
    );
    
    if (token2645Event) {
      console.log(`\n✅ 找到 Token ${tokenId} 的 Mint 事件:`);
      console.log(`   To: ${token2645Event.args?.to}`);
      console.log(`   Block: ${token2645Event.blockNumber}`);
    } else {
      console.log(`\n❌ 未找到 Token ${tokenId} 的 Mint 事件`);
    }

  } catch (error: any) {
    console.log(`❌ 无法查询事件: ${error.message}`);
  }

  // 6. 总结
  console.log(`\n📊 总结:`);
  console.log("=".repeat(70));
  console.log(`✅ 找到 ${existingTokens.length} 个存在的 tokens (前 ${maxCheck} 个)`);
  console.log(`✅ Token ${tokenId} 附近找到 ${nearbyTokens.length} 个 tokens`);
  
  if (existingTokens.includes(tokenId)) {
    console.log(`✅ Token ${tokenId} 存在于合约中`);
  } else {
    console.log(`❌ Token ${tokenId} 不存在于合约中`);
    console.log(`\n💡 建议:`);
    console.log(`   1. 检查前端使用的 token ID 是否正确`);
    console.log(`   2. 可能 NFT 已被销毁或转移`);
    console.log(`   3. 使用实际存在的 token ID 进行测试`);
    
    if (existingTokens.length > 0) {
      console.log(`\n   可用的 token IDs: ${existingTokens.slice(0, 5).join(", ")}${existingTokens.length > 5 ? "..." : ""}`);
    }
  }

  console.log("\n" + "=".repeat(70) + "\n");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

