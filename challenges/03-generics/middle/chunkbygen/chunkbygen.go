// Package chunkbygen — Gopher Workplace challenge.
package chunkbygen

// ChunkBy splits s wherever together(prev, cur) is false.
func ChunkBy[T any](s []T, together func(prev, cur T) bool) [][]T {
	// TODO(candidate): start a new group whenever together reports false.
	panic("not implemented")
}
