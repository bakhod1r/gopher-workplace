# Remove by Index

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Deleting an element while keeping order uses the append-of-two-halves idiom.

## Task

Implement `RemoveAt(xs, i)`; out-of-range `i` returns `xs` unchanged.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  xs=[1,2,3,4], i=1
Output: []int{1,3,4}
```

**Example 2:**

```
Input:  xs=[1,2,3], i=0
Output: []int{2,3}
```

**Example 3:**

```
Input:  xs=[1,2,3], i=5
Output: []int{1,2,3}
```

_Explanation:_ Out-of-range index returns xs unchanged.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two-half append** | `append(xs[:i], xs[i+1:]...)`. |
| 2 | **Bounds guard** | Check `0 <= i < len`. |
| 3 | **Variadic spread** | `...` expands the tail. |

## Hint

Guard bounds, then `return append(xs[:i], xs[i+1:]...)`.

## Validate

```bash
make verify
```
