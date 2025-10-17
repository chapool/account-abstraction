/**
 * Paymaster 集成示例
 * 
 * 展示如何在 Staking 项目中集成 BNB Chain Paymaster
 * 实现零 Gas 费用户体验
 */

import { ethers } from "ethers";

// ============================================
// 配置
// ============================================

const CONFIG = {
  // Paymaster 配置
  paymaster: {
    url: "https://open-platform.nodereal.io/paymaster",
    apiKey: process.env.PAYMASTER_API_KEY || "",
    sponsorAddress: process.env.SPONSOR_ADDRESS || "",
  },
  
  // 合约地址
  contracts: {
    staking: "0x...staking_contract_address",
    cpnft: "0x...cpnft_contract_address",
  },
  
  // BSC 配置
  bsc: {
    rpcUrl: "https://bsc-dataseed.binance.org/",
    chainId: 56,
  },
};

// ============================================
// Paymaster 服务类
// ============================================

class PaymasterService {
  private provider: ethers.providers.JsonRpcProvider;
  private paymasterUrl: string;
  private apiKey: string;

  constructor() {
    this.provider = new ethers.providers.JsonRpcProvider(CONFIG.bsc.rpcUrl);
    this.paymasterUrl = CONFIG.paymaster.url;
    this.apiKey = CONFIG.paymaster.apiKey;
  }

  /**
   * 发送被赞助的交易
   */
  async sendSponsoredTransaction(
    userAddress: string,
    contractAddress: string,
    functionData: string,
    gasLimit: number
  ): Promise<string> {
    try {
      // 1. 构建零 Gas 交易
      const tx: ethers.providers.TransactionRequest = {
        to: contractAddress,
        from: userAddress,
        data: functionData,
        gasPrice: 0, // 零 Gas Price
        gasLimit: gasLimit,
        chainId: CONFIG.bsc.chainId,
      };

      // 2. 获取 nonce
      const nonce = await this.provider.getTransactionCount(userAddress);
      tx.nonce = nonce;

      console.log("📝 Transaction prepared:", {
        to: tx.to,
        from: tx.from,
        gasPrice: tx.gasPrice,
        gasLimit: tx.gasLimit,
      });

      // 3. 提交到 Paymaster
      const response = await fetch(`${this.paymasterUrl}/v1/sponsor`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "X-API-Key": this.apiKey,
        },
        body: JSON.stringify({
          transaction: tx,
          sponsorPolicy: {
            type: "whitelist",
            contracts: [contractAddress],
          },
        }),
      });

      if (!response.ok) {
        throw new Error(`Paymaster rejected: ${response.statusText}`);
      }

      const result = await response.json();
      console.log("✅ Transaction sponsored:", result.transactionHash);

      return result.transactionHash;
    } catch (error) {
      console.error("❌ Sponsorship failed:", error);
      throw error;
    }
  }

  /**
   * 检查用户配额
   */
  async checkUserQuota(userAddress: string) {
    const response = await fetch(`${this.paymasterUrl}/v1/quota/${userAddress}`, {
      headers: { "X-API-Key": this.apiKey },
    });

    return response.json();
  }
}

// ============================================
// Staking 操作封装（带 Paymaster）
// ============================================

class StakingWithPaymaster {
  private paymaster: PaymasterService;
  private stakingInterface: ethers.utils.Interface;

  constructor() {
    this.paymaster = new PaymasterService();
    
    // Staking 合约 ABI（简化版）
    const stakingABI = [
      "function stake(uint256 tokenId)",
      "function unstake(uint256 tokenId)",
      "function claimRewards(uint256 tokenId)",
      "function batchStake(uint256[] calldata tokenIds)",
      "function batchClaimRewards(uint256[] calldata tokenIds)",
    ];
    
    this.stakingInterface = new ethers.utils.Interface(stakingABI);
  }

  /**
   * 质押 NFT（零 Gas 费）
   */
  async stake(userAddress: string, tokenId: number): Promise<string> {
    console.log(`\n🎯 Staking NFT #${tokenId} with Paymaster...`);

    const data = this.stakingInterface.encodeFunctionData("stake", [tokenId]);

    const txHash = await this.paymaster.sendSponsoredTransaction(
      userAddress,
      CONFIG.contracts.staking,
      data,
      150000 // Gas limit
    );

    console.log("✅ Stake successful! Zero gas fee paid by user.");
    console.log("💰 Platform sponsored ~$0.118");
    
    return txHash;
  }

  /**
   * 批量质押（零 Gas 费）
   */
  async batchStake(userAddress: string, tokenIds: number[]): Promise<string> {
    console.log(`\n🎯 Batch staking ${tokenIds.length} NFTs with Paymaster...`);

    const data = this.stakingInterface.encodeFunctionData("batchStake", [tokenIds]);

    const gasLimit = 100000 * tokenIds.length;
    const txHash = await this.paymaster.sendSponsoredTransaction(
      userAddress,
      CONFIG.contracts.staking,
      data,
      gasLimit
    );

    const sponsorCost = (gasLimit * 1e-9 * 1200).toFixed(3); // 1 Gwei, $1200/BNB
    console.log(`✅ Batch stake successful! Zero gas fee paid by user.`);
    console.log(`💰 Platform sponsored ~$${sponsorCost}`);
    
    return txHash;
  }

  /**
   * 领取奖励（零 Gas 费）
   */
  async claimRewards(userAddress: string, tokenId: number): Promise<string> {
    console.log(`\n🎯 Claiming rewards for NFT #${tokenId} with Paymaster...`);

    const data = this.stakingInterface.encodeFunctionData("claimRewards", [tokenId]);

    const txHash = await this.paymaster.sendSponsoredTransaction(
      userAddress,
      CONFIG.contracts.staking,
      data,
      150000
    );

    console.log("✅ Rewards claimed! Zero gas fee paid by user.");
    console.log("💰 Platform sponsored ~$0.142");
    
    return txHash;
  }

  /**
   * 批量领取奖励（零 Gas 费）
   */
  async batchClaimRewards(userAddress: string, tokenIds: number[]): Promise<string> {
    console.log(`\n🎯 Batch claiming rewards for ${tokenIds.length} NFTs with Paymaster...`);

    const data = this.stakingInterface.encodeFunctionData("batchClaimRewards", [tokenIds]);

    const gasLimit = 75000 * tokenIds.length;
    const txHash = await this.paymaster.sendSponsoredTransaction(
      userAddress,
      CONFIG.contracts.staking,
      data,
      gasLimit
    );

    const sponsorCost = (gasLimit * 1e-9 * 1200).toFixed(3);
    console.log(`✅ Batch claim successful! Zero gas fee paid by user.`);
    console.log(`💰 Platform sponsored ~$${sponsorCost}`);
    
    return txHash;
  }

  /**
   * 取消质押（可选：用户自付或赞助）
   */
  async unstake(
    userAddress: string, 
    tokenId: number, 
    usePaymaster: boolean = false
  ): Promise<string> {
    console.log(`\n🎯 Unstaking NFT #${tokenId}...`);

    const data = this.stakingInterface.encodeFunctionData("unstake", [tokenId]);

    if (usePaymaster) {
      const txHash = await this.paymaster.sendSponsoredTransaction(
        userAddress,
        CONFIG.contracts.staking,
        data,
        200000
      );
      
      console.log("✅ Unstake successful! Zero gas fee (sponsored).");
      return txHash;
    } else {
      // 降级到普通交易（用户自付）
      console.log("⚠️ User will pay gas fee for unstake.");
      const provider = new ethers.providers.Web3Provider(window.ethereum);
      const signer = provider.getSigner();
      
      const tx = await signer.sendTransaction({
        to: CONFIG.contracts.staking,
        data: data,
        gasLimit: 200000,
      });
      
      console.log("✅ Unstake successful! User paid ~$0.218");
      return tx.hash;
    }
  }
}

// ============================================
// 使用示例
// ============================================

async function exampleUsage() {
  const staking = new StakingWithPaymaster();
  const userAddress = "0x...user_wallet_address";

  console.log("=".repeat(60));
  console.log("Paymaster 集成示例 - Staking 操作");
  console.log("=".repeat(60));

  try {
    // 示例 1: 质押单个 NFT（免费）
    await staking.stake(userAddress, 1);

    // 示例 2: 批量质押 5 个 NFT（免费）
    await staking.batchStake(userAddress, [1, 2, 3, 4, 5]);

    // 示例 3: 领取奖励（免费）
    await staking.claimRewards(userAddress, 1);

    // 示例 4: 批量领取 10 个 NFT 奖励（免费）
    await staking.batchClaimRewards(userAddress, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);

    // 示例 5: 取消质押（用户自付）
    await staking.unstake(userAddress, 1, false);

    console.log("\n" + "=".repeat(60));
    console.log("✅ 所有操作完成！");
    console.log("=".repeat(60));

  } catch (error) {
    console.error("操作失败:", error);
  }
}

// ============================================
// 高级功能：智能降级
// ============================================

/**
 * 智能交易发送器
 * 自动尝试 Paymaster，失败则降级到普通交易
 */
class SmartTransactionSender {
  private staking: StakingWithPaymaster;

  constructor() {
    this.staking = new StakingWithPaymaster();
  }

  async sendTransaction(
    operation: 'stake' | 'claim' | 'unstake',
    userAddress: string,
    tokenId: number
  ) {
    try {
      // 1. 首先尝试 Paymaster
      console.log(`🎯 Attempting ${operation} with Paymaster...`);
      
      let txHash: string;
      switch (operation) {
        case 'stake':
          txHash = await this.staking.stake(userAddress, tokenId);
          break;
        case 'claim':
          txHash = await this.staking.claimRewards(userAddress, tokenId);
          break;
        case 'unstake':
          txHash = await this.staking.unstake(userAddress, tokenId, true);
          break;
      }

      console.log("✅ Transaction sent via Paymaster (FREE)");
      return { txHash, gasFeePaid: 0 };

    } catch (error) {
      console.log("⚠️ Paymaster failed, falling back to normal transaction...");
      
      // 2. 降级到普通交易
      const provider = new ethers.providers.Web3Provider(window.ethereum);
      const signer = provider.getSigner();
      
      const stakingContract = new ethers.Contract(
        CONFIG.contracts.staking,
        ['function stake(uint256)', 'function claimRewards(uint256)', 'function unstake(uint256)'],
        signer
      );

      let tx;
      switch (operation) {
        case 'stake':
          tx = await stakingContract.stake(tokenId);
          break;
        case 'claim':
          tx = await stakingContract.claimRewards(tokenId);
          break;
        case 'unstake':
          tx = await stakingContract.unstake(tokenId);
          break;
      }

      const receipt = await tx.wait();
      const gasUsed = receipt.gasUsed;
      const gasPrice = receipt.effectiveGasPrice;
      const gasCost = gasUsed.mul(gasPrice);
      const gasCostUSD = parseFloat(ethers.utils.formatEther(gasCost)) * 1200;

      console.log(`✅ Transaction sent via normal method`);
      console.log(`💰 User paid: $${gasCostUSD.toFixed(3)}`);
      
      return { txHash: tx.hash, gasFeePaid: gasCostUSD };
    }
  }
}

// ============================================
// React Hook 示例
// ============================================

/**
 * 在 React 应用中使用的 Hook
 */
export function usePaymasterStaking() {
  const sender = new SmartTransactionSender();

  const stakeWithPaymaster = async (tokenId: number) => {
    const provider = new ethers.providers.Web3Provider(window.ethereum);
    const userAddress = await provider.getSigner().getAddress();
    
    return sender.sendTransaction('stake', userAddress, tokenId);
  };

  const claimWithPaymaster = async (tokenId: number) => {
    const provider = new ethers.providers.Web3Provider(window.ethereum);
    const userAddress = await provider.getSigner().getAddress();
    
    return sender.sendTransaction('claim', userAddress, tokenId);
  };

  const unstakeWithPaymaster = async (tokenId: number) => {
    const provider = new ethers.providers.Web3Provider(window.ethereum);
    const userAddress = await provider.getSigner().getAddress();
    
    return sender.sendTransaction('unstake', userAddress, tokenId);
  };

  return {
    stakeWithPaymaster,
    claimWithPaymaster,
    unstakeWithPaymaster,
  };
}

// ============================================
// UI 组件示例（React + TypeScript）
// ============================================

/**
 * Stake 按钮组件
 */
export function StakeButton({ tokenId }: { tokenId: number }) {
  const [loading, setLoading] = React.useState(false);
  const [gasSaved, setGasSaved] = React.useState(0);
  const { stakeWithPaymaster } = usePaymasterStaking();

  const handleStake = async () => {
    try {
      setLoading(true);
      
      const result = await stakeWithPaymaster(tokenId);
      
      if (result.gasFeePaid === 0) {
        setGasSaved(0.118); // 节省的金额
        alert(`✅ Staked successfully!\n🎉 Gas fee sponsored by platform!\n💰 You saved $0.118`);
      } else {
        alert(`✅ Staked successfully!\n💰 Gas fee: $${result.gasFeePaid.toFixed(3)}`);
      }
    } catch (error) {
      alert(`❌ Stake failed: ${error.message}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="stake-card">
      <h3>NFT #{tokenId}</h3>
      <button onClick={handleStake} disabled={loading}>
        {loading ? "Processing..." : "🆓 Stake (Free Gas!)"}
      </button>
      {gasSaved > 0 && (
        <p className="savings">💰 Saved ${gasSaved.toFixed(3)} in gas!</p>
      )}
    </div>
  );
}

// ============================================
// 运行示例
// ============================================

async function main() {
  console.log("\n" + "=".repeat(60));
  console.log("Paymaster Integration Example");
  console.log("=".repeat(60));

  const staking = new StakingWithPaymaster();
  const userAddress = "0xa3B605fB633AD0A0DC4B74b10bBfc40fDB050d35";

  // 示例 1: 免费质押
  console.log("\n📌 Example 1: Stake with Paymaster");
  await staking.stake(userAddress, 1);

  // 示例 2: 免费批量质押
  console.log("\n📌 Example 2: Batch Stake with Paymaster");
  await staking.batchStake(userAddress, [1, 2, 3, 4, 5]);

  // 示例 3: 免费领取奖励
  console.log("\n📌 Example 3: Claim Rewards with Paymaster");
  await staking.claimRewards(userAddress, 1);

  // 示例 4: 免费批量领取
  console.log("\n📌 Example 4: Batch Claim with Paymaster");
  await staking.batchClaimRewards(userAddress, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);

  console.log("\n" + "=".repeat(60));
  console.log("✅ All operations completed with ZERO gas fee for users!");
  console.log("=".repeat(60));
}

// 如果直接运行此文件
if (require.main === module) {
  main()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}

export { PaymasterService, StakingWithPaymaster, SmartTransactionSender };

