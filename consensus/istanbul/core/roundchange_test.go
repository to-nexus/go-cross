package core

import (
	"crypto/ecdsa"
	"errors"
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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestPreprepare_RejectsStaleRoundChangeJustification(t *testing.T) {
	valSet, keys := generateValidatorSetAndKeys(t, 4)
	proposerAddr := valSet.GetByIndex(0).Address()
	validatorAddr := valSet.GetByIndex(1).Address()

	const (
		sequence     int64 = 1
		currentRound int64 = 2
		staleRound   int64 = 1
	)

	preparedProposal := makeBlockWithTime(sequence, 1)
	attackerProposal := makeBlockWithTime(sequence, 2)
	validatorCore, validatorBackend := newRoundChangeTestCore(
		t, validatorAddr, keys[validatorAddr], valSet, preparedProposal, sequence, currentRound,
	)
	validatorCore.current.preparedRound = big.NewInt(staleRound)
	validatorCore.current.preparedBlock = preparedProposal

	preprepare := protocols.NewPreprepare(big.NewInt(sequence), big.NewInt(currentRound), attackerProposal)
	preprepare.SetSource(proposerAddr)
	for i := 0; i < valSet.QuorumSize(); i++ {
		addr := valSet.GetByIndex(uint64(i)).Address()
		roundChange := createSignedRoundChangeMessage(t, keys[addr], addr, staleRound, 0, nil)
		preprepare.JustificationRoundChanges = append(preprepare.JustificationRoundChanges, &roundChange.SignedRoundChangePayload)
	}
	signProtocolMessage(t, preprepare, keys[proposerAddr])

	payload, err := rlp.EncodeToBytes(preprepare)
	if err != nil {
		t.Fatalf("failed to encode PRE-PREPARE: %v", err)
	}
	if err := validatorCore.handleEncodedMsg(protocols.PreprepareCode, payload); err != errInvalidPreparedBlock {
		t.Fatalf("expected stale ROUND-CHANGE justification to be rejected, got %v", err)
	}
	if validatorCore.state != StateAcceptRequest {
		t.Fatalf("expected validator to remain AcceptRequest, got %v", validatorCore.state)
	}
	if validatorCore.current.preparedBlock.Hash() != preparedProposal.Hash() {
		t.Fatal("expected validator to keep its prepared proposal")
	}
	if len(validatorBackend.broadcasts) != 0 {
		t.Fatalf("expected no PREPARE broadcast, got %d", len(validatorBackend.broadcasts))
	}
}

// ##CROSS: blob sidecars
func TestRoundChange_ReproposesPreparedBlockWithSidecars(t *testing.T) {
	valSet, keys := generateValidatorSetAndKeys(t, 4)
	proposerAddr := valSet.GetByIndex(0).Address()

	const (
		sequence      int64 = 1
		currentRound  int64 = 2
		preparedRound int64 = 1
	)
	blobTx := types.NewTx(&types.BlobTx{BlobHashes: []common.Hash{{0x01}}})
	prepared := makeBlockWithTime(sequence, 1).WithBody(types.Body{Transactions: []*types.Transaction{blobTx}})
	sidecars := types.BlobSidecars{{BlockNumber: big.NewInt(sequence), BlockHash: prepared.Hash(), TxHash: blobTx.Hash()}}

	// newCore returns a proposer whose backend accepts a blob block only with its sidecars, like the DA check.
	newCore := func(t *testing.T) (*Core, *roundChangeTestBackend) {
		c, b := newRoundChangeTestCore(t, proposerAddr, keys[proposerAddr], valSet, makeBlockWithTime(sequence, 2), sequence, currentRound)
		b.verify = func(p istanbul.Proposal) error {
			if len(p.Sidecars()) == 0 {
				return errors.New("unavailable blob data")
			}
			return nil
		}
		return c, b
	}
	// run sends a quorum of ROUND-CHANGE messages that justify the prepared block, the i-th one carrying
	// rcBlock(i), and returns the proposal of the PRE-PREPARE the proposer sends for the new round.
	run := func(t *testing.T, c *Core, b *roundChangeTestBackend, rcBlock func(i int) *types.Block) istanbul.Proposal {
		t.Helper()
		var prepares []*protocols.Prepare
		for i := 0; i < valSet.QuorumSize(); i++ {
			addr := valSet.GetByIndex(uint64(i)).Address()
			prepares = append(prepares, createSignedPrepareMessage(t, keys[addr], addr, preparedRound, prepared))
		}
		for i := 0; i < valSet.QuorumSize(); i++ {
			addr := valSet.GetByIndex(uint64(i)).Address()
			rc := createSignedRoundChangeMessage(t, keys[addr], addr, currentRound, preparedRound, rcBlock(i))
			rc.Justification = prepares
			payload, err := rlp.EncodeToBytes(rc)
			require.NoError(t, err)
			require.NoError(t, c.handleEncodedMsg(protocols.RoundChangeCode, payload))
		}
		for _, bc := range b.broadcasts {
			if bc.code == protocols.PreprepareCode {
				msg, err := protocols.Decode(bc.code, bc.payload)
				require.NoError(t, err)
				return msg.(*protocols.Preprepare).Proposal
			}
		}
		t.Fatal("no PRE-PREPARE broadcast")
		return nil
	}

	t.Run("uses sidecars carried by ROUND-CHANGE", func(t *testing.T) {
		c, b := newCore(t)
		proposal := run(t, c, b, func(int) *types.Block { return prepared.WithSidecars(sidecars) })
		assert.Equal(t, prepared.Hash(), proposal.Hash())
		assert.Len(t, proposal.Sidecars(), 1)
	})

	t.Run("uses own prepared block when ROUND-CHANGE lacks sidecars", func(t *testing.T) {
		c, b := newCore(t)
		c.current.preparedRound = big.NewInt(preparedRound)
		c.current.preparedBlock = prepared.WithSidecars(sidecars)
		proposal := run(t, c, b, func(int) *types.Block { return prepared })
		assert.Equal(t, prepared.Hash(), proposal.Hash())
		assert.Len(t, proposal.Sidecars(), 1)
	})

	t.Run("uses sidecars from a later ROUND-CHANGE when the first lacks them", func(t *testing.T) {
		c, b := newCore(t)
		proposal := run(t, c, b, func(i int) *types.Block {
			if i == 0 {
				return prepared
			}
			return prepared.WithSidecars(sidecars)
		})
		assert.Equal(t, prepared.Hash(), proposal.Hash())
		assert.Len(t, proposal.Sidecars(), 1)
	})
}

// ##

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
	forgotten  []common.Hash                 // hashes passed to ForgetMessage
	verify     func(istanbul.Proposal) error // optional Verify result
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
func (b *roundChangeTestBackend) Verify(proposal istanbul.Proposal) (time.Duration, error) {
	if b.verify != nil {
		return 0, b.verify(proposal)
	}
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
func (b *roundChangeTestBackend) ForgetMessage(hash common.Hash) {
	b.forgotten = append(b.forgotten, hash)
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
