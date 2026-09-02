// Package resultgen — Gopher Workplace challenge.
package resultgen

// Result carries either a value of T or a failure reason.
// Its zero value is a failure with an empty reason.
type Result[T any] struct {
	value  T
	ok     bool
	reason string
}

// Ok returns a successful result holding v.
func Ok[T any](v T) Result[T] {
	// TODO(candidate): build a successful result.
	panic("not implemented")
}

// Fail returns a failed result carrying reason.
func Fail[T any](reason string) Result[T] {
	// TODO(candidate): build a failed result.
	panic("not implemented")
}

// Unwrap returns the value and whether the result was successful.
func (r Result[T]) Unwrap() (T, bool) {
	// TODO(candidate): report the value and success flag.
	panic("not implemented")
}

// Reason returns the failure reason, or an empty string on success.
func (r Result[T]) Reason() string {
	// TODO(candidate): report the failure reason.
	panic("not implemented")
}
