// Package sliceresetreuse — Gopher Workplace challenge.
package sliceresetreuse

// Reset empties s for reuse while keeping its allocated array, so the next
// round of appends costs nothing. A nil slice resets to an empty, non-nil
// slice.
//
// Examples:
//
//	len(Reset([]int{1,2,3})) == 0, cap unchanged
func Reset(s []int) []int {
	panic("not implemented")
}

// FillEvens overwrites buf with the even numbers below 2*n, reusing buf's
// capacity when it is large enough.
//
// Examples:
//
//	FillEvens(nil, 3) => []int{0, 2, 4}
func FillEvens(buf []int, n int) []int {
	panic("not implemented")
}
