package wraponce

import (
	"errors"
	"testing"
)

func TestOnce(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := Once("save", nil); got != nil {
			t.Errorf("Once = %v, want nil", got)
		}
	})

	t.Run("wraps", func(t *testing.T) {
		got := Once("save", ErrA)
		if got.Error() != "save: a" {
			t.Errorf("message = %q, want %q", got.Error(), "save: a")
		}
	})

	t.Run("idempotent", func(t *testing.T) {
		once := Once("save", ErrA)
		twice := Once("save", once)
		if twice.Error() != "save: a" {
			t.Errorf("message = %q, want %q", twice.Error(), "save: a")
		}
		if twice != once {
			t.Error("second call produced a new value, want the input unchanged")
		}
	})

	t.Run("different_op_still_wraps", func(t *testing.T) {
		got := Once("outer", Once("save", ErrA))
		if got.Error() != "outer: save: a" {
			t.Errorf("message = %q, want %q", got.Error(), "outer: save: a")
		}
	})

	t.Run("matchable", func(t *testing.T) {
		if !errors.Is(Once("save", Once("save", ErrA)), ErrA) {
			t.Error("errors.Is = false, want true")
		}
	})
}
