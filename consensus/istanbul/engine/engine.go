package engine

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/holiman/uint256"
	"golang.org/x/crypto/sha3"
)

var (
	nilUncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.
)

type SignerFn func(data []byte) ([]byte, error)
type SignerTxFn func(accounts.Account, *types.Transaction, *big.Int) (*types.Transaction, error)

type Engine struct {
	cfg *istanbul.Config

	signer common.Address // Ethereum address of the signing key
	sign   SignerFn       // Signer function to authorize hashes with

	// ##CROSS: consensus system contract
	signTx       SignerTxFn
	consensus    consensus.Engine
	ethAPI       *ethapi.BlockChainAPI
	validatorSet abi.ABI
	stakeHub     abi.ABI
	// ##
}

// ##CROSS: consensus system contract
var systemContracts = map[common.Address]bool{
	contracts.ValidatorSetAddr: true,
	contracts.StakeHubAddr:     true,
}

// ##

func NewEngine(cfg *istanbul.Config, signer common.Address, sign SignerFn, signTx SignerTxFn, ce consensus.Engine, ethAPI *ethapi.BlockChainAPI) *Engine {
	return &Engine{
		cfg:    cfg,
		signer: signer,
		sign:   sign,
		// ##CROSS: consensus system contract
		signTx:       signTx,
		consensus:    ce,
		ethAPI:       ethAPI,
		validatorSet: contracts.ValidatorSetABI(),
		stakeHub:     contracts.StakeHubABI(),
		// ##
	}
}

// ##CROSS: consensus system contract
func (e *Engine) IsSystemTransaction(tx *types.Transaction, header *types.Header) (bool, error) {
	if to := tx.To(); to == nil || !isToSystemContract(*to) {
		return false, nil
	}
	if tx.GasPrice().Sign() != 0 {
		return false, nil
	}
	sender, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		return false, err
	}
	return sender == header.Coinbase, nil
}

func (e *Engine) IsSystemContract(to *common.Address) bool {
	if to == nil {
		return false
	}
	return isToSystemContract(*to)
}

func isToSystemContract(to common.Address) bool {
	return systemContracts[to]
}

// ##

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

	// Verify the existence / non-existence of cancun-specific header fields
	cancun := chain.Config().IsCancun(header.Number, header.Time)
	if !cancun {
		switch {
		case header.ExcessBlobGas != nil:
			return fmt.Errorf("invalid excessBlobGas: have %d, expected nil", header.ExcessBlobGas)
		case header.BlobGasUsed != nil:
			return fmt.Errorf("invalid blobGasUsed: have %d, expected nil", header.BlobGasUsed)
		case header.WithdrawalsHash != nil && header.Number.Uint64() > 0:
			// ##CROSS: fork
			// Withdrawals root was at genesis but not in blocks between genesis and cancun activation.
			// So we check if withdrawals root exists in block between genesis and cancun activation.
			return fmt.Errorf("invalid withdrawalsHash: have %x, expected nil", header.WithdrawalsHash)
		case header.ParentBeaconRoot != nil:
			return fmt.Errorf("invalid parentBeaconRoot, have %#x, expected nil", header.ParentBeaconRoot)
		}
	} else {
		switch {
		case !header.EmptyWithdrawalsHash():
			return errors.New("header has wrong WithdrawalsHash")
		case header.ParentBeaconRoot == nil || *header.ParentBeaconRoot != (common.Hash{}):
			return fmt.Errorf("invalid parentBeaconRoot, have %#x, expected zero hash", header.ParentBeaconRoot)
		}
	}

	prague := chain.Config().IsPrague(header.Number, header.Time)
	if !prague {
		if header.RequestsHash != nil {
			return fmt.Errorf("invalid RequestsHash, have %#x, expected nil", header.RequestsHash)
		}
	} else {
		if header.RequestsHash == nil {
			return errors.New("header has nil RequestsHash after Prague")
		}
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

	// Verify that the gas limit is <= 2^63-1
	if header.GasLimit > params.MaxGasLimit {
		return fmt.Errorf("invalid gasLimit: have %v, max %v", header.GasLimit, params.MaxGasLimit)
	}
	// Verify that the gasUsed is <= gasLimit
	if header.GasUsed > header.GasLimit {
		return fmt.Errorf("invalid gasUsed: have %d, gasLimit %d", header.GasUsed, header.GasLimit)
	}

	// Verify the header's EIP-1559 attributes.
	if err := eip1559.VerifyEIP1559Header(chain.Config(), parent, header); err != nil {
		return err
	}
	if chain.Config().IsCancun(header.Number, header.Time) {
		// Verify the header's EIP-4844 attributes.
		if err := eip4844.VerifyEIP4844Header(chain.Config(), parent, header); err != nil {
			return err
		}
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

	var validatorsList []common.Address
	if chain.Config().IsCrossway(header.Number, header.Time) {
		// ##CROSS: consensus system contract
		// after crossway, validator list is managed by the system contract
		vlist, err := e.getValidators(rpc.BlockNumberOrHashWithHash(header.ParentHash, false))
		if err != nil {
			return err
		}
		if vlist != nil {
			validatorsList = vlist
		} else {
			validatorsList = validator.SortedAddresses(validators.List())
		}
		// ##
	} else {
		validatorsList = validator.SortedAddresses(validators.List())
	}

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
//
// All system transactions will be consumed in this function.
func (e *Engine) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state vm.StateDB, body *types.Body, txs *[]*types.Transaction, receipts *types.Receipts, systemTxs *[]*types.Transaction, usedGas *uint64, tracer *tracing.Hooks) error {
	// ##CROSS: contract upgrade
	cx := &chainContext{Chain: chain, engine: e.consensus}
	parent := chain.GetHeaderByHash(header.ParentHash)
	if parent == nil {
		return errors.New("parent not found")
	}

	upgraded := contracts.TryUpdateSystemContract(chain.Config(), header, parent.Time, state)
	if upgraded {
		// consume system transactions
		initData := contracts.InitSystemContract(chain.Config(), header, parent.Time)
		if len(initData) > 0 {
			// all initData should be included in systemTxs, in same order
			if len(*systemTxs) < len(initData) {
				return errors.New("systemTxs is not enough")
			}

			// split transactions into common and system txs
			for _, data := range initData {
				log.Info("Finalize: initializing system contract", "blockNumber", header.Number.Uint64(), "contract", data.To, "data", common.Bytes2Hex(data.Data))
				msg := newSystemMessage(header.Coinbase, data.To, data.Data)
				err := e.applyReceivedSystemTransaction(msg, state, header, cx, txs, (*[]*types.Receipt)(receipts), systemTxs, usedGas, tracer)
				if err != nil {
					return err
				}
			}
		}
	}

	// TODO: slash validator

	// ##

	// Accumulate any block and uncle rewards and commit the final state root
	// e.accumulateRewards(chain, state, header)

	// ##CROSS: consensus system contract
	if chain.Config().IsCrossway(header.Number, header.Time) {
		// log.Info("Finalize: updating validator set", "blockNumber", header.Number.Uint64())
		if err := e.updateValidatorSet(header, state, cx, txs, (*[]*types.Receipt)(receipts), systemTxs, usedGas, tracer); err != nil {
			log.Error("Failed to update validator set", "error", err, "blockNumber", header.Number.Uint64())
			return err
		}
	}
	// ##

	return nil
}

// FinalizeAndAssemble implements consensus.Engine, ensuring no uncles are set,
// nor block rewards given, and returns the final block.
//
// All system transactions will be created and executed in this function.
func (e *Engine) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, body *types.Body, receipts []*types.Receipt, tracer *tracing.Hooks) (*types.Block, []*types.Receipt, error) {
	// ##CROSS: fork
	if body != nil {
		if body.Transactions == nil {
			body.Transactions = make([]*types.Transaction, 0)
		}
	}
	if receipts == nil {
		receipts = make([]*types.Receipt, 0)
	}
	// ##

	// ##CROSS: contract upgrade
	cx := &chainContext{Chain: chain, engine: e.consensus}
	parent := chain.GetHeaderByHash(header.ParentHash)
	if parent == nil {
		return nil, nil, errors.New("parent not found")
	}
	upgraded := contracts.TryUpdateSystemContract(chain.Config(), header, parent.Time, state)
	if upgraded {
		// init upgraded contracts
		initData := contracts.InitSystemContract(chain.Config(), header, parent.Time)
		for _, data := range initData {
			msg := newSystemMessage(e.signer, data.To, data.Data)
			log.Info("FinalizeAndAssemble: initializing system contract", "blockNumber", header.Number.Uint64(), "contract", data.To, "data", common.Bytes2Hex(data.Data))
			err := e.applyGeneratedSystemTransaction(msg, state, header, cx, &body.Transactions, &receipts, &header.GasUsed, tracer)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	// TODO: slash validator
	// TODO: distribute rewards
	// ##

	// ##CROSS: consensus system contract
	if chain.Config().IsCrossway(header.Number, header.Time) && e.cfg.OnNewEpoch(header.Number) {
		// log.Info("FinalizeAndAssemble: updating validator set", "blockNumber", header.Number.Uint64())
		if err := e.updateValidatorSet(header, state, cx, &body.Transactions, &receipts, nil, &header.GasUsed, tracer); err != nil {
			log.Error("Failed to update validator set", "error", err, "blockNumber", header.Number.Uint64())
			return nil, nil, err
		}
	}
	// ##

	// should not happen. Once happen, stopping the node is better than broadcasting the block
	if header.GasLimit < header.GasUsed {
		return nil, nil, errors.New("gas consumption of system txs exceed the gas limit")
	}

	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = nilUncleHash

	// Assemble and return the final block for sealing
	blk := types.NewBlock(header, body, receipts, trie.NewStackTrie(nil))
	return blk, receipts, nil
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

func (e *Engine) ExtractValidators(header *types.Header) ([]common.Address, error) {
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
func (e *Engine) accumulateRewards(chain consensus.ChainHeaderReader, state vm.StateDB, header *types.Header) {
}

// ##CROSS: consensus system contract

// applyGeneratedSystemTransaction applies a system transaction to the state.
// It creates a system transaction from the given message, signs it using the engine's signer function
// and applies it to the state. The created transaction and receipt is appended to the given slices.
//
// This function is called to generate system transactions and append them to the transaction list of the new block.
func (e *Engine) applyGeneratedSystemTransaction(
	msg *core.Message,
	state vm.StateDB,
	header *types.Header,
	chainContext core.ChainContext,
	txs *[]*types.Transaction,
	receipts *[]*types.Receipt,
	usedGas *uint64,
	tracer *tracing.Hooks,
) (err error) {
	nonce := state.GetNonce(msg.From)
	tx := types.NewTransaction(nonce, *msg.To, msg.Value, msg.GasLimit, msg.GasPrice, msg.Data)
	tx, err = e.signTx(accounts.Account{Address: msg.From}, tx, chainContext.Config().ChainID)
	if err != nil {
		return err
	}

	return e.applySystemTransaction(msg, tx, state, header, chainContext, e.signer, txs, receipts, usedGas, tracer)
}

// applyReceivedSystemTransaction applies a received system transaction to the state.
// It assumes that the given message is equal to the first element of systemTxs and applies it to the state.
// The applied transaction and receipt is appended to the given slices and removes the transaction from systemTxs.
//
// This function is called to apply system transactions of the received block to the state.
func (e *Engine) applyReceivedSystemTransaction(
	msg *core.Message,
	state vm.StateDB,
	header *types.Header,
	chainContext core.ChainContext,
	txs *[]*types.Transaction,
	receipts *[]*types.Receipt,
	systemTxs *[]*types.Transaction,
	usedGas *uint64,
	tracer *tracing.Hooks,
) error {
	signer := types.LatestSignerForChainID(chainContext.Config().ChainID)
	nonce := state.GetNonce(msg.From)
	expectedTx := types.NewTransaction(nonce, *msg.To, msg.Value, msg.GasLimit, msg.GasPrice, msg.Data)
	expectedHash := signer.Hash(expectedTx)

	if systemTxs == nil || len(*systemTxs) == 0 || (*systemTxs)[0] == nil {
		return errors.New("supposed to get a actual transaction, but get none")
	}
	actualTx := (*systemTxs)[0]
	if !bytes.Equal(signer.Hash(actualTx).Bytes(), expectedHash.Bytes()) {
		return fmt.Errorf("expected tx hash %v, get %v, nonce %d, to %s, value %s, gas %d, gasPrice %s, data %s",
			expectedHash.String(),
			actualTx.Hash().String(),
			expectedTx.Nonce(),
			expectedTx.To().String(),
			expectedTx.Value().String(),
			expectedTx.Gas(),
			expectedTx.GasPrice().String(),
			hexutil.Bytes(expectedTx.Data()).String(),
		)
	}
	// move to next
	*systemTxs = (*systemTxs)[1:]

	return e.applySystemTransaction(msg, actualTx, state, header, chainContext, header.Coinbase, txs, receipts, usedGas, tracer)
}

func (e *Engine) applySystemTransaction(
	msg *core.Message,
	tx *types.Transaction,
	state vm.StateDB,
	header *types.Header,
	chainContext core.ChainContext,
	author common.Address,
	txs *[]*types.Transaction,
	receipts *[]*types.Receipt,
	usedGas *uint64,
	tracer *tracing.Hooks,
) (err error) {
	state.SetTxContext(tx.Hash(), len(*txs))

	context := core.NewEVMBlockContext(header, chainContext, &author)
	evm := vm.NewEVM(context, state, chainContext.Config(), vm.Config{Tracer: tracer})
	evm.SetTxContext(core.NewEVMTxContext(msg))
	evm.StateDB.AddAddressToAccessList(*msg.To)

	var receipt *types.Receipt
	if tracer != nil {
		onSystemCallStart(tracer, evm.GetVMContext())
		if tracer.OnTxStart != nil {
			tracer.OnTxStart(evm.GetVMContext(), tx, msg.From)
		}
		if tracer.OnSystemCallEnd != nil {
			defer tracer.OnSystemCallEnd()
		}
		if tracer.OnTxEnd != nil {
			defer func() { tracer.OnTxEnd(receipt, err) }()
		}
	}

	gasUsed, err := applySystemMessage(msg, evm, state, header)
	if err != nil {
		return err
	}
	*txs = append(*txs, tx)

	// Update the state with pending changes.
	var root []byte
	if evm.ChainConfig().IsByzantium(header.Number) {
		state.Finalise(true)
	} else {
		root = state.IntermediateRoot(evm.ChainConfig().IsEIP158(header.Number)).Bytes()
	}
	*usedGas += gasUsed

	if state.AccessEvents() != nil {
		state.AccessEvents().Merge(evm.AccessEvents)
	}

	receipt = types.NewReceipt(root, false, *usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = gasUsed
	receipt.Logs = state.GetLogs(tx.Hash(), header.Number.Uint64(), header.Hash())
	receipt.Bloom = types.CreateBloom(receipt)
	receipt.BlockHash = header.Hash()
	receipt.BlockNumber = header.Number
	receipt.TransactionIndex = uint(state.TxIndex())
	*receipts = append(*receipts, receipt)
	return
}

func applySystemMessage(msg *core.Message, evm *vm.EVM, state vm.StateDB, header *types.Header) (uint64, error) {
	// Apply the transaction to the current state (included in the env)
	if evm.ChainConfig().IsCancun(header.Number, header.Time) {
		rules := evm.ChainConfig().Rules(evm.Context.BlockNumber, evm.Context.Random != nil, evm.Context.Time)
		state.Prepare(rules, msg.From, evm.Context.Coinbase, msg.To, vm.ActivePrecompiles(rules), msg.AccessList)
	}
	// Increment the nonce for the next transaction
	state.SetNonce(msg.From, state.GetNonce(msg.From)+1, tracing.NonceChangeEoACall)

	ret, returnGas, err := evm.Call(
		msg.From,
		*msg.To,
		msg.Data,
		msg.GasLimit,
		uint256.MustFromBig(msg.Value),
	)
	if err != nil {
		log.Error("apply message failed", "msg", string(ret), "err", err)
	}
	return msg.GasLimit - returnGas, err
}

func newSystemMessage(from, toAddress common.Address, data []byte) *core.Message {
	return &core.Message{
		From:     from,
		GasLimit: math.MaxUint64 / 2,
		GasPrice: big.NewInt(0),
		Value:    big.NewInt(0),
		To:       &toAddress,
		Data:     data,
	}
}

func onSystemCallStart(tracer *tracing.Hooks, ctx *tracing.VMContext) {
	if tracer.OnSystemCallStartV2 != nil {
		tracer.OnSystemCallStartV2(ctx)
	} else if tracer.OnSystemCallStart != nil {
		tracer.OnSystemCallStart()
	}
}

// chain context
type chainContext struct {
	Chain  consensus.ChainHeaderReader
	engine consensus.Engine
}

func (c chainContext) Engine() consensus.Engine {
	return c.engine
}

func (c chainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	return c.Chain.GetHeader(hash, number)
}

func (c chainContext) Config() *params.ChainConfig {
	return c.Chain.Config()
}

// ##
