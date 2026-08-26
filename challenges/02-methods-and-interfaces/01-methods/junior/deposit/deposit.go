// Package deposit — Gopher Workplace challenge.
package deposit

// Account holds a balance in cents.
type Account struct {
	Balance int
}

// Deposit adds amount to the balance. Amount must be positive; if not, it
// is ignored (no-op).
//
// Examples:
//
//	a := Account{100}; a.Deposit(50) // a.Balance == 150
//	a := Account{100}; a.Deposit(-1) // a.Balance == 100 (ignored)
func (a *Account) Deposit(amount int) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
