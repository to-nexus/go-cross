package locals

import (
	"errors"
	"io"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// ##CROSS: fix upstream
type failingCloser struct{ io.Writer }

func (failingCloser) Close() error { return errors.New("close failed") }

func TestJournal_Rotate(t *testing.T) {
	t.Run("recovers after a close error", func(t *testing.T) {
		j := newTxJournal(filepath.Join(t.TempDir(), "transactions.rlp"))
		j.writer = failingCloser{io.Discard}

		require.Error(t, j.rotate(nil))
		require.Nil(t, j.writer, "a closed writer must not be kept")

		require.NoError(t, j.rotate(nil))
		require.NotNil(t, j.writer)
		require.NoError(t, j.close())
	})
}

// ##
