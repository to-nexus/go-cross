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
	"fmt"
	"math"
	"math/big"
	"slices"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

// ##CROSS: freezer difficulty migration

// migrateDifficultyTable backfills the non-prunable total difficulty freezer table from frozen headers.
func (f *Freezer) migrateDifficultyTable() error {
	headers := f.tables[ChainFreezerHeaderTable]
	diffs := f.tables[ChainFreezerDifficultyTable]
	if headers == nil || diffs == nil {
		return nil
	}
	if headers.itemHidden.Load() != 0 || headers.itemOffset.Load() != 0 {
		return fmt.Errorf("ancient headers have removed items: hidden=%d offset=%d", headers.itemHidden.Load(), headers.itemOffset.Load())
	}

	target := f.findDifficultyMigrationHead()
	if diffs.items.Load() == target && diffs.itemHidden.Load() == 0 && diffs.itemOffset.Load() == 0 {
		// Already migrated.
		return nil
	}
	if f.readonly {
		return fmt.Errorf("readonly: reopen writable to migrate: have head=%d hidden=%d offset=%d want head=%d",
			diffs.items.Load(), diffs.itemHidden.Load(), diffs.itemOffset.Load(), target)
	}

	if diffs.itemHidden.Load() != 0 || diffs.itemOffset.Load() != 0 || diffs.items.Load() > target {
		// Difficulty table is pruned or bigger than the header table - reset and backfill from scratch.
		var err error
		diffs, err = f.resetDifficultyTable(diffs)
		if err != nil {
			return err
		}
	}

	start := diffs.items.Load()
	if start == target {
		return nil
	}

	var (
		err error
		td  *big.Int
	)
	if start > 0 {
		td, err = f.difficultyAt(diffs, start-1)
	} else {
		td = new(big.Int)
	}
	if err != nil {
		// Failed to read the last difficulty - reset and backfill from scratch.
		log.Warn("Regenerating ancient difficulty table", "err", err)
		diffs, err = f.resetDifficultyTable(diffs)
		if err != nil {
			return err
		}
		start = 0
		td = new(big.Int)
	}

	started := time.Now()
	log.Info("Migrating ancient difficulty table", "from", start, "to", target)
	if err := backfillDifficultyTable(headers, diffs, start, target, td); err != nil {
		return err
	}
	log.Info("Done migrating ancient difficulty table", "items", target-start, "elapsed", time.Since(started))
	return nil
}

func (f *Freezer) findDifficultyMigrationHead() uint64 {
	head := uint64(math.MaxUint64)
	for kind, table := range f.tables {
		if kind == ChainFreezerDifficultyTable {
			continue
		}
		if slices.Contains(additionTables, kind) && EmptyTable(table) {
			continue
		}
		head = min(head, table.items.Load())
	}
	if head == math.MaxUint64 {
		return 0
	}
	return head
}

func (f *Freezer) resetDifficultyTable(diffs *freezerTable) (*freezerTable, error) {
	next, err := diffs.resetItems(0)
	if err != nil {
		return nil, err
	}
	f.tables[ChainFreezerDifficultyTable] = next
	return next, nil
}

func (f *Freezer) difficultyAt(diffs *freezerTable, number uint64) (*big.Int, error) {
	data, err := diffs.Retrieve(number)
	if err != nil {
		return nil, fmt.Errorf("failed to read ancient difficulty %d: %w", number, err)
	}
	td := new(big.Int)
	if err := rlp.DecodeBytes(data, td); err != nil {
		return nil, fmt.Errorf("failed to decode ancient difficulty %d: %w", number, err)
	}
	return td, nil
}

func backfillDifficultyTable(headers, diffs *freezerTable, start, limit uint64, td *big.Int) error {
	const (
		difficultyMigrationSyncInterval = 1_000_000
		difficultyMigrationLogInterval  = 8 * time.Second
	)

	batch := diffs.newBatch()
	started := time.Now()
	logged := started
	for number := start; number < limit; number++ {
		header, err := retrieveAncientHeader(headers, number)
		if err != nil {
			return err
		}
		td.Add(td, header.Difficulty)
		if err := batch.Append(number, new(big.Int).Set(td)); err != nil {
			return fmt.Errorf("failed to append ancient difficulty %d: %w", number, err)
		}
		processed := number - start + 1
		if processed%difficultyMigrationSyncInterval == 0 {
			// Checkpoint
			if err := batch.commit(); err != nil {
				return err
			}
			if err := diffs.Sync(); err != nil {
				return err
			}
		}
		if time.Since(logged) >= difficultyMigrationLogInterval {
			// Log progress
			log.Info("Backfilling ancient difficulty table", "current", number+1, "target", limit, "elapsed", time.Since(started))
			logged = time.Now()
		}
	}
	if err := batch.commit(); err != nil {
		return err
	}
	return diffs.Sync()
}

func retrieveAncientHeader(headers *freezerTable, number uint64) (*types.Header, error) {
	data, err := headers.Retrieve(number)
	if err != nil {
		return nil, fmt.Errorf("failed to read ancient header %d: %w", number, err)
	}
	header := new(types.Header)
	if err := rlp.DecodeBytes(data, header); err != nil {
		return nil, fmt.Errorf("failed to decode ancient header %d: %w", number, err)
	}
	if header.Number == nil || header.Number.Uint64() != number {
		return nil, fmt.Errorf("ancient header number mismatch: have %v want %d", header.Number, number)
	}
	if header.Difficulty == nil {
		return nil, fmt.Errorf("ancient header %d has nil difficulty", number)
	}
	return header, nil
}
