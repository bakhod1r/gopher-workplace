# Saturating floats, and NaN

## The idea

Clamp with two comparisons — but NaN breaks the naive version, because every
comparison with NaN is false, so a `NaN` would slip through unchanged:

```go
if math.IsNaN(x) { return 0 }
if x < 0 { return 0 }
if x > 1 { return 1 }
return x
```

## Why it matters

Saturation guards colour channels, alpha, and probabilities. A leaked NaN
poisons downstream math, so mapping it to a safe default is essential.

## Why NaN slips through

`NaN < 0` and `NaN > 1` are both false, so a comparison-only clamp returns the
NaN untouched. You must test `math.IsNaN` (or order the checks so NaN can't
escape).

## Watch out

- Order matters: test NaN before the range checks.
- `min`/`max` builtins (Go 1.21+) also propagate NaN, so they don't fix this.
