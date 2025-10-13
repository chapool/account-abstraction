import { ethers } from "hardhat";

async function main() {
  const tokenId = "2645"; // 字符串形式，和前端一样

  console.log(`\n🔍 完全模拟前端调用方式测试\n`);
  console.log("=".repeat(60));

  // 新的 StakingReader 地址
  const stakingReaderAddress = "0x3eCA9E28583C469CAA654fcb49Bd94Ce25C0262C";

  const StakingReader = await ethers.getContractFactory("StakingReader");
  const stakingReader = StakingReader.attach(stakingReaderAddress);

  console.log(`StakingReader: ${stakingReaderAddress}`);
  console.log(`Token ID: ${tokenId}`);
  console.log(`调用函数: getPendingRewards(${tokenId})\n`);

  try {
    // 完全按照前端的方式调用
    const result = await stakingReader.getPendingRewards(tokenId);
    console.log(`✅ 成功! Pending Rewards: ${ethers.utils.formatEther(result)} CPOP`);
    console.log(`原始值: ${result.toString()}`);
    
  } catch (error: any) {
    console.error(`❌ 调用失败!`);
    console.error(`错误消息: ${error.message}`);
    
    if (error.reason) {
      console.error(`Revert 原因: ${error.reason}`);
    }
    
    if (error.data) {
      console.error(`Error data: ${error.data}`);
    }
    
    if (error.transaction) {
      console.error(`Transaction:`, error.transaction);
    }
    
    // 打印完整错误
    console.error(`\n完整错误对象:`, JSON.stringify(error, null, 2));
  }

  // 再测试一下直接调用 Staking 合约
  console.log(`\n\n🔍 直接测试 Staking 合约...\n`);
  
  const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
  const Staking = await ethers.getContractFactory("Staking");
  const staking = Staking.attach(stakingAddress);
  
  try {
    const result = await staking.calculatePendingRewards(tokenId);
    console.log(`✅ Staking.calculatePendingRewards 成功!`);
    console.log(`Pending Rewards: ${ethers.utils.formatEther(result)} CPOP`);
  } catch (error: any) {
    console.error(`❌ Staking.calculatePendingRewards 失败!`);
    console.error(`错误: ${error.message}`);
    if (error.reason) {
      console.error(`原因: ${error.reason}`);
    }
  }

  // 检查 stakes 状态
  console.log(`\n\n🔍 检查 stakes[${tokenId}] 状态...\n`);
  
  try {
    const stakeInfo = await staking.stakes(tokenId);
    console.log(`Stakes info:`);
    console.log(`  owner: ${stakeInfo.owner}`);
    console.log(`  tokenId: ${stakeInfo.tokenId}`);
    console.log(`  level: ${stakeInfo.level}`);
    console.log(`  isActive: ${stakeInfo.isActive}`);
    console.log(`  stakeTime: ${stakeInfo.stakeTime}`);
  } catch (error: any) {
    console.error(`❌ 无法获取 stakes: ${error.message}`);
  }

  console.log("\n" + "=".repeat(60) + "\n");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

