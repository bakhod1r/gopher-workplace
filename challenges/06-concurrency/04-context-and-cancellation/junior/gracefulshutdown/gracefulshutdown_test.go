package gracefulshutdown

import (
	"context"
	"testing"
	"time"
)

func sigterm() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func expiredWindow() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	cancel()
	return ctx
}

func missedDrainDeadline() context.Context {
	ctx, cancel := context.WithDeadline(context.Background(), time.Unix(0, 0))
	cancel()
	return ctx
}

func TestWaitForShutdown(t *testing.T) {
	cases := []struct {
		name string
		ctx  context.Context
		want error
	}{
		{"sigterm", sigterm(), context.Canceled},
		{"expired_drain_window", expiredWindow(), context.DeadlineExceeded},
		{"missed_drain_deadline", missedDrainDeadline(), context.DeadlineExceeded},
		{"child_of_sigterm", func() context.Context {
			ctx, cancel := context.WithCancel(sigterm())
			defer cancel()
			return ctx
		}(), context.Canceled},
		{"signal_handler_fired_twice", func() context.Context {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			cancel()
			return ctx
		}(), context.Canceled},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := WaitForShutdown(tc.ctx); got != tc.want {
				t.Errorf("WaitForShutdown() = %v, want %v", got, tc.want)
			}
		})
	}
}
