// Package requestmeta — Gopher Workplace challenge.
package requestmeta

import "context"

// Unexported key types keep each value in its own slot and out of reach of
// other packages.
type tenantKey struct{}

type traceKey struct{}

// WithMeta returns a copy of ctx carrying both the tenant ID and the trace ID
// that the edge middleware extracted from the request.
//
// Examples:
//
//	Meta(WithMeta(bg, "acme", "4bf9"))  => "acme", "4bf9"
//	Meta(context.Background())          => "", ""
//	Meta(WithMeta(WithMeta(bg, "a", "1"), "b", "2")) => "b", "2"
func WithMeta(ctx context.Context, tenant, trace string) context.Context {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Meta reports the tenant and trace IDs carried by ctx. Missing values come
// back as empty strings.
func Meta(ctx context.Context) (tenant, trace string) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
