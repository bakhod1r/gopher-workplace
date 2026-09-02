// Package tosetgen — Gopher Workplace challenge.
package tosetgen

// ToSet returns a set of the distinct values in s.
//
// Examples:
//
//	ToSet([]int{1, 1, 2}) => map[int]struct{}{1: {}, 2: {}}
//	ToSet([]int{})        => map[int]struct{}{}
func ToSet[T comparable](s []T) map[T]struct{} {
	// TODO(candidate): add every element as a key.
	panic("not implemented")
}
