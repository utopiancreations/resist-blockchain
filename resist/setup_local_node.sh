#!/bin/bash
# This script automates the setup of a local node for the Resist blockchain.
set -e

# --- Configuration ---
BINARY="/Users/joshuamiller/go/bin/resistd"
KEY_NAME="validator"
CHAIN_ID="resist-local"
KEYRING_BACKEND="test"
# This password is used for non-interactive key creation and signing.
# In a real-world scenario, this should be handled more securely.
PASSWORD="password"

# --- File Paths ---
KEY_FILE="validator_keys.json"
MNEMONIC_FILE="validator_mnemonic.txt"
NODE_HOME="$HOME/.resist"

# --- Script ---

echo "--- Removing old data from $NODE_HOME ---"
rm -rf $NODE_HOME

echo "--- Initializing node ---"
$BINARY init local-resist-node --chain-id $CHAIN_ID

echo "--- Creating validator key ---"
# The `keys add` command with --output json prints key info to stdout
# and the mnemonic to stderr. We capture both.
printf "%s\n%s\n" "$PASSWORD" "$PASSWORD" | $BINARY keys add $KEY_NAME --keyring-backend $KEYRING_BACKEND --output json > "$KEY_FILE" 2> "$MNEMONIC_FILE"

# Extract the address from the key file to use in the genesis command
VALIDATOR_ADDR=$($BINARY keys show $KEY_NAME -a --keyring-backend $KEYRING_BACKEND)

echo "Validator key information saved to $KEY_FILE"
echo "Validator mnemonic phrase saved to $MNEMONIC_FILE"
echo "IMPORTANT: Secure the '$MNEMONIC_FILE' as it provides full control over the account."

echo "--- Adding genesis account ---"
$BINARY genesis add-genesis-account $VALIDATOR_ADDR 1000000000stake --keyring-backend $KEYRING_BACKEND

echo "--- Creating genesis transaction ---"
$BINARY genesis gentx $KEY_NAME 1000000stake --chain-id $CHAIN_ID --keyring-backend $KEYRING_BACKEND

echo "--- Collecting genesis transactions ---"
$BINARY genesis collect-gentxs

echo ""
echo "--- Local node setup complete! ---"
echo "You can now start the node with the following command:"
echo ""
echo "$BINARY start"
echo ""