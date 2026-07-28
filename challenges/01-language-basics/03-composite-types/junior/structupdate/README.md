# Mutate via Pointer Receiver

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

To change a struct's field from a method, the receiver must be a pointer.

## Task

Implement `Deposit(amount)` on `*Account` to add to `Balance`.

## Examples

```go
a := &Account{100}; a.Deposit(50) // Balance == 150
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer receiver** | `func (a *Account)` can mutate. |
| 2 | **Value vs pointer** | Value receiver mutates a copy only. |
| 3 | **Auto-address** | `a.Deposit()` works on an addressable value. |

## Hint

`a.Balance += amount`.

## Validate

```bash
make verify
```
