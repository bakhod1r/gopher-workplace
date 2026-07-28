# Compose Two Funcs

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Function composition builds a new function by chaining two — the returned
closure captures both `f` and `g`.

## Task

Implement `Compose` in [compose.go](compose.go) so the result computes `f(g(x))`.

Do **not** change the function signature or the tests.

## Examples

```go
Compose(inc, dbl)(3) // => 7  (inc(dbl(3)))
Compose(dbl, inc)(3) // => 8  (dbl(inc(3)))
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Capture two funcs** | Closure remembers `f` and `g`. |
| 2 | **Application order** | `f(g(x))`: g first, then f. |
| 3 | **Return a func** | Result type `func(int) int`. |

## Hint

`return func(x int) int { return f(g(x)) }`.

## Validate

```bash
make verify
```
