# Dropping The Wrong End

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

The sampling loop discards warmup rounds, the constant is right, the guards are right — and the reported steady-state mean is still dominated by the cold-cache first run. The code is discarding samples from the wrong end of the slice.

## Task

Fix the single planted bug in [warmupcountedbug.go](warmupcountedbug.go):

1. Find and fix the one bug so the *first* `warmup` samples are discarded.
2. A non-positive warmup keeps everything; discarding every sample gives `0`.
3. The clamping already in place must keep working.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  StableMean([100 2 4], 1)
Output: 3
```

**Example 2:**

```
Input:  StableMean([100 90 2 4], 2)
Output: 3
```

**Example 3:**

```
Input:  StableMean([1 2], 9)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Warm-up is at the start** | Cold caches and lazy initialisation are paid on the first iterations. |
| 2 | **`s[n:]` versus `s[:len(s)-n]`** | Both compile, both have the right length, and only one drops the right samples. |
| 3 | **The right length is not the right data** | A slice of the correct size hides this bug from every length assertion. |

## Hint

Both slice expressions produce the same number of samples. Only one of them keeps the steady state.

## Validate

```bash
make verify
```
