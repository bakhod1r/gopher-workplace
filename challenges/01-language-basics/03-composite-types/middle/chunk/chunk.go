// Package chunk splits a slice into fixed-size chunks.
package chunk

// Chunk splits xs into consecutive slices of at most size elements. The last
// chunk may be shorter. size <= 0 returns an empty result.
//
// TODO(candidate): implement with slice expressions.
func Chunk(xs []int, size int) [][]int {
	panic("not implemented")
}
