// Package resultchaingen — Gopher Workplace challenge.
package resultchaingen

// Result carries either a value of T or a failure reason.
// Its zero value is a failure with an empty reason.
type Result[T any] struct {
	value  T
	ok     bool
	reason string
}

// Ok returns a successful result.
func Ok[T any](v T) Result[T] {
	// TODO(candidate): build a successful result.
	panic("not implemented")
}

// Fail returns a failed result carrying reason.
func Fail[T any](reason string) Result[T] {
	// TODO(candidate): build a failed result.
	panic("not implemented")
}

// Then applies f when r succeeded, and passes the failure
// through otherwise.
func Then[T, U any](r Result[T], f func(T) Result[U]) Result[U] {
	// TODO(candidate): apply f only on success, propagating failures.
	panic("not implemented")
}

// Unwrap returns the value and whether the result succeeded.
func (r Result[T]) Unwrap() (T, bool) {
	// TODO(candidate): report the value and the flag.
	panic("not implemented")
}
