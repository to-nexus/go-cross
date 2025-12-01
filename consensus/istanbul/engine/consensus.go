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
	"github.com/holiman/uint256"
)

// ##CROSS: istanbul param

// SyncIstanbulParam reads the istanbul parameters from the IstanbulParam contract.
// It reads all checkpoints up to the given timepoint and caches them.
func (e *Engine) SyncIstanbulParam(header *types.Header) error {
	timepoint := header.Number
	istanbulParamInstance := e.istanbulParam.Instance(e.contractBackend, contracts.IstanbulParamAddr)

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

		log.Info("Syncing istanbul param", "index", index, "number", result.Timepoint, "config", result)

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

// getCurrentValidators reads the active validators from the ValidatorSet contract.
func (e *Engine) getCurrentValidators(number uint64) ([]common.Address, error) {
	validatorSetInstance := e.validatorSet.Instance(e.contractBackend, contracts.ValidatorSetAddr)

	callopts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(number)}
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

type ValidatorInfo struct {
	address common.Address
	amount  *uint256.Int
}

// updateValidatorSet creates a new validator set and updates the ValidatorSet contract.
// It reads all staked validators from the StakeHub contract and sorts them by their staking amount,
// then it selects the top N validators as the active validators.
// Finally, it creates a system transaction and applies it to the given state.
func (e *Engine) updateValidatorSet(header *types.Header, state vm.StateDB, cx core.ChainContext, txs *[]*types.Transaction, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, usedGas *uint64, tracer *tracing.Hooks) error {
	if header.Number.Uint64() == 0 {
		return nil
	}
	number := header.Number.Uint64() - 1 // read from parent block

	// read active validators from StakeHub
	validatorInfos, err := e.getStakedValidatorInfo(number)
	if err != nil {
		return err
	}
	if len(validatorInfos) == 0 {
		return errors.New("no validator info")
	}

	threshold, err := e.getValidatorThreshold(number)
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
	msg := newSystemMessage(header.Coinbase, contracts.ValidatorSetAddr, data, nil)

	log.Info("Updating validator set", "number", header.Number.Uint64(), "validators", validators)
	if systemTxs != nil {
		err = e.applyReceivedSystemTransaction(msg, state, header, cx, txs, receipts, systemTxs, usedGas, tracer)
	} else {
		err = e.applyGeneratedSystemTransaction(msg, state, header, cx, txs, receipts, &header.GasUsed, tracer)
	}
	return err
}

// getStakedValidatorInfo reads all staked validators from the StakeHub contract.
func (e *Engine) getStakedValidatorInfo(number uint64) ([]ValidatorInfo, error) {
	stakeHubInstance := e.stakeHub.Instance(e.contractBackend, contracts.StakeHubAddr)

	callopts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(number)}
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
		amountU256, _ := uint256.FromBig(result.Amounts[i])
		validatorInfos[i] = ValidatorInfo{
			address: validator,
			amount:  amountU256,
		}
	}

	return validatorInfos, nil
}

// getValidatorThreshold reads the active validator threshold from the StakeHub contract.
func (e *Engine) getValidatorThreshold(number uint64) (uint64, error) {
	stakeHubInstance := e.stakeHub.Instance(e.contractBackend, contracts.StakeHubAddr)

	callopts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(number)}
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

// ##CROSS: validator slash
// slashValidators slashes validators who failed to propose blocks in the previous sequence
func (e *Engine) slashValidators(header *types.Header, state vm.StateDB, cx *chainContext, txs *[]*types.Transaction, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, usedGas *uint64, tracer *tracing.Hooks) error {
	if e.pos == nil || header.Number.Uint64() == 0 {
		return nil
	}
	number := header.Number.Uint64() - 1 // read from parent block

	idleValidators := e.pos.OfflineValidators(cx.Chain, number)
	for _, validator := range idleValidators {
		log.Info("Slashing idle validator", "number", header.Number.Uint64(), "validator", validator)

		data := e.validatorSlash.PackSlashOffline(validator)
		msg := newSystemMessage(header.Coinbase, contracts.ValidatorSlashAddr, data, nil)

		var err error
		if systemTxs != nil {
			err = e.applyReceivedSystemTransaction(msg, state, header, cx, txs, receipts, systemTxs, usedGas, tracer)
		} else {
			err = e.applyGeneratedSystemTransaction(msg, state, header, cx, txs, receipts, &header.GasUsed, tracer)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

// ##

// ##CROSS: validator reward
func (e *Engine) distributeRewards(header *types.Header, state vm.StateDB, cx core.ChainContext, txs *[]*types.Transaction, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, usedGas *uint64, tracer *tracing.Hooks) error {
	// move collected rewards to the validator
	coinbase := header.Coinbase
	if coinbase == (common.Address{}) {
		coinbase = e.signer
	}

	// Calculate total transaction fees
	// totalFee := state.GetBalance(params.SystemAddress)

	totalFee := new(uint256.Int)
	for i, tx := range *txs {
		if isSystemTx, err := e.IsSystemTransaction(tx, header); err != nil {
			return err
		} else if isSystemTx {
			// All remaining transactions are system transactions
			break
		}

		receipt := (*receipts)[i]

		gasPrice := new(big.Int).Add(header.BaseFee, tx.EffectiveGasTipValue(header.BaseFee)) // base fee is not burned
		fee := new(big.Int).Mul(big.NewInt(int64(receipt.GasUsed)), gasPrice)

		// Blob fee is also included in the total fee
		if receipt.BlobGasPrice != nil && receipt.BlobGasUsed > 0 {
			blobFee := new(big.Int).Mul(new(big.Int).SetUint64(receipt.BlobGasUsed), receipt.BlobGasPrice)
			fee.Add(fee, blobFee)
		}

		feeU256, _ := uint256.FromBig(fee)
		totalFee.Add(totalFee, feeU256)
	}

	if totalFee.IsZero() {
		return nil
	}

	// state.SubBalance(params.SystemAddress, totalFee, tracing.BalanceDecreaseDistributeReward)
	// state.AddBalance(coinbase, totalFee, tracing.BalanceIncreaseDistributeReward)
	state.AddBalance(coinbase, totalFee, tracing.BalanceIncreaseRewardTransactionFee)
	log.Trace("Distributing rewards", "number", header.Number.Uint64(), "validator", coinbase, "totalFee", totalFee, "usedGas", *usedGas)

	// call 'distributeReward' function to send the collected rewards to the validator share contract
	data := e.stakeHub.PackDistributeReward(coinbase)
	msg := newSystemMessage(coinbase, contracts.StakeHubAddr, data, totalFee.ToBig())

	var err error
	if systemTxs != nil {
		err = e.applyReceivedSystemTransaction(msg, state, header, cx, txs, receipts, systemTxs, usedGas, tracer)
	} else {
		err = e.applyGeneratedSystemTransaction(msg, state, header, cx, txs, receipts, &header.GasUsed, tracer)
	}
	if err != nil {
		return err
	}
	return nil
}

// ##
