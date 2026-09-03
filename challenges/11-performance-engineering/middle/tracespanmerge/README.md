# What The Trace Actually Covers

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

An execution trace is a pile of intervals: goroutines running, GC phases, syscalls. Asking "how long was the GC active" is not a sum — three concurrent mark workers running for 100ms each is 100ms of GC, not 300ms. Merging the intervals first is what turns overlapping spans into an answer about wall-clock time.

## Task

Implement both functions in [tracespanmerge.go](tracespanmerge.go):

1. `Merge` combines overlapping *and* touching spans into the fewest that cover the same time, ordered by start.
2. Drop spans whose end is at or before their start, and do not modify the input.
3. `Covered` returns the total time the merged spans occupy.

## Examples

**Example 1:**

```
Input:  Merge([{0 10} {5 20}])
Output: [{0 20}]
```

**Example 2:**

```
Input:  Merge([{0 100} {10 20}])
Output: [{0 100}]
```

**Example 3:**

```
Input:  Covered([{0 100} {0 100} {0 100}])
Output: 100
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sort by start, then sweep** | The classic interval merge: one sort and one linear pass. |
| 2 | **Nested spans must not shrink the outer one** | Extend the end with a maximum, never assign it. |
| 3 | **Coverage is not the sum** | Parallel work overlaps, and only merged intervals answer wall-clock questions. |

## Topics used again

`slices.Clone`, `slices.SortFunc`, `max`, interval sweeps.

## Hint

After sorting, a span either extends the current one or starts a new one — the test is `s.Start <= cur.End`.

## Validate

```bash
make verify
```
