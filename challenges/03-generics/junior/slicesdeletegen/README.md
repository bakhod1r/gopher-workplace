# Delete

**Level:** junior  
**Topic:** 03-generics

## Context

A user removes one row from a table. The row index comes from the UI and may be stale.

## Task

Implement the stub(s) in [slicesdeletegen.go](slicesdeletegen.go):

1. Implement `RemoveAt` using `slices.Delete`.
2. Return `s` unchanged for an out-of-range index.
3. Do not modify the caller's slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RemoveAt([]int{1, 2, 3}, 1)
Output: []int{1, 3}
```

**Example 2:**

```
Input:  RemoveAt([]int{1}, 0)
Output: []int{}
```

**Example 3:**

```
Input:  RemoveAt([]int{1}, 5)
Output: []int{1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.Delete`** | `Delete(s, i, j)` removes the half-open range `[i, j)`. |
| 2 | **Half-open ranges** | Removing one element means `Delete(s, i, i+1)`. |
| 3 | **Cloning first** | Reused from earlier: clone before an in-place helper when the caller must keep its data. |

## Hint

`Delete` takes a range, so one element is `i` to `i+1`.

## Validate

```bash
make verify
```
