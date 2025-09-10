// Copyright 2021 The go-ethereum Authors
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

package eip1559

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// VerifyEIP1559Header verifies some header attributes which were changed in EIP-1559,
// - gas limit check
// - basefee check
func VerifyEIP1559Header(config *params.ChainConfig, parent, header *types.Header) error {
	if types.IsIstanbulDigest(parent.MixDigest) { // ##CROSS: istanbul
		// Verify that the gas limit equals to the value from the last transition or the parent header
		gasLimit := config.GetGasLimit(header.Number)
		if gasLimit == nil {
			gasLimit = &parent.GasLimit
		}
		if *gasLimit != header.GasLimit {
			return fmt.Errorf("invalid gasLimit: have %d, want %d", header.GasLimit, *gasLimit)
		}
	} else {
		// Verify that the gas limit remains within allowed bounds
		parentGasLimit := parent.GasLimit
		if !config.IsLondon(parent.Number) {
			parentGasLimit = parent.GasLimit * config.ElasticityMultiplier()
		}
		if err := misc.VerifyGaslimit(parentGasLimit, header.GasLimit); err != nil {
			return err
		}
	}
	// Verify the header is not malformed
	if header.BaseFee == nil {
		return errors.New("header is missing baseFee")
	}
	// Verify the baseFee is correct based on the parent header.
	expectedBaseFee := CalcBaseFee(config, parent)
	if header.BaseFee.Cmp(expectedBaseFee) != 0 {
		return fmt.Errorf("invalid baseFee: have %s, want %s, parentBaseFee %s, parentGasUsed %d",
			header.BaseFee, expectedBaseFee, parent.BaseFee, parent.GasUsed)
	}
	return nil
}

// CalcBaseFee calculates the basefee of the header.
func CalcBaseFee(config *params.ChainConfig, parent *types.Header) *big.Int {
	if config.Clique != nil { // ##CROSS: for test
		return new(big.Int).Set(parent.BaseFee)
	}

	// If the current block is the first EIP-1559 block, return the InitialBaseFee.
	if !config.IsLondon(parent.Number) {
		return new(big.Int).SetUint64(params.InitialBaseFee)
	}

	// ##CROSS: basefee
	var elasticityMultiplier uint64
	if types.IsIstanbulDigest(parent.MixDigest) {
		elasticityMultiplier = config.GetElasticityMultiplier(parent.Number)
	} else {
		elasticityMultiplier = config.ElasticityMultiplier()
	}
	parentGasTarget := parent.GasLimit / elasticityMultiplier

	// If the parent gasUsed is the same as the target, the baseFee remains unchanged.
	if parent.GasUsed == parentGasTarget {
		return new(big.Int).Set(parent.BaseFee)
	}

	var (
		num   = new(big.Int)
		denom = new(big.Int)
	)

	baseFeeChangeDenominator := func() uint64 {
		if types.IsIstanbulDigest(parent.MixDigest) {
			if denominator := config.GetBaseFeeChangeDenominator(parent.Number); denominator > 0 {
				return denominator
			}
		}
		return config.BaseFeeChangeDenominator()
	}()

	if parent.GasUsed > parentGasTarget {
		// If the parent block used more gas than its target, the baseFee should increase.
		// max(1, parentBaseFee * gasUsedDelta / parentGasTarget / baseFeeChangeDenominator)
		num.SetUint64(parent.GasUsed - parentGasTarget)
		num.Mul(num, parent.BaseFee)
		num.Div(num, denom.SetUint64(parentGasTarget))
		num.Div(num, denom.SetUint64(baseFeeChangeDenominator))
		if num.Cmp(common.Big1) < 0 {
			num.Add(parent.BaseFee, common.Big1)
		} else {
			num.Add(parent.BaseFee, num)
		}

		if types.IsIstanbulDigest(parent.MixDigest) {
			if max := config.GetMaxBaseFee(parent.Number); max != nil && num.Cmp(max) > 0 {
				num.Set(max)
			}
		}
		return num
	} else {
		// Otherwise if the parent block used less gas than its target, the baseFee should decrease.
		// max(0, parentBaseFee * gasUsedDelta / parentGasTarget / baseFeeChangeDenominator)
		num.SetUint64(parentGasTarget - parent.GasUsed)
		num.Mul(num, parent.BaseFee)
		num.Div(num, denom.SetUint64(parentGasTarget))
		num.Div(num, denom.SetUint64(baseFeeChangeDenominator))

		baseFee := num.Sub(parent.BaseFee, num)
		if baseFee.Sign() < 0 {
			baseFee = common.Big0
		}

		if types.IsIstanbulDigest(parent.MixDigest) {
			if min := config.GetMinBaseFee(parent.Number); min != nil && baseFee.Cmp(min) < 0 {
				baseFee.Set(min)
			}
		}
		return baseFee
	}
	// ##
}
