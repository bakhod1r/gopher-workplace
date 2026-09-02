// Package satadd — Gopher Workplace challenge.
package satadd

// Unsigned is the set of unsigned integer types.
type Unsigned interface {
	~uint | ~uint64
}

// SatAdd returns a+b, or the maximum value of T on overflow.
func SatAdd[T Unsigned](a, b T) T {
	// TODO(candidate): add, detecting wrap-around.
	panic("not implemented")
}
