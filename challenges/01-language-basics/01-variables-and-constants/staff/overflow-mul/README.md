# Widen Before Multiply

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`int64(w * h)` computes `w*h` in `int32` first — it overflows and wraps, and the
widening conversion happens *after* the damage. Widen the operands first.

## Task

Fix the single line between the markers in [area.go](area.go) so large products
do not overflow.

## Examples

```go
Area(3, 4)             // => 12
Area(100000, 100000)   // => 10000000000
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Operand type wins** | `w*h` is `int32`; the result wraps before conversion. |
| 2 | **Widen first** | Convert each operand to `int64` before `*`. |
| 3 | **Silent overflow** | Signed wraparound gives a plausible wrong number. |

## Hint

`int64(w) * int64(h)`.

## Validate

```bash
make verify
```
