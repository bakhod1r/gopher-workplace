// Package stackofstacksgen — Gopher Workplace challenge.
package stackofstacksgen

// PlateStack is a stack split into inner stacks of at most cap items.
// Its zero value uses a capacity of 1.
type PlateStack[T any] struct {
	stacks [][]T
	cap    int
}

// Push adds v, starting a new inner stack when the current one
// is full.
func (s *PlateStack[T]) Push(v T) {
	// TODO(candidate): append to the last inner stack, starting one when needed.
	panic("not implemented")
}

// Pop removes and returns the most recently pushed element,
// dropping an inner stack once it empties.
func (s *PlateStack[T]) Pop() (T, bool) {
	// TODO(candidate): pop from the last inner stack, dropping it when empty.
	panic("not implemented")
}

// Stacks returns how many inner stacks are in use.
func (s *PlateStack[T]) Stacks() int {
	// TODO(candidate): report the number of inner stacks.
	panic("not implemented")
}

// Cap sets the capacity of each inner stack.
// It is provided for you.
func (s *PlateStack[T]) Cap(n int) {
	if n < 1 {
		n = 1
	}
	s.cap = n
}
