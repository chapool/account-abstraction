#!/bin/bash

# CPOP Contract ABI to Go bindings generator
# Usage: ./generate-bindings.sh

echo "Generating Go bindings for CPOP contracts..."

# Generate bindings for each contract directly in current directory
abigen --abi CPOPToken.abi.json --pkg cpop --type CPOPToken --out "cpop_token.go"
abigen --abi AAWallet.abi.json --pkg cpop --type AAWallet --out "aa_wallet.go"
abigen --abi WalletManager.abi.json --pkg cpop --type WalletManager --out "wallet_manager.go"
abigen --abi MasterAggregator.abi.json --pkg cpop --type MasterAggregator --out "master_aggregator.go"
abigen --abi GasPaymaster.abi.json --pkg cpop --type GasPaymaster --out "gas_paymaster.go"
abigen --abi GasPriceOracle.abi.json --pkg cpop --type GasPriceOracle --out "gas_price_oracle.go"
abigen --abi SessionKeyManager.abi.json --pkg cpop --type SessionKeyManager --out "session_key_manager.go"

echo "Go bindings generated successfully!"
echo "Import them in your Go code with: import \"github.com/HzBay/account-abstraction/cpop-abis\""