// Package bls provides the BLS12-381 keys and signatures used by CROSS.
//
// Other packages import this package instead of prysm, so the BLS backend is chosen in one place.
// It uses the prysm blst backend directly, because the prysm crypto/bls package also links herumi, which CROSS does not use.
package bls

import (
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/blst"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/common"
)

// SecretKey is a BLS secret key.
type SecretKey = common.SecretKey

// PublicKey is a BLS public key.
type PublicKey = common.PublicKey

// Signature is a BLS signature.
type Signature = common.Signature

// RandKey creates a new random secret key.
func RandKey() (SecretKey, error) {
	return blst.RandKey()
}

// SecretKeyFromBytes creates a secret key from a big-endian byte slice.
func SecretKeyFromBytes(privKey []byte) (SecretKey, error) {
	return blst.SecretKeyFromBytes(privKey)
}

// PublicKeyFromBytes creates a public key from its compressed bytes.
func PublicKeyFromBytes(pubKey []byte) (PublicKey, error) {
	return blst.PublicKeyFromBytes(pubKey)
}

// SignatureFromBytes creates a signature from its compressed bytes.
func SignatureFromBytes(sig []byte) (Signature, error) {
	return blst.SignatureFromBytes(sig)
}

// AggregateCompressedSignatures aggregates compressed signatures into one signature.
func AggregateCompressedSignatures(sigs [][]byte) (Signature, error) {
	return blst.AggregateCompressedSignatures(sigs)
}
