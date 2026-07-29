# Zero Through Pointer

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Assigning through a pointer resets the caller's variable.

## Task

Implement `Zero` in [setzero.go](setzero.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  x := 99; Zero(&x)
Output: x == 0
```

_Explanation:_ Assigning through the pointer clears the caller's variable.

**Example 2:**

```
Input:  x := 0; Zero(&x)
Output: x == 0
```

**Example 3:**

```
Input:  x := -7; Zero(&x)
Output: x == 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Write through pointer** | `*p = 0`. |
| 2 | **Value reset** | The caller sees the change. |
| 3 | **Pointer vs value** | A value parameter could not do this. |

## Hint

`*p = 0`.

## Validate

```bash
make verify
```
