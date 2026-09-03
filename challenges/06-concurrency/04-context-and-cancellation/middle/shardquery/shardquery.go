// Package shardquery — Gopher Workplace challenge.
package shardquery

import (
	"context"

	_ "golang.org/x/sync/errgroup"
)

// ShardQuery runs the analytics query against a single shard.
type ShardQuery func(ctx context.Context) error

// QueryAllShards runs one query per shard concurrently and waits for all of
// them. Every shard receives the errgroup's context, so the first shard that
// fails cancels the others and the dashboard request stops paying for scans
// whose results are already useless.
//
// It returns nil when every shard succeeded, otherwise the first error the
// group recorded.
//
// Examples:
//
//	QueryAllShards(ctx, nil)                    => nil
//	QueryAllShards(ctx, [ok, ok])               => nil
//	QueryAllShards(ctx, [broken, parked, parked]) => the broken shard's error
func QueryAllShards(ctx context.Context, shards []ShardQuery) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
