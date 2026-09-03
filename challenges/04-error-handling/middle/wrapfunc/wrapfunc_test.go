package wrapfunc

import (
	"errors"
	"testing"
)

func TestNamed(t *testing.T) {
	t.Run("lazy", func(t *testing.T) {
		calls := 0
		Named("load", func() error { calls++; return nil })
		if calls != 0 {
			t.Errorf("calls = %d, want 0 before invocation", calls)
		}
	})

	t.Run("success_passthrough", func(t *testing.T) {
		if got := Named("load", func() error { return nil })(); got != nil {
			t.Errorf("got = %v, want nil", got)
		}
	})

	t.Run("annotates", func(t *testing.T) {
		got := Named("load", func() error { return ErrBoom })()
		if got == nil {
			t.Fatal("got nil, want an error")
		}
		if got.Error() != "load: boom" {
			t.Errorf("message = %q, want %q", got.Error(), "load: boom")
		}
	})

	t.Run("matchable", func(t *testing.T) {
		got := Named("load", func() error { return ErrBoom })()
		if !errors.Is(got, ErrBoom) {
			t.Error("errors.Is = false, want true")
		}
	})

	t.Run("runs_each_call", func(t *testing.T) {
		calls := 0
		step := Named("load", func() error { calls++; return nil })
		step()
		step()
		if calls != 2 {
			t.Errorf("calls = %d, want 2", calls)
		}
	})
}
