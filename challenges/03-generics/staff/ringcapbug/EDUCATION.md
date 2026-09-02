# The Ring That Lends Out Its Buffer

## Intuition

A contiguous window can be returned without copying — but only if its capacity is capped. A two-index slice inherits the ring's full capacity, so the caller's `append` lands directly in the ring's next write slot.

## Approach

1. Return an empty slice when the ring is empty.
2. When the window does not wrap, return it with a three-index slice so `cap == len`.
3. When it wraps, copy the two segments into a fresh slice.

## Solution

```go
func (r *Ring[T]) Slice() []T {
	if r.n == 0 {
		return []T{}
	}
	end := r.head + r.n
	if end <= len(r.buf) {
		return r.buf[r.head:end:end]
	}
	out := make([]T, 0, r.n)
	out = append(out, r.buf[r.head:]...)
	out = append(out, r.buf[:end-len(r.buf)]...)
	return out
}

func (r *Ring[T]) Push(v T) {
	if r.n == len(r.buf) {
		r.buf[r.head] = v
		r.head = (r.head + 1) % len(r.buf)
		return
	}
	r.buf[(r.head+r.n)%len(r.buf)] = v
	r.n++
}

func (r *Ring[T]) Len() int {
	return r.n
}
```

## Walkthrough

With `cap` 8 and three elements, `r.buf[0:3]` has capacity 8. `append(snapshot, 99)` writes `99` into `r.buf[3]`, which is exactly where the next `Push` was going to store a sample — and the appended slice then observes that sample.

## Pitfalls

- Assuming the wrapped branch protects you — the bug only shows on the fast path.
- Fixing it by always copying, which throws away the zero-allocation fast path.
