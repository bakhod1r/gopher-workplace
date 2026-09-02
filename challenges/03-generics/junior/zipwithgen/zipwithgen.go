// Package zipwithgen — Gopher Workplace challenge.
package zipwithgen

// ZipWith combines a and b element-wise with f, stopping at the
// shorter of the two slices.
//
// Examples:
//
//	ZipWith([]int{1, 2}, []int{10, 20}, add) => []int{11, 22}
//	ZipWith([]int{1}, []int{}, add)          => []int{}
func ZipWith[T, U, R any](a []T, b []U, f func(T, U) R) []R {
	// TODO(candidate): combine matching positions until the shorter slice runs out.
	panic("not implemented")
}
