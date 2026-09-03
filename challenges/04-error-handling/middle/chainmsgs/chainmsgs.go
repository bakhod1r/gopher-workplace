// Package chainmsgs — Gopher Workplace challenge.
package chainmsgs

import "errors"

// ErrBase is a stand-in root failure used by the tests.
var ErrBase = errors.New("base failure")

// Messages returns the message of every error in err's chain.
//
// Examples:
//
//	Messages(ErrBase) => ["base failure"]
//	Messages(nil)     => nil
func Messages(err error) []string {
	// TODO(candidate): implement this.
	_ = errors.Unwrap
	panic("not implemented")
}
