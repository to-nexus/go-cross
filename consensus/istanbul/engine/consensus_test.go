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
		validatorSlash: breakpoint.NewValidatorSlash(),
		signTx: func(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
			return tx, nil
		},
	}

	var txs []*types.Transaction
	var receipts []*types.Receipt
	var usedGas uint64

	err := engine.slashValidators(h2, statedb, cx, &txs, &receipts, nil, &usedGas, nil)

	require.NoError(t, err)
	assert.Empty(t, txs)
	assert.Empty(t, receipts)
	assert.Zero(t, usedGas)
	assert.Zero(t, statedb.GetNonce(validators[0]))
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
