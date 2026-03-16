/**
 * Earn System Unit Tests — Sample Data Scenarios
 *
 * Based on docs/Earn-Feature-Spec.md §7.5:
 * - rewardRate ≈ 8.267 CPP/sec (platform 10,000 USDT/week)
 * - totalWeightedUSDT = 1,000,000
 * - 4 users A/B/C/D with different deposits and boosts
 */
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { Contract } from "ethers";

const parse = (n: string) => ethers.utils.parseEther(n);
const DAY = 24 * 60 * 60;

describe("Earn System — Sample Data Scenarios", function () {
  let owner: SignerWithAddress;
  let userA: SignerWithAddress;
  let userB: SignerWithAddress;
  let userC: SignerWithAddress;
  let userD: SignerWithAddress;
  let filler: SignerWithAddress;

  let usdt: Contract;
  let cppToken: Contract;
  let cpot: Contract;
  let cpnft: Contract;
  let accountManager: Contract;
  let vault: Contract;
  let locker: Contract;
  let nftBoost: Contract;
  let reader: Contract;
  let distributor: Contract;

  // rewardRate ≈ 8.267 CPP/sec (5M CPP per week)
  const REWARD_RATE_WEI = ethers.utils.parseEther("8.267");

  // Sample user params (from spec §7.5)
  const USER_A_USDT = parse("500");
  const USER_B_USDT = parse("5000");
  const USER_C_USDT = parse("50000");
  const USER_D_USDT = parse("10000");
  const ONE_M = parse("1000000");
  const FILLER_FOR_A = ONE_M.sub(USER_A_USDT);          // 999,500
  const FILLER_FOR_B = ONE_M.sub(parse("5062.5"));      // 994,937.5 (B: veUnits=1250→125bps)
  const FILLER_FOR_C = ONE_M.sub(parse("54250"));       // 945,750 (C: 500+350=850bps, unchanged)
  const FILLER_FOR_D = ONE_M.sub(parse("10520"));       // 989,480 (D: 500+20=520bps)

  beforeEach(async function () {
    [owner, userA, userB, userC, userD, filler] = await ethers.getSigners();

    // Mocks
    const MockUSDT = await ethers.getContractFactory("MockUSDT");
    usdt = await MockUSDT.deploy();
    await usdt.deployed();

    const CPPToken = await ethers.getContractFactory("CPPToken");
    cppToken = await CPPToken.deploy(owner.address, parse("0"));
    await cppToken.deployed();

    const MockCPOT = await ethers.getContractFactory("MockCPOT");
    cpot = await MockCPOT.deploy();
    await cpot.deployed();

    const EarnMockAccountManager = await ethers.getContractFactory("EarnMockAccountManager");
    accountManager = await EarnMockAccountManager.deploy();
    await accountManager.deployed();

    const MockCPNFTForEarn = await ethers.getContractFactory("MockCPNFTForEarn");
    cpnft = await MockCPNFTForEarn.deploy();
    await cpnft.deployed();

    // ChapoolEarnVault
    const ChapoolEarnVault = await ethers.getContractFactory("ChapoolEarnVault");
    vault = await upgrades.deployProxy(
      ChapoolEarnVault,
      [usdt.address, cppToken.address, accountManager.address, owner.address],
      { kind: "uups", initializer: "initialize" }
    );
    await vault.deployed();
    await cppToken.grantRole(vault.address, 2); // MINTER_ROLE

    // VeCPOTLocker
    const VeCPOTLocker = await ethers.getContractFactory("VeCPOTLocker");
    locker = await upgrades.deployProxy(VeCPOTLocker, [cpot.address, owner.address], {
      kind: "uups",
      initializer: "initialize",
    });
    await locker.deployed();
    await locker.setVault(vault.address);

    // NFTBoostController
    const NFTBoostController = await ethers.getContractFactory("NFTBoostController");
    nftBoost = await upgrades.deployProxy(NFTBoostController, [cpnft.address, owner.address], {
      kind: "uups",
      initializer: "initialize",
    });
    await nftBoost.deployed();

    vault.setVecpotLocker(locker.address);
    vault.setNftBoostController(nftBoost.address);
    vault.setRewardRate(REWARD_RATE_WEI);

    // ChapoolRewardDistributor (for reader)
    const ChapoolRewardDistributor = await ethers.getContractFactory("ChapoolRewardDistributor");
    distributor = await upgrades.deployProxy(ChapoolRewardDistributor, [owner.address], {
      kind: "uups",
      initializer: "initialize",
    });
    await distributor.deployed();
    await distributor.setEarnVault(vault.address);

    // ChapoolVaultReader
    const ChapoolVaultReader = await ethers.getContractFactory("ChapoolVaultReader");
    reader = await ChapoolVaultReader.deploy(
      vault.address,
      locker.address,
      nftBoost.address,
      distributor.address,
      owner.address
    );
    await reader.deployed();

    // Mint USDT to users
    await usdt.connect(owner).mint(userA.address, parse("100000"));
    await usdt.connect(owner).mint(userB.address, parse("100000"));
    await usdt.connect(owner).mint(userC.address, parse("100000"));
    await usdt.connect(owner).mint(userD.address, parse("100000"));
    await usdt.connect(owner).mint(filler.address, parse("1000000"));
    await usdt.connect(owner).mint(owner.address, parse("1000000"));

    // Mint CPOT to users (for locking)
    await cpot.connect(owner).transfer(userB.address, parse("10000"));
    await cpot.connect(owner).transfer(userC.address, parse("100000"));
    await cpot.connect(owner).transfer(userD.address, parse("50000"));
  });

  describe("User A — 500 USDT, no boost", function () {
    it("should have weightedUSDT = 500", async function () {
      await usdt.connect(userA).approve(vault.address, USER_A_USDT);
      await vault.connect(userA).deposit(USER_A_USDT, userA.address);
      const w = await vault.weightedUSDT(userA.address);
      expect(w).to.equal(USER_A_USDT);
    });

    it("should accrue ~357 CPP per day (with totalWeightedUSDT ≈ 1M)", async function () {
      await usdt.connect(filler).approve(vault.address, FILLER_FOR_A);
      await vault.connect(filler).deposit(FILLER_FOR_A, filler.address);

      await usdt.connect(userA).approve(vault.address, USER_A_USDT);
      await vault.connect(userA).deposit(USER_A_USDT, userA.address);

      await ethers.provider.send("evm_increaseTime", [DAY]);
      await ethers.provider.send("evm_mine", []);

      const pending = await vault.getPendingCPP(userA.address);
      const expected = 357;
      expect(pending).to.be.gte(parse(String(expected - 10)));
      expect(pending).to.be.lte(parse(String(expected + 50)));
    });
  });

  describe("User B — 5,000 USDT, lock 5,000 CPOT × 90d (+1.25% veCPOT boost)", function () {
    it("should have weightedUSDT = 5,062.5 after locking (veUnits=1,250 → 125bps)", async function () {
      await usdt.connect(userB).approve(vault.address, USER_B_USDT);
      await vault.connect(userB).deposit(USER_B_USDT, userB.address);

      await cpot.connect(userB).approve(locker.address, parse("10000"));
      await locker.connect(userB).lock(parse("5000"), 90);

      const w = await vault.weightedUSDT(userB.address);
      expect(w).to.equal(parse("5062.5"));
    });

    it("should accrue ~3,615 CPP per day", async function () {
      await usdt.connect(filler).approve(vault.address, FILLER_FOR_B);
      await vault.connect(filler).deposit(FILLER_FOR_B, filler.address);

      await usdt.connect(userB).approve(vault.address, USER_B_USDT);
      await vault.connect(userB).deposit(USER_B_USDT, userB.address);
      await cpot.connect(userB).approve(locker.address, parse("10000"));
      await locker.connect(userB).lock(parse("5000"), 90);

      await ethers.provider.send("evm_increaseTime", [DAY]);
      await ethers.provider.send("evm_mine", []);

      const pending = await vault.getPendingCPP(userB.address);
      const expected = 3615;
      expect(pending).to.be.gte(parse(String(expected - 50)));
      expect(pending).to.be.lte(parse(String(expected + 100)));
    });
  });

  describe("User C — 50,000 USDT, lock 360d + NFT SS (+8.5% boost)", function () {
    it("should have weightedUSDT = 54,250", async function () {
      await usdt.connect(userC).approve(vault.address, USER_C_USDT);
      await vault.connect(userC).deposit(USER_C_USDT, userC.address);
      await cpot.connect(userC).approve(locker.address, parse("100000"));
      await locker.connect(userC).lock(parse("50000"), 360);

      await cpnft.connect(owner).mint(userC.address, 100, 5); // tokenId 100, level 5 = SS
      await nftBoost.connect(userC).activateBoost(100);
      await vault.connect(userC).syncBoost(userC.address);

      const w = await vault.weightedUSDT(userC.address);
      expect(w).to.equal(parse("54250"));
    });

    it("should accrue ~38,720 CPP per day", async function () {
      await usdt.connect(filler).approve(vault.address, FILLER_FOR_C);
      await vault.connect(filler).deposit(FILLER_FOR_C, filler.address);

      await usdt.connect(userC).approve(vault.address, USER_C_USDT);
      await vault.connect(userC).deposit(USER_C_USDT, userC.address);
      await cpot.connect(userC).approve(locker.address, parse("100000"));
      await locker.connect(userC).lock(parse("50000"), 360);
      await cpnft.connect(owner).mint(userC.address, 100, 5);
      await nftBoost.connect(userC).activateBoost(100);
      await vault.connect(userC).syncBoost(userC.address);

      await ethers.provider.send("evm_increaseTime", [DAY]);
      await ethers.provider.send("evm_mine", []);

      const pending = await vault.getPendingCPP(userC.address);
      const expected = 38720;
      expect(pending).to.be.gte(parse(String(expected - 500)));
      expect(pending).to.be.lte(parse(String(expected + 1000)));
    });
  });

  describe("User D — 10,000 USDT, lock 20,000 CPOT × 180d + NFT A (+5.2% boost)", function () {
    it("should have weightedUSDT = 10,520 (veUnits=10,000 → 500bps cap + NFT A 20bps)", async function () {
      await usdt.connect(userD).approve(vault.address, USER_D_USDT);
      await vault.connect(userD).deposit(USER_D_USDT, userD.address);
      await cpot.connect(userD).approve(locker.address, parse("50000"));
      await locker.connect(userD).lock(parse("20000"), 180);

      await cpnft.connect(owner).mint(userD.address, 200, 3); // level 3 = A
      await nftBoost.connect(userD).activateBoost(200);
      await vault.connect(userD).syncBoost(userD.address);

      const w = await vault.weightedUSDT(userD.address);
      expect(w).to.equal(parse("10520"));
    });

    it("should accrue ~7,511 CPP per day", async function () {
      await usdt.connect(filler).approve(vault.address, FILLER_FOR_D);
      await vault.connect(filler).deposit(FILLER_FOR_D, filler.address);

      await usdt.connect(userD).approve(vault.address, USER_D_USDT);
      await vault.connect(userD).deposit(USER_D_USDT, userD.address);
      await cpot.connect(userD).approve(locker.address, parse("50000"));
      await locker.connect(userD).lock(parse("20000"), 180);
      await cpnft.connect(owner).mint(userD.address, 200, 3);
      await nftBoost.connect(userD).activateBoost(200);
      await vault.connect(userD).syncBoost(userD.address);

      await ethers.provider.send("evm_increaseTime", [DAY]);
      await ethers.provider.send("evm_mine", []);

      const pending = await vault.getPendingCPP(userD.address);
      const expected = 7511;
      expect(pending).to.be.gte(parse(String(expected - 100)));
      expect(pending).to.be.lte(parse(String(expected + 200)));
    });
  });

  describe("claimCPP", function () {
    it("should mint CPP to user AA account (EarnMockAccountManager returns user)", async function () {
      await usdt.connect(filler).approve(vault.address, FILLER_FOR_A);
      await vault.connect(filler).deposit(FILLER_FOR_A, filler.address);
      await usdt.connect(userA).approve(vault.address, USER_A_USDT);
      await vault.connect(userA).deposit(USER_A_USDT, userA.address);

      await ethers.provider.send("evm_increaseTime", [DAY]);
      await ethers.provider.send("evm_mine", []);

      const before = await cppToken.balanceOf(userA.address);
      await vault.connect(userA).claimCPP();
      const after = await cppToken.balanceOf(userA.address);
      expect(after.sub(before)).to.be.gte(parse("350"));
    });
  });

  describe("Staked NFT cannot boost", function () {
    it("should return 0 boost when NFT is staked (activate first, then stake)", async function () {
      await cpnft.connect(owner).mint(userA.address, 301, 4); // S level (index 4)
      await nftBoost.connect(userA).activateBoost(301);
      expect(await nftBoost.getNFTBoostBps(userA.address)).to.equal(50); // S = 50 bps
      await cpnft.connect(owner).setStakeStatus(301, true);
      const bps = await nftBoost.getNFTBoostBps(userA.address);
      expect(bps).to.equal(0);  // staked → no boost
    });

    it("should reject activateBoost for staked NFT", async function () {
      await cpnft.connect(owner).mint(userA.address, 302, 4);
      await cpnft.connect(owner).setStakeStatus(302, true);
      await expect(nftBoost.connect(userA).activateBoost(302)).to.be.revertedWith("NFT is staked");
    });
  });

  describe("ChapoolVaultReader", function () {
    it("getUserDashboard should return correct fields for user B", async function () {
      await usdt.connect(userB).approve(vault.address, USER_B_USDT);
      await vault.connect(userB).deposit(USER_B_USDT, userB.address);
      await cpot.connect(userB).approve(locker.address, parse("10000"));
      await locker.connect(userB).lock(parse("5000"), 90);

      const dash = await reader.getUserDashboard(userB.address);
      expect(dash.usdtBalance).to.equal(USER_B_USDT);
      expect(dash.weightedUSDT).to.equal(parse("5062.5"));
      expect(dash.vecpotBoostBps).to.equal(125); // veUnits=1250, (1250×1)/10 = 125 bps
    });
  });

  describe("Emergency withdrawal", function () {
    const EMERGENCY_TIMELOCK = 24 * 60 * 60; // 24 hours
    const emergencyAmount = parse("1000");

    it("should revert initiateEmergencyWithdraw when not in emergency mode", async function () {
      await usdt.connect(userA).approve(vault.address, emergencyAmount);
      await vault.connect(userA).deposit(emergencyAmount, userA.address);
      await expect(
        vault.connect(owner).initiateEmergencyWithdraw(emergencyAmount, owner.address)
      ).to.be.revertedWith("Not in emergency mode");
    });

    it("should set emergency mode and pause deposits", async function () {
      await vault.connect(owner).setEmergencyMode(true);
      expect(await vault.emergencyMode()).to.equal(true);
      expect(await vault.paused()).to.equal(true);
      await expect(
        vault.connect(userB).deposit(parse("100"), userB.address)
      ).to.be.reverted;
    });

    it("should initiate emergency withdraw and emit event", async function () {
      await usdt.connect(userA).approve(vault.address, emergencyAmount);
      await vault.connect(userA).deposit(emergencyAmount, userA.address);
      await vault.connect(owner).setEmergencyMode(true);

      const tx = await vault.connect(owner).initiateEmergencyWithdraw(emergencyAmount, owner.address);
      const receipt = await tx.wait();
      const executeAfter = (await ethers.provider.getBlock(receipt.blockNumber)).timestamp + EMERGENCY_TIMELOCK;
      await expect(tx).to.emit(vault, "EmergencyWithdrawInitiated").withArgs(emergencyAmount, owner.address, executeAfter);

      const pw = await vault.pendingEmergencyWithdraw();
      expect(pw.amount).to.equal(emergencyAmount);
      expect(pw.to).to.equal(owner.address);
      expect(pw.pending).to.equal(true);
    });

    it("should revert executeEmergencyWithdraw before timelock", async function () {
      await usdt.connect(userA).approve(vault.address, emergencyAmount);
      await vault.connect(userA).deposit(emergencyAmount, userA.address);
      await vault.connect(owner).setEmergencyMode(true);
      await vault.connect(owner).initiateEmergencyWithdraw(emergencyAmount, owner.address);

      await expect(vault.connect(owner).executeEmergencyWithdraw()).to.be.revertedWith(
        "Timelock not elapsed"
      );
    });

    it("should allow cancelEmergencyWithdraw and clear pending", async function () {
      await usdt.connect(userA).approve(vault.address, emergencyAmount);
      await vault.connect(userA).deposit(emergencyAmount, userA.address);
      await vault.connect(owner).setEmergencyMode(true);
      await vault.connect(owner).initiateEmergencyWithdraw(emergencyAmount, owner.address);

      await vault.connect(owner).cancelEmergencyWithdraw();
      const pw = await vault.pendingEmergencyWithdraw();
      expect(pw.pending).to.equal(false);
      expect(pw.amount).to.equal(0);
    });

    it("should initiate and execute after timelock", async function () {
      await usdt.connect(userA).approve(vault.address, emergencyAmount);
      await vault.connect(userA).deposit(emergencyAmount, userA.address);
      await vault.connect(owner).setEmergencyMode(true);
      await vault.connect(owner).initiateEmergencyWithdraw(emergencyAmount, owner.address);

      await ethers.provider.send("evm_increaseTime", [EMERGENCY_TIMELOCK + 1]);
      await ethers.provider.send("evm_mine", []);

      const ownerBefore = await usdt.balanceOf(owner.address);
      await expect(vault.connect(owner).executeEmergencyWithdraw())
        .to.emit(vault, "EmergencyWithdrawExecuted");
      const ownerAfter = await usdt.balanceOf(owner.address);
      expect(ownerAfter.sub(ownerBefore)).to.equal(emergencyAmount);
    });

    it("should revert execute when no pending withdraw", async function () {
      await expect(vault.connect(owner).executeEmergencyWithdraw()).to.be.revertedWith(
        "No pending withdraw"
      );
    });

    it("should allow setEmergencyMode(false) and clear pending if any", async function () {
      await usdt.connect(userD).approve(vault.address, parse("500"));
      await vault.connect(userD).deposit(parse("500"), userD.address);
      await vault.connect(owner).setEmergencyMode(true);
      await vault.connect(owner).initiateEmergencyWithdraw(parse("500"), userD.address);
      await vault.connect(owner).setEmergencyMode(false);
      expect(await vault.emergencyMode()).to.equal(false);
      expect(await vault.paused()).to.equal(false);
      const pw = await vault.pendingEmergencyWithdraw();
      expect(pw.pending).to.equal(false);
    });
  });
});
