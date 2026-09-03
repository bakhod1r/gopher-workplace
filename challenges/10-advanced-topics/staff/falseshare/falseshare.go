// Package falseshare — Gopher Workplace challenge.
package falseshare

import "sync"

// cacheLine is the coherence granule the counters must not share.
const cacheLine = 64

// counter is one worker's slot.
type counter struct {
	n   int64
	pad [cacheLine - 8]byte
}

// Count runs workers goroutines, each incrementing its own counter iters
// times, and returns the total.
//
// Each worker's counter must sit on its own cache line: adjacent counters
// put the cores into a write-invalidate storm over one line.
//
// Examples:
//
//	Count(4, 1000) => 4000
func Count(workers, iters int) int64 {
	panic("not implemented")
}
