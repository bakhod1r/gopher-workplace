# How Many Workers

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

One goroutine per item is the default answer and it is wrong twice over: a million goroutines is a million stacks, and CPU-bound work on eight cores goes no faster with a thousand runners than with eight. A bounded pool fixes both — and the right bound depends on whether the work computes or waits.

## Task

Implement both functions in [workerpoolsize.go](workerpoolsize.go):

1. `Map` applies `f` to every item with at most `workers` goroutines, returning results in input order.
2. A non-positive `workers` runs one goroutine; more workers than items is harmless; the result must be race-free under `-race`.
3. `Sizing` returns `cpus` for CPU-bound work and `cpus/(1-blocked)` when a `blocked` fraction is spent waiting, floored and never below 1; a `cpus` below 1 gives 1, and a `blocked` outside `[0,1)` is treated as 0.

## Examples

**Example 1:**

```
Input:  Map([1 2 3 4 5], 3, double)
Output: [2 4 6 8 10]
```

**Example 2:**

```
Input:  Sizing(8, 0)
Output: 8
```

**Example 3:**

```
Input:  Sizing(8, 0.9)
Output: 80
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded parallelism** | Beyond the core count, more goroutines add scheduling and cache pressure, not throughput. |
| 2 | **Distinct slots need no lock** | Each worker writing `out[i]` for its own `i` is race-free by construction. |
| 3 | **Waiting changes the arithmetic** | Work that blocks 90% of the time needs ten times the workers to keep the CPUs busy. |

## Topics used again

Goroutines, channels, `sync.WaitGroup`, slice indexing.

## Hint

Feed indexes down a channel and let each worker write `out[i]` — no mutex required.

## Validate

```bash
make verify
```
