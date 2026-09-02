// Package asserteqgen — Gopher Workplace challenge.
package asserteqgen

import (
	"testing"
)

// EqualSlice reports whether got and want match, recording a
// readable failure when they do not.
func EqualSlice[T comparable](t testing.TB, got, want []T) bool {
	// TODO(candidate): compare lengths and elements, reporting the first difference.
	panic("not implemented")
}
