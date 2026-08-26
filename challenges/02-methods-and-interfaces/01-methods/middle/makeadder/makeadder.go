// Package makeadder — Gopher Workplace challenge.
package makeadder

// Number holds an integer.
type Number struct {
	Val int
}

// Adder returns a function that adds x to Val and returns the result.
//
// Examples:
//
//	add := Number{5}.Adder()
//	add(3) => 8
func (n Number) Adder() func(int) int {
	// TODO(candidate): return a closure that adds x to n.Val.
	panic("not implemented")
}
