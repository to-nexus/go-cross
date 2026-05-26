package engine

import (
	"bytes"
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
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

	revertCode = []byte{0x60, 0x00, 0x60, 0x00, 0xfd}
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
		{
			name:   "validator signer length mismatch",
			number: 100,
			mockCaller: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					method: {packActiveValidatorsResponse(t,
						[]common.Address{addr1, addr2},
						[][]byte{signer1},
					)},
				},
			},
			expectedValidators: nil,
			expectedErr:        errValidatorSetLengthMismatch,
		},
		{
			name:   "invalid signer length",
			number: 100,
			mockCaller: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					method: {packActiveValidatorsResponse(t,
						[]common.Address{addr1},
						[][]byte{[]byte{1, 2, 3}},
					)},
				},
			},
			expectedValidators: nil,
			expectedErr:        errInvalidValidatorSigner,
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

func TestVerifyValidators(t *testing.T) {
	code := []byte{1, 2, 3}
	method := string(breakpoint.NewValidatorSet().PackGetActiveValidators()[:4])
	signerKey1 := types.BytesToBLSPublicKey(signer1)
	signerKey2 := types.BytesToBLSPublicKey(signer2)
	signerKey3 := types.BytesToBLSPublicKey(signer3)

	tests := []struct {
		name               string
		number             uint64
		headerValidators   []common.Address
		headerSigners      []types.BLSPublicKey
		contractValidators []common.Address
		contractSigners    [][]byte
		expectedErr        error
		expectedCalls      int
	}{
		{
			name:             "success",
			number:           100,
			headerValidators: []common.Address{addr1, addr2},
			headerSigners:    []types.BLSPublicKey{signerKey1, signerKey2},
			// The contract response is sorted by GetAcviteValidators before verification.
			contractValidators: []common.Address{addr2, addr1},
			contractSigners:    [][]byte{signer2, signer1},
			expectedCalls:      1,
		},
		{
			name:               "validator mismatch",
			number:             100,
			headerValidators:   []common.Address{addr1, addr3},
			headerSigners:      []types.BLSPublicKey{signerKey1, signerKey3},
			contractValidators: []common.Address{addr1, addr2},
			contractSigners:    [][]byte{signer1, signer2},
			expectedErr:        errValidatorSetMismatch,
			expectedCalls:      1,
		},
		{
			name:               "signer mismatch",
			number:             100,
			headerValidators:   []common.Address{addr1, addr2},
			headerSigners:      []types.BLSPublicKey{signerKey1, signerKey3},
			contractValidators: []common.Address{addr1, addr2},
			contractSigners:    [][]byte{signer1, signer2},
			expectedErr:        errValidatorSignerMismatch,
			expectedCalls:      1,
		},
		{
			name:             "not validator epoch",
			number:           101,
			headerValidators: []common.Address{addr1},
			headerSigners:    []types.BLSPublicKey{signerKey1},
		},
		{
			name:               "empty contract validators",
			number:             100,
			headerValidators:   []common.Address{addr1},
			headerSigners:      []types.BLSPublicKey{signerKey1},
			contractValidators: []common.Address{},
			contractSigners:    [][]byte{},
			expectedErr:        errValidatorSetMismatch,
			expectedCalls:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBackend := &mockContractBackend{codeAtBytes: code}
			if tt.contractValidators != nil {
				mockBackend.callResponses = map[string][][]byte{
					method: {packActiveValidatorsResponse(t, tt.contractValidators, tt.contractSigners)},
				}
			}
			engine := &Engine{
				cfg:             &istanbul.Config{ValidatorEpochLength: 100},
				contractBackend: mockBackend,
				validatorSet:    breakpoint.NewValidatorSet(),
			}
			header := createHeaderWithValidatorsAndSigners(t, tt.number, tt.headerValidators, tt.headerSigners)
			chain := &mockChainHeaderReader{
				headers: map[common.Hash]*types.Header{},
				config:  params.TestChainConfig,
			}

			err := engine.verifyValidators(chain, header)

			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expectedCalls, mockBackend.callCnt)
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

func TestUpdateValidatorSet(t *testing.T) {
	code := []byte{1, 2, 3}
	methodGetValidators := string(breakpoint.NewStakeHub().PackGetValidators(big.NewInt(0), big.NewInt(0))[:4])
	methodMaxCouncil := string(breakpoint.NewStakeHub().PackMaxCouncil()[:4])

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
					methodMaxCouncil: {packMaxCouncilResponse(t, 2)},
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
					methodMaxCouncil: {packMaxCouncilResponse(t, 2)},
				},
			},
			expectedErr: errors.New("empty staked validators"),
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
					methodMaxCouncil: {packMaxCouncilResponse(t, 0)},
				},
			},
			expectedErr: errors.New("zero validator threshold"),
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
			cx := &chainMock{config: chainConfig}

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

func packMaxCouncilResponse(t *testing.T, maxCouncil uint64) []byte {
	t.Helper()

	parsedABI, err := breakpoint.StakeHubMetaData.ParseABI()
	require.NoError(t, err)

	method, ok := parsedABI.Methods["maxCouncil"]
	require.True(t, ok)

	retPacked, err := method.Outputs.Pack(big.NewInt(int64(maxCouncil)))
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
				consensus:      &mockIstanbulEngine{validators: validators},
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

			if tt.expectedSlashed > 0 {
				assert.Len(t, txs, 1)
			} else {
				assert.Len(t, txs, 0)
			}
		})
	}
}

func TestOfflineProposers(t *testing.T) {
	validators := []common.Address{addr1, addr2, addr3}
	headers := make(map[common.Hash]*types.Header)
	h0 := createHeaderWithIstanbulExtra(t, 98, common.Hash{}, validators[0], validators, 0)
	h1 := createHeaderWithIstanbulExtra(t, 99, h0.Hash(), validators[1], validators, 3)
	headers[h0.Hash()] = h0
	headers[h1.Hash()] = h1

	mockChain := &mockChainHeaderReader{
		headers: headers,
		config:  params.TestChainConfig,
	}
	engine := &Engine{
		cfg: &istanbul.Config{
			ProposerPolicy: istanbul.NewRoundRobinProposerPolicy(),
		},
		consensus: &mockIstanbulEngine{validators: validators},
	}

	proposers, err := engine.offlineProposers(mockChain, h1)

	require.NoError(t, err)
	assert.Equal(t, []common.Address{validators[1], validators[2], validators[0]}, proposers)
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

type mockIstanbulEngine struct {
	consensus.Engine
	validators []common.Address
}

var _ consensus.Engine = (*mockIstanbulEngine)(nil)
var _ consensus.IstanbulEngine = (*mockIstanbulEngine)(nil)

func (m *mockIstanbulEngine) IsSystemTransaction(tx *types.Transaction, header *types.Header) (bool, error) {
	return false, nil
}
func (m *mockIstanbulEngine) IsSystemContract(to *common.Address) bool {
	return false
}
func (m *mockIstanbulEngine) EstimateGasForSystemTxs(chain consensus.ChainHeaderReader, header *types.Header) uint64 {
	return 0
}
func (m *mockIstanbulEngine) ValidatorsAt(chain consensus.ChainHeaderReader, header *types.Header) []common.Address {
	return append([]common.Address(nil), m.validators...)
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

func createHeaderWithValidatorsAndSigners(t *testing.T, number uint64, validators []common.Address, signers []types.BLSPublicKey) *types.Header {
	t.Helper()

	extra := &types.IstanbulExtra{
		VanityData:    make([]byte, types.IstanbulExtraVanity),
		Validators:    validators,
		Vote:          nil,
		Round:         0,
		CommittedSeal: [][]byte{},
		RandomReveal:  []byte{},
		Signers:       signers,
	}

	extraBytes, err := rlp.EncodeToBytes(extra)
	require.NoError(t, err)

	return &types.Header{
		Number:     big.NewInt(int64(number)),
		ParentHash: common.Hash{},
		Extra:      extraBytes,
		BaseFee:    big.NewInt(0),
		GasLimit:   10000000,
		Difficulty: big.NewInt(0),
	}
}

func TestSlashValidatorsRevert(t *testing.T) {
	validators := []common.Address{addr1, addr2, addr3}
	headers := make(map[common.Hash]*types.Header)
	h0 := createHeaderWithIstanbulExtra(t, 98, common.Hash{}, validators[0], validators, 0)
	h1 := createHeaderWithIstanbulExtra(t, 99, h0.Hash(), validators[2], validators, 1)
	h2 := createHeaderWithIstanbulExtra(t, 100, h1.Hash(), validators[0], validators, 0)
	headers[h0.Hash()] = h0
	headers[h1.Hash()] = h1
	headers[h2.Hash()] = h2

	memdb := rawdb.NewMemoryDatabase()
	triedb := triedb.NewDatabase(memdb, nil)
	sdb := state.NewDatabase(triedb, nil)
	statedb, _ := state.New(types.EmptyRootHash, sdb)
	statedb.SetCode(contracts.ValidatorSlashAddr, revertCode)

	mockChain := &mockChainHeaderReader{
		headers: headers,
		config:  params.TestChainConfig,
	}
	cx := &chainContext{Chain: mockChain}
	engine := &Engine{
		cfg: &istanbul.Config{
			ProposerPolicy: istanbul.NewRoundRobinProposerPolicy(),
		},
		signer:         validators[0],
		consensus:      &mockIstanbulEngine{validators: validators},
		validatorSlash: breakpoint.NewValidatorSlash(),
		signTx: func(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
			return tx, nil
		},
	}

	var txs []*types.Transaction
	var receipts []*types.Receipt
	var usedGas uint64

	err := engine.slashValidators(h2, statedb, cx, &txs, &receipts, nil, &usedGas, nil)

	require.ErrorContains(t, err, "execution reverted")
	assert.Empty(t, txs)
	assert.Empty(t, receipts)
	assert.Zero(t, usedGas)
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
			cx := &chainMock{config: params.TestChainConfig}

			txs := make([]*types.Transaction, len(tt.txs))
			copy(txs, tt.txs)
			receipts := make([]*types.Receipt, len(tt.receipts))
			copy(receipts, tt.receipts)
			var usedGas uint64

			signer := common.HexToAddress("0x1111111111111111111111111111111111111111")
			statedb.AddBalance(signer, uint256.MustFromBig(tt.expectedFee), tracing.BalanceChangeUnspecified)
			engine := &Engine{
				contractBackend: &mockContractBackend{codeAtBytes: []byte{1, 2, 3}},
				rewardHub:       breakpoint.NewRewardHub(),
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

				assert.Equal(t, contracts.RewardHubAddr, *tx.To())
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

func TestSelectActiveCouncil(t *testing.T) {
	pubFor := func(b byte) types.BLSPublicKey {
		bs := make([]byte, types.BLSPublicKeyLength)
		for i := range bs {
			bs[i] = b
		}
		return types.BytesToBLSPublicKey(bs)
	}
	mkInfo := func(addr common.Address, sig byte, power uint64) ValidatorInfo {
		return ValidatorInfo{address: addr, signer: pubFor(sig), power: uint256.NewInt(power)}
	}

	tests := []struct {
		name      string
		infos     []ValidatorInfo
		threshold uint64
		expected  []common.Address
	}{
		{
			name: "power sort and address asc",
			infos: []ValidatorInfo{
				mkInfo(addr1, 0xaa, 1000),
				mkInfo(addr2, 0xbb, 3000),
				mkInfo(addr3, 0xcc, 2000),
			},
			threshold: 5,
			// All three retained (threshold>count), final order by address asc.
			expected: []common.Address{addr1, addr2, addr3},
		},
		{
			name: "threshold cut keeps top N",
			infos: []ValidatorInfo{
				mkInfo(addr1, 0xaa, 1000),
				mkInfo(addr2, 0xbb, 3000),
				mkInfo(addr3, 0xcc, 2000),
			},
			threshold: 2,
			// Top 2 by power: addr2(3000) and addr3(2000); then address asc.
			expected: []common.Address{addr2, addr3},
		},
		{
			name: "zero power dropped",
			infos: []ValidatorInfo{
				mkInfo(addr1, 0xaa, 5000),
				mkInfo(addr2, 0xbb, 0),
				mkInfo(addr3, 0xcc, 1000),
			},
			threshold: 5,
			// addr2 dropped (zero power)
			expected: []common.Address{addr1, addr3},
		},
		{
			name: "stable sort preserves registration order on tie",
			infos: []ValidatorInfo{
				mkInfo(addr2, 0xbb, 1000),
				mkInfo(addr1, 0xaa, 1000),
				mkInfo(addr3, 0xcc, 1000),
			},
			threshold: 2,
			// All three have equal power -> stable sort keeps input order: addr2, addr1, addr3.
			// Top 2 = addr2, addr1; final address-asc sort = addr1, addr2.
			expected: []common.Address{addr1, addr2},
		},
		{
			name:      "empty input",
			infos:     nil,
			threshold: 3,
			expected:  []common.Address{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectActiveCouncil(tt.infos, tt.threshold)
			addrs := make([]common.Address, 0, len(got))
			for _, v := range got {
				addrs = append(addrs, v.address)
			}
			if tt.expected == nil {
				tt.expected = []common.Address{}
			}
			assert.Equal(t, tt.expected, addrs)
		})
	}
}

func TestComputeNextCouncil(t *testing.T) {
	code := []byte{1, 2, 3}
	methodGetValidators := string(breakpoint.NewStakeHub().PackGetValidators(big.NewInt(0), big.NewInt(0))[:4])
	methodMaxCouncil := string(breakpoint.NewStakeHub().PackMaxCouncil()[:4])

	type expected struct {
		validators []common.Address
		signers    []types.BLSPublicKey
		errSubstr  string
	}
	tests := []struct {
		name     string
		backend  *mockContractBackend
		expected expected
	}{
		{
			name: "top-2 then address-asc",
			backend: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{addr1, addr2, addr3},
						[][]byte{signer1, signer2, signer3},
						[]*big.Int{big.NewInt(1000), big.NewInt(3000), big.NewInt(2000)},
					)},
					methodMaxCouncil: {packMaxCouncilResponse(t, 2)},
				},
			},
			// Top-2 by power = addr2(3000), addr3(2000) -> address-asc -> addr2, addr3.
			expected: expected{
				validators: []common.Address{addr2, addr3},
				signers:    []types.BLSPublicKey{types.BytesToBLSPublicKey(signer2), types.BytesToBLSPublicKey(signer3)},
			},
		},
		{
			name: "empty staked validators",
			backend: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{},
						[][]byte{},
						[]*big.Int{},
					)},
					methodMaxCouncil: {packMaxCouncilResponse(t, 2)},
				},
			},
			expected: expected{errSubstr: "empty staked validators"},
		},
		{
			name: "zero validator threshold",
			backend: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{addr1},
						[][]byte{signer1},
						[]*big.Int{big.NewInt(1000)},
					)},
					methodMaxCouncil: {packMaxCouncilResponse(t, 0)},
				},
			},
			expected: expected{errSubstr: "zero validator threshold"},
		},
		{
			name: "all zero-power filtered to empty",
			backend: &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{addr1, addr2},
						[][]byte{signer1, signer2},
						[]*big.Int{big.NewInt(0), big.NewInt(0)},
					)},
					methodMaxCouncil: {packMaxCouncilResponse(t, 5)},
				},
			},
			expected: expected{errSubstr: "empty filtered validators"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				contractBackend: tt.backend,
				stakeHub:        breakpoint.NewStakeHub(),
			}
			validators, signers, err := e.computeNextCouncil(100)
			if tt.expected.errSubstr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expected.errSubstr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.expected.validators, validators)
			assert.Equal(t, tt.expected.signers, signers)
		})
	}
}

func TestComputeNextCouncilMatchesUpdateValidatorSet(t *testing.T) {
	code := []byte{1, 2, 3}
	methodGetValidators := string(breakpoint.NewStakeHub().PackGetValidators(big.NewInt(0), big.NewInt(0))[:4])
	methodMaxCouncil := string(breakpoint.NewStakeHub().PackMaxCouncil()[:4])

	// Two-call backend: once for computeNextCouncil, once for updateValidatorSet.
	backend := &mockContractBackend{
		codeAtBytes: code,
		callResponses: map[string][][]byte{
			methodGetValidators: {
				packGetValidatorsResponse(t,
					[]common.Address{addr1, addr2, addr3},
					[][]byte{signer1, signer2, signer3},
					[]*big.Int{big.NewInt(1000), big.NewInt(3000), big.NewInt(2000)},
				),
				packGetValidatorsResponse(t,
					[]common.Address{addr1, addr2, addr3},
					[][]byte{signer1, signer2, signer3},
					[]*big.Int{big.NewInt(1000), big.NewInt(3000), big.NewInt(2000)},
				),
			},
			methodMaxCouncil: {
				packMaxCouncilResponse(t, 2),
				packMaxCouncilResponse(t, 2),
			},
		},
	}

	e := &Engine{
		contractBackend: backend,
		validatorSet:    breakpoint.NewValidatorSet(),
		stakeHub:        breakpoint.NewStakeHub(),
		signer:          addr1,
		signTx: func(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
			return tx, nil
		},
	}

	previewValidators, previewSigners, err := e.computeNextCouncil(100)
	require.NoError(t, err)

	// Now run updateValidatorSet and decode the validators encoded into the system tx data.
	memdb := rawdb.NewMemoryDatabase()
	triedb := triedb.NewDatabase(memdb, nil)
	sdb := state.NewDatabase(triedb, nil)
	statedb, _ := state.New(types.EmptyRootHash, sdb)

	header := &types.Header{
		Number:     big.NewInt(101),
		ParentHash: common.HexToHash("0xdeadbeef"),
		Coinbase:   addr1,
		BaseFee:    big.NewInt(0),
		GasLimit:   10000000,
		Difficulty: big.NewInt(0),
		Time:       1000000000,
	}
	cx := &chainMock{config: params.TestChainConfig}
	var (
		txs      []*types.Transaction
		receipts []*types.Receipt
		usedGas  uint64
	)
	err = e.updateValidatorSet(header, statedb, cx, &txs, &receipts, nil, &usedGas, nil)
	require.NoError(t, err)
	require.Len(t, txs, 1, "updateValidatorSet must emit exactly one system tx")

	// Decode the system tx data and compare to the preview.
	tx := txs[0]
	updateValidatorsSelector := e.validatorSet.PackUpdateValidators(nil, nil)[:4]
	require.True(t, bytes.HasPrefix(tx.Data(), updateValidatorsSelector), "system tx is not an updateValidators call")

	parsedABI, err := breakpoint.ValidatorSetMetaData.ParseABI()
	require.NoError(t, err)
	method := parsedABI.Methods["updateValidators"]
	values, err := method.Inputs.Unpack(tx.Data()[4:])
	require.NoError(t, err)
	require.Len(t, values, 2)

	gotValidators, ok := values[0].([]common.Address)
	require.True(t, ok)
	gotSignersBytes, ok := values[1].([][]byte)
	require.True(t, ok)
	gotSigners := make([]types.BLSPublicKey, 0, len(gotSignersBytes))
	for _, b := range gotSignersBytes {
		gotSigners = append(gotSigners, types.BytesToBLSPublicKey(b))
	}

	assert.Equal(t, previewValidators, gotValidators, "validators previewed in Prepare/verify must match the council written by Finalize")
	assert.Equal(t, previewSigners, gotSigners, "signers previewed in Prepare/verify must match the council written by Finalize")
}

func TestVerifyValidatorsBoundaryCoincidence(t *testing.T) {
	code := []byte{1, 2, 3}
	methodGetActive := string(breakpoint.NewValidatorSet().PackGetActiveValidators()[:4])
	methodGetValidators := string(breakpoint.NewStakeHub().PackGetValidators(big.NewInt(0), big.NewInt(0))[:4])
	methodMaxCouncil := string(breakpoint.NewStakeHub().PackMaxCouncil()[:4])

	// Boundary scenario:
	//   - ValidatorEpochLength = 100, CouncilPeriod = 86400.
	//   - Block 100: validator epoch boundary AND first block of a new council period
	//     (parent.Time/86400 != header.Time/86400).
	//
	// OLD council (currently in ValidatorSet contract at parent): {addr1, addr2}
	// NEW council to be installed in Finalize via updateValidatorSet -- top-2 of StakeHub:
	//   StakeHub returns addr1(1000), addr2(0), addr3(2000) with maxCouncil=2.
	//   Top-2 then address-asc -> NEW = {addr1, addr3}.
	//
	// Therefore header.Extra MUST carry NEW = {addr1, addr3}.

	newCouncilValidators := []common.Address{addr1, addr3}
	newCouncilSigners := []types.BLSPublicKey{types.BytesToBLSPublicKey(signer1), types.BytesToBLSPublicKey(signer3)}
	oldCouncilValidators := []common.Address{addr1, addr2}
	oldCouncilSigners := []types.BLSPublicKey{types.BytesToBLSPublicKey(signer1), types.BytesToBLSPublicKey(signer2)}

	const (
		councilPeriod = uint64(86400)
		parentTime    = councilPeriod - 1 // last second of OLD council period
		headerTime    = councilPeriod     // first second of NEW council period
	)

	// Parent header registered in chain so verifyValidators can detect the coincident boundary.
	parentHeader := &types.Header{
		Number: big.NewInt(99),
		Time:   parentTime,
	}
	parentHash := parentHeader.Hash()

	mkBackend := func() *mockContractBackend {
		return &mockContractBackend{
			codeAtBytes: code,
			callResponses: map[string][][]byte{
				// computeNextCouncil path: StakeHub.getValidators + maxCouncil.
				methodGetValidators: {packGetValidatorsResponse(t,
					[]common.Address{addr1, addr2, addr3},
					[][]byte{signer1, signer2, signer3},
					[]*big.Int{big.NewInt(1000), big.NewInt(0), big.NewInt(2000)},
				)},
				methodMaxCouncil: {packMaxCouncilResponse(t, 2)},
				// If verifyValidators ever falls back to getCurrentValidators,
				// it would see the OLD council and the header (which carries NEW) would be rejected with errValidatorSetMismatch.
				methodGetActive: {packActiveValidatorsResponse(t, oldCouncilValidators, [][]byte{signer1, signer2})},
			},
		}
	}

	tests := []struct {
		name             string
		headerValidators []common.Address
		headerSigners    []types.BLSPublicKey
		registerParent   bool
		expectedErr      error
	}{
		{
			name:             "header carries NEW council -> accepted",
			headerValidators: newCouncilValidators,
			headerSigners:    newCouncilSigners,
			registerParent:   true,
		},
		{
			name:             "header carries OLD council (pre-fix output) -> rejected",
			headerValidators: oldCouncilValidators,
			headerSigners:    oldCouncilSigners,
			registerParent:   true,
			expectedErr:      errValidatorSetMismatch,
		},
		{
			name:             "missing parent header: falls back to current contract state (OLD)",
			headerValidators: oldCouncilValidators,
			headerSigners:    oldCouncilSigners,
			registerParent:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backend := mkBackend()
			engine := &Engine{
				cfg: &istanbul.Config{
					ValidatorEpochLength: 100,
					CouncilPeriod:        councilPeriod,
				},
				contractBackend: backend,
				validatorSet:    breakpoint.NewValidatorSet(),
				stakeHub:        breakpoint.NewStakeHub(),
			}

			header := createHeaderWithValidatorsAndSigners(t, 100, tt.headerValidators, tt.headerSigners)
			header.ParentHash = parentHash
			header.Time = headerTime

			headers := map[common.Hash]*types.Header{}
			if tt.registerParent {
				headers[parentHash] = parentHeader
			}
			chain := &mockChainHeaderReader{headers: headers, config: params.TestChainConfig}

			err := engine.verifyValidators(chain, header)
			if tt.expectedErr != nil {
				require.ErrorIs(t, err, tt.expectedErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestPrepareValidatorsBoundaryCoincidence(t *testing.T) {
	code := []byte{1, 2, 3}
	methodGetActive := string(breakpoint.NewValidatorSet().PackGetActiveValidators()[:4])
	methodGetValidators := string(breakpoint.NewStakeHub().PackGetValidators(big.NewInt(0), big.NewInt(0))[:4])
	methodMaxCouncil := string(breakpoint.NewStakeHub().PackMaxCouncil()[:4])

	// Same scenario as TestVerifyValidatorsBoundaryCoincidence
	const councilPeriod = uint64(86400)

	// Set up a PoSA-active chain config
	posaConfig := *params.TestChainConfig
	breakpointTime := uint64(0)
	posaConfig.BreakpointTime = &breakpointTime
	posaConfig.Istanbul = &params.IstanbulConfig{
		EpochLength: 100,
		PoSA: &params.PoSAConfig{
			CouncilPeriod:        councilPeriod,
			ValidatorEpochLength: 100,
		},
	}

	tests := []struct {
		name               string
		parentTime         uint64
		headerTime         uint64
		expectedValidators []common.Address
	}{
		{
			// parent in council bucket 0 (Time<86400), header in bucket 1 (Time==86400)
			// -> OnNewCouncilPeriod(parent, header) == true -> computeNextCouncil path.
			name:       "header.Extra gets NEW council from StakeHub",
			parentTime: councilPeriod - 1,
			headerTime: councilPeriod,
			// NEW council derived from StakeHub: top-2 then address-asc.
			// StakeHub returns addr1(1000), addr2(0), addr3(2000) -> {addr1, addr3}.
			expectedValidators: []common.Address{addr1, addr3},
		},
		{
			// Both in the same council bucket -> OnNewCouncilPeriod returns false ->
			// default path reads ValidatorSet.getActiveValidators -> {addr1, addr2}.
			name:               "header.Extra gets OLD council from ValidatorSet contract",
			parentTime:         100,
			headerTime:         200,
			expectedValidators: []common.Address{addr1, addr2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parentHeader := &types.Header{
				Number: big.NewInt(99),
				Time:   tt.parentTime,
			}
			parentHash := parentHeader.Hash()

			backend := &mockContractBackend{
				codeAtBytes: code,
				callResponses: map[string][][]byte{
					methodGetValidators: {packGetValidatorsResponse(t,
						[]common.Address{addr1, addr2, addr3},
						[][]byte{signer1, signer2, signer3},
						[]*big.Int{big.NewInt(1000), big.NewInt(0), big.NewInt(2000)},
					)},
					methodMaxCouncil: {packMaxCouncilResponse(t, 2)},
					methodGetActive: {packActiveValidatorsResponse(t,
						[]common.Address{addr1, addr2},
						[][]byte{signer1, signer2},
					)},
				},
			}

			engine := &Engine{
				cfg: &istanbul.Config{
					ValidatorEpochLength: 100,
					CouncilPeriod:        councilPeriod,
				},
				contractBackend: backend,
				validatorSet:    breakpoint.NewValidatorSet(),
				stakeHub:        breakpoint.NewStakeHub(),
			}

			header := &types.Header{
				Number:     big.NewInt(100),
				ParentHash: parentHash,
				Time:       tt.headerTime,
				BaseFee:    big.NewInt(0),
				GasLimit:   10000000,
				Difficulty: big.NewInt(0),
			}
			chain := &mockChainHeaderReader{
				headers: map[common.Hash]*types.Header{parentHash: parentHeader},
				config:  &posaConfig,
			}

			applies, err := engine.prepareValidators(chain, header, nil)
			require.NoError(t, err)
			require.NotEmpty(t, applies)

			// Apply the ApplyExtra functions to a fresh IstanbulExtra to inspect what would be written.
			extra := &types.IstanbulExtra{VanityData: make([]byte, types.IstanbulExtraVanity)}
			for _, apply := range applies {
				require.NoError(t, apply(extra))
			}

			assert.Equal(t, tt.expectedValidators, extra.Validators)
		})
	}
}

// ##
