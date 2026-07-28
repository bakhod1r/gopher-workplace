# Simultaneous Assignment

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _call-by-value_

## Context

Sequential assignments overwrite `a` before its old value can reach `c`. Go's
multiple assignment `a, b, c = b, c, a` evaluates the whole right-hand side
first, so all moves happen simultaneously.

## Task

Fix [rotate3.go](rotate3.go) using simultaneous assignment.

Do **not** change the function signature or the tests.

## Examples

```go
RotateLeft(1, 2, 3) // => 2, 3, 1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Simultaneous assignment** | RHS is fully evaluated before any assign. |
| 2 | **Clobbering** | Sequential moves lose the original value. |
| 3 | **No temp needed** | Parallel assignment avoids temporaries. |

## Hint

Use one parallel assignment: `a, b, c = b, c, a`.

## Validate

```bash
make verify
```
