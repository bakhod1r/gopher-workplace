# Percentile

**Level:** middle  
**Topic:** 03-generics

## Context

An SLO dashboard reports p95 latency. The definition must be the nearest-rank one so the number matches the alerting rules.

## Task

Implement the stub(s) in [percentilegen.go](percentilegen.go):

1. Implement `Percentile` using the nearest-rank method.
2. Clamp `p` into `[0, 100]`; return zero and `false` for an empty slice.
3. Leave the input untouched.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Percentile([]float64{1,2,3,4}, 50)
Output: 2, true
```

**Example 2:**

```
Input:  Percentile([]float64{1,2,3,4}, 100)
Output: 4, true
```

**Example 3:**

```
Input:  Percentile([]float64{}, 95)
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nearest rank** | `ceil(p/100 * n)` gives a 1-based rank; subtract one to index. |
| 2 | **Clamping twice** | Clamping `p` is not enough — `p = 0` yields rank 0, which must become 1. |
| 3 | **Definitions differ** | Interpolated percentiles give different answers; say which one you implement. |

## Hint

Rank is 1-based: convert with `rank-1`, and floor the rank at 1.

## Validate

```bash
make verify
```
