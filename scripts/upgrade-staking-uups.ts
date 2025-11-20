import { ethers, upgrades } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 升级Staking合约（UUPS代理） - 移除batchStake和batchUnstake的50个NFT限制
 */
async function main() {
  console.log("=".repeat(80));
  console.log("升级Staking合约（UUPS）- 移除批量操作NFT数量限制");
  console.log("=".repeat(80));

  const [deployer] = await ethers.getSigners();
  console.log("部署账户:", deployer.address);
  console.log("账户余额:", ethers.utils.formatEther(await deployer.getBalance()), "BNB\n");

  // 读取部署信息
  const deploymentPath = path.join(__dirname, "../deployments/bnbTestnet/core.json");
  const deployment = JSON.parse(fs.readFileSync(deploymentPath, "utf-8"));
  
  const stakingProxyAddress = deployment.contracts.Staking;
  console.log("当前Staking代理地址:", stakingProxyAddress);
  
  // 获取当前实现地址
  const currentImplAddress = await upgrades.erc1967.getImplementationAddress(stakingProxyAddress);
  console.log("当前实现地址:", currentImplAddress);
  
  console.log("\n" + "-".repeat(80));
  console.log("步骤 1: 编译并部署新实现合约");
  console.log("-".repeat(80));
  
  const StakingFactory = await ethers.getContractFactory("Staking");
  console.log("✅ Staking合约工厂已创建");
  
  console.log("正在部署新实现合约...");
  const newImplementation = await StakingFactory.deploy();
  await newImplementation.deployed();
  console.log("✅ 新实现合约已部署:", newImplementation.address);
  
  console.log("\n" + "-".repeat(80));
  console.log("步骤 2: 调用upgradeToAndCall升级代理（UUPS）");
  console.log("-".repeat(80));
  
  // 连接到代理合约
  const proxy = await ethers.getContractAt("Staking", stakingProxyAddress);
  
  // 准备空的初始化数据（如果不需要调用初始化函数）
  const initData = "0x";
  
  console.log("正在执行 upgradeToAndCall...");
  console.log(`  代理地址: ${stakingProxyAddress}`);
  console.log(`  新实现: ${newImplementation.address}`);
  console.log(`  初始化数据: ${initData}`);
  
  // 调用代理的upgradeToAndCall函数（UUPS特有）
  const tx = await proxy.upgradeToAndCall(newImplementation.address, initData, {
    gasLimit: 500000
  });
  
  console.log("升级交易已发送:", tx.hash);
  const receipt = await tx.wait();
  console.log("✅ 升级交易已确认");
  console.log(`  Gas使用: ${receipt.gasUsed.toString()}`);
  console.log(`  区块号: ${receipt.blockNumber}`);
  
  console.log("\n" + "-".repeat(80));
  console.log("步骤 3: 验证升级结果");
  console.log("-".repeat(80));
  
  // 获取新实现地址
  const newImplAddress = await upgrades.erc1967.getImplementationAddress(stakingProxyAddress);
  console.log("新实现地址:", newImplAddress);
  
  // 验证实现地址是否正确更新
  if (newImplAddress.toLowerCase() === newImplementation.address.toLowerCase()) {
    console.log("✅ 实现地址已正确更新");
  } else {
    console.log("❌ 实现地址不匹配!");
    console.log(`  期望: ${newImplementation.address}`);
    console.log(`  实际: ${newImplAddress}`);
    throw new Error("升级失败：实现地址未更新");
  }
  
  // 验证合约仍然可以正常调用
  try {
    const cpnftAddress = await proxy.cpnftContract();
    console.log("✅ CPNFT合约地址:", cpnftAddress);
    
    const totalStaked = await proxy.totalStakedCount();
    console.log("✅ 当前总质押数量:", totalStaked.toString());
    
    console.log("✅ 合约功能正常");
  } catch (error: any) {
    console.error("❌ 升级后验证失败:", error.message);
    throw error;
  }
  
  console.log("\n" + "=".repeat(80));
  console.log("📝 升级总结");
  console.log("=".repeat(80));
  console.log(`代理地址: ${stakingProxyAddress} (不变)`);
  console.log(`旧实现: ${currentImplAddress}`);
  console.log(`新实现: ${newImplementation.address}`);
  console.log(`代理类型: UUPS`);
  console.log(`升级方法: upgradeToAndCall`);
  console.log("\n✅ 升级成功完成!");
  console.log("\n📋 变更内容:");
  console.log("  - 移除了 batchStake 的 50 个NFT限制");
  console.log("  - 移除了 batchUnstake 的 50 个NFT限制");
  console.log("  - 现在可以一次性质押/解除质押任意数量的NFT（受区块gas限制）");
  console.log("\n⚠️  注意事项:");
  console.log("  - 建议单次操作不超过100个NFT以避免gas过高");
  console.log("  - 大批量操作建议分批进行以确保稳定性");
  console.log("=".repeat(80));
  
  // 更新部署文件
  const stakingConfigPath = path.join(__dirname, "../deployments/bnbTestnet/staking-config.json");
  let stakingConfig: any = {};
  
  if (fs.existsSync(stakingConfigPath)) {
    stakingConfig = JSON.parse(fs.readFileSync(stakingConfigPath, "utf-8"));
  }
  
  stakingConfig.stakingProxy = stakingProxyAddress;
  stakingConfig.stakingImplementation = newImplementation.address;
  stakingConfig.previousImplementation = currentImplAddress;
  stakingConfig.lastUpgrade = new Date().toISOString();
  stakingConfig.upgradeNote = "Removed 50 NFT limit from batchStake and batchUnstake";
  stakingConfig.upgradeMethod = "upgradeToAndCall (UUPS)";
  stakingConfig.proxyType = "UUPS";
  
  fs.writeFileSync(stakingConfigPath, JSON.stringify(stakingConfig, null, 2));
  console.log(`\n💾 升级信息已保存到: ${stakingConfigPath}`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

