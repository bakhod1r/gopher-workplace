package fetchinvoice

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

func TestFetchInvoice(t *testing.T) {
	live, cancelLive := context.WithCancel(context.Background())
	defer cancelLive()

	cases := []struct {
		name    string
		ctx     context.Context
		id      int
		want    string
		wantErr error
	}{
		{"found", live, 7, "invoice-7", nil},
		{"found_large_id", live, 100045, "invoice-100045", nil},
		{"zero_id", live, 0, "", ErrNotFound},
		{"negative_id", live, -3, "", ErrNotFound},
		{"client_disconnected", cancelled(), 7, "", context.Canceled},
		{"budget_expired", expired(), 7, "", context.DeadlineExceeded},
		{"context_beats_not_found", cancelled(), 0, "", context.Canceled},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FetchInvoice(tc.ctx, tc.id)
			if err != tc.wantErr {
				t.Fatalf("FetchInvoice() err = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("FetchInvoice() = %q, want %q", got, tc.want)
			}
		})
	}
}
