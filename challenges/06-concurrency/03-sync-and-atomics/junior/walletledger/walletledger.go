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
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Debit deducts amount when funds suffice and reports success.
//
// Examples:
//
//	w := NewWallet(100); w.Debit(30) => true
//	w := NewWallet(20); w.Debit(50)  => false
func (w *Wallet) Debit(amount int) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Balance returns the current balance.
func (w *Wallet) Balance() int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
