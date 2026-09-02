// Package tabletestgen — Gopher Workplace challenge.
package tabletestgen

import (
	"testing"
)

// Case is one table-driven test case.
type Case[I, W any] struct {
	Name string
	In   I
	Want W
}

// Run executes every case as a subtest, comparing fn's output
// with the expected value.
func Run[I, W comparable](t *testing.T, name string, cases []Case[I, W], fn func(I) W) {
	// TODO(candidate): run each case as a subtest and compare.
	panic("not implemented")
}
