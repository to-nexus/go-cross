package core

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
)

func TestSealIndexes(t *testing.T) {
	addrs := []common.Address{
		common.HexToAddress("0xc53f2189bf6d7bf56722731787127f90d319e112"),
		common.HexToAddress("0xed2d479591fe2c5626ce09bca4ed2a62e00e5bc2"),
		common.HexToAddress("0xc8417f834995aaeb35f342a67a4961e19cd4735c"),
		common.HexToAddress("0x784ae51f5013b51c8360afdf91c6bc5a16f586ea"),
		common.HexToAddress("0xecf0974e6f0630fd91ea4da8399cdb3f59e5220f"),
		common.HexToAddress("0x411c4d11acd714b82a5242667e36de14b9e1d10b"),
	}
	stringPolicy := istanbul.NewProposerPolicyByIdAndSortFunc(istanbul.RoundRobin, istanbul.ValidatorSortByString())
	valSet := validator.NewSet(addrs, nil, stringPolicy)
	if valSet.List()[3].Address() != addrs[1] {
		t.Fatalf("test fixture no longer differs between string and byte order")
	}

	indexes := sealIndexes(valSet)

	expected := map[common.Address]uint{
		addrs[5]: 0,
		addrs[3]: 1,
		addrs[0]: 2,
		addrs[2]: 3,
		addrs[4]: 4,
		addrs[1]: 5,
	}
	for addr, index := range expected {
		if indexes[addr] != index {
			t.Fatalf("unexpected canonical index for %s: have %d, want %d", addr, indexes[addr], index)
		}
	}
}
