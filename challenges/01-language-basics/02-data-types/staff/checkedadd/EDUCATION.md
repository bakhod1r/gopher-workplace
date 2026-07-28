# Detecting integer overflow before it happens

## The idea

You cannot test overflow by computing `a+b` — that already wrapped. Rearrange so
the check itself never overflows: for `b > 0`, the sum overflows exactly when

```go
a > math.MaxInt64 - b
```

and symmetrically for `b < 0`, when `a < math.MinInt64 - b`.

## Why it matters

Signed overflow in Go wraps silently (no panic). Counters, financial sums, and
size calculations that can exceed int64 need an explicit pre-check. `a > MaxInt64`
is a no-op guard — a real bug that lets the wrap through.

## Watch out

- Never test overflow using the overflowed result.
- `MaxInt64 - b` is safe because `b > 0`; `MinInt64 - b` is safe because `b < 0`.
- `math/bits.Add64` and Go 1.21's overflow-aware helpers exist for this.
