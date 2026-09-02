// Package aliasbug — Gopher Workplace challenge.
package aliasbug

// Chunk splits s into consecutive groups of at most size elements.
// Each group is independent of the input.
//
// Examples:
//
//	Chunk([]int{1, 2, 3}, 2) => [][]int{{1, 2}, {3}}
func Chunk[T any](s []T, size int) [][]T {
	// CHANGE CODE BELOW THIS LINE
	if size <= 0 {
		return [][]T{}
	}
	out := make([][]T, 0)
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		out = append(out, s[i:end])
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
