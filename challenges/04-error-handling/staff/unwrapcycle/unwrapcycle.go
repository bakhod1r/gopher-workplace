// Package unwrapcycle — Gopher Workplace challenge.
package unwrapcycle

import "errors"

// ErrA is a stand-in root failure used by the tests.
var ErrA = errors.New("a")

// Chain returns each error's message, stopping if the chain repeats.
//
// Examples:
//
//	Chain(nil) => nil
func Chain(err error) []string {
	// TODO(candidate): implement this.
	_ = errors.Unwrap
	panic("not implemented")
}
