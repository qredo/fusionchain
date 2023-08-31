# Fusion Chain Guide

This document leads through the features of Fusion Chain 
with the CLI interface and demonstrates the experience 
of interacting with a mocked MPC.

## Contents

* [Prerequisites](#prerequisites)
    * [Requirements](#requirements)
    * [Setup](#setup)
        * [Run the Chain](#run-the-chain)
        * [Faucet](#Faucet)
        * [Keyring](#keyring)
        * [Accounts](#accounts)
* [Start](#start)
    * [Basic Walkthrough](#basic-walkthrough)
    * [Broadcast](#broadcast)
        * [Pregenerated](#pregenerated)
        * [More Transactions](#more-tx)
    * [Manage Workspaces](#manage-workspaces)
    * [QAssets](#qassets)

## Prerequisites

### Requirements

- go 1.21+
- make
- docker (used to regenerate protobufs)

### Setup

#### Run the chain

Clone the repo:

```bash
git clone git@github.com:qredo/fusionchain.git
```

The `blockchain` directory contains the Cosmos SDK blockchain code. We can run
a local node with the following:

```bash
cd blockchain
./init.sh
```

This will run a local node with a couple of pre-funded accounts.

Resume the chain after stopping the daemon: `fusiond start`

#### Faucet

If you don't use the default key, you need to fund your account. Run the faucet:

```bash
cd blockchain
go run cmd/faucet/faucet.go
```

Fund your Fusion Chain wallet

```bash
curl localhost:8000 -XPOST -d'{"address":"qredo1ud49m3n00jkmtayj9w7k35zka3fqcl4lqp2j03"}'
```

#### Keyring

In a separate terminal, switch to the mocked keyring (`mokr`) and run it:

```bash
cd mokr && go run .
```

`mokr` will automatically monitor the chain and generate new keys and
signatures when requested.

#### Accounts

To interact with the chain you can use the `fusiond` CLI tool.

It's suggested to create an alias like this:

```bash
alias fchain="fusiond --node tcp://localhost:27657 --home ~/.fusiond/ --from shulgin --gas-prices 1000000000nQRDO"
```

that includes some common flags:

- `--node tcp://localhost:27657`, the Tendermint RPC endpoint
- `--home ~/.fusiond/`, the directory containing keys data
- `--from shulgin`, the account being used to sign transactions
- `--gas-prices 1000000000nQRDO`, the fee for transactions

## Start

### Basic Walkthrough

```bash
# create a new workspace
fchain tx identity new-workspace --yes

# check for newly created workspace
fchain q identity workspaces

# creata a new key of type `ecdsa`` for the workspace
fchain tx treasury new-key-request qredoworkspace14a2hpadpsy9h5m6us54 0 ecdsa --yes 

# wait for the MPC (keyring_id = 0) to pick up the request and generate a new key
# you can monitor all the requests with:
fchain q treasury key-requests 0 all

# and after the request in fulfilled you will find the public key:
fchain q treasury keys qredoworkspace14a2hpadpsy9h5m6us54

# let's use your new key to sign a payload
# payload must be a 32byte hash of arbitrary data to be singed
fchain tx treasury new-signature-request 1 '778f572f33acfab831365d52e563a0ddd2829ddd7060bec69719b7e41f6ef91c' --yes

# after a while you'll be able to retrieve the signature generated by the MPC
fchain q treasury signature-requests 0 all

# you can create a Ethereum-specific wallet using your public key
fchain tx treasury new-wallet-request sepolia 1 --yes

# retrieving the wallets will include the Ethereum address for that key:
fchain q treasury wallets
```

### Broadcast

For the wallet address generated in the previous step, make sure it is funded. For this example, we are using `0xC828Bf9126667972400E1ABE600BAAB877B1e674` as example. 

For the testnet Qredo runs its own watcher to broadcast the transaction. You can also run you own.

#### Pregenerated

This unsigned transaction only works for the first transaction to be sent from an address. The transaction looks as follows: 

```
nonce: 0
to: 0x993f45666B2A78434711D1a20D2A9733c07A5318
amount: 4000000000000000 WEI
gasLimit: 21000
gasPrice: 20000000000
data: 
```

Now request the transaction to be signed by the Fusion MPCs

```bash
# submit unsigned tx to be signed
fchain tx treasury new-sign-transaction-request 1 eb808504a817c80082520894993f45666b2a78434711d1a20d2a9733c07a5318870e35fa931a000080808080 --yes

# check the tx signature request has been fulfilled
fchain query treasury sign-transaction-requests ethereum
```

#### More Tx

To create more transactions for that wallet, generate the unsigned tx by yourself:

```bash
# switch to the relayer-eth
cd relayer-eth

# create a unsinged transaction, adjust the nonce
go run ./cmd/buildtx/ -nonce 0 -to 0x993f45666B2A78434711D1a20D2A9733c07A5318 -amount 4000000000000000

# increase the nonce for each new transaction you want to submit
```

### Manage Workspaces

You can add and remove owners or adjust policies inside a workspace. 

```bash
# add a new owner to the workspace
fchain tx identity add-workspace-owner qredoworkspace14a2hpadpsy9h5m6us54 qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5 --yes

# check the new owner has been added to the workspace
fchain q identity workspaces-by-owner qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5

# remove owner from the workspace
fchain tx identity remove-workspace-owner qredoworkspace14a2hpadpsy9h5m6us54 qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5 --yes

# check the new owner has been removed to the workspace
fchain q identity workspaces-by-owner qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5

# create a new workspace and append it to an existing one
fchain tx identity new-child-workspace qredoworkspace14a2hpadpsy9h5m6us54 --yes

# check the indicated workspace has a new workspace as child_workspace
fchain q identity workspaces
```

### QAssets

As of now, QAssets are just in a demo state. Anyone can freely mint any QAsset they like. This will change once we are onboarding verifiable oracles. 

```bash
# Mint an ETH-QAssets to a workspace
fchain tx qassets mint 1 qredoworkspace14a2hpadpsy9h5m6us54 false "" "" 1000000 --yes

# Check the workspace's balance
fchain q bank balances qredoworkspace14a2hpadpsy9h5m6us54

# Send QAssets
fchain tx qassets send qredoworkspace14a2hpadpsy9h5m6us54 qredoworkspace10j06zdk5gyl6vrss5d5 qETH-SEPOLIA 200000 --yes

# Check both the workspace's balance
fchain q bank balances qredoworkspace14a2hpadpsy9h5m6us54
fchain q bank balances qredoworkspace10j06zdk5gyl6vrss5d5

# Burn QAssets - This is also mocked right now
fchain tx qassets burn qredoworkspace14a2hpadpsy9h5m6us54 1 false "" "" 50000 --yes 

# Check both the workspace's balance
fchain q bank balances qredoworkspace14a2hpadpsy9h5m6us54
```