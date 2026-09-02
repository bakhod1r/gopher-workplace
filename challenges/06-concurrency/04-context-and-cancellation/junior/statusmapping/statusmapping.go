// Package statusmapping — Gopher Workplace challenge.
package statusmapping

// Status maps the error a handler finished with onto the label the access log
// and the metrics counter use.
//
//	nil                       => "ok"
//	context.Canceled          => "client_closed_request"
//	context.DeadlineExceeded  => "gateway_timeout"
//	anything else             => "internal_error"
//
// Wrapped errors count: an error wrapping context.DeadlineExceeded with %w
// still maps to "gateway_timeout".
//
// Examples:
//
//	Status(nil)                                        => "ok"
//	Status(context.Canceled)                           => "client_closed_request"
//	Status(fmt.Errorf("query: %w", context.DeadlineExceeded)) => "gateway_timeout"
func Status(err error) string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
