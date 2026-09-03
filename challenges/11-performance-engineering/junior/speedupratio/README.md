# "Twice As Fast" Means One Thing

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Halving a latency is a 2x speedup and a 50% reduction. Those are the same result described two ways, and mixing them up — "200% faster" — is how a modest win gets written up as a miracle. Compute both, from the same two numbers, and the ambiguity disappears.

## Task

Implement both functions in [speedupratio.go](speedupratio.go):

1. `Speedup` returns `base / candidate`, the multiplier.
2. `PercentChange` returns the signed change relative to base, negative when faster.
3. Non-positive inputs give `0`, except that a candidate of `0` in `PercentChange` is a `-100%` change.

## Examples

**Example 1:**

```
Input:  Speedup(100, 25)
Output: 4
```

**Example 2:**

```
Input:  PercentChange(100, 80)
Output: -20
```

**Example 3:**

```
Input:  PercentChange(100, 125)
Output: 25
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Ratio vs percentage** | A 2x speedup is −50%; they are reciprocal views of one measurement. |
| 2 | **The baseline is the denominator** | Both formulas divide by the original, never the new value. |
| 3 | **A slowdown is a speedup below 1** | The same formula handles regressions without a second branch. |

## Topics used again

Float arithmetic, guards, signed results.

## Hint

`PercentChange` divides the difference by base; `Speedup` divides base by candidate.

## Validate

```bash
make verify
```
