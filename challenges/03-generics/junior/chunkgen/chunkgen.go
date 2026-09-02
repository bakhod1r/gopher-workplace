// Package chunkgen — Gopher Workplace challenge.
package chunkgen

// Chunk splits s into consecutive groups of at most size elements.
// It returns an empty result when size <= 0 or s is empty.
//
// Examples:
//
//	Chunk([]int{1, 2, 3}, 2) => [][]int{{1, 2}, {3}}
//	Chunk([]int{1}, 0)       => [][]int{}
func Chunk[T any](s []T, size int) [][]T {
	// TODO(candidate): cut s into consecutive groups of at most size elements.
	panic("not implemented")
}
