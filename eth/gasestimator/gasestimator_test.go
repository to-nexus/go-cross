// Copyright 2026 The go-ethereum Authors
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

package gasestimator

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/holiman/uint256"
)

type testChain struct {
	config *params.ChainConfig
}

func (c testChain) Engine() consensus.Engine { return ethash.NewFaker() }

func (c testChain) GetHeader(common.Hash, uint64) *types.Header { return nil }

func (c testChain) Config() *params.ChainConfig { return c.config }

func TestEstimateAllowsFeePayerToCoverGas(t *testing.T) {
	config := params.MergedTestChainConfig
	from := common.HexToAddress("0x1000000000000000000000000000000000000001")
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	feePayer := common.HexToAddress("0x3000000000000000000000000000000000000003")

	db := state.NewDatabase(triedb.NewDatabase(rawdb.NewMemoryDatabase(), nil), nil)
	statedb, err := state.New(types.EmptyRootHash, db)
	if err != nil {
		t.Fatal(err)
	}
	statedb.SetBalance(feePayer, new(uint256.Int).SetAllOne(), tracing.BalanceChangeUnspecified)

	estimate, _, err := Estimate(context.Background(), &core.Message{
		From:             from,
		To:               &to,
		Value:            new(big.Int),
		GasPrice:         big.NewInt(params.InitialBaseFee),
		GasFeeCap:        big.NewInt(params.InitialBaseFee),
		GasTipCap:        new(big.Int),
		FeePayer:         &feePayer,
		SkipNonceChecks:  true,
		SkipFromEOACheck: true,
	}, &Options{
		Config: config,
		Chain:  &testChain{config: config},
		Header: &types.Header{
			Number:     big.NewInt(1),
			GasLimit:   30_000_000,
			BaseFee:    big.NewInt(params.InitialBaseFee),
			Difficulty: new(big.Int),
		},
		State: statedb,
	}, 0)
	if err != nil {
		t.Fatalf("expected estimate to ignore sender gas balance: %v", err)
	}
	if estimate != params.TxGas {
		t.Fatalf("unexpected estimate: have %d want %d", estimate, params.TxGas)
	}
}
