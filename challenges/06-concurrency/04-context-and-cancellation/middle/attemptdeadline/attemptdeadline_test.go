package attemptdeadline

import (
	"context"
	"errors"
	"testing"
	"time"
)

var errDeclined = errors.New("card declined")

// provider fails the first failFirst calls, records how many times it was
// called, and always reports its own context error first — a well-behaved
// upstream client.
func provider(failFirst int, calls *int, sawDeadline *bool) Charge {
	return func(ctx context.Context) error {
		*calls++
		if _, ok := ctx.Deadline(); ok {
			*sawDeadline = true
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		if *calls <= failFirst {
			return errDeclined
		}
		return nil
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

func TestChargeWithAttemptDeadline(t *testing.T) {
	live := context.Background()

	cases := []struct {
		name       string
		ctx        context.Context
		attempts   int
		perAttempt time.Duration
		failFirst  int
		wantErr    error
		wantCalls  int
	}{
		{"succeeds_first_try", live, 3, time.Hour, 0, nil, 1},
		{"succeeds_third_try", live, 3, time.Hour, 2, nil, 3},
		{"exhausts_attempts", live, 2, time.Hour, 5, errDeclined, 2},
		{"zero_attempts", live, 0, time.Hour, 5, nil, 0},
		{"attempt_deadline_expires", live, 3, 0, 0, context.DeadlineExceeded, 3},
		{"client_hung_up", hungUp(), 3, time.Hour, 0, context.Canceled, 0},
		{"request_budget_expired", budgetExpired(), 3, time.Hour, 0, context.DeadlineExceeded, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			sawDeadline := false

			err := ChargeWithAttemptDeadline(tc.ctx, tc.attempts, tc.perAttempt, provider(tc.failFirst, &calls, &sawDeadline))
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("ChargeWithAttemptDeadline() = %v, want %v", err, tc.wantErr)
			}
			if calls != tc.wantCalls {
				t.Errorf("charge called %d times, want %d", calls, tc.wantCalls)
			}
			if calls > 0 && !sawDeadline {
				t.Error("attempt context carried no deadline, want a per-attempt sub-deadline")
			}
		})
	}
}
