// Package sharedview — Gopher Workplace challenge.
package sharedview

import (
	"sync"
	"unsafe"
)

// SumParallel reinterprets b as int64 values and totals them using
// workers goroutines over disjoint chunks.
//
// The view shares b's storage, so nothing may write through it. Concurrent
// reads of shared memory need no synchronisation at all.
//
// Examples:
//
//	SumParallel(sixteenBytes, 2) => the two int64s summed, true
func SumParallel(b []byte, workers int) (int64, bool) {
	panic("not implemented")
}
