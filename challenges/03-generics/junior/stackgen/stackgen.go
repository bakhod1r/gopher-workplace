// Package stackgen — Gopher Workplace challenge.
package stackgen

// Stack is a last-in-first-out collection of T.
// Its zero value is an empty stack.
type Stack[T any] struct {
	items []T
}

// Push adds v to the top of the stack.
func (s *Stack[T]) Push(v T) {
	// TODO(candidate): append v to the items.
	panic("not implemented")
}

// Pop removes and returns the top element and true.
// It returns the zero value and false when the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	// TODO(candidate): remove and return the last item.
	panic("not implemented")
}

// Len returns the number of elements on the stack.
func (s *Stack[T]) Len() int {
	// TODO(candidate): report how many items are stored.
	panic("not implemented")
}
