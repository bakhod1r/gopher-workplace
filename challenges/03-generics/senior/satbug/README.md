# Saturation That Never Triggers

**Level:** senior  
**Topic:** 03-generics

## Context

A byte counter still wraps to near zero after a busy day, even though someone added saturation logic for exactly that case.

## Task

Fix the single planted bug in [satbug.go](satbug.go):

1. Find and fix the single bug so wrap-around is detected.
2. Do not widen to a fixed type — the maximum must come from `T`.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  SatAdd(uint(1), 2)
Output: 3
```

**Example 2:**

```
Input:  SatAdd(maxUint64, 1)
Output: maxUint64
```

**Example 3:**

```
Input:  SatAdd(uint(0), 0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Unsigned values are never negative** | `sum < 0` is dead code for every type in the set. |
| 2 | **Detecting the wrap** | After wrapping, the sum is smaller than either operand. |
| 3 | **Dead branches compile fine** | The compiler accepts the comparison; only the behaviour is wrong. |

## Hint

Can an unsigned sum ever be less than zero?

## Validate

```bash
make verify
```
