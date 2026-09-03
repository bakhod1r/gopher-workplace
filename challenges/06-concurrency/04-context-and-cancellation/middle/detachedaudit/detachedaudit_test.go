package detachedaudit

import (
	"context"
	"errors"
	"testing"
)

func capture(seen *string, ctxErr *error) func(context.Context, string) error {
	return func(ctx context.Context, actor string) error {
		*seen = actor
		*ctxErr = ctx.Err()
		return nil
	}
}

func TestRecord(t *testing.T) {
	cases := []struct {
		name      string
		ctx       func() context.Context
		wantActor string
	}{
		{"live_request", func() context.Context {
			return WithActor(context.Background(), "u1")
		}, "u1"},
		{"cancelled_request", func() context.Context {
			ctx, cancel := context.WithCancel(WithActor(context.Background(), "u2"))
			cancel()
			return ctx
		}, "u2"},
		{"expired_deadline", func() context.Context {
			ctx, cancel := context.WithTimeout(WithActor(context.Background(), "u3"), 0)
			_ = cancel
			return ctx
		}, "u3"},
		{"no_actor", context.Background, ""},
		{"nested_values", func() context.Context {
			return WithActor(WithActor(context.Background(), "old"), "u4")
		}, "u4"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var seen string
			var ctxErr error
			actor, err := Record(tc.ctx(), capture(&seen, &ctxErr))
			if err != nil {
				t.Fatalf("Record() error = %v, want nil", err)
			}
			if actor != tc.wantActor || seen != tc.wantActor {
				t.Errorf("actor = %q (write saw %q), want %q", actor, seen, tc.wantActor)
			}
			if ctxErr != nil {
				t.Errorf("write received a finished context: %v", ctxErr)
			}
		})
	}
}

func TestWriteErrorIsReturned(t *testing.T) {
	errWrite := errors.New("audit sink down")
	ctx, cancel := context.WithCancel(WithActor(context.Background(), "u1"))
	cancel()

	actor, err := Record(ctx, func(context.Context, string) error { return errWrite })
	if !errors.Is(err, errWrite) {
		t.Errorf("err = %v, want %v", err, errWrite)
	}
	if actor != "u1" {
		t.Errorf("actor = %q, want %q", actor, "u1")
	}
}
