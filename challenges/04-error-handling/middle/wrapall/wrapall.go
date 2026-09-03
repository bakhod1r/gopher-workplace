// Package wrapall — Gopher Workplace challenge.
package wrapall

import (
	"errors"
	"fmt"
)

// ErrJob is a stand-in job failure used by the tests.
var ErrJob = errors.New("job failed")

// WrapAll annotates every non-nil error with its index.
//
// Examples:
//
//	WrapAll([]error{nil, ErrJob}) => ["job 1: job failed"]
//	WrapAll(nil)                  => nil
func WrapAll(errs []error) []error {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
