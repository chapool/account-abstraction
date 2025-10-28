import { ethers } from "hardhat";
import * as dotenv from "dotenv";

dotenv.config();

async function main() {
  const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
  const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";
  
  const [signer] = await ethers.getSigners();
  const provider = signer.provider;
  
  const staking = await ethers.getContractAt("Staking", stakingAddress, signer);

  console.log("🔍 检查组合加成...\n");

  // 直接查询组合状态和加成
  const comboStatus = await staking.getComboStatus(userAddress, 3);
  const effectiveBonus = await staking.getEffectiveComboBonus(userAddress, 3);
  
  console.log("🎯 组合状态 (A 级 - Level 3):");
  console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
  console.log("  当前数量:", comboStatus.currentCount.toString());
  console.log("  目标数量:", comboStatus.targetCount.toString());
  console.log("  等待期到期:", new Date(Number(comboStatus.deadline) * 1000).toLocaleString());
  console.log("  是否有效:", comboStatus.isEffective);
  console.log("  加成值:", comboStatus.bonus.toString(), "基点");
  console.log("  是否等待中:", comboStatus.isPending);
  console.log();
  
  console.log("💰 当前生效加成:", effectiveBonus.toString(), "基点");
  const bonusPercent = Number(effectiveBonus) / 100;
  console.log("  加成百分比:", bonusPercent.toFixed(2), "%");
  console.log();
  
  // 判断问题
  if (comboStatus.currentCount.toNumber() >= 10) {
    if (effectiveBonus < 2000) {
      console.log("❌ 问题发现:");
      console.log("  10 个 A 级 NFT 应该获得 20% (2000基点) 加成");
      console.log("  实际获得:", bonusPercent.toFixed(2) + "%");
      console.log();
      
      const now = await staking.getCurrentTimestamp();
      if (Number(comboStatus.deadline) > Number(now)) {
        const waitDays = Math.floor((Number(comboStatus.deadline) - Number(now)) / 86400);
        console.log("  ⏳ 等待期未结束，还需", waitDays, "天");
        console.log("     到期时间:", new Date(Number(comboStatus.deadline) * 1000).toLocaleString());
      } else {
        console.log("  ⚠️  等待期已结束但仍未生效，检查 isEffective 状态");
      }
    } else {
      console.log("✅ 组合加成正常 (≥20%)");
    }
  } else {
    console.log("⚠️  当前只有", comboStatus.currentCount.toString(), "个 A 级 NFT，需要 10 个才能获得 20% 加成");
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

