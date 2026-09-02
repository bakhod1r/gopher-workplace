// Package withdraw — Gopher Workplace challenge.
package withdraw

import "errors"

// Withdrawal failures.
var (
	ErrInvalidAmount     = errors.New("amount must be positive")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

// Withdraw debits amount from balance and returns the new balance.
//
// Examples:
//
//	Withdraw(100, 30)  => 70, nil
//	Withdraw(100, 150) => 100, ErrInsufficientFunds
func Withdraw(balance, amount int) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
