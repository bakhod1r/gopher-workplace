// Package flatmapgen — Gopher Workplace challenge.
package flatmapgen

// FlatMap applies f to each element and concatenates the results.
func FlatMap[T, U any](s []T, f func(T) []U) []U {
	// TODO(candidate): concatenate the slices f returns.
	panic("not implemented")
}
