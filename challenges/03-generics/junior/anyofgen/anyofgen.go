// Package anyofgen — Gopher Workplace challenge.
package anyofgen

// Any reports whether at least one element satisfies pred.
// Any returns false for an empty slice.
//
// Examples:
//
//	Any([]int{1, 2}, isEven) => true
//	Any([]int{}, isEven)     => false
func Any[T any](s []T, pred func(T) bool) bool {
	// TODO(candidate): return true at the first element pred accepts.
	panic("not implemented")
}

// All reports whether every element satisfies pred.
// All returns true for an empty slice.
//
// Examples:
//
//	All([]int{2, 4}, isEven) => true
//	All([]int{}, isEven)     => true
func All[T any](s []T, pred func(T) bool) bool {
	// TODO(candidate): return false at the first element pred rejects.
	panic("not implemented")
}
