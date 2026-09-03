# ns/op That Drifts With `b.N`

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

A benchmark that builds a large fixture reports a different ns/op every run, and the number always drops as the machine gets faster at nothing in particular. Somebody has been chasing the variance for a week. The fixture is not the thing being measured.

## Task

Fix the single planted bug in [timerbeforesetupbug.go](timerbeforesetupbug.go):

1. Find and fix the one bug so the reported total charges only the per-iteration work.
2. The non-positive `n` guard already works and must keep working.
3. `PerOp` must return the same value for every `n`.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Measured(1000, 7, 3)
Output: 21
```

**Example 2:**

```
Input:  PerOp(1000000, 7, 10) and PerOp(1000000, 7, 1000000)
Output: the same number
```

**Example 3:**

```
Input:  Measured(500, 7, 0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`b.ResetTimer` discards** | Everything measured before the call is dropped, not amortised. |
| 2 | **Fixed cost divided by `b.N`** | It shrinks as `N` grows, which is why the number drifts between runs. |
| 3 | **Stability is the test** | A correct benchmark reports the same ns/op whatever iteration count the tool picks. |

## Hint

One of the two terms in the sum does not belong to any iteration.

## Validate

```bash
make verify
```
