# copy is bounded by length

## Intuition

`copy` moves `min(len(dst), len(src))` elements; a zero-length destination copies nothing regardless of capacity.

## Approach

1. `copy` transfers only `min(len(dst), len(src))` elements.
2. The bug makes `dst` length 0, so nothing copies.
3. Allocate `make([]int, len(xs))`.

## Solution

```go
func Clone(xs []int) []int {
	dst := make([]int, len(xs))
	copy(dst, xs)
	return dst
}
```

## Walkthrough

With a zero-length `dst`, `copy` writes nothing and returns an empty clone. Sizing `dst` to `len(xs)` copies all elements.

## Pitfalls

- `copy` uses LENGTH, not capacity, of both slices.
- For a clone use `make([]T, len(src))` (or `append([]T(nil), src...)`).
