package common

const (
	SecretKeyLength = 32
	PublicKeyLength = 48
	SignatureLength = 96
)

// ZeroSecretKey represents a zero secret key.
var ZeroSecretKey = [SecretKeyLength]byte{}

// InfinitePublicKey represents an infinite public key (G1 Point at Infinity).
var InfinitePublicKey = [PublicKeyLength]byte{0xC0}

// InfiniteSignature represents an infinite signature (G2 Point at Infinity).
var InfiniteSignature = [SignatureLength]byte{0xC0}
