// Package funcadapt — Gopher Workplace challenge.
package funcadapt

// Handler transforms a string.
type Handler interface {
	Handle(s string) string
}

// HandlerFunc adapts a plain function to Handler.
type HandlerFunc func(string) string

// Handle calls the underlying function.
//
// Examples:
//
//	HandlerFunc(strings.ToUpper).Handle("hi") => "HI"
func (f HandlerFunc) Handle(s string) string {
	// TODO(candidate): call the receiver.
	panic("not implemented")
}

// Run invokes h on s.
func Run(h Handler, s string) string {
	// TODO(candidate): delegate.
	panic("not implemented")
}

// Chain returns a Handler that applies every handler left to right.
//
// Examples:
//
//	Chain(upper, exclaim).Handle("hi") => "HI!"
//	Chain().Handle("hi")               => "hi"
func Chain(hs ...Handler) Handler {
	// TODO(candidate): return a HandlerFunc that folds over hs.
	panic("not implemented")
}
