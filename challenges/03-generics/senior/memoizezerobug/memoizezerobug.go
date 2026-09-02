// Package memoizezerobug — Gopher Workplace challenge.
package memoizezerobug

// Memoize wraps fn so each distinct argument is computed at most once.
//
// Examples:
//
//	f := Memoize(slow); f(1); f(1) // slow ran once
func Memoize[K comparable, V any](fn func(K) V) func(K) V {
	// CHANGE CODE BELOW THIS LINE
	cache := make(map[K]V)
	return func(k K) V {
		if v, ok := cache[k]; ok {
			return v
		}
		return fn(k)
	}
	// CHANGE CODE ABOVE THIS LINE
}
