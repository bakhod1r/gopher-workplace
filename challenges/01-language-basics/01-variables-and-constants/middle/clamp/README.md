# Clamp to Range

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Clamping needs valid bounds. When `lo > hi` you normalize them in a local
scope — a good place to practice `if`-init variables and block scope without
shadowing the parameters.

## Task

In [clamp.go](clamp.go):

1. If `lo > hi`, swap them (locally).
2. Return `v` limited to `[lo, hi]`.

## Examples

```go
Clamp(5, 0, 10)  // => 5
Clamp(-3, 0, 10) // => 0
Clamp(99, 0, 10) // => 10
Clamp(5, 10, 0)  // => 5  (bounds swapped)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Block scope** | Variables declared in a block vanish at its end. |
| 2 | **Shadowing risk** | `lo, hi :=` inside a block would shadow, not update, the params. |
| 3 | **Multiple assignment** | `lo, hi = hi, lo` swaps without a temp. |

## Hint

Reassign with `lo, hi = hi, lo` (use `=`, not `:=`, or you shadow the params).

## Validate

```bash
make verify
```
