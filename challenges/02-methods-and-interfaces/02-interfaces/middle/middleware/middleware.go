// Package middleware — Gopher Workplace challenge.
package middleware

// Handler transforms a request string.
type Handler interface {
	Handle(s string) string
}

// HandlerFunc adapts a function to Handler.
type HandlerFunc func(string) string

// Handle calls the underlying function.
func (f HandlerFunc) Handle(s string) string {
	// TODO(candidate): call the receiver.
	panic("not implemented")
}

// Middleware wraps a handler in extra behaviour.
type Middleware func(Handler) Handler

// WithCount returns a Middleware that increments *n on every call.
func WithCount(n *int) Middleware {
	// TODO(candidate): wrap, count, delegate.
	panic("not implemented")
}

// WithPrefix returns a Middleware that prefixes the handler's result.
//
// Examples:
//
//	WithPrefix("p:")(base).Handle("x") => "p:x"   // base echoes its input
func WithPrefix(p string) Middleware {
	// TODO(candidate): wrap and prefix the result.
	panic("not implemented")
}

// Apply wraps h so that the first middleware listed runs outermost.
//
// Examples:
//
//	Apply(base, WithPrefix("a:"), WithPrefix("b:")).Handle("x") => "a:b:x"
func Apply(h Handler, ms ...Middleware) Handler {
	// TODO(candidate): wrap back to front.
	panic("not implemented")
}
