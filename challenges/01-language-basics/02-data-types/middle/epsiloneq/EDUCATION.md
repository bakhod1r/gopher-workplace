# Comparing floats

## Intuition

Binary floats cannot represent most decimals exactly, so `0.1+0.2` is
`0.30000000000000004`. Compare with a tolerance:

```go
math.Abs(a-b) <= eps
```

## Approach

1. If either a or b is NaN, return false. 2. Otherwise return math.Abs(a-b) <= eps.

## Solution

```go
import "math"

func Equal(a, b, eps float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) {
		return false
	}
	return math.Abs(a-b) <= eps
}
```

## Walkthrough

Equal(0.1+0.2,0.3,1e-9): |0.30000000000000004-0.3| ~ 5.5e-17 <= 1e-9 -> true.

## Pitfalls

- If `a` or `b` is NaN, `a-b` is NaN and `NaN <= eps` is false — so NaN is
  correctly reported unequal, for free.
- Absolute tolerance fails across very different magnitudes; relative tolerance
  (`|a-b| <= eps*max(|a|,|b|)`) scales better.
- `eps == 0` reduces to exact equality (fine for identical values).
