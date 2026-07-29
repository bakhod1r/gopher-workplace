# Output parameters via pointers

## Intuition

Writing results through caller-supplied pointers is an alternative to multiple return values, common when interfacing with fixed signatures.

## Approach

1. Seed both through the pointers: `*min, *max = xs[0], xs[0]`.
2. Scan the rest, updating `*min`/`*max` when a smaller/larger value appears.
3. Results land in the caller's variables.

## Solution

```go
func MinMax(xs []int, min, max *int) {
	*min, *max = xs[0], xs[0]
	for _, v := range xs {
		if v < *min {
			*min = v
		}
		if v > *max {
			*max = v
		}
	}
}
```

## Walkthrough

For `[3, -1, 7]`: seed `lo = hi = 3`; `-1` lowers `lo` to `-1`; `7` raises `hi` to `7`.

## Pitfalls

- Dereference to write the results: `*min = v`.
- Go usually prefers multiple returns, but pointer outputs appear in stdlib.
