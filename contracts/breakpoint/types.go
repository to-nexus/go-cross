// Copyright 2025 The go-ethereum Authors
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

// Package breakpoint contains the contract runtime bytecode deployed at the ONE mainnet Breakpoint fork.
// Bytecode is frozen at go-cross commit 1646df77cdfcf32e5f5c828560672f1384f590d2; later upgrades need a separate package.
package breakpoint

import _ "embed"

// ##CROSS: contract upgrade

var (
	//go:embed cross_ex
	CrossExCode string
)
