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

package validator

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/stretchr/testify/assert"
)

// ##CROSS: UPSTREAM PR-1687
var (
	addr1    = common.HexToAddress("0xc53f2189bf6d7bf56722731787127f90d319e112")
	addr2    = common.HexToAddress("0xed2d479591fe2c5626ce09bca4ed2a62e00e5bc2")
	addr3    = common.HexToAddress("0xc8417f834995aaeb35f342a67a4961e19cd4735c")
	addr4    = common.HexToAddress("0x784ae51f5013b51c8360afdf91c6bc5a16f586ea")
	addr5    = common.HexToAddress("0xecf0974e6f0630fd91ea4da8399cdb3f59e5220f")
	addr6    = common.HexToAddress("0x411c4d11acd714b82a5242667e36de14b9e1d10b")
	addr7    = common.HexToAddress("0x681381b3D0DaaC179d95aCc9e22E23da2DA670f6")
	addrSet  = []common.Address{addr1, addr2, addr3, addr4, addr5, addr6}
	addrSet2 = []common.Address{addr7, addr1, addr2, addr3, addr4, addr5}
)

// ##

func TestProposerPolicy(t *testing.T) {
	addressSortedByByte := []common.Address{addr6, addr4, addr1, addr3, addr5, addr2}
	addressSortedByString := []common.Address{addr6, addr4, addr1, addr2, addr5, addr3}

	pp := istanbul.NewRoundRobinProposerPolicy()
	pp.Use(istanbul.ValidatorSortByByte())

	valSet := NewSet(addrSet, pp)
	valList := valSet.List()

	for i := 0; i < 6; i++ {
		assert.Equal(t, addressSortedByByte[i].Hex(), valList[i].String(), "validatorSet not byte sorted")
	}

	pp.Use(istanbul.ValidatorSortByString())
	for i := 0; i < 6; i++ {
		assert.Equal(t, addressSortedByString[i].Hex(), valList[i].String(), "validatorSet not string sorted")
	}
}

// ##CROSS: UPSTREAM PR-1687
// TestProposerPolicyRegistration verifies that the validator set registration does not exceed the maximum limit defined by istanbul.MaxValidatorSetInRegistry.
// The test procedure is as follows:
//  1. Two RoundRobinProposerPolicy instances (pp and pp2) are created along with two different validator sets (valSet and valSet2).
//  2. The test repeatedly registers valSet with pp for (MaxValidatorSetInRegistry + 100) times.
//  3. Then, valSet2 is registered.
//  4. Finally, the test asserts that the total number of registered validator sets is exactly equal to MaxValidatorSetInRegistry,
//     ensuring that any validator sets beyond this limit are dropped.
func TestProposerPolicyRegistration(t *testing.T) {
	// test that registration can't go beyond MaxValidatorSetInRegistry limit
	pp := istanbul.NewRoundRobinProposerPolicy()
	pp2 := istanbul.NewRoundRobinProposerPolicy()
	valSet := NewSet(addrSet, pp)
	valSet2 := NewSet(addrSet2, pp2)

	for i := 0; i < istanbul.MaxValidatorSetInRegistry+100; i++ {
		pp.RegisterValidatorSet(valSet)
	}
	pp.RegisterValidatorSet(valSet2)
	assert.Equal(t, istanbul.MaxValidatorSetInRegistry, pp.GetRegistrySize(), "validator set not dropped")
}

// ##
