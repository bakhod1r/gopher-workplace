# Counters That Fight Over A Cache Line

**Level:** staff
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A profiler shows a parallel counter loop running slower with eight cores than with one. No lock is contended and the race detector is silent.

## Task

Implement [falseshare.go](falseshare.go):

1. Run `workers` goroutines, each incrementing its own `counter` `iters` times.
2. Wait for all of them, then return the sum of the counters.
3. Each goroutine must own its slot exclusively — no shared writes, no race.
4. Non-positive `workers` or negative `iters` return 0.

Replace the stub body in [falseshare.go](falseshare.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Count(4, 1000)
Output: 4000
```

**Example 2:**

```
Input:  Count(8, 100000)
Output: 800000
```

_Explanation:_ Correct under real parallelism.

**Example 3:**

```
Input:  Count(0, 10)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **False sharing** | Independent variables on one cache line are not independent to the hardware. |
| 2 | **Padding** | `pad [56]byte` after an int64 pushes the next counter to its own line. |
| 3 | **sync.WaitGroup** | The join point that makes the accumulation safe to read. |
| 4 | **Loop-variable capture** | Pass the slot as a parameter rather than closing over the index. |

## Hint

Every worker writes to its own variable and the cores still serialise. What is the unit of coherence?

## Validate

```bash
make verify
```
