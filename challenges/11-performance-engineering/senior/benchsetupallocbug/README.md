# Setup That Moved Into The Loop

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

The encoder was written to own one buffer and reuse it, and the allocation test that used to pass now fails. Nothing about the API changed — the buffer is being recreated on every call, which is exactly the cost the design existed to avoid.

## Task

Fix the single planted bug in [benchsetupallocbug.go](benchsetupallocbug.go):

1. Find and fix the one bug so a warm `Encode` performs no allocations.
2. Each call must still replace the previous record rather than appending to it.
3. The buffer must still grow when a record needs more room.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Encode([a b], [1 2])
Output: "a=1;b=2;"
```

**Example 2:**

```
Input:  Encode(first), then Encode(second)
Output: "second=2;"
```

**Example 3:**

```
Input:  a warm Encode measured with AllocsPerRun
Output: 0 allocations
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`make` in the hot path** | A fresh buffer per call is one allocation per call, however small it is. |
| 2 | **Reset is not reallocate** | `buf[:0]` clears the length and keeps the array. |
| 3 | **Both do the same thing functionally** | Which is why the correctness tests pass and only the allocation test fails. |

## Hint

The line between the markers produces a correct empty buffer. It just produces a new one.

## Validate

```bash
make verify
```
