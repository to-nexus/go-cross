package minerconfig

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params"
)

// Config is the configuration parameters of mining.
type Config struct {
	Etherbase           common.Address `toml:",omitempty"` // Public address for block mining rewards
	PendingFeeRecipient common.Address `toml:"-"`          // Address for pending block rewards.
	ExtraData           hexutil.Bytes  `toml:",omitempty"` // Block extra data set by the miner
	GasFloor            uint64         // Target gas floor for mined blocks.
	GasCeil             uint64         // Target gas ceiling for mined blocks.
	GasPrice            *big.Int       // Minimum gas price for mining a transaction
	Recommit            time.Duration  // The time interval for miner to re-create mining work.

	NewPayloadTimeout time.Duration // The maximum time allowance for creating a new payload
}

// DefaultConfig contains default settings for miner.
var DefaultConfig = Config{
	GasCeil:  30000000,
	GasPrice: big.NewInt(params.GWei), // ##CROSS: min tip

	// The default recommit time is chosen as two seconds since
	// consensus-layer usually will wait a half slot of time(6s)
	// for payload generation. It should be enough for Geth to
	// run 3 rounds.
	Recommit:          10 * time.Second,
	NewPayloadTimeout: 2 * time.Second,
}
