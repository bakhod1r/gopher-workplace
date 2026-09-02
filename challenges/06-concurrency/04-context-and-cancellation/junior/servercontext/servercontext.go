// Package servercontext — Gopher Workplace challenge.
package servercontext

import "context"

// ServerContext returns the root context that main creates once at process
// start and passes to the HTTP server, the database pool and the metrics
// exporter. Nothing above it exists, so it is never cancelled and carries no
// deadline and no values.
//
// Examples:
//
//	ServerContext() != nil     => true
//	ServerContext().Err()      => nil
//	ServerContext().Done()     => nil (never closed)
func ServerContext() context.Context {
	// TODO(candidate): implement this.
	panic("not implemented")
}
