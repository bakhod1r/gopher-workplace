// Package wrapbudget — Gopher Workplace challenge.
package wrapbudget

import (
	"errors"
	"fmt"
)

// ErrBase is a stand-in root failure used by the tests.
var ErrBase = errors.New("base")

// Wrap annotates err with msg unless its chain already has max links.
//
// Examples:
//
//	Wrap(ErrBase, "a", 2) => "a: base"
//	Wrap(nil, "a", 2)     => nil
func Wrap(err error, msg string, max int) error {
	// TODO(candidate): implement this.
	_, _ = fmt.Errorf, errors.Unwrap
	panic("not implemented")
}
