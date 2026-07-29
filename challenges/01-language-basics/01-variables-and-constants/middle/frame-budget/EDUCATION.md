# Integer constant division

## Intuition

When both operands are integers, `/` is **integer division**: it truncates
toward zero, discarding any remainder.

```go
const TargetFPS = 60
const budgetMicros = 1_000_000 / TargetFPS // 16666, not 16666.67
```

The result is exact and computed at compile time, but the fraction is gone.

## Approach

1. `1_000_000 / TargetFPS` with integer division gives the budget.
2. `OverBudget` compares against it.

## Solution

```go
const TargetFPS = 60

func FrameBudgetMicros() int {
	return 1_000_000 / TargetFPS
}

func OverBudget(us int) bool {
	return us > FrameBudgetMicros()
}
```

## Walkthrough

At 60 FPS the budget is 16666 µs; a 20000 µs frame is over.

## Pitfalls

- `1000000 / 60` is 16666. If you want rounding, do it explicitly
  (`(x + d/2) / d`), do not sprinkle magic `+1`s.
- Mixing a float makes it float division: `1000000.0 / 60` is 16666.67.
- Digit separators (`1_000_000`) are cosmetic and improve readability.
