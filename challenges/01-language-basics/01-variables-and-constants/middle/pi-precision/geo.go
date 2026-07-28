// Package geo uses a high-precision untyped constant for circle math.
package geo

// Pi is an untyped constant with more precision than float64 can hold; the
// compiler rounds it at the point of use.
//
// TODO(candidate): define Pi to at least 20 significant digits as an untyped
// constant (no type on the left).
const Pi = 3.0

// Area returns Pi*r*r as float64.
//
// TODO(candidate): implement.
func Area(r float64) float64 {
	panic("not implemented")
}
