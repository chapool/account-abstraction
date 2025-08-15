import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import {
    MasterAggregator,
    AAWallet,
    WalletManager,
    EntryPoint,
    GasPaymaster,
    CPOPToken,
    MasterAggregator__factory,
    AAWallet__factory,
    WalletManager__factory,
    EntryPoint__factory,
    GasPaymaster__factory,
    CPOPToken__factory,
} from "../typechain";
import { PackedUserOperation } from "./utils/UserOperation";

describe("MasterAggregator - Complete Real-World Scenario Tests", function () {
    let owner: Signer;
    let masterSigner1: Signer;
    let masterSigner2: Signer;
    let user1: Signer;
    let user2: Signer;
    let user3: Signer;
    let beneficiary: Signer;

    let entryPoint: EntryPoint;
    let masterAggregator: MasterAggregator;
    let walletManager: WalletManager;
    let gasPaymaster: GasPaymaster;
    let cpopToken: CPOPToken;

    let wallet1: AAWallet;
    let wallet2: AAWallet;
    let wallet3: AAWallet;

    let ownerAddress: string;
    let masterSigner1Address: string;
    let masterSigner2Address: string;
    let user1Address: string;
    let user2Address: string;
    let user3Address: string;

    const MASTER_PRIVATE_KEY_1 = ethers.id("test_master_key_1");
    const MASTER_PRIVATE_KEY_2 = ethers.id("test_master_key_2");
    const SESSION_KEY_PRIVATE = ethers.id("test_session_key");

    before(async function () {
        // Get signers
        [owner, masterSigner1, masterSigner2, user1, user2, user3, beneficiary] = await ethers.getSigners();
        
        ownerAddress = await owner.getAddress();
        masterSigner1Address = await masterSigner1.getAddress();
        masterSigner2Address = await masterSigner2.getAddress();
        user1Address = await user1.getAddress();
        user2Address = await user2.getAddress();
        user3Address = await user3.getAddress();

        console.log("🚀 Setting up complete MasterAggregator test environment...");
    });

    beforeEach(async function () {
        // Deploy EntryPoint
        const EntryPointFactory = (await ethers.getContractFactory("EntryPoint")) as EntryPoint__factory;
        entryPoint = await EntryPointFactory.deploy();
        await entryPoint.waitForDeployment();

        // Deploy CPOP Token
        const CPOPTokenFactory = (await ethers.getContractFactory("CPOPToken")) as CPOPToken__factory;
        cpopToken = await CPOPTokenFactory.deploy(ownerAddress);
        await cpopToken.waitForDeployment();

        // Deploy GasPaymaster
        const GasPaymasterFactory = (await ethers.getContractFactory("GasPaymaster")) as GasPaymaster__factory;
        gasPaymaster = await GasPaymasterFactory.deploy(
            await entryPoint.getAddress(),
            await cpopToken.getAddress(),
            ethers.ZeroAddress, // No oracle for simplicity
            ethers.parseEther("1000"), // 1000 tokens per ETH
            true, // Burn tokens
            ethers.ZeroAddress // No beneficiary when burning
        );
        await gasPaymaster.waitForDeployment();

        // Deploy MasterAggregator
        const MasterAggregatorFactory = (await ethers.getContractFactory("MasterAggregator")) as MasterAggregator__factory;
        masterAggregator = await MasterAggregatorFactory.deploy();
        await masterAggregator.waitForDeployment();

        // Initialize MasterAggregator
        await masterAggregator.initialize(
            ownerAddress,
            [masterSigner1Address, masterSigner2Address]
        );

        // Deploy WalletManager
        const WalletManagerFactory = (await ethers.getContractFactory("WalletManager")) as WalletManager__factory;
        walletManager = await WalletManagerFactory.deploy();
        await walletManager.waitForDeployment();

        // Initialize WalletManager
        await walletManager.initialize(
            await entryPoint.getAddress(),
            await cpopToken.getAddress(),
            ownerAddress
        );

        // Setup tokens for users
        await cpopToken.mint(user1Address, ethers.parseEther("1000"));
        await cpopToken.mint(user2Address, ethers.parseEther("1000"));
        await cpopToken.mint(user3Address, ethers.parseEther("1000"));

        console.log("✅ Test environment setup complete");
    });

    describe("1. Master Signer Management", function () {
        it("Should initialize with authorized master signers", async function () {
            expect(await masterAggregator.authorizedMasters(masterSigner1Address)).to.be.true;
            expect(await masterAggregator.authorizedMasters(masterSigner2Address)).to.be.true;
        });

        it("Should allow owner to authorize/deauthorize master signers", async function () {
            const newMaster = await user1.getAddress();
            
            // Authorize new master
            await expect(masterAggregator.setMasterAuthorization(newMaster, true))
                .to.emit(masterAggregator, "MasterAuthorized")
                .withArgs(newMaster, true);
                
            expect(await masterAggregator.authorizedMasters(newMaster)).to.be.true;

            // Deauthorize master
            await expect(masterAggregator.setMasterAuthorization(newMaster, false))
                .to.emit(masterAggregator, "MasterAuthorized")
                .withArgs(newMaster, false);
                
            expect(await masterAggregator.authorizedMasters(newMaster)).to.be.false;
        });

        it("Should not allow non-owner to manage master signers", async function () {
            await expect(
                masterAggregator.connect(user1).setMasterAuthorization(user1Address, true)
            ).to.be.revertedWith("Ownable: caller is not the owner");
        });
    });

    describe("2. Wallet Creation & Authorization", function () {
        beforeEach(async function () {
            // Create wallets for each user with master signers
            const salt1 = ethers.id("user1_wallet");
            const salt2 = ethers.id("user2_wallet");
            const salt3 = ethers.id("user3_wallet");

            // Create wallets with master signers
            const tx1 = await walletManager.createAccountWithMasterSigner(
                user1Address,
                salt1,
                masterSigner1Address
            );
            const receipt1 = await tx1.wait();
            const wallet1Address = receipt1?.logs.find(log => 
                log.topics[0] === walletManager.interface.getEvent("AccountCreated")!.topicHash
            )?.topics[1];
            
            const tx2 = await walletManager.createAccountWithMasterSigner(
                user2Address,
                salt2,
                masterSigner1Address
            );
            const receipt2 = await tx2.wait();
            const wallet2Address = receipt2?.logs.find(log => 
                log.topics[0] === walletManager.interface.getEvent("AccountCreated")!.topicHash
            )?.topics[1];

            const tx3 = await walletManager.createAccountWithMasterSigner(
                user3Address,
                salt3,
                masterSigner2Address
            );
            const receipt3 = await tx3.wait();
            const wallet3Address = receipt3?.logs.find(log => 
                log.topics[0] === walletManager.interface.getEvent("AccountCreated")!.topicHash
            )?.topics[1];

            // Get wallet instances
            wallet1 = AAWallet__factory.connect(wallet1Address!, owner);
            wallet2 = AAWallet__factory.connect(wallet2Address!, owner);
            wallet3 = AAWallet__factory.connect(wallet3Address!, owner);
        });

        it("Should create wallets with correct master signers", async function () {
            expect(await wallet1.getMasterSigner()).to.equal(masterSigner1Address);
            expect(await wallet2.getMasterSigner()).to.equal(masterSigner1Address);
            expect(await wallet3.getMasterSigner()).to.equal(masterSigner2Address);
        });

        it("Should auto-authorize wallets based on master signer relationship", async function () {
            const wallet1Address = await wallet1.getAddress();
            const wallet2Address = await wallet2.getAddress();

            // Auto-authorize wallets
            await masterAggregator.autoAuthorizeWallet(wallet1Address, masterSigner1Address);
            await masterAggregator.autoAuthorizeWallet(wallet2Address, masterSigner1Address);

            // Check authorization
            expect(await masterAggregator.isWalletControlledByMaster(wallet1Address, masterSigner1Address)).to.be.true;
            expect(await masterAggregator.isWalletControlledByMaster(wallet2Address, masterSigner1Address)).to.be.true;
        });

        it("Should batch authorize multiple wallets", async function () {
            const walletAddresses = [
                await wallet1.getAddress(),
                await wallet2.getAddress()
            ];

            await expect(
                masterAggregator.batchSetWalletAuthorization(
                    masterSigner1Address,
                    walletAddresses,
                    true
                )
            ).to.emit(masterAggregator, "WalletAuthorized");

            for (const walletAddr of walletAddresses) {
                expect(await masterAggregator.isWalletControlledByMaster(walletAddr, masterSigner1Address)).to.be.true;
            }
        });
    });

    describe("3. Signature Aggregation - Real World Scenarios", function () {
        let wallet1Address: string;
        let wallet2Address: string;

        beforeEach(async function () {
            // Setup wallets
            const salt1 = ethers.id("aggregation_user1");
            const salt2 = ethers.id("aggregation_user2");

            await walletManager.createAccountWithMasterSigner(user1Address, salt1, masterSigner1Address);
            await walletManager.createAccountWithMasterSigner(user2Address, salt2, masterSigner1Address);

            wallet1Address = await walletManager.getAccountAddress(user1Address, salt1);
            wallet2Address = await walletManager.getAccountAddress(user2Address, salt2);

            wallet1 = AAWallet__factory.connect(wallet1Address, owner);
            wallet2 = AAWallet__factory.connect(wallet2Address, owner);

            // Auto-authorize wallets
            await masterAggregator.autoAuthorizeWallet(wallet1Address, masterSigner1Address);
            await masterAggregator.autoAuthorizeWallet(wallet2Address, masterSigner1Address);

            // Fund wallets with ETH for gas
            await owner.sendTransaction({
                to: wallet1Address,
                value: ethers.parseEther("1")
            });
            await owner.sendTransaction({
                to: wallet2Address,
                value: ethers.parseEther("1")
            });
        });

        it("Should create and validate master aggregated signatures", async function () {
            // Create mock user operations
            const userOps: PackedUserOperation[] = [
                {
                    sender: wallet1Address,
                    nonce: 0,
                    initCode: "0x",
                    callData: wallet1.interface.encodeFunctionData("execute", [
                        user1Address,
                        ethers.parseEther("0.1"),
                        "0x"
                    ]),
                    accountGasLimits: ethers.concat([
                        ethers.toBeHex(21000, 16), // verificationGasLimit
                        ethers.toBeHex(100000, 16) // callGasLimit
                    ]),
                    preVerificationGas: 21000,
                    gasFees: ethers.concat([
                        ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16), // maxPriorityFeePerGas
                        ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)  // maxFeePerGas
                    ]),
                    paymasterAndData: "0x",
                    signature: "0x"
                },
                {
                    sender: wallet2Address,
                    nonce: 0,
                    initCode: "0x",
                    callData: wallet2.interface.encodeFunctionData("execute", [
                        user2Address,
                        ethers.parseEther("0.1"),
                        "0x"
                    ]),
                    accountGasLimits: ethers.concat([
                        ethers.toBeHex(21000, 16),
                        ethers.toBeHex(100000, 16)
                    ]),
                    preVerificationGas: 21000,
                    gasFees: ethers.concat([
                        ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16),
                        ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)
                    ]),
                    paymasterAndData: "0x",
                    signature: "0x"
                }
            ];

            // Create master aggregated signature
            const aggregatedSig = await masterAggregator.createMasterAggregatedSignature(
                userOps,
                masterSigner1Address,
                MASTER_PRIVATE_KEY_1
            );

            console.log("📝 Created aggregated signature for 2 operations");

            // Validate the aggregated signature
            await expect(
                masterAggregator.validateSignatures(userOps, aggregatedSig)
            ).to.emit(masterAggregator, "AggregatedValidation");

            console.log("✅ Aggregated signature validation successful");
        });

        it("Should calculate gas savings correctly", async function () {
            const savings2Ops = await masterAggregator.calculateGasSavings(2);
            const savings5Ops = await masterAggregator.calculateGasSavings(5);
            const savings10Ops = await masterAggregator.calculateGasSavings(10);

            console.log(`💰 Gas savings for 2 ops: ${savings2Ops}`);
            console.log(`💰 Gas savings for 5 ops: ${savings5Ops}`);
            console.log(`💰 Gas savings for 10 ops: ${savings10Ops}`);

            expect(savings2Ops).to.be.greaterThan(0);
            expect(savings5Ops).to.be.greaterThan(savings2Ops);
            expect(savings10Ops).to.be.greaterThan(savings5Ops);
        });

        it("Should handle maximum operations per aggregation", async function () {
            const maxOps = await masterAggregator.maxAggregatedOps();
            console.log(`📊 Maximum operations per aggregation: ${maxOps}`);

            // Create maximum number of operations
            const userOps: PackedUserOperation[] = [];
            for (let i = 0; i < maxOps; i++) {
                userOps.push({
                    sender: wallet1Address,
                    nonce: i,
                    initCode: "0x",
                    callData: "0x",
                    accountGasLimits: ethers.concat([
                        ethers.toBeHex(21000, 16),
                        ethers.toBeHex(100000, 16)
                    ]),
                    preVerificationGas: 21000,
                    gasFees: ethers.concat([
                        ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16),
                        ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)
                    ]),
                    paymasterAndData: "0x",
                    signature: "0x"
                });
            }

            const aggregatedSig = await masterAggregator.createMasterAggregatedSignature(
                userOps,
                masterSigner1Address,
                MASTER_PRIVATE_KEY_1
            );

            await expect(
                masterAggregator.validateSignatures(userOps, aggregatedSig)
            ).to.emit(masterAggregator, "AggregatedValidation");

            console.log(`✅ Successfully aggregated ${maxOps} operations`);
        });
    });

    describe("4. Session Key Aggregation", function () {
        let wallet1Address: string;
        let wallet2Address: string;
        let sessionKeyAddress: string;

        beforeEach(async function () {
            // Create session key address from private key
            const sessionKeyWallet = new ethers.Wallet(SESSION_KEY_PRIVATE);
            sessionKeyAddress = sessionKeyWallet.address;

            // Setup wallets
            const salt1 = ethers.id("session_user1");
            const salt2 = ethers.id("session_user2");

            await walletManager.createAccountWithMasterSigner(user1Address, salt1, masterSigner1Address);
            await walletManager.createAccountWithMasterSigner(user2Address, salt2, masterSigner1Address);

            wallet1Address = await walletManager.getAccountAddress(user1Address, salt1);
            wallet2Address = await walletManager.getAccountAddress(user2Address, salt2);

            wallet1 = AAWallet__factory.connect(wallet1Address, masterSigner1);
            wallet2 = AAWallet__factory.connect(wallet2Address, masterSigner1);

            // Add session keys to both wallets
            const validAfter = Math.floor(Date.now() / 1000) - 3600; // 1 hour ago
            const validUntil = Math.floor(Date.now() / 1000) + 86400; // 24 hours from now
            const permissions = ethers.id("EXECUTE_PERMISSION");

            await wallet1.addSessionKey(sessionKeyAddress, validAfter, validUntil, permissions);
            await wallet2.addSessionKey(sessionKeyAddress, validAfter, validUntil, permissions);

            console.log(`🔑 Added session key ${sessionKeyAddress} to wallets`);
        });

        it("Should create and validate session key aggregated signatures", async function () {
            const userOps: PackedUserOperation[] = [
                {
                    sender: wallet1Address,
                    nonce: 0,
                    initCode: "0x",
                    callData: "0x",
                    accountGasLimits: ethers.concat([
                        ethers.toBeHex(21000, 16),
                        ethers.toBeHex(100000, 16)
                    ]),
                    preVerificationGas: 21000,
                    gasFees: ethers.concat([
                        ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16),
                        ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)
                    ]),
                    paymasterAndData: "0x",
                    signature: "0x"
                },
                {
                    sender: wallet2Address,
                    nonce: 0,
                    initCode: "0x",
                    callData: "0x",
                    accountGasLimits: ethers.concat([
                        ethers.toBeHex(21000, 16),
                        ethers.toBeHex(100000, 16)
                    ]),
                    preVerificationGas: 21000,
                    gasFees: ethers.concat([
                        ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16),
                        ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)
                    ]),
                    paymasterAndData: "0x",
                    signature: "0x"
                }
            ];

            // Create session key aggregated signature
            const sessionSig = await masterAggregator.createSessionKeyAggregatedSignature(
                userOps,
                sessionKeyAddress,
                SESSION_KEY_PRIVATE
            );

            console.log("🔐 Created session key aggregated signature");

            // Validate session key signature
            const isValid = await masterAggregator.validateSessionKeyAggregatedSignature(
                userOps,
                sessionSig
            );

            expect(isValid).to.be.true;
            console.log("✅ Session key aggregated signature validation successful");
        });
    });

    describe("5. Integration with EntryPoint", function () {
        it("Should stake and manage aggregator in EntryPoint", async function () {
            const stakeAmount = ethers.parseEther("1");
            const unstakeDelay = 86400; // 1 day

            // Add stake to EntryPoint
            await masterAggregator.addStake(entryPoint, unstakeDelay, { value: stakeAmount });
            
            console.log(`💰 Added stake of ${ethers.formatEther(stakeAmount)} ETH to EntryPoint`);

            // Check stake info
            const stakeInfo = await entryPoint.deposits(await masterAggregator.getAddress());
            expect(stakeInfo.stake).to.equal(stakeAmount);
            expect(stakeInfo.unstakeDelaySec).to.equal(unstakeDelay);

            console.log("✅ Stake management successful");
        });
    });

    describe("6. Configuration Management", function () {
        it("Should update aggregation configuration", async function () {
            const newMaxOps = 75;
            const newValidationWindow = 600; // 10 minutes

            await masterAggregator.updateConfig(newMaxOps, newValidationWindow);

            expect(await masterAggregator.maxAggregatedOps()).to.equal(newMaxOps);
            expect(await masterAggregator.validationWindow()).to.equal(newValidationWindow);

            console.log(`⚙️ Updated configuration: maxOps=${newMaxOps}, window=${newValidationWindow}s`);
        });

        it("Should maintain nonce progression for master signers", async function () {
            const initialNonce = await masterAggregator.getMasterNonce(masterSigner1Address);
            expect(initialNonce).to.equal(0);

            // Create and validate a signature (this increments nonce)
            const userOps: PackedUserOperation[] = [{
                sender: user1Address,
                nonce: 0,
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.concat([
                    ethers.toBeHex(21000, 16),
                    ethers.toBeHex(100000, 16)
                ]),
                preVerificationGas: 21000,
                gasFees: ethers.concat([
                    ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16),
                    ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)
                ]),
                paymasterAndData: "0x",
                signature: "0x"
            }];

            // Authorize wallet first
            await masterAggregator.setWalletAuthorization(masterSigner1Address, user1Address, true);

            const aggregatedSig = await masterAggregator.createMasterAggregatedSignature(
                userOps,
                masterSigner1Address,
                MASTER_PRIVATE_KEY_1
            );

            await masterAggregator.validateSignatures(userOps, aggregatedSig);

            const newNonce = await masterAggregator.getMasterNonce(masterSigner1Address);
            expect(newNonce).to.equal(1);

            console.log(`📊 Nonce progression: ${initialNonce} → ${newNonce}`);
        });
    });

    describe("7. Error Handling & Security", function () {
        it("Should reject unauthorized master signers", async function () {
            const unauthorizedMaster = await user3.getAddress();
            const userOps: PackedUserOperation[] = [{
                sender: user1Address,
                nonce: 0,
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.concat([
                    ethers.toBeHex(21000, 16),
                    ethers.toBeHex(100000, 16)
                ]),
                preVerificationGas: 21000,
                gasFees: ethers.concat([
                    ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16),
                    ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)
                ]),
                paymasterAndData: "0x",
                signature: "0x"
            }];

            const invalidSig = abi.encode(
                ["address", "uint256", "bytes"],
                [unauthorizedMaster, 0, "0x"]
            );

            await expect(
                masterAggregator.validateSignatures(userOps, invalidSig)
            ).to.be.revertedWith("MasterAggregator: unauthorized master");
        });

        it("Should reject operations exceeding maximum limit", async function () {
            const tooManyOps: PackedUserOperation[] = [];
            const maxOps = await masterAggregator.maxAggregatedOps();
            
            for (let i = 0; i <= maxOps; i++) {
                tooManyOps.push({
                    sender: user1Address,
                    nonce: i,
                    initCode: "0x",
                    callData: "0x",
                    accountGasLimits: ethers.concat([
                        ethers.toBeHex(21000, 16),
                        ethers.toBeHex(100000, 16)
                    ]),
                    preVerificationGas: 21000,
                    gasFees: ethers.concat([
                        ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16),
                        ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)
                    ]),
                    paymasterAndData: "0x",
                    signature: "0x"
                });
            }

            const aggregatedSig = await masterAggregator.createMasterAggregatedSignature(
                tooManyOps,
                masterSigner1Address,
                MASTER_PRIVATE_KEY_1
            );

            await expect(
                masterAggregator.validateSignatures(tooManyOps, aggregatedSig)
            ).to.be.revertedWith("MasterAggregator: too many operations");
        });

        it("Should reject wallets not controlled by master", async function () {
            const userOps: PackedUserOperation[] = [{
                sender: user1Address, // Not authorized for masterSigner1
                nonce: 0,
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.concat([
                    ethers.toBeHex(21000, 16),
                    ethers.toBeHex(100000, 16)
                ]),
                preVerificationGas: 21000,
                gasFees: ethers.concat([
                    ethers.toBeHex(ethers.parseUnits("20", "gwei"), 16),
                    ethers.toBeHex(ethers.parseUnits("30", "gwei"), 16)
                ]),
                paymasterAndData: "0x",
                signature: "0x"
            }];

            const aggregatedSig = await masterAggregator.createMasterAggregatedSignature(
                userOps,
                masterSigner1Address,
                MASTER_PRIVATE_KEY_1
            );

            await expect(
                masterAggregator.validateSignatures(userOps, aggregatedSig)
            ).to.be.revertedWith("MasterAggregator: wallet not controlled by master");
        });
    });

    after(async function () {
        console.log("🎯 MasterAggregator comprehensive testing completed!");
        console.log("📊 Test Summary:");
        console.log("   ✅ Master signer management");
        console.log("   ✅ Wallet authorization");
        console.log("   ✅ Signature aggregation");
        console.log("   ✅ Session key support");
        console.log("   ✅ EntryPoint integration");
        console.log("   ✅ Configuration management");
        console.log("   ✅ Security validations");
    });
});

// Helper function for ABI encoding
function abi(...args: any[]) {
    return {
        encode: (types: string[], values: any[]) => {
            return ethers.AbiCoder.defaultAbiCoder().encode(types, values);
        }
    };
}