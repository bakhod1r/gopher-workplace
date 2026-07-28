// Package nancheck detects non-finite floats.
package nancheck

// Finite reports whether x is a normal finite number (not NaN, not ±Inf).
//
// TODO(candidate): implement using math.IsNaN and math.IsInf.
func Finite(x float64) bool {
	panic("not implemented")
}
