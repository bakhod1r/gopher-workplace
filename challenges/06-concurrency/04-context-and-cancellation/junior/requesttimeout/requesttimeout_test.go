package requesttimeout

import (
	"context"
	"testing"
	"time"
)

func TestWithRequestTimeoutGenerous(t *testing.T) {
	cases := []struct {
		name string
		d    time.Duration
	}{
		{"one_hour", time.Hour},
		{"one_minute", time.Minute},
		{"ten_seconds", 10 * time.Second},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := WithRequestTimeout(context.Background(), tc.d)
			defer cancel()

			if err := ctx.Err(); err != nil {
				t.Errorf("Err() = %v, want nil for a generous budget", err)
			}
			if _, ok := ctx.Deadline(); !ok {
				t.Error("Deadline() ok = false, want true")
			}
			if ctx.Done() == nil {
				t.Error("Done() = nil, want a real channel")
			}

			cancel()
			<-ctx.Done()
			if err := ctx.Err(); err != context.Canceled {
				t.Errorf("after cancel, Err() = %v, want %v", err, context.Canceled)
			}
		})
	}
}

func TestWithRequestTimeoutExpired(t *testing.T) {
	cases := []struct {
		name string
		d    time.Duration
	}{
		{"zero_budget", 0},
		{"negative_budget", -time.Second},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := WithRequestTimeout(context.Background(), tc.d)
			defer cancel()

			<-ctx.Done()
			if err := ctx.Err(); err != context.DeadlineExceeded {
				t.Errorf("Err() = %v, want %v", err, context.DeadlineExceeded)
			}
		})
	}
}

func TestWithRequestTimeoutInheritsCancel(t *testing.T) {
	parent, cancelParent := context.WithCancel(context.Background())
	cancelParent()

	ctx, cancel := WithRequestTimeout(parent, time.Hour)
	defer cancel()

	<-ctx.Done()
	if err := ctx.Err(); err != context.Canceled {
		t.Errorf("Err() = %v, want %v — a dead parent must win over a generous timeout", err, context.Canceled)
	}
}
