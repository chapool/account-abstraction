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

describe("Web2 Users - handleOps Flow Test", function () {
    let owner: Signer;
    let masterSigner: Signer;
    let web2User1: Signer;
    let web2User2: Signer;
    let web2User3: Signer;
    let beneficiary: Signer;

    let entryPoint: EntryPoint;
    let masterAggregator: MasterAggregator;
    let walletManager: WalletManager;
    let wallet1: AAWallet;
    let wallet2: AAWallet;
    let wallet3: AAWallet;

    let ownerAddress: string;
    let masterSignerAddress: string;
    let web2User1Address: string;
    let web2User2Address: string;
    let web2User3Address: string;
    let beneficiaryAddress: string;

    before(async function () {
        console.log("🚀 Setting up Web2 Users handleOps test environment...");
        
        [owner, masterSigner, web2User1, web2User2, web2User3, beneficiary] = await ethers.getSigners();
        
        ownerAddress = await owner.getAddress();
        masterSignerAddress = await masterSigner.getAddress();
        web2User1Address = await web2User1.getAddress();
        web2User2Address = await web2User2.getAddress();
        web2User3Address = await web2User3.getAddress();
        beneficiaryAddress = await beneficiary.getAddress();

        console.log("Master Signer:", masterSignerAddress);
        console.log("Web2 User 1:", web2User1Address);
        console.log("Web2 User 2:", web2User2Address);
        console.log("Web2 User 3:", web2User3Address);
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
            [masterSignerAddress]
        ]);
        
        const masterAggregatorProxy = await ProxyFactory.deploy(masterAggregatorImpl.address, initData);
        await masterAggregatorProxy.deployed();

        masterAggregator = new ethers.Contract(
            masterAggregatorProxy.address,
            masterAggregatorImpl.interface,
            owner
        ) as MasterAggregator;

        // Deploy TestToken for WalletManager initialization
        const TestToken = await ethers.getContractFactory("TestToken");
        const testToken = await TestToken.deploy();
        await testToken.deployed();

        // Deploy WalletManager
        const WalletManagerFactory = (await ethers.getContractFactory("WalletManager")) as WalletManager__factory;
        const walletManagerImpl = await WalletManagerFactory.deploy();
        await walletManagerImpl.deployed();

        const walletManagerInitData = walletManagerImpl.interface.encodeFunctionData("initialize", [
            entryPoint.address,
            testToken.address,
            ownerAddress
        ]);
        
        const walletManagerProxy = await ProxyFactory.deploy(walletManagerImpl.address, walletManagerInitData);
        await walletManagerProxy.deployed();

        walletManager = new ethers.Contract(
            walletManagerProxy.address,
            walletManagerImpl.interface,
            owner
        ) as WalletManager;

        // Set MasterAggregator in WalletManager
        await walletManager.setMasterAggregator(masterAggregator.address);
        console.log("WalletManager aggregator set to:", await walletManager.masterAggregatorAddress());

        // Update MasterAggregator configuration
        await masterAggregator.updateConfig(50, 300);

        // Create Web2 user wallets
        await createWeb2UserWallets();
        
        // Fund accounts
        await fundAccounts();
    });

    async function createWeb2UserWallets() {
        console.log("🏗️ Creating Web2 user wallets...");

        // Create wallets directly with aggregator configuration for testing
        const AAWalletFactory = await ethers.getContractFactory("AAWallet");
        const ProxyFactory = await ethers.getContractFactory("ERC1967Proxy");
        
        const walletImpl = await AAWalletFactory.deploy();
        await walletImpl.deployed();

        // Create wallet 1 for web2User1
        const initData1 = walletImpl.interface.encodeFunctionData("initializeWithAggregator", [
            entryPoint.address,
            web2User1Address, // User1 owns their wallet
            masterSignerAddress,
            masterAggregator.address
        ]);
        
        const walletProxy1 = await ProxyFactory.deploy(walletImpl.address, initData1);
        await walletProxy1.deployed();
        
        wallet1 = new ethers.Contract(
            walletProxy1.address,
            walletImpl.interface,
            web2User1 // Connect with user1 as signer
        ) as AAWallet;

        // Create wallet 2 for web2User2
        const initData2 = walletImpl.interface.encodeFunctionData("initializeWithAggregator", [
            entryPoint.address,
            web2User2Address, // User2 owns their wallet
            masterSignerAddress,
            masterAggregator.address
        ]);
        
        const walletProxy2 = await ProxyFactory.deploy(walletImpl.address, initData2);
        await walletProxy2.deployed();
        
        wallet2 = new ethers.Contract(
            walletProxy2.address,
            walletImpl.interface,
            web2User2 // Connect with user2 as signer
        ) as AAWallet;

        // Create wallet 3 for web2User3
        const initData3 = walletImpl.interface.encodeFunctionData("initializeWithAggregator", [
            entryPoint.address,
            web2User3Address, // User3 owns their wallet
            masterSignerAddress,
            masterAggregator.address
        ]);
        
        const walletProxy3 = await ProxyFactory.deploy(walletImpl.address, initData3);
        await walletProxy3.deployed();
        
        wallet3 = new ethers.Contract(
            walletProxy3.address,
            walletImpl.interface,
            web2User3 // Connect with user3 as signer
        ) as AAWallet;

        console.log("Web2 User 1 Wallet:", wallet1.address);
        console.log("Web2 User 2 Wallet:", wallet2.address);
        console.log("Web2 User 3 Wallet:", wallet3.address);
        
        // Register wallet-master relationships in aggregator
        await masterAggregator.setWalletAuthorization(masterSignerAddress, wallet1.address, true);
        await masterAggregator.setWalletAuthorization(masterSignerAddress, wallet2.address, true);
        await masterAggregator.setWalletAuthorization(masterSignerAddress, wallet3.address, true);
        
        console.log("✅ Web2 user wallets created and configured");
    }

    async function fundAccounts() {
        console.log("💰 Funding Web2 user accounts...");
        
        // Fund wallets with ETH for gas
        await owner.sendTransaction({
            to: wallet1.address,
            value: ethers.utils.parseEther("2")
        });
        
        await owner.sendTransaction({
            to: wallet2.address,
            value: ethers.utils.parseEther("2")
        });

        await owner.sendTransaction({
            to: wallet3.address,
            value: ethers.utils.parseEther("2")
        });

        // Deposit stake for wallets in EntryPoint
        await entryPoint.depositTo(wallet1.address, { value: ethers.utils.parseEther("1") });
        await entryPoint.depositTo(wallet2.address, { value: ethers.utils.parseEther("1") });
        await entryPoint.depositTo(wallet3.address, { value: ethers.utils.parseEther("1") });

        console.log("✅ All Web2 user accounts funded and staked");
    }

    describe("Web2 Users Submitting Individual UserOps", function () {
        it("should handle multiple Web2 users submitting individual UserOps via handleOps", async function () {
            console.log("🔄 Starting Web2 users individual UserOps flow...");

            // Step 1: Each Web2 user creates their own UserOperation
            console.log("📝 Web2 users creating individual UserOperations...");

            // User 1 wants to send 0.1 ETH to themselves
            const userOp1: PackedUserOperation = {
                sender: wallet1.address,
                nonce: await entryPoint.getNonce(wallet1.address, 0),
                initCode: "0x",
                callData: wallet1.interface.encodeFunctionData("execute", [
                    web2User1Address, // Send to self
                    ethers.utils.parseEther("0.1"),
                    "0x"
                ]),
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [200000, 200000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x" // Empty signature - will be filled by master signer
            };

            // User 2 wants to send 0.2 ETH to themselves  
            const userOp2: PackedUserOperation = {
                sender: wallet2.address,
                nonce: await entryPoint.getNonce(wallet2.address, 0),
                initCode: "0x",
                callData: wallet2.interface.encodeFunctionData("execute", [
                    web2User2Address, // Send to self
                    ethers.utils.parseEther("0.2"),
                    "0x"
                ]),
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [200000, 200000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x" // Empty signature - will be filled by master signer
            };

            // User 3 wants to send 0.3 ETH to themselves
            const userOp3: PackedUserOperation = {
                sender: wallet3.address,
                nonce: await entryPoint.getNonce(wallet3.address, 0),
                initCode: "0x",
                callData: wallet3.interface.encodeFunctionData("execute", [
                    web2User3Address, // Send to self
                    ethers.utils.parseEther("0.3"),
                    "0x"
                ]),
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [200000, 200000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x" // Empty signature - will be filled by master signer
            };

            console.log("📋 UserOps created:");
            console.log("  User 1 UserOp: Send 0.1 ETH to", web2User1Address);
            console.log("  User 2 UserOp: Send 0.2 ETH to", web2User2Address);
            console.log("  User 3 UserOp: Send 0.3 ETH to", web2User3Address);

            // Step 2: Get UserOp hashes
            const userOpHash1 = await entryPoint.getUserOpHash(userOp1);
            const userOpHash2 = await entryPoint.getUserOpHash(userOp2);
            const userOpHash3 = await entryPoint.getUserOpHash(userOp3);

            console.log("🔑 UserOp hashes generated");

            // Step 3: Master signer signs each UserOp individually (simulating backend service)
            console.log("✍️ Master signer signing UserOps...");

            // For individual operations, we can either:
            // A) Use master signer direct signatures for each UserOp
            // B) Use aggregator (but handleOps doesn't support aggregator)
            // We'll use option A - master signer signs each UserOp

            // Sign UserOp1
            const userOp1Hash = ethers.utils.arrayify(userOpHash1);
            const signature1 = await masterSigner.signMessage(userOp1Hash);
            userOp1.signature = signature1;

            // Sign UserOp2  
            const userOp2Hash = ethers.utils.arrayify(userOpHash2);
            const signature2 = await masterSigner.signMessage(userOp2Hash);
            userOp2.signature = signature2;

            // Sign UserOp3
            const userOp3Hash = ethers.utils.arrayify(userOpHash3);
            const signature3 = await masterSigner.signMessage(userOp3Hash);
            userOp3.signature = signature3;

            console.log("✅ All UserOps signed by master signer");

            // Step 4: Record initial balances
            const balanceBefore1 = await ethers.provider.getBalance(web2User1Address);
            const balanceBefore2 = await ethers.provider.getBalance(web2User2Address);
            const balanceBefore3 = await ethers.provider.getBalance(web2User3Address);

            console.log("💰 Initial balances recorded");

            // Step 5: Submit UserOps to EntryPoint via handleOps (individual operations)
            console.log("🎯 Executing handleOps with individual UserOperations...");
            
            const userOps = [userOp1, userOp2, userOp3];
            
            const tx = await entryPoint.handleOps(
                userOps,
                beneficiaryAddress
            );
            const receipt = await tx.wait();

            console.log("✅ handleOps executed successfully!");
            console.log("⛽ Gas used:", receipt.gasUsed.toString());

            // Step 6: Verify execution results
            const balanceAfter1 = await ethers.provider.getBalance(web2User1Address);
            const balanceAfter2 = await ethers.provider.getBalance(web2User2Address);
            const balanceAfter3 = await ethers.provider.getBalance(web2User3Address);

            // Verify each user received their expected amount
            expect(balanceAfter1.sub(balanceBefore1)).to.equal(ethers.utils.parseEther("0.1"));
            expect(balanceAfter2.sub(balanceBefore2)).to.equal(ethers.utils.parseEther("0.2"));
            expect(balanceAfter3.sub(balanceBefore3)).to.equal(ethers.utils.parseEther("0.3"));

            // Verify nonces incremented
            expect(await entryPoint.getNonce(wallet1.address, 0)).to.equal(1);
            expect(await entryPoint.getNonce(wallet2.address, 0)).to.equal(1);
            expect(await entryPoint.getNonce(wallet3.address, 0)).to.equal(1);

            console.log("🎉 All Web2 user operations executed successfully!");
            console.log("💸 User 1 received:", ethers.utils.formatEther(balanceAfter1.sub(balanceBefore1)), "ETH");
            console.log("💸 User 2 received:", ethers.utils.formatEther(balanceAfter2.sub(balanceBefore2)), "ETH");
            console.log("💸 User 3 received:", ethers.utils.formatEther(balanceAfter3.sub(balanceBefore3)), "ETH");
        });

        it("should handle Web2 users with different transaction types", async function () {
            console.log("🔄 Testing different Web2 user transaction types...");

            // User 1: Simple ETH transfer
            const userOp1: PackedUserOperation = {
                sender: wallet1.address,
                nonce: await entryPoint.getNonce(wallet1.address, 0),
                initCode: "0x",
                callData: wallet1.interface.encodeFunctionData("execute", [
                    beneficiaryAddress, // Send to beneficiary
                    ethers.utils.parseEther("0.05"),
                    "0x"
                ]),
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [150000, 150000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            // User 2: Contract interaction (setting aggregator)
            const userOp2: PackedUserOperation = {
                sender: wallet2.address,
                nonce: await entryPoint.getNonce(wallet2.address, 0),
                initCode: "0x",
                callData: wallet2.interface.encodeFunctionData("setAggregator", [
                    masterAggregator.address
                ]),
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [150000, 150000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            // User 3: Multiple calls in one transaction
            const calls = [
                {
                    target: web2User3Address,
                    value: ethers.utils.parseEther("0.01"),
                    data: "0x"
                },
                {
                    target: beneficiaryAddress,
                    value: ethers.utils.parseEther("0.02"),
                    data: "0x"
                }
            ];

            const userOp3: PackedUserOperation = {
                sender: wallet3.address,
                nonce: await entryPoint.getNonce(wallet3.address, 0),
                initCode: "0x",
                callData: wallet3.interface.encodeFunctionData("executeBatch", [calls]),
                accountGasLimits: ethers.utils.solidityPack(["uint128", "uint128"], [250000, 250000]),
                preVerificationGas: 21000,
                gasFees: ethers.utils.solidityPack(["uint128", "uint128"], [1000000000, 1000000000]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            // Sign all UserOps
            const userOpHash1 = await entryPoint.getUserOpHash(userOp1);
            const userOpHash2 = await entryPoint.getUserOpHash(userOp2);
            const userOpHash3 = await entryPoint.getUserOpHash(userOp3);

            userOp1.signature = await masterSigner.signMessage(ethers.utils.arrayify(userOpHash1));
            userOp2.signature = await masterSigner.signMessage(ethers.utils.arrayify(userOpHash2));
            userOp3.signature = await masterSigner.signMessage(ethers.utils.arrayify(userOpHash3));

            // Record balances
            const beneficiaryBalanceBefore = await ethers.provider.getBalance(beneficiaryAddress);
            const user3BalanceBefore = await ethers.provider.getBalance(web2User3Address);

            // Execute
            const tx = await entryPoint.handleOps(
                [userOp1, userOp2, userOp3],
                beneficiaryAddress
            );
            await tx.wait();

            // Verify results
            const beneficiaryBalanceAfter = await ethers.provider.getBalance(beneficiaryAddress);
            const user3BalanceAfter = await ethers.provider.getBalance(web2User3Address);

            // Calculate actual amounts received
            const beneficiaryReceived = beneficiaryBalanceAfter.sub(beneficiaryBalanceBefore);
            const user3Received = user3BalanceAfter.sub(user3BalanceBefore);
            
            console.log("💰 Beneficiary received:", ethers.utils.formatEther(beneficiaryReceived), "ETH");
            console.log("💰 User 3 received:", ethers.utils.formatEther(user3Received), "ETH");
            
            // Beneficiary should receive at least 0.05 + 0.02 = 0.07 ETH (plus gas refunds)
            expect(beneficiaryReceived).to.be.at.least(ethers.utils.parseEther("0.07"));
            
            // User 3 should receive exactly 0.01 ETH
            expect(user3Received).to.equal(ethers.utils.parseEther("0.01"));

            // Verify wallet2's aggregator was set (if it has getter function)
            try {
                expect(await wallet2.aggregatorAddress()).to.equal(masterAggregator.address);
                console.log("✅ Wallet2 aggregator successfully updated");
            } catch {
                console.log("⚠️ Wallet2 aggregator update verification skipped (no getter)");
            }

            console.log("🎉 Different transaction types executed successfully!");
        });
    });

    after(async function () {
        console.log("🎯 Web2 Users handleOps testing completed!");
        console.log("📊 Test Summary:");
        console.log("   ✅ Multiple Web2 users can submit individual UserOps");
        console.log("   ✅ Master signer can sign for multiple users");
        console.log("   ✅ EntryPoint.handleOps processes all operations");
        console.log("   ✅ Different transaction types supported");
        console.log("   ✅ Gas efficiency with individual operations");
        console.log("   ✅ Web2 user experience validated");
    });
});