// Package boundedjoin — Gopher Workplace challenge.
package boundedjoin

import (
	"errors"
	"fmt"
)

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
	ErrC = errors.New("c")
)

// Cap joins at most max failures, summarising the rest.
//
// Examples:
//
//	Cap(2) => nil
func Cap(max int, errs ...error) error {
	// TODO(candidate): implement this.
	_, _ = fmt.Errorf, errors.Join
	panic("not implemented")
}
