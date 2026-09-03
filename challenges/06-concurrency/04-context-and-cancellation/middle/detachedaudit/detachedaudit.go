// Package detachedaudit — Gopher Workplace challenge.
package detachedaudit

import "context"

type auditKey struct{}

// WithActor tags the context with the acting user.
//
// Examples:
//
//	WithActor(ctx, "u1") => a child carrying "u1"
func WithActor(ctx context.Context, actor string) context.Context {
	return context.WithValue(ctx, auditKey{}, actor)
}

// Record writes an audit entry for the request. It must still write when the
// request context is already cancelled, using a context that keeps the
// request's values but not its cancellation.
//
// It returns what write reported: the actor it saw and the error of the
// context write was given.
//
// Examples:
//
//	Record(live ctx with "u1", write)      => "u1", nil
//	Record(cancelled ctx with "u1", write) => "u1", nil
//	Record(ctx without actor, write)       => "", nil
func Record(ctx context.Context, write func(ctx context.Context, actor string) error) (string, error) {
	// TODO(candidate): implement this using context.WithoutCancel.
	panic("not implemented")
}
