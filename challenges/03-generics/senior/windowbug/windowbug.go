// Package windowbug — Gopher Workplace challenge.
package windowbug

// Windows returns every consecutive window of n elements.
// It returns an empty result when n <= 0 or n > len(s).
//
// Examples:
//
//	Windows([]int{1, 2, 3}, 2) => [][]int{{1, 2}, {2, 3}}
func Windows[T any](s []T, n int) [][]T {
	// CHANGE CODE BELOW THIS LINE
	out := make([][]T, 0)
	if n <= 0 || n > len(s) {
		return out
	}
	for i := 0; i < len(s); i++ {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		w := make([]T, end-i)
		copy(w, s[i:end])
		out = append(out, w)
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
