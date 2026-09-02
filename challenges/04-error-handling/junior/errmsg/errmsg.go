// Package errmsg — Gopher Workplace challenge.
package errmsg

import "errors"

// ErrTimeout is a stand-in failure used by the tests.
var ErrTimeout = errors.New("timeout")

// Message returns err's text, or "" when err is nil.
//
// Examples:
//
//	Message(nil)        => ""
//	Message(ErrTimeout) => "timeout"
func Message(err error) string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
