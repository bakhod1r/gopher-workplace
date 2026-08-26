// Package withdraw — Gopher Workplace challenge.
package withdraw

// Account holds a balance in cents.
type Account struct {
	Balance int
}

// Withdraw subtracts amount from the balance if sufficient funds exist.
// Returns true on success, false if insufficient funds or invalid amount.
//
// Examples:
//
//	a := Account{100}; ok := a.Withdraw(30) // true, a.Balance == 70
//	a := Account{100}; ok := a.Withdraw(200) // false, a.Balance == 100
func (a *Account) Withdraw(amount int) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
