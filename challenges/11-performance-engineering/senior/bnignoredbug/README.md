# The Body That Runs Once

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

A benchmark reports a suspiciously round, suspiciously large ns/op and never varies. The harness raised the iteration count into the millions; the body did not notice.

## Task

Fix the single planted bug in [bnignoredbug.go](bnignoredbug.go):

1. Find and fix the one bug so the body runs exactly `n` times.
2. `work` must see indexes `0` through `n-1`.
3. A non-positive `n` performs no calls.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Run(3, record)
Output: 3, indexes [0 1 2]
```

**Example 2:**

```
Input:  Run(1000, sum)
Output: 1000, sum 499500
```

**Example 3:**

```
Input:  PerOp(300, 3, noop)
Output: 100
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`b.N` is the harness's decision** | It raises the count until the run is long enough to time; the body must follow. |
| 2 | **Ignoring it inflates ns/op** | The whole elapsed time is divided by a count the body did not honour. |
| 3 | **A constant bound is invisible** | The loop looks like a loop, and the number looks like a measurement. |

## Hint

Look at what the loop condition compares against.

## Validate

```bash
make verify
```
