// Package asinterface — Gopher Workplace challenge.
package asinterface

import (
	"errors"
	"fmt"
)

// ErrOther is a stand-in unrelated failure used by the tests.
var ErrOther = errors.New("other")

// Throttled reports a retry delay.
type Throttled struct {
	Seconds int
}

// Error implements the error interface.
func (e *Throttled) Error() string {
	return fmt.Sprintf("throttled for %ds", e.Seconds)
}

// RetryAfter reports the delay in seconds.
func (e *Throttled) RetryAfter() int {
	return e.Seconds
}

// DelayOf returns the retry hint carried anywhere in err.
//
// Examples:
//
//	DelayOf(&Throttled{Seconds: 5}) => 5, true
func DelayOf(err error) (int, bool) {
	// TODO(candidate): implement this.
	_ = errors.As
	panic("not implemented")
}
