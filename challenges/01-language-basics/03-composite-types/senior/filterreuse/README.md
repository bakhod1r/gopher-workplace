# Filter Corrupts Input

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`out := xs[:0]` reuses `xs`'s backing array. Appending the kept elements
overwrites `xs` from the front, so the input is corrupted mid-iteration.

## Task

Fix the initialization between the markers in
[filterreuse.go](filterreuse.go) to use a fresh slice.

## Examples

```go
xs := []int{1,2,3,4}; Evens(xs) // [2 4]; xs stays [1 2 3 4]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **xs[:0] aliases** | Same array, length 0. |
| 2 | **In-place filter** | Valid only when overwriting is intended. |
| 3 | **Fresh slice** | `[]int{}` or make for independence. |

## Hint

`out := []int{}` (don't reuse `xs`'s array).

## Validate

```bash
make verify
```
