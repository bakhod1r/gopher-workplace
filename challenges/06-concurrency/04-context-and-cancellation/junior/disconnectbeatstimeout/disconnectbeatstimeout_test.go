package disconnectbeatstimeout

import (
	"context"
	"errors"
	"testing"
)

func TestDisconnectDuringTimeout(t *testing.T) {
	cases := []struct {
		name  string
		check func(t *testing.T)
	}{
		{"reports_canceled", func(t *testing.T) {
			if got := DisconnectDuringTimeout(); got != context.Canceled {
				t.Errorf("DisconnectDuringTimeout() = %v, want %v", got, context.Canceled)
			}
		}},
		{"not_deadline_exceeded", func(t *testing.T) {
			if errors.Is(DisconnectDuringTimeout(), context.DeadlineExceeded) {
				t.Error("got DeadlineExceeded, want Canceled — cancel ran long before the deadline")
			}
		}},
		{"errors_is_canceled", func(t *testing.T) {
			if !errors.Is(DisconnectDuringTimeout(), context.Canceled) {
				t.Error("errors.Is(..., context.Canceled) = false, want true")
			}
		}},
		{"not_nil", func(t *testing.T) {
			if DisconnectDuringTimeout() == nil {
				t.Error("got nil, want context.Canceled")
			}
		}},
		{"deterministic", func(t *testing.T) {
			for i := 0; i < 20; i++ {
				if got := DisconnectDuringTimeout(); got != context.Canceled {
					t.Fatalf("call %d = %v, want %v", i, got, context.Canceled)
				}
			}
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, tc.check)
	}
}
