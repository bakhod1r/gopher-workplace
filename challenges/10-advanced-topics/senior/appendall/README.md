# Concatenate Slices With One Allocation

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A merge step concatenates a few dozen slices per batch. It starts from nil and grows, reallocating and copying everything it has so far at every doubling.

## Task

Fix the single planted bug in [appendall.go](appendall.go):

1. Concatenate the parts in order.
2. Fix the single bug so the result costs one allocation.
3. An empty input, or all-empty parts, returns an empty result.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  AppendAll([][]int{{1},{2,3}})
Output: [1 2 3]
```

**Example 2:**

```
Input:  64 parts of 4 elements
Output: 1 allocation, not several
```

**Example 3:**

```
Input:  AppendAll(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Known output size** | The sum of the parts' lengths is the exact final length. |
| 2 | **append's doubling** | Growing from nil to 256 elements reallocates about nine times. |
| 3 | **make with length 0 and a capacity** | Keeps `append` semantics while reserving the space. |

## Hint

One extra loop before the existing one.

## Validate

```bash
make verify
```
