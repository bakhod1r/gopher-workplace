// Package withstack — Gopher Workplace challenge.
package withstack

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
)

// ErrBase is a stand-in failure used by the tests.
var ErrBase = errors.New("base")

// Here annotates err with the caller's file and line.
//
// Examples:
//
//	Here(nil) => nil
func Here(err error) error {
	// TODO(candidate): implement this.
	_, _, _ = fmt.Errorf, filepath.Base, runtime.Caller
	panic("not implemented")
}
