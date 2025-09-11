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
	require.Equal(t, allCodeHash[:], common.Hex2Bytes("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"))
}
