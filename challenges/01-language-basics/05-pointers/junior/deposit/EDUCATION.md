# Pointer receivers by hand

## Intuition

A function taking `*Account` and updating a field is the mechanism behind pointer-receiver methods.

## Approach

1. `a` is an `*Account`.
2. `a.Balance += amount` mutates the caller's field through the pointer.

## Solution

```go
type Account struct{ Balance int }

func Deposit(a *Account, amount int) {
	a.Balance += amount
}
```

## Walkthrough

`Deposit(&a, 50)` with `Balance = 100`: the field becomes `150` in the caller's struct.

## Pitfalls

- A value parameter would update a copy, losing the change.
- `a.Balance += amount` auto-dereferences the pointer.
