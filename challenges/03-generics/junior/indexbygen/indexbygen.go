// Package indexbygen — Gopher Workplace challenge.
package indexbygen

// IndexBy maps key(v) to v. On duplicate keys the last element wins.
func IndexBy[T any, K comparable](s []T, key func(T) K) map[K]T {
	// TODO(candidate): build a lookup keyed by key(v).
	panic("not implemented")
}
