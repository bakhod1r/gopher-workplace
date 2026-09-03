# Sort A Small Set Without Reaching For The Heap

**Level:** senior
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A quicksort's pivot selection calls `sort.Ints` on a three-element slice. The slice escapes into `sort.Interface`, so the hottest function in the sort allocates twice per partition.

## Task

Implement [noescapesort.go](noescapesort.go):

1. Return the median of three ints.
2. Handle duplicates and negatives correctly.
3. Zero allocations — no slice, no interface, no sort package.

Replace the stub body in [noescapesort.go](noescapesort.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Median3(3, 1, 2)
Output: 2
```

**Example 2:**

```
Input:  Median3(5, 5, 1)
Output: 5
```

_Explanation:_ Duplicates make the median one of the repeated values.

**Example 3:**

```
Input:  Median3(-1, 0, 1)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sort.Interface boxes its argument** | A slice passed as an interface escapes to the heap. |
| 2 | **Comparison networks** | Three values need at most three comparisons and no storage. |
| 3 | **Hot-path specialisation** | A general tool is the wrong shape when n is fixed and tiny. |

## Hint

Three comparisons and two swaps of local variables. Nothing else.

## Validate

```bash
make verify
```
