// Package poolgen — Gopher Workplace challenge.
package poolgen

// Pool reuses values of T instead of allocating new ones.
// It is not safe for concurrent use. Use NewPool to create one.
type Pool[T any] struct {
	free []T
	make func() T
}

// NewPool returns a pool that builds new values with make.
func NewPool[T any](make func() T) *Pool[T] {
	// TODO(candidate): store the factory.
	panic("not implemented")
}

// Get returns a pooled value, or a freshly built one.
func (p *Pool[T]) Get() T {
	// TODO(candidate): reuse a stored value, or build one.
	panic("not implemented")
}

// Put returns v to the pool for reuse.
func (p *Pool[T]) Put(v T) {
	// TODO(candidate): store v for later reuse.
	panic("not implemented")
}

// Idle returns how many values are waiting to be reused.
func (p *Pool[T]) Idle() int {
	// TODO(candidate): report the number of pooled values.
	panic("not implemented")
}
