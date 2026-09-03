# Is The Heap Growing, Or Just Breathing?

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

A healthy Go heap is a sawtooth: it climbs to the GC target, collapses, climbs again. A leak looks the same except the troughs creep upward. Telling them apart needs two things — knowing where the next collection will trigger, and measuring the slope between the endpoints rather than reacting to the teeth.

## Task

Implement the three functions in [heapgrowthrate.go](heapgrowthrate.go):

1. `NextTarget` returns `live * (1 + GOGC/100)`, reporting `false` when GOGC is 0 or less.
2. `GrowthPerSec` returns the bytes-per-second slope between the first and last samples, reporting `false` for fewer than two samples or a non-increasing timestamp.
3. `Doubling` returns how many seconds the heap needs to double at that slope, reporting `false` when the slope is not positive.

## Examples

**Example 1:**

```
Input:  NextTarget(4MB, 100)
Output: 8MB, true
```

**Example 2:**

```
Input:  GrowthPerSec([{0 100} {1s 300}])
Output: 200, true
```

**Example 3:**

```
Input:  Doubling([{0 100} {1s 200}])
Output: 2, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **GOGC is a ratio, not a size** | `GOGC=100` means "collect when the heap has grown by 100% over live". |
| 2 | **Measure the trend, not the teeth** | The sawtooth between endpoints says nothing about whether memory is being retained. |
| 3 | **Doubling time is the actionable number** | "Grows 200 bytes a second" means nothing; "doubles every two seconds" is a pager. |

## Topics used again

Multiple return values, float slopes, guards, int64 time arithmetic.

## Hint

`Doubling` is `last.Live / slope` once the slope is in bytes per second.

## Validate

```bash
make verify
```
