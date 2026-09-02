package legacyclient

import (
	"context"
	"testing"
)

func TestLegacyClientContext(t *testing.T) {
	cases := []struct {
		name  string
		check func(t *testing.T)
	}{
		{"is_todo", func(t *testing.T) {
			if LegacyClientContext() != context.TODO() {
				t.Errorf("LegacyClientContext() = %v, want context.TODO()", LegacyClientContext())
			}
		}},
		{"not_background", func(t *testing.T) {
			if LegacyClientContext() == context.Background() {
				t.Error("LegacyClientContext() = context.Background(), want context.TODO()")
			}
		}},
		{"no_error", func(t *testing.T) {
			if err := LegacyClientContext().Err(); err != nil {
				t.Errorf("Err() = %v, want nil", err)
			}
		}},
		{"done_is_nil", func(t *testing.T) {
			if LegacyClientContext().Done() != nil {
				t.Error("Done() != nil, want nil")
			}
		}},
		{"no_deadline", func(t *testing.T) {
			if _, ok := LegacyClientContext().Deadline(); ok {
				t.Error("Deadline() ok = true, want false")
			}
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, tc.check)
	}
}
