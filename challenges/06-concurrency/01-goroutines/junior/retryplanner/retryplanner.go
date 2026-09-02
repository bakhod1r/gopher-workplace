// Package retryplanner — Gopher Workplace challenge.
package retryplanner

// Backoffs returns the exponential backoff delay for each attempt number.
//
// Examples:
//
//	Backoffs([]int{0, 1, 3}, 100)  => [100 200 800]
//	Backoffs([]int{-1}, 100)       => [100]
//	Backoffs(nil, 100)             => []
func Backoffs(attempts []int, baseMs int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
