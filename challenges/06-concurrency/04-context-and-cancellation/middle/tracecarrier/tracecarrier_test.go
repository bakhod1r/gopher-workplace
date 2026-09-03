package tracecarrier

import (
	"context"
	"reflect"
	"testing"
)

func TestTraceID(t *testing.T) {
	cases := []struct {
		name string
		ctx  func() context.Context
		want string
	}{
		{"tagged", func() context.Context { return WithTrace(context.Background(), "abc") }, "abc"},
		{"untagged", context.Background, ""},
		{"empty_id_ignored", func() context.Context { return WithTrace(context.Background(), "") }, ""},
		{"overwritten", func() context.Context {
			return WithTrace(WithTrace(context.Background(), "old"), "new")
		}, "new"},
		{"inherited_by_child", func() context.Context {
			ctx, cancel := context.WithCancel(WithTrace(context.Background(), "abc"))
			cancel()
			return ctx
		}, "abc"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := TraceID(tc.ctx()); got != tc.want {
				t.Errorf("TraceID() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestWithTraceKeepsParentOnEmptyID(t *testing.T) {
	parent := context.Background()
	if got := WithTrace(parent, ""); got != parent {
		t.Error("WithTrace with an empty ID should return the parent unchanged")
	}
}

func TestChain(t *testing.T) {
	got := Chain(context.Background(), "abc")
	want := []string{"", "abc", "abc"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Chain() = %v, want %v", got, want)
	}
}

func TestStringKeyDoesNotLeakIn(t *testing.T) {
	//nolint:staticcheck // deliberately using a plain string key from "another package"
	ctx := context.WithValue(context.Background(), "traceKey", "sneaky")
	if got := TraceID(ctx); got != "" {
		t.Errorf("TraceID() = %q, want %q — a string key must not collide", got, "")
	}
}
