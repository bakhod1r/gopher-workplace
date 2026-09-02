// Package atomicbitset — Gopher Workplace challenge.
package atomicbitset

import (
	"math/bits"
	"sync/atomic"
)

// Set is a concurrent bitset.
type Set interface {
	Set(i int) bool
	Clear(i int) bool
	Test(i int) bool
	Count() int
}

// Bitset stores bits in atomically updated 64-bit words.
type Bitset struct {
	words []atomic.Uint64
	n     int
}

// NewBitset returns a bitset holding n bits.
func NewBitset(n int) *Bitset {
	if n < 0 {
		n = 0
	}
	return &Bitset{words: make([]atomic.Uint64, (n+63)/64), n: n}
}

// Set turns bit i on and reports whether it changed.
//
// Examples:
//
//	Set(5) => true;  Set(5) again => false
func (b *Bitset) Set(i int) bool {
	// TODO(candidate): CAS loop ORing in the mask.
	panic("not implemented")
}

// Clear turns bit i off and reports whether it changed.
func (b *Bitset) Clear(i int) bool {
	// TODO(candidate): CAS loop clearing the mask.
	panic("not implemented")
}

// Test reports whether bit i is set.
func (b *Bitset) Test(i int) bool {
	// TODO(candidate): atomic load and mask.
	panic("not implemented")
}

// Count returns how many bits are set.
func (b *Bitset) Count() int {
	// TODO(candidate): popcount every word.
	panic("not implemented")
}

var _ = bits.OnesCount64
