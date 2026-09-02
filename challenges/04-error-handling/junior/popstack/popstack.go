// Package popstack — Gopher Workplace challenge.
package popstack

import "errors"

// ErrEmpty reports a pop from an empty stack.
var ErrEmpty = errors.New("empty stack")

// Pop removes and returns the last element of s.
//
// Examples:
//
//	Pop([]int{1, 2, 3}) => [1 2], 3, nil
//	Pop(nil)            => nil, 0, ErrEmpty
func Pop(s []int) ([]int, int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
