package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/contracts/predeploy"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ##CROSS: make genesis

func TestMakeGenesisCommand(t *testing.T) {
	t.Run("generates keys without input", func(t *testing.T) {
		dir := t.TempDir()
		t.Chdir(dir)
		geth := runGeth(t, "make-genesis", "--all", "--validators", "2", "--json")
		output := geth.Output()
		geth.WaitExit()
		require.Zero(t, geth.ExitStatus(), geth.StderrText())

		var keys struct {
			Admin struct {
				Address common.Address `json:"address"`
			} `json:"admin"`
			Beneficiary struct {
				Address common.Address `json:"address"`
			} `json:"beneficiary"`
			Validators []json.RawMessage `json:"validators"`
		}
		require.NoError(t, json.Unmarshal(output, &keys))
		assert.Len(t, keys.Validators, 2)

		data, err := os.ReadFile(filepath.Join(dir, "genesis.json"))
		require.NoError(t, err)

		genesis := new(core.Genesis)
		require.NoError(t, json.Unmarshal(data, genesis))
		assert.Equal(t, keys.Admin.Address, genesis.Config.Istanbul.PoSA.Admin)
		assert.Equal(t, keys.Beneficiary.Address, *genesis.Config.Istanbul.Beneficiary)
		assert.Zero(t, *genesis.Config.BreakpointTime)
		assert.Zero(t, *genesis.Config.BreakpointBravoTime)
		for _, addr := range []common.Address{contracts.CrossExAddr, contracts.BridgeAddr, contracts.Multicall3Addr, contracts.SingletonFactoryAddr, contracts.SingletonFactory2Addr} {
			assert.NotEmpty(t, genesis.Alloc[addr].Code, "code at %s", addr)
		}
		// Before initialize, the bridge proxy holds only its implementation slot.
		assert.Greater(t, len(genesis.Alloc[contracts.BridgeAddr].Storage), 1, "the bridge must be initialized")

		db := rawdb.NewMemoryDatabase()
		_, err = genesis.Commit(db, triedb.NewDatabase(db, nil))
		require.NoError(t, err)
	})

	t.Run("no breakpoint", func(t *testing.T) {
		dir := t.TempDir()
		t.Chdir(dir)
		geth := runGeth(t, "make-genesis", "--bridge", "--validators", "2")
		geth.WaitExit()
		require.Zero(t, geth.ExitStatus(), geth.StderrText())

		data, err := os.ReadFile(filepath.Join(dir, "genesis.json"))
		require.NoError(t, err)

		genesis := new(core.Genesis)
		require.NoError(t, json.Unmarshal(data, genesis))
		zero := uint64(0)
		assert.Equal(t, &zero, genesis.Config.CancunTime)
		assert.Equal(t, &zero, genesis.Config.PragueTime)
		assert.Nil(t, genesis.Config.BreakpointTime)
		assert.Nil(t, genesis.Config.BreakpointBravoTime)
		assert.Nil(t, genesis.Config.OsakaTime)
		assert.Equal(t, common.Hex2Bytes(predeploy.CrossExCode), genesis.Alloc[contracts.CrossExAddr].Code)
		assert.Equal(t, params.HistoryStorageCode, genesis.Alloc[params.HistoryStorageAddress].Code)
		assert.NotContains(t, genesis.Alloc, contracts.ValidatorSetAddr)
		assert.Greater(t, len(genesis.Alloc[contracts.BridgeAddr].Storage), 1, "the bridge must be initialized")

		db := rawdb.NewMemoryDatabase()
		_, err = genesis.Commit(db, triedb.NewDatabase(db, nil))
		require.NoError(t, err)
	})

	t.Run("rejects --validators with input", func(t *testing.T) {
		geth := runGeth(t, "make-genesis", "--validators", "2", "in.json", "out.json")
		geth.WaitExit()
		assert.NotZero(t, geth.ExitStatus())
		assert.Contains(t, geth.StderrText(), "only available without input arguments")
	})
}
