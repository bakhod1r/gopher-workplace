// Package falsesharing — Gopher Workplace challenge.
package falsesharing

import "sync/atomic"

// CacheLine is the padding target in bytes.
const CacheLine = 64

// Counters is a set of per-worker counters.
type Counters interface {
	Inc(i int)
	Total() int64
}

// paddedCell is one counter padded to a full cache line.
type paddedCell struct {
	v atomic.Int64
	_ [CacheLine - 8]byte
}

// PaddedCounters gives every counter its own cache line.
type PaddedCounters struct {
	cells []paddedCell
}

// NewPadded returns n padded counters.
func NewPadded(n int) *PaddedCounters {
	return &PaddedCounters{cells: make([]paddedCell, n)}
}

// Inc increments counter i. Out-of-range indexes are ignored.
func (p *PaddedCounters) Inc(i int) {
	// TODO(candidate): bounds-check, then atomically increment.
	panic("not implemented")
}

// Total sums every counter.
func (p *PaddedCounters) Total() int64 {
	// TODO(candidate): sum the cells.
	panic("not implemented")
}

// PackedCounters stores counters adjacently — they share cache lines.
type PackedCounters struct {
	vals []atomic.Int64
}

// NewPacked returns n packed counters.
func NewPacked(n int) *PackedCounters {
	return &PackedCounters{vals: make([]atomic.Int64, n)}
}

// Inc increments counter i. Out-of-range indexes are ignored.
func (p *PackedCounters) Inc(i int) {
	// TODO(candidate): bounds-check, then atomically increment.
	panic("not implemented")
}

// Total sums every counter.
func (p *PackedCounters) Total() int64 {
	// TODO(candidate): sum the values.
	panic("not implemented")
}
