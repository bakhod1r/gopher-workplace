// Package wrapfunc — Gopher Workplace challenge.
package wrapfunc

import (
	"errors"
	"fmt"
)

// ErrBoom is a stand-in step failure used by the tests.
var ErrBoom = errors.New("boom")

// Named returns f decorated so its failures are labelled with name.
//
// Examples:
//
//	Named("load", func() error { return nil })() => nil
func Named(name string, f func() error) func() error {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
