package shutdowntwice

import (
	"context"
	"errors"
	"testing"
)

func TestShutdownTwice(t *testing.T) {
	cases := []struct {
		name  string
		check func(t *testing.T)
	}{
		{"canceled", func(t *testing.T) {
			if got := ShutdownTwice(); got != context.Canceled {
				t.Errorf("ShutdownTwice() = %v, want %v", got, context.Canceled)
			}
		}},
		{"no_panic", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("ShutdownTwice() panicked: %v", r)
				}
			}()
			ShutdownTwice()
		}},
		{"errors_is", func(t *testing.T) {
			if !errors.Is(ShutdownTwice(), context.Canceled) {
				t.Error("errors.Is(..., context.Canceled) = false, want true")
			}
		}},
		{"not_nil", func(t *testing.T) {
			if ShutdownTwice() == nil {
				t.Error("got nil, want context.Canceled")
			}
		}},
		{"stable_across_calls", func(t *testing.T) {
			for i := 0; i < 20; i++ {
				if got := ShutdownTwice(); got != context.Canceled {
					t.Fatalf("call %d = %v, want %v", i, got, context.Canceled)
				}
			}
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, tc.check)
	}
}
