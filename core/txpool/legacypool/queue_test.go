package legacypool

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
)

// ##CROSS: fix upstream
func TestQueue_Bump(t *testing.T) {
	q := newQueue(DefaultConfig, types.LatestSigner(params.TestChainConfig))
	addr := common.Address{0x01}

	t.Run("ignores an account without queued txs", func(t *testing.T) {
		q.bump(addr)
		assert.NotContains(t, q.beats, addr)
	})

	t.Run("refreshes a queued account", func(t *testing.T) {
		old := time.Now().Add(-time.Minute)
		q.beats[addr] = old
		q.bump(addr)
		assert.True(t, q.beats[addr].After(old))
	})
}

// ##
