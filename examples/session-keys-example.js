/**
 * CPOP Session Keys Usage Examples
 * Demonstrates various session key scenarios and integrations
 */

const { ethers } = require("ethers");
const { CPOPSessionKeyManager } = require("../scripts/cpop-session-keys");

class SessionKeyExamples {
    constructor(provider, entryPointAddress, aggregatorAddress) {
        this.provider = provider;
        this.entryPointAddress = entryPointAddress;
        this.aggregatorAddress = aggregatorAddress;
    }

    /**
     * Example 1: DApp Integration with Session Keys
     * User grants temporary permission to a DApp for trading
     */
    async dappIntegrationExample() {
        console.log("\n=== DApp Integration Example ===");
        
        // Setup accounts
        const userWallet = ethers.Wallet.createRandom().connect(this.provider);
        const userAccountAddress = "0x" + "1".repeat(40); // Mock account address
        
        const sessionManager = new CPOPSessionKeyManager(
            this.provider,
            userAccountAddress
        );
        
        // 1. User creates session key for DApp interaction
        const dappAddress = "0x" + "2".repeat(40); // Mock DApp address
        const sessionKeyData = await sessionManager.createDAppSessionKey(
            dappAddress,
            2 * 60 * 60 // 2 hours
        );
        
        console.log("1. Created session key for DApp:");
        console.log(`   Session Key: ${sessionKeyData.sessionKey}`);
        console.log(`   Valid until: ${new Date(sessionKeyData.validUntil * 1000)}`);
        console.log(`   Permissions: ${sessionKeyData.permissions}`);
        
        // 2. User adds session key to their account
        // (In real scenario, this would be done with proper signer)
        console.log("\n2. Adding session key to account...");
        console.log(`   Account: ${userAccountAddress}`);
        console.log(`   Session Key: ${sessionKeyData.sessionKey}`);
        
        // 3. DApp can now use session key to execute operations
        const dappOperation = {
            sender: userAccountAddress,
            nonce: 1,
            callData: "0x", // DApp-specific call data
            accountGasLimits: "0x" + "1".repeat(64),
            preVerificationGas: 21000,
            gasFees: "0x" + "2".repeat(64),
            paymasterAndData: "0x"
        };
        
        // 4. Sign operation with session key
        const signature = await sessionManager.signWithSessionKey(
            dappOperation,
            sessionKeyData.sessionKeyPrivateKey
        );
        
        console.log("\n3. DApp signed operation with session key:");
        console.log(`   Signature: ${signature.slice(0, 20)}...`);
        
        return sessionKeyData;
    }

    /**
     * Example 2: Gaming Session Keys
     * Player grants game contract permission for in-game transactions
     */
    async gamingSessionExample() {
        console.log("\n=== Gaming Session Example ===");
        
        const playerWallet = ethers.Wallet.createRandom().connect(this.provider);
        const playerAccountAddress = "0x" + "3".repeat(40);
        const gameAddress = "0x" + "4".repeat(40);
        
        const sessionManager = new CPOPSessionKeyManager(
            this.provider,
            playerAccountAddress
        );
        
        // 1. Create game session key (7 days for long gaming sessions)
        const gameSessionKey = await sessionManager.createGameSessionKey(
            gameAddress,
            7 * 24 * 60 * 60 // 7 days
        );
        
        console.log("1. Created game session key:");
        console.log(`   Game Contract: ${gameAddress}`);
        console.log(`   Session Duration: 7 days`);
        console.log(`   Allowed Operations: NFT transfers, minting, trading`);
        
        // 2. Simulate game operations
        const gameOperations = [
            {
                type: "mint_item",
                target: gameAddress,
                data: "0x40c10f19" + "1".repeat(56) // mint(address,uint256)
            },
            {
                type: "transfer_nft",
                target: gameAddress,
                data: "0x42842e0e" + "2".repeat(56) // safeTransferFrom
            },
            {
                type: "approve_marketplace",
                target: gameAddress,
                data: "0x095ea7b3" + "3".repeat(56) // approve
            }
        ];
        
        console.log("\n2. Game operations that can be performed:");
        gameOperations.forEach((op, index) => {
            console.log(`   ${index + 1}. ${op.type} -> ${op.target}`);
        });
        
        return gameSessionKey;
    }

    /**
     * Example 3: Trading Bot with Session Keys
     * Automated trading with limited permissions and time constraints
     */
    async tradingBotExample() {
        console.log("\n=== Trading Bot Example ===");
        
        const traderWallet = ethers.Wallet.createRandom().connect(this.provider);
        const traderAccountAddress = "0x" + "5".repeat(40);
        
        const exchangeAddresses = [
            "0x" + "6".repeat(40), // Uniswap V3
            "0x" + "7".repeat(40), // SushiSwap
            "0x" + "8".repeat(40)  // Curve
        ];
        
        const sessionManager = new CPOPSessionKeyManager(
            this.provider,
            traderAccountAddress
        );
        
        // 1. Create trading session key (1 hour for high-frequency trading)
        const tradingSessionKey = await sessionManager.createTradingSessionKey(
            exchangeAddresses,
            60 * 60 // 1 hour
        );
        
        console.log("1. Created trading bot session key:");
        console.log(`   Trading Duration: 1 hour`);
        console.log(`   Allowed Exchanges: ${exchangeAddresses.length}`);
        console.log(`   Operations: swap, approve, transfer`);
        
        // 2. Simulate trading operations
        const tradingStrategies = [
            {
                name: "Arbitrage",
                description: "Buy low on Uniswap, sell high on SushiSwap",
                frequency: "Every 30 seconds",
                riskLevel: "Medium"
            },
            {
                name: "DCA (Dollar Cost Average)",
                description: "Regular purchases of ETH/USDC",
                frequency: "Every 5 minutes",
                riskLevel: "Low"
            },
            {
                name: "Momentum Trading",
                description: "Follow price trends across multiple DEXs",
                frequency: "Every 1 minute",
                riskLevel: "High"
            }
        ];
        
        console.log("\n2. Available trading strategies:");
        tradingStrategies.forEach((strategy, index) => {
            console.log(`   ${index + 1}. ${strategy.name}:`);
            console.log(`      Description: ${strategy.description}`);
            console.log(`      Frequency: ${strategy.frequency}`);
            console.log(`      Risk: ${strategy.riskLevel}`);
        });
        
        return tradingSessionKey;
    }

    /**
     * Example 4: Master Signer with Multiple Session Keys
     * Web2 company managing multiple wallets with different session keys
     */
    async masterSignerMultiWalletExample() {
        console.log("\n=== Master Signer Multi-Wallet Example ===");
        
        const masterSigner = ethers.Wallet.createRandom().connect(this.provider);
        const sessionKeyManagerAddress = "0x" + "9".repeat(40);
        
        // Multiple wallets controlled by master signer
        const wallets = [
            { address: "0x" + "A".repeat(40), purpose: "Trading Wallet" },
            { address: "0x" + "B".repeat(40), purpose: "Gaming Wallet" },
            { address: "0x" + "C".repeat(40), purpose: "DeFi Wallet" },
            { address: "0x" + "D".repeat(40), purpose: "NFT Wallet" }
        ];
        
        console.log("1. Master signer controls multiple wallets:");
        wallets.forEach((wallet, index) => {
            console.log(`   ${index + 1}. ${wallet.purpose}: ${wallet.address}`);
        });
        
        // Create session managers for each wallet
        const sessionManagers = wallets.map(wallet => 
            new CPOPSessionKeyManager(
                this.provider,
                wallet.address,
                sessionKeyManagerAddress
            )
        );
        
        // 2. Create different types of session keys for different purposes
        const sessionKeyConfigs = [
            {
                type: "trading",
                template: "TRADING",
                duration: 60 * 60, // 1 hour
                walletIndex: 0
            },
            {
                type: "gaming",
                template: "GAME_INTERACTION",
                duration: 7 * 24 * 60 * 60, // 7 days
                walletIndex: 1
            },
            {
                type: "defi",
                template: "DAPP_INTERACTION",
                duration: 24 * 60 * 60, // 24 hours
                walletIndex: 2
            },
            {
                type: "nft",
                template: "DAPP_INTERACTION",
                duration: 12 * 60 * 60, // 12 hours
                walletIndex: 3
            }
        ];
        
        console.log("\n2. Session key configurations:");
        sessionKeyConfigs.forEach((config, index) => {
            console.log(`   ${index + 1}. ${config.type.toUpperCase()} session key:`);
            console.log(`      Template: ${config.template}`);
            console.log(`      Duration: ${config.duration / 3600} hours`);
            console.log(`      Target Wallet: ${wallets[config.walletIndex].purpose}`);
        });
        
        // 3. Batch operations with session keys
        const batchOperations = sessionKeyConfigs.map(config => {
            const sessionKey = ethers.Wallet.createRandom().address;
            return {
                account: wallets[config.walletIndex].address,
                sessionKey: sessionKey,
                validAfter: Math.floor(Date.now() / 1000),
                validUntil: Math.floor(Date.now() / 1000) + config.duration,
                permissions: "0x" + "0".repeat(64) // Allow all for demo
            };
        });
        
        console.log("\n3. Batch session key deployment:");
        console.log(`   Total operations: ${batchOperations.length}`);
        console.log(`   Master signer can deploy all at once`);
        
        return { masterSigner, wallets, batchOperations };
    }

    /**
     * Example 5: Session Key with Aggregated Signatures
     * Multiple operations signed with one session key across multiple wallets
     */
    async sessionKeyAggregationExample() {
        console.log("\n=== Session Key Aggregation Example ===");
        
        const sessionKey = ethers.Wallet.createRandom();
        const walletAddresses = [
            "0x" + "E".repeat(40),
            "0x" + "F".repeat(40),
            "0x" + "10".repeat(39)
        ];
        
        console.log("1. Session key aggregation setup:");
        console.log(`   Session Key: ${sessionKey.address}`);
        console.log(`   Target Wallets: ${walletAddresses.length}`);
        
        // Create user operations for multiple wallets
        const userOperations = walletAddresses.map((address, index) => ({
            sender: address,
            nonce: index + 1,
            callData: `0x${"1".repeat(8 + index * 2)}`, // Different call data
            accountGasLimits: "0x" + "1".repeat(64),
            preVerificationGas: 21000,
            gasFees: "0x" + "2".repeat(64),
            paymasterAndData: "0x"
        }));
        
        console.log("\n2. User operations to be aggregated:");
        userOperations.forEach((op, index) => {
            console.log(`   ${index + 1}. Wallet: ${op.sender}`);
            console.log(`      Nonce: ${op.nonce}`);
            console.log(`      Call Data: ${op.callData.slice(0, 20)}...`);
        });
        
        // Calculate gas savings from aggregation
        const individualGasCost = userOperations.length * 3000; // 3000 gas per signature
        const aggregatedGasCost = 5000 + (userOperations.length * 500); // Fixed + marginal cost
        const gasSavings = individualGasCost - aggregatedGasCost;
        
        console.log("\n3. Gas optimization with session key aggregation:");
        console.log(`   Individual signatures cost: ${individualGasCost} gas`);
        console.log(`   Aggregated signature cost: ${aggregatedGasCost} gas`);
        console.log(`   Gas savings: ${gasSavings} gas (${Math.round(gasSavings/individualGasCost*100)}%)`);
        
        return { sessionKey, userOperations, gasSavings };
    }

    /**
     * Example 6: Session Key Lifecycle Management
     * Complete lifecycle from creation to expiration
     */
    async sessionKeyLifecycleExample() {
        console.log("\n=== Session Key Lifecycle Example ===");
        
        const accountAddress = "0x" + "11".repeat(20);
        const sessionManager = new CPOPSessionKeyManager(this.provider, accountAddress);
        
        // 1. Create session key
        const sessionKeyData = await sessionManager.createSessionKey({
            duration: 60 * 60, // 1 hour
            permissions: "0x" + "0".repeat(64)
        });
        
        console.log("1. Session Key Created:");
        console.log(`   Address: ${sessionKeyData.sessionKey}`);
        console.log(`   Valid from: ${new Date(sessionKeyData.validAfter * 1000)}`);
        console.log(`   Valid until: ${new Date(sessionKeyData.validUntil * 1000)}`);
        
        // 2. Simulate usage over time
        const usageScenarios = [
            { time: 0, action: "Created", status: "Active" },
            { time: 30 * 60, action: "DApp Transaction", status: "Active" },
            { time: 45 * 60, action: "Token Swap", status: "Active" },
            { time: 60 * 60, action: "Attempting Transaction", status: "Expired" },
            { time: 65 * 60, action: "Manual Revocation", status: "Revoked" }
        ];
        
        console.log("\n2. Session Key Usage Timeline:");
        usageScenarios.forEach((scenario, index) => {
            const timeStr = `${Math.floor(scenario.time / 60)}:${String(scenario.time % 60).padStart(2, '0')}`;
            console.log(`   ${timeStr} - ${scenario.action} (${scenario.status})`);
        });
        
        // 3. Cleanup and security
        console.log("\n3. Security and Cleanup:");
        console.log("   - Session keys automatically expire");
        console.log("   - Manual revocation available anytime");
        console.log("   - Private keys should be securely deleted after use");
        console.log("   - Monitor for unauthorized usage");
        
        return sessionKeyData;
    }

    /**
     * Run all examples
     */
    async runAllExamples() {
        console.log("🚀 CPOP Session Keys Examples");
        console.log("=" * 50);
        
        try {
            await this.dappIntegrationExample();
            await this.gamingSessionExample();
            await this.tradingBotExample();
            await this.masterSignerMultiWalletExample();
            await this.sessionKeyAggregationExample();
            await this.sessionKeyLifecycleExample();
            
            console.log("\n✅ All examples completed successfully!");
            console.log("\n📋 Summary of Session Key Benefits:");
            console.log("   • Temporary permissions for enhanced security");
            console.log("   • Automated operations without exposing main keys");
            console.log("   • Fine-grained permission control");
            console.log("   • Gas optimization through aggregation");
            console.log("   • Perfect for Web2 user experience");
            
        } catch (error) {
            console.error("❌ Error running examples:", error);
        }
    }
}

// Export for use in other scripts
module.exports = { SessionKeyExamples };

// Run examples if called directly
if (require.main === module) {
    async function main() {
        const provider = new ethers.providers.JsonRpcProvider(
            process.env.RPC_URL || "http://localhost:8545"
        );
        
        const entryPointAddress = process.env.ENTRYPOINT_ADDRESS || "0x" + "EP".repeat(20);
        const aggregatorAddress = process.env.AGGREGATOR_ADDRESS || "0x" + "AG".repeat(20);
        
        const examples = new SessionKeyExamples(provider, entryPointAddress, aggregatorAddress);
        await examples.runAllExamples();
    }
    
    main().catch(console.error);
}