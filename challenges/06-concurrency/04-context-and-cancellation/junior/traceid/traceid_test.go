package traceid

import (
	"context"
	"testing"
)

func TestTraceID(t *testing.T) {
	bg := context.Background()

	cases := []struct {
		name   string
		ctx    context.Context
		want   string
		wantOK bool
	}{
		{"no_trace_header", bg, "", false},
		{"trace_present", WithTraceID(bg, "4bf92f"), "4bf92f", true},
		{"empty_header_is_present", WithTraceID(bg, ""), "", true},
		{"innermost_hop_wins", WithTraceID(WithTraceID(bg, "a"), "b"), "b", true},
		{"survives_cancellable_derivation", func() context.Context {
			ctx, cancel := context.WithCancel(WithTraceID(bg, "deadbeef"))
			cancel()
			return ctx
		}(), "deadbeef", true},
		{"foreign_string_key_does_not_collide", context.WithValue(bg, "traceKey", "spoofed"), "", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := TraceID(tc.ctx)
			if ok != tc.wantOK {
				t.Fatalf("TraceID() ok = %v, want %v", ok, tc.wantOK)
			}
			if got != tc.want {
				t.Errorf("TraceID() = %q, want %q", got, tc.want)
			}
		})
	}
}
