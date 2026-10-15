package ethapi

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/stretchr/testify/require"
)

// ##CROSS: fix upstream 769c8c5 (#35248)
func TestTracer_OnEnter(t *testing.T) {
	for _, tc := range []struct {
		op   vm.OpCode
		want int
	}{
		{vm.CALL, 1},
		{vm.CALLCODE, 0}, // runs the target code on the caller's own balance
		{vm.DELEGATECALL, 0},
	} {
		t.Run(tc.op.String(), func(t *testing.T) {
			tracer := newTracer(true, 0, 0, common.Hash{}, common.Hash{}, 0)
			tracer.onEnter(0, byte(tc.op), common.Address{1}, common.Address{2}, nil, 0, big.NewInt(1))
			require.Len(t, tracer.logs[0], tc.want)
		})
	}
}

// ##
