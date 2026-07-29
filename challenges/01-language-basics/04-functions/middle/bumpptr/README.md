# Pointer Parameter Mutates

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _call-by-value_

## Context

A value parameter is a copy; a pointer parameter aliases the caller's variable,
so writing through it is visible outside.

## Task

Implement `Bump` in [bumpptr.go](bumpptr.go): increment `*p`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  x := 5; Bump(&x)
Output: x == 6
```

**Example 2:**

```
Input:  x := 0; Bump(&x)
Output: x == 1
```

**Example 3:**

```
Input:  x := -1; Bump(&x)
Output: x == 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer parameter** | `*int` refers to the caller's variable. |
| 2 | **Dereference to write** | `*p++` or `*p = *p + 1`. |
| 3 | **Value vs pointer** | Only pointers let a callee mutate a caller's int. |

## Hint

`*p++` (or `*p = *p + 1`).

## Validate

```bash
make verify
```
