# Mapping a pure transform

## Intuition

Producing a new slice from an old one, one transformed element at a time, is the functional map pattern in imperative Go.

## Approach

1. Range the slice.
2. Clamp each value to [lo, hi] and append.

## Solution

```go
func ClampAll(xs []int, lo, hi int) []int {
	out := make([]int, 0, len(xs))
	for _, v := range xs {
		if v < lo {
			v = lo
		}
		if v > hi {
			v = hi
		}
		out = append(out, v)
	}
	return out
}
```

## Walkthrough

For `[-1 5 99]` with bounds 0..10: -1→0, 5 stays, 99→10.

## Pitfalls

- Preallocating with `make([]int, 0, len(xs))` avoids regrowth but isn't required.
- Clamp each value before appending; don't mutate `xs[i]`.
