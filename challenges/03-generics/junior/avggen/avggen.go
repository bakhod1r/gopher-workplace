// Package avggen — Gopher Workplace challenge.
package avggen

// Number is the set of numeric types this package works with.
type Number interface {
	~int | ~int64 | ~float64
}

// Average returns the arithmetic mean of s as a float64.
// It returns 0 for an empty slice.
//
// Examples:
//
//	Average([]int{1, 2, 3}) => 2
//	Average([]int{})        => 0
func Average[T Number](s []T) float64 {
	// TODO(candidate): sum as float64, then divide by the count.
	panic("not implemented")
}
