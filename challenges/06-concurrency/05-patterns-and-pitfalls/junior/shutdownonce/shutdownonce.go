// Package shutdownonce — Gopher Workplace challenge.
package shutdownonce

// ShutdownOnce lets closers goroutines all request shutdown at the same time
// while the quit channel is closed exactly once. It returns after every
// requester has finished. With zero closers the channel is left open.
//
// Examples:
//
//	ShutdownOnce(quit, 1)   => quit closed, no panic
//	ShutdownOnce(quit, 10)  => quit closed exactly once
//	ShutdownOnce(quit, 0)   => quit left open
func ShutdownOnce(quit chan struct{}, closers int) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
