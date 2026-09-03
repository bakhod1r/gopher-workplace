// Package runsafe — Gopher Workplace challenge.
package runsafe

import (
	"errors"
	"fmt"
)

// ErrRuntime reports a recovered runtime panic.
var ErrRuntime = errors.New("runtime panic")

// Run executes f, converting panics into errors wrapping ErrRuntime.
//
// Examples:
//
//	Run(func() error { return nil }) => nil
func Run(f func() error) (err error) {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
