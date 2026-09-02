// Package percent — Gopher Workplace challenge.
package percent

import "errors"

// Percentage failures.
var (
	ErrZeroTotal = errors.New("total is zero")
	ErrNegative  = errors.New("negative argument")
)

// Percent returns part as a percentage of total.
//
// Examples:
//
//	Percent(1, 4) => 25, nil
//	Percent(1, 0) => 0, ErrZeroTotal
func Percent(part, total int) (float64, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
