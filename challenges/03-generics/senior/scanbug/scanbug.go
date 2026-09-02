// Package scanbug — Gopher Workplace challenge.
package scanbug

// Scan returns the accumulator's value after each element.
// The result has exactly len(s) elements.
//
// Examples:
//
//	Scan([]int{1, 2, 3}, 0, add) => []int{1, 3, 6}
func Scan[T, A any](s []T, init A, f func(A, T) A) []A {
	// CHANGE CODE BELOW THIS LINE
	out := make([]A, 0, len(s))
	acc := init
	for _, v := range s {
		out = append(out, acc)
		acc = f(acc, v)
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
