package abortupload

import (
	"context"
	"errors"
	"testing"
)

func TestAbortUpload(t *testing.T) {
	cases := []struct {
		name  string
		check func(t *testing.T)
	}{
		{"returns_canceled", func(t *testing.T) {
			if got := AbortUpload(); got != context.Canceled {
				t.Errorf("AbortUpload() = %v, want %v", got, context.Canceled)
			}
		}},
		{"errors_is_canceled", func(t *testing.T) {
			if !errors.Is(AbortUpload(), context.Canceled) {
				t.Error("errors.Is(AbortUpload(), context.Canceled) = false, want true")
			}
		}},
		{"not_deadline", func(t *testing.T) {
			if errors.Is(AbortUpload(), context.DeadlineExceeded) {
				t.Error("AbortUpload() reported DeadlineExceeded, want Canceled")
			}
		}},
		{"not_nil", func(t *testing.T) {
			if AbortUpload() == nil {
				t.Error("AbortUpload() = nil, want context.Canceled")
			}
		}},
		{"repeatable", func(t *testing.T) {
			for i := 0; i < 5; i++ {
				if got := AbortUpload(); got != context.Canceled {
					t.Fatalf("call %d: AbortUpload() = %v, want %v", i, got, context.Canceled)
				}
			}
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, tc.check)
	}
}
