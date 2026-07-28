# Recover Then Re-panic

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

A recover handler that absorbs EVERY panic hides bugs. It should absorb only
the known sentinel and re-panic anything else so real failures still surface.

## Task

Fix the recover handler in [repanic.go](repanic.go) to re-panic non-sentinel values.

Do **not** change the function signature or the tests.

## Examples

```go
Run(func(){})                     // => true
Run(func(){ panic(ErrSentinel) }) // => false
Run(func(){ panic("x") })         // re-panics "x"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Selective recovery** | Absorb only known panic values. |
| 2 | **Re-panic** | `panic(r)` rethrows the unhandled case. |
| 3 | **Don't mask bugs** | Blanket recover swallows real crashes. |

## Hint

Inside the handler: `if r == ErrSentinel { normal = false } else { panic(r) }`.

## Validate

```bash
make verify
```
