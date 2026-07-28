# Absolute Value

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

A single `if` flips the sign of negatives — the simplest branch there is.

## Task

Implement `Abs` in [absif.go](absif.go).

Do **not** change the function signature or the tests.

## Examples

```go
Abs(-5) // => 5
Abs(5)  // => 5
Abs(0)  // => 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Single if** | Negate when `n < 0`. |
| 2 | **Return value** | One int out. |
| 3 | **Zero** | Already non-negative. |

## Hint

`if n < 0 { return -n }; return n`.

## Validate

```bash
make verify
```
