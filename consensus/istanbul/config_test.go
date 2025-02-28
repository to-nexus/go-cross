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

package istanbul

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/naoina/toml"
	"github.com/stretchr/testify/assert"
)

func TestProposerPolicy_UnmarshalTOML(t *testing.T) {
	input := `id = 2
`
	expectedId := ProposerPolicyId(2)
	var p proposerPolicyToml
	assert.NoError(t, toml.Unmarshal([]byte(input), &p))
	assert.Equal(t, expectedId, p.Id, "ProposerPolicyId mismatch")
}

func TestProposerPolicy_MarshalTOML(t *testing.T) {
	output := `id = 1
`
	p := &ProposerPolicy{Id: 1}
	b, err := p.MarshalTOML()
	if err != nil {
		t.Errorf("error marshalling ProposerPolicy: %v", err)
	}
	assert.Equal(t, output, b, "ProposerPolicy MarshalTOML mismatch")
}

func TestGetConfig(t *testing.T) {
	if !reflect.DeepEqual(DefaultConfig.GetConfig(nil), *DefaultConfig) {
		t.Errorf("error default config:\nexpected: %v\n", DefaultConfig)
	}

	config1 := *DefaultConfig
	config1.Epoch = 40000
	config1.BlockPeriod = 1       // default value
	config1.EmptyBlockPeriod = 10 // default value
	config3 := config1
	config3.BlockPeriod = 5
	config3.EmptyBlockPeriod = 10
	config5 := config3
	config5.EmptyBlockPeriod = 5 // equals to block period because it can be lower than this active value (5) from block 3
	config5.RequestTimeout = 15000

	type test struct {
		blockNumber    int64
		expectedConfig Config
	}
	tests := []test{
		{1, config1},
		{2, config1},
		{3, config3},
		{4, config3},
		{5, config5},
		{10, config5},
		{100, config5},
	}

	for _, test := range tests {
		c := test.expectedConfig.GetConfig(big.NewInt(test.blockNumber))
		if !reflect.DeepEqual(c, test.expectedConfig) {
			t.Errorf("error mismatch:\nexpected: %v\ngot: %v\n", test.expectedConfig, c)
		}
	}
}
