# Sort Mutates Caller

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`out := xs` copies the slice **header**, not the data, so `sort.Ints(out)` sorts
the caller's backing array. "SortedCopy" corrupts its input.

## Task

Fix the copy between the markers in [sortmutates.go](sortmutates.go) to duplicate
the data first.

## Examples

```go
xs := []int{3,1,2}; SortedCopy(xs) // [1 2 3]; xs stays [3 1 2]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Header vs data** | `out := xs` shares the array. |
| 2 | **sort in place** | `sort.Ints` mutates. |
| 3 | **Copy first** | `append([]int{}, xs...)`. |

## Hint

`out := append([]int{}, xs...)`.

## Validate

```bash
make verify
```
