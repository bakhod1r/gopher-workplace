package requestid

import (
	"context"
	"testing"
)

func TestRequestID(t *testing.T) {
	bg := context.Background()

	cases := []struct {
		name string
		ctx  context.Context
		want string
	}{
		{"no_id_attached", bg, "unknown"},
		{"id_from_proxy", WithRequestID(bg, "req-8f21"), "req-8f21"},
		{"empty_id_falls_back", WithRequestID(bg, ""), "unknown"},
		{"inner_middleware_wins", WithRequestID(WithRequestID(bg, "old"), "new"), "new"},
		{"wrong_type_falls_back", context.WithValue(bg, requestIDKey{}, 42), "unknown"},
		{"survives_cancellable_derivation", func() context.Context {
			ctx, cancel := context.WithCancel(WithRequestID(bg, "req-2"))
			defer cancel()
			return ctx
		}(), "req-2"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := RequestID(tc.ctx); got != tc.want {
				t.Errorf("RequestID() = %q, want %q", got, tc.want)
			}
		})
	}
}
