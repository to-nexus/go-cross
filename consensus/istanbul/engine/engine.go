package engine

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/engine/bls"
	ecdsaengine "github.com/ethereum/go-ethereum/consensus/istanbul/engine/ecdsa"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

func NewEngine(cfg *istanbul.Config, useBls bool, privateKey *ecdsa.PrivateKey) Engine {
	if useBls {
		engine, err := bls.NewEngine(cfg, privateKey)
		if err != nil {
			panic(err)
		}
		return engine
	}
	return ecdsaengine.NewEngine(cfg, privateKey)
}

type Engine interface {
	Address() common.Address
	Author(header *types.Header) (common.Address, error)
	Signers(header *types.Header) ([]common.Address, error)

	CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int

	VerifyBlockProposal(chain consensus.ChainHeaderReader, block *types.Block, validators istanbul.ValidatorSet) (time.Duration, error)
	VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header, validators istanbul.ValidatorSet) error
	VerifyUncles(chain consensus.ChainReader, block *types.Block) error
	VerifySeal(chain consensus.ChainHeaderReader, header *types.Header, validators istanbul.ValidatorSet) error

	Prepare(chain consensus.ChainHeaderReader, header *types.Header, validators istanbul.ValidatorSet) error
	Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header)
	FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error)
	ExtractGenesisValidators(header *types.Header) ([]common.Address, error)

	Sign(data []byte) ([]byte, error)
	SignWithoutHashing(data []byte) ([]byte, error)
	CheckSignature(data []byte, address common.Address, sig []byte) error

	Seal(chain consensus.ChainHeaderReader, block *types.Block, validators istanbul.ValidatorSet) (*types.Block, error)
	SealHash(header *types.Header) common.Hash
	CommitHeader(header *types.Header, seals [][]byte, round *big.Int) error

	WriteVote(header *types.Header, candidate common.Address, authorize bool) error
	ReadVote(header *types.Header) (candidate common.Address, authorize bool, err error)
}
