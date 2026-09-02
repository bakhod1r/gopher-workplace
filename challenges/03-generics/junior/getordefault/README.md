# Get Or Default

**Level:** junior  
**Topic:** 03-generics

## Context

A settings loader falls back to a built-in default when a key is absent — but a key explicitly set to zero must keep its zero.

## Task

Implement the stub(s) in [getordefault.go](getordefault.go):

1. Implement `GetOr`, returning the stored value when `k` is present.
2. Return `def` only when the key is missing, not when the stored value is zero.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  GetOr(map[string]int{"a": 1}, "a", 9)
Output: 1
```

**Example 2:**

```
Input:  GetOr(map[string]int{"a": 0}, "a", 9)
Output: 0
```

**Example 3:**

```
Input:  GetOr(map[string]int{}, "a", 9)
Output: 9
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comma-ok lookup** | `v, ok := m[k]` is the only way to tell a stored zero from a missing key. |
| 2 | **Maps with type parameters** | A map key type must be comparable, so `K` needs the `comparable` constraint. |
| 3 | **`if` with init statement** | Reused from language basics: the lookup can live in the `if` header. |

## Hint

A plain `m[k]` cannot distinguish "absent" from "present and zero".

## Validate

```bash
make verify
```
