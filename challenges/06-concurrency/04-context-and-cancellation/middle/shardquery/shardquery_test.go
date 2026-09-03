package shardquery

import (
	"context"
	"errors"
	"testing"
)

var (
	errShardBroken = errors.New("shard 7 unavailable")
	errShardLate   = errors.New("shard 9 gave up")
)

func okShard() ShardQuery {
	return func(ctx context.Context) error { return nil }
}

func brokenShard() ShardQuery {
	return func(ctx context.Context) error { return errShardBroken }
}

// parkedShard blocks until the group context is cancelled, then reports the
// reason it observed.
func parkedShard(observed chan<- error) ShardQuery {
	return func(ctx context.Context) error {
		<-ctx.Done()
		observed <- ctx.Err()
		return ctx.Err()
	}
}

// lateShard also waits for cancellation but reports its own error, which the
// group must ignore because it arrives after the first one.
func lateShard() ShardQuery {
	return func(ctx context.Context) error {
		<-ctx.Done()
		return errShardLate
	}
}

func hungUp() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func budgetExpired() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	cancel()
	return ctx
}

func TestQueryAllShards(t *testing.T) {
	observed := make(chan error, 8)

	cases := []struct {
		name         string
		ctx          context.Context
		shards       []ShardQuery
		wantErr      error
		wantObserved int
	}{
		{"no_shards", context.Background(), nil, nil, 0},
		{"every_shard_succeeds", context.Background(), []ShardQuery{okShard(), okShard(), okShard()}, nil, 0},
		{"broken_shard_cancels_the_rest", context.Background(), []ShardQuery{parkedShard(observed), brokenShard(), parkedShard(observed)}, errShardBroken, 2},
		{"later_error_is_ignored", context.Background(), []ShardQuery{brokenShard(), lateShard()}, errShardBroken, 0},
		{"client_hung_up", hungUp(), []ShardQuery{parkedShard(observed)}, context.Canceled, 1},
		{"dashboard_budget_expired", budgetExpired(), []ShardQuery{parkedShard(observed)}, context.DeadlineExceeded, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := QueryAllShards(tc.ctx, tc.shards)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("QueryAllShards() = %v, want %v", err, tc.wantErr)
			}
			for i := 0; i < tc.wantObserved; i++ {
				reason := <-observed
				if reason != context.Canceled && reason != context.DeadlineExceeded {
					t.Errorf("parked shard %d stopped with %v, want a context error", i, reason)
				}
			}
		})
	}
}
