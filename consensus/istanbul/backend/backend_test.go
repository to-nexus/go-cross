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
	"context"
	"crypto/ecdsa"
	"math/big"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	istcore "github.com/ethereum/go-ethereum/consensus/istanbul/core"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/bls"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSign(t *testing.T) {
	b := newBackend()
	defer b.Stop()
	data := []byte("Here is a string....")
	sig, err := b.Sign(data)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}
	//Check signature recover
	hashData := crypto.Keccak256(data)
	pubkey, _ := crypto.Ecrecover(hashData, sig)
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])
	if signer != getAddress() {
		t.Errorf("address mismatch: have %v, want %s", signer.Hex(), getAddress().Hex())
	}
}

func TestCheckSignature(t *testing.T) {
	key, _ := generatePrivateKey()
	data := []byte("Here is a string....")
	hashData := crypto.Keccak256(data)
	sig, _ := crypto.Sign(hashData, key)
	b := newBackend()
	defer b.Stop()
	a := getAddress()
	err := b.CheckSignature(data, a, sig)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}
	a = getInvalidAddress()
	err = b.CheckSignature(data, a, sig)
	if err != istanbul.ErrInvalidSignature {
		t.Errorf("error mismatch: have %v, want %v", err, istanbul.ErrInvalidSignature)
	}
}

func TestCheckValidatorSignature(t *testing.T) {
	vset, keys := newTestValidatorSet(5)

	// 1. Positive test: sign with validator's key should succeed
	data := []byte("dummy data")
	hashData := crypto.Keccak256(data)
	for i, k := range keys {
		// Sign
		sig, err := crypto.Sign(hashData, k)
		if err != nil {
			t.Errorf("error mismatch: have %v, want nil", err)
		}
		// CheckValidatorSignature should succeed
		addr, err := istanbul.CheckValidatorSignature(vset, data, sig)
		if err != nil {
			t.Errorf("error mismatch: have %v, want nil", err)
		}
		validator := vset.GetByIndex(uint64(i))
		if addr != validator.Address() {
			t.Errorf("validator address mismatch: have %v, want %v", addr, validator.Address())
		}
	}

	// 2. Negative test: sign with any key other than validator's key should return error
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}
	// Sign
	sig, err := crypto.Sign(hashData, key)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	// CheckValidatorSignature should return ErrUnauthorizedAddress
	addr, err := istanbul.CheckValidatorSignature(vset, data, sig)
	if err != istanbul.ErrUnauthorizedAddress {
		t.Errorf("error mismatch: have %v, want %v", err, istanbul.ErrUnauthorizedAddress)
	}
	emptyAddr := common.Address{}
	if addr != emptyAddr {
		t.Errorf("address mismatch: have %v, want %v", addr, emptyAddr)
	}
}

// ##CROSS: bls seal
type signedSeal []byte

func (c signedSeal) Index() uint            { return 0 }
func (c signedSeal) Signer() common.Address { return common.Address{} }
func (c signedSeal) Signature() []byte      { return c }

// ##

func TestCommit(t *testing.T) {
	chain, backend := newBlockChain(1)
	defer backend.Stop()
	block := updateIstanbulBlock(makeBlockWithoutSeal(chain, backend, chain.Genesis()), backend.Address())
	// As the proposer, Commit hands the sealed block to Seal through commitCh.
	proposed := block.Hash()
	backend.proposedBlockHash.Store(&proposed)

	t.Run("sends the sealed block to the sealer", func(t *testing.T) {
		err := backend.Commit(block, []istanbul.SignedSeal{commitSeal(t, backend.privateKey, block)}, big.NewInt(0))
		require.NoError(t, err)
		assert.Equal(t, block.Hash(), (<-backend.commitCh).Hash())
		assert.False(t, backend.HasBadProposal(block.Hash()))
	})

	t.Run("rejects missing seals", func(t *testing.T) {
		err := backend.Commit(block, nil, big.NewInt(0))
		require.ErrorIs(t, err, istanbul.ErrInvalidCommittedSeals)
	})

	// ##CROSS: istanbul validation
	t.Run("marks a block that fails header verification as bad", func(t *testing.T) {
		other, err := crypto.GenerateKey()
		require.NoError(t, err)
		err = backend.Commit(block, []istanbul.SignedSeal{commitSeal(t, other, block)}, big.NewInt(0))
		require.ErrorIs(t, err, istanbul.ErrInvalidCommittedSeals)
		assert.Empty(t, backend.commitCh)
		assert.True(t, backend.HasBadProposal(block.Hash()))
	})
	// ##
}

// ##CROSS: istanbul validation
// commitSeal returns the round 0 committed seal of block signed by key.
func commitSeal(t *testing.T, key *ecdsa.PrivateKey, block *types.Block) signedSeal {
	t.Helper()
	seal, err := crypto.Sign(istcore.PrepareCommittedSeal(block.Header(), 0), key)
	require.NoError(t, err)
	return seal
}

// ##

func TestGetProposer(t *testing.T) {
	chain, engine := newBlockChain(1)
	defer engine.Stop()

	block := makeBlock(chain, engine, chain.Genesis())
	chain.InsertChain(types.Blocks{block})
	expected := engine.GetProposer(1)
	actual := engine.Address()
	if actual != expected {
		t.Errorf("proposer mismatch: have %v, want %v", actual.Hex(), expected.Hex())
	}
}

// TestTransitionDeadlock test whether a deadlock occurs when testIstanbulBlock is set to 1
// This was fixed as part of commit 2a8310663ecafc0233758ca7883676bf568e926e
func TestTransitionDeadlock(t *testing.T) {
	timeout := time.After(1 * time.Minute)
	done := make(chan bool)
	go func() {
		chain, engine := newBlockChain(1)
		defer engine.Stop()
		// Create an insert a new block into the chain.
		block := makeBlock(chain, engine, chain.Genesis())
		_, err := chain.InsertChain(types.Blocks{block})
		if err != nil {
			t.Errorf("Error inserting block: %v", err)
		}

		if err = engine.NewChainHead(); err != nil {
			t.Errorf("Error posting NewChainHead Event: %v", err)
		}

		done <- true
	}()

	select {
	case <-timeout:
		t.Fatal("Deadlock occurred")
	case <-done:
	}
}

/**
 * SimpleBackend
 * Private key: bb047e5940b6d83354d9432db7c449ac8fca2248008aaa7271369880f9f11cc1
 * Public key: 04a2bfb0f7da9e1b9c0c64e14f87e8fb82eb0144e97c25fe3a977a921041a50976984d18257d2495e7bfd3d4b280220217f429287d25ecdf2b0d7c0f7aae9aa624
 * Address: 0x70524d664ffe731100208a0154e556f9bb679ae6
 */
func getAddress() common.Address {
	return common.HexToAddress("0x70524d664ffe731100208a0154e556f9bb679ae6")
}

func getInvalidAddress() common.Address {
	return common.HexToAddress("0x9535b2e7faaba5288511d89341d94a38063a349b")
}

func generatePrivateKey() (*ecdsa.PrivateKey, error) {
	key := "bb047e5940b6d83354d9432db7c449ac8fca2248008aaa7271369880f9f11cc1"
	return crypto.HexToECDSA(key)
}

func newTestValidatorSet(n int) (istanbul.ValidatorSet, []*ecdsa.PrivateKey) {
	// generate validators
	keys := make(Keys, n)
	addrs := make([]common.Address, n)
	for i := 0; i < n; i++ {
		privateKey, _ := crypto.GenerateKey()
		keys[i] = privateKey
		addrs[i] = crypto.PubkeyToAddress(privateKey.PublicKey)
	}
	vset := validator.NewSet(addrs, nil, istanbul.NewRoundRobinProposerPolicy())
	sort.Sort(keys) //Keys need to be sorted by its public key address
	return vset, keys
}

type Keys []*ecdsa.PrivateKey

func (slice Keys) Len() int { return len(slice) }
func (slice Keys) Less(i, j int) bool {
	return strings.Compare(crypto.PubkeyToAddress(slice[i].PublicKey).String(), crypto.PubkeyToAddress(slice[j].PublicKey).String()) < 0
}
func (slice Keys) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func newBackend() (b *Backend) {
	_, b = newBlockChain(1)
	key, _ := generatePrivateKey()
	b.privateKey = key
	// ##CROSS: bls seal
	blsKey, _ := bls.RandKey()
	b.blsSecretKey = blsKey
	// ##
	return
}

// ##CROSS: blob sidecars
func TestBackend_Verify(t *testing.T) {
	chain, engine := newBlockChain(1)
	defer engine.Stop()
	block := makeBlock(chain, engine, chain.Genesis())

	t.Run("accepts block without blob data", func(t *testing.T) {
		_, err := engine.Verify(block)
		require.NoError(t, err)
	})

	t.Run("rejects sidecars that do not match blob transactions", func(t *testing.T) {
		sidecar := &types.BlobSidecar{BlockNumber: block.Number(), BlockHash: block.Hash()}
		_, err := engine.Verify(block.WithSidecars(types.BlobSidecars{sidecar}))
		require.ErrorIs(t, err, istanbul.ErrUnavailableBlobData)
	})
}

// ##

// ##CROSS: consensus peer permissioning
func TestBackend_IsValidatorAt(t *testing.T) {
	chain, engine := newBlockChain(1)
	defer engine.Stop()
	head := chain.CurrentHeader()

	t.Run("admits a member of the validator set", func(t *testing.T) {
		ok, err := engine.IsValidatorAt(chain, head, engine.Address())
		require.NoError(t, err)
		require.True(t, ok)
	})

	t.Run("rejects an unknown address", func(t *testing.T) {
		ok, err := engine.IsValidatorAt(chain, head, common.HexToAddress("0x1234"))
		require.NoError(t, err)
		require.False(t, ok)
	})

	t.Run("returns an error when the snapshot is unavailable", func(t *testing.T) {
		unknown := &types.Header{Number: big.NewInt(5), ParentHash: common.Hash{0x01}}
		ok, err := engine.IsValidatorAt(chain, unknown, engine.Address())
		require.Error(t, err)
		require.False(t, ok)
	})
}

// ##

// ##CROSS: istanbul status limit
func TestAPI_Status(t *testing.T) {
	chain, engine := newBlockChain(1)
	defer engine.Stop()

	parent := chain.Genesis()
	for range 3 {
		block := makeBlock(chain, engine, parent)
		_, err := chain.InsertChain(types.Blocks{block})
		require.NoError(t, err)
		require.NoError(t, engine.NewChainHead()) // move the core to the next block without waiting for a round change
		parent = block
	}

	api := &API{chain: chain, backend: engine}
	number := func(n int64) *rpc.BlockNumber {
		bn := rpc.BlockNumber(n)
		return &bn
	}

	t.Run("counts block authors", func(t *testing.T) {
		status, err := api.Status(context.Background(), number(1), number(3))
		require.NoError(t, err)
		require.Equal(t, uint64(2), status.NumBlocks)
		require.Equal(t, map[common.Address]int{engine.Address(): 2}, status.SigningStatus)
	})

	t.Run("rejects a range over the limit", func(t *testing.T) {
		_, err := api.Status(context.Background(), number(0), number(maxStatusBlocks+1))
		require.ErrorContains(t, err, "range too large")
	})
}

// ##
