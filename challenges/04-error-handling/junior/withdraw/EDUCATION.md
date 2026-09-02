# Account Withdrawal

## Intuition

A failed operation must not half-apply. Returning the untouched balance beside the error keeps the caller's state consistent whichever branch runs.

## Approach

1. Reject non-positive amounts.
2. Reject amounts larger than the balance.
3. Return `balance - amount, nil`.

## Solution

```go
if amount <= 0 {
	return balance, ErrInvalidAmount
}
if amount > balance {
	return balance, ErrInsufficientFunds
}
return balance - amount, nil
```

## Walkthrough

`Withdraw(100, 100)` is allowed — `100 > 100` is false — and empties the account to `0, nil`.

## Pitfalls

- Returning `0` on failure, wiping out the caller's balance.
- Using `amount >= balance`, which forbids a legal full withdrawal.
- Checking affordability before amount validity, so `-5` reads as affordable.
