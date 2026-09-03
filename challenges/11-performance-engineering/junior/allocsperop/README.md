# Counting Allocations Without A Profiler

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

"Does this function allocate?" is the cheapest performance question to answer and the one worth answering in a test, not a profile. The standard library exposes exactly that measurement, and it is deterministic enough to assert on.

## Task

Implement `AllocsOf` in [allocsperop.go](allocsperop.go):

1. Report the heap allocations performed by one call to `f`.
2. Average over `runs` samples and round to the nearest whole allocation.
3. Treat a non-positive `runs` as `1`.

## Examples

**Example 1:**

```
Input:  AllocsOf(100, func(){ sink = make([]byte, 1024) })
Output: 1
```

**Example 2:**

```
Input:  AllocsOf(100, func(){ sum += 1 })
Output: 0
```

**Example 3:**

```
Input:  AllocsOf(100, three makes)
Output: 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`testing.AllocsPerRun`** | Runs `f` a fixed number of times and returns the average allocation count. |
| 2 | **Allocation count is assertable** | Unlike timing, it is stable enough to fail a test on a regression. |
| 3 | **Escape analysis decides** | Work that stays on the stack contributes zero, which is what "zero-alloc" means. |

## Topics used again

Function values, float-to-int conversion.

## Hint

`testing.AllocsPerRun` returns a `float64` that is already a whole number in practice — round it, do not truncate.

## Validate

```bash
make verify
```
