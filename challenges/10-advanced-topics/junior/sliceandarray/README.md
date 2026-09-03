# A Slice Header Versus An Array

**Level:** junior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A memory estimate multiplies the record count by the size of a struct containing a slice. The estimate is off by the entire payload, because the slice's elements were never counted.

## Task

Implement [sliceandarray.go](sliceandarray.go):

1. Return `unsafe.Sizeof` for a `[8]int` and for an `[]int`.
2. Derive both from declared variables rather than writing numbers.

Replace the stub body in [sliceandarray.go](sliceandarray.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Sizes()
Output: 64, 24
```

_Explanation:_ Eight ints; a pointer, a length and a capacity.

**Example 2:**

```
Input:  Sizeof of a 100000-element slice
Output: still 24
```

_Explanation:_ The header does not grow.

**Example 3:**

```
Input:  array vs slice
Output: the array is larger
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Arrays carry their contents** | The length is in the type, so the size includes every element. |
| 2 | **Slices are three words** | Pointer, length and capacity — the elements live elsewhere. |
| 3 | **Sizeof measures the type** | Never what a pointer or header refers to. |

## Hint

Declare one of each and measure them.

## Validate

```bash
make verify
```
