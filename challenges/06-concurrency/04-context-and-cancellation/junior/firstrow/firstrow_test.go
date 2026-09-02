package firstrow

import (
	"context"
	"testing"
	"time"
)

func rowsWith(vals ...string) <-chan string {
	ch := make(chan string, len(vals))
	for _, v := range vals {
		ch <- v
	}
	return ch
}

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

func TestFirstRow(t *testing.T) {
	live, cancelLive := context.WithCancel(context.Background())
	defer cancelLive()

	cases := []struct {
		name    string
		ctx     context.Context
		rows    <-chan string
		want    string
		wantErr error
	}{
		{"first_of_many", live, rowsWith("alice", "bob"), "alice", nil},
		{"single_row", live, rowsWith("carol"), "carol", nil},
		{"empty_row_value", live, rowsWith(""), "", nil},
		{"client_disconnected", cancelled(), make(chan string), "", context.Canceled},
		{"budget_expired", expired(), make(chan string), "", context.DeadlineExceeded},
		{"disconnected_with_no_rows", cancelled(), make(chan string), "", context.Canceled},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FirstRow(tc.ctx, tc.rows)
			if err != tc.wantErr {
				t.Fatalf("FirstRow() err = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("FirstRow() = %q, want %q", got, tc.want)
			}
		})
	}
}
