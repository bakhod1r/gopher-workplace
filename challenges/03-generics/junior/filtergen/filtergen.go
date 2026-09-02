// Package filtergen — Gopher Workplace challenge.
package filtergen

// Filter returns the elements of s for which keep returns true,
// in their original order.
//
// Examples:
//
//	Filter([]int{1, 2, 3}, isEven) => []int{2}
//	Filter([]int{}, isEven)        => []int{}
func Filter[T any](s []T, keep func(T) bool) []T {
	// TODO(candidate): collect the elements keep accepts.
	panic("not implemented")
}
