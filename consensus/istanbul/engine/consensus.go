package engine

import (
	"bytes"
	"errors"
	"math/big"
	"slices"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/holiman/uint256"
)

// ##CROSS: istanbul param

// SyncIstanbulParam reads the istanbul parameters from the IstanbulParam contract.
// It reads all checkpoints up to the given timepoint and caches them.
func (e *Engine) SyncIstanbulParam(header *types.Header) error {
	timepoint := header.Number
	istanbulParamInstance := e.istanbulParam.Instance(e.contactBackend, contracts.IstanbulParamAddr)

	length, err := bind.Call(istanbulParamInstance, nil, e.istanbulParam.PackNumCheckpoints(), e.istanbulParam.UnpackNumCheckpoints)
	if err != nil {
		log.Error("Failed to call numCheckpoints", "error", err)
		return err
	}
	if length == 0 {
		return nil
	}

	// get the latest param index at this timepoint
	latestIndex, err := bind.Call(istanbulParamInstance, nil, e.istanbulParam.PackGetParamIndex(timepoint), e.istanbulParam.UnpackGetParamIndex)
	if err != nil {
		if errors.Is(err, bind.ErrNoCode) {
			// contract is not deployed
			return nil
		}
		log.Error("Failed to call getParamIndex", "error", err)
		return err
	}

	// if no params exist yet, nothing to sync
	if latestIndex == length {
		return nil
	}

	// iterate through all param indices to latestIndex
	for index := uint64(0); index <= latestIndex; index++ {
		if params.IstanbulConfigByIndex(index) != nil {
			// already cached
			continue
		}

		// read the config from the contract and cache it
		result, err := bind.Call(istanbulParamInstance, nil, e.istanbulParam.PackGetParamsByIndex(index), e.istanbulParam.UnpackGetParamsByIndex)
		if err != nil {
			log.Error("Failed to call getParamsByIndex", "index", index, "error", err)
			return err
		}

		log.Info("Syncing istanbul param", "index", index, "blockNumber", result.Timepoint, "config", result)

		config := params.IstanbulConfig{
			EpochLength:             result.EpochLength,
			BlockPeriodSeconds:      result.BlockPeriod,
			EmptyBlockPeriodSeconds: result.EmptyBlockPeriod,
			RequestTimeoutSeconds:   result.RequestTimeout,
			ProposerPolicy:          result.ProposerPolicy,
		}
		if result.MaxRequestTimeout != 0 {
			config.MaxRequestTimeoutSeconds = &result.MaxRequestTimeout
		}
		if result.GasLimit != 0 {
			config.GasLimit = &result.GasLimit
		}
		if result.ElasticityMultiplier != 0 {
			config.ElasticityMultiplier = &result.ElasticityMultiplier
		}
		if result.BaseFeeChangeDenominator != 0 {
			config.BaseFeeChangeDenominator = &result.BaseFeeChangeDenominator
		}
		if result.MaxBaseFee != nil && result.MaxBaseFee.Sign() > 0 {
			config.MaxBaseFee = (*math.HexOrDecimal256)(result.MaxBaseFee)
		}
		if result.MinBaseFee != nil && result.MinBaseFee.Sign() > 0 {
			config.MinBaseFee = (*math.HexOrDecimal256)(result.MinBaseFee)
		}

		params.SetIstanbulConfig(index, &config, result.Timepoint)
	}
	return nil
}

// ##

// ##CROSS: consensus system contract

type ValidatorInfo struct {
	address common.Address
	amount  *uint256.Int
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

// getValidators reads the active validators from the ValidatorSet contract.
func (e *Engine) getValidators(blockNr rpc.BlockNumberOrHash) ([]common.Address, error) {
	validatorSetInstance := e.validatorSet.Instance(e.contactBackend, contracts.ValidatorSetAddr)

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

// updateValidatorSet creates a new validator set and updates the ValidatorSet contract.
// It reads all staked validators from the StakeHub contract and sorts them by their staking amount,
// then it selects the top N validators as the active validators.
// Finally, it creates a system transaction and applies it to the given state.
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
		return b.amount.Cmp(a.amount)
	})

	validators := make([]common.Address, 0, threshold)
	for _, validator := range validatorInfos {
		if validator.amount.IsZero() || len(validators) >= int(threshold) {
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

// getValidatorInfo reads all staked validators from the StakeHub contract.
func (e *Engine) getValidatorInfo(blockNr rpc.BlockNumberOrHash) ([]ValidatorInfo, error) {
	stakeHubInstance := e.stakeHub.Instance(e.contactBackend, contracts.StakeHubAddr)

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
			amount:  uint256.MustFromBig(result.Amounts[i]),
		}
	}

	return validatorInfos, nil
}

// getValidatorThreshold reads the active validator threshold from the StakeHub contract.
func (e *Engine) getValidatorThreshold(blockNr rpc.BlockNumberOrHash) (uint64, error) {
	stakeHubInstance := e.stakeHub.Instance(e.contactBackend, contracts.StakeHubAddr)

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
