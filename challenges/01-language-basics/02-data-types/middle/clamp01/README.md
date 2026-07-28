# Saturate to [0,1]

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Colours and probabilities live in [0,1]. Saturating clamps out-of-range values —
and NaN must be pinned to a safe default.

## Task

Implement `Saturate(x)` limiting to [0,1]; NaN → 0.

## Examples

```go
Saturate(0.5) // => 0.5
Saturate(-1)  // => 0
Saturate(2)   // => 1
Saturate(NaN) // => 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Range clamp** | Two comparisons pin the value. |
| 2 | **NaN handling** | NaN fails `<` and `>`; test it explicitly. |
| 3 | **Ordering** | Check NaN (or low bound) first. |

## Hint

Handle NaN first (`math.IsNaN`), then `if x<0 {0}; if x>1 {1}; x`.

## Validate

```bash
make verify
```
