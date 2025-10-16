package engine

import (
	"bytes"
	"context"
	"errors"
	"math"
	"math/big"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

// ##CROSS: consensus system contract

type ValidatorInfo struct {
	address common.Address
	amount  int64
}

func (e *Engine) getValidators(blockNr rpc.BlockNumberOrHash) ([]common.Address, error) {
	const method = "getValidators"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	code, err := e.ethAPI.GetCode(ctx, contracts.ValidatorSetAddr, blockNr)
	if err != nil {
		log.Error("Failed to get ValidatorSet code", "error", err)
		return nil, err
	}
	if len(code) == 0 {
		// contract is not deployed
		return nil, nil
	}

	data, err := e.validatorSet.Pack(method)
	if err != nil {
		log.Error("Failed to pack tx for getValidators", "error", err)
		return nil, err
	}
	msgData := (hexutil.Bytes)(data)

	toAddress := contracts.ValidatorSetAddr
	gas := (hexutil.Uint64)(uint64(math.MaxUint64 / 2))

	result, err := e.ethAPI.Call(ctx, ethapi.TransactionArgs{
		Gas:  &gas,
		To:   &toAddress,
		Data: &msgData,
	}, &blockNr, nil, nil)
	if err != nil {
		log.Error("Failed to call getValidators", "error", err)
		return nil, err
	}

	var validators []common.Address
	if err := e.validatorSet.UnpackIntoInterface(&validators, method, result); err != nil {
		log.Error("Failed to unpack result of getValidators", "error", err)
		return nil, err
	}

	// sort ascending
	slices.SortFunc(validators, func(a, b common.Address) int {
		return bytes.Compare(a[:], b[:])
	})

	return validators, nil
}

func (e *Engine) updateValidatorSet(header *types.Header, state vm.StateDB, cx core.ChainContext, txs *[]*types.Transaction, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, usedGas *uint64, tracer *tracing.Hooks) error {
	blockNr := rpc.BlockNumberOrHashWithHash(header.ParentHash, false)

	// read active validators from StakeHub
	validatorInfos, err := e.getValidatorElectionInfo(blockNr)
	if err != nil {
		return err
	}
	if len(validatorInfos) == 0 {
		return errors.New("no validator election info")
	}

	threshold, err := e.getValidatorThreshold(blockNr)
	if err != nil {
		return err
	}
	if threshold == 0 {
		return errors.New("validator threshold is 0")
	}

	// cut top N validators by their staking amount
	slices.SortFunc(validatorInfos, func(a, b ValidatorInfo) int {
		return int(a.amount - b.amount)
	})
	if threshold < uint64(len(validatorInfos)) {
		validatorInfos = validatorInfos[:threshold]
	}

	// update validators in ValidatorSet
	validators := make([]common.Address, len(validatorInfos))
	for i, validator := range validatorInfos {
		validators[i] = validator.address
	}

	const method = "updateValidators"
	data, err := e.validatorSet.Pack(method, validators)
	if err != nil {
		log.Error("Failed to pack tx for updateValidators", "error", err)
		return err
	}

	msg := newSystemMessage(header.Coinbase, contracts.ValidatorSetAddr, data)
	log.Info("Updating validator set", "blockNumber", header.Number.Uint64(), "data", common.Bytes2Hex(data))
	if systemTxs != nil {
		err = e.applyReceivedSystemTransaction(msg, state, header, cx, txs, receipts, systemTxs, usedGas, tracer)
	} else {
		msg.From = e.signer // header.Coinbase is empty here
		err = e.applyGeneratedSystemTransaction(msg, state, header, cx, txs, receipts, &header.GasUsed, tracer)
	}
	return err
}

func (e *Engine) getValidatorElectionInfo(blockNr rpc.BlockNumberOrHash) ([]ValidatorInfo, error) {
	const method = "getValidatorElectionInfo"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, err := e.stakeHub.Pack(method, big.NewInt(0), big.NewInt(0))
	if err != nil {
		log.Error("Failed to pack tx for getValidatorElectionInfo", "error", err)
		return nil, err
	}
	msgData := (hexutil.Bytes)(data)

	toAddress := contracts.StakeHubAddr
	gas := (hexutil.Uint64)(uint64(math.MaxUint64 / 2))

	result, err := e.ethAPI.Call(ctx, ethapi.TransactionArgs{
		Gas:  &gas,
		To:   &toAddress,
		Data: &msgData,
	}, &blockNr, nil, nil)
	if err != nil {
		log.Error("Failed to call getValidatorElectionInfo", "error", err)
		return nil, err
	}

	var validators []common.Address
	var amounts []*big.Int
	var totalLength *big.Int
	if err := e.stakeHub.UnpackIntoInterface(&[]any{&validators, &amounts, &totalLength}, method, result); err != nil {
		log.Error("Failed to unpack result of getValidatorElectionInfo", "error", err)
		return nil, err
	}
	if length := totalLength.Uint64(); length != uint64(len(validators)) || length != uint64(len(amounts)) {
		return nil, errors.New("validator length mismatch")
	}

	validatorInfos := make([]ValidatorInfo, len(validators))
	for i, validator := range validators {
		validatorInfos[i] = ValidatorInfo{
			address: validator,
			amount:  amounts[i].Int64(),
		}
	}

	return validatorInfos, nil
}

func (e *Engine) getValidatorThreshold(blockNr rpc.BlockNumberOrHash) (uint64, error) {
	const method = "validatorThreshold"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, err := e.stakeHub.Pack(method)
	if err != nil {
		log.Error("Failed to pack tx for validatorThreshold", "error", err)
		return 0, err
	}
	msgData := (hexutil.Bytes)(data)

	toAddress := contracts.StakeHubAddr
	gas := (hexutil.Uint64)(uint64(math.MaxUint64 / 2))

	result, err := e.ethAPI.Call(ctx, ethapi.TransactionArgs{
		Gas:  &gas,
		To:   &toAddress,
		Data: &msgData,
	}, &blockNr, nil, nil)
	if err != nil {
		log.Error("Failed to call validatorThreshold", "error", err)
		return 0, err
	}

	var threshold *big.Int
	if err := e.stakeHub.UnpackIntoInterface(&threshold, method, result); err != nil {
		log.Error("Failed to unpack result of validatorThreshold", "error", err)
		return 0, err
	}

	return threshold.Uint64(), nil
}

// ##
