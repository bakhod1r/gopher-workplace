// Package assertequalhelper — Gopher Workplace challenge.
package assertequalhelper

import (
	"testing"
)

// Equal reports whether got and want match, and records a
// readable failure on t when they do not.
func Equal[T comparable](t testing.TB, got, want T) bool {
	// TODO(candidate): compare, reporting a failure when they differ.
	panic("not implemented")
}
