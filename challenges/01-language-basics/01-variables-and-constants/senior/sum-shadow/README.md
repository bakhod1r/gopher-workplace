# Shadowed Accumulator

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

The loop looks like it accumulates into `total`, but `SumPositive` always
returns 0. A `:=` inside the `if` block creates a new `total` that dies each
iteration — the outer one never changes.

## Task

Fix the single line between the markers in [accumulate.go](accumulate.go) so the
outer `total` accumulates. Keep the signature and the surrounding code.

## Examples

```go
SumPositive([]int{1,2,3})       // => 6
SumPositive([]int{-1,5,-2,4})   // => 9
SumPositive(nil)                // => 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shadowing** | `:=` in an inner block declares a fresh variable. |
| 2 | **Block scope** | The inner `total` is discarded at the block's end. |
| 3 | **`=` vs `:=`** | Use `=` to update the existing variable. |

## Hint

`total := total + x` shadows. Change `:=` to `+=` (and drop the `_ = total`).

## Validate

```bash
make verify
```
