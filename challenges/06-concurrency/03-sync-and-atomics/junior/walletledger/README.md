# Wallet Ledger

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A payments service debits and credits customer wallets from concurrent transaction handlers. A wallet must never go negative, so the balance check and the deduction have to happen as one indivisible step.

## Task

Implement the stubbed functions in [walletledger.go](walletledger.go) so that:

1. `Credit` adds an amount to the wallet.
2. `Debit` deducts the amount only if funds suffice, and reports success.
3. `Balance` returns the current balance.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  w := NewWallet(100); w.Credit(50); w.Balance()
Output: 150
```

**Example 2:**

```
Input:  w := NewWallet(100); w.Debit(30)
Output: true, balance 70
```

**Example 3:**

```
Input:  w := NewWallet(20); w.Debit(50)
Output: false, balance 20
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Mutex** | Check-then-act must live inside one lock hold. |
| 2 | **Invariants** | The lock protects a business rule (balance >= 0), not just a field. |
| 3 | **Multiple return values** | `Debit` reports whether the charge went through. |

## Hint

Take the lock at the top of `Debit` with `defer w.mu.Unlock()`, then check and deduct inside.

## Validate

```bash
make verify
go test -race ./...
```
