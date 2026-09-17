// Copyright 2016 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

var customGenesisTests = []struct {
	genesis string
	query   string
	result  string
}{
	// Genesis file with a mostly-empty chain configuration (ensure missing fields work)
	{
		genesis: `{
			"alloc"      : {},
			"coinbase"   : "0x0000000000000000000000000000000000000000",
			"difficulty" : "0x20000",
			"extraData"  : "",
			"gasLimit"   : "0x2fefd8",
			"nonce"      : "0x0000000000001338",
			"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
			"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
			"timestamp"  : "0x00",
			"config": {
				"terminalTotalDifficulty": 0
			}
		}`,
		query:  "eth.getBlock(0).nonce",
		result: "0x0000000000001338",
	},
	// Genesis file with specific chain configurations
	{
		genesis: `{
			"alloc"      : {},
			"coinbase"   : "0x0000000000000000000000000000000000000000",
			"difficulty" : "0x20000",
			"extraData"  : "",
			"gasLimit"   : "0x2fefd8",
			"nonce"      : "0x0000000000001339",
			"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
			"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
			"timestamp"  : "0x00",
			"config"     : {
				"homesteadBlock"                : 42,
				"daoForkBlock"                  : 141,
				"daoForkSupport"                : true,
				"terminalTotalDifficulty": 0
			}
		}`,
		query:  "eth.getBlock(0).nonce",
		result: "0x0000000000001339",
	},
}

func TestMakeDefaultBreakpointGenesis(t *testing.T) {
	t.Run("reject zero validators", func(t *testing.T) {
		_, _, err := makeDefaultBreakpointGenesis(0)
		require.ErrorContains(t, err, "greater than zero")
	})

	t.Run("generate requested validators", func(t *testing.T) {
		testMakeDefaultBreakpointGenesis(t, 5)
	})

	t.Run("print keys as JSON", func(t *testing.T) {
		_, keys, err := makeDefaultBreakpointGenesis(2)
		require.NoError(t, err)
		var output bytes.Buffer
		require.NoError(t, printBreakpointGenesisKeys(&output, keys, true))
		var decoded struct {
			Admin      map[string]string `json:"admin"`
			Validators []struct {
				ID        string            `json:"id"`
				Validator map[string]string `json:"validator"`
				Operator  map[string]string `json:"operator"`
				Signer    map[string]string `json:"signer"`
			} `json:"validators"`
		}
		require.NoError(t, json.Unmarshal(output.Bytes(), &decoded))
		require.Equal(t, crypto.PubkeyToAddress(keys.admin.PublicKey).Hex(), decoded.Admin["address"])
		require.Len(t, decoded.Validators, 2)
		require.Equal(t, "validator1", decoded.Validators[0].ID)
		require.NotEmpty(t, decoded.Validators[0].Validator["privateKey"])
		require.NotEmpty(t, decoded.Validators[0].Operator["privateKey"])
		require.NotEmpty(t, decoded.Validators[0].Signer["secretKey"])
	})
}

func testMakeDefaultBreakpointGenesis(t *testing.T, validatorCount uint) {
	originalAdmin := params.CrossDev3ChainConfig.Istanbul.PoSA.Admin
	genesis, keys, err := makeDefaultBreakpointGenesis(validatorCount)
	require.NoError(t, err)
	require.Len(t, keys.validators, int(validatorCount))
	require.Len(t, genesis.Config.Istanbul.Validators, int(validatorCount))
	require.Len(t, genesis.Config.Istanbul.PoSA.Validators, int(validatorCount))
	require.Equal(t, originalAdmin, params.CrossDev3ChainConfig.Istanbul.PoSA.Admin)
	expectedConfig := *params.CrossDev3ChainConfig
	expectedConfig.Istanbul = genesis.Config.Istanbul
	require.Equal(t, &expectedConfig, genesis.Config)
	expectedIstanbul := *params.CrossDev3ChainConfig.Istanbul
	expectedPoSA := *expectedIstanbul.PoSA
	expectedIstanbul.PoSA = &expectedPoSA
	expectedIstanbul.Validators = genesis.Config.Istanbul.Validators
	expectedPoSA.Admin = genesis.Config.Istanbul.PoSA.Admin
	expectedPoSA.Validators = genesis.Config.Istanbul.PoSA.Validators
	require.Equal(t, &expectedIstanbul, genesis.Config.Istanbul)

	admin := crypto.PubkeyToAddress(keys.admin.PublicKey)
	require.Equal(t, admin, genesis.Config.Istanbul.PoSA.Admin)
	require.Equal(t, new(big.Int).Mul(big.NewInt(100_000_000_000), big.NewInt(params.Ether)), genesis.Alloc[admin].Balance)

	addresses := make(map[common.Address]struct{})
	for i, validatorKeys := range keys.validators {
		validator := genesis.Config.Istanbul.PoSA.Validators[i]
		require.Equal(t, crypto.PubkeyToAddress(validatorKeys.validator.PublicKey), validator.Validator)
		require.Equal(t, crypto.PubkeyToAddress(validatorKeys.operator.PublicKey), validator.Operator)
		require.Equal(t, validatorKeys.signer.PublicKey().Marshal(), []byte(validator.Signer))
		require.Equal(t, validator.Validator, genesis.Config.Istanbul.Validators[i])
		addresses[validator.Validator] = struct{}{}
		addresses[validator.Operator] = struct{}{}
	}
	require.Len(t, addresses, int(validatorCount*2))
}

// Tests that initializing Geth with a custom genesis block and chain definitions
// work properly.
func TestCustomGenesis(t *testing.T) {
	t.Parallel()
	for i, tt := range customGenesisTests {
		// Create a temporary data directory to use and inspect later
		datadir := t.TempDir()

		// Initialize the data directory with the custom genesis block
		json := filepath.Join(datadir, "genesis.json")
		if err := os.WriteFile(json, []byte(tt.genesis), 0600); err != nil {
			t.Fatalf("test %d: failed to write genesis file: %v", i, err)
		}
		runGeth(t, "--datadir", datadir, "init", json).WaitExit()

		// Query the custom genesis block
		geth := runGeth(t, "--networkid", "1337", "--syncmode=full", "--cache", "16",
			"--datadir", datadir, "--maxpeers", "0", "--port", "0", "--authrpc.port", "0",
			"--nodiscover", "--nat", "none", "--ipcdisable",
			"--exec", tt.query, "console")
		geth.ExpectRegexp(tt.result)
		geth.ExpectExit()
	}
}

// TestCustomBackend that the backend selection and detection (leveldb vs pebble) works properly.
func TestCustomBackend(t *testing.T) {
	t.Parallel()
	// Test pebble, but only on 64-bit platforms
	if strconv.IntSize != 64 {
		t.Skip("Custom backends are only available on 64-bit platform")
	}
	genesis := `{
		"alloc"      : {},
		"coinbase"   : "0x0000000000000000000000000000000000000000",
			"difficulty" : "0x20000",
			"extraData"  : "",
			"gasLimit"   : "0x2fefd8",
			"nonce"      : "0x0000000000001338",
			"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
			"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
			"timestamp"  : "0x00",
			"config": {
				"terminalTotalDifficulty": 0
			}
		}`
	type backendTest struct {
		initArgs   []string
		initExpect string
		execArgs   []string
		execExpect string
	}
	testfunc := func(t *testing.T, tt backendTest) error {
		// Create a temporary data directory to use and inspect later
		datadir := t.TempDir()

		// Initialize the data directory with the custom genesis block
		json := filepath.Join(datadir, "genesis.json")
		if err := os.WriteFile(json, []byte(genesis), 0600); err != nil {
			return fmt.Errorf("failed to write genesis file: %v", err)
		}
		{ // Init
			args := append(tt.initArgs, "--datadir", datadir, "init", json)
			geth := runGeth(t, args...)
			geth.ExpectRegexp(tt.initExpect)
			geth.ExpectExit()
		}
		{ // Exec + query
			args := append(tt.execArgs, "--networkid", "1337", "--syncmode=full", "--cache", "16",
				"--datadir", datadir, "--maxpeers", "0", "--port", "0", "--authrpc.port", "0",
				"--nodiscover", "--nat", "none", "--ipcdisable",
				"--exec", "eth.getBlock(0).nonce", "console")
			geth := runGeth(t, args...)
			geth.ExpectRegexp(tt.execExpect)
			geth.ExpectExit()
		}
		return nil
	}
	for i, tt := range []backendTest{
		{ // When not specified, it should default to pebble
			execArgs:   []string{"--db.engine", "pebble"},
			execExpect: "0x0000000000001338",
		},
		{ // Explicit leveldb
			initArgs:   []string{"--db.engine", "leveldb"},
			execArgs:   []string{"--db.engine", "leveldb"},
			execExpect: "0x0000000000001338",
		},
		{ // Explicit leveldb first, then autodiscover
			initArgs:   []string{"--db.engine", "leveldb"},
			execExpect: "0x0000000000001338",
		},
		{ // Explicit pebble
			initArgs:   []string{"--db.engine", "pebble"},
			execArgs:   []string{"--db.engine", "pebble"},
			execExpect: "0x0000000000001338",
		},
		{ // Explicit pebble, then auto-discover
			initArgs:   []string{"--db.engine", "pebble"},
			execExpect: "0x0000000000001338",
		},
		{ // Can't start pebble on top of leveldb
			initArgs:   []string{"--db.engine", "leveldb"},
			execArgs:   []string{"--db.engine", "pebble"},
			execExpect: `Fatal: Failed to register the Ethereum service: db.engine choice was pebble but found pre-existing leveldb database in specified data directory`,
		},
		{ // Can't start leveldb on top of pebble
			initArgs:   []string{"--db.engine", "pebble"},
			execArgs:   []string{"--db.engine", "leveldb"},
			execExpect: `Fatal: Failed to register the Ethereum service: db.engine choice was leveldb but found pre-existing pebble database in specified data directory`,
		},
		{ // Reject invalid backend choice
			initArgs:   []string{"--db.engine", "mssql"},
			initExpect: `Fatal: Invalid choice for db.engine 'mssql', allowed 'leveldb' or 'pebble'`,
			// Since the init fails, this will return the (default) cross mainnet genesis
			// block nonce
			execExpect: `0x000000000000aaaa`, // ##CROSS: config
		},
	} {
		if err := testfunc(t, tt); err != nil {
			t.Fatalf("test %d-leveldb: %v", i, err)
		}
	}
}
