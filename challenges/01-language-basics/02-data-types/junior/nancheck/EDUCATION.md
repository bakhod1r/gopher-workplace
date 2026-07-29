# NaN and infinity

## Intuition

IEEE-754 floats include `NaN` and `±Inf`. Detect them with the `math` helpers,
never with `==`:

```go
!math.IsNaN(x) && !math.IsInf(x, 0) // finite?
```

`math.IsInf(x, 0)` tests either sign; pass `1` or `-1` for a specific one.

## Approach

1. Reject NaN with math.IsNaN(x).
2. Reject +/-Inf with math.IsInf(x, 0) (0 means either sign).
3. Return true only when neither holds.

## Solution

```go
import "math"

func Finite(x float64) bool {
	return !math.IsNaN(x) && !math.IsInf(x, 0)
}
```

## Walkthrough

Finite(math.Inf(1)): IsNaN=false, IsInf(x,0)=true, so !false && !true = false.

## Pitfalls

- `x != x` is true exactly when `x` is NaN — a valid but cryptic idiom; prefer
  `math.IsNaN`.
- Sorting or comparing slices with NaN gives inconsistent orderings.
- `math.Inf(1)` is +Inf, `math.Inf(-1)` is -Inf.
