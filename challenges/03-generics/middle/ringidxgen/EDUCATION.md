# Ring Buffer With Indices

## Intuition

Wrapping with arithmetic rather than reslicing is what makes this allocation-free — the same reason ring buffers are used in hot paths.

## Approach

1. `NewRing`: allocate `size` cells.
2. `Add`: write at the head, advance modulo the capacity, grow the count up to the capacity.
3. `Items`: start `n` positions behind the head and read forward, wrapping.

## Solution

```go
func NewRing[T any](size int) *Ring[T] {
	if size < 0 {
		size = 0
	}
	return &Ring[T]{buf: make([]T, size)}
}

func (r *Ring[T]) Add(v T) {
	if len(r.buf) == 0 {
		return
	}
	r.buf[r.head] = v
	r.head = (r.head + 1) % len(r.buf)
	if r.n < len(r.buf) {
		r.n++
	}
}

func (r *Ring[T]) Items() []T {
	out := make([]T, 0, r.n)
	if r.n == 0 {
		return out
	}
	start := (r.head - r.n + len(r.buf)) % len(r.buf)
	for i := 0; i < r.n; i++ {
		out = append(out, r.buf[(start+i)%len(r.buf)])
	}
	return out
}
```

## Walkthrough

With capacity 2, adding `1, 2, 3` overwrites cell 0 with `3`, and `Items` reads from the oldest position to yield `[2 3]`.

## Pitfalls

- Computing the start index without adding the capacity, producing a negative index.
- Letting the count exceed the capacity.
- Returning the raw buffer, which exposes stale cells and internal order.
