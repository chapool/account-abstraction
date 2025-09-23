import { ethers } from "hardhat";

async function main() {
  console.log("CPNFT Staking Usage Examples");

  // 获取合约实例
  const stakingAddress = "0x..."; // 替换为实际的质押合约地址
  const CPNFTStakingFactory = await ethers.getContractFactory("CPNFTStaking");
  const staking = CPNFTStakingFactory.attach(stakingAddress);

  // 获取用户账户
  const [owner, user1, user2] = await ethers.getSigners();
  console.log("User1 address:", user1.address);
  console.log("User2 address:", user2.address);

  // 示例1: 质押NFT
  console.log("\n=== 示例1: 质押NFT ===");
  const tokenId = 1;
  
  try {
    // 计算收益地址
    const beneficiary = await staking.calculateBeneficiaryAddress(tokenId, user1.address);
    console.log("Calculated beneficiary address:", beneficiary);
    
    // 计算masterSigner
    const masterSigner = await staking.getCalculatedMasterSigner(tokenId);
    console.log("Calculated master signer:", masterSigner);
    
    // 质押NFT
    console.log("Staking NFT...");
    const stakeTx = await staking.connect(user1).stake(tokenId);
    await stakeTx.wait();
    console.log("✓ NFT staked successfully");
    console.log("Transaction hash:", stakeTx.hash);
    
    // 获取质押详情
    const stakeDetails = await staking.getStakeDetails(tokenId);
    console.log("Stake details:");
    console.log("- Owner:", stakeDetails.stakeInfo.owner);
    console.log("- Beneficiary:", stakeDetails.stakeInfo.beneficiary);
    console.log("- Stake time:", new Date(stakeDetails.stakeInfo.stakeTime * 1000));
    console.log("- Is active:", stakeDetails.stakeInfo.isActive);
    console.log("- Calculated beneficiary:", stakeDetails.calculatedBeneficiary);
    
  } catch (error) {
    console.log("✗ Staking failed:", error.message);
  }

  // 示例2: 批量质押
  console.log("\n=== 示例2: 批量质押 ===");
  const tokenIds = [2, 3, 4];
  
  try {
    console.log("Batch staking NFTs:", tokenIds);
    const batchStakeTx = await staking.connect(user1).batchStake(tokenIds);
    await batchStakeTx.wait();
    console.log("✓ Batch staking completed");
    console.log("Transaction hash:", batchStakeTx.hash);
    
    // 检查每个NFT的质押状态
    for (const id of tokenIds) {
      const details = await staking.getStakeDetails(id);
      console.log(`NFT ${id}:`, {
        owner: details.stakeInfo.owner,
        beneficiary: details.stakeInfo.beneficiary,
        isActive: details.stakeInfo.isActive
      });
    }
    
  } catch (error) {
    console.log("✗ Batch staking failed:", error.message);
  }

  // 示例3: 查询用户质押信息
  console.log("\n=== 示例3: 查询用户质押信息 ===");
  
  try {
    // 获取用户质押的所有NFT
    const userStakeInfo = await staking.getUserStakeInfo(user1.address);
    console.log("User stake info:");
    console.log("- Token IDs:", userStakeInfo.tokenIds);
    console.log("- Total pending rewards:", userStakeInfo.totalPendingRewards.toString());
    
    // 获取用户待领取奖励
    const pendingRewards = await staking.getUserPendingRewards(user1.address);
    console.log("- Pending rewards:", pendingRewards.toString());
    
    // 检查收益地址的奖励
    if (userStakeInfo.tokenIds.length > 0) {
      const firstTokenId = userStakeInfo.tokenIds[0];
      const details = await staking.getStakeDetails(firstTokenId);
      const beneficiaryRewards = await staking.getBeneficiaryPendingRewards(details.stakeInfo.beneficiary);
      console.log("- Beneficiary pending rewards:", beneficiaryRewards.toString());
    }
    
  } catch (error) {
    console.log("✗ Query failed:", error.message);
  }

  // 示例4: 领取奖励
  console.log("\n=== 示例4: 领取奖励 ===");
  
  try {
    // 等待一段时间让奖励累积
    console.log("Waiting for rewards to accumulate...");
    await new Promise(resolve => setTimeout(resolve, 5000)); // 等待5秒
    
    // 领取单个NFT的奖励
    const claimTx = await staking.connect(user1).claimReward(tokenId);
    await claimTx.wait();
    console.log("✓ Single reward claimed");
    console.log("Transaction hash:", claimTx.hash);
    
    // 批量领取奖励
    const batchClaimTx = await staking.connect(user1).batchClaimRewards(tokenIds);
    await batchClaimTx.wait();
    console.log("✓ Batch rewards claimed");
    console.log("Transaction hash:", batchClaimTx.hash);
    
  } catch (error) {
    console.log("✗ Claim failed:", error.message);
  }

  // 示例5: 验证收益地址
  console.log("\n=== 示例5: 验证收益地址 ===");
  
  try {
    const isValid = await staking.verifyBeneficiaryAddress(tokenId);
    console.log("Is beneficiary address valid:", isValid);
    
    // 重新计算并比较
    const details = await staking.getStakeDetails(tokenId);
    const calculated = details.calculatedBeneficiary;
    const stored = details.stakeInfo.beneficiary;
    console.log("Calculated beneficiary:", calculated);
    console.log("Stored beneficiary:", stored);
    console.log("Addresses match:", calculated.toLowerCase() === stored.toLowerCase());
    
  } catch (error) {
    console.log("✗ Verification failed:", error.message);
  }

  // 示例6: 取消质押
  console.log("\n=== 示例6: 取消质押 ===");
  
  try {
    // 等待最小质押期
    console.log("Waiting for minimum stake period...");
    await new Promise(resolve => setTimeout(resolve, 2000)); // 等待2秒
    
    const unstakeTx = await staking.connect(user1).unstake(tokenId);
    await unstakeTx.wait();
    console.log("✓ NFT unstaked successfully");
    console.log("Transaction hash:", unstakeTx.hash);
    
    // 验证NFT不再被质押
    const details = await staking.getStakeDetails(tokenId);
    console.log("NFT is still active:", details.stakeInfo.isActive);
    
  } catch (error) {
    console.log("✗ Unstaking failed:", error.message);
  }

  // 示例7: 管理员功能
  console.log("\n=== 示例7: 管理员功能 ===");
  
  try {
    // 更新动态乘数
    console.log("Updating dynamic multiplier...");
    const updateTx = await staking.connect(owner).updateDynamicMultiplier(0, 1100); // 110% for level 0
    await updateTx.wait();
    console.log("✓ Dynamic multiplier updated");
    
    // 设置等级总供应量
    console.log("Setting level total supply...");
    const setSupplyTx = await staking.connect(owner).setLevelTotalSupply(0, 10000); // Level 0 total supply
    await setSupplyTx.wait();
    console.log("✓ Level total supply set");
    
    // 更新默认主签名者
    console.log("Updating default master signer...");
    const newMasterSigner = "0x1234567890123456789012345678901234567890";
    const updateMasterTx = await staking.connect(owner).updateDefaultMasterSigner(newMasterSigner);
    await updateMasterTx.wait();
    console.log("✓ Default master signer updated");
    
    // 暂停/恢复质押
    console.log("Pausing staking...");
    const pauseTx = await staking.connect(owner).setStakingPaused(true);
    await pauseTx.wait();
    console.log("✓ Staking paused");
    
    console.log("Resuming staking...");
    const resumeTx = await staking.connect(owner).setStakingPaused(false);
    await resumeTx.wait();
    console.log("✓ Staking resumed");
    
  } catch (error) {
    console.log("✗ Admin function failed:", error.message);
  }

  console.log("\n=== 所有示例执行完成 ===");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
