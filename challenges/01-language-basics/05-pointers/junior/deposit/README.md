# Deposit Into Account

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Mutating a struct's field through a pointer is exactly what a pointer-receiver
method does under the hood.

## Task

Implement `Deposit` in [deposit.go](deposit.go).

Do **not** change the function signature or the tests.

## Examples

```go
Deposit(&a, 50) // a.Balance += 50
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Struct pointer** | `*Account` aliases the caller's account. |
| 2 | **Field mutation** | `a.Balance += amount`. |
| 3 | **In-place update** | Caller sees the new balance. |

## Hint

`a.Balance += amount`.

## Validate

```bash
make verify
```
