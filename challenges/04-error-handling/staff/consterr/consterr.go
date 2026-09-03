// Package consterr — Gopher Workplace challenge.
package consterr

// Error is a string that implements the error interface.
type Error string

// Sentinels that cannot be reassigned.
const (
	ErrClosed Error = "closed"
	ErrBusy   Error = "busy"
)

// Error implements the error interface.
func (e Error) Error() string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
