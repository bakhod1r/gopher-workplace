// Package mustfind — Gopher Workplace challenge.
package mustfind

import "errors"

// ErrNotFound reports an absent value.
var ErrNotFound = errors.New("value not found")

// Find returns the index of target in s.
//
// Examples:
//
//	Find([]int{4, 7, 9}, 7) => 1, nil
//	Find([]int{4, 7}, 5)    => -1, ErrNotFound
func Find(s []int, target int) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
