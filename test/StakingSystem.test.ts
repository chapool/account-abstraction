import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { Contract } from "ethers";

describe("Staking System Integration Test", function () {
  let stakingConfig: Contract;
  let staking: Contract;
  let stakingReader: Contract;
  let cpnft: Contract;
  let cpopToken: Contract;
  let accountManager: Contract;
  
  let owner: any;
  let user1: any;
  let user2: any;

  beforeEach(async function () {
    [owner, user1, user2] = await ethers.getSigners();

    // Deploy StakingConfig (non-upgradeable)
    const StakingConfig = await ethers.getContractFactory("StakingConfig");
    stakingConfig = await StakingConfig.deploy();

    // Deploy mock contracts
    const MockCPNFT = await ethers.getContractFactory("CPNFT");
    cpnft = await upgrades.deployProxy(MockCPNFT, ["CPNFT", "CPNFT", "https://api.example.com/"], {
      kind: "uups",
      initializer: "initialize",
    });
    
    const MockCPOPToken = await ethers.getContractFactory("CPOPToken");
    cpopToken = await MockCPOPToken.deploy(owner.address, ethers.utils.parseEther("1000000"));
    
    const MockAccountManager = await ethers.getContractFactory("AccountManager");
    accountManager = await MockAccountManager.deploy();

    // Deploy Staking contract (upgradeable)
    const Staking = await ethers.getContractFactory("Staking");
    staking = await upgrades.deployProxy(
      Staking,
      [cpnft.address, cpopToken.address, accountManager.address, stakingConfig.address, owner.address],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    // Set staking contract address in CPNFT
    await cpnft.connect(owner).setStakingContract(staking.address);

    // Grant MINTER_ROLE to staking contract
    const MINTER_ROLE = await cpopToken.MINTER_ROLE();
    await cpopToken.connect(owner).grantRole(staking.address, MINTER_ROLE);

    // Deploy StakingReader (non-upgradeable)
    const StakingReader = await ethers.getContractFactory("StakingReader");
    stakingReader = await StakingReader.deploy(staking.address, stakingConfig.address, cpnft.address);

    // Set up total supply for each level in CPNFT contract
    await cpnft.connect(owner).setLevelSupply(1, 1000); // C level
    await cpnft.connect(owner).setLevelSupply(2, 500);  // B level
    await cpnft.connect(owner).setLevelSupply(3, 200);  // A level
    await cpnft.connect(owner).setLevelSupply(4, 100);  // S level
    await cpnft.connect(owner).setLevelSupply(5, 50);   // SS level
    await cpnft.connect(owner).setLevelSupply(6, 10);   // SSS level
  });

  describe("Deployment", function () {
    it("Should deploy all contracts successfully", async function () {
      expect(stakingConfig.address).to.not.equal(ethers.constants.AddressZero);
      expect(staking.address).to.not.equal(ethers.constants.AddressZero);
      expect(stakingReader.address).to.not.equal(ethers.constants.AddressZero);
      expect(cpnft.address).to.not.equal(ethers.constants.AddressZero);
      expect(cpopToken.address).to.not.equal(ethers.constants.AddressZero);
      expect(accountManager.address).to.not.equal(ethers.constants.AddressZero);
    });

    it("Should initialize with correct default values", async function () {
      const minStakeDays = await stakingConfig.getMinStakeDays();
      expect(minStakeDays).to.equal(7);

      const earlyWithdrawPenalty = await stakingConfig.getEarlyWithdrawPenalty();
      expect(earlyWithdrawPenalty).to.equal(5000); // 50%

      const sssReward = await stakingConfig.getDailyReward(6);
      expect(sssReward).to.equal(100);

      const cReward = await stakingConfig.getDailyReward(1);
      expect(cReward).to.equal(3);
    });
  });

  describe("NFT Staking", function () {
    beforeEach(async function () {
      // Mint some NFTs for testing
      await cpnft.connect(owner).mint(user1.address, 1); // C level
      await cpnft.connect(owner).mint(user1.address, 2); // B level
      await cpnft.connect(owner).mint(user1.address, 3); // A level
      await cpnft.connect(owner).mint(user1.address, 4); // S level
      await cpnft.connect(owner).mint(user1.address, 5); // SS level
      await cpnft.connect(owner).mint(user1.address, 6); // SSS level
    });

    it("Should allow staking valid level NFTs", async function () {
      // Stake C level NFT (tokenId 1)
      await cpnft.connect(user1).approve(staking.address, 1);
      await staking.connect(user1).stake(1);

      const stakeInfo = await stakingReader.getStakeDetails(1);
      expect(stakeInfo.owner).to.equal(user1.address);
      expect(stakeInfo.level).to.equal(1); // C level
      expect(stakeInfo.isActive).to.be.true;

    });

    it("Should not allow staking NORMAL level NFTs", async function () {
      // Mint a NORMAL level NFT (level 0)
      await cpnft.connect(owner).mint(user1.address, 0);
      
      await cpnft.connect(user1).approve(staking.address, 7);
      
      await expect(
        staking.connect(user1).stake(7)
      ).to.be.revertedWith("NORMAL level NFTs cannot be staked");
    });

    it("Should not allow staking non-owned NFTs", async function () {
      await cpnft.connect(user1).approve(staking.address, 1);
      
      await expect(
        staking.connect(user2).stake(1)
      ).to.be.revertedWith("Not the owner of this NFT");
    });

    it("Should allow batch staking", async function () {
      // Approve all NFTs
      await cpnft.connect(user1).setApprovalForAll(staking.address, true);
      
      // Batch stake multiple NFTs
      await staking.connect(user1).batchStake([1, 2, 3]);
      
      // Check that all NFTs are staked
      for (let i = 1; i <= 3; i++) {
        const stakeInfo = await stakingReader.getStakeDetails(i);
        expect(stakeInfo.isActive).to.be.true;
      }
    });
  });

  describe("Reward Calculation", function () {
    beforeEach(async function () {
      // Mint and stake an SSS level NFT
      await cpnft.connect(owner).mint(user1.address, 6);
      await cpnft.connect(user1).approve(staking.address, 1);
      await staking.connect(user1).stake(1);
    });

    it("Should calculate pending rewards correctly", async function () {
      // Fast forward time by 1 day
      await ethers.provider.send("evm_increaseTime", [86400]); // 1 day
      await ethers.provider.send("evm_mine", []);

      const pendingRewards = await stakingReader.getPendingRewards(1);
      expect(pendingRewards).to.equal(100); // SSS level daily reward
    });

    it("Should apply decay after decay interval", async function () {
      // Fast forward time by 180 days (SSS decay interval)
      await ethers.provider.send("evm_increaseTime", [180 * 86400]);
      await ethers.provider.send("evm_mine", []);

      const pendingRewards = await stakingReader.getPendingRewards(1);
      // Should be decayed by 10% (1000 basis points) for 180 days
      // 180 days * 90 CPP/day = 16200 CPP
      expect(pendingRewards).to.equal(16200);
    });

    it("Should provide detailed reward calculation", async function () {
      const details = await stakingReader.getDetailedRewardCalculation(1);
      expect(details.baseReward).to.equal(100); // SSS level
      expect(details.decayedReward).to.equal(100); // No decay yet
      expect(details.comboBonus).to.equal(0); // No combo bonus
      expect(details.dynamicMultiplier).to.equal(10000); // 1.0x
    });
  });

  describe("Combo Bonuses", function () {
    beforeEach(async function () {
      // Mint 5 A level NFTs for combo testing
      for (let i = 1; i <= 5; i++) {
        await cpnft.connect(owner).mint(user1.address, 3); // A level
      }
      
      await cpnft.connect(user1).setApprovalForAll(staking.address, true);
    });

    it("Should apply combo bonus for 5 same-level NFTs", async function () {
      // Stake 5 A level NFTs
      await staking.connect(user1).batchStake([1, 2, 3, 4, 5]);
      
      // Fast forward time by 1 day
      await ethers.provider.send("evm_increaseTime", [86400]);
      await ethers.provider.send("evm_mine", []);

      const pendingRewards = await stakingReader.getPendingRewards(1);
      // A level base reward is 15, with 10% combo bonus
      expect(pendingRewards).to.equal(15); // 15 CPP per day (combo bonus applied in calculation)
    });

    it("Should show combo information correctly", async function () {
      await staking.connect(user1).batchStake([1, 2, 3, 4, 5]);
      
      const comboInfo = await stakingReader.getComboInfo(user1.address, 3); // A level
      expect(comboInfo.count).to.equal(5);
      expect(comboInfo.currentBonus).to.equal(1000); // 10%
    });
  });

  describe("Unstaking", function () {
    beforeEach(async function () {
      // Mint and stake an SSS level NFT
      await cpnft.connect(owner).mint(user1.address, 6);
      await cpnft.connect(user1).approve(staking.address, 1);
      await staking.connect(user1).stake(1);
      
      // Fast forward time by 10 days
      await ethers.provider.send("evm_increaseTime", [10 * 86400]);
      await ethers.provider.send("evm_mine", []);
    });

    it("Should allow unstaking after minimum period", async function () {
      const initialBalance = await cpopToken.balanceOf(user1.address);
      
      await staking.connect(user1).unstake(1);
      
      const finalBalance = await cpopToken.balanceOf(user1.address);
      expect(finalBalance).to.be.gt(initialBalance);
      
      // Check that NFT is no longer staked
      const stakeInfo = await stakingReader.getStakeDetails(1);
      expect(stakeInfo.isActive).to.be.false;
    });

    it("Should apply early withdrawal penalty", async function () {
      // Stake another NFT and immediately try to unstake
      await cpnft.connect(owner).mint(user1.address, 1);
      await cpnft.connect(user1).approve(staking.address, 2);
      await staking.connect(user1).stake(2);
      
      const initialBalance = await cpopToken.balanceOf(user1.address);
      
      await staking.connect(user1).unstake(2);
      
      const finalBalance = await cpopToken.balanceOf(user1.address);
      const balanceIncrease = finalBalance.sub(initialBalance);
      
      // Should be much less due to early withdrawal penalty
      expect(balanceIncrease).to.be.lt(ethers.utils.parseEther("1"));
    });
  });

  describe("Platform Statistics", function () {
    beforeEach(async function () {
      // Mint and stake various level NFTs
      await cpnft.connect(owner).mint(user1.address, 1); // C level
      await cpnft.connect(owner).mint(user1.address, 2); // B level
      await cpnft.connect(owner).mint(user1.address, 6); // SSS level
      
      await cpnft.connect(user1).setApprovalForAll(staking.address, true);
      await staking.connect(user1).batchStake([1, 2, 3]);
    });

    it("Should provide correct platform statistics", async function () {
      const stats = await stakingReader.getPlatformStats();
      
      expect(stats.staked[1]).to.equal(1); // 1 C level NFT staked
      expect(stats.staked[2]).to.equal(1); // 1 B level NFT staked
      expect(stats.staked[6]).to.equal(1); // 1 SSS level NFT staked
      
      expect(stats.supply[1]).to.equal(1000); // C level total supply
      expect(stats.supply[2]).to.equal(500);  // B level total supply
      expect(stats.supply[6]).to.equal(10);   // SSS level total supply
    });

    it("Should calculate level statistics correctly", async function () {
      const levelStats = await stakingReader.getLevelStats();
      
      // Check staking ratios
      expect(levelStats.stakingRatios[1]).to.equal(10); // 1/1000 * 10000 = 10
      expect(levelStats.stakingRatios[2]).to.equal(20); // 1/500 * 10000 = 20
      expect(levelStats.stakingRatios[6]).to.equal(1000); // 1/10 * 10000 = 1000
    });
  });

  describe("Configuration Management", function () {
    it("Should allow owner to update rewards", async function () {
      await stakingConfig.connect(owner).updateRewards([3, 8, 15, 30, 50, 100]); // C, B, A, S, SS, SSS
      
      const sssReward = await stakingConfig.getDailyReward(6);
      expect(sssReward).to.equal(100);
      
      const cReward = await stakingConfig.getDailyReward(1);
      expect(cReward).to.equal(3);
    });

    it("Should not allow non-owner to update config", async function () {
      await expect(
        stakingConfig.connect(user1).updateRewards([200, 100, 60, 30, 16, 6])
      ).to.be.revertedWith("OwnableUnauthorizedAccount");
    });

    it("Should provide configuration data", async function () {
      const config = await stakingReader.getConfiguration();
      
      expect(config.minStakeDays).to.equal(7);
      expect(config.earlyWithdrawPenalty).to.equal(5000);
      expect(config.dailyRewards[5]).to.equal(100); // SSS level
      expect(config.dailyRewards[0]).to.equal(3);   // C level
    });
  });

  describe("Error Handling", function () {
    it("Should revert when staking already staked NFT", async function () {
      await cpnft.connect(owner).mint(user1.address, 1);
      await cpnft.connect(user1).approve(staking.address, 1);
      await staking.connect(user1).stake(1);
      
      await expect(
        staking.connect(user1).stake(1)
      ).to.be.revertedWith("NFT is already staked");
    });

    it("Should revert when unstaking non-staked NFT", async function () {
      await expect(
        staking.connect(user1).unstake(999)
      ).to.be.revertedWith("Not the owner of this stake");
    });

    it("Should revert when claiming rewards for non-staked NFT", async function () {
      await expect(
        staking.connect(user1).claimRewards(999)
      ).to.be.revertedWith("Not the owner of this stake");
    });
  });
});
