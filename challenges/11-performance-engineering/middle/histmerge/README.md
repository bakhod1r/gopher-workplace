# Merging Histograms Across Instances

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Fleet-wide latency comes from adding every instance's histogram together — which works only when they share bucket bounds. Add counts from mismatched bounds and you get a plausible-looking histogram describing nothing, which is why a merge must refuse rather than guess.

## Task

Implement both functions in [histmerge.go](histmerge.go):

1. `Valid` checks that the bounds are strictly ascending and the counts are exactly one longer.
2. `Merge` sums the counts of two valid histograms with identical bounds.
3. Mismatched or malformed input reports `false`; neither input may be modified or aliased.

## Examples

**Example 1:**

```
Input:  Merge({[1 5] [1 2 3]}, {[1 5] [10 20 30]})
Output: {[1 5] [11 22 33]}, true
```

**Example 2:**

```
Input:  Merge({[1 5] ...}, {[1 6] ...})
Output: false
```

**Example 3:**

```
Input:  Valid({[2 1] [0 0 0]})
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Buckets only add under identical bounds** | Different edges make the counts incomparable, not merely imprecise. |
| 2 | **The counts are one longer** | The extra slot is the overflow bucket past the final bound. |
| 3 | **Refuse rather than approximate** | A silently wrong merge is worse than no merge at all. |

## Topics used again

`slices.Equal`, `slices.Clone`, validation, multiple returns.

## Hint

Validate both inputs, compare the bounds with `slices.Equal`, then sum into fresh slices.

## Validate

```bash
make verify
```
