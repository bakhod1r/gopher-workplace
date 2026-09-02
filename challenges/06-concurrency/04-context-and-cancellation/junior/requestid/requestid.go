// Package requestid — Gopher Workplace challenge.
package requestid

import "context"

// requestIDKey is the unexported key under which the request ID is stored.
type requestIDKey struct{}

// WithRequestID returns a copy of ctx carrying the request ID that the edge
// proxy assigned to this request.
//
// Examples:
//
//	RequestID(WithRequestID(bg, "req-8f21")) => "req-8f21"
//	RequestID(context.Background())          => "unknown"
//	RequestID(WithRequestID(bg, ""))         => "unknown"
func WithRequestID(ctx context.Context, id string) context.Context {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// RequestID returns the request ID the logger should stamp on every line, or
// "unknown" when the context carries no usable ID.
func RequestID(ctx context.Context) string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
