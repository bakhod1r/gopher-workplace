# Ring Buffer Index Underflow

## Intuition

With `head = 1` and `n = 1` in a buffer of 3 the expression is `0`, but `head = 0, n = 2` gives `-2`, and Go's `%` leaves it negative.

## Approach

1. Add the capacity before taking the modulus.
2. Walk `n` positions forward from there, wrapping each index.

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

A ring of capacity 3 holding one element after a wrap computes `0 - 1 = -1`, which indexes out of range instead of wrapping to the last cell.

## Pitfalls

- Assuming Go's `%` behaves like a mathematical modulus.
- Clamping the negative index to 0, which silently returns the wrong element.
- Testing only with a completely full buffer.
