// Package signgen — Gopher Workplace challenge.
package signgen

// Signed is the set of signed numeric types.
type Signed interface {
	~int | ~int64 | ~float64
}

// Sign returns -1 for negative v, 0 for zero, and +1 for positive v.
//
// Examples:
//
//	Sign(-2)   => -1
//	Sign(0)    => 0
//	Sign(1.5)  => 1
func Sign[T Signed](v T) int {
	// TODO(candidate): return -1, 0, or +1.
	panic("not implemented")
}
