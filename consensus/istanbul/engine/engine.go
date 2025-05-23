package engine

import (
	"bytes"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	"golang.org/x/crypto/sha3"
)

var (
	nilUncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.
)

type SignerFn func(data []byte) ([]byte, error)

type Engine struct {
	cfg *istanbul.Config

	signer common.Address // Ethereum address of the signing key
	sign   SignerFn       // Signer function to authorize hashes with
}

func NewEngine(cfg *istanbul.Config, signer common.Address, sign SignerFn) *Engine {
	return &Engine{
		cfg:    cfg,
		signer: signer,
		sign:   sign,
	}
}

func (e *Engine) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (e *Engine) CommitHeader(header *types.Header, seals [][]byte, round *big.Int) error {
	return ApplyHeaderIstanbulExtra(
		header,
		writeCommittedSeals(seals),
		writeRoundNumber(round),
	)
}

// writeCommittedSeals writes the extra-data field of a block header with given committed seals.
func writeCommittedSeals(committedSeals [][]byte) ApplyExtra {
	return func(extra *types.IstanbulExtra) error {
		if len(committedSeals) == 0 {
			return istanbul.ErrInvalidCommittedSeals
		}

		for _, seal := range committedSeals {
			if len(seal) != types.IstanbulExtraSeal {
				return istanbul.ErrInvalidCommittedSeals
			}
		}

		extra.CommittedSeal = make([][]byte, len(committedSeals))
		copy(extra.CommittedSeal, committedSeals)

		return nil
	}
}

// writeRoundNumber writes the extra-data field of a block header with given round.
func writeRoundNumber(round *big.Int) ApplyExtra {
	return func(extra *types.IstanbulExtra) error {
		extra.Round = uint32(round.Uint64())
		return nil
	}
}

func (e *Engine) VerifyBlockProposal(chain consensus.ChainHeaderReader, block *types.Block, validators istanbul.ValidatorSet) (time.Duration, error) {
	// check block body
	txnHash := types.DeriveSha(block.Transactions(), trie.NewStackTrie(nil))
	if txnHash != block.Header().TxHash {
		return 0, istanbul.ErrMismatchTxhashes
	}

	uncleHash := types.CalcUncleHash(block.Uncles())
	if uncleHash != nilUncleHash {
		return 0, istanbul.ErrInvalidUncleHash
	}

	// verify the header of proposed block
	err := e.VerifyHeader(chain, block.Header(), nil, validators)
	if err == nil || err == istanbul.ErrEmptyCommittedSeals {
		// ignore errEmptyCommittedSeals error because we don't have the committed seals yet
		return 0, nil
	} else if err == consensus.ErrFutureBlock {
		return time.Until(time.Unix(int64(block.Header().Time), 0)), consensus.ErrFutureBlock
	}

	parentHeader := chain.GetHeaderByHash(block.ParentHash())
	config := e.cfg.GetConfig(parentHeader.Number)

	if config.EmptyBlockPeriod > config.BlockPeriod && len(block.Transactions()) == 0 {
		if block.Header().Time < parentHeader.Time+config.EmptyBlockPeriod {
			return 0, fmt.Errorf("empty block verification fail")
		}
	}

	return 0, err
}

func (e *Engine) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header, validators istanbul.ValidatorSet) error {
	return e.verifyHeader(chain, header, parents, validators)
}

// verifyHeader checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (e *Engine) verifyHeader(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header, validators istanbul.ValidatorSet) error {
	if header.Number == nil {
		return istanbul.ErrUnknownBlock
	}

	// Don't waste time checking blocks from the future (adjusting for allowed threshold)
	adjustedTimeNow := time.Now().Add(time.Duration(e.cfg.AllowedFutureBlockTime) * time.Second).Unix()
	if header.Time > uint64(adjustedTimeNow) {
		return consensus.ErrFutureBlock
	}

	if _, err := types.ExtractIstanbulExtra(header); err != nil {
		return istanbul.ErrInvalidExtraDataFormat
	}

	// Ensure that the mix digest is zero as we don't have fork protection currently
	if !types.IsIstanbulDigest(header.MixDigest) {
		return istanbul.ErrInvalidMixDigest
	}

	// Ensure that the block doesn't contain any uncles which are meaningless in Istanbul
	if header.UncleHash != nilUncleHash {
		return istanbul.ErrInvalidUncleHash
	}

	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if header.Difficulty == nil || header.Difficulty.Cmp(istanbul.DefaultDifficulty) != 0 {
		return istanbul.ErrInvalidDifficulty
	}

	return e.verifyCascadingFields(chain, header, validators, parents)
}

func (e *Engine) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool, validators istanbul.ValidatorSet) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))
	go func() {
		errored := false
		for i, header := range headers {
			var err error
			if errored {
				err = consensus.ErrUnknownAncestor
			} else {
				err = e.verifyHeader(chain, header, headers[:i], validators)
			}

			if err != nil {
				errored = true
			}

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

// verifyCascadingFields verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (e *Engine) verifyCascadingFields(chain consensus.ChainHeaderReader, header *types.Header, validators istanbul.ValidatorSet, parents []*types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}

	// Check parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}

	// Ensure that the block's parent has right number and hash
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}

	// Ensure that the block's timestamp isn't too close to it's parent
	// When the BlockPeriod is reduced it is reduced for the proposal.
	// e.g when blockperiod is 1 from block 10 the block period between 9 and 10 is 1
	if parent.Time+e.cfg.GetConfig(header.Number).BlockPeriod > header.Time {
		return istanbul.ErrInvalidTimestamp
	}

	// Verify signer
	if err := e.verifySigner(chain, header, parents, validators); err != nil {
		return err
	}

	return e.verifyCommittedSeals(chain, header, parents, validators)
}

func (e *Engine) verifySigner(chain consensus.ChainHeaderReader, header *types.Header, _ []*types.Header, validators istanbul.ValidatorSet) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return istanbul.ErrUnknownBlock
	}

	// Resolve the authorization key and check against signers
	signer, err := e.Author(header)
	if err != nil {
		return err
	}

	// Signer should be in the validator set of previous block's extraData.
	if _, v := validators.GetByAddress(signer); v == nil {
		return istanbul.ErrUnauthorized
	}

	if extra, err := types.ExtractIstanbulExtra(header); err != nil {
		return err
	} else {
		signature := extra.RandomReveal
		if pub, err := crypto.SigToPub(crypto.Keccak256(header.Number.Bytes(), chain.Config().ChainID.Bytes()), signature); err != nil {
			return err
		} else if crypto.PubkeyToAddress(*pub) != signer {
			return istanbul.ErrMismatchedRandomSignature
		} else if header.MixDigest != types.MakeIstanbulDigest(crypto.Keccak256Hash(signature)) {
			return istanbul.ErrInvalidMixDigest
		}
	}

	return nil
}

// verifyCommittedSeals checks whether every committed seal is signed by one of the parent's validators
func (e *Engine) verifyCommittedSeals(_ consensus.ChainHeaderReader, header *types.Header, _ []*types.Header, validators istanbul.ValidatorSet) error {
	number := header.Number.Uint64()

	if number == 0 {
		// We don't need to verify committed seals in the genesis block
		return nil
	}

	extra, err := types.ExtractIstanbulExtra(header)
	if err != nil {
		return err
	}
	committedSeal := extra.CommittedSeal

	// The length of Committed seals should be larger than 0
	if len(committedSeal) == 0 {
		return istanbul.ErrEmptyCommittedSeals
	}

	validatorsCpy := validators.Copy()

	// Check whether the committed seals are generated by validators
	validSeal := 0
	committers, err := e.Signers(header)
	if err != nil {
		return err
	}

	for _, addr := range committers {
		if validatorsCpy.RemoveValidator(addr) {
			validSeal++
			continue
		}
		return istanbul.ErrInvalidCommittedSeals
	}

	if validSeal < validators.QuorumSize() {
		return istanbul.ErrInvalidCommittedSeals
	}

	return nil
}

// VerifyUncles verifies that the given block's uncles conform to the consensus
// rules of a given engine.
func (e *Engine) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return istanbul.ErrInvalidUncleHash
	}
	return nil
}

// VerifySeal checks whether the crypto seal on a header is valid according to
// the consensus rules of the given engine.
func (e *Engine) VerifySeal(chain consensus.ChainHeaderReader, header *types.Header, validators istanbul.ValidatorSet) error {
	// get parent header and ensure the signer is in parent's validator set
	number := header.Number.Uint64()
	if number == 0 {
		return istanbul.ErrUnknownBlock
	}

	// ensure that the difficulty equals to istanbul.DefaultDifficulty
	if header.Difficulty.Cmp(istanbul.DefaultDifficulty) != 0 {
		return istanbul.ErrInvalidDifficulty
	}

	return e.verifySigner(chain, header, nil, validators)
}

func (e *Engine) Prepare(chain consensus.ChainHeaderReader, header *types.Header, validators istanbul.ValidatorSet) error {
	header.Coinbase = common.Address{}
	header.Nonce = istanbul.EmptyBlockNonce

	// make mixdigest
	signed, err := e.sign(append(header.Number.Bytes(), chain.Config().ChainID.Bytes()...))
	if err != nil {
		return err
	} else {
		header.MixDigest = types.MakeIstanbulDigest(crypto.Keccak256Hash(signed)) // ##CROSS: istanbul digest
	}

	// copy the parent extra data as the header extra data
	number := header.Number.Uint64()

	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	// use the same difficulty for all blocks
	header.Difficulty = istanbul.DefaultDifficulty

	// set header's timestamp
	header.Time = parent.Time + e.cfg.GetConfig(header.Number).BlockPeriod
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}

	// currentBlockNumber := big.NewInt(0).SetUint64(number - 1)
	// for _, transition := range e.cfg.Transitions {
	// 	if transition.Block.Cmp(currentBlockNumber) == 0 && len(transition.Validators) > 0 {
	// 		toRemove := make([]istanbul.Validator, 0, validators.Size())
	// 		l := validators.List()
	// 		toRemove = append(toRemove, l...)
	// 		for i := range toRemove {
	// 			validators.RemoveValidator(toRemove[i].Address())
	// 		}
	// 		for i := range transition.Validators {
	// 			validators.AddValidator(transition.Validators[i])
	// 		}
	// 		break
	// 	}
	// }
	validatorsList := validator.SortedAddresses(validators.List())
	// add validators in snapshot to extraData's validators section
	return ApplyHeaderIstanbulExtra(
		header,
		WriteValidators(validatorsList),
		WriteRandomReveal(signed),
	)
}

func WriteValidators(validators []common.Address) ApplyExtra {
	return func(extra *types.IstanbulExtra) error {
		extra.Validators = validators
		return nil
	}
}

func WriteRandomReveal(signature []byte) ApplyExtra {
	return func(extra *types.IstanbulExtra) error {
		if len(signature) != 65 {
			return istanbul.ErrInvalidSignature
		}
		extra.RandomReveal = signature
		return nil
	}
}

// Finalize runs any post-transaction state modifications (e.g. block rewards)
// and assembles the final block.
//
// Note, the block header and state database might be updated to reflect any
// consensus rules that happen at finalization (e.g. block rewards).
func (e *Engine) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	// Accumulate any block and uncle rewards and commit the final state root
	e.accumulateRewards(chain, state, header)

	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = nilUncleHash
}

// FinalizeAndAssemble implements consensus.Engine, ensuring no uncles are set,
// nor block rewards given, and returns the final block.
func (e *Engine) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	e.Finalize(chain, header, state, txs, uncles)
	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts, trie.NewStackTrie(nil)), nil
}

// Seal generates a new block for the given input block with the local miner's
// seal place on top.
func (e *Engine) Seal(chain consensus.ChainHeaderReader, block *types.Block, validators istanbul.ValidatorSet) (*types.Block, error) {
	if _, v := validators.GetByAddress(e.signer); v == nil {
		return block, istanbul.ErrUnauthorized
	}

	header := block.Header()
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		return block, consensus.ErrUnknownAncestor
	}

	// Set Coinbase
	header.Coinbase = e.signer

	return block.WithSeal(header), nil
}

func (e *Engine) SealHash(header *types.Header) common.Hash {
	header.Coinbase = e.signer
	return sigHash(header)
}

func (e *Engine) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	return new(big.Int)
}

func (e *Engine) ExtractGenesisValidators(header *types.Header) ([]common.Address, error) {
	extra, err := types.ExtractIstanbulExtra(header)
	if err != nil {
		return nil, err
	}

	return extra.Validators, nil
}

func (e *Engine) Signers(header *types.Header) ([]common.Address, error) {
	extra, err := types.ExtractIstanbulExtra(header)
	if err != nil {
		return []common.Address{}, err
	}
	committedSeal := extra.CommittedSeal
	proposalSeal := PrepareCommittedSeal(header, extra.Round)

	var addrs []common.Address
	// 1. Get committed seals from current header
	for _, seal := range committedSeal {
		// 2. Get the original address by seal and parent block hash
		addr, err := istanbul.GetSignatureAddressNoHashing(proposalSeal, seal)
		if err != nil {
			return nil, istanbul.ErrInvalidSignature
		}
		addrs = append(addrs, addr)
	}

	return addrs, nil
}

func (e *Engine) Address() common.Address {
	return e.signer
}

// FIXME: Need to update this for Istanbul
// sigHash returns the hash which is used as input for the Istanbul
// signing. It is the hash of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// Note, the method requires the extra data to be at least 65 bytes, otherwise it
// panics. This is done to avoid accidentally using both forms (signature present
// or not), which could be abused to produce different hashes for the same header.
func sigHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	rlp.Encode(hasher, types.IstanbulFilteredHeader(header))
	hasher.Sum(hash[:0])
	return hash
}

// PrepareCommittedSeal returns a committed seal for the given hash
func PrepareCommittedSeal(header *types.Header, round uint32) []byte {
	h := types.CopyHeader(header)
	return h.IstanbulHashWithRoundNumber(round).Bytes()
}

func (e *Engine) WriteVote(header *types.Header, candidate common.Address, authorize bool) error {
	return ApplyHeaderIstanbulExtra(
		header,
		WriteVote(candidate, authorize),
	)
}

func WriteVote(candidate common.Address, authorize bool) ApplyExtra {
	return func(extra *types.IstanbulExtra) error {
		voteType := types.IstanbulDropVote
		if authorize {
			voteType = types.IstanbulAuthVote
		}

		vote := &types.ValidatorVote{RecipientAddress: candidate, VoteType: voteType}
		extra.Vote = vote
		return nil
	}
}

func (e *Engine) ReadVote(header *types.Header) (candidate common.Address, authorize bool, err error) {
	extra, err := getExtra(header)
	if err != nil {
		return common.Address{}, false, err
	}

	var vote *types.ValidatorVote
	if extra.Vote == nil {
		vote = &types.ValidatorVote{RecipientAddress: common.Address{}, VoteType: types.IstanbulDropVote}
	} else {
		vote = extra.Vote
	}

	// Tally up the new vote from the validator
	switch {
	case vote.VoteType == types.IstanbulAuthVote:
		authorize = true
	case vote.VoteType == types.IstanbulDropVote:
		authorize = false
	default:
		return common.Address{}, false, istanbul.ErrInvalidVote
	}

	return vote.RecipientAddress, authorize, nil
}

func getExtra(header *types.Header) (*types.IstanbulExtra, error) {
	if len(header.Extra) < types.IstanbulExtraVanity {
		// In this scenario, the header extradata only contains client specific information, hence create a new extra and set vanity
		vanity := append(header.Extra, bytes.Repeat([]byte{0x00}, types.IstanbulExtraVanity-len(header.Extra))...)
		return &types.IstanbulExtra{
			VanityData:    vanity,
			Validators:    []common.Address{},
			CommittedSeal: [][]byte{},
			Round:         0,
			Vote:          nil,
		}, nil
	}

	// This is the case when Extra has already been set
	return types.ExtractIstanbulExtra(header)
}

func setExtra(h *types.Header, extra *types.IstanbulExtra) error {
	payload, err := rlp.EncodeToBytes(extra)
	if err != nil {
		return err
	}

	h.Extra = payload
	return nil
}

// AccumulateRewards credits the beneficiary of the given block with a reward.
func (e *Engine) accumulateRewards(chain consensus.ChainHeaderReader, state *state.StateDB, header *types.Header) {

}
