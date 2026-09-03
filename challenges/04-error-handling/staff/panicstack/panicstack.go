// Package panicstack — Gopher Workplace challenge.
package panicstack

import (
	"fmt"
	"runtime"
)

// Trace runs f, converting a panic into an error carrying a stack snippet.
//
// Examples:
//
//	Trace(func() {}) => nil
func Trace(f func()) (err error) {
	// TODO(candidate): implement this.
	_, _ = fmt.Errorf, runtime.Stack
	panic("not implemented")
}
