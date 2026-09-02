// Package everynthzerobug — Gopher Workplace challenge.
package everynthzerobug

// EveryNth returns every n-th element starting from index 0.
// It returns an empty slice when n is not positive.
//
// Examples:
//
//	EveryNth([]int{0, 1, 2, 3}, 2) => []int{0, 2}
func EveryNth[T any](s []T, n int) []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, 0)
	if n <= 0 {
		return out
	}
	for i := n; i < len(s); i += n {
		out = append(out, s[i])
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
