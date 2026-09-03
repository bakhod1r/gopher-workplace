// Package budgetsteps — Gopher Workplace challenge.
package budgetsteps

import "errors"

// ErrBudgetExceeded reports an exhausted operation budget.
var ErrBudgetExceeded = errors.New("budget exceeded")

// Spend runs steps while the budget allows.
//
// Examples:
//
//	Spend(0) => nil
func Spend(budget int, steps ...func() error) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
