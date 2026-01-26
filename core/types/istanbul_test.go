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

package types

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestHeaderHash(t *testing.T) {
	// 0xd848102c76ea4c0a814cd8501ee5e2f243d4d7a0c6a2bba68a4be23ca1f80965
	expectedExtra := common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000f89af8549444add0ec310f115a0e603b2d7db9f067778eaf8a94294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212946beaaed781d2d2ab6350f5c4566a2c6eaac407a6948be76812f765c24641ec63dc2852b378aba2b440b8410000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0")
	expectedHash := common.HexToHash("0xd848102c76ea4c0a814cd8501ee5e2f243d4d7a0c6a2bba68a4be23ca1f80965")

	// for istanbul consensus
	header := &Header{MixDigest: IstanbulDigest, Extra: expectedExtra}
	if !reflect.DeepEqual(header.Hash(), expectedHash) {
		t.Errorf("expected: %v, but got: %v", expectedHash.Hex(), header.Hash().Hex())
	}

	// append useless information to extra-data
	unexpectedExtra := append(expectedExtra, []byte{1, 2, 3}...)
	header.Extra = unexpectedExtra
	if !reflect.DeepEqual(header.Hash(), rlpHash(header)) {
		t.Errorf("expected: %v, but got: %v", rlpHash(header).Hex(), header.Hash().Hex())
	}
}

func TestExtractToIstanbulExtra(t *testing.T) {
	testCases := []struct {
		istRawData     []byte
		expectedResult *IstanbulExtra
		expectedErr    error
	}{
		{
			// normal case
			hexutil.MustDecode("0xf85b80f8549444add0ec310f115a0e603b2d7db9f067778eaf8a94294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212946beaaed781d2d2ab6350f5c4566a2c6eaac407a6948be76812f765c24641ec63dc2852b378aba2b440c080c080"),
			&IstanbulExtra{
				VanityData: []byte{},
				Validators: []common.Address{
					common.BytesToAddress(hexutil.MustDecode("0x44add0ec310f115a0e603b2d7db9f067778eaf8a")),
					common.BytesToAddress(hexutil.MustDecode("0x294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212")),
					common.BytesToAddress(hexutil.MustDecode("0x6beaaed781d2d2ab6350f5c4566a2c6eaac407a6")),
					common.BytesToAddress(hexutil.MustDecode("0x8be76812f765c24641ec63dc2852b378aba2b440")),
				},
				CommittedSeal: [][]byte{},
				Round:         0,
				Vote:          nil,
				RandomReveal:  []byte{},
			},
			nil,
		},
		{
			// ##CROSS: bls seal
			hexutil.MustDecode("0xf9012380f8549444add0ec310f115a0e603b2d7db9f067778eaf8a94294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212946beaaed781d2d2ab6350f5c4566a2c6eaac407a6948be76812f765c24641ec63dc2852b378aba2b440c080c080c10ff8c4b0b2280286d402014611fe75e86eb09c9d3128292d2f8a023e784e455155cdcfeb51b69a34ad5f1393a5b4685bddb831d7b08a40be1e6030ac0b0e34af8358487f8ce404bcfbd005e47621ecb06cd14ece0d9c9350a4db52b79e588c9ad9c889f322b095ef4ce1b76306d7cd92d2875239ff0b8475933a6e97464a70bb74fd07782a9f1b4ab1fbfad533105a64226b90bed5f5b0a57c8d7cfd9a068630908d1802dcff016fe8fe81e149589934e2f63c1665b6a368cb3dee3a3025384510036eedb47349"),
			&IstanbulExtra{
				VanityData: []byte{},
				Validators: []common.Address{
					common.BytesToAddress(hexutil.MustDecode("0x44add0ec310f115a0e603b2d7db9f067778eaf8a")),
					common.BytesToAddress(hexutil.MustDecode("0x294fc7e8f22b3bcdcf955dd7ff3ba2ed833f8212")),
					common.BytesToAddress(hexutil.MustDecode("0x6beaaed781d2d2ab6350f5c4566a2c6eaac407a6")),
					common.BytesToAddress(hexutil.MustDecode("0x8be76812f765c24641ec63dc2852b378aba2b440")),
				},
				CommittedSeal: [][]byte{},
				Round:         0,
				Vote:          nil,
				RandomReveal:  []byte{},
				SignersBitset: []uint64{0b1111},
				Signers: []BLSPublicKey{
					BLSPublicKey(hexutil.MustDecode("0xb2280286d402014611fe75e86eb09c9d3128292d2f8a023e784e455155cdcfeb51b69a34ad5f1393a5b4685bddb831d7")),
					BLSPublicKey(hexutil.MustDecode("0x8a40be1e6030ac0b0e34af8358487f8ce404bcfbd005e47621ecb06cd14ece0d9c9350a4db52b79e588c9ad9c889f322")),
					BLSPublicKey(hexutil.MustDecode("0x95ef4ce1b76306d7cd92d2875239ff0b8475933a6e97464a70bb74fd07782a9f1b4ab1fbfad533105a64226b90bed5f5")),
					BLSPublicKey(hexutil.MustDecode("0xa57c8d7cfd9a068630908d1802dcff016fe8fe81e149589934e2f63c1665b6a368cb3dee3a3025384510036eedb47349")),
				},
			},
			nil,
			// ##
		},
	}
	for _, test := range testCases {
		h := &Header{Extra: test.istRawData}
		istanbulExtra, err := ExtractIstanbulExtra(h)
		if err != test.expectedErr {
			t.Errorf("expected: %v, but got: %v", test.expectedErr, err)
		}
		if !reflect.DeepEqual(istanbulExtra, test.expectedResult) {
			t.Errorf("expected: %v, but got: %v", test.expectedResult, istanbulExtra)
		}
	}
}
