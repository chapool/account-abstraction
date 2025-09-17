import { ethers } from "hardhat";

async function main() {
    const [deployer] = await ethers.getSigners();
    
    // Contract addresses
    const PAYMENT_ADDRESS = "0xC81bac959087100BE02B5C599118Fdc04c56556d";
    const MOCK_USDT_ADDRESS = "0x8c3346F5A95cB5927fE09C6265b09eEA607887d6";
    
    // Payment parameters
    const ORDER_ID = 1; // Change this for different orders
    const AMOUNT_USDT = 1000; // Amount in USDT (will be converted to 6 decimals)
    
    console.log(`Making payment: Order ${ORDER_ID}, ${AMOUNT_USDT} USDT`);
    console.log("Payer:", deployer.address);
    
    try {
        // Get contracts
        const Payment = await ethers.getContractFactory("Payment");
        const payment = Payment.attach(PAYMENT_ADDRESS);
        
        const MockUSDT = await ethers.getContractFactory("MockUSDT");
        const mockUSDT = MockUSDT.attach(MOCK_USDT_ADDRESS);
        
        // Check if order already paid
        const isPaid = await payment.isOrderPaid(ORDER_ID);
        if (isPaid) {
            console.log("❌ Order already paid!");
            return;
        }
        
        // Convert amount to wei (6 decimals for USDT)
        const amount = ethers.utils.parseUnits(AMOUNT_USDT.toString(), 6);
        
        // Check balance
        const balance = await mockUSDT.balanceOf(deployer.address);
        if (balance.lt(amount)) {
            console.log("Getting USDT from faucet...");
            const faucetTx = await mockUSDT.faucet();
            await faucetTx.wait();
            console.log("✅ Faucet completed");
        }
        
        // Approve spending
        const allowance = await mockUSDT.allowance(deployer.address, PAYMENT_ADDRESS);
        if (allowance.lt(amount)) {
            console.log("Approving USDT spending...");
            const approveTx = await mockUSDT.approve(PAYMENT_ADDRESS, amount);
            await approveTx.wait();
            console.log("✅ Approval completed");
        }
        
        // Make payment
        console.log("Processing payment...");
        const tx = await payment.pay(ORDER_ID, MOCK_USDT_ADDRESS, amount);
        const receipt = await tx.wait();
        
        console.log("🎉 Payment successful!");
        console.log("Transaction:", tx.hash);
        console.log("Block:", receipt.blockNumber);
        console.log("Gas used:", receipt.gasUsed.toString());
        
        // Verify payment
        const paymentInfo = await payment.getPayment(ORDER_ID);
        console.log("\nPayment verified:");
        console.log("Order ID:", paymentInfo.orderId.toString());
        console.log("Amount:", ethers.utils.formatUnits(paymentInfo.amount, 6), "USDT");
        console.log("Timestamp:", new Date(paymentInfo.timestamp.toNumber() * 1000).toISOString());
        
    } catch (error) {
        console.error("❌ Payment failed:", error.message);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
