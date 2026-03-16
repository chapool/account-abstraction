import '@nomiclabs/hardhat-waffle'
import '@nomiclabs/hardhat-ethers'
import '@typechain/hardhat'
import '@openzeppelin/hardhat-upgrades'
import { HardhatUserConfig, task } from 'hardhat/config'
import 'hardhat-deploy'

import 'solidity-coverage'

import * as fs from 'fs'
import * as dotenv from 'dotenv'

// Load environment variables based on network
if (process.env.NODE_ENV === 'production' || process.env.SEPOLIA) {
  dotenv.config({ path: '.env.sepolia' })
} else if (process.env.HARDHAT_NETWORK === 'opbnbTestnet' || process.env.NETWORK === 'opbnbTestnet') {
  dotenv.config({ path: '.env.opBNBTestnet' })
} else if (process.env.HARDHAT_NETWORK === 'opbnb' || process.env.NETWORK === 'opbnb') {
  dotenv.config({ path: '.env.opBNB' })
} else if (process.env.HARDHAT_NETWORK === 'bnbTestnet' || process.env.NETWORK === 'bnbTestnet') {
  dotenv.config({ path: '.env.bnbTestnet' })
} else {
  dotenv.config()
}

const SALT = '0x0a59dbff790c23c976a548690c27297883cc66b4c67024f9117b0238995e35e9'
process.env.SALT = process.env.SALT ?? SALT

task('deploy', 'Deploy contracts')
  .addFlag('simpleAccountFactory', 'deploy sample factory (by default, enabled only on localhost)')

const mnemonicFileName = process.env.MNEMONIC_FILE!
let mnemonic = 'test '.repeat(11) + 'junk'
if (fs.existsSync(mnemonicFileName)) { mnemonic = fs.readFileSync(mnemonicFileName, 'ascii') }

function getNetwork1 (url: string): { url: string, accounts: { mnemonic: string } } {
  return {
    url,
    accounts: { mnemonic }
  }
}

function getNetwork (name: string): { url: string, accounts: { mnemonic: string } } {
  // RPC URLs should be configured via environment variables
  // For Sepolia: use SEPOLIA_RPC_URL or ETH_SEPOLIA_RPC_URL
  // For other networks: use INFURA_ID or network-specific RPC URL env vars
  const url = process.env.SEPOLIA_RPC_URL || 
              process.env.ETH_SEPOLIA_RPC_URL ||
              (name === 'sepolia' ? undefined : `https://${name}.infura.io/v3/${process.env.INFURA_ID}`)
  if (!url) {
    // Return a placeholder config instead of throwing error to allow other networks to work
    return { url: 'http://localhost:8545', accounts: { mnemonic } }
  }
  return getNetwork1(url)
}

function getNetworkWithPrivateKey (url: string, privateKey: string): { url: string, accounts: string[] } {
  return {
    url,
    accounts: [privateKey]
  }
}

const optimizedCompilerSettings = {
  version: '0.8.28',
  settings: {
    evmVersion: 'cancun',
    optimizer: { enabled: true, runs: 1000000 },
    viaIR: true
  }
}

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

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
      'contracts/core/EntryPoint.sol': optimizedCompilerSettings,
      'contracts/core/EntryPointSimulations.sol': optimizedCompilerSettings,
      'contracts/accounts/SimpleAccount.sol': optimizedCompilerSettings,
      'contracts/CPNFT/Staking.sol': optimizedCompilerSettings
    }
  },
  networks: {
    dev: { url: 'http://localhost:8545' },
    // github action starts localgeth service, for gas calculations
    localgeth: { url: 'http://localgeth:8545' },
    sepolia: getNetwork('sepolia'),
    proxy: getNetwork1('http://localhost:8545'),
    sepoliaCustom: {
      url: process.env.ETH_RPC_URL || process.env.SEPOLIA_RPC_URL || process.env.ETH_SEPOLIA_RPC_URL || 'http://localhost:8545',
      chainId: 11155111,
      gasPrice: process.env.SEPOLIA_GAS_PRICE ? parseInt(process.env.SEPOLIA_GAS_PRICE, 10) : 20000000000,
      accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : [],
      timeout: process.env.NETWORK_TIMEOUT ? parseInt(process.env.NETWORK_TIMEOUT, 10) : 60000
    },
    // BNB Chain Testnet (BSC Testnet) - Chain ID: 97
    // This is the primary testnet for BNB Smart Chain (BSC) development
    // RPC endpoint: Configure via BSC_TESTNET_RPC_URL or BNB_TESTNET_RPC_URL environment variable
    // Block explorer: https://testnet.bscscan.com
    bnbTestnet: {
      url: process.env.BSC_TESTNET_RPC_URL || process.env.BNB_TESTNET_RPC_URL || 'http://localhost:8545',
      chainId: 97,
      gasPrice: process.env.BSC_TESTNET_GAS_PRICE || process.env.BNB_TESTNET_GAS_PRICE 
        ? parseInt(process.env.BSC_TESTNET_GAS_PRICE || process.env.BNB_TESTNET_GAS_PRICE || '20000000000', 10) 
        : 20000000000,
      accounts: process.env.PRIVATE_KEY 
        ? [process.env.PRIVATE_KEY]
        : process.env.MNEMONIC 
          ? { 
              mnemonic: process.env.MNEMONIC,
              initialIndex: parseInt(process.env.ACCOUNT_INDEX || '0', 10),
              count: 1
            }
          : [],
      timeout: process.env.NETWORK_TIMEOUT ? parseInt(process.env.NETWORK_TIMEOUT, 10) : 60000
    },
    // opBNB Mainnet - Chain ID: 204
    // opBNB is an optimistic rollup on BNB Chain, providing enhanced scalability
    // This network is actively used for production deployments on BNB Chain ecosystem
    // RPC endpoint: Configure via OPBNB_RPC_URL or OPBNB_MAINNET_RPC_URL environment variable
    // Block explorer: https://opbnb.bscscan.com
    opbnb: {
      url: process.env.OPBNB_RPC_URL || process.env.OPBNB_MAINNET_RPC_URL || 'https://opbnb-mainnet.nodereal.io/v1/e59b1bd6360a4cdf9a39530a77e815ca',
      chainId: 204,
      gasPrice: process.env.OPBNB_GAS_PRICE ? parseInt(process.env.OPBNB_GAS_PRICE, 10) : 1000000000, // 1 gwei for opBNB
      accounts: process.env.PRIVATE_KEY 
        ? [process.env.PRIVATE_KEY]
        : process.env.MNEMONIC 
          ? { 
              mnemonic: process.env.MNEMONIC,
              initialIndex: parseInt(process.env.ACCOUNT_INDEX || '0', 10),
              count: 1
            }
          : [],
      timeout: process.env.NETWORK_TIMEOUT ? parseInt(process.env.NETWORK_TIMEOUT, 10) : 60000
    },
    // opBNB Testnet - Chain ID: 5611
    // Testnet for opBNB development and testing
    // RPC endpoint: Configure via OPBNB_TESTNET_RPC_URL environment variable
    opbnbTestnet: {
      url: process.env.OPBNB_TESTNET_RPC_URL || 'https://opbnb-testnet.nodereal.io/v1/e59b1bd6360a4cdf9a39530a77e815ca',
      chainId: 5611,
      gasPrice: process.env.OPBNB_TESTNET_GAS_PRICE || process.env.OPBNB_GAS_PRICE
        ? parseInt(process.env.OPBNB_TESTNET_GAS_PRICE || process.env.OPBNB_GAS_PRICE || '1000000000', 10)
        : 1000000000, // 1 gwei for opBNB
      accounts: process.env.PRIVATE_KEY 
        ? [process.env.PRIVATE_KEY]
        : process.env.MNEMONIC 
          ? { 
              mnemonic: process.env.MNEMONIC,
              initialIndex: parseInt(process.env.ACCOUNT_INDEX || '0', 10),
              count: 1
            }
          : [],
      timeout: process.env.NETWORK_TIMEOUT ? parseInt(process.env.NETWORK_TIMEOUT, 10) : 60000
    }
  },
  mocha: {
    timeout: 10000
  }
}

// coverage chokes on the "compilers" settings
if (process.env.COVERAGE != null) {
  // @ts-ignore
  config.solidity = config.solidity.compilers[0]
}

export default config

