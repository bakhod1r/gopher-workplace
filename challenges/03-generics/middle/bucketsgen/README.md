# Bucket Counts

**Level:** middle  
**Topic:** 03-generics

## Context

A histogram groups latencies into configured buckets, matching the boundaries the alerting system already uses.

## Task

Implement the stub(s) in [bucketsgen.go](bucketsgen.go):

1. Implement `Buckets`, returning one count per bucket.
2. Bucket `i` holds values in `[edges[i-1], edges[i])`; the last bucket holds everything from the final edge upward.
3. With no edges every element lands in the single bucket.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Buckets([]int{1,5,9}, []int{5})
Output: []int{1, 2}
```

**Example 2:**

```
Input:  Buckets([]int{1}, []int{})
Output: []int{1}
```

**Example 3:**

```
Input:  Buckets([]int{}, []int{5})
Output: []int{0, 0}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Half-open buckets** | `[lo, hi)` is what makes an element land in exactly one bucket. |
| 2 | **One more count than edges** | `n` edges cut the line into `n+1` regions. |
| 3 | **Only `<` is needed** | Expressing the test with `<` alone keeps the constraint minimal. |

## Hint

`n` edges give `n+1` buckets; advance while the value is not below the current edge.

## Validate

```bash
make verify
```
