# Rates From Cumulative Counters

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

`runtime.MemStats` exposes counters, not rates: `TotalAlloc` and `Mallocs` only ever go up. Every allocation-rate dashboard is built by subtracting two snapshots — and every one of them has, at some point, shipped the bug where the counters are `uint64`, the subtraction underflows, and the graph shows eighteen quintillion bytes per second.

## Task

Implement the three functions in [allocratecalc.go](allocratecalc.go):

1. `Delta` returns the bytes, mallocs and frees between two snapshots, reporting `false` when any counter decreased or the timestamp did not advance.
2. `BytesPerSec` converts that byte delta into a per-second rate.
3. `LiveObjects` returns `Mallocs - Frees`, reporting `false` rather than wrapping when frees exceed mallocs.

## Examples

**Example 1:**

```
Input:  Delta({0 100 10 4}, {1s 500 30 20})
Output: 400, 20, 16, true
```

**Example 2:**

```
Input:  BytesPerSec({0 0}, {2s 2000000})
Output: 1000000, true
```

**Example 3:**

```
Input:  LiveObjects({Mallocs: 4, Frees: 10})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Counters need differencing** | The absolute value only says how long the process has been up. |
| 2 | **Unsigned subtraction wraps** | `uint64(4) - uint64(10)` is a huge positive number, not an error. |
| 3 | **A counter reset means a restart** | The right response is to discard the sample, not to graph the jump. |

## Topics used again

Unsigned arithmetic, multiple return values, guards, float conversion.

## Hint

Compare before subtracting; every `uint64` subtraction in this file needs a guard above it.

## Validate

```bash
make verify
```
