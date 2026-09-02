// Package minmaxbuiltin — Gopher Workplace challenge.
package minmaxbuiltin

import (
	"cmp"
)

// Middle returns v clamped to the range [lo, hi] using the builtins.
func Middle[T cmp.Ordered](v, lo, hi T) T {
	// TODO(candidate): combine the min and max builtins.
	panic("not implemented")
}

// Spread returns the difference between the largest and smallest
// of the three values.
func Spread(a, b, c int) int {
	// TODO(candidate): use the builtins on three arguments.
	panic("not implemented")
}
