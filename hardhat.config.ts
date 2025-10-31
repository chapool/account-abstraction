import '@nomiclabs/hardhat-waffle'
import '@nomiclabs/hardhat-ethers'
import '@typechain/hardhat'
import '@openzeppelin/hardhat-upgrades'
import { HardhatUserConfig, task } from 'hardhat/config'
import 'hardhat-deploy'

import 'solidity-coverage'

import * as fs from 'fs'
import * as dotenv from 'dotenv'

// Load environment variables
if (process.env.NODE_ENV === 'production' || process.env.SEPOLIA) {
  dotenv.config({ path: '.env.sepolia' })
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
  // Use Alchemy by default for Sepolia, fallback to Infura
  const defaultSepoliaUrl = 'https://eth-sepolia.g.alchemy.com/v2/CDpjLA10IDFcyjqnNHlVE'
  const url = process.env.SEPOLIA_RPC_URL || 
              (name === 'sepolia' ? defaultSepoliaUrl : `https://${name}.infura.io/v3/${process.env.INFURA_ID}`)
  return getNetwork1(url)
  // return getNetwork1(`wss://${name}.infura.io/ws/v3/${process.env.INFURA_ID}`)
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
      url: process.env.ETH_RPC_URL || process.env.SEPOLIA_RPC_URL || 'https://eth-sepolia.g.alchemy.com/v2/CDpjLA10IDFcyjqnNHlVE',
      chainId: 11155111,
      gasPrice: 20000000000,
      accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : [],
      timeout: 60000
    },
    bnbTestnet: {
      url: 'https://bnb-testnet.g.alchemy.com/v2/CDpjLA10IDFcyjqnNHlVE',
      chainId: 97,
      gasPrice: 20000000000,
      accounts: ['0x7afae63a6cbe2fad617ede265676aa02d61b91cd8ca817bffbcd0fbb67e1f18a'],
      timeout: 60000
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
