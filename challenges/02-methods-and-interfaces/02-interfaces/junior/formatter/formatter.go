// Package formatter — Gopher Workplace challenge.
package formatter

// Formatter renders a log message.
type Formatter interface {
	Format(msg string) string
}

// Plain renders the message as-is.
type Plain struct{}

// Format returns the message unchanged.
func (p Plain) Format(msg string) string {
	// TODO(candidate): return msg.
	panic("not implemented")
}

// KeyValue renders the message as key=value.
type KeyValue struct{}

// Format returns "msg=<msg>".
//
// Examples:
//
//	KeyValue{}.Format("hi") => "msg=hi"
func (k KeyValue) Format(msg string) string {
	// TODO(candidate): prefix with "msg=".
	panic("not implemented")
}

// Render formats msg using f.
func Render(f Formatter, msg string) string {
	// TODO(candidate): delegate to the formatter.
	panic("not implemented")
}
