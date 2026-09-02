// Package querystate — Gopher Workplace challenge.
package querystate

// QueryState models the reporting endpoint's database query context. It
// returns the context's error twice: once while the client is still connected
// and the query is allowed to run, and once after the client disconnected and
// the handler cancelled it.
//
// Examples:
//
//	connected, _ := QueryState()      => connected is nil
//	_, disconnected := QueryState()   => disconnected is context.Canceled
//	connected != disconnected         => true
func QueryState() (connected, disconnected error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
