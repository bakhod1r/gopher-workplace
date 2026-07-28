# Sliding Window Sums

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

A metrics smoother reports the sum over each window of k samples.

## Task

Implement `Sums(xs, k)`; invalid k → empty.

## Examples

```go
Sums([]int{1,2,3,4}, 2) // => [3 5 7]
Sums([]int{1,2,3}, 3)   // => [6]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Window count** | `len(xs)-k+1` windows. |
| 2 | **Rolling sum** | Add the entering, subtract the leaving. |
| 3 | **Guard k** | Empty when k>len or k<=0. |

## Hint

Compute the first window sum, then roll: `sum += xs[i+k-1] - xs[i-1]`.

## Validate

```bash
make verify
```
