# Integer constant division

## The idea

When both operands are integers, `/` is **integer division**: it truncates
toward zero, discarding any remainder.

```go
const TargetFPS = 60
const budgetMicros = 1_000_000 / TargetFPS // 16666, not 16666.67
```

The result is exact and computed at compile time, but the fraction is gone.

## Why it matters

For a per-frame time budget, 16666 µs is the honest floor: a frame is "over
budget" only when it exceeds it. Silently rounding up would hide real overruns.

## Watch out

- `1000000 / 60` is 16666. If you want rounding, do it explicitly
  (`(x + d/2) / d`), do not sprinkle magic `+1`s.
- Mixing a float makes it float division: `1000000.0 / 60` is 16666.67.
- Digit separators (`1_000_000`) are cosmetic and improve readability.

## Try it yourself

```go
const perSecond = 1_000_000_000 / 60 // ns per frame at 60fps
const half = 7 / 2                    // 3
const halfF = 7.0 / 2                 // 3.5
```
