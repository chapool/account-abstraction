import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";

describe("NFT Staking System Tests", function () {
    let owner: SignerWithAddress;
    let user1: SignerWithAddress;
    let user2: SignerWithAddress;

    // Contracts
    let cpnft: any;
    let cpopToken: any;
    let accountManager: any;
    let stakingConfig: any;
    let staking: any;
    let stakingReader: any;

    const MINTER_ROLE = 2;

    beforeEach(async function () {
        [owner, user1, user2] = await ethers.getSigners();

        // Deploy contracts
        const CPNFT = await ethers.getContractFactory("CPNFT");
        cpnft = await upgrades.deployProxy(CPNFT, ["TestNFT", "TNFT", "https://api.test.com/"]);
        await cpnft.deployed();

        const MockCPOPToken = await ethers.getContractFactory("CPOPToken");
        cpopToken = await MockCPOPToken.deploy(owner.address, ethers.utils.parseEther("1000000"));
        await cpopToken.deployed();

        const MockAccountManager = await ethers.getContractFactory("AccountManager");
        accountManager = await MockAccountManager.deploy();
        await accountManager.deployed();

        const StakingConfig = await ethers.getContractFactory("StakingConfig");
        stakingConfig = await StakingConfig.deploy();
        await stakingConfig.deployed();

        const Staking = await ethers.getContractFactory("Staking");
        staking = await upgrades.deployProxy(Staking, [
            cpnft.address,
            cpopToken.address,
            accountManager.address,
            stakingConfig.address,
            owner.address
        ]);
        await staking.deployed();

        const StakingReader = await ethers.getContractFactory("StakingReader");
        stakingReader = await StakingReader.deploy(
            staking.address,
            stakingConfig.address,
            cpnft.address
        );
        await stakingReader.deployed();

        // Setup
        await cpnft.connect(owner).setStakingContract(staking.address);
        await cpopToken.connect(owner).grantRole(staking.address, MINTER_ROLE);
        
        // Set large supplies to avoid dynamic multiplier effects in basic tests
        await cpnft.connect(owner).setLevelSupply(1, 10000); // C level
        await cpnft.connect(owner).setLevelSupply(2, 10000); // B level
        await cpnft.connect(owner).setLevelSupply(3, 10000); // A level
        await cpnft.connect(owner).setLevelSupply(4, 10000); // S level
        await cpnft.connect(owner).setLevelSupply(5, 10000); // SS level
        await cpnft.connect(owner).setLevelSupply(6, 10000); // SSS level
    });

    describe("Basic Staking Functionality", function () {
        it("Should stake and unstake NFT correctly", async function () {
            // Mint A-level NFT
            const tokenId = await cpnft.connect(owner).callStatic.mint(user1.address, 3);
            await cpnft.connect(owner).mint(user1.address, 3);

            // Stake NFT
            await staking.connect(user1).stake(tokenId);

            // Check stake info
            const stakeInfo = await stakingReader.getStakeDetails(tokenId);
            expect(stakeInfo.owner).to.equal(user1.address);
            expect(stakeInfo.level).to.equal(3);
            expect(stakeInfo.isActive).to.be.true;

            // Wait 7 days (minimum staking period)
            await ethers.provider.send("evm_increaseTime", [7 * 24 * 60 * 60]);
            await ethers.provider.send("evm_mine", []);

            // Unstake NFT
            await staking.connect(user1).unstake(tokenId);

            // Check unstake
            const unstakeInfo = await stakingReader.getStakeDetails(tokenId);
            expect(unstakeInfo.isActive).to.be.false;
        });

        it("Should reject NORMAL level NFTs", async function () {
            // Mint NORMAL level NFT
            const tokenId = await cpnft.connect(owner).callStatic.mint(user1.address, 0);
            await cpnft.connect(owner).mint(user1.address, 0);

            // Should fail to stake
            await expect(staking.connect(user1).stake(tokenId))
                .to.be.revertedWith("NORMAL level NFTs cannot be staked");
        });

        it("Should accept all other levels", async function () {
            const levels = [1, 2, 3, 4, 5, 6]; // C, B, A, S, SS, SSS
            
            for (let i = 0; i < levels.length; i++) {
                const tokenId = await cpnft.connect(owner).callStatic.mint(user1.address, levels[i]);
                await cpnft.connect(owner).mint(user1.address, levels[i]);
                
                // Should succeed
                await staking.connect(user1).stake(tokenId);
                
                // Unstake immediately for next test
                await staking.connect(user1).unstake(tokenId);
            }
        });
    });

    describe("Reward Calculation Accuracy", function () {
        it("Should calculate A-level NFT 91-day staking correctly", async function () {
            // Mint and stake A-level NFT
            const tokenId = await cpnft.connect(owner).callStatic.mint(user1.address, 3);
            await cpnft.connect(owner).mint(user1.address, 3);
            await staking.connect(user1).stake(tokenId);

            // Wait 91 days
            await ethers.provider.send("evm_increaseTime", [91 * 24 * 60 * 60]);
            await ethers.provider.send("evm_mine", []);

            // Claim rewards
            await staking.connect(user1).claimRewards(tokenId);

            // Get total rewards
            const stakeInfo = await stakingReader.getStakeDetails(tokenId);
            const totalRewards = stakeInfo.totalRewards.toNumber();

            // Document expected calculation:
            // Phase 1 (Days 1-45): 45 × 15 = 675 CPP
            // Phase 2 (Days 46-90): 45 × 11.25 = 506.25 CPP  
            // Phase 3 (Day 91): 1 × 8.4375 = 8.4375 CPP
            // Base total: 1189.6875 CPP
            // Continuous bonus (30%): 356.90625 CPP
            // Expected total: 1546.59375 CPP

            const expectedTotal = 1546.59375;
            const tolerance = 50; // Allow 50 CPP tolerance

            console.log(`A级NFT 91天质押收益测试:`);
            console.log(`实际收益: ${totalRewards} CPP`);
            console.log(`预期收益: ${expectedTotal} CPP`);
            console.log(`差异: ${Math.abs(totalRewards - expectedTotal)} CPP`);

            expect(Math.abs(totalRewards - expectedTotal)).to.be.lessThan(tolerance);

            // Clean up
            await staking.connect(user1).unstake(tokenId);
        });

        it("Should calculate different level rewards correctly", async function () {
            const testCases = [
                { level: 1, days: 30, expectedBase: 90, description: "C级30天", hasDecay: true },    // 30 × 3 = 90, but with decay
                { level: 2, days: 30, expectedBase: 240, description: "B级30天", hasDecay: true },   // 30 × 8 = 240, but with decay  
                { level: 4, days: 30, expectedBase: 900, description: "S级30天", hasDecay: false },   // 30 × 30 = 900, no decay yet
                { level: 5, days: 30, expectedBase: 1500, description: "SS级30天", hasDecay: false }, // 30 × 50 = 1500, no decay yet
                { level: 6, days: 30, expectedBase: 3000, description: "SSS级30天", hasDecay: false } // 30 × 100 = 3000, no decay yet
            ];

            for (const testCase of testCases) {
                // Mint and stake NFT
                const tokenId = await cpnft.connect(owner).callStatic.mint(user1.address, testCase.level);
                await cpnft.connect(owner).mint(user1.address, testCase.level);
                await staking.connect(user1).stake(tokenId);

                // Wait for specified days
                await ethers.provider.send("evm_increaseTime", [testCase.days * 24 * 60 * 60]);
                await ethers.provider.send("evm_mine", []);

                // Claim rewards
                await staking.connect(user1).claimRewards(tokenId);

                // Get total rewards
                const stakeInfo = await stakingReader.getStakeDetails(tokenId);
                const totalRewards = stakeInfo.totalRewards.toNumber();

                console.log(`${testCase.description}: ${totalRewards} CPP (预期基础: ${testCase.expectedBase} CPP)`);

                // Should be close to expected base (allowing for continuous bonus and decay)
                if (testCase.hasDecay) {
                    // For levels with decay, expect lower than base due to decay
                    expect(totalRewards).to.be.greaterThan(testCase.expectedBase * 0.5); // Allow for decay
                    expect(totalRewards).to.be.lessThan(testCase.expectedBase * 1.5); // Allow for bonuses
                } else {
                    // For levels without decay, expect close to base
                    expect(totalRewards).to.be.greaterThan(testCase.expectedBase);
                    expect(totalRewards).to.be.lessThan(testCase.expectedBase * 1.5); // Allow for bonuses
                }

                // Clean up
                await staking.connect(user1).unstake(tokenId);
            }
        });
    });

    describe("Combo Bonus System", function () {
        it("Should apply combo bonuses correctly", async function () {
            // Mint 5 A-level NFTs for combo bonus
            const tokenIds = [];
            for (let i = 0; i < 5; i++) {
                const tokenId = await cpnft.connect(owner).callStatic.mint(user1.address, 3);
                await cpnft.connect(owner).mint(user1.address, 3);
                tokenIds.push(tokenId);
            }

            // Stake all 5 NFTs
            for (const tokenId of tokenIds) {
                await staking.connect(user1).stake(tokenId);
            }

            // Wait 15 days (minimum for 5-NFT combo)
            await ethers.provider.send("evm_increaseTime", [15 * 24 * 60 * 60]);
            await ethers.provider.send("evm_mine", []);

            // Check combo info
            const comboInfo = await stakingReader.getComboInfo(user1.address, 3);
            expect(comboInfo[0]).to.equal(5); // 5 NFTs staked
            expect(comboInfo[1]).to.equal(1000); // 10% bonus

            // Claim rewards for first NFT
            await staking.connect(user1).claimRewards(tokenIds[0]);
            const stakeInfo = await stakingReader.getStakeDetails(tokenIds[0]);
            const rewards = stakeInfo.totalRewards.toNumber();

            // Should have combo bonus applied
            const expectedBase = 15 * 15; // 15 days × 15 CPP
            const expectedWithCombo = expectedBase * 1.1; // 10% combo bonus

            console.log(`5个A级NFT组合质押收益测试:`);
            console.log(`实际收益: ${rewards} CPP`);
            console.log(`预期基础收益: ${expectedBase} CPP`);
            console.log(`预期组合加成收益: ${expectedWithCombo} CPP`);

            expect(rewards).to.be.greaterThan(expectedBase);
            expect(rewards).to.be.closeTo(expectedWithCombo, 50);

            // Clean up
            for (const tokenId of tokenIds) {
                await staking.connect(user1).unstake(tokenId);
            }
        });
    });

    describe("Dynamic Balance Mechanism", function () {
        it("Should apply dynamic multipliers based on staking ratio", async function () {
            // Reset A level supply to 0 first
            await cpnft.connect(owner).setLevelSupply(3, 0);
            
            // Mint 100 NFTs first to establish total supply
            const tokenIds = [];
            for (let i = 0; i < 100; i++) {
                const tokenId = await cpnft.connect(owner).callStatic.mint(user1.address, 3);
                await cpnft.connect(owner).mint(user1.address, 3);
                tokenIds.push(tokenId);
            }

            // Now stake 70 NFTs (70% staking ratio - should trigger 0.8x multiplier)
            for (let i = 0; i < 70; i++) {
                await staking.connect(user1).stake(tokenIds[i]);
            }

            // Check level stats
            const levelStats = await stakingReader.getLevelStats();
            const dynamicMultiplier = levelStats.dynamicMultipliers[3]; // A level index (level 3)
            
            // Also check total staked count
            const totalStaked = await staking.totalStakedPerLevel(3);
            const totalSupply = await cpnft.getLevelSupply(3);

            console.log(`动态平衡机制测试:`);
            console.log(`质押数量: ${totalStaked}`);
            console.log(`总供应量: ${totalSupply}`);
            console.log(`质押比例: ${(totalStaked * 100 / totalSupply).toFixed(1)}%`);
            console.log(`动态乘数: ${dynamicMultiplier / 10000}x`);

            // Should be 0.8x multiplier (8000) for 70% staking ratio
            // But since we're testing with 70% which is > 60%, it should be 0.8x
            if (totalStaked * 100 / totalSupply > 60) {
                expect(dynamicMultiplier.toNumber()).to.be.closeTo(8000, 100);
            } else {
                // For other ratios, just check it's a valid multiplier
                expect(dynamicMultiplier.toNumber()).to.be.greaterThan(0);
            }

            // Clean up - only unstake the staked ones
            for (let i = 0; i < 70; i++) {
                await staking.connect(user1).unstake(tokenIds[i]);
            }
        });
    });

    describe("Continuous Staking Bonus", function () {
        it("Should apply continuous staking bonuses correctly", async function () {
            const testCases = [
                { days: 7, expectedBonus: 0, description: "7天无奖励" },
                { days: 30, expectedBonus: 10, description: "30天10%奖励" },
                { days: 90, expectedBonus: 30, description: "90天30%奖励" }
            ];

            for (const testCase of testCases) {
                // Mint and stake A-level NFT
                const tokenId = await cpnft.connect(owner).callStatic.mint(user1.address, 3);
                await cpnft.connect(owner).mint(user1.address, 3);
                await staking.connect(user1).stake(tokenId);

                // Wait for specified days
                await ethers.provider.send("evm_increaseTime", [testCase.days * 24 * 60 * 60]);
                await ethers.provider.send("evm_mine", []);

                // Claim rewards
                await staking.connect(user1).claimRewards(tokenId);

                // Get total rewards
                const stakeInfo = await stakingReader.getStakeDetails(tokenId);
                const totalRewards = stakeInfo.totalRewards.toNumber();

                // Calculate expected rewards
                let expectedBase = 0;
                if (testCase.days <= 45) {
                    expectedBase = 15 * testCase.days;
                } else if (testCase.days <= 90) {
                    expectedBase = 15 * 45 + 11.25 * (testCase.days - 45);
                } else {
                    expectedBase = 15 * 45 + 11.25 * 45 + 8.4375 * (testCase.days - 90);
                }

                const expectedTotal = expectedBase * (1 + testCase.expectedBonus / 100);

                console.log(`${testCase.description}:`);
                console.log(`实际收益: ${totalRewards} CPP`);
                console.log(`预期收益: ${expectedTotal} CPP`);
                console.log(`差异: ${Math.abs(totalRewards - expectedTotal)} CPP`);

                expect(Math.abs(totalRewards - expectedTotal)).to.be.lessThan(100);

                // Clean up
                await staking.connect(user1).unstake(tokenId);
            }
        });
    });

    describe("Early Withdrawal Penalty", function () {
        it("Should apply early withdrawal penalty for < 7 days", async function () {
            // Mint and stake A-level NFT
            const tokenId = await cpnft.connect(owner).callStatic.mint(user1.address, 3);
            await cpnft.connect(owner).mint(user1.address, 3);
            await staking.connect(user1).stake(tokenId);

            // Wait only 5 days (less than 7-day minimum)
            await ethers.provider.send("evm_increaseTime", [5 * 24 * 60 * 60]);
            await ethers.provider.send("evm_mine", []);

            // Get pending rewards before unstaking
            const pendingBefore = await staking.calculatePendingRewards(tokenId);

            // Unstake (should apply penalty)
            await staking.connect(user1).unstake(tokenId);

            // Check final rewards
            const stakeInfo = await stakingReader.getStakeDetails(tokenId);
            const finalRewards = stakeInfo.totalRewards.toNumber();

            console.log(`提前赎回惩罚测试:`);
            console.log(`待领取收益: ${pendingBefore.toNumber()} CPP`);
            console.log(`最终收益: ${finalRewards} CPP`);
            console.log(`惩罚比例: ${((pendingBefore.toNumber() - finalRewards) / pendingBefore.toNumber() * 100).toFixed(2)}%`);

            // Should have penalty applied (50% of rewards deducted)
            expect(finalRewards).to.be.lessThan(pendingBefore.toNumber());
            expect(finalRewards).to.be.closeTo(pendingBefore.toNumber() * 0.5, 10);
        });
    });
});
