// Package roundgen — Gopher Workplace challenge.
package roundgen

// Float is the set of floating-point types.
type Float interface {
	~float32 | ~float64
}

// RoundHalfUp rounds v to the nearest whole number,
// with halves rounded away from zero.
func RoundHalfUp[T Float](v T) T {
	// TODO(candidate): round to the nearest whole number, halves away from zero.
	panic("not implemented")
}
