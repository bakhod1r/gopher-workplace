// Package paddedcounters — Gopher Workplace challenge.
package paddedcounters

import (
	"sync"
	"unsafe"
)

// LineSize is the coherence granule the counters must not share.
const LineSize = 64

// Slot is one worker's counter, padded to its own cache line.
type Slot struct {
	N   int64
	Pad [LineSize - unsafe.Sizeof(int64(0))]byte
}

// Run gives each of workers goroutines its own Slot, has each increment
// its counter iters times, and returns the total.
//
// Slot must be padded so no two counters share a cache line; the padding is
// computed from the counter's own size, not hard-coded.
//
// Examples:
//
//	Run(4, 1000) => 4000
func Run(workers, iters int) int64 {
	panic("not implemented")
}
