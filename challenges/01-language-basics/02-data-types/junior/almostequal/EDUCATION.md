# Comparing floats with a tolerance

## Intuition

Binary floats can't represent most decimals exactly, so `==` is the wrong tool.
Compare within a small tolerance instead:

```go
math.Abs(a-b) <= eps
```

## Approach

1. Import the math package.
2. Compute the absolute difference math.Abs(a-b).
3. Return true when that difference is strictly less than the epsilon 1e-9.

## Solution

```go
import "math"

func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}
```

## Walkthrough

AlmostEqual(0.1+0.2, 0.3): a-b = 4.4e-17, math.Abs = 4.4e-17, and 4.4e-17 < 1e-9 -> true.

## Pitfalls

- If `a` or `b` is NaN, `a-b` is NaN and the comparison is false — NaN is never
  "almost equal".
- Absolute tolerance suits values near the same magnitude; relative tolerance
  scales across magnitudes.
- The right `eps` depends on the operations performed, not a universal constant.
