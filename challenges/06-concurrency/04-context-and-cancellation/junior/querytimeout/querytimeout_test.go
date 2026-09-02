package querytimeout

import (
	"context"
	"errors"
	"testing"
)

func TestExhaustedQueryBudget(t *testing.T) {
	cases := []struct {
		name  string
		check func(t *testing.T)
	}{
		{"deadline_exceeded", func(t *testing.T) {
			if got := ExhaustedQueryBudget(); got != context.DeadlineExceeded {
				t.Errorf("ExhaustedQueryBudget() = %v, want %v", got, context.DeadlineExceeded)
			}
		}},
		{"errors_is", func(t *testing.T) {
			if !errors.Is(ExhaustedQueryBudget(), context.DeadlineExceeded) {
				t.Error("errors.Is(..., context.DeadlineExceeded) = false, want true")
			}
		}},
		{"not_canceled", func(t *testing.T) {
			if errors.Is(ExhaustedQueryBudget(), context.Canceled) {
				t.Error("got Canceled, want DeadlineExceeded")
			}
		}},
		{"not_nil", func(t *testing.T) {
			if ExhaustedQueryBudget() == nil {
				t.Error("ExhaustedQueryBudget() = nil, want context.DeadlineExceeded")
			}
		}},
		{"repeatable", func(t *testing.T) {
			for i := 0; i < 20; i++ {
				if got := ExhaustedQueryBudget(); got != context.DeadlineExceeded {
					t.Fatalf("call %d = %v, want %v", i, got, context.DeadlineExceeded)
				}
			}
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, tc.check)
	}
}
