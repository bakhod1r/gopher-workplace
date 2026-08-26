// Package popstk — Gopher Workplace challenge.
package popstk

// Stack is a stack of integers.
type Stack struct {
	Items []int
}

// Pop removes and returns the top element. If the stack is empty, it returns
// 0 and false.
//
// Examples:
//
//	s := Stack{Items: []int{1, 2, 3}}; v, ok := s.Pop() // 3, true
//	s = Stack{}; v, ok = s.Pop()                         // 0, false
func (s *Stack) Pop() (int, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
