# ONE Chain

**This repository** is the official execution client for the **ONE** blockchain by to-nexus.
It is built on top of the **go-ethereum v1.15** codebase and introduces a brand-new consensus engine — **PoSA (Proof of Staked Authority)** — together with the **Breakpoint** hard fork, while remaining fully compatible with modern Ethereum specifications such as EIP-1559, Cancun, and Prague.

> **Naming update (effective 2026-07-23):** The chain was renamed from **CROSS chain** to **ONE chain**, and its native coin from **CROSS** to **ONE**. All chain and coin references in this document use the new names. The mainnet flag is now **`--one`** (the old **`--cross`** still works as a backward-compatible alias). Technical identifiers that mirror the codebase and repositories — the `go-cross` / `cross-contracts` repositories, the `cross_*` RPC namespace, and system-contract names such as `CrossReserve` — are unchanged.

> This release **completely replaces the previous permissioned QBFT (Quorum BFT) consensus with PoSA**.
> Anyone who stakes the required amount can now participate as a validator, and ordinary users can earn block rewards through delegation.

## Related Repositories

| Repository | Purpose | Visibility |
|---|---|---|
| [`to-nexus/go-cross`](https://github.com/to-nexus/go-cross) | This repository — the Go Cross execution client (geth fork) | Public |
| [`to-nexus/cross-contracts`](https://github.com/to-nexus/cross-contracts) | System contracts that implement PoSA on-chain logic: `StakeHub`, `RewardHub`, `ValidatorSet`, `ValidatorSlash`, `DelegationPool`, `CrossReserve`, `ProtocolReserve`, `IstanbulParam`. Built with Foundry. | *Scheduled to become public.* |

---

## Table of Contents

1. [Key Facts](#1-key-facts)
2. [Version Information](#2-version-information)
3. [Networks / Chain IDs](#3-networks--chain-ids)
4. [Architecture Overview](#4-architecture-overview)
5. [PoSA Consensus Engine](#5-posa-consensus-engine)
6. [System Contracts](#6-system-contracts)
7. [Staking](#7-staking)
8. [Block Rewards / Fees / Revenue Model](#8-block-rewards--fees--revenue-model)
9. [Slashing and Validator Penalties](#9-slashing-and-validator-penalties)
10. [Installation and Build](#10-installation-and-build)
11. [Running a Node](#11-running-a-node)
12. [Operating a Validator](#12-operating-a-validator)
13. [BLS Key Management](#13-bls-key-management)
14. [Sample Configuration File](#14-sample-configuration-file)
15. [RPC Modules](#15-rpc-modules)
16. [Upgrade Notes (QBFT → PoSA)](#16-upgrade-notes-qbft--posa)
17. [License / Contributing](#17-license--contributing)

---

## 1. Key Facts

- **Consensus engine**: PoSA (Istanbul message protocol + BLS aggregate signatures + stake-weighted validator election)
- **Block time**: 1 second
- **Validator requirements**: ≥ 1,000,000 ONE self-delegation **plus on-chain whitelist registration after KYB (Know-Your-Business) verification**
- **Minimum delegator stake**: 5 ONE
- **Unbonding period**: 14 days
- **Council size**: up to 21 validators (top-staked)
- **Council rotation period**: 86,400 s (24 hours)
- **Block reward**: starts at 9.5129376 ONE/block, halved every year (10 tiers)
- **Validator commission**: 5 % by default (configurable, capped at 50 %)
- **Base fee (EIP-1559 burn)**: 100 % burned (`0x...dEaD`)

---

## 2. Version Information

| Item | Value |
|---|---|
| Execution Layer (EL) | **go-ethereum v1.15 base** |
| Consensus engine | **PoSA** (Istanbul Engine + BLS) |
| Hard forks | Homestead → Berlin → London → Shanghai → Adventure → Cancun → Prague → **Breakpoint** |
| Breakpoint activation (mainnet ONE) | `BreakpointTime = 1780282800` (2026-06-01 03:00:00 UTC) |
| Breakpoint activation (testnet ZoneZero) | `BreakpointTime = 1778814000` (2026-05-15 03:00:00 UTC) |
| Go version | **1.23 or higher** |
| RLP, DB, Trie | 1.15 line |

> The **Breakpoint** hard fork is the **PoSA activation point**.
> Before `BreakpointTime`, blocks are still produced by the legacy permissioned validator set.
> From the first block after that timestamp, stake-weighted validator election and on-chain reward distribution take effect.

---

## 3. Networks / Chain IDs

| Network | Flag | Chain ID | Breakpoint (UTC) | Notes |
|---|---|---|---|---|
| ONE Mainnet | `--one` (alias `--cross`) | **612055** | 2026-06-01 03:00:00 | Production |
| ZoneZero (testnet) | `--zonezero` | 612044 | 2026-05-15 03:00:00 | Official public testnet |

The native coin is **ONE**, and all staking, rewards, and slashing are denominated in the native coin.

---

## 4. Architecture Overview

```
┌────────────────────────────────────────────────────────────────────┐
│                          Go Cross (geth)                           │
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

- The **Istanbul Engine** retains the round-change / preprepare / prepare / commit message flow, but delegates validator-set decisions and reward distribution to on-chain system contracts.
- The **block header's Extra data** carries BLS signer information; after Breakpoint, blocks are committed with BLS aggregate signatures.
- At the end of every block, the node emits **system transactions** (gas price = 0, sender = coinbase) to invoke `RewardHub.distributeReward` / `ValidatorSlash.slashOffline` / `ValidatorSlash.mitigate` etc.

---

## 5. PoSA Consensus Engine

### 5.1 Overview

PoSA is a **stake-weighted authority consensus** in the spirit of BNB Chain, redesigned for the ONE environment.
Whereas the previous QBFT was a permissioned scheme in which “N pre-whitelisted validators ran BFT consensus”, PoSA works as follows:

1. **Anyone** can apply to become a validator through `StakeHub`.
2. The higher the combined stake (self + delegations), the higher the chance of being elected to the **Council**.
3. The elected council validators run **Istanbul BFT** rounds to finalize blocks.
4. Per-block inflation rewards flow into `DelegationPool` and are distributed pro-rata to delegators.

### 5.2 Configuration Parameters (mainnet defaults)

| Parameter | Value | Location / Change Authority |
|---|---|---|
| Block Period | 1 s | `params.IstanbulConfig.BlockPeriodSeconds` |
| Empty Block Period | 0 (empty blocks may be produced immediately) | same |
| Request Timeout | 5 s | same |
| Max Request Timeout | 60 s | same |
| Proposer Policy | RoundRobin (0) | same |
| Epoch Length | 86,400 blocks | same |
| Council Period | 86,400 s (24 h) | `IstanbulParam` / PoSA Config |
| Elasticity Multiplier | 3 | same |
| Base Fee Change Denom. | 8 | same |
| Max / Min Base Fee | 1 ONE / 1 Gwei | same |

### 5.3 Round Flow

The IBFT message flow under `consensus/istanbul/core/` is kept intact:

```
PREPREPARE  →  PREPARE (2f+1)  →  COMMIT (2f+1)  →  FINAL COMMIT
        └──── on failure, enter the next round via ROUND CHANGE ────┘
```

After PoSA activation, the commit phase is handled with **BLS aggregate signatures**.
A validator node must hold a **BLS Secret Key** separately from its operator ECDSA account.
The `signerProof` (Proof of Possession) submitted during validator registration is verified in `StakeHub._checkSigner` via the BLS12-381 precompile (`0xc0`).

### 5.4 Validator-Set Refresh Cycles

- **Council Period (`councilPeriod`, default 24 h)**:
  On the first block where `OnNewCouncilPeriod` is true, the council (the 21 block-producing validators) is rotated based on the stake ranking snapshot from `StakeHub`.

---

## 6. System Contracts

The following system contracts are predeployed once PoSA is active.
Source code lives in the `cross-contracts` repository — <https://github.com/to-nexus/cross-contracts> (scheduled to become public). It is built / tested with Foundry.

| Address | Contract | Role |
|---|---|---|
| `0x...1000` | `IstanbulParam` | Checkpointed consensus parameters (epoch / council period / base fee, etc.) |
| `0x...1001` | `ValidatorSet` | Stores active validator addresses + BLS signer keys; updated only by node system calls |
| `0x...1002` | **`StakeHub`** | Validator registration / revocation / jail-unjail; entry point for stake / unstake / restake |
| `0x...1003` | **`RewardHub`** | Block-reward distribution, emission table, commission settlement, log aggregation |
| `0x...1004` | `ValidatorSlash` | Offline slashing counter; calls `StakeHub.penalizeOffline` when threshold is hit; mitigate |
| (separate) | `DelegationPool` | Single-pool staking / reward accounting (Synthetix StakingRewards pattern) |
| (separate) | `CrossReserve` | Inflation reserve. Funds are pulled by DelegationPool during reward distribution |
| (separate) | `ProtocolReserve` | Holds slashed / forfeited / unpaid funds; supports time-locked withdrawal scheduling |

System-contract entry points are gated by **`msg.sender == block.coinbase && tx.gasprice == 0`** (`SystemBase._isSystemCall`); ordinary users cannot invoke them.

---

## 7. Staking

All staking actions are exposed as external functions of `StakeHub` (`0x...1002`).
Every function honors `whenNotPaused` and the blacklist check.

### 7.1 Delegator Actions

| Function | Description |
|---|---|
| `stake() payable` | Stake for the caller. **Minimum 5 ONE** |
| `stakeFor(address account) payable` | Stake on behalf of another address |
| `unstake(uint256 amount)` | Begin unbonding. Claimable after 14 days; up to 10 outstanding requests per account |
| `claimUnstake([recipient])` | Claim all matured unbonding requests in one call |
| `restakeReward()` | Compound pending rewards back into stake immediately (no withdrawal) |
| `getStakedAmount(account)` / `getEarnedReward(account)` / `getUnbondingInfo(account)` | Read balance / earned reward / unbonding queue |

Internally, all stakes are pooled into a single `DelegationPool`, and reward accounting uses the **Synthetix StakingRewards** scheme: the difference between the accumulated **`rewardPerStakeStored`** and the per-user **`rewardPerStakePaid`**.

```
earned(account)
  = balance · (rewardPerStakeStored − rewardPerStakePaid[account]) / 1e18
  + reward[account]
```

### 7.2 Validator Actions

| Function | Description |
|---|---|
| `applyValidator(validatorAddr, signerAddr, signerProof, id) payable` | Apply to become a validator. **Requires on-chain whitelist registration after KYB (Know-Your-Business) verification** + 3-9-char alphanumeric `id` + BLS Proof-of-Possession verification + cumulative stake ≥ **1,000,000 ONE** |
| `updateValidator(newValidatorAddr)` | Change the consensus node address. The old value cannot be reused; 1-day cooldown |
| `updateSigner(signerAddr, signerProof)` | Rotate the BLS signer key; 1-day cooldown |
| `revokeValidator()` | Voluntarily resign as validator. The id, validator address, and BLS key are permanently sealed and cannot be reused |
| `unjail()` | Unjail oneself after the jail period |
| `suspend(uint256 endTime)` | Voluntary suspension (e.g. maintenance). Immediately removed from the council |

### 7.3 Status Codes

`StakeHub.getValidatorDescriptions` returns the following status values:

| Value | Meaning |
|---|---|
| 0 | inactive |
| 1 | active candidate (registered but not in the council) |
| 2 | active council (actively producing blocks) |
| 3 | jailed (penalized via slashing) |
| 4 | suspended (voluntary) |

### 7.4 KYB Whitelist / Blacklist

- A would-be validator operator must first complete the **KYB (Know-Your-Business)** process, and only KYB-approved operator addresses are added to the on-chain whitelist.
- `applyValidator` checks only **whitelist membership** (`onlyWhiteListed`), so any address that has not passed KYB cannot register as a validator under any circumstance.
  → Regular stake / unstake operations are available to everyone without KYB.
- Blacklisted accounts are blocked from staking, unstaking, and claiming; if the account is a validator, it is jailed immediately (`StakeHub.addToBlackList`).

---

## 8. Block Rewards / Fees / Revenue Model

### 8.1 Reward Flow at a Glance

```
[block N committed]
   └─ node emits system tx: RewardHub.distributeReward(validator, tip)
        │   msg.value = baseFee + tip (the portion of gas the node collected)
        ▼
   ┌──────────────────────────────────────────────────────────────────┐
   │ 1) Entire baseFee  →  burned to 0x...dEaD (EIP-1559 burn)        │
   │ 2) Inflation reward = emission table[block N]                    │
   │ 3) commission = reward * 5%  +  tip                              │
   │ 4) RewardHub → DelegationPool.distributeReward()                 │
   │      a) pull `reward` from CrossReserve                          │
   │      b) (reward − commission) → added to pool's rewardPerStake   │
   │      c) commission (+tip) → auto-redeposited as operator's stake │
   └──────────────────────────────────────────────────────────────────┘
```

### 8.2 Inflation / Reward Curve

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

In short, a **10-tier halving schedule that halves every year**.
The table can be redefined by the admin via `addRewardTableEntry` / `removeRewardTableEntry`, and every change emits `RewardTableAdded` / `RewardTableRemoved`.

### 8.3 Commission and Tips

- **Default commission 5 %** (`commissionRate = 500 / 10_000`); cap 50 %.
- Commission = "block reward × commission rate + user-paid tip (priority fee)", and it is **immediately compounded** into the operator (`operatorAddr`) stake (no separate withdrawal needed).
- If the validator is inactive (jailed / suspended), the operator mapping is cleared and the commission flows to `ProtocolReserve` instead.

### 8.4 Estimated APR

`RewardHub.estimatedApr(blockNum, blocksPerYear)` returns exactly the formula below:

```
APR  =  blocksPerYear × perBlock × (1 − commissionRate)
        ───────────────────────────────────────────────────
                        total staked amount
```

A validator's effective yield is the above APR plus their own commission and the rewards on their self-delegated stake.

### 8.5 Base Fee Burn

- The full amount of `baseFee × gasUsed` (EIP-1559) flows in as the `baseFee` portion of `RewardHub.distributeReward`'s `msg.value`, and is sent directly to `0x000000000000000000000000000000000000dEaD` to be burned.
- If the burn transfer fails, the amount is classified as retained and a `RewardRetained(Burn, amount)` event is emitted.

---

## 9. Slashing and Validator Penalties

`ValidatorSlash` (`0x...1004`) is responsible for:

### 9.1 Offline Slashing

1. Validators that fail to produce a block (round timeout) are counted via the `slashOffline([...])` system call.
2. Once the cumulative count reaches **`offlineSlashThreshold` (default 50)**, `StakeHub.penalizeOffline` is invoked.
3. Penalty:
   - Slash **10,000 ONE** from the operator's stake (moved from DelegationPool to ProtocolReserve)
   - **2-day jail** (`offlineJailDuration`); removed from the council immediately
4. To exit jail, the validator must call `unjail()` themselves, and at that moment the remaining stake must be ≥ 1,000,000 ONE.

### 9.2 Mitigate

- `mitigate()` is periodically called as a system tx, decreasing every validator's offline counter by `offlineSlashThreshold / MITIGATE_RATE` (= 50 / 4 = 12.5).
- When the count reaches 0, the validator is removed from the slash tracker and returns to a clean state.

---

## 10. Installation and Build

### 10.1 Prerequisites

- Go **1.23+**
- GCC / make / pkg-config
- (optional) Docker / docker-compose

### 10.2 Build

```bash
git clone https://github.com/to-nexus/go-cross.git
cd go-cross

# Build the client only
make geth

# Build all auxiliary tools (puppeth, abigen, bootnode, devp2p, etc.)
make all
```

Build artifacts are placed under `./build/bin/`.

### 10.3 Docker Image

```bash
docker build -t to-nexus/go-cross:latest -f Dockerfile .
```

---

## 11. Running a Node

### 11.1 Initialize the Data Directory

Built-in networks are selected by name as the argument to `init` (one of `cross`, `zonezero`, `crossdev3`, `crossdev`):

```bash
./build/bin/geth --datadir <YOUR_DATADIR> init cross
```

To use a custom genesis, just pass the path to `genesis.json` instead.

### 11.2 Full / RPC Node

```bash
./build/bin/geth \
  --one \
  --datadir <YOUR_DATADIR> \
  --syncmode full \
  --http --http.addr 0.0.0.0 --http.port 22001 \
  --http.api eth,net,web3,istanbul,cross,debug \
  --ws   --ws.addr   0.0.0.0 --ws.port 32001 \
  --ws.api  eth,net,web3,istanbul,cross \
  --port 30001 --maxpeers 50
```

### 11.3 Archive Node

If you need to retain the full state trie, add:

```bash
  --gcmode archive --state.scheme hash
```

### 11.4 Metrics / Profiling

```bash
  --metrics --metrics.addr 0.0.0.0 --metrics.port 6060 \
  --pprof   --pprof.addr   127.0.0.1 --pprof.port 6061
```

---

## 12. Operating a Validator

### 12.1 Prerequisites

1. **Operator account** — the ECDSA account used to register and manage the validator. **Must have completed KYB (Know-Your-Business) and be registered in the on-chain whitelist.**
2. **Consensus node (validator) account** — the node's `etherbase`, used to sign BFT messages. Recommended to be separate from the operator.
3. **BLS signer key** — mandatory for committing blocks after Breakpoint.
4. Self-delegation or recruited delegations totaling at least **1,000,000 ONE**.

### 12.2 Validator Registration Flow

```text
(1) From the operator wallet:  StakeHub.stake { value: 1_000_000 ONE }
(2) Generate a BLS Proof of Possession:  geth bls generate-proof
(3) From the operator wallet:  StakeHub.applyValidator(
        validatorAddr,   // actual consensus node address
        signerAddr,      // 48-byte BLS public key
        signerProof,     // 96-byte PoP
        id               // 3-9 alphanumeric chars (e.g. "myval01")
    )
(4) Automatically added to the council at the start of the next Council Period
```

### 12.3 Node Flags (Validator)

```bash
./build/bin/geth \
  --one \
  --datadir <YOUR_DATADIR> \
  --syncmode full \
  --mine \
  --miner.etherbase <VALIDATOR_ADDRESS> \
  --unlock <VALIDATOR_ADDRESS> --password <PASSWORD_FILE> \
  --bls.keystore <YOUR_DATADIR>/bls-keystore \
  --bls.key <PATH_TO_BLS_SECRET_KEY_OPTIONAL>
```

- `--mine` is required to participate actively in rounds (you can enable it as a non-validator, but it is meaningless).
- Without a configured BLS secret key, the log will print `Istanbul: BLS secret key is not configured - node cannot participate in PoSA`, and the node cannot take part in consensus.
- `etherbase` **must match exactly** the `validatorAddr` used at registration.

### 12.4 Operational Checklist

- Poll `StakeHub.getValidatorInfo(operator)` regularly via `eth_call`.
- Track cumulative pooled / commission / fee / burned via `RewardHub.getValidatorLog(validatorAddr)`.
- Watch `ValidatorSlash.getSlashInfo(validatorAddr)` to see whether the offline counter is approaching the threshold.
- Reward distribution happens even on empty blocks (≥ 1 per day), so monitor round state via the `istanbul_*` RPC API.

### 12.5 Recovering Active Status

| Situation | Recovery |
|---|---|
| Voluntary `suspend(endTime)` | Automatically rejoins the council in the next Council Period after `endTime` |
| `jail` state | Call `unjail()` once the remaining stake is ≥ 1,000,000 ONE |
| Blacklisted | Keeper removes from the blacklist, then call `unjail()` as above |

---

## 13. BLS Key Management

The `geth bls` subcommand handles the full BLS key lifecycle.
By default, keys are stored as encrypted keystore files under `<DATADIR>/bls-keystore`.

| Command | Description |
|---|---|
| `geth bls list` | Print a summary of stored BLS keys |
| `geth bls new` | Generate a new key pair (interactive password, or `--password`) |
| `geth bls update <pubkey>` | Reset the password |
| `geth bls import <hex-file>` | Import a plaintext hex private key and store it encrypted |
| `geth bls decrypt <pubkey>` | Decrypt a key (for backup) |
| `geth bls generate-proof <operator-address> <pubkey> --chain-id <id>` | Produce the 96-byte PoP required by `applyValidator` |

> **Backup is essential**: keep the entire keystore directory in a safe place. If you lose the password, recovery is impossible.

---

## 14. Sample Configuration File

Usable as `./build/bin/geth --config config.toml`.

```toml
[Eth]
SyncMode = "full"
NetworkId = 612055           # ONE mainnet

[Eth.Miner]
Etherbase = "0x..."          # validator address (only when running as a validator)
GasPrice  = 0

[Node]
DataDir = "/data/cross"
HTTPHost = "0.0.0.0"
HTTPPort = 22001
HTTPCors = ["*"]
HTTPModules = ["net", "web3", "eth", "istanbul", "cross"]
WSHost = "0.0.0.0"
WSPort = 32001
WSOrigins = ["*"]
WSModules = ["net", "web3", "eth", "istanbul", "cross"]

[Node.P2P]
ListenAddr = ":30001"
MaxPeers = 50
NoDiscovery = false
DiscoveryV4 = true
StaticNodes = []

[Node.HTTPTimeouts]
ReadTimeout  = "30s"
WriteTimeout = "30s"
IdleTimeout  = "120s"
```

Command-line flags take precedence over the TOML configuration.

---

## 15. RPC Modules

On top of the standard EL modules (`eth`, `net`, `web3`, `txpool`, `debug`, `admin`, `personal*`), Go Cross adds:

### `istanbul_*`

- `istanbul_getSnapshot(blockNumber)` — validator snapshot (signers and sort policy included)
- `istanbul_getValidators(blockNumberOrTag)` — the active validator set at the given block
- `istanbul_candidates` — council candidates (legacy compatibility)
- `istanbul_isValidator` — whether the local node currently qualifies as a validator

### `cross_*`

- `cross_chainConfig` — ONE-specific metadata such as PoSA activation block and RewardStartBlock
- System-contract indexing / reward-log helpers (see `eth/api*` for the full list)

### Standard EVM Calls

Staking, validator registration, and reward queries are all standard `eth_call` / `eth_sendRawTransaction` flows.
Contract ABIs can be generated from the `cross-contracts` repository (<https://github.com/to-nexus/cross-contracts>, scheduled to become public):

```bash
# from inside cross-contracts
forge soldeer install
scripts/gen-binding StakeHub  breakpoint
scripts/gen-binding RewardHub breakpoint
```

The snippet above is run from inside the `cross-contracts` checkout: `forge soldeer install` fetches the Solidity dependencies (OpenZeppelin, etc.) via Foundry's soldeer manager, and each `scripts/gen-binding <Contract> breakpoint` invocation compiles the contract with `forge`, filters its ABI to a curated function set, then runs `abigen` to emit Go bindings into the `breakpoint` package — those generated files end up under `contracts/breakpoint/` in this repository (e.g. `stake_hub.go`, `reward_hub.go`) and are what the node uses to issue the per-block system transactions described earlier. Re-run these commands whenever the on-chain contracts change, then commit the regenerated `.go` files alongside the Go Cross update.

---

## 16. Upgrade Notes (QBFT → PoSA)

Key differences vs. the previous QBFT-based release:

| Item | QBFT (Before) | PoSA (After) |
|---|---|---|
| Validator eligibility | Fixed whitelist | **KYB (Know-Your-Business) verification + 1 M ONE stake** — anyone qualifies |
| Number of validators | Fixed N | **21-member council** + unlimited candidates |
| Block rewards | None / external distribution | **On-chain distribution (RewardHub)**, automatic 5 % commission |
| Slashing | None | Offline slashing (threshold 50, 10,000 ONE, 2-day jail) |
| Validator refresh | Governance / off-chain coordination | Council rotates **every 24 hours** based on the latest stake ranking |
| Delegators (stakers) | None | Anyone with 5 ONE or more; 14-day unbonding |
| Base-fee handling | (N/A) | Entire EIP-1559 base fee burned (`0x...dEaD`) |
| EL base version | go-ethereum 1.13 | **go-ethereum 1.15** |
| Activation | Genesis | **Breakpoint = `block.timestamp ≥ 1780282800`** (mainnet, 2026-06-01 03:00:00 UTC) |

> Existing data directories can be reused as-is.
> Pre-Breakpoint blocks are still produced by the legacy permissioned validator set, and the validator set is automatically swapped at the first council rotation after Breakpoint.

---

## 17. License / Contributing

- All code under `cmd/` is licensed under the **GNU GPL v3.0** (`COPYING`).
- All other library code is licensed under the **GNU LGPL v3.0** (`COPYING.LESSER`).

Issues and pull requests are welcome.
For security disclosures, please follow the policy in `SECURITY.md`.

General inquiries should be filed as GitHub Issues.
