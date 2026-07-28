# Nil-Safe Dereference

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _nil-pointer-dereference_

## Context

Dereferencing a nil pointer panics. Guard with a nil check before reading `*p`.

## Task

Implement `ValueOr` in [valueor.go](valueor.go).

Do **not** change the function signature or the tests.

## Examples

```go
ValueOr(nil, 5) // => 5
ValueOr(&9, 5)  // => 9
```

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
