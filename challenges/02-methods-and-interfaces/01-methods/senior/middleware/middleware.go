// Package middleware — Gopher Workplace challenge.
package middleware

// Handler represents an HTTP-like handler.
type Handler func(req string) string

// Middleware modifies or intercepts requests/responses.
type Middleware func(Handler) Handler

// Stack is an ordered list of middlewares, outermost first.
type Stack []Middleware

// Then wraps next with every middleware in the stack. The first middleware in
// the stack ends up outermost, so it sees the request first.
func (s Stack) Then(next Handler) Handler {
	// TODO(candidate): apply the middlewares in reverse order so s[0] is outermost.
	panic("not implemented")
}
