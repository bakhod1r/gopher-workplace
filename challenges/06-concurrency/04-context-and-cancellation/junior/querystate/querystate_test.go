package querystate

import (
	"context"
	"testing"
)

func TestQueryState(t *testing.T) {
	cases := []struct {
		name  string
		check func(t *testing.T, connected, disconnected error)
	}{
		{"connected_is_nil", func(t *testing.T, connected, disconnected error) {
			if connected != nil {
				t.Errorf("connected = %v, want nil", connected)
			}
		}},
		{"disconnected_is_canceled", func(t *testing.T, connected, disconnected error) {
			if disconnected != context.Canceled {
				t.Errorf("disconnected = %v, want %v", disconnected, context.Canceled)
			}
		}},
		{"they_differ", func(t *testing.T, connected, disconnected error) {
			if connected == disconnected {
				t.Error("connected == disconnected, want nil then context.Canceled")
			}
		}},
		{"not_deadline", func(t *testing.T, connected, disconnected error) {
			if disconnected == context.DeadlineExceeded {
				t.Error("disconnected = DeadlineExceeded, want Canceled")
			}
		}},
		{"stable_across_calls", func(t *testing.T, connected, disconnected error) {
			c2, d2 := QueryState()
			if c2 != connected || d2 != disconnected {
				t.Errorf("second call = (%v, %v), want (%v, %v)", c2, d2, connected, disconnected)
			}
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			connected, disconnected := QueryState()
			tc.check(t, connected, disconnected)
		})
	}
}
