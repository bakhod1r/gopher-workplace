# The Ranking That Changes Every Run

**Level:** staff  
**Topic:** 03-generics

## Context

A leaderboard endpoint is served from a map. Its golden-file test passes locally and fails on CI roughly half the time; the same input produces a different ordering of the tied entries on every process start.

## Task

Fix the single planted bug in [rankvalueonlybug.go](rankvalueonlybug.go):

1. Find and fix the single bug so the ordering is fully determined by the data.
2. Higher values must still come first.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Rank({a:2, b:1})
Output: [a b]
```

**Example 2:**

```
Input:  Rank({b:1, a:1})
Output: [a b]
```

**Example 3:**

```
Input:  Rank({})
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deterministic output** | Map iteration order is randomised, so any ordering the caller depends on must be fully specified by the comparator. |
| 2 | **A stable sort of an unstable input** | Stability preserves the *input* order — and that order came from randomised map iteration. |
| 3 | **Total orders** | A comparator that returns 0 for distinct elements leaves their order to whatever produced the slice. |

## Hint

What decides the order of two keys with the same value?

## Validate

```bash
make verify
```
