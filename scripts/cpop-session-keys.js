/**
 * CPOP Session Keys Management Tool
 * Provides utilities for creating, managing, and using session keys
 */

const { ethers } = require("ethers");
const { keccak256, defaultAbiCoder } = ethers.utils;

class CPOPSessionKeyManager {
    constructor(provider, accountAddress, sessionKeyManagerAddress) {
        this.provider = provider;
        this.accountAddress = accountAddress;
        this.sessionKeyManagerAddress = sessionKeyManagerAddress;
        
        // Initialize contract instances
        this.accountContract = new ethers.Contract(
            accountAddress,
            [
                "function addSessionKey(address sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions) external",
                "function revokeSessionKey(address sessionKey) external",
                "function getSessionKeyInfo(address sessionKey) external view returns (bool isValid, uint48 validAfter, uint48 validUntil, bytes32 permissions)",
                "function canSessionKeyExecute(address sessionKey, address target, bytes4 selector) external view returns (bool canExecute)"
            ],
            provider
        );
        
        if (sessionKeyManagerAddress) {
            this.managerContract = new ethers.Contract(
                sessionKeyManagerAddress,
                [
                    "function addSessionKeyWithTemplate(address masterSigner, address sessionKey, string calldata templateName, uint48 customDuration) external",
                    "function batchAddSessionKeys((address account, address sessionKey, uint48 validAfter, uint48 validUntil, bytes32 permissions)[] calldata operations) external",
                    "function batchRevokeSessionKey(address sessionKey) external",
                    "function getRegisteredAccounts(address masterSigner) external view returns (address[] memory accounts)",
                    "function getTemplate(string calldata templateName) external view returns ((string name, uint48 defaultDuration, bytes32 permissions, bool isActive) memory template)"
                ],
                provider
            );
        }
    }

    /**
     * Create a new session key
     * @param {Object} options Session key options
     * @returns {Object} Session key data
     */
    async createSessionKey(options = {}) {
        const {
            validAfter = 0,
            duration = 24 * 60 * 60, // 24 hours default
            permissions = "0x0000000000000000000000000000000000000000000000000000000000000000", // Allow all
            generateNewKey = true
        } = options;

        let sessionKey;
        let sessionKeyPrivateKey;

        if (generateNewKey) {
            const wallet = ethers.Wallet.createRandom();
            sessionKey = wallet.address;
            sessionKeyPrivateKey = wallet.privateKey;
        } else {
            sessionKey = options.sessionKey;
            sessionKeyPrivateKey = options.sessionKeyPrivateKey;
        }

        const validAfterTime = validAfter || Math.floor(Date.now() / 1000);
        const validUntilTime = validAfterTime + duration;

        return {
            sessionKey,
            sessionKeyPrivateKey,
            validAfter: validAfterTime,
            validUntil: validUntilTime,
            permissions,
            duration
        };
    }

    /**
     * Add session key to account
     * @param {string} sessionKey Session key address
     * @param {Object} options Session key options
     * @param {ethers.Signer} signer Account owner signer
     */
    async addSessionKey(sessionKey, options, signer) {
        const {
            validAfter = 0,
            validUntil,
            permissions = "0x0000000000000000000000000000000000000000000000000000000000000000"
        } = options;

        const contract = this.accountContract.connect(signer);
        
        console.log(`Adding session key ${sessionKey} to account ${this.accountAddress}`);
        console.log(`Valid from: ${new Date(validAfter * 1000)}`);
        console.log(`Valid until: ${new Date(validUntil * 1000)}`);

        const tx = await contract.addSessionKey(
            sessionKey,
            validAfter,
            validUntil,
            permissions
        );

        const receipt = await tx.wait();
        console.log(`Session key added. Tx hash: ${receipt.transactionHash}`);
        
        return receipt;
    }

    /**
     * Add session key using template
     * @param {string} sessionKey Session key address
     * @param {string} templateName Template name
     * @param {number} customDuration Custom duration (optional)
     * @param {ethers.Signer} masterSigner Master signer
     */
    async addSessionKeyWithTemplate(sessionKey, templateName, customDuration, masterSigner) {
        if (!this.managerContract) {
            throw new Error("SessionKeyManager contract not initialized");
        }

        const contract = this.managerContract.connect(masterSigner);
        
        console.log(`Adding session key ${sessionKey} using template ${templateName}`);

        const tx = await contract.addSessionKeyWithTemplate(
            await masterSigner.getAddress(),
            sessionKey,
            templateName,
            customDuration || 0
        );

        const receipt = await tx.wait();
        console.log(`Session key added with template. Tx hash: ${receipt.transactionHash}`);
        
        return receipt;
    }

    /**
     * Batch add session keys to multiple accounts
     * @param {Array} operations Array of batch operations
     * @param {ethers.Signer} masterSigner Master signer
     */
    async batchAddSessionKeys(operations, masterSigner) {
        if (!this.managerContract) {
            throw new Error("SessionKeyManager contract not initialized");
        }

        const contract = this.managerContract.connect(masterSigner);
        
        console.log(`Batch adding ${operations.length} session keys`);

        const tx = await contract.batchAddSessionKeys(operations);
        const receipt = await tx.wait();
        
        console.log(`Batch session keys added. Tx hash: ${receipt.transactionHash}`);
        
        return receipt;
    }

    /**
     * Revoke session key
     * @param {string} sessionKey Session key address
     * @param {ethers.Signer} signer Account owner signer
     */
    async revokeSessionKey(sessionKey, signer) {
        const contract = this.accountContract.connect(signer);
        
        console.log(`Revoking session key ${sessionKey} from account ${this.accountAddress}`);

        const tx = await contract.revokeSessionKey(sessionKey);
        const receipt = await tx.wait();
        
        console.log(`Session key revoked. Tx hash: ${receipt.transactionHash}`);
        
        return receipt;
    }

    /**
     * Get session key information
     * @param {string} sessionKey Session key address
     */
    async getSessionKeyInfo(sessionKey) {
        const [isValid, validAfter, validUntil, permissions] = await this.accountContract.getSessionKeyInfo(sessionKey);
        
        return {
            sessionKey,
            isValid,
            validAfter: validAfter.toNumber(),
            validUntil: validUntil.toNumber(),
            permissions,
            validAfterDate: new Date(validAfter.toNumber() * 1000),
            validUntilDate: new Date(validUntil.toNumber() * 1000),
            isCurrentlyValid: isValid && 
                Date.now() >= validAfter.toNumber() * 1000 && 
                Date.now() <= validUntil.toNumber() * 1000
        };
    }

    /**
     * Check if session key can execute specific operation
     * @param {string} sessionKey Session key address
     * @param {string} target Target contract address
     * @param {string} selector Function selector
     */
    async canSessionKeyExecute(sessionKey, target, selector) {
        return await this.accountContract.canSessionKeyExecute(sessionKey, target, selector);
    }

    /**
     * Create permission encoding for specific targets and selectors
     * @param {Array} allowedTargets Array of allowed contract addresses
     * @param {Array} allowedSelectors Array of allowed function selectors
     */
    createPermissions(allowedTargets = [], allowedSelectors = []) {
        if (allowedTargets.length === 0 && allowedSelectors.length === 0) {
            // Allow all operations
            return "0x0000000000000000000000000000000000000000000000000000000000000000";
        }

        // Simplified permission encoding
        // In production, would implement more sophisticated permission system
        let permissions = "0x";
        
        if (allowedTargets.length > 0) {
            const targetHash = keccak256(defaultAbiCoder.encode(["address[]"], [allowedTargets]));
            permissions += targetHash.slice(2, 34); // 16 bytes
        } else {
            permissions += "00000000000000000000000000000000"; // 16 bytes of zeros
        }
        
        if (allowedSelectors.length > 0) {
            const selectorHash = keccak256(defaultAbiCoder.encode(["bytes4[]"], [allowedSelectors]));
            permissions += selectorHash.slice(2, 10); // 4 bytes
        } else {
            permissions += "00000000"; // 4 bytes of zeros
        }
        
        permissions += "000000000000000000000000"; // 12 bytes padding
        
        return permissions;
    }

    /**
     * Create session key for DApp interaction
     * @param {string} dappAddress DApp contract address
     * @param {number} duration Session duration in seconds
     */
    async createDAppSessionKey(dappAddress, duration = 24 * 60 * 60) {
        const sessionKeyData = await this.createSessionKey({ duration });
        
        // Create permissions for DApp interaction
        const permissions = this.createPermissions([dappAddress], []);
        
        return {
            ...sessionKeyData,
            permissions,
            type: "DAPP_INTERACTION",
            targetContract: dappAddress
        };
    }

    /**
     * Create session key for trading operations
     * @param {Array} exchangeAddresses Array of exchange contract addresses
     * @param {number} duration Session duration in seconds
     */
    async createTradingSessionKey(exchangeAddresses, duration = 60 * 60) {
        const sessionKeyData = await this.createSessionKey({ duration });
        
        // Common trading function selectors
        const tradingSelectors = [
            "0xa9059cbb", // transfer(address,uint256)
            "0x095ea7b3", // approve(address,uint256)
            "0x7ff36ab5", // swapExactTokensForETH(uint256,uint256,address[],address,uint256)
            "0x38ed1739"  // swapExactTokensForTokens(uint256,uint256,address[],address,uint256)
        ];
        
        const permissions = this.createPermissions(exchangeAddresses, tradingSelectors);
        
        return {
            ...sessionKeyData,
            permissions,
            type: "TRADING",
            exchangeAddresses,
            allowedSelectors: tradingSelectors
        };
    }

    /**
     * Create session key for game interactions
     * @param {string} gameAddress Game contract address
     * @param {number} duration Session duration in seconds
     */
    async createGameSessionKey(gameAddress, duration = 7 * 24 * 60 * 60) {
        const sessionKeyData = await this.createSessionKey({ duration });
        
        // Common game function selectors
        const gameSelectors = [
            "0xa9059cbb", // transfer(address,uint256)
            "0x095ea7b3", // approve(address,uint256)
            "0x40c10f19", // mint(address,uint256)
            "0x42842e0e"  // safeTransferFrom(address,address,uint256)
        ];
        
        const permissions = this.createPermissions([gameAddress], gameSelectors);
        
        return {
            ...sessionKeyData,
            permissions,
            type: "GAME_INTERACTION",
            gameContract: gameAddress,
            allowedSelectors: gameSelectors
        };
    }

    /**
     * Sign transaction with session key
     * @param {Object} userOp User operation
     * @param {string} sessionKeyPrivateKey Session key private key
     */
    async signWithSessionKey(userOp, sessionKeyPrivateKey) {
        const sessionKeyWallet = new ethers.Wallet(sessionKeyPrivateKey, this.provider);
        
        // Create user operation hash
        const userOpHash = this.getUserOperationHash(userOp);
        
        // Sign with session key
        const signature = await sessionKeyWallet.signMessage(ethers.utils.arrayify(userOpHash));
        
        return signature;
    }

    /**
     * Get user operation hash for signing
     * @param {Object} userOp User operation
     */
    getUserOperationHash(userOp) {
        // Simplified hash calculation
        // In production, would use exact EntryPoint hash calculation
        const encoded = defaultAbiCoder.encode(
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
        );
        
        return keccak256(encoded);
    }

    /**
     * Get template information
     * @param {string} templateName Template name
     */
    async getTemplate(templateName) {
        if (!this.managerContract) {
            throw new Error("SessionKeyManager contract not initialized");
        }

        const template = await this.managerContract.getTemplate(templateName);
        
        return {
            name: template.name,
            defaultDuration: template.defaultDuration.toNumber(),
            permissions: template.permissions,
            isActive: template.isActive,
            durationHours: Math.floor(template.defaultDuration.toNumber() / 3600),
            durationDays: Math.floor(template.defaultDuration.toNumber() / (24 * 3600))
        };
    }

    /**
     * List all session keys for account (helper function)
     * Note: In production, would maintain an index or use events
     */
    async listSessionKeys(sessionKeys) {
        const results = [];
        
        for (const sessionKey of sessionKeys) {
            try {
                const info = await this.getSessionKeyInfo(sessionKey);
                results.push(info);
            } catch (error) {
                console.warn(`Failed to get info for session key ${sessionKey}:`, error.message);
            }
        }
        
        return results;
    }
}

module.exports = { CPOPSessionKeyManager };

// Example usage
if (require.main === module) {
    async function example() {
        // Setup
        const provider = new ethers.providers.JsonRpcProvider("http://localhost:8545");
        const accountAddress = "0x1234567890123456789012345678901234567890";
        const managerAddress = "0x0987654321098765432109876543210987654321";
        
        const sessionManager = new CPOPSessionKeyManager(provider, accountAddress, managerAddress);
        
        // Create a DApp session key
        const dappSessionKey = await sessionManager.createDAppSessionKey(
            "0xdAppContract1234567890123456789012345678901234567890",
            24 * 60 * 60 // 24 hours
        );
        
        console.log("Created DApp session key:", dappSessionKey);
        
        // Create a trading session key
        const tradingSessionKey = await sessionManager.createTradingSessionKey(
            [
                "0xUniswapRouter12345678901234567890123456789012345678",
                "0xSushiswapRouter123456789012345678901234567890123456"
            ],
            60 * 60 // 1 hour
        );
        
        console.log("Created trading session key:", tradingSessionKey);
        
        // Create a game session key
        const gameSessionKey = await sessionManager.createGameSessionKey(
            "0xGameContract1234567890123456789012345678901234567890",
            7 * 24 * 60 * 60 // 7 days
        );
        
        console.log("Created game session key:", gameSessionKey);
    }
    
    example().catch(console.error);
}