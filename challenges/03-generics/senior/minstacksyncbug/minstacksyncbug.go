// Package minstacksyncbug — Gopher Workplace challenge.
package minstacksyncbug

import (
	"cmp"
)

// MinStack is a stack with O(1) Min.
type MinStack[T cmp.Ordered] struct {
	items []T
	mins  []T
}

// Pop removes and returns the top element and true.
// It returns the zero value and false when empty.
func (s *MinStack[T]) Pop() (T, bool) {
	// CHANGE CODE BELOW THIS LINE
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
	// CHANGE CODE ABOVE THIS LINE
}

// Push adds v to the stack.
func (s *MinStack[T]) Push(v T) {
	s.items = append(s.items, v)
	if len(s.mins) == 0 || v <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, v)
	}
}

// Min returns the smallest element still on the stack and true.
func (s *MinStack[T]) Min() (T, bool) {
	if len(s.mins) == 0 {
		var zero T
		return zero, false
	}
	return s.mins[len(s.mins)-1], true
}
