package scwallet

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/require"
)

// ##CROSS: fix upstream 46d8047 (#33606)
func TestSecureChannelSession_DecryptAPDU(t *testing.T) {
	s := &SecureChannelSession{sessionEncKey: make([]byte, 32), iv: make([]byte, aes.BlockSize)}
	for _, n := range []int{0, 1, 15, 17} {
		_, err := s.decryptAPDU(make([]byte, n))
		require.ErrorContains(t, err, "invalid ciphertext length", "length %d", n)
	}
}

// ##
