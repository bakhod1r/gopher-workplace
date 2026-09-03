// Package repanic — Gopher Workplace challenge.
package repanic

import (
	"fmt"
	"runtime"
)

// Handle runs f, converting application panics into errors and
// re-panicking on runtime faults.
//
// Examples:
//
//	Handle(func() {}) => nil
func Handle(f func()) (err error) {
	// TODO(candidate): implement this.
	_, _ = fmt.Errorf, runtime.Gosched
	panic("not implemented")
}
