# Setup That Moved Into The Loop

## Intuition

`make([]byte, 0, 64)` and `e.buf[:0]` both give you an empty buffer. Only one of them gives you the buffer you already had.

## Approach

1. Reset the existing buffer's length instead of allocating a new one.

## Solution

```go
func (e *Encoder) Encode(names, values []string) []byte {
	e.buf = e.buf[:0]
	...
}
```

## Walkthrough

`e.buf[:0]` on the zero `Encoder` is a nil slice resliced to zero length, which is legal — so the first call still allocates through `append`, and every later call reuses that array. The fixed capacity in the bug also caps nothing: `append` grows past 64 anyway, adding a second allocation for larger records.

## Pitfalls

- Any `make` inside a function documented as reusing a buffer.
- `e.buf = nil` as the reset, which is the same bug spelled differently.
- Assuming a capacity hint inside the hot path helps; the allocation is the cost, not the growth.
