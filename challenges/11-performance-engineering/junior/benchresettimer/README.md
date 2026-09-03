# The Setup That Should Not Count

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Building a 10MB fixture before the loop is not the thing you are benchmarking, but the timer is already running when the body starts. `b.ResetTimer()` throws away everything measured so far. Model that arithmetic.

## Task

Implement `Measured` in [benchresettimer.go](benchresettimer.go):

1. Charge `workNS` for each of the `n` iterations.
2. Charge nothing at all for `setupNS`.
3. A non-positive `n` measures `0`.

## Examples

**Example 1:**

```
Input:  Measured(1000, 7, 3)
Output: 21
```

**Example 2:**

```
Input:  Measured(500, 7, 0)
Output: 0
```

**Example 3:**

```
Input:  Measured(1, 2, 1000000)
Output: 2000000
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`b.ResetTimer` discards, not pauses** | Everything before the call is dropped from the total. |
| 2 | **Setup pollutes ns/op** | A fixed cost divided by `b.N` shrinks as `N` grows, so results drift run to run. |
| 3 | **Per-iteration cost is what you report** | Only work inside the loop belongs in the number. |

## Hint

`setupNS` appears in the signature but not in the answer.

## Validate

```bash
make verify
```
