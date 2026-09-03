# Concatenating Without Regrowing

## Intuition

Concatenation has a known length before it starts. Allocating that length once gives you both the performance win and the aliasing safety for free.

## Approach

1. `make([]int, 0, len(a)+len(b))`.
2. Append `a...`, then `b...`.

## Solution

```go
func Merge(a, b []int) []int {
	out := make([]int, 0, len(a)+len(b))
	out = append(out, a...)
	return append(out, b...)
}
```

## Walkthrough

Because `out` starts with exactly the needed capacity, neither append reallocates, and because it is a fresh array, no write can reach `a` or `b`.

## Pitfalls

- `append(a, b...)`, which may write into `a`'s spare capacity.
- `make([]int, len(a)+len(b))` plus `append`, which leaves zeros in front.
- Two separate `copy` calls into a slice created with the wrong length.
