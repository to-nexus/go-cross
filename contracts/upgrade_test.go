package contracts

import (
	"crypto/sha256"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// ##CROSS: contract upgrade

func TestAllCodesHash(t *testing.T) {
	allCodes := make([]byte, 0, 10_000_000)
	for _, upgrade := range upgrades {
		if upgrade != nil {
			for _, addressConfig := range upgrade.Configs {
				allCodes = append(allCodes, addressConfig.ContractAddr[:]...)
				allCodes = append(allCodes, addressConfig.Code[:]...)
			}
		}
	}
	allCodeHash := sha256.Sum256(allCodes)
	require.Equal(t, common.Hex2Bytes("f159fb30b48993083f2a2a09ab8a8f55fbad1905815f348607acdc7c0c4645a7"), allCodeHash[:])
}
