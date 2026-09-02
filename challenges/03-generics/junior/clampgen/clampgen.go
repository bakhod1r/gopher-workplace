// Package clampgen — Gopher Workplace challenge.
package clampgen

import (
	"cmp"
)

// Clamp returns v limited to the range [lo, hi].
//
// Examples:
//
//	Clamp(5, 0, 3) => 3
//	Clamp(-1, 0, 3) => 0
func Clamp[T cmp.Ordered](v, lo, hi T) T {
	// TODO(candidate): pull v inside the range.
	panic("not implemented")
}
