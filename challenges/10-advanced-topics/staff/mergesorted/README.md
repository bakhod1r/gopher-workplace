# Merge Runs Into The Caller's Buffer

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

An external sort merges its runs by concatenating and re-sorting. The concatenation doubles peak memory and throws away the ordering the runs already have.

## Task

Implement [mergesorted.go](mergesorted.go):

1. Append every element of the sorted `runs` to `dst` in ascending order.
2. Preserve duplicates; an empty input returns `dst` unchanged.
3. With room in `dst` and a modest number of runs, allocate nothing.

Replace the stub body in [mergesorted.go](mergesorted.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Merge(nil, [][]int{{1,3,5},{2,4}})
Output: [1 2 3 4 5]
```

**Example 2:**

```
Input:  Merge([]int{0}, [][]int{{2},{1}})
Output: [0 1 2]
```

_Explanation:_ dst is extended.

**Example 3:**

```
Input:  Merge(nil, [][]int{{1,1},{1}})
Output: [1 1 1]
```

_Explanation:_ Duplicates survive.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Merging preserves work** | The runs are already sorted; re-sorting throws that away. |
| 2 | **Cursors as the only state** | One index per run, no data movement. |
| 3 | **A stack array for small cases** | A fixed local array avoids allocating the cursor slice. |
| 4 | **Append-style output** | The caller owns the destination. |

## Hint

One cursor per run. Each step picks the smallest element still under a cursor.

## Validate

```bash
make verify
```
