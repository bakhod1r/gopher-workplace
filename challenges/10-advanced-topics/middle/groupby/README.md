# Bucket Pairs Without Growing Every Bucket

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A grouping step over a few million rows spends most of its time in `growslice`: every bucket doubles its way up from nothing, and there are thousands of buckets.

## Task

Implement [groupby.go](groupby.go):

1. Collect each pair's second element into a bucket keyed by its first.
2. Preserve input order within a bucket.
3. Size the map and every bucket up front — each bucket's capacity must equal its final length.
4. An empty input returns an empty, non-nil map.

Replace the stub body in [groupby.go](groupby.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Group([][2]int{{1,10},{2,20},{1,11}})
Output: map[1:[10 11] 2:[20]]
```

**Example 2:**

```
Input:  Group(nil)
Output: map[]
```

_Explanation:_ Empty, not nil.

**Example 3:**

```
Input:  cap of each bucket
Output: equal to its length
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two-pass sizing** | Counting first turns every bucket's growth into one allocation. |
| 2 | **Map size hints** | The distinct-key count is known after the counting pass. |
| 3 | **Append order** | A single forward pass preserves the input order for free. |

## Hint

Count first, allocate second, fill third.

## Validate

```bash
make verify
```
