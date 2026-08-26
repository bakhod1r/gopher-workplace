# Aggregation Methods

## Intuition

Methods like `Sum`, `Average`, `Min`, `Max` compute aggregate values from
collections. They're read-only, so value receivers work perfectly.

## Approach

1. Initialize `total := 0.0`.
2. Range over `s.Values`, add each.
3. Return `total`.

## Solution

```go
func (s Stats) Sum() float64 {
	total := 0.0
	for _, v := range s.Values {
		total += v
	}
	return total
}
```

## Walkthrough

For `[]float64{1, 2, 3}`:
- total = 0 + 1 = 1.
- total = 1 + 2 = 3.
- total = 3 + 3 = 6.

## Pitfalls

- Ranging over a nil slice is safe — the loop body never executes.
- Floating-point precision: `0.1 + 0.2 + 0.3` may not equal `0.6` exactly.
  The test allows a tiny epsilon.
