package replicafanout

import (
	"context"
	"errors"
	"testing"
	"time"
)

var errReplicaDown = errors.New("replica down")

func okReplica(row string) Replica {
	return func(ctx context.Context) (string, error) { return row, nil }
}

func downReplica() Replica {
	return func(ctx context.Context) (string, error) { return "", errReplicaDown }
}

// blockedReplica models a replica that never answers: it parks until the
// fan-out context is cancelled and reports the reason it saw.
func blockedReplica(observed chan<- error) Replica {
	return func(ctx context.Context) (string, error) {
		<-ctx.Done()
		observed <- ctx.Err()
		return "", ctx.Err()
	}
}

func hungUp() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func budgetExpired() context.Context {
	ctx, cancel := context.WithDeadline(context.Background(), time.Unix(0, 0))
	cancel()
	return ctx
}

func TestReadFromReplicas(t *testing.T) {
	observed := make(chan error, 8)

	cases := []struct {
		name         string
		ctx          context.Context
		replicas     []Replica
		want         string
		wantErr      error
		wantObserved int
	}{
		{"empty_replica_set", context.Background(), nil, "", ErrNoReplicas, 0},
		{"single_replica_answers", context.Background(), []Replica{okReplica("row-1")}, "row-1", nil, 0},
		{"losers_are_cancelled", context.Background(), []Replica{blockedReplica(observed), okReplica("row-2"), blockedReplica(observed)}, "row-2", nil, 2},
		{"failure_does_not_stop_the_read", context.Background(), []Replica{downReplica(), okReplica("row-3")}, "row-3", nil, 0},
		{"every_replica_down", context.Background(), []Replica{downReplica(), downReplica()}, "", errReplicaDown, 0},
		{"client_hung_up", hungUp(), []Replica{okReplica("row-4")}, "", context.Canceled, 0},
		{"request_budget_expired", budgetExpired(), []Replica{okReplica("row-5")}, "", context.DeadlineExceeded, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ReadFromReplicas(tc.ctx, tc.replicas)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("ReadFromReplicas() error = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("ReadFromReplicas() row = %q, want %q", got, tc.want)
			}
			for i := 0; i < tc.wantObserved; i++ {
				if reason := <-observed; reason != context.Canceled {
					t.Errorf("losing replica %d stopped with %v, want %v", i, reason, context.Canceled)
				}
			}
		})
	}
}
