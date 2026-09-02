// Package retryattempt — Gopher Workplace challenge.
package retryattempt

import "context"

// attemptKey is the unexported key holding the current retry attempt number.
type attemptKey struct{}

// WithAttempt returns a copy of ctx tagged with the retry attempt number, so
// the HTTP client can stamp it on the outbound request and the logger can show
// which try failed.
//
// Examples:
//
//	Attempt(WithAttempt(bg, 2))                => 2
//	Attempt(context.Background())              => 0
//	Attempt(WithAttempt(WithAttempt(bg, 1), 2)) => 2
func WithAttempt(ctx context.Context, n int) context.Context {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Attempt reports the retry attempt recorded on ctx, or 0 when none was set.
func Attempt(ctx context.Context) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
