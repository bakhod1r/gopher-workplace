// Package batcherbug — Gopher Workplace challenge.
package batcherbug

// Batches splits s into groups of exactly size, plus a final short group.
//
// Examples:
//
//	Batches([]int{1, 2, 3}, 2) => [][]int{{1, 2}, {3}}
func Batches[T any](s []T, size int) [][]T {
	// CHANGE CODE BELOW THIS LINE
	if size <= 0 {
		return [][]T{}
	}
	out := make([][]T, 0)
	cur := make([]T, 0, size)
	for _, v := range s {
		cur = append(cur, v)
		if len(cur) == size {
			out = append(out, cur)
			cur = make([]T, 0, size)
		}
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
