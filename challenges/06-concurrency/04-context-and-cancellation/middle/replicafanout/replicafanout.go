// Package replicafanout — Gopher Workplace challenge.
package replicafanout

import (
	"context"
	"errors"
)

// Replica reads the requested row from one database replica.
type Replica func(ctx context.Context) (string, error)

// ErrNoReplicas reports that the read was issued against an empty replica set.
var ErrNoReplicas = errors.New("no replicas configured")

// ReadFromReplicas asks every replica for the same row at once and returns the
// first answer that comes back. The replicas share one derived context, so as
// soon as a winner is found the losers are cancelled and stop burning replica
// CPU on a row nobody will read.
//
// If the request context is already finished it returns ctx.Err() without
// touching a replica. If every replica fails it returns the failure reported by
// replicas[0]; an empty replica set returns ErrNoReplicas.
//
// Examples:
//
//	ReadFromReplicas(ctx, nil)                       => "", ErrNoReplicas
//	ReadFromReplicas(ctx, [slow, fast])              => the fast replica's row
//	ReadFromReplicas(cancelled ctx, [fast])          => "", context.Canceled
func ReadFromReplicas(ctx context.Context, replicas []Replica) (string, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
