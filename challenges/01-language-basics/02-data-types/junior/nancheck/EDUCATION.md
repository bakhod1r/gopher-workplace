# NaN and infinity

## The idea

IEEE-754 floats include `NaN` and `±Inf`. Detect them with the `math` helpers,
never with `==`:

```go
!math.IsNaN(x) && !math.IsInf(x, 0) // finite?
```

`math.IsInf(x, 0)` tests either sign; pass `1` or `-1` for a specific one.

## Why it matters

`NaN` is famously **not equal to itself**, so `x == math.NaN()` is always false —
you cannot detect NaN by comparison. Unchecked, a single NaN propagates through
every subsequent calculation and silently corrupts results.

## Watch out

- `x != x` is true exactly when `x` is NaN — a valid but cryptic idiom; prefer
  `math.IsNaN`.
- Sorting or comparing slices with NaN gives inconsistent orderings.
- `math.Inf(1)` is +Inf, `math.Inf(-1)` is -Inf.

## Try it yourself

```go
math.NaN() == math.NaN() // false
math.IsNaN(math.NaN())   // true
1.0 / 0.0                // +Inf (untyped constant division is a compile error; use vars)
```
