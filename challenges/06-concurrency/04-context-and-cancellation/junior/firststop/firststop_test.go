package firststop

import (
	"context"
	"testing"
	"time"
)

func cancelled() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func expired() context.Context {
	ctx, cancel := context.WithDeadline(context.Background(), time.Unix(0, 0))
	cancel()
	return ctx
}

func TestFirstStop(t *testing.T) {
	liveReq, cancelReq := context.WithCancel(context.Background())
	defer cancelReq()

	liveShutdown, cancelShutdown := context.WithCancel(context.Background())
	defer cancelShutdown()

	cases := []struct {
		name     string
		req      context.Context
		shutdown context.Context
		want     error
	}{
		{"client_hung_up", cancelled(), liveShutdown, context.Canceled},
		{"request_budget_gone", expired(), liveShutdown, context.DeadlineExceeded},
		{"process_shutting_down", liveReq, cancelled(), context.Canceled},
		{"drain_window_expired", liveReq, expired(), context.DeadlineExceeded},
		{"request_beats_never_ending_shutdown", cancelled(), context.Background(), context.Canceled},
		{"shutdown_beats_never_ending_request", context.Background(), expired(), context.DeadlineExceeded},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FirstStop(tc.req, tc.shutdown); got != tc.want {
				t.Errorf("FirstStop() = %v, want %v", got, tc.want)
			}
		})
	}
}
