// Package opsperiter — Gopher Workplace challenge.
package opsperiter

// OpsPerIter reports how many logical operations one benchmark iteration
// performed, rounding half away from zero: a body that processes a batch of
// records still wants a per-record number. A non-positive iters reports 0.
//
// Examples:
//
//	OpsPerIter(10, 4) => 3
func OpsPerIter(totalOps int, iters int) int {
	panic("not implemented")
}
