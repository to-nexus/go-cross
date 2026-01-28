// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package backend

import (
	"crypto/ecdsa"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/core"
	"github.com/ethereum/go-ethereum/consensus/istanbul/engine"
	"github.com/ethereum/go-ethereum/consensus/istanbul/protocols"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/bls"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
)

const (
	// fetcherID is the ID indicates the block is from Istanbul engine
	fetcherID = "istanbul"
)

type Backend struct {
	config *istanbul.Config

	privateKey *ecdsa.PrivateKey
	address    common.Address

	core *core.Core

	engine *engine.Engine

	istanbulEventMux *event.TypeMux

	logger log.Logger

	db ethdb.Database

	chain        consensus.ChainHeaderReader
	currentBlock func() *types.Block
	hasBadBlock  func(db ethdb.Reader, hash common.Hash) bool

	// the channels for istanbul engine notifications
	commitCh          chan *types.Block
	proposedBlockHash common.Hash
	sealMu            sync.Mutex
	coreStarted       bool
	coreMu            sync.RWMutex

	// Current list of candidates we are pushing
	candidates map[common.Address]bool
	// Protects the signer fields
	candidatesLock sync.RWMutex
	// Snapshots for recent block to speed up reorgs
	recents *lru.Cache[common.Hash, *Snapshot]

	// event subscription for ChainHeadEvent event
	broadcaster consensus.IstanbulBroadcaster

	recentMessages *lru.Cache[common.Address, *lru.Cache[common.Hash, bool]] // the cache of peer's messages
	knownMessages  *lru.Cache[common.Hash, bool]                             // the cache of self messages

	blsSecretKey bls.SecretKey // ##CROSS: bls seal
}

// New creates an Ethereum backend for Istanbul core engine.
func New(config *istanbul.Config, privateKey *ecdsa.PrivateKey, blsSecretKey bls.SecretKey, db ethdb.Database, contractBackend bind.ContractBackend) *Backend {
	// Allocate the snapshot caches and create the engine
	recents := lru.NewCache[common.Hash, *Snapshot](inmemorySnapshots)
	recentMessages := lru.NewCache[common.Address, *lru.Cache[common.Hash, bool]](inmemoryPeers)
	knownMessages := lru.NewCache[common.Hash, bool](inmemoryMessages)

	sb := &Backend{
		config:           config,
		istanbulEventMux: new(event.TypeMux),
		privateKey:       privateKey,
		address:          crypto.PubkeyToAddress(privateKey.PublicKey),
		logger:           log.New(),
		db:               db,
		commitCh:         make(chan *types.Block, 1),
		recents:          recents,
		candidates:       make(map[common.Address]bool),
		coreStarted:      false,
		recentMessages:   recentMessages,
		knownMessages:    knownMessages,
		blsSecretKey:     blsSecretKey, // ##CROSS: bls seal
	}

	sb.engine = engine.NewEngine(sb.config, sb.address, sb.Sign, sb.SignTx, sb.SignBLS, sb, contractBackend)
	return sb
}

func (sb *Backend) Engine() *engine.Engine {
	return sb.engine
}

// zekun: HACK
func (sb *Backend) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	return sb.Engine().CalcDifficulty(chain, time, parent)
}

// Address implements istanbul.Backend.Address
func (sb *Backend) Address() common.Address {
	return sb.Engine().Address()
}

// Validators implements istanbul.Backend.Validators
func (sb *Backend) Validators(proposal istanbul.Proposal) istanbul.ValidatorSet {
	return sb.getValidators(proposal.Number().Uint64(), proposal.Hash())
}

// Broadcast implements istanbul.Backend.Broadcast
func (sb *Backend) Broadcast(valSet istanbul.ValidatorSet, code uint64, payload []byte) error {
	// send to others
	sb.Gossip(valSet, code, payload)
	// send to self
	msg := istanbul.MessageEvent{
		Code:    code,
		Payload: payload,
	}
	go sb.istanbulEventMux.Post(msg)
	return nil
}

// Gossip implements istanbul.Backend.Gossip
func (sb *Backend) Gossip(valSet istanbul.ValidatorSet, code uint64, payload []byte) error {
	hash := istanbul.RLPHash(payload)
	sb.knownMessages.Add(hash, true)

	targets := make(map[common.Address]bool)
	for _, val := range valSet.List() {
		if val.Address() != sb.Address() {
			targets[val.Address()] = true
		}
	}
	if sb.broadcaster != nil && len(targets) > 0 {
		ps := sb.broadcaster.FindPeers(targets)
		for addr, p := range ps {
			m, ok := sb.recentMessages.Get(addr)
			if ok {
				if _, ok := m.Get(hash); ok {
					// This peer had this event, skip it
					continue
				}
			} else {
				m = lru.NewCache[common.Hash, bool](inmemoryMessages)
			}

			m.Add(hash, true)
			sb.recentMessages.Add(addr, m)

			var outboundCode uint64 = istanbulMsg
			if _, ok := protocols.MessageCodes()[code]; ok {
				outboundCode = code
			}
			go p.SendIstanbulConsensus(outboundCode, payload)
		}
	}
	return nil
}

// Commit implements istanbul.Backend.Commit
func (sb *Backend) Commit(proposal istanbul.Proposal, seals []istanbul.SignedSeal, round *big.Int) (err error) {
	// Check if the proposal is a valid block
	block, ok := proposal.(*types.Block)
	if !ok {
		sb.logger.Error("Istanbul: invalid block proposal", "proposal", proposal)
		return istanbul.ErrInvalidProposal
	}

	// Commit header
	header := block.Header()
	err = sb.Engine().CommitHeader(sb.chain, header, seals, round)
	if err != nil {
		return
	}

	// Remove ValidatorSet added to ProposerPolicy registry, if not done, the registry keeps increasing size with each block height
	sb.config.ProposerPolicy.ClearRegistry()

	// update block's header
	block = block.WithSeal(header)

	logger := sb.logger.With("author", sb.Address(), "hash", proposal.Hash(), "number", proposal.Number().Uint64())
	// ##CROSS: bls seal
	if sb.chain.Config().IsIstanbulPoSA(header.Number, header.Time) {
		extra, _ := types.ExtractIstanbulExtra(header)
		var committedSeal []byte
		if len(extra.CommittedSeal) > 0 {
			committedSeal = extra.CommittedSeal[0]
		}
		logger = logger.With("committedSeal", common.PrettyBytes(committedSeal), "validators", extra.Validators, "signersBitset", extra.SignersBitset, "signers", extra.Signers)
	}
	// ##
	logger.Info("Istanbul: block proposal committed")

	// - if the proposed and committed blocks are the same, send the proposed hash
	//   to commit channel, which is being watched inside the engine.Seal() function.
	// - otherwise, we try to insert the block.
	// -- if success, the ChainHeadEvent event will be broadcasted, try to build
	//    the next block and the previous Seal() will be stopped.
	// -- otherwise, a error will be returned and a round change event will be fired.
	if sb.proposedBlockHash == block.Hash() {
		// feed block hash to Seal() and wait the Seal() result
		sb.commitCh <- block
		return nil
	}

	if sb.broadcaster != nil {
		sb.broadcaster.Enqueue(fetcherID, block)
	}

	return nil
}

// EventMux implements istanbul.Backend.EventMux
func (sb *Backend) EventMux() *event.TypeMux {
	return sb.istanbulEventMux
}

// Verify implements istanbul.Backend.Verify
func (sb *Backend) Verify(proposal istanbul.Proposal) (time.Duration, error) {
	// Check if the proposal is a valid block
	block, ok := proposal.(*types.Block)
	if !ok {
		sb.logger.Error("Istanbul: invalid block proposal", "proposal", proposal)
		return 0, istanbul.ErrInvalidProposal
	}

	// check bad block
	if sb.HasBadProposal(block.Hash()) {
		sb.logger.Warn("Istanbul: bad block proposal", "proposal", proposal)
		return 0, istanbul.ErrBlacklistedHash
	}

	header := block.Header()
	snap, err := sb.snapshot(sb.chain, header.Number.Uint64()-1, header.ParentHash, nil)
	if err != nil {
		log.Error("Istanbul: failed to get snapshot", "err", err, "number", header.Number.Uint64()-1, "hash", header.ParentHash.String())
		return 0, err
	}

	return sb.Engine().VerifyBlockProposal(sb.chain, block, snap.ValSet)
}

// Sign implements istanbul.Backend.Sign
func (sb *Backend) Sign(data []byte) ([]byte, error) {
	hashData := crypto.Keccak256(data)
	return crypto.Sign(hashData, sb.privateKey)
}

// SignWithoutHashing implements istanbul.Backend.SignWithoutHashing and signs input data with the backend's private key without hashing the input data
func (sb *Backend) SignWithoutHashing(data []byte) ([]byte, error) {
	return crypto.Sign(data, sb.privateKey)
}

// ##CROSS: bls random reveal
// SignBLS signs the data using the BLS secret key
func (sb *Backend) SignBLS(data []byte) ([]byte, error) {
	if sb.blsSecretKey == nil {
		return nil, istanbul.ErrInvalidSecretKey
	}
	return sb.blsSecretKey.Sign(data).Marshal(), nil
}

// ##

// ##CROSS: bls seal
// SignSeal implements istanbul.Backend.SignSeal and signs the seal with the backend's private key
func (sb *Backend) SignSeal(header *types.Header, data []byte) ([]byte, error) {
	// If Istanbul PoSA is active, return the BLS signature
	if sb.chain.Config().IsIstanbulPoSA(header.Number, header.Time) {
		return sb.SignBLS(data)
	}
	// Fallback to the former method
	return sb.SignWithoutHashing(data)
}

// SealSize returns the size of the seal for the given proposal
func (sb *Backend) SealSize(proposal istanbul.Proposal) int {
	if block, ok := proposal.(*types.Block); ok {
		if sb.chain.Config().IsIstanbulPoSA(block.Number(), block.Time()) {
			return types.IstanbulExtraSealBLS
		}
	}
	return types.IstanbulExtraSeal
}

// VerifyCommittedSeal checks if the seal is valid signature by verifying it with the public key of the signer.
func (sb *Backend) VerifyCommittedSeal(signature []byte, signer common.Address, proposal istanbul.Proposal, round uint32, validatorSet istanbul.ValidatorSet) error {
	var header *types.Header
	if block, ok := proposal.(*types.Block); ok {
		header = block.Header()
	} else {
		return istanbul.ErrInvalidProposal
	}

	if !sb.chain.Config().IsIstanbulPoSA(header.Number, header.Time) {
		return nil
	}

	_, validator := validatorSet.GetByAddress(signer)
	if validator == nil {
		return istanbul.ErrUnauthorized
	}

	signerAddr := validator.SignerAddress()
	if signerAddr == (types.BLSPublicKey{}) {
		return istanbul.ErrInvalidPublicKey
	}
	pubkey, err := bls.PublicKeyFromBytes(signerAddr.Bytes())
	if err != nil {
		return err
	}

	sig, err := bls.SignatureFromBytes(signature)
	if err != nil {
		return err
	}
	if !sig.Verify(pubkey, core.PrepareCommittedSeal(header, round)) {
		return istanbul.ErrInvalidSignature
	}
	return nil
}

// ##

// ##CROSS: consensus system contract
// SignTx signs the transaction with the backend's private key
func (sb *Backend) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	signer := types.LatestSignerForChainID(chainID)
	return types.SignTx(tx, signer, sb.privateKey)
}

// ##

// CheckSignature implements istanbul.Backend.CheckSignature
func (sb *Backend) CheckSignature(data []byte, address common.Address, sig []byte) error {
	signer, err := istanbul.GetSignatureAddress(data, sig)
	if err != nil {
		return err
	}
	// Compare derived addresses
	if signer != address {
		return istanbul.ErrInvalidSignature
	}

	return nil
}

// HasPropsal implements istanbul.Backend.HashBlock
func (sb *Backend) HasPropsal(hash common.Hash, number *big.Int) bool {
	return sb.chain.GetHeader(hash, number.Uint64()) != nil
}

// GetProposer implements istanbul.Backend.GetProposer
func (sb *Backend) GetProposer(number uint64) common.Address {
	if h := sb.chain.GetHeaderByNumber(number); h != nil {
		a, _ := sb.Author(h)
		return a
	}
	return common.Address{}
}

// ParentValidators implements istanbul.Backend.GetParentValidators
func (sb *Backend) ParentValidators(proposal istanbul.Proposal) istanbul.ValidatorSet {
	if block, ok := proposal.(*types.Block); ok {
		return sb.getValidators(block.Number().Uint64()-1, block.ParentHash())
	}
	return validator.NewSet(nil, nil, sb.config.GetConfig(proposal.Number()).ProposerPolicy) // ##CROSS: istanbul param
}

func (sb *Backend) getValidators(number uint64, hash common.Hash) istanbul.ValidatorSet {
	snap, err := sb.snapshot(sb.chain, number, hash, nil)
	if err != nil {
		return validator.NewSet(nil, nil, sb.config.GetConfig(new(big.Int).SetUint64(number)).ProposerPolicy) // ##CROSS: istanbul param
	}
	return snap.ValSet
}

func (sb *Backend) LastProposal() (istanbul.Proposal, common.Address) {
	block := sb.currentBlock()

	var proposer common.Address
	if block.Number().Cmp(common.Big0) > 0 {
		var err error
		proposer, err = sb.Author(block.Header())
		if err != nil {
			sb.logger.Error("Istanbul: last block proposal invalid", "err", err)
			return nil, common.Address{}
		}
	}

	// Return header only block here since we don't need block body
	return block, proposer
}

func (sb *Backend) HasBadProposal(hash common.Hash) bool {
	if sb.hasBadBlock == nil {
		return false
	}
	return sb.hasBadBlock(sb.db, hash)
}

func (sb *Backend) Close() error {
	return nil
}

func (sb *Backend) startIstanbul() error {
	sb.logger.Info("Istanbul: activate")

	// ##CROSS: bls seal
	if sb.chain != nil && sb.chain.Config().Istanbul != nil && sb.chain.Config().Istanbul.PoSAActivationSeconds != nil && sb.blsSecretKey == nil {
		sb.logger.Warn("Istanbul: BLS secret key is not configured - node cannot participate in PoSA")
	}
	// ##

	sb.logger.Trace("Istanbul: set ProposerPolicy sorter to ValidatorSortByByteFunc")
	sb.config.ProposerPolicy.Use(istanbul.ValidatorSortByByte())

	sb.core = core.New(sb, sb.config)
	if err := sb.core.Start(); err != nil {
		sb.logger.Error("Istanbul: failed to activate", "err", err)
		return err
	}

	return nil
}

func (sb *Backend) stop() error {
	core := sb.core
	sb.core = nil

	if core != nil {
		sb.logger.Info("Istanbul: deactivate")
		if err := core.Stop(); err != nil {
			sb.logger.Error("Istanbul: failed to deactivate", "err", err)
			return err
		}
	}

	return nil
}

// ##CROSS: istanbul posa
// ValidatorsAt returns validators from the snapshot at the given block number.
func (sb *Backend) ValidatorsAt(chain consensus.ChainHeaderReader, header *types.Header) []common.Address {
	if chain == nil || header == nil {
		return nil
	}
	snap, err := sb.snapshot(chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil
	}
	return snap.validators()
}

// ##
