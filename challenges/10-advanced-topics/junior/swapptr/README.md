# Exchange Two Values Through Pointers

**Level:** junior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A sort helper takes the values instead of pointers. It swaps its own copies, the slice never changes, and the sort loops forever.

## Task

Implement [swapptr.go](swapptr.go):

1. Exchange the values `a` and `b` point at.
2. Do nothing when either pointer is nil.
3. Zero allocations.

Replace the stub body in [swapptr.go](swapptr.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  x, y := 1, 2; Swap(&x, &y)
Output: x is 2, y is 1
```

**Example 2:**

```
Input:  Swap(&x, &x)
Output: x unchanged
```

_Explanation:_ Swapping a value with itself is a no-op.

**Example 3:**

```
Input:  Swap(&x, nil)
Output: nothing happens
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Writing through a pointer** | `*p = v` reaches the caller's variable. |
| 2 | **Tuple assignment** | `*a, *b = *b, *a` evaluates the right side before assigning. |
| 3 | **Aliasing is fine here** | The tuple form is correct even when both pointers are the same. |

## Hint

One statement does it, and it happens to handle the aliased case too.

## Validate

```bash
make verify
```
