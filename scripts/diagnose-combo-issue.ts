import { ethers } from "hardhat";

async function main() {
  const userAddress = "0xDf3715f4693CC308c961AaF0AacD56400E229F43";
  const stakingAddress = "0x51a07dE2Bd277F0E6412452e3B54982Fc32CA6E5";

  const [signer] = await ethers.getSigners();
  const staking = await ethers.getContractAt("Staking", stakingAddress);

  console.log("🔍 诊断组合加成问题...\n");
  console.log("用户地址:", userAddress);
  console.log("Staking 合约:", stakingAddress);
  console.log();

  try {
    // 1. 查询当前数量
    const currentCount = await staking.userLevelCounts(userAddress, 3);
    console.log("📊 A 级 NFT 数量:", currentCount.toString());
    console.log();

    // 2. 查询组合状态
    console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
    console.log("🎯 组合状态详情");
    console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
    
    const comboStatus = await staking.getComboStatus(userAddress, 3);
    console.log("当前数量 (count):", comboStatus.count.toString());
    console.log("级别 (level):", comboStatus.level.toString());
    console.log("加成值 (bonus):", comboStatus.bonus.toString(), "基点");
    console.log("有效时间 (effectiveFrom):", new Date(Number(comboStatus.effectiveFrom) * 1000).toLocaleString());
    console.log("是否等待中 (isPending):", comboStatus.isPending);
    console.log();

    // 3. 查询当前生效的加成
    const effectiveBonus = await staking.getEffectiveComboBonus(userAddress, 3);
    console.log("💰 当前生效加成:", effectiveBonus.toString(), "基点");
    const bonusPercent = Number(effectiveBonus) / 100;
    console.log("加成百分比:", bonusPercent.toFixed(2) + "%");
    console.log();

    // 4. 诊断问题
    console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
    console.log("🔬 问题诊断");
    console.log("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");

    const now = Number(await staking.getCurrentTimestamp());
    const effectiveFrom = Number(comboStatus.effectiveFrom);

    console.log("当前时间:", new Date(now * 1000).toLocaleString());
    console.log();

    if (comboStatus.count.toNumber() >= 10) {
      console.log("✅ 数量足够 (≥10 个 A 级 NFT)");
      
      if (effectiveBonus.toNumber() < 2000) {
        console.log("❌ 加成不足 (应该是 2000 基点 = 20%)");
        console.log();
        
        if (!comboStatus.isEffective) {
          console.log("⚠️  问题 1: isEffective = false (加成未生效)");
        }
        
        if (comboStatus.isPending && effectiveFrom > now) {
          const waitDays = Math.floor((effectiveFrom - now) / 86400);
          console.log("⚠️  问题 2: 等待期未结束");
          console.log("   还需等待:", waitDays, "天");
          console.log("   生效时间:", new Date(effectiveFrom * 1000).toLocaleString());
        }
        
        if (comboStatus.bonus.toNumber() === 0) {
          console.log("⚠️  问题 3: bonus = 0 (可能是配置问题)");
        }
        
        
        // 查询配置
        const configAddress = await staking.configContract();
        console.log("\n📋 查看配置...");
        console.log("Config 合约:", configAddress);
        
        const config = await ethers.getContractAt("StakingConfig", configAddress);
        
        const thresholds = await config.getComboThresholds();
        const bonuses = await config.getComboBonuses();
        const minDays = await config.getComboMinDays();
        
        console.log("\n组合加成配置:");
        console.log("阈值:", thresholds.map(t => t.toString()).join(", "));
        console.log("加成:", bonuses.map(b => b.toString()).join(", "));
        console.log("最小天数:", minDays.map(d => d.toString()).join(", "));
        
      } else {
        console.log("✅ 组合加成正常");
      }
    } else {
      console.log("⚠️  数量不足:", currentCount.toString(), "个 A 级 NFT");
      console.log("需要至少 10 个才能获得 20% 加成");
    }

  } catch (error: any) {
    console.error("❌ 查询失败:", error.message);
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

