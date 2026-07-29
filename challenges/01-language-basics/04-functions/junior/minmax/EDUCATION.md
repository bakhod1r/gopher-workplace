# Multiple return values

## Intuition

A function's result list can hold several typed values, returned as a tuple and destructured at the call site.

## Approach

1. Seed both `min` and `max` from `xs[0]`.
2. Scan the rest, lowering `min` and raising `max`.
3. Return the named results.

## Solution

```go
func MinMax(xs []int) (min, max int) {
	min, max = xs[0], xs[0]
	for _, v := range xs[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}
```

## Walkthrough

For `[3 -1 7]`: seed 3,3; `-1` drops min to -1; `7` raises max to 7.

## Pitfalls

- Seeding `min` to `0` breaks on all-positive slices.
- Named results are zero-initialised; you still must set them.
