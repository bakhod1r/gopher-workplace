# Metrics Dashboard

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A dashboard downsamples a raw metric series into the points it can actually
draw: each pixel column shows the peak of a fixed-size window of samples.
Windows are independent, so each is reduced in its own goroutine.

## Task

Implement `PeakPerWindow` in [metricsdashboard.go](metricsdashboard.go) so that:

1. Return `nil` when `window <= 0`.
2. Split `samples` into consecutive windows of `window` values; the last window may be shorter.
3. Write the maximum of window `c` to `out[c]`, one goroutine per window.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PeakPerWindow([]int{1, 9, 3, 4}, 2)
Output: [9 4]
```

**Example 2:**

```
Input:  PeakPerWindow([]int{5, 2, 8}, 2)
Output: [5 8]
```

**Example 3:**

```
Input:  PeakPerWindow([]int{1}, 0)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Windows never overlap** | The boundary maths guarantees each window is non-empty and disjoint, so `part[0]` is always safe. |

## Hint

Seed the running peak from `part[0]` rather than `0`, or an all-negative series
will draw as a flat line at zero.

## Validate

```bash
make verify
```
