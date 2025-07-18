package ethapi

import (
	"context"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
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
	api := NewCrossAPI(newTestBackend(t, genBlocks, genesis, beacon.New(ethash.NewFaker()), func(i int, b *core.BlockGen) {
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
				hexutil.MustDecode("0xd03a06f3f369354d90e2b618173dbc90e86108b73dc44d0ace3675afd03ebcfd"),
				hexutil.MustDecode("0x60bfe5a78e7fc3038c4fafc335a81731f5c602f48eda27d8d588990108a4df32"),
				hexutil.MustDecode("0xe6da2f0464c2ec6e5fb0acb315bbac9b2587dfca4318767c4a077f6972fa67a9"),
				hexutil.MustDecode("0xf3638bbf0b61b2f1b7a7a369c9e65795984534cb24c33ae721e6666dd37e6a8f"),
				hexutil.MustDecode("0xd003763c7d9dad98e026ffa209cc1a2911659721562cc2a8937d31767c27d470"),
				hexutil.MustDecode("0xddf098e7044f701f4a974d9cb58d42e5d41733f088b1fb49f5ac41f192b858c7"),
				hexutil.MustDecode("0x552259cc768c2472fe9275a4e00ebd8ebf18b6c76661acf629f505a632c1d486"),
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
}
