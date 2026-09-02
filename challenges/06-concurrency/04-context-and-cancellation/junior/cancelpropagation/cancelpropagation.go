// Package cancelpropagation — Gopher Workplace challenge.
package cancelpropagation

// QueryErrAfterRequestCancel builds the per-request context, derives the
// database query context from it, cancels the request (the client hung up),
// and returns the error the query context reports.
//
// Examples:
//
//	QueryErrAfterRequestCancel()                              => context.Canceled
//	errors.Is(QueryErrAfterRequestCancel(), context.Canceled) => true
//	the result is never nil
func QueryErrAfterRequestCancel() error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
