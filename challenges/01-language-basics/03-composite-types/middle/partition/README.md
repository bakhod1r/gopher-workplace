# Partition by Predicate

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Splitting a stream into two buckets (pass/fail, even/odd) in one pass.

## Task

Implement `Partition(xs)` → (evens, odds), order preserved.

## Examples

```go
Partition([]int{1,2,3,4,5,6}) // => [2 4 6], [1 3 5]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two accumulators** | One slice per bucket. |
| 2 | **Predicate** | `x%2 == 0` selects the bucket. |
| 3 | **Multiple returns** | Return both slices. |

## Hint

`for _, x := range xs { if x%2==0 { evens=append(evens,x) } else { odds=append(odds,x) } }`.

## Validate

```bash
make verify
```
