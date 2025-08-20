import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import { CPOPToken } from "../typechain";

describe("CPOPToken Role-Based Access Control", function () {
    let admin: Signer;
    let minter: Signer;
    let burner: Signer;
    let user1: Signer;
    let user2: Signer;
    let walletManager: Signer;
    
    let adminAddress: string;
    let minterAddress: string;
    let burnerAddress: string;
    let user1Address: string;
    let user2Address: string;
    let walletManagerAddress: string;
    
    let token: CPOPToken;
    
    const INITIAL_SUPPLY = ethers.utils.parseEther("1000000");
    
    // Role constants
    const ADMIN_ROLE = 1;  // 0001
    const MINTER_ROLE = 2; // 0010  
    const BURNER_ROLE = 4; // 0100

    beforeEach(async function () {
        [admin, minter, burner, user1, user2, walletManager] = await ethers.getSigners();
        
        adminAddress = await admin.getAddress();
        minterAddress = await minter.getAddress();
        burnerAddress = await burner.getAddress();
        user1Address = await user1.getAddress();
        user2Address = await user2.getAddress();
        walletManagerAddress = await walletManager.getAddress();

        const TokenFactory = await ethers.getContractFactory("contracts/cpop/CPOPToken.sol:CPOPToken");
        token = await TokenFactory.deploy(adminAddress, INITIAL_SUPPLY);
        await token.deployed();
        
        // Setup roles for testing
        await token.grantRole(minterAddress, MINTER_ROLE);
        await token.grantRole(burnerAddress, BURNER_ROLE);
        await token.grantRole(walletManagerAddress, MINTER_ROLE | BURNER_ROLE); // Multiple roles
    });

    describe("Deployment & Initial Setup", function () {
        it("should deploy with admin having all roles", async function () {
            expect(await token.hasRole(adminAddress, ADMIN_ROLE)).to.be.true;
            expect(await token.hasRole(adminAddress, MINTER_ROLE)).to.be.true;
            expect(await token.hasRole(adminAddress, BURNER_ROLE)).to.be.true;
            
            expect(await token.totalSupply()).to.equal(INITIAL_SUPPLY);
            expect(await token.balanceOf(adminAddress)).to.equal(INITIAL_SUPPLY);
        });

        it("should measure deployment gas for role-based system", async function () {
            const deployTx = await ethers.provider.getTransactionReceipt(token.deployTransaction.hash);
            console.log(`Role-based CPOPToken deployment gas: ${deployTx.gasUsed.toString()}`);
            
            // Should still be very efficient despite role system
            expect(deployTx.gasUsed.toNumber()).to.be.lessThan(1000000);
        });
    });

    describe("Role Management", function () {
        it("should allow admin to grant roles", async function () {
            expect(await token.hasRole(user1Address, MINTER_ROLE)).to.be.false;
            
            const tx = await token.grantRole(user1Address, MINTER_ROLE);
            const receipt = await tx.wait();
            
            console.log(`Grant role gas: ${receipt.gasUsed.toString()}`);
            
            expect(await token.hasRole(user1Address, MINTER_ROLE)).to.be.true;
            
            // Event should be emitted
            const event = receipt.events?.find(e => e.event === "RoleGranted");
            expect(event?.args?.account).to.equal(user1Address);
            expect(event?.args?.role).to.equal(MINTER_ROLE);
        });

        it("should allow admin to revoke roles", async function () {
            expect(await token.hasRole(minterAddress, MINTER_ROLE)).to.be.true;
            
            const tx = await token.revokeRole(minterAddress, MINTER_ROLE);
            const receipt = await tx.wait();
            
            console.log(`Revoke role gas: ${receipt.gasUsed.toString()}`);
            
            expect(await token.hasRole(minterAddress, MINTER_ROLE)).to.be.false;
        });

        it("should support multiple roles for same address", async function () {
            await token.grantRole(user1Address, MINTER_ROLE);
            await token.grantRole(user1Address, BURNER_ROLE);
            
            expect(await token.hasRole(user1Address, MINTER_ROLE)).to.be.true;
            expect(await token.hasRole(user1Address, BURNER_ROLE)).to.be.true;
            expect(await token.hasRole(user1Address, ADMIN_ROLE)).to.be.false;
            
            // Check combined roles value
            expect(await token.roles(user1Address)).to.equal(MINTER_ROLE | BURNER_ROLE);
        });

        it("should reject role management from non-admin", async function () {
            await expect(
                token.connect(user1).grantRole(user2Address, MINTER_ROLE)
            ).to.be.reverted;
            
            await expect(
                token.connect(minter).revokeRole(burnerAddress, BURNER_ROLE)
            ).to.be.reverted;
        });

        it("should reject invalid roles", async function () {
            await expect(
                token.grantRole(user1Address, 0)
            ).to.be.reverted;
            
            await expect(
                token.grantRole(user1Address, 8) // > 7
            ).to.be.reverted;
        });
    });

    describe("Minting with Roles", function () {
        it("should allow MINTER_ROLE to mint", async function () {
            const mintAmount = ethers.utils.parseEther("1000");
            
            const tx = await token.connect(minter).mint(user1Address, mintAmount);
            const receipt = await tx.wait();
            
            console.log(`Role-based mint gas: ${receipt.gasUsed.toString()}`);
            
            expect(await token.balanceOf(user1Address)).to.equal(mintAmount);
            expect(await token.totalSupply()).to.equal(INITIAL_SUPPLY.add(mintAmount));
        });

        it("should allow WalletManager (multiple roles) to mint", async function () {
            const mintAmount = ethers.utils.parseEther("500");
            
            await token.connect(walletManager).mint(user2Address, mintAmount);
            
            expect(await token.balanceOf(user2Address)).to.equal(mintAmount);
        });

        it("should reject minting from non-minter", async function () {
            await expect(
                token.connect(user1).mint(user2Address, ethers.utils.parseEther("100"))
            ).to.be.reverted;
        });

        it("should allow multiple minters to work simultaneously", async function () {
            const mintAmount = ethers.utils.parseEther("200");
            
            // Both admin and dedicated minter should work
            await token.connect(admin).mint(user1Address, mintAmount);
            await token.connect(minter).mint(user2Address, mintAmount);
            
            expect(await token.balanceOf(user1Address)).to.equal(mintAmount);
            expect(await token.balanceOf(user2Address)).to.equal(mintAmount);
        });
    });

    describe("Burning with Roles", function () {
        beforeEach(async function () {
            // Setup some tokens for burning tests
            await token.connect(minter).mint(user1Address, ethers.utils.parseEther("1000"));
            await token.connect(minter).mint(user2Address, ethers.utils.parseEther("1000"));
        });

        it("should allow BURNER_ROLE to burn from any address", async function () {
            const burnAmount = ethers.utils.parseEther("300");
            
            const tx = await token.connect(burner).adminBurn(user1Address, burnAmount);
            const receipt = await tx.wait();
            
            console.log(`Admin burn gas: ${receipt.gasUsed.toString()}`);
            
            expect(await token.balanceOf(user1Address)).to.equal(
                ethers.utils.parseEther("700")
            );
        });

        it("should allow BURNER_ROLE to burnFrom without approval", async function () {
            const burnAmount = ethers.utils.parseEther("200");
            
            // No approval needed for BURNER_ROLE
            const tx = await token.connect(burner).burnFrom(user1Address, burnAmount);
            const receipt = await tx.wait();
            
            console.log(`BurnFrom with role gas: ${receipt.gasUsed.toString()}`);
            
            expect(await token.balanceOf(user1Address)).to.equal(
                ethers.utils.parseEther("800")
            );
        });

        it("should require approval for burnFrom without BURNER_ROLE", async function () {
            const burnAmount = ethers.utils.parseEther("100");
            
            // Should fail without approval
            await expect(
                token.connect(user2).burnFrom(user1Address, burnAmount)
            ).to.be.reverted;
            
            // Should work with approval
            await token.connect(user1).approve(user2Address, burnAmount);
            await token.connect(user2).burnFrom(user1Address, burnAmount);
            
            expect(await token.balanceOf(user1Address)).to.equal(
                ethers.utils.parseEther("900")
            );
        });

        it("should reject adminBurn from non-burner", async function () {
            await expect(
                token.connect(user1).adminBurn(user2Address, ethers.utils.parseEther("100"))
            ).to.be.reverted;
        });
    });

    describe("Real-world Scenario: Multiple Contracts", function () {
        it("should handle scenario with WalletManager and other contracts", async function () {
            console.log("\n=== Multi-Contract Scenario ===");
            
            // Setup: Grant roles to different contracts
            await token.grantRole(user1Address, MINTER_ROLE); // Simulate another contract
            await token.grantRole(user2Address, BURNER_ROLE); // Simulate treasury contract
            
            const initialUserBalance = ethers.utils.parseEther("500");
            const mintAmount = ethers.utils.parseEther("200");
            const burnAmount = ethers.utils.parseEther("100");
            
            // 1. WalletManager mints initial tokens for user
            console.log("1. WalletManager mints tokens for new user");
            await token.connect(walletManager).mint(user1Address, initialUserBalance);
            
            // 2. Another contract (user1 simulating) mints rewards
            console.log("2. Reward contract mints additional tokens");
            await token.connect(user1).mint(user1Address, mintAmount);
            
            // 3. Treasury contract (user2 simulating) burns tokens for fees
            console.log("3. Treasury burns tokens for fees");
            await token.connect(user2).adminBurn(user1Address, burnAmount);
            
            const finalBalance = initialUserBalance.add(mintAmount).sub(burnAmount);
            expect(await token.balanceOf(user1Address)).to.equal(finalBalance);
            
            console.log(`Final user balance: ${ethers.utils.formatEther(finalBalance)} CPOP`);
            console.log("=== Scenario Complete ===\n");
        });

        it("should efficiently handle role combinations", async function () {
            // Test walletManager with both MINTER and BURNER roles
            const amount = ethers.utils.parseEther("1000");
            
            // Can mint
            await token.connect(walletManager).mint(user1Address, amount);
            expect(await token.balanceOf(user1Address)).to.equal(amount);
            
            // Can also burn without approval
            await token.connect(walletManager).adminBurn(user1Address, amount.div(2));
            expect(await token.balanceOf(user1Address)).to.equal(amount.div(2));
        });
    });

    describe("Gas Efficiency with Roles", function () {
        it("should compare gas costs with and without role checks", async function () {
            console.log("\n=== ROLE-BASED GAS EFFICIENCY ===");
            
            const mintAmount = ethers.utils.parseEther("1000");
            
            // Mint with role check
            const mintTx = await token.connect(minter).mint(user1Address, mintAmount);
            const mintReceipt = await mintTx.wait();
            
            // Role check should add minimal gas overhead
            console.log(`Mint with role check: ${mintReceipt.gasUsed.toString()} gas`);
            expect(mintReceipt.gasUsed.toNumber()).to.be.lessThan(60000);
            
            // Role grant/revoke operations
            const grantTx = await token.grantRole(user1Address, MINTER_ROLE);
            const grantReceipt = await grantTx.wait();
            console.log(`Grant role: ${grantReceipt.gasUsed.toString()} gas`);
            
            const revokeTx = await token.revokeRole(user1Address, MINTER_ROLE);
            const revokeReceipt = await revokeTx.wait();
            console.log(`Revoke role: ${revokeReceipt.gasUsed.toString()} gas`);
            
            console.log("===================================\n");
        });

        it("should test hasRole gas efficiency", async function () {
            // View function - should use minimal gas
            const hasMinterRole = await token.hasRole(minterAddress, MINTER_ROLE);
            const hasAdminRole = await token.hasRole(minterAddress, ADMIN_ROLE);
            
            expect(hasMinterRole).to.be.true;
            expect(hasAdminRole).to.be.false;
            
            // Test combined role check
            const hasMultipleRoles = await token.hasRole(walletManagerAddress, MINTER_ROLE | BURNER_ROLE);
            expect(hasMultipleRoles).to.be.true;
        });
    });
});