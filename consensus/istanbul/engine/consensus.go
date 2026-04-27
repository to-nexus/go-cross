package engine

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
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
	timepoint := header.Number.Uint64()
	istanbulParamInstance := e.istanbulParam.Instance(e.contractBackend, contracts.IstanbulParamAddr)

	// get the latest param index at this timepoint
	latestIndex, err := bind.Call(istanbulParamInstance, nil, e.istanbulParam.PackGetParamIndex(timepoint), e.istanbulParam.UnpackGetParamIndex)
	if err != nil {
		if errors.Is(err, bind.ErrNoCode) {
			// contract is not deployed
			return nil
		}
		return fmt.Errorf("failed to call getParamIndex: %w", err)
	}

	// get the number of cached configs
	length := params.IstanbulConfigLength()

	// read uncached configs from the contract
	for index := uint64(length); index <= latestIndex; index++ {
		result, err := bind.Call(istanbulParamInstance, nil, e.istanbulParam.PackGetParamsByIndex(index), e.istanbulParam.UnpackGetParamsByIndex)
		if err != nil {
			return fmt.Errorf("failed to call getParamsByIndex(%d): %w", index, err)
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
		if result.CouncilPeriod != 0 {
			config.CouncilPeriod = &result.CouncilPeriod
		}

		params.AddIstanbulConfig(&config, result.Timepoint)
	}
	return nil
}

// ##

// ##CROSS: consensus system contract

type validatorSet breakpoint.GetActiveValidatorsOutput

var (
	errValidatorSetLengthMismatch = errors.New("validator set length mismatch")
	errInvalidValidatorSigner     = errors.New("invalid validator signer")
)

func (v validatorSet) Len() int { return len(v.ValidatorAddrs) }
func (v validatorSet) Less(i, j int) bool {
	return bytes.Compare(v.ValidatorAddrs[i][:], v.ValidatorAddrs[j][:]) < 0
}
func (v validatorSet) Swap(i, j int) {
	v.ValidatorAddrs[i], v.ValidatorAddrs[j] = v.ValidatorAddrs[j], v.ValidatorAddrs[i]
	v.SignerAddrs[i], v.SignerAddrs[j] = v.SignerAddrs[j], v.SignerAddrs[i]
}

// getCurrentValidators reads the active validators from the ValidatorSet contract.
func (e *Engine) getCurrentValidators(number uint64) ([]common.Address, []types.BLSPublicKey, error) {
	validatorSetInstance := e.validatorSet.Instance(e.contractBackend, contracts.ValidatorSetAddr)

	callopts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(number)}
	calldata := e.validatorSet.PackGetActiveValidators()
	result, err := bind.Call(validatorSetInstance, callopts, calldata, e.validatorSet.UnpackGetActiveValidators)
	if err != nil {
		if errors.Is(err, bind.ErrNoCode) {
			// contract is not deployed
			return nil, nil, nil
		}
		return nil, nil, fmt.Errorf("failed to call getActiveValidators: %w", err)
	}

	if len(result.ValidatorAddrs) != len(result.SignerAddrs) {
		return nil, nil, fmt.Errorf("%w: validators %d signers %d", errValidatorSetLengthMismatch, len(result.ValidatorAddrs), len(result.SignerAddrs))
	}
	for i, signer := range result.SignerAddrs {
		if len(signer) != types.BLSPublicKeyLength {
			return nil, nil, fmt.Errorf("%w: signer %d has length %d", errInvalidValidatorSigner, i, len(signer))
		}
	}

	// sort ascending
	sort.Sort(validatorSet(result))

	signers := make([]types.BLSPublicKey, 0, len(result.SignerAddrs))
	for _, signer := range result.SignerAddrs {
		signers = append(signers, types.BytesToBLSPublicKey(signer))
	}
	return result.ValidatorAddrs, signers, nil
}

type ValidatorInfo struct {
	address common.Address
	signer  types.BLSPublicKey
	power   *uint256.Int
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
		// do not update validator set, it will use the legacy validators from the header
		log.Warn("No validator info", "number", header.Number.Uint64())
		return nil
	}

	threshold, err := e.getMaxCouncil(number)
	if err != nil {
		return err
	}
	if threshold == 0 {
		return errors.New("zero validator threshold")
	}

	// sort descending by their staking power
	slices.SortFunc(validatorInfos, func(a, b ValidatorInfo) int {
		return b.power.Cmp(a.power)
	})

	// cut top N validators by their staking power
	if len(validatorInfos) > int(threshold) {
		validatorInfos = validatorInfos[:threshold]
	}

	// sort ascending by their address
	slices.SortFunc(validatorInfos, func(a, b ValidatorInfo) int {
		return bytes.Compare(a.address[:], b.address[:])
	})

	validators := make([]common.Address, 0, threshold)
	// ##CROSS: bls seal
	signersBytes := make([][]byte, 0, threshold)
	signers := make([]types.BLSPublicKey, 0, threshold)
	// ##
	for _, validator := range validatorInfos {
		if validator.power.IsZero() || len(validators) >= int(threshold) {
			// no more active validators or enough validators are collected
			break
		}
		validators = append(validators, validator.address)
		// ##CROSS: bls seal
		signersBytes = append(signersBytes, validator.signer.Bytes())
		signers = append(signers, validator.signer)
		// ##
	}

	data := e.validatorSet.PackUpdateValidators(validators, signersBytes)
	msg := newSystemMessage(header.Coinbase, contracts.ValidatorSetAddr, data, nil)

	log.Info("Updating validator set", "number", header.Number.Uint64(), "validators", validators, "signers", signers, "isMining", systemTxs == nil)
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
		return nil, fmt.Errorf("failed to call getValidators: %w", err)
	}

	if length := result.TotalLength.Uint64(); length != uint64(len(result.ValidatorAddrs)) || length != uint64(len(result.SignerAddrs)) || length != uint64(len(result.Powers)) {
		return nil, errors.New("validator length mismatch")
	}

	validatorInfos := make([]ValidatorInfo, len(result.ValidatorAddrs))
	for i, validator := range result.ValidatorAddrs {
		powerU256, _ := uint256.FromBig(result.Powers[i])
		signer := result.SignerAddrs[i]
		validatorInfos[i] = ValidatorInfo{
			address: validator,
			signer:  types.BytesToBLSPublicKey(signer),
			power:   powerU256,
		}
	}

	return validatorInfos, nil
}

// getMaxCouncil reads the max number of council from the StakeHub contract.
func (e *Engine) getMaxCouncil(number uint64) (uint64, error) {
	stakeHubInstance := e.stakeHub.Instance(e.contractBackend, contracts.StakeHubAddr)

	callopts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(number)}
	calldata := e.stakeHub.PackMaxCouncil()
	cnt, err := bind.Call(stakeHubInstance, callopts, calldata, e.stakeHub.UnpackMaxCouncil)
	if err != nil {
		if errors.Is(err, bind.ErrNoCode) {
			// contract is not deployed
			return 0, nil
		}
		return 0, fmt.Errorf("failed to call maxCouncil: %w", err)
	}

	return cnt.Uint64(), nil
}

// ##

// ##CROSS: validator slash
// slashValidators slashes validators who failed to propose blocks in the previous sequence
func (e *Engine) slashValidators(header *types.Header, state vm.StateDB, cx *chainContext, txs *[]*types.Transaction, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, usedGas *uint64, tracer *tracing.Hooks) error {
	if header == nil || header.Number.Uint64() == 0 {
		return nil
	}

	// Slash offline proposers of the parent block
	parent := cx.Chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil || parent.Number.Uint64() == 0 {
		return nil
	}

	slashed, err := e.offlineProposers(cx.Chain, parent)
	if err != nil {
		return err
	}
	if len(slashed) == 0 {
		return nil
	}

	slashedCount := map[common.Address]int{}
	for round, proposer := range slashed {
		slashedCount[proposer]++
		log.Info("Slashing offline proposer",
			"number", header.Number.Uint64(),
			"block", parent.Number.Uint64(),
			"round", round,
			"proposer", proposer,
			"isMining", systemTxs == nil,
		)
	}

	// Slashing same proposer multiple times will be reverted; reduce them
	target := make([]common.Address, 0, len(slashedCount))
	for proposer, count := range slashedCount {
		target = append(target, proposer)
		if count > 1 {
			log.Warn("Multiple slashed proposer",
				"number", header.Number.Uint64(),
				"block", parent.Number.Uint64(),
				"proposer", proposer,
				"count", count,
			)
		}
	}

	data := e.validatorSlash.PackSlashOffline(target)
	msg := newSystemMessage(header.Coinbase, contracts.ValidatorSlashAddr, data, nil)

	if systemTxs != nil {
		_, err = e.applyOptionalReceivedSystemTransaction("validator slash", msg, state, header, cx, txs, receipts, systemTxs, usedGas, tracer)
		return err
	}
	e.applyOptionalGeneratedSystemTransaction("validator slash", msg, state, header, cx, txs, receipts, &header.GasUsed, tracer)
	return nil
}

// offlineProposers collects the offline proposers of the given header.
func (e *Engine) offlineProposers(chain consensus.ChainHeaderReader, header *types.Header) ([]common.Address, error) {
	if header == nil || header.Number.Uint64() == 0 {
		return nil, nil
	}
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil || parent.Number.Uint64() == 0 {
		return nil, nil
	}

	// Only round robin policy is supported
	policy := e.cfg.GetConfig(header.Number).ProposerPolicy
	if policy == nil || policy.Id != istanbul.RoundRobin {
		return nil, nil
	}

	extra, err := types.ExtractIstanbulExtra(header)
	if err != nil {
		return nil, err
	}
	if extra.Round == 0 {
		// No extra round, so no proposers to slash
		return nil, nil
	}

	// If extra.Round > 0, collect proposers of previous rounds
	validators := extra.Validators
	if len(validators) == 0 {
		// Fallback to IstanbulEngine to get validators
		if ie := consensus.ToIstanbulEngine(e.consensus); ie != nil {
			validators = ie.ValidatorsAt(chain, header)
		}
	}
	if len(validators) == 0 {
		return nil, fmt.Errorf("no validators at block %d", header.Number.Uint64())
	}

	valSet := validator.NewSet(validators, nil, policy)
	lastProposer, _ := e.Author(parent)
	proposers := make([]common.Address, 0, extra.Round)
	for round := uint32(0); round < extra.Round; round++ {
		valSet.CalcProposer(lastProposer, uint64(round))
		proposer := valSet.GetProposer()
		proposers = append(proposers, proposer.Address())
		lastProposer = proposer.Address()
	}

	return proposers, nil
}

// mitigateSlashedValidators mitigates the slashed validators by configured rate
func (e *Engine) mitigateSlashedValidators(header *types.Header, state vm.StateDB, cx *chainContext, txs *[]*types.Transaction, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, usedGas *uint64, tracer *tracing.Hooks) error {
	if header == nil || header.Number.Uint64() == 0 {
		return nil
	}

	validatorSlashInstance := e.validatorSlash.Instance(e.contractBackend, contracts.ValidatorSlashAddr)

	number := header.Number.Uint64() - 1 // read from parent block
	callopts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(number)}
	calldata := e.validatorSlash.PackGetSlashedValidators()
	slashed, err := bind.Call(validatorSlashInstance, callopts, calldata, e.validatorSlash.UnpackGetSlashedValidators)
	if err != nil {
		if errors.Is(err, bind.ErrNoCode) {
			// contract is not deployed
			return nil
		}
		return fmt.Errorf("failed to call getSlashedValidators: %w", err)
	}

	if len(slashed) == 0 {
		// No slashed validators
		return nil
	}

	data := e.validatorSlash.PackMitigate()
	msg := newSystemMessage(header.Coinbase, contracts.ValidatorSlashAddr, data, nil)

	log.Info("Mitigating slashed validators", "number", header.Number.Uint64(), "slashed", slashed, "isMining", systemTxs == nil)
	if systemTxs != nil {
		_, err = e.applyOptionalReceivedSystemTransaction("mitigate slashed validators", msg, state, header, cx, txs, receipts, systemTxs, usedGas, tracer)
		return err
	}
	e.applyOptionalGeneratedSystemTransaction("mitigate slashed validators", msg, state, header, cx, txs, receipts, &header.GasUsed, tracer)
	return nil
}

// ##

// ##CROSS: validator reward
func (e *Engine) distributeRewards(header *types.Header, state vm.StateDB, cx core.ChainContext, txs *[]*types.Transaction, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, usedGas *uint64, tracer *tracing.Hooks) error {
	coinbase := header.Coinbase
	if coinbase == (common.Address{}) {
		coinbase = e.signer
	}

	// Calculate total tip and total fee by summing up the tip and fee of all transactions
	totalTip := new(big.Int)
	totalFee := new(big.Int)
	for i, tx := range *txs {
		if isSystemTx, err := e.IsSystemTransaction(tx, header); err != nil {
			return err
		} else if isSystemTx {
			// All remaining transactions are system transactions
			break
		}
		if i >= len(*receipts) {
			break
		}

		receipt := (*receipts)[i]

		gasTipCap, err := tx.EffectiveGasTip(header.BaseFee)
		if err != nil {
			return err
		}
		gasPrice := new(big.Int).Add(header.BaseFee, gasTipCap)
		total := new(big.Int).Mul(big.NewInt(int64(receipt.GasUsed)), gasPrice)
		tip := new(big.Int).Mul(big.NewInt(int64(receipt.GasUsed)), gasTipCap)

		// Blob fee is also included in the total fee
		if receipt.BlobGasPrice != nil && receipt.BlobGasUsed > 0 {
			blobFee := new(big.Int).Mul(new(big.Int).SetUint64(receipt.BlobGasUsed), receipt.BlobGasPrice)
			total.Add(total, blobFee)
		}

		totalFee.Add(totalFee, total)
		totalTip.Add(totalTip, tip)
	}

	log.Trace("Distributing rewards",
		"number", header.Number.Uint64(),
		"validator", coinbase,
		"totalFee", totalFee,
		"totalTip", totalTip,
		"usedGas", *usedGas,
		"isMining", systemTxs == nil)

	snapshot := newSystemStateSnapshot(state, txs, receipts, usedGas)

	totalFeeU256, _ := uint256.FromBig(totalFee)
	if totalFeeU256.Sign() != 0 {
		state.AddBalance(coinbase, totalFeeU256, tracing.BalanceIncreaseRewardTransactionFee)
	}

	// Call 'distributeReward' function to send the collected rewards to the validator share contract
	data := e.rewardHub.PackDistributeReward(coinbase, totalTip)
	msg := newSystemMessage(coinbase, contracts.RewardHubAddr, data, totalFee)

	if systemTxs != nil {
		included, err := e.applyOptionalReceivedSystemTransaction("distribute rewards", msg, state, header, cx, txs, receipts, systemTxs, usedGas, tracer)
		if err != nil || !included {
			snapshot.Revert()
		}
		if err == nil && !included {
			// `distributeRewards` is omitted by the proposer, fallback total fee to RewardHub
			fallbackReward(header, state, coinbase, totalFeeU256)
		}
		return err
	}

	if !e.applyOptionalGeneratedSystemTransaction("distribute rewards", msg, state, header, cx, txs, receipts, usedGas, tracer) {
		// Failed to distribute rewards, fallback total fee to RewardHub
		snapshot.Revert()
		fallbackReward(header, state, coinbase, totalFeeU256)
	}
	return nil
}

// fallbackReward sends the reward to the RewardHub contract, in case the distributeRewards transaction fails.
func fallbackReward(header *types.Header, state vm.StateDB, coinbase common.Address, amount *uint256.Int) {
	if amount == nil || amount.IsZero() {
		return
	}
	log.Warn("Falling back reward balance to RewardHub",
		"number", header.Number.Uint64(),
		"validator", coinbase,
		"rewardHub", contracts.RewardHubAddr,
		"amount", amount,
	)
	state.AddBalance(contracts.RewardHubAddr, amount, tracing.BalanceIncreaseRewardTransactionFee)
}

// ##
