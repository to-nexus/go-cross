// Copyright 2025 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package contracts

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/breakpoint"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/params/forks"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ##CROSS: contract upgrade

func TestSystemContractUpgrade_BreakpointMainnet(t *testing.T) {
	// Hashes come from the mainnet deployment artifacts in 1646df77cdfcf32e5f5c828560672f1384f590d2.
	const proxyHash = "0xcb3127f93664205d6698822f13c463ceff3a816614aa2b9ef5f7ca829ab338e0"
	expected := map[common.Address]string{
		CrossExAddr:            "0xb377190fa28822db6f642913214aac30fcb3d2d016cbe1145a6c691dff5d2684",
		ValidatorSetImplAddr:   "0xae5fe123044247938f39b75c97626330f3a453bdabf3dc4852c947bbb48a93fa",
		StakeHubImplAddr:       "0x7ff6b3d4253ce7ec4a480deed975af89c4d37f867c7f7a0486e31f3bd6a3cbac",
		RewardHubImplAddr:      "0xcf52900dfa03359b36ef0abdec7c84bee00ab851a9be87bea15f350d50478d3b",
		ValidatorSlashImplAddr: "0x6be02b209efaa8064b46fbd1dac9910f6a60f7e8574a9f39a4efa4eff18409b9",
		ValidatorSetAddr:       proxyHash,
		StakeHubAddr:           proxyHash,
		RewardHubAddr:          proxyHash,
		ValidatorSlashAddr:     proxyHash,
	}
	config := params.CrossChainConfig
	upgrade := getUpgrade(forks.Breakpoint, config.ChainID.Uint64())
	require.Len(t, upgrade.Configs, len(expected))
	statedb, err := state.New(types.EmptyRootHash, state.NewDatabaseForTesting())
	require.NoError(t, err)
	header := &types.Header{Number: big.NewInt(100), Time: *config.BreakpointTime}
	require.True(t, TryUpdateSystemContract(config, header, header.Time-1, statedb, false))
	for addr, hash := range expected {
		assert.Equal(t, common.HexToHash(hash), crypto.Keccak256Hash(statedb.GetCode(addr)), "code at %s", addr)
	}
}

func TestTryUpdateSystemContract(t *testing.T) {
	forkTime := uint64(time.Now().Unix())
	config := params.AllDevChainProtocolChanges
	config.PragueTime = &forkTime
	config.BreakpointTime = &forkTime
	config.Istanbul = &params.IstanbulConfig{PoSA: &params.PoSAConfig{}}

	statedb, _ := state.New(types.EmptyRootHash, state.NewDatabaseForTesting())

	header := &types.Header{
		Number: big.NewInt(100),
		Time:   forkTime + 1,
	}

	upgraded := TryUpdateSystemContract(config, header, forkTime-1, statedb, false)
	require.True(t, upgraded)

	t.Run("Prague upgrade", func(t *testing.T) {
		code := statedb.GetCode(params.HistoryStorageAddress)
		assert.Equal(t, params.HistoryStorageCode, code)
		nonce := statedb.GetNonce(params.HistoryStorageAddress)
		assert.Equal(t, uint64(1), nonce)
	})

	t.Run("Breakpoint upgrade", func(t *testing.T) {
		code := statedb.GetCode(ValidatorSetImplAddr)
		assert.Equal(t, common.FromHex(breakpoint.ValidatorSetMetaData.BinRuntime), code)
		nonce := statedb.GetNonce(ValidatorSetAddr)
		assert.Equal(t, uint64(1), nonce)
	})
}

// ##CROSS: consensus system contract
func TestInitSystemContract(t *testing.T) {
	validators := []common.Address{
		common.HexToAddress("0x000000000000000000000000000000000000aaaa"),
		common.HexToAddress("0x000000000000000000000000000000000000bbbb"),
		common.HexToAddress("0x000000000000000000000000000000000000cccc"),
	}
	// ##CROSS: bls seal
	signers := make([][]byte, 0, len(validators))
	for range len(validators) {
		signers = append(signers, types.BLSPublicKey{}.Bytes())
	}
	// ##
	extra := &types.IstanbulExtra{
		VanityData:    []byte{},
		Validators:    validators,
		CommittedSeal: [][]byte{},
		Round:         0,
		Vote:          nil,
		RandomReveal:  []byte{},
	}
	payload, err := rlp.EncodeToBytes(extra)
	require.NoError(t, err)

	admin := common.HexToAddress("0x000000000000000000000000000000000000ffff")
	pool := common.HexToAddress("0x0000000000000000000000000000000000001001")
	config := &params.ChainConfig{
		Istanbul: &params.IstanbulConfig{
			EpochLength:             86400,
			BlockPeriodSeconds:      1,
			EmptyBlockPeriodSeconds: 0,
			RequestTimeoutSeconds:   10,
			ProposerPolicy:          0,
			PoSA: &params.PoSAConfig{
				Admin:            admin,
				DelegationPool:   pool,
				RewardStartBlock: new(big.Int).SetUint64(100),
			},
		},
	}

	operators := validators
	ids := make([]string, 0, len(validators))
	for i := range validators {
		ids = append(ids, fmt.Sprintf("test%02d", i+1))
		config.Istanbul.PoSA.Validators = append(config.Istanbul.PoSA.Validators, params.PoSAValidator{
			ID:        ids[i],
			Operator:  operators[i],
			Validator: validators[i],
			Signer:    signers[i],
		})
	}

	tests := []struct {
		name     string
		fork     forks.Fork
		header   *types.Header
		expected []ContractInitData
	}{
		{
			name: "Breakpoint",
			fork: forks.Breakpoint,
			header: &types.Header{
				Number:   big.NewInt(100),
				Time:     uint64(time.Now().Unix()),
				GasLimit: 2e10,
				Extra:    payload,
			},
			expected: []ContractInitData{
				{
					To:   ValidatorSetAddr,
					Data: breakpoint.NewValidatorSet().PackInitialize(admin, extra.Validators, signers),
				},
				{
					To:   StakeHubAddr,
					Data: breakpoint.NewStakeHub().PackInitialize(pool, admin, operators, validators, signers, ids),
				},
				{
					To:   RewardHubAddr,
					Data: breakpoint.NewRewardHub().PackInitialize(pool, admin, big.NewInt(100)),
				},
				{
					To:   ValidatorSlashAddr,
					Data: breakpoint.NewValidatorSlash().PackInitialize(admin),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initData := getSystemContractInitialization(upgrades[tt.fork][0], config, tt.header)
			assert.Equal(t, tt.expected, initData)
		})
	}
}

// ##
