# Per Operation, Not Per Iteration

## Intuition

Truncating division always rounds down, so add half the divisor first and the boundary lands where you want it.

## Approach

1. Return `0` for a non-positive `iters`.
2. Return `(totalOps + iters/2) / iters`.

## Solution

```go
func OpsPerIter(totalOps int, iters int) int {
	if iters <= 0 {
		return 0
	}
	return (totalOps + iters/2) / iters
}
```

## Walkthrough

`OpsPerIter(10, 4)` computes `(10 + 2) / 4 = 3`, while `OpsPerIter(9, 4)` computes `(9 + 2) / 4 = 2` — exactly the half-up boundary.

## Pitfalls

- Plain `totalOps / iters`, which reports `2` for `2.5`.
- Converting to `float64` and back, which reintroduces `math.Round`'s banker's-rounding confusion.
- Dividing before the guard.
