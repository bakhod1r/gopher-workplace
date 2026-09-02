// Package groupsizegen — Gopher Workplace challenge.
package groupsizegen

// Partition splits s into the elements pred accepts and the rest,
// each in original order.
//
// Examples:
//
//	Partition([]int{1, 2, 3}, isEven) => []int{2}, []int{1, 3}
//	Partition([]int{}, isEven)        => []int{}, []int{}
func Partition[T any](s []T, pred func(T) bool) ([]T, []T) {
	// TODO(candidate): split into accepted and rejected, preserving order.
	panic("not implemented")
}
