// Package structupdate mutates a struct through a pointer.
package structupdate

// Account holds a balance in cents.
type Account struct {
	Balance int
}

// Deposit adds amount to the account's balance, mutating the receiver.
//
// TODO(candidate): use a pointer receiver to mutate in place.
func (a *Account) Deposit(amount int) {
	panic("not implemented")
}
