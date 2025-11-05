package contracts

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
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

		maxBaseFee := math.NewHexOrDecimal256(1e18)
		minBaseFee := math.NewHexOrDecimal256(1e8)
		config := &params.ChainConfig{
			Istanbul: &params.IstanbulConfig{
				EpochLength:             86400,
				BlockPeriodSeconds:      1,
				EmptyBlockPeriodSeconds: 0,
				RequestTimeoutSeconds:   10,
				ProposerPolicy:          0,
				MaxBaseFee:              maxBaseFee,
				MinBaseFee:              minBaseFee,
			},
		}
		header := &types.Header{
			Number:   big.NewInt(100),
			Time:     uint64(time.Now().Unix()),
			GasLimit: 2e10,
			Extra:    payload,
		}

		expected := []ContractInitData{
			{
				To: IstanbulParamAddr,
				Data: func() []byte {
					// pack initialize data for IstanbulParam
					return breakpoint.NewIstanbulParam().PackInitialize(
						86400,                             // epochLength
						1,                                 // blockPeriod
						0,                                 // emptyBlockPeriod
						10,                                // requestTimeout
						0,                                 // maxRequestTimeout
						config.ElasticityMultiplier(),     // elasticityMultiplier
						config.BaseFeeChangeDenominator(), // baseFeeChangeDenominator
						(*big.Int)(maxBaseFee),            // maxBaseFee
						(*big.Int)(minBaseFee),            // minBaseFee
						0,                                 // proposerPolicy
						2e10,                              // gasLimit
					)
				}(),
			},
			{
				To: ValidatorSetAddr,
				Data: func() []byte {
					return breakpoint.NewValidatorSet().PackUpdateValidators(extra.Validators)
				}(),
			},
			{
				To:   StakeHubAddr,
				Data: common.FromHex("c4d66de80000000000000000000000000000000000000000000000000000000000000000"), // initialize(address(0))
			},
			{
				To:   GovernorAddr,
				Data: common.FromHex("c4d66de80000000000000000000000000000000000000000000000000000000000000000"), // initialize(address(0))
			},
			{
				To:   GovernanceTokenAddr,
				Data: common.FromHex("8129fc1c"), // initialize()
			},
			{
				To:   GovernanceTimelockAddr,
				Data: common.FromHex("8129fc1c"), // initialize()
			},
		}

		initData := getSystemContractInitialization(upgrades[forks.Breakpoint], config, header)
		assert.Equal(t, expected, initData)
	})
	// ##
}
