// Package requiretext — Gopher Workplace challenge.
package requiretext

import (
	"errors"
	"strings"
)

// ErrRequired reports a missing value.
var ErrRequired = errors.New("value required")

// Require reports whether s carries any non-whitespace content.
//
// Examples:
//
//	Require("hello") => nil
//	Require("   ")   => ErrRequired
func Require(s string) error {
	// TODO(candidate): implement this.
	_ = strings.TrimSpace
	panic("not implemented")
}
