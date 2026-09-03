// Package stackframes — Gopher Workplace challenge.
package stackframes

import "runtime"

// Frames returns up to max function names from the caller upward.
//
// Examples:
//
//	Frames(0) => []
func Frames(max int) []string {
	// TODO(candidate): implement this.
	_ = runtime.Callers
	panic("not implemented")
}
