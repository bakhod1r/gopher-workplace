# Merge Sorted Slices

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

The merge step of merge sort: combine two sorted runs into one, in O(n).

## Task

Implement `Merge(a, b)` (ascending, duplicates kept).

## Examples

```go
Merge([]int{1,3,5}, []int{2,4,6}) // => [1 2 3 4 5 6]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two pointers** | Advance the smaller side. |
| 2 | **Drain remainder** | Append the leftover tail. |
| 3 | **Stable dupes** | Keep equal elements. |

## Hint

Walk `i,j`; append the smaller; then append whichever slice remains.

## Validate

```bash
make verify
```
