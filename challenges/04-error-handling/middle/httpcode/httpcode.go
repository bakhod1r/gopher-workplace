// Package httpcode — Gopher Workplace challenge.
package httpcode

import (
	"errors"
	"fmt"
)

// HTTPError carries a transport status code.
type HTTPError struct {
	Code int
}

// Error implements the error interface.
func (e *HTTPError) Error() string {
	return fmt.Sprintf("http status %d", e.Code)
}

// CodeOf returns the status code carried by err, if any.
//
// Examples:
//
//	CodeOf(&HTTPError{Code: 404}) => 404, true
//	CodeOf(errors.New("boom"))    => 0, false
func CodeOf(err error) (int, bool) {
	// TODO(candidate): implement this.
	_ = errors.As
	panic("not implemented")
}
