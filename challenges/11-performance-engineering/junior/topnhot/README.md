# The Top Listing

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`pprof -top` shows the ten hottest functions, and everyone stops reading after the first three. Producing that list from a map means confronting the fact that Go's map iteration order is deliberately randomised — without a total order on the rows, the same profile prints differently every run.

## Task

Implement `TopN` in [topnhot.go](topnhot.go):

1. Order by `Value` descending.
2. Break ties by `Func` ascending, so the output is deterministic.
3. Return at most `n` rows; a non-positive `n` returns an empty, non-nil slice.

## Examples

**Example 1:**

```
Input:  TopN({a:3 b:9 c:3 d:1}, 2)
Output: [{b 9} {a 3}]
```

**Example 2:**

```
Input:  TopN({z:5 a:5 m:5}, 3)
Output: [{a 5} {m 5} {z 5}]
```

**Example 3:**

```
Input:  TopN({a:1}, 10)
Output: [{a 1}]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Map order is randomised** | Any report built from a map needs an explicit tiebreaker to be reproducible. |
| 2 | **Descending by value, ascending by name** | Two keys, opposite directions, one comparison function. |
| 3 | **Clamping the slice** | `n` larger than the data must not slice out of range. |

## Topics used again

`slices.SortFunc`, `cmp.Compare`, map iteration, slice bounds.

## Hint

Collect into a slice, sort with a comparison that falls through to the name, then cut with `min(n, len(entries))`.

## Validate

```bash
make verify
```
