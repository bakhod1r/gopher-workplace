# Sliding Window Maximum

**Level:** middle  
**Topic:** 03-generics

## Context

A monitoring chart shows the peak of a rolling window over a long series. Re-scanning each window is quadratic and too slow.

## Task

Implement the stub(s) in [windowmaxgen.go](windowmaxgen.go):

1. Implement `WindowMax`, returning the maximum of each consecutive window of `n`.
2. Return an empty result when `n <= 0` or `n` exceeds the length.
3. Run in linear time — do not scan each window.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  WindowMax([]int{1,3,2}, 2)
Output: [3 3]
```

**Example 2:**

```
Input:  WindowMax([]int{1,2,3}, 3)
Output: [3]
```

**Example 3:**

```
Input:  WindowMax([]int{1}, 2)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Monotonic deque** | Indexes are kept in decreasing value order, so the front is always the window maximum. |
| 2 | **Two eviction rules** | Drop indexes that fell out of the window, and drop those that can never win again. |
| 3 | **Amortised linear** | Each index is appended and removed at most once. |

## Hint

Store indexes, not values — you need them to know when a candidate leaves the window.

## Validate

```bash
make verify
```
