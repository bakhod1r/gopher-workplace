// Package almostequal — Gopher Workplace challenge.
package almostequal

// AlmostEqual reports whether a and b are within a small tolerance of each
// other. Floating-point arithmetic is inexact — 0.1 + 0.2 is not exactly 0.3 —
// so floats must be compared with a tolerance (epsilon), never with ==.
//
// Two values count as equal when their absolute difference is < 1e-9. Use the
// standard library's math.Abs for the absolute value (add `import "math"`).
//
// Examples:
//
//	AlmostEqual(0.1+0.2, 0.3) => true   // == would be false here
//	AlmostEqual(1.0, 1.0)     => true
//	AlmostEqual(1.0, 1.0001)  => false
//	AlmostEqual(-2.5, -2.5)   => true
func AlmostEqual(a, b float64) bool {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
