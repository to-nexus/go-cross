// Copyright 2023 The go-ethereum Authors
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

package eth

import (
	"compress/gzip"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/latest"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

// AdminAPI is the collection of Ethereum full node related APIs for node
// administration.
type AdminAPI struct {
	eth *Ethereum
}

// NewAdminAPI creates a new instance of AdminAPI.
func NewAdminAPI(eth *Ethereum) *AdminAPI {
	return &AdminAPI{eth: eth}
}

// ExportChain exports the current blockchain into a local file,
// or a range of blocks if first and last are non-nil.
func (api *AdminAPI) ExportChain(file string, first *uint64, last *uint64) (bool, error) {
	if first == nil && last != nil {
		return false, errors.New("last cannot be specified without first")
	}
	if first != nil && last == nil {
		head := api.eth.BlockChain().CurrentHeader().Number.Uint64()
		last = &head
	}
	if _, err := os.Stat(file); err == nil {
		// File already exists. Allowing overwrite could be a DoS vector,
		// since the 'file' may point to arbitrary paths on the drive.
		return false, errors.New("location would overwrite an existing file")
	}
	// Make sure we can create the file to export into
	out, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return false, err
	}
	defer out.Close()

	var writer io.Writer = out
	if strings.HasSuffix(file, ".gz") {
		writer = gzip.NewWriter(writer)
		defer writer.(*gzip.Writer).Close()
	}

	// Export the blockchain
	if first != nil {
		if err := api.eth.BlockChain().ExportN(writer, *first, *last); err != nil {
			return false, err
		}
	} else if err := api.eth.BlockChain().Export(writer); err != nil {
		return false, err
	}
	return true, nil
}

func hasAllBlocks(chain *core.BlockChain, bs []*types.Block) bool {
	for _, b := range bs {
		if !chain.HasBlock(b.Hash(), b.NumberU64()) {
			return false
		}
	}

	return true
}

// ImportChain imports a blockchain from a local file.
func (api *AdminAPI) ImportChain(file string) (bool, error) {
	// Make sure the can access the file to import
	in, err := os.Open(file)
	if err != nil {
		return false, err
	}
	defer in.Close()

	var reader io.Reader = in
	if strings.HasSuffix(file, ".gz") {
		if reader, err = gzip.NewReader(reader); err != nil {
			return false, err
		}
	}

	// Run actual the import in pre-configured batches
	stream := rlp.NewStream(reader, 0)

	blocks, index := make([]*types.Block, 0, 2500), 0
	for batch := 0; ; batch++ {
		// Load a batch of blocks from the input file
		for len(blocks) < cap(blocks) {
			block := new(types.Block)
			if err := stream.Decode(block); err == io.EOF {
				break
			} else if err != nil {
				return false, fmt.Errorf("block %d: failed to parse: %v", index, err)
			}
			// ignore the genesis block when importing blocks
			if block.NumberU64() == 0 {
				continue
			}
			blocks = append(blocks, block)
			index++
		}
		if len(blocks) == 0 {
			break
		}

		if hasAllBlocks(api.eth.BlockChain(), blocks) {
			blocks = blocks[:0]
			continue
		}
		// Import the batch and reset the buffer
		if _, err := api.eth.BlockChain().InsertChain(blocks); err != nil {
			return false, fmt.Errorf("batch %d: failed to insert: %v", batch, err)
		}
		blocks = blocks[:0]
	}
	return true, nil
}

// ##CROSS: istanbul posa
const (
	suspendTxAttempts = 3                // the number of sends when system txs of this node take the nonce
	suspendTxTimeout  = 30 * time.Second // the time to wait for the suspend tx to be included
)

// SuspendIstanbulValidator sends a transaction that calls StakeHub.suspendByValidator, signed with the node key.
// The node key address must be the current validator address and must hold enough balance for the gas fee.
// The validator is suspended until endTime (unix seconds); an endTime in the past clears the suspension.
// It returns after the tx is included in a block, with an error when the tx reverts.
// When system txs of this node take the nonce, the tx is sent again.
func (api *AdminAPI) SuspendIstanbulValidator(ctx context.Context, endTime uint64) (common.Hash, error) {
	bc := api.eth.BlockChain()
	// Without the StakeHub code, the call would succeed as a plain transfer and only burn gas.
	statedb, err := bc.State()
	if err != nil {
		return common.Hash{}, err
	}
	if statedb.GetCodeSize(contracts.StakeHubAddr) == 0 {
		return common.Hash{}, errors.New("stake hub is not deployed")
	}

	key := api.eth.nodeConfig.NodeKey()
	from, to := crypto.PubkeyToAddress(key.PublicKey), contracts.StakeHubAddr
	input := hexutil.Bytes(latest.NewStakeHub().PackSuspendByValidator(new(big.Int).SetUint64(endTime)))
	args := ethapi.TransactionArgs{From: &from, To: &to, Input: &input}

	// Subscribe before sending, so no new head is missed.
	heads := make(chan core.ChainHeadEvent, 128)
	sub := bc.SubscribeChainHeadEvent(heads)
	defer sub.Unsubscribe()
	ctx, cancel := context.WithTimeout(ctx, suspendTxTimeout)
	defer cancel()

	for attempt := 1; ; attempt++ {
		tx, err := api.sendSuspendTx(ctx, key, args)
		if err != nil {
			return common.Hash{}, err
		}
		included, err := api.waitSuspendTx(ctx, tx, from, heads)
		if included || err != nil {
			return tx.Hash(), err
		}
		if attempt == suspendTxAttempts {
			return common.Hash{}, fmt.Errorf("suspend tx is dropped %d times, as system txs took the nonce", attempt)
		}
		log.Warn("Suspend tx is dropped, as a system tx took the nonce; sending it again", "hash", tx.Hash(), "nonce", tx.Nonce())

		// The pool may still hold the dropped tx, so take the nonce from the chain.
		statedb, err := bc.State()
		if err != nil {
			return common.Hash{}, err
		}
		nonce := hexutil.Uint64(statedb.GetNonce(from))
		args.Nonce = &nonce
	}
}

func (api *AdminAPI) sendSuspendTx(ctx context.Context, key *ecdsa.PrivateKey, args ethapi.TransactionArgs) (*types.Transaction, error) {
	// FillTransaction estimates the gas, so a call that reverts fails here before sending.
	// The result is a dynamic fee tx after London, so it is never taken as a system tx (legacy with zero gas price).
	filled, err := ethapi.NewTransactionAPI(api.eth.APIBackend, new(ethapi.AddrLocker), "").FillTransaction(ctx, args)
	if err != nil {
		return nil, withStakeHubRevertReason(err)
	}
	tx, err := types.SignTx(filled.Tx, types.LatestSignerForChainID(api.eth.BlockChain().Config().ChainID), key)
	if err != nil {
		return nil, err
	}
	if _, err := ethapi.SubmitTransaction(ctx, api.eth.APIBackend, tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// waitSuspendTx waits until tx is included in a block, and returns an error with the reason when tx fails.
// It returns false when another tx took the nonce of tx.
func (api *AdminAPI) waitSuspendTx(ctx context.Context, tx *types.Transaction, from common.Address, heads <-chan core.ChainHeadEvent) (bool, error) {
	bc := api.eth.BlockChain()
	for {
		// Read the state before the tx lookup. A block writes its tx lookups before it becomes the head,
		// so a nonce taken in this state means tx is either found below or dropped.
		statedb, err := bc.State()
		if err != nil {
			return false, err
		}
		if lookup, _ := bc.GetCanonicalTransaction(tx.Hash()); lookup != nil {
			// The state can change after the gas estimation, so the included tx may still revert.
			if bc.GetReceiptsByHash(lookup.BlockHash)[lookup.Index].Status == types.ReceiptStatusFailed {
				return true, fmt.Errorf("suspend tx %s failed: %w", tx.Hash().Hex(), api.suspendTxFailure(ctx, tx, lookup))
			}
			return true, nil
		}
		if statedb.GetNonce(from) > tx.Nonce() {
			return false, nil
		}
		select {
		case <-heads:
		case <-ctx.Done():
			return false, fmt.Errorf("suspend tx %s is not included yet: %w", tx.Hash().Hex(), ctx.Err())
		}
	}
}

// suspendTxFailure runs the included tx again on the state before it, and returns why it failed.
func (api *AdminAPI) suspendTxFailure(ctx context.Context, tx *types.Transaction, lookup *rawdb.LegacyTxLookupEntry) error {
	bc := api.eth.BlockChain()
	block := bc.GetBlock(lookup.BlockHash, lookup.BlockIndex)
	if block == nil {
		return fmt.Errorf("block %s is not found", lookup.BlockHash.Hex())
	}
	_, blockCtx, statedb, release, err := api.eth.stateAtTransaction(ctx, block, int(lookup.Index), 0)
	if err != nil {
		return err
	}
	defer release()
	msg, err := core.TransactionToMessage(tx, types.MakeSigner(bc.Config(), block.Number(), block.Time()), block.BaseFee())
	if err != nil {
		return err
	}
	statedb.SetTxContext(tx.Hash(), int(lookup.Index))
	res, err := core.ApplyMessage(vm.NewEVM(blockCtx, statedb, bc.Config(), vm.Config{}), msg, new(core.GasPool).AddGas(tx.Gas()))
	switch {
	case err != nil:
		return err
	case res.Err == nil:
		return errors.New("tx succeeds when it runs again")
	case len(res.Revert()) > 0:
		return fmt.Errorf("%w: %s", res.Err, stakeHubRevertReason(res.Revert()))
	default:
		return res.Err // e.g. out of gas
	}
}

type (
	rpcRevertError interface {
		rpc.Error
		rpc.DataError
	}
	stakeHubRevertError struct {
		rpcRevertError
		reason string
	}
)

// Error returns the message in the same form as the RPC revert error, with the StakeHub reason.
func (e *stakeHubRevertError) Error() string {
	return vm.ErrExecutionReverted.Error() + ": " + e.reason
}

var stakeHubSuspendErrors = func() map[[4]byte]string {
	errs := make(map[[4]byte]string)
	for _, sig := range []string{
		"EnforcedPause()",
		"InBlackList(address)",
		"ValidatorNotExists()",
		"InvalidValidator()",
		"JailedValidator()",
		"NotEnoughValidators()",
		"SafeCastOverflowedUintDowncast(uint8,uint256)",
	} {
		errs[[4]byte(crypto.Keccak256([]byte(sig)))] = sig
	}
	return errs
}()

// withStakeHubRevertReason adds the StakeHub reason to the RPC revert error, keeping its code and data.
func withStakeHubRevertReason(err error) error {
	var revertErr rpcRevertError
	if !errors.As(err, &revertErr) {
		return err
	}
	data, _ := revertErr.ErrorData().(string)
	revert, decodeErr := hexutil.Decode(data)
	if decodeErr != nil || len(revert) == 0 {
		return err
	}
	return &stakeHubRevertError{revertErr, stakeHubRevertReason(revert)}
}

// stakeHubRevertReason returns the StakeHub error signature, the message of Error(string) or Panic(uint256), or the hex data.
func stakeHubRevertReason(revert []byte) string {
	if len(revert) >= 4 {
		if sig, ok := stakeHubSuspendErrors[[4]byte(revert)]; ok {
			return sig
		}
	}
	if reason, err := abi.UnpackRevert(revert); err == nil {
		return reason
	}
	return hexutil.Encode(revert)
}

// ##
