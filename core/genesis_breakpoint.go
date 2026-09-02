// Copyright 2026 The go-ethereum Authors
// This file is part of the go-ethereum library.

package core

import (
	"errors"
	"fmt"
	"maps"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/params/predeploy"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/holiman/uint256"
)

// ##CROSS: fork breakpoint

const initialOperatorStake = 1_000_000

var (
	reserveGenesisABI = mustGenesisABI(`[
		{"type":"function","name":"initialize","inputs":[{"name":"owner","type":"address"},{"name":"account","type":"address"}]},
		{"type":"function","name":"addManager","inputs":[{"name":"manager","type":"address"}]}
	]`)
	poolGenesisABI = mustGenesisABI(`[
		{"type":"function","name":"initialize","inputs":[{"name":"owner","type":"address"},{"name":"router","type":"address"},{"name":"crossReserve","type":"address"},{"name":"protocolReserve","type":"address"},{"name":"startBlock","type":"uint256"}]},
		{"type":"function","name":"setPosaActivation","inputs":[{"name":"blockNumber","type":"uint256"}]}
	]`)
	stakeHubGenesisABI = mustGenesisABI(`[
		{"type":"function","name":"stakeFor","inputs":[{"name":"account","type":"address"}]}
	]`)
)

// MakeBreakpointGenesis adds initialized Breakpoint contracts to a custom genesis.
func MakeBreakpointGenesis(genesis *Genesis) error {
	return makeBreakpointGenesis(genesis, nil)
}

func clonePoSAChainConfig(config *params.ChainConfig) *params.ChainConfig {
	cloned := *config
	if config.Istanbul != nil {
		istanbul := *config.Istanbul
		cloned.Istanbul = &istanbul
		if config.Istanbul.PoSA != nil {
			posa := *config.Istanbul.PoSA
			posa.Validators = append([]params.PoSAValidator(nil), config.Istanbul.PoSA.Validators...)
			cloned.Istanbul.PoSA = &posa
		}
	}
	return &cloned
}

func cloneAlloc(alloc types.GenesisAlloc) types.GenesisAlloc {
	cloned := make(types.GenesisAlloc, len(alloc))
	for addr, account := range alloc {
		if account.Balance != nil {
			account.Balance = new(big.Int).Set(account.Balance)
		}
		account.Code = common.CopyBytes(account.Code)
		account.Storage = maps.Clone(account.Storage)
		cloned[addr] = account
	}
	return cloned
}

// makeBreakpointGenesis creates a Breakpoint genesis allocation with initialized contracts and state.
func makeBreakpointGenesis(genesis *Genesis, registeredValidators []params.PoSAValidator) error {
	if err := prepareBreakpointConfig(genesis); err != nil {
		return err
	}
	if err := addBreakpointGenesisCode(genesis); err != nil {
		return err
	}
	statedb, err := makeBreakpointState(genesis.Alloc)
	if err != nil {
		return err
	}
	if err := initializeBreakpointState(genesis, statedb, registeredValidators); err != nil {
		return err
	}
	alloc, err := dumpBreakpointState(statedb, genesis.Config)
	if err != nil {
		return err
	}
	genesis.Alloc = alloc
	return genesis.validateBreakpointAlloc()
}

// prepareBreakpointConfig checks the genesis config for Breakpoint requirements and sets default fork times and PoSA parameters.
func prepareBreakpointConfig(genesis *Genesis) error {
	if genesis == nil || genesis.Config == nil {
		return errors.New("genesis chain config is required")
	}

	// Check config.
	config := genesis.Config
	if config.ChainID == nil {
		return errors.New("genesis chain ID is required")
	}
	if !config.IsLondon(common.Big0) {
		return errors.New("requires London active at genesis")
	}
	if config.Istanbul == nil || config.Istanbul.PoSA == nil {
		return errors.New("requires Istanbul PoSA config")
	}

	// Check PoSA config.
	posa := config.Istanbul.PoSA
	if posa.Admin == (common.Address{}) {
		return errors.New("requires PoSA admin")
	}
	if posa.CouncilPeriod == 0 || posa.ValidatorEpochLength == 0 {
		return errors.New("requires PoSA council period and validator epoch length to be positive")
	}
	if len(posa.Validators) == 0 {
		return errors.New("requires at least one PoSA validator")
	}
	for i, validator := range posa.Validators {
		if validator.Operator == (common.Address{}) || validator.Validator == (common.Address{}) || validator.ID == "" {
			return fmt.Errorf("requires PoSA validator %d to be complete", i)
		}
		if len(validator.Signer) != types.BLSPublicKeyLength {
			return fmt.Errorf("validator %d signer has length %d, want %d", i, len(validator.Signer), types.BLSPublicKeyLength)
		}
	}

	// Set default fork times and PoSA parameters.
	zero := uint64(0)
	config.ShanghaiTime = &zero
	config.AdventureTime = &zero
	config.CancunTime = &zero
	config.PragueTime = &zero
	config.BreakpointTime = &zero
	posa.DelegationPool = contracts.DelegationPoolAddr
	posa.RewardStartBlock = big.NewInt(1)
	return config.CheckConfigForkOrder()
}

// addBreakpointGenesisCode adds the Breakpoint contracts to the genesis allocation.
func addBreakpointGenesisCode(genesis *Genesis) error {
	if genesis.Alloc == nil {
		genesis.Alloc = make(types.GenesisAlloc)
	}

	// Replacements.
	delete(genesis.Alloc, contracts.CrossExAddr)

	// List up the Breakpoint contracts to add.
	tmpl := types.GenesisAlloc{
		contracts.CrossExAddr:            {Balance: new(big.Int), Nonce: 1, Code: common.Hex2Bytes(breakpoint.CrossExCode)},
		params.BeaconRootsAddress:        {Balance: new(big.Int), Nonce: 1, Code: params.BeaconRootsCode},
		params.HistoryStorageAddress:     {Balance: new(big.Int), Nonce: 1, Code: params.HistoryStorageCode},
		params.WithdrawalQueueAddress:    {Balance: new(big.Int), Nonce: 1, Code: params.WithdrawalQueueCode},
		params.ConsolidationQueueAddress: {Balance: new(big.Int), Nonce: 1, Code: params.ConsolidationQueueCode},
	}
	maps.Insert(tmpl, maps.All(predeploy.BreakpointGenesisAlloc()))

	// Add the Breakpoint contracts to the genesis allocation.
	for addr, account := range tmpl {
		if existing, ok := genesis.Alloc[addr]; ok {
			// If the existing account has non-zero nonce, code, or storage, consider it as a conflict.
			if existing.Nonce != 0 || len(existing.Code) != 0 || len(existing.Storage) != 0 {
				return fmt.Errorf("genesis allocation conflicts with Breakpoint contract at %s", addr)
			}
			// Preserve the existing balance if it is non-zero.
			if existing.Balance != nil {
				account.Balance = new(big.Int).Set(existing.Balance)
			}
		}
		genesis.Alloc[addr] = account
	}
	return nil
}

// makeBreakpointState creates a new state database initialized with the given genesis allocation.
func makeBreakpointState(alloc types.GenesisAlloc) (*state.StateDB, error) {
	db := rawdb.NewMemoryDatabase()
	tdb := triedb.NewDatabase(db, &triedb.Config{Preimages: true})
	statedb, err := state.New(types.EmptyRootHash, state.NewDatabase(tdb, nil))
	if err != nil {
		return nil, err
	}

	for addr, account := range alloc {
		balance := account.Balance
		if balance == nil {
			balance = new(big.Int)
		}
		statedb.SetBalance(addr, uint256.MustFromBig(balance), tracing.BalanceIncreaseGenesisBalance)
		statedb.SetNonce(addr, account.Nonce, tracing.NonceChangeGenesis)
		statedb.SetCode(addr, account.Code, tracing.CodeChangeGenesis)
		for key, value := range account.Storage {
			statedb.SetState(addr, key, value)
		}
	}
	return statedb, nil
}

// initializeBreakpointState runs the initialization calls for every Breakpoint contract.
func initializeBreakpointState(genesis *Genesis, statedb *state.StateDB, registeredValidators []params.PoSAValidator) error {
	posa := genesis.Config.Istanbul.PoSA
	admin := posa.Admin
	one := big.NewInt(1)
	if registeredValidators == nil {
		registeredValidators = posa.Validators
	}

	// Run the PoSA contract initialization calls.
	calls := []genesisCall{
		{admin, contracts.CrossReserveAddr, reserveGenesisABI, "initialize", nil, []any{admin, contracts.DelegationPoolAddr}},
		{admin, contracts.ProtocolReserveAddr, reserveGenesisABI, "initialize", nil, []any{admin, contracts.DelegationPoolAddr}},
		{admin, contracts.DelegationPoolAddr, poolGenesisABI, "initialize", nil, []any{admin, contracts.StakeHubAddr, contracts.CrossReserveAddr, contracts.ProtocolReserveAddr, one}},
		{admin, contracts.CrossReserveAddr, reserveGenesisABI, "addManager", nil, []any{admin}},
		{admin, contracts.ProtocolReserveAddr, reserveGenesisABI, "addManager", nil, []any{admin}},
		{admin, contracts.DelegationPoolAddr, poolGenesisABI, "setPosaActivation", nil, []any{one}},
	}
	for _, call := range calls {
		if err := runGenesisCall(genesis, statedb, call); err != nil {
			return err
		}
	}

	// Collect parameters from the registered validators.
	var (
		operators  = make([]common.Address, 0, len(registeredValidators))
		validators = make([]common.Address, 0, len(registeredValidators))
		signers    = make([][]byte, 0, len(registeredValidators))
		ids        = make([]string, 0, len(registeredValidators))
	)
	for _, validator := range registeredValidators {
		operators = append(operators, validator.Operator)
		validators = append(validators, validator.Validator)
		signers = append(signers, common.CopyBytes(validator.Signer))
		ids = append(ids, validator.ID)
	}

	// Run the Breakpoint system contract initialization calls.
	system := params.SystemAddress
	systemCalls := []struct {
		to   common.Address
		data []byte
	}{
		{contracts.ValidatorSetAddr, packValidatorSetInitialize(admin, posa.Validators)},
		{contracts.StakeHubAddr, breakpoint.NewStakeHub().PackInitialize(contracts.DelegationPoolAddr, admin, operators, validators, signers, ids)},
		{contracts.RewardHubAddr, breakpoint.NewRewardHub().PackInitialize(contracts.DelegationPoolAddr, admin, one)},
		{contracts.ValidatorSlashAddr, breakpoint.NewValidatorSlash().PackInitialize(admin)},
	}
	for _, call := range systemCalls {
		if err := runPackedGenesisCall(genesis, statedb, system, call.to, nil, call.data); err != nil {
			return err
		}
	}

	// Stake the initial operator stake for each registered validator.
	stake := new(big.Int).Mul(big.NewInt(initialOperatorStake), big.NewInt(params.Ether))
	totalStake := new(big.Int).Mul(stake, new(big.Int).SetUint64(uint64(len(operators))))
	statedb.AddBalance(system, uint256.MustFromBig(totalStake), tracing.BalanceIncreaseGenesisBalance)
	for _, operator := range operators {
		call := genesisCall{system, contracts.StakeHubAddr, stakeHubGenesisABI, "stakeFor", stake, []any{operator}}
		if err := runGenesisCall(genesis, statedb, call); err != nil {
			return err
		}
	}
	return nil
}

func packValidatorSetInitialize(admin common.Address, validatorConfig []params.PoSAValidator) []byte {
	validators := make([]common.Address, 0, len(validatorConfig))
	signers := make([][]byte, 0, len(validatorConfig))
	for _, validator := range validatorConfig {
		validators = append(validators, validator.Validator)
		signers = append(signers, common.CopyBytes(validator.Signer))
	}
	return breakpoint.NewValidatorSet().PackInitialize(admin, validators, signers)
}

type genesisCall struct {
	caller common.Address
	to     common.Address
	abi    abi.ABI
	method string
	value  *big.Int
	args   []any
}

func runGenesisCall(genesis *Genesis, statedb *state.StateDB, call genesisCall) error {
	data, err := call.abi.Pack(call.method, call.args...)
	if err != nil {
		return fmt.Errorf("pack %s: %w", call.method, err)
	}
	return runPackedGenesisCall(genesis, statedb, call.caller, call.to, call.value, data)
}

func runPackedGenesisCall(genesis *Genesis, statedb *state.StateDB, caller, to common.Address, value *big.Int, data []byte) error {
	if value == nil {
		value = new(big.Int)
	}
	baseFee := genesis.BaseFee
	if baseFee == nil {
		baseFee = new(big.Int)
	}
	random := common.Hash{}
	blockContext := vm.BlockContext{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		GetHash:     func(uint64) common.Hash { return common.Hash{} },
		Coinbase:    params.SystemAddress,
		BlockNumber: new(big.Int),
		Time:        genesis.Timestamp,
		Difficulty:  new(big.Int),
		BaseFee:     baseFee,
		BlobBaseFee: big.NewInt(params.BlobTxMinBlobGasprice),
		GasLimit:    math.MaxUint64,
		Random:      &random,
	}
	evm := vm.NewEVM(blockContext, statedb, genesis.Config, vm.Config{})
	evm.SetTxContext(vm.TxContext{Origin: caller, GasPrice: new(big.Int)})
	rules := genesis.Config.Rules(common.Big0, true, genesis.Timestamp)
	statedb.Prepare(rules, caller, params.SystemAddress, &to, vm.ActivePrecompiles(rules), nil)
	ret, _, err := evm.Call(caller, to, data, math.MaxUint64, uint256.MustFromBig(value))
	if err != nil {
		if reason, unpackErr := abi.UnpackRevert(ret); unpackErr == nil {
			return fmt.Errorf("genesis call to %s reverted: %s", to, reason)
		}
		return fmt.Errorf("genesis call to %s failed: %w", to, err)
	}
	return nil
}

// dumpBreakpointState dumps the state of the given state database into a genesis allocation.
func dumpBreakpointState(statedb *state.StateDB, config *params.ChainConfig) (types.GenesisAlloc, error) {
	root, err := statedb.Commit(0, config.IsEIP158(common.Big0), false)
	if err != nil {
		return nil, err
	}
	statedb, err = state.New(root, statedb.Database())
	if err != nil {
		return nil, err
	}
	alloc := make(genesisAllocDump)
	statedb.DumpToCollector(alloc, &state.DumpConfig{OnlyWithAddresses: true})
	return types.GenesisAlloc(alloc), nil
}

type genesisAllocDump types.GenesisAlloc

func (genesisAllocDump) OnRoot(common.Hash) {}

func (alloc genesisAllocDump) OnAccount(addr *common.Address, account state.DumpAccount) {
	if addr == nil {
		return
	}
	balance, _ := new(big.Int).SetString(account.Balance, 0)
	storage := make(map[common.Hash]common.Hash, len(account.Storage))
	for key, value := range account.Storage {
		storage[key] = common.HexToHash(value)
	}
	alloc[*addr] = types.Account{Balance: balance, Nonce: account.Nonce, Code: account.Code, Storage: storage}
}

func mustGenesisABI(input string) abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(input))
	if err != nil {
		panic(err)
	}
	return parsed
}
