// Package workerpoolsize — Gopher Workplace challenge.
package workerpoolsize

// Map applies f to every item using at most workers goroutines and returns
// the results in input order. A non-positive workers runs one goroutine, and
// more workers than items is harmless. No shared state may be written without
// synchronisation — writing results into distinct slice slots is enough.
//
// Examples:
//
//	Map([]int{1, 2, 3}, 2, double) => []int{2, 4, 6}
func Map(items []int, workers int, f func(int) int) []int {
	panic("not implemented")
}

// Sizing returns a sensible worker count for a workload: cpus for CPU-bound
// work, and cpus*blocked for work that spends a blocked fraction of its time
// waiting on I/O, rounded down but never below 1. blocked is clamped into
// [0,1).
//
// Examples:
//
//	Sizing(8, 0) => 8
func Sizing(cpus int, blocked float64) int {
	panic("not implemented")
}
