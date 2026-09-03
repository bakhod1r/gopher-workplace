# Waiting For A Core

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Scheduling latency is the time between a goroutine becoming runnable and actually running. In a healthy program it is nanoseconds. When it grows, adding goroutines makes things worse rather than better — the work is already queued and the cores are already busy. `go tool trace` reports it, and this is the arithmetic behind that view.

## Task

Implement the three functions in [schedlatency.go](schedlatency.go):

1. `Delay` returns one event's wait, reporting `false` when `Running` precedes `Runnable`.
2. `Delays` returns the valid delays in input order.
3. `Stats` returns the mean and worst delay, reporting `false` when there are no valid events.

## Examples

**Example 1:**

```
Input:  Delay({1 100 150})
Output: 50, true
```

**Example 2:**

```
Input:  Delays([{1 0 10} {2 50 40} {3 0 5}])
Output: [10 5]
```

**Example 3:**

```
Input:  Stats([{1 0 10} {2 0 20}])
Output: 15, 20, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Runnable is not running** | The gap between them is queueing, and it is invisible in a CPU profile. |
| 2 | **Rising latency means saturation** | More goroutines cannot help when the cores are already the constraint. |
| 3 | **The mean hides the tail** | Ninety-nine fast schedules and one 10µs stall barely move the average. |

## Topics used again

Multiple return values, filtering, mean and maximum in one pass, guards.

## Hint

`Stats` can build on `Delays` rather than re-validating the events.

## Validate

```bash
make verify
```
