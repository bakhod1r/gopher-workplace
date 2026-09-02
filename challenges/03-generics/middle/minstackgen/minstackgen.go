// Package minstackgen — Gopher Workplace challenge.
package minstackgen

import (
	"cmp"
)

// MinStack is a stack of T that can report its minimum in O(1).
// Its zero value is an empty stack.
type MinStack[T cmp.Ordered] struct {
	items []T
	mins  []T
}

// Push adds v to the stack.
func (s *MinStack[T]) Push(v T) {
	// TODO(candidate): store v and track the running minimum.
	panic("not implemented")
}

// Pop removes and returns the top element and true.
func (s *MinStack[T]) Pop() (T, bool) {
	// TODO(candidate): remove the top element and the matching minimum.
	panic("not implemented")
}

// Min returns the smallest stored element and true.
func (s *MinStack[T]) Min() (T, bool) {
	// TODO(candidate): report the current minimum in constant time.
	panic("not implemented")
}
