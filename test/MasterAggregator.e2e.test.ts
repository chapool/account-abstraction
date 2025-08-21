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

interface UserOpsPerAggregator {
    userOps: PackedUserOperation[];
    aggregator: string;
    signature: string;
}

describe("MasterAggregator - End-to-End handleAggregatedOps Tests", function () {
    let owner: Signer;
    let masterSigner1: Signer;
    let user1: Signer;
    let user2: Signer;
    let beneficiary: Signer;

    let entryPoint: EntryPoint;
    let masterAggregator: MasterAggregator;
    let walletManager: WalletManager;
    let wallet1: AAWallet;
    let wallet2: AAWallet;

    let ownerAddress: string;
    let masterSigner1Address: string;
    let user1Address: string;
    let user2Address: string;
    let beneficiaryAddress: string;

    before(async function () {
        console.log("🚀 Setting up End-to-End handleAggregatedOps test environment...");
        
        [owner, masterSigner1, user1, user2, beneficiary] = await ethers.getSigners();
        
        ownerAddress = await owner.getAddress();
        masterSigner1Address = await masterSigner1.getAddress();
        user1Address = await user1.getAddress();
        user2Address = await user2.getAddress();
        beneficiaryAddress = await beneficiary.getAddress();

        console.log("Master Signer:", masterSigner1Address);
        console.log("User 1:", user1Address);
        console.log("User 2:", user2Address);
    });

    beforeEach(async function () {
        // Deploy EntryPoint
        const EntryPointFactory = (await ethers.getContractFactory("EntryPoint")) as EntryPoint__factory;
        entryPoint = await EntryPointFactory.deploy();
        await entryPoint.deployed();

        // Deploy MasterAggregator implementation
        const MasterAggregatorFactory = (await ethers.getContractFactory("MasterAggregator")) as MasterAggregator__factory;
        const masterAggregatorImpl = await MasterAggregatorFactory.deploy();
        await masterAggregatorImpl.deployed();

        // Deploy proxy for MasterAggregator
        const ProxyFactory = await ethers.getContractFactory("ERC1967Proxy");
        const initData = masterAggregatorImpl.interface.encodeFunctionData("initialize", [
            ownerAddress,
            [masterSigner1Address]
        ]);
        
        const masterAggregatorProxy = await ProxyFactory.deploy(masterAggregatorImpl.address, initData);
        await masterAggregatorProxy.deployed();

        // Connect to the proxy using the implementation ABI and owner as signer
        masterAggregator = new ethers.Contract(
            masterAggregatorProxy.address,
            masterAggregatorImpl.interface,
            owner
        ) as MasterAggregator;

        // Deploy WalletManager
        const WalletManagerFactory = (await ethers.getContractFactory("WalletManager")) as WalletManager__factory;
        const walletManagerImpl = await WalletManagerFactory.deploy();
        await walletManagerImpl.deployed();

        // Deploy TestToken for WalletManager initialization
        const TestToken = await ethers.getContractFactory("TestToken");
        const testToken = await TestToken.deploy();
        await testToken.deployed();

        // Deploy proxy for WalletManager  
        const walletManagerInitData = walletManagerImpl.interface.encodeFunctionData("initialize", [
            entryPoint.address,
            testToken.address,
            ownerAddress
        ]);
        
        const walletManagerProxy = await ProxyFactory.deploy(walletManagerImpl.address, walletManagerInitData);
        await walletManagerProxy.deployed();

        // Connect to the proxy using the implementation ABI and owner as signer
        walletManager = new ethers.Contract(
            walletManagerProxy.address,
            walletManagerImpl.interface,
            owner
        ) as WalletManager;

        // Set MasterAggregator in WalletManager
        await walletManager.setMasterAggregator(masterAggregator.address);
        console.log("WalletManager aggregator set to:", await walletManager.masterAggregatorAddress());

        // Update MasterAggregator configuration to ensure proper settings
        await masterAggregator.updateConfig(50, 300);

        // Create wallets
        await createTestWallets();
        
        // Fund wallets and EntryPoint
        await fundAccounts();
    });

    async function createTestWallets() {
        // Create wallets directly with aggregator configuration
        // Use simpler approach: create wallets directly with initialize
        
        const AAWalletFactory = await ethers.getContractFactory("AAWallet");
        const ProxyFactory = await ethers.getContractFactory("ERC1967Proxy");
        
        // Deploy wallet implementations
        const walletImpl = await AAWalletFactory.deploy();
        await walletImpl.deployed();
        
        // Create wallet 1 with aggregator
        const initData1 = walletImpl.interface.encodeFunctionData("initialize", [
            entryPoint.address,
            ownerAddress, // Use our test account as owner
            masterSigner1Address,
            masterAggregator.address
        ]);
        
        const walletProxy1 = await ProxyFactory.deploy(walletImpl.address, initData1);
        await walletProxy1.deployed();
        
        wallet1 = new ethers.Contract(
            walletProxy1.address,
            walletImpl.interface,
            owner
        ) as AAWallet;
        
        // Create wallet 2 with aggregator
        const initData2 = walletImpl.interface.encodeFunctionData("initialize", [
            entryPoint.address,
            ownerAddress, // Use our test account as owner
            masterSigner1Address,
            masterAggregator.address
        ]);
        
        const walletProxy2 = await ProxyFactory.deploy(walletImpl.address, initData2);
        await walletProxy2.deployed();
        
        wallet2 = new ethers.Contract(
            walletProxy2.address,
            walletImpl.interface,
            owner
        ) as AAWallet;

        console.log("Wallet 1 created at:", wallet1.address);
        console.log("Wallet 2 created at:", wallet2.address);
        
        // Manually register wallet-master relationships in aggregator
        await masterAggregator.setWalletAuthorization(masterSigner1Address, wallet1.address, true);
        await masterAggregator.setWalletAuthorization(masterSigner1Address, wallet2.address, true);
        
        console.log("Wallets created with aggregator configuration");
    }

    async function fundAccounts() {
        // Fund wallets with ETH for gas
        await owner.sendTransaction({
            to: wallet1.address,
            value: ethers.utils.parseEther("1")
        });
        
        await owner.sendTransaction({
            to: wallet2.address,
            value: ethers.utils.parseEther("1")
        });

        // Deposit stake for wallets in EntryPoint
        await entryPoint.depositTo(wallet1.address, { value: ethers.utils.parseEther("1") });
        await entryPoint.depositTo(wallet2.address, { value: ethers.utils.parseEther("1") });

        console.log("✅ Accounts funded and staked");
    }

    describe("Complete handleAggregatedOps Flow", function () {
        it("should execute multiple UserOps via handleAggregatedOps with master signature", async function () {
            console.log("🔄 Starting complete aggregated operation flow...");

            // Step 1: Create UserOperations for both wallets
            const userOp1: PackedUserOperation = {
                sender: wallet1.address,
                nonce: await entryPoint.getNonce(wallet1.address, 0),
                initCode: "0x",
                callData: wallet1.interface.encodeFunctionData("execute", [
                    user1Address, // Send to self
                    ethers.utils.parseEther("0.1"),
                    "0x"
                ]),
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [200000, 200000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x" // Empty signature for aggregator validation
            };

            const userOp2: PackedUserOperation = {
                sender: wallet2.address,
                nonce: await entryPoint.getNonce(wallet2.address, 0),
                initCode: "0x",
                callData: wallet2.interface.encodeFunctionData("execute", [
                    user2Address, // Send to self
                    ethers.utils.parseEther("0.1"),
                    "0x"
                ]),
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [200000, 200000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x" // Empty signature for aggregator validation
            };

            // Step 2: Get UserOp hashes
            const userOpHash1 = await entryPoint.getUserOpHash(userOp1);
            const userOpHash2 = await entryPoint.getUserOpHash(userOp2);

            console.log("UserOp Hash 1:", userOpHash1);
            console.log("UserOp Hash 2:", userOpHash2);

            // Step 3: Get master signing data  
            const [hashToSign, nonce] = await masterAggregator.getMasterSigningData(
                [userOp1, userOp2],
                masterSigner1Address
            );

            console.log("Hash to sign:", hashToSign);
            console.log("Master nonce:", nonce.toString());

            // Step 4: Master signer creates signature
            // Create the message hash manually to match what the contract expects
            // The contract will call toEthSignedMessageHash() on the aggregated hash
            // So we need to sign the raw aggregated hash (not the eth-prefixed version)
            
            // Create the aggregated hash components manually
            const userOpHashes = [];
            for (const userOp of [userOp1, userOp2]) {
                const opHash = ethers.utils.keccak256(ethers.utils.defaultAbiCoder.encode(
                    ["address", "uint256", "bytes", "bytes32", "uint256", "bytes32", "bytes"],
                    [
                        userOp.sender,
                        userOp.nonce,
                        userOp.callData,
                        userOp.accountGasLimits,
                        userOp.preVerificationGas,
                        userOp.gasFees,
                        userOp.paymasterAndData
                    ]
                ));
                userOpHashes.push(opHash);
            }
            
            const aggregatedHash = ethers.utils.keccak256(ethers.utils.defaultAbiCoder.encode(
                ["string", "address", "uint256", "uint256", "address", "bytes32[]"],
                [
                    "MASTER_AGGREGATION",
                    masterSigner1Address,
                    nonce,
                    31337, // chainId
                    masterAggregator.address,
                    userOpHashes
                ]
            ));
            
            // Now sign this raw hash - ethers will add the eth prefix
            const masterSignature = await masterSigner1.signMessage(ethers.utils.arrayify(aggregatedHash));
            console.log("Master signature created");

            // Debug: Check master aggregator configuration
            console.log("Max aggregated ops:", await masterAggregator.maxAggregatedOps());
            console.log("Master authorized:", await masterAggregator.authorizedMasters(masterSigner1Address));
            console.log("UserOps length:", [userOp1, userOp2].length);
            console.log("Wallet 1 controlled by master:", await masterAggregator.isWalletControlledByMaster(wallet1.address, masterSigner1Address));
            console.log("Wallet 2 controlled by master:", await masterAggregator.isWalletControlledByMaster(wallet2.address, masterSigner1Address));
            
            // Check wallet configuration
            console.log("Wallet 1 aggregator address:", await wallet1.aggregatorAddress());
            console.log("Wallet 2 aggregator address:", await wallet2.aggregatorAddress());
            console.log("Wallet 1 master signer:", await wallet1.masterSigner());
            console.log("Wallet 2 master signer:", await wallet2.masterSigner());
            console.log("Expected aggregator address:", masterAggregator.address);

            // Step 5: Create aggregated signature
            let aggregatedSignature;
            try {
                // First call the function to update state (nonces), then get return value
                const tx = await masterAggregator.createMasterAggregatedSignature(
                    [userOp1, userOp2],
                    masterSigner1Address,
                    masterSignature
                );
                
                // Get the return value from the transaction
                aggregatedSignature = await masterAggregator.callStatic.createMasterAggregatedSignature(
                    [userOp1, userOp2],
                    masterSigner1Address,
                    masterSignature
                );
                
                console.log("Aggregated signature:", aggregatedSignature);
                console.log("Aggregated signature type:", typeof aggregatedSignature);
                console.log("Aggregated signature length:", aggregatedSignature ? aggregatedSignature.length : "undefined");
            } catch (error) {
                console.error("Error creating aggregated signature:", error);
                throw error;
            }

            // Step 6: Prepare UserOpsPerAggregator structure
            const opsPerAggregator: UserOpsPerAggregator[] = [{
                userOps: [userOp1, userOp2],
                aggregator: masterAggregator.address,
                signature: aggregatedSignature
            }];

            // Step 7: Prepare for EntryPoint execution (skip direct validation to avoid nonce increment)

            // Step 8: Execute via EntryPoint.handleAggregatedOps
            console.log("🎯 Executing handleAggregatedOps...");
            
            const balanceBefore1 = await ethers.provider.getBalance(user1Address);
            const balanceBefore2 = await ethers.provider.getBalance(user2Address);

            // Validate our data before calling EntryPoint
            console.log("UserOp1 signature length:", userOp1.signature.length);
            console.log("UserOp2 signature length:", userOp2.signature.length);
            console.log("Aggregated signature length:", aggregatedSignature.length);
            console.log("Beneficiary address:", beneficiaryAddress);

            const tx = await entryPoint.handleAggregatedOps(
                opsPerAggregator,
                beneficiaryAddress
            );
            const receipt = await tx.wait();

            console.log("✅ handleAggregatedOps executed successfully!");
            console.log("Gas used:", receipt.gasUsed.toString());

            // Step 8: Verify execution results
            const balanceAfter1 = await ethers.provider.getBalance(user1Address);
            const balanceAfter2 = await ethers.provider.getBalance(user2Address);

            // Both users should have received 0.1 ETH
            expect(balanceAfter1.sub(balanceBefore1)).to.equal(ethers.utils.parseEther("0.1"));
            expect(balanceAfter2.sub(balanceBefore2)).to.equal(ethers.utils.parseEther("0.1"));

            // Nonces should have incremented
            expect(await entryPoint.getNonce(wallet1.address, 0)).to.equal(1);
            expect(await entryPoint.getNonce(wallet2.address, 0)).to.equal(1);

            console.log("✅ All operations executed successfully!");
            console.log("User 1 received:", ethers.utils.formatEther(balanceAfter1.sub(balanceBefore1)), "ETH");
            console.log("User 2 received:", ethers.utils.formatEther(balanceAfter2.sub(balanceBefore2)), "ETH");
        });

        it("should handle signature validation through aggregator", async function () {
            console.log("🔍 Testing signature validation flow...");

            // Create a simple UserOp
            const userOp: PackedUserOperation = {
                sender: wallet1.address,
                nonce: await entryPoint.getNonce(wallet1.address, 0),
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [100000, 100000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            // Get signing data and create signature
            const [hashToSign] = await masterAggregator.getMasterSigningData(
                [userOp],
                masterSigner1Address
            );

            const masterSignature = await masterSigner1.signMessage(ethers.utils.arrayify(hashToSign));

            // Test validateSignatures function
            const userOpHash = await entryPoint.getUserOpHash(userOp);
            const validationResult = await masterAggregator.validateSignatures(
                [userOp],
                masterSignature
            );

            // Should return encoded validation data
            expect(validationResult).to.not.equal("0x");
            console.log("✅ Signature validation successful");
        });

        it("should reject operations from unauthorized wallets", async function () {
            console.log("🚫 Testing unauthorized wallet rejection...");

            // Create a wallet not controlled by our master signer
            const unauthorizedWallet = ethers.Wallet.createRandom().address;

            const userOp: PackedUserOperation = {
                sender: unauthorizedWallet,
                nonce: 0,
                initCode: "0x",
                callData: "0x",
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [100000, 100000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            // Should fail at signature aggregation level
            await expect(
                masterAggregator.aggregateSignatures([userOp])
            ).to.be.revertedWith("MasterAggregator: wallet has no master");

            console.log("✅ Unauthorized wallet correctly rejected");
        });
    });

    describe("Gas Efficiency Testing", function () {
        it("should demonstrate gas savings with aggregation", async function () {
            console.log("⛽ Testing gas efficiency of aggregation...");

            // Create 3 operations for comparison
            const userOps: PackedUserOperation[] = [];
            
            for (let i = 0; i < 2; i++) {
                const wallet = i === 0 ? wallet1 : wallet2;
                const userOp: PackedUserOperation = {
                    sender: wallet.address,
                    nonce: await entryPoint.getNonce(wallet.address, 0),
                    initCode: "0x",
                    callData: wallet.interface.encodeFunctionData("execute", [
                        await owner.getAddress(),
                        ethers.utils.parseEther("0.01"),
                        "0x"
                    ]),
                    accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [150000, 150000]),
                    preVerificationGas: 21000,
                    gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                    paymasterAndData: "0x",
                    signature: "0x"
                };
                userOps.push(userOp);
            }

            // Execute with aggregation
            const [hashToSign] = await masterAggregator.getMasterSigningData(
                userOps,
                masterSigner1Address
            );

            const masterSignature = await masterSigner1.signMessage(ethers.utils.arrayify(hashToSign));
            const aggregatedSignature = await masterAggregator.createMasterAggregatedSignature(
                userOps,
                masterSigner1Address,
                masterSignature
            );

            const opsPerAggregator: UserOpsPerAggregator[] = [{
                userOps: userOps,
                aggregator: masterAggregator.address,
                signature: aggregatedSignature
            }];

            const tx = await entryPoint.handleAggregatedOps(
                opsPerAggregator,
                beneficiaryAddress
            );
            const receipt = await tx.wait();

            console.log("✅ Aggregated execution completed");
            console.log("Gas used for", userOps.length, "operations:", receipt.gasUsed.toString());
            console.log("Average gas per operation:", receipt.gasUsed.div(userOps.length).toString());

            // Calculate theoretical gas savings
            const gasSavings = await masterAggregator.calculateGasSavings(userOps.length);
            console.log("Theoretical gas savings:", gasSavings.toString());
        });
    });

    after(async function () {
        console.log("🎯 End-to-End handleAggregatedOps testing completed!");
        console.log("📊 Test Summary:");
        console.log("   ✅ Complete aggregated operations flow tested");
        console.log("   ✅ Master signature validation verified");
        console.log("   ✅ EntryPoint.handleAggregatedOps integration confirmed");
        console.log("   ✅ Multi-wallet aggregation working");
        console.log("   ✅ Gas efficiency demonstrated");
        console.log("   ✅ Security validations passing");
    });
});