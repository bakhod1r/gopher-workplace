// Package stackembed — Gopher Workplace challenge.
package stackembed

// Stack is a last-in-first-out collection of T.
type Stack[T any] struct {
	items []T
}

// TracedStack is a Stack that also records push activity.
// Its zero value is ready to use.
type TracedStack[T any] struct {
	Stack[T]
	last   T
	pushes int
}

// Push adds v and records it as the last pushed value.
func (t *TracedStack[T]) Push(v T) {
	// TODO(candidate): delegate to the embedded stack, then record v.
	panic("not implemented")
}

// Pushes returns how many times Push was called.
func (t *TracedStack[T]) Pushes() int {
	// TODO(candidate): report the recorded push count.
	panic("not implemented")
}

// Last returns the value most recently pushed and true,
// or the zero value and false when nothing was pushed.
func (t *TracedStack[T]) Last() (T, bool) {
	// TODO(candidate): report the last pushed value.
	panic("not implemented")
}

// Push adds v to the stack. It is provided for you.
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Len returns the stack size. It is provided for you.
func (s *Stack[T]) Len() int {
	return len(s.items)
}
