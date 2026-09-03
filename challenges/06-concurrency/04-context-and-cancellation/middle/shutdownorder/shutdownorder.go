// Package shutdownorder — Gopher Workplace challenge.
package shutdownorder

import (
	"context"
)

// Service is one subsystem of the API server, in startup order.
type Service struct {
	Name string
	Stop func(ctx context.Context) error
}

// ShutdownServices stops services in reverse startup order — the HTTP listener
// before the cache, the cache before the database — so that nothing is torn
// down while something above it still depends on it. Every Stop shares the
// drain context, and the sequence aborts as soon as the drain window closes or
// a Stop reports a failure.
//
// It returns the names of the services that stopped cleanly, in the order they
// were stopped, along with the reason the sequence ended (nil if it completed).
//
// Examples:
//
//	ShutdownServices(ctx, [db, cache, http])     => ["http" "cache" "db"], nil
//	ShutdownServices(ctx, [db, brokenCache, http]) => ["http"], errStopFailed
//	ShutdownServices(cancelled ctx, [db, http])  => [], context.Canceled
func ShutdownServices(ctx context.Context, services []Service) ([]string, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
