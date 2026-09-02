# Typed Buffer Pool

## Intuition

`sync.Pool` is a per-P cache of reusable objects: `Get` prefers the current
processor's local free list, so in the common case it is nearly free of
contention. It exists to reduce allocation *rate*, not to guarantee reuse — the
runtime clears pools during garbage collection, and objects can disappear
between a `Put` and the next `Get`.

Wrapping it is about types: `Get() any` at every call site invites assertions
scattered through the codebase.

## Approach

1. `Get`: pull from the pool and assert once, here.
2. `Put`: hand the object back.

## Solution

```go
func (p *BufferPool) Get() *Buffer {
	return p.pool.Get().(*Buffer)
}

func (p *BufferPool) Put(b *Buffer) {
	p.pool.Put(b)
}
```

## Walkthrough

`NewPool` sets a `New` function, so an empty pool manufactures a fresh
`&Buffer{Data: make([]byte, 1024)}` rather than returning nil. That is why the
unchecked assertion is safe: `Get` can only ever produce a `*Buffer`.

Two `Get` calls with no `Put` in between must return different objects — the
pool removes what it hands out. After a `Put`, the next `Get` on the same
processor very likely returns the same pointer, which is what makes the pool
worth having.

## Pitfalls

- **Pooling without a `New` func.** `Get` then returns nil on an empty pool and
  the assertion panics. Always set `New`, or use the comma-ok form.
- **Not resetting on `Put` (or on `Get`).** The next user inherits the previous
  contents — a data leak between requests. `bytes.Buffer` pools call `Reset` for
  exactly this reason.
- **Putting back a buffer someone still holds.** Two goroutines then share one
  object; this is the pooling bug that produces impossible-looking corruption.
- **Storing pointers to pooled memory elsewhere.** Once returned, the object may
  be reused at any moment.
- **Pooling small objects.** Below a few hundred bytes, Go's allocator usually
  beats the pool's bookkeeping. Measure before pooling.

## Why entries disappear

Go's GC clears every pool's victim cache on each cycle, so a pool never keeps
garbage alive across two collections. That is a deliberate trade: pools help
steady-state throughput without turning into unbounded memory.
