# Sliding Windows In One Pass

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Smoothing a noisy latency series is the standard way to see a trend, and it is also the standard way to write an accidental O(n·k) loop: re-summing each window looks natural and costs a hundred times more than it needs to. The fix is to remember the previous window's sum.

## Task

Implement both functions in [movingavg.go](movingavg.go):

1. `Window` returns the mean of every `n`-sample sliding window, `len(samples)-n+1` values in total.
2. Do it in one pass — add the entering sample, subtract the leaving one — not by re-summing each window.
3. A non-positive `n`, or one larger than the input, gives an empty non-nil slice; `Smoothest` returns the index of the smallest mean, `-1` when there are none.

## Examples

**Example 1:**

```
Input:  Window([1 2 3 4], 2)
Output: [1.5 2.5 3.5]
```

**Example 2:**

```
Input:  Window([1 2 3 4], 4)
Output: [2.5]
```

**Example 3:**

```
Input:  Smoothest([3 1 1 5])
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Incremental beats recomputing** | The new window differs from the old one by two samples, not `n`. |
| 2 | **The window count is `m-n+1`** | Off by one here and the last window silently disappears. |
| 3 | **Smoothing hides spikes** | A wide window can average away the exact event you are hunting. |

## Topics used again

Sliding windows, running sums, slice bounds, guards.

## Hint

Sum the first window before the loop; each later step is one add and one subtract.

## Validate

```bash
make verify
```
