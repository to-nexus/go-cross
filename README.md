# 🚀 Go Cross

Go Cross is a **blockchain client based on Ethereum v1.14.13**, integrating ConsenSys Quorum's latest BFT consensus algorithm. It features **Adventure Fork, a custom fork that incorporates elements from the Shanghai and Cancun forks**.

## 🔢 Version Information

- **Execution Layer: Ethereum**: v1.14.13
- **Consensus Layer: ConsenSys Quorum**: Latest BFT-based consensus algorithm
- **Fork Version**: Adventure Fork (Shanghai + partial Cancun features applied)

## 🛠 Installation and Execution Guide

### **1️⃣ Prerequisites**

- **Go 1.20 or higher**
- **Git**
- **Docker (Optional for containerized deployment)**

### **2️⃣ Installation and Execution Steps**

#### **① Clone the Repository**

```bash
git clone https://github.com/to-nexus/go-cross.git
cd go-cross
```

#### **② Build the Client**

```bash
make geth
```

#### **③ Initialize a New Blockchain**

```bash
./build/bin/geth --datadir <your-datadir> init cross
```

#### **④ Start the Node**

```bash
./build/bin/geth --config config.toml
```

## ⚙️ Sample `config.toml`

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

The go-cross library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The go-cross binaries (i.e. all code inside of the `cmd` directory) are licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.

## 🤝 Contributing

Go Cross is an open-source project, and contributions are welcome! If you have improvements or fixes, feel free to **open an issue or submit a pull request**.

## 📞 Contact & Support

For questions or collaboration inquiries, please reach out via GitHub Issues. 🚀

