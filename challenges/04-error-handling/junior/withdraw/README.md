# Account Withdrawal

**Level:** junior
**Topic:** 04-error-handling

## Context

A ledger debits an account. Overdrafts are forbidden, and so are non-positive amounts.

## Task

Implement `Withdraw` in [withdraw.go](withdraw.go):

1. Return the new balance and nil for a valid withdrawal.
2. Return the unchanged balance and `ErrInvalidAmount` when `amount <= 0`.
3. Return the unchanged balance and `ErrInsufficientFunds` when `amount > balance`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Withdraw(100, 30)
Output: 70, nil
```

**Example 2:**

```
Input:  Withdraw(100, 0)
Output: 100, ErrInvalidAmount
```

**Example 3:**

```
Input:  Withdraw(100, 150)
Output: 100, ErrInsufficientFunds
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Ordered validation** | Amount sanity is checked before affordability. |
| 2 | **Unchanged state on failure** | A rejected operation leaves the balance alone. |
| 3 | **Multiple sentinels** | Two rules, two errors. |

## Hint

On failure the caller still needs a usable balance — return the original, not zero.

## Validate

```bash
make verify
```
