# Reading A Goroutine Dump

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

`SIGQUIT` a hung Go program and you get every goroutine's state and stack. The diagnosis is almost always in the histogram: fifty thousand goroutines blocked in `chan receive` at one line is a leak; a thousand `runnable` is a program with more work than cores. Counting is the whole technique.

## Task

Implement the three functions in [goroutinestates.go](goroutinestates.go):

1. `Count` tallies goroutines by state, counting an empty state as `"unknown"`.
2. `Blocked` counts everything that is neither `"running"` nor `"runnable"`.
3. `LeakSuspects` returns the wait sites with at least `threshold` blocked goroutines in the same state, ordered by count descending then state then site; a threshold below 1 means 1.

## Examples

**Example 1:**

```
Input:  Count([{running x} {chan receive y} {chan receive y}])
Output: {running:1 chan receive:2}
```

**Example 2:**

```
Input:  Blocked([{running a} {runnable b} {chan receive c} {select d} {IO wait e}])
Output: 3
```

**Example 3:**

```
Input:  150 blocked in main.consume, 50 in main.dispatch; LeakSuspects(_, 100)
Output: [main.consume]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The histogram is the diagnosis** | State counts localise a hang faster than reading any single stack. |
| 2 | **Runnable is not blocked** | It means ready to run and waiting for a core — a scheduling signal, not a leak. |
| 3 | **A leak has a single site** | Thousands of goroutines with the identical top frame is the fingerprint. |

## Topics used again

Map aggregation with composite keys, sorting with tie-breaks, guards.

## Hint

Group by the `{state, site}` pair, then filter by count and sort.

## Validate

```bash
make verify
```
