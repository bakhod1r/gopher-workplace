// Package vecgen — Gopher Workplace challenge.
package vecgen

// Number is the numeric set used for vector arithmetic.
type Number interface {
	~int | ~int64 | ~float64
}

// Add returns the element-wise sum of a and b.
// It returns an empty slice when the lengths differ.
func Add[T Number](a, b []T) []T {
	// TODO(candidate): add element-wise, rejecting mismatched lengths.
	panic("not implemented")
}

// Dot returns the dot product of a and b, and whether the
// lengths matched.
func Dot[T Number](a, b []T) (T, bool) {
	// TODO(candidate): sum the element-wise products.
	panic("not implemented")
}
