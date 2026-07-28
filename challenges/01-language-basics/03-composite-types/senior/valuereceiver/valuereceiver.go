// Package valuereceiver credits an account. A planted value receiver drops the
// mutation.
package valuereceiver

// Wallet holds a balance in cents.
type Wallet struct {
	Balance int
}

// Credit adds amount to the wallet's balance.
// CHANGE CODE BELOW THIS LINE
func (w Wallet) Credit(amount int) {
	// CHANGE CODE ABOVE THIS LINE
	w.Balance += amount
}
