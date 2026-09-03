package wrapnonnil

import (
	"errors"
	"testing"
)

func TestWrapNonNil(t *testing.T) {
	t.Run("nil_stays_nil", func(t *testing.T) {
		if got := WrapNonNil("step", nil); got != nil {
			t.Errorf("WrapNonNil = %v, want nil", got)
		}
	})

	t.Run("empty_msg_nil", func(t *testing.T) {
		if got := WrapNonNil("", nil); got != nil {
			t.Errorf("WrapNonNil = %v, want nil", got)
		}
	})

	t.Run("message", func(t *testing.T) {
		got := WrapNonNil("step", ErrX)
		if got.Error() != "step: boom" {
			t.Errorf("message = %q, want %q", got.Error(), "step: boom")
		}
	})

	t.Run("empty_msg", func(t *testing.T) {
		got := WrapNonNil("", ErrX)
		if got.Error() != ": boom" {
			t.Errorf("message = %q, want %q", got.Error(), ": boom")
		}
	})

	t.Run("matchable", func(t *testing.T) {
		if !errors.Is(WrapNonNil("step", ErrX), ErrX) {
			t.Error("errors.Is = false, want true")
		}
	})
}
