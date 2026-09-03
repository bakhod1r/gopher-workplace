// Package concatalias — Gopher Workplace challenge.
package concatalias

// Concat joins the slices into one new slice.
// The result never shares storage with any input.
//
// Examples:
//
//	Concat([]int{1}, []int{2}) => []int{1, 2}
func Concat[T any](ss ...[]T) []T {
	// CHANGE CODE BELOW THIS LINE
	if len(ss) == 1 {
		return ss[0]
	}
	n := 0
	for _, s := range ss {
		n += len(s)
	}
	out := make([]T, 0, n)
	for _, s := range ss {
		out = append(out, s...)
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
