# Copy Value Between Pointers

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Assigning `*dst = *src` copies the value; the two variables stay independent
afterward.

## Task

Implement `CopyInto` in [copyval.go](copyval.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a, b := 1, 9; CopyInto(&a, &b)
Output: a == 9, b == 9
```

**Example 2:**

```
Input:  a, b := 5, 5; CopyInto(&a, &b)
Output: a == 5, b == 5
```

**Example 3:**

```
Input:  a, b := 0, -2; CopyInto(&a, &b)
Output: a == -2, b == -2
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Deref both sides** | `*dst = *src`. |
| 2 | **Value copy** | Independent after copy. |
| 3 | **Direction** | src into dst. |

## Hint

`*dst = *src`.

## Validate

```bash
make verify
```
