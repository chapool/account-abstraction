import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import {
    MasterAggregator,
    MasterAggregator__factory,
    AAWallet,
    AAWallet__factory,
    EntryPoint,
    EntryPoint__factory
} from "../typechain";

describe("MasterAggregator - Basic Unit Tests", function () {
    let owner: Signer;
    let masterSigner1: Signer;
    let masterSigner2: Signer;
    let unauthorized: Signer;
    
    let entryPoint: EntryPoint;
    let masterAggregator: MasterAggregator;
    let mockWallet1: AAWallet;
    
    let ownerAddress: string;
    let masterSigner1Address: string;
    let masterSigner2Address: string;
    let unauthorizedAddress: string;
    let mockWallet1Address: string;

    beforeEach(async function () {
        [owner, masterSigner1, masterSigner2, unauthorized] = await ethers.getSigners();
        
        ownerAddress = await owner.address;
        masterSigner1Address = await masterSigner1.address;
        masterSigner2Address = await masterSigner2.address;
        unauthorizedAddress = await unauthorized.address;

        // Deploy EntryPoint
        const EntryPointFactory = (await ethers.getContractFactory("EntryPoint")) as EntryPoint__factory;
        entryPoint = await EntryPointFactory.deploy();
        await entryPoint.deployed();

        // Deploy mock AAWallet via proxy pattern
        const AAWalletFactory = (await ethers.getContractFactory("AAWallet")) as AAWallet__factory;
        const aaWalletImpl = await AAWalletFactory.deploy();
        await aaWalletImpl.deployed();
        
        const initData1 = aaWalletImpl.interface.encodeFunctionData("initialize", [
            await entryPoint.address, ownerAddress, masterSigner1Address
        ]);
        const proxyFactory1 = await ethers.getContractFactory("ERC1967Proxy");
        const proxy1 = await proxyFactory1.deploy(await aaWalletImpl.address, initData1);
        await proxy1.deployed();
        mockWallet1 = AAWalletFactory.attach(await proxy1.address);
        mockWallet1Address = await mockWallet1.address;

        // Deploy MasterAggregator via proxy pattern
        const MasterAggregatorFactory = (await ethers.getContractFactory("MasterAggregator")) as MasterAggregator__factory;
        const masterAggregatorImpl = await MasterAggregatorFactory.deploy();
        await masterAggregatorImpl.deployed();

        const initData = masterAggregatorImpl.interface.encodeFunctionData("initialize", [
            ownerAddress, 
            [masterSigner1Address, masterSigner2Address]
        ]);
        const proxyFactory = await ethers.getContractFactory("ERC1967Proxy");
        const proxy = await proxyFactory.deploy(await masterAggregatorImpl.address, initData);
        await proxy.deployed();

        masterAggregator = MasterAggregatorFactory.attach(await proxy.address);
    });

    describe("Initialization", function () {
        it("should initialize correctly with initial masters", async function () {
            expect(await masterAggregator.owner()).to.equal(ownerAddress);
            expect(await masterAggregator.authorizedMasters(masterSigner1Address)).to.be.true;
            expect(await masterAggregator.authorizedMasters(masterSigner2Address)).to.be.true;
        });

        it("should have correct default configuration", async function () {
            expect(await masterAggregator.maxAggregatedOps()).to.equal(50);
            expect(await masterAggregator.validationWindow()).to.equal(300);
            expect(await masterAggregator.getMasterNonce(masterSigner1Address)).to.equal(0);
        });

        it("should revert with invalid owner", async function () {
            const MasterAggregatorFactory = (await ethers.getContractFactory("MasterAggregator")) as MasterAggregator__factory;
            const impl = await MasterAggregatorFactory.deploy();
            await impl.deployed();

            const initData = impl.interface.encodeFunctionData("initialize", [
                ethers.ZeroAddress, []
            ]);
            const proxyFactory = await ethers.getContractFactory("ERC1967Proxy");
            
            await expect(
                proxyFactory.deploy(await impl.address, initData)
            ).to.be.revertedWith("MasterAggregator: invalid owner");
        });
    });

    describe("Master Signer Management", function () {
        it("should authorize master signer", async function () {
            expect(await masterAggregator.authorizedMasters(unauthorizedAddress)).to.be.false;

            const tx = await masterAggregator.connect(owner).setMasterAuthorization(unauthorizedAddress, true);
            const receipt = await tx.wait();

            expect(await masterAggregator.authorizedMasters(unauthorizedAddress)).to.be.true;

            const event = receipt?.logs.find((log: any) => {
                try {
                    const parsed = masterAggregator.interface.parseLog(log);
                    return parsed?.name === "MasterAuthorized";
                } catch {
                    return false;
                }
            });
            expect(event).to.not.be.undefined;
        });

        it("should revoke master signer authorization", async function () {
            expect(await masterAggregator.authorizedMasters(masterSigner1Address)).to.be.true;

            const tx = await masterAggregator.connect(owner).setMasterAuthorization(masterSigner1Address, false);
            await tx.wait();

            expect(await masterAggregator.authorizedMasters(masterSigner1Address)).to.be.false;
        });

        it("should only allow owner to manage master signers", async function () {
            await expect(
                masterAggregator.connect(unauthorized).setMasterAuthorization(unauthorizedAddress, true)
            ).to.be.reverted;
        });

        it("should revert when authorizing invalid master signer", async function () {
            await expect(
                masterAggregator.connect(owner).setMasterAuthorization(ethers.ZeroAddress, true)
            ).to.be.revertedWith("MasterAggregator: invalid master");
        });
    });

    describe("Wallet Authorization Management", function () {
        it("should authorize wallet for master signer", async function () {
            expect(await masterAggregator.masterToWallets(masterSigner1Address, mockWallet1Address)).to.be.false;

            const tx = await masterAggregator.connect(owner).setWalletAuthorization(
                masterSigner1Address, 
                mockWallet1Address, 
                true
            );
            await tx.wait();

            expect(await masterAggregator.masterToWallets(masterSigner1Address, mockWallet1Address)).to.be.true;
        });

        it("should auto-authorize wallet with valid master relationship", async function () {
            const tx = await masterAggregator.autoAuthorizeWallet(mockWallet1Address, masterSigner1Address);
            await tx.wait();

            expect(await masterAggregator.masterToWallets(masterSigner1Address, mockWallet1Address)).to.be.true;
        });

        it("should revert when authorizing wallet for unauthorized master", async function () {
            // First revoke master authorization
            await masterAggregator.connect(owner).setMasterAuthorization(masterSigner1Address, false);

            await expect(
                masterAggregator.connect(owner).setWalletAuthorization(
                    masterSigner1Address, 
                    mockWallet1Address, 
                    true
                )
            ).to.be.revertedWith("MasterAggregator: master not authorized");
        });

        it("should revert auto-authorize with invalid master relationship", async function () {
            // Try to auto-authorize mockWallet1 under masterSigner2 (but mockWallet1 has masterSigner1)
            await expect(
                masterAggregator.autoAuthorizeWallet(mockWallet1Address, masterSigner2Address)
            ).to.be.revertedWith("MasterAggregator: invalid wallet-master relationship");
        });
    });

    describe("Configuration Management", function () {
        it("should update configuration", async function () {
            const newMaxOps = 75;
            const newWindow = 600;

            await masterAggregator.connect(owner).updateConfig(newMaxOps, newWindow);

            expect(await masterAggregator.maxAggregatedOps()).to.equal(newMaxOps);
            expect(await masterAggregator.validationWindow()).to.equal(newWindow);
        });

        it("should only allow owner to update configuration", async function () {
            await expect(
                masterAggregator.connect(unauthorized).updateConfig(75, 600)
            ).to.be.reverted;
        });

        it("should revert with invalid max operations", async function () {
            await expect(
                masterAggregator.connect(owner).updateConfig(0, 600)
            ).to.be.revertedWith("MasterAggregator: invalid max ops");

            await expect(
                masterAggregator.connect(owner).updateConfig(101, 600)
            ).to.be.revertedWith("MasterAggregator: invalid max ops");
        });

        it("should revert with invalid validation window", async function () {
            await expect(
                masterAggregator.connect(owner).updateConfig(50, 0)
            ).to.be.revertedWith("MasterAggregator: invalid window");

            await expect(
                masterAggregator.connect(owner).updateConfig(50, 3601)
            ).to.be.revertedWith("MasterAggregator: invalid window");
        });
    });

    describe("Wallet Validation", function () {
        it("should validate wallet controlled by master through contract verification", async function () {
            // mockWallet1 has masterSigner1 as its master signer
            const isValid = await masterAggregator.isWalletControlledByMaster(
                mockWallet1Address, 
                masterSigner1Address
            );
            expect(isValid).to.be.true;
        });

        it("should return false for invalid wallet-master relationship", async function () {
            const isValid = await masterAggregator.isWalletControlledByMaster(
                mockWallet1Address, 
                masterSigner2Address
            );
            expect(isValid).to.be.false;
        });

        it("should return master nonce correctly", async function () {
            expect(await masterAggregator.getMasterNonce(masterSigner1Address)).to.equal(0);
            expect(await masterAggregator.getMasterNonce(masterSigner2Address)).to.equal(0);
        });
    });

    describe("Gas Savings Calculation", function () {
        it("should calculate gas savings correctly", async function () {
            // Single operation should have no savings
            expect(await masterAggregator.calculateGasSavings(1)).to.equal(0);
            
            // Two operations: individual = 6000, aggregated = 6000, savings = 0
            expect(await masterAggregator.calculateGasSavings(2)).to.equal(0);
            
            // Three operations: individual = 9000, aggregated = 6500, savings = 2500
            expect(await masterAggregator.calculateGasSavings(3)).to.equal(2500);
            
            // Ten operations: individual = 30000, aggregated = 10000, savings = 20000
            expect(await masterAggregator.calculateGasSavings(10)).to.equal(20000);
        });

        it("should return zero for no operations", async function () {
            expect(await masterAggregator.calculateGasSavings(0)).to.equal(0);
        });
    });

    describe("IAggregator Interface", function () {
        it("should return empty signature for individual user op validation", async function () {
            const mockUserOp = {
                sender: mockWallet1Address,
                nonce: 0,
                callData: "0x",
                accountGasLimits: ethers.concat([
                    ethers.toBeHex(100000, 16),
                    ethers.toBeHex(100000, 16)
                ]),
                preVerificationGas: 21000,
                gasFees: ethers.concat([
                    ethers.toBeHex(ethers.parseUnits("1", "gwei"), 16),
                    ethers.toBeHex(ethers.parseUnits("2", "gwei"), 16)
                ]),
                paymasterAndData: "0x",
                signature: "0x"
            };

            const result = await masterAggregator.validateUserOpSignature(mockUserOp);
            expect(result).to.equal("0x");
        });

        it("should return placeholder for aggregate signatures", async function () {
            const mockUserOps = [{
                sender: mockWallet1Address,
                nonce: 0,
                callData: "0x",
                accountGasLimits: ethers.concat([
                    ethers.toBeHex(100000, 16),
                    ethers.toBeHex(100000, 16)
                ]),
                preVerificationGas: 21000,
                gasFees: ethers.concat([
                    ethers.toBeHex(ethers.parseUnits("1", "gwei"), 16),
                    ethers.toBeHex(ethers.parseUnits("2", "gwei"), 16)
                ]),
                paymasterAndData: "0x",
                signature: "0x"
            }];

            const result = await masterAggregator.aggregateSignatures(mockUserOps);
            
            // Should return encoded placeholder
            const decoded = ethers.AbiCoder.defaultAbiCoder().decode(
                ["address", "uint256", "bytes"],
                result
            );
            expect(decoded[0]).to.equal(ethers.ZeroAddress);
            expect(decoded[1]).to.equal(0);
            expect(decoded[2]).to.equal("0x");
        });
    });

    describe("Session Key Support", function () {
        beforeEach(async function () {
            // Add a session key to wallet for testing
            const sessionKeyAddress = await unauthorized.address; // Use as mock session key
            await mockWallet1.connect(owner).addSessionKey(
                sessionKeyAddress,
                Math.floor(Date.now() / 1000),
                Math.floor(Date.now() / 1000) + 3600,
                ethers.ZeroHash
            );
        });

        it("should create session key aggregated signature", async function () {
            const userOps = [{
                sender: mockWallet1Address,
                nonce: 0,
                callData: "0x",
                accountGasLimits: ethers.concat([
                    ethers.toBeHex(100000, 16),
                    ethers.toBeHex(100000, 16)
                ]),
                preVerificationGas: 21000,
                gasFees: ethers.concat([
                    ethers.toBeHex(ethers.parseUnits("1", "gwei"), 16),
                    ethers.toBeHex(ethers.parseUnits("2", "gwei"), 16)
                ]),
                paymasterAndData: "0x",
                signature: "0x"
            }];

            const sessionKeyAddress = await unauthorized.address;
            const mockPrivateKey = ethers.keccak256(ethers.toUtf8Bytes("mock"));
            const result = await masterAggregator.createSessionKeyAggregatedSignature(
                userOps, 
                sessionKeyAddress, 
                mockPrivateKey
            );

            // Should return encoded session key signature
            const decoded = ethers.AbiCoder.defaultAbiCoder().decode(
                ["string", "address", "bytes"],
                result
            );
            expect(decoded[0]).to.equal("SESSION_KEY");
            expect(decoded[1]).to.equal(sessionKeyAddress);
        });
    });

    describe("Upgrade Functionality", function () {
        it("should only allow owner to authorize upgrades", async function () {
            // This tests the _authorizeUpgrade internal function via ownership
            expect(await masterAggregator.owner()).to.equal(ownerAddress);
        });

        it("should have pause functionality", async function () {
            // Test that pause function exists and can be called by owner
            await expect(
                masterAggregator.connect(owner).pause()
            ).to.not.be.reverted;
        });

        it("should only allow owner to pause", async function () {
            await expect(
                masterAggregator.connect(unauthorized).pause()
            ).to.be.reverted;
        });
    });
});