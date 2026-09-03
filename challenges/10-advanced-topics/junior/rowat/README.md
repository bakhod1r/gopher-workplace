# Reach Into A Slice Of Slices

**Level:** junior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A grid helper indexes rows directly. A request with a bad row number takes the whole service down with an index-out-of-range panic.

## Task

Implement [rowat.go](rowat.go):

1. Return the `i`-th row of `g` and whether it exists.
2. An out-of-range index reports false instead of panicking.
3. The row is a view — writes through it must reach `g`.

Replace the stub body in [rowat.go](rowat.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Row([][]int{{1},{2}}, 1)
Output: [2], true
```

**Example 2:**

```
Input:  Row(g, -1)
Output: nil, false
```

**Example 3:**

```
Input:  row[0] = 99 after Row(g, 0)
Output: g[0][0] is 99
```

_Explanation:_ The row is not a copy.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounds checks you write yourself** | The runtime's check panics; the caller wanted an answer. |
| 2 | **Slice of slices** | The outer slice holds headers; indexing it copies a header, not the elements. |
| 3 | **Views share storage** | Writing through the returned row reaches the grid. |

## Hint

Check both ends of the range before indexing.

## Validate

```bash
make verify
```
