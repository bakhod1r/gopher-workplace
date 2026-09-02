// Package requirelengen — Gopher Workplace challenge.
package requirelengen

import (
	"testing"
)

// RequireLen fails the test immediately when s has the wrong
// length. Use it before indexing into a result.
func RequireLen[T any](t testing.TB, s []T, want int) {
	// TODO(candidate): stop the test when the length is wrong.
	panic("not implemented")
}
