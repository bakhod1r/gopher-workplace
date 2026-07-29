# Insert Shift Direction

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

To insert at `i`, the elements from `i` onward must shift **right** by one. The
code does `copy(xs[i:], xs[i+1:])`, which shifts left, dropping the tail.

## Task

Fix the copy between the markers in [insertshift.go](insertshift.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  xs=[1 2 3], i=1, v=9
Output: [1 9 2 3]
```

**Example 2:**

```
Input:  xs=[1 2 3], i=0, v=0
Output: [0 1 2 3]
```

**Example 3:**

```
Input:  xs=[1 2 3], i=3, v=4
Output: [1 2 3 4]
```

_Explanation:_ insert at end.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shift right** | `copy(xs[i+1:], xs[i:])`. |
| 2 | **Overlap-safe** | `copy` handles the overlap. |
| 3 | **Grow then shift** | Append a slot, then open the gap. |

## Hint

`copy(xs[i+1:], xs[i:])`.

## Validate

```bash
make verify
```
