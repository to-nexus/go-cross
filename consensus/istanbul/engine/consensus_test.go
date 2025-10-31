package engine

import (
	"context"
	"errors"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/triedb"
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

// BlockHashContractCaller methods
func (m *mockContractBackend) CodeAtHash(ctx context.Context, contract common.Address, hash common.Hash) ([]byte, error) {
	return m.codeAtBytes, m.codeAtErr
}

func (m *mockContractBackend) CallContractAtHash(ctx context.Context, call ethereum.CallMsg, hash common.Hash) ([]byte, error) {
	return m.callContract(call)
}

// ContractTransactor methods
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

// ContractFilterer methods
func (m *mockContractBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}
func (m *mockContractBackend) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, nil
}

func TestGetValidators(t *testing.T) {
	code := []byte{1, 2, 3}
	method := string(breakpoint.NewValidatorSet().PackGetValidators()[:4])

	tests := []struct {
		name               string
		blockNr            rpc.BlockNumberOrHash
		mockCaller         *mockContractBackend
		expectedValidators []common.Address
		expectedErr        error
	}{
		{
			name:    "success with validators",
			blockNr: rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(100)),
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
			name:    "success with empty validators",
			blockNr: rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(100)),
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
			name:    "contract not deployed (ErrNoCode)",
			blockNr: rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(100)),
			mockCaller: &mockContractBackend{
				codeAtBytes: []byte{}, // No code
			},
			expectedValidators: nil,
			expectedErr:        nil,
		},
		{
			name:    "error from CallContract",
			blockNr: rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(100)),
			mockCaller: &mockContractBackend{
				codeAtBytes:     []byte{1, 2, 3},
				callContractErr: errors.New("call contract error"),
			},
			expectedValidators: nil,
			expectedErr:        errors.New("call contract error"),
		},
		{
			name:    "success with block hash",
			blockNr: rpc.BlockNumberOrHashWithHash(common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"), false),
			mockCaller: &mockContractBackend{
				codeAtBytes: []byte{1, 2, 3},
				callContractResponses: map[string][][]byte{
					method: {packValidatorsResponse(t, []common.Address{
						common.HexToAddress("0x5555555555555555555555555555555555555555"),
					})},
				},
			},
			expectedValidators: []common.Address{
				common.HexToAddress("0x5555555555555555555555555555555555555555"),
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := &Engine{
				contactBackend: tt.mockCaller,
				validatorSet:   breakpoint.NewValidatorSet(),
			}

			validators, err := engine.getValidators(tt.blockNr)

			if tt.expectedErr != nil {
				if err == nil {
					t.Errorf("expected error %v, got nil", tt.expectedErr)
					return
				}
				if err.Error() != tt.expectedErr.Error() {
					t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// getValidators() already sorts validators, so compare directly
			if !reflect.DeepEqual(validators, tt.expectedValidators) {
				t.Errorf("validators mismatch:\nexpected: %v\ngot: %v", tt.expectedValidators, validators)
			}
		})
	}
}

// packValidatorsResponse packs the validators response for testing
func packValidatorsResponse(t *testing.T, validators []common.Address) []byte {
	t.Helper()

	parsedABI, err := breakpoint.ValidatorSetMetaData.ParseABI()
	if err != nil {
		t.Fatalf("failed to parse ABI: %v", err)
	}
	method, ok := parsedABI.Methods["getValidators"]
	if !ok {
		t.Fatalf("getValidators method not found")
	}
	retPacked, err := method.Outputs.Pack(validators)
	if err != nil {
		t.Fatalf("failed to pack return value: %v", err)
	}

	return retPacked
}

// mockChainContext implements core.ChainContext for testing
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
			name: "success with validators - applyGeneratedSystemTransaction",
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
			var tracer *tracing.Hooks

			// create engine with mock backend
			engine := &Engine{
				contactBackend: tt.mockBackend,
				validatorSet:   breakpoint.NewValidatorSet(),
				stakeHub:       breakpoint.NewStakeHub(),
				signer:         common.HexToAddress("0x1111111111111111111111111111111111111111"),
				signTx: func(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
					return tx, nil
				},
			}

			// call updateValidatorSet
			err := engine.updateValidatorSet(header, statedb, cx, &txs, &receipts, nil, &usedGas, tracer)

			// verify error
			if tt.expectedErr != nil {
				if err == nil {
					t.Errorf("expected error %v, got nil", tt.expectedErr)
					return
				}
				if err.Error() != tt.expectedErr.Error() {
					t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// for success cases, verify that a transaction was created
			if len(txs) == 0 {
				t.Errorf("expected at least one transaction to be created")
			}
		})
	}
}

// packGetValidatorsResponse packs the StakeHub getValidators response for testing
func packGetValidatorsResponse(t *testing.T, validatorAddrs []common.Address, amounts []*big.Int) []byte {
	t.Helper()

	parsedABI, err := breakpoint.StakeHubMetaData.ParseABI()
	if err != nil {
		t.Fatalf("failed to parse ABI: %v", err)
	}
	method, ok := parsedABI.Methods["getValidators"]
	if !ok {
		t.Fatalf("getValidators method not found")
	}
	totalLength := big.NewInt(int64(len(validatorAddrs)))
	retPacked, err := method.Outputs.Pack(validatorAddrs, amounts, totalLength)
	if err != nil {
		t.Fatalf("failed to pack return value: %v", err)
	}
	return retPacked
}

// packValidatorThresholdResponse packs the StakeHub validatorThreshold response for testing
func packValidatorThresholdResponse(t *testing.T, threshold uint64) []byte {
	t.Helper()

	parsedABI, err := breakpoint.StakeHubMetaData.ParseABI()
	if err != nil {
		t.Fatalf("failed to parse ABI: %v", err)
	}
	method, ok := parsedABI.Methods["validatorThreshold"]
	if !ok {
		t.Fatalf("validatorThreshold method not found")
	}
	retPacked, err := method.Outputs.Pack(big.NewInt(int64(threshold)))
	if err != nil {
		t.Fatalf("failed to pack return value: %v", err)
	}

	return retPacked
}

func TestSyncIstanbulParam(t *testing.T) {
	code := []byte{1, 2, 3}
	methodNumCheckpoints := string(breakpoint.NewIstanbulParam().PackNumCheckpoints()[:4])
	methodGetParamIndex := string(breakpoint.NewIstanbulParam().PackGetParamIndex(big.NewInt(100))[:4])
	methodGetParamsByIndex := string(breakpoint.NewIstanbulParam().PackGetParamsByIndex(0)[:4])

	tests := []struct {
		name        string
		header      *types.Header
		mockBackend *mockContractBackend
		expectedErr error
		validate    func(t *testing.T)
	}{
		{
			name:   "success - sync single param",
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
						Beneficiary:              common.HexToAddress("0x1111111111111111111111111111111111111111"),
						ProposerPolicy:           1,
						GasLimit:                 8000000,
					})},
				},
			},
			expectedErr: nil,
			validate: func(t *testing.T) {
				config := params.IstanbulConfigByIndex(0)
				if config == nil {
					t.Errorf("expected config to be cached at index 0")
					return
				}
				if config.EpochLength != 1000 {
					t.Errorf("expected EpochLength 1000, got %d", config.EpochLength)
				}
				if config.BlockPeriodSeconds != 5 {
					t.Errorf("expected BlockPeriodSeconds 5, got %d", config.BlockPeriodSeconds)
				}
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
				callContractErr: errors.New("numCheckpoints error"),
			},
			expectedErr: errors.New("numCheckpoints error"),
			validate:    func(t *testing.T) {},
		},
		{
			name:   "getParamIndex returns ErrNoCode",
			header: &types.Header{Number: big.NewInt(100)},
			mockBackend: &mockContractBackend{
				codeAtBytes: []byte{},
				callContractResponses: map[string][][]byte{
					methodNumCheckpoints: {packNumCheckpointsResponse(t, 1)},
					// getParamIndex response is missing, so CallContract returns nil
					// bind.Call will then check CodeAt, find empty code, and return ErrNoCode
				},
			},
			expectedErr: nil,
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
				callContractErr: errors.New("getParamIndex error"),
			},
			expectedErr: errors.New("getParamIndex error"),
			validate:    func(t *testing.T) {},
		},
		{
			name:   "latestIndex equals length - already synced",
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
				callContractErr: errors.New("getParamsByIndex error"),
			},
			expectedErr: errors.New("getParamsByIndex error"),
			validate:    func(t *testing.T) {},
		},
		{
			name:   "success - sync multiple params",
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
							Beneficiary:              common.HexToAddress("0x1111111111111111111111111111111111111111"),
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
							Beneficiary:              common.HexToAddress("0x2222222222222222222222222222222222222222"),
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
							Beneficiary:              common.Address{},
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
				if config0 == nil {
					t.Errorf("expected config to be cached at index 0")
					return
				}
				if config0.EpochLength != 1000 {
					t.Errorf("expected EpochLength 1000, got %d", config0.EpochLength)
				}

				config1 := params.IstanbulConfigByIndex(1)
				if config1 == nil {
					t.Errorf("expected config to be cached at index 1")
					return
				}
				if config1.EpochLength != 2000 {
					t.Errorf("expected EpochLength 2000, got %d", config1.EpochLength)
				}

				config2 := params.IstanbulConfigByIndex(2)
				if config2 == nil {
					t.Errorf("expected config to be cached at index 2")
					return
				}
				if config2.EpochLength != 3000 {
					t.Errorf("expected EpochLength 3000, got %d", config2.EpochLength)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := &Engine{
				contactBackend: tt.mockBackend,
				istanbulParam:  breakpoint.NewIstanbulParam(),
			}

			params.ClearCachedIstanbulConfigs()
			err := engine.SyncIstanbulParam(tt.header)

			if tt.expectedErr != nil {
				if err == nil {
					t.Errorf("expected error %v, got nil", tt.expectedErr)
					return
				}
				if err.Error() != tt.expectedErr.Error() {
					t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			tt.validate(t)
		})
	}
}

// packNumCheckpointsResponse packs the IstanbulParam numCheckpoints response for testing
func packNumCheckpointsResponse(t *testing.T, count uint64) []byte {
	t.Helper()

	parsedABI, err := breakpoint.IstanbulParamMetaData.ParseABI()
	if err != nil {
		t.Fatalf("failed to parse ABI: %v", err)
	}
	method, ok := parsedABI.Methods["numCheckpoints"]
	if !ok {
		t.Fatalf("numCheckpoints method not found")
	}
	retPacked, err := method.Outputs.Pack(count)
	if err != nil {
		t.Fatalf("failed to pack return value: %v", err)
	}

	return retPacked
}

// packGetParamIndexResponse packs the IstanbulParam getParamIndex response for testing
func packGetParamIndexResponse(t *testing.T, index uint64) []byte {
	t.Helper()

	parsedABI, err := breakpoint.IstanbulParamMetaData.ParseABI()
	if err != nil {
		t.Fatalf("failed to parse ABI: %v", err)
	}
	method, ok := parsedABI.Methods["getParamIndex"]
	if !ok {
		t.Fatalf("getParamIndex method not found")
	}
	retPacked, err := method.Outputs.Pack(index)
	if err != nil {
		t.Fatalf("failed to pack return value: %v", err)
	}

	return retPacked
}

// packGetParamsByIndexResponse packs the IstanbulParam getParamsByIndex response for testing
func packGetParamsByIndexResponse(t *testing.T, result breakpoint.GetParamsByIndexOutput) []byte {
	t.Helper()

	parsedABI, err := breakpoint.IstanbulParamMetaData.ParseABI()
	if err != nil {
		t.Fatalf("failed to parse ABI: %v", err)
	}
	method, ok := parsedABI.Methods["getParamsByIndex"]
	if !ok {
		t.Fatalf("getParamsByIndex method not found")
	}
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
		result.Beneficiary,
		result.ProposerPolicy,
		result.GasLimit,
	)
	if err != nil {
		t.Fatalf("failed to pack return value: %v", err)
	}

	return retPacked
}
