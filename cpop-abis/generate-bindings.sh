#!/bin/bash

# CPOP Contract ABI to Go bindings generator
# Usage: ./generate-bindings.sh

echo "Generating Go bindings for CPOP contracts..."

# Generate bindings for each contract directly in current directory
abigen --abi CPOPToken.abi.json --pkg cpop --type CPOPToken --out "cpop_token.go"
abigen --abi AAccount.abi.json --pkg cpop --type AAccount --out "a_account.go"
abigen --abi AccountManager.abi.json --pkg cpop --type AccountManager --out "account_manager.go"
abigen --abi MasterAggregator.abi.json --pkg cpop --type MasterAggregator --out "master_aggregator.go"
abigen --abi GasPaymaster.abi.json --pkg cpop --type GasPaymaster --out "gas_paymaster.go"
abigen --abi GasPriceOracle.abi.json --pkg cpop --type GasPriceOracle --out "gas_price_oracle.go"
abigen --abi SessionKeyManager.abi.json --pkg cpop --type SessionKeyManager --out "session_key_manager.go"

echo "Go bindings generated successfully!"
echo "Import them in your Go code with: import \"github.com/HzBay/account-abstraction/cpop-abis\""