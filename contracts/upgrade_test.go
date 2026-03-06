package contracts

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/contracts/erc8004"
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
	header := &types.Header{
		Number: big.NewInt(100),
		Time:   forkTime + 1,
	}

	t.Run("Testnet", func(t *testing.T) {
		statedb, _ := state.New(types.EmptyRootHash, state.NewDatabaseForTesting())
		config := params.AllDevChainProtocolChanges
		config.PragueTime = &forkTime
		config.BreakpointTime = &forkTime

		upgraded := TryUpdateSystemContract(config, header, forkTime-1, statedb)
		require.True(t, upgraded)

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

		// ##CROSS: erc8004
		t.Run("ERC-8004 upgrade", func(t *testing.T) {
			code := statedb.GetCode(ERC8004IdentityRegistryAddrTestnet)
			assert.Equal(t, common.FromHex(erc8004.IdentityRegistryCode), code)
			nonce := statedb.GetNonce(ERC8004IdentityRegistryAddrTestnet)
			assert.Equal(t, uint64(1), nonce)

			// because chain ID is 1337, contracts are deployed to testnet addresses; mainnet addresses are empty
			code = statedb.GetCode(ERC8004ReputationRegistryAddr)
			assert.Empty(t, code)
		})
		// ##
	})

	t.Run("Mainnet", func(t *testing.T) {
		statedb, _ := state.New(types.EmptyRootHash, state.NewDatabaseForTesting())
		config := params.CrossChainConfig
		config.PragueTime = &forkTime
		config.BreakpointTime = &forkTime

		upgraded := TryUpdateSystemContract(config, header, forkTime-1, statedb)
		require.True(t, upgraded)

		// ##CROSS: erc8004
		t.Run("ERC-8004 upgrade", func(t *testing.T) {
			code := statedb.GetCode(ERC8004IdentityRegistryAddr)
			assert.Equal(t, common.FromHex(erc8004.IdentityRegistryCode), code)
			nonce := statedb.GetNonce(ERC8004IdentityRegistryAddr)
			assert.Equal(t, uint64(1), nonce)

			code = statedb.GetCode(ERC8004ReputationRegistryAddrTestnet)
			assert.Empty(t, code)
		})
		// ##
	})
}

// ##CROSS: consensus system contract
func TestInitSystemContract(t *testing.T) {
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
	// ##CROSS: bls seal
	signers := make([][]byte, 0, len(extra.Validators))
	for range len(extra.Validators) {
		signers = append(signers, types.BLSPublicKey{}.Bytes())
	}
	// ##
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

	initialize := common.FromHex("8129fc1c")

	tests := []struct {
		name     string
		fork     forks.Fork
		header   *types.Header
		expected []ContractInitData
	}{
		{
			name: "Breakpoint",
			fork: forks.Breakpoint,
			header: &types.Header{
				Number:   big.NewInt(100),
				Time:     uint64(time.Now().Unix()),
				GasLimit: 2e10,
				Extra:    payload,
			},
			expected: []ContractInitData{
				{
					To: IstanbulParamAddr,
					Data: breakpoint.NewIstanbulParam().PackInitialize(
						300,                               // epochLength
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
						86400,                             // councilPeriod
					),
				},
				{
					To:   ValidatorSetAddr,
					Data: breakpoint.NewValidatorSet().PackUpdateValidators(extra.Validators, signers),
				},
				{
					To:   StakeHubAddr,
					Data: initialize,
				},
				{
					To:   ValidatorSlashAddr,
					Data: initialize,
				},
				{
					To:   GovernorAddr,
					Data: initialize,
				},
				{
					To:   GovernanceTokenAddr,
					Data: initialize,
				},
				{
					To:   GovernanceTimelockAddr,
					Data: initialize,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initData := getSystemContractInitialization(upgrades[tt.fork][0], config, tt.header)
			assert.Equal(t, tt.expected, initData)
		})
	}
}

// ##
