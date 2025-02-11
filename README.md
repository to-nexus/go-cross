# Go Cross

## Overview

This project is a customized blockchain client integrating an improved **Byzantine Fault Tolerant (BFT) consensus algorithm**. The modifications enhance transaction processing efficiency and usability by enabling validator-based consensus.

## Versions

- **Ethereum**: 1.13.15
- **ConsenSys Quorum**: v24.4.1

## Consensus Algorithm: Byzantine Fault Tolerant (BFT) Protocols

This blockchain client implements an improved version of the IBFT (Istanbul Byzantine Fault Tolerance) consensus mechanism, known as QBFT, addressing several limitations and enhancing overall network performance. Key improvements include:

- **Enhanced Scalability**: The updated algorithm reduces message complexity, allowing for efficient operation with a larger validator set.
- **Lower Network Overhead**: Optimized block proposal mechanisms minimize unnecessary message exchanges, improving network efficiency.
- **Faster Finality**: The new consensus method reaches finality with lower latency, especially in networks with a higher number of participating nodes.
- **Improved Fault Tolerance**: Better handling of faulty or malicious nodes ensures greater network stability and security.
- **Optimized Validator Rotation**: A more efficient round-robin mechanism prevents congestion and improves fairness in validator selection.

### Consensus Process

The consensus process follows these key steps:

1. **Block Proposal**: A designated validator proposes a new block to be added to the chain.
2. **Pre-Prepare Phase**: Other validators receive the proposed block and verify its validity.
3. **Prepare Phase**: Validators broadcast their agreement on the proposed block.
4. **Commit Phase**: If a sufficient number of validators agree, they finalize the block and append it to the blockchain.
5. **New Round**: The next validator in the round-robin schedule is selected to propose the next block.

These improvements make the updated BFT protocol more suitable for decentralized blockchain networks requiring high performance and resilience.

## Installation Guide

To set up and run the customized client:

### Prerequisites

- Go (1.18 or higher)
- Git
- Docker (optional, for containerized deployment)

### Steps

1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo-name.git
   cd your-repo-name
   ```
2. Build the client:
   ```sh
   make geth
   ```
3. Initialize a new blockchain:
   ```sh
   ./build/bin/geth --datadir <your-datadir> init <genesis.json>
   ```
4. Start the node:
   ```sh
   ./build/bin/geth --config config.toml
   ```

## Sample `config.toml`

Below is a sample `config.toml` file for running the node:

```toml
[Eth]
SyncMode = "full" 

[Node]
DataDir = <your-datadir>
HTTPHost = "0.0.0.0"
HTTPPort = 22001
HTTPCors = ["*"]
HTTPModules = ["net", "web3", "eth", "istanbul"]
WSHost = "0.0.0.0"
WSPort = 32001
WSOrigins = ["*"]
WSModules = ["net", "web3", "eth", "istanbul"]

[Node.P2P]
ListenAddr = ":30001" 
MaxPeers = 50
NoDiscovery = false
aDiscoveryV4 = false
StaticNodes = []

```

## License

This project is released under the **MIT License**. See the [LICENSE](./LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you'd like to improve the project.

## Contact

For questions or collaboration inquiries, reach out via [GitHub Issues](https://github.com/your-repo-name/issues).

