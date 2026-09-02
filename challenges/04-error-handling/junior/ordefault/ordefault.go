// Package ordefault — Gopher Workplace challenge.
package ordefault

import "errors"

// ErrMissing is a stand-in failure used by the tests.
var ErrMissing = errors.New("missing")

// OrDefault returns v when err is nil, otherwise def.
//
// Examples:
//
//	OrDefault("8080", nil, "80")       => "8080"
//	OrDefault("", ErrMissing, "80")    => "80"
func OrDefault(v string, err error, def string) string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
