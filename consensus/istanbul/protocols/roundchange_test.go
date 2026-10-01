package protocols

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ##CROSS: blob sidecars
func TestRoundChange_EncodeRLP(t *testing.T) {
	block := types.NewBlockWithHeader(&types.Header{Number: big.NewInt(1)})
	sidecars := types.BlobSidecars{{BlockNumber: big.NewInt(1), BlockHash: block.Hash(), TxHash: common.Hash{0x01}}}

	t.Run("carries prepared block sidecars", func(t *testing.T) {
		payload, err := rlp.EncodeToBytes(NewRoundChange(big.NewInt(1), big.NewInt(2), big.NewInt(1), block.WithSidecars(sidecars)))
		require.NoError(t, err)

		msg, err := Decode(RoundChangeCode, payload)
		require.NoError(t, err)
		decoded := msg.(*RoundChange)
		assert.Equal(t, block.Hash(), decoded.PreparedBlock.Hash())
		require.Len(t, decoded.PreparedBlock.Sidecars(), 1)
		assert.Equal(t, sidecars[0].TxHash, decoded.PreparedBlock.Sidecars()[0].TxHash)
	})

	t.Run("keeps the old encoding without sidecars", func(t *testing.T) {
		payload, err := rlp.EncodeToBytes(NewRoundChange(big.NewInt(1), big.NewInt(2), big.NewInt(1), block))
		require.NoError(t, err)

		// Old nodes expect exactly [signed payload, prepared block, justification].
		var fields []rlp.RawValue
		require.NoError(t, rlp.DecodeBytes(payload, &fields))
		assert.Len(t, fields, 3)

		msg, err := Decode(RoundChangeCode, payload)
		require.NoError(t, err)
		assert.Empty(t, msg.(*RoundChange).PreparedBlock.Sidecars())
	})
}

// ##
