// Package tracecarrier — Gopher Workplace challenge.
package tracecarrier

import "context"

// traceKey is the unexported context key type for the trace ID. Using a named
// empty struct type — not a string — keeps other packages from colliding.
type traceKey struct{}

// WithTrace returns a child context carrying the trace ID. An empty ID is
// ignored: the parent is returned unchanged.
//
// Examples:
//
//	WithTrace(ctx, "abc") => a child carrying "abc"
//	WithTrace(ctx, "")    => ctx unchanged
func WithTrace(ctx context.Context, id string) context.Context {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// TraceID returns the trace ID on the context, or "" when there is none.
//
// Examples:
//
//	TraceID(WithTrace(ctx, "abc"))   => "abc"
//	TraceID(context.Background())    => ""
func TraceID(ctx context.Context) string {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Chain simulates a call chain: it tags the context, then reports what each
// downstream layer observes. Layers that run before the tag must see "".
//
// Examples:
//
//	Chain(ctx, "abc") => ["" "abc" "abc"]
//	Chain(ctx, "")    => ["" "" ""]
func Chain(ctx context.Context, id string) []string {
	// TODO(candidate): read the ID before tagging, then twice after.
	panic("not implemented")
}
