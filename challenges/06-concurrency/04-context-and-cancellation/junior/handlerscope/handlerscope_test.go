package handlerscope

import (
	"context"
	"errors"
	"testing"
)

var errDB = errors.New("query failed")

func TestServeRequest(t *testing.T) {
	cases := []struct {
		name    string
		handler func(ctx context.Context) error
		want    error
	}{
		{"ok", func(ctx context.Context) error { return nil }, nil},
		{"handler_error", func(ctx context.Context) error { return errDB }, errDB},
		{"ctx_alive_during_handler", func(ctx context.Context) error { return ctx.Err() }, nil},
		{"ctx_is_cancellable", func(ctx context.Context) error {
			if ctx.Done() == nil {
				return errors.New("ctx.Done() is nil, want a cancellable request context")
			}
			return nil
		}, nil},
		{"ctx_is_derived", func(ctx context.Context) error {
			if ctx == context.Background() {
				return errors.New("handler received context.Background(), want a derived per-request context")
			}
			return nil
		}, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ServeRequest(tc.handler); !errors.Is(got, tc.want) {
				t.Errorf("ServeRequest() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestServeRequestCancelsOnReturn(t *testing.T) {
	var captured context.Context
	if err := ServeRequest(func(ctx context.Context) error {
		captured = ctx
		return nil
	}); err != nil {
		t.Fatalf("ServeRequest() = %v, want nil", err)
	}

	<-captured.Done()
	if got := captured.Err(); got != context.Canceled {
		t.Errorf("after ServeRequest returned, ctx.Err() = %v, want %v", got, context.Canceled)
	}
}
