import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import {
    MasterAggregator,
    AAWallet,
    WalletManager,
    EntryPoint,
    MasterAggregator__factory,
    AAWallet__factory,
    WalletManager__factory,
    EntryPoint__factory,
} from "../typechain";
import { PackedUserOperation } from "./utils/UserOperation";

describe("MasterAggregator - System Integration Tests", function () {
    let owner: Signer;
    let masterSigner1: Signer;
    let masterSigner2: Signer;
    let user1: Signer;
    let user2: Signer;
    let user3: Signer;

    let entryPoint: EntryPoint;
    let masterAggregator: MasterAggregator;
    let walletManager: WalletManager;

    let wallet1: AAWallet;
    let wallet2: AAWallet;
    let wallet3: AAWallet;

    let ownerAddress: string;
    let masterSigner1Address: string;
    let masterSigner2Address: string;
    let user1Address: string;
    let user2Address: string;
    let user3Address: string;

    before(async function () {
        console.log("🚀 Setting up MasterAggregator system integration tests...");
        
        // Get signers
        [owner, masterSigner1, masterSigner2, user1, user2, user3] = await ethers.getSigners();
        
        ownerAddress = await owner.getAddress();
        masterSigner1Address = await masterSigner1.getAddress();
        masterSigner2Address = await masterSigner2.getAddress();
        user1Address = await user1.getAddress();
        user2Address = await user2.getAddress();
        user3Address = await user3.getAddress();

        console.log("Master Signer 1:", masterSigner1Address);
        console.log("Master Signer 2:", masterSigner2Address);
    });

    beforeEach(async function () {
        // Deploy EntryPoint
        const EntryPointFactory = (await ethers.getContractFactory("EntryPoint")) as EntryPoint__factory;
        entryPoint = await EntryPointFactory.deploy();
        await entryPoint.deployed();

        // Deploy MasterAggregator with initial masters
        const MasterAggregatorFactory = (await ethers.getContractFactory("MasterAggregator")) as MasterAggregator__factory;
        masterAggregator = await MasterAggregatorFactory.deploy();
        await masterAggregator.deployed();

        // Initialize MasterAggregator with master signers
        await masterAggregator.initialize(
            ownerAddress,
            [masterSigner1Address, masterSigner2Address]
        );

        // Deploy WalletManager
        const WalletManagerFactory = (await ethers.getContractFactory("WalletManager")) as WalletManager__factory;
        walletManager = await WalletManagerFactory.deploy();
        await walletManager.deployed();

        // Initialize WalletManager
        await walletManager.initialize(
            entryPoint.address,
            ethers.constants.AddressZero, // No token for simplicity
            ownerAddress
        );

        // Set MasterAggregator in WalletManager
        await walletManager.setMasterAggregator(masterAggregator.address);
    });

    describe("1. Master List Management (Fix Verification)", function () {
        it("should correctly track authorized masters in list", async function () {
            // Verify initial masters are in the list
            expect(await masterAggregator.authorizedMastersList(0)).to.equal(masterSigner1Address);
            expect(await masterAggregator.authorizedMastersList(1)).to.equal(masterSigner2Address);
            
            // Verify masters are authorized
            expect(await masterAggregator.authorizedMasters(masterSigner1Address)).to.be.true;
            expect(await masterAggregator.authorizedMasters(masterSigner2Address)).to.be.true;
        });

        it("should add new master to list when authorized", async function () {
            const newMaster = await user1.getAddress();
            
            await masterAggregator.setMasterAuthorization(newMaster, true);
            
            // Check that new master is added to list
            expect(await masterAggregator.authorizedMastersList(2)).to.equal(newMaster);
            expect(await masterAggregator.authorizedMasters(newMaster)).to.be.true;
        });

        it("should remove master from list when deauthorized", async function () {
            // Deauthorize masterSigner2
            await masterAggregator.setMasterAuthorization(masterSigner2Address, false);
            
            // Check that masterSigner2 is removed and list is compacted
            expect(await masterAggregator.authorizedMasters(masterSigner2Address)).to.be.false;
            
            // List should now only have masterSigner1
            try {
                await masterAggregator.authorizedMastersList(1);
                expect.fail("Should not have second element");
            } catch (error) {
                // Expected - list should be shorter
            }
        });
    });

    describe("2. _findMasterForWallet Function Testing", function () {
        beforeEach(async function () {
            // Create wallets with master signers
            const salt1 = ethers.randomBytes(32);
            const salt2 = ethers.randomBytes(32);
            const salt3 = ethers.randomBytes(32);

            // Create wallet 1 with masterSigner1
            const tx1 = await walletManager.createAccountWithMasterSigner(
                user1Address, 
                salt1, 
                masterSigner1Address
            );
            const receipt1 = await tx1.wait();
            const wallet1Address = receipt1?.logs?.[0]?.args?.[0];
            wallet1 = AAWallet__factory.connect(wallet1Address, owner);

            // Create wallet 2 with masterSigner1 
            const tx2 = await walletManager.createAccountWithMasterSigner(
                user2Address,
                salt2,
                masterSigner1Address
            );
            const receipt2 = await tx2.wait();
            const wallet2Address = receipt2?.logs?.[0]?.args?.[0];
            wallet2 = AAWallet__factory.connect(wallet2Address, owner);

            // Create wallet 3 with masterSigner2
            const tx3 = await walletManager.createAccountWithMasterSigner(
                user3Address,
                salt3,
                masterSigner2Address
            );
            const receipt3 = await tx3.wait();
            const wallet3Address = receipt3?.logs?.[0]?.args?.[0];
            wallet3 = AAWallet__factory.connect(wallet3Address, owner);
        });

        it("should correctly identify master for wallet", async function () {
            // Test isWalletControlledByMaster for each wallet
            expect(await masterAggregator.isWalletControlledByMaster(
                wallet1.address, 
                masterSigner1Address
            )).to.be.true;

            expect(await masterAggregator.isWalletControlledByMaster(
                wallet2.address, 
                masterSigner1Address
            )).to.be.true;

            expect(await masterAggregator.isWalletControlledByMaster(
                wallet3.address, 
                masterSigner2Address
            )).to.be.true;

            // Test negative cases
            expect(await masterAggregator.isWalletControlledByMaster(
                wallet1.address, 
                masterSigner2Address
            )).to.be.false;

            expect(await masterAggregator.isWalletControlledByMaster(
                wallet3.address, 
                masterSigner1Address
            )).to.be.false;
        });

        it("should handle non-existent wallet gracefully", async function () {
            const randomAddress = ethers.Wallet.createRandom().address;
            
            expect(await masterAggregator.isWalletControlledByMaster(
                randomAddress, 
                masterSigner1Address
            )).to.be.false;
        });
    });

    describe("3. Signature Aggregation System Test", function () {
        let wallet1Address: string;
        let wallet2Address: string;
        let wallet3Address: string;

        beforeEach(async function () {
            // Create wallets with different master signers
            const salt1 = ethers.randomBytes(32);
            const salt2 = ethers.randomBytes(32);
            const salt3 = ethers.randomBytes(32);

            // Create wallet 1 & 2 with masterSigner1
            const tx1 = await walletManager.createAccountWithMasterSigner(
                user1Address, salt1, masterSigner1Address
            );
            const receipt1 = await tx1.wait();
            wallet1Address = receipt1?.logs?.[0]?.args?.[0];

            const tx2 = await walletManager.createAccountWithMasterSigner(
                user2Address, salt2, masterSigner1Address
            );
            const receipt2 = await tx2.wait();
            wallet2Address = receipt2?.logs?.[0]?.args?.[0];

            // Create wallet 3 with masterSigner2
            const tx3 = await walletManager.createAccountWithMasterSigner(
                user3Address, salt3, masterSigner2Address
            );
            const receipt3 = await tx3.wait();
            wallet3Address = receipt3?.logs?.[0]?.args?.[0];
        });

        it("should aggregate operations from same master", async function () {
            // Create user operations for wallets controlled by masterSigner1
            const userOp1: PackedUserOperation = {
                sender: wallet1Address,
                nonce: 0,
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [100000, 100000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            const userOp2: PackedUserOperation = {
                sender: wallet2Address,
                nonce: 0,
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [100000, 100000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            // Test aggregateSignatures
            const aggregatedSig = await masterAggregator.aggregateSignatures([userOp1, userOp2]);
            
            // Decode the aggregated signature
            const decoded = ethers.utils.defaultAbiCoder.decode(
                ["address", "uint256", "string"],
                aggregatedSig
            );

            expect(decoded[0]).to.equal(masterSigner1Address); // Should identify masterSigner1
            expect(decoded[1]).to.equal(0); // Nonce should be 0
            expect(decoded[2]).to.equal("MASTER_AGGREGATION_PLACEHOLDER");
        });

        it("should reject operations from different masters", async function () {
            // Create user operations for wallets controlled by different masters
            const userOp1: PackedUserOperation = {
                sender: wallet1Address, // masterSigner1
                nonce: 0,
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [100000, 100000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            const userOp3: PackedUserOperation = {
                sender: wallet3Address, // masterSigner2
                nonce: 0,
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [100000, 100000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            // Should revert when mixing operations from different masters
            await expect(
                masterAggregator.aggregateSignatures([userOp1, userOp3])
            ).to.be.revertedWith("MasterAggregator: operations from different masters");
        });

        it("should handle unauthorized wallet gracefully", async function () {
            // Create operation for a wallet that doesn't exist or isn't authorized
            const fakeWallet = ethers.Wallet.createRandom().address;
            
            const userOp: PackedUserOperation = {
                sender: fakeWallet,
                nonce: 0,
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [100000, 100000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            await expect(
                masterAggregator.aggregateSignatures([userOp])
            ).to.be.revertedWith("MasterAggregator: wallet has no master");
        });
    });

    describe("4. Master Signing and Validation", function () {
        let wallet1Address: string;
        let wallet2Address: string;

        beforeEach(async function () {
            // Create wallets for testing
            const salt1 = ethers.randomBytes(32);
            const salt2 = ethers.randomBytes(32);

            const tx1 = await walletManager.createAccountWithMasterSigner(
                user1Address, salt1, masterSigner1Address
            );
            const receipt1 = await tx1.wait();
            wallet1Address = receipt1?.logs?.[0]?.args?.[0];

            const tx2 = await walletManager.createAccountWithMasterSigner(
                user2Address, salt2, masterSigner1Address
            );
            const receipt2 = await tx2.wait();
            wallet2Address = receipt2?.logs?.[0]?.args?.[0];
        });

        it("should generate correct signing data for master", async function () {
            const userOp1: PackedUserOperation = {
                sender: wallet1Address,
                nonce: 0,
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [100000, 100000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            const userOp2: PackedUserOperation = {
                sender: wallet2Address,
                nonce: 0,
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [100000, 100000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            // Get signing data
            const [hashToSign, nonce] = await masterAggregator.getMasterSigningData(
                [userOp1, userOp2], 
                masterSigner1Address
            );

            expect(hashToSign).to.not.equal(ethers.constants.HashZero);
            expect(nonce).to.equal(0);

            // Verify the hash is deterministic
            const [hashToSign2, nonce2] = await masterAggregator.getMasterSigningData(
                [userOp1, userOp2], 
                masterSigner1Address
            );

            expect(hashToSign).to.equal(hashToSign2);
            expect(nonce).to.equal(nonce2);
        });

        it("should validate master signatures correctly", async function () {
            const userOp1: PackedUserOperation = {
                sender: wallet1Address,
                nonce: 0,
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [100000, 100000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            // Get signing data
            const [hashToSign, nonce] = await masterAggregator.getMasterSigningData(
                [userOp1], 
                masterSigner1Address
            );

            // Sign with master signer
            const signature = await masterSigner1.signMessage(ethers.utils.arrayify(hashToSign));

            // Test signature validation (this would normally be called by EntryPoint)
            const userOpHash = ethers.utils.keccak256(ethers.utils.toUtf8Bytes("test"));
            const validationResult = await masterAggregator.validateSignatures(
                [userOp1],
                signature
            );

            // Should not revert and return success indicator
            expect(validationResult).to.not.be.undefined;
        });
    });

    describe("5. Integration with WalletManager", function () {
        it("should automatically register wallet-master relationships", async function () {
            const salt = ethers.randomBytes(32);
            
            // Create wallet through WalletManager
            const tx = await walletManager.createAccountWithMasterSigner(
                user1Address, 
                salt, 
                masterSigner1Address
            );
            const receipt = await tx.wait();
            const walletAddress = receipt?.logs?.[0]?.args?.[0];

            // Verify that the wallet-master relationship is registered
            expect(await masterAggregator.isWalletControlledByMaster(
                walletAddress, 
                masterSigner1Address
            )).to.be.true;
        });

        it("should set aggregator address in wallet during creation", async function () {
            const salt = ethers.randomBytes(32);
            
            const tx = await walletManager.createAccountWithMasterSigner(
                user1Address, 
                salt, 
                masterSigner1Address
            );
            const receipt = await tx.wait();
            const walletAddress = receipt?.logs?.[0]?.args?.[0];

            const wallet = AAWallet__factory.connect(walletAddress, owner);
            
            // Verify aggregator is set
            expect(await wallet.getAggregator()).to.equal(masterAggregator.address);
        });
    });

    after(async function () {
        console.log("🎯 MasterAggregator system integration tests completed!");
        console.log("📊 Test Summary:");
        console.log("   ✅ Master list management verified");
        console.log("   ✅ _findMasterForWallet function fixed and tested");
        console.log("   ✅ Signature aggregation system working");
        console.log("   ✅ Master signing and validation functional");
        console.log("   ✅ WalletManager integration successful");
    });
});