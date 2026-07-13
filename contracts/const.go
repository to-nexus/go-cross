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

package contracts

import "github.com/ethereum/go-ethereum/common"

// ##CROSS: contract upgrade

var (
	// predeploy
	BridgeImplSlot = common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")

	CrossExAddr    = common.HexToAddress("0xfeed00000000000000000000000000000000C0DE")
	BridgeAddr     = common.HexToAddress("0xb81d6e000000000000000000000000000000C0de")
	BridgeImplAddr = common.HexToAddress("0xB81D6e000000000000000000000000000000AAaA")
	Multicall3Addr = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")

	// breakpoint
	// ##CROSS: consensus system contract
	ValidatorSetAddr       = common.HexToAddress("0x0000000000000000000000000000000000001001")
	StakeHubAddr           = common.HexToAddress("0x0000000000000000000000000000000000001002")
	RewardHubAddr          = common.HexToAddress("0x0000000000000000000000000000000000001003")
	ValidatorSlashAddr     = common.HexToAddress("0x0000000000000000000000000000000000001004")
	ValidatorSetImplAddr   = common.HexToAddress("0x000000000000000000000000000000000000A001")
	StakeHubImplAddr       = common.HexToAddress("0x000000000000000000000000000000000000A002")
	RewardHubImplAddr      = common.HexToAddress("0x000000000000000000000000000000000000A003")
	ValidatorSlashImplAddr = common.HexToAddress("0x000000000000000000000000000000000000A004")
	// ##
)

var systemContracts = map[common.Address]bool{
	ValidatorSetAddr:   true,
	StakeHubAddr:       true,
	RewardHubAddr:      true,
	ValidatorSlashAddr: true,
}

// IsSystemContract checks if the address is a system contract.
func IsSystemContract(addr common.Address) bool {
	return systemContracts[addr]
}
