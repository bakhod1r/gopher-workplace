# Move Elements Over Themselves

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A queue drains its head by allocating a new slice for the tail. The tail is almost the whole queue, so the drain costs a full copy into fresh memory every time.

## Task

Implement [shiftleft.go](shiftleft.go):

1. Drop the first `n` elements by moving the rest to the front of the same array.
2. Clamp `n` into `[0, len(s)]`; return the shortened slice.
3. Zero allocations, and correct when the ranges overlap.

Replace the stub body in [shiftleft.go](shiftleft.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Shift([]int{1,2,3,4}, 2)
Output: [3 4]
```

**Example 2:**

```
Input:  Shift([]int{1,2}, 9)
Output: []
```

_Explanation:_ n is clamped to the length.

**Example 3:**

```
Input:  Shift(s, 1) over 1000 elements
Output: every element moved down by one
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **copy handles overlap** | It behaves like memmove, so a left shift over itself is well defined. |
| 2 | **copy returns the count** | The number of elements moved is the new length. |
| 3 | **In-place over reallocation** | The array is already there; only the view changes. |

## Hint

`copy(s, s[n:])` — and its return value is the answer's length.

## Validate

```bash
make verify
```
