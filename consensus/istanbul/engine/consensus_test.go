package engine

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ core.ChainContext = (*mockChainContext)(nil)

var (
	addr1 = common.HexToAddress("0x1111111111111111111111111111111111111111")
	addr2 = common.HexToAddress("0x2222222222222222222222222222222222222222")
	addr3 = common.HexToAddress("0x3333333333333333333333333333333333333333")
	// ##CROSS: bls seal
	signer1 = common.Hex2Bytes("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	signer2 = common.Hex2Bytes("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	signer3 = common.Hex2Bytes("cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc")
	signer4 = common.Hex2Bytes("dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd")
	// ##
)

// mockContractBackend implements bind.ContractBackend for testing
type mockContractBackend struct {
	codeAtBytes   []byte
	codeAtErr     error
	callResponses map[string][][]byte // method selector => responses
	callErr       error
	callCnt       int
}

func (m *mockContractBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return m.codeAtBytes, m.codeAtErr
}
func (m *mockContractBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return m.callContract(call)
}
func (m *mockContractBackend) callContract(call ethereum.CallMsg) ([]byte, error) {
	if m.callResponses != nil && len(call.Data) >= 4 {
		methodSelector := string(call.Data[:4])
		if responses, ok := m.callResponses[methodSelector]; ok && len(responses) > 0 {
			response := responses[0]
			m.callResponses[methodSelector] = responses[1:]
			m.callCnt++
			return response, m.callErr
		}
	}
	return nil, m.callErr
}

func (m *mockContractBackend) CodeAtHash(ctx context.Context, contract common.Address, hash common.Hash) ([]byte, error) {
	return m.codeAtBytes, m.codeAtErr
}
func (m *mockContractBackend) CallContractAtHash(ctx context.Context, call ethereum.CallMsg, hash common.Hash) ([]byte, error) {
	return m.callContract(call)
}
func (m *mockContractBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return &types.Header{}, nil
}
func (m *mockContractBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return m.codeAtBytes, m.codeAtErr
}
func (m *mockContractBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return 0, nil
}
func (m *mockContractBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	return 0, nil
}
func (m *mockContractBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}
func (m *mockContractBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}
func (m *mockContractBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return nil
}
func (m *mockContractBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}
func (m *mockContractBackend) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, nil
}

func TestGetCurrentValidators(t *testing.T) {
	code := []byte{1, 2, 3}
	method := string(breakpoint.NewValidatorSet().PackGetActiveValidators()[:4])
	callContractErr := errors.New("call contract error")

	tests := []struct {
		name               string
		number             uint64
		mockCaller         *mockContractBackend
		expectedValidators []common.Address
		expectedSigners    []types.BLSPublicKey // ##CROSS: bls seal
		expectedErr        error
	}{
		{
			name:   "success with validators",
			number: 100,
			mockCaller: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					method: {packActiveValidatorsResponse(t,
						[]common.Address{addr2, addr1, addr3},
						[][]byte{signer2, signer1, signer3},
					)},
				},
			},
			expectedValidators: []common.Address{addr1, addr2, addr3},
			expectedSigners:    []types.BLSPublicKey{types.BytesToBLSPublicKey(signer1), types.BytesToBLSPublicKey(signer2), types.BytesToBLSPublicKey(signer3)},
			expectedErr:        nil,
		},
		{
			name:   "success with empty validators",
			number: 100,
			mockCaller: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					method: {packActiveValidatorsResponse(t, []common.Address{}, [][]byte{})},
				},
			},
			expectedValidators: []common.Address{},
			expectedSigners:    []types.BLSPublicKey{},
			expectedErr:        nil,
		},
		{
			name:   "contract not deployed",
			number: 100,
			mockCaller: &mockContractBackend{
				codeAtBytes: []byte{}, // No code
			},
			expectedValidators: nil,
			expectedErr:        nil,
		},
		{
			name:   "error from CallContract",
			number: 100,
			mockCaller: &mockContractBackend{
				codeAtBytes: []byte{1, 2, 3},
				callErr:     callContractErr,
			},
			expectedValidators: nil,
			expectedErr:        callContractErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := &Engine{
				contractBackend: tt.mockCaller,
				validatorSet:    breakpoint.NewValidatorSet(),
			}

			validators, signers, err := engine.getCurrentValidators(tt.number)

			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expectedValidators, validators)
			assert.Equal(t, tt.expectedSigners, signers) // ##CROSS: bls seal
		})
	}
}

func packActiveValidatorsResponse(t *testing.T, validators []common.Address, signers [][]byte) []byte {
	t.Helper()

	parsedABI, err := breakpoint.ValidatorSetMetaData.ParseABI()
	require.NoError(t, err)

	method, ok := parsedABI.Methods["getActiveValidators"]
	require.True(t, ok)

	retPacked, err := method.Outputs.Pack(validators, signers)
	require.NoError(t, err)

	return retPacked
}

// mockChainContext implements core.ChainContext and consensus.ChainHeaderReader for testing
type mockChainContext struct {
	config *params.ChainConfig
}

func (m *mockChainContext) Engine() consensus.Engine {
	return nil
}
func (m *mockChainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	return nil
}
func (m *mockChainContext) Config() *params.ChainConfig {
	return m.config
}
func (m *mockChainContext) CurrentHeader() *types.Header {
	return nil
}
func (m *mockChainContext) GetHeaderByNumber(number uint64) *types.Header {
	return nil
}
func (m *mockChainContext) GetHeaderByHash(hash common.Hash) *types.Header {
	return nil
}
func (m *mockChainContext) GetTd(hash common.Hash, number uint64) *big.Int {
	return nil
}

func TestUpdateValidatorSet(t *testing.T) {
	code := []byte{1, 2, 3}
	methodGetValidators := string(breakpoint.NewStakeHub().PackGetValidators(big.NewInt(0), big.NewInt(0))[:4])
	methodValidatorThreshold := string(breakpoint.NewStakeHub().PackValidatorThreshold()[:4])

	tests := []struct {
		name        string
		threshold   uint64
		mockBackend *mockContractBackend
		expectedErr error
		expectTx    bool
	}{
		{
			name:      "success",
			threshold: 2,
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{addr1, addr2, addr3},
						[][]byte{signer1, signer2, signer3},
						[]*big.Int{big.NewInt(1000), big.NewInt(2000), big.NewInt(3000)},
					)},
					methodValidatorThreshold: {packValidatorThresholdResponse(t, 2)},
				},
			},
			expectedErr: nil,
			expectTx:    true,
		},
		{
			name:      "no validator info",
			threshold: 2,
			mockBackend: &mockContractBackend{
				codeAtBytes: []byte{1, 2, 3},
				callResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{},
						[][]byte{},
						[]*big.Int{},
					)},
					methodValidatorThreshold: {packValidatorThresholdResponse(t, 2)},
				},
			},
			expectedErr: nil,
		},
		{
			name:      "zero threshold",
			threshold: 0,
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{addr1},
						[][]byte{signer1},
						[]*big.Int{big.NewInt(1000)},
					)},
					methodValidatorThreshold: {packValidatorThresholdResponse(t, 0)},
				},
			},
			expectedErr: errors.New("zero validator threshold"),
		},
		{
			name:      "error from getValidatorInfo",
			threshold: 2,
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callErr:     errors.New("getValidatorInfo error"),
			},
			expectedErr: errors.New("getValidatorInfo error"),
		},
		{
			name:      "error from getValidatorThreshold",
			threshold: 2,
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{addr1},
						[][]byte{signer1},
						[]*big.Int{big.NewInt(1000)},
					)},
				},
				callErr: errors.New("getValidatorThreshold error"),
			},
			expectedErr: errors.New("getValidatorThreshold error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memdb := rawdb.NewMemoryDatabase()
			triedb := triedb.NewDatabase(memdb, nil)
			sdb := state.NewDatabase(triedb, nil)
			statedb, _ := state.New(types.EmptyRootHash, sdb)

			header := &types.Header{
				Number:     big.NewInt(100),
				ParentHash: common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
				Coinbase:   common.HexToAddress("0x0000000000000000000000000000000000000000"),
				BaseFee:    big.NewInt(0),
				GasLimit:   10000000,
				Difficulty: big.NewInt(0),
				Time:       1000000000,
			}

			chainConfig := params.TestChainConfig
			cx := &mockChainContext{config: chainConfig}

			var txs []*types.Transaction
			var receipts []*types.Receipt
			var usedGas uint64

			engine := &Engine{
				contractBackend: tt.mockBackend,
				validatorSet:    breakpoint.NewValidatorSet(),
				stakeHub:        breakpoint.NewStakeHub(),
				signer:          common.HexToAddress("0x1111111111111111111111111111111111111111"),
				signTx: func(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
					return tx, nil
				},
			}

			err := engine.updateValidatorSet(header, statedb, cx, &txs, &receipts, nil, &usedGas, nil)

			if tt.expectedErr != nil {
				require.ErrorContains(t, err, tt.expectedErr.Error())
				return
			}
			require.NoError(t, err)

			if tt.expectTx {
				assert.NotEmpty(t, txs)
			} else {
				assert.Empty(t, txs)
			}
		})
	}
}

func packGetValidatorsResponse(t *testing.T, validatorAddrs []common.Address, signerAddrs [][]byte, amounts []*big.Int) []byte {
	t.Helper()

	parsedABI, err := breakpoint.StakeHubMetaData.ParseABI()
	require.NoError(t, err)

	method, ok := parsedABI.Methods["getValidators"]
	require.True(t, ok)

	totalLength := big.NewInt(int64(len(validatorAddrs)))
	retPacked, err := method.Outputs.Pack(validatorAddrs, signerAddrs, amounts, totalLength)
	require.NoError(t, err)

	return retPacked
}

func packValidatorThresholdResponse(t *testing.T, threshold uint64) []byte {
	t.Helper()

	parsedABI, err := breakpoint.StakeHubMetaData.ParseABI()
	require.NoError(t, err)

	method, ok := parsedABI.Methods["validatorThreshold"]
	require.True(t, ok)

	retPacked, err := method.Outputs.Pack(big.NewInt(int64(threshold)))
	require.NoError(t, err)

	return retPacked
}

// ##CROSS: istanbul param
func TestSyncIstanbulParam(t *testing.T) {
	code := []byte{1, 2, 3}
	methodGetParamIndex := string(breakpoint.NewIstanbulParam().PackGetParamIndex(100)[:4])
	methodGetParamsByIndex := string(breakpoint.NewIstanbulParam().PackGetParamsByIndex(0)[:4])
	callContractErr := errors.New("call contract error")

	tests := []struct {
		name        string
		header      *types.Header
		mockBackend *mockContractBackend
		precached   int
		expectedErr error
		validate    func(t *testing.T)
	}{
		{
			name:   "sync single param",
			header: &types.Header{Number: big.NewInt(100)},
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					methodGetParamIndex: {packGetParamIndexResponse(t, 0)},
					methodGetParamsByIndex: {packGetParamsByIndexResponse(t, breakpoint.GetParamsByIndexOutput{
						Timepoint:                100,
						CouncilPeriod:            86400,
						BlockPeriod:              5,
						EmptyBlockPeriod:         10,
						RequestTimeout:           3000,
						MaxRequestTimeout:        6000,
						ElasticityMultiplier:     3,
						BaseFeeChangeDenominator: 8,
						MaxBaseFee:               big.NewInt(1e18),
						MinBaseFee:               big.NewInt(1e9),
						ProposerPolicy:           1,
						GasLimit:                 8000000,
					})},
				},
			},
			expectedErr: nil,
			validate: func(t *testing.T) {
				config := params.IstanbulConfigAt(100)
				require.NotNil(t, config)
				require.NotNil(t, config.CouncilPeriod)
				assert.EqualValues(t, 86400, *config.CouncilPeriod)
				assert.EqualValues(t, 5, config.BlockPeriodSeconds)
				assert.EqualValues(t, 10, config.EmptyBlockPeriodSeconds)
				assert.EqualValues(t, 3000, config.RequestTimeoutSeconds)
				require.NotNil(t, config.MaxRequestTimeoutSeconds)
				assert.EqualValues(t, 6000, *config.MaxRequestTimeoutSeconds)
				require.NotNil(t, config.ElasticityMultiplier)
				assert.EqualValues(t, 3, *config.ElasticityMultiplier)
				require.NotNil(t, config.BaseFeeChangeDenominator)
				assert.EqualValues(t, 8, *config.BaseFeeChangeDenominator)
				require.NotNil(t, config.MaxBaseFee)
				assert.EqualValues(t, 1e18, (*big.Int)(config.MaxBaseFee).Uint64())
				require.NotNil(t, config.MinBaseFee)
				assert.EqualValues(t, 1e9, (*big.Int)(config.MinBaseFee).Uint64())
				assert.EqualValues(t, 1, config.ProposerPolicy)
				require.NotNil(t, config.GasLimit)
				assert.EqualValues(t, 8000000, *config.GasLimit)
			},
		},
		{
			name:   "error from getParamIndex",
			header: &types.Header{Number: big.NewInt(100)},
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callErr:     callContractErr,
			},
			expectedErr: callContractErr,
			validate:    func(t *testing.T) {},
		},
		{
			name:   "already synced",
			header: &types.Header{Number: big.NewInt(100)},
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					methodGetParamIndex: {packGetParamIndexResponse(t, 2)},
				},
			},
			precached:   3,
			expectedErr: nil,
			validate:    func(t *testing.T) {},
		},
		{
			name:   "error from getParamsByIndex",
			header: &types.Header{Number: big.NewInt(100)},
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					methodGetParamIndex: {packGetParamIndexResponse(t, 0)},
				},
				callErr: callContractErr,
			},
			expectedErr: callContractErr,
			validate:    func(t *testing.T) {},
		},
		{
			name:   "sync multiple params",
			header: &types.Header{Number: big.NewInt(200)},
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					methodGetParamIndex: {packGetParamIndexResponse(t, 2)},
					methodGetParamsByIndex: {
						packGetParamsByIndexResponse(t, breakpoint.GetParamsByIndexOutput{
							Timepoint:                100,
							CouncilPeriod:            86400,
							BlockPeriod:              5,
							EmptyBlockPeriod:         10,
							RequestTimeout:           3000,
							MaxRequestTimeout:        6000,
							ElasticityMultiplier:     3,
							BaseFeeChangeDenominator: 8,
							MaxBaseFee:               big.NewInt(1e18),
							MinBaseFee:               big.NewInt(1e9),
							ProposerPolicy:           1,
							GasLimit:                 8000000,
						}),
						packGetParamsByIndexResponse(t, breakpoint.GetParamsByIndexOutput{
							Timepoint:                150,
							CouncilPeriod:            86400,
							BlockPeriod:              3,
							EmptyBlockPeriod:         5,
							RequestTimeout:           2000,
							MaxRequestTimeout:        5000,
							ElasticityMultiplier:     2,
							BaseFeeChangeDenominator: 10,
							MaxBaseFee:               big.NewInt(2e18),
							MinBaseFee:               big.NewInt(2e9),
							ProposerPolicy:           0,
							GasLimit:                 10000000,
						}),
						packGetParamsByIndexResponse(t, breakpoint.GetParamsByIndexOutput{
							Timepoint:                200,
							CouncilPeriod:            86400,
							BlockPeriod:              2,
							EmptyBlockPeriod:         3,
							RequestTimeout:           1000,
							MaxRequestTimeout:        4000,
							ElasticityMultiplier:     1,
							BaseFeeChangeDenominator: 12,
							MaxBaseFee:               big.NewInt(3e18),
							MinBaseFee:               big.NewInt(3e9),
							ProposerPolicy:           2,
							GasLimit:                 12000000,
						}),
					},
				},
			},
			expectedErr: nil,
			validate: func(t *testing.T) {
				config0 := params.IstanbulConfigAt(100)
				require.NotNil(t, config0)
				assert.EqualValues(t, 5, config0.BlockPeriodSeconds)
				require.NotNil(t, config0.GasLimit)
				assert.EqualValues(t, 8000000, *config0.GasLimit)

				config1 := params.IstanbulConfigAt(150)
				require.NotNil(t, config1)
				assert.EqualValues(t, 3, config1.BlockPeriodSeconds)
				require.NotNil(t, config1.GasLimit)
				assert.EqualValues(t, 10000000, *config1.GasLimit)

				config2 := params.IstanbulConfigAt(200)
				require.NotNil(t, config2)
				assert.EqualValues(t, 2, config2.BlockPeriodSeconds)
				require.NotNil(t, config2.GasLimit)
				assert.EqualValues(t, 12000000, *config2.GasLimit)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params.ClearCachedIstanbulConfigs()

			engine := &Engine{
				contractBackend: tt.mockBackend,
				istanbulParam:   breakpoint.NewIstanbulParam(),
			}
			expResponse := 0
			for _, responses := range tt.mockBackend.callResponses {
				expResponse += len(responses)
			}
			for i := range tt.precached {
				params.AddIstanbulConfig(&params.IstanbulConfig{}, uint64(i+1))
			}

			err := engine.SyncIstanbulParam(tt.header)

			require.ErrorIs(t, err, tt.expectedErr)

			assert.Equal(t, expResponse, tt.mockBackend.callCnt)
			tt.validate(t)
		})
	}
}

func packGetParamIndexResponse(t *testing.T, index uint64) []byte {
	t.Helper()

	parsedABI, err := breakpoint.IstanbulParamMetaData.ParseABI()
	require.NoError(t, err)

	method, ok := parsedABI.Methods["getParamIndex"]
	require.True(t, ok)

	retPacked, err := method.Outputs.Pack(index)
	require.NoError(t, err)

	return retPacked
}

func packGetParamsByIndexResponse(t *testing.T, result breakpoint.GetParamsByIndexOutput) []byte {
	t.Helper()

	parsedABI, err := breakpoint.IstanbulParamMetaData.ParseABI()
	require.NoError(t, err)

	method, ok := parsedABI.Methods["getParamsByIndex"]
	require.True(t, ok)

	retPacked, err := method.Outputs.Pack(
		result.Timepoint,
		result.EpochLength,
		result.BlockPeriod,
		result.EmptyBlockPeriod,
		result.RequestTimeout,
		result.MaxRequestTimeout,
		result.ElasticityMultiplier,
		result.BaseFeeChangeDenominator,
		result.MaxBaseFee,
		result.MinBaseFee,
		result.ProposerPolicy,
		result.GasLimit,
		result.CouncilPeriod,
	)
	require.NoError(t, err)

	return retPacked
}

// ##CROSS: validator slash
func TestSlashValidators(t *testing.T) {
	validators := []common.Address{
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
	}

	tests := []struct {
		name            string
		setupHeaders    func() map[common.Hash]*types.Header
		proposerPolicy  *istanbul.ProposerPolicy
		expectedSlashed int
	}{
		{
			name: "proposer policy is not RoundRobin",
			setupHeaders: func() map[common.Hash]*types.Header {
				headers := make(map[common.Hash]*types.Header)
				h0 := createHeaderWithIstanbulExtra(t, 98, common.Hash{}, validators[0], validators, 0)
				h1 := createHeaderWithIstanbulExtra(t, 99, h0.Hash(), validators[2], validators, 1)
				h2 := createHeaderWithIstanbulExtra(t, 100, h1.Hash(), validators[0], validators, 0)
				headers[h0.Hash()] = h0
				headers[h1.Hash()] = h1
				headers[h2.Hash()] = h2
				return headers
			},
			proposerPolicy:  istanbul.NewStickyProposerPolicy(),
			expectedSlashed: 0,
		},
		{
			name: "round 0",
			setupHeaders: func() map[common.Hash]*types.Header {
				headers := make(map[common.Hash]*types.Header)
				h0 := createHeaderWithIstanbulExtra(t, 98, common.Hash{}, validators[0], validators, 0)
				h1 := createHeaderWithIstanbulExtra(t, 99, h0.Hash(), validators[1], validators, 0)
				h2 := createHeaderWithIstanbulExtra(t, 100, h1.Hash(), validators[2], validators, 0)
				headers[h0.Hash()] = h0
				headers[h1.Hash()] = h1
				headers[h2.Hash()] = h2
				return headers
			},
			proposerPolicy:  istanbul.NewRoundRobinProposerPolicy(),
			expectedSlashed: 0,
		},
		{
			name: "single slash",
			setupHeaders: func() map[common.Hash]*types.Header {
				headers := make(map[common.Hash]*types.Header)
				h0 := createHeaderWithIstanbulExtra(t, 98, common.Hash{}, validators[0], validators, 0)
				h1 := createHeaderWithIstanbulExtra(t, 99, h0.Hash(), validators[2], validators, 1)
				h2 := createHeaderWithIstanbulExtra(t, 100, h1.Hash(), validators[0], validators, 0)
				headers[h0.Hash()] = h0
				headers[h1.Hash()] = h1
				headers[h2.Hash()] = h2
				return headers
			},
			proposerPolicy:  istanbul.NewRoundRobinProposerPolicy(),
			expectedSlashed: 1,
		},
		{
			name: "multiple slashes",
			setupHeaders: func() map[common.Hash]*types.Header {
				headers := make(map[common.Hash]*types.Header)
				h0 := createHeaderWithIstanbulExtra(t, 98, common.Hash{}, validators[0], validators, 0)
				h1 := createHeaderWithIstanbulExtra(t, 99, h0.Hash(), validators[0], validators, 3)
				h2 := createHeaderWithIstanbulExtra(t, 100, h1.Hash(), validators[1], validators, 0)
				headers[h0.Hash()] = h0
				headers[h1.Hash()] = h1
				headers[h2.Hash()] = h2
				return headers
			},
			proposerPolicy:  istanbul.NewRoundRobinProposerPolicy(),
			expectedSlashed: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memdb := rawdb.NewMemoryDatabase()
			triedb := triedb.NewDatabase(memdb, nil)
			sdb := state.NewDatabase(triedb, nil)
			statedb, _ := state.New(types.EmptyRootHash, sdb)

			mockChain := &mockChainHeaderReader{
				headers: tt.setupHeaders(),
				config:  params.TestChainConfig,
			}
			cx := &chainContext{Chain: mockChain}

			policy := tt.proposerPolicy
			if policy == nil {
				policy = istanbul.NewRoundRobinProposerPolicy()
			}
			cfg := &istanbul.Config{
				ProposerPolicy: policy,
			}

			engine := &Engine{
				cfg:            cfg,
				signer:         validators[0],
				validatorSlash: breakpoint.NewValidatorSlash(),
				signTx: func(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
					return tx, nil
				},
			}

			header := mockChain.GetHeaderByNumber(100)
			require.NotNil(t, header)

			var txs []*types.Transaction
			var receipts []*types.Receipt
			var usedGas uint64

			err := engine.slashValidators(header, statedb, cx, &txs, &receipts, nil, &usedGas, nil)
			require.NoError(t, err)

			assert.Equal(t, tt.expectedSlashed, len(txs))
			for _, tx := range txs {
				assert.Equal(t, contracts.ValidatorSlashAddr, *tx.To())
			}
		})
	}
}

// mockChainHeaderReader implements consensus.ChainHeaderReader for testing
type mockChainHeaderReader struct {
	headers map[common.Hash]*types.Header
	config  *params.ChainConfig
}

func (m *mockChainHeaderReader) Config() *params.ChainConfig {
	return m.config
}
func (m *mockChainHeaderReader) CurrentHeader() *types.Header {
	return nil
}
func (m *mockChainHeaderReader) GetHeader(hash common.Hash, number uint64) *types.Header {
	return m.headers[hash]
}
func (m *mockChainHeaderReader) GetHeaderByNumber(number uint64) *types.Header {
	for _, header := range m.headers {
		if header.Number.Uint64() == number {
			return header
		}
	}
	return nil
}
func (m *mockChainHeaderReader) GetHeaderByHash(hash common.Hash) *types.Header {
	return m.headers[hash]
}
func (m *mockChainHeaderReader) GetTd(hash common.Hash, number uint64) *big.Int {
	return nil
}

// createHeaderWithIstanbulExtra creates a header with Istanbul extra data
func createHeaderWithIstanbulExtra(t *testing.T, number uint64, parentHash common.Hash, coinbase common.Address, validators []common.Address, round uint32) *types.Header {
	t.Helper()

	extra := &types.IstanbulExtra{
		VanityData:    make([]byte, types.IstanbulExtraVanity),
		Validators:    validators,
		Vote:          nil,
		Round:         round,
		CommittedSeal: [][]byte{},
		RandomReveal:  []byte{},
	}

	extraBytes, err := rlp.EncodeToBytes(extra)
	require.NoError(t, err)

	return &types.Header{
		Number:     big.NewInt(int64(number)),
		ParentHash: parentHash,
		Coinbase:   coinbase,
		Extra:      extraBytes,
		BaseFee:    big.NewInt(0),
		GasLimit:   10000000,
		Difficulty: big.NewInt(0),
	}
}

// ##

// ##CROSS: validator reward
func TestDistributeRewards(t *testing.T) {
	tests := []struct {
		name        string
		header      *types.Header
		txs         []*types.Transaction
		receipts    []*types.Receipt
		expectedFee *big.Int
	}{
		{
			name: "no transactions",
			header: &types.Header{
				Number:     big.NewInt(100),
				ParentHash: common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
				BaseFee:    big.NewInt(1000000000),
				GasLimit:   10000000,
				Difficulty: big.NewInt(0),
				Time:       1000000000,
			},
			txs:         []*types.Transaction{},
			receipts:    []*types.Receipt{},
			expectedFee: big.NewInt(0),
		},
		{
			name: "success with single tx",
			header: &types.Header{
				Number:     big.NewInt(100),
				ParentHash: common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
				BaseFee:    big.NewInt(1000000000),
				GasLimit:   10000000,
				Difficulty: big.NewInt(0),
				Time:       1000000000,
			},
			txs: []*types.Transaction{
				types.NewTx(&types.DynamicFeeTx{
					Nonce:     0,
					GasTipCap: big.NewInt(1000000000),
					GasFeeCap: big.NewInt(2000000000),
					Gas:       21000,
					To:        ptrAddr(common.HexToAddress("0x2222222222222222222222222222222222222222")),
					Value:     big.NewInt(0),
				}),
			},
			receipts: []*types.Receipt{
				{
					Status:  types.ReceiptStatusSuccessful,
					GasUsed: 21000,
				},
			},
			expectedFee: new(big.Int).Mul(big.NewInt(21000), big.NewInt(2000000000)), // 21000 * (1+1) Gwei
		},
		{
			name: "success with multiple txs",
			header: &types.Header{
				Number:     big.NewInt(100),
				ParentHash: common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
				BaseFee:    big.NewInt(1000000000),
				GasLimit:   10000000,
				Difficulty: big.NewInt(0),
				Time:       1000000000,
			},
			txs: []*types.Transaction{
				types.NewTx(&types.DynamicFeeTx{
					Nonce:     0,
					GasTipCap: big.NewInt(1000000000),
					GasFeeCap: big.NewInt(2000000000),
					Gas:       21000,
					To:        ptrAddr(common.HexToAddress("0x2222222222222222222222222222222222222222")),
					Value:     big.NewInt(0),
				}),
				types.NewTx(&types.DynamicFeeTx{
					Nonce:     1,
					GasTipCap: big.NewInt(2000000000),
					GasFeeCap: big.NewInt(3000000000),
					Gas:       42000,
					To:        ptrAddr(common.HexToAddress("0x3333333333333333333333333333333333333333")),
					Value:     big.NewInt(0),
				}),
			},
			receipts: []*types.Receipt{
				{
					Status:  types.ReceiptStatusSuccessful,
					GasUsed: 21000,
				},
				{
					Status:  types.ReceiptStatusSuccessful,
					GasUsed: 42000,
				},
			},
			expectedFee: new(big.Int).Add(
				new(big.Int).Mul(big.NewInt(21000), big.NewInt(2000000000)),
				new(big.Int).Mul(big.NewInt(42000), big.NewInt(3000000000)),
			), // 21000 * (1+1) Gwei + 42000 * (1+2) Gwei
		},
		{
			name: "legacy tx",
			header: &types.Header{
				Number:     big.NewInt(100),
				ParentHash: common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
				BaseFee:    big.NewInt(1000000000),
				GasLimit:   10000000,
				Difficulty: big.NewInt(0),
				Time:       1000000000,
			},
			txs: []*types.Transaction{
				types.NewTx(&types.LegacyTx{
					Nonce:    0,
					GasPrice: big.NewInt(2000000000), // 2 Gwei
					Gas:      21000,
					To:       ptrAddr(common.HexToAddress("0x2222222222222222222222222222222222222222")),
					Value:    big.NewInt(0),
				}),
			},
			receipts: []*types.Receipt{
				{
					Status:  types.ReceiptStatusSuccessful,
					GasUsed: 21000,
				},
			},
			expectedFee: new(big.Int).Mul(big.NewInt(21000), big.NewInt(2000000000)), // 21000 * 2 Gwei
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memdb := rawdb.NewMemoryDatabase()
			triedb := triedb.NewDatabase(memdb, nil)
			sdb := state.NewDatabase(triedb, nil)
			statedb, _ := state.New(types.EmptyRootHash, sdb)
			cx := &mockChainContext{config: params.TestChainConfig}

			txs := make([]*types.Transaction, len(tt.txs))
			copy(txs, tt.txs)
			receipts := make([]*types.Receipt, len(tt.receipts))
			copy(receipts, tt.receipts)
			var usedGas uint64

			signer := common.HexToAddress("0x1111111111111111111111111111111111111111")
			engine := &Engine{
				contractBackend: &mockContractBackend{codeAtBytes: []byte{1, 2, 3}},
				stakeHub:        breakpoint.NewStakeHub(),
				signer:          signer,
				signTx: func(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
					require.Equal(t, account.Address, signer)
					return tx, nil
				},
			}

			err := engine.distributeRewards(tt.header, statedb, cx, &txs, &receipts, nil, &usedGas, nil)

			require.NoError(t, err)
			if len(tt.txs) > 0 {
				// 1 system tx should be added
				assert.Equal(t, len(tt.txs)+1, len(txs))
				tx := txs[len(txs)-1]

				assert.Equal(t, contracts.StakeHubAddr, *tx.To())
				// verify fee
				assert.Equal(t, tt.expectedFee.Uint64(), tx.Value().Uint64())
			}
			assert.Equal(t, len(txs), len(receipts))
		})
	}
}

func ptrAddr(addr common.Address) *common.Address {
	return &addr
}

// ##
