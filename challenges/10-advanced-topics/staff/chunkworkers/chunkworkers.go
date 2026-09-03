// Package chunkworkers — Gopher Workplace challenge.
package chunkworkers

import "sync"

// SumParallel sums s using workers goroutines over disjoint chunks of
// the input and returns the total.
//
// The input must not be copied: each worker gets a view. Parallelism has to
// be real — no locking on a shared accumulator per element.
//
// Examples:
//
//	SumParallel([]int{1, 2, 3, 4}, 2) => 10
func SumParallel(s []int, workers int) int64 {
	panic("not implemented")
}
