// Package groupbybug — Gopher Workplace challenge.
package groupbybug

// GroupBy buckets elements by key, preserving input order within a bucket.
//
// Examples:
//
//	GroupBy([]int{1, 2, 3}, parity) => map[odd:[1 3] even:[2]]
func GroupBy[T any, K comparable](s []T, key func(T) K) map[K][]T {
	// CHANGE CODE BELOW THIS LINE
	out := make(map[K][]T)
	for _, v := range s {
		k := key(v)
		out[k] = []T{v}
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
