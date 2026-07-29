# Negate In Place

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Negation in place reads and rewrites through the pointer.

## Task

Implement `Negate` in [negate.go](negate.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  x := 5; Negate(&x)
Output: x == -5
```

**Example 2:**

```
Input:  x := -3; Negate(&x)
Output: x == 3
```

**Example 3:**

```
Input:  x := 0; Negate(&x)
Output: x == 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Write through pointer** | `*p = -*p`. |
| 2 | **Involution** | Applying twice restores. |
| 3 | **In-place** | Caller's variable changes. |

## Hint

`*p = -*p`.

## Validate

```bash
make verify
```
