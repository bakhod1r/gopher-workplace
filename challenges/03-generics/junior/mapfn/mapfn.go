// Package mapfn — Gopher Workplace challenge.
package mapfn

// Map applies f to every element of s and returns the results.
//
// Examples:
//
//	Map([]int{1, 2}, func(n int) int { return n * 2 })    => []int{2, 4}
//	Map([]int{1}, func(n int) string { return "x" })    => []string{"x"}
func Map[T, U any](s []T, f func(T) U) []U {
	// TODO(candidate): apply f to each element, collecting the results.
	panic("not implemented")
}
