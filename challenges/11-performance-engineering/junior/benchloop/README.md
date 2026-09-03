# The Loop `b.N` Expects

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Every Go benchmark body is the same shape: run the work exactly `b.N` times, no more and no less. Get the loop wrong and the reported ns/op is a fiction. Here you write that loop by hand.

## Task

Implement `Run` in [benchloop.go](benchloop.go):

1. Call `work` with `i` counting `0, 1, ... n-1`.
2. Return the number of calls made.
3. A non-positive `n` performs no calls and returns `0`.

## Examples

**Example 1:**

```
Input:  Run(3, record)
Output: 3, work saw indexes [0 1 2]
```

**Example 2:**

```
Input:  Run(0, record)
Output: 0, work never called
```

**Example 3:**

```
Input:  Run(-5, record)
Output: 0, work never called
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`b.N` is a count, not a size** | The harness picks `N`; the body runs the same work `N` times. |
| 2 | **Off-by-one ruins ns/op** | Elapsed time is divided by the count you promised to run. |
| 3 | **Guarding the loop** | A non-positive count must be a no-op, never a panic. |

## Hint

One `for` loop, one counter, one guard.

## Validate

```bash
make verify
```
