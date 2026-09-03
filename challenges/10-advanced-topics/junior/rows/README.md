# One Allocation For The Whole Grid

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

An image filter builds its working buffer as a slice of independently allocated rows. The rows land all over the heap and the filter's inner loop misses cache on every row change.

## Task

Implement [rows.go](rows.go):

1. Return an `r` by `c` grid of zeros.
2. All rows must be windows into one flat array — at most two allocations total.
3. Return nil when `r <= 0` or `c <= 0`.

Replace the stub body in [rows.go](rows.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Rows(2, 3)
Output: [[0 0 0] [0 0 0]]
```

**Example 2:**

```
Input:  &g[0][2] and &g[1][0]
Output: adjacent addresses
```

_Explanation:_ Row 1 begins where row 0 ends.

**Example 3:**

```
Input:  Rows(0, 3)
Output: <nil>
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Backing arrays** | Many slices can be windows into one allocation. |
| 2 | **Three-index slicing** | `flat[a:b:b]` caps each row so an append cannot spill into the next. |
| 3 | **Locality** | Contiguous memory is what makes the traversal cache-friendly. |

## Hint

Allocate `r*c` ints first, then hand out windows of `c`.

## Validate

```bash
make verify
```
