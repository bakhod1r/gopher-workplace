// Package disconnectbeatstimeout — Gopher Workplace challenge.
package disconnectbeatstimeout

// DisconnectDuringTimeout models a request that was given a generous timeout
// but whose client hung up long before the deadline: the handler's cancel func
// runs first. It returns the reason the context ultimately recorded.
//
// Examples:
//
//	DisconnectDuringTimeout()                              => context.Canceled
//	errors.Is(DisconnectDuringTimeout(), context.Canceled) => true
//	the result is never context.DeadlineExceeded
func DisconnectDuringTimeout() error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
