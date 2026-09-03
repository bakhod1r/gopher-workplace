// Package lazyzerobug — Gopher Workplace challenge.
package lazyzerobug

// Lazy computes Fn once and caches the result.
type Lazy[T comparable] struct {
	Fn    func() T
	value T
	done  bool
}

// Get returns the memoised value, computing it on first use.
func (l *Lazy[T]) Get() T {
	// CHANGE CODE BELOW THIS LINE
	var zero T
	if any(l.value) == any(zero) {
		l.value = l.Fn()
	}
	return l.value
	// CHANGE CODE ABOVE THIS LINE
}
