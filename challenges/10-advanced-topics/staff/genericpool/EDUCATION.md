# A Typed Pool With No Assertions

## Intuition

A pool's danger is stale state and its friction is type assertions. Generics remove the second, and a single assignment of the zero value removes the first for every type at once.

## Approach

1. `p.inner.Get().(*T)` — safe because only `*T` values ever enter the pool.
2. Overwrite with the zero `T`.
3. Return the pointer.

## Solution

```go
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
// 	p := NewPool[Buffer](); p.Get() => a zeroed *Buffer
func (p *Pool[T]) Get() *T {
	v := p.inner.Get().(*T)
	var zero T
	*v = zero
	return v
}
```

## Walkthrough

A `buffer` put back with `Name` set comes out of the pool with that name still in it; assigning the zero value clears every field without the wrapper knowing what they are.

## Pitfalls

- Skipping the reset and documenting that callers must do it — one of them will not.
- Putting a value back while still holding a pointer to it.
- Assuming the pool retains what you put; entries can be dropped at any collection.
