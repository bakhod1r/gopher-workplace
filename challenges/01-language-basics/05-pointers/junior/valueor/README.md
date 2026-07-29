# Nil-Safe Dereference

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _nil-pointer-dereference_

## Context

Dereferencing a nil pointer panics. Guard with a nil check before reading `*p`.

## Task

Implement `ValueOr` in [valueor.go](valueor.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ValueOr(nil, 5)
Output: 5
```

_Explanation:_ Nil pointer → the default is returned, no dereference.

**Example 2:**

```
Input:  n := 9; ValueOr(&n, 5)
Output: 9
```

**Example 3:**

```
Input:  n := 0; ValueOr(&n, 5)
Output: 0
```

_Explanation:_ A real pointer to 0 still returns 0, not the default.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil check** | `if p == nil` guards the deref. |
| 2 | **Safe dereference** | Only read `*p` when non-nil. |
| 3 | **Default fallback** | Return def for nil. |

## Hint

`if p == nil { return def }; return *p`.

## Validate

```bash
make verify
```
