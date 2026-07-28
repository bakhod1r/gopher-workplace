// Package subslice — Gopher Workplace challenge.
package subslice

// Head returns the first n elements of s as an INDEPENDENT copy: later writes
// to the result must not affect s, and writes to s must not affect the result.
// A plain slice expression s[:n] would share s's backing array — this must not.
// If n is greater than len(s), all of s is returned; n <= 0 gives an empty slice.
//
// Examples:
//
//	Head([]int{1, 2, 3, 4}, 2) => []int{1, 2}   // independent of the input
//	Head([]int{1, 2}, 5)       => []int{1, 2}
//	Head([]int{1, 2, 3}, 0)    => []int{}
func Head(s []int, n int) []int {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
