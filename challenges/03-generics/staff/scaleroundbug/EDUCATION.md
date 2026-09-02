# Scaling That Divides First

## Intuition

Integer division is not associative with multiplication. `(v / 100) * pct` discards up to 99 units of `v` before scaling, so the error is multiplied by `pct` — and any `v` below 100 collapses to zero.

## Approach

1. Multiply the value by the percentage first, at the full width of `T`.
2. Add or subtract half the divisor to round away from zero.
3. Divide by 100 once, at the end.

## Solution

```go
func ScaleAll[T Integer](vals []T, pct T) []T {
	out := make([]T, len(vals))
	for i, v := range vals {
		n := v * pct
		if n >= 0 {
			out[i] = (n + 50) / 100
		} else {
			out[i] = (n - 50) / 100
		}
	}
	return out
}
```

## Walkthrough

`ScaleAll([]int{7}, 300)` computes `7 / 100`, which is `0`, and then `0 * 300` — losing the entire line item.

## Pitfalls

- Rounding with `(n + 50) / 100` for negative `n`, which rounds towards positive infinity.
- Reordering into `v * (pct / 100)`, which is the same defect wearing a different hat.
