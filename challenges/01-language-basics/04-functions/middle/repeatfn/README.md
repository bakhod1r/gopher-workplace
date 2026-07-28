# Apply N Times

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Capturing both a function and a count produces a closure that self-iterates —
function exponentiation.

## Task

Implement `Repeat` in [repeatfn.go](repeatfn.go).

Do **not** change the function signature or the tests.

## Examples

```go
Repeat(inc, 3)(0) // => 3
Repeat(inc, 0)(9) // => 9
Repeat(dbl, 2)(1) // => 4
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Capture f and n** | Both are remembered by the closure. |
| 2 | **Loop application** | Apply f n times inside. |
| 3 | **Identity at 0** | No applications returns the input. |

## Hint

`return func(x int) int { for i := 0; i < n; i++ { x = f(x) }; return x }`.

## Validate

```bash
make verify
```
