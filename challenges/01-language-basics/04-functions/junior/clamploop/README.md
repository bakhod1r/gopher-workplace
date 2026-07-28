# Clamp Each Element

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

Mapping a transform over a slice into a fresh result keeps the caller's data
intact while a `for range` applies the per-element rule.

## Task

Implement `ClampAll` in [clamploop.go](clamploop.go).

Do **not** change the function signature or the tests.

## Examples

```go
ClampAll([]int{-5,3,20}, 0, 10) // => [0 3 10]
ClampAll(nil, 0, 10)            // => []
ClampAll([]int{5}, 0, 10)       // => [5]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Map over a slice** | Build a new slice by appending transformed values. |
| 2 | **Per-element clamp** | Reuse the two-bound clamp logic. |
| 3 | **Non-destructive** | Never write into the input. |

## Hint

Range `xs`; for each, append the clamped value (`if v<lo v=lo; if v>hi v=hi`) to `out`.

## Validate

```bash
make verify
```
