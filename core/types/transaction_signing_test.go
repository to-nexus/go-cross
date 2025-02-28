// Copyright 2016 The go-ethereum Authors
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
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
)

// TestFeeDelegationSigning verifies the fee-delegation (gas fee sponsorship) mechanism in EIP-1559 transactions.
// It demonstrates how an original DynamicFeeTx can be wrapped into a FeeDelegatedDynamicFeeTx,
// allowing a separate payer address to cover gas fees without changing the original sender.
func TestFeeDelegationSigning(t *testing.T) {
	var (
		senderKey    *ecdsa.PrivateKey
		sender       common.Address
		payerKey     *ecdsa.PrivateKey
		payer        common.Address
		londonSigner Signer // Replace with your actual type or interface
		feepaySigner Signer // Replace with your actual type or interface
		senderTx     *Transaction
		feepayTx     *Transaction
	)

	// generate key pairs
	var err error
	senderKey, err = crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	sender = crypto.PubkeyToAddress(senderKey.PublicKey)
	t.Log("Sender Address:", sender)

	payerKey, err = crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	payer = crypto.PubkeyToAddress(payerKey.PublicKey)
	t.Log("Payer Address:", payer)

	// initialize signers (e.g., chain ID 18)
	londonSigner = NewLondonSigner(big.NewInt(18))
	feepaySigner = NewFeeDelegationSigner(big.NewInt(18))

	// Create and sign a DynamicFeeTx by the sender
	t.Run("Sign sender's EIP-1559 transaction", func(t *testing.T) {
		var err error
		senderTx, err = SignTx(NewTx(&DynamicFeeTx{
			Nonce:     10,
			GasTipCap: big.NewInt(10),
			GasFeeCap: big.NewInt(20),
			Gas:       20,
			To:        nil,
			Value:     big.NewInt(30),
			Data:      []byte{},
		}), londonSigner, senderKey)
		if err != nil {
			t.Fatal(err)
		}
	})

	// Wrap the sender's transaction into a FeeDelegatedDynamicFeeTx, set fee payer, and sign
	t.Run("Sign fee-delegated transaction", func(t *testing.T) {
		var err error
		feepayTx, err = SignTx(NewTx(func() *FeeDelegatedDynamicFeeTx {
			data := &FeeDelegatedDynamicFeeTx{}
			data.FeePayer = &payer
			// Copy the original senderTx's inner DynamicFeeTx
			data.SetSenderTx(*senderTx.inner.(*DynamicFeeTx))
			return data
		}()), feepaySigner, payerKey)
		if err != nil {
			t.Fatal(err)
		}
	})

	// Verify that both the original and fee-delegated transactions share the same hash
	t.Run("Verify transaction hashes match", func(t *testing.T) {
		if senderTx.Hash() != feepayTx.Hash() {
			t.Fatal("Mismatch in hash:",
				"senderTx", senderTx.Hash(),
				"feepayTx", feepayTx.Hash())
		}
	})

	// Check the sender address of the original transaction using the LondonSigner
	t.Run("Check sender of senderTx with LondonSigner", func(t *testing.T) {
		got, err := londonSigner.Sender(senderTx)
		if err != nil {
			t.Fatal(err)
		}
		if got != sender {
			t.Fatalf("Mismatch sender: got %v, want %v", got, sender)
		}
	})

	// Check the fee payer of the fee-delegated transaction using FeeDelegationSigner
	t.Run("Check payer of feepayTx with FeeDelegationSigner", func(t *testing.T) {
		got, err := feepaySigner.Sender(feepayTx)
		if err != nil {
			t.Fatal(err)
		}
		if got != payer {
			t.Fatalf("Mismatch payer: got %v, want %v", got, payer)
		}
	})

	// Compare LondonSigner.Sender(feepayTx) with FeePayer() for feepayTx
	t.Run("Compare LondonSigner.Sender(feepayTx) with FeePayer()", func(t *testing.T) {
		expectedSender, err := londonSigner.Sender(feepayTx)
		if err != nil {
			t.Fatal(err)
		}
		expectedPayer, err := FeePayer(londonSigner, feepayTx)
		if err != nil {
			t.Fatal(err)
		}

		if expectedSender != expectedPayer {
			t.Fatal("Mismatch in payer between londonSigner.Sender and FeePayer:",
				"expectedSender", expectedSender,
				"expectedPayer", expectedPayer)
		}

		// The actual fee payer should not match the original sender
		if expectedPayer != sender {
			t.Fatal("Mismatch payer:", "expectedPayer", expectedPayer, "want", payer)
		}
	})

	// Verify the sender of the original senderTx with the FeeDelegationSigner
	t.Run("Check senderTx with FeeDelegationSigner", func(t *testing.T) {
		got, err := feepaySigner.Sender(senderTx)
		if err != nil {
			t.Fatal(err)
		}
		if got != sender {
			t.Fatalf("Mismatch sender with feepaySigner: got %v, want %v", got, sender)
		}
	})

	// Verify the sender of the fee-delegated transaction using the LondonSigner
	t.Run("Check feepayTx with LondonSigner", func(t *testing.T) {
		got, err := londonSigner.Sender(feepayTx)
		if err != nil {
			t.Fatal(err)
		}
		if got != sender {
			t.Fatalf("Mismatch sender: got %v, want %v", got, sender)
		}
	})
}

func TestEIP155Signing(t *testing.T) {
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)

	signer := NewEIP155Signer(big.NewInt(18))
	tx, err := SignTx(NewTransaction(0, addr, new(big.Int), 0, new(big.Int), nil), signer, key)
	if err != nil {
		t.Fatal(err)
	}

	from, err := Sender(signer, tx)
	if err != nil {
		t.Fatal(err)
	}
	if from != addr {
		t.Errorf("expected from and address to be equal. Got %x want %x", from, addr)
	}
}

func TestEIP155ChainId(t *testing.T) {
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)

	signer := NewEIP155Signer(big.NewInt(18))
	tx, err := SignTx(NewTransaction(0, addr, new(big.Int), 0, new(big.Int), nil), signer, key)
	if err != nil {
		t.Fatal(err)
	}
	if !tx.Protected() {
		t.Fatal("expected tx to be protected")
	}

	if tx.ChainId().Cmp(signer.chainId) != 0 {
		t.Error("expected chainId to be", signer.chainId, "got", tx.ChainId())
	}

	tx = NewTransaction(0, addr, new(big.Int), 0, new(big.Int), nil)
	tx, err = SignTx(tx, HomesteadSigner{}, key)
	if err != nil {
		t.Fatal(err)
	}

	if tx.Protected() {
		t.Error("didn't expect tx to be protected")
	}

	if tx.ChainId().Sign() != 0 {
		t.Error("expected chain id to be 0 got", tx.ChainId())
	}
}

func TestEIP155SigningVitalik(t *testing.T) {
	// Test vectors come from http://vitalik.ca/files/eip155_testvec.txt
	for i, test := range []struct {
		txRlp, addr string
	}{
		{"f864808504a817c800825208943535353535353535353535353535353535353535808025a0044852b2a670ade5407e78fb2863c51de9fcb96542a07186fe3aeda6bb8a116da0044852b2a670ade5407e78fb2863c51de9fcb96542a07186fe3aeda6bb8a116d", "0xf0f6f18bca1b28cd68e4357452947e021241e9ce"},
		{"f864018504a817c80182a410943535353535353535353535353535353535353535018025a0489efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bcaa0489efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6", "0x23ef145a395ea3fa3deb533b8a9e1b4c6c25d112"},
		{"f864028504a817c80282f618943535353535353535353535353535353535353535088025a02d7c5bef027816a800da1736444fb58a807ef4c9603b7848673f7e3a68eb14a5a02d7c5bef027816a800da1736444fb58a807ef4c9603b7848673f7e3a68eb14a5", "0x2e485e0c23b4c3c542628a5f672eeab0ad4888be"},
		{"f865038504a817c803830148209435353535353535353535353535353535353535351b8025a02a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4e0a02a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de", "0x82a88539669a3fd524d669e858935de5e5410cf0"},
		{"f865048504a817c80483019a28943535353535353535353535353535353535353535408025a013600b294191fc92924bb3ce4b969c1e7e2bab8f4c93c3fc6d0a51733df3c063a013600b294191fc92924bb3ce4b969c1e7e2bab8f4c93c3fc6d0a51733df3c060", "0xf9358f2538fd5ccfeb848b64a96b743fcc930554"},
		{"f865058504a817c8058301ec309435353535353535353535353535353535353535357d8025a04eebf77a833b30520287ddd9478ff51abbdffa30aa90a8d655dba0e8a79ce0c1a04eebf77a833b30520287ddd9478ff51abbdffa30aa90a8d655dba0e8a79ce0c1", "0xa8f7aba377317440bc5b26198a363ad22af1f3a4"},
		{"f866068504a817c80683023e3894353535353535353535353535353535353535353581d88025a06455bf8ea6e7463a1046a0b52804526e119b4bf5136279614e0b1e8e296a4e2fa06455bf8ea6e7463a1046a0b52804526e119b4bf5136279614e0b1e8e296a4e2d", "0xf1f571dc362a0e5b2696b8e775f8491d3e50de35"},
		{"f867078504a817c807830290409435353535353535353535353535353535353535358201578025a052f1a9b320cab38e5da8a8f97989383aab0a49165fc91c737310e4f7e9821021a052f1a9b320cab38e5da8a8f97989383aab0a49165fc91c737310e4f7e9821021", "0xd37922162ab7cea97c97a87551ed02c9a38b7332"},
		{"f867088504a817c8088302e2489435353535353535353535353535353535353535358202008025a064b1702d9298fee62dfeccc57d322a463ad55ca201256d01f62b45b2e1c21c12a064b1702d9298fee62dfeccc57d322a463ad55ca201256d01f62b45b2e1c21c10", "0x9bddad43f934d313c2b79ca28a432dd2b7281029"},
		{"f867098504a817c809830334509435353535353535353535353535353535353535358202d98025a052f8f61201b2b11a78d6e866abc9c3db2ae8631fa656bfe5cb53668255367afba052f8f61201b2b11a78d6e866abc9c3db2ae8631fa656bfe5cb53668255367afb", "0x3c24d7329e92f84f08556ceb6df1cdb0104ca49f"},
	} {
		signer := NewEIP155Signer(big.NewInt(1))

		var tx *Transaction
		err := rlp.DecodeBytes(common.Hex2Bytes(test.txRlp), &tx)
		if err != nil {
			t.Errorf("%d: %v", i, err)
			continue
		}

		from, err := Sender(signer, tx)
		if err != nil {
			t.Errorf("%d: %v", i, err)
			continue
		}

		addr := common.HexToAddress(test.addr)
		if from != addr {
			t.Errorf("%d: expected %x got %x", i, addr, from)
		}
	}
}

func TestChainId(t *testing.T) {
	key, _ := defaultTestKey()

	tx := NewTransaction(0, common.Address{}, new(big.Int), 0, new(big.Int), nil)

	var err error
	tx, err = SignTx(tx, NewEIP155Signer(big.NewInt(1)), key)
	if err != nil {
		t.Fatal(err)
	}

	_, err = Sender(NewEIP155Signer(big.NewInt(2)), tx)
	if !errors.Is(err, ErrInvalidChainId) {
		t.Error("expected error:", ErrInvalidChainId, err)
	}

	_, err = Sender(NewEIP155Signer(big.NewInt(1)), tx)
	if err != nil {
		t.Error("expected no error")
	}
}

type nilSigner struct {
	v, r, s *big.Int
	Signer
}

func (ns *nilSigner) SignatureValues(tx *Transaction, sig []byte) (r, s, v *big.Int, err error) {
	return ns.v, ns.r, ns.s, nil
}

// TestNilSigner ensures a faulty Signer implementation does not result in nil signature values or panics.
func TestNilSigner(t *testing.T) {
	key, _ := crypto.GenerateKey()
	innerSigner := LatestSignerForChainID(big.NewInt(1))
	for i, signer := range []Signer{
		&nilSigner{v: nil, r: nil, s: nil, Signer: innerSigner},
		&nilSigner{v: big.NewInt(1), r: big.NewInt(1), s: nil, Signer: innerSigner},
		&nilSigner{v: big.NewInt(1), r: nil, s: big.NewInt(1), Signer: innerSigner},
		&nilSigner{v: nil, r: big.NewInt(1), s: big.NewInt(1), Signer: innerSigner},
	} {
		t.Run(fmt.Sprintf("signer_%d", i), func(t *testing.T) {
			t.Run("legacy", func(t *testing.T) {
				legacyTx := createTestLegacyTxInner()
				_, err := SignNewTx(key, signer, legacyTx)
				if !errors.Is(err, ErrInvalidSig) {
					t.Fatal("expected signature values error, no nil result or panic")
				}
			})
			// test Blob tx specifically, since the signature value types changed
			t.Run("blobtx", func(t *testing.T) {
				blobtx := createEmptyBlobTxInner(false)
				_, err := SignNewTx(key, signer, blobtx)
				if !errors.Is(err, ErrInvalidSig) {
					t.Fatal("expected signature values error, no nil result or panic")
				}
			})
		})
	}
}

func createTestLegacyTxInner() *LegacyTx {
	return &LegacyTx{
		Nonce:    uint64(0),
		To:       nil,
		Value:    big.NewInt(0),
		Gas:      params.TxGas,
		GasPrice: big.NewInt(params.GWei),
		Data:     nil,
	}
}
