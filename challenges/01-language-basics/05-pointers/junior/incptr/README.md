# Increment Through Pointer

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

A pointer parameter `*int` refers to the caller's variable. Writing through it
with `*p` changes the original.

## Task

Implement `Inc` in [incptr.go](incptr.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  x := 41; Inc(&x)
Output: x == 42
```

_Explanation:_ `Inc` follows the pointer to `x` and adds 1 in place.

**Example 2:**

```
Input:  x := 0; Inc(&x)
Output: x == 1
```

**Example 3:**

```
Input:  x := -1; Inc(&x)
Output: x == 0
```

_Explanation:_ Works for negatives too — dereference, add one, store back.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer parameter** | `*int` aliases the caller's variable. |
| 2 | **Dereference to write** | `*p` reads/writes the pointee. |
| 3 | **Address-of** | The caller passes `&x`. |

## Hint

`*p++` (or `*p = *p + 1`).

## Validate

```bash
make verify
```
