// Package nilinterface — Gopher Workplace challenge.
package nilinterface

// OpError reports a failed operation.
type OpError struct {
	Op string
}

// Error implements the error interface.
func (e *OpError) Error() string {
	return e.Op + " failed"
}

// Wrap converts e into an error, mapping a nil pointer to a nil error.
//
// Examples:
//
//	Wrap(nil)                     => nil
//	Wrap(&OpError{Op: "read"})    => "read failed"
func Wrap(e *OpError) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
