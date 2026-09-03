// Package uniquekeybug — Gopher Workplace challenge.
package uniquekeybug

// UniqueBy keeps the first element for each distinct key,
// in input order.
//
// Examples:
//
//	UniqueBy(rows, idOf) => one row per id, the first
func UniqueBy[T any, K comparable](s []T, key func(T) K) []T {
	// CHANGE CODE BELOW THIS LINE
	byKey := make(map[K]T, len(s))
	order := make([]K, 0, len(s))
	for _, v := range s {
		k := key(v)
		if _, ok := byKey[k]; !ok {
			order = append(order, k)
		}
		byKey[k] = v
	}
	out := make([]T, 0, len(order))
	for _, k := range order {
		out = append(out, byKey[k])
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
