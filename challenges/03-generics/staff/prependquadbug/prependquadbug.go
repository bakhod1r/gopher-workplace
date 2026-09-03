// Package prependquadbug — Gopher Workplace challenge.
package prependquadbug

// Reversed returns a new slice holding s's elements in reverse order.
//
// It must run in linear time and allocate exactly once.
//
// Examples:
//
//	Reversed([]int{1, 2, 3}) => []int{3, 2, 1}
func Reversed[T any](s []T) []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, 0, len(s))
	for _, v := range s {
		out = append([]T{v}, out...)
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
