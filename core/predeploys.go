package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	ERC1967ProxyImplementationSlot = common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")
)

const (
	Bridge     = "0xB81D6E000000000000000000000000000000FACE"
	BridgeImpl = "0xB81D6E000000000000000000000000000000C0DE"
)

var (
	BridgeAddr     = common.HexToAddress(Bridge)
	BridgeImplAddr = common.HexToAddress(BridgeImpl)

	Predeploys = map[common.Address]types.Account{}
)

func init() {
	Predeploys[BridgeAddr] = types.Account{
		Balance: new(big.Int).Mul(big.NewInt(100_000_000_000), big.NewInt(1e18)),
		Code:    ERC1967ProxyCode,
		Storage: map[common.Hash]common.Hash{
			ERC1967ProxyImplementationSlot: common.BytesToHash(BridgeImplAddr.Bytes()),
		},
	}

	Predeploys[BridgeImplAddr] = types.Account{
		Code: BridgeCode,
	}
}
