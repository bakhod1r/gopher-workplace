# Variadic Average

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _variadic_

## Context

Averaging combines a variadic parameter with a divide-by-count guard — the
empty case must not divide by zero.

## Task

Implement `Average` in [averagef.go](averagef.go). Return `0, false` for no arguments.

Do **not** change the function signature or the tests.

## Examples

```go
Average()          // => 0, false
Average(2, 4)      // => 3, true
Average(1, 2, 3, 4) // => 2.5, true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Variadic + guard** | Check `len(nums) == 0` before dividing. |
| 2 | **Integer vs float division** | Divide by `float64(len(nums))`. |
| 3 | **Comma-ok** | Report validity with the bool. |

## Hint

Sum the slice, then return `sum / float64(len(nums)), true`; guard the empty case first.

## Validate

```bash
make verify
```
