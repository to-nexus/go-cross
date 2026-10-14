package core

import (
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/protocols"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rlp"
)

func TestPreprepareRejectsMismatchedProposalNumber(t *testing.T) {
	valSet, keys := generateValidatorSetAndKeys(t, 4)
	maliciousProposer := valSet.GetByIndex(0).Address()
	honestVictim := valSet.GetByIndex(1).Address()

	const (
		currentSequence int64 = 2
		currentRound    int64 = 0
	)

	freshProposal := makeBlockWithTime(currentSequence, 2)
	staleProposal := makeBlockWithTime(currentSequence-1, 1)
	if freshProposal.NumberU64() == staleProposal.NumberU64() {
		t.Fatalf("test setup invalid: fresh and stale proposals must use different heights")
	}

	backend := &preprepareNumberMismatchPoCTestBackend{
		addr:         honestVictim,
		key:          keys[honestVictim],
		valSet:       valSet,
		mux:          new(event.TypeMux),
		lastProposal: makeBlockWithTime(currentSequence-1, 0),
		lastProposer: common.Address{},
	}

	c := New(backend, istanbul.DefaultConfig)
	c.valSet = valSet
	c.current = newRoundState(
		&istanbul.View{Sequence: big.NewInt(currentSequence), Round: big.NewInt(currentRound)},
		valSet,
		nil,
		nil,
		nil,
		&Request{Proposal: freshProposal},
		func(common.Hash) bool { return false },
	)
	c.roundChangeSet = newRoundChangeSet(valSet)
	c.roundChangeSet.NewRound(big.NewInt(currentRound))

	preprepare := protocols.NewPreprepare(big.NewInt(currentSequence), big.NewInt(currentRound), staleProposal)
	preprepare.SetSource(maliciousProposer)

	signProtocolMessage(t, preprepare, keys[maliciousProposer])

	payload, err := rlp.EncodeToBytes(preprepare)
	if err != nil {
		t.Fatalf("failed to encode PRE-PREPARE: %v", err)
	}
	if err := c.handleEncodedMsg(protocols.PreprepareCode, payload); err != errInvalidMessage {
		t.Fatalf("expected mismatched-number PRE-PREPARE to be rejected, got %v", err)
	}
	if c.state != StateAcceptRequest {
		t.Fatalf("expected validator to remain AcceptRequest, got %v", c.state)
	}
	if c.current.Proposal() != nil {
		t.Fatalf("expected current proposal to remain unset")
	}
	if c.current.preparedBlock != nil {
		t.Fatalf("expected stale proposal not to be recorded as preparedBlock")
	}
	if backend.verifyCalls != 0 {
		t.Fatalf("expected mismatched proposal to be rejected before backend verification, got %d calls", backend.verifyCalls)
	}
}

func TestRoundChangeRejectsMismatchedPreparedBlockNumber(t *testing.T) {
	valSet, keys := generateValidatorSetAndKeys(t, 4)
	honestVictim := valSet.GetByIndex(1).Address()

	const (
		currentSequence int64 = 2
		currentRound    int64 = 1
		preparedRound   int64 = 0
	)

	freshProposal := makeBlockWithTime(currentSequence, 2)
	staleProposal := makeBlockWithTime(currentSequence-1, 1)
	if freshProposal.NumberU64() == staleProposal.NumberU64() {
		t.Fatalf("test setup invalid: fresh and stale proposals must use different heights")
	}

	backend := &preprepareNumberMismatchPoCTestBackend{
		addr:         honestVictim,
		key:          keys[honestVictim],
		valSet:       valSet,
		mux:          new(event.TypeMux),
		lastProposal: makeBlockWithTime(currentSequence-1, 0),
		lastProposer: common.Address{},
	}

	c := New(backend, istanbul.DefaultConfig)
	c.valSet = valSet
	c.current = newRoundState(
		&istanbul.View{Sequence: big.NewInt(currentSequence), Round: big.NewInt(currentRound)},
		valSet,
		nil,
		nil,
		nil,
		&Request{Proposal: freshProposal},
		func(common.Hash) bool { return false },
	)
	c.roundChangeSet = newRoundChangeSet(valSet)
	c.roundChangeSet.NewRound(big.NewInt(currentRound))

	prepareSenders := []common.Address{
		valSet.GetByIndex(0).Address(),
		honestVictim,
		valSet.GetByIndex(2).Address(),
	}
	prepareMessages := make([]*protocols.Prepare, 0, len(prepareSenders))
	for _, from := range prepareSenders {
		prepare := protocols.NewPrepareWithSigAndSource(big.NewInt(currentSequence), big.NewInt(preparedRound), staleProposal.Hash(), nil, from)
		signProtocolMessage(t, prepare, keys[from])
		prepareMessages = append(prepareMessages, prepare)
	}

	roundChange := protocols.NewRoundChange(big.NewInt(currentSequence), big.NewInt(currentRound), big.NewInt(preparedRound), staleProposal)
	roundChange.SetSource(prepareSenders[0])
	signProtocolMessage(t, roundChange, keys[prepareSenders[0]])
	roundChange.Justification = prepareMessages

	payload, err := rlp.EncodeToBytes(roundChange)
	if err != nil {
		t.Fatalf("failed to encode ROUND-CHANGE: %v", err)
	}
	if err := c.handleEncodedMsg(protocols.RoundChangeCode, payload); err != errInvalidPreparedBlock {
		t.Fatalf("expected stale prepared block to be rejected, got %v", err)
	}
	if highestPreparedRound, highestPreparedBlock := c.highestPrepared(big.NewInt(currentRound)); highestPreparedRound != nil || highestPreparedBlock != nil {
		t.Fatalf("expected stale prepared block not to be stored, got round %v block %v", highestPreparedRound, highestPreparedBlock)
	}
	if len(backend.broadcasts) != 0 {
		t.Fatalf("expected no PRE-PREPARE broadcast for stale prepared block, got %d broadcasts", len(backend.broadcasts))
	}
}

type preprepareNumberMismatchPoCTestBackend struct {
	addr         common.Address
	key          *ecdsa.PrivateKey
	valSet       istanbul.ValidatorSet
	mux          *event.TypeMux
	broadcasts   []testBroadcast
	verifyCalls  int
	lastProposal istanbul.Proposal
	lastProposer common.Address
}

func (b *preprepareNumberMismatchPoCTestBackend) Address() common.Address {
	return b.addr
}
func (b *preprepareNumberMismatchPoCTestBackend) Validators(istanbul.Proposal) istanbul.ValidatorSet {
	return b.valSet
}
func (b *preprepareNumberMismatchPoCTestBackend) EventMux() *event.TypeMux {
	return b.mux
}
func (b *preprepareNumberMismatchPoCTestBackend) Broadcast(_ istanbul.ValidatorSet, code uint64, payload []byte) error {
	b.broadcasts = append(b.broadcasts, testBroadcast{code: code, payload: append([]byte(nil), payload...)})
	return nil
}
func (b *preprepareNumberMismatchPoCTestBackend) Gossip(istanbul.ValidatorSet, uint64, []byte) error {
	return nil
}
func (b *preprepareNumberMismatchPoCTestBackend) Commit(istanbul.Proposal, []istanbul.SignedSeal, *big.Int) error {
	return nil
}
func (b *preprepareNumberMismatchPoCTestBackend) Verify(istanbul.Proposal) (time.Duration, error) {
	b.verifyCalls++
	return 0, nil
}
func (b *preprepareNumberMismatchPoCTestBackend) Sign(data []byte) ([]byte, error) {
	return crypto.Sign(crypto.Keccak256(data), b.key)
}
func (b *preprepareNumberMismatchPoCTestBackend) SignWithoutHashing(data []byte) ([]byte, error) {
	return crypto.Sign(data, b.key)
}
func (b *preprepareNumberMismatchPoCTestBackend) SignSeal(*types.Header, []byte) ([]byte, error) {
	return nil, nil
}
func (b *preprepareNumberMismatchPoCTestBackend) SealSize(istanbul.Proposal) int {
	return 0
}
func (b *preprepareNumberMismatchPoCTestBackend) VerifyCommittedSeal([]byte, common.Address, istanbul.Proposal, uint32, istanbul.ValidatorSet) error {
	return nil
}
func (b *preprepareNumberMismatchPoCTestBackend) CheckSignature([]byte, common.Address, []byte) error {
	return nil
}
func (b *preprepareNumberMismatchPoCTestBackend) LastProposal() (istanbul.Proposal, common.Address) {
	return b.lastProposal, b.lastProposer
}
func (b *preprepareNumberMismatchPoCTestBackend) HasPropsal(common.Hash, *big.Int) bool {
	return false
}
func (b *preprepareNumberMismatchPoCTestBackend) GetProposer(uint64) common.Address {
	return common.Address{}
}
func (b *preprepareNumberMismatchPoCTestBackend) ParentValidators(istanbul.Proposal) istanbul.ValidatorSet {
	return b.valSet
}
func (b *preprepareNumberMismatchPoCTestBackend) HasBadProposal(common.Hash) bool {
	return false
}
func (b *preprepareNumberMismatchPoCTestBackend) Close() error {
	return nil
}

func generateValidatorSetAndKeys(t *testing.T, n int) (istanbul.ValidatorSet, map[common.Address]*ecdsa.PrivateKey) {
	t.Helper()

	vals := make([]common.Address, n)
	signers := make([]types.BLSPublicKey, n)
	keys := make(map[common.Address]*ecdsa.PrivateKey)
	for i := 0; i < n; i++ {
		privateKey, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(privateKey.PublicKey)
		vals[i] = addr
		keys[addr] = privateKey
	}
	valSet := validator.NewSet(vals, signers, istanbul.NewRoundRobinProposerPolicy())
	return valSet, keys
}

func makeBlockWithTime(number int64, ts uint64) *types.Block {
	header := &types.Header{
		Difficulty: big.NewInt(0),
		Number:     big.NewInt(number),
		GasLimit:   0,
		GasUsed:    0,
		Time:       ts,
	}
	block := &types.Block{}
	return block.WithSeal(header)
}

func signProtocolMessage(t *testing.T, msg protocols.Message, key *ecdsa.PrivateKey) {
	t.Helper()

	encodedPayload, err := msg.EncodePayloadForSigning()
	if err != nil {
		t.Fatal("failed to encode payload", "error", err)
	}

	hashData := crypto.Keccak256(encodedPayload)
	signature, err := crypto.Sign(hashData, key)
	if err != nil {
		t.Fatal("failed to sign data", "error", err)
	}

	msg.SetSignature(signature)
}
