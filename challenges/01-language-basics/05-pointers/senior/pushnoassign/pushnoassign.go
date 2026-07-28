// Package pushnoassign pushes onto a slice-backed stack. A planted bug discards
// the append result, so growth past capacity is lost.
package pushnoassign

type Stack struct{ data []int }

// Push appends v to the stack.
func (s *Stack) Push(v int) {
	// CHANGE CODE BELOW THIS LINE
	_ = append(s.data, v)
	// CHANGE CODE ABOVE THIS LINE
}
