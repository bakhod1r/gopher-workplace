# Clamp the Sub-slice End

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`xs[:n]` panics when `n > len(xs)`. A "take up to n" helper must clamp the end to
the length.

## Task

Fix the return between the markers in
[subsliceclamp.go](subsliceclamp.go) to clamp `n`.

## Examples

```go
Take([]int{1,2,3}, 10) // => [1 2 3] (no panic)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice bounds** | High index must be ≤ len. |
| 2 | **Clamp** | `min(n, len(xs))`. |
| 3 | **Panic safety** | Out-of-range slicing panics. |

## Hint

`if n > len(xs) { n = len(xs) }; return xs[:n]`.

## Validate

```bash
make verify
```
