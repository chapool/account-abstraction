import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";

describe("SS Level NFT Decay Test", function () {
    let owner: SignerWithAddress;
    let user1: SignerWithAddress;

    // Contracts
    let cpnft: any;
    let cpopToken: any;
    let accountManager: any;
    let stakingConfig: any;
    let staking: any;
    let stakingReader: any;

    const MINTER_ROLE = 2;

    beforeEach(async function () {
        [owner, user1] = await ethers.getSigners();

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
        
        // Set large supplies to avoid dynamic multiplier effects
        await cpnft.connect(owner).setLevelSupply(5, 10000); // SS level
    });

    describe("SS Level Decay Calculation", function () {
        it("Should calculate SS level decay phases correctly", async function () {
            console.log("\n=== SS级NFT分阶段衰减收益计算测试 ===");
            
            const testPhases = [
                { days: 90, expectedRate: 50, description: "第1-90天 (无衰减)" },
                { days: 180, expectedRate: 42.5, description: "第91-180天 (15%衰减)" },
                { days: 270, expectedRate: 36.1, description: "第181-270天 (累计30%衰减)" },
                { days: 360, expectedRate: 30, description: "第271-360天 (累计40%衰减，最低值)" }
            ];

            for (const phase of testPhases) {
                // Mint fresh SS-level NFT for each test
                const tokenId = await cpnft.connect(owner).callStatic.mint(user1.address, 5);
                await cpnft.connect(owner).mint(user1.address, 5);
                await staking.connect(user1).stake(tokenId);

                // Wait for the specified days
                await ethers.provider.send("evm_increaseTime", [phase.days * 24 * 60 * 60]);
                await ethers.provider.send("evm_mine", []);

                // Claim rewards
                await staking.connect(user1).claimRewards(tokenId);

                // Get total rewards
                const stakeInfo = await stakingReader.getStakeDetails(tokenId);
                const totalRewards = stakeInfo.totalRewards.toNumber();

                // Calculate expected rewards based on document
                let expectedTotal = 0;
                if (phase.days <= 90) {
                    // Phase 1: No decay
                    expectedTotal = 50 * phase.days;
                } else if (phase.days <= 180) {
                    // Phase 1 + Phase 2
                    expectedTotal = 50 * 90 + 42.5 * (phase.days - 90);
                } else if (phase.days <= 270) {
                    // Phase 1 + Phase 2 + Phase 3
                    expectedTotal = 50 * 90 + 42.5 * 90 + 36.1 * (phase.days - 180);
                } else {
                    // Phase 1 + Phase 2 + Phase 3 + Phase 4
                    expectedTotal = 50 * 90 + 42.5 * 90 + 36.1 * 90 + 30 * (phase.days - 270);
                }

                // Add continuous staking bonus (10% for 30+ days, 30% for 90+ days)
                let continuousBonus = 0;
                if (phase.days >= 90) {
                    continuousBonus = expectedTotal * 0.30; // 30% bonus for 90+ days
                } else if (phase.days >= 30) {
                    continuousBonus = expectedTotal * 0.10; // 10% bonus for 30+ days
                }
                expectedTotal += continuousBonus;

                const difference = Math.abs(totalRewards - expectedTotal);
                const percentageDiff = (difference / expectedTotal * 100).toFixed(2);

                console.log(`\n${phase.description}:`);
                console.log(`  质押天数: ${phase.days}天`);
                console.log(`  当前日收益率: ${phase.expectedRate} CPP/天`);
                console.log(`  实际总收益: ${totalRewards} CPP`);
                console.log(`  预期总收益: ${expectedTotal} CPP`);
                console.log(`  差异: ${difference} CPP (${percentageDiff}%)`);

                // Verify accuracy (allow 5% tolerance)
                expect(difference / expectedTotal).to.be.lessThan(0.05);

                // Clean up
                await staking.connect(user1).unstake(tokenId);
            }
        });

        it("Should verify SS level decay configuration", async function () {
            console.log("\n=== SS级NFT衰减配置验证 ===");
            
            // Get SS level configuration (level 5)
            const levelConfig = await stakingConfig.getLevelConfig(5);
            
            console.log(`SS级NFT配置:`);
            console.log(`  初始日收益: ${levelConfig[0]} CPP`);
            console.log(`  衰减周期: ${levelConfig[1]} 天`);
            console.log(`  衰减率: ${levelConfig[2]} (${levelConfig[2]/100}%)`);
            console.log(`  最大衰减率: ${levelConfig[3]} (${levelConfig[3]/100}%)`);

            // Verify configuration matches document
            expect(levelConfig[0]).to.equal(50); // Initial daily reward: 50 CPP
            expect(levelConfig[1]).to.equal(90); // Decay interval: 90 days
            expect(levelConfig[2]).to.equal(1500); // Decay rate: 15% (1500/10000)
            expect(levelConfig[3]).to.equal(4000); // Max decay rate: 40% (4000/10000)
        });

        it("Should test daily reward progression for SS level", async function () {
            console.log("\n=== SS级NFT日收益递减验证 ===");
            
            // Test phase transitions
            const testCases = [
                { days: 90, description: "第90天 (无衰减期结束)" },
                { days: 180, description: "第180天 (15%衰减期结束)" },
                { days: 270, description: "第270天 (30%衰减期结束)" }
            ];

            for (const testCase of testCases) {
                // Mint fresh SS-level NFT for each test
                const tokenId = await cpnft.connect(owner).callStatic.mint(user1.address, 5);
                await cpnft.connect(owner).mint(user1.address, 5);
                await staking.connect(user1).stake(tokenId);

                // Wait for the specified days
                await ethers.provider.send("evm_increaseTime", [testCase.days * 24 * 60 * 60]);
                await ethers.provider.send("evm_mine", []);

                // Get pending rewards
                const pendingRewards = await staking.calculatePendingRewards(tokenId);
                const avgDailyReward = pendingRewards.toNumber() / testCase.days;

                console.log(`${testCase.description}:`);
                console.log(`  实际平均日收益: ${avgDailyReward.toFixed(2)} CPP`);
                console.log(`  总质押天数: ${testCase.days}天`);

                // Clean up
                await staking.connect(user1).unstake(tokenId);
            }
        });
    });
});
