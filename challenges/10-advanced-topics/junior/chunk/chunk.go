// Package chunk — Gopher Workplace challenge.
package chunk

// Chunk splits s into consecutive groups of at most n elements.
//
// The last group holds the remainder. For n <= 0 the result is nil. The
// groups are views into s — no element is copied.
//
// Examples:
//
//	Chunk([]int{1, 2, 3}, 2) => [][]int{{1, 2}, {3}}
func Chunk(s []int, n int) [][]int {
	panic("not implemented")
}
