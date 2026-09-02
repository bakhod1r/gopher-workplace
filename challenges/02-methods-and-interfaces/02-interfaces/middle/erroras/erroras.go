// Package erroras — Gopher Workplace challenge.
package erroras

// HTTPError carries a status code.
type HTTPError struct {
	Status int
}

// Error renders "http <status>".
func (e *HTTPError) Error() string {
	// TODO(candidate): "http <Status>".
	panic("not implemented")
}

// Call performs a request; any non-200 status is a wrapped *HTTPError.
//
// Examples:
//
//	Call(200) => nil
//	Call(500) => error "call: http 500"
func Call(status int) error {
	// TODO(candidate): wrap an *HTTPError for non-200.
	panic("not implemented")
}

// StatusOf digs the status code out of an error chain, or 0.
func StatusOf(err error) int {
	// TODO(candidate): errors.As into an *HTTPError.
	panic("not implemented")
}

// Retryable reports whether the error carries a 5xx status.
func Retryable(err error) bool {
	// TODO(candidate): 500 and above.
	panic("not implemented")
}
