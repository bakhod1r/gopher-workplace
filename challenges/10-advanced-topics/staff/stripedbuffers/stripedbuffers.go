// Package stripedbuffers — Gopher Workplace challenge.
package stripedbuffers

import (
	"sync"
	"unsafe"
)

// lineSize is the coherence granule the shards must not share.
const lineSize = 64

// stripe is one shard: a lock, a scratch buffer, and padding to a line.
type stripe struct {
	mu  sync.Mutex
	buf []byte
	_   [lineSize - unsafe.Sizeof(sync.Mutex{}) - unsafe.Sizeof([]byte(nil))]byte
}

// Striped hands out per-shard scratch buffers.
type Striped struct {
	stripes []stripe
}

// NewStriped returns a Striped with n shards, each holding a size-byte buffer.
func NewStriped(n, size int) *Striped {
	if n < 1 {
		n = 1
	}
	s := &Striped{stripes: make([]stripe, n)}
	for i := range s.stripes {
		s.stripes[i].buf = make([]byte, 0, size)
	}
	return s
}

// With runs fn on the shard's scratch buffer and returns the number of
// bytes fn left in it.
//
// Each shard has its own buffer and its own lock, padded so neighbouring
// shards do not share a cache line.
//
// Examples:
//
//	s.With(0, func(b []byte) []byte { return append(b, 'x') }) => 1
func (s *Striped) With(id int, fn func(buf []byte) []byte) int {
	panic("not implemented")
}
