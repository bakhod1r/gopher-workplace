# Slices share backing arrays

## Intuition

A slice is a header (pointer, length, capacity) over a backing array. Copying the
header (`b := a`) shares the array. An independent copy allocates a new array:

```go
out := make([]int, len(xs))
copy(out, xs)
```

## Approach

1. Allocate make([]int, len(xs)) — a fresh backing array; for nil xs this is a non-nil empty slice.
2. copy(result, xs) copies element-by-element.
3. Return result, which shares nothing with xs.

## Solution

```go
func Clone(xs []int) []int {
	result := make([]int, len(xs))
	copy(result, xs)
	return result
}
```

## Walkthrough

Clone([1,2,3]): make len-3 slice, copy in 1,2,3. Writing result[0]=99 touches only the new array; xs[0] stays 1.

## Pitfalls

- `copy` copies `min(len(dst), len(src))` elements.
- `make([]int, len(xs))` is non-nil even when `len` is 0.
- `append` to a shared slice may alias depending on capacity.
