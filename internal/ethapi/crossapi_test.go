package ethapi

import (
	"bytes"
	"context"
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCrossAPI_GetCheckpointProof(t *testing.T) {
	t.Parallel()

	var (
		genBlocks = 100
		genesis   = &core.Genesis{
			Config: params.MergedTestChainConfig,
			Alloc:  types.GenesisAlloc{},
		}
	)
	api := NewCrossAPI(newTestBackend(t, genBlocks, genesis, beacon.NewFaker(), func(i int, b *core.BlockGen) {
		b.SetPoS()
	}))

	tests := []struct {
		blockNumber rpc.BlockNumber
		startNumber rpc.BlockNumber
		endNumber   rpc.BlockNumber
		expectErr   string
		want        []hexutil.Bytes
	}{
		{
			blockNumber: rpc.BlockNumber(genBlocks / 2),
			startNumber: rpc.BlockNumber(0),
			endNumber:   rpc.BlockNumber(genBlocks - 1),
			expectErr:   "",
			want: []hexutil.Bytes{
				hexutil.MustDecode("0xb00f2343bc11fc5af8797215d15bccd5ba8a7f9b481a4b8ebdee8ba040806052"),
				hexutil.MustDecode("0xf046032fa0163ee498bc37dc6fab720f9194d3f7b166129938e0fd24e892c956"),
				hexutil.MustDecode("0x0a101fbdb6781f2c65fafded8de7f07f92b1d36dff828cd2246b406e870ef413"),
				hexutil.MustDecode("0xa6c19ea6be764111e82440c85f70e819e8e171079eef88230205834f66237883"),
				hexutil.MustDecode("0x80e6b653aea327e0978adad5df333572c9c5c1b2907ab3342c2a03c319e45746"),
				hexutil.MustDecode("0x71eb085c7899c18f0514977e5cd6fa314f573c916c841dee35594371c74967ac"),
				hexutil.MustDecode("0x48141c3f48a5891a961c551be88b6f6db131058c46e4899ef3ef666bc9459512"),
			},
		},
		{
			blockNumber: rpc.BlockNumber(-1),
			startNumber: rpc.BlockNumber(0),
			endNumber:   rpc.BlockNumber(genBlocks - 1),
			expectErr:   "invalid block number",
			want:        nil,
		},
		{
			blockNumber: rpc.BlockNumber(genBlocks),
			startNumber: rpc.BlockNumber(0),
			endNumber:   rpc.BlockNumber(genBlocks / 2),
			expectErr:   "invalid range",
			want:        nil,
		},
		{
			blockNumber: rpc.BlockNumber(0),
			startNumber: rpc.BlockNumber(0),
			endNumber:   rpc.BlockNumber(100000),
			expectErr:   "range too large",
			want:        nil,
		},
		{
			// Blocks above the head must not be hashed as nil headers.
			blockNumber: rpc.BlockNumber(genBlocks),
			startNumber: rpc.BlockNumber(genBlocks - 1),
			endNumber:   rpc.BlockNumber(genBlocks + 5),
			expectErr:   "header 101 not found",
			want:        nil,
		},
		{
			// A range ending at MaxInt64 must not loop forever.
			blockNumber: rpc.BlockNumber(math.MaxInt64 - 1),
			startNumber: rpc.BlockNumber(math.MaxInt64 - 1),
			endNumber:   rpc.BlockNumber(math.MaxInt64),
			expectErr:   "header 9223372036854775806 not found",
			want:        nil,
		},
		{
			blockNumber: rpc.BlockNumber(0),
			startNumber: rpc.BlockNumber(0),
			endNumber:   rpc.BlockNumber(math.MaxInt64),
			expectErr:   "range too large",
			want:        nil,
		},
	}

	for i, tc := range tests {
		proof, err := api.GetCheckpointProof(context.Background(), tc.blockNumber, tc.startNumber, tc.endNumber)
		if tc.expectErr != "" {
			if err == nil {
				t.Errorf("test %d: want error %v, have nothing", i, tc.expectErr)
				continue
			}
			if err.Error() != tc.expectErr {
				t.Errorf("test %d: error mismatch, want %v, have %v", i, tc.expectErr, err)
			}
			continue
		}
		if err != nil {
			t.Errorf("test %d: want no error, have %v", i, err)
			continue
		}
		if len(tc.want) > 0 {
			if !reflect.DeepEqual(tc.want, proof) {
				t.Errorf("test %d: proof mismatch, want %v, have %v", i, tc.want, proof)
			}
		}
	}

	t.Run("leaf and proof match the CrossCheck contract", func(t *testing.T) {
		leaf := checkpointLeaf(1000000, common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111"))
		require.Equal(t, "0xaa040d2e337b40490365f32debbf80fd48474d6750fc750de6e0e5cd01357723", hexutil.Encode(leaf))

		// Blocks 100..104.
		leaves := make([][]byte, 0, 5)
		for i := range 5 {
			leaves = append(leaves, checkpointLeaf(uint64(100+i), common.BytesToHash(bytes.Repeat([]byte{byte(i + 1)}, 32))))
		}
		tree, err := generateMerkleTree(leaves)
		require.NoError(t, err)
		assert.Equal(t, "0x2840dff574efea4e4fd15d64a1954654b4a2afede395bbfbc40a9c5a083d24ed", hexutil.Encode(tree.Root().Hash))

		proof, err := generateMerkleProof(tree, 2) // block 102
		require.NoError(t, err)
		assert.Equal(t, []hexutil.Bytes{
			hexutil.MustDecode("0x3f17efbd3067581b2213b27d5f8def82471e2741e531136bf91f7af939f21f83"),
			hexutil.MustDecode("0x586c4ebfc18d6ff34a06f7461e78a69ea31f2b28fcd970c4ffa59aa54c9005cd"),
			hexutil.MustDecode("0x232e601b094e77857dd0ed5d4e4a491fd1679cc786724f7676cb872bd4e482f9"),
		}, proof)

		proof, err = generateMerkleProof(tree, 4) // block 104
		require.NoError(t, err)
		assert.Equal(t, []hexutil.Bytes{
			hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
			hexutil.MustDecode("0x098345d67c99d856cac50938e331f76af6626e26e5cbbe14ec5c75137d70d633"),
		}, proof)
	})

	t.Run("range ending at MaxInt64 stops at the end", func(t *testing.T) {
		// The backend returns a header for every number, so only the loop bound can stop it.
		backend := &anyHeaderBackend{}
		start := rpc.BlockNumber(math.MaxInt64 - 1)
		_, err := NewCrossAPI(backend).GetCheckpointProof(context.Background(), start, start, math.MaxInt64)
		if err != nil {
			t.Fatalf("want no error, have %v", err)
		}
		want := []rpc.BlockNumber{math.MaxInt64 - 1, math.MaxInt64}
		if !reflect.DeepEqual(want, backend.calls) {
			t.Fatalf("calls mismatch, want %v, have %v", want, backend.calls)
		}
	})
}

// anyHeaderBackend returns an empty header for any block number and records the requested numbers.
type anyHeaderBackend struct {
	Backend
	calls []rpc.BlockNumber
}

func (b *anyHeaderBackend) HeaderByNumber(_ context.Context, number rpc.BlockNumber) (*types.Header, error) {
	b.calls = append(b.calls, number)
	if len(b.calls) > 2 {
		return nil, errors.New("too many calls")
	}
	return &types.Header{}, nil
}
