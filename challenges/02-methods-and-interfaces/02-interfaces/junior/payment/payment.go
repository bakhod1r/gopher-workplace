// Package payment — Gopher Workplace challenge.
package payment

// Payment attempts to charge an amount.
type Payment interface {
	Charge(amount int) bool
}

// Card pays up to a credit limit.
type Card struct {
	Limit int
}

// Charge succeeds when amount fits in the limit.
//
// Examples:
//
//	Card{Limit: 100}.Charge(50)  => true
//	Card{Limit: 100}.Charge(150) => false
func (c Card) Charge(amount int) bool {
	// TODO(candidate): amount must fit the limit.
	panic("not implemented")
}

// Cash pays from a wallet.
type Cash struct {
	Available int
}

// Charge succeeds when amount fits the wallet.
func (c Cash) Charge(amount int) bool {
	// TODO(candidate): amount must fit Available.
	panic("not implemented")
}

// Checkout returns "paid" or "declined".
func Checkout(p Payment, amount int) string {
	// TODO(candidate): charge, then map the result to a word.
	panic("not implemented")
}
