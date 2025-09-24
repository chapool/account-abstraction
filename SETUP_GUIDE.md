# CPNFT Deployment Setup Guide

## Issue Resolution: "Invalid JSON-RPC response received: invalid project id"

This error occurs because the environment variables for network configuration are missing or incorrect.

## Quick Fix Options

### Option 1: Create Environment File (Recommended)

1. **Create `.env.sepolia` file** in the project root:
```bash
touch .env.sepolia
```

2. **Add the following content** to `.env.sepolia`:
```env
# Infura Project ID (get from https://infura.io/)
INFURA_ID=your_infura_project_id_here

# Private key of the wallet that will deploy the contract
PRIVATE_KEY=your_private_key_here

# Etherscan API Key for contract verification (optional)
ETHERSCAN_API_KEY=your_etherscan_api_key_here

# Alternative RPC URL
ETH_RPC_URL=https://sepolia.infura.io/v3/your_infura_project_id_here
```

3. **Replace the placeholder values**:
   - Get Infura Project ID from [https://infura.io/](https://infura.io/)
   - Use your wallet's private key (keep it secure!)
   - Get Etherscan API key from [https://etherscan.io/apis](https://etherscan.io/apis)

### Option 2: Use Alternative RPC Provider

If you don't want to use Infura, you can use other RPC providers:

**Alchemy** (Alternative to Infura):
```env
ETH_RPC_URL=https://eth-sepolia.g.alchemy.com/v2/your_alchemy_api_key
```

**QuickNode**:
```env
ETH_RPC_URL=https://your_quicknode_endpoint
```

**Public RPC** (Less reliable):
```env
ETH_RPC_URL=https://rpc.sepolia.org
```

### Option 3: Deploy to Local Network First

Test the deployment locally first:
```bash
# Start local network
npx hardhat node

# In another terminal, deploy to local network
npx hardhat run scripts/deploy-cpnft-with-config.ts --network localhost
```

## Deployment Commands

### After setting up environment:

**Using the new config script:**
```bash
npx hardhat run scripts/deploy-cpnft-with-config.ts --network sepolia
```

**Using hardhat deploy:**
```bash
npx hardhat deploy --network sepolia --tags CPNFT
```

**Using simple script:**
```bash
npx hardhat run scripts/deploy-cpnft-simple.ts --network sepolia
```

## Getting Sepolia ETH

You need Sepolia ETH for gas fees. Get it from:
- [Sepolia Faucet](https://sepoliafaucet.com/)
- [Alchemy Faucet](https://sepoliafaucet.com/)
- [Chainlink Faucet](https://faucets.chain.link/sepolia)

## Network Configuration Check

Verify your network configuration in `hardhat.config.ts`:

```typescript
networks: {
  sepolia: {
    url: process.env.ETH_RPC_URL || `https://sepolia.infura.io/v3/${process.env.INFURA_ID}`,
    accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : [],
    chainId: 11155111
  }
}
```

## Troubleshooting Steps

1. **Check environment variables are loaded:**
   ```bash
   node -e "require('dotenv').config({path: '.env.sepolia'}); console.log('INFURA_ID:', process.env.INFURA_ID ? 'Set' : 'Not set')"
   ```

2. **Test network connection:**
   ```bash
   npx hardhat console --network sepolia
   # Then in console: ethers.provider.getNetwork()
   ```

3. **Check wallet balance:**
   ```bash
   npx hardhat run --network sepolia -e "console.log(ethers.utils.formatEther(await ethers.provider.getBalance(await ethers.getSigners().then(s => s[0].getAddress()))))"
   ```

## Security Notes

- ⚠️ **Never commit your `.env.sepolia` file to git**
- ⚠️ **Keep your private key secure**
- ⚠️ **Use testnet ETH only for testing**

## Alternative: Direct RPC Deployment

If you continue having issues, you can deploy using a direct RPC call by modifying the hardhat config temporarily:

```typescript
// In hardhat.config.ts, add:
sepoliaDirect: {
  url: 'https://sepolia.infura.io/v3/YOUR_PROJECT_ID',
  accounts: ['YOUR_PRIVATE_KEY'],
  chainId: 11155111
}
```

Then deploy with:
```bash
npx hardhat run scripts/deploy-cpnft-with-config.ts --network sepoliaDirect
```
