import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import { CPOPToken } from "../typechain";

describe("CPOPToken Batch Transfer Operations", function () {
    let admin: Signer;
    let user1: Signer;
    let user2: Signer;
    let user3: Signer;
    let user4: Signer;
    let user5: Signer;
    
    let adminAddress: string;
    let user1Address: string;
    let user2Address: string;
    let user3Address: string;
    let user4Address: string;
    let user5Address: string;
    
    let token: CPOPToken;
    
    const INITIAL_SUPPLY = ethers.utils.parseEther("1000000");
    
    // Role constants
    const ADMIN_ROLE = 1;  // 0001
    const MINTER_ROLE = 2; // 0010  
    const BURNER_ROLE = 4; // 0100

    beforeEach(async function () {
        [admin, user1, user2, user3, user4, user5] = await ethers.getSigners();
        
        adminAddress = await admin.getAddress();
        user1Address = await user1.getAddress();
        user2Address = await user2.getAddress();
        user3Address = await user3.getAddress();
        user4Address = await user4.getAddress();
        user5Address = await user5.getAddress();

        const TokenFactory = await ethers.getContractFactory("contracts/cpop/CPOPToken.sol:CPOPToken");
        token = await TokenFactory.deploy(adminAddress, INITIAL_SUPPLY);
        await token.deployed();
        
        // Give some tokens to users for testing
        const transferAmount = ethers.utils.parseEther("10000");
        await token.transfer(user1Address, transferAmount);
        await token.transfer(user2Address, transferAmount);
        await token.transfer(user3Address, transferAmount);
    });

    describe("Batch Transfer Operations", function () {
        it("should transfer tokens to multiple recipients in one transaction", async function () {
            const recipients = [user4Address, user5Address];
            const amounts = [
                ethers.utils.parseEther("100"),
                ethers.utils.parseEther("200")
            ];
            
            const initialSenderBalance = await token.balanceOf(user1Address);
            const initialRecipient1Balance = await token.balanceOf(user4Address);
            const initialRecipient2Balance = await token.balanceOf(user5Address);
            
            const tx = await token.connect(user1).batchTransfer(recipients, amounts);
            const receipt = await tx.wait();
            
            console.log(`Batch transfer gas for 2 recipients: ${receipt.gasUsed.toString()}`);
            
            // Check sender balance
            const expectedSenderBalance = initialSenderBalance.sub(amounts[0]).sub(amounts[1]);
            expect(await token.balanceOf(user1Address)).to.equal(expectedSenderBalance);
            
            // Check recipient balances
            expect(await token.balanceOf(user4Address)).to.equal(initialRecipient1Balance.add(amounts[0]));
            expect(await token.balanceOf(user5Address)).to.equal(initialRecipient2Balance.add(amounts[1]));
            
            // Check Transfer events
            const transferEvents = receipt.events?.filter(e => e.event === "Transfer");
            expect(transferEvents?.length).to.equal(2);
        });

        it("should handle large batch transfer operations efficiently", async function () {
            const recipients = [user2Address, user3Address, user4Address, user5Address];
            const amounts = [
                ethers.utils.parseEther("50"),
                ethers.utils.parseEther("75"),
                ethers.utils.parseEther("100"),
                ethers.utils.parseEther("125")
            ];
            
            const tx = await token.connect(user1).batchTransfer(recipients, amounts);
            const receipt = await tx.wait();
            
            console.log(`Batch transfer gas for 4 recipients: ${receipt.gasUsed.toString()}`);
            
            // Should be efficient
            expect(receipt.gasUsed.toNumber()).to.be.lessThan(200000);
            
            // Verify all recipients received tokens
            for (let i = 0; i < recipients.length; i++) {
                const balance = await token.balanceOf(recipients[i]);
                expect(balance).to.be.gte(amounts[i]); // At least the transferred amount
            }
        });

        it("should reject batch transfer with insufficient balance", async function () {
            const recipients = [user4Address, user5Address];
            const amounts = [
                ethers.utils.parseEther("5000"), // User1 has 10000 tokens
                ethers.utils.parseEther("6000")  // Total 11000 > 10000
            ];
            
            await expect(
                token.connect(user1).batchTransfer(recipients, amounts)
            ).to.be.reverted;
        });

        it("should reject batch transfer with mismatched array lengths", async function () {
            const recipients = [user4Address, user5Address];
            const amounts = [ethers.utils.parseEther("100")]; // Length mismatch
            
            await expect(
                token.connect(user1).batchTransfer(recipients, amounts)
            ).to.be.reverted;
        });

        it("should reject empty batch transfer", async function () {
            const recipients: string[] = [];
            const amounts: any[] = [];
            
            await expect(
                token.connect(user1).batchTransfer(recipients, amounts)
            ).to.be.reverted;
        });
    });

    describe("Batch TransferFrom Operations", function () {
        beforeEach(async function () {
            // Setup approvals for testing
            await token.connect(user1).approve(user2Address, ethers.utils.parseEther("1000"));
            await token.connect(user1).approve(adminAddress, ethers.utils.parseEther("1000"));
            await token.connect(user2).approve(adminAddress, ethers.utils.parseEther("1000"));
            await token.connect(user3).approve(user2Address, ethers.utils.parseEther("1000"));
            await token.connect(user3).approve(adminAddress, ethers.utils.parseEther("1000"));
        });

        it("should transfer from multiple sources with proper allowances", async function () {
            const from = [user1Address, user3Address];
            const to = [user4Address, user5Address];
            const amounts = [
                ethers.utils.parseEther("100"),
                ethers.utils.parseEther("200")
            ];
            
            const tx = await token.connect(user2).batchTransferFrom(from, to, amounts);
            const receipt = await tx.wait();
            
            console.log(`Batch transferFrom gas for 2 transfers: ${receipt.gasUsed.toString()}`);
            
            // Check balances
            expect(await token.balanceOf(user4Address)).to.equal(amounts[0]);
            expect(await token.balanceOf(user5Address)).to.equal(amounts[1]);
            
            // Check allowances were reduced
            expect(await token.allowance(user1Address, user2Address)).to.equal(
                ethers.utils.parseEther("900") // 1000 - 100
            );
            expect(await token.allowance(user3Address, user2Address)).to.equal(
                ethers.utils.parseEther("800") // 1000 - 200
            );
        });

        it("should allow admin to transfer without allowance restrictions", async function () {
            const from = [user1Address, user2Address];
            const to = [user4Address, user5Address];
            const amounts = [
                ethers.utils.parseEther("500"),
                ethers.utils.parseEther("600")
            ];
            
            // Check initial balances
            const initialUser1Balance = await token.balanceOf(user1Address);
            const initialUser2Balance = await token.balanceOf(user2Address);
            
            // Admin can transfer without allowance restrictions
            const tx = await token.connect(admin).batchTransferFrom(from, to, amounts);
            const receipt = await tx.wait();
            
            console.log(`Admin batch transferFrom gas: ${receipt.gasUsed.toString()}`);
            
            // Check transfers succeeded
            expect(await token.balanceOf(user4Address)).to.equal(amounts[0]);
            expect(await token.balanceOf(user5Address)).to.equal(amounts[1]);
            
            // Check from addresses had tokens deducted
            expect(await token.balanceOf(user1Address)).to.equal(initialUser1Balance.sub(amounts[0]));
            expect(await token.balanceOf(user2Address)).to.equal(initialUser2Balance.sub(amounts[1]));
            
            // Admin operations don't affect allowances since admin bypasses them
            expect(await token.allowance(user1Address, adminAddress)).to.equal(
                ethers.utils.parseEther("1000") // Should remain unchanged
            );
            expect(await token.allowance(user2Address, adminAddress)).to.equal(
                ethers.utils.parseEther("1000") // Should remain unchanged
            );
        });

        it("should reject batch transferFrom with insufficient allowance", async function () {
            const from = [user1Address];
            const to = [user4Address];
            const amounts = [ethers.utils.parseEther("1500")]; // More than approved (1000)
            
            await expect(
                token.connect(user2).batchTransferFrom(from, to, amounts)
            ).to.be.reverted;
        });

        it("should reject batch transferFrom with mismatched array lengths", async function () {
            const from = [user1Address, user3Address];
            const to = [user4Address]; // Length mismatch
            const amounts = [ethers.utils.parseEther("100"), ethers.utils.parseEther("200")];
            
            await expect(
                token.connect(user2).batchTransferFrom(from, to, amounts)
            ).to.be.reverted;
        });

        it("should handle unlimited allowance correctly", async function () {
            // Set unlimited allowance
            await token.connect(user1).approve(user2Address, ethers.constants.MaxUint256);
            
            const from = [user1Address];
            const to = [user4Address];
            const amounts = [ethers.utils.parseEther("500")];
            
            const initialAllowance = await token.allowance(user1Address, user2Address);
            
            await token.connect(user2).batchTransferFrom(from, to, amounts);
            
            // Unlimited allowance should remain unchanged
            expect(await token.allowance(user1Address, user2Address)).to.equal(initialAllowance);
        });
    });

    describe("Gas Efficiency Comparison", function () {
        it("should compare batch vs individual transfer operations", async function () {
            console.log("\\n=== BATCH vs INDIVIDUAL TRANSFER COMPARISON ===");
            
            const recipients = [user2Address, user3Address, user4Address];
            const amounts = [
                ethers.utils.parseEther("100"),
                ethers.utils.parseEther("200"),
                ethers.utils.parseEther("300")
            ];
            
            // Reset balances for fair comparison
            const initialBalance = await token.balanceOf(user1Address);
            
            // Individual transfers
            let individualGas = 0;
            for (let i = 0; i < recipients.length; i++) {
                const tx = await token.connect(user1).transfer(recipients[i], amounts[i]);
                const receipt = await tx.wait();
                individualGas += receipt.gasUsed.toNumber();
            }
            
            // Transfer tokens back for batch test
            for (let i = 0; i < recipients.length; i++) {
                await token.connect(ethers.provider.getSigner(recipients[i])).transfer(user1Address, amounts[i]);
            }
            
            // Batch transfer
            const batchTx = await token.connect(user1).batchTransfer(recipients, amounts);
            const batchReceipt = await batchTx.wait();
            const batchGas = batchReceipt.gasUsed.toNumber();
            
            console.log(`Individual transfers (3x): ${individualGas} gas`);
            console.log(`Batch transfer (3x): ${batchGas} gas`);
            console.log(`Gas savings: ${individualGas - batchGas} gas (${((individualGas - batchGas) / individualGas * 100).toFixed(2)}%)`);
            
            expect(batchGas).to.be.lessThan(individualGas);
            console.log("=======================================\\n");
        });
    });

    describe("Real-world Scenarios", function () {
        it("should handle reward distribution scenario", async function () {
            console.log("\\n=== REWARD DISTRIBUTION SCENARIO ===");
            
            // Admin distributes rewards to multiple users
            const rewardRecipients = [user1Address, user2Address, user3Address, user4Address];
            const rewardAmounts = [
                ethers.utils.parseEther("50"),
                ethers.utils.parseEther("75"),
                ethers.utils.parseEther("100"),
                ethers.utils.parseEther("125")
            ];
            
            const tx = await token.connect(admin).batchTransfer(rewardRecipients, rewardAmounts);
            const receipt = await tx.wait();
            
            console.log(`Reward distribution to 4 users gas: ${receipt.gasUsed.toString()}`);
            console.log(`Average gas per recipient: ${Math.round(receipt.gasUsed.toNumber() / 4)}`);
            
            // Verify rewards distributed
            for (let i = 0; i < rewardRecipients.length; i++) {
                const balance = await token.balanceOf(rewardRecipients[i]);
                expect(balance).to.be.gte(rewardAmounts[i]);
            }
            
            console.log("=====================================\\n");
        });

        it("should handle admin batch redistribution scenario", async function () {
            console.log("\\n=== ADMIN BATCH REDISTRIBUTION SCENARIO ===");
            
            // Admin redistributes tokens between users without allowance
            const from = [user1Address, user2Address];
            const to = [user4Address, user5Address];
            const amounts = [
                ethers.utils.parseEther("200"),
                ethers.utils.parseEther("300")
            ];
            
            const tx = await token.connect(admin).batchTransferFrom(from, to, amounts);
            const receipt = await tx.wait();
            
            console.log(`Admin redistribution gas: ${receipt.gasUsed.toString()}`);
            console.log(`Average gas per redistribution: ${Math.round(receipt.gasUsed.toNumber() / 2)}`);
            
            // Verify redistribution
            expect(await token.balanceOf(user4Address)).to.equal(amounts[0]);
            expect(await token.balanceOf(user5Address)).to.equal(amounts[1]);
            
            console.log("==========================================\\n");
        });
    });
});