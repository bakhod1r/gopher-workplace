# Auth Check

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A gateway accepts several credential kinds. Each one decides for itself whether it grants access.

## Task

Implement the stub(s) in [authcheck.go](authcheck.go):

1. Implement `Allow` on `Token` — allow when the token is non-empty and not `"expired"`.
2. Implement `Allow` on `Guest` — a guest is never allowed.
3. Implement `CanEnter`, which reports whether the given `Authenticator` allows access.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CanEnter(Token{Value: "abc"})
Output: true
```

**Example 2:**

```
Input:  CanEnter(Token{Value: "expired"})
Output: false
```

**Example 3:**

```
Input:  CanEnter(Guest{})
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface satisfaction** | A type implements an interface by having its methods — no `implements` keyword. |
| 2 | **Value receivers** | `Token` and `Guest` both satisfy `Authenticator` with value receivers. |
| 3 | **String comparison** | Reused from language basics: `==` and `len` on strings. |

## Hint

`CanEnter` just calls `a.Allow()` — the dynamic type picks the method.

## Validate

```bash
make verify
```
