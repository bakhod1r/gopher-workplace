// Package sentinelerr — Gopher Workplace challenge.
package sentinelerr

import "errors"

// MaxCount is the largest count Validate accepts.
const MaxCount = 1000

// The conditions Validate can report.
var (
	ErrNegative = errors.New("count is negative")
	ErrTooLarge = errors.New("count is too large")
)

// Validate reports whether n is a usable count: it must be non-negative
// and no greater than MaxCount.
//
// The failures are fixed conditions, so they must be reported with the
// package's sentinel errors rather than a freshly formatted one.
//
// Examples:
//
//	Validate(-1) => ErrNegative
func Validate(n int) error {
	panic("not implemented")
}
