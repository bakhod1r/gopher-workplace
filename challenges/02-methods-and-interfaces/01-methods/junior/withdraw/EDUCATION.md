# Conditional Mutation

## Intuition

`Withdraw` is a good example of conditional mutation: check preconditions
before changing state. The boolean return tells the caller whether the
operation succeeded.

## Approach

1. Guard: `amount <= 0` or `amount > Balance` → return `false`.
2. Subtract and return `true`.

## Solution

```go
func (a *Account) Withdraw(amount int) bool {
	if amount <= 0 || amount > a.Balance {
		return false
	}
	a.Balance -= amount
	return true
}
```

## Walkthrough

For `Account{100}.Withdraw(30)`:
- `30 > 0` and `30 <= 100` → proceed.
- `a.Balance -= 30` → 70.
- Returns `true`.

## Pitfalls

- Forgetting the `amount <= 0` check → negative withdrawal adds money.
- Not using pointer receiver → balance reverts after method returns.
