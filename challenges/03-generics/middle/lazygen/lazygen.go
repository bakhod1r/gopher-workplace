// Package lazygen — Gopher Workplace challenge.
package lazygen

// Lazy computes a value of T on first use and remembers it.
// It is not safe for concurrent use. Use NewLazy to create one.
type Lazy[T any] struct {
	compute func() T
	value   T
	done    bool
}

// NewLazy returns a value computed on first use.
func NewLazy[T any](compute func() T) *Lazy[T] {
	// TODO(candidate): store the computation.
	panic("not implemented")
}

// Get returns the value, computing it at most once.
func (l *Lazy[T]) Get() T {
	// TODO(candidate): compute on first use, then reuse.
	panic("not implemented")
}

// Done reports whether the value has been computed.
func (l *Lazy[T]) Done() bool {
	// TODO(candidate): report whether the value exists yet.
	panic("not implemented")
}
