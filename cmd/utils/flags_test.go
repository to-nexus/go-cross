// Copyright 2019 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// Package utils contains internal helper functions for go-ethereum commands.
package utils

import (
	"flag"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func Test_SplitTagsFlag(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		args string
		want map[string]string
	}{
		{
			"2 tags case",
			"host=localhost,bzzkey=123",
			map[string]string{
				"host":   "localhost",
				"bzzkey": "123",
			},
		},
		{
			"1 tag case",
			"host=localhost123",
			map[string]string{
				"host": "localhost123",
			},
		},
		{
			"empty case",
			"",
			map[string]string{},
		},
		{
			"garbage",
			"smth=smthelse=123",
			map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := SplitTagsFlag(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitTagsFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ##CROSS: fix upstream
func TestFlags_SetP2PConfig(t *testing.T) {
	newCtx := func(t *testing.T, args ...string) *cli.Context {
		t.Helper()
		set := flag.NewFlagSet("test", flag.ContinueOnError)
		require.NoError(t, NoDiscoverFlag.Apply(set))
		require.NoError(t, set.Parse(args))
		return cli.NewContext(cli.NewApp(), set, nil)
	}
	for _, tt := range []struct {
		args []string
		want bool
	}{
		{[]string{"--nodiscover"}, true},
		{[]string{"--nodiscover=true"}, true},
		{[]string{"--nodiscover=false"}, false},
	} {
		t.Run(tt.args[0], func(t *testing.T) {
			cfg := &p2p.Config{}
			SetP2PConfig(newCtx(t, tt.args...), cfg)
			assert.Equal(t, tt.want, cfg.NoDiscovery)
		})
	}
}

// ##
