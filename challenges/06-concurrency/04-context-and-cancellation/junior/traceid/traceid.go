// Package traceid — Gopher Workplace challenge.
package traceid

import "context"

// traceKey is an unexported key type, so no other package can collide with the
// trace ID this package stores.
type traceKey struct{}

// WithTraceID returns a copy of ctx carrying the distributed-trace ID read
// from the incoming gRPC metadata.
//
// Examples:
//
//	TraceID(WithTraceID(context.Background(), "4bf92f"))  => "4bf92f", true
//	TraceID(context.Background())                         => "", false
//	TraceID(WithTraceID(WithTraceID(bg, "a"), "b"))       => "b", true
func WithTraceID(ctx context.Context, id string) context.Context {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// TraceID reports the trace ID carried by ctx, if any. Outbound clients call it
// to stamp the header on the next hop.
func TraceID(ctx context.Context) (string, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
