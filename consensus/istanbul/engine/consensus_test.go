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
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ core.ChainContext = (*mockChainContext)(nil)

// mockContractBackend implements bind.ContractBackend for testing
type mockContractBackend struct {
	codeAtBytes           []byte
	codeAtErr             error
	callContractResponses map[string][][]byte // method selector => responses
	callContractErr       error
}

func (m *mockContractBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return m.codeAtBytes, m.codeAtErr
}
func (m *mockContractBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return m.callContract(call)
}
func (m *mockContractBackend) callContract(call ethereum.CallMsg) ([]byte, error) {
	if m.callContractResponses != nil && len(call.Data) >= 4 {
		methodSelector := string(call.Data[:4])
		if responses, ok := m.callContractResponses[methodSelector]; ok && len(responses) > 0 {
			response := responses[0]
			m.callContractResponses[methodSelector] = responses[1:]
			return response, m.callContractErr
		}
	}
	return nil, m.callContractErr
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
	method := string(breakpoint.NewValidatorSet().PackGetValidators()[:4])
	callContractErr := errors.New("call contract error")

	tests := []struct {
		name               string
		number             uint64
		mockCaller         *mockContractBackend
		expectedValidators []common.Address
		expectedErr        error
	}{
		{
			name:   "success with validators",
			number: 100,
			mockCaller: &mockContractBackend{
				codeAtBytes: code,
				callContractResponses: map[string][][]byte{
					method: {packValidatorsResponse(t, []common.Address{
						common.HexToAddress("0x2222222222222222222222222222222222222222"),
						common.HexToAddress("0x1111111111111111111111111111111111111111"),
						common.HexToAddress("0x3333333333333333333333333333333333333333"),
					})},
				},
			},
			expectedValidators: []common.Address{
				common.HexToAddress("0x1111111111111111111111111111111111111111"),
				common.HexToAddress("0x2222222222222222222222222222222222222222"),
				common.HexToAddress("0x3333333333333333333333333333333333333333"),
			},
			expectedErr: nil,
		},
		{
			name:   "success with empty validators",
			number: 100,
			mockCaller: &mockContractBackend{
				codeAtBytes: code,
				callContractResponses: map[string][][]byte{
					method: {packValidatorsResponse(t, []common.Address{})},
				},
			},
			expectedValidators: []common.Address{},
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
				codeAtBytes:     []byte{1, 2, 3},
				callContractErr: callContractErr,
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

			validators, err := engine.getCurrentValidators(tt.number)

			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expectedValidators, validators)
		})
	}
}

func packValidatorsResponse(t *testing.T, validators []common.Address) []byte {
	t.Helper()

	parsedABI, err := breakpoint.ValidatorSetMetaData.ParseABI()
	require.NoError(t, err)

	method, ok := parsedABI.Methods["getValidators"]
	require.True(t, ok)

	retPacked, err := method.Outputs.Pack(validators)
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
		name           string
		validatorInfos []struct {
			addr   common.Address
			amount int64
		}
		threshold   uint64
		mockBackend *mockContractBackend
		expectedErr error
	}{
		{
			name: "success",
			validatorInfos: []struct {
				addr   common.Address
				amount int64
			}{
				{addr: common.HexToAddress("0x1111111111111111111111111111111111111111"), amount: 1000},
				{addr: common.HexToAddress("0x2222222222222222222222222222222222222222"), amount: 2000},
				{addr: common.HexToAddress("0x3333333333333333333333333333333333333333"), amount: 3000},
			},
			threshold: 2,
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callContractResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{
							common.HexToAddress("0x1111111111111111111111111111111111111111"),
							common.HexToAddress("0x2222222222222222222222222222222222222222"),
							common.HexToAddress("0x3333333333333333333333333333333333333333"),
						},
						[]*big.Int{big.NewInt(1000), big.NewInt(2000), big.NewInt(3000)},
					)},
					methodValidatorThreshold: {packValidatorThresholdResponse(t, 2)},
				},
			},
			expectedErr: nil,
		},
		{
			name: "no validator info",
			validatorInfos: []struct {
				addr   common.Address
				amount int64
			}{},
			threshold: 2,
			mockBackend: &mockContractBackend{
				codeAtBytes: []byte{1, 2, 3},
				callContractResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{},
						[]*big.Int{},
					)},
					methodValidatorThreshold: {packValidatorThresholdResponse(t, 2)},
				},
			},
			expectedErr: errors.New("no validator info"),
		},
		{
			name: "zero threshold",
			validatorInfos: []struct {
				addr   common.Address
				amount int64
			}{
				{addr: common.HexToAddress("0x1111111111111111111111111111111111111111"), amount: 1000},
			},
			threshold: 0,
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callContractResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{
							common.HexToAddress("0x1111111111111111111111111111111111111111"),
						},
						[]*big.Int{big.NewInt(1000)},
					)},
					methodValidatorThreshold: {packValidatorThresholdResponse(t, 0)},
				},
			},
			expectedErr: errors.New("zero validator threshold"),
		},
		{
			name: "error from getValidatorInfo",
			validatorInfos: []struct {
				addr   common.Address
				amount int64
			}{},
			threshold: 2,
			mockBackend: &mockContractBackend{
				codeAtBytes:     code,
				callContractErr: errors.New("getValidatorInfo error"),
			},
			expectedErr: errors.New("getValidatorInfo error"),
		},
		{
			name: "error from getValidatorThreshold",
			validatorInfos: []struct {
				addr   common.Address
				amount int64
			}{
				{addr: common.HexToAddress("0x1111111111111111111111111111111111111111"), amount: 1000},
			},
			threshold: 2,
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callContractResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{
							common.HexToAddress("0x1111111111111111111111111111111111111111"),
						},
						[]*big.Int{big.NewInt(1000)},
					)},
				},
				callContractErr: errors.New("getValidatorThreshold error"),
			},
			expectedErr: errors.New("getValidatorThreshold error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
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

			// for success cases, verify that a transaction was created
			assert.NotEmpty(t, txs)
		})
	}
}

func packGetValidatorsResponse(t *testing.T, validatorAddrs []common.Address, amounts []*big.Int) []byte {
	t.Helper()

	parsedABI, err := breakpoint.StakeHubMetaData.ParseABI()
	require.NoError(t, err)

	method, ok := parsedABI.Methods["getValidators"]
	require.True(t, ok)

	totalLength := big.NewInt(int64(len(validatorAddrs)))
	retPacked, err := method.Outputs.Pack(validatorAddrs, amounts, totalLength)
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
	methodNumCheckpoints := string(breakpoint.NewIstanbulParam().PackNumCheckpoints()[:4])
	methodGetParamIndex := string(breakpoint.NewIstanbulParam().PackGetParamIndex(big.NewInt(100))[:4])
	methodGetParamsByIndex := string(breakpoint.NewIstanbulParam().PackGetParamsByIndex(0)[:4])
	callContractErr := errors.New("call contract error")

	tests := []struct {
		name        string
		header      *types.Header
		mockBackend *mockContractBackend
		expectedErr error
		validate    func(t *testing.T)
	}{
		{
			name:   "sync single param",
			header: &types.Header{Number: big.NewInt(100)},
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callContractResponses: map[string][][]byte{
					methodNumCheckpoints: {packNumCheckpointsResponse(t, 1)},
					methodGetParamIndex:  {packGetParamIndexResponse(t, 0)},
					methodGetParamsByIndex: {packGetParamsByIndexResponse(t, breakpoint.GetParamsByIndexOutput{
						Timepoint:                100,
						EpochLength:              1000,
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
				config := params.IstanbulConfigByIndex(0)
				require.NotNil(t, config)
				assert.EqualValues(t, 1000, config.EpochLength)
				assert.EqualValues(t, 5, config.BlockPeriodSeconds)
			},
		},
		{
			name:   "numCheckpoints is 0",
			header: &types.Header{Number: big.NewInt(100)},
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callContractResponses: map[string][][]byte{
					methodNumCheckpoints: {packNumCheckpointsResponse(t, 0)},
				},
			},
			expectedErr: nil,
			validate:    func(t *testing.T) {},
		},
		{
			name:   "error from numCheckpoints",
			header: &types.Header{Number: big.NewInt(100)},
			mockBackend: &mockContractBackend{
				codeAtBytes:     code,
				callContractErr: callContractErr,
			},
			expectedErr: callContractErr,
			validate:    func(t *testing.T) {},
		},
		{
			name:   "error from getParamIndex",
			header: &types.Header{Number: big.NewInt(100)},
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callContractResponses: map[string][][]byte{
					methodNumCheckpoints: {packNumCheckpointsResponse(t, 1)},
				},
				callContractErr: callContractErr,
			},
			expectedErr: callContractErr,
			validate:    func(t *testing.T) {},
		},
		{
			name:   "already synced",
			header: &types.Header{Number: big.NewInt(100)},
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callContractResponses: map[string][][]byte{
					methodNumCheckpoints: {packNumCheckpointsResponse(t, 2)},
					methodGetParamIndex:  {packGetParamIndexResponse(t, 2)},
				},
			},
			expectedErr: nil,
			validate:    func(t *testing.T) {},
		},
		{
			name:   "error from getParamsByIndex",
			header: &types.Header{Number: big.NewInt(100)},
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callContractResponses: map[string][][]byte{
					methodNumCheckpoints: {packNumCheckpointsResponse(t, 1)},
					methodGetParamIndex:  {packGetParamIndexResponse(t, 0)},
				},
				callContractErr: callContractErr,
			},
			expectedErr: callContractErr,
			validate:    func(t *testing.T) {},
		},
		{
			name:   "sync multiple params",
			header: &types.Header{Number: big.NewInt(200)},
			mockBackend: &mockContractBackend{
				codeAtBytes: code,
				callContractResponses: map[string][][]byte{
					methodNumCheckpoints: {packNumCheckpointsResponse(t, 3)},
					methodGetParamIndex:  {packGetParamIndexResponse(t, 2)},
					methodGetParamsByIndex: {
						packGetParamsByIndexResponse(t, breakpoint.GetParamsByIndexOutput{
							Timepoint:                100,
							EpochLength:              1000,
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
							EpochLength:              2000,
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
							EpochLength:              3000,
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
				// Verify configs are cached (0, 1, 2)
				config0 := params.IstanbulConfigByIndex(0)
				require.NotNil(t, config0)
				assert.EqualValues(t, 1000, config0.EpochLength)

				config1 := params.IstanbulConfigByIndex(1)
				require.NotNil(t, config1)
				assert.EqualValues(t, 2000, config1.EpochLength)

				config2 := params.IstanbulConfigByIndex(2)
				require.NotNil(t, config2)
				assert.EqualValues(t, 3000, config2.EpochLength)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := &Engine{
				contractBackend: tt.mockBackend,
				istanbulParam:   breakpoint.NewIstanbulParam(),
			}

			params.ClearCachedIstanbulConfigs()
			err := engine.SyncIstanbulParam(tt.header)

			require.ErrorIs(t, err, tt.expectedErr)

			tt.validate(t)
		})
	}
}

func packNumCheckpointsResponse(t *testing.T, count uint64) []byte {
	t.Helper()

	parsedABI, err := breakpoint.IstanbulParamMetaData.ParseABI()
	require.NoError(t, err)

	method, ok := parsedABI.Methods["numCheckpoints"]
	require.True(t, ok)

	retPacked, err := method.Outputs.Pack(count)
	require.NoError(t, err)

	return retPacked
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
	)
	require.NoError(t, err)

	return retPacked
}

// ##CROSS: validator slash
type mockIstanbulPoS struct {
	offlineValidators []common.Address
}

func (m *mockIstanbulPoS) IsSystemTransaction(tx *types.Transaction, header *types.Header) (bool, error) {
	return false, nil
}
func (m *mockIstanbulPoS) IsSystemContract(to *common.Address) bool {
	return false
}
func (m *mockIstanbulPoS) SyncIstanbulParam(header *types.Header) error {
	return nil
}
func (m *mockIstanbulPoS) OfflineValidators(chain consensus.ChainHeaderReader, number uint64) []common.Address {
	return m.offlineValidators
}

func TestSlashValidators(t *testing.T) {
	tests := []struct {
		name          string
		header        *types.Header
		pos           *mockIstanbulPoS
		expectedTxLen int
	}{
		{
			name:          "no offline validators",
			header:        &types.Header{Number: big.NewInt(100)},
			pos:           &mockIstanbulPoS{offlineValidators: []common.Address{}},
			expectedTxLen: 0,
		},
		{
			name: "success single",
			header: &types.Header{
				Number:     big.NewInt(100),
				ParentHash: common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
				BaseFee:    big.NewInt(0),
				GasLimit:   10000000,
				Difficulty: big.NewInt(0),
				Time:       1000000000,
			},
			pos:           &mockIstanbulPoS{offlineValidators: []common.Address{common.HexToAddress("0x1111111111111111111111111111111111111111")}},
			expectedTxLen: 1,
		},
		{
			name: "success multiple",
			header: &types.Header{
				Number:     big.NewInt(100),
				ParentHash: common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
				BaseFee:    big.NewInt(0),
				GasLimit:   10000000,
				Difficulty: big.NewInt(0),
				Time:       1000000000,
			},
			pos: &mockIstanbulPoS{offlineValidators: []common.Address{
				common.HexToAddress("0x1111111111111111111111111111111111111111"),
				common.HexToAddress("0x2222222222222222222222222222222222222222"),
				common.HexToAddress("0x3333333333333333333333333333333333333333"),
			}},
			expectedTxLen: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
			memdb := rawdb.NewMemoryDatabase()
			triedb := triedb.NewDatabase(memdb, nil)
			sdb := state.NewDatabase(triedb, nil)
			statedb, _ := state.New(types.EmptyRootHash, sdb)

			chainConfig := params.TestChainConfig
			cx := &chainContext{
				Chain:  &mockChainContext{config: chainConfig},
				engine: nil,
			}

			var txs []*types.Transaction
			var receipts []*types.Receipt
			var usedGas uint64

			signer := common.HexToAddress("0x1111111111111111111111111111111111111111")
			engine := &Engine{
				contractBackend: &mockContractBackend{codeAtBytes: []byte{1, 2, 3}},
				validatorSlash:  breakpoint.NewValidatorSlash(),
				signer:          signer,
				signTx: func(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
					require.Equal(t, account.Address, signer)
					return tx, nil
				},
				pos: tt.pos,
			}

			err := engine.slashValidators(tt.header, statedb, cx, &txs, &receipts, nil, &usedGas, nil)

			require.NoError(t, err)
			assert.Equal(t, tt.expectedTxLen, len(txs))
			assert.Equal(t, tt.expectedTxLen, len(receipts))
		})
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
			// setup
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
