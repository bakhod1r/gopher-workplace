# Reading A Mutex Profile

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

A CPU profile shows a program working; a mutex profile shows it waiting. Every record is a call site with two numbers — how many times a goroutine blocked there and how long it waited in total — and the two rank very differently. A lock taken ten thousand times for a microsecond each is a different problem from one taken twice for four milliseconds.

## Task

Implement both functions in [contentionrank.go](contentionrank.go):

1. `Rank` aggregates records per site, computing the total count, total delay, and mean delay per blocking event.
2. Order by total delay descending, then by site ascending; drop records with a non-positive count or negative delay.
3. `Worst` returns the highest-delay site, reporting `false` when there is none.

## Examples

**Example 1:**

```
Input:  Rank([{a 2 10} {a 2 30}])
Output: [{a 4 40 10}]
```

**Example 2:**

```
Input:  Rank([{small 100 100} {big 1 5000} {medium 10 1000}])
Output: big, medium, small
```

**Example 3:**

```
Input:  Worst([{a 1 5} {b 1 50}])
Output: site b, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two numbers, two stories** | Total delay finds the biggest waste; mean delay finds the slow critical section. |
| 2 | **Contention is measured, not guessed** | `runtime.SetMutexProfileFraction` turns this data on; without it you are speculating. |
| 3 | **Deterministic reports** | Aggregating through a map means the output needs a total order. |

## Topics used again

Map aggregation, `slices.SortFunc`, multi-key comparison, guards.

## Hint

`Worst` is the first row of `Rank`, if there is one.

## Validate

```bash
make verify
```
