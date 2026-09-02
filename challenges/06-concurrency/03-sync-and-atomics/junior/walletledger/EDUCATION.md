# Wallet Ledger

## Intuition

If you check the balance, release the lock, and then deduct, another charge can slip in between and both succeed - overdrawing the wallet. Holding the lock across the whole decision keeps the invariant true.

## Approach

1. `Debit` locks and defers unlock.
2. If `amount > balance`, return false.
3. Otherwise deduct and return true.

## Solution

```go
// Package walletledger - Gopher Workplace challenge.
package walletledger

import "sync"

// Wallet is a customer balance safe for concurrent transactions.
type Wallet struct {
	mu      sync.Mutex
	balance int
}

// NewWallet returns a wallet with the given opening balance.
func NewWallet(opening int) *Wallet {
	return &Wallet{balance: opening}
}

// Credit adds amount to the wallet.
//
// Examples:
//
//	w := NewWallet(100); w.Credit(50); w.Balance() => 150
func (w *Wallet) Credit(amount int) {
	w.mu.Lock()
	w.balance += amount
	w.mu.Unlock()
}

// Debit deducts amount when funds suffice and reports success.
//
// Examples:
//
//	w := NewWallet(100); w.Debit(30) => true
//	w := NewWallet(20); w.Debit(50)  => false
func (w *Wallet) Debit(amount int) bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	if amount > w.balance {
		return false
	}
	w.balance -= amount
	return true
}

// Balance returns the current balance.
func (w *Wallet) Balance() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.balance
}
```

## Walkthrough

Balance 100, two concurrent `Debit(80)` charges. The first holds the lock, sees 100, deducts to 20. The second then sees 20, declines, returns false. The wallet never goes negative.

## Pitfalls

- Splitting the check and the deduction across two lock holds.
- Locking `Debit` but not `Balance`.
- Returning the balance from `Debit` and letting callers act on a value that is already stale.
