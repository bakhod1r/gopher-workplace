package drainnotifications

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func closedChan(vals ...string) <-chan string {
	ch := make(chan string, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
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

func TestDrain(t *testing.T) {
	live, cancelLive := context.WithCancel(context.Background())
	defer cancelLive()

	cases := []struct {
		name    string
		ctx     context.Context
		ch      <-chan string
		want    []string
		wantErr error
	}{
		{"two_notifications", live, closedChan("a", "b"), []string{"a", "b"}, nil},
		{"preserves_order", live, closedChan("first", "second", "third"), []string{"first", "second", "third"}, nil},
		{"single", live, closedChan("only"), []string{"only"}, nil},
		{"none_pending", live, closedChan(), []string{}, nil},
		{"subscriber_cancelled", cancelled(), make(chan string), []string{}, context.Canceled},
		{"subscriber_expired", expired(), make(chan string), []string{}, context.DeadlineExceeded},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Drain(tc.ctx, tc.ch)
			if err != tc.wantErr {
				t.Fatalf("Drain() err = %v, want %v", err, tc.wantErr)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Drain() = %v, want %v", got, tc.want)
			}
		})
	}
}
