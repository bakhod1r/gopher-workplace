// Package retryattempts — Gopher Workplace challenge.
package retryattempts

import (
	"errors"
	"fmt"
)

// ErrNoAttempts reports a non-positive attempt budget.
var ErrNoAttempts = errors.New("no attempts allowed")

// Retry calls f until it succeeds or the attempt budget runs out.
//
// Examples:
//
//	Retry(3, func() error { return nil }) => nil
//	Retry(0, f)                           => ErrNoAttempts
func Retry(attempts int, f func() error) error {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
