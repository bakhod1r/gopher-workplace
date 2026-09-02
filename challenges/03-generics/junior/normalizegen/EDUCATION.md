# Normalize

## Intuition

Restricting the set to floats is a design decision, not an accident: with `~int` in the set, `v/peak` would truncate every result to `0` or `1`.

## Approach

1. Scan once for the largest magnitude.
2. Return the elements unchanged when that peak is zero.
3. Otherwise append `v / peak` for each element.

## Solution

```go
func Normalize[T Float](s []T) []T {
	out := make([]T, 0, len(s))
	var peak T
	for _, v := range s {
		m := v
		if m < 0 {
			m = -m
		}
		if m > peak {
			peak = m
		}
	}
	for _, v := range s {
		if peak == 0 {
			out = append(out, v)
			continue
		}
		out = append(out, v/peak)
	}
	return out
}
```

## Walkthrough

`Normalize([]float64{-4, 2})` finds a peak of `4`, then returns `-1` and `0.5`, preserving sign.

## Pitfalls

- Allowing `~int` in the constraint, which truncates every quotient.
- Dividing by the maximum value rather than the maximum magnitude, which flips signs for negative peaks.
- Dividing by zero on an all-zero input, producing NaN.
