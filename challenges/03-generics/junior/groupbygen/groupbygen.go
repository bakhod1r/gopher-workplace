// Package groupbygen — Gopher Workplace challenge.
package groupbygen

// GroupBy buckets the elements of s by key(v), keeping order
// within each bucket.
func GroupBy[T any, K comparable](s []T, key func(T) K) map[K][]T {
	// TODO(candidate): append each element into the bucket for its key.
	panic("not implemented")
}
