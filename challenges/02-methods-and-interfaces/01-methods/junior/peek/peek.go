// Package peek — Gopher Workplace challenge.
package peek

// Stack is a stack of integers.
type Stack struct {
	Items []int
}

// Peek returns the top element without removing it. Returns 0 and false if
// the stack is empty.
//
// Examples:
//
//	Stack{Items: []int{1, 2, 3}}.Peek() => (3, true)
//	Stack{}.Peek()                       => (0, false)
func (s Stack) Peek() (int, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
