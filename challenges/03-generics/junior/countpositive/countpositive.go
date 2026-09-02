// Package countpositive — Gopher Workplace challenge.
package countpositive

// Number is the set of numeric types this package works with.
type Number interface {
	~int | ~int64 | ~float64
}

// CountPositive returns how many elements of s are greater than zero.
//
// Examples:
//
//	CountPositive([]int{-1, 0, 2}) => 1
func CountPositive[T Number](s []T) int {
	// TODO(candidate): count the elements greater than zero.
	panic("not implemented")
}
