# Guard Clauses in Methods

## Intuition

A guard clause validates input before acting. Here, negative deposits are
silently ignored — a design choice. An alternative is returning an error, but
this puzzle focuses on method mutation.

## Approach

1. Check `amount > 0`.
2. Add to balance.

## Solution

```go
func (a *Account) Deposit(amount int) {
	if amount > 0 {
		a.Balance += amount
	}
}
```

## Walkthrough

For `Account{100}.Deposit(50)`:
- `50 > 0` → true.
- `a.Balance += 50` → 150.

## Pitfalls

- Forgetting the guard → negative "deposits" reduce the balance (unexpected).
- Value receiver → balance change is lost.
