package downloader

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ##CROSS: blob sidecars
// Sidecars must match the blob txs before they enter the result cache,
// and they count against the cache memory budget.
func TestQueue_DeliverBodies(t *testing.T) {
	legacyTx := types.NewTx(&types.LegacyTx{})
	blobTx := types.NewTx(&types.BlobTx{BlobHashes: []common.Hash{{0x01}}})
	sidecar := func(header *types.Header, txIndex uint64, txHash common.Hash, blobs int) *types.BlobSidecar {
		return &types.BlobSidecar{
			BlobTxSidecar: types.BlobTxSidecar{
				Blobs:       make([]kzg4844.Blob, blobs),
				Commitments: make([]kzg4844.Commitment, blobs),
				Proofs:      make([]kzg4844.Proof, blobs),
			},
			BlockNumber: header.Number,
			BlockHash:   header.Hash(),
			TxIndex:     txIndex,
			TxHash:      txHash,
		}
	}

	for _, tc := range []struct {
		name     string
		txs      types.Transactions
		sidecars func(header *types.Header) types.BlobSidecars
		valid    bool
	}{
		{
			name:     "accepts sidecars that match the blob txs",
			txs:      types.Transactions{legacyTx, blobTx},
			sidecars: func(h *types.Header) types.BlobSidecars { return types.BlobSidecars{sidecar(h, 1, blobTx.Hash(), 1)} },
			valid:    true,
		},
		{
			name:     "accepts a pruned body without sidecars",
			txs:      types.Transactions{legacyTx, blobTx},
			sidecars: func(*types.Header) types.BlobSidecars { return nil },
			valid:    true,
		},
		{
			name:     "rejects a sidecar for a block without blob txs",
			txs:      types.Transactions{legacyTx},
			sidecars: func(h *types.Header) types.BlobSidecars { return types.BlobSidecars{sidecar(h, 0, legacyTx.Hash(), 0)} },
		},
		{
			name:     "rejects a sidecar with more blobs than its tx",
			txs:      types.Transactions{legacyTx, blobTx},
			sidecars: func(h *types.Header) types.BlobSidecars { return types.BlobSidecars{sidecar(h, 1, blobTx.Hash(), 64)} },
		},
		{
			name:     "rejects a sidecar for another tx",
			txs:      types.Transactions{legacyTx, blobTx},
			sidecars: func(h *types.Header) types.BlobSidecars { return types.BlobSidecars{sidecar(h, 0, legacyTx.Hash(), 1)} },
		},
		{
			name: "rejects an extra sidecar",
			txs:  types.Transactions{legacyTx, blobTx},
			sidecars: func(h *types.Header) types.BlobSidecars {
				return types.BlobSidecars{sidecar(h, 1, blobTx.Hash(), 1), sidecar(h, 1, blobTx.Hash(), 1)}
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			q := newQueue(1, 1)
			q.Prepare(1, FullSync)
			header := &types.Header{
				Number:     big.NewInt(1),
				Difficulty: big.NewInt(1),
				TxHash:     types.DeriveSha(tc.txs, trie.NewStackTrie(nil)),
				UncleHash:  types.EmptyUncleHash,
			}
			blobGas := uint64(len(blobTx.BlobHashes())) * params.BlobTxBlobGasPerBlob
			if len(tc.txs) == 1 {
				blobGas = 0
			}
			header.BlobGasUsed = &blobGas
			require.Len(t, q.Schedule([]*types.Header{header}, []common.Hash{header.Hash()}, 1), 1)
			peer := dummyPeer("peer")
			req, _, _ := q.ReserveBodies(peer, 1)
			require.NotNil(t, req)

			sidecars := tc.sidecars(header)
			accepted, err := q.DeliverBodies(peer.id, [][]*types.Transaction{tc.txs}, []common.Hash{header.TxHash},
				make([][]*types.Header, 1), []common.Hash{types.EmptyUncleHash},
				make([][]*types.Withdrawal, 1), nil, []types.BlobSidecars{sidecars})
			if !tc.valid {
				require.ErrorIs(t, err, errInvalidBody)
				assert.Zero(t, accepted)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, 1, accepted)

			// The cache budget counts the sidecars of delivered results.
			results := q.Results(false)
			require.Len(t, results, 1)
			assert.Equal(t, sidecars, results[0].Sidecars)
			if len(sidecars) > 0 {
				assert.Greater(t, float64(q.resultSize), blockCacheSizeWeight*float64(sidecars[0].Size()))
			}
		})
	}
}
