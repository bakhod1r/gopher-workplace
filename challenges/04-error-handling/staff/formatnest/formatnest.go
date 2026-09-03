// Package formatnest — Gopher Workplace challenge.
package formatnest

import (
	"errors"
	"strings"
)

// ErrA is a stand-in root failure used by the tests.
var ErrA = errors.New("a")

// Verbose renders err's chain, one line per layer.
//
// Examples:
//
//	Verbose(nil)  => ""
//	Verbose(ErrA) => "a"
func Verbose(err error) string {
	// TODO(candidate): implement this.
	_, _ = errors.Unwrap, strings.Join
	panic("not implemented")
}
