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
	"bytes"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core/types"
)

func New(addr common.Address, signerAddr types.BLSPublicKey) istanbul.Validator {
	return &blsValidator{
		address:    addr,
		signerAddr: signerAddr, // ##CROSS: bls seal
	}
}

func NewSet(addrs []common.Address, signerAddrs []types.BLSPublicKey, policy *istanbul.ProposerPolicy) istanbul.ValidatorSet { // ##CROSS: bls seal
	return newDefaultSet(addrs, signerAddrs, policy)
}

// Deprecated: use types.ExtractIstanbulExtra instead
func ExtractValidators(extraData []byte) []common.Address {
	// get the validator addresses
	addrs := make([]common.Address, (len(extraData) / common.AddressLength))
	for i := 0; i < len(addrs); i++ {
		copy(addrs[i][:], extraData[i*common.AddressLength:])
	}

	return addrs
}

// Check whether the extraData is presented in prescribed form
// Deprecated: use types.ExtractIstanbulExtra instead
func ValidExtraData(extraData []byte) bool {
	return len(extraData)%common.AddressLength == 0
}

func SortedAddresses(validators []istanbul.Validator) []common.Address {
	addrs := make([]common.Address, len(validators))
	for i, validator := range validators {
		addrs[i] = validator.Address()
	}

	slices.SortFunc(addrs, func(a, b common.Address) int {
		return bytes.Compare(a[:], b[:])
	})

	return addrs
}
