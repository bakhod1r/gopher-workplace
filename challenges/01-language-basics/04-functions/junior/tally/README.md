# Count and Sum Evens

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _variadic_

## Context

A single pass can compute several aggregates at once, returned together.

## Task

Implement `EvenStats` in [tally.go](tally.go): count the even arguments and sum them.

Do **not** change the function signature or the tests.

## Examples

```go
EvenStats()                 // => 0, 0
EvenStats(1, 2, 3, 4, 5, 6) // => 3, 12
EvenStats(2, 4)             // => 2, 6
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Variadic input** | `xs ...int` is a slice. |
| 2 | **Even test** | `n%2 == 0` identifies evens. |
| 3 | **Two accumulators** | Update count and total in the same loop. |

## Hint

Range over `xs`; when `x%2 == 0`, increment `count` and add to `total`.

## Validate

```bash
make verify
```
