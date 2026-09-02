package cancelpropagation

import (
	"context"
	"errors"
	"testing"
)

func TestQueryErrAfterRequestCancel(t *testing.T) {
	cases := []struct {
		name  string
		check func(t *testing.T)
	}{
		{"query_is_canceled", func(t *testing.T) {
			if got := QueryErrAfterRequestCancel(); got != context.Canceled {
				t.Errorf("QueryErrAfterRequestCancel() = %v, want %v", got, context.Canceled)
			}
		}},
		{"errors_is", func(t *testing.T) {
			if !errors.Is(QueryErrAfterRequestCancel(), context.Canceled) {
				t.Error("errors.Is(..., context.Canceled) = false, want true")
			}
		}},
		{"not_nil", func(t *testing.T) {
			if QueryErrAfterRequestCancel() == nil {
				t.Error("got nil, want context.Canceled — cancellation must reach the child")
			}
		}},
		{"not_deadline", func(t *testing.T) {
			if errors.Is(QueryErrAfterRequestCancel(), context.DeadlineExceeded) {
				t.Error("got DeadlineExceeded, want Canceled")
			}
		}},
		{"repeatable", func(t *testing.T) {
			for i := 0; i < 20; i++ {
				if got := QueryErrAfterRequestCancel(); got != context.Canceled {
					t.Fatalf("call %d = %v, want %v", i, got, context.Canceled)
				}
			}
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, tc.check)
	}
}
