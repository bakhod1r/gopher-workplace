// Package sumifgen — Gopher Workplace challenge.
package sumifgen

// Number is the set of numeric types this package works with.
type Number interface {
	~int | ~int64 | ~float64
}

// SumIf returns the total of the elements keep accepts.
//
// Examples:
//
//	SumIf([]int{1, 2, 3}, isEven) => 2
func SumIf[T Number](s []T, keep func(T) bool) T {
	// TODO(candidate): add only the elements keep accepts.
	panic("not implemented")
}
