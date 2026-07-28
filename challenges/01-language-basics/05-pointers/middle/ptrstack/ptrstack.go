// Package ptrstack — Gopher Workplace challenge.
package ptrstack

type Stack struct{ data []int }

func (s *Stack) Len() int { return len(s.data) }

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}

// Push pushes v onto the stack (pointer receiver so it mutates).
func (s *Stack) Push(v int) {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
