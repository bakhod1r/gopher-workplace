package publishevent

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

func fullChan() chan string {
	ch := make(chan string, 1)
	ch <- "backlog"
	return ch
}

func TestPublish(t *testing.T) {
	live, cancelLive := context.WithCancel(context.Background())
	defer cancelLive()

	cases := []struct {
		name    string
		ctx     context.Context
		out     chan string
		event   string
		wantErr error
		wantVal string
	}{
		{"buffered_room", live, make(chan string, 1), "order.created", nil, "order.created"},
		{"empty_event", live, make(chan string, 1), "", nil, ""},
		{"writer_backed_up_and_disconnected", cancelled(), fullChan(), "order.created", context.Canceled, ""},
		{"writer_backed_up_and_expired", expired(), fullChan(), "order.paid", context.DeadlineExceeded, ""},
		{"no_buffer_and_cancelled", cancelled(), make(chan string), "order.shipped", context.Canceled, ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := Publish(tc.ctx, tc.out, tc.event)
			if err != tc.wantErr {
				t.Fatalf("Publish() = %v, want %v", err, tc.wantErr)
			}
			if tc.wantErr != nil {
				return
			}
			got := <-tc.out
			if got != tc.wantVal {
				t.Errorf("published %q, want %q", got, tc.wantVal)
			}
		})
	}
}
