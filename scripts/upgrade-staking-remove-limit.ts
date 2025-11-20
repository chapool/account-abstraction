import { ethers, upgrades } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * 升级Staking合约 - 移除batchStake和batchUnstake的50个NFT限制
 */
async function main() {
  console.log("=".repeat(80));
  console.log("升级Staking合约 - 移除批量操作NFT数量限制");
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
  console.log("步骤 1: 编译新的Staking合约");
  console.log("-".repeat(80));
  
  const StakingFactory = await ethers.getContractFactory("Staking");
  console.log("✅ Staking合约工厂已创建");
  
  console.log("\n" + "-".repeat(80));
  console.log("步骤 2: 导入并验证代理合约");
  console.log("-".repeat(80));
  
  try {
    // 尝试导入现有代理
    console.log("正在导入现有代理合约...");
    await upgrades.forceImport(stakingProxyAddress, StakingFactory);
    console.log("✅ 代理合约已导入");
  } catch (error: any) {
    // 如果已经导入过，会报错但可以继续
    console.log("⚠️  导入警告:", error.message);
    console.log("继续验证升级兼容性...");
  }
  
  try {
    await upgrades.validateUpgrade(stakingProxyAddress, StakingFactory);
    console.log("✅ 升级验证通过 - 合约兼容");
  } catch (error: any) {
    console.error("❌ 升级验证失败:", error.message);
    throw error;
  }
  
  console.log("\n" + "-".repeat(80));
  console.log("步骤 3: 部署新实现合约");
  console.log("-".repeat(80));
  
  // 部署新实现
  console.log("正在部署新的Staking实现合约...");
  const newImplementation = await StakingFactory.deploy();
  await newImplementation.deployed();
  console.log("✅ 新实现合约已部署:", newImplementation.address);
  
  console.log("\n" + "-".repeat(80));
  console.log("步骤 4: 使用upgradeToAndCall升级代理");
  console.log("-".repeat(80));
  
  // 获取ProxyAdmin
  const proxyAdminAddress = await upgrades.erc1967.getAdminAddress(stakingProxyAddress);
  console.log("ProxyAdmin地址:", proxyAdminAddress);
  
  // 连接到ProxyAdmin合约（使用ABI）
  const proxyAdminAbi = [
    "function upgradeAndCall(address proxy, address implementation, bytes memory data) public payable"
  ];
  const proxyAdmin = new ethers.Contract(proxyAdminAddress, proxyAdminAbi, deployer);
  
  // 准备空的初始化数据（如果不需要调用初始化函数，使用空数据）
  // 如果需要调用某个函数，可以编码函数调用：
  // const initData = StakingFactory.interface.encodeFunctionData("someInitFunction", [params]);
  const initData = "0x";
  
  console.log("正在执行 upgradeAndCall...");
  console.log(`  代理地址: ${stakingProxyAddress}`);
  console.log(`  新实现: ${newImplementation.address}`);
  console.log(`  初始化数据: ${initData}`);
  
  // 使用ProxyAdmin的upgradeAndCall方法
  const tx = await proxyAdmin.upgradeAndCall(
    stakingProxyAddress,
    newImplementation.address,
    initData,
    { gasLimit: 5000000 } // 设置足够的gas limit
  );
  
  console.log("升级交易已发送:", tx.hash);
  const receipt = await tx.wait();
  console.log("✅ 升级交易已确认");
  console.log(`  Gas使用: ${receipt.gasUsed.toString()}`);
  console.log(`  区块号: ${receipt.blockNumber}`);
  
  // 获取新实现地址
  const newImplAddress = await upgrades.erc1967.getImplementationAddress(stakingProxyAddress);
  console.log("\n新实现地址:", newImplAddress);
  
  // 验证实现地址是否正确更新
  if (newImplAddress.toLowerCase() === newImplementation.address.toLowerCase()) {
    console.log("✅ 实现地址已正确更新");
  } else {
    console.log("⚠️  实现地址不匹配!");
    console.log(`  期望: ${newImplementation.address}`);
    console.log(`  实际: ${newImplAddress}`);
  }
  
  console.log("\n" + "-".repeat(80));
  console.log("步骤 5: 验证升级结果");
  console.log("-".repeat(80));
  
  const staking = await ethers.getContractAt("Staking", stakingProxyAddress);
  
  // 验证合约仍然可以正常调用
  try {
    const cpnftAddress = await staking.cpnftContract();
    console.log("✅ CPNFT合约地址:", cpnftAddress);
    
    const totalStaked = await staking.totalStakedCount();
    console.log("✅ 当前总质押数量:", totalStaked.toString());
    
    console.log("✅ 合约升级成功，功能正常");
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
  console.log(`升级方法: upgradeAndCall`);
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
  stakingConfig.upgradeMethod = "upgradeAndCall";
  
  fs.writeFileSync(stakingConfigPath, JSON.stringify(stakingConfig, null, 2));
  console.log(`\n💾 升级信息已保存到: ${stakingConfigPath}`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });

