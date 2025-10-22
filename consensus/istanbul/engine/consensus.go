package engine

import (
	"bytes"
	"errors"
	"math/big"
	"slices"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

// ##CROSS: consensus system contract

type ValidatorInfo struct {
	address common.Address
	amount  int64
}

func withBlockNumberOrHash(blockNr rpc.BlockNumberOrHash) *bind.CallOpts {
	callOpts := bind.CallOpts{}
	if hash, ok := blockNr.Hash(); ok {
		callOpts.BlockHash = hash
	} else if number, ok := blockNr.Number(); ok {
		callOpts.BlockNumber = big.NewInt(int64(number))
	}
	return &callOpts
}

func (e *Engine) getValidators(blockNr rpc.BlockNumberOrHash) ([]common.Address, error) {
	validatorSetInstance := e.validatorSet.Instance(e.ethClient, contracts.ValidatorSetAddr)

	callopts := withBlockNumberOrHash(blockNr)
	calldata := e.validatorSet.PackGetValidators()
	validators, err := bind.Call(validatorSetInstance, callopts, calldata, e.validatorSet.UnpackGetValidators)
	if err != nil {
		if errors.Is(err, bind.ErrNoCode) {
			// contract is not deployed
			return nil, nil
		}
		log.Error("Failed to call getValidators", "error", err)
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
	validatorInfos, err := e.getValidatorInfo(blockNr)
	if err != nil {
		return err
	}
	if len(validatorInfos) == 0 {
		return errors.New("no validator info")
	}

	threshold, err := e.getValidatorThreshold(blockNr)
	if err != nil {
		return err
	}
	if threshold == 0 {
		return errors.New("zero validator threshold")
	}

	// cut top N validators by their staking amount
	slices.SortFunc(validatorInfos, func(a, b ValidatorInfo) int {
		return int(a.amount - b.amount)
	})

	validators := make([]common.Address, 0, threshold)
	for _, validator := range validatorInfos {
		if validator.amount <= 0 || len(validators) >= int(threshold) {
			// no more active validators or enough validators are collected
			break
		}
		validators = append(validators, validator.address)
	}

	data := e.validatorSet.PackUpdateValidators(validators)
	msg := newSystemMessage(header.Coinbase, contracts.ValidatorSetAddr, data)

	log.Info("Updating validator set", "blockNumber", header.Number.Uint64(), "validators", validators)
	if systemTxs != nil {
		err = e.applyReceivedSystemTransaction(msg, state, header, cx, txs, receipts, systemTxs, usedGas, tracer)
	} else {
		msg.From = e.signer // header.Coinbase is empty here
		err = e.applyGeneratedSystemTransaction(msg, state, header, cx, txs, receipts, &header.GasUsed, tracer)
	}
	return err
}

func (e *Engine) getValidatorInfo(blockNr rpc.BlockNumberOrHash) ([]ValidatorInfo, error) {
	stakeHubInstance := e.stakeHub.Instance(e.ethClient, contracts.StakeHubAddr)

	callopts := withBlockNumberOrHash(blockNr)
	calldata := e.stakeHub.PackGetValidators(big.NewInt(0), big.NewInt(0))
	result, err := bind.Call(stakeHubInstance, callopts, calldata, e.stakeHub.UnpackGetValidators)
	if err != nil {
		if errors.Is(err, bind.ErrNoCode) {
			// contract is not deployed
			return nil, nil
		}
		log.Error("Failed to call getValidators", "error", err)
		return nil, err
	}

	if length := result.TotalLength.Uint64(); length != uint64(len(result.ValidatorAddrs)) || length != uint64(len(result.Amounts)) {
		return nil, errors.New("validator length mismatch")
	}

	validatorInfos := make([]ValidatorInfo, len(result.ValidatorAddrs))
	for i, validator := range result.ValidatorAddrs {
		validatorInfos[i] = ValidatorInfo{
			address: validator,
			amount:  result.Amounts[i].Int64(),
		}
	}

	return validatorInfos, nil
}

func (e *Engine) getValidatorThreshold(blockNr rpc.BlockNumberOrHash) (uint64, error) {
	stakeHubInstance := e.stakeHub.Instance(e.ethClient, contracts.StakeHubAddr)

	callopts := withBlockNumberOrHash(blockNr)
	calldata := e.stakeHub.PackValidatorThreshold()
	threshold, err := bind.Call(stakeHubInstance, callopts, calldata, e.stakeHub.UnpackValidatorThreshold)
	if err != nil {
		if errors.Is(err, bind.ErrNoCode) {
			// contract is not deployed
			return 0, nil
		}
		log.Error("Failed to call validatorThreshold", "error", err)
		return 0, err
	}

	return threshold.Uint64(), nil
}

// ##
