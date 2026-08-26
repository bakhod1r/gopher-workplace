# Negate

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A ledger system negates transactions for reversals. The method mutates the
value in place.

## Task

Implement `Negate` on `*MyInt` in [negate.go](negate.go):

1. Flip the sign: positive → negative, negative → positive.
2. Pointer receiver — mutation must persist.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  n := MyInt(5); n.Negate()
Output: n == -5
```

**Example 2:**

```
Input:  n := MyInt(-3); n.Negate()
Output: n == 3
```

**Example 3:**

```
Input:  n := MyInt(0); n.Negate()
Output: n == 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer receiver on defined type** | `*MyInt` lets you mutate the underlying value. |
| 2 | **Dereference and assign** | `*n = -(*n)` or equivalent. |

## Hint

`*n = -(*n)` — dereference, negate, store back.

## Validate

```bash
make verify
```
