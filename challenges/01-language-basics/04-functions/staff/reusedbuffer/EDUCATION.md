# Returning reused buffers

## Intuition

Reusing a buffer with `buf[:0]` avoids allocations, but returning it aliases shared memory that the next call overwrites; hand out a copy when the caller keeps the result.

## Approach

1. Returning the internal `buf` shares it across calls, so later calls overwrite earlier results.
2. Return a copy: `append([]int(nil), buf...)`.

## Solution

```go
func Reader() func(vals ...int) []int {
	buf := make([]int, 0, 16)
	return func(vals ...int) []int {
		buf = buf[:0]
		buf = append(buf, vals...)
		return append([]int(nil), buf...)
	}
}
```

## Walkthrough

The reused buffer means `b`'s call mutates the array `a` still references. Returning a fresh copy each call keeps results stable.

## Pitfalls

- `buf[:0]` reuses the same array; returning it shares mutable memory.
- Copy with `append([]T(nil), buf...)` when the caller retains the slice.
