// Copyright 2017 The go-ethereum Authors
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

package backend

import (
	"bytes"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/engine"
	istanbulEngine "github.com/ethereum/go-ethereum/consensus/istanbul/engine"
	"github.com/ethereum/go-ethereum/consensus/istanbul/testutils"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testerVote struct {
	validator string
	voted     string
	auth      bool
}

// testerAccountPool is a pool to maintain currently active tester accounts,
// mapped from textual names used in the tests below to actual Ethereum private
// keys capable of signing transactions.
type testerAccountPool struct {
	accounts map[string]*ecdsa.PrivateKey
}

func newTesterAccountPool() *testerAccountPool {
	return &testerAccountPool{
		accounts: make(map[string]*ecdsa.PrivateKey),
	}
}

func (ap *testerAccountPool) writeValidatorVote(header *types.Header, validator string, recipientAddress string, authorize bool) error {
	return engine.ApplyHeaderIstanbulExtra(
		header,
		engine.WriteVote(ap.address(recipientAddress), authorize),
	)
}

func (ap *testerAccountPool) address(account string) common.Address {
	// Ensure we have a persistent key for the account
	if ap.accounts[account] == nil {
		ap.accounts[account], _ = crypto.GenerateKey()
	}
	// Resolve and return the Ethereum address
	return crypto.PubkeyToAddress(ap.accounts[account].PublicKey)
}

// Tests that voting is evaluated correctly for various simple and complex scenarios.
func TestVoting(t *testing.T) {
	// Define the various voting scenarios to test
	tests := []struct {
		epoch      uint64
		validators []string
		votes      []testerVote
		results    []string
	}{
		{
			// Single validator, no votes cast
			validators: []string{"A"},
			votes:      []testerVote{{validator: "A"}},
			results:    []string{"A"},
		}, {
			// Single validator, voting to add two others (only accept first, second needs 2 votes)
			validators: []string{"A"},
			votes: []testerVote{
				{validator: "A", voted: "B", auth: true},
				{validator: "B"},
				{validator: "A", voted: "C", auth: true},
			},
			results: []string{"A", "B"},
		}, {
			// Two validators, voting to add three others (only accept first two, third needs 3 votes already)
			validators: []string{"A", "B"},
			votes: []testerVote{
				{validator: "A", voted: "C", auth: true},
				{validator: "B", voted: "C", auth: true},
				{validator: "A", voted: "D", auth: true},
				{validator: "B", voted: "D", auth: true},
				{validator: "C"},
				{validator: "A", voted: "E", auth: true},
				{validator: "B", voted: "E", auth: true},
			},
			results: []string{"A", "B", "C", "D"},
		}, {
			// Single validator, dropping itself (weird, but one less cornercase by explicitly allowing this)
			validators: []string{"A"},
			votes: []testerVote{
				{validator: "A", voted: "A", auth: false},
			},
			results: []string{},
		}, {
			// Two validators, actually needing mutual consent to drop either of them (not fulfilled)
			validators: []string{"A", "B"},
			votes: []testerVote{
				{validator: "A", voted: "B", auth: false},
			},
			results: []string{"A", "B"},
		}, {
			// Two validators, actually needing mutual consent to drop either of them (fulfilled)
			validators: []string{"A", "B"},
			votes: []testerVote{
				{validator: "A", voted: "B", auth: false},
				{validator: "B", voted: "B", auth: false},
			},
			results: []string{"A"},
		}, {
			// Three validators, two of them deciding to drop the third
			validators: []string{"A", "B", "C"},
			votes: []testerVote{
				{validator: "A", voted: "C", auth: false},
				{validator: "B", voted: "C", auth: false},
			},
			results: []string{"A", "B"},
		}, {
			// Four validators, consensus of two not being enough to drop anyone
			validators: []string{"A", "B", "C", "D"},
			votes: []testerVote{
				{validator: "A", voted: "C", auth: false},
				{validator: "B", voted: "C", auth: false},
			},
			results: []string{"A", "B", "C", "D"},
		}, {
			// Four validators, consensus of three already being enough to drop someone
			validators: []string{"A", "B", "C", "D"},
			votes: []testerVote{
				{validator: "A", voted: "D", auth: false},
				{validator: "B", voted: "D", auth: false},
				{validator: "C", voted: "D", auth: false},
			},
			results: []string{"A", "B", "C"},
		}, {
			// Authorizations are counted once per validator per target
			validators: []string{"A", "B"},
			votes: []testerVote{
				{validator: "A", voted: "C", auth: true},
				{validator: "B"},
				{validator: "A", voted: "C", auth: true},
				{validator: "B"},
				{validator: "A", voted: "C", auth: true},
			},
			results: []string{"A", "B"},
		}, {
			// Authorizing multiple accounts concurrently is permitted
			validators: []string{"A", "B"},
			votes: []testerVote{
				{validator: "A", voted: "C", auth: true},
				{validator: "B"},
				{validator: "A", voted: "D", auth: true},
				{validator: "B"},
				{validator: "A"},
				{validator: "B", voted: "D", auth: true},
				{validator: "A"},
				{validator: "B", voted: "C", auth: true},
			},
			results: []string{"A", "B", "C", "D"},
		}, {
			// Deauthorizations are counted once per validator per target
			validators: []string{"A", "B"},
			votes: []testerVote{
				{validator: "A", voted: "B", auth: false},
				{validator: "B"},
				{validator: "A", voted: "B", auth: false},
				{validator: "B"},
				{validator: "A", voted: "B", auth: false},
			},
			results: []string{"A", "B"},
		}, {
			// Deauthorizing multiple accounts concurrently is permitted
			validators: []string{"A", "B", "C", "D"},
			votes: []testerVote{
				{validator: "A", voted: "C", auth: false},
				{validator: "B"},
				{validator: "C"},
				{validator: "A", voted: "D", auth: false},
				{validator: "B"},
				{validator: "C"},
				{validator: "A"},
				{validator: "B", voted: "D", auth: false},
				{validator: "C", voted: "D", auth: false},
				{validator: "A"},
				{validator: "B", voted: "C", auth: false},
			},
			results: []string{"A", "B"},
		}, {
			// Votes from deauthorized validators are discarded immediately (deauth votes)
			validators: []string{"A", "B", "C"},
			votes: []testerVote{
				{validator: "C", voted: "B", auth: false},
				{validator: "A", voted: "C", auth: false},
				{validator: "B", voted: "C", auth: false},
				{validator: "A", voted: "B", auth: false},
			},
			results: []string{"A", "B"},
		}, {
			// Votes from deauthorized validators are discarded immediately (auth votes)
			validators: []string{"A", "B", "C"},
			votes: []testerVote{
				{validator: "C", voted: "B", auth: false},
				{validator: "A", voted: "C", auth: false},
				{validator: "B", voted: "C", auth: false},
				{validator: "A", voted: "B", auth: false},
			},
			results: []string{"A", "B"},
		}, {
			// Cascading changes are not allowed, only the the account being voted on may change
			validators: []string{"A", "B", "C", "D"},
			votes: []testerVote{
				{validator: "A", voted: "C", auth: false},
				{validator: "B"},
				{validator: "C"},
				{validator: "A", voted: "D", auth: false},
				{validator: "B", voted: "C", auth: false},
				{validator: "C"},
				{validator: "A"},
				{validator: "B", voted: "D", auth: false},
				{validator: "C", voted: "D", auth: false},
			},
			results: []string{"A", "B", "C"},
		}, {
			// Changes reaching consensus out of bounds (via a deauth) execute on touch
			validators: []string{"A", "B", "C", "D"},
			votes: []testerVote{
				{validator: "A", voted: "C", auth: false},
				{validator: "B"},
				{validator: "C"},
				{validator: "A", voted: "D", auth: false},
				{validator: "B", voted: "C", auth: false},
				{validator: "C"},
				{validator: "A"},
				{validator: "B", voted: "D", auth: false},
				{validator: "C", voted: "D", auth: false},
				{validator: "A"},
				{validator: "C", voted: "C", auth: true},
			},
			results: []string{"A", "B"},
		}, {
			// Changes reaching consensus out of bounds (via a deauth) may go out of consensus on first touch
			validators: []string{"A", "B", "C", "D"},
			votes: []testerVote{
				{validator: "A", voted: "C", auth: false},
				{validator: "B"},
				{validator: "C"},
				{validator: "A", voted: "D", auth: false},
				{validator: "B", voted: "C", auth: false},
				{validator: "C"},
				{validator: "A"},
				{validator: "B", voted: "D", auth: false},
				{validator: "C", voted: "D", auth: false},
				{validator: "A"},
				{validator: "B", voted: "C", auth: true},
			},
			results: []string{"A", "B", "C"},
		}, {
			// Ensure that pending votes don't survive authorization status changes. This
			// corner case can only appear if a validator is quickly added, remove and then
			// readded (or the inverse), while one of the original voters dropped. If a
			// past vote is left cached in the system somewhere, this will interfere with
			// the final validator outcome.
			validators: []string{"A", "B", "C", "D", "E"},
			votes: []testerVote{
				{validator: "A", voted: "F", auth: true}, // Authorize F, 3 votes needed
				{validator: "B", voted: "F", auth: true},
				{validator: "C", voted: "F", auth: true},
				{validator: "D", voted: "F", auth: false}, // Deauthorize F, 4 votes needed (leave A's previous vote "unchanged")
				{validator: "E", voted: "F", auth: false},
				{validator: "B", voted: "F", auth: false},
				{validator: "C", voted: "F", auth: false},
				{validator: "D", voted: "F", auth: true}, // Almost authorize F, 2/3 votes needed
				{validator: "E", voted: "F", auth: true},
				{validator: "B", voted: "A", auth: false}, // Deauthorize A, 3 votes needed
				{validator: "C", voted: "A", auth: false},
				{validator: "D", voted: "A", auth: false},
				{validator: "B", voted: "F", auth: true}, // Finish authorizing F, 3/3 votes needed
			},
			results: []string{"B", "C", "D", "E", "F"},
		}, {
			// Epoch transitions reset all votes to allow chain checkpointing
			epoch:      3,
			validators: []string{"A", "B"},
			votes: []testerVote{
				{validator: "A", voted: "C", auth: true},
				{validator: "B"},
				{validator: "A"}, // Checkpoint block, (don't vote here, it's validated outside of snapshots)
				{validator: "B", voted: "C", auth: true},
			},
			results: []string{"A", "B"},
		},
	}

	// Run through the scenarios and test them
	for i, tt := range tests {
		// Create the account pool and generate the initial set of validators
		accounts := newTesterAccountPool()

		validators := make([]common.Address, len(tt.validators))
		for j, validator := range tt.validators {
			validators[j] = accounts.address(validator)
		}
		for j := 0; j < len(validators); j++ {
			for k := j + 1; k < len(validators); k++ {
				if bytes.Compare(validators[j][:], validators[k][:]) > 0 {
					validators[j], validators[k] = validators[k], validators[j]
				}
			}
		}

		genesis := testutils.Genesis(validators)
		config := new(istanbul.Config)
		*config = *istanbul.DefaultConfig
		if tt.epoch != 0 {
			config.Epoch = tt.epoch
		}

		chain, backend := newBlockchainFromConfig(
			genesis,
			[]*ecdsa.PrivateKey{accounts.accounts[tt.validators[0]]},
			config,
		)

		// Assemble a chain of headers from the cast votes
		headers := make([]*types.Header, len(tt.votes))
		for j, vote := range tt.votes {
			blockNumber := big.NewInt(int64(j) + 1)
			headers[j] = &types.Header{
				Number:     blockNumber,
				Time:       uint64(int64(j) * int64(config.GetConfig(blockNumber).BlockPeriod)),
				Coinbase:   accounts.address(vote.validator),
				Difficulty: istanbul.DefaultDifficulty,
				MixDigest:  types.IstanbulDigest,
			}

			headers[j].Extra = make([]byte, len(genesis.ExtraData))
			copy(headers[j].Extra, genesis.ExtraData)

			_ = engine.ApplyHeaderIstanbulExtra(
				headers[j],
				engine.WriteValidators(validators),
			)

			if len(vote.voted) > 0 {
				if err := accounts.writeValidatorVote(headers[j], vote.validator, vote.voted, vote.auth); err != nil {
					t.Errorf("Error writeValidatorVote test: %d, validator: %s, voteType: %v (err=%v)", j, vote.voted, vote.auth, err)
				}
			}
		}

		// Set ParentHash after all headers are fully configured
		genesisBlock := chain.GetBlockByNumber(0)
		if genesisBlock == nil {
			t.Errorf("test %d: failed to get genesis block", i)
			backend.Stop()
			continue
		}
		for j := range headers {
			if j == 0 {
				headers[j].ParentHash = genesisBlock.Hash()
			} else {
				headers[j].ParentHash = headers[j-1].Hash()
			}
		}

		// Pass all the headers through clique and ensure tallying succeeds
		head := headers[len(headers)-1]

		snap, err := backend.snapshot(chain, head.Number.Uint64(), head.Hash(), headers)
		if err != nil {
			t.Errorf("test %d: failed to create voting snapshot: %v", i, err)
			backend.Stop()
			continue
		}
		// Verify the final list of validators against the expected ones
		validators = make([]common.Address, len(tt.results))
		for j, validator := range tt.results {
			validators[j] = accounts.address(validator)
		}
		for j := 0; j < len(validators); j++ {
			for k := j + 1; k < len(validators); k++ {
				if bytes.Compare(validators[j][:], validators[k][:]) > 0 {
					validators[j], validators[k] = validators[k], validators[j]
				}
			}
		}
		result := snap.validators()
		if len(result) != len(validators) {
			t.Errorf("test %d: validators mismatch: have %x, want %x", i, result, validators)
			backend.Stop()
			continue
		}
		for j := 0; j < len(result); j++ {
			if !bytes.Equal(result[j][:], validators[j][:]) {
				t.Errorf("test %d, validator %d: validator mismatch: have %x, want %x", i, j, result[j], validators[j])
			}
		}
		backend.Stop()
	}
}

func TestSnapshot(t *testing.T) {
	t.Run("injects PoSA signers when batch parent is not stored", func(t *testing.T) {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)

		validatorAddress := crypto.PubkeyToAddress(key.PublicKey)
		expectedSigner := types.BytesToBLSPublicKey([]byte{1})
		breakpointTime := uint64(2)

		genesis := testutils.Genesis([]common.Address{validatorAddress})
		chainConfig := *genesis.Config
		istanbulConfig := *genesis.Config.Istanbul
		istanbulConfig.PoSA = &params.PoSAConfig{
			Validators: []params.PoSAValidator{
				{
					Validator: validatorAddress,
					Signer:    expectedSigner[:],
				},
			},
		}
		chainConfig.BreakpointTime = &breakpointTime
		chainConfig.ShanghaiTime = new(uint64)
		chainConfig.Istanbul = &istanbulConfig
		genesis.Config = &chainConfig

		config := copyConfig(istanbul.DefaultConfig)
		chain, backend := newBlockchainFromConfig(genesis, []*ecdsa.PrivateKey{key}, config)
		defer backend.Stop()

		genesisHeader := chain.GetHeaderByNumber(0)
		require.NotNil(t, genesisHeader)

		baseHeader := &types.Header{
			ParentHash: genesisHeader.Hash(),
			Number:     big.NewInt(1),
			Time:       breakpointTime - 1,
			Coinbase:   validatorAddress,
			Difficulty: istanbul.DefaultDifficulty,
			MixDigest:  types.IstanbulDigest,
			Extra:      append([]byte(nil), genesis.ExtraData...),
		}
		breakpointHeader := &types.Header{
			ParentHash: baseHeader.Hash(),
			Number:     big.NewInt(2),
			Time:       breakpointTime,
			Coinbase:   validatorAddress,
			Difficulty: istanbul.DefaultDifficulty,
			MixDigest:  types.IstanbulDigest,
			Extra:      append([]byte(nil), genesis.ExtraData...),
		}

		baseSnapshot := newSnapshot(
			config.Epoch,
			baseHeader.Number.Uint64(),
			baseHeader.Hash(),
			validator.NewSet(
				[]common.Address{validatorAddress},
				nil,
				config.ProposerPolicy,
			),
		)
		backend.recents.Add(baseSnapshot.Hash, baseSnapshot)

		require.Nil(t, chain.GetHeader(baseHeader.Hash(), baseHeader.Number.Uint64()))

		snapshot, err := backend.snapshot(
			chain,
			breakpointHeader.Number.Uint64(),
			breakpointHeader.Hash(),
			[]*types.Header{baseHeader, breakpointHeader},
		)
		require.NoError(t, err)

		_, actualValidator := snapshot.ValSet.GetByAddress(validatorAddress)
		require.NotNil(t, actualValidator)
		assert.Equal(t, expectedSigner, actualValidator.SignerAddress())
	})
}

func TestSaveAndLoad(t *testing.T) {
	validators := []common.Address{
		common.BytesToAddress([]byte("1234567894")),
		common.BytesToAddress([]byte("1234567895")),
	}
	newTestSnapshot := func(signers []types.BLSPublicKey) *Snapshot {
		return &Snapshot{
			Epoch:  5,
			Number: 10,
			Hash:   common.HexToHash("1234567890"),
			Votes: []*Vote{
				{
					Validator: common.BytesToAddress([]byte("1234567891")),
					Block:     15,
					Address:   common.BytesToAddress([]byte("1234567892")),
					Authorize: false,
				},
			},
			Tally: map[common.Address]Tally{
				common.BytesToAddress([]byte("1234567893")): {
					Authorize: false,
					Votes:     20,
				},
			},
			ValSet: validator.NewSet(validators, signers, istanbul.NewRoundRobinProposerPolicy()),
		}
	}
	assertSnapshot := func(t *testing.T, expected, actual *Snapshot) {
		t.Helper()

		assert.Equal(t, expected.Epoch, actual.Epoch)
		assert.Equal(t, expected.Number, actual.Number)
		assert.Equal(t, expected.Hash, actual.Hash)
		assert.Equal(t, expected.Votes, actual.Votes)
		assert.Equal(t, expected.Tally, actual.Tally)
		assert.Equal(t, expected.ValSet.Policy().Id, actual.ValSet.Policy().Id)
		require.Equal(t, expected.ValSet.Size(), actual.ValSet.Size())
		for i, expectedVal := range expected.ValSet.List() {
			actualVal := actual.ValSet.List()[i]
			assert.Equal(t, expectedVal.Address(), actualVal.Address())
			assert.Equal(t, expectedVal.SignerAddress(), actualVal.SignerAddress())
		}
	}

	t.Run("no signers", func(t *testing.T) {
		snap := newTestSnapshot(nil)
		db := rawdb.NewMemoryDatabase()
		require.NoError(t, snap.store(db))

		// Stored JSON does not contain "signers" field.
		blob, err := db.Get(append([]byte(dbKeySnapshotPrefix), snap.Hash[:]...))
		require.NoError(t, err)
		assert.NotContains(t, string(blob), `"signers"`)

		loaded, err := loadSnapshot(snap.Epoch, db, snap.Hash)
		require.NoError(t, err)
		assertSnapshot(t, snap, loaded)
	})

	// ##CROSS: bls seal
	t.Run("with signers", func(t *testing.T) {
		snap := newTestSnapshot([]types.BLSPublicKey{
			types.BytesToBLSPublicKey([]byte{1}),
			types.BytesToBLSPublicKey([]byte{2}),
		})
		db := rawdb.NewMemoryDatabase()
		require.NoError(t, snap.store(db))

		loaded, err := loadSnapshot(snap.Epoch, db, snap.Hash)
		require.NoError(t, err)
		assertSnapshot(t, snap, loaded)
	})
	// ##
}

// ##CROSS: validator sort policy
func TestSnapshotLoadValidatorsByteSorted(t *testing.T) {
	// These two addresses are byte-ordered as 0x...0b < 0x...0c, but their
	// EIP-55 string forms sort in the opposite order because "C" < "b".
	addrA := common.HexToAddress("0x000000000000000000000000000000000000000b")
	addrB := common.HexToAddress("0x000000000000000000000000000000000000000c")
	signerA := types.BytesToBLSPublicKey(bytes.Repeat([]byte{0x11}, types.BLSPublicKeyLength))
	signerB := types.BytesToBLSPublicKey(bytes.Repeat([]byte{0x22}, types.BLSPublicKeyLength))

	// Using string policy, it should be string-ordered
	stringPolicy := istanbul.NewRoundRobinProposerPolicy()
	stringPolicy.Use(istanbul.ValidatorSortByString())
	stringSet := validator.NewSet(
		[]common.Address{addrA, addrB},
		[]types.BLSPublicKey{signerA, signerB},
		stringPolicy,
	)
	stringList := stringSet.List()
	if stringList[0].Address() != addrB || stringList[1].Address() != addrA {
		t.Fatalf("test fixture must differ between string and byte order, got %v then %v", stringList[0].Address(), stringList[1].Address())
	}

	// Using byte policy, it should be byte-ordered
	bytePolicy := istanbul.NewRoundRobinProposerPolicy()
	bytePolicy.Use(istanbul.ValidatorSortByByte())

	liveSet := validator.NewSet(
		[]common.Address{addrA, addrB},
		[]types.BLSPublicKey{signerA, signerB},
		bytePolicy,
	)
	liveList := liveSet.List()
	if liveList[0].Address() != addrA || liveList[1].Address() != addrB {
		t.Fatalf("expected live validator set to stay byte-ordered, got %v then %v", liveList[0].Address(), liveList[1].Address())
	}

	// Restore from DB, it should be byte-ordered
	db := rawdb.NewMemoryDatabase()
	snap := &Snapshot{
		Epoch:  10,
		Number: 1,
		Hash:   common.HexToHash("0x01"),
		ValSet: liveSet,
	}
	if err := snap.store(db); err != nil {
		t.Fatalf("store snapshot: %v", err)
	}
	restored, err := loadSnapshot(10, db, snap.Hash)
	if err != nil {
		t.Fatalf("load snapshot: %v", err)
	}

	restoredList := restored.ValSet.List()
	if restoredList[0].Address() != addrA || restoredList[1].Address() != addrB {
		t.Fatalf("expected restored validator set to stay byte-ordered, got %v then %v", restoredList[0].Address(), restoredList[1].Address())
	}

	// Restore signers from the header, they should be byte-ordered
	header := &types.Header{}
	if err := istanbulEngine.ApplyHeaderIstanbulExtra(header, func(extra *types.IstanbulExtra) error {
		extra.SignersBitset = []uint64{1} // Select byte-sorted validator index 0.
		return nil
	}); err != nil {
		t.Fatalf("set signers bitset: %v", err)
	}

	var e istanbulEngine.Engine
	liveAddrs, livePubs, err := e.BLSSigners(header, liveSet)
	if err != nil {
		t.Fatalf("decode signers with live set: %v", err)
	}

	reloadedAddrs, reloadedPubs, err := e.BLSSigners(header, restored.ValSet)
	if err != nil {
		t.Fatalf("decode signers with restored set: %v", err)
	}

	if len(liveAddrs) != 1 || len(livePubs) != 1 || liveAddrs[0] != addrA || livePubs[0] != signerA {
		t.Fatalf("expected live decode to resolve bit 0 to %v/%v, got %v/%v", addrA, signerA, liveAddrs, livePubs)
	}
	if len(reloadedAddrs) != 1 || len(reloadedPubs) != 1 || reloadedAddrs[0] != addrA || reloadedPubs[0] != signerA {
		t.Fatalf("expected restored decode to resolve bit 0 to %v/%v, got %v/%v", addrA, signerA, reloadedAddrs, reloadedPubs)
	}
}

// ##
