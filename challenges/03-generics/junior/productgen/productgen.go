// Package productgen — Gopher Workplace challenge.
package productgen

// Number is the set of numeric types this package works with.
type Number interface {
	~int | ~int64 | ~float64
}

// Product returns the product of s. It returns 1 for an empty slice.
//
// Examples:
//
//	Product([]int{2, 3}) => 6
//	Product([]int{})     => 1
func Product[T Number](s []T) T {
	// TODO(candidate): multiply every element into a running product.
	panic("not implemented")
}
