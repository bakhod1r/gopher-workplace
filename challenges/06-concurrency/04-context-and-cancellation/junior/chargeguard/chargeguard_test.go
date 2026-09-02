package chargeguard

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

func TestCharge(t *testing.T) {
	live, cancelLive := context.WithCancel(context.Background())
	defer cancelLive()

	cases := []struct {
		name      string
		ctx       context.Context
		result    string
		want      string
		wantErr   error
		wantCalls int
	}{
		{"captures_when_alive", live, "captured", "captured", nil, 1},
		{"captures_other_receipt", live, "ch_1a2b", "ch_1a2b", nil, 1},
		{"empty_receipt_still_ok", live, "", "", nil, 1},
		{"client_gone_no_charge", cancelled(), "captured", "", context.Canceled, 0},
		{"budget_gone_no_charge", expired(), "captured", "", context.DeadlineExceeded, 0},
		{"background_is_alive", context.Background(), "captured", "captured", nil, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			capture := func() string {
				calls++
				return tc.result
			}

			got, err := Charge(tc.ctx, capture)
			if err != tc.wantErr {
				t.Fatalf("Charge() err = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("Charge() = %q, want %q", got, tc.want)
			}
			if calls != tc.wantCalls {
				t.Errorf("capture called %d times, want %d", calls, tc.wantCalls)
			}
		})
	}
}
