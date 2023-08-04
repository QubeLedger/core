#!/bin/bash

KEY="main"
CHAINID="qube-1"
MONIKER="localtestnet"
KEYRING="os"
LOGLEVEL="info"
TRACE="--trace"
# TRACE=""

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# remove existing daemon and client
rm -rf ~/.quadrate*

# make install

./qubed config keyring-backend $KEYRING
./qubed config chain-id $CHAINID

# if $KEY exists it should be deleted
./qubed keys add $KEY --keyring-backend $KEYRING

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
./qubed init $MONIKER --chain-id $CHAINID

cat $HOME/.quadrate/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="uqube"' > $HOME/.quadrate/config/tmp_genesis.json && mv $HOME/.quadrate/config/tmp_genesis.json $HOME/.quadrate/config/genesis.json
cat $HOME/.quadrate/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="uqube"' > $HOME/.quadrate/config/tmp_genesis.json && mv $HOME/.quadrate/config/tmp_genesis.json $HOME/.quadrate/config/genesis.json
cat $HOME/.quadrate/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uqube"' > $HOME/.quadrate/config/tmp_genesis.json && mv $HOME/.quadrate/config/tmp_genesis.json $HOME/.quadrate/config/genesis.json
cat $HOME/.quadrate/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="uqube"' > $HOME/.quadrate/config/tmp_genesis.json && mv $HOME/.quadrate/config/tmp_genesis.json $HOME/.quadrate/config/genesis.json

# Set gas limit in genesis
cat $HOME/.quadrate/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="20000000"' > $HOME/.quadrate/config/tmp_genesis.json && mv $HOME/.quadrate/config/tmp_genesis.json $HOME/.quadrate/config/genesis.json

# Allocate genesis accounts (cosmos formatted addresses)
./qubed add-genesis-account $KEY 100000000000000000000000000uqube --keyring-backend $KEYRING

# Sign genesis transaction
./qubed gentx $KEY 1000000000000000000000uqube --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
./qubed collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
./qubed validate-genesis

# disable produce empty block and enable prometheus metrics
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.quadrate/config/config.toml
    sed -i '' 's/prometheus = false/prometheus = true/' $HOME/.quadrate/config/config.toml
    sed -i '' 's/prometheus-retention-time = 0/prometheus-retention-time  = 1000000000000/g' $HOME/.quadrate/config/app.toml
    sed -i '' 's/enabled = false/enabled = true/g' $HOME/.quadrate/config/app.toml
else
    sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.quadrate/config/config.toml
    sed -i 's/prometheus = false/prometheus = true/' $HOME/.quadrate/config/config.toml
    sed -i 's/prometheus-retention-time  = "0"/prometheus-retention-time  = "1000000000000"/g' $HOME/.quadrate/config/app.toml
    sed -i 's/enabled = false/enabled = true/g' $HOME/.quadrate/config/app.toml
fi

if [[ $1 == "pending" ]]; then
    echo "pending mode is on, please wait for the first block committed."
    if [[ $OSTYPE == "darwin"* ]]; then
        sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.quadrate/config/config.toml
        sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.quadrate/config/config.toml
        sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.quadrate/config/config.toml
        sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.quadrate/config/config.toml
        sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.quadrate/config/config.toml
        sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.quadrate/config/config.toml
        sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.quadrate/config/config.toml
        sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.quadrate/config/config.toml
        sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.quadrate/config/config.toml
    else
        sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.quadrate/config/config.toml
        sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.quadrate/config/config.toml
        sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.quadrate/config/config.toml
        sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.quadrate/config/config.toml
        sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.quadrate/config/config.toml
        sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.quadrate/config/config.toml
        sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.quadrate/config/config.toml
        sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.quadrate/config/config.toml
        sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.quadrate/config/config.toml
    fi
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
./qubed start --pruning=nothing  --log_level $LOGLEVEL  --minimum-gas-prices=0.0001uqube
