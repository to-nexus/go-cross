package blst

import (
	"crypto/rand"
	"encoding/binary"
	mrand "math/rand"
	"sync"
)

type source struct{}

var lock sync.RWMutex
var _ mrand.Source64 = (*source)(nil)

// Seed does nothing when crypto/rand is used as source.
func (_ *source) Seed(_ int64) {}

// Int63 returns uniformly-distributed random (as in CSPRNG) int64 value within [0, 1<<63) range.
// Panics if random generator reader cannot return data.
func (s *source) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

// Uint64 returns uniformly-distributed random (as in CSPRNG) uint64 value within [0, 1<<64) range.
// Panics if random generator reader cannot return data.
func (_ *source) Uint64() (val uint64) {
	lock.RLock()
	defer lock.RUnlock()
	if err := binary.Read(rand.Reader, binary.BigEndian, &val); err != nil {
		panic(err) // lint:nopanic -- Panic risk is communicated in the godoc commentary.
	}
	return
}

func newGenerator() *mrand.Rand {
	return mrand.New(&source{})
}
