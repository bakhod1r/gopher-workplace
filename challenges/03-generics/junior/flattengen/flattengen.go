// Package flattengen — Gopher Workplace challenge.
package flattengen

// Flatten concatenates the groups in order into one slice.
//
// Examples:
//
//	Flatten([][]int{{1, 2}, {3}}) => []int{1, 2, 3}
//	Flatten([][]int{})            => []int{}
func Flatten[T any](groups [][]T) []T {
	// TODO(candidate): append every group in order into one slice.
	panic("not implemented")
}
