// Package genericpool — Gopher Workplace challenge.
package genericpool

import "sync"

// Pool is a typed wrapper around sync.Pool.
type Pool[T any] struct {
	inner sync.Pool
}

// NewPool returns a pool of T values.
func NewPool[T any]() *Pool[T] {
	return &Pool[T]{inner: sync.Pool{New: func() any { return new(T) }}}
}

// Put returns a value to the pool.
func (p *Pool[T]) Put(v *T) {
	if v == nil {
		return
	}
	p.inner.Put(v)
}

// Get returns a pointer to a zeroed T from the pool, or a new one when
// the pool is empty.
//
// The type parameter keeps the values typed on the way in and out, so no
// caller ever writes a type assertion.
//
// Examples:
//
//	p := NewPool[Buffer](); p.Get() => a zeroed *Buffer
func (p *Pool[T]) Get() *T {
	panic("not implemented")
}
