// Package satclampbug — Gopher Workplace challenge.
package satclampbug

// Signed is the set of signed integer types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// AddSat adds a and b, clamping at the limits of T instead of wrapping.
//
// Examples:
//
//	AddSat(int8(100), int8(100)) => 127
func AddSat[T Signed](a, b T) T {
	// CHANGE CODE BELOW THIS LINE
	minV, maxV := limits[T]()
	sum := a + b
	if sum > maxV {
		return maxV
	}
	if sum < minV {
		return minV
	}
	return sum
	// CHANGE CODE ABOVE THIS LINE
}

// limits returns the smallest and largest value representable in T.
func limits[T Signed]() (T, T) {
	var maxV T
	for x := T(1); x > 0; x <<= 1 {
		maxV |= x
	}
	return -maxV - 1, maxV
}
