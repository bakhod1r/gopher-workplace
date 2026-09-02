// Package reducegen — Gopher Workplace challenge.
package reducegen

// Reduce folds s into a single accumulator value, starting
// from init and applying f left to right.
//
// Examples:
//
//	Reduce([]int{1, 2, 3}, 0, add)  => 6
//	Reduce([]int{}, 5, add)         => 5
func Reduce[T, A any](s []T, init A, f func(A, T) A) A {
	// TODO(candidate): fold the elements into the accumulator, left to right.
	panic("not implemented")
}
