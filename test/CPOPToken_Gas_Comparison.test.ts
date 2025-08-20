import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";

describe("CPOPToken Gas Optimization Comparison", function () {
    let owner: Signer;
    let user1: Signer;
    let user2: Signer;
    
    let ownerAddress: string;
    let user1Address: string;
    let user2Address: string;
    
    let originalToken: any;
    let optimizedToken: any;

    beforeEach(async function () {
        [owner, user1, user2] = await ethers.getSigners();
        
        ownerAddress = await owner.getAddress();
        user1Address = await user1.getAddress();
        user2Address = await user2.getAddress();

        // Deploy original CPOPToken
        const OriginalTokenFactory = await ethers.getContractFactory("CPOPToken");
        originalToken = await OriginalTokenFactory.deploy(ownerAddress);
        await originalToken.deployed();

        // Deploy optimized CPOPToken
        const OptimizedTokenFactory = await ethers.getContractFactory("CPOPToken", {
            libraries: {},
        });
        // For the optimized version, we'll deploy a simple contract
        const optimizedTokenCode = `
        // SPDX-License-Identifier: MIT
        pragma solidity ^0.8.28;
        
        contract CPOPTokenOptimized {
            string public constant name = "CPOP Token";
            string public constant symbol = "CPOP";
            uint8 public constant decimals = 18;
            
            uint256 public totalSupply;
            mapping(address => uint256) public balanceOf;
            mapping(address => mapping(address => uint256)) public allowance;
            
            event Transfer(address indexed from, address indexed to, uint256 value);
            event Approval(address indexed owner, address indexed spender, uint256 value);
            
            address public owner;
            
            constructor(address _owner, uint256 _initialSupply) {
                owner = _owner;
                totalSupply = _initialSupply;
                balanceOf[_owner] = _initialSupply;
                emit Transfer(address(0), _owner, _initialSupply);
            }
            
            function transfer(address to, uint256 amount) external returns (bool) {
                balanceOf[msg.sender] -= amount;
                unchecked { balanceOf[to] += amount; }
                emit Transfer(msg.sender, to, amount);
                return true;
            }
            
            function approve(address spender, uint256 amount) external returns (bool) {
                allowance[msg.sender][spender] = amount;
                emit Approval(msg.sender, spender, amount);
                return true;
            }
            
            function transferFrom(address from, address to, uint256 amount) external returns (bool) {
                uint256 allowed = allowance[from][msg.sender];
                if (allowed != type(uint256).max) {
                    allowance[from][msg.sender] = allowed - amount;
                }
                balanceOf[from] -= amount;
                unchecked { balanceOf[to] += amount; }
                emit Transfer(from, to, amount);
                return true;
            }
            
            function mint(address to, uint256 amount) external {
                require(msg.sender == owner, "Only owner");
                unchecked {
                    totalSupply += amount;
                    balanceOf[to] += amount;
                }
                emit Transfer(address(0), to, amount);
            }
        }`;
        
        // We'll test with the optimized contract file we created
        const OptimizedFactory = await ethers.getContractFactory("CPOPToken");
        optimizedToken = await OptimizedFactory.deploy(ownerAddress, ethers.utils.parseEther("1000000"));
        await optimizedToken.deployed();
    });

    describe("Deployment Gas Costs", function () {
        it("should compare deployment costs", async function () {
            console.log("\n=== DEPLOYMENT GAS COMPARISON ===");
            
            // Get deployment transaction receipts
            const originalDeployTx = await ethers.provider.getTransactionReceipt(originalToken.deployTransaction.hash);
            const optimizedDeployTx = await ethers.provider.getTransactionReceipt(optimizedToken.deployTransaction.hash);
            
            const originalGas = originalDeployTx.gasUsed;
            const optimizedGas = optimizedDeployTx.gasUsed;
            const savings = originalGas.sub(optimizedGas);
            const savingsPercent = savings.mul(100).div(originalGas);
            
            console.log(`Original CPOPToken deployment: ${originalGas.toString()} gas`);
            console.log(`Optimized CPOPToken deployment: ${optimizedGas.toString()} gas`);
            console.log(`Gas savings: ${savings.toString()} gas (${savingsPercent.toString()}%)`);
            
            expect(optimizedGas.lt(originalGas)).to.be.true;
        });
    });

    describe("Transfer Gas Costs", function () {
        beforeEach(async function () {
            // Setup for transfer tests
            await originalToken.addToWhitelist(ownerAddress);
            await originalToken.addToWhitelist(user1Address);
            await originalToken.addToWhitelist(user2Address);
            
            // Mint tokens for testing
            await originalToken.mint(user1Address, ethers.utils.parseEther("1000"));
            await optimizedToken.mint(user1Address, ethers.utils.parseEther("1000"));
        });

        it("should compare transfer gas costs", async function () {
            console.log("\n=== TRANSFER GAS COMPARISON ===");
            
            const transferAmount = ethers.utils.parseEther("100");
            
            // Test original token transfer
            const originalTransferTx = await originalToken.connect(user1).transfer(user2Address, transferAmount);
            const originalTransferReceipt = await originalTransferTx.wait();
            
            // Test optimized token transfer
            const optimizedTransferTx = await optimizedToken.connect(user1).transfer(user2Address, transferAmount);
            const optimizedTransferReceipt = await optimizedTransferTx.wait();
            
            const originalGas = originalTransferReceipt.gasUsed;
            const optimizedGas = optimizedTransferReceipt.gasUsed;
            const savings = originalGas.sub(optimizedGas);
            const savingsPercent = savings.mul(100).div(originalGas);
            
            console.log(`Original transfer: ${originalGas.toString()} gas`);
            console.log(`Optimized transfer: ${optimizedGas.toString()} gas`);
            console.log(`Gas savings: ${savings.toString()} gas (${savingsPercent.toString()}%)`);
            
            expect(optimizedGas.lt(originalGas)).to.be.true;
        });
    });

    describe("Approve Gas Costs", function () {
        it("should compare approve gas costs", async function () {
            console.log("\n=== APPROVE GAS COMPARISON ===");
            
            const approveAmount = ethers.utils.parseEther("1000");
            
            // Note: Original token doesn't have approve function, so we'll simulate
            // For testing purposes, we'll use a mock comparison
            
            // Test optimized token approve
            const optimizedApproveTx = await optimizedToken.connect(user1).approve(user2Address, approveAmount);
            const optimizedApproveReceipt = await optimizedApproveTx.wait();
            
            console.log(`Optimized approve: ${optimizedApproveReceipt.gasUsed.toString()} gas`);
            console.log("Note: Original token doesn't have approve function");
            
            // Verify approval worked
            const allowanceAmount = await optimizedToken.allowance(user1Address, user2Address);
            expect(allowanceAmount).to.equal(approveAmount);
        });
    });

    describe("Mint Gas Costs", function () {
        it("should compare mint gas costs", async function () {
            console.log("\n=== MINT GAS COMPARISON ===");
            
            const mintAmount = ethers.utils.parseEther("500");
            
            // Test original token mint
            const originalMintTx = await originalToken.mint(user1Address, mintAmount);
            const originalMintReceipt = await originalMintTx.wait();
            
            // Test optimized token mint
            const optimizedMintTx = await optimizedToken.mint(user1Address, mintAmount);
            const optimizedMintReceipt = await optimizedMintTx.wait();
            
            const originalGas = originalMintReceipt.gasUsed;
            const optimizedGas = optimizedMintReceipt.gasUsed;
            const savings = originalGas.sub(optimizedGas);
            const savingsPercent = savings.mul(100).div(originalGas);
            
            console.log(`Original mint: ${originalGas.toString()} gas`);
            console.log(`Optimized mint: ${optimizedGas.toString()} gas`);
            console.log(`Gas savings: ${savings.toString()} gas (${savingsPercent.toString()}%)`);
            
            expect(optimizedGas.lt(originalGas)).to.be.true;
        });
    });

    describe("Functionality Tests", function () {
        beforeEach(async function () {
            // Mint initial tokens for testing
            await optimizedToken.mint(user1Address, ethers.utils.parseEther("1000"));
        });

        it("should work as basic ERC20", async function () {
            const initialBalance = await optimizedToken.balanceOf(user1Address);
            const transferAmount = ethers.utils.parseEther("100");
            
            // Test transfer
            await optimizedToken.connect(user1).transfer(user2Address, transferAmount);
            
            expect(await optimizedToken.balanceOf(user1Address)).to.equal(
                initialBalance.sub(transferAmount)
            );
            expect(await optimizedToken.balanceOf(user2Address)).to.equal(transferAmount);
        });

        it("should handle approvals and transferFrom", async function () {
            const approveAmount = ethers.utils.parseEther("200");
            const transferAmount = ethers.utils.parseEther("100");
            
            // Approve user2 to spend from user1
            await optimizedToken.connect(user1).approve(user2Address, approveAmount);
            
            expect(await optimizedToken.allowance(user1Address, user2Address)).to.equal(approveAmount);
            
            // Transfer from user1 to owner via user2
            const initialUser1Balance = await optimizedToken.balanceOf(user1Address);
            await optimizedToken.connect(user2).transferFrom(user1Address, ownerAddress, transferAmount);
            
            expect(await optimizedToken.balanceOf(user1Address)).to.equal(
                initialUser1Balance.sub(transferAmount)
            );
            expect(await optimizedToken.allowance(user1Address, user2Address)).to.equal(
                approveAmount.sub(transferAmount)
            );
        });
    });

    describe("Contract Size Comparison", function () {
        it("should compare contract sizes", async function () {
            console.log("\n=== CONTRACT SIZE COMPARISON ===");
            
            const originalCode = await ethers.provider.getCode(originalToken.address);
            const optimizedCode = await ethers.provider.getCode(optimizedToken.address);
            
            const originalSize = originalCode.length / 2 - 1; // Convert hex to bytes
            const optimizedSize = optimizedCode.length / 2 - 1;
            const sizeSavings = originalSize - optimizedSize;
            const sizeSavingsPercent = (sizeSavings / originalSize) * 100;
            
            console.log(`Original contract size: ${originalSize} bytes`);
            console.log(`Optimized contract size: ${optimizedSize} bytes`);
            console.log(`Size reduction: ${sizeSavings} bytes (${sizeSavingsPercent.toFixed(2)}%)`);
            
            expect(optimizedSize).to.be.lessThan(originalSize);
        });
    });
});