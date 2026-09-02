// Package absgen — Gopher Workplace challenge.
package absgen

// Signed is the set of signed numeric types.
type Signed interface {
	~int | ~int64 | ~float64
}

// Abs returns the magnitude of v.
//
// Examples:
//
//	Abs(-3)    => 3
//	Abs(2.5)   => 2.5
func Abs[T Signed](v T) T {
	// TODO(candidate): drop the sign.
	panic("not implemented")
}
