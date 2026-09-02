// Package satbug — Gopher Workplace challenge.
package satbug

// Unsigned is the set of unsigned integer types.
type Unsigned interface {
	~uint | ~uint64
}

// SatAdd returns a+b, or the maximum value of T on overflow.
//
// Examples:
//
//	SatAdd(maxUint, 1) => maxUint
func SatAdd[T Unsigned](a, b T) T {
	// CHANGE CODE BELOW THIS LINE
	sum := a + b
	if sum < 0 {
		return ^T(0)
	}
	return sum
	// CHANGE CODE ABOVE THIS LINE
}
