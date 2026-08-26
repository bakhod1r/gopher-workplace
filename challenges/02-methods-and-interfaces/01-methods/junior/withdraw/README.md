# Withdraw

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

The banking service also debits funds. Overdrafts and negative amounts are
rejected.

## Task

Implement `Withdraw` on `*Account` in [withdraw.go](withdraw.go):

1. If `amount > 0` and `amount <= Balance`, subtract and return `true`.
2. Otherwise, return `false` (balance unchanged).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a := Account{100}; a.Withdraw(30)
Output: true, a.Balance == 70
```

**Example 2:**

```
Input:  a := Account{100}; a.Withdraw(200)
Output: false, a.Balance == 100
```

**Example 3:**

```
Input:  a := Account{100}; a.Withdraw(100)
Output: true, a.Balance == 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer receiver** | Balance mutation must persist. |
| 2 | **Boolean return** | Success/failure signal. |
| 3 | **Guard clauses** | Validate before mutating. |

## Hint

Check both `amount > 0` and `amount <= a.Balance` before subtracting.

## Validate

```bash
make verify
```
