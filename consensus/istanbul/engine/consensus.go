package engine

import (
	"bytes"
	"cmp"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
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
	"github.com/holiman/uint256"
)

// ##CROSS: consensus system contract

type validatorSet breakpoint.GetActiveValidatorsOutput

var (
	errValidatorSetLengthMismatch = errors.New("validator set length mismatch")
	errValidatorSetMismatch       = errors.New("validator set mismatch")
	errValidatorSignerMismatch    = errors.New("validator signer mismatch")
	errInvalidValidatorSigner     = errors.New("invalid validator signer")
)

// verifyValidators checks the validators and signers in the header against the current active validators from the ValidatorSet contract.
func (e *Engine) verifyValidators(chain consensus.ChainHeaderReader, header *types.Header) error {
	if header == nil || header.Number.Uint64() == 0 {
		return nil
	}
	// Validator set is only updated on the beginning of a new epoch
	if !e.cfg.GetConfig(header.Number).OnNewValidatorEpoch(header.Number.Uint64()) {
		return nil
	}

	extra, err := types.ExtractIstanbulExtra(header)
	if err != nil {
		return err
	}

	var (
		validatorList []common.Address
		signerList    []types.BLSPublicKey
	)

	// When a validator epoch block is also a council period rollover, the ValidatorSet contract is updated later in this block's Finalize.
	// Reading the contract at parent here would return the pre-rollover council.
	// We pre-compute the post-rollover council manually to match what updateValidatorSet will produce in Finalize.
	if chain != nil {
		parent := chain.GetHeaderByHash(header.ParentHash)
		if parent != nil && e.cfg.GetConfig(header.Number).OnNewCouncilPeriod(parent.Time, header.Time) {
			validatorList, signerList, err = e.computeNextCouncil(header.Number.Uint64() - 1)
			log.Warn("New epoch + new council period: computing next council manually",
				"number", header.Number.Uint64(),
				"validatorList", validatorList,
				"signerList", signerList)
			if err != nil {
				return err
			}
		}
	}

	// Default path: read the current contract state at parent.
	if len(validatorList) == 0 {
		validatorList, signerList, err = e.getCurrentValidators(header.Number.Uint64() - 1)
		if err != nil {
			return err
		}
	}

	if len(validatorList) == 0 {
		// Fallback: when the contract has not been initialized yet, the header keeps carrying the current snapshot validators
		log.Warn("verifyValidators: fallback to snapshot validators", "number", header.Number.Uint64())
		if ie := consensus.ToIstanbulEngine(e.consensus); ie != nil && chain != nil {
			parent := chain.GetHeaderByHash(header.ParentHash)
			if parent == nil {
				return consensus.ErrUnknownAncestor
			}
			validatorList = ie.ValidatorsAt(chain, parent)
		}
		if !slices.Equal(extra.Validators, validatorList) {
			return fmt.Errorf("%w: header %v, snapshot %v", errValidatorSetMismatch, extra.Validators, validatorList)
		}
		return nil
	}

	if !slices.Equal(extra.Validators, validatorList) {
		return fmt.Errorf("%w: header %v, contract %v", errValidatorSetMismatch, extra.Validators, validatorList)
	}
	if !slices.Equal(extra.Signers, signerList) {
		return fmt.Errorf("%w: header %v, contract %v", errValidatorSignerMismatch, extra.Signers, signerList)
	}
	return nil
}

func (v validatorSet) Len() int { return len(v.ValidatorAddrs) }
func (v validatorSet) Less(i, j int) bool {
	return bytes.Compare(v.ValidatorAddrs[i][:], v.ValidatorAddrs[j][:]) < 0
}
func (v validatorSet) Swap(i, j int) {
	v.ValidatorAddrs[i], v.ValidatorAddrs[j] = v.ValidatorAddrs[j], v.ValidatorAddrs[i]
	v.SignerAddrs[i], v.SignerAddrs[j] = v.SignerAddrs[j], v.SignerAddrs[i]
}

// getCurrentValidators reads the current active validators from the ValidatorSet contract.
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
	address    common.Address
	signer     types.BLSPublicKey
	power      *uint256.Int
	createTime uint64
}

func (v ValidatorInfo) String() string {
	return fmt.Sprintf("{address: %s, power: %s, createTime: %d}", v.address.Hex(), v.power.String(), v.createTime)
}

// computeNextCouncil computes the (validators, signers) that updateValidatorSet would write into the ValidatorSet contract.
func (e *Engine) computeNextCouncil(number uint64) ([]common.Address, []types.BLSPublicKey, error) {
	validatorInfos, err := e.getStakedValidatorInfo(number)
	if err != nil {
		return nil, nil, err
	}
	if len(validatorInfos) == 0 {
		return nil, nil, errors.New("empty staked validators")
	}

	threshold, err := e.getMaxCouncil(number)
	if err != nil {
		return nil, nil, err
	}
	if threshold == 0 {
		return nil, nil, errors.New("zero validator threshold")
	}

	selected := selectActiveCouncil(validatorInfos, threshold)
	if len(selected) == 0 {
		return nil, nil, errors.New("empty filtered validators")
	}

	validators := make([]common.Address, 0, len(selected))
	signers := make([]types.BLSPublicKey, 0, len(selected))
	for _, v := range selected {
		validators = append(validators, v.address)
		signers = append(signers, v.signer)
	}
	return validators, signers, nil
}

// selectActiveCouncil derives the active council from a list of staked validators by:
//   - sort by power descending, breaking ties by createTime ascending then address ascending
//   - cut top N and drop zero-power entries
func selectActiveCouncil(validatorInfos []ValidatorInfo, threshold uint64) []ValidatorInfo {
	log.Debug("Selecting active council", "threshold", threshold, "candidates", validatorInfos)

	// sort descending by power -> ascending by createTime -> ascending by address
	slices.SortStableFunc(validatorInfos, func(a, b ValidatorInfo) int {
		if c := b.power.Cmp(a.power); c != 0 {
			return c
		}
		if c := cmp.Compare(a.createTime, b.createTime); c != 0 {
			return c
		}
		return bytes.Compare(a.address[:], b.address[:])
	})

	// cut top N and drop zero-power entries
	for n := range validatorInfos {
		if n >= int(threshold) || validatorInfos[n].power.IsZero() {
			validatorInfos = validatorInfos[:n]
			break
		}
	}

	// sort ascending by address for stable on-chain layout
	slices.SortFunc(validatorInfos, func(a, b ValidatorInfo) int {
		return bytes.Compare(a.address[:], b.address[:])
	})
	return validatorInfos
}

// updateValidatorSet creates a new validator set and updates the ValidatorSet contract.
// It reads all staked validators from the StakeHub contract at parent block and organizes active council from them.
// Finally, it creates a system transaction and applies it to the given state.
func (e *Engine) updateValidatorSet(header *types.Header, state vm.StateDB, cx core.ChainContext, txs *[]*types.Transaction, receipts *[]*types.Receipt, systemTxs *[]*types.Transaction, usedGas *uint64, tracer *tracing.Hooks) error {
	if header.Number.Uint64() == 0 {
		return nil
	}
	number := header.Number.Uint64() - 1 // read from parent block

	validators, signers, err := e.computeNextCouncil(number)
	if err != nil {
		return err
	}

	signersBytes := make([][]byte, 0, len(signers))
	for _, s := range signers {
		signersBytes = append(signersBytes, s.Bytes())
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
			address:    validator,
			signer:     types.BytesToBLSPublicKey(signer),
			power:      powerU256,
			createTime: result.CreateTimes[i].Uint64(),
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
	slices.SortFunc(target, func(a, b common.Address) int {
		return bytes.Compare(a[:], b[:])
	})

	data := e.validatorSlash.PackSlashOffline(target)
	msg := newSystemMessage(header.Coinbase, contracts.ValidatorSlashAddr, data, nil)

	if systemTxs != nil {
		err = e.applyReceivedSystemTransaction(msg, state, header, cx, txs, receipts, systemTxs, usedGas, tracer)
	} else {
		err = e.applyGeneratedSystemTransaction(msg, state, header, cx, txs, receipts, &header.GasUsed, tracer)
	}
	return err
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
	// ##CROSS: validator sort policy
	// Validators are sorted by byte
	policy = istanbul.NewProposerPolicyByIdAndSortFunc(policy.Id, istanbul.ValidatorSortByByte())
	// ##

	extra, err := types.ExtractIstanbulExtra(header)
	if err != nil {
		return nil, err
	}
	if extra.Round == 0 {
		// No extra round, so no proposers to slash
		return nil, nil
	}

	// If extra.Round > 0, collect proposers of previous rounds; select proposers from the parent snapshot validator set
	var validators []common.Address
	if ie := consensus.ToIstanbulEngine(e.consensus); ie != nil {
		validators = ie.ValidatorsAt(chain, parent)
	}
	if len(validators) == 0 {
		log.Warn("No validators at block", "number", header.Number.Uint64())
		return nil, nil
	}

	valSet := validator.NewSet(validators, nil, policy)
	lastProposer, _ := e.Author(parent)
	proposers := make([]common.Address, 0, extra.Round)
	for round := uint32(0); round < extra.Round; round++ {
		valSet.CalcProposer(lastProposer, uint64(round))
		proposer := valSet.GetProposer()
		proposers = append(proposers, proposer.Address())
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
		log.Warn("Failed to call getSlashedValidators", "error", err, "number", header.Number.Uint64())
		return nil
	}

	if len(slashed) == 0 {
		// No slashed validators
		return nil
	}

	data := e.validatorSlash.PackMitigate()
	msg := newSystemMessage(header.Coinbase, contracts.ValidatorSlashAddr, data, nil)

	log.Info("Mitigating slashed validators", "number", header.Number.Uint64(), "slashed", slashed, "isMining", systemTxs == nil)
	if systemTxs != nil {
		err = e.applyReceivedSystemTransaction(msg, state, header, cx, txs, receipts, systemTxs, usedGas, tracer)
	} else {
		err = e.applyGeneratedSystemTransaction(msg, state, header, cx, txs, receipts, &header.GasUsed, tracer)
	}
	return err
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

		gasTip, err := tx.EffectiveGasTip(header.BaseFee)
		if err != nil {
			return err
		}
		gasPrice := new(big.Int).Add(header.BaseFee, gasTip)
		total := new(big.Int).Mul(big.NewInt(int64(receipt.GasUsed)), gasPrice)
		tip := new(big.Int).Mul(big.NewInt(int64(receipt.GasUsed)), gasTip)

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

	// Call 'distributeReward' function to send the collected rewards to the validator share contract
	// Total fee is already collected to the coinbase's balance in the state transition.
	data := e.rewardHub.PackDistributeReward(coinbase, totalTip)
	msg := newSystemMessage(coinbase, contracts.RewardHubAddr, data, totalFee)

	var err error
	if systemTxs != nil {
		err = e.applyReceivedSystemTransaction(msg, state, header, cx, txs, receipts, systemTxs, usedGas, tracer)
	} else {
		err = e.applyGeneratedSystemTransaction(msg, state, header, cx, txs, receipts, usedGas, tracer)
	}
	return err
}

// ##
