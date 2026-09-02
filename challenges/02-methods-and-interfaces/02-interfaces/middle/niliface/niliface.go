// Package niliface — Gopher Workplace challenge.
package niliface

// OpError reports a failed operation.
type OpError struct {
	Op string
}

// Error renders the failure.
func (e *OpError) Error() string { return e.Op + " failed" }

// Run performs an operation. It returns a nil error on success.
//
// Examples:
//
//	Run(false) => nil
//	Run(true)  => &OpError{Op: "op"}
func Run(fail bool) error {
	// TODO(candidate): return a genuinely nil error when fail is false.
	panic("not implemented")
}

// IsNil reports whether err is the nil interface value.
func IsNil(err error) bool {
	// TODO(candidate): compare against nil.
	panic("not implemented")
}

// FailedCount counts the non-nil errors.
func FailedCount(errs []error) int {
	// TODO(candidate): count the failures.
	panic("not implemented")
}
