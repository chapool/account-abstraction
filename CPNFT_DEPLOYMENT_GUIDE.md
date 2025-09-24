# CPNFT Contract Deployment Guide

This guide explains how to deploy the CPNFT contract to Sepolia testnet.

## Prerequisites

1. **Environment Setup**
   - Node.js and npm/yarn installed
   - Hardhat configured with Sepolia network
   - Sepolia ETH for gas fees (get from faucet: https://sepoliafaucet.com/)

2. **Environment Variables**
   Create a `.env.sepolia` file with:
   ```
   INFURA_ID=your_infura_project_id
   PRIVATE_KEY=your_wallet_private_key
   ETHERSCAN_API_KEY=your_etherscan_api_key (optional, for verification)
   ```

## Deployment Methods

### Method 1: Using Hardhat Deploy Plugin (Recommended)

This method uses the hardhat-deploy plugin and automatically saves deployment information.

```bash
# Compile contracts first
npx hardhat compile

# Deploy to Sepolia
npx hardhat deploy --network sepolia --tags CPNFT
```

### Method 2: Using Standalone Script

This method runs a simple deployment script directly.

```bash
# Compile contracts first
npx hardhat compile

# Deploy to Sepolia
npx hardhat run scripts/deploy-cpnft-simple.ts --network sepolia
```

### Method 3: Using Hardhat Deploy with Specific Script

```bash
# Deploy only the CPNFT contract
npx hardhat deploy --network sepolia --tags CPNFT
```

## Contract Details

- **Name**: Chapool NFT
- **Symbol**: CPNFT
- **Base URI**: http://chapool.net:8080/api/v1/meta/
- **Type**: Upgradeable (UUPS Proxy)
- **Network**: Sepolia Testnet

## Post-Deployment Steps

1. **Verify Contract** (Optional)
   ```bash
   npx hardhat verify --network sepolia <CONTRACT_ADDRESS>
   ```

2. **Test Contract Functions**
   - Check contract name and symbol
   - Test minting functionality (owner only)
   - Test level setting functionality
   - Test staking status functions

3. **Set Up Staking Contract** (If needed)
   ```solidity
   // Call setStakingContract with your staking contract address
   cpnft.setStakingContract(STAKING_CONTRACT_ADDRESS);
   ```

## Deployment Information

After successful deployment, you'll get:
- **Proxy Address**: The main contract address to use
- **Implementation Address**: The logic contract address
- **Admin Address**: The proxy admin address
- **Transaction Hash**: The deployment transaction hash

## Troubleshooting

### Common Issues

1. **Insufficient Gas**
   - Ensure you have enough Sepolia ETH
   - Check gas price settings in hardhat.config.ts

2. **Network Connection Issues**
   - Verify your Infura project ID
   - Check network configuration

3. **Compilation Errors**
   - Run `npx hardhat clean` and `npx hardhat compile`
   - Check Solidity version compatibility

4. **Verification Fails**
   - Wait a few minutes after deployment
   - Try manual verification with the contract address

### Getting Help

- Check Hardhat documentation: https://hardhat.org/docs
- Verify on Etherscan: https://sepolia.etherscan.io/
- Check deployment logs for detailed error messages

## Contract Functions

### Owner Functions
- `mint(address to, NFTLevel level)`: Mint a single NFT
- `batchMint(address[] to, NFTLevel[] levels)`: Mint multiple NFTs
- `burn(uint256 tokenId)`: Burn a single NFT
- `batchBurn(uint256[] tokenIds)`: Burn multiple NFTs
- `setTokenLevel(uint256 tokenId, NFTLevel level)`: Set NFT level
- `setStakingContract(address _stakingContract)`: Set staking contract
- `setBaseURI(string baseURI)`: Set base URI for metadata

### Public Functions
- `getTokenLevel(uint256 tokenId)`: Get NFT level
- `isStaked(uint256 tokenId)`: Check if NFT is staked
- `tokenURI(uint256 tokenId)`: Get token metadata URI

### NFT Levels
- C (0), B (1), A (2), S (3), SS (4), SSS (5)

## Security Notes

- The contract uses UUPS upgradeable pattern
- Only the contract owner can upgrade
- Staked tokens cannot be transferred
- All batch operations are owner-only for security
