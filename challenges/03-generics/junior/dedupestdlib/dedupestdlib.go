// Package dedupestdlib — Gopher Workplace challenge.
package dedupestdlib

import (
	"cmp"
)

// Distinct returns the distinct elements of s in ascending order.
// The input is not modified.
func Distinct[T cmp.Ordered](s []T) []T {
	// TODO(candidate): clone, sort, then collapse the runs.
	panic("not implemented")
}
