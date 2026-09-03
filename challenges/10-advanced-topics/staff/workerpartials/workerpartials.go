// Package workerpartials — Gopher Workplace challenge.
package workerpartials

import "sync"

// Histogram counts data into buckets bins by value modulo buckets, using
// workers goroutines over disjoint chunks.
//
// Workers must not share a counter: each accumulates privately and the
// results are folded once, after the join.
//
// Examples:
//
//	Histogram([]int{0, 1, 2, 3}, 2, 2) => []int64{2, 2}
func Histogram(data []int, buckets, workers int) []int64 {
	panic("not implemented")
}
