# Saturating floats, and NaN

## Intuition

Clamp with two comparisons — but NaN breaks the naive version, because every
comparison with NaN is false, so a `NaN` would slip through unchanged:

```go
if math.IsNaN(x) { return 0 }
if x < 0 { return 0 }
if x > 1 { return 1 }
return x
```

## Approach

1. If x is NaN (math.IsNaN) return 0. 2. If x<0 return 0. 3. If x>1 return 1. 4. Otherwise return x.

## Solution

```go
import "math"

func Saturate(x float64) float64 {
	if math.IsNaN(x) {
		return 0
	}
	if x < 0 {
		return 0
	}
	if x > 1 {
		return 1
	}
	return x
}
```

## Walkthrough

Saturate(2): not NaN, not <0, but 2>1 -> return 1.

## Pitfalls

- Order matters: test NaN before the range checks.
- `min`/`max` builtins (Go 1.21+) also propagate NaN, so they don't fix this.
