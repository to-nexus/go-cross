# ONE Chain

**This repository** is the official execution client for the **ONE** blockchain by to-nexus.
It is based on **go-ethereum v1.16.9**, with ONE-specific consensus, networking, and fee handling.
The **Breakpoint** hard fork activates **PoSA (Proof of Staked Authority)** using Istanbul BFT and BLS aggregate signatures.
Ethereum execution features are enabled according to each network's fork configuration.

The mainnet flag is `--one`; `--cross` remains a backward-compatible alias.
See [Naming Compatibility](#naming-compatibility) for historical identifiers.

## Table of Contents

1. [Overview](#overview)
2. [Networks](#networks)
3. [Build and Installation](#build-and-installation)
4. [Running a Node](#running-a-node)
5. [Operating a Validator](#operating-a-validator)
6. [Architecture and Consensus](#architecture-and-consensus)
7. [System Contracts and Staking](#system-contracts-and-staking)
8. [RPC and Contract Integration](#rpc-and-contract-integration)
9. [Development](#development)
10. [Upgrading](#upgrading)
11. [Security, Contributing, and License](#security-contributing-and-license)

---

## Overview

- **Consensus**: PoSA using Istanbul BFT and BLS aggregate signatures.
- **Target block time**: 1 second.
- **Networks**: ONE mainnet and ZoneZero testnet; the native coin is ONE.
- **Staking**: a shared pool with stake-based council selection and on-chain rewards.
- **Operation**: full/RPC nodes, archive nodes, and validators with separate node and BLS keys.

### Version Information

| Item | Value |
|---|---|
| Client version | **1.4.0-stable** (see [version/version.go](version/version.go)) |
| Upstream execution client base | **go-ethereum v1.16.9**, with subsequent backports |
| Fork configuration | Network-specific; see [Networks](#networks) and [params/config.go](params/config.go) |
| Go version | **1.25 or higher** |

Use `./build/bin/geth version` to check the version and commit of the installed binary.

### Related Repositories

| Repository | Purpose |
|---|---|
| [`to-nexus/go-cross`](https://github.com/to-nexus/go-cross) | This repository — the ONE execution client (geth fork) |
| [`to-nexus/cross-contracts`](https://github.com/to-nexus/cross-contracts) | Foundry-based system contracts for staking, rewards, validator management, and reserves; repository access may be required |

---

## Networks

| Network | Flag | Chain ID | Breakpoint (UTC) | Notes |
|---|---|---|---|---|
| ONE Mainnet | `--one` | **612055** | 2026-06-01 03:00:00 | Production |
| ZoneZero (testnet) | `--zonezero` | 612044 | 2026-05-15 03:00:00 | Official public testnet |

Selected fork settings for mainnet and ZoneZero are:

| Network | Cancun / Prague (UTC) | Breakpoint timestamp |
|---|---|---|
| ONE Mainnet | 2026-06-01 03:00:00 | `1780282800` |
| ZoneZero | 2026-05-15 03:00:00 | `1778814000` |

The native coin is **ONE**, and all staking, rewards, and slashing are denominated in the native coin.

---

## Build and Installation

### Build Requirements

- Go **1.25+**
- GCC / make / pkg-config
- (optional) Docker / docker-compose

### Build

```bash
git clone https://github.com/to-nexus/go-cross.git
cd go-cross

# Build the client only
make geth

# Build the client and auxiliary tools (abigen, devp2p, evm, etc.)
make all
```

Build artifacts are placed under `./build/bin/`.

### Docker Image

```bash
docker build -t to-nexus/go-cross:latest -f Dockerfile .
```

---

## Running a Node

### Initialize the Data Directory

Initialize mainnet with `init one`:

```bash
./build/bin/geth --datadir <YOUR_DATADIR> init one
```

To create a custom genesis and initialize a separate chain, follow [Create a Custom Chain](#create-a-custom-chain).

### Full / RPC Node

```bash
./build/bin/geth \
  --one \
  --datadir <YOUR_DATADIR> \
  --syncmode full \
  --http --http.addr 127.0.0.1 --http.port 22001 \
  --http.api eth,net,web3,istanbul,cross \
  --ws   --ws.addr   127.0.0.1 --ws.port 32001 \
  --ws.api  eth,net,web3,istanbul,cross \
  --port 30001 --maxpeers 50
```

This example binds RPC to localhost.
Use `0.0.0.0` only when remote access is intended and a firewall or authenticated reverse proxy controls access.
Keep `debug` and administrative methods on trusted interfaces, and restrict the method set exposed by a public proxy.
Set explicit HTTP CORS domains and WebSocket origins only for browser clients that need them; `*` allows any origin and provides no authentication.

### Configuration File

This is a full/RPC node configuration for ONE mainnet.
Set `DataDir` to a writable directory, then run `./build/bin/geth --one --config config.toml`.
For validator operation, also pass `--mine --nodekey <PATH_TO_NODE_KEY> --bls.key <PATH_TO_BLS_SECRET_KEY>`.

```toml
[Eth]
SyncMode = "full"
NetworkId = 612055           # ONE mainnet

[Node]
DataDir = "/data/one"
HTTPHost = "127.0.0.1"
HTTPPort = 22001
HTTPModules = ["net", "web3", "eth", "istanbul", "cross"]
WSHost = "127.0.0.1"
WSPort = 32001
WSModules = ["net", "web3", "eth", "istanbul", "cross"]

[Node.P2P]
ListenAddr = ":30001"
MaxPeers = 50
NoDiscovery = false
DiscoveryV4 = true
DiscoveryV5 = false
StaticNodes = []

[Node.HTTPTimeouts]
# Durations are integer nanoseconds.
ReadTimeout  = 30000000000
WriteTimeout = 30000000000
IdleTimeout  = 120000000000
```

Explicit command-line flags take precedence over the TOML configuration.
Network preset flags also select the network's genesis and network ID; `NetworkId` alone does not define a custom genesis.
Generate a complete configuration for this binary with `./build/bin/geth --one --datadir <YOUR_DATADIR> dumpconfig config.toml`.
See [Full / RPC Node](#full--rpc-node) before enabling remote RPC access.

### Archive Node

To retain all historical account and storage states with the path-based database, add:

```bash
  --gcmode archive --state.scheme path --history.state 0
```

Without `--history.state 0`, state history is limited to the most recent 300,000 blocks by default, even with `--gcmode archive`.
Path-based historical state storage does not retain every historical trie node for arbitrary historical Merkle proofs.
For a hash-based archive that retains trie nodes, use `--gcmode archive --state.scheme hash` from the start of synchronization.
Changing retention settings cannot restore data already pruned; a complete archive may require a fresh full synchronization.
Keep the existing database's state scheme when restarting it; switching schemes requires a separate data directory and synchronization.

### Metrics / Profiling

```bash
  --metrics --metrics.addr 127.0.0.1 --metrics.port 6060 \
  --pprof   --pprof.addr   127.0.0.1 --pprof.port 6061
```

### Create a Custom Chain

`geth make-genesis` creates a genesis file with initialized contracts.
It can generate fresh validator keys or use an existing genesis configuration.
The example below creates an isolated PoSA chain with one validator and requires `jq` in addition to the built client.

#### Generate the Genesis and Keys

From the repository root, create a new working directory outside the checkout:

```bash
GETH_BIN="$PWD/build/bin/geth"
umask 077
mkdir -p <CUSTOM_CHAIN_DIR>
cd <CUSTOM_CHAIN_DIR>

"$GETH_BIN" make-genesis --breakpoint --validators 1 --json > keys.json
```

Without positional arguments, the command writes `genesis.json` in the current directory.
`--json` writes the generated admin, beneficiary, operator, validator, and BLS keys to standard output, captured here in `keys.json`.
The key file contains plaintext secrets; keep it private and out of source control, and use a new directory for each generated chain.
The generated admin account is funded, and `--breakpoint` initializes the system contracts, registers the initial validators, and credits each operator with its initial stake.
These genesis validators do not need the separate registration workflow described in [Operating a Validator](#operating-a-validator).

| Option | Effect |
|---|---|
| `--validators N` | Generate `N` validators; default 3, minimum 1; available only without input files |
| `--json` | Print generated keys as JSON; available only without input files |
| `--breakpoint` | Initialize PoSA system contracts at genesis; rewards start at block 1 |
| `--bridge` | Add and initialize the bridge proxy and implementation |
| `--multicall` | Add Multicall3 |
| `--singleton-factory` / `--sf` | Add deterministic deployment proxies |
| `--all` | Enable all four contract options above |
| `--force` / `-f` | Allow overwriting the output genesis file |

`CrossEx` is included regardless of these options.
Bridge ownership defaults to the PoSA admin; `--bridge.owner`, `--bridge.dev`, and `--bridge.threshold` override the owner, reward wallet, and signature threshold when the bridge is enabled.
Without `--breakpoint` or `--all`, generating keys alone does not initialize the PoSA contracts.

#### Initialize and Start the Validator

Extract the first validator's node key and BLS secret key, then initialize a fresh data directory:

```bash
jq -r '.validators[0].validator.privateKey | ltrimstr("0x")' keys.json > nodekey
jq -r '.validators[0].signer.secretKey | ltrimstr("0x")' keys.json > bls.key
chmod 600 keys.json nodekey bls.key

"$GETH_BIN" --datadir ./data init ./genesis.json

CUSTOM_CHAIN_ID="$(jq -r '.config.chainId' genesis.json)"
"$GETH_BIN" \
  --datadir ./data --networkid "$CUSTOM_CHAIN_ID" \
  --syncmode full --mine \
  --nodekey ./nodekey --bls.key ./bls.key \
  --nodiscover --bootnodes "" --maxpeers 0 \
  --port 30303 \
  --http --http.addr 127.0.0.1 --http.port 8545 \
  --http.api eth,net,web3,istanbul,cross
```

Use the validator's `privateKey` for `--nodekey`, not the operator's wallet key.
Run the custom chain without a built-in network preset such as `--one` or `--zonezero`.
The generated chain ID is currently `612088`; `--networkid` selects the P2P network ID and does not change `config.chainId`.
Use the input-file mode below when choosing a different chain ID or initial configuration.

In a second terminal, query the local RPC endpoint to check the chain ID and confirm that the block height increases:

```bash
curl -sS http://127.0.0.1:8545 \
  -H 'Content-Type: application/json' \
  --data '[{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1},{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":2}]'
```

For multiple validators, generate the full set once, initialize every node with the same `genesis.json`, and extract a different `validators[i]` key pair for each node.
Use separate data directories and ports, allow peers instead of `--maxpeers 0`, and connect the nodes explicitly through static peers or `admin.addPeer` over local IPC.
Enough validators must be online and connected to reach the [consensus quorum](#round-flow).

#### Use an Existing Genesis Configuration

Prepare an uninitialized genesis input with your chosen `config.chainId`, funded accounts in `alloc`, and Istanbul/PoSA settings, then run:

```bash
"$GETH_BIN" make-genesis --breakpoint ./genesis-input.json ./custom-genesis.json
"$GETH_BIN" --datadir ./custom-data init ./custom-genesis.json
```

For `--breakpoint`, the input must have London active at genesis, a nonzero PoSA admin, positive council and validator-epoch periods, and at least one complete validator entry with an ID of 3 to 9 ASCII letters or digits, operator address, validator address, and 48-byte BLS public key.
This mode uses the configured identities and does not generate private keys; supply matching node and BLS keys when starting the nodes.
`--validators` and `--json` are not accepted with input files.
The builder preserves account balances, initializes the selected contract state, and sets the required execution/Breakpoint fork times to zero and `rewardStartBlock` to 1 for PoSA.
Do not feed an already initialized PoSA genesis back into this mode; conflicting contract code or storage is rejected.
Finalize the genesis before initializing any nodes, and use a fresh data directory when changing it.

---

## Operating a Validator

### Validator Requirements

1. **Operator account** — the wallet used to register and manage the validator, meeting the [onboarding requirements](#kyb-whitelist--blacklist).
2. **Consensus node key** — the ECDSA node key used for P2P identity and consensus signing; its derived address is the registered `validatorAddr`.
   Keep this key separate from the operator wallet.
3. **BLS signer key** — mandatory for committing blocks after Breakpoint.
4. At least `minValidatorStake` credited to the **operator account** in the shared pool (initially 1,000,000 ONE).

### BLS Key Preparation and Management

The `geth bls` subcommand handles the full BLS key lifecycle.
By default, keys are stored as encrypted keystore files under `<DATADIR>/bls-keystore`.

| Command | Description |
|---|---|
| `geth bls list` | Print a summary of stored BLS keys |
| `geth bls new` | Generate a new key pair (interactive password, or `--password`) |
| `geth bls update <pubkey>` | Reset the password |
| `geth bls import <hex-file>` | Import a plaintext hex private key and store it encrypted |
| `geth bls decrypt <pubkey>` | Decrypt a key (for backup) |
| `geth bls generate-proof --chain-id <id> <operator-address> <pubkey>` | Produce the 96-byte PoP required by `applyValidator` |

Place command options before positional arguments, and use the same `--datadir` for creation, proof generation, and decryption.
Use the numeric chain ID from [Networks](#networks); the proof is bound to the operator address, BLS public key, and chain ID.

```bash
./build/bin/geth bls new --datadir <YOUR_DATADIR>
./build/bin/geth bls generate-proof --datadir <YOUR_DATADIR> \
  --chain-id 612055 <OPERATOR_ADDRESS> <BLS_PUBLIC_KEY>

# Export the runtime key without displaying it in the terminal.
# Use a new output path and keep this plaintext file readable only by its owner.
umask 077
./build/bin/geth bls decrypt --datadir <YOUR_DATADIR> \
  --password <BLS_PASSWORD_FILE> <BLS_PUBLIC_KEY> > <PATH_TO_BLS_SECRET_KEY>
chmod 600 <PATH_TO_BLS_SECRET_KEY>
```

Start the validator with `--bls.key <PATH_TO_BLS_SECRET_KEY>` after successful decryption.
The password file must contain the keystore password on its first line and be readable only by its owner.
Use `--password` when redirecting decrypted output: interactive prompts are also written to standard output and would corrupt the exported key file.
Keep the plaintext runtime key out of source control and protect it alongside the node key.

> **Backup is essential**: preserve both the node key and the encrypted BLS keystore.
> Recovering an encrypted BLS key requires its password.

### Node Identity and Registration

Start and synchronize a node using the [full/RPC node command](#full--rpc-node), adding `--nodekey <PATH_TO_NODE_KEY>` and `--bls.key <PATH_TO_BLS_SECRET_KEY>` for the keys prepared above.
If you use the node's persisted key instead of an explicit `--nodekey`, preserve that key for subsequent validator runs.

Before registration:

1. Query `istanbul_nodeAddress` and `istanbul_signerAddress` through the local [RPC interface](#rpc-and-contract-integration).
2. Use these values as `validatorAddr` and `signerAddr`, and check that the BLS public key matches the one used for the Proof of Possession.
3. From the operator wallet, call `StakeHub.stake` to meet `minValidatorStake`, then submit the registration call:

```text
StakeHub.applyValidator(
    validatorAddr,   // consensus address derived from the node key
    signerAddr,      // 48-byte BLS public key
    signerProof,     // 96-byte PoP bound to the operator and chain ID
    id               // 3-9 alphanumeric characters, e.g. "myval01"
)
```

Use the deployed contract's [complete ABI](#standard-evm-calls) for these wallet transactions.
Registration makes the validator eligible for the next council selection; selection still depends on its stake ranking and `maxCouncil`.

### Start the Validator

```bash
./build/bin/geth \
  --one \
  --datadir <YOUR_DATADIR> \
  --syncmode full \
  --mine \
  --nodekey <PATH_TO_NODE_KEY> \
  --bls.key <PATH_TO_BLS_SECRET_KEY>
```

- `--mine` is required to participate actively in rounds (you can enable it as a non-validator, but it is meaningless).
- Without a configured BLS secret key, the log will print `Istanbul: BLS secret key is not configured - node cannot participate in PoSA`, and the node cannot take part in consensus.
- The address derived from `--nodekey` **must match exactly** the registered `validatorAddr`.
- If `--nodekey` is omitted, the node loads or generates its persisted node key in the data directory; preserve it across restarts and migrations.
- Istanbul sets `etherbase` from the node key, overriding `--miner.etherbase`.
- `--bls.key` loads a 32-byte plaintext hex secret key; `--bls.keystore` is used by key-management commands and does not load a consensus key at startup.

### Operational Checklist

- Poll `StakeHub.getValidatorInfo(operator)` regularly via `eth_call`.
- Track cumulative pooled / commission / fee / burned via `RewardHub.getValidatorLog(validatorAddr)`.
- Watch `ValidatorSlash.getSlashInfo(validatorAddr)` for the contract revision listed under [Contract Versions and Addresses](#contract-versions-and-addresses); newer implementations may index this query by operator address, so use the deployed contract's ABI and semantics.
- Compare `istanbul_nodeAddress` and `istanbul_signerAddress` with the consensus address and BLS public key registered in `StakeHub`.
- Query `istanbul_isValidator` to check current set membership, and `istanbul_getConfig` for consensus timing and `rewardStartBlock`.
- Membership alone does not prove that the BLS key is configured or the node is successfully committing; monitor `istanbul_status` and consensus logs.
- Reward distribution also runs in blocks without user transactions; the target block period is one second, without an extra empty-block delay by default.

### Recovering Active Status

| Situation | Recovery |
|---|---|
| Voluntary `suspend(endTime)` | Becomes eligible for selection again after `endTime`; re-entry depends on the next council rotation and stake ranking |
| `jail` state | After the jail period expires, call `unjail()` with stake ≥ `minValidatorStake`, then wait for selection |
| Blacklisted | Keeper removes from the blacklist, then call `unjail()` as above |

Recovery behavior follows the [contract versions](#contract-versions-and-addresses) described below; check the deployed implementation before maintenance.

---

## Architecture and Consensus

### Architecture

```
┌────────────────────────────────────────────────────────────────────┐
│                                                                    │
│  ┌─────────────┐   ┌──────────────────┐   ┌─────────────────────┐  │
│  │  P2P / eth  │   │ Istanbul Engine  │   │  EVM (Prague-ready) │  │
│  │  txpool     │◄──┤  + PoSA module   │──►│  Cancun blob, etc.  │  │
│  └─────────────┘   │  + BLS seal      │   └─────────────────────┘  │
│         ▲          └────────┬─────────┘             │              │
│         │                   │ system tx             │ system call  │
│         │                   ▼                       ▼              │
│         │          ┌──────────────────────────────────────────┐    │
│         │          │  System Contracts (predeployed)          │    │
│         │          │  - StakeHub        0x...1002             │    │
│         │          │  - RewardHub       0x...1003             │    │
│         │          │  - ValidatorSet    0x...1001             │    │
│         │          │  - ValidatorSlash  0x...1004             │    │
│         │          │  - DelegationPool / CrossReserve /       │    │
│         │          │    ProtocolReserve                       │    │
│         │          └──────────────────────────────────────────┘    │
│         ▼                                                          │
│   JSON-RPC (eth / net / web3 / istanbul / cross / debug ...)       │
└────────────────────────────────────────────────────────────────────┘
```

- The **Istanbul Engine** retains the round-change / preprepare / prepare / commit message flow, selects the council from on-chain stake data, and invokes system contracts to update the validator set and distribute rewards.
- The **block header's Extra data** carries BLS signer information; after Breakpoint, blocks are committed with BLS aggregate signatures.
- During PoSA block finalization, the node emits **system transactions** (gas price = 0, sender = coinbase) for reward distribution and, when required, offline slashing and counter mitigation.

### PoSA Activation

The **Breakpoint** hard fork is the **PoSA activation point**.
Before `BreakpointTime`, blocks are still produced by the legacy permissioned validator set.
PoSA requires a configured Istanbul PoSA section and a non-genesis block with `timestamp >= BreakpointTime + 1`.
The configured initial validator set is used until council rotation selects validators by stake.
Reward emission also depends on `RewardStartBlock` and the contract's reward table.

### Validator Selection and Council Rotation

PoSA selects validators by operator stake and finalizes blocks through Istanbul BFT:

1. Eligible operators register through `StakeHub` using the requirements in [Staking and Validator Management](#staking-and-validator-management).
2. The engine selects up to `maxCouncil` eligible validators by operator stake, descending; ties are resolved by registration time, then validator address, ascending.
3. The elected council validators run **Istanbul BFT** rounds to finalize blocks.
4. Per-block inflation rewards flow into `DelegationPool` and are distributed pro-rata to delegators.

Council selection runs on the first block of each new `councilPeriod` (24 hours by default).
It uses the stake ranking snapshot in `StakeHub` and selects up to 21 validators by default.

### Consensus Parameters

| Parameter | Value | Location / Change Authority |
|---|---|---|
| Block Period | 1 s | `params.IstanbulConfig.BlockPeriodSeconds` |
| Empty Block Period | 0 (no additional delay beyond the block period) | `params.IstanbulConfig.EmptyBlockPeriodSeconds` |
| Request Timeout | 5 s | `params.IstanbulConfig.RequestTimeoutSeconds` |
| Max Request Timeout | 60 s | `params.IstanbulConfig.MaxRequestTimeoutSeconds` |
| Proposer Policy | RoundRobin (0) | `params.IstanbulConfig.ProposerPolicy` |
| Epoch Length | 86,400 blocks | `params.IstanbulConfig.EpochLength` |
| Council Period | 86,400 s (24 h) | `params.PoSAConfig.CouncilPeriod` |
| Validator Epoch Length | 300 blocks | `params.PoSAConfig.ValidatorEpochLength` |
| Elasticity Multiplier | 3 | `params.IstanbulConfig.ElasticityMultiplier` |
| Base Fee Change Denom. | 8 | `params.IstanbulConfig.BaseFeeChangeDenominator` |
| Max / Min Base Fee | 1 ONE / 1 Gwei | `params.IstanbulConfig.MaxBaseFee` / `MinBaseFee` |

### Round Flow

The IBFT message flow under `consensus/istanbul/core/` is kept intact:

```
PREPREPARE  →  PREPARE (quorum)  →  COMMIT (quorum)  →  FINAL COMMIT
        └──── on failure, enter the next round via ROUND CHANGE ────┘
```

Quorum is `ceil(2N / 3)`, where `N` is the current validator-set size: 14 signatures for 21 validators.

After PoSA activation, the commit phase is handled with **BLS aggregate signatures**.
A validator node must hold a **BLS Secret Key** separately from its operator ECDSA account.
The `signerProof` (Proof of Possession) submitted during validator registration is verified in `StakeHub._checkSigner` via the BLS12-381 precompile (`0xc0`).

---

## System Contracts and Staking

### Contract Versions and Addresses

The following system contracts are predeployed once PoSA is active.
Source code lives in the [cross-contracts repository](https://github.com/to-nexus/cross-contracts), built and tested with Foundry.
The client bundles runtime bytecode and selected Go bindings under [contracts/breakpoint](contracts/breakpoint).
This guide uses initialization defaults from the source revisions recorded in [contracts/upgrade.go](contracts/upgrade.go):

- [StakeHub / ValidatorSet / ValidatorSlash: `5cbc608`](https://github.com/to-nexus/cross-contracts/tree/5cbc608bd77ef186120308ec8219a016197efba2/src)
- [RewardHub: `373dfdd`](https://github.com/to-nexus/cross-contracts/blob/373dfdd9ec4b53f7f3efc0516a7b7cfaa7cab0f6/src/RewardHub.sol)

Contract values in this guide are initialization defaults, not a live snapshot of either network.
Administrators can change staking limits, council size, rewards, commission, and penalty parameters.
Read the deployed contracts before registering or staking.
Use the ABI matching the deployed implementation; newer contract branches may change functions and recovery behavior.

| Address | Contract | Role |
|---|---|---|
| `0x...1001` | `ValidatorSet` | Stores council addresses and BLS signer keys; updated by node system calls and validator lifecycle operations |
| `0x...1002` | **`StakeHub`** | Validator registration / revocation / jail-unjail; entry point for stake / unstake / restake |
| `0x...1003` | **`RewardHub`** | Block-reward distribution, emission table, commission settlement, log aggregation |
| `0x...1004` | `ValidatorSlash` | Offline slashing counter; calls `StakeHub.penalizeOffline` when threshold is hit; mitigate |
| (separate) | `DelegationPool` | Single-pool staking / reward accounting (Synthetix StakingRewards pattern) |
| (separate) | `CrossReserve` | Inflation reserve. Funds are pulled by DelegationPool during reward distribution |
| (separate) | `ProtocolReserve` | Holds slashed / forfeited / unpaid funds; supports time-locked withdrawal scheduling |

Functions marked `systemCall` require **`msg.sender == block.coinbase && tx.gasprice == 0`** (`SystemBase._isSystemCall`).
User staking and validator-management functions have separate access checks.

### Staking and Validator Management

All staking actions are exposed as external functions of `StakeHub` (`0x...1002`).
The state-changing staking actions below check pause and blacklist status; read-only queries remain available.

#### Delegator Actions

| Function | Description |
|---|---|
| `stake() payable` | Stake for the caller; initial minimum deposit **10 ONE** (`minStake`) |
| `stakeFor(address account) payable` | Stake on behalf of another address |
| `unstake(uint256 amount)` | Begin unbonding. Claimable after 14 days; up to 10 outstanding requests per account |
| `claimUnstake([recipient])` | Claim all matured unbonding requests in one call |
| `restakeReward()` | Compound pending rewards back into stake immediately (no withdrawal) |
| `getStakedAmount(account)` / `getEarnedReward(account)` / `getUnbondingInfo(account)` | Read balance / earned reward / unbonding queue |

Internally, all stakes are pooled into a single `DelegationPool`, and reward accounting uses the **Synthetix StakingRewards** scheme: the difference between the accumulated **`rewardPerStakeStored`** and the per-user **`rewardPerStakePaid`**.
Ordinary staking does not assign delegation weight to a selected validator.
Validator eligibility and ranking use the balance credited to the operator account.
`stakeFor(operator)` credits that operator's balance; it does not preserve a separate stake owned by the payer.

```
earned(account)
  = balance · (rewardPerStakeStored − rewardPerStakePaid[account]) / 1e18
  + reward[account]
```

#### Validator Actions

| Function | Description |
|---|---|
| `applyValidator(validatorAddr, signerAddr, signerProof, id) payable` | Register with the whitelist policy below, a 3-9-character alphanumeric `id`, BLS Proof of Possession, and operator stake ≥ `minValidatorStake` (initially **1,000,000 ONE**) |
| `updateValidator(newValidatorAddr)` | Change the consensus node address. The old value cannot be reused; 1-day cooldown |
| `updateSigner(signerAddr, signerProof)` | Rotate the BLS signer key; 1-day cooldown |
| `revokeValidator()` | Voluntarily resign as validator. The id, validator address, and BLS key are permanently sealed and cannot be reused |
| `unjail()` | Unjail oneself after the jail period |
| `suspend(uint256 endTime)` | Voluntary suspension (e.g. maintenance). Immediately removed from the council |

#### Status Codes

`StakeHub.getValidatorDescriptions` returns the following status values:

| Value | Meaning |
|---|---|
| 0 | inactive |
| 1 | active candidate (registered but not in the council) |
| 2 | active council (actively producing blocks) |
| 3 | jailed (penalized via slashing) |
| 4 | suspended (voluntary) |

#### KYB Whitelist / Blacklist

- Validator onboarding requires **KYB (Know-Your-Business)** approval and operator whitelist registration under the network's onboarding policy.
- The contract enforces whitelist membership through `onlyWhiteListed` while the whitelist is enabled; it does not verify KYB documents itself.
- The whitelist manager can disable the whitelist check, so read `isWhiteListEnabled()` and `isWhiteListed(operator)` to verify the on-chain requirement.
- Ordinary staking does not require whitelist membership, but pause and blacklist restrictions still apply.
- Blacklisted accounts are blocked from staking, unstaking, and claiming; if the account is a validator, it is jailed immediately (`StakeHub.addToBlackList`).

### Rewards and Fees

Emission starts according to `RewardStartBlock` and the configured reward table.

#### Reward Flow at a Glance

```
[block N committed]
   └─ node emits system tx: RewardHub.distributeReward(validator, tip)
        │   msg.value = execution base fee + blob fee + tip
        ▼
   ┌──────────────────────────────────────────────────────────────────┐
   │ 1) Execution base fee + blob fee → burned to 0x...dEaD          │
   │ 2) Inflation reward = emission table[block N]                    │
   │ 3) block commission = reward * commissionRate                   │
   │ 4) RewardHub → DelegationPool.distributeReward()                 │
   │      a) pull `reward` from CrossReserve                          │
   │      b) reward − block commission → pool's rewardPerStake        │
   │      c) block commission + tip → operator's stake               │
   └──────────────────────────────────────────────────────────────────┘
```

#### Inflation / Reward Curve

`RewardHub` is initialized with the following emission table (with 1 year = `BLOCKS_PER_YEAR = 31_536_000` blocks, assuming 1-second blocks):

| Start block (offset) | Per-block reward |
|---|---|
| `startBlock + 0y` | **9.5129376 ONE** |
| `+1y` | 4.7564688 |
| `+2y` | 2.3782344 |
| `+3y` | 1.1891172 |
| `+4y` | 0.5945586 |
| `+5y` | 0.2972793 |
| `+6y` | 0.14863965 |
| `+7y` | 0.074319825 |
| `+8y` | 0.0371599125 |
| `+9y` | 0.01857995625 |

The schedule halves the reward every 31,536,000 blocks, approximately one year at the target block interval.
It is based on block height rather than elapsed calendar time.
The table can be redefined by the admin via `addRewardTableEntry` / `removeRewardTableEntry`, and every change emits `RewardTableAdded` / `RewardTableRemoved`.

#### Commission and Tips

- **Default commission 5 %** (`commissionRate = 500 / 10_000`); cap 50 %.
- Commission = "block reward × commission rate + user-paid tip (priority fee)", and it is **immediately compounded** into the operator (`operatorAddr`) stake (no separate withdrawal needed).
- If the validator is inactive (jailed / suspended), `validatorToActiveOperator` returns the zero address and the commission flows to `ProtocolReserve` instead.

#### Estimated APR

`RewardHub.estimatedApr(blockNum, blocksPerYear)` estimates the pool's annual return from the selected reward tier:

```
netPerBlock = floor(perBlock × (10_000 − commissionRate) / 10_000)
aprWad = floor(blocksPerYear × netPerBlock × 1e18 / totalStaked)
```

The return value uses **WAD units: `1e18 = 100%`**; divide by `1e16` to display a percentage.
For example, `5e16` represents 5% APR.
The function returns zero when the pool, reward tier, annual block count, or total stake is unavailable or zero.
It assumes successful per-block distribution and does not forecast future reward tiers or compounding.
A validator also earns block commissions in addition to the pool rewards on its operator stake.

#### Base Fee Burn

- The full amount of `baseFee × gasUsed` (EIP-1559) flows in as the `baseFee` portion of `RewardHub.distributeReward`'s `msg.value`, and is sent directly to `0x000000000000000000000000000000000000dEaD` to be burned.
- Blob fees are included in the non-tip amount sent for burning.
- If the burn transfer fails, the amount is classified as retained and a `RewardRetained(Burn, amount)` event is emitted.

### Slashing and Validator Penalties

`ValidatorSlash` (`0x...1004`) is responsible for:

#### Offline Slashing

1. Validators that fail to produce a block (round timeout) are counted via the `slashOffline([...])` system call.
2. Once the cumulative count reaches **`offlineSlashThreshold` (default 50)**, `StakeHub.penalizeOffline` is invoked.
3. Penalty:
   - Slash **10,000 ONE** from the operator's stake (moved from DelegationPool to ProtocolReserve)
   - **2-day jail** (`offlineJailDuration`); removed from the council immediately
4. After the jail period expires, the operator must call `unjail()` with at least `minValidatorStake` remaining (initially 1,000,000 ONE).

#### Mitigate

- `mitigate()` is called as a system transaction at council-period boundaries, decreasing tracked offline counters by integer division `offlineSlashThreshold / mitigateRate` (initially `50 / 4 = 12`).
- Both `offlineSlashThreshold` and `mitigateRate` are configurable contract parameters.
- When the count reaches 0, the validator is removed from the slash tracker and returns to a clean state.

---

## RPC and Contract Integration

In addition to the standard modules (`eth`, `net`, `web3`, `txpool`, `debug`, `admin`), this client registers `istanbul` and `cross`.
HTTP and WebSocket availability depends on `--http.api` and `--ws.api` (or their TOML equivalents).

### `istanbul_*`

- `istanbul_getSnapshot(blockNumber)` — validator snapshot, including the validator set and vote state
- `istanbul_getValidators(blockNumberOrTag)` — the active validator set at the given block
- `istanbul_nodeAddress` — consensus ECDSA address derived from the node key
- `istanbul_signerAddress` — configured BLS public key
- `istanbul_isValidator(blockNumber)` — whether the node address belongs to the validator set; an omitted block number uses the current head
- `istanbul_getConfig(blockNumber)` — consensus parameters, including council period, validator epoch length, and reward start block
- `istanbul_status(startBlock, endBlock)` — validator activity over a block range, limited to 86,400 blocks
- `istanbul_candidates` — local legacy validator-vote proposals; use `StakeHub` to query PoSA candidates

For example, check the configured node identity on the local HTTP endpoint:

```bash
curl -sS http://127.0.0.1:22001 \
  -H 'Content-Type: application/json' \
  --data '{"jsonrpc":"2.0","method":"istanbul_nodeAddress","params":[],"id":1}'
```

### `cross_*`

- `cross_getCheckpointProof(blockNumber, startBlock, endBlock)` — Merkle proof for a block hash in an inclusive block range

Pass numeric block quantities (hex strings), not tags such as `latest`.
The range must satisfy `0 <= startBlock < endBlock`, contain `blockNumber`, and contain at most 86,400 blocks.
See [internal/ethapi/crossapi.go](internal/ethapi/crossapi.go) for the implementation.
Read consensus settings with `istanbul_getConfig` and reward data through contract calls.

### Standard EVM Calls

Staking, validator registration, and reward queries are all standard `eth_call` / `eth_sendRawTransaction` flows.
Use the complete ABI matching the deployed contract for wallet and application calls.
For example, from the matching `cross-contracts` checkout:

```bash
forge soldeer install
forge inspect StakeHub abi --json > StakeHub.abi.json
forge inspect RewardHub abi --json > RewardHub.abi.json
```

---

## Development

### Maintaining Node Bindings

The `scripts/gen-binding` helper generates a filtered function ABI and runtime bytecode for node integration.
Its output is not a complete application ABI.
It requires Foundry, `jq`, and this client's `abigen` on `PATH`.

```bash
# In go-cross, build the matching tools.
make all
export PATH="$PWD/build/bin:$PATH"

# In a sibling cross-contracts checkout at the intended contract revision.
cd ../cross-contracts
forge soldeer install
scripts/gen-binding StakeHub RewardHub breakpoint

# Generated files are in the cross-contracts root, not the client directory.
cp stake_hub.go reward_hub.go ../go-cross/contracts/breakpoint/
```

Review generated ABI and runtime-bytecode changes together with the target network's upgrade configuration.
Regenerate bindings only as part of an intended client/contract update; do not replace bundled bytecode merely because a newer contract branch exists.

---

## Upgrading

### Upgrade Procedure

1. Record the installed client version, network, database backend, and state scheme.
2. Stop the node cleanly and back up its data directory, node key, BLS keys, and configuration.
3. Check the target release's migration instructions for the exact source version and network.
   This README does not guarantee database compatibility or rollback across arbitrary versions.
4. Preserve the existing state scheme when reusing a compatible database.
   Use a separate data directory and full synchronization when changing schemes or rebuilding missing archive history.
5. After restart, verify chain ID, sync progress, node/BLS identities, and validator activity before resuming normal operation.

### Legacy QBFT Migration

Breakpoint transitions the chain from the legacy permissioned validator set to stake-based council selection with BLS commits.
It introduces shared-pool staking, on-chain reward distribution, and validator penalties.
See [PoSA Activation](#posa-activation) for the transition conditions and [System Contracts and Staking](#system-contracts-and-staking) for the current behavior.

### Naming Compatibility

**Naming update (effective 2026-07-23):** The chain was renamed from **CROSS chain** to **ONE chain**, and its native coin from **CROSS** to **ONE**.
The mainnet flag is **`--one`**; **`--cross`** remains a backward-compatible alias.
Technical identifiers such as the `go-cross` / `cross-contracts` repositories, the `cross_*` RPC namespace, and `CrossReserve` retain their existing names.

---

## Security, Contributing, and License

### Security

For security disclosures, follow [SECURITY.md](SECURITY.md) and use a verified private reporting channel for ONE-specific issues.
Do not include vulnerability details in public issues or pull requests.

### Contributing

Issues and pull requests are welcome.
File general inquiries as GitHub Issues.

### License

- All code under `cmd/` is licensed under the **GNU GPL v3.0** (`COPYING`).
- All other library code is licensed under the **GNU LGPL v3.0** (`COPYING.LESSER`).
