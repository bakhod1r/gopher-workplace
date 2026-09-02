# The Ranking That Changes Between Runs

**Level:** staff  
**Topic:** 03-generics

## Context

A "top tags" panel is cached by content hash. The hash keeps changing between identical rebuilds, so the cache never hits and reviewers see the tag order shuffle on every deploy.

## Task

Fix the single planted bug in [orderleakbug.go](orderleakbug.go):

1. Find and fix the single bug so equal counts are ordered by ascending value.
2. The primary ordering by descending count must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  RankByCount([b a b c])
Output: [b a c]
```

**Example 2:**

```
Input:  all counts equal
Output: values in ascending order, every run
```

**Example 3:**

```
Input:  RankByCount([]string{})
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deterministic output** | Map iteration order is randomised; anything ordered must be sorted by a total order. |
| 2 | **Stable is not deterministic** | A stable sort preserves the *input* order — and here the input order came from a map. |
| 3 | **Total orders** | A comparator that returns 0 for distinct elements leaves their order to the caller's luck. |

## Hint

The sort is stable. Where did the slice it stabilises come from?

## Validate

```bash
make verify
```
