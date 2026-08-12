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

package rawdb

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
)

// ##CROSS: freezer difficulty migration

func TestFreezer_MigrateDifficultyTable(t *testing.T) {
	headers := makeHeaders(5)

	t.Run("migrates", func(t *testing.T) {
		legacyTables := map[string]freezerTableConfig{
			ChainFreezerHeaderTable:  {noSnappy: false, prunable: false},
			ChainFreezerHashTable:    {noSnappy: true, prunable: false},
			ChainFreezerBodiesTable:  {noSnappy: false, prunable: true},
			ChainFreezerReceiptTable: {noSnappy: false, prunable: true},
		}

		f, dir := newFreezerForTesting(t, legacyTables)

		require.NoError(t, writeTestAncientChain(f, headers, false))
		require.NoError(t, f.Close())

		f = newFreezerForTestingWithDir(t, chainFreezerTableConfigs, dir)
		defer f.Close()

		require.Equal(t, uint64(len(headers)), f.frozen.Load())
		require.Equal(t, uint64(len(headers)), f.tables[ChainFreezerDifficultyTable].items.Load())
		require.Equal(t, uint64(0), f.tables[ChainFreezerDifficultyTable].itemHidden.Load())
		checkAncientDifficulty(t, f, 0, 1)
		checkAncientDifficulty(t, f, 1, 3)
		checkAncientDifficulty(t, f, 4, 15)
	})

	t.Run("regenerates pruned table", func(t *testing.T) {
		legacyTables := map[string]freezerTableConfig{
			ChainFreezerHeaderTable:      {noSnappy: false, prunable: false},
			ChainFreezerHashTable:        {noSnappy: true, prunable: false},
			ChainFreezerBodiesTable:      {noSnappy: false, prunable: true},
			ChainFreezerReceiptTable:     {noSnappy: false, prunable: true},
			ChainFreezerDifficultyTable:  {noSnappy: true, prunable: true},
			ChainFreezerBlobSidecarTable: {noSnappy: false, prunable: true},
		}

		f, dir := newFreezerForTesting(t, legacyTables)

		require.NoError(t, writeTestAncientChain(f, headers, true))
		_, err := f.TruncateTail(2)
		require.NoError(t, err)
		require.NoError(t, f.Close())

		f = newFreezerForTestingWithDir(t, chainFreezerTableConfigs, dir)
		defer f.Close()

		require.Equal(t, uint64(len(headers)), f.frozen.Load())
		require.Equal(t, uint64(len(headers)), f.tables[ChainFreezerDifficultyTable].items.Load())
		require.Equal(t, uint64(0), f.tables[ChainFreezerDifficultyTable].itemHidden.Load())
		checkAncientDifficulty(t, f, 0, 1)
		checkAncientDifficulty(t, f, 4, 15)
	})
}

func makeHeaders(count int) []*types.Header {
	headers := make([]*types.Header, count)
	for i := range count {
		headers[i] = &types.Header{
			Number:     big.NewInt(int64(i)),
			Difficulty: big.NewInt(int64(i + 1)),
		}
	}
	return headers
}

func writeTestAncientChain(f *Freezer, headers []*types.Header, withDifficulty bool) error {
	_, err := f.ModifyAncients(func(op ethdb.AncientWriteOp) error {
		td := new(big.Int)
		for number, header := range headers {
			num := uint64(number)
			td.Add(td, header.Difficulty)
			if err := op.AppendRaw(ChainFreezerHashTable, num, header.Hash().Bytes()); err != nil {
				return err
			}
			if err := op.Append(ChainFreezerHeaderTable, num, header); err != nil {
				return err
			}
			if err := op.AppendRaw(ChainFreezerBodiesTable, num, rlp.EmptyList); err != nil {
				return err
			}
			if err := op.AppendRaw(ChainFreezerReceiptTable, num, rlp.EmptyList); err != nil {
				return err
			}
			if withDifficulty {
				if err := op.Append(ChainFreezerDifficultyTable, num, new(big.Int).Set(td)); err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

func checkAncientDifficulty(t *testing.T, f *Freezer, number uint64, want int64) {
	t.Helper()

	data, err := f.Ancient(ChainFreezerDifficultyTable, number)
	require.NoError(t, err)

	td := new(big.Int)
	require.NoError(t, rlp.DecodeBytes(data, td))
	require.Zero(t, td.Cmp(big.NewInt(want)))
}
