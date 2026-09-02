// Package trackergen — Gopher Workplace challenge.
package trackergen

import (
	"cmp"
)

// Tracker records the smallest and largest values it has seen.
// Its zero value is ready to use.
type Tracker[T cmp.Ordered] struct {
	lo T
	hi T
	n  int
}

// Add records v.
func (t *Tracker[T]) Add(v T) {
	// TODO(candidate): record v, updating the smallest and largest seen.
	panic("not implemented")
}

// Bounds returns the smallest and largest recorded values and true.
// It returns zero values and false before anything is recorded.
func (t *Tracker[T]) Bounds() (T, T, bool) {
	// TODO(candidate): report the recorded range.
	panic("not implemented")
}
