// Package shutdownflag - Gopher Workplace challenge.
package shutdownflag

import "sync/atomic"

// ShutdownFlag is a one-way graceful-shutdown flag safe for concurrent use.
type ShutdownFlag struct {
	down atomic.Bool
}

// Request marks shutdown as requested.
//
// Examples:
//
//	var f ShutdownFlag; f.Request(); f.Requested() => true
func (f *ShutdownFlag) Request() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Requested reports whether shutdown has been requested.
//
// Examples:
//
//	var f ShutdownFlag; f.Requested() => false
func (f *ShutdownFlag) Requested() bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// ClaimShutdown requests shutdown and reports whether this call started it.
//
// Examples:
//
//	var f ShutdownFlag; f.ClaimShutdown() => true
//	f.ClaimShutdown()                     => false
func (f *ShutdownFlag) ClaimShutdown() bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
