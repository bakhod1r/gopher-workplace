# Swap Through Pointers

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

With two pointers you can swap the caller's variables in place — the classic
demonstration that pointers alias real storage.

## Task

Implement `Swap` in [swapptr.go](swapptr.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  x, y := 1, 2; Swap(&x, &y)
Output: x == 2, y == 1
```

_Explanation:_ Parallel assignment through both pointers exchanges the values.

**Example 2:**

```
Input:  x, y := 5, 5; Swap(&x, &y)
Output: x == 5, y == 5
```

**Example 3:**

```
Input:  x, y := -1, 9; Swap(&x, &y)
Output: x == 9, y == -1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two pointer params** | Each aliases a caller variable. |
| 2 | **Dereference both** | `*a, *b = *b, *a`. |
| 3 | **Parallel assignment** | RHS evaluated before assigning. |

## Hint

`*a, *b = *b, *a`.

## Validate

```bash
make verify
```
