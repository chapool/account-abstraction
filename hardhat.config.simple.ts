import '@nomicfoundation/hardhat-ethers'
import '@openzeppelin/hardhat-upgrades'
import '@nomicfoundation/hardhat-verify'
import { HardhatUserConfig } from 'hardhat/config'
import * as dotenv from 'dotenv'

// Load environment variables
dotenv.config({ path: '.env.sepolia' })

const config: HardhatUserConfig = {
  solidity: {
    compilers: [{
      version: '0.8.28',
      settings: {
        evmVersion: 'cancun',
        viaIR: true,
        optimizer: { enabled: true, runs: 1000000 }
      }
    }],
    overrides: {
      'contracts/core/EntryPoint.sol': {
        version: '0.8.28',
        settings: {
          evmVersion: 'cancun',
          optimizer: { enabled: true, runs: 1000000 },
          viaIR: true
        }
      }
    }
  },
  networks: {
    sepoliaCustom: {
      url: process.env.ETH_RPC_URL!,
      chainId: 11155111,
      gasPrice: 20000000000,
      accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : [],
      timeout: 60000
    }
  },
  mocha: {
    timeout: 10000
  }
}

export default config