// Package dedupe — Gopher Workplace challenge.
package dedupe

// Dedupe returns the elements of in with duplicates removed, keeping the order
// in which each value first appeared. The input slice must not be mutated, and
// a nil or empty input returns an empty (non-nil) slice.
//
// Examples:
//
//	Dedupe([]int{1, 1, 2, 3, 3, 3}) => []int{1, 2, 3}
//	Dedupe([]int{5, 4, 5, 4})       => []int{5, 4}   // first-appearance order
//	Dedupe([]int{7, 7, 7})          => []int{7}
//	Dedupe(nil)                     => []int{}
func Dedupe(in []int) []int {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
