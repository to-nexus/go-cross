package ethapi

import (
	"context"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/beacon"
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
				hexutil.MustDecode("0x23b77de72d3b74b419b297b97447149eb2686fa78e9d2422e0d6413221ab8e62"),
				hexutil.MustDecode("0x59636c033d8bcdfc74002797ce17bca6f15ea598fe376ce33db07b5a54b68679"),
				hexutil.MustDecode("0xecbd8162c0309bed4ca4acf70faf2311f77451f8da2a889f48c871bb9c7c6dc1"),
				hexutil.MustDecode("0xebc2c9d482409851b56360a2fa323a11a48972a0dabdf26483b682f2f9495790"),
				hexutil.MustDecode("0xacd5e8d210f20f60de45ec0911cbe8cc18774e671db64986291540fb2cdb1281"),
				hexutil.MustDecode("0x1bf83590e26f20453618409f24445cea1cc327335bf368752436efe60ca81437"),
				hexutil.MustDecode("0x8c0d0ddf20af5c3476155157ad7b286654dd12434a62f22545efe6425b1be350"),
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
