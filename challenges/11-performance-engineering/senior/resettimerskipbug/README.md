# A Reset That Resets Nothing

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

The benchmark calls `ResetTimer` in the right place, the code reads correctly, and the setup cost is still in every result. The call is there; it just does not do anything.

## Task

Fix the single planted bug in [resettimerskipbug.go](resettimerskipbug.go):

1. Find and fix the one bug so `Reset` discards everything measured so far.
2. Measurements added after a reset must still accumulate normally.
3. `Benchmark` must report only the per-iteration work.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Add(500); Reset(); Elapsed()
Output: 0
```

**Example 2:**

```
Input:  Add(500); Reset(); Add(3); Add(4)
Output: 7
```

**Example 3:**

```
Input:  Benchmark(500, 7, 3)
Output: 21
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reset marks, it does not clear** | The total keeps growing; what changes is the point measurement counts from. |
| 2 | **Reset means zero** | Not "pause", not "note the position" — the accumulated measurement is discarded. |
| 3 | **The call being present proves nothing** | Test the behaviour, not the presence of the line. |

## Hint

`Elapsed` subtracts the mark from the total. What should the mark be set to?

## Validate

```bash
make verify
```
