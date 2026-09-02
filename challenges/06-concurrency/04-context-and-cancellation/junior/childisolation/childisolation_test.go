package childisolation

import (
	"context"
	"testing"
)

func TestRequestErrAfterQueryCancel(t *testing.T) {
	cases := []struct {
		name  string
		check func(t *testing.T)
	}{
		{"request_still_alive", func(t *testing.T) {
			if got := RequestErrAfterQueryCancel(); got != nil {
				t.Errorf("RequestErrAfterQueryCancel() = %v, want nil", got)
			}
		}},
		{"not_canceled", func(t *testing.T) {
			if RequestErrAfterQueryCancel() == context.Canceled {
				t.Error("cancelling the query cancelled the whole request, want the request untouched")
			}
		}},
		{"not_deadline", func(t *testing.T) {
			if RequestErrAfterQueryCancel() == context.DeadlineExceeded {
				t.Error("got DeadlineExceeded, want nil")
			}
		}},
		{"repeatable", func(t *testing.T) {
			for i := 0; i < 20; i++ {
				if got := RequestErrAfterQueryCancel(); got != nil {
					t.Fatalf("call %d = %v, want nil", i, got)
				}
			}
		}},
		{"deterministic", func(t *testing.T) {
			a, b := RequestErrAfterQueryCancel(), RequestErrAfterQueryCancel()
			if a != b {
				t.Errorf("two calls disagreed: %v vs %v", a, b)
			}
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, tc.check)
	}
}
