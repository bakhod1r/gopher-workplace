// Package shutdowntwice — Gopher Workplace challenge.
package shutdowntwice

// ShutdownTwice models an operator sending SIGTERM and then, impatiently,
// SIGINT: the signal handler calls the shutdown cancel func twice. It returns
// the shutdown context's error afterwards.
//
// Examples:
//
//	ShutdownTwice()                              => context.Canceled
//	calling cancel twice must not panic          => true
//	the second call must not change the reason   => true
func ShutdownTwice() error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
