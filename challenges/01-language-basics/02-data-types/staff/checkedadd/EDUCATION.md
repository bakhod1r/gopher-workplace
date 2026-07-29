# Detecting integer overflow before it happens

## Intuition

You cannot test overflow by computing `a+b` — that already wrapped. Rearrange so
the check itself never overflows: for `b > 0`, the sum overflows exactly when

```go
a > math.MaxInt64 - b
```

and symmetrically for `b < 0`, when `a < math.MinInt64 - b`.

## Approach

1. Bug: `a > math.MaxInt64` is never true (a is int64), so positive overflow is missed.
2. Overflow when a+b > MaxInt64, i.e. a > MaxInt64 - b (no overflow in the check itself since b>0).
3. Fix: if b > 0 && a > math.MaxInt64-b.

## Solution

```go
import "math"

func Add(a, b int64) (int64, bool) {
	if b > 0 && a > math.MaxInt64-b {
		return 0, false
	}
	if b < 0 && a < math.MinInt64-b {
		return 0, false
	}
	return a + b, true
}
```

## Walkthrough

MaxInt64 + 1: b=1>0, a=MaxInt64 > MaxInt64-1 -> true -> (0,false).

## Pitfalls

- Never test overflow using the overflowed result.
- `MaxInt64 - b` is safe because `b > 0`; `MinInt64 - b` is safe because `b < 0`.
- `math/bits.Add64` and Go 1.21's overflow-aware helpers exist for this.
