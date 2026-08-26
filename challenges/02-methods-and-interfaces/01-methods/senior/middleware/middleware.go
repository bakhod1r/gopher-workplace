// Package middleware — Gopher Workplace challenge.
package middleware

// Handler represents an HTTP-like handler.
type Handler func(req string) string

// Middleware modifies or intercepts requests/responses.
type Middleware func(Handler) Handler

// Chain combines multiple middlewares into a single Middleware.
// The first middleware in the slice is applied first (outermost).
func Chain(mws ...Middleware) Middleware {
	return func(next Handler) Handler {
		// TODO(candidate): apply mws in reverse order so mws[0] is outermost.
		panic("not implemented")
	}
}
