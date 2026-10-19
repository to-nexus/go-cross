package core_test

import (
	"context"
	"encoding/json"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/contracts/latest"
	"github.com/ethereum/go-ethereum/contracts/predeploy"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ##CROSS: make genesis

func TestMakeGenesis(t *testing.T) {
	funded := common.HexToAddress("0xf00d")
	attacker := common.HexToAddress("0xbad")
	bridgeABI, err := abi.JSON(strings.NewReader(`[
		{"type":"function","name":"initialize","inputs":[{"type":"address"},{"type":"address"},{"type":"uint8"}]},
		{"type":"function","name":"upgradeToAndCall","inputs":[{"type":"address"},{"type":"bytes"}]},
		{"type":"function","name":"hasRole","inputs":[{"type":"bytes32"},{"type":"address"}],"outputs":[{"type":"bool"}]}
	]`))
	require.NoError(t, err)

	// legacyGenesis has Breakpoint scheduled after genesis.
	legacyGenesis := func() *core.Genesis {
		config := *params.CrossDev3ChainConfig
		return &core.Genesis{
			Config:     &config,
			GasLimit:   30_000_000,
			Difficulty: big.NewInt(1),
			Alloc:      types.GenesisAlloc{funded: {Balance: big.NewInt(1)}},
		}
	}
	// posaGenesis can become a Breakpoint genesis.
	posaGenesis := func() *core.Genesis {
		genesis := legacyGenesis()
		istanbul := *genesis.Config.Istanbul
		posa := *istanbul.PoSA
		genesis.Config.Istanbul = &istanbul
		istanbul.PoSA = &posa
		posa.Validators = append([]params.PoSAValidator(nil), posa.Validators[:2]...)
		genesis.Config.BreakpointTime = nil
		return genesis
	}
	// roundTrip checks that the genesis survives the JSON file that geth init reads.
	roundTrip := func(t *testing.T, genesis *core.Genesis) {
		t.Helper()
		data, err := json.Marshal(genesis)
		require.NoError(t, err)
		require.NoError(t, json.Unmarshal(data, new(core.Genesis)))
	}
	call := func(t *testing.T, backend *simulated.Backend, from, to common.Address, method string, args ...any) ([]byte, error) {
		t.Helper()
		data, err := bridgeABI.Pack(method, args...)
		require.NoError(t, err)
		return backend.Client().CallContract(context.Background(), ethereum.CallMsg{From: from, To: &to, Data: data}, nil)
	}

	t.Run("adds only CrossEx by default", func(t *testing.T) {
		genesis := legacyGenesis()
		input := genesis.Alloc
		require.NoError(t, core.MakeGenesis(genesis, core.GenesisContracts{}))

		assert.Len(t, genesis.Alloc, 2)
		assert.Equal(t, common.Hex2Bytes(predeploy.CrossExCode), genesis.Alloc[contracts.CrossExAddr].Code)
		assert.Equal(t, big.NewInt(1), genesis.Alloc[funded].Balance)
		assert.NotContains(t, input, contracts.CrossExAddr, "the input allocation must stay unchanged")
		roundTrip(t, genesis)
	})

	t.Run("initializes the Breakpoint contracts", func(t *testing.T) {
		genesis := posaGenesis()
		posa := genesis.Config.Istanbul.PoSA
		posa.DelegationPool, posa.RewardStartBlock = common.Address{}, nil
		coinbase := common.HexToAddress("0x1234")
		reserveBalance, systemBalance, crossExBalance := big.NewInt(12345), big.NewInt(7), big.NewInt(3)
		genesis.Coinbase = coinbase
		genesis.Alloc[contracts.CrossReserveAddr] = types.Account{Balance: reserveBalance}
		genesis.Alloc[params.SystemAddress] = types.Account{Balance: systemBalance}
		genesis.Alloc[contracts.CrossExAddr] = types.Account{Balance: crossExBalance, Nonce: 1, Code: common.Hex2Bytes(predeploy.CrossExCode)}
		require.NoError(t, core.MakeGenesis(genesis, core.GenesisContracts{Breakpoint: true}))

		assert.Equal(t, coinbase, genesis.Coinbase)
		assert.Zero(t, *genesis.Config.BreakpointTime)
		assert.Equal(t, contracts.DelegationPoolAddr, posa.DelegationPool)
		assert.Equal(t, big.NewInt(1), posa.RewardStartBlock)
		assert.Equal(t, reserveBalance, genesis.Alloc[contracts.CrossReserveAddr].Balance)
		assert.Equal(t, systemBalance, genesis.Alloc[params.SystemAddress].Balance)
		assert.Equal(t, common.Hex2Bytes(breakpoint.CrossExCode), genesis.Alloc[contracts.CrossExAddr].Code)
		assert.Equal(t, crossExBalance, genesis.Alloc[contracts.CrossExAddr].Balance)
		assert.Equal(t, common.FromHex(latest.ValidatorSetMetaData.BinRuntime), genesis.Alloc[contracts.ValidatorSetImplAddr].Code)
		assert.Equal(t, common.FromHex(latest.StakeHubMetaData.BinRuntime), genesis.Alloc[contracts.StakeHubImplAddr].Code)
		assert.Equal(t, common.FromHex(latest.RewardHubMetaData.BinRuntime), genesis.Alloc[contracts.RewardHubImplAddr].Code)
		assert.Equal(t, common.FromHex(latest.ValidatorSlashMetaData.BinRuntime), genesis.Alloc[contracts.ValidatorSlashImplAddr].Code)
		roundTrip(t, genesis)

		backend := simulated.NewBackend(genesis.Alloc)
		defer backend.Close()
		validatorSet := latest.NewValidatorSet()
		active, err := bind.Call(
			validatorSet.Instance(backend.Client(), contracts.ValidatorSetAddr),
			&bind.CallOpts{},
			validatorSet.PackGetActiveValidators(),
			validatorSet.UnpackGetActiveValidators,
		)
		require.NoError(t, err)
		assert.Len(t, active.ValidatorAddrs, 2)

		stakeHub := latest.NewStakeHub()
		initialStake := new(big.Int).Mul(big.NewInt(1_000_000), big.NewInt(params.Ether))
		for _, validator := range posa.Validators {
			staked, err := bind.Call(
				stakeHub.Instance(backend.Client(), contracts.StakeHubAddr),
				&bind.CallOpts{},
				stakeHub.PackGetStakedAmount(validator.Operator),
				stakeHub.UnpackGetStakedAmount,
			)
			require.NoError(t, err)
			assert.Equal(t, initialStake, staked)
		}
		assert.Equal(t, new(big.Int).Mul(initialStake, big.NewInt(2)), genesis.Alloc[contracts.DelegationPoolAddr].Balance)
	})

	t.Run("adds HistoryStorage if Prague is active at genesis", func(t *testing.T) {
		genesis := legacyGenesis()
		zero := uint64(0)
		genesis.Config.CancunTime, genesis.Config.PragueTime = &zero, &zero
		require.NoError(t, core.MakeGenesis(genesis, core.GenesisContracts{}))
		assert.Equal(t, params.HistoryStorageCode, genesis.Alloc[params.HistoryStorageAddress].Code)
	})

	t.Run("adds Multicall3 and the singleton factories", func(t *testing.T) {
		genesis := legacyGenesis()
		require.NoError(t, core.MakeGenesis(genesis, core.GenesisContracts{Multicall: true, SingletonFactory: true}))

		assert.Equal(t, common.Hex2Bytes(predeploy.Multicall3Code), genesis.Alloc[contracts.Multicall3Addr].Code)
		for _, addr := range []common.Address{contracts.SingletonFactoryAddr, contracts.SingletonFactory2Addr} {
			assert.Equal(t, common.Hex2Bytes(predeploy.SingletonFactoryCode), genesis.Alloc[addr].Code)
			assert.Equal(t, uint64(1), genesis.Alloc[addr].Nonce)
		}
		roundTrip(t, genesis)
	})

	t.Run("initializes the bridge", func(t *testing.T) {
		genesis := posaGenesis()
		balance := big.NewInt(12345)
		genesis.Alloc[contracts.BridgeAddr] = types.Account{Balance: balance}
		opts := core.GenesisContracts{Breakpoint: true, Bridge: true, Multicall: true, SingletonFactory: true}
		require.NoError(t, core.MakeGenesis(genesis, opts))
		assert.Equal(t, balance, genesis.Alloc[contracts.BridgeAddr].Balance)

		// The genesis must commit like a node would on geth init.
		db := rawdb.NewMemoryDatabase()
		_, err := genesis.Commit(db, triedb.NewDatabase(db, nil))
		require.NoError(t, err)

		backend := simulated.NewBackend(genesis.Alloc)
		defer backend.Close()
		admin := genesis.Config.Istanbul.PoSA.Admin
		ret, err := call(t, backend, attacker, contracts.BridgeAddr, "hasRole", common.Hash{}, admin)
		require.NoError(t, err)
		assert.Equal(t, common.LeftPadBytes([]byte{1}, 32), ret, "the PoSA admin must own the bridge")

		for _, to := range []common.Address{contracts.BridgeAddr, contracts.BridgeImplAddr} {
			_, err = call(t, backend, attacker, to, "initialize", attacker, attacker, uint8(1))
			assert.Error(t, err, "initialize on %s", to)
		}
		_, err = call(t, backend, attacker, contracts.BridgeAddr, "upgradeToAndCall", contracts.Multicall3Addr, []byte{})
		assert.Error(t, err)
	})

	t.Run("the shipped allocation lets anyone initialize the bridge", func(t *testing.T) {
		// Control for the subtest above: the same call succeeds without genesis initialization.
		shipped := core.DefaultCrossDev3GenesisBlock().Alloc
		backend := simulated.NewBackend(types.GenesisAlloc{
			contracts.BridgeAddr:     shipped[contracts.BridgeAddr],
			contracts.BridgeImplAddr: shipped[contracts.BridgeImplAddr],
		})
		defer backend.Close()
		_, err := call(t, backend, attacker, contracts.BridgeAddr, "initialize", attacker, attacker, uint8(1))
		assert.NoError(t, err)
	})

	t.Run("uses the given bridge owner", func(t *testing.T) {
		genesis := legacyGenesis()
		owner := common.HexToAddress("0x0123")
		require.NoError(t, core.MakeGenesis(genesis, core.GenesisContracts{Bridge: true, BridgeOwner: owner}))

		backend := simulated.NewBackend(genesis.Alloc)
		defer backend.Close()
		ret, err := call(t, backend, attacker, contracts.BridgeAddr, "hasRole", common.Hash{}, owner)
		require.NoError(t, err)
		assert.Equal(t, common.LeftPadBytes([]byte{1}, 32), ret)
	})

	t.Run("requires a bridge owner without a PoSA admin", func(t *testing.T) {
		genesis := legacyGenesis()
		istanbul := *genesis.Config.Istanbul
		istanbul.PoSA = nil
		genesis.Config.Istanbul = &istanbul
		assert.ErrorContains(t, core.MakeGenesis(genesis, core.GenesisContracts{Bridge: true}), "bridge owner is required")
	})

	t.Run("rejects code already in the allocation", func(t *testing.T) {
		genesis := legacyGenesis()
		genesis.Alloc[contracts.Multicall3Addr] = types.Account{Code: []byte{0x00}}
		assert.ErrorContains(t, core.MakeGenesis(genesis, core.GenesisContracts{Multicall: true}), "conflicts with contract")
	})
}
