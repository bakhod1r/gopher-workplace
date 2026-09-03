# The Pause Budget

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Go's collector is concurrent, and its stop-the-world pauses are sub-millisecond by design — but "sub-millisecond" only means something next to your latency budget. Five milliseconds of pause per second is invisible in a batch job and fatal for a 10ms p99. The budget has two halves: how much total time, and how long any single stop is allowed to be.

## Task

Implement the four functions in [gcpausesum.go](gcpausesum.go):

1. `Total` sums the pauses, ignoring negative entries; `Worst` returns the longest and its index, `-1` when there are none.
2. `FractionOf` divides the total by a wall-clock window, returning 0 for a non-positive window.
3. `WithinBudget` checks both the total fraction and the single-pause limit, inclusively.

## Examples

**Example 1:**

```
Input:  FractionOf([5000000], 1000000000)
Output: 0.005
```

**Example 2:**

```
Input:  Worst([100 500 200])
Output: 500, 1
```

**Example 3:**

```
Input:  WithinBudget([1000000 2000000], 1e9, 0.01, 1500000)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two limits, not one** | Total pause time and worst single pause fail in different ways. |
| 2 | **A pause is relative to the window** | The same 5ms is noise in a second and everything in a 10ms request. |
| 3 | **Inclusive limits** | "At most 1ms" has to accept exactly 1ms, or every budget is off by an epsilon. |

## Topics used again

Sums with guards, maximum with earliest index, float division, boolean combination.

## Hint

`WithinBudget` is `FractionOf` and `Worst` combined — do not recompute either.

## Validate

```bash
make verify
```
