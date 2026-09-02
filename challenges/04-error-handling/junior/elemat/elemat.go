// Package elemat — Gopher Workplace challenge.
package elemat

import "errors"

// ErrOutOfRange reports an index outside the slice.
var ErrOutOfRange = errors.New("index out of range")

// At returns s[i], or ErrOutOfRange when i is invalid.
//
// Examples:
//
//	At([]int{1, 2, 3}, 0) => 1, nil
//	At([]int{1, 2, 3}, 3) => 0, ErrOutOfRange
func At(s []int, i int) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
