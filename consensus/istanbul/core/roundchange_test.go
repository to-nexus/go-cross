package core

import (
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/protocols"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rlp"
)

// ##CROSS: istanbul validation

func TestRoundChange_RejectsDuplicatedPrepareJustification(t *testing.T) {
	valSet, keys := generateValidatorSetAndKeys(t, 4)
	proposerAddr := valSet.GetByIndex(0).Address()
	maliciousAddr := valSet.GetByIndex(1).Address()
	honestAddr1 := valSet.GetByIndex(2).Address()
	honestAddr2 := valSet.GetByIndex(3).Address()

	const (
		sequence      int64 = 1
		currentRound  int64 = 10
		preparedRound int64 = 5
	)

	benignProposal := makeBlockWithTime(sequence, 1)
	attackerProposal := makeBlockWithTime(sequence, 2)

	if benignProposal.Hash() == attackerProposal.Hash() {
		t.Fatalf("test setup invalid: proposals must differ")
	}

	proposerCore, proposerBackend := newRoundChangeTestCore(t, proposerAddr, keys[proposerAddr], valSet, benignProposal, sequence, currentRound)
	duplicatedPrepare := createSignedPrepareMessage(t, keys[maliciousAddr], maliciousAddr, preparedRound, attackerProposal)
	poisonedRoundChange := createSignedRoundChangeMessage(t, keys[maliciousAddr], maliciousAddr, currentRound, preparedRound, attackerProposal)
	poisonedRoundChange.Justification = []*protocols.Prepare{duplicatedPrepare, duplicatedPrepare, duplicatedPrepare}
	nilRoundChange1 := createSignedRoundChangeMessage(t, keys[honestAddr1], honestAddr1, currentRound, 0, nil)
	nilRoundChange2 := createSignedRoundChangeMessage(t, keys[honestAddr2], honestAddr2, currentRound, 0, nil)

	poisonedPayload, err := rlp.EncodeToBytes(poisonedRoundChange)
	if err != nil {
		t.Fatalf("failed to encode poisoned ROUND-CHANGE message: %v", err)
	}
	if err := proposerCore.handleEncodedMsg(protocols.RoundChangeCode, poisonedPayload); err != errInvalidPreparedBlock {
		t.Fatalf("expected poisoned ROUND-CHANGE to be rejected, got %v", err)
	}

	for _, msg := range []*protocols.RoundChange{nilRoundChange1, nilRoundChange2} {
		payload, err := rlp.EncodeToBytes(msg)
		if err != nil {
			t.Fatalf("failed to encode ROUND-CHANGE message: %v", err)
		}
		if err := proposerCore.handleEncodedMsg(protocols.RoundChangeCode, payload); err != nil {
			t.Fatalf("expected ROUND-CHANGE to be accepted, got %v", err)
		}
	}

	if len(proposerBackend.broadcasts) != 0 {
		t.Fatalf("expected proposer to reject poisoned highest prepared block, got %d broadcasts", len(proposerBackend.broadcasts))
	}
	if highestPreparedRound, highestPreparedBlock := proposerCore.highestPrepared(big.NewInt(currentRound)); highestPreparedRound != nil || highestPreparedBlock != nil {
		t.Fatalf("expected duplicated prepare justification to be ignored, got round %v block %v", highestPreparedRound, highestPreparedBlock)
	}

	preprepare := protocols.NewPreprepare(big.NewInt(sequence), big.NewInt(currentRound), attackerProposal)
	preprepare.SetSource(proposerAddr)
	preprepare.JustificationRoundChanges = []*protocols.SignedRoundChangePayload{
		&poisonedRoundChange.SignedRoundChangePayload,
		&nilRoundChange1.SignedRoundChangePayload,
		&nilRoundChange2.SignedRoundChangePayload,
	}
	preprepare.JustificationPrepares = []*protocols.Prepare{duplicatedPrepare, duplicatedPrepare, duplicatedPrepare}
	signProtocolMessage(t, preprepare, keys[proposerAddr])

	payload, err := rlp.EncodeToBytes(preprepare)
	if err != nil {
		t.Fatalf("failed to encode PRE-PREPARE: %v", err)
	}

	validatorCore, _ := newRoundChangeTestCore(t, honestAddr1, keys[honestAddr1], valSet, benignProposal, sequence, currentRound)
	if err := validatorCore.handleEncodedMsg(protocols.PreprepareCode, payload); err != errInvalidPreparedBlock {
		t.Fatalf("expected honest validator to reject poisoned PRE-PREPARE, got %v", err)
	}
	if validatorCore.state != StateAcceptRequest {
		t.Fatalf("expected honest validator to remain AcceptRequest, got %v", validatorCore.state)
	}
}

func newRoundChangeTestCore(
	t *testing.T,
	addr common.Address,
	key *ecdsa.PrivateKey,
	valSet istanbul.ValidatorSet,
	pendingProposal istanbul.Proposal,
	sequence int64,
	round int64,
) (*Core, *roundChangeTestBackend) {
	t.Helper()

	backend := &roundChangeTestBackend{
		addr:   addr,
		key:    key,
		valSet: valSet,
		mux:    new(event.TypeMux),
	}
	c := New(backend, istanbul.DefaultConfig)
	c.valSet = valSet
	c.current = newRoundState(
		&istanbul.View{Sequence: big.NewInt(sequence), Round: big.NewInt(round)},
		valSet,
		nil,
		nil,
		nil,
		&Request{Proposal: pendingProposal},
		func(common.Hash) bool { return false },
	)
	c.roundChangeSet = newRoundChangeSet(valSet)
	c.roundChangeSet.NewRound(big.NewInt(round))
	return c, backend
}

type roundChangeTestBackend struct {
	addr       common.Address
	key        *ecdsa.PrivateKey
	valSet     istanbul.ValidatorSet
	mux        *event.TypeMux
	broadcasts []testBroadcast
}

type testBroadcast struct {
	code    uint64
	payload []byte
}

func (b *roundChangeTestBackend) Address() common.Address {
	return b.addr
}
func (b *roundChangeTestBackend) Validators(istanbul.Proposal) istanbul.ValidatorSet {
	return b.valSet
}
func (b *roundChangeTestBackend) EventMux() *event.TypeMux {
	return b.mux
}
func (b *roundChangeTestBackend) Broadcast(_ istanbul.ValidatorSet, code uint64, payload []byte) error {
	b.broadcasts = append(b.broadcasts, testBroadcast{code: code, payload: append([]byte(nil), payload...)})
	return nil
}
func (b *roundChangeTestBackend) Gossip(istanbul.ValidatorSet, uint64, []byte) error {
	return nil
}
func (b *roundChangeTestBackend) Commit(istanbul.Proposal, []istanbul.SignedSeal, *big.Int) error {
	return nil
}
func (b *roundChangeTestBackend) Verify(istanbul.Proposal) (time.Duration, error) {
	return 0, nil
}
func (b *roundChangeTestBackend) Sign(data []byte) ([]byte, error) {
	return crypto.Sign(crypto.Keccak256(data), b.key)
}
func (b *roundChangeTestBackend) SignWithoutHashing(data []byte) ([]byte, error) {
	return crypto.Sign(data, b.key)
}
func (b *roundChangeTestBackend) SignSeal(*types.Header, []byte) ([]byte, error) {
	return nil, nil
}
func (b *roundChangeTestBackend) SealSize(istanbul.Proposal) int {
	return 0
}
func (b *roundChangeTestBackend) VerifyCommittedSeal([]byte, common.Address, istanbul.Proposal, uint32, istanbul.ValidatorSet) error {
	return nil
}
func (b *roundChangeTestBackend) CheckSignature([]byte, common.Address, []byte) error {
	return nil
}
func (b *roundChangeTestBackend) LastProposal() (istanbul.Proposal, common.Address) {
	return makeBlockWithTime(0, 0), common.Address{}
}
func (b *roundChangeTestBackend) HasPropsal(common.Hash, *big.Int) bool {
	return false
}
func (b *roundChangeTestBackend) GetProposer(uint64) common.Address {
	return common.Address{}
}
func (b *roundChangeTestBackend) ParentValidators(istanbul.Proposal) istanbul.ValidatorSet {
	return b.valSet
}
func (b *roundChangeTestBackend) HasBadProposal(common.Hash) bool {
	return false
}
func (b *roundChangeTestBackend) Close() error {
	return nil
}

func createSignedPrepareMessage(t *testing.T, key *ecdsa.PrivateKey, from common.Address, round int64, preparedBlock istanbul.Proposal) *protocols.Prepare {
	t.Helper()
	msg := protocols.NewPrepareWithSigAndSource(big.NewInt(1), big.NewInt(round), preparedBlock.Hash(), nil, from)
	signProtocolMessage(t, msg, key)
	return msg
}

func createSignedRoundChangeMessage(t *testing.T, key *ecdsa.PrivateKey, from common.Address, round, preparedRound int64, preparedBlock istanbul.Proposal) *protocols.RoundChange {
	t.Helper()
	msg := protocols.NewRoundChange(big.NewInt(1), big.NewInt(round), big.NewInt(preparedRound), preparedBlock)
	msg.SetSource(from)
	signProtocolMessage(t, msg, key)
	return msg
}
