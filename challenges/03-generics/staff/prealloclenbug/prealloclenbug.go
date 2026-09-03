// Package prealloclenbug — Gopher Workplace challenge.
package prealloclenbug

// Map applies f to every element and returns the results in order.
//
// Examples:
//
//	Map([]int{1, 2}, double) => []int{2, 4}
func Map[T, U any](s []T, f func(T) U) []U {
	// CHANGE CODE BELOW THIS LINE
	out := make([]U, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
