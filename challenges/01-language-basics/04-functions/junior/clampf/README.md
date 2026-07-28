# Clamp to Range

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _functions-basics_

## Context

Clamping keeps a value inside bounds — below the floor it becomes the floor,
above the ceiling it becomes the ceiling.

## Task

Implement `Clamp` in [clampf.go](clampf.go) with an inclusive range.

Do **not** change the function signature or the tests.

## Examples

```go
Clamp(5, 0, 10)   // => 5
Clamp(-3, 0, 10)  // => 0
Clamp(99, 0, 10)  // => 10
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Conditional logic** | Two `if` checks cover both bounds. |
| 2 | **Inclusive bounds** | Exactly `hi` stays `hi`, not clamped further. |
| 3 | **Single return** | One `int` out. |

## Hint

If `v < lo` return `lo`; if `v > hi` return `hi`; otherwise return `v`.

## Validate

```bash
make verify
```
