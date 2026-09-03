# How Parallel Was It, Really

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

`go tool trace` shows a row per processor and lets you see, at a glance, whether your carefully parallelised program ever ran more than one thing at a time. The computation behind that view is a sweep over span starts and ends: `+1` when something begins, `−1` when it ends, and the running maximum is the peak parallelism.

## Task

Implement the three functions in [tracespanoverlap.go](tracespanoverlap.go):

1. `Intersect` returns the overlap of two spans and whether one exists; touching spans do not overlap.
2. `Concurrency` returns the maximum number of spans active at any instant, ignoring empty spans.
3. `BusiestAt` returns the earliest timestamp at which that peak is reached, and `false` when there are no valid spans.

## Examples

**Example 1:**

```
Input:  Intersect({0 10}, {5 20})
Output: {5 10}, true
```

**Example 2:**

```
Input:  Concurrency([{0 10} {5 20} {6 7}])
Output: 3
```

**Example 3:**

```
Input:  BusiestAt([{0 10} {5 20}])
Output: 5, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The sweep-line counter** | Sort the endpoints, walk them, keep a running count and its maximum. |
| 2 | **Ends before starts at the same instant** | Half-open intervals mean a span ending at `t` is already gone when one starts at `t`. |
| 3 | **Peak parallelism is a bound** | It tells you what the program achieved, not what the machine offered. |

## Topics used again

Sorting composite keys, sweep lines, `max`, multiple return values.

## Hint

Build a slice of `{time, delta}` events, sort by time with `−1` before `+1` on ties, then sweep.

## Validate

```bash
make verify
```
