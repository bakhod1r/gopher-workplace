// Package customas — Gopher Workplace challenge.
package customas

// Modern is the structured error type callers migrate to.
type Modern struct {
	Op   string
	Code int
}

// Error implements the error interface.
func (e *Modern) Error() string {
	return e.Op + " failed"
}

// LegacyError is the old representation.
type LegacyError struct {
	Op  string
	Num int
}

// Error implements the error interface.
func (e *LegacyError) Error() string {
	return e.Op + " failed (legacy)"
}

// As converts the legacy error into a *Modern when asked for one.
func (e *LegacyError) As(target any) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
