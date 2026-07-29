# Simultaneous Assignment

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _call-by-value_

## Context

Sequential assignments overwrite `a` before its old value can reach `c`. Go's
multiple assignment `a, b, c = b, c, a` evaluates the whole right-hand side
first, so all moves happen simultaneously.

## Task

Fix [rotate3.go](rotate3.go) using simultaneous assignment.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RotateLeft(1, 2, 3)
Output: 2, 3, 1
```

**Example 2:**

```
Input:  RotateLeft(5, 5, 5)
Output: 5, 5, 5
```

**Example 3:**

```
Input:  RotateLeft(0, 1, 2)
Output: 1, 2, 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Simultaneous assignment** | RHS is fully evaluated before any assign. |
| 2 | **Clobbering** | Sequential moves lose the original value. |
| 3 | **No temp needed** | Parallel assignment avoids temporaries. |

## Hint

Use one parallel assignment: `a, b, c = b, c, a`.

## Validate

```bash
make verify
```
