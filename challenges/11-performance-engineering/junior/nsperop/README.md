# ns/op And "Is It Actually Faster?"

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

ns/op is the headline number of every benchmark line, and the question that follows it is always the same: did this change clear the bar we set? Both halves are arithmetic worth getting exactly right.

## Task

Implement both functions in [nsperop.go](nsperop.go):

1. `NsPerOp` divides `elapsedNS` by `iters`, truncating toward zero.
2. `NsPerOp` returns `0` when `iters` is non-positive or `elapsedNS` is negative.
3. `Faster` reports whether `candidate` is at least `pct` percent below `base`; a non-positive `base` is never beaten.

## Examples

**Example 1:**

```
Input:  NsPerOp(1000, 3)
Output: 333
```

**Example 2:**

```
Input:  Faster(100, 80, 20)
Output: true
```

**Example 3:**

```
Input:  Faster(100, 81, 20)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **ns/op is a truncated mean** | It hides variance completely; the number alone never proves a win. |
| 2 | **Improvement is relative to the base** | `(base-candidate)/base*100`, not a difference in nanoseconds. |
| 3 | **Boundary inclusive** | "At least 20% faster" must accept exactly 20%. |

## Topics used again

Integer division, float comparison.

## Hint

Express the threshold as `float64(candidate) <= float64(base)*(1-pct/100)` and the boundary takes care of itself.

## Validate

```bash
make verify
```
