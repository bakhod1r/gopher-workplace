// Package handlerscope — Gopher Workplace challenge.
package handlerscope

import "context"

// ServeRequest derives a per-request cancellable context, runs the handler with
// it, and guarantees the context is cancelled once the response has been
// written — so background lookups started by the handler are torn down instead
// of leaking for the life of the process.
//
// ServeRequest returns whatever the handler returns.
//
// Examples:
//
//	ServeRequest(func(ctx context.Context) error { return nil })     => nil
//	ServeRequest(func(ctx context.Context) error { return errDB })   => errDB
//	after ServeRequest returns, the handler's ctx has Err() == context.Canceled
func ServeRequest(handler func(ctx context.Context) error) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
