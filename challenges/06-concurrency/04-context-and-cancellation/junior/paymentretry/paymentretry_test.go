package paymentretry

import (
	"context"
	"errors"
	"testing"
	"time"
)

var errDeclined = errors.New("card declined")

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

// failThen returns a charge func that fails n times, then succeeds.
func failThen(n int, calls *int) func() error {
	return func() error {
		*calls++
		if *calls <= n {
			return errDeclined
		}
		return nil
	}
}

func TestChargeWithRetry(t *testing.T) {
	live, cancelLive := context.WithCancel(context.Background())
	defer cancelLive()

	cases := []struct {
		name      string
		ctx       context.Context
		attempts  int
		failFirst int
		want      error
		wantCalls int
	}{
		{"succeeds_first_try", live, 3, 0, nil, 1},
		{"succeeds_third_try", live, 3, 2, nil, 3},
		{"exhausts_attempts", live, 2, 5, errDeclined, 2},
		{"zero_attempts", live, 0, 0, nil, 0},
		{"client_hung_up", cancelled(), 3, 5, context.Canceled, 0},
		{"budget_expired", expired(), 3, 5, context.DeadlineExceeded, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			got := ChargeWithRetry(tc.ctx, tc.attempts, failThen(tc.failFirst, &calls))
			if got != tc.want {
				t.Fatalf("ChargeWithRetry() = %v, want %v", got, tc.want)
			}
			if calls != tc.wantCalls {
				t.Errorf("charge called %d times, want %d", calls, tc.wantCalls)
			}
		})
	}
}
