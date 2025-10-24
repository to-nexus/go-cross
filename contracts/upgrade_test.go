package contracts

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/params/forks"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ##CROSS: contract upgrade

func TestTryUpdateSystemContract(t *testing.T) {
	forkTime := uint64(time.Now().Unix())
	config := params.AllDevChainProtocolChanges
	config.PragueTime = &forkTime
	config.BreakpointTime = &forkTime

	statedb, _ := state.New(types.EmptyRootHash, state.NewDatabaseForTesting())

	header := &types.Header{
		Number: big.NewInt(100),
		Time:   forkTime + 1,
	}

	upgraded := TryUpdateSystemContract(config, header, forkTime-1, statedb)
	require.True(t, upgraded)

	// Prague upgrade
	t.Run("Prague upgrade", func(t *testing.T) {
		code := statedb.GetCode(params.HistoryStorageAddress)
		assert.Equal(t, params.HistoryStorageCode, code)
		nonce := statedb.GetNonce(params.HistoryStorageAddress)
		assert.Equal(t, uint64(1), nonce)
	})

	t.Run("Breakpoint upgrade", func(t *testing.T) {
		code := statedb.GetCode(ValidatorSetAddr)
		assert.Equal(t, common.FromHex(breakpoint.ValidatorSetMetaData.BinRuntime), code)
		nonce := statedb.GetNonce(ValidatorSetAddr)
		assert.Equal(t, uint64(1), nonce)
	})
}

func TestInitSystemContract(t *testing.T) {
	t.Run("ValidatorSet", func(t *testing.T) {
		// ##CROSS: consensus system contract
		extra := &types.IstanbulExtra{
			VanityData: []byte{},
			Validators: []common.Address{
				common.HexToAddress("0x000000000000000000000000000000000000aaaa"),
				common.HexToAddress("0x000000000000000000000000000000000000bbbb"),
				common.HexToAddress("0x000000000000000000000000000000000000cccc"),
			},
			CommittedSeal: [][]byte{},
			Round:         0,
			Vote:          nil,
			RandomReveal:  []byte{},
		}
		payload, err := rlp.EncodeToBytes(extra)
		require.NoError(t, err)

		beneficiary := common.HexToAddress("0x000000000000000000000000000000000000aaaa")
		config := &params.ChainConfig{
			Istanbul: &params.IstanbulConfig{
				Beneficiary: &beneficiary,
			},
		}
		header := &types.Header{
			Number: big.NewInt(100),
			Time:   uint64(time.Now().Unix()),
			Extra:  payload,
		}

		expected := []ContractInitData{
			{
				To:   IstanbulParamAddr,
				Data: common.FromHex("c4d66de8000000000000000000000000000000000000000000000000000000000000aaaa"),
			},
			{
				To:   ValidatorSetAddr,
				Data: common.FromHex("e71731e400000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000aaaa000000000000000000000000000000000000000000000000000000000000bbbb000000000000000000000000000000000000000000000000000000000000cccc"),
			},
			{
				To:   StakeHubAddr,
				Data: common.FromHex("c4d66de80000000000000000000000000000000000000000000000000000000000000000"),
			},
			{
				To:   GovernanceTokenAddr,
				Data: common.FromHex("8129fc1c"),
			},
		}

		initData := getSystemContractInitialization(upgrades[forks.Breakpoint], config, header)
		assert.Equal(t, expected, initData)
	})
	// ##
}
