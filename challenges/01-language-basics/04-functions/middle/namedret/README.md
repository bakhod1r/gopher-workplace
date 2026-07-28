# Named Return Adjust

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _named-return_

## Context

Named results are pre-declared, zero-valued variables. A bare `return` sends
back their current values — handy for guard-and-default patterns.

## Task

Implement `SafeDiv` in [namedret.go](namedret.go) using named returns and a bare return for the zero-division case.

Do **not** change the function signature or the tests.

## Examples

```go
SafeDiv(10, 2) // => 5, true
SafeDiv(1, 0)  // => 0, false
SafeDiv(9, 3)  // => 3, true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Named results** | `(q int, ok bool)` are variables. |
| 2 | **Bare return** | `return` sends current named values. |
| 3 | **Guard clause** | Return the zero-value pair on b==0. |

## Hint

`if b == 0 { return }` (q=0, ok=false), else `q, ok = a/b, true; return`.

## Validate

```bash
make verify
```
