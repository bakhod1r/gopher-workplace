# Filter Positives

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Filtering builds a new slice with `append`, keeping only elements that pass a
test.

## Task

Implement `Positives(xs)` returning a new slice of the positive elements. Empty
result must be non-nil.

## Examples

```go
Positives([]int{1,-2,3,0,4}) // => [1 3 4]
Positives([]int{-1,-2})      // => []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **append** | Grows a slice, returning the new header. |
| 2 | **Non-nil empty** | Start with `[]int{}` (or make) to avoid nil. |
| 3 | **Predicate** | Keep when `x > 0`. |

## Hint

`out := []int{}; for _, x := range xs { if x > 0 { out = append(out, x) } }`.

## Validate

```bash
make verify
```
