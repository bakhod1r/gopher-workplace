// Package memorize — Gopher Workplace challenge.
package memorize

// Memoizer caches results of expensive calls.
type Memoizer struct {
	cache map[string]string
	fn    func(string) string
}

// New creates a Memoizer for fn.
func New(fn func(string) string) *Memoizer {
	return &Memoizer{cache: make(map[string]string), fn: fn}
}

// Get returns the cached result, or calls fn, caches, and returns it.
func (m *Memoizer) Get(key string) string {
	// TODO(candidate): implement memoization.
	panic("not implemented")
}
